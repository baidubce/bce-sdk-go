package bec

import (
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// ////////////////////////////////////////////
// vpc test
// ////////////////////////////////////////////

func TestCreateVpc(t *testing.T) {
	req := &api.CreateVpcRequest{
		Name:        "vpc-del-zyc-gosdk-fin-93",
		RegionId:    "cn-nanning-cm",
		Cidr:        "192.168.92.0/24",
		Description: "gogogo",
		Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}},
	}
	res, err := CLIENT.CreateVpc(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestUpdateVpc(t *testing.T) {
	req := &api.UpdateVpcRequest{
		Name:        "vpc-zyc-gosdk-4",
		Description: "vpc-zyc-gosdk-3-desc-4",
	}
	res, err := CLIENT.UpdateVpc("vpc-8hqintn2y5ms", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetVpcList(t *testing.T) {
	req := &api.ListRequest{}
	res, err := CLIENT.GetVpcList(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetVpcDetail(t *testing.T) {
	res, err := CLIENT.GetVpcDetail("vpc-8hqintn2y5ms")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestDeleteVpc(t *testing.T) {
	res, err := CLIENT.DeleteVpc("vpc-8hqintn2y5ms")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}
