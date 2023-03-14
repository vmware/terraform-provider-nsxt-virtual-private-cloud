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

func TestNSXTDataSourceDhcpV4StaticBindingConfigBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSDhcpV4StaticBindingConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "nsx_id", "test-dhcpv4staticbinding-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "display_name", "test-dhcpv4staticbindingconfig-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "description", "DHCP v4 static binding config description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "resource_type", "DhcpV4StaticBindingConfig"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "gateway_address", "30.30.30.17"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "mac_address", "11:22:33:44:55:67"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "host_name", "vm1.vmware.com"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig", "ip_address", "30.30.30.16"),
				),
			},
		},
	})
}

const testAccNSXTDSDhcpV4StaticBindingConfigConfig = `

    resource "nsxt_vpc_dhcp_v4_static_binding_config" "testDhcpV4StaticBindingConfig" {
      	parent_path = nsxt_vpc_subnet.testVpcSubnet.path
	nsx_id = "test-dhcpv4staticbinding-abc"
	display_name = "test-dhcpv4staticbindingconfig-abc"
	description = "DHCP v4 static binding config description"
	resource_type = "DhcpV4StaticBindingConfig"
	gateway_address = "30.30.30.17"
	mac_address = "11:22:33:44:55:67"
	host_name = "vm1.vmware.com"
	ip_address = "30.30.30.16"
	options {
		option121 {
			static_routes {
	network = "10.22.12.1/23"
	next_hop = "2.2.2.2"
}
static_routes {
	network = "129.0.0.1/32"
	next_hop = "0.0.0.0"
}
		}
	}
}
    resource "nsxt_vpc_subnet" "testVpcSubnet" {
      	nsx_id = "test-vpcsubnet-abc"
	display_name = "test-vpcsubnet-abc"
	description = "VpcSubnet description"
	ipv4_subnet_size = 16
	access_mode = "Public"
}

data "nsxt_vpc_dhcp_v4_static_binding_config" "testDhcpV4StaticBindingConfig" {
  display_name = nsxt_vpc_dhcp_v4_static_binding_config.testDhcpV4StaticBindingConfig.display_name
}
`
