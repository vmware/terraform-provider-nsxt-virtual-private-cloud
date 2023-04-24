/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

//nolint
package nsxt

import (
	"encoding/json"
	"fmt"
	"log"

	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVmTagsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"virtual_machine_id": {
			Type:        schema.TypeString,
			Description: "External ID of the VM",
			Required:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     resourceTagSchema(),
		},
	}
}

func resourceNsxtVpcVmTags() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcVmTagsCreate,
		Read:   resourceNsxtVpcVmTagsRead,
		Update: resourceNsxtVpcVmTagsUpdate,
		Delete: resourceNsxtVpcVmTagsDelete,
		Schema: resourceVmTagsSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceVmTagsImporter,
		},
	}
}

func resourceVmTagsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if d.Id() != "" {
		// return the ID based import
		return []*schema.ResourceData{d}, nil
	}
	return nil, nil
}

func resourceNsxtVpcVmTagsRead(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	vm_external_id := d.Id()
	var obj interface{}
	if vm_external_id == "" {
		return fmt.Errorf("error obtaining Virtual Machine external ID")
	}
	url := nsxtClient.Config.BasePath + "/search?query=resource_type:VirtualMachine%20AND%20external_id:" + vm_external_id + "&context=vpcs:/orgs/" +
		nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
	err := nsxtClient.NsxtSession.Get(url, &obj)
	if err != nil {
		d.SetId("")
		log.Printf("[ERROR] in reading VM %v\n", err)
	} else {
		d.SetId(vm_external_id)
	}
	return err
}

func resourceNsxtVpcVmTagsCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceVmTagsSchema()
	external_id := d.Get("virtual_machine_id").(string)
	var robj interface{}
	if data, err := SchemaToNsxtData(d, s); err == nil {
		nsxtClient := meta.(*nsxtclient.NsxtClient)
		url := nsxtClient.Config.BasePath + "/infra/realized-state/enforcement-points/" + nsxtClient.NsxtSession.GetEnforcementPoint() + "/virtual-machines?action=update_tags"
		err := nsxtClient.NsxtSession.Post(url, data, &robj)
		if err != nil {
			log.Printf("[ERROR] VM tag addition failed %v\n", err)
		} else {
			d.SetId(external_id)
			resourceNsxtVpcVmTagsRead(d, meta)
		}
		return err
	} else {
		log.Printf("[ERROR] VM tag SchemaToNsxtData: Error %v", err)
		return err
	}
}

func resourceNsxtVpcVmTagsUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceNsxtVpcVmTagsCreate(d, meta)
}

func resourceNsxtVpcVmTagsDelete(d *schema.ResourceData, meta interface{}) error {
	tags_deleted := true
	var robj interface{}

	// Build JSON interface with empty value in tags, so VM gets updated with all tags removed
	dataMap := map[string]interface{}{"virtual_machine_id": d.Id(), "tags": make([]interface{}, 0)}
	jsonData, _ := json.Marshal(dataMap)
	var jsonInterface interface{}
	json.Unmarshal(jsonData, &jsonInterface)

	nsxtClient := meta.(*nsxtclient.NsxtClient)
	url := nsxtClient.Config.BasePath + "/infra/realized-state/enforcement-points/" + nsxtClient.NsxtSession.GetEnforcementPoint() + "/virtual-machines?action=update_tags"
	err := nsxtClient.NsxtSession.Post(url, jsonInterface, &robj)
	if err != nil {
		d.SetId("")
		log.Printf("[ERROR] VM tag deletion failed %v\n", err)
	} else {
		// Check if all tags are deleted successfully
		url := nsxtClient.Config.BasePath + "/search?query=resource_type:VirtualMachine%20AND%20external_id:" + d.Id() + "&context=vpcs:/orgs/" +
			nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
		err := nsxtClient.NsxtSession.Get(url, &robj)
		if err == nil {
			objMap, ok := robj.(map[string]interface{})
			if !ok {
				return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
			}
			// get results of type []interface {}
			results := objMap["results"]
			if objMap["result_count"].(float64) == 1 {
				for key := range results.([]interface{})[0].(map[string]interface{}) {
					if key == "tags" {
						tags_deleted = false
						break
					}
				}
			}
			if tags_deleted {
				d.SetId("")
				return nil
			} else {
				return fmt.Errorf("VM tag deletion failed")
			}
		}
	}
	return err
}

//nolint
