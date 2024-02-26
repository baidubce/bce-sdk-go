package bcm

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_BlockCustomAlarmConfig(t *testing.T) {
	params := BlockCustomAlarmPolicyParams{
		UserId:    "453bf***************c984129090dc",
		Namespace: "cyn-928",
		AlarmName: "test-custom",
	}
	err := bcmClient.BlockCustomAlarmConfig(&params)
	if err != nil {
		fmt.Print("block failed.")
		fmt.Print(err)
	}
	fmt.Print("block successfully.")
}

func TestClient_CreateCustomAlarmPolicy(t *testing.T) {
	rule := CustomAlarmRule{
		Index:              1,
		MetricName:         "test",
		Dimensions:         []MetricDimensions{},
		Statistics:         "average",
		Threshold:          "1",
		ComparisonOperator: ">",
		Cycle:              60,
		Count:              1,
		Function:           "THRESHOLD",
	}
	config := CustomAlarmConfig{
		UserID:        "453bf***************c984129090dc",
		ActionEnabled: true,
		AlarmActions:  []string{"e3b8e777-4f35-48ed-abf6-bfcf6316ae2c"},
		Region:        "bj",
		Namespace:     "cyn-928",
		AlarmName:     "zsli_test_5",
		Level:         "MAJOR",
		Rules:         []CustomAlarmRule{rule},
	}

	err := bcmClient.CreateCustomAlarmPolicy(&config)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("create custom alarm policy success")
}

func TestClient_DeleteCustomAlarmPolicy(t *testing.T) {
	var policys []AlarmPolicyBatch
	policys = append(policys, AlarmPolicyBatch{
		UserId: "453bf***************c984129090dc",
		Scope:  "cyn-928",
		AlarmName: []string{
			"test-113017",
		},
	})
	policyToDelete := AlarmPolicyBatchList{
		CustomAlarmList: policys,
	}
	err := bcmClient.DeleteCustomAlarmPolicy(&policyToDelete)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("delete successfully.")

}

func TestClient_DetailCustomAlarmConfig(t *testing.T) {
	params := DetailCustomAlarmPolicyParams{
		UserId:    "453bf***************c984129090dc",
		Namespace: "cyn-928",
		AlarmName: "test-custom",
	}
	resp, err := bcmClient.DetailCustomAlarmConfig(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp)
	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(jsonData))
}

func TestClient_ListCustomAlarmPolicy(t *testing.T) {
	params := ListCustomAlarmPolicyParams{
		UserId:   "453bf***************c984129090dc",
		PageNo:   1,
		PageSize: 10,
	}
	resp, err := bcmClient.ListCustomAlarmPolicy(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp)
	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(jsonData))
}

func TestClient_UnblockCustomAlarmConfig(t *testing.T) {

	params := UnblockCustomAlarmPolicyParams{
		UserId:    "453bf***************c984129090dc",
		Namespace: "cyn-928",
		AlarmName: "test-custom",
	}
	err := bcmClient.UnblockCustomAlarmConfig(&params)
	if err != nil {
		fmt.Print("unblock failed.")
		fmt.Print(err)
	}
	fmt.Print("unblock successfully.")
}

func TestClient_UpdateCustomAlarmPolicy(t *testing.T) {
	rule := CustomAlarmRule{
		Index:              1,
		MetricName:         "test",
		Dimensions:         []MetricDimensions{},
		Statistics:         "average",
		Threshold:          "1",
		ComparisonOperator: ">",
		Cycle:              60,
		Count:              1,
		Function:           "THRESHOLD",
	}
	config := CustomAlarmConfig{
		UserID:        "453bf***************c984129090dc",
		ActionEnabled: true,
		AlarmActions:  []string{"e3b8e777-4f35-48ed-abf6-bfcf6316ae2c"},
		Region:        "bj",
		Namespace:     "cyn-928",
		AlarmName:     "zsli_test_5",
		Level:         "MAJOR",
		Rules:         []CustomAlarmRule{rule},
	}

	err := bcmClient.UpdateCustomAlarmPolicy(&config)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("update custom alarm policy success")
}
