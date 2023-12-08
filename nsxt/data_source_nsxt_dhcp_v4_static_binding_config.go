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
	additionalSchemaItems := map[string]*schema.Schema{
		"parent_path": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}

	return &schema.Resource{
		Read:   dataSourceNsxtVpcDhcpV4StaticBindingConfigRead,
		Schema: getDataSourceCommonSchema(additionalSchemaItems),
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
