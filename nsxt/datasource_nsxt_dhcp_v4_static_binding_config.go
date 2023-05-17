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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceNsxtVpcDhcpV4StaticBindingConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtVpcDhcpV4StaticBindingConfigRead,
		Schema: map[string]*schema.Schema{
			"nsx_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNsxtVpcDhcpV4StaticBindingConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtVpcDhcpV4StaticBindingConfig()
	camelcasedName := "DhcpV4StaticBindingConfig"
	err := DatasourceRead(d, meta, camelcasedName, s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}
