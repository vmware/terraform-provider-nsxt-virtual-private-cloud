/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

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

func TestNSXTSubnetPortBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTSubnetPortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTSubnetPortConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTSubnetPortExists("nsxt_vpc_subnet_port.testSubnetPort"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testSubnetPort", "nsx_id", "test-port-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testSubnetPort", "display_name", "test-segmentport-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testSubnetPort", "description", "SegmentPort description"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testSubnetPort", "address_bindings.*", map[string]string{
						"ip_address": "1.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testSubnetPort", "address_bindings.*", map[string]string{
						"mac_address": "aa:bb:cc:dd:ee:ff"}),
				),
			},
			{
				Config: testAccNSXTSubnetPortupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTSubnetPortExists("nsxt_vpc_subnet_port.testSubnetPort"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testSubnetPort", "nsx_id", "test-port-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testSubnetPort", "display_name", "test-segmentport-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet_port.testSubnetPort", "description", "updated SegmentPort description"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testSubnetPort", "address_bindings.*", map[string]string{
						"ip_address": "1.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_subnet_port.testSubnetPort", "address_bindings.*", map[string]string{
						"mac_address": "aa:bb:cc:dd:ee:ff"}),
				),
			},
			{
				ResourceName:      "nsxt_vpc_subnet_port.testSubnetPort",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTSubnetPortConfig,
			},
		},
	})
}

func testAccCheckNSXTSubnetPortExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT SubnetPort policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTSubnetPortDestroy(s *terraform.State) error {
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
			return fmt.Errorf("NSXT SubnetPort still exists")
		}
	}
	return nil
}

//nolint

const testAccNSXTSubnetPortConfig = `
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
`

const testAccNSXTSubnetPortupdatedConfig = `
    resource "nsxt_vpc_subnet_port" "testSubnetPort" {
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
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc-updated"
	description = "updated VpcSubnet description"
	ipv4_subnet_size = 16
	access_mode = "Public"
}
`
