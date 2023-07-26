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

func dataSourceNsxtVpcIpAddressPool() *schema.Resource {
	additionalSchemaItems := map[string]*schema.Schema{} // Define any additional schema items specific to the data source

	return &schema.Resource{
		Read:   dataSourceNsxtVpcIpAddressPoolRead,
		Schema: getDataSourceCommonSchema(additionalSchemaItems),
	}
}

func dataSourceNsxtVpcIpAddressPoolRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtVpcIpAddressPool()
	err := DatasourceRead(d, meta, "IpAddressPool", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}
