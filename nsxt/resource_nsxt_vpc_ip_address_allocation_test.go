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
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
)

func TestNSXTVpcIpAddressAllocationBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTVpcIpAddressAllocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTVpcIpAddressAllocationConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTVpcIpAddressAllocationExists("nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation"),
					resource.TestCheckResourceAttr("nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation", "nsx_id", "test-vpcipaddressallocation-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation", "display_name", "test-vpcipaddressallocation-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation", "description", "Vpc IP address allocation description"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_ip_address_allocation.testVpcIpAddressAllocation",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTVpcIpAddressAllocationConfig,
			},
		},
	})
}

func testAccCheckNSXTVpcIpAddressAllocationExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT VpcIpAddressAllocation policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTVpcIpAddressAllocationDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_ip_address_allocation" {
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
			return fmt.Errorf("NSXT VpcIpAddressAllocation still exists")
		}
	}
	return nil
}

const testAccNSXTVpcIpAddressAllocationConfig = `
    resource "nsxt_vpc_ip_address_allocation" "testVpcIpAddressAllocation" {
      	nsx_id = "test-vpcipaddressallocation-abc"
	display_name = "test-vpcipaddressallocation-abc"
	description = "Vpc IP address allocation description"
}
`
