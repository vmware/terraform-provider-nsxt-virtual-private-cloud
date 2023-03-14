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

func resourceSubnetPortSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"address_bindings": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     resourcePortAddressBindingEntrySchema(),
		},
		"admin_state": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"UP", "DOWN"}, false),
			Default:      "UP",
		},
		"attachment": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     resourcePortAttachmentSchema(),
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
		"extra_configs": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     resourceSegmentExtraConfigSchema(),
		},
		"ignored_address_bindings": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     resourcePortAddressBindingEntrySchema(),
		},
		"init_state": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"UNBLOCKED_VLAN", "RESTORE_VIF"}, false),
			Computed:     true,
		},
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
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

func resourceNsxtVpcSubnetPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcSubnetPortCreate,
		Read:   resourceNsxtVpcSubnetPortRead,
		Update: resourceNsxtVpcSubnetPortUpdate,
		Delete: resourceNsxtVpcSubnetPortDelete,
		Schema: resourceSubnetPortSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceSubnetPortImporter,
		},
	}
}

func resourceSubnetPortImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceSubnetPortSchema()
	return ResourceImporter(d, m, "SubnetPort", s, d.Id())
}

func resourceNsxtVpcSubnetPortRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceSubnetPortSchema()
	err := APIRead(d, meta, "SubnetPort", s)
	if err != nil {
		log.Printf("[ERROR] Error occurred in reading object SubnetPort %v\n", err)
	}
	return err
}

func resourceNsxtVpcSubnetPortCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSubnetPortSchema()
	err := APICreateOrUpdate(d, meta, "SubnetPort", s)
	if err == nil {
		err = resourceNsxtVpcSubnetPortRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetPortUpdate(d *schema.ResourceData, meta interface{}) error {
	s := resourceSubnetPortSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "SubnetPort", s)
	if err == nil {
		err = resourceNsxtVpcSubnetPortRead(d, meta)
	}
	return err
}

func resourceNsxtVpcSubnetPortDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Resource SubnetPort not found\n")
			return err
		}
		d.SetId("")
	}
	return nil
}

//nolint