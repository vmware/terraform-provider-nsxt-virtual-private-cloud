/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package nsxt

import (
	"log"
	"os"
	"testing"
)

func testNsxtSession(t *testing.T, nsxtsess *NsxtSession) {
	//GET
	var res interface{}
	errInGet := nsxtsess.Get("policy/api/v1/infra/domains/default/groups", &res)
	if errInGet != nil || res == nil {
		t.Error("Group GET failed: ", errInGet)
		return
	}

	// create a group
	group := make(map[string]interface{})
	group["display_name"] = "testgroupviasession"
	err := nsxtsess.Patch("policy/api/v1/infra/domains/default/groups/group-test-1", group, &res)
	if err != nil {
		t.Error("Group Creation failed: ", err)
		return
	}

	// check group is created well
	err = nsxtsess.Get("policy/api/v1/infra/domains/default/groups/group-test-1", &res)
	if err != nil {
		t.Error("Group GET after PATCH failed: ", err)
		return
	}
	resp := res.(map[string]interface{})
	log.Printf("[DEBUG] Response for GET after PATCH *******: %+v", resp)
	if resp["display_name"] != "testgroupviasession" {
		t.Error("Expected group not fetched")
		return
	}

	// update  group
	resp["display_name"] = "updated_group_test_1"
	err = nsxtsess.Put("policy/api/v1/infra/domains/default/groups/group-test-1", resp, &res)
	if err != nil {
		t.Error("Group update failed: ", err)
		return
	}
	// check group is created well
	err = nsxtsess.Get("policy/api/v1/infra/domains/default/groups/group-test-1", &res)
	if err != nil {
		t.Error("Group GET after PUT failed: ", err)
		return
	}
	resp = res.(map[string]interface{})
	log.Printf("[DEBUG] Response for GET after PUT *******: %+v", resp)
	if resp["display_name"] != "updated_group_test_1" {
		t.Error("Expected group not fetched")
		return
	}

	// delete the group
	err = nsxtsess.Delete("policy/api/v1/infra/domains/default/groups/group-test-1")
	log.Printf("[ERROR] Error occurred, err: %s", err)
	if err != nil {
		t.Error("Deletion failed")
		return
	}
	// check group is deleted well
	err = nsxtsess.Get("policy/api/v1/infra/domains/default/groups/group-test-1", &res)
	if err == nil {
		t.Errorf("Expecting no group with that id")
		return
	}
}

func TestNsxtSession(t *testing.T) {
	var err error
	var session *NsxtSession
	// Get below details from user environment variables
	var NsxtManager = os.Getenv("NSXT_MANAGER_HOST")
	var NsxtUsername = os.Getenv("NSXT_USERNAME")
	var NsxtPassword = os.Getenv("NSXT_PASSWORD")
	if NsxtManager == "" {
		t.Fatalf("NSXT_MANAGER_HOST env variable must be set for this test")
	}

	if NsxtUsername == "" {
		t.Fatalf("NSXT_USERNAME env variable must be set for this test")
	}

	if NsxtPassword == "" {
		t.Fatalf("NSXT_PASSWORD env variable must be set for this test")
	}
	session, err = NewNsxtSession(NsxtManager, NsxtUsername, nil, true, SetPassword(NsxtPassword), SetInsecure(true))

	if err != nil {
		log.Printf("[DEBUG] Ignoring login error: %s", err)
	}
	testNsxtSession(t, session)
}
