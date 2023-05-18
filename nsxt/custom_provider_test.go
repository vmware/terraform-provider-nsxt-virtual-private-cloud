/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package nsxt

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
	nsxtsession "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/session"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"nsxt": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	log.Printf("start of TestProvider")
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
	// Validating Schema for Provider

	var configs = map[string]interface{}{"host": "", "username": "", "password": "", "org": "", "project": "", "vpc": ""}

	diags := Provider().Validate(
		&terraform.ResourceConfig{Config: configs})
	if diags.HasError() {
		t.Fatalf("err in provider Validate: %v", diags)
	}

	// Validating vpc_subnet resource in nsxt provider and datasource for orgroot
	var vpcsubnetconfigsData = map[string]interface{}{"display_name": ""}

	diags = Provider().ValidateDataSource("nsxt_vpc_subnet",
		&terraform.ResourceConfig{Config: vpcsubnetconfigsData})
	if diags.HasError() {
		t.Fatalf("err in provider ValidateDataSource: %v", diags)
	}

	var vpcSubnetConfigsRes = map[string]interface{}{"nsx_id": "", "description": "", "display_name": "", "resource_type": "", "tags": make([]string, 0)}

	diags = Provider().ValidateResource("nsxt_vpc_subnet",
		&terraform.ResourceConfig{Config: vpcSubnetConfigsRes})
	if diags.HasError() {
		t.Fatalf("err in provider ValidateResource: %v", diags)
	}

	testAccPreCheck(t)
	log.Printf("end of TestProvider. Validation successful")
}

func testAccEnvDefined(t *testing.T, envVar string) {
	if len(os.Getenv(envVar)) == 0 {
		t.Skipf("This test requires %s environment variable to be set", envVar)
	}
}

func testAccPreCheck(t *testing.T) {
	log.Printf("start of testAccPreCheck")
	// Get below environment variables as set by user
	config := Configuration{
		NsxManagerHost: os.Getenv("NSXT_MANAGER_HOST"),
		Username:       os.Getenv("NSXT_USERNAME"),
		Password:       os.Getenv("NSXT_PASSWORD"),
		OrgId:          os.Getenv("NSXT_ORG"),
		ProjectId:      os.Getenv("NSXT_PROJECT"),
		VpcId:          os.Getenv("NSXT_VPC"),
	}

	if config.NsxManagerHost == "" {
		t.Fatalf("NSXT_MANAGER_HOST must be set for acceptance test")
	}

	if config.Username == "" {
		t.Fatalf("NSXT_USERNAME must be set for acceptance test")
	}

	if config.Password == "" {
		t.Fatalf("NSXT_PASSWORD must be set for acceptance test")
	}

	if config.OrgId == "" {
		t.Fatalf("NSXT_ORG must be set for acceptance test")
	}

	if config.ProjectId == "" {
		t.Fatalf("NSXT_PROJECT must be set for acceptance test")
	}

	if config.VpcId == "" {
		t.Fatalf("NSXT_VPC must be set for acceptance test")
	}

	_, err := nsxtclient.NewNsxtClient(config.NsxManagerHost, config.Username, config.OrgId, config.ProjectId, config.VpcId,
		"", "", "", "", "", "", true, nsxtsession.SetPassword(config.Password), nsxtsession.SetInsecure(false),
		nsxtsession.SetMaxAPIRetries(5), nsxtsession.SetTimeout(time.Duration(90*int(time.Second))))

	if err != nil {
		fmt.Println("Error occurred in creating new Nsxt client")
		t.Fatalf(fmt.Sprintf("%+v", err))
	}
	log.Printf("end of testAccPreCheck")
}
