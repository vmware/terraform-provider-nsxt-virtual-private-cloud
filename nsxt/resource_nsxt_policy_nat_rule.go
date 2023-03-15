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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePolicyNatRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"action": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"destination_network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"firewall_match": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"MATCH_EXTERNAL_ADDRESS", "MATCH_INTERNAL_ADDRESS", "BYPASS"}, false),
			Default:      "MATCH_INTERNAL_ADDRESS",
		},
		"logging": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"policy_based_vpn_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"BYPASS", "MATCH"}, false),
			Computed:     true,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"scope": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"sequence_number": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"service": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"source_network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     resourceTagSchema(),
		},
		"translated_network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"translated_ports": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceNsxtVpcPolicyNatRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcPolicyNatRuleCreate,
		Read:   resourceNsxtVpcPolicyNatRuleRead,
		Update: resourceNsxtVpcPolicyNatRuleUpdate,
		Delete: resourceNsxtVpcPolicyNatRuleDelete,
		Schema: resourcePolicyNatRuleSchema(),
		Importer: &schema.ResourceImporter{
			State: resourcePolicyNatRuleImporter,
		},
	}
}

func resourcePolicyNatRuleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourcePolicyNatRuleSchema()
	return ResourceImporter(d, m, "PolicyNatRule", s, d.Id())
}

func resourceNsxtVpcPolicyNatRuleRead(d *schema.ResourceData, meta interface{}) error {
	s := resourcePolicyNatRuleSchema()
	err := APIRead(d, meta, "PolicyNatRule", s)
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object PolicyNatRule %v\n", err)
	}
	return err
}

func resourceNsxtVpcPolicyNatRuleCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourcePolicyNatRuleSchema()
	err := APICreateOrUpdate(d, meta, "PolicyNatRule", s)
	if err == nil {
		err = resourceNsxtVpcPolicyNatRuleRead(d, meta)
	}
	return err
}

func resourceNsxtVpcPolicyNatRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourcePolicyNatRuleSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "PolicyNatRule", s)
	if err == nil {
		err = resourceNsxtVpcPolicyNatRuleRead(d, meta)
	}
	return err
}

func resourceNsxtVpcPolicyNatRuleDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource PolicyNatRule not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint
