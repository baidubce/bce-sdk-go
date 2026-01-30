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

package et

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	EtClient *Client
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
	EtClient, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
}

// TestClient_GetEtChannel tests GetEtChannel method of EtClient
func TestClient_GetEtChannel(t *testing.T) {
	args := &GetEtChannelArgs{
		ClientToken: getClientToken(),
		EtId:        "dcphy-xxxxxxxxxxxx",
	}
	result, err := EtClient.GetEtChannel(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

// TestClient_RecommitEtChannel 测试客户端的RecommitEtChannel方法
func TestClient_RecommitEtChannel(t *testing.T) {
	args := &RecommitEtChannelArgs{
		ClientToken: getClientToken(),
		EtId:        "dcphy-xxxxxxxxxxxx",
		EtChannelId: "dedicatedconn-xxxxxxxxxxxx",
		Result: RecommitEtChannelResult{ // recommit et channel result
			AuthorizedUsers:     []string{"xxxxxxxxxxxxxxxxxx"}, // authorized users
			Description:         "test Description",             // description
			BaiduAddress:        "Your BaiduAddress",            // baidu address
			Name:                "test Name",                    // name
			Networks:            []string{"Your Networks"},      // networks
			CustomerAddress:     "Your CustomerAddress",         // customer address
			RouteType:           "Your RouteType",               // route type
			VlanId:              "1",                            // vlan id
			Status:              "Your Status",                  // status
			EnableIpv6:          0,                              // enable ipv6
			BaiduIpv6Address:    "Your BaiduIpv6Address",        // baidu ipv6 address
			Ipv6Networks:        []string{"Your Ipv6Networks"},  // ipv6 networks
			CustomerIpv6Address: "Your CustomerIpv6Address",     // customer ipv6 address
		},
	}
	err := EtClient.RecommitEtChannel(args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_UpdateEtChannel 测试客户端的更新ET通道函数
func TestClient_UpdateEtChannel(t *testing.T) {
	args := &UpdateEtChannelArgs{
		ClientToken: getClientToken(),
		EtId:        "dcphy-xxxxxxxxxxxx",
		EtChannelId: "dedicatedconn-xxxxxxxxxxxx",
		Result: UpdateEtChannelResult{ // update et channel result
			Name:        "testname",       // name
			Description: "testdecription", // description
		},
	}
	err := EtClient.UpdateEtChannel(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteEtChannel(t *testing.T) {
	args := &DeleteEtChannelArgs{
		ClientToken: getClientToken(),
		EtId:        "dcphy-xxxxxxxxxxxx",
		EtChannelId: "dedicatedconn-xxxxxxxxxxxx",
	}
	err := EtClient.DeleteEtChannel(args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_EnableEtChannelIPv6 测试函数
func TestClient_EnableEtChannelIPv6(t *testing.T) {
	args := &EnableEtChannelIPv6Args{
		ClientToken: getClientToken(),
		EtId:        "dcphy-xxxxxxxxxxxx",
		EtChannelId: "dedicatedconn-xxxxxxxxxxxx",
		Result: EnableEtChannelIPv6Result{ // enable et channel ipv6 result
			BaiduIpv6Address:    "2001:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:0001", // baidu ipv6 address
			Ipv6Networks:        []string{"2001:xxxx:xxxx:xxxx::/64"},      // ipv6 networks
			CustomerIpv6Address: "2001:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:0002", // customer ipv6 address
		},
	}
	err := EtClient.EnableEtChannelIPv6(args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_DisableEtChannelIPv6 测试函数
func TestClient_DisableEtChannelIPv6(t *testing.T) {
	args := &DisableEtChannelIPv6Args{
		ClientToken: getClientToken(),
		EtId:        "dcphy-tm25m1reihvw",
		EtChannelId: "dedicatedconn-ybffmxnpygcx",
	}
	err := EtClient.DisableEtChannelIPv6(args)
	ExpectEqual(t.Errorf, nil, err)
}

// getClientToken 函数返回一个长度为32的字符串作为客户端令牌。
func getClientToken() string {
	return util.NewUUID()
}

// TestClient_CreateEtDcphy 测试Client_CreateEtDcphy函数，创建一个ET DCphy实例
// 参数：t *testing.T - 单元测试对象指针，用于输出错误信息
func TestClient_CreateEtDcphy(t *testing.T) {
	args := &CreateEtDcphyArgs{
		ClientToken: getClientToken(),
		Name:        "test_InitEt",
		Isp:         "ISP_CMCC",
		IntfType:    "1G",
		ApType:      "SINGLE",
		ApAddr:      "BJYZ",
		UserName:    "test",
		UserPhone:   "18266666666",
		UserEmail:   "18266666666@baidu.com",
		UserIdc:     "北京|市辖区|东城区|百度科技园K2",
	}

	r, err := EtClient.CreateEtDcphy(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Id) != 0)
}

// TestClient_UpdateEtDcphy 测试Client_UpdateEtDcphy方法，更新ET DCphy信息
// 参数：t *testing.T - 类型为*testing.T的指针，表示测试对象
// 返回值：nil - 无返回值
func TestClient_UpdateEtDcphy(t *testing.T) {
	args := &UpdateEtDcphyArgs{
		Name:        "test_Update",
		Description: "new",
		UserName:    "testUpdate",
		UserPhone:   "18266666667",
		UserEmail:   "18266666667@baidu.com",
	}

	err := EtClient.UpdateEtDcphy("dcphy-23451", args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_ListEtDcphy 测试Client_ListEtDcphy方法，用于获取ET DC/PHY列表
// 参数：t *testing.T - 类型为*testing.T的指针，表示测试对象
// 返回值：nil - 无返回值
func TestClient_ListEtDcphy(t *testing.T) {
	args := &ListEtDcphyArgs{
		Marker:  "your marker",
		MaxKeys: 10000,
		Status:  "established",
	}

	r, err := EtClient.ListEtDcphy(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Ets) != 0)
}

func TestClient_ListEtDcphyDetail(t *testing.T) {
	r, err := EtClient.ListEtDcphyDetail("dcphy-23451")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Name) != 0)
}

// TestClient_CreateEtChannel 测试Client_CreateEtChannel方法，创建一个ET通道。
// 参数：*testing.T - t，表示单元测试的对象，用于报错和判定。
// 返回值：nil - 无返回值，只是进行单元测试。
func TestClient_CreateEtChannel(t *testing.T) {
	args := &CreateEtChannelArgs{
		ClientToken:         getClientToken(),
		EtId:                "dcphy-234r5",
		Description:         "test",
		BaiduAddress:        "172.1.1.1/24",
		Name:                "testChannel",
		Networks:            []string{"192.168.0.0/16"},
		CustomerAddress:     "172.1.1.2/24",
		RouteType:           "static-route",
		VlanId:              100,
		EnableIpv6:          1,
		BaiduIpv6Address:    "2400:da00:e003:0:1eb:200::1/88",
		CustomerIpv6Address: "2400:da00:e003:0:0:200::1/88",
		Ipv6Networks:        []string{"2400:da00:e003:0:15f::/87"},
	}

	r, err := EtClient.CreateEtChannel(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Id) != 0)
}

func TestCreateEtGatewayRouteRule(t *testing.T) {
	req := &CreateEtChannelRouteRuleArgs{
		EtId:        "dcphy-tm25m1reihvw",
		EtChannelId: "dedicatedconn-ybffmxnpygcx",
		ClientToken: getClientToken(),
		DestAddress: "11.11.12.14/32",
		NextHopType: "etChannel",
		NextHopId:   "dedicatedconn-ybffmxnpygcx",
	}
	response, err := EtClient.CreateEtChannelRouteRule(req)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(response)
	fmt.Println(string(r))
}

func TestListEtChannelRouteRule(t *testing.T) {
	req := &ListEtChannelRouteRuleArgs{
		EtId:        "dcphy-tm25m1reihvw",
		EtChannelId: "dedicatedconn-ybffmxnpygcx",
	}
	response, err := EtClient.ListEtChannelRouteRule(req)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(response)
	fmt.Println(string(r))
}

func TestUpdateEtChannelRouteRule(t *testing.T) {
	req := &UpdateEtChannelRouteRuleArgs{
		ClientToken: getClientToken(),
		EtId:        "dcphy-tm25m1reihvw", // 专线ID
		EtChannelId: "dedicatedconn-ybffmxnpygcx",
		RouteRuleId: "dcrr-07a5967b-84a",
		Description: "test",
	}
	err := EtClient.UpdateEtChannelRouteRule(req)
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteEtChannelRouteRule(t *testing.T) {
	req := &DeleteEtChannelRouteRuleArgs{
		ClientToken: getClientToken(),
		EtId:        "dcphy-tm25m1reihvw",
		EtChannelId: "dedicatedconn-ybffmxnpygcx",
		RouteRuleId: "dcrr-6378ed8b-a2d",
	}
	err := EtClient.DeleteEtChannelRouteRule(req)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_ListEtDcphyWithTag 是一个用于测试EtClient的ListEtDcphy方法返回是否带有Tags的单元测试
func TestClient_ListEtDcphyWithTag(t *testing.T) {
	args := &ListEtDcphyArgs{
		Marker:  "",
		MaxKeys: 10000,
		Status:  "established",
	}

	r, err := EtClient.ListEtDcphy(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Ets[0].Tags) > 0)
}

// TestClient_CreateEtDcphyWithTag 测试Client_CreateEtDcphy函数，带有Tag标签
// 参数：t *testing.T - 单元测试对象指针，用于输出错误信息
func TestClient_CreateEtDcphyWithTag(t *testing.T) {
	args := &CreateEtDcphyArgs{
		ClientToken: getClientToken(),
		Name:        "test_InitEt",
		Isp:         "ISP_CMCC",
		IntfType:    "1G",
		ApType:      "SINGLE",
		ApAddr:      "BJYZ",
		UserName:    "test",
		UserPhone:   "18266666666",
		UserEmail:   "18266666666@baidu.com",
		UserIdc:     "北京|市辖区|东城区|百度科技园K2",
		Tags:        []Tag{{"tag-chongkuitai", "chongkuitai"}},
	}

	r, err := EtClient.CreateEtDcphy(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Id) != 0)
}

// TestClient_ListEtDcphyDetailWithTag 测试带有Tag标签的专线详情查询
func TestClient_ListEtDcphyDetailWithTag(t *testing.T) {
	r, err := EtClient.ListEtDcphyDetail("dcphy-3r2uaz4psic5")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Name) != 0)
}

// TestClient_CreateEtChannelWithTag 测试Client_CreateEtChannel方法，创建一个带有tag的ET通道。
// 参数：*testing.T - t，表示单元测试的对象，用于报错和判定。
// 返回值：nil - 无返回值，只是进行单元测试。
func TestClient_CreateEtChannelWithTag(t *testing.T) {
	args := &CreateEtChannelArgs{
		ClientToken:         getClientToken(),
		EtId:                "dcphy-axibreesn6af",
		Description:         "test",
		BaiduAddress:        "11.11.11.1/24",
		Name:                "testChannel",
		Networks:            []string{"192.168.0.0/16"},
		CustomerAddress:     "11.11.11.2/24",
		RouteType:           "static-route",
		VlanId:              56,
		EnableIpv6:          1,
		BaiduIpv6Address:    "2400:da00:e003:0:1eb:200::1/88",
		CustomerIpv6Address: "2400:da00:e003:0:1eb:201::1/88",
		Ipv6Networks:        []string{"2400:da00:e003:0:15f::/87"},
		Tags:                []Tag{{"tag-chongkuitai", "chongkuitai"}},
	}

	r, err := EtClient.CreateEtChannel(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(r.Id) != 0)
}

// TestClient_GetEtChannelWithTag 测试带有tag的ET通道查询
func TestClient_GetEtChannelWithTag(t *testing.T) {
	args := &GetEtChannelArgs{
		ClientToken: getClientToken(),
		EtId:        "dcphy-xxxxxxxxxxx",
		EtChannelId: "dedicatedconn-xxxxxxxxxxx",
	}
	result, err := EtClient.GetEtChannel(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	log.Debug(string(r))
}

func TestClient_CreateEtChannelBfd(t *testing.T) {
	createEtChannelBfdRequest := &CreateEtChannelBfdRequest{
		EtId:             util.PtrString("dcphy-xxxxxxxxxxx"),
		EtChannelId:      util.PtrString("dedicatedconn-xxxxxxxxxxx"),
		SendInterval:     util.PtrInt32(int32(300)),
		ReceivInterval:   util.PtrInt32(int32(300)),
		DetectMultiplier: util.PtrInt32(int32(7)),
	}
	err := EtClient.CreateEtChannelBfd(createEtChannelBfdRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteEtChannelBfd(t *testing.T) {
	deleteEtChannelBfdRequest := &DeleteEtChannelBfdRequest{
		EtId:        util.PtrString("dcphy-xxxxxxxxxxx"),
		EtChannelId: util.PtrString("dedicatedconn-xxxxxxxxxxx"),
	}
	err := EtClient.DeleteEtChannelBfd(deleteEtChannelBfdRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateEtChannelBfd(t *testing.T) {
	updateEtChannelBfdRequest := &UpdateEtChannelBfdRequest{
		EtId:             util.PtrString("dcphy-xxxxxxxxxxx"),
		EtChannelId:      util.PtrString("dedicatedconn-xxxxxxxxxxx"),
		SendInterval:     util.PtrInt32(int32(301)),
		ReceivInterval:   util.PtrInt32(int32(301)),
		DetectMultiplier: util.PtrInt32(int32(8)),
	}
	err := EtClient.UpdateEtChannelBfd(updateEtChannelBfdRequest)
	ExpectEqual(t.Errorf, nil, err)
}
