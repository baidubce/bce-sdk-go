package bec

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
	"testing"
)

// ////////////////////////////////////////////
// vmService API
// ////////////////////////////////////////////
func TestCreateVmServiceOnly(t *testing.T) {
	getReq := &api.CreateVmServiceArgs{ServiceName: "wcw-test"}
	res, err := CLIENT.CreateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmService(t *testing.T) {
	getReq := &api.CreateVmServiceArgs{KeyConfig: &api.KeyConfig{
		Type:             "bccKeyPair",
		BccKeyPairIdList: []string{"k-lVBDKoDj"},
	}, ImageId: "m-f0ZRR9qB", Bandwidth: 100, ImageType: api.ImageTypeBec, SystemVolume: &api.SystemVolumeConfig{SizeInGB: 40, VolumeType: api.DiskTypeNVME, Name: "sys"},
		NetworkConfigList: &[]api.NetworkConfig{api.NetworkConfig{NodeType: "SINGLE", NetworksList: &[]api.Networks{api.Networks{NetType: "INTERNAL_IP", NetName: "eth0"},
			api.Networks{NetType: "PUBLIC_IP", NetName: "eth1"}}}}, DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{RegionId: "cn-langfang-ct", Replicas: 1,
			NetworkType: "classic"}}, DisableIntranet: false, NeedPublicIp: true, NeedIpv6PublicIp: false, SecurityGroupIds: []string{"sg-219mosrn"},
		DnsConfig: &api.DnsConfig{
			DnsType: "DEFAULT",
		}, Cpu: 2, Memory: 4, PaymentMethod: "postpay"}
	res, err := CLIENT.CreateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmServicePostPayTag(t *testing.T) {
	createReq := &api.CreateVmServiceArgs{ServiceName: "zyc-postpay-del-gosdk-userdata-49", VmName: "zyc-postpay-del-gosdk-userdata-49",
		AdminPass: "vWdUKxsdx$xBE8v%*",
		ImageId:   "m-uCGlYurQ", Bandwidth: 100, ImageType: api.ImageTypeBec,
		SystemVolume: &api.SystemVolumeConfig{SizeInGB: 40, VolumeType: api.DiskTypeNVME, Name: "sys"},
		NetworkConfigList: &[]api.NetworkConfig{api.NetworkConfig{NodeType: "SINGLE",
			NetworksList: &[]api.Networks{api.Networks{NetType: "INTERNAL_IP", NetName: "eth0"}}}},
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{
			RegionId:    "cn-nanning-cm",
			Replicas:    1,
			NetworkType: "vpc"}}, DisableIntranet: false, NeedPublicIp: false, NeedIpv6PublicIp: false,
		SecurityGroupIds: []string{"sg-219mosrn"},
		Reservation: &api.Reservation{
			TimeUnit: "month",
		},
		AutoRenew: &api.AutoRenew{
			Length:   4,
			TimeUnit: "month",
		},
		HostnameGenMethod: api.HostnameGenMethodRandom,
		DirectPay:         true,
		ActionType:        api.ActionTypeNew,
		DnsConfig: &api.DnsConfig{
			DnsType: "DEFAULT",
		}, Cpu: 1, Memory: 4, PaymentMethod: "postpay", Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}}, DeploysetIdList: []string{"dset-1j7ewwjb"},
		UserData: "dXNlcl9pbmplY3RlZF9kYXRhOiBJeUV2WW1sdUwzTm9DbVZqYUc4Z0lsZGxiR052YldVZ2RH" +
			"OGdRbUZwWkhVZ1FVa2dRMnh2ZFdRdUlpQjhJSFJsWlNBdmNtOXZkQzkxYzJWeVJHRjBZVVpwYkdVMA=="}
	t.Logf("%+v", createReq)
	req := TransJsonData(createReq)
	fmt.Println("req=" + req)
	res, err := CLIENT.CreateVmService(createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmServicePrepay(t *testing.T) {
	createReq := &api.CreateVmServiceArgs{ServiceName: "zyc-test-del-gosdk-prepay-23", VmName: "zyc-test-del-gosdk-prepay-23",
		AdminPass: "vWdUKxsdx$xBE8v%*",
		ImageId:   "m-uCGlYurQ", Bandwidth: 100, ImageType: api.ImageTypeBec,
		SystemVolume: &api.SystemVolumeConfig{SizeInGB: 40, VolumeType: api.DiskTypeNVME, Name: "sys"},
		NetworkConfigList: &[]api.NetworkConfig{api.NetworkConfig{NodeType: "SINGLE",
			NetworksList: &[]api.Networks{api.Networks{NetType: "INTERNAL_IP", NetName: "eth0"}}}},
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{
			RegionId:    "cn-nanning-cm",
			Replicas:    1,
			NetworkType: "vpc"}}, DisableIntranet: false, NeedPublicIp: false, NeedIpv6PublicIp: false,
		SecurityGroupIds: []string{"sg-219mosrn"},
		Reservation: &api.Reservation{
			Length:   1,
			TimeUnit: "month",
		},
		AutoRenew: &api.AutoRenew{
			Length:   1,
			TimeUnit: "month",
		},
		DirectPay: true,
		DnsConfig: &api.DnsConfig{
			DnsType: "DEFAULT",
		},
		Cpu: 1, Memory: 4,
		PaymentMethod: "prepay",
		Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}},
		UserData: "dXNlcl9pbmplY3RlZF9kYXRhOiBJeUV2WW1sdUwzTm9DbVZqYUc4Z0lsZGxiR052YldVZ2RH" +
			"OGdRbUZwWkhVZ1FVa2dRMnh2ZFdRdUlpQjhJSFJsWlNBdmNtOXZkQzkxYzJWeVJHRjBZVVpwYkdVMA=="}

	t.Logf("%+v", createReq)
	req := TransJsonData(createReq)
	fmt.Println("req=" + req)
	res, err := CLIENT.CreateVmService(createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmServiceCuda(t *testing.T) {
	createReq := &api.CreateVmServiceArgs{ServiceName: "zyc-test-del-gosdk-cuda-51", VmName: "zyc-test-del-gosdk-cuda-51",
		AdminPass: "vWdUKxsdx$xBE8v%*",
		ImageId:   "m-uCGlYurQ", Bandwidth: 100, ImageType: api.ImageTypeBec,
		SystemVolume: &api.SystemVolumeConfig{SizeInGB: 40, VolumeType: api.DiskTypeNVME, Name: "sys"},
		NetworkConfigList: &[]api.NetworkConfig{api.NetworkConfig{NodeType: "SINGLE",
			NetworksList: &[]api.Networks{api.Networks{NetType: "INTERNAL_IP", NetName: "eth0"}}}},
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{RegionId: "cn-baoding-ct", Replicas: 1,
			NetworkType: "vpc"}}, DisableIntranet: false, NeedPublicIp: false, NeedIpv6PublicIp: false,
		SecurityGroupIds: []string{"sg-219mosrn"},
		Reservation: &api.Reservation{
			TimeUnit: "month",
		},
		AutoRenew: &api.AutoRenew{
			Length:   3,
			TimeUnit: "month",
		},
		DnsConfig: &api.DnsConfig{
			DnsType: "DEFAULT",
		},
		Cpu: 14, Memory: 120,
		PaymentMethod: "postpay",
		Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}},
		DeploysetIdList: []string{"dset-1j7ewwjb"},
		Spec:            "bec.gn6.c14m120.1gac90",
		CudaVersion:     "12.5.1",
		CudnnVersion:    "9.6.0",
		DriverVersion:   "550.144.03",
		UserData: "dXNlcl9pbmplY3RlZF9kYXRhOiBJeUV2WW1sdUwzTm9DbVZqYUc4Z0lsZGxiR052YldVZ2RH" +
			"OGdRbUZwWkhVZ1FVa2dRMnh2ZFdRdUlpQjhJSFJsWlNBdmNtOXZkQzkxYzJWeVJHRjBZVVpwYkdVMA=="}
	res, err := CLIENT.CreateVmService(createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateVmServiceCuda2(t *testing.T) {
	createReq := &api.CreateVmServiceArgs{ServiceName: "zyc-test-del-gosdk-cuda-53", VmName: "zyc-test-del-gosdk-cuda-53",
		AdminPass: "vWdUKxsdx$xBE8v%*",
		ImageId:   "m-uCGlYurQ", Bandwidth: 100, ImageType: api.ImageTypeBec,
		SystemVolume: &api.SystemVolumeConfig{SizeInGB: 40, VolumeType: api.DiskTypeNVME, Name: "sys"},
		NetworkConfigList: &[]api.NetworkConfig{api.NetworkConfig{NodeType: "SINGLE",
			NetworksList: &[]api.Networks{api.Networks{NetType: "INTERNAL_IP", NetName: "eth0"}}}},
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{RegionId: "cn-baoding-ct", Replicas: 1,
			NetworkType: "vpc"}}, DisableIntranet: false, NeedPublicIp: false, NeedIpv6PublicIp: false,
		SecurityGroupIds: []string{"sg-219mosrn"},
		Reservation: &api.Reservation{
			TimeUnit: "month",
		},
		AutoRenew: &api.AutoRenew{
			Length:   3,
			TimeUnit: "month",
		},
		DnsConfig: &api.DnsConfig{
			DnsType: "DEFAULT",
		},
		Cpu: 14, Memory: 120,
		PaymentMethod: "postpay",
		Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}},
		DeploysetIdList: []string{"dset-1j7ewwjb"},
		Spec:            "bec.gn6.c14m120.1gac90",
		UserData: "dXNlcl9pbmplY3RlZF9kYXRhOiBJeUV2WW1sdUwzTm9DbVZqYUc4Z0lsZGxiR052YldVZ2RH" +
			"OGdRbUZwWkhVZ1FVa2dRMnh2ZFdRdUlpQjhJSFJsWlNBdmNtOXZkQzkxYzJWeVJHRjBZVVpwYkdVMA=="}
	res, err := CLIENT.CreateVmService(createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateVmService(t *testing.T) {
	getReq := &api.UpdateVmServiceArgs{
		UpdateBecVmForm: api.UpdateBecVmForm{
			Type: api.UpdateVmReplicas,
			KeyConfig: &api.KeyConfig{
				Type: "bccKeyPair", BccKeyPairIdList: []string{"k-lVBDKoDj"},
			},
		}, DeployInstances: &[]api.DeploymentInstance{
			api.DeploymentInstance{
				RegionId:    "cn-jinan-cm",
				Replicas:    1,
				NetworkType: "vpc",
			},
		}, ReplicaTemplate: api.ReplicaTemplate{
			Type:       "template",
			TemplateId: "tmpl-gc4maqay",
		},
	}
	res, err := CLIENT.UpdateVmService("s-dstkrmda", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmServiceList(t *testing.T) {
	getReq := &api.ListVmServiceArgs{}
	res, err := CLIENT.GetVmServiceList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmServiceDetail(t *testing.T) {
	res, err := CLIENT.GetVmServiceDetail("s-m5qrjnvr")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestVmServiceAction(t *testing.T) {
	res, err := CLIENT.VmServiceAction("s-dstkrmda", api.VmServiceActionStart)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetVmServiceMetrics(t *testing.T) {
	res, err := CLIENT.GetVmServiceMetrics("s-mifwgtju", "", 1660147200, 1660233600, 1, api.MetricsTypeCpu)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCrateVmPrivateIp(t *testing.T) {
	// 添加虚机辅助IP
	crateVmPrivateIpReq := &api.CreateVmPrivateIpForm{SecondaryPrivateIpAddressCount: 1}
	res, err := CLIENT.CreateVmPrivateIp("vm-czpgb91c-cn-langfang-ct-wgbem", crateVmPrivateIpReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteVmPrivateIp(t *testing.T) {
	// 删除虚机辅助IP
	deleteVmPrivateIpReq := &api.DeleteVmPrivateIpForm{PrivateIps: []string{"172.18.176.54"}}
	delRes, err := CLIENT.DeleteVmPrivateIp("vm-czpgb91c-cn-langfang-ct-wgbem", deleteVmPrivateIpReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", delRes)
}

func TestBatchOperateVmServiceStop(t *testing.T) {
	getReq := &api.VmServiceBatchActionArgs{IdList: []string{"s-xxxxx-1", "s-xxxxx-2"},
		Action: api.VmServiceBatchStart}
	res, err := CLIENT.BatchOperateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBatchOperateVmServiceStart(t *testing.T) {
	getReq := &api.VmServiceBatchActionArgs{IdList: []string{"s-bu5xjidw"}, Action: api.VmServiceBatchStop}
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

func TransJsonData(v interface{}) string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("serialization fail:", err)
		return ""
	}
	return string(jsonData)
}
