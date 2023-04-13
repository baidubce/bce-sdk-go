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

package havip

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
	HAVIP_CLIENT *Client
	region       string
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
	//for i := 0; i < 7; i++ {
	//	f = filepath.Dir(f)
	//}
	f = filepath.Dir(f)
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	_ = decoder.Decode(confObj)
	HAVIP_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

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

func TestClient_CreateHaVip(t *testing.T) {
	args := &CreateHaVipArgs{
		Name:             "havipGoSdkTest",
		Description:      "go sdk test",
		SubnetId:         "sbn-mnnhbyd2tbqr",
		PrivateIpAddress: "192.168.1.3",
		ClientToken:      getClientToken(),
	}
	result, err := HAVIP_CLIENT.CreateHaVip(args)
	ExpectEqual(t.Errorf, nil, err)
	HaVipId := result.HaVipId
	log.Debug(HaVipId)
}

func TestClient_ListHaVips(t *testing.T) {
	args := &ListHaVipArgs{}
	res, err := HAVIP_CLIENT.ListHaVip(args)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(res)
	fmt.Println(string(r))
}

func TestClient_GetHaVipDetail(t *testing.T) {
	result, err := HAVIP_CLIENT.GetHaVipDetail("havip-ied0wq4fs8va")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestClient_UpdateHaVip(t *testing.T) {
	args := &UpdateHaVipArgs{
		HaVipId:     "havip-yxz7bx3tskqs",
		ClientToken: getClientToken(),
		Name:        "havipGoSdkTest2",
	}
	err := HAVIP_CLIENT.UpdateHaVip(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteHaVip(t *testing.T) {
	args := &DeleteHaVipArgs{
		HaVipId:     "havip-qzctfgy0uadj",
		ClientToken: getClientToken(),
	}
	err := HAVIP_CLIENT.DeleteHaVip(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_HaVipAttachInstance(t *testing.T) {
	args := &HaVipInstanceArgs{
		HaVipId:     "havip-yxz7bx3tskqs",
		ClientToken: getClientToken(),
		InstanceIds: []string{
			"eni-xgg3pfhw384n",
		},
		InstanceType: "ENI",
	}
	err := HAVIP_CLIENT.HaVipAttachInstance(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_HaVipDetachInstance(t *testing.T) {
	args := &HaVipInstanceArgs{
		HaVipId:     "havip-yxz7bx3tskqs",
		ClientToken: getClientToken(),
		InstanceIds: []string{
			"eni-xgg3pfhw384n",
		},
		InstanceType: "ENI",
	}
	err := HAVIP_CLIENT.HaVipDetachInstance(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_HaVipBindPublicIp(t *testing.T) {
	args := &HaVipBindPublicIpArgs{
		HaVipId:         "havip-ied0wq4fs8va",
		ClientToken:     getClientToken(),
		PublicIpAddress: "110.242.73.217",
	}
	err := HAVIP_CLIENT.HaVipBindPublicIp(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_HaVipUnbindPublicIp(t *testing.T) {
	args := &HaVipUnbindPublicIpArgs{
		HaVipId:     "havip-ied0wq4fs8va",
		ClientToken: getClientToken(),
	}
	err := HAVIP_CLIENT.HaVipUnbindPublicIp(args)
	ExpectEqual(t.Errorf, nil, err)
}
