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

func TestNSXTSecurityPolicyBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNSXTSecurityPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTSecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTSecurityPolicyExists("nsxt_vpc_security_policy.testSecurityPolicy"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "nsx_id", "test-securitypolicy-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "display_name", "test-securitypolicy-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "description", "SecurityPolicy description"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "sequence_number", "0"),
				),
			},
			{
				Config: testAccNSXTSecurityPolicyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTSecurityPolicyExists("nsxt_vpc_security_policy.testSecurityPolicy"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "nsx_id", "test-securitypolicy-abc"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "display_name", "test-secutitypolicy-abc-updated"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "description", "updated SecurityPolicy description"),
					resource.TestCheckResourceAttr("nsxt_vpc_security_policy.testSecurityPolicy", "sequence_number", "0"),
				),
			},
			{
				ResourceName:      "nsxt_vpc_security_policy.testSecurityPolicy",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTSecurityPolicyConfig,
			},
		},
	})
}

func testAccCheckNSXTSecurityPolicyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		path := nsxtClient.Config.BasePath + rs.Primary.Attributes["path"]
		if path == "" {
			return fmt.Errorf("No NSXT SecurityPolicy policy path is set")
		}
		err := nsxtClient.NsxtSession.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckNSXTSecurityPolicyDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_security_policy" {
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
			return fmt.Errorf("NSXT SecurityPolicy still exists")
		}
	}
	return nil
}

const testAccNSXTSecurityPolicyConfig = `
    resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
      	nsx_id = "test-securitypolicy-abc"
	display_name = "test-securitypolicy-abc"
	description = "SecurityPolicy description"
	sequence_number = 0
}
`

const testAccNSXTSecurityPolicyupdatedConfig = `
    resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
      	nsx_id = "test-securitypolicy-abc"
	display_name = "test-secutitypolicy-abc-updated"
	description = "updated SecurityPolicy description"
	sequence_number = 0
}
`
