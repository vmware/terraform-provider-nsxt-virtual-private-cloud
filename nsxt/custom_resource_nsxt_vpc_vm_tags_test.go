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
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	nsxtclient "github.com/vmware/terraform-provider-nsxt-virtual-private-cloud/nsxt/clients"
)

func TestNSXTVpcVmTagsBasic(t *testing.T) {
	testCaseVMID := os.Getenv("NSXT_TEST_VM_ID")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccEnvDefined(t, "NSXT_TEST_VM_ID") },
		Providers:    testAccProviders,
		CheckDestroy: testAccNSXTVpcVmTagsCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNSXTVpcVmTagsConfig(testCaseVMID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTVpcVmTagsExists("nsxt_vpc_vm_tags.testVpcVmTags"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_vm_tags.testVpcVmTags", "tags.*", map[string]string{
						"scope": "os"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_vm_tags.testVpcVmTags", "tags.*", map[string]string{
						"tag": "windows"}),
				),
			},
			{
				Config: testAccNSXTVpcVmTagsUpdatedConfig(testCaseVMID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNSXTVpcVmTagsExists("nsxt_vpc_vm_tags.testVpcVmTags"),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_vm_tags.testVpcVmTags", "tags.*", map[string]string{
						"scope": "os"}),
					resource.TestCheckTypeSetElemNestedAttrs("nsxt_vpc_vm_tags.testVpcVmTags", "tags.*", map[string]string{
						"tag": "linux"}),
				),
			},
			{
				ResourceName:      "nsxt_vpc_vm_tags.testVpcVmTags",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccNSXTVpcVmTagsConfig(testCaseVMID),
			},
		},
	})
}

func testAccCheckNSXTVpcVmTagsExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		vm_id := rs.Primary.ID
		if vm_id == "" {
			return fmt.Errorf("VM ID not set")
		} else {
			url := nsxtClient.Config.BasePath + "/search?query=resource_type:VirtualMachine%20AND%20external_id:" + vm_id + "&context=vpcs:/orgs/" +
				nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
			err := nsxtClient.NsxtSession.Get(url, &obj)
			if err != nil {
				log.Printf("[ERROR] in reading VM %v\n", err)
				return err
			}
		}
		return nil
	}
}

func testAccNSXTVpcVmTagsCheckDestroy(s *terraform.State) error {
	nsxtClient := testAccProvider.Meta().(*nsxtclient.NsxtClient)
	var obj interface{}
	tags_deleted := true
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "nsxt_vpc_vm_tags" {
			continue
		}
		vm_id := rs.Primary.ID
		url := nsxtClient.Config.BasePath + "/search?query=resource_type:VirtualMachine%20AND%20external_id:" + vm_id + "&context=vpcs:/orgs/" +
			nsxtClient.Config.OrgID + "/projects/" + nsxtClient.Config.ProjectID + "/vpcs/" + nsxtClient.Config.VpcID
		err := nsxtClient.NsxtSession.Get(url, &obj)
		if err != nil {
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			objMap, ok := obj.(map[string]interface{})
			if !ok {
				return fmt.Errorf("want type map[string]interface{};  got %T", objMap)
			}
			// get results of type []interface {}
			results := objMap["results"]
			if objMap["result_count"].(float64) == 1 {
				for key := range results.([]interface{})[0].(map[string]interface{}) {
					if key == "tags" {
						tags_deleted = false
						break
					}
				}
			}
			if !tags_deleted {
				return fmt.Errorf("VM %s still has tags, although nsxt_vpc_vm_tags was deleted", vm_id)
			}
		}
	}
	return nil
}

func testAccNSXTVpcVmTagsConfig(vmId string) string {
	return fmt.Sprintf(`
	resource "nsxt_vpc_vm_tags" "testVpcVmTags" {
    virtual_machine_id = "%s"
  	tags {
			scope = "os"
			tag = "windows"
  	}
}
`, vmId)
}

func testAccNSXTVpcVmTagsUpdatedConfig(vmId string) string {
	return fmt.Sprintf(`
	resource "nsxt_vpc_vm_tags" "testVpcVmTags" {
		virtual_machine_id = "%s"
  	tags {
			scope = "os"
			tag = "linux"
  	}
}
`, vmId)
}
