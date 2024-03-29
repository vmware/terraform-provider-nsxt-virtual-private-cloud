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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestNSXTDataSourcePolicyContextProfileBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_TEST_CONTEXT_PROFILE_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_CONTEXT_PROFILE_NAME")
	testCaseScope := os.Getenv("NSXT_TEST_CONTEXT_PROFILE_SCOPE")
	testResourceName := "data.nsxt_vpc_policy_context_profile.testPolicyContextProfile"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSPolicyContextProfileConfigTemplate(testCaseNsxID, testCaseDisplayName, testCaseScope),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSPolicyContextProfileConfigTemplate(testCaseNsxID string, testCaseDisplayName string, testCaseScope string) string {
	return fmt.Sprintf(`
  data "nsxt_vpc_policy_context_profile" "testPolicyContextProfile" {
		nsx_id				 = "%s"
    display_name   = "%s"
		context {
			scope = "%s"
		}
}`, testCaseNsxID, testCaseDisplayName, testCaseScope)
}
