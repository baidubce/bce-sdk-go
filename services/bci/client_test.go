/*
 * Copyright 2023 Baidu, Inc.
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
package bci

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BCI_CLIENT    *Client
	BciInstanceId string
)

type Conf struct {
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Endpoint string `json:"Endpoint"`
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

	BCI_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateInstance(t *testing.T) {
	args := &CreateInstanceArgs{
		Name:               "cym-bci-test",
		ZoneName:           "zoneC",
		SubnetIds:          []string{"sbn-g463qx0aqu7q"},
		SecurityGroupIds:   []string{"g-59gf44p4jmwe"},
		RestartPolicy:      "Always",
		AutoCreateEip:      false,
		EipRouteType:       "BGP",
		EipBandwidthInMbps: 150,
		EipPaymentTiming:   "Postpaid",
		EipBillingMethod:   "ByTraffic",
		Tags: []Tag{
			{
				TagKey:   "mytag",
				TagValue: "serverglen",
			},
		},
		ImageRegistryCredentials: []ImageRegistryCredential{
			{
				Server:   "docker.io/wywcoder",
				UserName: "wywcoder",
				Password: "Qaz123456",
			},
		},
		Containers: []Container{
			{
				Name:            "container01",
				Image:           "registry.baidubce.com/bci-zjm-public/ubuntu:18.04",
				Memory:          1,
				CPU:             0.5,
				WorkingDir:      "",
				ImagePullPolicy: "IfNotPresent",
				Commands: []string{
					"/bin/sh",
					"-c",
					"sleep 30000 && exit 0",
				},
				VolumeMounts: []VolumeMount{
					{
						MountPath: "/usr/local/share",
						ReadOnly:  false,
						Name:      "emptydir",
						Type:      "EmptyDir",
					},
					{
						MountPath: "/config",
						ReadOnly:  false,
						Name:      "configfile",
						Type:      "ConfigFile",
					},
				},
				Ports: []Port{
					{
						Port:     8099,
						Protocol: "TCP",
					},
				},
				EnvironmentVars: []Environment{
					{
						Key:   "java",
						Value: "/usr/local/jre",
					},
				},
			}},
		Volume: &Volume{
			Nfs: []NfsVolume{},
			EmptyDir: []EmptyDirVolume{
				{
					Name: "emptydir",
				},
			},
			ConfigFile: []ConfigFileVolume{
				{
					Name: "configfile",
					ConfigFiles: []ConfigFileDetail{
						{
							Path: "podconfig",
							File: "filenxx",
						},
					},
				}},
		},
	}
	// res, _ := json.Marshal(args)
	// fmt.Println(string(res))
	result, err := BCI_CLIENT.CreateInstance(args)
	fmt.Printf("CreateInstance result: %+v, err: %+v \n", result, err)
	ExpectEqual(t.Errorf, nil, err)
	BciInstanceId = result.InstanceId
}

func TestClient_ListInstance(t *testing.T) {
	args := &ListInstanceArgs{
		MaxKeys: 5,
	}
	result, err := BCI_CLIENT.ListInstances(args)
	fmt.Printf("ListInstances result: %+v, err: %+v \n", result, err)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetInstance(t *testing.T) {
	args := &GetInstanceArgs{
		InstanceId: BciInstanceId,
	}
	result, err := BCI_CLIENT.GetInstance(args)
	fmt.Printf("GetInstance err: %+v, result: %+v \n", err, *result.Instance)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteInstance(t *testing.T) {
	args := &DeleteInstanceArgs{
		InstanceId:         BciInstanceId,
		RelatedReleaseFlag: true,
	}
	err := BCI_CLIENT.DeleteInstance(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BatchDeleteInstance(t *testing.T) {
	args := &BatchDeleteInstanceArgs{
		InstanceIds:        []string{BciInstanceId},
		RelatedReleaseFlag: true,
	}
	err := BCI_CLIENT.BatchDeleteInstance(args)
	ExpectEqual(t.Errorf, nil, err)
}
