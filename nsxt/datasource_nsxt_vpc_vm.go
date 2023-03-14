/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
package nsxt

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNsxtVpcVM() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceNsxtPolicyVpcVMRead,
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func DataSourceNsxtPolicyVpcVMRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtVpcVM()
	err := DatasourceReadForVM(d, meta, "VirtualMachine", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}
