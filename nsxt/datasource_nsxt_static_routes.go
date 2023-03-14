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

func dataSourceNsxtVpcStaticRoutes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtVpcStaticRoutesRead,
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

func dataSourceNsxtVpcStaticRoutesRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtVpcStaticRoutes()
	err := DatasourceRead(d, meta, "StaticRoutes", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}