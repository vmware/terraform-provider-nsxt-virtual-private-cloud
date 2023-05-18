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

func TestNSXTDataSourceSecurityPolicyRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSSecurityPolicyRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "display_name", "test-rule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "description", "Rule description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "nsx_id", "test-rule-abc"),
					resource.TestCheckTypeSetElemAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "services.*", "ANY"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "action", "ALLOW"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "sequence_number", "0"),
				),
			},
		},
	})
}

const testAccNSXTDSSecurityPolicyRuleConfig = `

    resource "nsxt_vpc_security_policy_rule" "testSecurityPolicyRule" {
      	display_name = "test-rule-abc"
	description = "Rule description"
	parent_path = nsxt_vpc_security_policy.testSecurityPolicy.path
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
	nsx_id = "test-group-abc-1"
	display_name = "test-group-abc-1"
	description = "Group 1 description"
}
    resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
      	nsx_id = "test-securitypolicy-abc-1"
	display_name = "test-securitypolicy-abc-1"
	description = "SecurityPolicy 1 description"
	sequence_number = 0
}

data "nsxt_vpc_security_policy_rule" "testSecurityPolicyRule" {
  display_name = nsxt_vpc_security_policy_rule.testSecurityPolicyRule.display_name
}
`
