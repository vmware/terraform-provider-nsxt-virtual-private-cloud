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

func TestNSXTDataSourceInfraL2BridgeEndpointProfileBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_TEST_L2_BRIDGE_ENDPOINT_PROFILE_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_L2_BRIDGE_ENDPOINT_PROFILE_NAME")
	testResourceName := "data.nsxt_shared_infra_l2_bridge_endpoint_profile.testInfraL2BridgeEndpointProfile"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSInfraL2BridgeEndpointProfileConfigTemplate(testCaseNsxID, testCaseDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSInfraL2BridgeEndpointProfileConfigTemplate(testCaseNsxID string, testCaseDisplayName string) string {
	return fmt.Sprintf(`
  data "nsxt_shared_infra_l2_bridge_endpoint_profile" "testInfraL2BridgeEndpointProfile" {
		nsx_id				 = "%s"
    display_name   = "%s"
}`, testCaseNsxID, testCaseDisplayName)
}
