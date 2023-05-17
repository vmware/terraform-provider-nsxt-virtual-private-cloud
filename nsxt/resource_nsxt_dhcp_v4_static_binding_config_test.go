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

func TestNSXTDhcpV4StaticBindingConfigBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTDhcpV4StaticBindingConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDhcpV4StaticBindingConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTDhcpV4StaticBindingConfigExists("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "display_name", "test-dhcpv4staticbindingconfig-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "description", "DHCP v4 static binding config description"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "nsx_id", "test-dhcpv4staticbinding-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "host_name", "vm1.vmware.com"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "mac_address", "11:22:33:44:55:67"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "ip_address", "192.168.4.32"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "resource_type", "DhcpV4StaticBindingConfig"),
				),
			},
			{
				Config: testAccNSXTDhcpV4StaticBindingConfigupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTDhcpV4StaticBindingConfigExists("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "display_name", "test-dhcpv4staticbindingconfig-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "description", "DHCP v4 static binding config description"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "nsx_id", "test-dhcpv4staticbinding-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "host_name", "vm1.vmware.com"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "mac_address", "11:22:33:44:55:67"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "ip_address", "192.168.4.32"),
					resource.TestCheckResourceAttr("nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "resource_type", "DhcpV4StaticBindingConfig"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTDhcpV4StaticBindingConfigConfig,
			},
		},
	})
}

func testAccCheckNSXTDhcpV4StaticBindingConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT DhcpV4StaticBindingConfig policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTDhcpV4StaticBindingConfigDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_dhcp_v4_static_binding_config" {
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
			return fmt.Errorf("NSXT DhcpV4StaticBindingConfig still exists")
		}
	}
	return nil
}

const testAccNSXTDhcpV4StaticBindingConfigConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc"
	description = "VpcSubnet description"
	access_mode = "Public"
}
    resource "nsxt_vpc_dhcp_v4_static_binding_config" "testDhcpV4StaticBindingConfig" {
      	display_name = "test-dhcpv4staticbindingconfig-abc"
	description = "DHCP v4 static binding config description"
	parent_path = nsxt_vpc_subnet.testVpcSubnet.path
	nsx_id = "test-dhcpv4staticbinding-abc"
	host_name = "vm1.vmware.com"
	mac_address = "11:22:33:44:55:67"
	ip_address = "192.168.4.32"
	options {
		option121 {
			static_routes {
	next_hop = "2.2.2.2"
	network = "10.22.12.1/23"
}
static_routes {
	next_hop = "0.0.0.0"
	network = "129.0.0.1/32"
}
		}
	}
	resource_type = "DhcpV4StaticBindingConfig"
}
`

const testAccNSXTDhcpV4StaticBindingConfigupdatedConfig = `
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	ipv4_subnet_size = 16
	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc-updated"
	description = "updated VpcSubnet description"
	access_mode = "Public"
}
    resource "nsxt_vpc_dhcp_v4_static_binding_config" "testDhcpV4StaticBindingConfig" {
      	display_name = "test-dhcpv4staticbindingconfig-abc-updated"
	description = "DHCP v4 static binding config description"
	parent_path = nsxt_vpc_subnet.testVpcSubnet.path
	nsx_id = "test-dhcpv4staticbinding-abc"
	host_name = "vm1.vmware.com"
	mac_address = "11:22:33:44:55:67"
	ip_address = "192.168.4.32"
	resource_type = "DhcpV4StaticBindingConfig"
}
`
