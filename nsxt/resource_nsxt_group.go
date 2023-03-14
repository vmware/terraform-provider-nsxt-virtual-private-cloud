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

func resourceGroupSchema() map[string]*schema.Schema {
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
		"expression": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     resourceExpressionCustomSchema(),
		},
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"IN_PROGRESS", "SUCCESS", "FAILURE"}, false),
			Computed:     true,
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

func resourceNsxtVpcGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcGroupCreate,
		Read:   resourceNsxtVpcGroupRead,
		Update: resourceNsxtVpcGroupUpdate,
		Delete: resourceNsxtVpcGroupDelete,
		Schema: resourceGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceGroupImporter,
		},
	}
}

func resourceGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceGroupSchema()
	return ResourceImporter(d, m, "Group", s, d.Id())
}

func resourceNsxtVpcGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceGroupSchema()
	err := APIRead(d, meta, "Group", s)
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object Group %v\n", err)
	}
	return err
}

func resourceNsxtVpcGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGroupSchema()
	err := APICreateOrUpdate(d, meta, "Group", s)
	if err == nil {
		err = resourceNsxtVpcGroupRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "Group", s)
	if err == nil {
		err = resourceNsxtVpcGroupRead(d, meta)
	}
	return err
}

func resourceNsxtVpcGroupDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource Group not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint