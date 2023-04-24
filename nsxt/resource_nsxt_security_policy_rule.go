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
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityPolicyRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tag": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
			Computed: true,
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
			Computed: true,
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

func resourceNsxtVpcSecurityPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcSecurityPolicyRuleCreate,
		Read:   resourceNsxtVpcSecurityPolicyRuleRead,
		Update: resourceNsxtVpcSecurityPolicyRuleUpdate,
		Delete: resourceNsxtVpcSecurityPolicyRuleDelete,
		Schema: resourceSecurityPolicyRuleSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceSecurityPolicyRuleImporter,
		},
	}
}

func resourceSecurityPolicyRuleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceSecurityPolicyRuleSchema()
	return ResourceImporter(d, m, "SecurityPolicyRule", s, d.Id())
}

func resourceNsxtVpcSecurityPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceSecurityPolicyRuleSchema()
	err := APIRead(d, meta, "SecurityPolicyRule", s)
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object SecurityPolicyRule %v\n", err)
	}
	return err
}

func resourceNsxtVpcSecurityPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSecurityPolicyRuleSchema()
	err := APICreateOrUpdate(d, meta, "SecurityPolicyRule", s)
	if err == nil {
		err = resourceNsxtVpcSecurityPolicyRuleRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSecurityPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSecurityPolicyRuleSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "SecurityPolicyRule", s)
	if err == nil {
		err = resourceNsxtVpcSecurityPolicyRuleRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSecurityPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource SecurityPolicyRule not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}
