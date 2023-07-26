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

func TestNSXTDataSourcePolicyVpcNatRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSPolicyVpcNatRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "translated_network", "192.168.4.0"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "display_name", "test-natrule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "description", "NatRule description"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "nsx_id", "test-natrule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "action", "SNAT"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "logging", "false"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "firewall_match", "MATCH_EXTERNAL_ADDRESS"),
					resource.TestCheckResourceAttr(
						"nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "sequence_number", "10"),
				),
			},
		},
	})
}

const testAccNSXTDSPolicyVpcNatRuleConfig = `

    resource "nsxt_policy_vpc_nat_rule" "testPolicyVpcNatRule" {
      	translated_network = "192.168.4.0"
	display_name = "test-natrule-abc"
	description = "NatRule description"
	enabled = true
	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/nat/USER"
	nsx_id = "test-natrule-abc"
	action = "SNAT"
	logging = false
	firewall_match = "MATCH_EXTERNAL_ADDRESS"
	sequence_number = 10
}

data "nsxt_policy_vpc_nat_rule" "testPolicyVpcNatRule" {
  display_name = nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule.display_name
  context_info {
    context = "vpc"
  }
}
`
