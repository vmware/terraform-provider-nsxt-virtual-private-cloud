/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// Auto generated code. DO NOT EDIT.

//nolint
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
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "nsx_id", "test-gatewaypolicy-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "display_name", "test-gatewaypolicy-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "description", "GatewayPolicy description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy.testGatewayPolicy", "sequence_number", "0"),
				),
			},
		},
	})
}

const testAccNSXTDSGatewayPolicyConfig = `

    resource "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
      	nsx_id = "test-gatewaypolicy-abc"
	display_name = "test-gatewaypolicy-abc"
	description = "GatewayPolicy description"
	sequence_number = 0
}

data "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
  display_name = nsxt_vpc_gateway_policy.testGatewayPolicy.display_name
}
`
