package bec

import (
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// ////////////////////////////////////////////
// route table test
// ////////////////////////////////////////////

func TestUpdateRoute(t *testing.T) {
	req := &api.UpdateRouteTableRequest{
		TableId:   "rtb-g1gbmjfxnyt0",
		TableName: "default-table-zyc-gosdk-3-desc-del-91",
	}
	res, err := CLIENT.UpdateRouteTable("rtb-g1gbmjfxnyt0", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetRouteTableList(t *testing.T) {
	req := &api.ListRequest{}
	res, err := CLIENT.GetRouteTableList(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetRouteTableDetail(t *testing.T) {
	res, err := CLIENT.GetRouteTableDetail("rtb-g1gbmjfxnyt0")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestCreateRouteRule(t *testing.T) {
	req := &api.CreateRouteRuleRequest{
		TableId:            "rtb-g1gbmjfxnyt0",
		IpVersion:          4,
		SourceAddress:      "192.168.2.0/24",
		DestinationAddress: "192.168.31.0/24",
		Nexthop:            "vm-5cvfvvqr-cn-nanning-cm-o5ws3",
		RouteType:          "CUSTOM",
		Description:        "rtb-rule-gosdk",
	}
	res, err := CLIENT.CreateRouteRule(req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestGetRouteRuleList(t *testing.T) {
	req := &api.ListRequest{}
	res, err := CLIENT.GetRouteRuleList("rtb-g1gbmjfxnyt0", req)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}

func TestDeleteRouteRule(t *testing.T) {
	res, err := CLIENT.DeleteRouteRule("rtr-ulsrccd9oje1")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
	jsonRes := TransJsonData(res)
	fmt.Printf("result = %v", jsonRes)
}
