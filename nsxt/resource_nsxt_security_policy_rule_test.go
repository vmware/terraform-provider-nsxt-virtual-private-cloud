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

func TestNSXTSecurityPolicyRuleBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTSecurityPolicyRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTSecurityPolicyRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTSecurityPolicyRuleExists("nsxt_vpc_security_policy_rule.testSecurityPolicyRule"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "nsx_id", "test-rule-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "display_name", "test-rule-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "description", "Rule description"),
					resource.TestCheckTypeSetElemAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "services.*", "ANY"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "action", "ALLOW"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "sequence_number", "0"),
				),
			},
			{
				Config: testAccNSXTSecurityPolicyRuleupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTSecurityPolicyRuleExists("nsxt_vpc_security_policy_rule.testSecurityPolicyRule"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "nsx_id", "test-rule-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "display_name", "test-rule-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "description", "updated Rule description"),
					resource.TestCheckTypeSetElemAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "services.*", "ANY"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "action", "ALLOW"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy_rule.testSecurityPolicyRule", "sequence_number", "0"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_security_policy_rule.testSecurityPolicyRule",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTSecurityPolicyRuleConfig,
			},
		},
	})
}

func testAccCheckNSXTSecurityPolicyRuleExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT SecurityPolicyRule policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTSecurityPolicyRuleDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_security_policy_rule" {
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
			return fmt.Errorf("NSXT SecurityPolicyRule still exists")
		}
	}
	return nil
}

//nolint

const testAccNSXTSecurityPolicyRuleConfig = `
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
`

const testAccNSXTSecurityPolicyRuleupdatedConfig = `
    resource "nsxt_vpc_security_policy_rule" "testSecurityPolicyRule" {
      	parent_path = nsxt_vpc_security_policy.testSecurityPolicy.path
	nsx_id = "test-rule-abc"
	display_name = "test-rule-abc-updated"
	description = "updated Rule description"
	source_groups = [nsxt_vpc_group.testGroup.path]
	destination_groups = [nsxt_vpc_group.testGroup.path]
	services = ["ANY"]
	action = "ALLOW"
	sequence_number = 0
}
    resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
      	nsx_id = "test-securitypolicy-abc"
	display_name = "test-secutitypolicy-abc-updated"
	description = "updated SecurityPolicy description"
	sequence_number = 0
}
    resource "nsxt_vpc_group" "testGroup" {
      	nsx_id = "test-group-abc"
	display_name = "test-group-abc-updated"
	description = "updated Group description"
}
`
