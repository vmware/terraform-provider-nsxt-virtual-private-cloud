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

func TestNSXTDataSourcePolicyNatRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSPolicyNatRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_network", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "display_name", "test-natrule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "description", "NatRule description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "service", "/infra/services/AD_Server"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "enabled", "True"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "nsx_id", "test-natrule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_ports", "80-82"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "action", "DNAT"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "logging", "False"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "firewall_match", "MATCH_EXTERNAL_ADDRESS"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "destination_network", "10.117.5.19"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "sequence_number", "10"),
				),
			},
		},
	})
}

const testAccNSXTDSPolicyNatRuleConfig = `

    resource "nsxt_vpc_policy_nat_rule" "testPolicyNatRule" {
      	translated_network = "192.168.1.1"
	display_name = "test-natrule-abc"
	description = "NatRule description"
	service = "/infra/services/AD_Server"
	enabled = true
	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/nat/USER"
	nsx_id = "test-natrule-abc"
	translated_ports = "80-82"
	action = "DNAT"
	logging = false
	firewall_match = "MATCH_EXTERNAL_ADDRESS"
	destination_network = "10.117.5.19"
	sequence_number = 10
}

data "nsxt_vpc_policy_nat_rule" "testPolicyNatRule" {
  display_name = nsxt_vpc_policy_nat_rule.testPolicyNatRule.display_name
}
`
