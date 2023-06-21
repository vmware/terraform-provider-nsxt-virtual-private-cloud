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

func resourceSubnetIpAddressAllocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"sync_realization": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"allocation_ip": {
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
		"tags": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 30,
			Elem:     resourceTagSchema(),
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

func resourceNsxtVpcSubnetIpAddressAllocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcSubnetIpAddressAllocationCreate,
		Read:   resourceNsxtVpcSubnetIpAddressAllocationRead,
		Update: resourceNsxtVpcSubnetIpAddressAllocationUpdate,
		Delete: resourceNsxtVpcSubnetIpAddressAllocationDelete,
		Schema: resourceSubnetIpAddressAllocationSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceIpAddressAllocationImporter,
		},
	}
}

func resourceIpAddressAllocationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceSubnetIpAddressAllocationSchema()
	return ResourceImporter(d, m, "IpAddressAllocation", s, d.Id())
}

func resourceNsxtVpcSubnetIpAddressAllocationRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceSubnetIpAddressAllocationSchema()
	err := APIRead(d, meta, "IpAddressAllocation", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object IpAddressAllocation %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtVpcSubnetIpAddressAllocationCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSubnetIpAddressAllocationSchema()
	err := APICreateOrUpdate(d, meta, "IpAddressAllocation", s)
	if err == nil {
		err = resourceNsxtVpcSubnetIpAddressAllocationRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetIpAddressAllocationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSubnetIpAddressAllocationSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "IpAddressAllocation", s)
	if err == nil {
		err = resourceNsxtVpcSubnetIpAddressAllocationRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetIpAddressAllocationDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if object not found errors occur, terraform should swallow it and not fail apply on object
		if err != nil && (strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[WARNING] Resource IpAddressAllocation not found on backend\n")
			return nil
		} else if err != nil {
			return err
		}
		d.SetId("")
	}
	return nil
}
