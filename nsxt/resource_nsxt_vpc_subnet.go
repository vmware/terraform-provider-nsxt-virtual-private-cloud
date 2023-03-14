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

func resourceVpcSubnetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"access_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"Private", "Public", "Isolated"}, false),
			Default:      "Private",
		},
		"advanced_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     resourceSubnetAdvancedConfigSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dhcp_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     resourceDhcpConfigSchema(),
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ipv4_subnet_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"skip_ipam": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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

func resourceNsxtVpcSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcSubnetCreate,
		Read:   resourceNsxtVpcSubnetRead,
		Update: resourceNsxtVpcSubnetUpdate,
		Delete: resourceNsxtVpcSubnetDelete,
		Schema: resourceVpcSubnetSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceVpcSubnetImporter,
		},
	}
}

func resourceVpcSubnetImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceVpcSubnetSchema()
	return ResourceImporter(d, m, "VpcSubnet", s, d.Id())
}

func resourceNsxtVpcSubnetRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcSubnetSchema()
	err := APIRead(d, meta, "VpcSubnet", s)
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object VpcSubnet %v\n", err)
	}
	return err
}

func resourceNsxtVpcSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcSubnetSchema()
	err := APICreateOrUpdate(d, meta, "VpcSubnet", s)
	if err == nil {
		err = resourceNsxtVpcSubnetRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcSubnetSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "VpcSubnet", s)
	if err == nil {
		err = resourceNsxtVpcSubnetRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource VpcSubnet not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint