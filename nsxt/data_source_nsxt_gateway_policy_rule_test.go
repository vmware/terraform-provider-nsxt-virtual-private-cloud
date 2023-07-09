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

func TestNSXTDataSourceGatewayPolicyRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSGatewayPolicyRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "display_name", "test-rule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "description", "Rule description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "nsx_id", "test-rule-abc"),
					resource.TestCheckTypeSetElemAttr("nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "services.*", "ANY"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "action", "ALLOW"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "sequence_number", "0"),
				),
			},
		},
	})
}

const testAccNSXTDSGatewayPolicyRuleConfig = `

    resource "nsxt_vpc_gateway_policy_rule" "testGatewayPolicyRule" {
      	display_name = "test-rule-abc"
	description = "Rule description"
	parent_path = nsxt_vpc_gateway_policy.testGatewayPolicy.path
	nsx_id = "test-rule-abc"
	destination_groups = [nsxt_vpc_group.testGroup.path]
	services = ["ANY"]
	action = "ALLOW"
	sequence_number = 0
	source_groups = [nsxt_vpc_group.testGroup.path]
}
    resource "nsxt_vpc_group" "testGroup" {
      	expression {
	expressions {
	key = "Name"
	operator = "CONTAINS"
	resource_type = "Condition"
	value = "vm_1"
	member_type = "VirtualMachine"
}
expressions {
	conjunction_operator = "AND"
	resource_type = "ConjunctionOperator"
}
expressions {
	key = "Tag"
	operator = "EQUALS"
	resource_type = "Condition"
	value = "London"
	member_type = "VirtualMachine"
}
	resource_type = "NestedExpression"
	tags {
	scope = "scope1"
	tag = "webvm"
}
}
expression {
	conjunction_operator = "OR"
	resource_type = "ConjunctionOperator"
}
expression {
	ip_addresses = ["10.112.10.1"]
	resource_type = "IPAddressExpression"
}
expression {
	conjunction_operator = "OR"
	resource_type = "ConjunctionOperator"
}
expression {
	paths = ["/orgs/default/projects/Dev_project/vpcs/dev_vpc/groups/default"]
	resource_type = "PathExpression"
}
	nsx_id = "test-group-abc-3"
	display_name = "test-group-abc-3"
	description = "Group 3 description"
}
    resource "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
      	nsx_id = "test-gatewaypolicy-abc-2"
	display_name = "test-gatewaypolicy-abc-2"
	description = "GatewayPolicy 2 description"
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

data "nsxt_vpc_gateway_policy_rule" "testGatewayPolicyRule" {
  parent_path = nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule.parent_path
}
`
