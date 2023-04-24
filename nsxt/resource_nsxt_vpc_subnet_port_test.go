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

func TestNSXTVpcSubnetPortBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTVpcSubnetPortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTVpcSubnetPortConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTVpcSubnetPortExists("nsxt_vpc_subnet_port.testVpcSubnetPort"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testVpcSubnetPort", "nsx_id", "test-port-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testVpcSubnetPort", "display_name", "test-segmentport-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testVpcSubnetPort", "description", "SegmentPort description"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testVpcSubnetPort", "address_bindings.*", map[string]string{
						"ip_address": "1.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testVpcSubnetPort", "address_bindings.*", map[string]string{
						"mac_address": "aa:bb:cc:dd:ee:ff"}),
				),
			},
			{
				Config: testAccNSXTVpcSubnetPortupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTVpcSubnetPortExists("nsxt_vpc_subnet_port.testVpcSubnetPort"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testVpcSubnetPort", "nsx_id", "test-port-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testVpcSubnetPort", "display_name", "test-segmentport-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testVpcSubnetPort", "description", "updated SegmentPort description"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testVpcSubnetPort", "address_bindings.*", map[string]string{
						"ip_address": "1.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testVpcSubnetPort", "address_bindings.*", map[string]string{
						"mac_address": "aa:bb:cc:dd:ee:ff"}),
				),
			},
			{
				ResourceName:      "nsxt_vpc_subnet_port.testVpcSubnetPort",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTVpcSubnetPortConfig,
			},
		},
	})
}

func testAccCheckNSXTVpcSubnetPortExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT VpcSubnetPort policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTVpcSubnetPortDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_subnet_port" {
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
			return fmt.Errorf("NSXT VpcSubnetPort still exists")
		}
	}
	return nil
}

const testAccNSXTVpcSubnetPortConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc"
	description = "VpcSubnet description"
	access_mode = "Public"
}
    resource "nsxt_vpc_subnet_port" "testVpcSubnetPort" {
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
`

const testAccNSXTVpcSubnetPortupdatedConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc-updated"
	description = "updated VpcSubnet description"
	access_mode = "Public"
}
    resource "nsxt_vpc_subnet_port" "testVpcSubnetPort" {
      	parent_path = nsxt_vpc_subnet.testVpcSubnet.path
	nsx_id = "test-port-abc"
	display_name = "test-segmentport-abc-updated"
	description = "updated SegmentPort description"
	address_bindings {
	ip_address = "1.1.1.1"
	mac_address = "aa:bb:cc:dd:ee:ff"
}
address_bindings {
	ip_address = "1.1.1.2"
	mac_address = "aa:bb:cc:dd:ee:f1"
}
}
`
