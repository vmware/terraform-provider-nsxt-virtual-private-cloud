/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
package nsxt

import (
	"log"
	"strings"

	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSubnetIpAddressAllocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"allocation_ip": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
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
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object IpAddressAllocation %v\n", err)
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
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource IpAddressAllocation not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint
