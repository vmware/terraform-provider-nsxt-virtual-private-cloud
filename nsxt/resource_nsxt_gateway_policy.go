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

func resourceGatewayPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"category": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"comments": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"tcp_strict": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"scheduler_path": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"stateful": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"locked": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
		"sequence_number": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tags": {
			Type:     schema.TypeSet,
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

func resourceNsxtVpcGatewayPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcGatewayPolicyCreate,
		Read:   resourceNsxtVpcGatewayPolicyRead,
		Update: resourceNsxtVpcGatewayPolicyUpdate,
		Delete: resourceNsxtVpcGatewayPolicyDelete,
		Schema: resourceGatewayPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: resourceGatewayPolicyImporter,
		},
	}
}

func resourceGatewayPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceGatewayPolicySchema()
	return ResourceImporter(d, m, "GatewayPolicy", s, d.Id())
}

func resourceNsxtVpcGatewayPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceGatewayPolicySchema()
	err := APIRead(d, meta, "GatewayPolicy", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object GatewayPolicy %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtVpcGatewayPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGatewayPolicySchema()
	err := APICreateOrUpdate(d, meta, "GatewayPolicy", s)
	if err == nil {
		err = resourceNsxtVpcGatewayPolicyRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGatewayPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGatewayPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "GatewayPolicy", s)
	if err == nil {
		err = resourceNsxtVpcGatewayPolicyRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGatewayPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if object not found errors occur, terraform should swallow it and not fail apply on object
		if err != nil && (strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[WARNING] Resource GatewayPolicy not found on backend\n")
			return nil
		} else if err != nil {
			return err
		}
		d.SetId("")
	}
	return nil
}
