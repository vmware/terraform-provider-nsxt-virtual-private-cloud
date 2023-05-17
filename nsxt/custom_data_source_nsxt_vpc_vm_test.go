/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// nolint
package nsxt

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestNSXTDataSourceVpcVMBasic(t *testing.T) {
	testCaseExternalID := os.Getenv("NSXT_TEST_VM_EXTERNAL_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_VM_DISPLAY_NAME")
	testCasePowerState := os.Getenv("NSXT_TEST_VM_POWER_STATE")
	testResourceName := "data.nsxt_vpc_vm.testVM"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSVpcVMConfigTemplate(testCaseExternalID, testCaseDisplayName, testCasePowerState),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "external_id"),
					resource.TestCheckResourceAttrSet(testResourceName, "power_state"),
				),
			},
		},
	})
}

func testAccNSXTDSVpcVMConfigTemplate(testCaseExternalID string, testCaseDisplayName string, testCasePowerState string) string {
	return fmt.Sprintf(`
  data "nsxt_vpc_vm" "testVM" {
		external_id		 = "%s"
    display_name   = "%s"
		power_state		 = "%s"
}`, testCaseExternalID, testCaseDisplayName, testCasePowerState)
}
