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

func TestNSXTStaticRoutesBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTStaticRoutesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTStaticRoutesConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTStaticRoutesExists("nsxt_vpc_static_routes.testStaticRoutes"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_static_routes.testStaticRoutes", "next_hops.*", map[string]string{
						"ip_address": "41.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_static_routes.testStaticRoutes", "next_hops.*", map[string]string{
						"admin_distance": "1"}),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "nsx_id", "test-staticroutes-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "display_name", "test-staticroutes-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "description", "StaticRoutes description"),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "network", "45.1.1.0/24"),
				),
			},
			{
				Config: testAccNSXTStaticRoutesupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTStaticRoutesExists("nsxt_vpc_static_routes.testStaticRoutes"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_static_routes.testStaticRoutes", "next_hops.*", map[string]string{
						"ip_address": "41.1.1.1"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_static_routes.testStaticRoutes", "next_hops.*", map[string]string{
						"admin_distance": "1"}),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "nsx_id", "test-staticroutes-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "display_name", "test-staticroutes-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "description", "StaticRoutes description"),
					resource.TestCheckResourceAttr("nsxt_vpc_static_routes.testStaticRoutes", "network", "45.1.1.0/24"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_static_routes.testStaticRoutes",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTStaticRoutesConfig,
			},
		},
	})
}

func testAccCheckNSXTStaticRoutesExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT StaticRoutes policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTStaticRoutesDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_static_routes" {
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
			return fmt.Errorf("NSXT StaticRoutes still exists")
		}
	}
	return nil
}

const testAccNSXTStaticRoutesConfig = `
    resource "nsxt_vpc_static_routes" "testStaticRoutes" {
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
	nsx_id = "test-staticroutes-abc"
	display_name = "test-staticroutes-abc"
	description = "StaticRoutes description"
	network = "45.1.1.0/24"
}
`

const testAccNSXTStaticRoutesupdatedConfig = `
    resource "nsxt_vpc_static_routes" "testStaticRoutes" {
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
	nsx_id = "test-staticroutes-abc"
	display_name = "test-staticroutes-abc-updated"
	description = "StaticRoutes description"
	network = "45.1.1.0/24"
}
`
