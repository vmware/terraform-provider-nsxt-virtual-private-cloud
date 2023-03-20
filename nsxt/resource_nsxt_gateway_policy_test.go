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

func TestNSXTGatewayPolicyBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTGatewayPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTGatewayPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGatewayPolicyExists("nsxt_vpc_gateway_policy.testGatewayPolicy"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "nsx_id", "test-gatewaypolicy-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "display_name", "test-gatewaypolicy-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "description", "GatewayPolicy description"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "sequence_number", "0"),
				),
			},
			{
				Config: testAccNSXTGatewayPolicyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTGatewayPolicyExists("nsxt_vpc_gateway_policy.testGatewayPolicy"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "nsx_id", "test-gatewaypolicy-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "display_name", "test-gatewaypolicy-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "description", "updated GatewayPolicy description"),
					resource.TestCheckResourceAttr("nsxt_vpc_gateway_policy.testGatewayPolicy", "sequence_number", "0"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_gateway_policy.testGatewayPolicy",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTGatewayPolicyConfig,
			},
		},
	})
}

func testAccCheckNSXTGatewayPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT GatewayPolicy policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTGatewayPolicyDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_gateway_policy" {
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
			return fmt.Errorf("NSXT GatewayPolicy still exists")
		}
	}
	return nil
}

//nolint

const testAccNSXTGatewayPolicyConfig = `
    resource "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
      	nsx_id = "test-gatewaypolicy-abc"
	display_name = "test-gatewaypolicy-abc"
	description = "GatewayPolicy description"
	sequence_number = 0
}
`

const testAccNSXTGatewayPolicyupdatedConfig = `
    resource "nsxt_vpc_gateway_policy" "testGatewayPolicy" {
      	nsx_id = "test-gatewaypolicy-abc"
	display_name = "test-gatewaypolicy-abc-updated"
	description = "updated GatewayPolicy description"
	sequence_number = 0
}
`
