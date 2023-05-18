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

func TestNSXTDataSourceGroupBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
							"key": "Name"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
							"operator": "CONTAINS"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
							"resource_type": "Condition"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
							"value": "vm_1"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
							"member_type": "VirtualMachine"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*", map[string]string{
							"resource_type": "NestedExpression"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
							"scope": "scope1"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
							"tag": "webvm"}),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc-2"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_group.testGroup", "display_name", "test-group-abc-2"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_group.testGroup", "description", "Group 2 description"),
				),
			},
		},
	})
}

const testAccNSXTDSGroupConfig = `

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
	nsx_id = "test-group-abc-2"
	display_name = "test-group-abc-2"
	description = "Group 2 description"
}

data "nsxt_vpc_group" "testGroup" {
  display_name = nsxt_vpc_group.testGroup.display_name
}
`
