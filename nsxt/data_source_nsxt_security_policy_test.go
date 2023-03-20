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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
						"nsxt_vpc_security_policy.testSecurityPolicy", "nsx_id", "test-securitypolicy-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy.testSecurityPolicy", "display_name", "test-securitypolicy-abc"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy.testSecurityPolicy", "description", "SecurityPolicy description"),
					resource.TestCheckResourceAttr(
						"nsxt_vpc_security_policy.testSecurityPolicy", "sequence_number", "0"),
				),
			},
		},
	})
}

const testAccNSXTDSSecurityPolicyConfig = `

    resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
      	nsx_id = "test-securitypolicy-abc"
	display_name = "test-securitypolicy-abc"
	description = "SecurityPolicy description"
	sequence_number = 0
}

data "nsxt_vpc_security_policy" "testSecurityPolicy" {
  display_name = nsxt_vpc_security_policy.testSecurityPolicy.display_name
}
`
