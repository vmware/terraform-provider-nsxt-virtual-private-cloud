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

func TestNSXTDataSourceIpAddressPoolBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_TEST_IP_ADDRESS_POOL_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_IP_ADDRESS_POOL_NAME")
	testCaseContext := os.Getenv("NSXT_TEST_IP_ADDRESS_POOL_CONTEXT")
	testResourceName := "data.nsxt_vpc_ip_address_pool.testIpAddressPool"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSIpAddressPoolConfigTemplate(testCaseNsxID, testCaseDisplayName, testCaseContext),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSIpAddressPoolConfigTemplate(testCaseNsxID string, testCaseDisplayName string, testCaseContext string) string {
	return fmt.Sprintf(`
  data "nsxt_vpc_ip_address_pool" "testIpAddressPool" {
		nsx_id				 = "%s"
    display_name   = "%s"
		context_info {
			context = "%s"
		}
}`, testCaseNsxID, testCaseDisplayName, testCaseContext)
}
