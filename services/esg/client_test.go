/*
 * Copyright  Baidu, Inc.
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

package esg

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	ESG_CLIENT *Client

	region string
)

type Conf struct {
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Endpoint string `json:"Endpoint"`
}

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
	ESG_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

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

func TestClient_CreateEsg(t *testing.T) {
	args := &CreateEsgArgs{
		Name: "esgGoSdkTest",
		Rules: []EnterpriseSecurityGroupRule{
			{
				Action:    "deny",
				Direction: "ingress",
				Ethertype: "IPv4",
				PortRange: "1-65535",
				Priority:  1000,
				Protocol:  "udp",
				Remark:    "go sdk test",
				SourceIp:  "all",
			},
			{
				Action:    "allow",
				Direction: "ingress",
				Ethertype: "IPv4",
				PortRange: "1-65535",
				Priority:  1000,
				Protocol:  "icmp",
				Remark:    "go sdk test",
				SourceIp:  "all",
			},
		},
		Desc:        "go sdk test",
		ClientToken: getClientToken(),
		Tags: []model.TagModel{
			{
				TagKey:   "test",
				TagValue: "",
			},
		},
	}
	result, err := ESG_CLIENT.CreateEsg(args)
	ExpectEqual(t.Errorf, nil, err)
	EnterpriseSecurityGroupId := result.EnterpriseSecurityGroupId
	log.Debug(EnterpriseSecurityGroupId)
}

func TestClient_ListEsgs(t *testing.T) {
	args := &ListEsgArgs{}
	res, err := ESG_CLIENT.ListEsg(args)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(res)
	fmt.Println(string(r))
}

func TestClient_DeleteEsg(t *testing.T) {
	args := &DeleteEsgArgs{
		EnterpriseSecurityGroupId: "esg-s91awqpw73un",
		ClientToken:               getClientToken(),
	}
	err := ESG_CLIENT.DeleteEsg(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateEsgRules(t *testing.T) {
	args := &CreateEsgRuleArgs{
		Rules: []EnterpriseSecurityGroupRule{
			{
				Action:    "deny",
				Direction: "ingress",
				Ethertype: "IPv4",
				PortRange: "1-65535",
				Priority:  1000,
				Protocol:  "udp",
				Remark:    "go sdk test",
				SourceIp:  "all",
			},
		},
		EnterpriseSecurityGroupId: "esg-v99qnxx7uh83",
		ClientToken:               getClientToken(),
	}
	err := ESG_CLIENT.CreateEsgRules(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteEsgRule(t *testing.T) {
	args := &DeleteEsgRuleArgs{
		EnterpriseSecurityGroupRuleId: "esgr-ak7b51zzgptc",
	}
	err := ESG_CLIENT.DeleteEsgRule(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateEsgRule(t *testing.T) {
	args := &UpdateEsgRuleArgs{
		Priority:                      900,
		Remark:                        "go sdk test update",
		ClientToken:                   getClientToken(),
		EnterpriseSecurityGroupRuleId: "esgr-ahm3xxi11s20",
	}
	err := ESG_CLIENT.UpdateEsgRule(args)
	ExpectEqual(t.Errorf, nil, err)
}
