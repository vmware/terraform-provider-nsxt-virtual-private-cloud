/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package nsxt

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func getDataSourceCommonSchema(additionalSchemaItems map[string]*schema.Schema) map[string]*schema.Schema {
	baseSchema := map[string]*schema.Schema{
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
			Computed: true,
		},
		"path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	for key, value := range additionalSchemaItems {
		baseSchema[key] = value
	}

	return baseSchema
}
