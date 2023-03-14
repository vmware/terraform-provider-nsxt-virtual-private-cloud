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

func TestNSXTDataSourceProjectInfraPolicyContextProfileBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_TEST_CONTEXT_PROFILE_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_CONTEXT_PROFILE_NAME")
	testResourceName := "data.nsxt_shared_project_infra_policy_context_profile.testProjectInfraPolicyContextProfile"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSProjectInfraPolicyContextProfileConfigTemplate(testCaseNsxID, testCaseDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSProjectInfraPolicyContextProfileConfigTemplate(testCaseNsxID string, testCaseDisplayName string) string {
	return fmt.Sprintf(`
  data "nsxt_shared_project_infra_policy_context_profile" "testProjectInfraPolicyContextProfile" {
		nsx_id				 = "%s"
    display_name   = "%s"
}`, testCaseNsxID, testCaseDisplayName)
}
