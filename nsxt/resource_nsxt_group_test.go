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
	nsxtclient "github.com/vmware/terraform-provider-nsxt-virtual-private-cloud/nsxt/clients"
)

func TestNSXTGroupBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGroupExists("nsxt_vpc_group.testGroup"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"key": "Name"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"operator": "CONTAINS"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"resource_type": "Condition"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"value": "vm_1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"member_type": "VirtualMachine"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*", map[string]string{
						"resource_type": "NestedExpression"}),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc-2"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "display_name", "test-group-abc-2"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "description", "Group 2 description"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "tags.*", map[string]string{
						"scope": "scope1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "tags.*", map[string]string{
						"tag": "webvm"}),
				),
			},
			{
				Config: testAccNSXTGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGroupExists("nsxt_vpc_group.testGroup"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc-2"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "display_name", "test-group-abc-2-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "description", "updated Group 2 description"),
				),
			},
			{
				Config: testAccNSXTGroupUpdate1Config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGroupExists("nsxt_vpc_group.testGroup"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"key": "Name"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"operator": "CONTAINS"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"resource_type": "Condition"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"value": "vm_1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"member_type": "VirtualMachine"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*", map[string]string{
						"resource_type": "NestedExpression"}),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc-2"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "display_name", "test-group-abc-updated-2"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "tags.*", map[string]string{
						"scope": "scope1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "tags.*", map[string]string{
						"tag": "webvm"}),
				),
			},
			{
				Config: testAccNSXTGroupUpdate2Config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGroupExists("nsxt_vpc_group.testGroup"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc-2"),
					resource.TestCheckTypeSetElemAttr("nsxt_vpc_group.testGroup", "expression.*.ip_addresses.*", "10.112.10.1"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*", map[string]string{
						"resource_type": "IPAddressExpression"}),
				),
			},
			{
				ResourceName:      "nsxt_vpc_group.testGroup",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTGroupConfig,
			},
		},
	})
}

func testAccCheckNSXTGroupExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT Group policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTGroupDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_group" {
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
			return fmt.Errorf("NSXT Group still exists")
		}
	}
	return nil
}

const testAccNSXTGroupConfig = `
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
	nsx_id = "test-group-abc-2"
	display_name = "test-group-abc-2"
	description = "Group 2 description"
	tags {
	scope = "scope1"
	tag = "webvm"
}
}
`

const testAccNSXTGroupupdatedConfig = `
    resource "nsxt_vpc_group" "testGroup" {
      	nsx_id = "test-group-abc-2"
	display_name = "test-group-abc-2-updated"
	description = "updated Group 2 description"
}
`

const testAccNSXTGroupUpdate1Config = `
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
	nsx_id = "test-group-abc-2"
	display_name = "test-group-abc-updated-2"
	tags {
	scope = "scope1"
	tag = "webvm"
}
}
`
const testAccNSXTGroupUpdate2Config = `
  resource "nsxt_vpc_group" "testGroup" {
  	nsx_id = "test-group-abc-2"
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
}
`
