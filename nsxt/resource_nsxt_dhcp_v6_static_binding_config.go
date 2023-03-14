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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
)

func resourceDhcpV6StaticBindingConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_revision": {
			Type:     schema.TypeInt,
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
		"dns_nameservers": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"domain_names": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ip_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"lease_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  86400,
		},
		"mac_address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"preferred_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"sntp_servers": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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
		"parent_path": {
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

func resourceNsxtVpcDhcpV6StaticBindingConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcDhcpV6StaticBindingConfigCreate,
		Read:   resourceNsxtVpcDhcpV6StaticBindingConfigRead,
		Update: resourceNsxtVpcDhcpV6StaticBindingConfigUpdate,
		Delete: resourceNsxtVpcDhcpV6StaticBindingConfigDelete,
		Schema: resourceDhcpV6StaticBindingConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceDhcpV6StaticBindingConfigImporter,
		},
	}
}

func resourceDhcpV6StaticBindingConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceDhcpV6StaticBindingConfigSchema()
	return ResourceImporter(d, m, "DhcpV6StaticBindingConfig", s, d.Id())
}

func resourceNsxtVpcDhcpV6StaticBindingConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceDhcpV6StaticBindingConfigSchema()
	err := APIRead(d, meta, "DhcpV6StaticBindingConfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceNsxtVpcDhcpV6StaticBindingConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceDhcpV6StaticBindingConfigSchema()
	err := APICreateOrUpdate(d, meta, "DhcpV6StaticBindingConfig", s)
	if err == nil {
		err = resourceNsxtVpcDhcpV6StaticBindingConfigRead(d, meta)
	}
	return err
}

func resourceNsxtVpcDhcpV6StaticBindingConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceDhcpV6StaticBindingConfigSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "DhcpV6StaticBindingConfig", s)
	if err == nil {
		err = resourceNsxtVpcDhcpV6StaticBindingConfigRead(d, meta)
	}
	return err
}

func resourceNsxtVpcDhcpV6StaticBindingConfigDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] resourceNsxtVpcDhcpV6StaticBindingConfigDelete not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}
