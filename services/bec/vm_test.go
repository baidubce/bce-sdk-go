package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// vmService API
//////////////////////////////////////////////
func TestCreateVmServiceOnly(t *testing.T) {
	getReq := &api.CreateVmServiceArgs{ServiceName: "xxxxx"}
	res, err := CLIENT.CreateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmServiceInstance(t *testing.T) {
	getReq := &api.CreateVmServiceArgs{ServiceName: "xxx", ImageId: "im-dikfttnj-3-u-guangzhou", Cpu: 1, Memory: 2,
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{City: "HANGZHOU", Region: "EAST_CHINA",
			ServiceProvider: "CHINA_MOBILE", Replicas: 1}}, AdminPass: "xxxxxx111xxB@", ImageType: api.ImageTypeBec}
	res, err := CLIENT.CreateVmServiceInstance("s-migkfcrh", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmService(t *testing.T) {
	getReq := &api.CreateVmServiceArgs{ServiceName: "xxxxxxx@", ImageId: "im-dikfttnj-3-u-guangzhou", AdminPass: "x123xxx@",
		SystemVolume: &api.SystemVolumeConfig{VolumeType: api.DiskTypeNVME}, Cpu: 1, Memory: 2,
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{City: "HANGZHOU", Region: "EAST_CHINA",
			ServiceProvider: "CHINA_MOBILE", Replicas: 1}}, ImageType: "bec", KeyConfig: &api.KeyConfig{Type: "password", AdminPass: "xxxx123@"}}
	res, err := CLIENT.CreateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateVmService(t *testing.T) {
	getReq := &api.UpdateVmServiceArgs{UpdateBecVmForm: api.UpdateBecVmForm{Type: api.UpdateVmTypeServiceName,
		VmName: "vm-xxxxx"}, ServiceName: "xxxxtest-2", DeployInstances: &[]api.DeploymentInstance{
		api.DeploymentInstance{Region: "SOUTH_CHINA", City: "GUANGZHOU", Replicas: 1, ServiceProvider: api.ServiceChinaUnicom}}}
	res, err := CLIENT.UpdateVmService("s-xxx", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmServiceList(t *testing.T) {
	getReq := &api.ListVmServiceArgs{ServiceId: "s-xxxxx"}
	res, err := CLIENT.GetVmServiceList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmServiceDetail(t *testing.T) {
	res, err := CLIENT.GetVmServiceDetail("s-xxxxxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestVmServiceAction(t *testing.T) {
	res, err := CLIENT.VmServiceAction("s-xxxxxxx", api.VmServiceActionStart)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmServiceMetrics(t *testing.T) {
	res, err := CLIENT.GetVmServiceMetrics("s-xxx", api.MetricsTypeBandwidthReceive, 259200, "")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBatchOperateVmServiceStop(t *testing.T) {
	getReq := &api.VmServiceBatchActionArgs{IdList: []string{"s-xxxxx-1", "s-xxxxx-2"},
		Action: api.VmServiceBatchStart}
	res, err := CLIENT.BatchOperateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBatchOperateVmServiceStart(t *testing.T) {
	getReq := &api.VmServiceBatchActionArgs{IdList: []string{"s-xxxxx"}, Action: api.VmServiceBatchStart}
	res, err := CLIENT.BatchOperateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteVmService(t *testing.T) {
	res, err := CLIENT.DeleteVmService("s-xxxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBatchDeleteVmService(t *testing.T) {
	getReq := &[]string{"s-xxxx-1", "s-xxxx-2"}
	res, err := CLIENT.BatchDeleteVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
