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

package resmanager

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

var resClient *Client

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given: %+v\n", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)
	resClient, _ = NewClientWithEndpoint(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.DEBUG)
}

func NewClientWithEndpoint(ak, sk, endpoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
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

func TestGetResGroupBatch(t *testing.T) {
	args := &ResGroupDetailRequest{
		ResourceBrief: []ResourceBrief{
			{
				ResourceId:     "weaxsey.org",
				ResourceType:   "CDN",
				ResourceRegion: "global",
			}, {
				ResourceId:     "vpc-vbgiiq5bmz2w",
				ResourceType:   "NETWORK",
				ResourceRegion: "bj",
			}, {
				ResourceId:     "vpc-zz1w1i8m1j5d",
				ResourceType:   "NETWORK",
				ResourceRegion: "bj",
			},
		},
	}
	res, err := resClient.GetResGroupBatch(args)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf(string(jsonRes))
}

func TestRemoveResourceFromGroup(t *testing.T) {
	args := &BindResourceToGroupArgs{
		Bindings: []Binding{
			{
				ResourceId:     "weaxsey.org",
				ResourceType:   "CDN",
				ResourceRegion: "global",
				GroupId:        "RESG-kvst7Bqqygx",
			},
		},
	}
	err := resClient.RemoveResourceFromGroup(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBindResourceToGroup(t *testing.T) {
	args := &BindResourceToGroupArgs{
		Bindings: []Binding{
			{
				ResourceId:     "weaxsey.org",
				ResourceType:   "CDN",
				ResourceRegion: "global",
				GroupId:        "RESG-kvst7Bqqygx",
			},
		},
	}
	res, err := resClient.BindResourceToGroup(args)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf(string(jsonRes))
}

func TestChangeResourceGroup(t *testing.T) {
	args := &ChangeResourceGroupArgs{
		MoveResModels: []MoveResModel{
			{
				TargetGroupId: "RESG-a088c216",
				OldGroupResInfo: OldGroupResInfo{
					ResourceId:     "weaxsey.org",
					ResourceType:   "CDN",
					ResourceRegion: "global",
					GroupId:        "RESG-kvst7Bqqygx",
				},
			},
		},
	}
	res, err := resClient.ChangeResourceGroup(args)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf(string(jsonRes))
}

func TestQueryGroupList(t *testing.T) {
	var args = "test"
	res, err := resClient.QueryGroupList(args)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf(string(jsonRes))
}

func TestQueryFullListArgsTooLarge(t *testing.T) {
	var groupId = "RESG-kvst7Bqqygx"
	var id = "weaxsey.org"
	var types = "CDN"
	var regions = "global,bj"
	var name = "weaxsey.org"
	var isCurrent = false
	var tags = "默认项目:"
	var pageNo = 1
	var pageSize = 10
	res, err := resClient.getGroupRes(groupId, id, types, regions, name, isCurrent, tags, pageNo, pageSize)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf("response jsonRes:%s", string(jsonRes))
}
