package bec

import (
	"fmt"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// ////////////////////////////////////////////
// vmService API
// ////////////////////////////////////////////
// 定义参数 城市、运营商中文名、镜像名称、服务名称
func TestVmServiceSce(t *testing.T) {
	cityName := "杭州"
	serviceProviderName := "移动"
	image := "fyy-test"
	name := "bec-tst"
	//创建无实例虚机服务，并获取服务ID
	getReq := &api.CreateVmServiceArgs{ServiceName: name}
	resVm, err := CLIENT.CreateVmService(getReq)
	serviceId := resVm.Details.ServiceId
	// 根据城市与运营商中文名称获取城市与运营商编号还有REGION编号
	var city string
	var provider api.ServiceProvider
	var region api.Region
	resNode, err := CLIENT.GetBecAvailableNodeInfoVo("vm")
	if err == nil {
		for _, v := range resNode.RegionList {
			for _, cityInfo := range v.CityList {
				if cityInfo.Name == cityName {
					for _, providerInfo := range cityInfo.ServiceProviderList {
						if providerInfo.Name == serviceProviderName {
							city = cityInfo.City
							provider = providerInfo.ServiceProvider
							region = v.Region
						}
					}
				}
			}
		}
	}

	// 根据镜像名称获取镜像ID
	imageReq := &api.ListVmImageArgs{KeywordType: "name", Keyword: image}
	res, err := CLIENT.ListVmImage(imageReq)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", res)
		return
	}
	if len(res.Result) == 0 {
		err = fmt.Errorf("no such image name %s", image)
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", "no such image name")
		return
	}
	imageId := res.Result[0].ImageId

	// 创建虚机，并获取虚机ID
	createVmiReq := &api.CreateVmServiceArgs{ImageId: imageId, Cpu: 1, Memory: 2, NeedPublicIp: true, Bandwidth: 50,
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{City: city, Region: region,
			ServiceProvider: provider, Replicas: 1}}, AdminPass: "xxxxxx111xxB@", ImageType: api.ImageTypeBec}
	createVmiRes, err := CLIENT.CreateVmServiceInstance(serviceId, createVmiReq)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", res)
		return
	}

	vmiId := createVmiRes.Details.Instances[0].VmId

	// 每隔10秒获取一次虚机创建状态，并打印虚机状态，直到虚机进入运行中
	var stateRes *api.VmInstanceDetailsVo
	ticker := time.NewTicker(10 * time.Second)
	errTime := 3
	in := true
	for in {
		select {
		case <-ticker.C:
			//stateRes, err := CLIENT.GetVmServiceDetail(vmId)
			stateRes, err = CLIENT.GetVirtualMachine(vmiId)
			if err != nil && errTime > 0 {
				errTime--
				continue
			}
			if errTime == 0 {
				in = false
				break
			}
			fmt.Println("vm status is ", stateRes.Status)
			if stateRes != nil && stateRes.Status == api.ResourceStatusRunning {
				in = false
				break
			}
		}
	}

	// 打印指定虚机的内外网IP、虚机名称、ID
	vmStatusRes, err := CLIENT.GetVirtualMachine(vmiId)
	fmt.Printf("ip:%s, internalIp:%s, vmName:%s, vmId: %s\n", vmStatusRes.PublicIp,
		vmStatusRes.InternalIp, vmStatusRes.VmName, vmStatusRes.VmId)

	// 添加虚机辅助IP
	crateVmPrivateIpReq := &api.CreateVmPrivateIpForm{SecondaryPrivateIpAddressCount: 1}
	_, err = CLIENT.CreateVmPrivateIp(vmiId, crateVmPrivateIpReq)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		return
	}

	// 删除虚机辅助IP
	GetVmRes, err := CLIENT.GetVirtualMachine(vmiId)
	privateIp := GetVmRes.PrivateIps
	deleteVmPrivateIpReq := &api.DeleteVmPrivateIpForm{PrivateIps: privateIp}
	delRes, err := CLIENT.DeleteVmPrivateIp(vmiId, deleteVmPrivateIpReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", delRes)

	// 重置虚机密码
	updateVmReq := &api.UpdateVmInstanceArgs{Type: "password", KeyConfig: &api.KeyConfig{Type: "password", AdminPass: "12345asdf@"}}
	updateVmRes, err := CLIENT.UpdateVmInstance(vmiId, updateVmReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", updateVmRes)

	// 每隔10秒获取一次虚机创建状态，并打印虚机状态，直到虚机进入运行中
	ticker = time.NewTicker(10 * time.Second)
	errTime = 3
	in = true
	for in {
		select {
		case <-ticker.C:
			//stateRes, err := CLIENT.GetVmServiceDetail(vmId)
			stateRes, err = CLIENT.GetVirtualMachine(vmiId)
			if err != nil && errTime > 0 {
				errTime--
				continue
			}
			if errTime == 0 {
				in = false
				break
			}
			fmt.Println("vm status is ", stateRes.Status)
			if stateRes != nil && stateRes.Status == api.ResourceStatusRunning {
				in = false
				break
			}
		}
	}

	// 关机虚机，每隔10秒获取一次虚机创建状态，并打印虚机状态，直到虚机关机
	operateVmRes, err := CLIENT.OperateVmDeployment(vmiId, api.VmInstanceBatchOperateStop)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", operateVmRes)
		return
	}

	var vmStateRes *api.VmInstanceDetailsVo
	ticker = time.NewTicker(10 * time.Second)
	errTime = 3
	in = true
	for in {
		select {
		case <-ticker.C:
			//stateRes, err := CLIENT.GetVmServiceDetail(vmId)
			vmStateRes, err = CLIENT.GetVirtualMachine(vmiId)
			if err != nil && errTime > 0 {
				errTime--
				continue
			}
			if errTime == 0 {
				in = false
				break
			}
			fmt.Println("vm status is ", vmStateRes.Status)
			if vmStateRes != nil && vmStateRes.Status == api.ResourceStatusStopped {
				in = false
				break
			}
		}
	}
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", vmStateRes)
	// 删除虚机
	deleteVmRes, err := CLIENT.DeleteVmInstance(vmiId)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", deleteVmRes)
	// 删除服务
	deleteVmServiceRes, err := CLIENT.DeleteVmService(serviceId)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", deleteVmServiceRes)
}

// 创建LocalDNS虚机
func TestCreateVmServiceWithLocalDns(t *testing.T) {
	getReq := &api.CreateVmServiceArgs{ServiceName: "xxxxxxx@-local", ImageId: "im-dikfttnj-3-u-guangzhou", AdminPass: "x123xxx@",
		SystemVolume: &api.SystemVolumeConfig{VolumeType: api.DiskTypeNVME}, Cpu: 1, Memory: 2,
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{City: "HANGZHOU", Region: api.RegionEastChina,
			ServiceProvider: api.ServiceChinaMobile, Replicas: 1}}, ImageType: api.ImageTypeBec, KeyConfig: &api.KeyConfig{Type: "password", AdminPass: "xxxx123@"},
		DnsConfig: &api.DnsConfig{DnsType: "LOCAL"}}
	res, err := CLIENT.CreateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

// 创建密钥对虚机
func TestCreateVmServiceWithKeypair(t *testing.T) {
	getReq := &api.CreateVmServiceArgs{ServiceName: "xxxxxxx@-key", ImageId: "im-dikfttnj-3-u-guangzhou",
		SystemVolume: &api.SystemVolumeConfig{VolumeType: api.DiskTypeNVME}, Cpu: 1, Memory: 2,
		DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{City: "HANGZHOU", Region: api.RegionEastChina,
			ServiceProvider: api.ServiceChinaMobile, Replicas: 1}}, ImageType: api.ImageTypeBec, KeyConfig: &api.KeyConfig{Type: "bccKeyPair", BccKeyPairIdList: []string{"k-r4FmM6flink"}}}
	res, err := CLIENT.CreateVmService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBlbScenario(t *testing.T) {
	// 创建负载均衡器
	// 每隔10秒获取一次负载均衡器状态，并打印负载均衡器状态与IP，直到创建完成
	getReq := &api.CreateBlbArgs{BlbName: "xxxx-test", Region: api.RegionEastChina,
		City: "HANGZHOU", LbType: "vm", ServiceProvider: api.ServiceChinaMobile}
	res, err := CLIENT.CreateBlb(getReq)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", res)
		return
	}

	lbId := res.Details.BlbId
	var stateRes *api.BlbInstanceVo
	ticker := time.NewTicker(10 * time.Second)
	errTime := 3
	in := true
	for in {
		select {
		case <-ticker.C:
			//stateRes, err := CLIENT.GetVmServiceDetail(vmId)
			stateRes, err = CLIENT.GetBlbDetail(lbId)
			if err != nil && errTime > 0 {
				errTime--
				continue
			}
			if errTime == 0 {
				in = false
				break
			}
			fmt.Printf("lb status is %s, ip is %s", stateRes.Status, stateRes.PublicIp)
			if stateRes != nil && stateRes.Status == api.ResourceStatusRunning {
				in = false
				break
			}
		}
	}

	// 创建TCP监听
	str := ""
	blbMonitorReq := &api.BlbMonitorArgs{LbMode: api.LbModeWrr, FrontendPort: &api.Port{Protocol: api.ProtocolTcp, Port: 80},
		BackendPort: 80, HealthCheck: &api.HealthCheck{HealthCheckString: &str, HealthCheckType: "tcp",
			HealthyThreshold: 1000, UnhealthyThreshold: 1000, TimeoutInSeconds: 60, IntervalInSeconds: 3}}
	blbMonitorRes, err := CLIENT.CreateBlbMonitorPort(lbId, blbMonitorReq)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", blbMonitorRes)
		return
	}

	vmiId := "vm-brcvrwvt-1-m-hangzhou-ehkrm"
	// 按虚机添加后端服务器
	createBlbBindingReq := &api.CreateBlbBindingArgs{BindingForms: &[]api.BlbBindingForm{
		api.BlbBindingForm{DeploymentId: vmiId, PodWeight: &[]api.Backends{
			api.Backends{Name: vmiId, Weight: 100}},
		}}}
	createBlbBindingRes, err := CLIENT.CreateBlbBinding(lbId, createBlbBindingReq)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", createBlbBindingRes)
		return
	}

	// 查询负载均衡器状态，打印负载均衡器后端服务器
	getBlbBackendPodRes, err := CLIENT.GetBlbBackendPodList(lbId, 0, 0)
	if len(getBlbBackendPodRes.Result) > 0 {
		for _, v := range getBlbBackendPodRes.Result {
			fmt.Println("backend rs is ", v.PodName)
		}
	}

	// 从后端服务器删除虚机
	deleteBlbBindPodReq := &api.DeleteBlbBindPodArgs{PodWeightList: &[]api.Backends{
		api.Backends{Name: vmiId}}}
	deleteBlbBindPodRes, err := CLIENT.DeleteBlbBindPod(lbId, deleteBlbBindPodReq)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", deleteBlbBindPodRes)
	}

	// 删除负载均衡器
	deleteRes, err := CLIENT.DeleteBlb(lbId)
	if err != nil {
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", deleteRes)
	}
}
