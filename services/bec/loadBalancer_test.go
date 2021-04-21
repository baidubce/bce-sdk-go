package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// Loadbalancer API
//////////////////////////////////////////////
func TestCreateBlb(t *testing.T) {
	getReq := &api.CreateBlbArgs{BlbName: "xxxx-test", Region: api.RegionEastChina,
		City: "HANGZHOU", LbType: "vm", ServiceProvider: api.ServiceChinaMobile}
	res, err := CLIENT.CreateBlb(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
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
}

func TestGetBlbDetail(t *testing.T) {
	res, err := CLIENT.GetBlbDetail("lb-uf1rv3o5")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateBlb(t *testing.T) {
	getReq := &api.UpdateBlbArgs{BlbName: "xxxxx-test-update", BandwidthInMbpsLimit: 1000}
	res, err := CLIENT.UpdateBlb("xxxxx", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateBlbMonitorPort(t *testing.T) {
	getReq := &api.BlbMonitorArgs{LbMode: api.LbModeWrr, FrontendPort: &api.Port{Protocol: api.ProtocolUdp, Port: 80},
		BackendPort: 80, HealthCheck: &api.HealthCheck{HealthCheckString: "", HealthCheckType: "udp",
			HealthyThreshold: 1000, UnhealthyThreshold: 1000, TimeoutInSeconds: 900, IntervalInSeconds: 3}}
	res, err := CLIENT.CreateBlbMonitorPort("lb-xxx", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteBlbMonitorPort(t *testing.T) {
	req := []api.Port{{Protocol: api.ProtocolUdp, Port: 80}}
	res, err := CLIENT.DeleteBlbMonitorPort("lb-xxx", &req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbMonitorPortList(t *testing.T) {
	res, err := CLIENT.GetBlbMonitorPortList("lb-xxx", 1, 100)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateBlbMonitorPort(t *testing.T) {
	getReq := &api.BlbMonitorArgs{LbMode: api.LbModeWrr, FrontendPort: &api.Port{Protocol: api.ProtocolTcp, Port: 80},
		BackendPort: 8080, HealthCheck: &api.HealthCheck{HealthCheckString: "", HealthCheckType: "tcp",
			HealthyThreshold: 1000, UnhealthyThreshold: 1000, TimeoutInSeconds: 900, IntervalInSeconds: 3}}
	res, err := CLIENT.UpdateBlbMonitorPort("lb-uf1rv3o5", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbMonitorPortDetails(t *testing.T) {
	res, err := CLIENT.GetBlbMonitorPortDetails("lb-xxx", api.ProtocolUdp, 80)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbBackendPodList(t *testing.T) {
	res, err := CLIENT.GetBlbBackendPodList("lb-xxx", 0, 0)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbBackendBindingStsList(t *testing.T) {
	res, err := CLIENT.GetBlbBackendBindingStsList("lb-xxx", 0, 0, "", "")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetBlbBindingPodListWithSts(t *testing.T) {
	res, err := CLIENT.GetBlbBindingPodListWithSts("lb-xxx", "sts-xxxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateBlbBinding(t *testing.T) {
	getReq := &api.CreateBlbBindingArgs{BindingForms: &[]api.BlbBindingForm{
		api.BlbBindingForm{DeploymentId: "xxxx", PodWeight: &[]api.Backends{
			api.Backends{Name: "xxxxxx", Ip: "172.xx.x.xx", Weight: 100}},
		}}}
	res, err := CLIENT.CreateBlbBinding("lb-xxx", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteBlbBindPod(t *testing.T) {
	getReq := &api.DeleteBlbBindPodArgs{PodWeightList: &[]api.Backends{
		api.Backends{Name: "vm-xxx", Ip: "172.16.9xxx.xxx", Weight: 10}},
		DeploymentIds: []string{"vmrs-xxxx"}}
	res, err := CLIENT.DeleteBlbBindPod("lb-xxxx", getReq)
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
	res, err := CLIENT.GetBlbMetrics("lb-xxxxxx", "extranet", "",
		"", 1, api.MetricsTypeBandwidthReceive)
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
	getReq := &api.BatchCreateBlbMonitorArg{Protocol: api.ProtocolUdp, LbMode: api.LbModeWrr, HealthCheck: &api.HealthCheck{HealthCheckString: "", HealthCheckType: "udp",
		HealthyThreshold: 1000, UnhealthyThreshold: 1000, TimeoutInSeconds: 900, IntervalInSeconds: 3}, PortGroups: &[]api.PortGroup{api.PortGroup{
		Port: 80, BackendPort: 80}, api.PortGroup{Port: 443, BackendPort: 443}}}
	res, err := CLIENT.BatchCreateBlbMonitor("lb-xxxx", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
