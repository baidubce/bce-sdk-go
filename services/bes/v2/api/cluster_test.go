package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	clusterTestClusterId = "644734225693675520"
	clusterTestOrderId   = "order-8f0d1b9c"
)

// clusterWriteBadRequest writes a 400 service error. A 4xx status is used on purpose: the test
// client mirrors the production DEFAULT_RETRY_POLICY, so a 5xx would trigger backoff retries.
func clusterWriteBadRequest(t *testing.T, w nethttp.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(nethttp.StatusBadRequest)
	if _, err := w.Write([]byte(`{"code":"BadRequest","message":"invalid cluster request"}`)); err != nil {
		t.Fatalf("write error response failed: %v", err)
	}
}

// clusterUnreachableHandler fails the test when a request reaches the server, used by cases that
// must fail during local validation.
func clusterUnreachableHandler(t *testing.T) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Errorf("unexpected request reached the server: %s", r.URL.Path)
		clusterWriteBadRequest(t, w)
	}
}

// clusterAssertRequest checks the parts of the request every cluster API builds the same way.
func clusterAssertRequest(t *testing.T, r *nethttp.Request, path string) {
	t.Helper()
	if r.URL.Path != path {
		t.Fatalf("unexpected path: %s", r.URL.Path)
	}
	if r.Method != nethttp.MethodPost {
		t.Fatalf("unexpected method: %s", r.Method)
	}
	if got := r.Header.Get(HEADER_REGION); got != testRegion {
		t.Fatalf("unexpected region header: %s", got)
	}
}

// The clusterFullXxxRequest builders return fully populated requests, including every optional
// field, so the required-field cases below only have to clear one field each.
func clusterFullCreateRequest() *CreateClusterRequest {
	return &CreateClusterRequest{
		Name:            "postpaynew1",
		Password:        testPassword,
		SecurityGroupId: "g-s77w7h0ierks",
		SubnetUuid:      "sbn-38zqasty2i8n",
		AvailableZone:   "cn-bd-b",
		VpcId:           "vpc-5dd5bib4h0vc",
		IsOldPackage:    true,
		Version:         "7.4.2",
		Modules: []ModuleInfo{{
			Type:        "es_node",
			InstanceNum: 2,
			SlotType:    "bes.g3.c2m8",
			DiskSlotInfo: &DiskSlotInfo{
				Size:       50,
				Type:       "premium_ssd",
				CdsExtraIo: 1,
			},
		}},
		Billing: &Billing{
			PaymentType:     "prepay",
			Time:            1,
			EnableAutoRenew: true,
			AutoRenewInfo:   &AutoRenewInfo{RenewTimeUnit: "month", RenewTime: 1},
			Coupon:          "coupon-1",
		},
		Tags:       []Tag{{TagKey: "env", TagValue: "test"}},
		EnableSSL:  true,
		ResGroupId: "RESG-oxSCFbYA",
	}
}

func clusterFullListRequest() *ListClustersRequest {
	return &ListClustersRequest{PageNo: 1, PageSize: 10}
}

func clusterFullDeleteRequest() *DeleteClusterRequest {
	return &DeleteClusterRequest{ClusterId: clusterTestClusterId, RefundReason: "no longer needed"}
}

func clusterFullRestartRequest() *RestartClusterRequest {
	return &RestartClusterRequest{ClusterId: clusterTestClusterId, Mode: "full_restart"}
}

func clusterFullResizeRequest() *ResizeClusterRequest {
	return &ResizeClusterRequest{
		ClusterId: clusterTestClusterId,
		Modules: []ResizeModuleInfo{{
			Type:              "es_node",
			Version:           "7.4.2",
			DesireInstanceNum: 3,
			SlotType:          "bes.g3.c2m8",
			DiskSlotInfo: &ResizeDiskSlotInfo{
				Type:       "premium_ssd",
				Size:       100,
				CdsExtraIo: 1,
			},
		}},
		PaymentType: "postpay",
		Coupon:      "coupon-1",
		ResizeMode:  "SCROLL",
		IsShrink:    true,
	}
}

func clusterFullAddModuleRequest() *AddClusterModuleRequest {
	return &AddClusterModuleRequest{
		ClusterId: clusterTestClusterId,
		Modules: []AddModuleInfo{{
			Type:              "es_coordinate_node",
			Version:           "7.4.2",
			DesireInstanceNum: 1,
			SlotType:          "bes.g3.c2m8",
			DiskSlotInfo:      &AddModuleDiskSlotInfo{Type: "ssd", Size: 20},
		}},
		ResizeMode:  "SCROLL",
		PaymentType: "postpay",
	}
}

func clusterFullResetPasswordRequest() *ResetClusterPasswordRequest {
	return &ResetClusterPasswordRequest{
		ClusterId:       clusterTestClusterId,
		NewPassword:     testPassword,
		ConfirmPassword: testPassword,
	}
}

func clusterFullToggleHTTPSRequest() *ToggleClusterHTTPSRequest {
	return &ToggleClusterHTTPSRequest{ClusterId: clusterTestClusterId, EnableSSL: true}
}

func clusterFullBindEIPRequest() *BindClusterEIPRequest {
	return &BindClusterEIPRequest{
		InstanceId:   clusterTestClusterId,
		InstanceType: "es_node",
		Eip:          "10.11.12.13",
	}
}

func clusterFullUnbindEIPRequest() *UnbindClusterEIPRequest {
	return &UnbindClusterEIPRequest{
		DeployId:           clusterTestClusterId,
		ModuleTemplateName: "es_node",
	}
}

func clusterFullToggleMonitorRequest() *ToggleClusterMonitorRequest {
	return &ToggleClusterMonitorRequest{ClusterId: clusterTestClusterId, EnableMonitor: boolPtr(true)}
}

func clusterFullDataSizeTendencyRequest() *GetClusterDataSizeTendencyRequest {
	return &GetClusterDataSizeTendencyRequest{
		ClusterId:   clusterTestClusterId,
		IndexPrefix: "logstash-",
		DatePattern: "yyyy-MM-dd",
		Times:       7,
		TimeUnit:    "DAYS",
	}
}

func clusterFullAssessRequest() *AssessClusterSourceRequest {
	return &AssessClusterSourceRequest{
		UseType:            "COMMON",
		OriginDataSize:     100,
		OriginDataSizeUnit: "GIB",
		DataAdd:            10,
		DataAddUnit:        "GIB_DAY",
		StorageDays:        30,
		WriteThroughput:    100,
		ReadThroughput:     50,
		Replica:            "1",
		NeedBos:            true,
		BosStorageDays:     60,
		NeedVector:         true,
		VectorDims:         768,
		VectorType:         "float",
	}
}

func TestCreateCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/create")
		body := readJSONBody(t, r)
		if body["name"] != "postpaynew1" || body["password"] != testPassword {
			t.Fatalf("unexpected credentials fields: %+v", body)
		}
		if body["securityGroupId"] != "g-s77w7h0ierks" || body["subnetUuid"] != "sbn-38zqasty2i8n" {
			t.Fatalf("unexpected network fields: %+v", body)
		}
		if body["availableZone"] != "cn-bd-b" || body["vpcId"] != "vpc-5dd5bib4h0vc" {
			t.Fatalf("unexpected zone fields: %+v", body)
		}
		if body["version"] != "7.4.2" || body["isOldPackage"] != true {
			t.Fatalf("unexpected version fields: %+v", body)
		}
		if body["enableSSL"] != true || body["resGroupId"] != "RESG-oxSCFbYA" {
			t.Fatalf("unexpected optional fields: %+v", body)
		}
		clusterAssertCreateModules(t, body)
		clusterAssertCreateBilling(t, body)
		clusterAssertCreateTags(t, body)
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"clusterId": clusterTestClusterId},
		})
	}))
	defer server.Close()

	result, err := CreateCluster(cli, testRegion, clusterFullCreateRequest())
	if err != nil {
		t.Fatalf("create cluster failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	created, ok := result.Result.(map[string]interface{})
	if !ok || created["clusterId"] != clusterTestClusterId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

// clusterAssertCreateModules / clusterAssertCreateBilling / clusterAssertCreateTags keep
// TestCreateCluster readable while still checking every nested field the builder populates.
func clusterAssertCreateModules(t *testing.T, body map[string]interface{}) {
	t.Helper()
	modules, ok := body["modules"].([]interface{})
	if !ok || len(modules) != 1 {
		t.Fatalf("unexpected modules: %v", body["modules"])
	}
	module, ok := modules[0].(map[string]interface{})
	if !ok || module["type"] != "es_node" || module["slotType"] != "bes.g3.c2m8" {
		t.Fatalf("unexpected module: %v", modules[0])
	}
	if module["instanceNum"] != float64(2) {
		t.Fatalf("unexpected instanceNum: %v", module["instanceNum"])
	}
	disk, ok := module["diskSlotInfo"].(map[string]interface{})
	if !ok || disk["size"] != float64(50) || disk["type"] != "premium_ssd" {
		t.Fatalf("unexpected diskSlotInfo: %v", module["diskSlotInfo"])
	}
	if disk["cdsExtraIo"] != float64(1) {
		t.Fatalf("unexpected cdsExtraIo: %v", disk["cdsExtraIo"])
	}
}

func clusterAssertCreateBilling(t *testing.T, body map[string]interface{}) {
	t.Helper()
	billing, ok := body["billing"].(map[string]interface{})
	if !ok || billing["paymentType"] != "prepay" || billing["time"] != float64(1) {
		t.Fatalf("unexpected billing: %v", body["billing"])
	}
	if billing["enableAutoRenew"] != true || billing["coupon"] != "coupon-1" {
		t.Fatalf("unexpected billing renew fields: %v", billing)
	}
	autoRenew, ok := billing["autoRenewInfo"].(map[string]interface{})
	if !ok || autoRenew["renewTimeUnit"] != "month" || autoRenew["renewTime"] != float64(1) {
		t.Fatalf("unexpected autoRenewInfo: %v", billing["autoRenewInfo"])
	}
}

func clusterAssertCreateTags(t *testing.T, body map[string]interface{}) {
	t.Helper()
	tags, ok := body["tags"].([]interface{})
	if !ok || len(tags) != 1 {
		t.Fatalf("unexpected tags: %v", body["tags"])
	}
	tag, ok := tags[0].(map[string]interface{})
	if !ok || tag["tagKey"] != "env" || tag["tagValue"] != "test" {
		t.Fatalf("unexpected tag: %v", tags[0])
	}
}

func TestListClusters(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/v2/list")
		body := readJSONBody(t, r)
		if body["pageNo"] != float64(1) || body["pageSize"] != float64(10) {
			t.Fatalf("unexpected page fields: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"page": map[string]interface{}{
				"pageNo":     1,
				"pageSize":   10,
				"totalCount": 1,
				"result": []map[string]interface{}{{
					"clusterId":     clusterTestClusterId,
					"clusterName":   "postpaynew1",
					"createTime":    "2026-07-01T03:03:54Z",
					"actualStatus":  "RUNNING",
					"runningTime":   "10d",
					"region":        "bd",
					"version":       "7.4.2",
					"clusterHealth": map[string]interface{}{"status": "green"},
					"billing":       map[string]interface{}{"paymentType": "postpay"},
					"tags":          []map[string]interface{}{{"tagKey": "env", "tagValue": "test"}},
					"resGroupList":  []map[string]interface{}{{"groupId": "g-1", "groupName": "default"}},
					"innerAccessIp": "192.168.0.1",
					"accessEip":     "10.11.12.13",
					"vpcId":         "vpc-5dd5bib4h0vc",
					"subnet":        []map[string]interface{}{{"availableZone": "cn-bd-b", "subnetName": "sn"}},
					"network":       []map[string]interface{}{{"subnetId": "sbn-38zqasty2i8n"}},
				}},
			},
		})
	}))
	defer server.Close()

	result, err := ListClusters(cli, testRegion, clusterFullListRequest())
	if err != nil {
		t.Fatalf("list clusters failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Page == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Page.PageNo != 1 || result.Page.PageSize != 10 || result.Page.TotalCount != 1 {
		t.Fatalf("unexpected page: %+v", result.Page)
	}
	if len(result.Page.Result) != 1 {
		t.Fatalf("unexpected page result: %+v", result.Page.Result)
	}
	item := result.Page.Result[0]
	if item.ClusterId != clusterTestClusterId || item.ClusterName != "postpaynew1" {
		t.Fatalf("unexpected cluster item: %+v", item)
	}
	if item.ClusterHealth == nil || item.ClusterHealth.Status != "green" {
		t.Fatalf("unexpected cluster health: %+v", item.ClusterHealth)
	}
	if item.Billing == nil || item.Billing.PaymentType != "postpay" {
		t.Fatalf("unexpected billing: %+v", item.Billing)
	}
	if len(item.Subnet) != 1 || item.Subnet[0].SubnetName != "sn" {
		t.Fatalf("unexpected subnet: %+v", item.Subnet)
	}
	if len(item.Network) != 1 || item.Network[0].SubnetId != "sbn-38zqasty2i8n" {
		t.Fatalf("unexpected network: %+v", item.Network)
	}
	if len(item.ResGroupList) != 1 || item.ResGroupList[0].GroupName != "default" {
		t.Fatalf("unexpected resGroupList: %+v", item.ResGroupList)
	}
}

// clusterDetailPayload is the detail result used by TestGetClusterDetail. It fills every optional
// block of ClusterDetail so the response decoding is exercised end to end.
func clusterDetailPayload() map[string]interface{} {
	return map[string]interface{}{
		"expireTime":      "2026-11-27T03:03:54Z",
		"clusterId":       clusterTestClusterId,
		"clusterName":     "postpaynew1",
		"adminUsername":   "superuser",
		"actualStatus":    "RUNNING",
		"desireStatus":    "RUNNING",
		"esUrl":           "http://192.168.0.1:9200",
		"kibanaUrl":       "http://192.168.0.1:5601",
		"exporterUrl":     "http://192.168.0.1:9114",
		"enableEsSSL":     true,
		"enableKibanaSSL": true,
		"esEip":           "10.11.12.13",
		"kibanaEip":       "10.11.12.14",
		"modules": []map[string]interface{}{{
			"type":              "es_node",
			"version":           "7.4.2",
			"slotType":          "bes.g3.c2m8",
			"slotDescription":   "2C8G",
			"actualInstanceNum": 2,
		}},
		"instances": []map[string]interface{}{{
			"instanceId":    "i-1",
			"status":        "RUNNING",
			"moduleType":    "es_node",
			"moduleVersion": "7.4.2",
			"hostIp":        "192.168.0.1",
		}},
		"region":            "bd",
		"inspectAuthorized": "true",
		"vpc":               "default",
		"vpcId":             "vpc-5dd5bib4h0vc",
		"subnet":            "sn",
		"availableZone":     "cn-bd-b",
		"securityGroup":     "g-s77w7h0ierks",
		"clusterHealth": map[string]interface{}{
			"status":              "green",
			"totalDiskSize":       "100GB",
			"usedDiskSize":        "10GB",
			"nodeDiskMaxSizeInGB": map[string]interface{}{"es_node": 100, "es_cold_tier_node": 200},
			"indexDiskSize":       "5GB",
		},
		"billing":      map[string]interface{}{"paymentType": "prepay"},
		"network":      []map[string]interface{}{{"subnetId": "sbn-1", "subnet": "sn", "availableZone": "cn-bd-b"}},
		"tags":         []map[string]interface{}{{"tagKey": "env", "tagValue": "test"}},
		"log":          map[string]interface{}{"mainLog": true, "gcLog": true, "indexSlowLog": true, "searchSlowLog": true, "auditLog": true},
		"resGroupList": []map[string]interface{}{{"groupId": "g-1", "groupName": "default"}},
		"bdVersion":    "2.0",
	}
}

func TestGetClusterDetail(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/v2/detail")
		if got := readJSONBody(t, r)["clusterId"]; got != clusterTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  clusterDetailPayload(),
		})
	}))
	defer server.Close()

	result, err := GetClusterDetail(cli, testRegion, &GetClusterDetailRequest{
		ClusterId: clusterTestClusterId,
	})
	if err != nil {
		t.Fatalf("get cluster detail failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
	detail := result.Result
	if detail.ClusterId != clusterTestClusterId || detail.ClusterName != "postpaynew1" {
		t.Fatalf("unexpected detail: %+v", detail)
	}
	if detail.AdminUsername != "superuser" || !detail.EnableEsSSL || !detail.EnableKibanaSSL {
		t.Fatalf("unexpected detail flags: %+v", detail)
	}
	if !detail.InspectAuthorized.Bool() {
		t.Fatalf("unexpected inspectAuthorized: %v", detail.InspectAuthorized)
	}
	if len(detail.Modules) != 1 || detail.Modules[0].ActualInstanceNum != 2 {
		t.Fatalf("unexpected modules: %+v", detail.Modules)
	}
	if len(detail.Instances) != 1 || detail.Instances[0].HostIp != "192.168.0.1" {
		t.Fatalf("unexpected instances: %+v", detail.Instances)
	}
	if detail.ClusterHealth == nil || detail.ClusterHealth.Status != "green" {
		t.Fatalf("unexpected cluster health: %+v", detail.ClusterHealth)
	}
	if detail.ClusterHealth.NodeDiskMaxSizeInGB == nil ||
		detail.ClusterHealth.NodeDiskMaxSizeInGB.EsNode.String() != "100" ||
		detail.ClusterHealth.NodeDiskMaxSizeInGB.EsColdTierNode.String() != "200" {
		t.Fatalf("unexpected node disk sizes: %+v", detail.ClusterHealth.NodeDiskMaxSizeInGB)
	}
	if detail.Log == nil || !detail.Log.MainLog || !detail.Log.AuditLog {
		t.Fatalf("unexpected log setting: %+v", detail.Log)
	}
	if len(detail.Network) != 1 || detail.Network[0].SubnetId != "sbn-1" {
		t.Fatalf("unexpected network: %+v", detail.Network)
	}
	if detail.Billing == nil || detail.Billing.PaymentType != "prepay" || detail.BdVersion != "2.0" {
		t.Fatalf("unexpected billing/version: %+v", detail)
	}
}

func TestDeleteCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/v2/delete")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["refundReason"] != "no longer needed" {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"orderId": clusterTestOrderId},
		})
	}))
	defer server.Close()

	result, err := DeleteCluster(cli, testRegion, clusterFullDeleteRequest())
	if err != nil {
		t.Fatalf("delete cluster failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result == nil || result.Result.OrderId != clusterTestOrderId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestStartCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/start")
		if got := r.Header.Get("Content-Type"); got != bce.DEFAULT_CONTENT_TYPE {
			t.Fatalf("unexpected content type: %s", got)
		}
		if got := readJSONBody(t, r)["clusterId"]; got != clusterTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := StartCluster(cli, testRegion, &StartClusterRequest{ClusterId: clusterTestClusterId})
	if err != nil {
		t.Fatalf("start cluster failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestStopCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/stop")
		if got := readJSONBody(t, r)["clusterId"]; got != clusterTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := StopCluster(cli, testRegion, &StopClusterRequest{ClusterId: clusterTestClusterId})
	if err != nil {
		t.Fatalf("stop cluster failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestRestartCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/restart")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["mode"] != "full_restart" {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := RestartCluster(cli, testRegion, clusterFullRestartRequest())
	if err != nil {
		t.Fatalf("restart cluster failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestResizeCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/resize")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["paymentType"] != "postpay" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["resizeMode"] != "SCROLL" || body["coupon"] != "coupon-1" || body["isShrink"] != true {
			t.Fatalf("unexpected resize fields: %+v", body)
		}
		modules, ok := body["modules"].([]interface{})
		if !ok || len(modules) != 1 {
			t.Fatalf("unexpected modules: %v", body["modules"])
		}
		module, ok := modules[0].(map[string]interface{})
		if !ok || module["type"] != "es_node" || module["version"] != "7.4.2" {
			t.Fatalf("unexpected module: %v", modules[0])
		}
		if module["desireInstanceNum"] != float64(3) || module["slotType"] != "bes.g3.c2m8" {
			t.Fatalf("unexpected module size: %v", module)
		}
		disk, ok := module["diskSlotInfo"].(map[string]interface{})
		if !ok || disk["type"] != "premium_ssd" || disk["size"] != float64(100) {
			t.Fatalf("unexpected diskSlotInfo: %v", module["diskSlotInfo"])
		}
		if disk["cdsExtraIo"] != float64(1) {
			t.Fatalf("unexpected cdsExtraIo: %v", disk["cdsExtraIo"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"orderId": clusterTestOrderId},
		})
	}))
	defer server.Close()

	result, err := ResizeCluster(cli, testRegion, clusterFullResizeRequest())
	if err != nil {
		t.Fatalf("resize cluster failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result == nil || result.Result.OrderId != clusterTestOrderId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestAddClusterModule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/addModule")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["resizeMode"] != "SCROLL" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["paymentType"] != "postpay" {
			t.Fatalf("unexpected paymentType: %v", body["paymentType"])
		}
		modules, ok := body["modules"].([]interface{})
		if !ok || len(modules) != 1 {
			t.Fatalf("unexpected modules: %v", body["modules"])
		}
		module, ok := modules[0].(map[string]interface{})
		if !ok || module["type"] != "es_coordinate_node" || module["version"] != "7.4.2" {
			t.Fatalf("unexpected module: %v", modules[0])
		}
		if module["desireInstanceNum"] != float64(1) || module["slotType"] != "bes.g3.c2m8" {
			t.Fatalf("unexpected module size: %v", module)
		}
		disk, ok := module["diskSlotInfo"].(map[string]interface{})
		if !ok || disk["type"] != "ssd" || disk["size"] != float64(20) {
			t.Fatalf("unexpected diskSlotInfo: %v", module["diskSlotInfo"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"orderId": clusterTestOrderId},
		})
	}))
	defer server.Close()

	result, err := AddClusterModule(cli, testRegion, clusterFullAddModuleRequest())
	if err != nil {
		t.Fatalf("add cluster module failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result == nil || result.Result.OrderId != clusterTestOrderId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

// TestResetClusterPassword also pins the string-shaped success flag the real API answers with.
func TestResetClusterPassword(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/password/reset")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["newPassword"] != testPassword {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["confirmPassword"] != testPassword {
			t.Fatalf("unexpected confirmPassword: %v", body["confirmPassword"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  "result",
		})
	}))
	defer server.Close()

	result, err := ResetClusterPassword(cli, testRegion, clusterFullResetPasswordRequest())
	if err != nil {
		t.Fatalf("reset cluster password failed: %v", err)
	}
	if !result.Success.Bool() || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestToggleClusterHTTPS(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/v2/esProtocolTransform")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["enableSSL"] != true {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := ToggleClusterHTTPS(cli, testRegion, clusterFullToggleHTTPSRequest())
	if err != nil {
		t.Fatalf("toggle cluster https failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestBindClusterEIP(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/eip/bind")
		body := readJSONBody(t, r)
		if body["instanceId"] != clusterTestClusterId || body["instanceType"] != "es_node" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["eip"] != "10.11.12.13" {
			t.Fatalf("unexpected eip: %v", body["eip"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := BindClusterEIP(cli, testRegion, clusterFullBindEIPRequest())
	if err != nil {
		t.Fatalf("bind cluster eip failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestUnbindClusterEIP(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/eip/unbind")
		body := readJSONBody(t, r)
		if body["deployId"] != clusterTestClusterId || body["moduleTemplateName"] != "es_node" {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := UnbindClusterEIP(cli, testRegion, clusterFullUnbindEIPRequest())
	if err != nil {
		t.Fatalf("unbind cluster eip failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestToggleClusterMonitor(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/monitor_state")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["enableMonitor"] != true {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  "result",
		})
	}))
	defer server.Close()

	result, err := ToggleClusterMonitor(cli, testRegion, clusterFullToggleMonitorRequest())
	if err != nil {
		t.Fatalf("toggle cluster monitor failed: %v", err)
	}
	if !result.Success.Bool() || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

// TestToggleClusterMonitorDisable pins the false branch of the *bool flag: an explicit false must
// still be sent, and must not be mistaken for a missing field.
func TestToggleClusterMonitorDisable(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if got := readJSONBody(t, r)["enableMonitor"]; got != false {
			t.Fatalf("unexpected enableMonitor: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": false, "status": 200})
	}))
	defer server.Close()

	result, err := ToggleClusterMonitor(cli, testRegion, &ToggleClusterMonitorRequest{
		ClusterId:     clusterTestClusterId,
		EnableMonitor: boolPtr(false),
	})
	if err != nil {
		t.Fatalf("toggle cluster monitor failed: %v", err)
	}
	if result.Success.Bool() {
		t.Fatalf("unexpected success: %v", result.Success)
	}
}

// TestGetClusterTasks uses the shapes captured from the gray environment: clusterId and the two
// timestamps come back as bare numbers, and subTasks carries a realProgress the API doc omits.
func TestGetClusterTasks(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/cluster_tasks")
		if got := readJSONBody(t, r)["clusterId"]; got != clusterTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"clusterId": 1210468255077109760,
				"appTasks": []map[string]interface{}{{
					"type":            "CREATE",
					"status":          "FINISH",
					"progress":        100,
					"detail":          map[string]interface{}{"operator": "console"},
					"start_timestamp": 1787206164689,
					"end_timestamp":   1787206403859,
					"subTasks": []map[string]interface{}{{
						"type":            "INIT",
						"status":          "FINISH",
						"progress":        100,
						"realProgress":    -1,
						"detail":          map[string]interface{}{"step": "1"},
						"start_timestamp": 1787206164689,
						"end_timestamp":   1787206214506,
						"taskList": []map[string]interface{}{{
							"name":          "index-check",
							"type":          "CHECK",
							"level":         "WARN",
							"status":        "FINISH",
							"finishTime":    "2026-07-01T03:03:54Z",
							"message":       "ok",
							"failedIndices": []string{"idx-1"},
						}},
					}},
				}},
			},
		})
	}))
	defer server.Close()

	result, err := GetClusterTasks(cli, testRegion, &GetClusterTasksRequest{
		ClusterId: clusterTestClusterId,
	})
	if err != nil {
		t.Fatalf("get cluster tasks failed: %v", err)
	}
	clusterAssertTasksResult(t, result)
}

func clusterAssertTasksResult(t *testing.T, result *GetClusterTasksResponse) {
	t.Helper()
	if !result.Success.Bool() || result.Status != 200 || result.Result == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result.ClusterId.String() != "1210468255077109760" {
		t.Fatalf("unexpected clusterId: %s", result.Result.ClusterId)
	}
	if len(result.Result.AppTasks) != 1 {
		t.Fatalf("unexpected appTasks: %+v", result.Result.AppTasks)
	}
	appTask := result.Result.AppTasks[0]
	if appTask.Type != "CREATE" || appTask.Status != "FINISH" || appTask.Progress != 100 {
		t.Fatalf("unexpected app task: %+v", appTask)
	}
	if appTask.Detail["operator"] != "console" {
		t.Fatalf("unexpected app task detail: %+v", appTask.Detail)
	}
	if appTask.StartTimestamp != 1787206164689 || appTask.EndTimestamp != 1787206403859 {
		t.Fatalf("unexpected app task timestamps: %+v", appTask)
	}
	if len(appTask.SubTasks) != 1 {
		t.Fatalf("unexpected sub tasks: %+v", appTask.SubTasks)
	}
	subTask := appTask.SubTasks[0]
	if subTask.Type != "INIT" || subTask.Progress != 100 || subTask.RealProgress != -1 {
		t.Fatalf("unexpected sub task: %+v", subTask)
	}
	if subTask.Detail["step"] != "1" || subTask.EndTimestamp != 1787206214506 {
		t.Fatalf("unexpected sub task detail: %+v", subTask)
	}
	if len(subTask.TaskList) != 1 {
		t.Fatalf("unexpected task list: %+v", subTask.TaskList)
	}
	check := subTask.TaskList[0]
	if check.Name != "index-check" || check.Level != "WARN" || check.Message != "ok" {
		t.Fatalf("unexpected check item: %+v", check)
	}
	if len(check.FailedIndices) != 1 || check.FailedIndices[0] != "idx-1" {
		t.Fatalf("unexpected failed indices: %+v", check.FailedIndices)
	}
}

func TestGetClusterDataSizeTendency(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/data_size_tendency")
		body := readJSONBody(t, r)
		if body["clusterId"] != clusterTestClusterId || body["indexPrefix"] != "logstash-" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["datePattern"] != "yyyy-MM-dd" || body["timeUnit"] != "DAYS" {
			t.Fatalf("unexpected date fields: %+v", body)
		}
		if body["times"] != float64(7) {
			t.Fatalf("unexpected times: %v", body["times"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"clusterId": 1210468255077109760,
				"data": []map[string]interface{}{{
					"byteSize":  "10737418240",
					"size":      "10GB",
					"time":      "2026-07-01",
					"timeMills": "1782000000000",
					"indices": []map[string]interface{}{{
						"byteSize": "1073741824",
						"index":    "logstash-2026.07.01",
						"size":     "1GB",
					}},
				}},
			},
		})
	}))
	defer server.Close()

	result, err := GetClusterDataSizeTendency(cli, testRegion, clusterFullDataSizeTendencyRequest())
	if err != nil {
		t.Fatalf("get cluster data size tendency failed: %v", err)
	}
	if !result.Success.Bool() || result.Status != 200 || result.Result == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result.ClusterId.String() != "1210468255077109760" || len(result.Result.Data) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	item := result.Result.Data[0]
	if item.Size != "10GB" || item.Time != "2026-07-01" || item.TimeMills != "1782000000000" {
		t.Fatalf("unexpected data item: %+v", item)
	}
	if len(item.Indices) != 1 || item.Indices[0].Index != "logstash-2026.07.01" {
		t.Fatalf("unexpected indices: %+v", item.Indices)
	}
}

// TestListAvailableCoupons is the only cluster API without a request parameter: it posts an empty
// JSON object, which is asserted here so the struct{} payload stays visible.
func TestListAvailableCoupons(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/coupon/avail")
		if body := readJSONBody(t, r); len(body) != 0 {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"coupons": []map[string]interface{}{{
					"id":                         1,
					"name":                       "coupon-1",
					"couponType":                 "CASH",
					"productType":                "BES",
					"productTypes":               []string{"BES"},
					"productRuleDescription":     "BES only",
					"totalAmount":                100,
					"balance":                    "100",
					"amountOrDiscount":           "100",
					"usedAmount":                 "0",
					"beginTime":                  1782000000000,
					"endTime":                    1783000000000,
					"couponStatus":               "AVAILABLE",
					"status":                     "NORMAL",
					"region":                     "bd",
					"conditionArgsMap":           map[string]interface{}{"minAmount": "10"},
					"effectArgsMap":              map[string]interface{}{"amount": "100"},
					"conditionEffectDescription": "spend 10 get 100",
				}},
			},
		})
	}))
	defer server.Close()

	result, err := ListAvailableCoupons(cli, testRegion)
	if err != nil {
		t.Fatalf("list available coupons failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
	if len(result.Result.Coupons) != 1 {
		t.Fatalf("unexpected coupons: %+v", result.Result.Coupons)
	}
	coupon := result.Result.Coupons[0]
	if coupon.Id != 1 || coupon.Name != "coupon-1" || coupon.CouponType != "CASH" {
		t.Fatalf("unexpected coupon: %+v", coupon)
	}
	if coupon.Balance != "100" || coupon.CouponStatus != "AVAILABLE" || coupon.Region != "bd" {
		t.Fatalf("unexpected coupon state: %+v", coupon)
	}
	if coupon.ConditionArgsMap["minAmount"] != "10" || coupon.EffectArgsMap["amount"] != "100" {
		t.Fatalf("unexpected coupon maps: %+v", coupon)
	}
}

func TestAssessClusterSource(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterAssertRequest(t, r, "/api/bes/cluster/source_assess")
		clusterAssertAssessBody(t, readJSONBody(t, r))
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result": map[string]interface{}{
				"zoneList": []map[string]interface{}{{"name": "cn-bj-a", "sellOut": "false"}},
				"packageStatusList": []map[string]interface{}{{
					"moduleType":     "es_node",
					"packageVersion": "bes.g3.c2m8",
					"diskType":       "premium_ssd",
					"diskSize":       100,
					"moduleNum":      3,
				}},
				"otherConfig": map[string]interface{}{"bosDataSize": 200},
			},
		})
	}))
	defer server.Close()

	result, err := AssessClusterSource(cli, testRegion, clusterFullAssessRequest())
	if err != nil {
		t.Fatalf("assess cluster source failed: %v", err)
	}
	if !result.Success.Bool() || result.Status != 200 || result.Result == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
	if len(result.Result.ZoneList) != 1 || result.Result.ZoneList[0].Name != "cn-bj-a" {
		t.Fatalf("unexpected zoneList: %+v", result.Result.ZoneList)
	}
	if result.Result.ZoneList[0].SellOut.Bool() {
		t.Fatalf("unexpected sellOut: %v", result.Result.ZoneList[0].SellOut)
	}
	if len(result.Result.PackageStatusList) != 1 {
		t.Fatalf("unexpected packageStatusList: %+v", result.Result.PackageStatusList)
	}
	pkg := result.Result.PackageStatusList[0]
	if pkg.ModuleType != "es_node" || pkg.DiskSize.String() != "100" || pkg.ModuleNum.String() != "3" {
		t.Fatalf("unexpected package: %+v", pkg)
	}
	if result.Result.OtherConfig == nil || result.Result.OtherConfig.BosDataSize.String() != "200" {
		t.Fatalf("unexpected otherConfig: %+v", result.Result.OtherConfig)
	}
}

func clusterAssertAssessBody(t *testing.T, body map[string]interface{}) {
	t.Helper()
	if body["useType"] != "COMMON" || body["originDataSizeUnit"] != "GIB" {
		t.Fatalf("unexpected use fields: %+v", body)
	}
	if body["originDataSize"] != float64(100) || body["dataAdd"] != float64(10) {
		t.Fatalf("unexpected data fields: %+v", body)
	}
	if body["dataAddUnit"] != "GIB_DAY" || body["storageDays"] != float64(30) {
		t.Fatalf("unexpected storage fields: %+v", body)
	}
	if body["writeThroughput"] != float64(100) || body["readThroughput"] != float64(50) {
		t.Fatalf("unexpected throughput fields: %+v", body)
	}
	if body["replica"] != "1" || body["needBos"] != true || body["bosStorageDays"] != float64(60) {
		t.Fatalf("unexpected bos fields: %+v", body)
	}
	if body["needVector"] != true || body["vectorDims"] != float64(768) || body["vectorType"] != "float" {
		t.Fatalf("unexpected vector fields: %+v", body)
	}
}

// clusterAPICall is one exported cluster API invoked with a fully populated request. The
// nil-client and the server-error tables walk the same list, so a new API cannot be forgotten.
type clusterAPICall struct {
	name string
	call func(bce.Client) error
}

func clusterAPICalls() []clusterAPICall {
	return []clusterAPICall{
		{"CreateCluster", func(c bce.Client) error {
			_, err := CreateCluster(c, testRegion, clusterFullCreateRequest())
			return err
		}},
		{"ListClusters", func(c bce.Client) error {
			_, err := ListClusters(c, testRegion, clusterFullListRequest())
			return err
		}},
		{"GetClusterDetail", func(c bce.Client) error {
			_, err := GetClusterDetail(c, testRegion, &GetClusterDetailRequest{
				ClusterId: clusterTestClusterId,
			})
			return err
		}},
		{"DeleteCluster", func(c bce.Client) error {
			_, err := DeleteCluster(c, testRegion, clusterFullDeleteRequest())
			return err
		}},
		{"StartCluster", func(c bce.Client) error {
			_, err := StartCluster(c, testRegion, &StartClusterRequest{ClusterId: clusterTestClusterId})
			return err
		}},
		{"StopCluster", func(c bce.Client) error {
			_, err := StopCluster(c, testRegion, &StopClusterRequest{ClusterId: clusterTestClusterId})
			return err
		}},
		{"RestartCluster", func(c bce.Client) error {
			_, err := RestartCluster(c, testRegion, clusterFullRestartRequest())
			return err
		}},
		{"ResizeCluster", func(c bce.Client) error {
			_, err := ResizeCluster(c, testRegion, clusterFullResizeRequest())
			return err
		}},
		{"AddClusterModule", func(c bce.Client) error {
			_, err := AddClusterModule(c, testRegion, clusterFullAddModuleRequest())
			return err
		}},
		{"ResetClusterPassword", func(c bce.Client) error {
			_, err := ResetClusterPassword(c, testRegion, clusterFullResetPasswordRequest())
			return err
		}},
		{"ToggleClusterHTTPS", func(c bce.Client) error {
			_, err := ToggleClusterHTTPS(c, testRegion, clusterFullToggleHTTPSRequest())
			return err
		}},
		{"BindClusterEIP", func(c bce.Client) error {
			_, err := BindClusterEIP(c, testRegion, clusterFullBindEIPRequest())
			return err
		}},
		{"UnbindClusterEIP", func(c bce.Client) error {
			_, err := UnbindClusterEIP(c, testRegion, clusterFullUnbindEIPRequest())
			return err
		}},
		{"ToggleClusterMonitor", func(c bce.Client) error {
			_, err := ToggleClusterMonitor(c, testRegion, clusterFullToggleMonitorRequest())
			return err
		}},
		{"GetClusterTasks", func(c bce.Client) error {
			_, err := GetClusterTasks(c, testRegion, &GetClusterTasksRequest{
				ClusterId: clusterTestClusterId,
			})
			return err
		}},
		{"GetClusterDataSizeTendency", func(c bce.Client) error {
			_, err := GetClusterDataSizeTendency(c, testRegion, clusterFullDataSizeTendencyRequest())
			return err
		}},
		{"ListAvailableCoupons", func(c bce.Client) error {
			_, err := ListAvailableCoupons(c, testRegion)
			return err
		}},
		{"AssessClusterSource", func(c bce.Client) error {
			_, err := AssessClusterSource(c, testRegion, clusterFullAssessRequest())
			return err
		}},
	}
}

func TestClusterAPIRejectsNilClient(t *testing.T) {
	for _, tc := range clusterAPICalls() {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(nil); !errors.Is(err, ErrNilClient) {
				t.Fatalf("expected ErrNilClient, got %v", err)
			}
		})
	}
}

func TestClusterAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		clusterWriteBadRequest(t, w)
	}))
	defer server.Close()

	for _, tc := range clusterAPICalls() {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected the server error to be propagated")
			}
		})
	}
}

// TestClusterAPIRejectsNilRequest covers the nil-request guard of every cluster API that takes a
// request. ListAvailableCoupons is absent because it has no request parameter at all.
//
// v2 ListClusters is included on purpose: unlike its v3 twin (services/bes/v3/api/cluster.go:147,
// which dereferences request.Region before any nil check and therefore panics), the v2 version
// does guard the nil request at cluster.go:74.
func TestClusterAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, clusterUnreachableHandler(t))
	defer server.Close()

	cases := []clusterAPICall{
		{"CreateCluster", func(c bce.Client) error {
			_, err := CreateCluster(c, testRegion, nil)
			return err
		}},
		{"ListClusters", func(c bce.Client) error {
			_, err := ListClusters(c, testRegion, nil)
			return err
		}},
		{"GetClusterDetail", func(c bce.Client) error {
			_, err := GetClusterDetail(c, testRegion, nil)
			return err
		}},
		{"DeleteCluster", func(c bce.Client) error {
			_, err := DeleteCluster(c, testRegion, nil)
			return err
		}},
		{"StartCluster", func(c bce.Client) error {
			_, err := StartCluster(c, testRegion, nil)
			return err
		}},
		{"StopCluster", func(c bce.Client) error {
			_, err := StopCluster(c, testRegion, nil)
			return err
		}},
		{"RestartCluster", func(c bce.Client) error {
			_, err := RestartCluster(c, testRegion, nil)
			return err
		}},
		{"ResizeCluster", func(c bce.Client) error {
			_, err := ResizeCluster(c, testRegion, nil)
			return err
		}},
		{"AddClusterModule", func(c bce.Client) error {
			_, err := AddClusterModule(c, testRegion, nil)
			return err
		}},
		{"ResetClusterPassword", func(c bce.Client) error {
			_, err := ResetClusterPassword(c, testRegion, nil)
			return err
		}},
		{"ToggleClusterHTTPS", func(c bce.Client) error {
			_, err := ToggleClusterHTTPS(c, testRegion, nil)
			return err
		}},
		{"BindClusterEIP", func(c bce.Client) error {
			_, err := BindClusterEIP(c, testRegion, nil)
			return err
		}},
		{"UnbindClusterEIP", func(c bce.Client) error {
			_, err := UnbindClusterEIP(c, testRegion, nil)
			return err
		}},
		{"ToggleClusterMonitor", func(c bce.Client) error {
			_, err := ToggleClusterMonitor(c, testRegion, nil)
			return err
		}},
		{"GetClusterTasks", func(c bce.Client) error {
			_, err := GetClusterTasks(c, testRegion, nil)
			return err
		}},
		{"GetClusterDataSizeTendency", func(c bce.Client) error {
			_, err := GetClusterDataSizeTendency(c, testRegion, nil)
			return err
		}},
		{"AssessClusterSource", func(c bce.Client) error {
			_, err := AssessClusterSource(c, testRegion, nil)
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected error for nil request")
			}
		})
	}
}

// clusterCreateFieldCases covers the 13 required-field branches of CreateCluster, including the
// three checks inside the modules loop.
func clusterCreateFieldCases() []clusterAPICall {
	create := func(mutate func(*CreateClusterRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullCreateRequest()
			mutate(request)
			_, err := CreateCluster(c, testRegion, request)
			return err
		}
	}
	return []clusterAPICall{
		{"CreateCluster empty name", create(func(r *CreateClusterRequest) { r.Name = "" })},
		{"CreateCluster empty password", create(func(r *CreateClusterRequest) { r.Password = "" })},
		{"CreateCluster empty securityGroupId", create(func(r *CreateClusterRequest) {
			r.SecurityGroupId = ""
		})},
		{"CreateCluster empty subnetUuid", create(func(r *CreateClusterRequest) { r.SubnetUuid = "" })},
		{"CreateCluster empty availableZone", create(func(r *CreateClusterRequest) {
			r.AvailableZone = ""
		})},
		{"CreateCluster empty vpcId", create(func(r *CreateClusterRequest) { r.VpcId = "" })},
		{"CreateCluster empty version", create(func(r *CreateClusterRequest) { r.Version = "" })},
		{"CreateCluster empty modules", create(func(r *CreateClusterRequest) { r.Modules = nil })},
		{"CreateCluster empty modules.type", create(func(r *CreateClusterRequest) {
			r.Modules[0].Type = ""
		})},
		{"CreateCluster empty modules.slotType", create(func(r *CreateClusterRequest) {
			r.Modules[0].SlotType = ""
		})},
		{"CreateCluster zero modules.instanceNum", create(func(r *CreateClusterRequest) {
			r.Modules[0].InstanceNum = 0
		})},
		{"CreateCluster nil billing", create(func(r *CreateClusterRequest) { r.Billing = nil })},
		{"CreateCluster empty billing.paymentType", create(func(r *CreateClusterRequest) {
			r.Billing.PaymentType = ""
		})},
	}
}

// clusterSimpleFieldCases covers the APIs whose only required fields are the paging pair or a
// single clusterId.
func clusterSimpleFieldCases() []clusterAPICall {
	return []clusterAPICall{
		{"ListClusters zero pageNo", func(c bce.Client) error {
			request := clusterFullListRequest()
			request.PageNo = 0
			_, err := ListClusters(c, testRegion, request)
			return err
		}},
		{"ListClusters zero pageSize", func(c bce.Client) error {
			request := clusterFullListRequest()
			request.PageSize = 0
			_, err := ListClusters(c, testRegion, request)
			return err
		}},
		{"GetClusterDetail empty clusterId", func(c bce.Client) error {
			_, err := GetClusterDetail(c, testRegion, &GetClusterDetailRequest{})
			return err
		}},
		{"DeleteCluster empty clusterId", func(c bce.Client) error {
			request := clusterFullDeleteRequest()
			request.ClusterId = ""
			_, err := DeleteCluster(c, testRegion, request)
			return err
		}},
		{"StartCluster empty clusterId", func(c bce.Client) error {
			_, err := StartCluster(c, testRegion, &StartClusterRequest{})
			return err
		}},
		{"StopCluster empty clusterId", func(c bce.Client) error {
			_, err := StopCluster(c, testRegion, &StopClusterRequest{})
			return err
		}},
		{"RestartCluster empty clusterId", func(c bce.Client) error {
			request := clusterFullRestartRequest()
			request.ClusterId = ""
			_, err := RestartCluster(c, testRegion, request)
			return err
		}},
		{"ToggleClusterHTTPS empty clusterId", func(c bce.Client) error {
			request := clusterFullToggleHTTPSRequest()
			request.ClusterId = ""
			_, err := ToggleClusterHTTPS(c, testRegion, request)
			return err
		}},
		{"GetClusterTasks empty clusterId", func(c bce.Client) error {
			_, err := GetClusterTasks(c, testRegion, &GetClusterTasksRequest{})
			return err
		}},
	}
}

// clusterScaleFieldCases covers the two scaling APIs, whose validation order differs:
// ResizeCluster checks paymentType before resizeMode, AddClusterModule the other way round.
func clusterScaleFieldCases() []clusterAPICall {
	resize := func(mutate func(*ResizeClusterRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullResizeRequest()
			mutate(request)
			_, err := ResizeCluster(c, testRegion, request)
			return err
		}
	}
	addModule := func(mutate func(*AddClusterModuleRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullAddModuleRequest()
			mutate(request)
			_, err := AddClusterModule(c, testRegion, request)
			return err
		}
	}
	return []clusterAPICall{
		{"ResizeCluster empty clusterId", resize(func(r *ResizeClusterRequest) { r.ClusterId = "" })},
		{"ResizeCluster empty modules", resize(func(r *ResizeClusterRequest) { r.Modules = nil })},
		{"ResizeCluster empty paymentType", resize(func(r *ResizeClusterRequest) {
			r.PaymentType = ""
		})},
		{"ResizeCluster empty resizeMode", resize(func(r *ResizeClusterRequest) { r.ResizeMode = "" })},
		{"AddClusterModule empty clusterId", addModule(func(r *AddClusterModuleRequest) {
			r.ClusterId = ""
		})},
		{"AddClusterModule empty modules", addModule(func(r *AddClusterModuleRequest) {
			r.Modules = nil
		})},
		{"AddClusterModule empty resizeMode", addModule(func(r *AddClusterModuleRequest) {
			r.ResizeMode = ""
		})},
		{"AddClusterModule empty paymentType", addModule(func(r *AddClusterModuleRequest) {
			r.PaymentType = ""
		})},
	}
}

// clusterAccessFieldCases covers password reset, the two EIP APIs and the monitor toggle, whose
// enableMonitor guard is a nil-pointer check rather than an empty-value check.
func clusterAccessFieldCases() []clusterAPICall {
	password := func(mutate func(*ResetClusterPasswordRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullResetPasswordRequest()
			mutate(request)
			_, err := ResetClusterPassword(c, testRegion, request)
			return err
		}
	}
	bind := func(mutate func(*BindClusterEIPRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullBindEIPRequest()
			mutate(request)
			_, err := BindClusterEIP(c, testRegion, request)
			return err
		}
	}
	unbind := func(mutate func(*UnbindClusterEIPRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullUnbindEIPRequest()
			mutate(request)
			_, err := UnbindClusterEIP(c, testRegion, request)
			return err
		}
	}
	monitor := func(mutate func(*ToggleClusterMonitorRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullToggleMonitorRequest()
			mutate(request)
			_, err := ToggleClusterMonitor(c, testRegion, request)
			return err
		}
	}
	return []clusterAPICall{
		{"ResetClusterPassword empty clusterId", password(func(r *ResetClusterPasswordRequest) {
			r.ClusterId = ""
		})},
		{"ResetClusterPassword empty newPassword", password(func(r *ResetClusterPasswordRequest) {
			r.NewPassword = ""
		})},
		{"ResetClusterPassword empty confirmPassword", password(func(r *ResetClusterPasswordRequest) {
			r.ConfirmPassword = ""
		})},
		{"BindClusterEIP empty instanceId", bind(func(r *BindClusterEIPRequest) { r.InstanceId = "" })},
		{"BindClusterEIP empty instanceType", bind(func(r *BindClusterEIPRequest) {
			r.InstanceType = ""
		})},
		{"BindClusterEIP empty eip", bind(func(r *BindClusterEIPRequest) { r.Eip = "" })},
		{"UnbindClusterEIP empty deployId", unbind(func(r *UnbindClusterEIPRequest) {
			r.DeployId = ""
		})},
		{"UnbindClusterEIP empty moduleTemplateName", unbind(func(r *UnbindClusterEIPRequest) {
			r.ModuleTemplateName = ""
		})},
		{"ToggleClusterMonitor empty clusterId", monitor(func(r *ToggleClusterMonitorRequest) {
			r.ClusterId = ""
		})},
		{"ToggleClusterMonitor nil enableMonitor", monitor(func(r *ToggleClusterMonitorRequest) {
			r.EnableMonitor = nil
		})},
	}
}

// clusterAnalyticsFieldCases covers the data size tendency and capacity assessment APIs.
func clusterAnalyticsFieldCases() []clusterAPICall {
	tendency := func(mutate func(*GetClusterDataSizeTendencyRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullDataSizeTendencyRequest()
			mutate(request)
			_, err := GetClusterDataSizeTendency(c, testRegion, request)
			return err
		}
	}
	assess := func(mutate func(*AssessClusterSourceRequest)) func(bce.Client) error {
		return func(c bce.Client) error {
			request := clusterFullAssessRequest()
			mutate(request)
			_, err := AssessClusterSource(c, testRegion, request)
			return err
		}
	}
	return []clusterAPICall{
		{"GetClusterDataSizeTendency empty clusterId", tendency(func(r *GetClusterDataSizeTendencyRequest) {
			r.ClusterId = ""
		})},
		{"GetClusterDataSizeTendency empty datePattern", tendency(func(r *GetClusterDataSizeTendencyRequest) {
			r.DatePattern = ""
		})},
		{"GetClusterDataSizeTendency zero times", tendency(func(r *GetClusterDataSizeTendencyRequest) {
			r.Times = 0
		})},
		{"GetClusterDataSizeTendency empty timeUnit", tendency(func(r *GetClusterDataSizeTendencyRequest) {
			r.TimeUnit = ""
		})},
		{"AssessClusterSource empty useType", assess(func(r *AssessClusterSourceRequest) {
			r.UseType = ""
		})},
		{"AssessClusterSource zero originDataSize", assess(func(r *AssessClusterSourceRequest) {
			r.OriginDataSize = 0
		})},
		{"AssessClusterSource empty originDataSizeUnit", assess(func(r *AssessClusterSourceRequest) {
			r.OriginDataSizeUnit = ""
		})},
		{"AssessClusterSource zero dataAdd", assess(func(r *AssessClusterSourceRequest) {
			r.DataAdd = 0
		})},
		{"AssessClusterSource empty dataAddUnit", assess(func(r *AssessClusterSourceRequest) {
			r.DataAddUnit = ""
		})},
		{"AssessClusterSource zero storageDays", assess(func(r *AssessClusterSourceRequest) {
			r.StorageDays = 0
		})},
		{"AssessClusterSource zero writeThroughput", assess(func(r *AssessClusterSourceRequest) {
			r.WriteThroughput = 0
		})},
		{"AssessClusterSource empty replica", assess(func(r *AssessClusterSourceRequest) {
			r.Replica = ""
		})},
	}
}

// TestClusterAPIRejectsMissingRequiredFields covers all 52 required-field validation branches of
// cluster.go. ListAvailableCoupons contributes none: it takes no request at all, so its only local
// guard is the nil-client check.
func TestClusterAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, clusterUnreachableHandler(t))
	defer server.Close()

	var cases []clusterAPICall
	cases = append(cases, clusterCreateFieldCases()...)
	cases = append(cases, clusterSimpleFieldCases()...)
	cases = append(cases, clusterScaleFieldCases()...)
	cases = append(cases, clusterAccessFieldCases()...)
	cases = append(cases, clusterAnalyticsFieldCases()...)
	if len(cases) != 52 {
		t.Fatalf("expected 52 required-field cases, got %d", len(cases))
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected a required-field validation error")
			}
		})
	}
}
