package bec

import (
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// ////////////////////////////////////////////
// subnet test
// ////////////////////////////////////////////

func TestCreateSubnet(t *testing.T) {
	req := &api.CreateSubnetRequest{
		Name:        "sb-del-zyc-gosdk-dele-91",
		VpcId:       "vpc-07szd7om4iu5",
		Cidr:        "192.168.31.96/27",
		Description: "gogogosb-91",
		Tags: &[]api.Tag{api.Tag{TagKey: "bec-zyc-key",
			TagValue: "bec-zyc-key-val"}},
	}
	res, err := CLIENT.CreateSubnet(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestUpdateSubnet(t *testing.T) {
	req := &api.UpdateSubnetRequest{
		Name:        "sb-zyc-gosdk-3-31",
		Description: "sb-zyc-gosdk-3-desc-31",
	}
	err := CLIENT.UpdateSubnet("sbn-uap7y8mlb7iz", req)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetSubnetList(t *testing.T) {
	req := &api.ListRequest{}
	res, err := CLIENT.GetSubnetList(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetSubnetDetail(t *testing.T) {
	res, err := CLIENT.GetSubnetDetail("sbn-uap7y8mlb7iz")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestDeleteSubnet(t *testing.T) {
	err := CLIENT.DeleteSubnet("sbn-uap7y8mlb7iz")
	ExpectEqual(t.Errorf, nil, err)
}
