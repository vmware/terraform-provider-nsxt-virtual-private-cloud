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

func TestNSXTDataSourceVpcIpAddressAllocationBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSVpcIpAddressAllocationConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation", "nsx_id", "test-vpcipaddressallocation-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation", "display_name", "test-vpcipaddressallocation-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation", "description", "Vpc IP address allocation description"),
				),
			},
		},
	})
}

const testAccNSXTDSVpcIpAddressAllocationConfig = `

    resource "nsxt_vpc_ip_address_allocation" "testVpcIpAddressAllocation" {
      	nsx_id = "test-vpcipaddressallocation-abc"
	display_name = "test-vpcipaddressallocation-abc"
	description = "Vpc IP address allocation description"
}

data "nsxt_vpc_ip_address_allocation" "testVpcIpAddressAllocation" {
  display_name = nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation.display_name
  context {
    scope = "vpc"
  }
}
`
