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

func resourceSecurityPolicySchema() map[string]*schema.Schema {
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
		"connectivity_preference": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"ALLOWLIST", "DENYLIST", "ALLOWLIST_ENABLE_LOGGING", "DENYLIST_ENABLE_LOGGING", "NONE"}, false),
		},
		"logging_enabled": {
			Type:       schema.TypeBool,
			Optional:   true,
			Default:    false,
			Deprecated: "This is a deprecated property. Please refer the documentation for details, and refrain from use as this will be removed in future versions.",
		},
		"connectivity_strategy": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"WHITELIST", "BLACKLIST", "WHITELIST_ENABLE_LOGGING", "BLACKLIST_ENABLE_LOGGING", "NONE"}, false),
			Deprecated:   "This is a deprecated property. Please refer the documentation for details, and refrain from use as this will be removed in future versions.",
		},
		"sequence_number": {
			Type:     schema.TypeInt,
			Required: true,
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
		"application_connectivity_strategy": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 3,
			Elem:     resourceApplicationConnectivityStrategySchema(),
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

func resourceNsxtVpcSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcSecurityPolicyCreate,
		Read:   resourceNsxtVpcSecurityPolicyRead,
		Update: resourceNsxtVpcSecurityPolicyUpdate,
		Delete: resourceNsxtVpcSecurityPolicyDelete,
		Schema: resourceSecurityPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: resourceSecurityPolicyImporter,
		},
	}
}

func resourceSecurityPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceSecurityPolicySchema()
	return ResourceImporter(d, m, "SecurityPolicy", s, d.Id())
}

func resourceNsxtVpcSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceSecurityPolicySchema()
	err := APIRead(d, meta, "SecurityPolicy", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object SecurityPolicy %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtVpcSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSecurityPolicySchema()
	err := APICreateOrUpdate(d, meta, "SecurityPolicy", s)
	if err == nil {
		err = resourceNsxtVpcSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSecurityPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "SecurityPolicy", s)
	if err == nil {
		err = resourceNsxtVpcSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if 'object not found' or 'forbidden' or 'success with no response' response occurs, terraform should swallow it and not fail apply on object, else throw error and fail
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Error occurred in Delete for resource SecurityPolicy \n")
			return err
		}
		d.SetId("")
	}
	return nil
}
