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

func TestNSXTDataSourceStaticRoutesBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSStaticRoutesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_static_routes.testStaticRoutes", "nsx_id", "test-staticroutes-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_static_routes.testStaticRoutes", "display_name", "test-staticroutes-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_static_routes.testStaticRoutes", "description", "StaticRoutes description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_static_routes.testStaticRoutes", "network", "45.1.1.0/24"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_static_routes.testStaticRoutes", "next_hops.*", map[string]string{
							"ip_address": "41.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"nsxt_vpc_static_routes.testStaticRoutes", "next_hops.*", map[string]string{
							"admin_distance": "1"}),
				),
			},
		},
	})
}

const testAccNSXTDSStaticRoutesConfig = `

    resource "nsxt_vpc_static_routes" "testStaticRoutes" {
      	nsx_id = "test-staticroutes-abc"
	display_name = "test-staticroutes-abc"
	description = "StaticRoutes description"
	network = "45.1.1.0/24"
	next_hops {
	ip_address = "41.1.1.1"
	admin_distance = 1
}
next_hops {
	ip_address = "42.1.1.1"
	admin_distance = 2
}
next_hops {
	ip_address = "43.1.2.3"
	admin_distance = 3
}
}

data "nsxt_vpc_static_routes" "testStaticRoutes" {
  display_name = nsxt_vpc_static_routes.testStaticRoutes.display_name
}
`
