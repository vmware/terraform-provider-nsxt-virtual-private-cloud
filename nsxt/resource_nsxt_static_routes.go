/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// Auto generated code. DO NOT EDIT. KIRAN

// nolint
package nsxt

import (
	nsxtclient "github.com/vmware/terraform-provider-nsxt-virtual-private-cloud/nsxt/clients"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceStaticRoutesSchema() map[string]*schema.Schema {
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
		"network": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 30,
			Elem:     resourceTagSchema(),
		},
		"next_hops": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem:     resourceRouterNexthopSchema(),
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
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object StaticRoutes %v\n", err)
		return nil
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
		// if 'object not found' or 'forbidden' or 'success with no response' response occurs, terraform should swallow it and not fail apply on object, else throw error and fail
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Error occurred in Delete for resource StaticRoutes \n")
			return err
		}
		d.SetId("")
	}
	return nil
}
