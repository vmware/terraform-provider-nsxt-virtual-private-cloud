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

func TestNSXTIpAddressAllocationBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTIpAddressAllocationDestroy,
		Steps: []resource.TestStep{
			{
				Config:             testAccNSXTIpAddressAllocationConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTIpAddressAllocationExists("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "nsx_id", "test-ipallocation-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "display_name", "test-ipallocation-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "description", "IpAllocation description"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "allocation_ip", "192.168.12.2"),
				),
			},
			{
				Config:             testAccNSXTIpAddressAllocationupdatedConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTIpAddressAllocationExists("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "nsx_id", "test-ipallocation-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "display_name", "test-ipallocation-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_ip_address_allocation.testIpAddressAllocation", "description", "IpAllocation description"),
				),
			},
			{
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				ImportStateVerify:  false,
				Config:             testAccNSXTIpAddressAllocationConfig,
			},
		},
	})
}

func testAccCheckNSXTIpAddressAllocationExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT IpAddressAllocation policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTIpAddressAllocationDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_subnet_ip_address_allocation" {
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
			return fmt.Errorf("NSXT IpAddressAllocation still exists")
		}
	}
	return nil
}

//nolint

const testAccNSXTIpAddressAllocationConfig = `
    resource "nsxt_vpc_subnet_ip_address_allocation" "testSubnetIpAddressAllocation" {
      	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/subnets/test-vpcsubnet-abc/ip-pools/static-ipv4-default"
	nsx_id = "test-ipallocation-abc"
	display_name = "test-ipallocation-abc"
	description = "IpAllocation description"
	allocation_ip = "192.168.12.2"
}
`

const testAccNSXTIpAddressAllocationupdatedConfig = `
    resource "nsxt_vpc_subnet_ip_address_allocation" "testSubnetIpAddressAllocation" {
      	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/subnets/test-vpcsubnet-abc/ip-pools/static-ipv4-default"
	nsx_id = "test-ipallocation-abc"
	display_name = "test-ipallocation-abc-updated"
	description = "IpAllocation description"
}
`
