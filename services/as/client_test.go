package as

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
	asClient *Client
	asConf   *Conf
)

type Conf struct {
	AK         string `json:"AK"`
	SK         string `json:"SK"`
	Endpoint   string `json:"Endpoint"`
	InstanceId string `json:"InstanceId"`
	UserId     string `json:"UserId"`
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
	asConf = &Conf{}
	_ = decoder.Decode(asConf)

	asClient, _ = NewClient(asConf.AK, asConf.SK, asConf.Endpoint)
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

func Test_GetAsGroupList(t *testing.T) {
	req := &ListAsGroupRequest{
		KeyWordType: "groupName",
		KeyWord:     "djw-test",
	}
	resp, err := asClient.ListAsGroup(req)
	fmt.Printf("err is %v\n", err)
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func TestClient_GetAsGroupDetail(t *testing.T) {
	req := &GetAsGroupRequest{
		GroupId: "asg-wqksXo95",
	}
	resp, err := asClient.GetAsGroup(req)
	fmt.Println(resp)
	ExpectEqual(t.Errorf, nil, err)
}

func Test_GetAsNodeList(t *testing.T) {
	req := &ListAsGroupRequest{
		GroupId: "asg-FKsD6xmT",
	}
	resp, err := asClient.ListAsNode(req)
	fmt.Printf("err is %v\n", err)
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func TestClient_IncreaseAs(t *testing.T) {
	req := &IncreaseAsGroupRequest{
		GroupId:   "asg-Hhm2ucIK",
		Zone:      []string{"zoneB"},
		NodeCount: 1,
	}
	err := asClient.IncreaseAsGroup(req)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DecreaseAs(t *testing.T) {
	req := &DecreaseAsGroupRequest{
		GroupId: "asg-Hhm2ucIK",
		Nodes:   []string{"i-z0PXqFD3"},
	}
	err := asClient.DecreaseAsGroup(req)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AdjustAs(t *testing.T) {
	req := &AdjustAsGroupRequest{
		GroupId:   "asg-Hhm2ucIK",
		AdjustNum: 1,
	}
	err := asClient.AdjustAsGroup(req)
	ExpectEqual(t.Errorf, nil, err)
}

func Test_CreateAsGroup(t *testing.T) {
	req := &CreateAsGroupRequest{
		GroupName: "yyy-test-gosdk2",
		Config: Config{
			MinNodeNum:    0,
			ExpectNum:     -1,
			MaxNodeNum:    2000,
			CooldownInSec: 300,
		},
		HealthCheck: HealthCheck{
			HealthCheckInterval: 15,
			GraceTime:           300,
		},
		ShrinkageStrategy: "Earlier",
		ZoneInfo: []ZoneInfo{
			{
				Zone:     "zoneA",
				SubnetID: "sbn-8mghgkzs3ch9",
			},
		},
		AssignTagInfo: AssignTagInfo{
			RelationTag: false,
			Tags: []Tag{
				{
					TagKey:   "默认项目",
					TagValue: "",
				},
			},
		},
		Nodes: []NodeInfo{
			{
				CpuCount:           8,
				MemoryCapacityInGB: 32,
				SysDiskType:        "enhanced_ssd_pl1",
				SysDiskInGB:        20,
				InstanceType:       13,
				ProductType:        "bidding",
				ImageId:            "24e80264-8a6d-49c1-b415-116d9cf38a75",
				ImageType:          "custom",
				OsType:             "linux",
				SecurityGroupId:    "g-yhryv5vyapb4",
				Spec:               "bcc.g4.c8m32",
				Priorities:         1,
				ZoneSubnet:         "[{\"zone\":\"zoneA\",\"subnetId\":\"sbn-8mghgkzs3ch9\",\"subnetName\":\"lyz2（192.168.0.0/24）\",\"subnetUuid\":\"5911e194-528f-4153-99a3-3c63b7bc7d7c\"}]",
				TotalCount:         1,
				BidModel:           "custom",
				BidPrice:           0.0264944,
			},
		},
		Eip: Eip{
			IfBindEip: false,
		},
		Billing: Billing{
			PaymentTiming: "bidding",
		},
		CmdConfig: CmdConfig{
			HasDecreaseCmd: false,
			DecCmdStrategy: "Proceed",
			DecCmdTimeout:  3600,
			DecCmdManual:   true,
			HasIncreaseCmd: false,
			IncCmdStrategy: "Proceed",
			IncCmdTimeout:  3600,
			IncCmdManual:   true,
		},
		BccNameConfig: BccNameConfig{
			BccName:     "",
			BccHostname: "",
		},
	}
	resp, err := asClient.CreateAsGroup(req)
	fmt.Printf("err is %v\n", err)
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
	ExpectEqual(t.Errorf, nil, err)
}

func Test_DeleteAsGroup(t *testing.T) {
	req := &DeleteAsGroupRequest{
		GroupIds: []string{"asg-CC1hjOJM"},
	}
	err := asClient.DeleteAsGroup(req)
	fmt.Printf("err is %v\n", err)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AdjustNode(t *testing.T) {

	res, err := asClient.AdjustNode("asg-sTufLpId", 2)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestClient_AttachNode(t *testing.T) {
	req := NodeRequest{
		Nodes: []string{"i-oukQ3mJd"},
	}
	err := asClient.AttachNode("asg-sTufLpId", &req)
	if err != nil {
		fmt.Print(err)
	}
}

func TestClient_DetachNode(t *testing.T) {
	req := NodeRequest{
		Nodes: []string{"i-mPkY5ZG5"},
	}
	err := asClient.DetachNode("asg-mPWFLu1E", &req)
	if err != nil {
		fmt.Print(err)
	}
}

func TestClient_ExecRule(t *testing.T) {

	res, err := asClient.ExecRule("asg-sTufLpId", "asrule-Z5l71OKv")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestClient_ListRecords(t *testing.T) {

	req := ListRecordsRequest{
		GroupID: "asg-sTufLpId",
		PageNo:  1,
		OrderBy: "startTime",
	}
	records, err := asClient.ListRecords(&req)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("create custom alarm policy success\n")
		fmt.Printf("%+v", records)
	}

}

func TestClient_ScalingDown(t *testing.T) {

	req := ScalingDownRequest{
		Nodes: []string{"i-oukQ3mJd"},
	}
	err := asClient.ScalingDown("asg-sTufLpId", &req)
	if err != nil {
		fmt.Print(err)
	}
}

func TestClient_ScalingUp(t *testing.T) {

	req := ScalingUpRequest{
		NodeCount: 1,
		Zone:      []string{"cn-bj-a"},
	}
	res, err := asClient.ScalingUp("asg-sTufLpId", &req)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestClient_CreateRule(t *testing.T) {
	req := RuleRequest{
		RuleName:        "testRule",
		GroupID:         "asg-mPWFLu1E",
		Type:            "PERIOD",
		ActionType:      "INCREASE",
		ActionNum:       1,
		CooldownInSec:   300,
		State:           "ENABLE",
		PeriodType:      "WEEK",
		PeriodStartTime: "2023-12-11T11:00:00Z",
		PeriodEndTime:   "2023-12-21T11:00:00Z",
		CronTime:        "12:30",
		PeriodValue:     2,
	}

	res, err := asClient.CreateRule(&req)
	if err != nil {
		t.Errorf("create group rule err: %v \n", err)
	}
	content, _ := json.Marshal(res)
	fmt.Println(string(content))
}

func TestClient_UpdateRule(t *testing.T) {
	req := RuleRequest{
		RuleName:        "testRule_update",
		GroupID:         "asg-mPWFLu1E",
		Type:            "PERIOD",
		ActionType:      "INCREASE",
		ActionNum:       1,
		CooldownInSec:   300,
		State:           "ENABLE",
		PeriodType:      "WEEK",
		PeriodStartTime: "2023-12-11T11:00:00Z",
		PeriodEndTime:   "2023-12-21T11:00:00Z",
		CronTime:        "12:40",
		PeriodValue:     2,
	}

	err := asClient.UpdateRule("asrule-dCtdLHRH", &req)
	if err != nil {
		t.Errorf("update group rule err: %v \n", err)
	}
}

func TestClient_GetRuleList(t *testing.T) {
	req := RuleListQuery{
		PageNo:   1,
		PageSize: 10,
		GroupID:  "asg-mPWFLu1E",
		OrderBy:  "createTime",
	}
	res, err := asClient.GetRuleList(&req)
	if err != nil {
		t.Errorf("get group rule list err: %v \n", err)
	}
	content, _ := json.Marshal(res)
	fmt.Println(string(content))
}

func TestClient_GetRuleDetail(t *testing.T) {
	res, err := asClient.GetRule("asrule-dCtdLHRH")
	if err != nil {
		t.Errorf("get group rule detail err: %v \n", err)
	}
	content, _ := json.Marshal(res)
	fmt.Println(string(content))
}

func TestClient_DeleteRule(t *testing.T) {
	req := RuleDelRequest{
		RuleIds: []string{"asrule-dCtdLHRH"},
	}
	err := asClient.DeleteRule(&req)
	if err != nil {
		t.Errorf("delete group rule err: %v \n", err)
	}
}
