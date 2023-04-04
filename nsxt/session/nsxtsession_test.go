/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package nsxt

import (
	"testing"

	"github.com/golang/glog"
)

func testNsxtSession(t *testing.T, nsxtsess *NsxtSession) {
	//GET
	var res interface{}
	errInGet := nsxtsess.Get("policy/api/v1/infra/domains/default/groups", &res)
	if errInGet != nil || res == nil {
		t.Error("Group GET failed: ", errInGet)
		return
	}
	resp := res.(map[string]interface{})
	glog.Infof("resp: %+v", resp)

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
	resp = res.(map[string]interface{})
	glog.Infof("res for GET after PATCH *******: %+v", resp)
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
	glog.Infof("res for GET after PUT *******: %+v", resp)
	if resp["display_name"] != "updated_group_test_1" {
		t.Error("Expected group not fetched")
		return
	}

	// delete the group
	err = nsxtsess.Delete("policy/api/v1/infra/domains/default/groups/group-test-1")
	glog.Infof("err: %s", err)
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
	var NsxtManager = "10.191.153.238"
	var NsxtUsername = "admin"
	var NsxtPassword = "e6+DgNzYA5D*"
	session, err = NewNsxtSession(NsxtManager, NsxtUsername, nil, true, SetPassword(NsxtPassword), SetInsecure(true))

	if err != nil {
		glog.Infof("ignoring login err: %s", err)
	}
	testNsxtSession(t, session)
}
