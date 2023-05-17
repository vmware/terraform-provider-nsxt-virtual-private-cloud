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
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
)

func resourceDhcpV4StaticBindingConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"options": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     resourceDhcpV4OptionsSchema(),
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
		"host_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"gateway_address": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip_address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"mac_address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 30,
			Elem:     resourceTagSchema(),
		},
		"lease_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  86400,
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

func resourceNsxtVpcDhcpV4StaticBindingConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcDhcpV4StaticBindingConfigCreate,
		Read:   resourceNsxtVpcDhcpV4StaticBindingConfigRead,
		Update: resourceNsxtVpcDhcpV4StaticBindingConfigUpdate,
		Delete: resourceNsxtVpcDhcpV4StaticBindingConfigDelete,
		Schema: resourceDhcpV4StaticBindingConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceDhcpV4StaticBindingConfigImporter,
		},
	}
}

func resourceDhcpV4StaticBindingConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceDhcpV4StaticBindingConfigSchema()
	return ResourceImporter(d, m, "DhcpV4StaticBindingConfig", s, d.Id())
}

func resourceNsxtVpcDhcpV4StaticBindingConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceDhcpV4StaticBindingConfigSchema()
	err := APIRead(d, meta, "DhcpV4StaticBindingConfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceNsxtVpcDhcpV4StaticBindingConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceDhcpV4StaticBindingConfigSchema()
	err := APICreateOrUpdate(d, meta, "DhcpV4StaticBindingConfig", s)
	if err == nil {
		err = resourceNsxtVpcDhcpV4StaticBindingConfigRead(d, meta)
	}
	return err
}

func resourceNsxtVpcDhcpV4StaticBindingConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceDhcpV4StaticBindingConfigSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "DhcpV4StaticBindingConfig", s)
	if err == nil {
		err = resourceNsxtVpcDhcpV4StaticBindingConfigRead(d, meta)
	}
	return err
}

func resourceNsxtVpcDhcpV4StaticBindingConfigDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] resourceNsxtVpcDhcpV4StaticBindingConfigDelete not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}
