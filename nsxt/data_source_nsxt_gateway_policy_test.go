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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestNSXTDataSourceGatewayPolicyBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSGatewayPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "nsx_id", "test-gatewaypolicy-abc-1"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "display_name", "test-gatewaypolicy-abc-1"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "description", "GatewayPolicy 1 description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "sequence_number", "0"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "tags.*", map[string]string{
							"scope": "scope1"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "tags.*", map[string]string{
							"tag": "webvm1"}),
				),
			},
		},
	})
}

const testAccNSXTDSGatewayPolicyConfig = `

    resource "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
      	nsx_id = "test-gatewaypolicy-abc-1"
	display_name = "test-gatewaypolicy-abc-1"
	description = "GatewayPolicy 1 description"
	sequence_number = 0
	tags {
	scope = "scope1"
	tag = "webvm1"
}
tags {
	scope = "scope2"
	tag = "webvm2"
}
}

data "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
  display_name = nsxt_vpc_gateway_policy.testGatewayPolicy.display_name
  context_info {
    context = "vpc"
  }
}
`
