package dts

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

var (
	DTS_CLIENT *Client
	DTS_ID     string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 1; i++ {
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

	DTS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateDts(t *testing.T) {
	args := &CreateDtsArgs{
		ProductType:        "postpay",
		Type:               "migration",
		Standard:           "Large",
		SourceInstanceType: "public",
		TargetInstanceType: "public",
		CrossRegionTag:     0,
		DirectionType: "single",
	}
	result, err := DTS_CLIENT.CreateDts(args)
	ExpectEqual(t.Errorf, nil, err)

	DTS_ID = result.DtsTasks[0].DtsId
}

func TestClient_GetDetail(t *testing.T) {
	//result, err := DTS_CLIENT.GetDetail("dtsmris3cu2k0uw3fuo0")
	result, err := DTS_CLIENT.GetDetail("dtsbukshhbsvd6yo7z96")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "unConfig", result.Status)
}

func TestClient_ListDts(t *testing.T) {
	var count = 0
	args := &ListDtsArgs{
		Type: "migration",
		MaxKeys: 10,
	}
	for true {
		result, err := DTS_CLIENT.ListDts(args)
		ExpectEqual(t.Errorf, nil, err)
		args.Marker = result.NextMarker
		for _, e := range result.Task {
			fmt.Println("dtsId: ", e.DtsId)
		}
		count += len(result.Task)
		if !result.IsTruncated {
			break
		}
	}
	fmt.Println("count: ", count)
}

func TestClient_ListDtsWithPage(t *testing.T) {
	var count = 0
	args := &ListDtsWithPageArgs{
		Types: []string{"bidirect"},
		Filters: []ListFilter{
			{
				Keyword: "he",
				KeywordType: "taskName",
			},
		},
		PageNo: 1,
		PageSize: 10,
		Order: "desc",
		OrderBy: "createTime",
	}
	for true {
		result, err := DTS_CLIENT.ListDtsWithPage(args)
		ExpectEqual(t.Errorf, nil, err)
		for _, e := range result.Result {
			fmt.Println("dtsId: ", e.DtsId)
		}
		count += len(result.Result)
		if len(result.Result) < 10 {
			break
		}
	}
	fmt.Println("count: ", count)
}

func TestClient_DeleteDts(t *testing.T) {
	err := DTS_CLIENT.DeleteDts("dtsbbpvtjcb6ploqexvf")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ConfigDts(t *testing.T) {
	args := &ConfigArgs {
		Type: "migration",
		TaskName: "go-sdk-1",
		DataType: []string{"increment"},
		SrcConnection: Connection{
			InstanceType: "bcerds",
			DbType: "mysql",
			InstanceId: "rdsm9744332",
			Region: "bj",
			FieldWhitelist: "",
			FieldBlacklist: "",
		},
		DstConnection: Connection{
			InstanceType: "bcerds",
			DbType: "mysql",
			InstanceId: "rdsmiu336698",
			Region: "bj",
			SqlType: "I,U",
		},
		SchemaMapping: []Schema{
			{
				Type: "table",
				Src: "hello.user",
				Dst: "hello.user",
				Where: "",
			},
		},
		Granularity: "dbtb",
		InitPosition: InitPosition{
			Type: "binlog",
			Position: "",
		},
	}
	result, err := DTS_CLIENT.ConfigDts("dtsmro61533558", args)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("result dtsId: ", result.DtsId)
}

func TestClient_PreCheck(t *testing.T) {
	result, err := DTS_CLIENT.PreCheck("dtsmro61533558")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("result success: ", result.Success)
	fmt.Println("result message: ", result.Message)
}

func TestClient_GetPreCheck(t *testing.T) {
	result, err := DTS_CLIENT.GetPreCheck("dtsmro61533558")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("result success: ", result.Success)
	for _, e := range result.Result {
		fmt.Println("name: ", e.Name, "status: ", e.Status, "Message: ", e.Message,
			"Subscription: ", e.Subscription)
	}
}

func TestClient_SkipPreCheck(t *testing.T) {
	response, err := DTS_CLIENT.SkipPreCheck("dtsmro61533558")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("response success: ", response.Success)
	fmt.Println("response result: ", response.Result)
}

func TestClient_StartDts(t *testing.T) {
	err := DTS_CLIENT.StartDts("dtsmro61533558")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_PauseDts(t *testing.T) {
	err := DTS_CLIENT.PauseDts("dtsmro61533558")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ShutdownDts(t *testing.T) {
	err := DTS_CLIENT.ShutdownDts("dtsmro61533558")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSchema(t *testing.T) {
	args := &GetSchemaArgs {
		Connection: Connection{
			InstanceType:   "bcerds",
			DbType:         "mysql",
			InstanceId:     "rdsm9744332",
			Region:         "bj",
			FieldWhitelist: "",
			FieldBlacklist: "",
		},
	}
	response, err := DTS_CLIENT.GetSchema(args)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("response success: ", response.Success)
	fmt.Println("response result: ", response.Result)
}

func TestClient_UpdateTaskName(t *testing.T) {
	args := &UpdateTaskNameArgs {
		TaskName: "go-sdkkk",
	}
	err := DTS_CLIENT.UpdateTaskName("dtsbe35xxxx8xzw365", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ResizeTaskStandard(t *testing.T) {
	args := &ResizeTaskStandardArgs {
		Standard: "Xlarge",
	}
	response, err := DTS_CLIENT.ResizeTaskStandard("dtsbexxxxxzw365", args)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("response orderId: ", response.OrderId)
}