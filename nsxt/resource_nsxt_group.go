/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// Auto generated code. DO NOT EDIT.

// nolint
package nsxt

import (
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"expression": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     resourceExpressionCustomSchema(),
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 30,
			Elem:     resourceTagSchema(),
		},
		"nsx_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func resourceNsxtVpcGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcGroupCreate,
		Read:   resourceNsxtVpcGroupRead,
		Update: resourceNsxtVpcGroupUpdate,
		Delete: resourceNsxtVpcGroupDelete,
		Schema: resourceGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceGroupImporter,
		},
	}
}

func resourceGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceGroupSchema()
	return ResourceImporter(d, m, "Group", s, d.Id())
}

func resourceNsxtVpcGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceGroupSchema()
	err := APIRead(d, meta, "Group", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object Group %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtVpcGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGroupSchema()
	err := APICreateOrUpdate(d, meta, "Group", s)
	if err == nil {
		err = resourceNsxtVpcGroupRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "Group", s)
	if err == nil {
		err = resourceNsxtVpcGroupRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGroupDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if 'object not found' or 'forbidden' or 'success with no response' response occurs, terraform should swallow it and not fail apply on object, else throw error and fail
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Error occurred in Delete for resource Group \n")
			return err
		}
		d.SetId("")
	}
	return nil
}
