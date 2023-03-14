/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
package nsxt

import (
	"log"
	"strings"

	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"application_connectivity_strategy": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     resourceApplicationConnectivityStrategySchema(),
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
		"connectivity_preference": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"ALLOWLIST", "DENYLIST", "ALLOWLIST_ENABLE_LOGGING", "DENYLIST_ENABLE_LOGGING", "NONE"}, false),
			Computed:     true,
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
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object SecurityPolicy %v\n", err)
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
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource SecurityPolicy not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint