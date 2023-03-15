/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// Auto generated code. DO NOT EDIT.

//nolint
package nsxt

import (
	"log"
	"strings"

	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceStaticRoutesSchema() map[string]*schema.Schema {
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
		"network": {
			Type:     schema.TypeString,
			Required: true,
		},
		"next_hops": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     resourceRouterNexthopSchema(),
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

func resourceNsxtVpcStaticRoutes() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcStaticRoutesCreate,
		Read:   resourceNsxtVpcStaticRoutesRead,
		Update: resourceNsxtVpcStaticRoutesUpdate,
		Delete: resourceNsxtVpcStaticRoutesDelete,
		Schema: resourceStaticRoutesSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceStaticRoutesImporter,
		},
	}
}

func resourceStaticRoutesImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceStaticRoutesSchema()
	return ResourceImporter(d, m, "StaticRoutes", s, d.Id())
}

func resourceNsxtVpcStaticRoutesRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceStaticRoutesSchema()
	err := APIRead(d, meta, "StaticRoutes", s)
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object StaticRoutes %v\n", err)
	}
	return err
}

func resourceNsxtVpcStaticRoutesCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceStaticRoutesSchema()
	err := APICreateOrUpdate(d, meta, "StaticRoutes", s)
	if err == nil {
		err = resourceNsxtVpcStaticRoutesRead(d, meta)
	}
	return err
}

func resourceNsxtVpcStaticRoutesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceStaticRoutesSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "StaticRoutes", s)
	if err == nil {
		err = resourceNsxtVpcStaticRoutesRead(d, meta)
	}
	return err
}

func resourceNsxtVpcStaticRoutesDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource StaticRoutes not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint
