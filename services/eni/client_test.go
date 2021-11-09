/*
 * Copyright 2021 Baidu, Inc.
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

package eni

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	ENI_CLIENT *Client
	region     string
)

type Conf struct {
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Endpoint string `json:"Endpoint"`
}

const (
	VPC_ID    string = "vpc-87n1d60i2vas"
	SUBNET_ID string = "sbn-cxwi5hmf8ugx"
)

func init() {
	log.SetLogHandler(log.STDERR)
	log.SetLogLevel(log.DEBUG)
	_, f, _, _ := runtime.Caller(0)
	// Get the directory of GOPATH, the config file should be under the directory.
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)
	ENI_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	region = confObj.Endpoint[4:6]

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

func getClientToken() string {
	return util.NewUUID()
}

func TestClient_CreateEni(t *testing.T) {
	args := &CreateEniArgs{
		Name:     "hzb_3",
		SubnetId: SUBNET_ID,
		SecurityGroupIds: []string{
			"g-eqhqsbs84yww",
		},
		PrivateIpSet: []PrivateIp{
			{
				Primary:          true,
				PrivateIpAddress: "192.168.0.54",
			},
		},
		Description: "go sdk test",
		ClientToken: getClientToken(),
	}
	result, err := ENI_CLIENT.CreateEni(args)
	ExpectEqual(t.Errorf, nil, err)
	EniId := result.EniId
	log.Debug(EniId)
}

func TestClient_UpdateEni(t *testing.T) {
	args := &UpdateEniArgs{
		EniId:       "eni-mmwvvbvfjch3",
		ClientToken: getClientToken(),
		Name:        "hzb_3_1",
		Description: "go sdk 2",
	}
	err := ENI_CLIENT.UpdateEni(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteEni(t *testing.T) {
	args := &DeleteEniArgs{
		EniId:       "eni-darynwwu5xk0",
		ClientToken: getClientToken(),
	}
	err := ENI_CLIENT.DeleteEni(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListEnis(t *testing.T) {
	args := &ListEniArgs{
		VpcId:      VPC_ID,
		InstanceId: "i-FodWXDUY",
		Name:       "eni",
	}
	res, err := ENI_CLIENT.ListEni(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(res)
	fmt.Println(string(r))
}

func TestClient_GetEniDetail(t *testing.T) {
	result, err := ENI_CLIENT.GetEniDetail("eni-mmwvvbvfjch3")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestClient_AddPrivateIp(t *testing.T) {
	args := &EniPrivateIpArgs{
		EniId:            "eni-mmwvvbvfjch3",
		ClientToken:      getClientToken(),
		PrivateIpAddress: "192.168.0.53",
	}
	result, err := ENI_CLIENT.AddPrivateIp(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestClient_DeletePrivateIp(t *testing.T) {
	args := &EniPrivateIpArgs{
		EniId:            "eni-mmwvvbvfjch3",
		ClientToken:      getClientToken(),
		PrivateIpAddress: "192.168.0.53",
	}
	err := ENI_CLIENT.DeletePrivateIp(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AttachEniInstance(t *testing.T) {
	args := &EniInstance{
		EniId:       "eni-pw4tqi2f9gvq",
		ClientToken: getClientToken(),
		InstanceId:  "i-KdT8ptiu",
	}
	err := ENI_CLIENT.AttachEniInstance(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DetachEniInstance(t *testing.T) {
	args := &EniInstance{
		EniId:       "eni-pw4tqi2f9gvq",
		ClientToken: getClientToken(),
		InstanceId:  "i-KdT8ptiu",
	}
	err := ENI_CLIENT.DetachEniInstance(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindEniPublicIp(t *testing.T) {
	args := &BindEniPublicIpArgs{
		EniId:            "eni-mmwvvbvfjch3",
		ClientToken:      getClientToken(),
		PrivateIpAddress: "192.168.0.54",
		PublicIpAddress:  "100.88.8.55",
	}
	err := ENI_CLIENT.BindEniPublicIp(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnBindEniPublicIp(t *testing.T) {
	args := &UnBindEniPublicIpArgs{
		EniId:           "eni-mmwvvbvfjch3",
		ClientToken:     getClientToken(),
		PublicIpAddress: "100.88.8.55",
	}
	err := ENI_CLIENT.UnBindEniPublicIp(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateEniSecurityGroup(t *testing.T) {
	args := &UpdateEniSecurityGroupArgs{
		EniId:       "eni-mmwvvbvfjch3",
		ClientToken: getClientToken(),
		SecurityGroupIds: []string{
			"g-eqhqsbs84yww",
			"g-8bfifey0s4h3",
		},
	}
	err := ENI_CLIENT.UpdateEniSecurityGroup(args)
	ExpectEqual(t.Errorf, nil, err)
}
