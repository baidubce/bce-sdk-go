package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// vmInstance API
//////////////////////////////////////////////

func TestCreateVmServiceInstance(t *testing.T) {
	// 使用vpc网络
	getReq := &api.CreateVmServiceArgs{
		Cpu: 1, Memory: 2,
		NeedIpv6PublicIp: false, NeedPublicIp: true, Bandwidth: 100,
		DeployInstances: &[]api.DeploymentInstance{
			{
				RegionId:    "cn-hangzhou-cm",
				Replicas:    1,
				NetworkType: "vpc",
			},
		},
		ImageType: api.ImageTypeBec,
		ImageId:   "im-6btnw2x2",
		SystemVolume: &api.SystemVolumeConfig{
			SizeInGB:   40,
			VolumeType: api.DiskTypeNVME,
			Name:       "sys"},
		KeyConfig: &api.KeyConfig{
			Type:             "bccKeyPair",
			BccKeyPairIdList: []string{"k-CIg4d2cC"},
		},
	}
	res, err := CLIENT.CreateVmServiceInstance("s-bu5xjidw", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmServiceInstanceTripleLine(t *testing.T) {
	// 测试三线节点制定运营商
	getReq1 := &api.CreateVmServiceArgs{
		Cpu: 1, Memory: 2,
		NeedIpv6PublicIp: false, NeedPublicIp: true, Bandwidth: 100,
		DeployInstances: &[]api.DeploymentInstance{
			{
				RegionId:            "cn-changchun-ix",
				Replicas:            1,
				NetworkType:         "vpc",
				SubServiceProviders: []string{"cm"},
			},
		},
		ImageType: api.ImageTypeBec,
		ImageId:   "im-6btnw2x2",
		SystemVolume: &api.SystemVolumeConfig{
			SizeInGB:   40,
			VolumeType: api.DiskTypeNVME,
			Name:       "sys"},
		KeyConfig: &api.KeyConfig{
			Type:             "bccKeyPair",
			BccKeyPairIdList: []string{"k-CIg4d2cC"},
		},
	}
	res1, err1 := CLIENT.CreateVmServiceInstance("s-bu5xjidw", getReq1)
	ExpectEqual(t.Errorf, nil, err1)
	t.Logf("%+v", res1)
}

func TestGetVmInstanceList(t *testing.T) {
	getReq := &api.ListRequest{ServiceId: "s-vrowm5qt"}
	res, err := CLIENT.GetVmInstanceList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVirtualMachine(t *testing.T) {
	res, err := CLIENT.GetVirtualMachine("vm-vrowm5qt-cn-baoding-ix-dvaqv")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestOperateVmDeployment(t *testing.T) {
	res, err := CLIENT.OperateVmDeployment("vm-czpgb91c-cn-langfang-ct-wgbem", api.VmInstanceBatchOperateStart)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmInstanceMetrics(t *testing.T) {
	res, err := CLIENT.GetVmInstanceMetrics("vm-2qbftgbf-cn-langfang-ct-f7p9j", "", 1660147200, 1660233600, 0, api.MetricsTypeCpu)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmConfig(t *testing.T) {
	res, err := CLIENT.GetVmConfig("vm-czpgb91c-cn-langfang-ct-wgbem")
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
	req := &api.UpdateVmInstanceArgs{Type: "resource", Cpu: 2, Memory: 4, VmName: "vm-czpgb91c-cn-langfang-ct-wgbem"}
	res, err := CLIENT.UpdateVmInstance("vm-czpgb91c-cn-langfang-ct-wgbem", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBindSecurityGroup(t *testing.T) {
	req := &api.BindSecurityGroupInstances{Instances: []api.InstancesBinding{
		api.InstancesBinding{
			InstanceId:       "vm-czpgb91c-cn-langfang-ct-wgbem",
			SecurityGroupIds: []string{"sg-itlgtacv", "sg-219mosrn"},
		},
	}}
	// action:  bind  or  unbind
	res, err := CLIENT.BindSecurityGroup("bind", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteVmInstance(t *testing.T) {
	res, err := CLIENT.DeleteVmInstance("vm-czpgb91c-cn-langfang-ct-wgbem")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
