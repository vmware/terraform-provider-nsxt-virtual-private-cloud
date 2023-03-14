/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
package nsxt

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
)

func TestNSXTPolicyNatRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTPolicyNatRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTPolicyNatRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTPolicyNatRuleExists("nsxt_vpc_policy_nat_rule.testPolicyNatRule"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "nsx_id", "test-natrule-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "display_name", "test-natrule-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "description", "NatRule description"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "action", "DNAT"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "destination_network", "10.117.5.19"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "service", "/infra/services/AD_Server"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_network", "192.168.1.1"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_ports", "80-82"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "sequence_number", "10"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "enabled", "true"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "logging", "false"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "firewall_match", "MATCH_EXTERNAL_ADDRESS"),
				),
			},
			{
				Config: testAccNSXTPolicyNatRuleupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTPolicyNatRuleExists("nsxt_vpc_policy_nat_rule.testPolicyNatRule"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "nsx_id", "test-natrule-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "display_name", "test-natrule-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "description", "NatRule description"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "action", "DNAT"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "destination_network", "10.117.5.19"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "service", "/infra/services/AD_Server"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_network", "192.168.1.1"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "translated_ports", "80-82"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "sequence_number", "10"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "enabled", "true"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "logging", "false"),
					resource.TestCheckResourceAttr("nsxt_vpc_policy_nat_rule.testPolicyNatRule", "firewall_match", "MATCH_EXTERNAL_ADDRESS"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_policy_nat_rule.testPolicyNatRule",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTPolicyNatRuleConfig,
			},
		},
	})
}

func testAccCheckNSXTPolicyNatRuleExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT PolicyNatRule policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTPolicyNatRuleDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_policy_nat_rule" {
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
			return fmt.Errorf("NSXT PolicyNatRule still exists")
		}
	}
	return nil
}

//nolint

const testAccNSXTPolicyNatRuleConfig = `
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
`

const testAccNSXTPolicyNatRuleupdatedConfig = `
    resource "nsxt_vpc_policy_nat_rule" "testPolicyNatRule" {
      	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/nat/USER"
	nsx_id = "test-natrule-abc"
	display_name = "test-natrule-abc-updated"
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
`
