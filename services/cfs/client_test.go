/*
 * Copyright 2022 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package cfs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CFS_CLIENT *Client
	CFS_ID     string
	MOUNT_ID   string
	USER_ID    string

	// set these values before start test
	VPC_TEST_ID    = ""
	SUBNET_TEST_ID = ""
	TEST_PROTOCOL  = "nfs"
	TEST_ZONE      = "zoneA"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK        string
	SK        string
	USER_ID   string
	VPC_ID    string
	SUBNET_ID string
	Endpoint  string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	CFS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	VPC_TEST_ID = confObj.VPC_ID
	SUBNET_TEST_ID = confObj.SUBNET_ID
	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_CreateFS(t *testing.T) {
	createArgs := &CreateFSArgs{
		ClientToken: getClientToken(),
		Name:        "sdkCFS",
		VpcID:       VPC_TEST_ID,
		Protocol:    TEST_PROTOCOL,
		Zone:        TEST_ZONE,
	}

	createResult, err := CFS_CLIENT.CreateFS(createArgs)
	ExpectEqual(t.Errorf, nil, err)

	CFS_ID = createResult.FSID
}

func TestClient_UpdateFS(t *testing.T) {
	updateArgs := &UpdateFSArgs{
		FSID:   CFS_ID,
		FSName: "sdCFS_newname",
	}

	err := CFS_CLIENT.UpdateFS(updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeFS(t *testing.T) {
	describeArgs := &DescribeFSArgs{
		UserId: USER_ID,
	}
	res, err := CFS_CLIENT.DescribeFS(describeArgs)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(time.Duration(1) * time.Second)
}

func TestClient_DescribeFSWithFS(t *testing.T) {
	describeArgs := &DescribeFSArgs{
		FSID: CFS_ID,
	}
	res, err := CFS_CLIENT.DescribeFS(describeArgs)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(time.Duration(1) * time.Second)
}

func TestClient_CreateMountTarget(t *testing.T) {
	createArgs := &CreateMountTargetArgs{
		FSID:     CFS_ID,
		SubnetId: SUBNET_TEST_ID,
		VpcID:    VPC_TEST_ID,
	}

	createResult, err := CFS_CLIENT.CreateMountTarget(createArgs)
	fmt.Print(createResult)
	ExpectEqual(t.Errorf, nil, err)
	MOUNT_ID = createResult.MountID
}

func TestClient_DescribeMountTarget(t *testing.T) {
	describeArgs := &DescribeMountTargetArgs{
		FSID: CFS_ID,
	}
	res, err := CFS_CLIENT.DescribeMountTarget(describeArgs)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(time.Duration(1) * time.Second)
}

func TestClient_DropMountTarget(t *testing.T) {
	dropArgs := &DropMountTargetArgs{
		FSID:    CFS_ID,
		MountId: MOUNT_ID,
	}
	err := CFS_CLIENT.DropMountTarget(dropArgs)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(time.Duration(5) * time.Second)
}

func TestClient_DropFS(t *testing.T) {
	dropArgs := &DropFSArgs{
		FSID: CFS_ID,
	}
	err := CFS_CLIENT.DropFS(dropArgs)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(time.Duration(1) * time.Second)
}

func getClientToken() string {
	return util.NewUUID()
}
