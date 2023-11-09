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
	nsxtclient "github.com/vmware/terraform-provider-nsxt-virtual-private-cloud/nsxt/clients"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGatewayPolicyRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tag": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"disabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"logged": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"destinations_excluded": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"sources_excluded": {
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
		"notes": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"action": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"ALLOW", "DROP", "REJECT", "JUMP_TO_APPLICATION"}, false),
			Computed:     true,
		},
		"ip_protocol": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"IPV4", "IPV6", "IPV4_IPV6"}, false),
			Computed:     true,
		},
		"direction": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"IN", "OUT", "IN_OUT"}, false),
			Default:      "IN_OUT",
		},
		"sequence_number": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"profiles": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 128,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"services": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 128,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"service_entries": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 128,
			Elem:     resourceServiceEntryCustomSchema(),
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 30,
			Elem:     resourceTagSchema(),
		},
		"scope": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 128,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"destination_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 128,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"source_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 128,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceNsxtVpcGatewayPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcGatewayPolicyRuleCreate,
		Read:   resourceNsxtVpcGatewayPolicyRuleRead,
		Update: resourceNsxtVpcGatewayPolicyRuleUpdate,
		Delete: resourceNsxtVpcGatewayPolicyRuleDelete,
		Schema: resourceGatewayPolicyRuleSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceGatewayPolicyRuleImporter,
		},
	}
}

func resourceGatewayPolicyRuleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceGatewayPolicyRuleSchema()
	return ResourceImporter(d, m, "GatewayPolicyRule", s, d.Id())
}

func resourceNsxtVpcGatewayPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceGatewayPolicyRuleSchema()
	err := APIRead(d, meta, "GatewayPolicyRule", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object GatewayPolicyRule %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtVpcGatewayPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGatewayPolicyRuleSchema()
	err := APICreateOrUpdate(d, meta, "GatewayPolicyRule", s)
	if err == nil {
		err = resourceNsxtVpcGatewayPolicyRuleRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGatewayPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGatewayPolicyRuleSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "GatewayPolicyRule", s)
	if err == nil {
		err = resourceNsxtVpcGatewayPolicyRuleRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGatewayPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if 'object not found' or 'forbidden' or 'success with no response' response occurs, terraform should swallow it and not fail apply on object, else throw error and fail
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Error occurred in Delete for resource GatewayPolicyRule \n")
			return err
		}
		d.SetId("")
	}
	return nil
}
