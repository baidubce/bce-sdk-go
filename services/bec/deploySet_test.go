package bec

import (
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// deploy set test
//////////////////////////////////////////////

func TestCreateDeploySet(t *testing.T) {
	getReq := &api.CreateDeploySetArgs{
		Name: "wcw_test",
		Desc: "wcw-test",
	}
	res, err := CLIENT.CreateDeploySet(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateDeploySet(t *testing.T) {
	getReq := &api.CreateDeploySetArgs{
		Name: "wcw_test",
		Desc: "wcw-test-gosdk",
	}
	err := CLIENT.UpdateDeploySet("dset-y4tumnel", getReq)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetDeploySetList(t *testing.T) {
	getReq := &api.ListRequest{}
	res, err := CLIENT.GetDeploySetList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetDeploySetDetail(t *testing.T) {
	res, err := CLIENT.GetDeploySetDetail("dset-1j7eswxx")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestUpdateVmInstanceDeploySet(t *testing.T) {
	getReq := &api.UpdateVmDeploySetArgs{
		InstanceId:      "vm-dstkrmda-cn-langfang-ct-4thbz",
		DeploysetIdList: []string{"dset-y4tumnel"},
	}
	err := CLIENT.UpdateVmInstanceDeploySet(getReq)
	ExpectEqual(t.Errorf, nil, err)

}

func TestDeleteVmInstanceFromDeploySet(t *testing.T) {
	getReq := &api.DeleteVmDeploySetArgs{
		DeploysetId:    "dset-y4tumnel",
		InstanceIdList: []string{"vm-dstkrmda-cn-langfang-ct-4thbz"},
	}
	err := CLIENT.DeleteVmInstanceFromDeploySet(getReq)
	ExpectEqual(t.Errorf, nil, err)

}

func TestDeleteDeploySet(t *testing.T) {
	err := CLIENT.DeleteDeploySet("dset-y4tumnel")
	ExpectEqual(t.Errorf, nil, err)

}
