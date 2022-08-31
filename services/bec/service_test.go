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
	res, err := CLIENT.GetService("s-f9ngbkbc")
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
				VolumeMounts: []api.V1VolumeMount{
					api.V1VolumeMount{
						MountPath: "/temp",
						Name:      "emptydir01",
					},
				}},
		},
		DeployInstances: &[]api.DeploymentInstance{
			api.DeploymentInstance{
				Replicas: 1,
				RegionId: "cn-langfang-ct",
			},
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
	res, err := CLIENT.GetServiceMetrics("s-f9ngbkbc", api.MetricsTypeMemory, "", 1661270400, 1661356800, 0)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateService(t *testing.T) {
	getReq := &api.UpdateServiceArgs{ServiceName: "s-f9ngbkbc", Type: api.UpdateServiceTypeReplicas, DeployInstances: &[]api.DeploymentInstance{
		api.DeploymentInstance{Region: api.RegionEastChina, Replicas: 1, City: "HANGZHOU", ServiceProvider: api.ServiceChinaMobile},
	}}
	res, err := CLIENT.UpdateService("s-f9ngbkbc", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestStopService(t *testing.T) {
	res, err := CLIENT.ServiceAction("s-f9ngbkbc", api.ServiceActionStop)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestStartService(t *testing.T) {
	res, err := CLIENT.ServiceAction("s-f9ngbkbc", api.ServiceActionStart)
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

func TestGetPodDeployment(t *testing.T) {
	res, err := CLIENT.GetPodDeployment("sts-f9ngbkbc-cn-langfang-ct-uxe4z")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestGetPodDeploymentMetrics(t *testing.T) {
	res, err := CLIENT.GetPodDeploymentMetrics("sts-f9ngbkbc-cn-langfang-ct-uxe4z", api.MetricsTypeMemory, "", 1661270400, 1661356800, 0)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestUpdatePodDeploymentReplicas(t *testing.T) {
	getReq := &api.UpdateDeploymentReplicasRequest{
		Replicas: 2,
	}
	res, err := CLIENT.UpdatePodDeploymentReplicas("sts-f9ngbkbc-cn-langfang-ct-uxe4z", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestDeletePodDeployment(t *testing.T) {
	getReq := &[]string{"sts-f9ngbkbc-cn-jinan-un-0cloi"}
	res, err := CLIENT.DeletePodDeployment(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetPodList(t *testing.T) {
	res, err := CLIENT.GetPodList(1, 100,
		"", "", "", "", "")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestGetPodMetrics(t *testing.T) {
	res, err := CLIENT.GetPodMetrics("sts-f9ngbkbc-cn-langfang-ct-uxe4z-0", api.MetricsTypeMemory, "", 1661270400, 1661356800, 0)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestGetPodDetail(t *testing.T) {
	res, err := CLIENT.GetPodDetail("sts-f9ngbkbc-cn-langfang-ct-uxe4z-0")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestRestartPod(t *testing.T) {
	err := CLIENT.RestartPod("sts-f9ngbkbc-cn-langfang-ct-uxe4z-0")
	ExpectEqual(t.Errorf, nil, err)
}
