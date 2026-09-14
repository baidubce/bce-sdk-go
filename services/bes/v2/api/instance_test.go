package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	instanceTestClusterId  = "218185657699405824"
	instanceTestInstanceId = "276316903033671680"
	instanceTestModuleType = "es_node"
	instanceTestOrderId    = "186f2566f2884534a65ed8cdc2499d82"
	instanceTestRequestId  = "9485f297f8a64d8da79dde31791c08a6"
)

// instanceWriteBadRequest writes a 400 service error. A 4xx status is used on purpose: the test
// client mirrors the production DEFAULT_RETRY_POLICY, so a 5xx would trigger backoff retries.
func instanceWriteBadRequest(t *testing.T, w nethttp.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(nethttp.StatusBadRequest)
	if _, err := w.Write([]byte(`{"code":"BadRequest","message":"invalid instance request"}`)); err != nil {
		t.Fatalf("write error response failed: %v", err)
	}
}

// instanceUnreachableHandler fails the test when a request reaches the server, used by cases that
// must fail during local validation.
func instanceUnreachableHandler(t *testing.T) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Errorf("unexpected request reached the server: %s", r.URL.Path)
		instanceWriteBadRequest(t, w)
	}
}

// instanceAssertRequest checks the parts of the request every instance API builds the same way.
func instanceAssertRequest(t *testing.T, r *nethttp.Request, path string) {
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

// The instanceFullXxxRequest builders return fully populated requests, so the required-field
// cases below only have to clear one field each.
func instanceFullOperationRequest() *InstanceOperationRequest {
	return &InstanceOperationRequest{
		ClusterId:  instanceTestClusterId,
		InstanceId: instanceTestInstanceId,
	}
}

func instanceFullBatchRequest() *BatchInstanceOperationRequest {
	return &BatchInstanceOperationRequest{
		ClusterId:      instanceTestClusterId,
		InstanceIdList: []string{instanceTestInstanceId},
	}
}

func instanceFullDeleteRequest() *DeleteInstancesRequest {
	return &DeleteInstancesRequest{
		ClusterId:      instanceTestClusterId,
		InstanceIdList: []string{instanceTestInstanceId},
		ModuleType:     instanceTestModuleType,
	}
}

func instanceFullListScaleInRequest() *ListScaleInInstancesRequest {
	return &ListScaleInInstancesRequest{
		ClusterId:  instanceTestClusterId,
		ModuleType: instanceTestModuleType,
	}
}

func instanceFullConfirmMigrationRequest() *ConfirmDataMigrationRequest {
	return &ConfirmDataMigrationRequest{
		ClusterId:      instanceTestClusterId,
		ModuleType:     instanceTestModuleType,
		InstanceIdList: []string{instanceTestInstanceId},
	}
}

func instanceFullRollbackMigrationRequest() *RollbackDataMigrationRequest {
	return &RollbackDataMigrationRequest{
		ClusterId: instanceTestClusterId,
		RequestId: instanceTestRequestId,
	}
}

func instanceFullListMigrationRequest() *ListDataMigrationInstancesRequest {
	return &ListDataMigrationInstancesRequest{
		ClusterId:  instanceTestClusterId,
		ModuleType: instanceTestModuleType,
	}
}

func instanceFullSuggestMigrationRequest() *SuggestDataMigrationInstancesRequest {
	return &SuggestDataMigrationInstancesRequest{
		ClusterId:    instanceTestClusterId,
		ModuleType:   instanceTestModuleType,
		MigrateCount: 2,
	}
}

func TestStartInstance(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/v2/instance/start")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["instanceId"] != instanceTestInstanceId {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"taskId": "start-1"},
		})
	}))
	defer server.Close()

	result, err := StartInstance(cli, testRegion, instanceFullOperationRequest())
	if err != nil {
		t.Fatalf("start instance failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	payload, ok := result.Result.(map[string]interface{})
	if !ok || payload["taskId"] != "start-1" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestStopInstance(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/v2/instance/stop")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["instanceId"] != instanceTestInstanceId {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"taskId": "stop-1"},
		})
	}))
	defer server.Close()

	result, err := StopInstance(cli, testRegion, instanceFullOperationRequest())
	if err != nil {
		t.Fatalf("stop instance failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	payload, ok := result.Result.(map[string]interface{})
	if !ok || payload["taskId"] != "stop-1" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestBatchStartInstances(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/v2/instance/batchStart")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId {
			t.Fatalf("unexpected clusterId: %v", body["clusterId"])
		}
		ids, ok := body["instanceIdList"].([]interface{})
		if !ok || len(ids) != 1 || ids[0] != instanceTestInstanceId {
			t.Fatalf("unexpected instanceIdList: %v", body["instanceIdList"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"taskId": "batch-start-1"},
		})
	}))
	defer server.Close()

	result, err := BatchStartInstances(cli, testRegion, instanceFullBatchRequest())
	if err != nil {
		t.Fatalf("batch start instances failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	payload, ok := result.Result.(map[string]interface{})
	if !ok || payload["taskId"] != "batch-start-1" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestBatchStopInstances(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/v2/instance/batchStop")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId {
			t.Fatalf("unexpected clusterId: %v", body["clusterId"])
		}
		ids, ok := body["instanceIdList"].([]interface{})
		if !ok || len(ids) != 1 || ids[0] != instanceTestInstanceId {
			t.Fatalf("unexpected instanceIdList: %v", body["instanceIdList"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"taskId": "batch-stop-1"},
		})
	}))
	defer server.Close()

	result, err := BatchStopInstances(cli, testRegion, instanceFullBatchRequest())
	if err != nil {
		t.Fatalf("batch stop instances failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	payload, ok := result.Result.(map[string]interface{})
	if !ok || payload["taskId"] != "batch-stop-1" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestDeleteInstances(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/v2/instance/delete")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["moduleType"] != instanceTestModuleType {
			t.Fatalf("unexpected body: %+v", body)
		}
		ids, ok := body["instanceIdList"].([]interface{})
		if !ok || len(ids) != 1 || ids[0] != instanceTestInstanceId {
			t.Fatalf("unexpected instanceIdList: %v", body["instanceIdList"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"orderId": instanceTestOrderId},
		})
	}))
	defer server.Close()

	result, err := DeleteInstances(cli, testRegion, instanceFullDeleteRequest())
	if err != nil {
		t.Fatalf("delete instances failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result == nil || result.Result.OrderId != instanceTestOrderId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestListScaleInInstances(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/v2/scalein/instances")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["moduleType"] != instanceTestModuleType {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"instanceInfoList": []map[string]interface{}{
					{
						"availableZone": "zoneA",
						"instanceList": []map[string]interface{}{
							{
								"instanceId": "798832817483157504",
								"hostIp":     "192.0.2.10",
								"isMaster":   true,
								"isEmpty":    true,
							},
						},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListScaleInInstances(cli, testRegion, instanceFullListScaleInRequest())
	if err != nil {
		t.Fatalf("list scale-in instances failed: %v", err)
	}
	if result.Result == nil || len(result.Result.InstanceInfoList) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	info := result.Result.InstanceInfoList[0]
	if info.AvailableZone != "zoneA" || len(info.InstanceList) != 1 {
		t.Fatalf("unexpected instance info: %+v", info)
	}
	item := info.InstanceList[0]
	if item.InstanceId != "798832817483157504" || item.HostIp != "192.0.2.10" {
		t.Fatalf("unexpected instance item: %+v", item)
	}
	if !item.IsMaster || !item.IsEmpty {
		t.Fatalf("unexpected instance flags: %+v", item)
	}
}

func TestConfirmDataMigration(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/migrate/v1/confirm")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["moduleType"] != instanceTestModuleType {
			t.Fatalf("unexpected body: %+v", body)
		}
		ids, ok := body["instanceIdList"].([]interface{})
		if !ok || len(ids) != 1 || ids[0] != instanceTestInstanceId {
			t.Fatalf("unexpected instanceIdList: %v", body["instanceIdList"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"requestId": instanceTestRequestId},
		})
	}))
	defer server.Close()

	result, err := ConfirmDataMigration(cli, testRegion, instanceFullConfirmMigrationRequest())
	if err != nil {
		t.Fatalf("confirm data migration failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	payload, ok := result.Result.(map[string]interface{})
	if !ok || payload["requestId"] != instanceTestRequestId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestRollbackDataMigration(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/migrate/v1/rollback")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["requestId"] != instanceTestRequestId {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  "rollback submitted",
		})
	}))
	defer server.Close()

	result, err := RollbackDataMigration(cli, testRegion, instanceFullRollbackMigrationRequest())
	if err != nil {
		t.Fatalf("rollback data migration failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result != "rollback submitted" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestListDataMigrationInstances(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/migrate/v1/instances")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["moduleType"] != instanceTestModuleType {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"instanceInfoList": []map[string]interface{}{
					{
						"availableZone": "zoneA",
						"instanceList": []map[string]interface{}{
							{
								"instanceId": "798832817483157502",
								"hostIp":     "192.0.2.11",
								"isMaster":   true,
							},
						},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListDataMigrationInstances(cli, testRegion, instanceFullListMigrationRequest())
	if err != nil {
		t.Fatalf("list data migration instances failed: %v", err)
	}
	if result.Result == nil || len(result.Result.InstanceInfoList) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	info := result.Result.InstanceInfoList[0]
	if info.AvailableZone != "zoneA" || len(info.InstanceList) != 1 {
		t.Fatalf("unexpected instance info: %+v", info)
	}
	item := info.InstanceList[0]
	if item.InstanceId != "798832817483157502" || item.HostIp != "192.0.2.11" || !item.IsMaster {
		t.Fatalf("unexpected instance item: %+v", item)
	}
}

func TestSuggestDataMigrationInstances(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceAssertRequest(t, r, "/api/bes/cluster/migrate/v1/suggestInstances")
		body := readJSONBody(t, r)
		if body["clusterId"] != instanceTestClusterId || body["moduleType"] != instanceTestModuleType {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["migrateCount"] != float64(2) {
			t.Fatalf("unexpected migrateCount: %v", body["migrateCount"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"suggestList": []map[string]interface{}{
					{
						"availableZone": "zoneA",
						"instanceList": []map[string]interface{}{
							{"instanceId": "798832817483157502", "hostIp": "192.0.2.12"},
						},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := SuggestDataMigrationInstances(cli, testRegion, instanceFullSuggestMigrationRequest())
	if err != nil {
		t.Fatalf("suggest data migration instances failed: %v", err)
	}
	if result.Result == nil || len(result.Result.SuggestList) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	suggest := result.Result.SuggestList[0]
	if suggest.AvailableZone != "zoneA" || len(suggest.InstanceList) != 1 {
		t.Fatalf("unexpected suggest info: %+v", suggest)
	}
	item := suggest.InstanceList[0]
	if item.InstanceId != "798832817483157502" || item.HostIp != "192.0.2.12" {
		t.Fatalf("unexpected suggest item: %+v", item)
	}
}

func TestInstanceAPIRejectsNilClient(t *testing.T) {
	cases := []struct {
		name string
		call func() error
	}{
		{"StartInstance", func() error {
			_, err := StartInstance(nil, testRegion, instanceFullOperationRequest())
			return err
		}},
		{"StopInstance", func() error {
			_, err := StopInstance(nil, testRegion, instanceFullOperationRequest())
			return err
		}},
		{"BatchStartInstances", func() error {
			_, err := BatchStartInstances(nil, testRegion, instanceFullBatchRequest())
			return err
		}},
		{"BatchStopInstances", func() error {
			_, err := BatchStopInstances(nil, testRegion, instanceFullBatchRequest())
			return err
		}},
		{"DeleteInstances", func() error {
			_, err := DeleteInstances(nil, testRegion, instanceFullDeleteRequest())
			return err
		}},
		{"ListScaleInInstances", func() error {
			_, err := ListScaleInInstances(nil, testRegion, instanceFullListScaleInRequest())
			return err
		}},
		{"ConfirmDataMigration", func() error {
			_, err := ConfirmDataMigration(nil, testRegion, instanceFullConfirmMigrationRequest())
			return err
		}},
		{"RollbackDataMigration", func() error {
			_, err := RollbackDataMigration(nil, testRegion, instanceFullRollbackMigrationRequest())
			return err
		}},
		{"ListDataMigrationInstances", func() error {
			_, err := ListDataMigrationInstances(nil, testRegion, instanceFullListMigrationRequest())
			return err
		}},
		{"SuggestDataMigrationInstances", func() error {
			_, err := SuggestDataMigrationInstances(nil, testRegion, instanceFullSuggestMigrationRequest())
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(); !errors.Is(err, ErrNilClient) {
				t.Fatalf("expected ErrNilClient, got %v", err)
			}
		})
	}
}

// TestInstanceAPIRejectsNilRequest covers the nil-request guard of every instance API: all ten
// take a mandatory request, so none of them has a positive nil-request case.
func TestInstanceAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, instanceUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"StartInstance", func(c bce.Client) error {
			_, err := StartInstance(c, testRegion, nil)
			return err
		}},
		{"StopInstance", func(c bce.Client) error {
			_, err := StopInstance(c, testRegion, nil)
			return err
		}},
		{"BatchStartInstances", func(c bce.Client) error {
			_, err := BatchStartInstances(c, testRegion, nil)
			return err
		}},
		{"BatchStopInstances", func(c bce.Client) error {
			_, err := BatchStopInstances(c, testRegion, nil)
			return err
		}},
		{"DeleteInstances", func(c bce.Client) error {
			_, err := DeleteInstances(c, testRegion, nil)
			return err
		}},
		{"ListScaleInInstances", func(c bce.Client) error {
			_, err := ListScaleInInstances(c, testRegion, nil)
			return err
		}},
		{"ConfirmDataMigration", func(c bce.Client) error {
			_, err := ConfirmDataMigration(c, testRegion, nil)
			return err
		}},
		{"RollbackDataMigration", func(c bce.Client) error {
			_, err := RollbackDataMigration(c, testRegion, nil)
			return err
		}},
		{"ListDataMigrationInstances", func(c bce.Client) error {
			_, err := ListDataMigrationInstances(c, testRegion, nil)
			return err
		}},
		{"SuggestDataMigrationInstances", func(c bce.Client) error {
			_, err := SuggestDataMigrationInstances(c, testRegion, nil)
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

// TestInstanceAPIRejectsMissingRequiredFields covers all 23 required-field validation branches in
// instance.go: 2 for start/stop, 2 for each batch start/stop, 3 for delete, 2 for list scale-in,
// 3 for confirm, 2 for rollback, 2 for list migration and 3 for suggest.
func TestInstanceAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, instanceUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"StartInstance empty clusterId", func(c bce.Client) error {
			request := instanceFullOperationRequest()
			request.ClusterId = ""
			_, err := StartInstance(c, testRegion, request)
			return err
		}},
		{"StartInstance empty instanceId", func(c bce.Client) error {
			request := instanceFullOperationRequest()
			request.InstanceId = ""
			_, err := StartInstance(c, testRegion, request)
			return err
		}},
		{"StopInstance empty clusterId", func(c bce.Client) error {
			request := instanceFullOperationRequest()
			request.ClusterId = ""
			_, err := StopInstance(c, testRegion, request)
			return err
		}},
		{"StopInstance empty instanceId", func(c bce.Client) error {
			request := instanceFullOperationRequest()
			request.InstanceId = ""
			_, err := StopInstance(c, testRegion, request)
			return err
		}},
		{"BatchStartInstances empty clusterId", func(c bce.Client) error {
			request := instanceFullBatchRequest()
			request.ClusterId = ""
			_, err := BatchStartInstances(c, testRegion, request)
			return err
		}},
		{"BatchStartInstances empty instanceIdList", func(c bce.Client) error {
			request := instanceFullBatchRequest()
			request.InstanceIdList = nil
			_, err := BatchStartInstances(c, testRegion, request)
			return err
		}},
		{"BatchStopInstances empty clusterId", func(c bce.Client) error {
			request := instanceFullBatchRequest()
			request.ClusterId = ""
			_, err := BatchStopInstances(c, testRegion, request)
			return err
		}},
		{"BatchStopInstances empty instanceIdList", func(c bce.Client) error {
			request := instanceFullBatchRequest()
			request.InstanceIdList = nil
			_, err := BatchStopInstances(c, testRegion, request)
			return err
		}},
		{"DeleteInstances empty clusterId", func(c bce.Client) error {
			request := instanceFullDeleteRequest()
			request.ClusterId = ""
			_, err := DeleteInstances(c, testRegion, request)
			return err
		}},
		{"DeleteInstances empty instanceIdList", func(c bce.Client) error {
			request := instanceFullDeleteRequest()
			request.InstanceIdList = nil
			_, err := DeleteInstances(c, testRegion, request)
			return err
		}},
		{"DeleteInstances empty moduleType", func(c bce.Client) error {
			request := instanceFullDeleteRequest()
			request.ModuleType = ""
			_, err := DeleteInstances(c, testRegion, request)
			return err
		}},
		{"ListScaleInInstances empty clusterId", func(c bce.Client) error {
			request := instanceFullListScaleInRequest()
			request.ClusterId = ""
			_, err := ListScaleInInstances(c, testRegion, request)
			return err
		}},
		{"ListScaleInInstances empty moduleType", func(c bce.Client) error {
			request := instanceFullListScaleInRequest()
			request.ModuleType = ""
			_, err := ListScaleInInstances(c, testRegion, request)
			return err
		}},
		{"ConfirmDataMigration empty clusterId", func(c bce.Client) error {
			request := instanceFullConfirmMigrationRequest()
			request.ClusterId = ""
			_, err := ConfirmDataMigration(c, testRegion, request)
			return err
		}},
		{"ConfirmDataMigration empty moduleType", func(c bce.Client) error {
			request := instanceFullConfirmMigrationRequest()
			request.ModuleType = ""
			_, err := ConfirmDataMigration(c, testRegion, request)
			return err
		}},
		{"ConfirmDataMigration empty instanceIdList", func(c bce.Client) error {
			request := instanceFullConfirmMigrationRequest()
			request.InstanceIdList = nil
			_, err := ConfirmDataMigration(c, testRegion, request)
			return err
		}},
		{"RollbackDataMigration empty clusterId", func(c bce.Client) error {
			request := instanceFullRollbackMigrationRequest()
			request.ClusterId = ""
			_, err := RollbackDataMigration(c, testRegion, request)
			return err
		}},
		{"RollbackDataMigration empty requestId", func(c bce.Client) error {
			request := instanceFullRollbackMigrationRequest()
			request.RequestId = ""
			_, err := RollbackDataMigration(c, testRegion, request)
			return err
		}},
		{"ListDataMigrationInstances empty clusterId", func(c bce.Client) error {
			request := instanceFullListMigrationRequest()
			request.ClusterId = ""
			_, err := ListDataMigrationInstances(c, testRegion, request)
			return err
		}},
		{"ListDataMigrationInstances empty moduleType", func(c bce.Client) error {
			request := instanceFullListMigrationRequest()
			request.ModuleType = ""
			_, err := ListDataMigrationInstances(c, testRegion, request)
			return err
		}},
		{"SuggestDataMigrationInstances empty clusterId", func(c bce.Client) error {
			request := instanceFullSuggestMigrationRequest()
			request.ClusterId = ""
			_, err := SuggestDataMigrationInstances(c, testRegion, request)
			return err
		}},
		{"SuggestDataMigrationInstances empty moduleType", func(c bce.Client) error {
			request := instanceFullSuggestMigrationRequest()
			request.ModuleType = ""
			_, err := SuggestDataMigrationInstances(c, testRegion, request)
			return err
		}},
		{"SuggestDataMigrationInstances zero migrateCount", func(c bce.Client) error {
			request := instanceFullSuggestMigrationRequest()
			request.MigrateCount = 0
			_, err := SuggestDataMigrationInstances(c, testRegion, request)
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected a required-field validation error")
			}
		})
	}
}

func TestInstanceAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		instanceWriteBadRequest(t, w)
	}))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"StartInstance", func(c bce.Client) error {
			_, err := StartInstance(c, testRegion, instanceFullOperationRequest())
			return err
		}},
		{"StopInstance", func(c bce.Client) error {
			_, err := StopInstance(c, testRegion, instanceFullOperationRequest())
			return err
		}},
		{"BatchStartInstances", func(c bce.Client) error {
			_, err := BatchStartInstances(c, testRegion, instanceFullBatchRequest())
			return err
		}},
		{"BatchStopInstances", func(c bce.Client) error {
			_, err := BatchStopInstances(c, testRegion, instanceFullBatchRequest())
			return err
		}},
		{"DeleteInstances", func(c bce.Client) error {
			_, err := DeleteInstances(c, testRegion, instanceFullDeleteRequest())
			return err
		}},
		{"ListScaleInInstances", func(c bce.Client) error {
			_, err := ListScaleInInstances(c, testRegion, instanceFullListScaleInRequest())
			return err
		}},
		{"ConfirmDataMigration", func(c bce.Client) error {
			_, err := ConfirmDataMigration(c, testRegion, instanceFullConfirmMigrationRequest())
			return err
		}},
		{"RollbackDataMigration", func(c bce.Client) error {
			_, err := RollbackDataMigration(c, testRegion, instanceFullRollbackMigrationRequest())
			return err
		}},
		{"ListDataMigrationInstances", func(c bce.Client) error {
			_, err := ListDataMigrationInstances(c, testRegion, instanceFullListMigrationRequest())
			return err
		}},
		{"SuggestDataMigrationInstances", func(c bce.Client) error {
			_, err := SuggestDataMigrationInstances(c, testRegion, instanceFullSuggestMigrationRequest())
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected the server error to be propagated")
			}
		})
	}
}
