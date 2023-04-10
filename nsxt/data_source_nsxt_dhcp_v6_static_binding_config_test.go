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

func TestNSXTDataSourceDhcpV6StaticBindingConfigBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSDhcpV6StaticBindingConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "display_name", "test-dhcpv6staticbindingconfig-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "description", "DHCP v6 static binding config description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "nsx_id", "test-dhcpv6staticbinding-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "mac_address", "11:22:33:44:55:67"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig", "resource_type", "DhcpV6StaticBindingConfig"),
				),
			},
		},
	})
}

const testAccNSXTDSDhcpV6StaticBindingConfigConfig = `

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

data "nsxt_vpc_dhcp_v6_static_binding_config" "testDhcpV6StaticBindingConfig" {
  display_name = nsxt_vpc_dhcp_v6_static_binding_config.testDhcpV6StaticBindingConfig.display_name
}
`
