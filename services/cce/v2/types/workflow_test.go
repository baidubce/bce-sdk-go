package types

import (
	"encoding/json"
	"testing"
)

func TestUpgradeNodesWorkflowConfigGPUDriverJSON(t *testing.T) {
	cfg := WorkflowConfig{
		UpgradeNodesWorkflowConfig: &UpgradeNodesWorkflowConfig{
			CCEInstanceIDList:    []string{"i-1", "i-2"},
			InstanceGroupID:      "ig-1",
			NodeUpgradeBatchSize: 2,
			IsPreCheck:           true,
			Components: []Component{
				{
					Name: ComponentGPUDriver,
					GPUTargetVersion: &GPUVersion{
						Driver: "575.57.08",
						CUDA:   "12.9.1",
						CuDNN:  "9.8.0",
					},
				},
			},
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal workflow config failed: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal workflow config failed: %v", err)
	}

	upgradeCfg, ok := got["upgradeNodesWorkflowConfig"].(map[string]interface{})
	if !ok {
		t.Fatalf("upgradeNodesWorkflowConfig missing: %s", string(data))
	}
	components, ok := upgradeCfg["components"].([]interface{})
	if !ok || len(components) != 1 {
		t.Fatalf("components missing: %s", string(data))
	}
	component, ok := components[0].(map[string]interface{})
	if !ok {
		t.Fatalf("component is not object: %s", string(data))
	}
	if component["name"] != string(ComponentGPUDriver) {
		t.Fatalf("unexpected component name: %v", component["name"])
	}
	gpuTarget, ok := component["gpuTargetVersion"].(map[string]interface{})
	if !ok {
		t.Fatalf("gpuTargetVersion missing: %s", string(data))
	}
	if gpuTarget["driver"] != "575.57.08" || gpuTarget["cuda"] != "12.9.1" || gpuTarget["cuDNN"] != "9.8.0" {
		t.Fatalf("unexpected gpuTargetVersion: %#v", gpuTarget)
	}
}

func TestUpgradeComponentsGPUDriverJSON(t *testing.T) {
	resp := UpgradeComponents{
		GPUDriver: UpgradeVersionList{
			GPUCurrentVersion: &GPUVersion{
				Driver: "570.124.06",
				CUDA:   "12.4.1",
				CuDNN:  "9.0.0",
			},
			ComponentVersions: []ComponentVersion{
				{
					NeedDrainNode: true,
					GPUTargetVersion: &GPUVersion{
						Driver: "575.57.08",
						CUDA:   "12.9.1",
						CuDNN:  "9.8.0",
					},
				},
			},
		},
	}

	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal upgrade components failed: %v", err)
	}

	var decoded UpgradeComponents
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal upgrade components failed: %v", err)
	}
	if decoded.GPUDriver.GPUCurrentVersion == nil {
		t.Fatalf("gpu current version missing after round trip")
	}
	if decoded.GPUDriver.GPUCurrentVersion.Driver != "570.124.06" {
		t.Fatalf("unexpected current driver: %q", decoded.GPUDriver.GPUCurrentVersion.Driver)
	}
	if len(decoded.GPUDriver.ComponentVersions) != 1 || decoded.GPUDriver.ComponentVersions[0].GPUTargetVersion == nil {
		t.Fatalf("gpu target version missing after round trip")
	}
	if decoded.GPUDriver.ComponentVersions[0].GPUTargetVersion.CuDNN != "9.8.0" {
		t.Fatalf("unexpected target cudnn: %q", decoded.GPUDriver.ComponentVersions[0].GPUTargetVersion.CuDNN)
	}
}

func TestUpgradeKubeletConfigJSON(t *testing.T) {
	readOnlyPort := int32(10255)
	runtimeRequestTimeout := int32(120)

	cfg := WorkflowConfig{
		UpgradeKubeletConfig: &UpgradeKubeletConfig{
			InstanceGroupID:      "ig-1",
			NodeUpgradeBatchSize: 2,
			DeployCustomConfig: KubeletDeployCustomConfig{
				ReadOnlyPort:          &readOnlyPort,
				Runtimerequesttimeout: &runtimeRequestTimeout,
				MemoryManagerPolicy:   MemoryManagerPolicyStatic,
			},
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal workflow config failed: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal workflow config failed: %v", err)
	}

	upgradeKubeletConfig, ok := got["upgradeKubeletConfig"].(map[string]interface{})
	if !ok {
		t.Fatalf("upgradeKubeletConfig missing: %s", string(data))
	}
	deployCustomConfig, ok := upgradeKubeletConfig["deployCustomConfig"].(map[string]interface{})
	if !ok {
		t.Fatalf("deployCustomConfig missing: %s", string(data))
	}
	if deployCustomConfig["runtimerequesttimeout"] != float64(120) {
		t.Fatalf("unexpected runtimerequesttimeout: %#v", deployCustomConfig["runtimerequesttimeout"])
	}
	if deployCustomConfig["memoryManagerPolicy"] != string(MemoryManagerPolicyStatic) {
		t.Fatalf("unexpected memoryManagerPolicy: %#v", deployCustomConfig["memoryManagerPolicy"])
	}
}
