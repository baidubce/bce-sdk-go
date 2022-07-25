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

package cfw

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	CfwClient *Client
)

type Conf struct {
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Endpoint string `json:"Endpoint"`
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

func init() {
	log.SetLogHandler(log.STDERR)
	log.SetLogLevel(log.DEBUG)
	_, f, _, _ := runtime.Caller(0)
	f = filepath.Dir(f)
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)
	CfwClient, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
}

func TestClient_CreateCfw(t *testing.T) {
	args := &CreateCfwRequest{
		Name:        "cfw_1",
		Description: "desc",
		CfwRules: []CreateRule{
			{
				IpVersion:     4,
				Priority:      4,
				Protocol:      "TCP",
				Direction:     "in",
				SourceAddress: "192.168.0.4",
				DestAddress:   "192.168.0.5",
				SourcePort:    "80",
				DestPort:      "88",
				Action:        "allow",
			},
		},
	}
	result, err := CfwClient.CreateCfw(args)
	ExpectEqual(t.Errorf, nil, err)
	CfwId := result.CfwId
	log.Debug(CfwId)
}

func TestClient_ListCfw(t *testing.T) {
	args := &ListCfwArgs{}
	result, err := CfwClient.ListCfw(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_GetCfw(t *testing.T) {
	result, err := CfwClient.GetCfw("cfw-xxxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_UpdateCfw(t *testing.T) {
	args := &UpdateCfwRequest{
		Name:        "cfw_2",
		Description: "desc",
	}
	err := CfwClient.UpdateCfw("cfw-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteCfw(t *testing.T) {
	err := CfwClient.DeleteCfw("cfw-xxxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateCfwRule(t *testing.T) {
	args := &CreateCfwRuleRequest{
		CfwRules: []CreateRule{
			{
				IpVersion:     4,
				Priority:      5,
				Protocol:      "TCP",
				Direction:     "in",
				SourceAddress: "192.168.0.3",
				DestAddress:   "192.168.0.4",
				SourcePort:    "80",
				DestPort:      "88",
				Action:        "allow",
			},
		},
	}
	err := CfwClient.CreateCfwRule("cfw-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateCfwRule(t *testing.T) {
	args := &UpdateCfwRuleRequest{
		IpVersion:     4,
		Priority:      2,
		Protocol:      "TCP",
		Direction:     "in",
		SourceAddress: "192.168.0.1",
		DestAddress:   "192.168.0.2",
		SourcePort:    "80",
		DestPort:      "88",
		Action:        "allow",
	}
	err := CfwClient.UpdateCfwRule("cfw-xxxxxxxxxxxx", "cfwr-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteCfwRule(t *testing.T) {
	args := &DeleteCfwRuleRequest{
		CfwRuleIds: []string{
			"cfwr-xxxxxxxxxxxx",
		},
	}
	err := CfwClient.DeleteCfwRule("cfw-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListInstance(t *testing.T) {
	args := &ListInstanceRequest{
		InstanceType: "csn",
	}
	result, err := CfwClient.ListInstance(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_BindCfw(t *testing.T) {
	args := &BindCfwRequest{
		InstanceType: "csn",
		Instances: []CfwBind{
			{
				Region:     "bj",
				InstanceId: "csn-xxxxxxxxxxxx",
				MemberId:   "vpc-xxxxxxxxxxxx",
			},
		},
	}
	err := CfwClient.BindCfw("cfw-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnbindCfw(t *testing.T) {
	args := &UnbindCfwRequest{
		InstanceType: "csn",
		Instances: []CfwBind{
			{
				Region:     "bj",
				InstanceId: "csn-xxxxxxxxxxxx",
				MemberId:   "vpc-xxxxxxxxxxxx",
			},
		},
	}
	err := CfwClient.UnbindCfw("cfw-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_EnableCfw(t *testing.T) {
	args := &EnableCfwRequest{
		InstanceId: "csn-xxxxxxxxxxxx",
		MemberId:   "vpc-xxxxxxxxxxxx",
	}
	err := CfwClient.EnableCfw("cfw-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DisableCfw(t *testing.T) {
	args := &DisableCfwRequest{
		InstanceId: "csn-xxxxxxxxxxxx",
		MemberId:   "vpc-xxxxxxxxxxxx",
	}
	err := CfwClient.DisableCfw("cfw-xxxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
}
