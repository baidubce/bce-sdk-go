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

package csn

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	CsnClient *Client
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
	CsnClient, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
}

func TestClient_CreateCsn(t *testing.T) {
	desc := "desc"
	args := &CreateCsnRequest{
		Name:        "csn_api_1",
		Description: &desc,
	}

	result, err := CsnClient.CreateCsn(args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
	CsnId := result.CsnId
	log.Debug(CsnId)
}

func TestClient_UpdateCsn(t *testing.T) {
	name := "csn_api_2"
	args := &UpdateCsnRequest{
		Name: &name,
	}

	err := CsnClient.UpdateCsn("csn-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteCsn(t *testing.T) {
	err := CsnClient.DeleteCsn("csn-xxxxxxxxxxx", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListCsn(t *testing.T) {
	args := &ListCsnArgs{}
	result, err := CsnClient.ListCsn(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_GetCsn(t *testing.T) {
	result, err := CsnClient.GetCsn("csn-xxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_ListInstance(t *testing.T) {
	args := &ListInstanceArgs{}
	result, err := CsnClient.ListInstance("csn-xxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_AttachInstance(t *testing.T) {
	args := &AttachInstanceRequest{
		InstanceId:     "vpc-xxxxxxxxxxx",
		InstanceType:   "vpc",
		InstanceRegion: "gz",
	}
	err := CsnClient.AttachInstance("csn-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DetachInstance(t *testing.T) {
	args := &DetachInstanceRequest{
		InstanceId:     "vpc-xxxxxxxxxxx",
		InstanceType:   "vpc",
		InstanceRegion: "gz",
	}
	err := CsnClient.DetachInstance("csn-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListRouteTable(t *testing.T) {
	args := &ListRouteTableArgs{}
	result, err := CsnClient.ListRouteTable("csn-xxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_CreateRouteRule(t *testing.T) {
	args := &CreateRouteRuleRequest{
		AttachId:    "tgwAttach-xxxxxxxxxxx",
		DestAddress: "0.0.0.0/0",
		RouteType:   "custom",
	}
	err := CsnClient.CreateRouteRule("csnRt-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListRouteRuleArgs(t *testing.T) {
	args := &ListRouteRuleArgs{}
	result, err := CsnClient.ListRouteRule("csnRt-xxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_DeleteRouteRule(t *testing.T) {
	err := CsnClient.DeleteRouteRule("csnRt-xxxxxxxxxxx",
		"xxxxxxxxxxx-32ac-4949-a3ca-xxxxxxxxxxx", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreatePropagation(t *testing.T) {
	desc := "desc"
	args := &CreatePropagationRequest{
		AttachId:    "tgwAttach-xxxxxxxxxxx",
		Description: &desc,
	}
	err := CsnClient.CreatePropagation("csnRt-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListPropagation(t *testing.T) {
	result, err := CsnClient.ListPropagation("csnRt-xxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_DeletePropagation(t *testing.T) {
	err := CsnClient.DeletePropagation("csnRt-xxxxxxxxxxx",
		"tgwAttach-xxxxxxxxxxx", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAssociation(t *testing.T) {
	desc := "desc"
	args := &CreateAssociationRequest{
		AttachId:    "tgwAttach-xxxxxxxxxxx",
		Description: &desc,
	}
	err := CsnClient.CreateAssociation("csnRt-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListAssociation(t *testing.T) {
	result, err := CsnClient.ListAssociation("csnRt-xxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_DeleteAssociation(t *testing.T) {
	err := CsnClient.DeleteAssociation("csnRt-xxxxxxxxxxx",
		"tgwAttach-xxxxxxxxxxx", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateCsnBp(t *testing.T) {
	instanceType := "center"
	args := &CreateCsnBpRequest{
		Name:          "csnBp_api_1",
		Bandwidth:     100,
		InterworkType: &instanceType,
		GeographicA:   "China",
		GeographicB:   "China",
		Billing: Billing{
			PaymentTiming: "Postpaid",
		},
	}
	result, err := CsnClient.CreateCsnBp(args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_UpdateCsnBp(t *testing.T) {
	args := &UpdateCsnBpRequest{
		Name: "csnBp_api_2",
	}
	err := CsnClient.UpdateCsnBp("csnBp-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteCsnBp(t *testing.T) {
	err := CsnClient.DeleteCsnBp("csnBp-xxxxxxxxxxx", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListCsnBp(t *testing.T) {
	args := &ListCsnBpArgs{}
	result, err := CsnClient.ListCsnBp(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_GetCsnBp(t *testing.T) {
	result, err := CsnClient.GetCsnBp("csnBp-xxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_ResizeCsnBp(t *testing.T) {
	args := &ResizeCsnBpRequest{
		Bandwidth: 50,
	}
	err := CsnClient.ResizeCsnBp("csnBp-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindCsnBp(t *testing.T) {
	args := &BindCsnBpRequest{
		CsnId: "csn-xxxxxxxxxxx",
	}
	err := CsnClient.BindCsnBp("csnBp-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnbindCsnBpRequest(t *testing.T) {
	args := &UnbindCsnBpRequest{
		CsnId: "csn-xxxxxxxxxxx",
	}
	err := CsnClient.UnbindCsnBp("csnBp-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateCsnBpLimit(t *testing.T) {
	args := &CreateCsnBpLimitRequest{
		LocalRegion: "bj",
		PeerRegion:  "gz",
		Bandwidth:   10,
	}
	err := CsnClient.CreateCsnBpLimit("csnBp-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListCsnBpLimit(t *testing.T) {
	result, err := CsnClient.ListCsnBpLimit("csnBp-xxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_DeleteCsnBpLimit(t *testing.T) {
	args := &DeleteCsnBpLimitRequest{
		LocalRegion: "bj",
		PeerRegion:  "gz",
	}
	err := CsnClient.DeleteCsnBpLimit("csnBp-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateCsnBpLimit(t *testing.T) {
	args := &UpdateCsnBpLimitRequest{
		LocalRegion: "bj",
		PeerRegion:  "gz",
		Bandwidth:   20,
	}
	err := CsnClient.UpdateCsnBpLimit("csnBp-xxxxxxxxxxx", args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListCsnBpLimitByCsnId(t *testing.T) {
	result, err := CsnClient.ListCsnBpLimitByCsnId("csn-xxxxxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_ListTgw(t *testing.T) {
	args := &ListTgwArgs{}
	result, err := CsnClient.ListTgw("csn-xxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_UpdateTgw(t *testing.T) {
	name := "tgw_1"
	desc := "desc"
	args := &UpdateTgwRequest{
		Name:        &name,
		Description: &desc,
	}
	err := CsnClient.UpdateTgw("csn-xxxxxxxxxxx", "tgw-xxxxxxxxxxx",
		args, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListTgwRuleArgs(t *testing.T) {
	args := &ListTgwRuleArgs{}
	result, err := CsnClient.ListTgwRule("csn-xxxxxxxxxxx", "tgw-xxxxxxxxxxx", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func getClientToken() string {
	return util.NewUUID()
}
