package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// image API
//////////////////////////////////////////////

func TestCreateVmImage(t *testing.T) {
	getReq := &api.CreateVmImageArgs{VmId: "vm-xxxx-1", Name: "xxxx-test"}
	res, err := CLIENT.CreateVmImage(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateVmImage(t *testing.T) {
	getReq := &api.UpdateVmImageArgs{Name: "xxxx-test-update"}
	res, err := CLIENT.UpdateVmImage("xxxxxx-i", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestListVmImage(t *testing.T) {
	getReq := &api.ListVmImageArgs{OsName: "CentOS", Status: "IMAGING"}
	res, err := CLIENT.ListVmImage(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteVmImage(t *testing.T) {
	req := []string{"xxxxxx-1", "xxxxxx-2"}
	res, err := CLIENT.DeleteVmImage(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
