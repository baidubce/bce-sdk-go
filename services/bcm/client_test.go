package bcm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/services/bcm/model"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	bcmClient *Client
	bcmConf   *Conf
)

type Conf struct {
	AK         string `json:"AK"`
	SK         string `json:"SK"`
	Endpoint   string `json:"Endpoint"`
	InstanceId string `json:"InstanceId"`
	UserId     string `json:"UserID"`
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
	bcmConf = &Conf{}
	_ = decoder.Decode(bcmConf)

	bcmClient, _ = NewClient(bcmConf.AK, bcmConf.SK, bcmConf.Endpoint)
	log.SetLogLevel(log.WARN)
}

func TestClient_GetMetricData(t *testing.T) {
	dimensions := map[string]string{
		"InstanceId": bcmConf.InstanceId,
	}
	req := &model.GetMetricDataRequest{
		UserId:         bcmConf.UserId,
		Scope:          "BCE_BCC",
		MetricName:     "vCPUUsagePercent",
		Dimensions:     dimensions,
		Statistics:     strings.Split(Average+","+SampleCount+","+Sum+","+Minimum+","+Maximum, ","),
		PeriodInSecond: 60,
		StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
		EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
	}
	resp, err := bcmClient.GetMetricData(req)
	if err != nil {
		t.Errorf("bcm get metric data error with %v\n", err)
	}
	if resp.Code != "OK" {
		t.Errorf("bcm get metric data response code error with %v\n", resp.Code)
	}
	if len(resp.DataPoints) < 1 {
		t.Error("bcm get metric data response dataPoints size not be greater 0\n")
	} else {
		if resp.DataPoints[0].Sum <= 0 {
			t.Error("bcm get metric data response dataPoints[0] sum not be greater 0\n")
		}
		if resp.DataPoints[0].Average <= 0 {
			t.Error("bcm get metric data response dataPoints[0] average not be greater 0\n")
		}
		if resp.DataPoints[0].SampleCount <= 0 {
			t.Error("bcm get metric data response dataPoints[0] sampleCount not be greater 0\n")
		}
		if resp.DataPoints[0].Minimum <= 0 {
			t.Error("bcm get metric data response dataPoints[0] minimum not be greater 0\n")
		}
		if resp.DataPoints[0].Maximum <= 0 {
			t.Error("bcm get metric data response dataPoints[0] maximum not be greater 0\n")
		}
	}
}

func TestClient_BatchGetMetricData(t *testing.T) {
	dimensions := map[string]string{
		"InstanceId": bcmConf.InstanceId,
	}
	req := &model.BatchGetMetricDataRequest{
		UserId:         bcmConf.UserId,
		Scope:          "BCE_BCC",
		MetricNames:    []string{"vCPUUsagePercent", "CpuIdlePercent"},
		Dimensions:     dimensions,
		Statistics:     strings.Split(Average+","+SampleCount+","+Sum+","+Minimum+","+Maximum, ","),
		PeriodInSecond: 60,
		StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
		EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
	}
	resp, err := bcmClient.BatchGetMetricData(req)
	if err != nil {
		t.Errorf("bcm batch get metric data error with %v\n", err)
	}
	if resp.Code != "OK" {
		t.Errorf("bcm batch get metric data response code error with %v\n", resp.Code)
	}
	if len(resp.ErrorList) > 0 {
		t.Error("bcm batch get metric data response errorList size not be lower 1\n")
	}
	if len(resp.SuccessList) < 2 {
		t.Error("bcm batch get metric data response successList size not be greater 1\n")
	} else {
		if len(resp.SuccessList[0].DataPoints) <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints size not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Sum <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] sum not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Average <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] average not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].SampleCount <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] sampleCount not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Minimum <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] minimum not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Maximum <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] maximum not be greater 0\n")
		}
	}

}

func TestClient_CreateNamespace(t *testing.T) {
	ns := &model.Namespace{
		Name:           "Test01",
		NamespaceAlias: "test",
		Comment:        "test",
		UserId:         bcmConf.UserId,
	}
	err := bcmClient.CreateNamespace(ns)
	if err != nil {
		t.Errorf("bcm create namespace error with %v\n", err)
	}
}

func TestClient_BatchDeleteNamespaces(t *testing.T) {
	cns := &model.CustomBatchNames{
		Names:  []string{"Test01"},
		UserId: bcmConf.UserId,
	}
	err := bcmClient.BatchDeleteNamespaces(cns)
	if err != nil {
		t.Errorf("bcm batch delete namespace error with %v\n", err)
	}
}

func TestClient_UpdateNamespace(t *testing.T) {
	ns := &model.Namespace{
		Name:           "Test01",
		NamespaceAlias: "test1",
		Comment:        "test1",
		UserId:         bcmConf.UserId,
	}
	err := bcmClient.UpdateNamespace(ns)
	if err != nil {
		t.Errorf("bcm update namespace error with %v\n", err)
	}
}

func TestClient_ListNamespaces(t *testing.T) {
	req := &model.ListNamespacesRequest{
		UserId:   bcmConf.UserId,
		PageNo:   1,
		PageSize: 10,
	}
	res, err := bcmClient.ListNamespaces(req)
	if err != nil {
		t.Errorf("bcm list namespaces error with %v\n", err)
		return
	}
	if res.PageNo != 1 {
		t.Errorf("bcm list namespace pageNo != 0 with pageNo: %d", res.PageNo)
	}
	if res.PageSize != 10 {
		t.Errorf("bcm list namespace pageSize != 10 with pageSize: %d", res.PageSize)
	}
	fmt.Println(res)

	req = &model.ListNamespacesRequest{
		UserId:   bcmConf.UserId,
		Name:     "Test01",
		PageNo:   1,
		PageSize: 10,
	}
	res, err = bcmClient.ListNamespaces(req)
	if err != nil {
		t.Errorf("bcm list namespaces error with %v\n", err)
		return
	}
	if res.PageNo != 1 {
		t.Errorf("bcm list namespace pageNo != 0 with pageNo: %d", res.PageNo)
	}
	if res.PageSize != 10 {
		t.Errorf("bcm list namespace pageSize != 10 with pageSize: %d", res.PageSize)
	}
	fmt.Println(res)
}

func TestClient_CreateNamespaceMetric(t *testing.T) {
	nm := &model.NamespaceMetric{
		UserId:      bcmConf.UserId,
		Namespace:   "Test01",
		MetricName:  "TestMetric01",
		MetricAlias: "test",
		Unit:        "sec",
		Cycle:       60,
	}
	err := bcmClient.CreateNamespaceMetric(nm)
	if err != nil {
		t.Errorf("bcm create namespace metric error with %v\n", err)
	}
}

func TestClient_BatchDeleteNamespaceMetrics(t *testing.T) {
	cns := &model.CustomBatchIds{
		UserId:    bcmConf.UserId,
		Namespace: "Test01",
		Ids:       []int64{1710},
	}
	err := bcmClient.BatchDeleteNamespaceMetrics(cns)
	if err != nil {
		t.Errorf("bcm batch delete namespace metric error with %v\n", err)
	}
}

func TestClient_UpdateNamespaceMetric(t *testing.T) {
	nm := &model.NamespaceMetric{
		UserId:      bcmConf.UserId,
		Namespace:   "Test01",
		MetricName:  "TestMetric02",
		MetricAlias: "test01",
		Unit:        "sec",
		Cycle:       60,
	}
	err := bcmClient.UpdateNamespaceMetric(nm)
	if err != nil {
		t.Errorf("bcm update namespace metric error with %v\n", err)
	}
}

func TestClient_ListNamespaceMetrics(t *testing.T) {
	req := &model.ListNamespaceMetricsRequest{
		UserId:    bcmConf.UserId,
		Namespace: "Test01",
		PageNo:    1,
		PageSize:  10,
	}
	res, err := bcmClient.ListNamespaceMetrics(req)
	if err != nil {
		t.Errorf("bcm list namespace metrics error with %v\n", err)
		return
	}
	if res.PageNo != 1 {
		t.Errorf("bcm list namespace metrics pageNo != 0 with pageNo: %d", res.PageNo)
	}
	if res.PageSize != 10 {
		t.Errorf("bcm list namespace metrics pageSize != 10 with pageSize: %d", res.PageSize)
	}
	fmt.Println(res)

	req = &model.ListNamespaceMetricsRequest{
		UserId:      bcmConf.UserId,
		Namespace:   "Test01",
		MetricName:  "TestMetric01",
		MetricAlias: "test",
		PageNo:      1,
		PageSize:    10,
	}
	res, err = bcmClient.ListNamespaceMetrics(req)
	if err != nil {
		t.Errorf("bcm list namespace metrics error with %v\n", err)
		return
	}
	if res.PageNo != 1 {
		t.Errorf("bcm list namespace metrics pageNo != 0 with pageNo: %d", res.PageNo)
	}
	if res.PageSize != 10 {
		t.Errorf("bcm list namespace metrics pageSize != 10 with pageSize: %d", res.PageSize)
	}
	fmt.Println(res)
}

func TestClient_GetCustomMetric(t *testing.T) {
	userId := bcmConf.UserId
	namespace := "Test01"
	metricName := "TestMetric01"
	nm, err := bcmClient.GetCustomMetric(userId, namespace, metricName)
	if err != nil {
		t.Errorf("bcm get custom metric error with %v\n", err)
		return
	}
	if nm.UserId != userId {
		t.Errorf("bcm get custom metric userid != %s with userId: %s", userId, nm.UserId)
	}
	if nm.Namespace != namespace {
		t.Errorf("bcm get custom metric namespace != %s with namespace: %s", namespace, nm.Namespace)
	}
	if nm.MetricName != metricName {
		t.Errorf("bcm get custom metric metricName != %s with metricName: %s", metricName, nm.MetricName)
	}
	fmt.Println(nm)
}

func TestClient_CreateNamespaceEvent(t *testing.T) {
	ne := &model.NamespaceEvent{
		UserId:         bcmConf.UserId,
		Namespace:      "Test01",
		EventName:      "TestEvent01",
		EventNameAlias: "test",
		EventLevel:     WarningEventLevel,
		Comment:        "event",
	}
	err := bcmClient.CreateNamespaceEvent(ne)
	if err != nil {
		t.Errorf("bcm create namespace event error with %v\n", err)
	}
}

func TestClient_BatchDeleteNamespaceEvents(t *testing.T) {
	ces := &model.CustomBatchEventNames{
		UserId:    bcmConf.UserId,
		Namespace: "Test01",
		Names:     []string{"TestEvent01"},
	}
	err := bcmClient.BatchDeleteNamespaceEvents(ces)
	if err != nil {
		t.Errorf("bcm batch delete namespace event error with %v\n", err)
	}
}

func TestClient_UpdateNamespaceEvent(t *testing.T) {
	ne := &model.NamespaceEvent{
		UserId:         bcmConf.UserId,
		Namespace:      "Test01",
		EventName:      "TestEvent01",
		EventNameAlias: "test01",
		EventLevel:     WarningEventLevel,
		Comment:        "event01",
	}
	err := bcmClient.UpdateNamespaceEvent(ne)
	if err != nil {
		t.Errorf("bcm update namespace event error with %v\n", err)
	}
}

func TestClient_ListNamespaceEvents(t *testing.T) {
	req := &model.ListNamespaceEventsRequest{
		UserId:    bcmConf.UserId,
		Namespace: "Test01",
		PageNo:    1,
		PageSize:  10,
	}
	res, err := bcmClient.ListNamespaceEvents(req)
	if err != nil {
		t.Errorf("bcm list namespace metrics error with %v\n", err)
		return
	}
	if res.PageNo != 1 {
		t.Errorf("bcm list namespace metrics pageNo != 0 with pageNo: %d", res.PageNo)
	}
	if res.PageSize != 10 {
		t.Errorf("bcm list namespace metrics pageSize != 10 with pageSize: %d", res.PageSize)
	}
	fmt.Println(res)

	req = &model.ListNamespaceEventsRequest{
		UserId:     bcmConf.UserId,
		Namespace:  "Test01",
		Name:       "TestEvent01",
		EventLevel: WarningEventLevel,
		PageNo:     1,
		PageSize:   10,
	}
	res, err = bcmClient.ListNamespaceEvents(req)
	if err != nil {
		t.Errorf("bcm list namespace metrics error with %v\n", err)
		return
	}
	if res.PageNo != 1 {
		t.Errorf("bcm list namespace metrics pageNo != 0 with pageNo: %d", res.PageNo)
	}
	if res.PageSize != 10 {
		t.Errorf("bcm list namespace metrics pageSize != 10 with pageSize: %d", res.PageSize)
	}
	fmt.Println(res)
}

func TestClient_GetCustomEvent(t *testing.T) {
	userId := bcmConf.UserId
	namespace := "Test01"
	eventName := "TestEvent01"
	ne, err := bcmClient.GetCustomEvent(userId, namespace, eventName)
	if err != nil {
		t.Errorf("bcm get custom event error with %v\n", err)
		return
	}
	if ne.UserId != userId {
		t.Errorf("bcm get custom event userid != %s with userId: %s", userId, ne.UserId)
	}
	if ne.Namespace != namespace {
		t.Errorf("bcm get custom event namespace != %s with namespace: %s", namespace, ne.Namespace)
	}
	if ne.EventName != eventName {
		t.Errorf("bcm get custom event metricName != %s with metricName: %s", eventName, ne.EventName)
	}
	fmt.Println(ne)
}

func TestCreateApplicationData(t *testing.T) {
	req := &model.ApplicationInfoRequest{
		Alias:       "test",
		Name:        "test_1207",
		Type:        "BCC",
		UserID:      "453bf9588c9e488f9ba2c984129090dc",
		Description: "test",
	}
	data, err := bcmClient.CreateApplicationData(req)
	if err != nil {
		t.Errorf("bcm create application data error with %v\n", err)
	}
	marshal, _ := json.Marshal(data)
	if len(marshal) <= 0 {
		t.Errorf("bcm create application data empty")
	}
	t.Logf("bcm create application data with %v\n", string(marshal))
}

func TestGetApplicationDataList(t *testing.T) {
	userId := bcmConf.UserId
	pageSize := 10
	pageNo := 1
	searchName := "test"
	req, err := bcmClient.GetApplicationDataList(userId, searchName, pageSize, pageNo)
	if err != nil {
		t.Errorf("bcm get application data list error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm get application data list empty")
	}
	t.Logf("bcm get application data list with %v\n", string(marshal))
}

func TestClient_UpdateApplicationData(t *testing.T) {
	res := &model.ApplicationInfoUpdateRequest{
		ID:          5336,
		Alias:       "test1206",
		Name:        "test_1206",
		Type:        "BCC",
		UserID:      "453bf9588c9e488f9ba2c984129090dc",
		Description: "testD",
	}
	req, err := bcmClient.UpdateApplicationData(res)
	if err != nil {
		t.Errorf("bcm update application data error with %v\n", err)
	}
	if req.Alias != "test1206" {
		t.Errorf("bcm update application data alias != test1206 with %s", req.Alias)
	}
	t.Logf("bcm update application data success")
}

func TestClient_DeleteApplicationData(t *testing.T) {
	res := &model.ApplicationInfoDeleteRequest{
		Name: "test_1206",
	}
	userId := bcmConf.UserId
	err := bcmClient.DeleteApplicationData(userId, res)
	if err != nil {
		t.Errorf("bcm delete application data error with %v\n", err)
	}
	t.Logf("bcm delete application data success")
}

func TestClient_GetApplicationInstanceList(t *testing.T) {
	res := &model.ApplicationInstanceListRequest{
		PageNo:      1,
		PageSize:    10,
		AppName:     "test-1130",
		SearchName:  "name",
		SearchValue: "",
		Region:      "bj",
	}
	userId := bcmConf.UserId
	req, err := bcmClient.GetApplicationInstanceList(userId, res)
	if err != nil {
		t.Errorf("bcm get application instance list error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm get application instance list empty")
	}
	t.Logf("bcm get application instance list with %v\n", string(marshal))
}

func TestClient_CreateApplicationInstance(t *testing.T) {
	infos := []*model.HostInstanceInfo{
		{
			InstanceID: "5cd74fe6-f508-4238-b20e-bddb62243a33",
			Region:     "bj",
		},
		{
			InstanceID: "2427ad4f-ac45-48b2-9a22-b92109b3fd97",
			Region:     "bj",
		},
	}
	res := &model.ApplicationInstanceCreateRequest{
		AppName:  "test-1130",
		UserID:   bcmConf.UserId,
		HostList: infos,
	}
	err := bcmClient.CreateApplicationInstance(res)
	if err != nil {
		t.Errorf("bcm create application instance error with %v\n", err)
	}
	t.Logf("bcm create application instance success")
}

func TestClient_GetApplicationInstanceCreatedList(t *testing.T) {
	res := &model.ApplicationInstanceCreatedListRequest{
		UserID:  bcmConf.UserId,
		AppName: "test-1130",
		Region:  "bj",
	}
	req, err := bcmClient.GetApplicationInstanceCreatedList(res)
	if err != nil {
		t.Errorf("bcm get application instance created list error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm get application instance created list empty")
	}
	t.Logf("bcm get application instance created list with %v\n", string(marshal))
}

func TestClient_DeleteApplicationInstance(t *testing.T) {
	res := &model.ApplicationInstanceDeleteRequest{
		ID:      "6980",
		AppName: "test-1130",
	}
	err := bcmClient.DeleteApplicationInstance(bcmConf.UserId, res)
	if err != nil {
		t.Errorf("bcm delete application instance error with %v\n", err)
	}
	t.Logf("bcm delete application instance success")
}

func TestClient_CreateApplicationMonitorTask01(t *testing.T) {
	res := &model.ApplicationMonitorTaskInfoRequest{
		AppName:     "test-1130",
		Type:        0,
		AliasName:   "test-proc",
		Cycle:       60,
		Target:      "/proc",
		Description: "test-1207",
	}
	req, err := bcmClient.CreateApplicationInstanceTask(bcmConf.UserId, res)
	if err != nil {
		t.Errorf("bcm create application monitor task error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm create application monitor task empty")
	}
	t.Logf("bcm create application monitor task with %v\n", string(marshal))
}

func TestClient_CreateApplicationMonitorTask02(t *testing.T) {
	extractResult := []*model.LogExtractResult{
		{
			ExtractFieldName:  "namespace",
			ExtractFieldValue: "04b91096-a294-477d-bd11-1a7bcfb5a921",
			DimensionMapTable: "namespaceTable",
		},
	}
	tags := []*model.AggTag{
		{
			Range: "App",
			Tags:  "",
		},
		{
			Range: "App",
			Tags:  "namespace",
		},
	}
	metrics := []*model.Metric{
		{
			MetricName:       "space",
			SaveInstanceData: 1,
			ValueFieldType:   0,
			MetricAlias:      "",
			MetricUnit:       "",
			ValueFieldName:   "",
			AggrTags:         tags,
		},
	}
	res := &model.ApplicationMonitorTaskInfoLogRequest{
		AppName:       "test-1130",
		Type:          2,
		AliasName:     "test-log-1207",
		Cycle:         60,
		Target:        "/opt/bcm-agent/log/bcm-agent.INFO",
		Description:   "test-LOG-1207",
		Rate:          5,
		ExtractResult: extractResult,
		LogExample:    "namespace:04b91096-a294-477d-bd11-1a7bcfb5a921\n",
		MatchRule:     "namespace:(?P<namespace>[0-9a-fA-F-]+)",
		UserID:        bcmConf.UserId,
		Metrics:       metrics,
	}
	req, err := bcmClient.CreateApplicationMonitorLogTask(bcmConf.UserId, res)
	if err != nil {
		t.Errorf("bcm create application monitor task error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm create application monitor task empty")
	}
	t.Logf("bcm create application monitor task with %v\n", string(marshal))
}

func TestClient_GetApplicationMonitorTaskDetail(t *testing.T) {
	res := &model.ApplicationMonitorTaskDetailRequest{
		AppName:  "test-1130",
		UserID:   bcmConf.UserId,
		TaskName: "d917e9963d6349909e3793b101e90333",
	}
	req, err := bcmClient.GetApplicationMonitorTaskDetail(res)
	if err != nil {
		t.Errorf("bcm get application monitor task detail error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm get application monitor task detail empty")
	}
	t.Logf("bcm get application monitor task detail with %v\n", string(marshal))
}

func TestClient_GetApplicationMonitorTaskList(t *testing.T) {
	res := &model.ApplicationMonitorTaskListRequest{
		AppName: "test-1130",
		UserID:  bcmConf.UserId,
	}
	req, err := bcmClient.GetApplicationMonitorTaskList(res)
	if err != nil {
		t.Errorf("bcm get application monitor task list error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm get application monitor task list empty")
	}
	t.Logf("bcm get application monitor task list with %v\n", string(marshal))
}

func TestClient_UpdateApplicationMonitorTask01(t *testing.T) {
	extractResult := []*model.LogExtractResult{
		{
			ExtractFieldName:  "namespace",
			ExtractFieldValue: "04b91096-a294-477d-bd11-1a7bcfb5a921",
			DimensionMapTable: "namespaceTable",
		},
	}
	tags := []*model.AggTag{
		{
			Range: "App",
			Tags:  "",
		},
		{
			Range: "App",
			Tags:  "namespace",
		},
	}
	metrics := []*model.Metric{
		{
			MetricName:       "space",
			SaveInstanceData: 1,
			ValueFieldType:   0,
			MetricAlias:      "",
			MetricUnit:       "",
			ValueFieldName:   "",
			AggrTags:         tags,
		},
	}
	res := &model.ApplicationMonitorTaskInfoUpdateRequest{
		ID:            "3922",
		Name:          "d917e9963d6349909e3793b101e90333",
		AppName:       "test-1130",
		Type:          2,
		AliasName:     "test-log-1207",
		Cycle:         60,
		Target:        "/opt/bcm-agent/log/bcm-agent.INFO",
		Description:   "test-LOG1207",
		Rate:          5,
		ExtractResult: extractResult,
		LogExample:    "namespace:04b91096-a294-477d-bd11-1a7bcfb5a921\n",
		MatchRule:     "namespace:(?P<namespace>[0-9a-fA-F-]+)",
		UserID:        bcmConf.UserId,
		Metrics:       metrics,
	}
	req, err := bcmClient.UpdateApplicationMonitorLogTask(bcmConf.UserId, res)
	if err != nil {
		t.Errorf("bcm create application monitor task error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm update application monitor task empty")
	}
	t.Logf("bcm update application monitor task with %v\n", string(marshal))
}

func TestClient_UpdateApplicationMonitorTask02(t *testing.T) {
	res := &model.ApplicationMonitorTaskInfoUpdateRequest{
		Name:        "456cfa48356845b6bac1d29abc85d8f4",
		AppName:     "test-1130",
		Type:        0,
		AliasName:   "test-proc-update02",
		Cycle:       60,
		Target:      "/proc/exec",
		Description: "test-1207",
	}
	req, err := bcmClient.UpdateApplicationMonitorTask(bcmConf.UserId, res)
	if err != nil {
		t.Errorf("bcm create application monitor task error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm update application monitor task empty")
	}
	t.Logf("bcm update application monitor task with %v\n", string(marshal))
}

func TestClient_DeleteApplicationMonitorTask(t *testing.T) {
	res := &model.ApplicationMonitorTaskDeleteRequest{
		Name:    "d917e9963d6349909e3793b101e90333",
		AppName: "test-1130",
		UserID:  bcmConf.UserId,
	}
	err := bcmClient.DeleteApplicationMonitorTask(res)
	if err != nil {
		t.Errorf("bcm delete application monitor task error with %v\n", err)
	}
}

func TestClient_CreateApplicationDimensionTable(t *testing.T) {
	res := &model.ApplicationDimensionTableInfoRequest{
		UserID:         bcmConf.UserId,
		AppName:        "test-1130",
		TableName:      "test-table",
		MapContentJSON: "a=>1\\nb=>2",
	}
	req, err := bcmClient.CreateApplicationDimensionTable(res)
	if err != nil {
		t.Errorf("bcm create application dimension table error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm create application dimension table empty")
	}
	t.Logf("bcm create application dimension table with %v\n", string(marshal))
}

func TestClient_GetApplicationDimensionTableList(t *testing.T) {
	res := &model.ApplicationDimensionTableListRequest{
		UserID:     bcmConf.UserId,
		AppName:    "test-1130",
		SearchName: "space",
	}
	req, err := bcmClient.GetApplicationDimensionTableList(res)
	if err != nil {
		t.Errorf("bcm get application dimension table list error with %v\n", err)
	}
	marshal, _ := json.Marshal(req)
	if len(marshal) <= 0 {
		t.Errorf("bcm get application dimension table list empty")
	}
	t.Logf("bcm get application dimension table list with %v\n", string(marshal))
}

func TestClient_UpdateApplicationDimensionTable(t *testing.T) {
	res := &model.ApplicationDimensionTableInfoRequest{
		UserID:         bcmConf.UserId,
		AppName:        "test-1130",
		TableName:      "test-table",
		MapContentJSON: "a=>1",
	}
	err := bcmClient.UpdateApplicationDimensionTable(res)
	if err != nil {
		t.Errorf("bcm update application dimension table error with %v\n", err)
	}
}

func TestClient_DeleteApplicationDimensionTable(t *testing.T) {
	res := &model.ApplicationDimensionTableDeleteRequest{
		UserID:    bcmConf.UserId,
		AppName:   "test-1130",
		TableName: "test-table",
	}
	err := bcmClient.DeleteApplicationDimensionTable(res)
	if err != nil {
		t.Errorf("bcm delete application dimension table error with %v\n", err)
	}
}

func TestClient_ListNotifyGroups(t *testing.T) {
	req := &model.ListNotifyGroupsRequest{
		Name:     "test",
		PageNo:   1,
		PageSize: 5,
	}
	res, err := bcmClient.ListNotifyGroups(req)
	if err != nil {
		t.Errorf("bcm list notify groups error with %v\n", err)
	}
	if res.Status != 200 {
		t.Errorf("bcm list notify groups error with %v\n", err)
	}
}

func TestClient_ListNotifyParties(t *testing.T) {
	req := &model.ListNotifyPartiesRequest{
		Name:     "test",
		PageNo:   1,
		PageSize: 5,
	}
	res, err := bcmClient.ListNotifyParty(req)
	if err != nil {
		t.Errorf("bcm list notify groups error with %v\n", err)
	}
	if res.Status != 200 {
		t.Errorf("bcm list notify groups error with status %d\n", res.Status)
	}
}

func TestClient_CreateAction(t *testing.T) {
	notification := model.ActionNotification{
		Type:     model.ActionNotificationTypeEmail,
		Receiver: "",
	}
	member := model.ActionMember{
		Type: "notifyParty",
		Id:   "56c9e0e2138c4f",
		Name: "lzs",
	}
	req := &model.CreateActionRequest{
		UserId:          bcmConf.UserId,
		Notifications:   []model.ActionNotification{notification},
		Members:         []model.ActionMember{member},
		Alias:           "test_wjr",
		DisableTimes:    nil,
		ActionCallBacks: nil,
	}
	err := bcmClient.CreateAction(req)
	if err != nil {
		t.Errorf("bcm create action error with %v\n", err)
	}
}

func TestClient_DeleteAction(t *testing.T) {
	req := &model.DeleteActionRequest{
		UserId: bcmConf.UserId,
		Name:   "b90d86da-e3a0-4c63-9bd2-a7e210d2027f",
	}
	err := bcmClient.DeleteAction(req)
	if err != nil {
		t.Errorf("bcm delete action error with %v\n", err)
	}
}

func TestClient_ListActions(t *testing.T) {
	req := &model.ListActionsRequest{
		UserId:   bcmConf.UserId,
		PageNo:   1,
		PageSize: 10,
	}
	resp, err := bcmClient.ListActions(req)
	if err != nil {
		t.Errorf("bcm list action error with %v\n", err)
	}
	if !resp.Success {
		t.Errorf("bcm list action error")
	}
}

func TestClient_UpdateAction(t *testing.T) {
	notification := model.ActionNotification{
		Type:     model.ActionNotificationTypeEmail,
		Receiver: "",
	}
	member := model.ActionMember{
		Type: "notifyParty",
		Id:   "56c9e0e2138c4f",
		Name: "lzs",
	}
	req := &model.UpdateActionRequest{
		UserId:          bcmConf.UserId,
		Name:            "4e9630d6-6348-450d-aab8-ea5f4003d6a4",
		Notifications:   []model.ActionNotification{notification},
		Members:         []model.ActionMember{member},
		Alias:           "test_wjr",
		DisableTimes:    nil,
		ActionCallBacks: nil,
	}
	err := bcmClient.UpdateAction(req)
	if err != nil {
		t.Errorf("bcm update action error with %v\n", err)
	}
}

func TestClient_LogExtract(t *testing.T) {
	req := &model.LogExtractRequest{
		UserId: bcmConf.UserId,
		ExtractRule: "800] \"(?<method>(GET|POST|PUT|DELETE)) .*/v1/dashboard/metric/(?<widget>(cycle|trend|report|billboard|gaugechart)) HTTP/1.1\".* " +
			"(?<resTime>[0-9]+)ms",
		LogExample: "10.157.16.207 - - [09/Apr/2020:20:45:33 +0800] \"POST /v1/dashboard/metric/gaugechart HTTP/1.1\" 200 117 109ms\n10.157.16.207" +
			" - - [09/Apr/2020:20:45:33 +0800] \"GET /v1/dashboard/metric/report HTTP/1.1\" 200 117 19ms",
	}

	resp, err := bcmClient.LogExtract(req)
	if err != nil {
		t.Errorf("bcm log extract error with %v\n", err)
	}
	if len(resp) <= 0 {
		t.Errorf("bcm log extract error")
	}
}

func TestClient_GetMetricMetaForApplication(t *testing.T) {
	req := &model.GetMetricMetaForApplicationRequest{
		UserId:        bcmConf.UserId,
		AppName:       "test14",
		TaskName:      "task13",
		MetricName:    "log.responseTime",
		Instances:     []string{"0.test14"},
		DimensionKeys: []string{"method"},
	}
	_, err := bcmClient.GetMetricMetaForApplication(req)
	if err != nil {
		t.Errorf("bcm get metric meta for application error with %v\n", err)
	}
}

func TestClient_GetMetricDataForApplication(t *testing.T) {
	req := &model.GetMetricDataForApplicationRequest{
		UserId:     bcmConf.UserId,
		AppName:    "gjm-test",
		TaskName:   "bbceac2807014fce920e92e31debf092",
		MetricName: "port.err_code",
		Instances:  []string{"8.gjm-test"},
		StartTime:  "2023-12-07T01:10:48Z",
		EndTime:    "2023-12-07T02:10:48Z",
		Cycle:      0,
		Statistics: []string{"average"},
		Dimensions: nil,
		AggrData:   false,
	}
	resp, err := bcmClient.GetMetricDataForApplication(req)
	if err != nil {
		t.Errorf("bcm get metric data for application error with %v\n", err)
	}
	fmt.Println(len(resp))
}

func TestClient_CreateAppMonitorAlarmConfig(t *testing.T) {
	req := &model.AppMonitorAlarmConfig{
		AlarmDescription:  "",
		AlarmName:         "test_wjr",
		UserId:            bcmConf.UserId,
		AppName:           "zmq-log-1115",
		MonitorObjectType: "APP",
		MonitorObject: model.MonitorObject{
			Id: 4030,
			MonitorObjectView: []model.MonitorObjectViewModel{{
				MonitorObjectName: "ab3b543f41974e26ab984d94fc7b9b92",
			}},
			MonitorObjectType: "APP",
		},
		SrcName:       "ab3b543f41974e26ab984d94fc7b9b92",
		SrcType:       "LOG",
		Type:          "INSTANCE",
		Level:         model.AlarmLevelMajor,
		ActionEnabled: true,
		PolicyEnabled: true,
		Rules: [][]model.AppMonitorAlarmRule{{{
			Metric:             "log.log_metric2",
			MetricAlias:        "log_metric2",
			Cycle:              60,
			Statistics:         "average",
			Threshold:          5,
			ComparisonOperator: ">",
			Count:              1,
			Function:           "THRESHOLD",
			Sequence:           0,
		}}},
		IncidentActions:     []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
		ResumeAction:        []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
		InsufficientActions: []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
		RepeatAlarmCycle:    300,
		MaxRepeatCount:      1,
	}

	resp, err := bcmClient.CreateAppMonitorAlarmConfig(req)
	if err != nil {
		t.Errorf("bcm create alarm config for application error with %v\n", err)
	}
	fmt.Println(*resp)
}

func TestClient_UpdateAppMonitorAlarmConfig(t *testing.T) {
	req := &model.AppMonitorAlarmConfig{
		AlarmDescription:  "",
		AlarmName:         "test_wjr",
		UserId:            bcmConf.UserId,
		AppName:           "zmq-log-1115",
		MonitorObjectType: "APP",
		MonitorObject: model.MonitorObject{
			Id: 4030,
			MonitorObjectView: []model.MonitorObjectViewModel{{
				MonitorObjectName: "ab3b543f41974e26ab984d94fc7b9b92",
			}},
			MonitorObjectType: "APP",
		},
		SrcName:       "ab3b543f41974e26ab984d94fc7b9b92",
		SrcType:       "LOG",
		Type:          "INSTANCE",
		Level:         model.AlarmLevelMajor,
		ActionEnabled: true,
		PolicyEnabled: true,
		Rules: [][]model.AppMonitorAlarmRule{[]model.AppMonitorAlarmRule{{
			Metric:             "log.log_metric2",
			MetricAlias:        "log_metric2",
			Cycle:              60,
			Statistics:         "average",
			Threshold:          5,
			ComparisonOperator: ">",
			Count:              1,
			Function:           "THRESHOLD",
			Sequence:           0,
		}}},
		IncidentActions:     []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
		ResumeAction:        []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
		InsufficientActions: []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
		RepeatAlarmCycle:    300,
		MaxRepeatCount:      1,
	}
	resp, err := bcmClient.UpdateAppMonitorAlarmConfig(req)
	if err != nil {
		t.Errorf("bcm update alarm config for application error with %v\n", err)
	}
	fmt.Println(*resp)
}

func TestClient_ListAppMonitorAlarmConfigs(t *testing.T) {
	req := &model.ListAppMonitorAlarmConfigsRequest{
		UserId:    bcmConf.UserId,
		AlarmName: "test",
		SrcType:   model.SrcTypeProc,
		PageNo:    1,
		PageSize:  10,
	}
	resp, err := bcmClient.ListAppMonitorAlarmConfigs(req)
	if err != nil {
		t.Errorf("bcm list alarm configs for application error with %v\n", err)
	}
	fmt.Println(*resp)
}

func TestClient_DeleteAppMonitorAlarmConfig(t *testing.T) {
	req := &model.DeleteAppMonitorAlarmConfigRequest{
		UserId:    bcmConf.UserId,
		AppName:   "zmq-log-1115",
		AlarmName: "test_wjr",
	}
	err := bcmClient.DeleteAppMonitorAlarmConfig(req)
	if err != nil {
		t.Errorf("bcm delete alarm config for application error with %v\n", err)
	}
}

func TestClient_GetAppMonitorAlarmConfig(t *testing.T) {
	req := &model.GetAppMonitorAlarmConfigDetailRequest{
		UserId:    bcmConf.UserId,
		AlarmName: "config-yyy",
		AppName:   "yyy-test",
	}
	resp, err := bcmClient.GetAppMonitorAlarmConfig(req)
	if err != nil {
		t.Errorf("bcm get alarm config for application error with %v\n", err)
	}
	fmt.Println(*resp)
}

func TestClient_GetAlarmMetricsForApplication(t *testing.T) {
	req := &model.GetAppMonitorAlarmMetricsRequest{
		UserId:   bcmConf.UserId,
		AppName:  "gjm-app-test",
		TaskName: "eecb448f083447498012cff4473c7ea1",
	}
	resp, err := bcmClient.GetAlarmMetricsForApplication(req)
	if err != nil {
		t.Errorf("bcm get alarm metrics for application error with %v\n", err)
	}
	fmt.Println(resp)
}

func TestClient_CreateDashboard(t *testing.T) {
	req := &model.DashboardRequest{
		Configure: "{\"tabs\":[{\"dimensions\":[],\"metric\":[],\"name\":\"\",\"namespace\":[],\"widgets\":" +
			"[[{\"name\":\"_54382_54383\"},{\"name\":\"_54382_54384\"}," +
			"{\"name\":\"_54382_54385\"}],[{\"name\":\"_54382_54386\"}]]}]}",
		Title:  "yyy-test-new",
		Type:   "common",
		UserID: bcmConf.UserId,
	}
	resp, err := bcmClient.CreateDashboard(req)
	if err != nil {
		t.Errorf("bcm create dashboard error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_DeleteDashboard(t *testing.T) {
	req := &model.DashboardRequest{
		UserID:        bcmConf.UserId,
		DashboardName: "_54439",
	}
	resp, err := bcmClient.DeleteDashboard(req)
	if err != nil {
		t.Errorf("bcm delete dashboard error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateDashboard(t *testing.T) {
	req := &model.DashboardRequest{
		UserID:        bcmConf.UserId,
		DashboardName: "_54405",
		Configure: "{\"tabs\":[{\"dimensions\":[],\"metric\":[],\"name\":\"\",\"namespace\":[]," +
			"\"widgets\":[[{\"name\":\"_54382_54383\"},{\"name\":\"_54382_54384\"}," +
			"{\"name\":\"_54382_54385\"}],[{\"name\":\"_54382_54386\"}]]}]}",
		Title: "yyy-test-goSdk-update_new",
	}
	resp, err := bcmClient.UpdateDashboard(req)
	if err != nil {
		t.Errorf("bcm update dashboard error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboard(t *testing.T) {
	req := &model.DashboardRequest{
		UserID:        bcmConf.UserId,
		DashboardName: "_54405",
	}
	resp, err := bcmClient.GetDashboard(req)
	if err != nil {
		t.Errorf("bcm get dashboard error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_DuplicateDashboard(t *testing.T) {
	req := &model.DashboardRequest{
		UserID:        bcmConf.UserId,
		DashboardName: "_54405",
	}
	resp, err := bcmClient.DuplicateDashboard(req)
	if err != nil {
		t.Errorf("bcm duplicate dashboard error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateDashboardWidget(t *testing.T) {
	req := &model.DashboardWidgetRequest{
		DashboardName: "_54440",
		UserID:        bcmConf.UserId,
	}
	resp, err := bcmClient.CreateDashboardWidget(req)
	if err != nil {
		t.Errorf("bcm create dashboard widget error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardWidget(t *testing.T) {
	req := &model.DashboardWidgetRequest{
		DashboardName: "_54440",
		UserID:        bcmConf.UserId,
		WidgetName:    "_54440_54445",
	}
	resp, err := bcmClient.GetDashboardWidget(req)
	if err != nil {
		t.Errorf("bcm get dashboard widget error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_DeleteDashboardWidget(t *testing.T) {
	req := &model.DashboardWidgetRequest{
		DashboardName: "_54440",
		UserID:        bcmConf.UserId,
		WidgetName:    "_54440_54453",
	}
	resp, err := bcmClient.DeleteDashboardWidget(req)
	if err != nil {
		t.Errorf("bcm delete dashboard widget error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateDashboardWidget(t *testing.T) {

	ConfigJsonStr := "{\"data\":[{\"metric\":[{\"name\":\"CpuIdlePercent\",\"unit\":\"%\"," +
		"\"alias\":\"CPU空闲率\",\"contrast\":[],\"timeContrast\":[],\"statistics\":\"avg\"}]," +
		"\"monitorObject\":[{\"instanceName\":\"zmq-as0001 \",\"id\":\"i-yq8qU5Qf\"}]," +
		"\"scope\":\"BCE_BCC\",\"subService\":\"linux\",\"region\":\"bj\"," +
		"\"scopeValue\":{\"name\":\"BCC\",\"value\":\"BCE_BCC\",\"hasChildren\":false}," +
		"\"resourceType\":\"Instance\",\"monitorType\":\"scope\"," +
		"\"namespace\":[{\"namespaceType\":\"instance\",\"transfer\":\"\",\"filter\":\"\"," +
		"\"name\":\"i-yq8qU5Qf___bj.BCE_BCC.453bf9588c9e488f9ba2c984129090dc\"," +
		"\"instanceName\":\"zmq-as0001 \",\"region\":\"bj\",\"bcmService\":\"BCE_BCC\"," +
		"\"subService\":[{\"name\":\"serviceType\",\"value\":\"linux\"}]}]," +
		"\"product\":\"453bf9588c9e488f9ba2c984129090dc\"}],\"style\":{\"displayType\":\"line\"," +
		"\"nullPointMode\":\"zero\",\"threshold\":0,\"decimals\":2,\"isEdit\":true,\"unit\":\"%\"}," +
		"\"title\":\"bccNew3\",\"timeRange\":{\"timeType\":\"dashboard\",\"unit\":\"minutes\"," +
		"\"number\":1,\"relative\":\"today()\"},\"time\":\"\",\"monitorType\":\"scope\"}"

	var data model.DashboardWidgetConfigure
	err := json.Unmarshal([]byte(ConfigJsonStr), &data)
	req := &model.DashboardWidgetRequest{
		DashboardName: "_54440",
		UserID:        bcmConf.UserId,
		WidgetName:    "_54440_54445",
		Title:         "bccNew3",
		Configure:     data,
		Type:          "trend",
	}
	resp, err := bcmClient.UpdateDashboardWidget(req)
	fmt.Printf("err is %v\n", err)
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateDashboardWidget2(t *testing.T) {
	req := &model.DashboardWidgetRequest{
		DashboardName: "_54440",
		UserID:        bcmConf.UserId,
		WidgetName:    "_54440_54445",
		Title:         "bccNew4",
		Configure: model.DashboardWidgetConfigure{
			Data: []model.Data{
				{
					Metric: []model.DashboardMetric{
						{
							Name:         "CpuIdlePercent",
							Unit:         "%",
							Alias:        "CPU空闲率",
							Contrast:     []string{},
							TimeContrast: []string{},
							Statistics:   "avg",
						},
					},
					MonitorObject: []model.DashboardMonitorObject{
						{
							InstanceName: "zmq-as0001",
							Id:           "i-yq8qU5Qf",
						},
					},
					Scope:      "BCE_BCC",
					SubService: "linux",
					Region:     "bj",
					ScopeValue: model.ScopeValue{
						Name:        "BCC",
						Value:       "BCE_BCC",
						HasChildren: false,
					},
					ResourceType: "Instance",
					MonitorType:  "scope",
					Namespace: []model.DashboardNamespace{
						{
							NamespaceType: "instance",
							Transfer:      "",
							Filter:        "",
							Name:          "i-yq8qU5Qf___bj.BCE_BCC.453bf9588c9e488f9ba2c984129090dc",
							InstanceName:  "zmq-as0001 ",
							Region:        "bj",
							BcmService:    "BCE_BCC",
							SubService: []model.SubService{
								{
									Name:  "serviceType",
									Value: "linux",
								},
							},
						},
					},
					Product: bcmConf.UserId,
				},
			},
			Style: model.Style{
				DisplayType:   "line",
				NullPointMode: "zero",
				Threshold:     0,
				Decimals:      2,
				IsEdit:        true,
				Unit:          "%",
			},
			Title: "bccNew4",
			TimeRange: model.TimeRange{
				TimeType: "dashboard",
				Unit:     "minutes",
				Number:   1,
				Relative: "today()",
			},
			Time:        "",
			MonitorType: "scope",
		},
		Type: "trend",
	}
	resp, err := bcmClient.UpdateDashboardWidget(req)
	if err != nil {
		t.Errorf("bcm update dashboard widget error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_DuplicateDashboardWidget(t *testing.T) {
	req := &model.DashboardWidgetRequest{
		DashboardName: "_54440",
		UserID:        bcmConf.UserId,
		WidgetName:    "_54440_54441",
	}
	resp, err := bcmClient.DuplicateDashboardWidget(req)
	if err != nil {
		t.Errorf("bcm duplicate dashboard widget error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardReportSingleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-08 09:10:59|2023-12-08 10:10:59",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "CpuIdlePercent",
						Unit:         "%",
						Alias:        "CPU空闲率",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "instance-xcy9049y ",
						Id:           "i-isvkUW76",
					},
				},
				Scope:      "BCE_BCC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "app",
						Transfer:      "",
						Filter:        "",
						Name:          "i-isvkUW76___bj.BCE_BCC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName:  "instance-xcy9049y ",
						Region:        "bj",
						BcmService:    "BCE_BCC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	reqStr, _ := json.Marshal(req)
	fmt.Println(string(reqStr))
	resp, err := bcmClient.GetDashboardReportData(req)
	if err != nil {
		t.Errorf("bcm get dashboard report data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardReportMultipleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-08 09:10:59|2023-12-08 10:10:59",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "vNicInBytes",
						Unit:         "Bytes",
						Alias:        "网卡输入流量",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
						Cycle:        60,
						Dimensions: []string{
							"eth1", "eth2",
						},
						MetricDimensions: []model.MetricDimensions{
							{
								Name: "nicName",
								Values: []string{
									"eth1",
									"eth2",
								},
							},
						},
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "prod.nmp.nn.yd1",
						Id:           "41b372b8-3acc-423c-a6b0-af5c69fd1c41",
					},
				},
				Scope:      "BCE_BEC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "app",
						Transfer:      "",
						Filter:        "",
						Name: "41b372b8-3acc-423c-a6b0-af5c69fd1c41___bj." +
							"BCE_BEC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName: "prod.nmp.nn.yd1 ",
						Region:       "bj",
						BcmService:   "BCE_BEC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardReportData(req)
	if err != nil {
		t.Errorf("bcm get dashboard report data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardTrendSingleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-08 13:14:55|2023-12-08 13:15:55",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "CPUUsagePercent",
						Unit:         "%",
						Alias:        "CPU使用率",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "instance-j7wp0ieh ",
						Id:           "i-CUL6HOp1",
					},
				},
				Scope:      "BCE_BCC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "instance",
						Transfer:      "",
						Filter:        "",
						Name:          "i-CUL6HOp1___bj.BCE_BCC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName:  "instance-j7wp0ieh ",
						Region:        "bj",
						BcmService:    "BCE_BCC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardTrendData(req)
	if err != nil {
		t.Errorf("bcm get dashboard trend data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardTrendMultipleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-07 13:36:15|2023-12-07 13:38:15",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "vNicInBytes",
						Unit:         "Bytes",
						Alias:        "网卡输入流量",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
						Cycle:        60,
						Dimensions: []string{
							"eth1", "eth2",
						},
						MetricDimensions: []model.MetricDimensions{
							{
								Name: "nicName",
								Values: []string{
									"eth1",
									"eth2",
								},
							},
						},
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "prod.nmp.nn.yd1",
						Id:           "41b372b8-3acc-423c-a6b0-af5c69fd1c41",
					},
				},
				Scope:      "BCE_BEC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "app",
						Transfer:      "",
						Filter:        "",
						Name: "41b372b8-3acc-423c-a6b0-af5c69fd1c41___bj." +
							"BCE_BEC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName: "prod.nmp.nn.yd1 ",
						Region:       "bj",
						BcmService:   "BCE_BEC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardTrendData(req)
	if err != nil {
		t.Errorf("bcm get dashboard trend data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardGaugechartSingleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-08 13:14:55|2023-12-08 13:15:55",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "CPUUsagePercent",
						Unit:         "%",
						Alias:        "CPU使用率",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "instance-j7wp0ieh ",
						Id:           "i-CUL6HOp1",
					},
				},
				Scope:      "BCE_BCC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "instance",
						Transfer:      "",
						Filter:        "",
						Name:          "i-CUL6HOp1___bj.BCE_BCC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName:  "instance-j7wp0ieh ",
						Region:        "bj",
						BcmService:    "BCE_BCC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardGaugeChartData(req)
	if err != nil {
		t.Errorf("bcm get dashboard gaugechart data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardGaugechartMultipleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-07 13:36:15|2023-12-07 14:38:15",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "vNicInBytes",
						Unit:         "Bytes",
						Alias:        "网卡输入流量",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
						Cycle:        60,
						Dimensions: []string{
							"eth1", "eth0",
						},
						MetricDimensions: []model.MetricDimensions{
							{
								Name: "nicName",
								Values: []string{
									"eth1",
									"eth0",
								},
							},
						},
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "prod.nmp.nn.yd1",
						Id:           "41b372b8-3acc-423c-a6b0-af5c69fd1c41",
					},
				},
				Scope:      "BCE_BEC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "app",
						Transfer:      "",
						Filter:        "",
						Name: "41b372b8-3acc-423c-a6b0-af5c69fd1c41___bj." +
							"BCE_BEC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName: "prod.nmp.nn.yd1 ",
						Region:       "bj",
						BcmService:   "BCE_BEC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardGaugeChartData(req)
	if err != nil {
		t.Errorf("bcm get dashboard gaugechart data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardBillboardSingleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-08 13:14:55|2023-12-08 13:15:55",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "CPUUsagePercent",
						Unit:         "%",
						Alias:        "CPU使用率",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "instance-j7wp0ieh ",
						Id:           "i-CUL6HOp1",
					},
				},
				Scope:      "BCE_BCC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "instance",
						Transfer:      "",
						Filter:        "",
						Name:          "i-CUL6HOp1___bj.BCE_BCC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName:  "instance-j7wp0ieh ",
						Region:        "bj",
						BcmService:    "BCE_BCC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardBillboardData(req)
	if err != nil {
		t.Errorf("bcm get dashboard billboard data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardBillboardMultipleData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-07 13:36:15|2023-12-07 14:38:15",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "vNicInBytes",
						Unit:         "Bytes",
						Alias:        "网卡输入流量",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
						Cycle:        60,
						Dimensions: []string{
							"eth1", "eth0",
						},
						MetricDimensions: []model.MetricDimensions{
							{
								Name: "nicName",
								Values: []string{
									"eth1",
									"eth0",
								},
							},
						},
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "prod.nmp.nn.yd1",
						Id:           "41b372b8-3acc-423c-a6b0-af5c69fd1c41",
					},
				},
				Scope:      "BCE_BEC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "app",
						Transfer:      "",
						Filter:        "",
						Name: "41b372b8-3acc-423c-a6b0-af5c69fd1c41___bj" +
							".BCE_BEC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName: "prod.nmp.nn.yd1 ",
						Region:       "bj",
						BcmService:   "BCE_BEC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardBillboardData(req)
	if err != nil {
		t.Errorf("bcm get dashboard billboard data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardTrendSeniorData(t *testing.T) {
	req := &model.DashboardDataRequest{
		Time: "2023-12-08 14:41:50|2023-12-08 14:42:50",
		Data: []model.Data{
			{
				Metric: []model.DashboardMetric{
					{
						Name:         "CPUUsagePercent",
						Unit:         "%",
						Alias:        "CPU使用率",
						Contrast:     []string{},
						TimeContrast: []string{},
						Statistics:   "avg",
					},
				},
				MonitorObject: []model.DashboardMonitorObject{
					{
						InstanceName: "instance-j7wp0ieh ",
						Id:           "i-CUL6HOp1",
					},
				},
				Scope:      "BCE_BCC",
				SubService: "linux",
				Region:     "bj",
				ScopeValue: model.ScopeValue{
					Name:        "BCC",
					Value:       "BCE_BCC",
					HasChildren: false,
				},
				MonitorType: "scope",
				Namespace: []model.DashboardNamespace{
					{
						NamespaceType: "instance",
						Transfer:      "",
						Filter:        "",
						Name:          "i-CUL6HOp1___bj.BCE_BCC.a0d04d7c202140cb80155ff7b6752ce4",
						InstanceName:  "instance-j7wp0ieh ",
						Region:        "bj",
						BcmService:    "BCE_BCC",
						SubService: []model.SubService{
							{
								Name:  "serviceType",
								Value: "linux",
							},
						},
					},
				},
				Product: bcmConf.UserId,
			},
		},
	}
	resp, err := bcmClient.GetDashboardTrendSeniorData(req)
	if err != nil {
		t.Errorf("bcm get dashboard trend senior data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetDashboardDimensionsData(t *testing.T) {
	req := &model.DashboardDimensionsRequest{
		UserID:     bcmConf.UserId,
		Dimensions: "nicName",
		Service:    "BCE_BEC",
		Region:     "bj",
		ResourceID: "7744b3f3-ec04-459a-b3ae-4379111534ff",
		MetricName: "vNicInBytes",
	}
	resp, err := bcmClient.GetDashboardDimensions(req)
	if err != nil {
		t.Errorf("bcm get dashboard dimensions data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateAlarmPolicy(t *testing.T) {
	req := &model.AlarmConfig{
		AliasName: "test_alarm_policy_01",
		Level:     model.AlarmLevelCritical,
		MonitorObject: &model.AlarmMonitorObject{
			MonitorType: model.MonitorObjectInstance,
			Names:       []string{"InstanceId:i-mPkY5Z**"},
			TypeName:    "Instance",
		},
		IncidentActions:     []string{"d242711b-****-****-****-b8ac175f8e7d"},
		ResumeAction:        nil,
		InsufficientActions: nil,
		Region:              "bj",
		Scope:               "BCE_BCC",
		UserId:              bcmConf.UserId,
		Rules: [][]*model.AlarmRule{
			{
				{
					Index:                 1,
					Metric:                "CPUUsagePercent",
					PeriodInSecond:        60,
					Statistics:            "average",
					Threshold:             "123456",
					ComparisonOperator:    ">",
					EvaluationPeriodCount: 1,
				},
			},
		},
		AlarmType:        "NORMAL",
		AlarmDescription: "这是一个测试的报警策略",
		SrcType:          "INSTANCE",
	}
	err := bcmClient.CreateAlarmPolicy(req)
	if err != nil {
		t.Errorf("bcm create alarm policy error with %v\n", err)
	}
}

func Test_UpdateAlarmPolicy(t *testing.T) {
	req := &model.AlarmConfig{
		AliasName: "test_alarm_policy_02",
		AlarmName: "030814********************3bc898",
		Level:     model.AlarmLevelCritical,
		MonitorObject: &model.AlarmMonitorObject{
			MonitorType: model.MonitorObjectInstance,
			Names:       []string{"InstanceId:i-mPkY5Z**"},
			TypeName:    "Instance",
		},
		IncidentActions:     []string{"d242711b-****-****-****-b8ac175f8e7d"},
		ResumeAction:        nil,
		InsufficientActions: nil,
		Region:              "bj",
		Scope:               "BCE_BCC",
		UserId:              bcmConf.UserId,
		Rules: [][]*model.AlarmRule{
			{
				{
					Index:                 1,
					Metric:                "CPUUsagePercent",
					PeriodInSecond:        60,
					Statistics:            "average",
					Threshold:             "123456",
					ComparisonOperator:    ">",
					EvaluationPeriodCount: 1,
				},
			},
		},
		AlarmType:        "NORMAL",
		AlarmDescription: "这是一个测试的报警策略",
		SrcType:          "INSTANCE",
	}
	err := bcmClient.UpdateAlarmPolicy(req)
	if err != nil {
		t.Errorf("bcm update alarm policy error with %v\n", err)
	}
}

func Test_BlockAlarmPolicy(t *testing.T) {
	req := &model.CommonAlarmConfigRequest{
		AlarmName: "030814********************3bc898",
		Scope:     "BCE_BCC",
		UserId:    bcmConf.UserId,
	}
	err := bcmClient.BlockAlarmPolicy(req)
	if err != nil {
		t.Errorf("bcm block alarm policy error with %v\n", err)
	}
}

func Test_UnblockAlarmPolicy(t *testing.T) {
	req := &model.CommonAlarmConfigRequest{
		AlarmName: "030814********************3bc898",
		Scope:     "BCE_BCC",
		UserId:    bcmConf.UserId,
	}
	err := bcmClient.UnblockAlarmPolicy(req)
	if err != nil {
		t.Errorf("bcm unblock alarm policy error with %v\n", err)
	}
}

func Test_GetAlarmPolicyDetail(t *testing.T) {
	req := &model.CommonAlarmConfigRequest{
		AlarmName: "030814********************3bc898",
		Scope:     "BCE_BCC",
		UserId:    bcmConf.UserId,
	}
	resp, err := bcmClient.GetAlarmPolicyDetail(req)
	if err != nil {
		t.Errorf("bcm get alarm policy detail error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_ListSingleInstanceAlarmConfigs(t *testing.T) {
	req := &model.ListSingleInstanceAlarmConfigsRequest{
		Scope:    "BCE_BCC",
		UserId:   bcmConf.UserId,
		PageNo:   1,
		PageSize: 10,
	}
	resp, err := bcmClient.ListSingleInstanceAlarmConfigs(req)
	if err != nil {
		t.Errorf("bcm get single instance alarm policys error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_ListAlarmMetrics(t *testing.T) {
	req := &model.ListAlarmMetricsRequest{
		Scope:  "BCE_BCC",
		UserId: bcmConf.UserId,
		Region: "bj",
	}
	resp, err := bcmClient.ListAlarmMetrics(req)
	if err != nil {
		t.Errorf("bcm list alarm metrics error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_DeleteAlarmPolicy(t *testing.T) {
	req := &model.CommonAlarmConfigRequest{
		AlarmName: "030814********************3bc898",
		Scope:     "BCE_BCC",
		UserId:    bcmConf.UserId,
	}
	err := bcmClient.DeleteAlarmPolicy(req)
	if err != nil {
		t.Errorf("bcm delete alarm policy error with %v\n", err)
	}
}

func Test_CreateAlarmPolicyV2(t *testing.T) {
	req := &model.AlarmConfigV2{
		AliasName:    "test_alarm_policy_01",
		ResourceType: "Instance",
		AlarmLevel:   model.AlarmLevelCritical,
		TargetType:   model.TargetTypeMultiInstances,
		Region:       "bj",
		Scope:        "BCE_BCC",
		UserId:       bcmConf.UserId,
		TargetInstances: []*model.AlarmInstance{
			{
				Region: "bj",
				Identifiers: []*model.KV{
					{
						Key:   "InstanceId",
						Value: "i-mPkY5Z**",
					},
				},
			},
		},
		Policies: []*model.AlarmPolicy{
			{
				AlarmPendingPeriodCount: 1,
				Rules: []*model.AlarmRuleV2{
					{
						MetricName: "CPUUsagePercent",
						Operator:   ">",
						Statistics: "average",
						Threshold:  12345.0,
						Window:     60,
					},
				},
			},
		},
		Actions: []*model.AlarmAction{
			{
				Name: "test_yangmoda",
			},
		},
	}
	resp, err := bcmClient.CreateAlarmPolicyV2(req)
	if err != nil {
		t.Errorf("bcm create alarm policy v2 error with %v\n", err)
	}
	fmt.Println(resp.Result.AlarmName)
}

func Test_UpdateAlarmPolicyV2(t *testing.T) {
	req := &model.AlarmConfigV2{
		AlarmName:    "030814********************3bc898",
		AliasName:    "test_alarm_policy_02",
		ResourceType: "Instance",
		AlarmLevel:   model.AlarmLevelCritical,
		TargetType:   model.TargetTypeMultiInstances,
		Region:       "bj",
		Scope:        "BCE_BCC",
		UserId:       bcmConf.UserId,
		TargetInstances: []*model.AlarmInstance{
			{
				Region: "bj",
				Identifiers: []*model.KV{
					{
						Key:   "InstanceId",
						Value: "i-mPkY5Z**",
					},
				},
			},
		},
		Policies: []*model.AlarmPolicy{
			{
				AlarmPendingPeriodCount: 1,
				Rules: []*model.AlarmRuleV2{
					{
						MetricName: "CPUUsagePercent",
						Operator:   ">",
						Statistics: "average",
						Threshold:  12345.0,
						Window:     60,
					},
				},
			},
		},
		Actions: []*model.AlarmAction{
			{
				Name: "test_yangmoda",
			},
		},
	}
	err := bcmClient.UpdateAlarmPolicyV2(req)
	if err != nil {
		t.Errorf("bcm update alarm policy v2 error with %v\n", err)
	}
}

func Test_BlockAlarmPolicyV2(t *testing.T) {
	req := &model.CommonAlarmConfigRequest{
		AlarmName: "030814********************3bc898",
		Scope:     "BCE_BCC",
		UserId:    bcmConf.UserId,
	}
	err := bcmClient.BlockAlarmPolicyV2(req)
	if err != nil {
		t.Errorf("bcm block alarm policy v2 error with %v\n", err)
	}
}

func Test_UnblockAlarmPolicyV2(t *testing.T) {
	req := &model.CommonAlarmConfigRequest{
		AlarmName: "030814********************3bc898",
		Scope:     "BCE_BCC",
		UserId:    bcmConf.UserId,
	}
	err := bcmClient.UnblockAlarmPolicyV2(req)
	if err != nil {
		t.Errorf("bcm unblock alarm policy v2 error with %v\n", err)
	}
}

func Test_GetAlarmPolicyDetailV2(t *testing.T) {
	req := &model.CommonAlarmConfigRequest{
		AlarmName: "030814********************3bc898",
		Scope:     "BCE_BCC",
		UserId:    bcmConf.UserId,
	}
	resp, err := bcmClient.GetAlarmPolicyDetailV2(req)
	if err != nil {
		t.Errorf("bcm get alarm policy detail v2 error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSiteHttpTask(t *testing.T) {
	req := &model.CreateHTTPTask{
		UserID:        bcmConf.UserId,
		TaskName:      "taskName",
		Address:       "address",
		Method:        "get",
		Cycle:         60,
		Idc:           "beijing-UNICOM",
		Timeout:       20,
		IPType:        "ipv4",
		AdvanceConfig: strconv.FormatBool(false),
	}
	resp, err := bcmClient.CreateSiteHttpTask(req)
	if err != nil {
		t.Errorf("create http site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateSiteHttpTask(t *testing.T) {
	req := &model.CreateHTTPTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		Address:  "address",
		Method:   "get",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  20,
	}
	resp, err := bcmClient.UpdateSiteHttpTask(req)
	if err != nil {
		t.Errorf("update http site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteHttpTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "xEvohDfLjMWZzkjgFiuHtqqrlKhdrYaI",
	}
	resp, err := bcmClient.GetSiteHttpTask(req)
	if err != nil {
		t.Errorf("get http site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSiteHttpsTask(t *testing.T) {
	req := &model.CreateHTTPSTask{
		UserID:        bcmConf.UserId,
		TaskName:      "taskName",
		Address:       "address",
		Method:        "get",
		Cycle:         60,
		Idc:           "beijing-UNICOM",
		Timeout:       20,
		IPType:        "ipv4",
		AdvanceConfig: strconv.FormatBool(false),
	}
	resp, err := bcmClient.CreateSiteHttpsTask(req)
	if err != nil {
		t.Errorf("create https site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateSiteHttpsTask(t *testing.T) {
	req := &model.CreateHTTPTask{
		UserID:        bcmConf.UserId,
		TaskID:        "mPzOMwuQsOcfGATDBIJHHWXsgcPCOMnd",
		TaskName:      "taskName",
		Address:       "address",
		Method:        "get",
		Cycle:         60,
		Idc:           "beijing-UNICOM",
		Timeout:       10,
		AdvanceConfig: strconv.FormatBool(false),
	}
	resp, err := bcmClient.UpdateSiteHttpTask(req)
	if err != nil {
		t.Errorf("update https site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteHttpsTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
	}
	resp, err := bcmClient.GetSiteHttpsTask(req)
	if err != nil {
		t.Errorf("get https site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSitePingTask(t *testing.T) {
	req := &model.CreatePingTask{
		UserID:      bcmConf.UserId,
		TaskName:    "taskName",
		Address:     "www.baidu.com",
		Cycle:       60,
		Idc:         "beijing-UNICOM",
		Timeout:     20,
		IPType:      "ipv4",
		PacketCount: 1,
	}
	resp, err := bcmClient.CreateSitePingTask(req)
	if err != nil {
		t.Errorf("create ping site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateSitePingTask(t *testing.T) {
	req := &model.CreatePingTask{
		UserID:      bcmConf.UserId,
		TaskName:    "taskName",
		TaskID:      "UQQmhYlGbPeqSfAFRSvFXMfuvuvUOmpZ",
		Address:     "www.baidu.com",
		Cycle:       60,
		Idc:         "beijing-UNICOM",
		Timeout:     10,
		IPType:      "ipv4",
		PacketCount: 1,
	}
	resp, err := bcmClient.UpdateSitePingTask(req)
	if err != nil {
		t.Errorf("update ping site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSitePingTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "UQQmhYlGbPeqSfAFRSvFXMfuvuvUOmpZ",
	}
	resp, err := bcmClient.GetSitePingTask(req)
	if err != nil {
		t.Errorf("get ping site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSiteTCPTask(t *testing.T) {
	req := &model.CreateTCPTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  5,
		IPType:   "ipv4",
		Port:     80,
	}
	resp, err := bcmClient.CreateSiteTcpTask(req)
	if err != nil {
		t.Errorf("create tcp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateSiteTCPTask(t *testing.T) {
	req := &model.CreateTCPTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		TaskID:   "bFcjzwTdSsILYsVsXdzRsqngOUzQhLSx",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  6,
		IPType:   "ipv4",
		Port:     80,
	}
	resp, err := bcmClient.UpdateSiteTcpTask(req)
	if err != nil {
		t.Errorf("update tcp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteTCPTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "bFcjzwTdSsILYsVsXdzRsqngOUzQhLSx",
	}
	resp, err := bcmClient.GetSiteTcpTask(req)
	if err != nil {
		t.Errorf("get tcp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSiteUDPTask(t *testing.T) {
	req := &model.CreateUDPTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  5,
		IPType:   "ipv4",
		Port:     80,
	}
	resp, err := bcmClient.CreateSiteUdpTask(req)
	if err != nil {
		t.Errorf("create udp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateSiteUDPTask(t *testing.T) {
	req := &model.CreateUDPTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		TaskID:   "EEZpttAeAsvmGWKpGPALkdqytWgMOgxK",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  6,
		IPType:   "ipv4",
		Port:     80,
	}
	resp, err := bcmClient.UpdateSiteUdpTask(req)
	if err != nil {
		t.Errorf("update udp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteUDPTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "EEZpttAeAsvmGWKpGPALkdqytWgMOgxK",
	}
	resp, err := bcmClient.GetSiteTcpTask(req)
	if err != nil {
		t.Errorf("get udp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSiteFptTask(t *testing.T) {
	req := &model.CreateFtpTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  5,
		IPType:   "ipv4",
		Port:     80,
	}
	resp, err := bcmClient.CreateSiteFtpTask(req)
	if err != nil {
		t.Errorf("create ftp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateSiteFtpTask(t *testing.T) {
	req := &model.CreateFtpTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		TaskID:   "GbraRSihmGMDTQPZntYAfagGfBYZyyfO",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  6,
		IPType:   "ipv4",
		Port:     80,
	}
	resp, err := bcmClient.UpdateSiteFtpTask(req)
	if err != nil {
		t.Errorf("update ftp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteFtpTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "GbraRSihmGMDTQPZntYAfagGfBYZyyfO",
	}
	resp, err := bcmClient.GetSiteFtpTask(req)
	if err != nil {
		t.Errorf("get ftp site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSiteDnsTask(t *testing.T) {
	req := &model.CreateDNSTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  5,
		IPType:   "ipv4",
	}
	resp, err := bcmClient.CreateSiteDnsTask(req)
	if err != nil {
		t.Errorf("create dns site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_UpdateSiteDnsTask(t *testing.T) {
	req := &model.CreateDNSTask{
		UserID:   bcmConf.UserId,
		TaskName: "taskName",
		TaskID:   "PSFowPmtxFQAnauTQBpKovcfDXVnWYnR",
		Address:  "www.baidu.com",
		Cycle:    60,
		Idc:      "beijing-UNICOM",
		Timeout:  6,
		IPType:   "ipv4",
	}
	resp, err := bcmClient.UpdateSiteDnsTask(req)
	if err != nil {
		t.Errorf("update dns site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteDnsTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "PSFowPmtxFQAnauTQBpKovcfDXVnWYnR",
	}
	resp, err := bcmClient.GetSiteDnsTask(req)
	if err != nil {
		t.Errorf("get dns site task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteTaskList(t *testing.T) {
	req := &model.GetTaskListRequest{
		UserID:   bcmConf.UserId,
		PageNo:   1,
		PageSize: 10,
	}
	resp, err := bcmClient.GetSiteTaskList(req)
	if err != nil {
		t.Errorf("get site task list error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_DeleteSiteTask(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "kaNmzlWBbBdOyywuHrnTHnsRRUpQJNaQ",
	}
	resp, err := bcmClient.DeleteSiteTask(req)
	if err != nil {
		t.Errorf("delete sita task error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteTaskDetail(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "tTeqkwJeqNTxJnoPJoltkQZCoNfKJhnS",
	}
	resp, err := bcmClient.GetSiteTaskDetail(req)
	if err != nil {
		t.Errorf("get site task detail error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateSiteAlarmConfig(t *testing.T) {
	rule := &model.SiteAlarmRule{
		Metric:             "connectTime",
		MetricAlias:        "建连时间",
		Statistics:         "average",
		Threshold:          "10000",
		ComparisonOperator: ">",
		Cycle:              60,
		Count:              1,
		Function:           "THRESHOLD",
		ActOnIdcs:          []string{"average"},
	}
	req := &model.CreateSiteAlarmConfigRequest{
		UserID:              bcmConf.UserId,
		TaskID:              "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
		AliasName:           "test_site",
		InsufficientActions: []string{"37eddd21-a44c-42c6-b0bb-a3e9d738a091"},
		Level:               model.LevelNotice,
		Namespace:           "BCM_SITE",
		Region:              "bj",
		ActionEnabled:       strconv.FormatBool(false),
		Rules:               []model.SiteAlarmRule{*rule},
		Cycle:               60,
	}
	err := bcmClient.CreateSiteAlarmConfig(req)
	if err != nil {
		t.Errorf("create site alarm config error with %v\n", err)
	}
}

func Test_UpdateSiteAlarmConfig(t *testing.T) {
	rule := &model.SiteAlarmRule{
		Metric:             "connectTime",
		MetricAlias:        "建连时间",
		Statistics:         "average",
		Threshold:          "20000",
		ComparisonOperator: ">",
		Cycle:              60,
		Count:              1,
		Function:           "THRESHOLD",
		ActOnIdcs:          []string{"average"},
	}
	req := &model.CreateSiteAlarmConfigRequest{
		UserID:              bcmConf.UserId,
		AlarmName:           "ec23c80fab8a4b9e8276f12945c09e9f",
		TaskID:              "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
		AliasName:           "test_site",
		InsufficientActions: []string{"37eddd21-a44c-42c6-b0bb-a3e9d738a091"},
		Level:               model.LevelNotice,
		Namespace:           "BCM_SITE",
		Region:              "bj",
		ActionEnabled:       strconv.FormatBool(false),
		Rules:               []model.SiteAlarmRule{*rule},
		Cycle:               60,
	}
	err := bcmClient.UpdateSiteAlarmConfig(req)
	if err != nil {
		t.Errorf("update site alarm config error with %v\n", err)
	}
}

func Test_DeleteSiteAlarmConfig(t *testing.T) {
	req := &model.DeleteSiteAlarmConfigRequest{
		UserID:     bcmConf.UserId,
		AlarmNames: []string{"ec23c80fab8a4b9e8276f12945c09e9f"},
	}
	err := bcmClient.DeleteSiteAlarmConfig(req)
	if err != nil {
		t.Errorf("delete site alarm config error with %v\n", err)
	}
}

func Test_GetSiteAlarmConfigDetail(t *testing.T) {
	req := &model.GetSiteAlarmConfigRequest{
		UserID:    bcmConf.UserId,
		AlarmName: "08f7d132e65547d8b3795e3b046d24a9",
	}
	resp, err := bcmClient.GetSiteAlarmConfigDetail(req)
	if err != nil {
		t.Errorf("get site alarm config detail error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteAlarmConfigList(t *testing.T) {
	req := &model.GetSiteAlarmConfigListRequest{
		UserID:   bcmConf.UserId,
		TaskID:   "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
		PageNo:   1,
		PageSize: 10,
	}
	resp, err := bcmClient.GetSiteAlarmConfigList(req)
	if err != nil {
		t.Errorf("get site alarm config list error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_BlockSiteAlarmConfig(t *testing.T) {
	req := &model.GetSiteAlarmConfigRequest{
		UserID:    bcmConf.UserId,
		AlarmName: "08f7d132e65547d8b3795e3b046d24a9",
		Namespace: "BCM_SITE",
	}
	err := bcmClient.BlockSiteAlarmConfig(req)
	if err != nil {
		t.Errorf("block site alarm config error with %v\n", err)
	}
}

func Test_UnBlockSiteAlarmConfig(t *testing.T) {
	req := &model.GetSiteAlarmConfigRequest{
		UserID:    bcmConf.UserId,
		AlarmName: "08f7d132e65547d8b3795e3b046d24a9",
		Namespace: "BCM_SITE",
	}
	err := bcmClient.UnBlockSiteAlarmConfig(req)
	if err != nil {
		t.Errorf("unblock site alarm config error with %v\n", err)
	}
}

func Test_GetTaskByAlarmName(t *testing.T) {
	req := &model.GetSiteAlarmConfigRequest{
		UserID:    bcmConf.UserId,
		AlarmName: "08f7d132e65547d8b3795e3b046d24a9",
		Namespace: "BCM_SITE",
	}
	resp, err := bcmClient.GetTaskByAlarmName(req)
	if err != nil {
		t.Errorf("get site task by alarm error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteMetricData(t *testing.T) {
	req := &model.GetSiteMetricDataRequest{
		UserID:     bcmConf.UserId,
		TaskID:     "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
		MetricName: "success",
		StartTime:  "2024-01-25T08:17:54Z",
		EndTime:    "2024-01-25T09:17:54Z",
		Statistics: []string{"average"},
		Cycle:      60,
	}
	resp, err := bcmClient.GetSiteMetricData(req)
	if err != nil {
		t.Errorf("get site metric data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteOverallView(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
	}
	resp, err := bcmClient.GetSiteOverallView(req)
	if err != nil {
		t.Errorf("get site overall view data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteProvincialView(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
		Isp:    "guangxi",
	}
	resp, err := bcmClient.GetSiteProvincialView(req)
	if err != nil {
		t.Errorf("get site provivcia view data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteAgentList(t *testing.T) {
	req := &model.GetSiteAgentListRequest{
		UserID: bcmConf.UserId,
	}
	resp, err := bcmClient.GetSiteAgentList(req)
	if err != nil {
		t.Errorf("get site agent list data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetSiteAgentListByTaskId(t *testing.T) {
	req := &model.GetTaskDetailRequest{
		UserID: bcmConf.UserId,
		TaskID: "kLuDwMzBzihGNAIYChFLSltPIzdPjTNQ",
	}
	resp, err := bcmClient.GetSiteAgentListByTaskId(req)
	if err != nil {
		t.Errorf("get site agent list data by taskId error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetCloudEventData(t *testing.T) {
	req := &model.EventDataRequest{
		AccountID: bcmConf.UserId,
		StartTime: "2023-10-01T00:00:00Z",
		EndTime:   "2023-11-01T01:00:00Z",
		PageNo:    1,
		PageSize:  10,
	}
	resp, err := bcmClient.GetCloudEventData(req)
	if err != nil {
		t.Errorf("get cloud event data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_GetPlatformEventData(t *testing.T) {
	req := &model.EventDataRequest{
		AccountID: bcmConf.UserId,
		StartTime: "2023-10-01T00:00:00Z",
		EndTime:   "2023-11-01T01:00:00Z",
		PageNo:    1,
		PageSize:  10,
	}
	resp, err := bcmClient.GetPlatformEventData(req)
	if err != nil {
		t.Errorf("get platform event data error with %v\n", err)
	}
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}

func Test_CreateEventPolicy(t *testing.T) {
	eventFilter := &model.EventFilter{
		EventLevel:    "*",
		EventTypeList: []string{"*"},
	}
	resource := &model.EventResourceFilter{
		Region:            "bj",
		Type:              "Instance",
		MonitorObjectType: "ALL",
	}
	req := &model.EventPolicy{
		AccountID:       bcmConf.UserId,
		ServiceName:     "BCE_BCC",
		Name:            "sdk_policy",
		BlockStatus:     "NORMAL",
		EventFilter:     *eventFilter,
		Resource:        *resource,
		IncidentActions: []string{"2fc6e953-331a-4404-8ce7-1c05975dbd9c"},
	}

	err := bcmClient.CreateEventPolicy(req)
	if err != nil {
		t.Errorf("create event policy error with %v\n", err)
	}
}

func Test_CreateInstanceGroup(t *testing.T) {
	req := &model.MergedGroup{
		UserId:      bcmConf.UserId,
		Region:      "bj",
		ServiceName: "BCE_BCC",
		TypeName:    "Instance",
		Name:        "sdk_test_group",
	}
	response, err := bcmClient.CreateInstanceGroup(req)
	if err != nil {
		t.Errorf("create instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_UpdateInstanceGroup(t *testing.T) {
	req := &model.InstanceGroup{
		UserID:      bcmConf.UserId,
		Region:      "bj",
		ServiceName: "BCE_BCC",
		TypeName:    "Instance",
		Name:        "sdk_test_group_update",
		ID:          7917,
	}

	response, err := bcmClient.UpdateInstanceGroup(req)
	if err != nil {
		t.Errorf("update instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_DeleteInstanceGroup(t *testing.T) {
	req := &model.InstanceGroupBase{
		UserID: bcmConf.UserId,
		ID:     "7917",
	}
	response, err := bcmClient.DeleteInstanceGroup(req)
	if err != nil {
		t.Errorf("delete instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_GetInstanceGroup(t *testing.T) {
	req := &model.InstanceGroupBase{
		UserID: bcmConf.UserId,
		ID:     "7917",
	}
	response, err := bcmClient.GetInstanceGroup(req)
	if err != nil {
		t.Errorf("get instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_ListInstanceGroup(t *testing.T) {
	req := &model.InstanceGroupQuery{
		UserID:      bcmConf.UserId,
		ServiceName: "BCE_BCC",
		PageNo:      1,
		PageSize:    10,
	}
	response, err := bcmClient.GetInstanceGroupList(req)
	if err != nil {
		t.Errorf("get instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_AddInstanceToInstanceGroup(t *testing.T) {
	monitorResource := &model.MonitorResource{
		Region:      "bj",
		UserId:      bcmConf.UserId,
		ServiceName: "BCE_BCC",
		TypeName:    "Instance",
		ResourceID:  "InstanceId:dd0109a3-a7fe-4ffb-b2ae-3c6aa0b63705",
	}

	req := &model.MergedGroup{
		ID:             "7917",
		UserId:         bcmConf.UserId,
		Region:         "bj",
		ServiceName:    "BCE_BCC",
		TypeName:       "Instance",
		Name:           "sdk_test_group_update",
		ResourceIDList: []model.MonitorResource{*monitorResource},
	}

	response, err := bcmClient.AddInstanceToInstanceGroup(req)
	if err != nil {
		t.Errorf("add instance to instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_RemoveInstanceFromInstanceGroup(t *testing.T) {
	monitorResource := &model.MonitorResource{
		Region:      "bj",
		UserId:      bcmConf.UserId,
		ServiceName: "BCE_BCC",
		TypeName:    "Instance",
		ResourceID:  "InstanceId:dd0109a3-a7fe-4ffb-b2ae-3c6aa0b63705",
	}

	req := &model.MergedGroup{
		ID:             "7917",
		UserId:         bcmConf.UserId,
		Region:         "bj",
		ServiceName:    "BCE_BCC",
		TypeName:       "Instance",
		Name:           "sdk_test_group_update",
		ResourceIDList: []model.MonitorResource{*monitorResource},
	}

	response, err := bcmClient.RemoveInstanceFromInstanceGroup(req)
	if err != nil {
		t.Errorf("remove instance from instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_GetInstanceGroupInstanceList(t *testing.T) {
	req := &model.IGInstanceQuery{
		UserID:      bcmConf.UserId,
		ID:          "7917",
		ServiceName: "BCE_BCC",
		TypeName:    "Instance",
		Region:      "bj",
		ViewType:    "DETAIL_VIEW",
		PageNo:      1,
		PageSize:    10,
	}
	response, err := bcmClient.GetInstanceGroupInstanceList(req)
	if err != nil {
		t.Errorf("get instance group instance list error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_GetAllInstanceForInstanceGroup(t *testing.T) {
	req := &model.IGInstanceQuery{
		UserID:      bcmConf.UserId,
		ServiceName: "BCE_BCC",
		TypeName:    "Instance",
		Region:      "bj",
		ViewType:    "LIST_VIEW",
		PageNo:      1,
		PageSize:    10,
		KeywordType: "name",
		Keyword:     "",
	}
	response, err := bcmClient.GetAllInstanceForInstanceGroup(req)
	if err != nil {
		t.Errorf("get all instance for instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func Test_GetFilterInstanceForInstanceGroup(t *testing.T) {
	req := &model.IGInstanceQuery{
		UserID:      bcmConf.UserId,
		ServiceName: "BCE_BCC",
		TypeName:    "Instance",
		Region:      "bj",
		ViewType:    "LIST_VIEW",
		PageNo:      1,
		PageSize:    10,
		ID:          "7917",
		UUID:        "7945f04a-0d5f-4d8b-a41c-5d5ef087884f",
	}
	response, err := bcmClient.GetFilterInstanceForInstanceGroup(req)
	if err != nil {
		t.Errorf("get filter instance for instance group error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func TestClient_GetMultiDimensionLatestMetrics(t *testing.T) {
	req := &model.MultiDimensionalLatestMetricsRequest{
		UserID: bcmConf.UserId,
		Scope:  "BCE_BLB",
		Region: "bj",
		Dimensions: []model.Dimension{
			{
				Name:  "BlbId",
				Value: "lb-****e1a0",
			},
		},
		Statistics:  []string{"average", "sum"},
		Timestamp:   "2024-03-18T06:01:00Z",
		MetricNames: []string{"ActiveConnCount"},
	}
	response, err := bcmClient.GetMultiDimensionLatestMetrics(req)
	if err != nil {
		t.Errorf("Get Multi-Dimension latest metrics error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func TestClient_GetMetricsByPartialDimensions(t *testing.T) {
	req := &model.MetricsByPartialDimensionsRequest{
		UserID: bcmConf.UserId,
		Scope:  "BCE_BLB",
		Region: "su",
		Dimensions: []model.Dimension{
			{
				Name:  "BlbPortType",
				Value: "TCP",
			},
		},
		Statistics:   []string{"sum", "average", "minimum", "maximum"},
		ResourceType: "Blb",
		MetricName:   "ActiveConnCount",
		StartTime:    "2024-03-20T02:21:17Z",
		EndTime:      "2024-03-20T03:21:17Z",
		Cycle:        30,
		PageNo:       2,
		PageSize:     2,
	}
	response, err := bcmClient.GetMetricsByPartialDimensions(req)
	if err != nil {
		t.Errorf("Get metricsByPartialDimensions error with %v\n", err)
	}
	//fmt.Printf("%+v", response)
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func TestClient_GetMetricsByPartialDimensions_atLeastParam(t *testing.T) {
	req := &model.MetricsByPartialDimensionsRequest{
		UserID:     bcmConf.UserId,
		Scope:      "BCE_BCC",
		Statistics: []string{"sum", "average", "minimum", "maximum"},
		MetricName: "CpuIdlePercent",
		StartTime:  "2024-03-20T02:21:17Z",
		EndTime:    "2024-03-20T03:21:17Z",
	}
	response, err := bcmClient.GetMetricsByPartialDimensions(req)
	if err != nil {
		t.Errorf("Get metricsByPartialDimensions error with %v\n", err)
	}
	//fmt.Printf("%+v", response)
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func TestClient_GetMetricsAllDataV2(t *testing.T) {
	req := &model.TsdbMetricAllDataQueryRequest{
		UserID: bcmConf.UserId,
		Scope:  "BCE_BCC",
		Region: "bj",
		Dimensions: [][]model.Dimension{
			{
				{
					Name:  "InstanceId",
					Value: "i-DMxr6UxX",
				},
			},
			{
				{
					Name:  "InstanceId",
					Value: "i-Y8NAmymd",
				},
			},
		},
		Statistics:  []string{"average", "sum"},
		StartTime:   "2024-03-20T07:01:00Z",
		EndTime:     "2024-03-20T07:05:00Z",
		MetricNames: []string{"CPUUsagePercent", "MemUsedPercent"},
	}
	response, err := bcmClient.GetMetricsAllDataV2(req)
	if err != nil {
		t.Errorf("Get all data metrics error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func TestClient_BatchGetMetricsAllDataV2(t *testing.T) {
	req := &model.TsdbMetricAllDataQueryRequest{
		UserID: bcmConf.UserId,
		Scope:  "BCE_MQ_KAFKA",
		Region: "bj",
		Type:   "Node",
		Dimensions: [][]model.Dimension{
			{
				{
					Name:  "ClusterId",
					Value: "efe456d667c64******0652c93812a79",
				},
				{
					Name:  "NodeId",
					Value: "i-Um1V8Haq",
				},
			},
		},
		Statistics:  []string{"average", "sum"},
		StartTime:   "2024-03-21T06:33:50Z",
		EndTime:     "2024-03-21T07:33:50Z",
		MetricNames: []string{"CpuUsedPercent", "CpuIdlePercent"},
	}
	response, err := bcmClient.BatchGetMetricsAllDataV2(req)
	if err != nil {
		t.Errorf("Get all data metrics error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func TestGetMetricDimensionTop(t *testing.T) {
	req := &model.TsdbDimensionTopQuery{
		UserID: bcmConf.UserId,
		Region: "bj",
		Scope:  "BCE_PFS",
		Dimensions: map[string]string{
			"InstanceId": "pfs-1*****7",
		},
		MetricName: "WriteIO",
		Statistics: "average",
		StartTime:  "2024-04-27T07:10:01Z",
		EndTime:    "2024-04-27T07:20:01Z",
		Labels: []string{
			"FilesetId",
		},
	}
	response, err := bcmClient.GetMetricDimensionTop(req)
	if err != nil {
		t.Errorf("Get metric dimensions top error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}

func TestGetMetricDimensionTopData(t *testing.T) {
	req := &model.TsdbDimensionTopQuery{
		UserID: bcmConf.UserId,
		Region: "bj",
		Scope:  "BCE_PFS",
		Dimensions: map[string]string{
			"InstanceId": "pfs-1******7",
		},
		MetricName: "WriteIO",
		Statistics: "average",
		StartTime:  "2024-07-10T07:10:01Z",
		EndTime:    "2024-07-10T07:20:01Z",
		Labels: []string{
			"FilesetId",
		},
	}
	response, err := bcmClient.GetMetricDimensionTopData(req)
	if err != nil {
		t.Errorf("Get metric dimensions top and data error with %v\n", err)
	}
	content, _ := json.Marshal(response)
	fmt.Println(string(content))
}
