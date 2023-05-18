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

func TestNSXTVpcSubnetBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTVpcSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTVpcSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTVpcSubnetExists("nsxt_vpc_subnet.testVpcSubnet"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "ipv4_subnet_size", "16"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "nsx_id", "test-vpcsubnet-abc-2"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "display_name", "test-vpcsubnet-abc-2"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "description", "VpcSubnet 2 description"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "access_mode", "Public"),
				),
			},
			{
				Config: testAccNSXTVpcSubnetupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTVpcSubnetExists("nsxt_vpc_subnet.testVpcSubnet"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "ipv4_subnet_size", "16"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "nsx_id", "test-vpcsubnet-abc-2"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "display_name", "test-vpcsubnet-abc-2-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "description", "updated VpcSubnet 2 description"),
					resource.TestCheckResourceAttr("nsxt_vpc_subnet.testVpcSubnet", "access_mode", "Public"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_subnet.testVpcSubnet",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTVpcSubnetConfig,
			},
		},
	})
}

func testAccCheckNSXTVpcSubnetExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT VpcSubnet policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTVpcSubnetDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_subnet" {
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
			return fmt.Errorf("NSXT VpcSubnet still exists")
		}
	}
	return nil
}

const testAccNSXTVpcSubnetConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc-2"
	display_name = "test-vpcsubnet-abc-2"
	description = "VpcSubnet 2 description"
	access_mode = "Public"
}
`

const testAccNSXTVpcSubnetupdatedConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc-2"
	display_name = "test-vpcsubnet-abc-2-updated"
	description = "updated VpcSubnet 2 description"
	access_mode = "Public"
}
`
