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

func TestNSXTDataSourceSubnetPortBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSSubnetPortConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet_port.testSubnetPort", "nsx_id", "test-port-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet_port.testSubnetPort", "display_name", "test-segmentport-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet_port.testSubnetPort", "description", "SegmentPort description"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_subnet_port.testSubnetPort", "address_bindings.*", map[string]string{
							"ip_address": "1.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_subnet_port.testSubnetPort", "address_bindings.*", map[string]string{
							"mac_address": "aa:bb:cc:dd:ee:ff"}),
				),
			},
		},
	})
}

const testAccNSXTDSSubnetPortConfig = `

    resource "nsxt_vpc_subnet_port" "testSubnetPort" {
      	parent_path = nsxt_vpc_subnet.testVpcSubnet.path
	nsx_id = "test-port-abc"
	display_name = "test-segmentport-abc"
	description = "SegmentPort description"
	address_bindings {
	ip_address = "1.1.1.1"
	mac_address = "aa:bb:cc:dd:ee:ff"
}
address_bindings {
	ip_address = "1.1.1.2"
	mac_address = "aa:bb:cc:dd:ee:f1"
}
}
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc"
	description = "VpcSubnet description"
	ipv4_subnet_size = 16
	access_mode = "Public"
}

data "nsxt_vpc_subnet_port" "testSubnetPort" {
  display_name = nsxt_vpc_subnet_port.testSubnetPort.display_name
}
`
