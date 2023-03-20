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

func TestNSXTDataSourcePolicyNatRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSPolicyNatRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "nsx_id", "test-natrule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "display_name", "test-natrule-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "description", "NatRule description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "action", "DNAT"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "destination_network", "10.117.5.19"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "service", "/infra/services/AD_Server"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_network", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_ports", "80-82"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "sequence_number", "10"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "logging", "false"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_policy_nat_rule.testPolicyNatRule", "firewall_match", "MATCH_EXTERNAL_ADDRESS"),
				),
			},
		},
	})
}

const testAccNSXTDSPolicyNatRuleConfig = `

    resource "nsxt_vpc_policy_nat_rule" "testPolicyNatRule" {
      	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/nat/USER"
	nsx_id = "test-natrule-abc"
	display_name = "test-natrule-abc"
	description = "NatRule description"
	action = "DNAT"
	destination_network = "10.117.5.19"
	service = "/infra/services/AD_Server"
	translated_network = "192.168.1.1"
	translated_ports = "80-82"
	sequence_number = 10
	enabled = true
	logging = false
	firewall_match = "MATCH_EXTERNAL_ADDRESS"
}

data "nsxt_vpc_policy_nat_rule" "testPolicyNatRule" {
  display_name = nsxt_vpc_policy_nat_rule.testPolicyNatRule.display_name
}
`
