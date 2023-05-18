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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestNSXTDataSourceSecurityPolicyBasic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSSecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy.testSecurityPolicy", "nsx_id", "test-securitypolicy-abc-2"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy.testSecurityPolicy", "display_name", "test-securitypolicy-abc-2"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy.testSecurityPolicy", "description", "SecurityPolicy 2 description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy.testSecurityPolicy", "sequence_number", "0"),
				),
			},
		},
	})
}

const testAccNSXTDSSecurityPolicyConfig = `

    resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
      	nsx_id = "test-securitypolicy-abc-2"
	display_name = "test-securitypolicy-abc-2"
	description = "SecurityPolicy 2 description"
	sequence_number = 0
}

data "nsxt_vpc_security_policy" "testSecurityPolicy" {
  display_name = nsxt_vpc_security_policy.testSecurityPolicy.display_name
}
`
