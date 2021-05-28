package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// service API
//////////////////////////////////////////////
func TestListService(t *testing.T) {
	res, err := CLIENT.ListService(1, 100,
		"", "", "", "", "")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetService(t *testing.T) {
	res, err := CLIENT.GetService("s-xxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestCreateService(t *testing.T) {
	getReq := &api.CreateServiceArgs{ServiceName: "xxxx-1-test",
		PaymentMethod: "postpay", ContainerGroupName: "cg1",
		Bandwidth:    100,
		NeedPublicIp: false,
		Containers: &[]api.ContainerDetails{
			api.ContainerDetails{
				Name:         "container01",
				Cpu:          1,
				Memory:       2,
				ImageAddress: "hub.baidubce.com/public/mysql",
				ImageVersion: "5.7",
				Commands: []string{"sh",
					"-c",
					"echo OK!&& sleep 3660"},
				VolumeMounts: &[]api.V1VolumeMount{
					api.V1VolumeMount{
						MountPath: "/temp",
						Name:      "emptydir01",
					},
				}},
		},
		DeployInstances: &[]api.DeploymentInstance{
			api.DeploymentInstance{
				Region:          "EAST_CHINA",
				Replicas:        1,
				City:            "SHANGHAI",
				ServiceProvider: "CHINA_TELECOM"},
		},
		Volumes: &api.Volume{
			EmptyDir: &[]api.EmptyDir{
				api.EmptyDir{Name: "emptydir01"},
			},
		},
		Tags: &[]api.Tag{
			api.Tag{
				TagKey:   "a",
				TagValue: "1"},
		},
	}
	res, err := CLIENT.CreateService(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetServiceMetrics(t *testing.T) {
	res, err := CLIENT.GetServiceMetrics("s-xxx", api.MetricsTypeBandwidthReceive, api.ServiceChinaMobile, 87200)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateService(t *testing.T) {
	getReq := &api.UpdateServiceArgs{ServiceName: "xxxx-1-test", Type: api.UpdateServiceTypeReplicas, DeployInstances: &[]api.DeploymentInstance{
		api.DeploymentInstance{Region: api.RegionEastChina, Replicas: 4, City: "HANGZHOU", ServiceProvider: api.ServiceChinaMobile},
	}}
	res, err := CLIENT.UpdateService("xxxx-1", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestStopService(t *testing.T) {
	res, err := CLIENT.ServiceAction("xxxx-1", api.ServiceActionStop)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestStartService(t *testing.T) {
	res, err := CLIENT.ServiceAction("xxxx-1", api.ServiceActionStart)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteService(t *testing.T) {
	res, err := CLIENT.DeleteService("xxxx-1")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestClient_ServiceBatchOperate(t *testing.T) {
	getReq := &api.ServiceBatchOperateArgs{IdList: []string{"xxxx-1", "xxx-2"}, Action: "start"}
	res, err := CLIENT.ServiceBatchOperate(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)

}

func TestClient_ServiceBatchDelete(t *testing.T) {
	getReq := &[]string{"xxxx-1", "xxx-2"}
	res, err := CLIENT.ServiceBatchDelete(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
