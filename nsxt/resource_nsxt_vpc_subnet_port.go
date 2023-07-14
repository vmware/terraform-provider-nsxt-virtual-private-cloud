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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpcSubnetPortSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"attachment": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem:     resourcePortAttachmentSchema(),
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
		"init_state": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"UNBLOCKED_VLAN", "RESTORE_VIF"}, false),
			Computed:     true,
		},
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"admin_state": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"UP", "DOWN"}, false),
			Default:      "UP",
		},
		"extra_configs": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     resourceSegmentExtraConfigSchema(),
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 30,
			Elem:     resourceTagSchema(),
		},
		"address_bindings": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 512,
			Elem:     resourcePortAddressBindingEntrySchema(),
		},
		"ignored_address_bindings": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MinItems: 0,
			MaxItems: 16,
			Elem:     resourcePortAddressBindingEntrySchema(),
		},
		"parent_path": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
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

func resourceNsxtVpcSubnetPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcSubnetPortCreate,
		Read:   resourceNsxtVpcSubnetPortRead,
		Update: resourceNsxtVpcSubnetPortUpdate,
		Delete: resourceNsxtVpcSubnetPortDelete,
		Schema: resourceVpcSubnetPortSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceVpcSubnetPortImporter,
		},
	}
}

func resourceVpcSubnetPortImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceVpcSubnetPortSchema()
	return ResourceImporter(d, m, "VpcSubnetPort", s, d.Id())
}

func resourceNsxtVpcSubnetPortRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcSubnetPortSchema()
	err := APIRead(d, meta, "VpcSubnetPort", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object VpcSubnetPort %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtVpcSubnetPortCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcSubnetPortSchema()
	err := APICreateOrUpdate(d, meta, "VpcSubnetPort", s)
	if err == nil {
		err = resourceNsxtVpcSubnetPortRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetPortUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcSubnetPortSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "VpcSubnetPort", s)
	if err == nil {
		err = resourceNsxtVpcSubnetPortRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetPortDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if 'object not found' or 'forbidden' or 'success with no response' response occurs, terraform should swallow it and not fail apply on object, else throw error and fail
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Error occurred in Delete for resource VpcSubnetPort \n")
			return err
		}
		d.SetId("")
	}
	return nil
}
