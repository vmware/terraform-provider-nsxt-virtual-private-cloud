/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
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
					resource.TestCheckResourceAttr(
						"nsxt_vpc_group.testGroup", "nsx_id", "test-group-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_group.testGroup", "display_name", "test-group-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_group.testGroup", "description", "Group description"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*", map[string]string{
							"resource_type": "NestedExpression"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
							"member_type": "VirtualMachine"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.expressions.*", map[string]string{
							"value": "vm_1"}),
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
						"nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
							"scope": "scope1"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_group.testGroup", "expression.*.tags.*", map[string]string{
							"tag": "webvm"}),
				),
			},
		},
	})
}

const testAccNSXTDSGroupConfig = `

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

data "nsxt_vpc_group" "testGroup" {
  display_name = nsxt_vpc_group.testGroup.display_name
}
`
