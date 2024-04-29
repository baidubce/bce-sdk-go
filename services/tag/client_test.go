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

package tag

import (
	"encoding/json"
	"fmt"
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

var tagClient *Client

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
	tagClient, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.DEBUG)
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

func TestCreateTags(t *testing.T) {
	args := &TagsRequest{
		Tags: []Tag{
			{
				TagKey:   "go_sdk_key",
				TagValue: "go_sdk_val",
			}, {
				TagKey:   "go_sdk_key",
				TagValue: "go_sdk_val1",
			},
		},
	}
	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	err := tagClient.CreateTags(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteTags(t *testing.T) {
	args := &TagsRequest{
		Tags: []Tag{
			{
				TagKey:   "go_sdk_key",
				TagValue: "go_sdk_val-del",
			},
			{
				TagKey:   "go_sdk_key",
				TagValue: "go_sdk_val-del1",
			},
		},
	}
	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	err := tagClient.CreateTags(args)
	ExpectEqual(t.Errorf, err, nil)

	t.Logf("request args:%s", string(jsonArgs))
	err = tagClient.DeleteTags(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestTagsResources(t *testing.T) {
	tagKey := "go_sdk_key"
	tagValue := "go_sdk_val"
	region := "global"
	resourceType := "APIGW"
	res, err := tagClient.TagsResources(tagKey, tagValue, region, resourceType)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf("response jsonRes:%s", string(jsonRes))
}

func TestUserList(t *testing.T) {
	tagKey := "go_sdk_key"
	tagValue := "go_sdk_val"
	res, err := tagClient.UserTagList(tagKey, tagValue)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf("response jsonRes:%s", string(jsonRes))
}

func TestCreateAssociationsByTag(t *testing.T) {
	args := &CreateAssociationsByTagRequest{
		TagKey:      "go_sdk_key",
		TagValue:    "go_sdk_val",
		ServiceType: "APIGW",
		Resource: []Resource{
			{
				ResourceId:  "58426",
				ServiceType: "APIGW",
				Region:      "global",
			},
			{
				ResourceId:  "58428",
				ServiceType: "APIGW",
				Region:      "global",
			},
		},
	}
	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	err := tagClient.CreateAssociationsByTag(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteAssociationsByTag(t *testing.T) {
	args := &DeleteAssociationsByTagRequest{
		TagKey:      "go_sdk_key",
		TagValue:    "go_sdk_val",
		ServiceType: "APIGW",
		Resource: []Resource{
			{
				ResourceId:  "58426",
				ServiceType: "APIGW",
				Region:      "global",
			},
			{
				ResourceId:  "58428",
				ServiceType: "APIGW",
				Region:      "global",
			},
		},
	}
	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	err := tagClient.DeleteAssociationsByTag(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateAndAssign(t *testing.T) {
	args := &CreateAndAssignTagRequest{
		ResourceWithTag: []ResourceWithTag{
			{
				ResourceId:   "58426",
				ResourceUuid: "GWGP-w79NnAxhhoj",
				ServiceType:  "APIGW",
				Region:       "global",
				Tags: []Tag{
					{
						TagKey:   "go_sdk_key",
						TagValue: "go_sdk_val",
					},
					{
						TagKey:   "go_sdk_key1",
						TagValue: "go_sdk_val",
					},
				},
			},
		},
	}
	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	err := tagClient.CreateAndAssign(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteTagAssociation(t *testing.T) {

	args := &DeleteTagAssociationRequest{
		Resource: &Resource{
			ResourceId:  "58426",
			ServiceType: "APIGW",
			Region:      "global",
		},
	}

	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	err := tagClient.DeleteTagAssociation(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestQueryFullList(t *testing.T) {
	var strongAssociation = true
	args := &FullTagListRequest{
		TagKey:   "go_sdk_key",
		TagValue: "go_sdk_val",
		Regions: []string{
			"global",
		},
		ServiceTypes: []string{
			"APIGW",
		},
		ResourceIds: []string{},
	}
	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	res, err := tagClient.QueryFullList(strongAssociation, args)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf("response jsonRes:%s", string(jsonRes))
}

func TestQueryFullListArgsTooLarge(t *testing.T) {
	var strongAssociation = true
	var resourceIds []string
	for i := 0; i < 201; i++ {
		resourceIds = append(resourceIds, fmt.Sprintf("bcm5-qasandbox.sys-qa.com%d", i))
	}

	args := &FullTagListRequest{
		TagKey:   "test_zmq_1213",
		TagValue: "value_zmq",
		Regions: []string{
			"bj",
		},
		ServiceTypes: []string{
			"CDN",
		},
		ResourceIds: resourceIds,
	}
	jsonArgs, _ := json.Marshal(args)
	t.Logf("request args:%s", string(jsonArgs))
	res, err := tagClient.QueryFullList(strongAssociation, args)
	ExpectEqual(t.Errorf, err, nil)
	jsonRes, _ := json.Marshal(res)
	t.Logf("response jsonRes:%s", string(jsonRes))
}
