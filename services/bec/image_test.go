package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// image API
//////////////////////////////////////////////

func TestCreateVmImage(t *testing.T) {
	getReq := &api.CreateVmImageArgs{VmId: "vm-dstkrmda-cn-jinan-cm-235ew", Name: "wcw-test"}
	res, err := CLIENT.CreateVmImage(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateVmImage(t *testing.T) {
	getReq := &api.UpdateVmImageArgs{Name: "wcw-test1"}
	res, err := CLIENT.UpdateVmImage("im-dqoicjmz", getReq)
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
	req := []string{"im-dqoicjmz"}
	res, err := CLIENT.DeleteVmImage(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
