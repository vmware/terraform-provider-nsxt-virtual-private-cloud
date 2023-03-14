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

func dataSourceNsxtVpcGatewayPolicyRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtVpcGatewayPolicyRuleRead,
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
				Computed: true,
			},
		},
	}
}

func dataSourceNsxtVpcGatewayPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtVpcGatewayPolicyRule()
	err := DatasourceRead(d, meta, "GatewayPolicyRule", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}
