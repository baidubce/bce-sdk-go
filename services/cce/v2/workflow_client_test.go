package v2

import (
	"encoding/json"
	"io"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
)

func newWorkflowTestClient(t *testing.T, handler func(nethttp.ResponseWriter, *nethttp.Request, []byte)) *Client {
	t.Helper()

	server := httptest.NewServer(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body failed: %v", err)
		}

		handler(w, r, body)
	}))
	t.Cleanup(server.Close)

	client, err := NewClient("test-ak", "test-sk", server.URL)
	if err != nil {
		t.Fatalf("create test client failed: %v", err)
	}
	client.Config.Credentials = nil

	return client
}

func TestGetInstanceGroupUpgradeComponentVersionsValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    *GetInstanceGroupUpgradeComponentVersionsArgs
		wantErr string
	}{
		{
			name:    "nil args",
			args:    nil,
			wantErr: "args is nil",
		},
		{
			name: "empty cluster id",
			args: &GetInstanceGroupUpgradeComponentVersionsArgs{
				InstanceGroupID: "ig-1",
			},
			wantErr: "clusterID is empty",
		},
		{
			name: "empty instance group id",
			args: &GetInstanceGroupUpgradeComponentVersionsArgs{
				ClusterID: "c-1",
			},
			wantErr: "instanceGroupID is empty",
		},
	}

	client := &Client{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.GetInstanceGroupUpgradeComponentVersions(tt.args)
			if err == nil || err.Error() != tt.wantErr {
				t.Fatalf("expected error %q, got %v", tt.wantErr, err)
			}
			if resp != nil {
				t.Fatalf("expected nil response, got %#v", resp)
			}
		})
	}
}

func TestCreateWorkflowValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    *CreateWorkflowArgs
		wantErr string
	}{
		{
			name:    "nil args",
			args:    nil,
			wantErr: "args is nil",
		},
		{
			name: "empty cluster id",
			args: &CreateWorkflowArgs{
				Request: &CreateWorkflowRequest{},
			},
			wantErr: "clusterID is empty",
		},
		{
			name: "nil request",
			args: &CreateWorkflowArgs{
				ClusterID: "c-1",
			},
			wantErr: "request is nil",
		},
	}

	client := &Client{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.CreateWorkflow(tt.args)
			if err == nil || err.Error() != tt.wantErr {
				t.Fatalf("expected error %q, got %v", tt.wantErr, err)
			}
			if resp != nil {
				t.Fatalf("expected nil response, got %#v", resp)
			}
		})
	}
}

func TestGetWorkflowValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    *GetWorkflowArgs
		wantErr string
	}{
		{
			name:    "nil args",
			args:    nil,
			wantErr: "args is nil",
		},
		{
			name: "empty cluster id",
			args: &GetWorkflowArgs{
				WorkflowID: "wf-1",
			},
			wantErr: "clusterID is empty",
		},
		{
			name: "empty workflow id",
			args: &GetWorkflowArgs{
				ClusterID: "c-1",
			},
			wantErr: "workflowID is empty",
		},
	}

	client := &Client{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.GetWorkflow(tt.args)
			if err == nil || err.Error() != tt.wantErr {
				t.Fatalf("expected error %q, got %v", tt.wantErr, err)
			}
			if resp != nil {
				t.Fatalf("expected nil response, got %#v", resp)
			}
		})
	}
}

func TestUpdateWorkflowValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    *UpdateWorkflowArgs
		wantErr string
	}{
		{
			name:    "nil args",
			args:    nil,
			wantErr: "args is nil",
		},
		{
			name: "empty cluster id",
			args: &UpdateWorkflowArgs{
				WorkflowID: "wf-1",
				Request:    &UpdateWorkflowRequest{},
			},
			wantErr: "clusterID is empty",
		},
		{
			name: "empty workflow id",
			args: &UpdateWorkflowArgs{
				ClusterID: "c-1",
				Request:   &UpdateWorkflowRequest{},
			},
			wantErr: "workflowID is empty",
		},
		{
			name: "nil request",
			args: &UpdateWorkflowArgs{
				ClusterID:  "c-1",
				WorkflowID: "wf-1",
			},
			wantErr: "request is nil",
		},
	}

	client := &Client{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.UpdateWorkflow(tt.args)
			if err == nil || err.Error() != tt.wantErr {
				t.Fatalf("expected error %q, got %v", tt.wantErr, err)
			}
			if resp != nil {
				t.Fatalf("expected nil response, got %#v", resp)
			}
		})
	}
}

func TestClientGetInstanceGroupUpgradeComponentVersions(t *testing.T) {
	client := newWorkflowTestClient(t, func(w nethttp.ResponseWriter, r *nethttp.Request, body []byte) {
		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != getInstanceGroupUpgradeComponentVersionsURI("c-1", "ig-1") {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if len(body) != 0 {
			t.Fatalf("unexpected request body: %s", string(body))
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&GetInstanceGroupUpgradeComponentVersionsResponse{
			RequestID: "req-1",
			Result: &types.UpgradeComponents{
				GPUDriver: types.UpgradeVersionList{
					GPUCurrentVersion: &types.GPUVersion{
						Driver: "570.148.08",
						CUDA:   "12.8.0",
						CuDNN:  "9.8.0",
					},
					ComponentVersions: []types.ComponentVersion{
						{
							NeedDrainNode: true,
							GPUTargetVersion: &types.GPUVersion{
								Driver: "570.172.08",
								CUDA:   "12.8.0",
								CuDNN:  "9.8.0",
							},
						},
					},
				},
			},
		}); err != nil {
			t.Fatalf("encode response failed: %v", err)
		}
	})

	resp, err := client.GetInstanceGroupUpgradeComponentVersions(&GetInstanceGroupUpgradeComponentVersionsArgs{
		ClusterID:       "c-1",
		InstanceGroupID: "ig-1",
	})
	if err != nil {
		t.Fatalf("GetInstanceGroupUpgradeComponentVersions failed: %v", err)
	}
	if resp.RequestID != "req-1" {
		t.Fatalf("unexpected requestID: %q", resp.RequestID)
	}
	if resp.Result == nil || resp.Result.GPUDriver.GPUCurrentVersion == nil {
		t.Fatalf("gpu current version missing in response: %#v", resp.Result)
	}
	if resp.Result.GPUDriver.GPUCurrentVersion.Driver != "570.148.08" {
		t.Fatalf("unexpected gpu current driver: %q", resp.Result.GPUDriver.GPUCurrentVersion.Driver)
	}
	if len(resp.Result.GPUDriver.ComponentVersions) != 1 || resp.Result.GPUDriver.ComponentVersions[0].GPUTargetVersion == nil {
		t.Fatalf("gpu target version missing in response: %#v", resp.Result.GPUDriver.ComponentVersions)
	}
	if resp.Result.GPUDriver.ComponentVersions[0].GPUTargetVersion.Driver != "570.172.08" {
		t.Fatalf("unexpected gpu target driver: %q", resp.Result.GPUDriver.ComponentVersions[0].GPUTargetVersion.Driver)
	}
}

func TestClientCreateWorkflow(t *testing.T) {
	client := newWorkflowTestClient(t, func(w nethttp.ResponseWriter, r *nethttp.Request, body []byte) {
		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != getClusterWorkflowURI("c-1") {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var req CreateWorkflowRequest
		if err := json.Unmarshal(body, &req); err != nil {
			t.Fatalf("unmarshal request failed: %v", err)
		}
		if req.WorkflowType != types.WorkflowTypeUpgradeNodes {
			t.Fatalf("unexpected workflow type: %q", req.WorkflowType)
		}
		if req.WorkflowConfig.UpgradeNodesWorkflowConfig == nil {
			t.Fatalf("upgradeNodesWorkflowConfig missing in request")
		}
		if req.WorkflowConfig.UpgradeNodesWorkflowConfig.InstanceGroupID != "ig-1" {
			t.Fatalf("unexpected instanceGroupID: %q", req.WorkflowConfig.UpgradeNodesWorkflowConfig.InstanceGroupID)
		}
		if len(req.WorkflowConfig.UpgradeNodesWorkflowConfig.Components) != 1 {
			t.Fatalf("unexpected component count: %d", len(req.WorkflowConfig.UpgradeNodesWorkflowConfig.Components))
		}
		component := req.WorkflowConfig.UpgradeNodesWorkflowConfig.Components[0]
		if component.Name != types.ComponentGPUDriver {
			t.Fatalf("unexpected component name: %q", component.Name)
		}
		if component.GPUTargetVersion == nil || component.GPUTargetVersion.Driver != "570.172.08" {
			t.Fatalf("unexpected gpu target version: %#v", component.GPUTargetVersion)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&CreateWorkflowResponse{
			WorkflowID: "wf-1",
			RequestID:  "req-2",
		}); err != nil {
			t.Fatalf("encode response failed: %v", err)
		}
	})

	pausePolicy := types.NotPause
	drainNodeBeforeUpgrade := true
	batchIntervalMinutes := 5

	resp, err := client.CreateWorkflow(&CreateWorkflowArgs{
		ClusterID: "c-1",
		Request: &CreateWorkflowRequest{
			WorkflowType: types.WorkflowTypeUpgradeNodes,
			WorkflowConfig: types.WorkflowConfig{
				UpgradeNodesWorkflowConfig: &types.UpgradeNodesWorkflowConfig{
					InstanceGroupID:        "ig-1",
					NodeUpgradeBatchSize:   1,
					DrainNodeBeforeUpgrade: &drainNodeBeforeUpgrade,
					PausePolicy:            &pausePolicy,
					BatchIntervalMinutes:   &batchIntervalMinutes,
					Components: []types.Component{
						{
							Name: types.ComponentGPUDriver,
							GPUTargetVersion: &types.GPUVersion{
								Driver: "570.172.08",
								CUDA:   "12.8.0",
								CuDNN:  "9.8.0",
							},
						},
					},
				},
			},
			WatchDogConfig: types.WatchDogConfig{
				UnhealthyPodsPercent: 20,
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateWorkflow failed: %v", err)
	}
	if resp.WorkflowID != "wf-1" || resp.RequestID != "req-2" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestClientGetWorkflow(t *testing.T) {
	client := newWorkflowTestClient(t, func(w nethttp.ResponseWriter, r *nethttp.Request, body []byte) {
		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != getClusterWorkflowWithIDURI("c-1", "wf-1") {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if len(body) != 0 {
			t.Fatalf("unexpected request body: %s", string(body))
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&GetWorkflowResponse{
			Workflow: &Workflow{
				Spec: &types.WorkflowSpec{
					WorkflowID:   "wf-1",
					ClusterID:    "c-1",
					WorkflowType: types.WorkflowTypeUpgradeNodes,
					WorkflowConfig: types.WorkflowConfig{
						UpgradeNodesWorkflowConfig: &types.UpgradeNodesWorkflowConfig{
							InstanceGroupID: "ig-1",
						},
					},
				},
				Status: &types.WorkflowStatus{
					WorkflowPhase:     types.WorkflowPhaseSucceeded,
					TotalTaskCount:    1,
					FinishedTaskCount: 1,
				},
			},
		}); err != nil {
			t.Fatalf("encode response failed: %v", err)
		}
	})

	resp, err := client.GetWorkflow(&GetWorkflowArgs{
		ClusterID:  "c-1",
		WorkflowID: "wf-1",
	})
	if err != nil {
		t.Fatalf("GetWorkflow failed: %v", err)
	}
	if resp.Workflow == nil || resp.Workflow.Spec == nil || resp.Workflow.Status == nil {
		t.Fatalf("workflow payload missing: %#v", resp)
	}
	if resp.Workflow.Spec.WorkflowType != types.WorkflowTypeUpgradeNodes {
		t.Fatalf("unexpected workflow type: %q", resp.Workflow.Spec.WorkflowType)
	}
	if resp.Workflow.Spec.WorkflowConfig.UpgradeNodesWorkflowConfig == nil ||
		resp.Workflow.Spec.WorkflowConfig.UpgradeNodesWorkflowConfig.InstanceGroupID != "ig-1" {
		t.Fatalf("unexpected workflow config: %#v", resp.Workflow.Spec.WorkflowConfig)
	}
	if resp.Workflow.Status.WorkflowPhase != types.WorkflowPhaseSucceeded {
		t.Fatalf("unexpected workflow phase: %q", resp.Workflow.Status.WorkflowPhase)
	}
}

func TestClientUpdateWorkflow(t *testing.T) {
	client := newWorkflowTestClient(t, func(w nethttp.ResponseWriter, r *nethttp.Request, body []byte) {
		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != getClusterWorkflowWithIDURI("c-1", "wf-1") {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var req UpdateWorkflowRequest
		if err := json.Unmarshal(body, &req); err != nil {
			t.Fatalf("unmarshal request failed: %v", err)
		}
		if req.Action != UpdateWorkflowActionUpdateSpec {
			t.Fatalf("unexpected action: %q", req.Action)
		}
		if req.Spec == nil {
			t.Fatalf("spec missing in update request")
		}
		if req.Spec.WorkflowType != types.WorkflowTypeUpgradeNodes {
			t.Fatalf("unexpected workflow type: %q", req.Spec.WorkflowType)
		}
		if req.Spec.WorkflowConfig.UpgradeNodesWorkflowConfig == nil ||
			req.Spec.WorkflowConfig.UpgradeNodesWorkflowConfig.InstanceGroupID != "ig-1" {
			t.Fatalf("unexpected workflow config: %#v", req.Spec.WorkflowConfig)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&UpdateWorkflowResponse{
			ClusterID:  "c-1",
			WorkflowID: "wf-1",
			RequestID:  "req-3",
		}); err != nil {
			t.Fatalf("encode response failed: %v", err)
		}
	})

	resp, err := client.UpdateWorkflow(&UpdateWorkflowArgs{
		ClusterID:  "c-1",
		WorkflowID: "wf-1",
		Request: &UpdateWorkflowRequest{
			Action: UpdateWorkflowActionUpdateSpec,
			Spec: &types.WorkflowSpec{
				WorkflowID:   "wf-1",
				ClusterID:    "c-1",
				WorkflowType: types.WorkflowTypeUpgradeNodes,
				WorkflowConfig: types.WorkflowConfig{
					UpgradeNodesWorkflowConfig: &types.UpgradeNodesWorkflowConfig{
						InstanceGroupID:      "ig-1",
						NodeUpgradeBatchSize: 1,
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("UpdateWorkflow failed: %v", err)
	}
	if resp.ClusterID != "c-1" || resp.WorkflowID != "wf-1" || resp.RequestID != "req-3" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}
