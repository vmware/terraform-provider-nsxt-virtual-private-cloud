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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpcIpAddressAllocationSchema() map[string]*schema.Schema {
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
		"ip_address_block_visibility": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"PUBLIC", "PRIVATE"}, false),
			Default:      "PUBLIC",
		},
		"ip_address_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"IPV4", "IPV6"}, false),
			Default:      "IPV4",
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

func resourceNsxtVpcIpAddressAllocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcIpAddressAllocationCreate,
		Read:   resourceNsxtVpcIpAddressAllocationRead,
		Update: resourceNsxtVpcIpAddressAllocationUpdate,
		Delete: resourceNsxtVpcIpAddressAllocationDelete,
		Schema: resourceVpcIpAddressAllocationSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceVpcIpAddressAllocationImporter,
		},
	}
}

func resourceVpcIpAddressAllocationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceVpcIpAddressAllocationSchema()
	return ResourceImporter(d, m, "VpcIpAddressAllocation", s, d.Id())
}

func resourceNsxtVpcIpAddressAllocationRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcIpAddressAllocationSchema()
	err := APIRead(d, meta, "VpcIpAddressAllocation", s)
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object VpcIpAddressAllocation %v\n", err)
	}
	return err
}

func resourceNsxtVpcIpAddressAllocationCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcIpAddressAllocationSchema()
	err := APICreateOrUpdate(d, meta, "VpcIpAddressAllocation", s)
	if err == nil {
		err = resourceNsxtVpcIpAddressAllocationRead(d, meta)
	}
	return err
}

func resourceNsxtVpcIpAddressAllocationUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceNsxtVpcIpAddressAllocationDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource VpcIpAddressAllocation not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint