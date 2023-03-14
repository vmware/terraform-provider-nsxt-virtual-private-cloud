/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceNsxtSharedInfraIpAddressBlock() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtSharedInfraIpAddressBlockRead,
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
		},
	}
}

func dataSourceNsxtSharedInfraIpAddressBlockRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtSharedInfraIpAddressBlock()
	err := DatasourceRead(d, meta, "InfraIpAddressBlock", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}
