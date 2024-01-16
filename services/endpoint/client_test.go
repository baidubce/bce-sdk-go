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

package endpoint

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
	ENDPOINT_CLIENT *Client

	region string

	EndpointId string
)

type Conf struct {
	AK          string `json:"AK"`
	SK          string `json:"SK"`
	VPCEndpoint string `json:"Endpoint"`
}

const (
	PAYMENT_TIMING_POSTPAID PaymentTimingType = "Postpaid"

	VPC_ID    string = "vpc-q1hcnhf7nmve"
	SUBNET_ID string = "sbn-crqu2vxzj049"
	SERVICE   string = "77.uservice-a7f5795b.beijing.baidubce.com"
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
	ENDPOINT_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.VPCEndpoint)

	region = confObj.VPCEndpoint[4:6]

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

func TestClient_GetService(t *testing.T) {
	result, err := ENDPOINT_CLIENT.GetServices()
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestClient_CreateEndpoint(t *testing.T) {
	args := &CreateEndpointArgs{
		VpcId:       VPC_ID,
		Name:        "go-sdk-create-1",
		SubnetId:    SUBNET_ID,
		Service:     "77.uservice-a7f5795b.beijing.baidubce.com",
		Description: "go sdk test",
		Billing: &Billing{
			PaymentTiming: PAYMENT_TIMING_POSTPAID,
		},
		ClientToken: getClientToken(),
	}
	result, err := ENDPOINT_CLIENT.CreateEndpoint(args)
	ExpectEqual(t.Errorf, nil, err)
	EndpointId := result.Id
	log.Debug(EndpointId)
}

func TestClient_DeleteEndpoint(t *testing.T) {
	err := ENDPOINT_CLIENT.DeleteEndpoint("endpoint-f0f191f0", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateEndpoint(t *testing.T) {
	args := &UpdateEndpointArgs{
		ClientToken: getClientToken(),
		Name:        "go-sdk-2",
		Description: "go sdk 2",
	}
	err := ENDPOINT_CLIENT.UpdateEndpoint("endpoint-3c7e02cb", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListEndpoints(t *testing.T) {
	args := &ListEndpointArgs{
		VpcId:    VPC_ID,
		SubnetId: SUBNET_ID,
	}
	res, err := ENDPOINT_CLIENT.ListEndpoints(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(res)
	fmt.Println(string(r))
}

func TestClient_GetEndpointDetail(t *testing.T) {
	result, err := ENDPOINT_CLIENT.GetEndpointDetail("endpoint-3c7e02cb")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}
func TestClient_UpdateEndpointNSG(t *testing.T) {
	args := &UpdateEndpointNSGArgs{
		SecurityGroupIds: []string{"g-wmxijt06y5um"},
	}
	err := ENDPOINT_CLIENT.UpdateEndpointNormalSecurityGroup("endpoint-3c7e02cb", args)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateEndpointESG(t *testing.T) {
	args := &UpdateEndpointESGArgs{
		EnterpriseSecurityGroupIds: []string{"esg-nhwrebdqi4q2"},
	}
	err := ENDPOINT_CLIENT.UpdateEndpointEnterpriseSecurityGroup("endpoint-3c7e02cb", args)
	ExpectEqual(t.Errorf, nil, err)
}
