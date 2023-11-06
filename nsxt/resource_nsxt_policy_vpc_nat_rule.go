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

func resourcePolicyVpcNatRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"logging": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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
		"destination_network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"translated_network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"source_network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"sequence_number": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"firewall_match": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"MATCH_EXTERNAL_ADDRESS", "MATCH_INTERNAL_ADDRESS", "BYPASS"}, false),
			Default:      "MATCH_INTERNAL_ADDRESS",
		},
		"action": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 30,
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

func resourceNsxtPolicyVpcNatRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtPolicyVpcNatRuleCreate,
		Read:   resourceNsxtPolicyVpcNatRuleRead,
		Update: resourceNsxtPolicyVpcNatRuleUpdate,
		Delete: resourceNsxtPolicyVpcNatRuleDelete,
		Schema: resourcePolicyVpcNatRuleSchema(),
		Importer: &schema.ResourceImporter{
			State: resourcePolicyVpcNatRuleImporter,
		},
	}
}

func resourcePolicyVpcNatRuleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourcePolicyVpcNatRuleSchema()
	return ResourceImporter(d, m, "PolicyVpcNatRule", s, d.Id())
}

func resourceNsxtPolicyVpcNatRuleRead(d *schema.ResourceData, meta interface{}) error {
	s := resourcePolicyVpcNatRuleSchema()
	err := APIRead(d, meta, "PolicyVpcNatRule", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object PolicyVpcNatRule %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtPolicyVpcNatRuleCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourcePolicyVpcNatRuleSchema()
	err := APICreateOrUpdate(d, meta, "PolicyVpcNatRule", s)
	if err == nil {
		err = resourceNsxtPolicyVpcNatRuleRead(d, meta)
	}
	return err
}

func resourceNsxtPolicyVpcNatRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourcePolicyVpcNatRuleSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "PolicyVpcNatRule", s)
	if err == nil {
		err = resourceNsxtPolicyVpcNatRuleRead(d, meta)
	}
	return err
}

func resourceNsxtPolicyVpcNatRuleDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if 'object not found' or 'forbidden' or 'success with no response' response occurs, terraform should swallow it and not fail apply on object, else throw error and fail
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Error occurred in Delete for resource PolicyVpcNatRule \n")
			return err
		}
		d.SetId("")
	}
	return nil
}
