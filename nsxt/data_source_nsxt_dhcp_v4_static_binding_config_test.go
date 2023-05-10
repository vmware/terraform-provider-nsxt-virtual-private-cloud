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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestNSXTDataSourceDhcpV4StaticBindingConfigBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSDhcpV4StaticBindingConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "display_name", "test-dhcpv4staticbindingconfig-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "description", "DHCP v4 static binding config description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "nsx_id", "test-dhcpv4staticbinding-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "host_name", "vm1.vmware.com"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "mac_address", "11:22:33:44:55:67"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "ip_address", "192.168.4.32"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "resource_type", "DhcpV4StaticBindingConfig"),
				),
			},
		},
	})
}

const testAccNSXTDSDhcpV4StaticBindingConfigConfig = `

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

data "nsxt_vpc_dhcp_v4_static_binding_config" "testDhcpV4StaticBindingConfig" {
  display_name = nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig.display_name
}
`
