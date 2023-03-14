/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
package nsxt

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestNSXTDataSourceInfraIpAddressPoolBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_TEST_IP_ADDRESS_POOL_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_IP_ADDRESS_POOL_NAME")
	testResourceName := "data.nsxt_shared_infra_ip_address_pool.testInfraIpAddressPool"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSInfraIpAddressPoolConfigTemplate(testCaseNsxID, testCaseDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSInfraIpAddressPoolConfigTemplate(testCaseNsxID string, testCaseDisplayName string) string {
	return fmt.Sprintf(`
  data "nsxt_shared_infra_ip_address_pool" "testInfraIpAddressPool" {
		nsx_id				 = "%s"
    display_name   = "%s"
}`, testCaseNsxID, testCaseDisplayName)
}