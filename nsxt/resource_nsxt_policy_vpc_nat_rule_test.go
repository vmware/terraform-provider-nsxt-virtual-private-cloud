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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
)

func TestNSXTPolicyVpcNatRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTPolicyVpcNatRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTPolicyVpcNatRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTPolicyVpcNatRuleExists("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "translated_network", "192.168.4.0"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "display_name", "test-natrule-abc"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "description", "NatRule description"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "enabled", "true"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "nsx_id", "test-natrule-abc"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "action", "SNAT"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "logging", "false"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "firewall_match", "MATCH_EXTERNAL_ADDRESS"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "sequence_number", "10"),
				),
			},
			{
				Config: testAccNSXTPolicyVpcNatRuleupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTPolicyVpcNatRuleExists("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "translated_network", "192.168.4.0"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "display_name", "test-natrule-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "description", "NatRule description"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "enabled", "true"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "nsx_id", "test-natrule-abc"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "action", "SNAT"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "logging", "false"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "firewall_match", "MATCH_EXTERNAL_ADDRESS"),
					resource.TestCheckResourceAttr("nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule", "sequence_number", "10"),
				),
			},
			{
				ResourceName:      "nsxt_policy_vpc_nat_rule.testPolicyVpcNatRule",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTPolicyVpcNatRuleConfig,
			},
		},
	})
}

func testAccCheckNSXTPolicyVpcNatRuleExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT PolicyVpcNatRule policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTPolicyVpcNatRuleDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_policy_vpc_nat_rule" {
			continue
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "400") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("NSXT PolicyVpcNatRule still exists")
		}
	}
	return nil
}

const testAccNSXTPolicyVpcNatRuleConfig = `
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
`

const testAccNSXTPolicyVpcNatRuleupdatedConfig = `
  resource "nsxt_policy_vpc_nat_rule" "testPolicyVpcNatRule" {
  translated_network = "192.168.4.0"
	display_name = "test-natrule-abc-updated"
	description = "NatRule description"
	enabled = true
	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/nat/USER"
	nsx_id = "test-natrule-abc"
	action = "SNAT"
	logging = false
	firewall_match = "MATCH_EXTERNAL_ADDRESS"
	sequence_number = 10
}
`
