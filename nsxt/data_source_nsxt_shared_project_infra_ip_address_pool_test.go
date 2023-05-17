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

func TestNSXTDataSourceProjectInfraIpAddressPoolBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_PROJECT_INFRA_TEST_IP_ADDRESS_POOL_ID")
	testCaseDisplayName := os.Getenv("NSXT_PROJECT_INFRA_TEST_IP_ADDRESS_POOL_NAME")
	testResourceName := "data.nsxt_shared_project_infra_ip_address_pool.testProjectInfraIpAddressPool"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSProjectInfraIpAddressPoolConfigTemplate(testCaseNsxID, testCaseDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSProjectInfraIpAddressPoolConfigTemplate(testCaseNsxID string, testCaseDisplayName string) string {
	return fmt.Sprintf(`
  data "nsxt_shared_project_infra_ip_address_pool" "testProjectInfraIpAddressPool" {
		nsx_id				 = "%s"
    display_name   = "%s"
}`, testCaseNsxID, testCaseDisplayName)
}
