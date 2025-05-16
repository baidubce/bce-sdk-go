package bec

import (
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// ////////////////////////////////////////////
// Loadbalancer API
// ////////////////////////////////////////////
func TestCreateBlb(t *testing.T) {
	getReq := &api.CreateBlbArgs{
		BlbName:     "gosdk-test",
		RegionId:    "cn-hangzhou-cm",
		LbType:      "vm",
		NetworkType: "vpc",
	}
	res, err := CLIENT.CreateBlb(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateBlbVpc(t *testing.T) {
	getReq := &api.CreateBlbArgs{
		BlbName:             "gosdk-test1",
		RegionId:            "cn-baoding-ix",
		LbType:              "vm",
		NetworkType:         "vpc",
		SubServiceProviders: []string{"cm"},
	}
	res, err := CLIENT.CreateBlb(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateBlbVpcTripleTag(t *testing.T) {
	cReq := &api.CreateBlbArgs{
		BlbName:             "gosdk-test1-zyc-del",
		RegionId:            "cn-huhehaote-ix",
		LbType:              "vm",
		NetworkType:         "vpc",
		SubnetId:            "sbn-6s0hyohf3gdl",
		VpcId:               "vpc-jlhljrppmtqn",
		SubServiceProviders: []string{"cm"},
		NeedPublicIp:        true,
		Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}},
	}
	res, err := CLIENT.CreateBlb(cReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestDeleteBlb(t *testing.T) {
	res, err := CLIENT.DeleteBlb("xxxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbList(t *testing.T) {
	res, err := CLIENT.GetBlbList("vm", "desc", "createTime", "", "",
		"", "", 1, 100)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetBlbDetail(t *testing.T) {
	res, err := CLIENT.GetBlbDetail("applb-cn-huhehaote-ix-udmtylek")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestUpdateBlb(t *testing.T) {
	getReq := &api.UpdateBlbArgs{BlbName: "applb-cn-hangzhou-cm-wkfdcbin", BandwidthInMbpsLimit: 1000, Type: "bandwidth"}
	res, err := CLIENT.UpdateBlb("applb-cn-hangzhou-cm-wkfdcbin", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateBlbMonitorPort(t *testing.T) {

	str := ""
	t.Logf("%v", &str)
	getReq := &api.BlbMonitorArgs{LbMode: api.LbModeWrr, FrontendPort: &api.Port{Protocol: api.ProtocolTcp, Port: 80},
		BackendPort: 80, HealthCheck: &api.HealthCheck{HealthCheckString: &str, HealthCheckType: "tcp",
			HealthyThreshold: 1000, UnhealthyThreshold: 1000, TimeoutInSeconds: 50, IntervalInSeconds: 3}}
	res, err := CLIENT.CreateBlbMonitorPort("applb-cn-hangzhou-cm-wkfdcbin", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteBlbMonitorPort(t *testing.T) {
	req := []api.Port{{Protocol: api.ProtocolUdp, Port: 81}}
	res, err := CLIENT.DeleteBlbMonitorPort("applb-cn-hangzhou-cm-3vkqupsp", &req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbMonitorPortList(t *testing.T) {
	res, err := CLIENT.GetBlbMonitorPortList("applb-cn-hangzhou-cm-wkfdcbin", 1, 100)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateBlbMonitorPort(t *testing.T) {
	str := ""
	getReq := &api.BlbMonitorArgs{LbMode: api.LbModeWrr, FrontendPort: &api.Port{Protocol: api.ProtocolTcp, Port: 80},
		BackendPort: 8090, HealthCheck: &api.HealthCheck{HealthCheckString: &str, HealthCheckType: "tcp",
			HealthyThreshold: 1000, UnhealthyThreshold: 1000, TimeoutInSeconds: 60, IntervalInSeconds: 3}}
	res, err := CLIENT.UpdateBlbMonitorPort("applb-cn-hangzhou-cm-wkfdcbin", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbMonitorPortDetails(t *testing.T) {
	res, err := CLIENT.GetBlbMonitorPortDetails("applb-cn-hangzhou-cm-3vkqupsp", api.ProtocolTcp, 80)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbBackendPodList(t *testing.T) {
	res, err := CLIENT.GetBlbBackendPodList("applb-cn-langfang-ct-whak2ian", 0, 0)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbBackendBindingStsList(t *testing.T) {
	res, err := CLIENT.GetBlbBackendBindingStsList("applb-cn-langfang-ct-whak2ian", 0, 0, "", "")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbBindingPodListWithSts(t *testing.T) {
	res, err := CLIENT.GetBlbBindingPodListWithSts("lb-pclalzin", "sts-loy3f9tk-2-t-langfang")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateBlbBinding(t *testing.T) {
	getReq := &api.CreateBlbBindingArgs{BindingForms: &[]api.BlbBindingForm{
		api.BlbBindingForm{DeploymentId: "vm-xqjitfy1-cn-langfang-ct", PodWeight: &[]api.Backends{
			api.Backends{Name: "vm-xqjitfy1-cn-langfang-ct-d1f7x", Ip: "172.16.8.22", Weight: 100}},
		}}}
	res, err := CLIENT.CreateBlbBinding("lb-jgvg4pbz", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteBlbBindPod(t *testing.T) {
	getReq := &api.DeleteBlbBindPodArgs{PodWeightList: &[]api.Backends{
		api.Backends{Name: "vm-xqjitfy1-cn-langfang-ct-d1f7x", Ip: "172.16.8.22", Weight: 100}},
		DeploymentIds: []string{"vvm-xqjitfy1-cn-langfang-ct"}}
	res, err := CLIENT.DeleteBlbBindPod("lb-jgvg4pbz", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateBlbBindPodWeight(t *testing.T) {
	getReq := &api.UpdateBindPodWeightArgs{PodWeightList: &[]api.Backends{
		api.Backends{Name: "vm-xxxxx", Ip: "172.16.xx.xx", Weight: 10}},
		DeploymentIds: []string{"vmrs-xxxxxx"}}
	res, err := CLIENT.UpdateBlbBindPodWeight("lb-xxxxx", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbMetrics(t *testing.T) {
	res, err := CLIENT.GetBlbMetrics("applb-cn-langfang-ct-lpcfbjv6", "extranet", "", "",
		1657613108, 1660291508, 1, api.MetricsTypeTrafficReceive)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBatchCreateBlb(t *testing.T) {
	getReq := &api.BatchCreateBlbArgs{BlbName: "xxxx-test", DeployInstances: &[]api.DeploymentInstance{
		api.DeploymentInstance{Region: api.RegionSouthChina, Replicas: 1, City: "GUANGZHOU", ServiceProvider: api.ServiceChinaUnicom},
	}, LbType: "vm"}
	res, err := CLIENT.BatchCreateBlb(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBatchDeleteBlb(t *testing.T) {
	res, err := CLIENT.BatchDeleteBlb([]string{"lb-xxxx", "lb-xxxx-2", "lb-xxxx-3"})
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestBatchCreateBlbMonitor(t *testing.T) {
	str := "\\00\\01\\01\\00\\01\\00\\00\\00\\00\\00\\00\\05baidu\\03com\\00\\00\\01\\00\\01"
	getReq := &api.BatchCreateBlbMonitorArg{Protocol: api.ProtocolUdp, LbMode: api.LbModeWrr, HealthCheck: &api.HealthCheck{HealthCheckString: &str, HealthCheckType: "udp",
		HealthyThreshold: 1000, UnhealthyThreshold: 1000, TimeoutInSeconds: 60, IntervalInSeconds: 3}, PortGroups: &[]api.PortGroup{api.PortGroup{
		Port: 82, BackendPort: 82}, api.PortGroup{Port: 443, BackendPort: 443}}}
	res, err := CLIENT.BatchCreateBlbMonitor("applb-cn-hangzhou-cm-3vkqupsp", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
