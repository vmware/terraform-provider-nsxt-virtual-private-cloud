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

func TestNSXTDataSourceVpcSubnetBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSVpcSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet.testVpcSubnet", "ipv4_subnet_size", "16"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet.testVpcSubnet", "nsx_id", "test-vpcsubnet-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet.testVpcSubnet", "display_name", "test-vpcsubnet-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet.testVpcSubnet", "description", "VpcSubnet description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet.testVpcSubnet", "access_mode", "Public"),
				),
			},
		},
	})
}

const testAccNSXTDSVpcSubnetConfig = `

    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc"
	description = "VpcSubnet description"
	access_mode = "Public"
}

data "nsxt_vpc_subnet" "testVpcSubnet" {
  display_name = nsxt_vpc_subnet.testVpcSubnet.display_name
}
`
