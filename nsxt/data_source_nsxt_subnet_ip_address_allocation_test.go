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

func TestNSXTDataSourceIpAddressAllocationBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             testAccNSXTDSIpAddressAllocationConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "nsx_id", "test-ipallocation-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "display_name", "test-ipallocation-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "description", "IpAllocation description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "allocation_ip", "192.168.12.2"),
				),
			},
		},
	})
}

const testAccNSXTDSIpAddressAllocationConfig = `

    resource "nsxt_vpc_subnet_ip_address_allocation" "testIpAddressAllocation" {
      	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/subnets/test-vpcsubnet-abc/ip-pools/static-ipv4-default"
	nsx_id = "test-ipallocation-abc"
	display_name = "test-ipallocation-abc"
	description = "IpAllocation description"
	allocation_ip = "192.168.12.2"
}

data "nsxt_vpc_subnet_ip_address_allocation" "testIpAddressAllocation" {
  display_name = nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation.display_name
}
`