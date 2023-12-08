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

func dataSourceNsxtVpcService() *schema.Resource {
	additionalSchemaItems := map[string]*schema.Schema{} // Define any additional schema items specific to the data source

	return &schema.Resource{
		Read:   dataSourceNsxtVpcServiceRead,
		Schema: getDataSourceCommonSchema(additionalSchemaItems),
	}
}

func dataSourceNsxtVpcServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtVpcService()
	err := DatasourceRead(d, meta, "Service", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}
