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
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
)

// It is generic API to create and update any Nsxt REST resource. If the resource does not exist it will try to
// create it. In case, it is present then automatically converts to PATCH semantics.
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
		pathWithoutQueryParamsTrimmed := strings.TrimPrefix(pathWithoutQueryParams, "policy/api/v1")

		if resourceID != "" {
			// resource with id already present in NSX, this is an update request
			log.Printf("[INFO] APICreateOrUpdate: Updating obj %v with path %s\n", objType, path)
			err = nsxtClient.NsxtSession.Patch(path, data, &robj)
			if err != nil {
				log.Printf("[ERROR] APICreateOrUpdate updation failed %v\n", err)
			} else {
				d.SetId(pathWithoutQueryParams)
				d.Set("path", pathWithoutQueryParamsTrimmed)
			}
		} else {
			// resource is new, this is a create request
			log.Printf("[INFO] APICreateOrUpdate: Creating obj %v schema %v data %v with path %s\n", objType, d, data, path)
			err = nsxtClient.NsxtSession.Patch(path, data, &robj)
			if err != nil {
				log.Printf("[ERROR] APICreateOrUpdate creation failed %v\n", err)
			} else {
				d.SetId(pathWithoutQueryParams)
				d.Set("path", pathWithoutQueryParamsTrimmed)
			}
		}
		return err
	} else { //nolint
		log.Printf("[ERROR] APICreateOrUpdate: Error %v", err)
		return err
	}
}

func APIRead(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	var obj interface{}
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	// calculate the policy path of resource
	id := d.Id()
	resourceName := d.Get("display_name").(string)

	if id != "" {
		// Read using ID as complete policy path
		log.Printf("[DEBUG] APIRead reading object with id %v\n", id)
		err := nsxtClient.NsxtSession.Get(id, &obj)
		log.Printf("json unmarshal response: %v\n", obj)
		if err == nil && obj != nil {
			d.SetId(id)
			// set readonly property _revision for terraform show
			objMap, ok := obj.(map[string]interface{})
			if !ok {
				return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
			}
		} else {
			log.Printf("Read not successful for " + objType)
			d.SetId("")
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
			log.Printf("Read not successful for " + objType)
			d.SetId("")
		}
	} else {
		log.Printf("[DEBUG] APIRead object ID not present for resource %s\n", objType)
		d.SetId("")
	}

	if localData, err := SchemaToNsxtData(d, s); err == nil {
		modAPIRes, err := SetDefaultsInAPIRes(obj, localData, s)
		if err != nil {
			log.Printf("[ERROR] APIRead in modifying api response object %v\n", err)
		}
		if _, err := APIDataToSchema(modAPIRes, d, s); err != nil {
			log.Printf("[ERROR] APIRead in setting read object %v\n", err)
			d.SetId("")
		}
		log.Printf("[DEBUG] type: %v localData : %v", objType, localData)
		log.Printf("[DEBUG] type: %v modAPIRes: %v", objType, modAPIRes)
	}

	return nil
}

func DatasourceRead(d *schema.ResourceData, meta interface{}, objType string, s *schema.Resource) error {
	var obj interface{}
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	displayName := d.Get("display_name").(string)
	nsxID := d.Get("nsx_id").(string)
	isInfraObject := false
	isProjectInfra := false
	var url string
	if strings.HasPrefix(objType, "ProjectInfra") {
		objType = strings.TrimPrefix(objType, "ProjectInfra")
		isProjectInfra = true
	} else if strings.HasPrefix(objType, "Infra") {
		objType = strings.TrimPrefix(objType, "Infra")
		isInfraObject = true
	}
	// Some objTypes (should be resource_type value) are differently worded in .yaml, hence correcting them for usage in Search query
	switch objType {
	case "SecurityPolicyRule":
		objType = "Rule"
	case "GatewayPolicyRule":
		objType = "Rule"
	}

	// Get the object from NSX using hidden full text search API
	url = nsxtClient.Config.BasePath + "/search?query=resource_type:" + objType
	if nsxID != "" {
		url += "%20AND%20id:" + nsxID
	}
	if displayName != "" {
		url += "%20AND%20display_name:" + displayName
	}
	if d.Get("parent_path") != nil {
		parentPath := d.Get("parent_path").(string)
		if parentPath != "" {
			url += "%20AND%20parent_path:\"" + parentPath + "\""
		}
	}
	if isProjectInfra {
		// If it is project/infra object, search in context of project
		url += "&context=projects:" + "/orgs/" + nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID
	} else if !isInfraObject && !isProjectInfra {
		// If not infra and project/infra object, search in context of VPC
		url += "&context=vpcs:" + "/orgs/" + nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
	}
	err := nsxtClient.NsxtSession.Get(url, &obj)
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
			for key, value := range results.([]interface{})[0].(map[string]interface{}) {
				// This API returns all default objects + project shared objects + VPC objects. We need to filter based on path to get only what datasource desires
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
		} else if objMap["result_count"].(float64) == 0 {
			// get the entity using listing api, for some VPC resources(eg VpcIpAddressAllocation), search API isn't indexed
			url = ComputePolicyPath(d, objType, false, nsxtClient, true)
			// keep maximum allowed page size for pagination
			url += "?page_size=1000"
			err := nsxtClient.NsxtSession.Get(url, &obj)
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
					for index, itemObject := range results.([]interface{}) {
						log.Printf("itemObject: %v at index %d\n", itemObject, index)
						for key, value := range itemObject.(map[string]interface{}) {
							log.Printf("key: %s at value %v\n", key, value)
							// Filter based on nsx_id or display_name to get only what datasource desires
							//TODO: Check how to refine search in case of multiple entries found with same display_name
							if (nsxID != "" && key == "id" && value == nsxID) || (displayName != "" && key == "display_name" && value == displayName) {
								perfectMatchFound = true
								break
							}
						}
						if perfectMatchFound {
							for key, value := range itemObject.(map[string]interface{}) {
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
							break
						}
					}
					// If match found, or all records processed, break
					if perfectMatchFound || cursor == 0 {
						break
					} else {
						// get next page using cursor
						url += "&cursor=" + strconv.Itoa(cursor)
						err = nsxtClient.NsxtSession.Get(url, &obj)
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
					return fmt.Errorf("no record found for %s with display_name '%s'", objType, displayName)
				}
			}
		} else {
			return fmt.Errorf("either multiple records found for %s with display_name '%s', or object is not shared with Project", objType, displayName)
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
	var url string

	// Get the object from NSX using hidden full text search API for specific VPC
	url = nsxtClient.Config.BasePath + "/search?query=resource_type:" + objType
	if externalID != "" {
		url += "%20AND%20external_id:" + externalID
	}
	if displayName != "" {
		url += "%20AND%20display_name:" + displayName
	}
	if powerState != "" {
		url += "%20AND%20power_state:" + powerState
	}
	url += "&context=vpcs:" + "/orgs/" + nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
	err := nsxtClient.NsxtSession.Get(url, &obj)
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
			return fmt.Errorf("either multiple records found for %s with display_name '%s', or object is not shared with Project", objType, displayName)
		}
	}
	return err
}

// It takes the Nsxt JSON data and fills in the terraform data during API read.
// It takes input as the top level schema and it uses that to properly create the corresponding terraform resource data
// It also checks whether a given nsxt key is defined in the schema before attempting to fill the data.
func APIDataToSchema(adata interface{}, d interface{}, t map[string]*schema.Schema) (interface{}, error) {
	switch adata.(type) {
	default:
	case map[string]interface{}:
		// resolve d interface into a set
		if t == nil {
			m := map[string]interface{}{}
			for k, v := range adata.(map[string]interface{}) {
				if obj, err := APIDataToSchema(v, nil, nil); err == nil {
					m[k] = obj
				} else if err != nil {
					log.Printf("[ERROR] APIDataToSchema %v in converting k: %v v: %v", err, k, v)
				}
			}
			objs := []interface{}{}
			objs = append(objs, m)
			s := schema.NewSet(CommonHash, objs)
			return s, nil
		} else { //nolint
			for k, v := range adata.(map[string]interface{}) {
				if _, ok := t[k]; ok {
					// found in the schema
					if obj, err := APIDataToSchema(v, nil, nil); err == nil {
						switch dType := d.(type) {
						default:
						case *schema.ResourceData:
							objType := reflect.TypeOf(obj).Kind()
							if objType != reflect.Map && objType != reflect.Array && objType != reflect.Slice {
								// object is primitive type, set directy into schema if tf resource has the key in the schema
								if err := d.(*schema.ResourceData).Set(k, obj); err != nil {
									log.Printf("[ERROR] APIDataToSchema %v in setting %v type %v", err, obj, dType)
								}
							} else {
								objList := obj.([]interface{})
								isObjectMap := false
								for _, v := range objList {
									if m, ok := v.(map[string]interface{}); ok {
										isObjectMap = true
										for k, v := range m {
											if nestedMap, ok := v.(map[string]interface{}); ok {
												for nestedKey, nestedValue := range nestedMap {
													fmt.Printf("%s.%s.%s = %v\n", k, nestedKey, reflect.TypeOf(nestedValue), nestedValue)
													if _, ok := d.(*schema.ResourceData).State().Attributes[nestedKey]; ok {
														// Set the value in the schema
														if err := d.(*schema.ResourceData).Set(nestedKey, nestedValue); err != nil {
															log.Printf("[ERROR] APIDataToSchema %v in setting %v   type %v", err, obj, dType)
														}
													}
												}
											}
										}
									}
								}
								if !isObjectMap {
									if err := d.(*schema.ResourceData).Set(k, obj); err != nil {
										log.Printf("[ERROR] APIDataToSchema %v in setting %v type %v", err, obj, dType)
									}
								}
							}
						case map[string]interface{}:
							d.(map[string]interface{})[k] = obj
						}
					}
				}
			}
			return d, nil
		}
	case []interface{}:
		var objList []interface{}
		varray := adata.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := APIDataToSchema(varray[i], nil, nil)
			if err == nil {
				switch objType := obj.(type) {
				default:
					log.Printf("Appending objtype %v to list %v", objType, objList)
					objList = append(objList, obj)
				case *schema.Set:
					objList = append(objList, obj.(*schema.Set).List()[0])
				}
			} else {
				log.Printf("[ERROR] APIDataToSchema %v", err)
			}
		}
		return objList, nil
		/** Return the same object as there is nothing special about **/
	}
	return adata, nil
}

func CommonHash(v interface{}) int {
	return schema.HashString("nsxt")
}

// It takes the terraform plan data and schema and converts it into Nsxt JSON
// It recursively resolves the data type of the terraform schema and converts scalar to scalar, Set to dictionary,
// and list to list.
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
			log.Printf("%v", sType)
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
		var objList []interface{}
		varray := d.(*schema.Set).List()
		var setSchema interface{}
		switch setSchemaType := s.(*schema.Schema).Elem.(type) {
		default:
			log.Printf("%v", setSchemaType)
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
