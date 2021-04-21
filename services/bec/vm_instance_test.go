package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// vmInstance API
//////////////////////////////////////////////
func TestGetVmInstanceList(t *testing.T) {
	getReq := &api.ListRequest{KeywordType: "instanceId", Keyword: "vm-4xxxxxx"}
	res, err := CLIENT.GetVmInstanceList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetNodeVmInstanceList(t *testing.T) {
	getReq := &api.ListRequest{KeywordType: "instanceId", Keyword: "vm-xxx"}
	res, err := CLIENT.GetNodeVmInstanceList(getReq, "EAST_CHINA", "CHINA_MOBILE", "HANGZHOU")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVirtualMachine(t *testing.T) {
	res, err := CLIENT.GetVirtualMachine("vm-xxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestOperateVmDeployment(t *testing.T) {
	res, err := CLIENT.OperateVmDeployment("vm-4xxxxxxx", api.VmInstanceBatchOperateStart)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmInstanceMetrics(t *testing.T) {
	res, err := CLIENT.GetVmInstanceMetrics("vm-xxx", "", 86400,
		api.MetricsTypeBandwidthReceive)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmConfig(t *testing.T) {
	res, err := CLIENT.GetVmConfig("vm-xxxxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestReinstallVmInstance(t *testing.T) {
	req := &api.ReinstallVmInstanceArg{ImageId: "im-dikfxxxx", AdminPass: "1xxAxxx@", ImageType: api.ImageTypeBec, KeyConfig: &api.KeyConfig{Type: "password", AdminPass: "1xxAxxx@"}}
	res, err := CLIENT.ReinstallVmInstance("vm-pxxx", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateVmDeployment(t *testing.T) {
	req := &api.UpdateVmDeploymentArgs{Type: "resource", Cpu: 2, VmName: "xxxxx-test",
		SystemVolume: &api.SystemVolumeConfig{VolumeType: api.DiskTypeNVME, Name: "sys", PvcName: "lvm-xxxxxx-rootfs"}}
	res, err := CLIENT.UpdateVmDeployment("vm-wxxxxx", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteVmInstance(t *testing.T) {
	res, err := CLIENT.DeleteVmInstance("vm-4xxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
