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

func resourceGatewayPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"category": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"comments": {
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
		"locked": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"scheduler_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"sequence_number": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"stateful": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     resourceTagSchema(),
		},
		"tcp_strict": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
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
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object GatewayPolicy %v\n", err)
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
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource GatewayPolicy not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint
