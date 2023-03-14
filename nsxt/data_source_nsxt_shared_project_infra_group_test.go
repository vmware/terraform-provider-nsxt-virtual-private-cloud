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

func TestNSXTDataSourceProjectInfraGroupBasic(t *testing.T) {
	testCaseNsxID := os.Getenv("NSXT_TEST_GROUP_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_GROUP_NAME")
	testResourceName := "data.nsxt_shared_project_infra_group.testProjectInfraGroup"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSProjectInfraGroupConfigTemplate(testCaseNsxID, testCaseDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNSXTDSProjectInfraGroupConfigTemplate(testCaseNsxID string, testCaseDisplayName string) string {
	return fmt.Sprintf(`
  data "nsxt_shared_project_infra_group" "testProjectInfraGroup" {
		nsx_id				 = "%s"
    display_name   = "%s"
}`, testCaseNsxID, testCaseDisplayName)
}
