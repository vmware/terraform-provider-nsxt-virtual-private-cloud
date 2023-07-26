/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func getDataSourceCommonSchema(additionalSchemaItems map[string]*schema.Schema) map[string]*schema.Schema {
	baseSchema := map[string]*schema.Schema{
		"context": {
			Type:        schema.TypeList,
			Description: "Datasource context information",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"scope": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"vpc", "project", "infra"}, false),
						Description:  "The scope in which the object exists or spans out of. It can be any one of vpc, project or infra.",
					},
					"domain": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "default",
						Description: "The domain ID of the object. Applicable for groups.",
					},
				},
			},
		},
		"nsx_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the object on NSX.",
		},
		"display_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The display name of the object.",
		},
		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The description of the object.",
		},
		"path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The complete policy path of the object.",
		},
	}

	for key, value := range additionalSchemaItems {
		baseSchema[key] = value
	}

	return baseSchema
}
