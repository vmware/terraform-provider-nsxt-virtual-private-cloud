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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "nsx_id", "test-rule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "display_name", "test-rule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "description", "Rule description"),
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
      	parent_path = nsxt_vpc_security_policy.testSecurityPolicy.path
	nsx_id = "test-rule-abc"
	display_name = "test-rule-abc"
	description = "Rule description"
	source_groups = [nsxt_vpc_group.testGroup.path]
	destination_groups = [nsxt_vpc_group.testGroup.path]
	services = ["ANY"]
	action = "ALLOW"
	sequence_number = 0
}
    resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
      	nsx_id = "test-securitypolicy-abc"
	display_name = "test-securitypolicy-abc"
	description = "SecurityPolicy description"
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

data "nsxt_vpc_security_policy_rule" "testSecurityPolicyRule" {
  display_name = nsxt_vpc_security_policy_rule.testSecurityPolicyRule.display_name
}
`
