/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExpressionsCustomSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exclude": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceExcludedMembersListSchema(),
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operator": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope_operator": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conjunction_operator": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"marked_for_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceExpressionCustomSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceTagSchema(),
				MaxItems: 30,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"expressions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceExpressionsCustomSchema(),
				MinItems: 1,
			},
			"conjunction_operator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				MinItems: 1,
				MaxItems: 4000,
			},
			"paths": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"external_ids": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				MinItems: 1,
			},
			"exclude": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceExcludedMembersListSchema(),
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"member_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope_operator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"marked_for_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceServiceEntryCustomSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsx_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// all properties from each resource_type enum value schemas included
			"protocol_number": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmp_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmp_type": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"source_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"l4_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ether_type": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nested_service_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
