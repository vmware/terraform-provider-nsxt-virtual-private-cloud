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

func TestNSXTDhcpV6StaticBindingConfigBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTDhcpV6StaticBindingConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDhcpV6StaticBindingConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTDhcpV6StaticBindingConfigExists("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "display_name", "test-dhcpv6staticbindingconfig-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "description", "DHCP v6 static binding config description"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "nsx_id", "test-dhcpv6staticbinding-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "mac_address", "11:22:33:44:55:67"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "resource_type", "DhcpV6StaticBindingConfig"),
				),
			},
			{
				Config: testAccNSXTDhcpV6StaticBindingConfigupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTDhcpV6StaticBindingConfigExists("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "display_name", "test-dhcpv6staticbindingconfig-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "description", "DHCP v6 static binding config description"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "nsx_id", "test-dhcpv6staticbinding-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "mac_address", "11:22:33:44:55:67"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "resource_type", "DhcpV6StaticBindingConfig"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTDhcpV6StaticBindingConfigConfig,
			},
		},
	})
}

func testAccCheckNSXTDhcpV6StaticBindingConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT DhcpV6StaticBindingConfig policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTDhcpV6StaticBindingConfigDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_dhcp_v6_static_binding_config" {
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
			return fmt.Errorf("NSXT DhcpV6StaticBindingConfig still exists")
		}
	}
	return nil
}

const testAccNSXTDhcpV6StaticBindingConfigConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc"
	description = "VpcSubnet description"
	access_mode = "Public"
}
    resource "nsxt_vpc_dhcp_v6_static_binding_config" "testDhcpV6StaticBindingConfig" {
      	display_name = "test-dhcpv6staticbindingconfig-abc"
	description = "DHCP v6 static binding config description"
	parent_path = nsxt_vpc_subnet.testVpcSubnet.path
	nsx_id = "test-dhcpv6staticbinding-abc"
	mac_address = "11:22:33:44:55:67"
	resource_type = "DhcpV6StaticBindingConfig"
}
`

const testAccNSXTDhcpV6StaticBindingConfigupdatedConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc-updated"
	description = "updated VpcSubnet description"
	access_mode = "Public"
}
    resource "nsxt_vpc_dhcp_v6_static_binding_config" "testDhcpV6StaticBindingConfig" {
      	display_name = "test-dhcpv6staticbindingconfig-abc-updated"
	description = "DHCP v6 static binding config description"
	parent_path = nsxt_vpc_subnet.testVpcSubnet.path
	nsx_id = "test-dhcpv6staticbinding-abc"
	mac_address = "11:22:33:44:55:67"
	resource_type = "DhcpV6StaticBindingConfig"
}
`
