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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
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
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "display_name", "test-group-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "description", "Group description"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*", map[string]string{
						"resource_type": "NestedExpression"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"member_type": "VirtualMachine"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"value": "vm_1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"key": "Name"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"operator": "CONTAINS"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"resource_type": "Condition"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
						"scope": "scope1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
						"tag": "webvm"}),
				),
			},
			{
				Config: testAccNSXTGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGroupExists("nsxt_vpc_group.testGroup"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "display_name", "test-group-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "description", "updated Group description"),
				),
			},
			{
				Config: testAccNSXTGroupUpdate1Config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGroupExists("nsxt_vpc_group.testGroup"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "display_name", "test-group-abc-updated-1"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*", map[string]string{
						"resource_type": "NestedExpression"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"member_type": "VirtualMachine"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"value": "vm_1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"key": "Name"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"operator": "CONTAINS"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
						"resource_type": "Condition"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
						"scope": "scope1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
						"tag": "webvm"}),
				),
			},
			{
				Config: testAccNSXTGroupUpdate2Config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGroupExists("nsxt_vpc_group.testGroup"),
					resource.TestCheckResourceAttr("nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_group.testGroup", "expression.*", map[string]string{
						"resource_type": "IPAddressExpression"}),
					resource.TestCheckTypeSetElemAttr("nsxt_vpc_group.testGroup", "expression.*.ip_addresses.*", "10.112.10.1"),
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

//nolint

const testAccNSXTGroupConfig = `
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

const testAccNSXTGroupupdatedConfig = `
    resource "nsxt_vpc_group" "testGroup" {
      	nsx_id = "test-group-abc"
	display_name = "test-group-abc-updated"
	description = "updated Group description"
}
`

const testAccNSXTGroupUpdate1Config = `
  resource "nsxt_vpc_group" "testGroup" {
  	nsx_id = "test-group-abc"
	display_name = "test-group-abc-updated-1"
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
const testAccNSXTGroupUpdate2Config = `
  resource "nsxt_vpc_group" "testGroup" {
  	nsx_id = "test-group-abc"
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
