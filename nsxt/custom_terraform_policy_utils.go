/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package nsxt

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
)

/* Function to create and update NSXT resources. If the resource does not exist it will try to
create it. In case it is present then it updates the resource. */
func APICreateOrUpdate(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema,
	opts ...bool) error {
	log.Printf("[DEBUG] start of APICreateOrUpdate")
	var robj interface{}
	obj := d
	nsxtClient := meta.(*nsxtclient.NsxtClient)

	if data, err := SchemaToNsxtData(obj, s); err == nil {
		resourceID := d.Id()
		// calculate the policy path of resource
		path := ComputePolicyPath(d, objType, false, nsxtClient, false)
		pathWithoutQueryParams := strings.Split(path, "?")[0]
		// trim policy/api/v1 from pathWithoutQueryParams and save as path, this way path can also be used to refer dependant resource policypath
		pathWithoutQueryParamsTrimmed := strings.TrimPrefix(pathWithoutQueryParams, nsxtClient.Config.BasePath)

		if resourceID != "" {
			// resource with id already present in NSX, this is an update request
			log.Printf("[INFO] APICreateOrUpdate: Updating obj %v with path %s\n", objType, path)
			err = nsxtClient.NsxtSession.Patch(path, data, &robj)
			if err != nil {
				log.Printf("[ERROR] APICreateOrUpdate updation failed %v\n", err)
			} else {
				d.SetId(pathWithoutQueryParamsTrimmed)
				d.Set("path", pathWithoutQueryParamsTrimmed)
			}
		} else {
			// resource is new, this is a create request
			log.Printf("[INFO] APICreateOrUpdate: Creating obj %v schema %v data %v with path %s\n", objType, d, data, path)
			if objType == "SecurityPolicyRule" || objType == "GatewayPolicyRule" {
				/* For Rule- source_groups, destination_groups, service, action and scope are optional and defaulted in create, but not defaulted in update. These are
				kept as optional in yaml without mentioning defaults. Hence, terraform resource for Rule has optional against them.
				Here adding the manual check to match API behaviour. Check each property value given or not. If not given, default it. */
				dataMap := data.(map[string]interface{})
				if dataMap["action"] == nil {
					dataMap["action"] = "ALLOW"
				}
				keys_with_any_default_value := []string{"source_groups", "destination_groups", "services", "scope"}
				for _, key := range keys_with_any_default_value {
					if dataMap[key] == nil {
						dataMap[key] = []interface{}{"ANY"}
					}
				}
			}
			err = nsxtClient.NsxtSession.Patch(path, data, &robj)
			if err != nil {
				log.Printf("[ERROR] APICreateOrUpdate creation failed %v\n", err)
			} else {
				d.SetId(pathWithoutQueryParamsTrimmed)
				d.Set("path", pathWithoutQueryParamsTrimmed)
			}
		}
		return err
	} else {
		log.Printf("[ERROR] APICreateOrUpdate: Error %v", err)
		return err
	}
}

/* Funtion for reading the current state of the resource in NSX.
Called during multiple phases in lifecycle-
1. Initial Resource Creation: When Terraform creates a new resource for the first time, the Read function is called to fetch the initial state of the resource. This state is then stored in the Terraform state file.
2. Refresh Operation: During a refresh operation, Terraform calls the Read function to retrieve the current state of the resource from NSX. This allows Terraform to compare the current state with the desired state defined in the Terraform configuration and determine if any updates are required.
(The refresh operation in Terraform is triggered explicitly by running the terraform refresh command or as part of other Terraform operations like terraform apply or terraform plan).
3. Post Update operation: To fetch latest changes and update state file accordingly.
4. Post Delete operation: To fetch the final state of the resource, which will typically be null state to indicate that the resource no longer exists, the update the state file with the final state of the resource, ensuring that the state accurately reflects the state of the infrastructure.
*/
func APIRead(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	var obj interface{}
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	// calculate the policy path of resource
	id := d.Id()
	policyPath := nsxtClient.Config.BasePath + id
	resourceName := d.Get("display_name").(string)

	if id != "" {
		// Read using ID as complete policy path
		log.Printf("[DEBUG] APIRead reading object with id %v\n", id)
		err := nsxtClient.NsxtSession.Get(policyPath, &obj)
		log.Printf("[DEBUG] json unmarshal response: %v\n", obj)
		if err == nil && obj != nil {
			d.SetId(id)
			// set readonly property _revision for terraform show
			objMap, ok := obj.(map[string]interface{})
			if !ok {
				return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
			}
		} else {
			log.Printf("[INFO] Read not successful for " + objType)
			// Set resource ID to "" to signal Terraform that the resource should be treated as non-existent or in an error state.
			d.SetId("")
			return err
		}
	} else if resourceName != "" {
		// compute policy path, and read using display_name
		policyPath := ComputePolicyPath(d, objType, true, nsxtClient, false)
		log.Printf("[DEBUG] APIRead reading object with display_name %v\n", resourceName)
		err := nsxtClient.NsxtSession.Get(policyPath, &obj)
		if err == nil && obj != nil {
			d.SetId(policyPath)
			// set readonly property _revision for terraform show
			objMap, ok := obj.(map[string]interface{})
			if !ok {
				return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
			}
		} else {
			log.Printf("[INFO] Read not successful for " + objType)
			d.SetId("")
			return err
		}
	} else {
		log.Printf("[DEBUG] APIRead object ID not present for resource %s\n", objType)
		d.SetId("")
	}

	if localData, err := SchemaToNsxtData(d, s); err == nil {
		modAPIRes, err := SetDefaultsInAPIRes(obj, localData, s)
		if err != nil {
			log.Printf("[ERROR] APIRead in modifying api response object %v\n", err)
			return err
		}
		if _, err := APIDataToSchema(modAPIRes, d, s); err != nil {
			log.Printf("[ERROR] APIRead in setting read object %v\n", err)
			d.SetId("")
			return err
		}
	} else {
		return err
	}

	return nil
}

func setAttrsInDatasourceSchema(mapObject interface{}, d *schema.ResourceData, objType string) {
	for key, value := range mapObject.(map[string]interface{}) {
		// Filter based on nsx_id or display_name to get only what datasource desires
		if key == "id" {
			d.Set("nsx_id", value.(string))
		}
		if key == "path" {
			d.SetId(value.(string)) // because terraform ID is policyPath value
			d.Set("path", value.(string))
		}
		if key == "display_name" {
			d.Set("display_name", value.(string))
		}
		if key == "description" {
			d.Set("description", value.(string))
		}
		if key == "parent_path" && d.Get("parent_path") != nil {
			d.Set("parent_path", value.(string))
		}
		if (objType == "VpcIpAddressAllocation" || objType == "IpAddressAllocation") && key == "allocation_ip" {
			d.Set("allocation_ip", value.(string))
		}
	}
}

// Function to do READ on datasource, using id, display_name, parent_path
func DatasourceRead(d *schema.ResourceData, meta interface{}, objType string, s *schema.Resource) error {
	var obj interface{}
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	var contextInfoMap map[string]interface{}
	var scope string
	var domain string
	contextInfo := d.Get("context").([]interface{})
	if len(contextInfo) > 0 {
		contextInfoMap = contextInfo[0].(map[string]interface{})
	}
	scope = contextInfoMap["scope"].(string)
	domain = contextInfoMap["domain"].(string)

	displayName := d.Get("display_name").(string)
	nsxID := d.Get("nsx_id").(string)
	var uri string
	// Some objTypes (should be resource_type value) are differently worded in .yaml, hence correcting them for usage in Search query
	switch objType {
	case "SecurityPolicyRule":
		objType = "Rule"
	case "GatewayPolicyRule":
		objType = "Rule"
	}

	// Get the object from NSX using hidden full text search API
	uri = nsxtClient.Config.BasePath + "/search?query=resource_type:" + objType
	if nsxID != "" {
		encodedNsxId := url.QueryEscape(nsxID)
		uri += "%20AND%20id:" + encodedNsxId
	}
	if displayName != "" {
		encodedDisplayName := url.QueryEscape(displayName)
		uri += "%20AND%20display_name:" + encodedDisplayName
	}
	if d.Get("parent_path") != nil {
		parentPath := d.Get("parent_path").(string)
		if parentPath != "" {
			uri += "%20AND%20parent_path:%22" + parentPath + "%22"
		}
	}
	if scope == "project" {
		// If it is project/infra object, search in context of project
		uri += "&context=projects:" + "/orgs/" + nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID
	} else if scope == "vpc" {
		// Search in context of Vpc
		uri += "&context=vpcs:" + "/orgs/" + nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
	} else {
		// No context for infra based Search, include domain info
		if objType == "Group" {
			domainPolicyPath := "/infra/domains/" + domain
			uri += "%20AND%20parent_path:%22" + domainPolicyPath + "%22"
		}
	}
	err := nsxtClient.NsxtSession.Get(uri, &obj)
	if err != nil {
		return fmt.Errorf("[ERROR] Search failed for object %s %v", objType, err)
	} else {
		objMap, ok := obj.(map[string]interface{})
		if !ok {
			return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
		}
		// get results of type []interface {} to get id and policypath
		results := objMap["results"]
		if objMap["result_count"].(float64) == 1 {
			setAttrsInDatasourceSchema(results.([]interface{})[0], d, objType)
		} else if objMap["result_count"].(float64) == 0 {
			// get the entity using listing api, for some VPC resources(eg VpcIpAddressAllocation), search API isn't indexed
			uri = ComputePolicyPath(d, objType, false, nsxtClient, true)
			// keep maximum allowed page size for pagination
			uri += "?page_size=1000"
			err := nsxtClient.NsxtSession.Get(uri, &obj)
			if err != nil {
				return fmt.Errorf("[ERROR] GET failed for object %s %v", objType, err)
			} else {
				objMap, ok := obj.(map[string]interface{})
				if !ok {
					return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
				}
				// get results of type []interface {}
				results := objMap["results"]
				perfectMatchFound := false
				cursor := 0
				if objMap["cursor"] != nil {
					cursor = objMap["cursor"].(int)
				}
				for {
					for _, itemObject := range results.([]interface{}) {
						for key, value := range itemObject.(map[string]interface{}) {
							// Filter based on nsx_id or display_name to get only what datasource desires
							//TODO: Check how to refine search in case of multiple entries found with same display_name
							if (nsxID != "" && key == "id" && value == nsxID) || (displayName != "" && key == "display_name" && value == displayName) {
								perfectMatchFound = true
								break
							}
						}
						if perfectMatchFound {
							setAttrsInDatasourceSchema(itemObject, d, objType)
							break
						}
					}
					// If match found, or all records processed, break
					if perfectMatchFound || cursor == 0 {
						break
					} else {
						// get next page using cursor
						uri += "&cursor=" + strconv.Itoa(cursor)
						err = nsxtClient.NsxtSession.Get(uri, &obj)
						if err != nil {
							return fmt.Errorf("[ERROR] GET failed for object %s %v", objType, err)
						} else {
							objMap, ok = obj.(map[string]interface{})
							if !ok {
								return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
							}
							// get results of type []interface {}
							results = objMap["results"]
							cursor = objMap["cursor"].(int)
						}
					}
				}
				if !perfectMatchFound {
					return fmt.Errorf("no record found for %s with id '%s', display_name '%s' in scope '%s'", objType, nsxID, displayName, scope)
				}
			}
		} else {
			return fmt.Errorf("either multiple records found for this %s, or the %s is not shared with Project", objType, objType)
		}
	}
	return err
}

func DatasourceReadForVM(d *schema.ResourceData, meta interface{}, objType string, s *schema.Resource) error {
	var obj interface{}
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	displayName := d.Get("display_name").(string)
	externalID := d.Get("external_id").(string)
	powerState := d.Get("power_state").(string)
	var uri string

	// Get the object from NSX using hidden full text search API for specific VPC
	uri = nsxtClient.Config.BasePath + "/search?query=resource_type:" + objType
	if externalID != "" {
		uri += "%20AND%20external_id:" + externalID
	}
	if displayName != "" {
		encodedDisplayName := url.QueryEscape(displayName)
		uri += "%20AND%20display_name:" + encodedDisplayName
	}
	if powerState != "" {
		uri += "%20AND%20power_state:" + powerState
	}
	uri += "&context=vpcs:" + "/orgs/" + nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
	err := nsxtClient.NsxtSession.Get(uri, &obj)
	if err != nil {
		log.Printf("[ERROR] Search failed for object %s %v\n", objType, err)
	} else {
		objMap, ok := obj.(map[string]interface{})
		if !ok {
			return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
		}
		// get results of type []interface {} to get id and policypath
		results := objMap["results"]
		if objMap["result_count"].(float64) == 1 {
			for key, value := range results.([]interface{})[0].(map[string]interface{}) {
				if key == "external_id" {
					d.Set("external_id", value.(string))
					d.SetId(value.(string)) // set terraform ID
				}
				if key == "power_state" {
					d.Set("power_state", value.(string))
				}
				if key == "display_name" {
					d.Set("display_name", value.(string))
				}
			}
		} else {
			return fmt.Errorf("either multiple records found for %s with external_id '%s', display_name '%s', or object is not shared with Project", objType, externalID, displayName)
		}
	}
	return err
}

func convertToSchemaMap(dataMap map[string]interface{}) map[string]*schema.Schema {
	schemaMap := make(map[string]*schema.Schema)

	for key, value := range dataMap {
		// Create a new schema based on the value's type
		var fieldSchema *schema.Schema

		switch value.(type) {
		case string:
			fieldSchema = &schema.Schema{Type: schema.TypeString}
		case int:
			fieldSchema = &schema.Schema{Type: schema.TypeInt}
		case float64:
			fieldSchema = &schema.Schema{Type: schema.TypeFloat}
		case bool:
			fieldSchema = &schema.Schema{Type: schema.TypeBool}
		// Add additional cases for other types as needed
		default:
			// Handle unsupported types or custom schema generation logic
			// we can set the fieldSchema to a generic schema type or define a custom schema based on requirements
			fieldSchema = &schema.Schema{Type: schema.TypeString}
		}

		// Add the schema to the new schema map
		schemaMap[key] = fieldSchema
	}
	return schemaMap
}

/* Populate data from JSON to terraform schema properties. Check for dataType, chec whether property is present in tf schema, if present, then populate. */
func populateTerraformData(key string, value interface{}, fieldSchema *schema.Schema, terraformDataMap map[string]interface{}, terraformDataData *schema.ResourceData, schema_s map[string]*schema.Schema) (interface{}, error) {
	switch fieldSchema.Type {
	case schema.TypeString:
		strValue, ok := value.(string)
		if ok {
			if terraformDataMap != nil {
				terraformDataMap[key] = strValue
			} else {
				terraformDataData.Set(key, strValue)
			}
		} else {
			return nil, fmt.Errorf("invalid data type for field %s: %T", key, value)
		}
	case schema.TypeInt:
		intValue, ok := value.(int)
		if ok {
			if terraformDataMap != nil {
				terraformDataMap[key] = intValue
			} else {
				terraformDataData.Set(key, intValue)
			}
		} else if floatValue, ok := value.(float64); ok {
			if terraformDataMap != nil {
				terraformDataMap[key] = int(floatValue)
			} else {
				terraformDataData.Set(key, int(floatValue))
			}
		} else {
			return nil, fmt.Errorf("invalid data type for field %s: %T", key, value)
		}
	case schema.TypeFloat:
		floatValue, ok := value.(float64)
		if ok {
			if terraformDataMap != nil {
				terraformDataMap[key] = floatValue
			} else {
				terraformDataData.Set(key, floatValue)
			}
		} else {
			return nil, fmt.Errorf("invalid data type for field %s: %T", key, value)
		}
	case schema.TypeBool:
		boolValue, ok := value.(bool)
		if ok {
			if terraformDataMap != nil {
				terraformDataMap[key] = boolValue
			} else {
				terraformDataData.Set(key, boolValue)
			}
		} else {
			return nil, fmt.Errorf("invalid data type for field %s: %T", key, value)
		}
	case schema.TypeList:
		listValue, ok := value.([]interface{})
		if ok {
			listData := make([]interface{}, 0)
			for _, item := range listValue {
				if item != nil {
					if terraformDataMap != nil {
						if itemResource, isResource := item.(*schema.Schema); isResource {
							itemData, err := APIDataToSchema(item, make(map[string]interface{}), itemResource.Elem.(*schema.Resource).Schema)
							if err != nil {
								return nil, err
							}
							listData = append(listData, itemData)
						} else {
							switch item.(type) {
							case string, int, float64, bool:
								listData = append(listData, item)
							case map[string]interface{}:
								itemData, err := APIDataToSchema(item, make(map[string]interface{}), schema_s)
								if err != nil {
									return nil, err
								}
								listData = append(listData, itemData)
							default:
								return nil, fmt.Errorf("invalid data type for field %s: %v", key, item)
							}
						}
					} else {
						// process terraformDataData
						itemData, err := APIDataToSchema(item, make(map[string]interface{}), fieldSchema.Elem.(*schema.Resource).Schema)
						if err != nil {
							return nil, err
						}
						listData = append(listData, itemData)
					}

				}
			}
			if terraformDataMap != nil {
				terraformDataMap[key] = listData
			} else {
				terraformDataData.Set(key, listData)
			}
		} else {
			return nil, fmt.Errorf("invalid data type for field %s: %T", key, value)
		}
	case schema.TypeSet:
		setValue, ok := value.([]interface{})
		if ok {
			setData := make([]interface{}, 0)
			for _, item := range setValue {
				if item != nil {
					switch item.(type) {
					case string, int, float64, bool:
						setData = append(setData, item)
					case map[string]interface{}:
						mapSchemaItem := convertToSchemaMap(item.(map[string]interface{}))
						// Handle the map type item here, recursively call the APIDataToSchema function on the map item
						itemData, err := APIDataToSchema(item, make(map[string]interface{}), mapSchemaItem)
						if err != nil {
							return nil, err
						}
						setData = append(setData, itemData)
					case *schema.Schema:
						schemaItem := item.(*schema.Schema)
						itemData, err := APIDataToSchema(item, make(map[string]interface{}), map[string]*schema.Schema{
							"schema": schemaItem,
						})
						if err != nil {
							return nil, err
						}
						setData = append(setData, itemData)
					case *schema.Resource:
						resourceItem := item.(*schema.Resource)
						itemData, err := APIDataToSchema(item, make(map[string]interface{}), resourceItem.Schema)
						if err != nil {
							return nil, err
						}
						setData = append(setData, itemData)
					default:
						return nil, fmt.Errorf("invalid data type for field %s: %T, expected primitive type or map or *schema.Schema or *schema.Resource", key, item)
					}
				}
			}
			if terraformDataMap != nil {
				terraformDataMap[key] = schema.NewSet(CommonHash, setData)
			} else {
				terraformDataData.Set(key, schema.NewSet(CommonHash, setData))
			}
		}
	case schema.TypeMap:
		mapValue, ok := value.(map[string]interface{})
		if ok {
			mapData, err := APIDataToSchema(mapValue, make(map[string]interface{}), fieldSchema.Elem.(*schema.Schema).Elem.(*schema.Resource).Schema)
			if err != nil {
				return nil, err
			}
			if terraformDataMap != nil {
				terraformDataMap[key] = mapData
			} else {
				terraformDataData.Set(key, mapData)
			}
		} else {
			return nil, fmt.Errorf("invalid data type for field %s: %T", key, value)
		}
	default:
		return nil, fmt.Errorf("unsupported schema type for field %s: %s", key, fieldSchema.Type)
	}
	return nil, nil
}

/* It takes the NSXT JSON data and fills in the terraform data during API read.
It takes input as the top level schema and it uses that to properly create the corresponding terraform resource data
It also checks whether a given nsxt key is defined in the schema before attempting to fill the data. */
func APIDataToSchema(jsonData interface{}, terraformData interface{}, schema_s map[string]*schema.Schema) (interface{}, error) {
	jsonDataMap, ok := jsonData.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid JSON data type: %T", jsonData)
	}

	switch v := terraformData.(type) {
	case map[string]interface{}:
		terraformDataMap := v
		for key, value := range jsonDataMap {
			if fieldSchema, exists := schema_s[key]; exists {
				// Because there is no state during import operation
				populateTerraformData(key, value, fieldSchema, terraformDataMap, nil, schema_s)
			}
		}
		return terraformDataMap, nil

	case *schema.ResourceData:
		terraformDataData := v
		for key, value := range jsonDataMap {
			// Process the key if its present in terraform schema. We don't want all properties from API response, only want the ones in tf resource schema
			if fieldSchema, exists := schema_s[key]; exists {
				populateTerraformData(key, value, fieldSchema, nil, terraformDataData, schema_s)
			}
		}
		return terraformDataData, nil

	default:
		return nil, fmt.Errorf("invalid Terraform data type: %T", terraformData)
	}
}

func CommonHash(v interface{}) int {
	return schema.HashString("nsxt")
}

/* Function that takes the terraform plan data and schema and converts it into NSXT JSON
It recursively resolves the data type of the terraform schema and converts scalar to scalar, Set to dictionary and list to list. */
func SchemaToNsxtData(d interface{}, s interface{}) (interface{}, error) {
	switch dType := d.(type) {
	default:
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range d.(map[string]interface{}) {
			if obj, err := SchemaToNsxtData(v, s.(map[string]*schema.Schema)[k]); err == nil && obj != nil && obj != "" {
				m[k] = obj
			} else if err != nil {
				log.Printf("[ERROR] SchemaToNsxtData %v in parsing k: %v v: %v type: %v", err, k, v, dType)
			}
		}
		return m, nil

	case []interface{}:
		var objList []interface{}
		varray := d.([]interface{})
		var listSchema interface{}
		switch sType := s.(*schema.Schema).Elem.(type) {
		default:
			log.Printf("[DEBUG] Element schema type: %v", sType)
		case *schema.Resource:
			listSchema = s.(*schema.Schema).Elem.(*schema.Resource).Schema
		case *schema.Schema:
			listSchema = s.(*schema.Schema).Elem.(*schema.Schema)
		}
		for i := 0; i < len(varray); i++ {
			obj, err := SchemaToNsxtData(varray[i], listSchema)
			if err == nil && obj != nil {
				objList = append(objList, obj)
			}
		}
		if len(objList) == 0 {
			return nil, nil
		}
		return objList, nil

	case *schema.Set:
		if len(d.(*schema.Set).List()) == 0 {
			return nil, nil
		}
		/**
		Because both Object type and Set type are defined as TypeSet in terraform schema and there is no way to differentiate between the two,
		Get maxItems from schema for that Set. If its 1, means its an object type in schema, otherwise a Set with multiple items allowed.
		**/
		maxItems := s.(*schema.Schema).MaxItems
		if maxItems == 1 {
			// Its an object
			obj, err := SchemaToNsxtData(d.(*schema.Set).List()[0], s.(*schema.Schema).Elem.(*schema.Resource).Schema)
			return obj, err
		}
		// Its a Set
		var objList []interface{}
		varray := d.(*schema.Set).List()
		var setSchema interface{}
		switch setSchemaType := s.(*schema.Schema).Elem.(type) {
		default:
			log.Printf("[DEBUG] Set schema type: %v", setSchemaType)
		case *schema.Resource:
			setSchema = s.(*schema.Schema).Elem.(*schema.Resource).Schema
		case *schema.Schema:
			setSchema = s.(*schema.Schema).Elem.(*schema.Schema)
		}
		for i := 0; i < len(varray); i++ {
			obj, err := SchemaToNsxtData(varray[i], setSchema)
			if err == nil && obj != nil {
				objList = append(objList, obj)
			}
		}
		if len(objList) == 0 {
			return nil, nil
		}
		return objList, nil

	case *schema.ResourceData:
		// In this case the top level schema should be present.
		m := make(map[string]interface{})
		r := d.(*schema.ResourceData)
		for k, v := range s.(map[string]*schema.Schema) {
			if data, ok := r.GetOkExists(k); ok {
				if obj, err := SchemaToNsxtData(data, v); err == nil && obj != nil && obj != "" {
					m[k] = obj
				} else if err != nil {
					log.Printf("[ERROR] SchemaToNsxtData %v in converting k: %v v: %v", err, k, v)
				}
			}
		}
		return m, nil
	}
	/** Return the same object as there is nothing special about **/
	return d, nil
}

// Function to import an existing entity on NSXT into terraform management, so we can manage it using terraform.
func ResourceImporter(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema, path string) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] ResourceImporter objType %v using policy path %v\n", objType, d.Id())
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	if d.Id() != "" {
		// return the ID based import
		return []*schema.ResourceData{d}, nil
	}
	var data interface{}
	results := make([]*schema.ResourceData, 1)
	if path == "" {
		return nil, fmt.Errorf("error in computing policy path for object %s in terraform_policy_utils.ResourceImporter()", objType)
	}
	err1 := nsxtClient.NsxtSession.Get(path, &data)
	if err1 != nil || data == nil {
		log.Printf("[ERROR] ResourceImporter %v in GET of path %v\n", err1, path)
		return nil, err1
	}
	obj := data.(map[string]interface{})
	log.Printf("[DEBUG] ResourceImporter processing obj %v\n", obj)
	result := new(schema.ResourceData)
	if _, err := APIDataToSchema(obj, result, s); err == nil {
		log.Printf("[DEBUG] ResourceImporter Processing obj %v\n", obj)
		id := obj["id"].(string)
		result.SetId(id)
		results[0] = result
	}
	return results, nil
}

// Calculate policy path for different resources in the provider
func ComputePolicyPath(d *schema.ResourceData, objType string, isReadRequest bool, nsxtClient *nsxtclient.NsxtClient, isListingRequest bool) string {
	var url string
	var orgID string
	var projectID string
	var vpcID string
	basePath := nsxtClient.Config.BasePath
	orgID = nsxtClient.Config.OrgID
	projectID = nsxtClient.Config.ProjectID
	vpcID = nsxtClient.Config.VpcID

	switch objType {
	case "Group":
		if isListingRequest {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/groups/"
		} else {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/groups/" + d.Get("nsx_id").(string)
		}
	case "VpcSubnet":
		if isListingRequest {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/subnets/"
		} else {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/subnets/" + d.Get("nsx_id").(string)
		}
	case "VpcSubnetPort":
		if isListingRequest {
			url = basePath + d.Get("parent_path").(string) + "/ports/"
		} else {
			url = basePath + d.Get("parent_path").(string) + "/ports/" + d.Get("nsx_id").(string)
		}
	case "SecurityPolicy":
		if isListingRequest {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/security-policies/"
		} else {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/security-policies/" + d.Get("nsx_id").(string)
		}
	case "SecurityPolicyRule":
		if isListingRequest {
			url = basePath + d.Get("parent_path").(string) + "/rules/"
		} else {
			url = basePath + d.Get("parent_path").(string) + "/rules/" + d.Get("nsx_id").(string)
		}
	case "GatewayPolicy":
		if isListingRequest {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/gateway-policies/"
		} else {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/gateway-policies/" + d.Get("nsx_id").(string)
		}
	case "GatewayPolicyRule":
		if isListingRequest {
			url = basePath + d.Get("parent_path").(string) + "/rules/"
		} else {
			url = basePath + d.Get("parent_path").(string) + "/rules/" + d.Get("nsx_id").(string)
		}
	case "PolicyVpcNatRule":
		if isListingRequest {
			url = basePath + d.Get("parent_path").(string) + "/nat-rules/"
		} else {
			url = basePath + d.Get("parent_path").(string) + "/nat-rules/" + d.Get("nsx_id").(string)
		}
	case "IpAddressAllocation":
		if isListingRequest {
			url = basePath + d.Get("parent_path").(string) + "/ip-allocations/"
		} else {
			url = basePath + d.Get("parent_path").(string) + "/ip-allocations/" + d.Get("nsx_id").(string)
		}
	case "VpcIpAddressAllocation":
		if isListingRequest {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/ip-address-allocations/"
		} else {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/ip-address-allocations/" + d.Get("nsx_id").(string)
		}
	case "DhcpV4StaticBindingConfig":
		if isListingRequest {
			url = basePath + d.Get("parent_path").(string) + "/dhcp-static-binding-configs/"
		} else {
			url = basePath + d.Get("parent_path").(string) + "/dhcp-static-binding-configs/" + d.Get("nsx_id").(string)
		}
	case "DhcpV6StaticBindingConfig":
		if isListingRequest {
			url = basePath + d.Get("parent_path").(string) + "/dhcp-static-binding-configs/"
		} else {
			url = basePath + d.Get("parent_path").(string) + "/dhcp-static-binding-configs/" + d.Get("nsx_id").(string)
		}
	case "StaticRoutes":
		if isListingRequest {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/static-routes/"
		} else {
			url = basePath + "/orgs/" + orgID + "/projects/" + projectID + "/vpcs/" + vpcID + "/static-routes/" + d.Get("nsx_id").(string)
		}
	}
	return url
}

// It sets default values in the terraform resources to avoid diffs for scalars.
func SetDefaultsInAPIRes(apiRes interface{}, dLocal interface{}, s map[string]*schema.Schema) (interface{}, error) {
	if apiRes == nil {
		log.Printf("[ERROR] SetDefaultsInAPIRes got nil for %v", s)
		return apiRes, nil
	}
	switch dLocal.(type) {
	default:
	case map[string]interface{}:
		for k, v := range dLocal.(map[string]interface{}) {
			switch v.(type) {
			// Getting key, value for given dLocal
			default:
				if _, ok := apiRes.(map[string]interface{})[k]; !ok {
					// Cheking if field is present in schema
					if dval, ok := s[k]; ok {
						// Getting default values from schema
						defaultVal, err := dval.DefaultValue()
						if err != nil {
							log.Printf("[ERROR] SetDefaultsInAPIRes did not get default value from schema err %v %v",
								err, defaultVal)
						} else {
							if defaultVal != nil {
								apiRes.(map[string]interface{})[k] = defaultVal
							}
						}
					}
				}
			// dLocal nested dictionary.
			case map[string]interface{}:
				if s2, ok := s[k]; ok {
					switch s2.Elem.(type) {
					default:
					case *schema.Resource:
						if apiRes.(map[string]interface{})[k] != nil {
							apiRes1, err := SetDefaultsInAPIRes(apiRes.(map[string]interface{})[k], v, s2.Elem.(*schema.Resource).Schema)
							if err != nil {
								log.Printf("[ERROR] SetDefaultsInAPIRes %v", err)
							} else {
								apiRes.(map[string]interface{})[k] = apiRes1
							}
						} else {
							apiRes.(map[string]interface{})[k] = v
						}
					}
				}
			// dLocal is array of dictionaries.
			case []interface{}:
				var objList []interface{}
				if apiRes.(map[string]interface{})[k] != nil {
					varray2 := apiRes.(map[string]interface{})[k].([]interface{})
					//getting schema for nested object.
					s2, err := s[k]
					var dst, src []interface{}
					//As err returned is boolean value
					if !err {
						log.Printf("[ERROR] SetDefaultsInAPIRes in fetching k %v err %v", k, err)
					}
					if len(varray2) > len(v.([]interface{})) {
						dst = varray2
						src = v.([]interface{})
					} else {
						dst = v.([]interface{})
						src = varray2
					}
					for x, y := range src {
						switch s2.Elem.(type) {
						default:
						case *schema.Resource:
							obj, err := SetDefaultsInAPIRes(dst[x], y, s2.Elem.(*schema.Resource).Schema)
							if err != nil {
								log.Printf("[ERROR] SetDefaultsInAPIRes err %v in x %v y %v", err, x, y)
							} else {
								objList = append(objList, obj)
							}
						case *schema.Schema:
							objList = append(objList, src[x])
						}
					}
				}
				apiRes.(map[string]interface{})[k] = objList
			}
		}
	}
	return apiRes, nil
}
