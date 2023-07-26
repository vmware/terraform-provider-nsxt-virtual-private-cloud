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

func TestNSXTDataSourceGroupBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_TEST_GROUP_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_GROUP_NAME")
	testCaseContext := os.Getenv("NSXT_TEST_GROUP_CONTEXT")
	testCaseDomain := os.Getenv("NSXT_TEST_GROUP_DOMAIN")
	testResourceName := "data.nsxt_vpc_group.testGroup"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSGroupConfigTemplate(testCaseNsxID, testCaseDisplayName, testCaseContext, testCaseDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSGroupConfigTemplate(testCaseNsxID string, testCaseDisplayName string, testCaseContext string, testCaseDomain string) string {
	return fmt.Sprintf(`
  data "nsxt_vpc_group" "testGroup" {
		nsx_id				 = "%s"
    display_name   = "%s"
		context_info {
			context = "%s"
	domain = "%s"
		}
}`, testCaseNsxID, testCaseDisplayName, testCaseContext, testCaseDomain)
}
