package bec

import (
	"fmt"
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

func TestCreateVmServiceInstanceTag(t *testing.T) {
	// 使用vpc网络
	createReq := &api.CreateVmServiceArgs{VmName: "zyc-test-del-gosdk-userdata-instance-4",
		KeyConfig: &api.KeyConfig{
			Type:             "bccKeyPair",
			BccKeyPairIdList: []string{"k-rxoIlhIg"},
		}, ImageId: "m-uCGlYurQ", Bandwidth: 100, ImageType: api.ImageTypeBec,
		SystemVolume: &api.SystemVolumeConfig{SizeInGB: 40, VolumeType: api.DiskTypeNVME, Name: "sys"},
		NetworkConfigList: &[]api.NetworkConfig{api.NetworkConfig{NodeType: "SINGLE",
			NetworksList: &[]api.Networks{api.Networks{NetType: "INTERNAL_IP", NetName: "eth0"}}}},
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{RegionId: "cn-nanning-cm", Replicas: 1,
			NetworkType: "vpc"}}, DisableIntranet: false, NeedPublicIp: false, NeedIpv6PublicIp: false,
		SecurityGroupIds: []string{"sg-219mosrn"},
		DnsConfig: &api.DnsConfig{
			DnsType: "DEFAULT",
		}, Cpu: 1, Memory: 4, PaymentMethod: "postpay", Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}}, DeploysetIdList: []string{"dset-1j7ewwjb"},
		UserData: "dXNlcl9pbmplY3RlZF9kYXRhOiBJeUV2WW1sdUwzTm9DbVZqYUc4Z0lsZGxiR052YldVZ2RH" +
			"OGdRbUZwWkhVZ1FVa2dRMnh2ZFdRdUlpQjhJSFJsWlNBdmNtOXZkQzkxYzJWeVJHRjBZVVpwYkdVMA=="}

	res, err := CLIENT.CreateVmServiceInstance("s-m5qrjnvr", createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
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
	getReq := &api.ListRequest{ServiceId: "s-img3b4zz"}
	res, err := CLIENT.GetVmInstanceList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetVmInstanceListVpcSubnet(t *testing.T) {
	getReq := &api.ListRequest{KeywordType: "instanceName", Keyword: "zyc"}
	res, err := CLIENT.GetVmInstanceList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetVmInstanceListDeploySet(t *testing.T) {
	getReq := &api.ListRequest{KeywordType: "instanceName", Keyword: "zyc"}
	res, err := CLIENT.GetVmInstanceList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetVirtualMachineDetail(t *testing.T) {
	res, err := CLIENT.GetVirtualMachine("vm-nmjqmm8l-cn-nanning-cm-zj1wd")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
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

func TestReinstallVmInstanceUsrData(t *testing.T) {
	req := &api.ReinstallVmInstanceArg{ImageId: "m-3Wa7ZIgH", AdminPass: "ikuI5g", ImageType: api.ImageTypeBec,
		UserData: "dXNlcl9pbmplY3RlZF9kYXRhOiBJeUV2WW1sdUwzTm9DbVZqYUc4Z0lsZGxiR052YldVZ2RHOGdRbUZwWkhVZ1FVa2dRMnh2Z" +
			"FdRdUlpQjhJSFJsWlNBdmNtOXZkQzkxYzJWeVJHRjBZVVpwYkdVMA=="}
	res, err := CLIENT.ReinstallVmInstance("vm-baynoc9f-cn-nanning-cm-y8q5y", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestReinstallVmInstanceCuda(t *testing.T) {
	req := &api.ReinstallVmInstanceArg{ImageId: "m-3Wa7ZIgH", AdminPass: "ikuI5g", ImageType: api.ImageTypeBec,
		CudaVersion:   "12.5.1",
		CudnnVersion:  "9.6.0",
		DriverVersion: "550.144.03",
		UserData: "dXNlcl9pbmplY3RlZF9kYXRhOiBJeUV2WW1sdUwzTm9DbVZqYUc4Z0lsZGxiR052YldVZ2RHOGdRbUZwWkhVZ1FVa2dRMnh2Z" +
			"FdRdUlpQjhJSFJsWlNBdmNtOXZkQzkxYzJWeVJHRjBZVVpwYkdVMA=="}
	res, err := CLIENT.ReinstallVmInstance("vm-jdwctnqq-cn-baoding-ct-xtbul", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
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
