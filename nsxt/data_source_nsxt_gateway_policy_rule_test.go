/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
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
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "nsx_id", "test-rule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "display_name", "test-rule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule", "description", "Rule description"),
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
      	parent_path = nsxt_vpc_gateway_policy.testGatewayPolicy.path
	nsx_id = "test-rule-abc"
	display_name = "test-rule-abc"
	description = "Rule description"
	source_groups = [nsxt_vpc_group.testGroup.path]
	destination_groups = [nsxt_vpc_group.testGroup.path]
	services = ["ANY"]
	action = "ALLOW"
	sequence_number = 0
}
    resource "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
      	nsx_id = "test-gatewaypolicy-abc"
	display_name = "test-gatewaypolicy-abc"
	description = "GatewayPolicy description"
	sequence_number = 0
}
    resource "nsxt_vpc_group" "testGroup" {
      	nsx_id = "test-group-abc"
	display_name = "test-group-abc"
	description = "Group description"
	expression {
	resource_type = "NestedExpression"
	expressions {
	member_type = "VirtualMachine"
	value = "vm_1"
	key = "Name"
	operator = "CONTAINS"
	resource_type = "Condition"
}
expressions {
	resource_type = "ConjunctionOperator"
	conjunction_operator = "AND"
}
expressions {
	member_type = "VirtualMachine"
	value = "London"
	key = "Tag"
	operator = "EQUALS"
	resource_type = "Condition"
}
	tags {
	scope = "scope1"
	tag = "webvm"
}
}
expression {
	resource_type = "ConjunctionOperator"
	conjunction_operator = "OR"
}
expression {
	resource_type = "IPAddressExpression"
	ip_addresses = ["10.112.10.1"]
}
expression {
	resource_type = "ConjunctionOperator"
	conjunction_operator = "OR"
}
expression {
	resource_type = "PathExpression"
	paths = ["/orgs/default/projects/Dev_project/vpcs/dev_vpc/groups/default"]
}
}

data "nsxt_vpc_gateway_policy_rule" "testGatewayPolicyRule" {
  display_name = nsxt_vpc_gateway_policy_rule.testGatewayPolicyRule.display_name
}
`
