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

func TestNSXTDataSourceVpcVMBasic(t *testing.T) {
	testCaseExternalID := os.Getenv("NSXT_TEST_VM_EXTERNAL_ID")
	testCaseDisplayName := os.Getenv("NSXT_TEST_VM_DISPLAY_NAME")
	testResourceName := "data.nsxt_vpc_vm.testVM"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTDSVpcVMConfigTemplate(testCaseExternalID, testCaseDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(testResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(testResourceName, "external_id"),
					resource.TestCheckResourceAttrSet(testResourceName, "power_state"),
				),
			},
		},
	})
}

func testAccNSXTDSVpcVMConfigTemplate(testCaseExternalID string, testCaseDisplayName string) string {
	return fmt.Sprintf(`
  data "nsxt_vpc_vm" "testVM" {
		external_id		 = "%s"
    display_name   = "%s"
}`, testCaseExternalID, testCaseDisplayName)
}