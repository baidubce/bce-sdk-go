package cfc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/services/cfc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CfcClient       *Client
	FunctionName01  string
	FunctionName02  string
	AliasName01     string
	AliasName02     string
	FunctionBRN     string
	CodeSha256      string
	RelationId      string
	zipFilePython   string
	zipFileNodejs01 string
	zipFileNodejs02 string
	LayerName01     string
	LayerName02     string
	LayerVersionBRN string
	ServiceName01   string
	ServiceName02   string
)

const (
	invokeTestReturnPayload = "Hello World"
)

var (
	logSuccess = true
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

type PayloadDemo struct {
	A string
	B int
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given:(%s) err(%v)\n", conf, err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)
	FunctionName01 = fmt.Sprintf("sdktest-function01-%s", time.Now().Format("2006-01-02T150405"))
	FunctionName02 = fmt.Sprintf("sdktest-function02-%s", time.Now().Format("2006-01-02T150405"))
	zipFilePython = filepath.Join(filepath.Dir(f), "./python.zip")
	zipFileNodejs01 = filepath.Join(filepath.Dir(f), "./nodejs.zip")
	zipFileNodejs02 = filepath.Join(filepath.Dir(f), "./nodejs2.zip")

	AliasName01 = fmt.Sprintf("sdktest-alias01-%s", time.Now().Format("2006-01-02T150405"))
	AliasName02 = fmt.Sprintf("sdktest-alias02-%s", time.Now().Format("2006-01-02T150405"))

	LayerName01 = fmt.Sprintf("sdktest-layer01-%s", time.Now().Format("2006-01-02T150405"))
	LayerName02 = fmt.Sprintf("sdktest-layer02-%s", time.Now().Format("2006-01-02T150405"))

	ServiceName01 = fmt.Sprintf("sdktest-service01-%s", time.Now().Format("2006-01-02T150405"))
	ServiceName02 = fmt.Sprintf("sdktest-service02-%s", time.Now().Format("2006-01-02T150405"))

	CfcClient, err = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	if err != nil {
		panic(err)
	}
	log.SetLogHandler(log.FILE)
	//log.SetLogHandler(log.STDERR | log.FILE)
	//log.SetRotateType(log.ROTATE_SIZE)
	log.SetLogLevel(log.DEBUG)
}

func TestCreateFunction(t *testing.T) {
	codeFile, err := ioutil.ReadFile(zipFilePython)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}

	codeFile2, err := ioutil.ReadFile(zipFilePython)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	// This function return Hello World
	cases := []api.CreateFunctionArgs{
		{
			Code:         &api.CodeFile{ZipFile: codeFile},
			FunctionName: FunctionName01,
			Handler:      "index.handler",
			Runtime:      "python3",
			MemorySize:   128,
			Timeout:      3,
			Description:  "Description",
		},
		{
			Code:         &api.CodeFile{ZipFile: codeFile2},
			FunctionName: FunctionName02,
			Handler:      "index.handler",
			Runtime:      "python3",
			MemorySize:   256,
			Timeout:      3,
			Description:  "Description",
		},
	}
	for _, args := range cases {
		res, err := CfcClient.CreateFunction(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}
		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestCreateFunctionByBlueprint(t *testing.T) {
	cases := []api.CreateFunctionByBlueprintArgs{
		{
			FunctionName: "f2-1",
			ServiceName:  "default",
			BlueprintID:  "dd4372ef-0a8c-43b6-8c77-723da09ce439",
			Environment: &api.Environment{
				Variables: map[string]string{
					"k1": "v1",
				},
			},
		},
	}
	for _, args := range cases {
		res, err := CfcClient.CreateFunctionByBlueprint(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}
		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestListFunctions(t *testing.T) {
	args := &api.ListFunctionsArgs{
		FunctionVersion: "1",
		Marker:          1,
		MaxItems:        2,
	}
	res, err := CfcClient.ListFunctions(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestGetFunction(t *testing.T) {
	res, err := CfcClient.GetFunction(&api.GetFunctionArgs{
		FunctionName: FunctionName01,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	FunctionBRN = res.Configuration.FunctionBrn
	fmt.Printf(FunctionBRN)
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestInvocations(t *testing.T) {
	cases := []struct {
		args        *api.InvocationsArgs
		respPayload string
		err         error
	}{
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeRequestResponse,
				Payload:        nil,
			},
			respPayload: invokeTestReturnPayload,
			err:         nil,
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeEvent,
				Payload:        `[{"a":1},{"a":2}]`,
			},
			respPayload: "",
			err:         nil,
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeRequestResponse,
				Payload:        `[{"a":,{"a":2}]`,
			},
			respPayload: "",
			err:         fmt.Errorf("could not parse payload into json"),
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeEvent,
				Payload:        []byte(`{"a":1}`),
			},
			respPayload: "",
			err:         nil,
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeRequestResponse,
				Payload:        []*PayloadDemo{&PayloadDemo{A: "1", B: 2}, &PayloadDemo{A: "3", B: 4}},
			},
			respPayload: invokeTestReturnPayload,
			err:         nil,
		},
	}
	for _, tc := range cases {
		t.Run("invoke", func(t *testing.T) {
			res, err := CfcClient.Invocations(tc.args)
			if err == nil && tc.err != nil {
				t.Errorf("Expected err to be %v, but got nil", tc.err)
			} else if err != nil && tc.err == nil {
				t.Errorf("Expected err to be nil, but got %v", err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected err to be %v, but got %v", tc.err, err)
			} else if res != nil && res.Payload != tc.respPayload {
				t.Errorf("Expected Payload to be %s, but got %s", tc.respPayload, res.Payload)
			}
		})
	}
}

func TestUpdateFunctionCode(t *testing.T) {
	codeFile, err := ioutil.ReadFile(zipFileNodejs02)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	res, err := CfcClient.UpdateFunctionCode(&api.UpdateFunctionCodeArgs{
		FunctionName: FunctionName02,
		ZipFile:      codeFile,
		Publish:      false,
		DryRun:       false,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestUpdateFunctionConfiguration(t *testing.T) {
	res, err := CfcClient.UpdateFunctionConfiguration(&api.UpdateFunctionConfigurationArgs{
		FunctionName: FunctionName02,
		Timeout:      5,
		Description:  "wooo cool",
		Handler:      "index.handler",
		Runtime:      "nodejs12",
		Environment: &api.Environment{
			Variables: map[string]string{
				"name": "Test",
			},
		},
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestGetFunctionConfiguration(t *testing.T) {
	res, err := CfcClient.GetFunctionConfiguration(&api.GetFunctionConfigurationArgs{
		FunctionName: FunctionName02,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestPublishVersion(t *testing.T) {
	res, err := CfcClient.GetFunction(&api.GetFunctionArgs{
		FunctionName: FunctionName02,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	CodeSha256 = res.Configuration.CodeSha256
	fmt.Printf(FunctionBRN)
	result, err := CfcClient.PublishVersion(&api.PublishVersionArgs{
		FunctionName: FunctionName02,
		Description:  "test",
		CodeSha256:   CodeSha256,
	})
	if logSuccess && err == nil {
		t.Logf("res %v ", result)
	}
}

func TestListVersionsByFunction(t *testing.T) {
	res, err := CfcClient.ListVersionsByFunction(&api.ListVersionsByFunctionArgs{
		FunctionName: FunctionName02,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestCreateAlias(t *testing.T) {
	cases := []api.CreateAliasArgs{
		{
			FunctionName:    FunctionName02,
			FunctionVersion: "$LATEST",
			Name:            AliasName01,
			Description:     "test alias",
		},
		{
			FunctionName:    FunctionName02,
			FunctionVersion: "$LATEST",
			Name:            AliasName02,
			Description:     "test alias",
		},
	}
	for _, args := range cases {
		res, err := CfcClient.CreateAlias(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}
		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestGetAlias(t *testing.T) {
	args := &api.GetAliasArgs{
		FunctionName: FunctionName02,
		AliasName:    AliasName01,
	}
	res, err := CfcClient.GetAlias(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestUpdateAlias(t *testing.T) {
	args := &api.UpdateAliasArgs{
		FunctionName:    FunctionName02,
		AliasName:       AliasName01,
		FunctionVersion: "$LATEST",
		Description:     "test alias " + AliasName01,
	}
	res, err := CfcClient.UpdateAlias(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestListAliases(t *testing.T) {
	args := &api.ListAliasesArgs{
		FunctionName:    FunctionName02,
		FunctionVersion: "$LATEST",
		Marker:          0,
		MaxItems:        2,
	}
	res, err := CfcClient.ListAliases(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestDeleteAlias(t *testing.T) {
	args := &api.DeleteAliasArgs{
		FunctionName: FunctionName02,
		AliasName:    AliasName02,
	}
	err := CfcClient.DeleteAlias(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

func TestCreateTrigger(t *testing.T) {
	cases := []api.CreateTriggerArgs{
		{
			Target: FunctionBRN,
			Source: api.SourceTypeHTTP,
			Data: struct {
				ResourcePath string
				Method       string
				AuthType     string
			}{
				ResourcePath: fmt.Sprintf("tr01-%s", time.Now().Format("2006-01-02T150405")),
				Method:       "GET",
				AuthType:     "anonymous",
			},
		}, {
			Target: FunctionBRN,
			Source: api.SourceTypeHTTP,
			Data: struct {
				ResourcePath string
				Method       string
				AuthType     string
			}{
				ResourcePath: fmt.Sprintf("tr02-%s", time.Now().Format("2006-01-02T150405")),
				Method:       "GET",
				AuthType:     "anonymous",
			},
		},
	}
	for _, args := range cases {
		res, err := CfcClient.CreateTrigger(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}

		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestListTriggers(t *testing.T) {
	args := &api.ListTriggersArgs{
		FunctionBrn: FunctionBRN,
	}
	res, err := CfcClient.ListTriggers(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	if len(res.Relation) > 0 {
		RelationId = res.Relation[0].RelationId
	}
	t.Logf("res %v", res)
	resStr, err := json.Marshal(res)
	if err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestUpdateTrigger(t *testing.T) {
	args := &api.UpdateTriggerArgs{
		RelationId: RelationId,
		Target:     FunctionBRN,
		Source:     api.SourceTypeHTTP,
		Data: struct {
			ResourcePath string
			Method       string
			AuthType     string
		}{
			ResourcePath: fmt.Sprintf("tr99-%s", time.Now().Format("2006-01-02T150405")),
			Method:       "GET",
			AuthType:     "anonymous",
		},
	}
	res, err := CfcClient.UpdateTrigger(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	if res.Relation != nil {
		RelationId = res.Relation.RelationId
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestDeleteTrigger(t *testing.T) {
	listArgs := &api.ListTriggersArgs{
		FunctionBrn: FunctionBRN,
	}
	res, err := CfcClient.ListTriggers(listArgs)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	if len(res.Relation) > 0 {
		RelationId = res.Relation[0].RelationId
	}
	args := &api.DeleteTriggerArgs{
		RelationId: RelationId,
		Target:     FunctionBRN,
		Source:     api.SourceTypeHTTP,
	}
	t.Logf("args (%+v)", args)
	err = CfcClient.DeleteTrigger(args)
	if err != nil {
		t.Errorf("err (%v)", err)
	}
}

func TestDeleteFunction(t *testing.T) {
	args := &api.DeleteFunctionArgs{
		FunctionName: FunctionName01,
	}
	err := CfcClient.DeleteFunction(args)
	if err != nil {
		t.Logf("res (%v)", err)
	}
}

// TODO test fail
func TestListEventSource(t *testing.T) {
	FunctionBRN = ""
	args := &api.ListEventSourceArgs{
		FunctionName: FunctionBRN,
		Marker:       0,
		MaxItems:     100,
	}
	res, err := CfcClient.ListEventSource(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
	resStr, err := json.Marshal(res)
	if err == nil {
		t.Logf("res %s ", resStr)
	}
}

// test pass
func TestGetEventSource(t *testing.T) {
	args := &api.GetEventSourceArgs{
		UUID: "uuid",
	}
	res, err := CfcClient.GetEventSource(args)
	if err != nil {
		t.Logf("res (%v)", err)
	}
	t.Logf("res %+v", res)
	resStr, err := json.Marshal(res)
	if err == nil {
		t.Logf("res %s ", resStr)
	}
}

// test pass
func TestCreateEventSource(t *testing.T) {
	unEnabled := false
	FunctionBRN = ""
	args := &api.CreateEventSourceArgs{
		Enabled:      &unEnabled,
		BatchSize:    3,
		Type:         api.TypeEventSourceDatahubTopic,
		FunctionName: FunctionBRN,
		DatahubConfig: api.DatahubConfig{
			MetaHostEndpoint: "endpoint",
			MetaHostPort:     2181,
			ClusterName:      "clusterName",
			PipeName:         "pipeName",
			PipeletNum:       1,
			StartPoint:       -1,
			AclName:          "aclName",
			AclPassword:      "aclPassword",
		},
	}
	res, err := CfcClient.CreateEventSource(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}
func TestUpdateEventSource(t *testing.T) {
	FunctionBRN = ""
	unEnabled := false
	args := &api.UpdateEventSourceArgs{
		UUID: "uuid",
		FuncEventSource: api.FuncEventSource{
			Enabled:      &unEnabled,
			BatchSize:    3,
			Type:         api.TypeEventSourceDatahubTopic,
			FunctionName: FunctionBRN,
			DatahubConfig: api.DatahubConfig{
				MetaHostEndpoint: "10.155.195.11",
				MetaHostPort:     2181,
				ClusterName:      "clusterName",
				PipeName:         "pipeName",
				PipeletNum:       1,
				StartPoint:       -1,
				AclName:          "aclName",
				AclPassword:      "aclPassword",
			},
		},
	}
	res, err := CfcClient.UpdateEventSource(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestDeleteEventSource(t *testing.T) {
	args := &api.DeleteEventSourceArgs{
		UUID: "uuid",
	}
	err := CfcClient.DeleteEventSource(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

func TestListFlow(t *testing.T) {
	res, err := CfcClient.ListFlow()
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res.Flows[0])
}

func TestCreateFlow(t *testing.T) {
	res, err := CfcClient.CreateFlow(&api.CreateUpdateFlowArgs{
		Name:        "demo-x2",
		Type:        "FDL",
		Definition:  "name: demo\nstart: initData\nstates:\n  - type: pass\n    name: initData\n    data:\n      hello: world\n    end: true",
		Description: "ut test demo",
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
}

func TestUpdateFlow(t *testing.T) {
	res, err := CfcClient.UpdateFlow(&api.CreateUpdateFlowArgs{
		Name:        "demo-x2",
		Type:        "FDL",
		Definition:  "name: demo\nstart: initData2\nstates:\n  - type: pass\n    name: initData2\n    data:\n      hello: world\n    end: true",
		Description: "ut test demo2",
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
}

func TestDescribeFlow(t *testing.T) {
	res, err := CfcClient.DescribeFlow("demo-x2")
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
}

func TestDeleteFlow(t *testing.T) {
	err := CfcClient.DeleteFlow("demo-x2")
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

func TestStartExecution(t *testing.T) {
	res, err := CfcClient.StartExecution(&api.StartExecutionArgs{
		FlowName:      "demo-x2",
		ExecutionName: "s3",
		Input:         "{\"fruits\":[\"apple\", \"banana\"]}",
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
}

func TestStopExecution(t *testing.T) {
	res, err := CfcClient.StopExecution("demo-x2", "s3")
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
}

func TestDescribeExecution(t *testing.T) {
	res, err := CfcClient.DescribeExecution("demo-x2", "s3")
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
}

func TestListExecutions(t *testing.T) {
	res, err := CfcClient.ListExecutions("demo-x2")
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v", res)
}

func TestGetExecutionHistory(t *testing.T) {
	res, err := CfcClient.GetExecutionHistory(&api.GetExecutionHistoryArgs{
		FlowName:      "demo-x2",
		ExecutionName: "s3",
		Limit:         40,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	for _, e := range res.Events {
		t.Logf("event %+v", e)
	}
}

// Layer related tests
func TestPublishLayer(t *testing.T) {
	codeFile, err := ioutil.ReadFile(zipFilePython)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	base64CodeFile := base64.StdEncoding.EncodeToString(codeFile)

	cases := []api.PublishLayerVersionInput{
		{
			LayerName:   LayerName01,
			Description: "Test layer for SDK testing",
			Content: &api.LayerVersionContentInput{
				ZipFile: base64CodeFile,
			},
			CompatibleRuntimes: []string{"python3", "python3.10"},
		},
		{
			LayerName:   LayerName02,
			Description: "Another test layer for SDK testing",
			Content: &api.LayerVersionContentInput{
				ZipFile: base64CodeFile,
			},
			CompatibleRuntimes: []string{"python3"},
		},
	}

	for _, args := range cases {
		res, err := CfcClient.PublishLayer(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}
		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestGetLayerVersion(t *testing.T) {
	args := &api.GetLayerVersionArgs{
		LayerName:     LayerName01,
		VersionNumber: "1",
	}
	res, err := CfcClient.GetLayerVersion(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestListLayerVersions(t *testing.T) {
	args := &api.ListLayerVersionsInput{
		LayerName: "sdktest-layer01-2025-06-17T205914",
		ListCondition: &api.ListCondition{
			PageNo:   1,
			PageSize: 10,
		},
	}
	res, err := CfcClient.ListLayerVersions(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestListLayers(t *testing.T) {
	args := &api.ListLayerInput{
		ListCondition: &api.ListCondition{
			PageNo:   1,
			PageSize: 10,
		},
	}
	res, err := CfcClient.ListLayers(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestDeleteLayerVersion(t *testing.T) {
	args := &api.DeleteLayerVersionArgs{
		LayerName:     LayerName02,
		VersionNumber: "1",
	}
	err := CfcClient.DeleteLayerVersion(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

func TestDeleteLayer(t *testing.T) {
	args := &api.DeleteLayerArgs{
		LayerName: LayerName01,
	}
	err := CfcClient.DeleteLayer(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

// Service related tests
func TestCreateService(t *testing.T) {
	args := &api.CreateServiceArgs{
		ServiceName: "test-service",
		ServiceDesc: stringPtr("Test service description from SDK"),
	}
	res, err := CfcClient.CreateService(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestGetService(t *testing.T) {
	// Use an existing service from the list
	res, err := CfcClient.GetService(&api.GetServiceArgs{
		ServiceName: "default", // Use the default service that always exists
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestUpdateService(t *testing.T) {
	// Use an existing service for update test
	args := &api.UpdateServiceArgs{
		ServiceName: "default", // Use the default service
		ServiceDesc: stringPtr("Updated test service description from SDK"),
	}
	res, err := CfcClient.UpdateService(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestListServices(t *testing.T) {

	res, err := CfcClient.ListServices()
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res.Services, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}

	// Test the direct Services slice
	if logSuccess {
		t.Logf("Found %d services", len(res.Services))
		for i, service := range res.Services {
			t.Logf("Service %d: %+v", i, service)
		}
	}
}

func TestDeleteService(t *testing.T) {
	// Test delete with a non-essential service (one of the test services)
	// First get the list to find a test service to delete
	listRes, err := CfcClient.ListServices()
	if err != nil {
		t.Fatalf("Failed to list services: %v", err)
	}

	// Find a test service (not the default one) to delete
	var serviceToDelete string
	for _, service := range listRes.Services {
		if service.ServiceName != "default" && len(service.ServiceName) > 10 {
			serviceToDelete = service.ServiceName
			break
		}
	}

	if serviceToDelete == "" {
		t.Skip("No test service found to delete")
		return
	}

	t.Logf("Deleting service: %s", serviceToDelete)
	err = CfcClient.DeleteService(&api.DeleteServiceArgs{
		ServiceName: serviceToDelete,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("Successfully deleted service: %s", serviceToDelete)
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}
