package vpc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/eip"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	VPC_CLIENT *Client
	EIP_CLIENT *eip.Client

	region string

	VPCID        string
	SubnetID     string
	RouteTableID string
	RouteRuleID  string
	AclRuleID    string
	NatID        string
	PeerConnID   string
	LocalIfID    string
	PeerVPCID    string
	EIPAddress   string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK          string `json:"AK"`
	SK          string `json:"SK"`
	VPCEndpoint string `json:"VPC"`
	EIPEndpoint string `json:"EIP"`
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	// Get the directory of GOPATH, the config file should be under the directory.
	//for i := 0; i < 7; i++ {
	//	f = filepath.Dir(f)
	//}
	f = filepath.Dir(f)
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	VPC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.VPCEndpoint)
	EIP_CLIENT, _ = eip.NewClient(confObj.AK, confObj.SK, confObj.EIPEndpoint)
	log.SetLogLevel(log.WARN)

	region = confObj.VPCEndpoint[4:6]
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestCreateVPCDhcp(t *testing.T) {
	VPCID = ""
	args := &CreateVpcDhcpArgs{
		DomainNameServers: "3.3.3.5",
		ClientToken:       getClientToken(),
	}
	err := VPC_CLIENT.CreateVPCDhcp(VPCID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestUpdateVPCDhcp(t *testing.T) {
	VPCID = ""
	args := &UpdateVpcDhcpArgs{
		DomainNameServers: "",
		ClientToken:       getClientToken(),
	}
	err := VPC_CLIENT.UpdateVPCDhcp(VPCID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetVPCDhcp(t *testing.T) {
	VPCID = ""
	result, err := VPC_CLIENT.GetVPCDhcpInfo(VPCID)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestCreateVPC(t *testing.T) {
	args := &CreateVPCArgs{
		Name:        "TestSDK-VPC",
		Description: "vpc test",
		Cidr:        "192.168.0.0/16",
		Tags: []model.TagModel{
			{
				TagKey:   "tagK",
				TagValue: "tagV",
			},
		},
		ClientToken: getClientToken(),
	}
	result, err := VPC_CLIENT.CreateVPC(args)
	ExpectEqual(t.Errorf, nil, err)

	VPCID = result.VPCID
}

func TestListVPC(t *testing.T) {
	args := &ListVPCArgs{
		MaxKeys:   1000,
		IsDefault: "false",
	}
	res, err := VPC_CLIENT.ListVPC(args)
	fmt.Println(res.VPCs[0].CreatedTime)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(res)
	fmt.Println(string(r))
}

func TestGetVPCDetail(t *testing.T) {
	result, err := VPC_CLIENT.GetVPCDetail(VPCID)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
	ExpectEqual(t.Errorf, "TestSDK-VPC", result.VPC.Name)
	ExpectEqual(t.Errorf, "vpc test", result.VPC.Description)
	ExpectEqual(t.Errorf, "192.168.0.0/16", result.VPC.Cidr)
	ExpectEqual(t.Errorf, 1, len(result.VPC.Tags))
	ExpectEqual(t.Errorf, "tagK", result.VPC.Tags[0].TagKey)
	ExpectEqual(t.Errorf, "tagV", result.VPC.Tags[0].TagValue)
}

func TestOpenRelay(t *testing.T) {
	VPCID = ""
	args := &UpdateVpcRelayArgs{
		VpcId: VPCID,
	}
	err := VPC_CLIENT.OpenRelay(args)
	ExpectEqual(t.Errorf, nil, err)

}

func TestShutdownRelay(t *testing.T) {
	VPCID = ""
	args := &UpdateVpcRelayArgs{
		VpcId: VPCID,
	}
	err := VPC_CLIENT.ShutdownRelay(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestUpdateVPC(t *testing.T) {
	args := &UpdateVPCArgs{
		Name:        "TestSDK-VPC-update",
		Description: "vpc update",
		ClientToken: getClientToken(),
	}
	err := VPC_CLIENT.UpdateVPC(VPCID, args)
	ExpectEqual(t.Errorf, nil, err)

	result, err := VPC_CLIENT.GetVPCDetail(VPCID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "TestSDK-VPC-update", result.VPC.Name)
	ExpectEqual(t.Errorf, "vpc update", result.VPC.Description)
}

func TestGetPrivateIpAddressInfo(t *testing.T) {
	args := &GetVpcPrivateIpArgs{
		VpcId:              "vpc-2pa2x0bjt26i",
		PrivateIpAddresses: []string{"192.168.0.1,192.168.0.2"},
		PrivateIpRange:     "192.168.0.0-192.168.0.45",
	}
	result, err := VPC_CLIENT.GetPrivateIpAddressesInfo(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}
func TestCreateSubnet(t *testing.T) {
	args := &CreateSubnetArgs{
		Name:        "TestSDK-Subnet",
		ZoneName:    fmt.Sprintf("cn-%s-a", region),
		Cidr:        "192.168.1.0/24",
		VpcId:       VPCID,
		SubnetType:  SUBNET_TYPE_BCC,
		Description: "test subnet",
		EnableIpv6:  true,
		Tags: []model.TagModel{
			{
				TagKey:   "tagK",
				TagValue: "tagV",
			},
		},
		ClientToken: getClientToken(),
	}
	result, err := VPC_CLIENT.CreateSubnet(args)
	ExpectEqual(t.Errorf, nil, err)

	SubnetID = result.SubnetId
}

func TestListSubnets(t *testing.T) {
	args := &ListSubnetArgs{
		VpcId:      VPCID,
		SubnetType: SUBNET_TYPE_BCC,
	}
	_, err := VPC_CLIENT.ListSubnets(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetSubnetDetail(t *testing.T) {
	result, err := VPC_CLIENT.GetSubnetDetail(SubnetID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "TestSDK-Subnet", result.Subnet.Name)
	ExpectEqual(t.Errorf, fmt.Sprintf("cn-%s-a", region), result.Subnet.ZoneName)
	ExpectEqual(t.Errorf, "192.168.1.0/24", result.Subnet.Cidr)
	ExpectEqual(t.Errorf, VPCID, result.Subnet.VPCId)
	ExpectEqual(t.Errorf, SUBNET_TYPE_BCC, result.Subnet.SubnetType)
	ExpectEqual(t.Errorf, "test subnet", result.Subnet.Description)
	ExpectEqual(t.Errorf, 1, len(result.Subnet.Tags))
	ExpectEqual(t.Errorf, "tagK", result.Subnet.Tags[0].TagKey)
	ExpectEqual(t.Errorf, "tagV", result.Subnet.Tags[0].TagValue)
}

func TestUpdateSubnet(t *testing.T) {
	args := &UpdateSubnetArgs{
		ClientToken: getClientToken(),
		Name:        "TestSDK-Subnet-update",
		Description: "subnet update",
		EnableIpv6:  true,
	}
	err := VPC_CLIENT.UpdateSubnet(SubnetID, args)
	ExpectEqual(t.Errorf, nil, err)

	result, err := VPC_CLIENT.GetSubnetDetail(SubnetID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "TestSDK-Subnet-update", result.Subnet.Name)
	ExpectEqual(t.Errorf, "subnet update", result.Subnet.Description)
}

func TestListAclEntrys(t *testing.T) {
	_, err := VPC_CLIENT.ListAclEntrys(VPCID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestCreateAclRule(t *testing.T) {
	args := &CreateAclRuleArgs{
		ClientToken: getClientToken(),
		AclRules: []AclRuleRequest{
			{
				SubnetId:             SubnetID,
				Description:          "test acl rule",
				Protocol:             ACL_RULE_PROTOCOL_TCP,
				SourceIpAddress:      "192.168.1.1",
				DestinationIpAddress: "172.17.1.1",
				SourcePort:           "5555",
				DestinationPort:      "6666",
				Position:             10,
				Direction:            ACL_RULE_DIRECTION_EGRESS,
				Action:               ACL_RULE_ACTION_ALLOW,
			},
		},
	}
	err := VPC_CLIENT.CreateAclRule(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestListAclRules(t *testing.T) {
	args := &ListAclRulesArgs{
		SubnetId: SubnetID,
	}
	result, err := VPC_CLIENT.ListAclRules(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 3, len(result.AclRules))
	for _, acl := range result.AclRules {
		if acl.Position == 10 {
			ExpectEqual(t.Errorf, SubnetID, acl.SubnetId)
			ExpectEqual(t.Errorf, "test acl rule", acl.Description)
			ExpectEqual(t.Errorf, ACL_RULE_PROTOCOL_TCP, acl.Protocol)
			ExpectEqual(t.Errorf, "192.168.1.1/32", acl.SourceIpAddress)
			ExpectEqual(t.Errorf, "172.17.1.1/32", acl.DestinationIpAddress)
			ExpectEqual(t.Errorf, "5555", acl.SourcePort)
			ExpectEqual(t.Errorf, "6666", acl.DestinationPort)
			ExpectEqual(t.Errorf, ACL_RULE_DIRECTION_EGRESS, acl.Direction)
			ExpectEqual(t.Errorf, ACL_RULE_ACTION_ALLOW, acl.Action)
			AclRuleID = acl.Id
			break
		}
	}
	if AclRuleID == "" {
		t.Errorf("Test acl rule failed.")
	}
}

func TestUpdateAclRule(t *testing.T) {
	args := &UpdateAclRuleArgs{
		ClientToken:          getClientToken(),
		Description:          "acl rule update",
		Protocol:             ACL_RULE_PROTOCOL_UDP,
		SourceIpAddress:      "192.168.1.2",
		DestinationIpAddress: "172.17.1.2",
		SourcePort:           "7777",
		DestinationPort:      "8888",
		Position:             10,
		Action:               ACL_RULE_ACTION_DENY,
	}
	err := VPC_CLIENT.UpdateAclRule(AclRuleID, args)
	ExpectEqual(t.Errorf, nil, err)

	listAclArgs := &ListAclRulesArgs{SubnetId: SubnetID}
	result, err := VPC_CLIENT.ListAclRules(listAclArgs)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 3, len(result.AclRules))
	isExist := false
	for _, acl := range result.AclRules {
		if acl.Position == 10 {
			isExist = true
			ExpectEqual(t.Errorf, SubnetID, acl.SubnetId)
			ExpectEqual(t.Errorf, "acl rule update", acl.Description)
			ExpectEqual(t.Errorf, ACL_RULE_PROTOCOL_UDP, acl.Protocol)
			ExpectEqual(t.Errorf, "192.168.1.2/32", acl.SourceIpAddress)
			ExpectEqual(t.Errorf, "172.17.1.2/32", acl.DestinationIpAddress)
			ExpectEqual(t.Errorf, "7777", acl.SourcePort)
			ExpectEqual(t.Errorf, "8888", acl.DestinationPort)
			ExpectEqual(t.Errorf, ACL_RULE_ACTION_DENY, acl.Action)
			break
		}
	}
	if !isExist {
		t.Errorf("Test acl rule failed.")
	}
}

func TestCreateDefaultNatGateway(t *testing.T) {
	args := &CreateNatGatewayArgs{
		ClientToken: getClientToken(),
		Name:        "Test-SDK-NatGateway",
		VpcId:       VPCID,
		Spec:        NAT_GATEWAY_SPEC_SMALL,
		Billing: &Billing{
			PaymentTiming: PAYMENT_TIMING_POSTPAID,
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",
				TagValue: "tagValue",
			},
		},
	}
	result, err := VPC_CLIENT.CreateNatGateway(args)
	ExpectEqual(t.Errorf, nil, err)
	NatID = result.NatId

	err = waitStateForNatGateway(NatID, NAT_STATUS_UNCONFIGURED)
	ExpectEqual(t.Errorf, nil, err)
}

func TestCreateEnhanceNatGateway(t *testing.T) {
	args := &CreateNatGatewayArgs{
		ClientToken: getClientToken(),
		Name:        "Test-SDK-NatGateway-CU",
		VpcId:       VPCID,
		CuNum:       "3",
		Billing: &Billing{
			PaymentTiming: PAYMENT_TIMING_POSTPAID,
		},
	}
	result, err := VPC_CLIENT.CreateNatGateway(args)
	ExpectEqual(t.Errorf, nil, err)
	NatID = result.NatId

	err = waitStateForNatGateway(NatID, NAT_STATUS_UNCONFIGURED)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetRouteTableDetail(t *testing.T) {
	RouteTableID = ""
	result, err := VPC_CLIENT.GetRouteTableDetail(RouteTableID, VPCID)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
	ExpectEqual(t.Errorf, nil, err)
	RouteTableID = result.RouteTableId
}

func TestCreateRouteRule(t *testing.T) {
	args := &CreateRouteRuleArgs{
		ClientToken:        getClientToken(),
		RouteTableId:       RouteTableID,
		SourceAddress:      "192.168.1.0/24",
		DestinationAddress: "172.17.0.0/16",
		NexthopType:        NEXTHOP_TYPE_NAT,
		NexthopId:          NatID,
		Description:        "test route rule",
	}
	result, err := VPC_CLIENT.CreateRouteRule(args)
	ExpectEqual(t.Errorf, nil, err)

	RouteRuleID = result.RouteRuleId

	routeTable, err := VPC_CLIENT.GetRouteTableDetail("", VPCID)
	ExpectEqual(t.Errorf, nil, err)
	isExist := false
	for _, rule := range routeTable.RouteRules {
		if rule.RouteRuleId == result.RouteRuleId {
			isExist = true
			ExpectEqual(t.Errorf, RouteTableID, rule.RouteTableId)
			ExpectEqual(t.Errorf, "192.168.1.0/24", rule.SourceAddress)
			ExpectEqual(t.Errorf, "172.17.0.0/16", rule.DestinationAddress)
			ExpectEqual(t.Errorf, NEXTHOP_TYPE_NAT, rule.NexthopType)
			ExpectEqual(t.Errorf, NatID, rule.NexthopId)
			ExpectEqual(t.Errorf, "test route rule", rule.Description)
		}
	}
	if !isExist {
		t.Errorf("Test route rule failed.")
	}
}

func TestCreateEtGatewayRouteRule(t *testing.T) {
	RouteTableID = ""
	var SourceAddress = "12.0.0.0/25"
	var DestinationAddress = "2.2.2.6/32"
	var Description = "sdk test etGateway route rule"
	var NexthopType = NEXTHOP_TYPE_ETGATEWAY

	mulargs := &CreateRouteRuleArgs{
		ClientToken:        getClientToken(),
		RouteTableId:       RouteTableID,
		SourceAddress:      SourceAddress,
		DestinationAddress: DestinationAddress,
		NextHopList: []NextHop{
			{
				NexthopId:   "",
				NexthopType: NexthopType,
				PathType:    "ha:active",
			}, {
				NexthopId:   "",
				NexthopType: NexthopType,
				PathType:    "ha:standby",
			},
		},
		Description: Description,
	}
	result, err := VPC_CLIENT.CreateRouteRule(mulargs)

	r, err := json.Marshal(result)
	fmt.Println(string(r))
	fmt.Println(err)

	ExpectEqual(t.Errorf, nil, err)

	var RouteRuleIds = result.RouteRuleIds
	routeTable, err := VPC_CLIENT.GetRouteTableDetail(RouteTableID, VPCID)
	ExpectEqual(t.Errorf, nil, err)
	isExist := false
	for _, rule := range routeTable.RouteRules {
		for _, RouteRuleId := range RouteRuleIds {
			if rule.RouteRuleId == RouteRuleId {
				isExist = true
				ExpectEqual(t.Errorf, RouteTableID, rule.RouteTableId)
				ExpectEqual(t.Errorf, SourceAddress, rule.SourceAddress)
				ExpectEqual(t.Errorf, DestinationAddress, rule.DestinationAddress)
				ExpectEqual(t.Errorf, NexthopType, rule.NexthopType)
				ExpectEqual(t.Errorf, Description, rule.Description)
			}
		}
	}
	if !isExist {
		t.Errorf("Test route rule failed.")
	}
}

func TestListNatGateway(t *testing.T) {
	args := &ListNatGatewayArgs{
		VpcId: VPCID,
		NatId: NatID,
	}
	result, err := VPC_CLIENT.ListNatGateway(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(result.Nats))
}

func TestGetNatGatewayDetail(t *testing.T) {
	result, err := VPC_CLIENT.GetNatGatewayDetail("nat-bzrav7t2ppb5")
	tags := []model.TagModel{
		{
			TagKey:   "tagKey",
			TagValue: "tagValue",
		},
	}
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Test-SDK-NatGateway", result.Name)
	ExpectEqual(t.Errorf, VPCID, result.VpcId)
	ExpectEqual(t.Errorf, NAT_GATEWAY_SPEC_SMALL, result.Spec)
	ExpectEqual(t.Errorf, PAYMENT_TIMING_POSTPAID, result.PaymentTiming)
	ExpectEqual(t.Errorf, tags, result.Tags)
}

func TestUpdateNatGateway(t *testing.T) {
	args := &UpdateNatGatewayArgs{
		ClientToken: getClientToken(),
		Name:        "Test-SDK-NatGateway-update",
	}
	err := VPC_CLIENT.UpdateNatGateway(NatID, args)
	ExpectEqual(t.Errorf, nil, err)

	result, err := VPC_CLIENT.GetNatGatewayDetail(NatID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Test-SDK-NatGateway-update", result.Name)
}

func TestResizeNatGateway(t *testing.T) {
	args := &ResizeNatGatewayArgs{
		ClientToken: getClientToken(),
		CuNum:       8,
	}
	err := VPC_CLIENT.ResizeNatGateway(NatID, args)
	ExpectEqual(t.Errorf, nil, err)
	err = waitCuNumForNatGateway(NatID, args.CuNum)
	ExpectEqual(t.Errorf, nil, err)
}

func TestBindEips(t *testing.T) {
	// create eip first
	args := &eip.CreateEipArgs{
		Name:            "sdk-eip",
		BandWidthInMbps: 10,
		Billing: &eip.Billing{
			PaymentTiming: "Postpaid",
			BillingMethod: "ByTraffic",
		},
		ClientToken: getClientToken(),
	}
	result, err := EIP_CLIENT.CreateEip(args)
	ExpectEqual(t.Errorf, nil, err)
	EIPAddress = result.Eip

	// wait until the eip available
	err = waitStateForEIP(EIPAddress, "available")
	ExpectEqual(t.Errorf, nil, err)

	// bind eip
	bindEipArgs := &BindEipsArgs{
		ClientToken: getClientToken(),
		Eips:        []string{EIPAddress},
	}
	err = VPC_CLIENT.BindEips(NatID, bindEipArgs)
	ExpectEqual(t.Errorf, nil, err)

	// wait until the eip bind completed
	err = waitStateForNatGateway(NatID, NAT_STATUS_ACTIVE)
	ExpectEqual(t.Errorf, nil, err)
}

func TestUnBindEips(t *testing.T) {
	// unbind eip
	args := &UnBindEipsArgs{
		ClientToken: getClientToken(),
		Eips:        []string{EIPAddress},
	}
	err := VPC_CLIENT.UnBindEips(NatID, args)
	ExpectEqual(t.Errorf, nil, err)

	// wait until the eip unbind completed
	err = waitStateForNatGateway(NatID, NAT_STATUS_UNCONFIGURED)
	ExpectEqual(t.Errorf, nil, err)

	// delete eip
	err = EIP_CLIENT.DeleteEip(EIPAddress, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestRenewNatGateway(t *testing.T) {
	args := &RenewNatGatewayArgs{
		ClientToken: getClientToken(),
		Billing: &Billing{
			PaymentTiming: PAYMENT_TIMING_POSTPAID,
		},
	}
	err := VPC_CLIENT.RenewNatGateway(NatID, args)
	if err == nil {
		t.Errorf("Test RenewNatGateway failed.")
	}
}

func TestCreateNatGatewaySnatRule(t *testing.T) {
	args := &CreateNatGatewaySnatRuleArgs{
		RuleName:          "sdk-test",
		PublicIpAddresses: []string{"100.88.10.84"},
		SourceCIDR:        "192.168.3.33",
	}
	result, err := VPC_CLIENT.CreateNatGatewaySnatRule("nat-b1jb3b5e34tc", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestBatchCreateNatGatewaySnatRule(t *testing.T) {
	snatargs1 := SnatRuleArgs{
		RuleName:          "sdk-test-b1",
		PublicIpAddresses: []string{"100.88.9.91", "100.88.2.154"},
		SourceCIDR:        "192.168.3.6",
	}
	snatargs2 := SnatRuleArgs{
		RuleName:          "sdk-test-b2",
		PublicIpAddresses: []string{"100.88.10.84", "100.88.2.154"},
		SourceCIDR:        "192.168.3.7",
	}

	args := &BatchCreateNatGatewaySnatRuleArgs{
		NatId: "nat-b1jb3b5e34tc",
	}

	args.SnatRules = append(args.SnatRules, snatargs1)
	args.SnatRules = append(args.SnatRules, snatargs2)
	result, err := VPC_CLIENT.BatchCreateNatGatewaySnatRule(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestUpdateNatGatewaySnatRule(t *testing.T) {
	args := &UpdateNatGatewaySnatRuleArgs{
		RuleName:   "sdk-test-1",
		SourceCIDR: "192.168.3.66",
	}
	result := VPC_CLIENT.UpdateNatGatewaySnatRule("nat-b1jb3b5e34tc", "rule-7zr5941yngks", args)
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
}

func TestDeleteNatGatewaySnatRule(t *testing.T) {
	result := VPC_CLIENT.DeleteNatGatewaySnatRule("nat-b1jb3b5e34tc", "rule-7zr5941yngks", getClientToken())
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
}

func TestListNatGatewaySnatRules(t *testing.T) {
	args := &ListNatGatewaySnatRuleArgs{
		NatId: "nat-b1jb3b5e34tc",
	}
	result, err := VPC_CLIENT.ListNatGatewaySnatRules(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestCreateNatGatewayDnatRule(t *testing.T) {
	args := &CreateNatGatewayDnatRuleArgs{
		RuleName:         "sdk-test",
		PublicIpAddress:  "100.88.0.217",
		PrivateIpAddress: "192.168.1.4",
		Protocol:         "TCP",
		PublicPort:       "121",
		PrivatePort:      "122",
	}
	result, err := VPC_CLIENT.CreateNatGatewayDnatRule("nat-b1jb3b5e34tc", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestBatchCreateNatGatewayDnatRule(t *testing.T) {
	snatargs1 := DnatRuleArgs{
		RuleName:         "sdk-test-db1",
		PublicIpAddress:  "100.88.0.217",
		PrivateIpAddress: "192.168.1.21",
		Protocol:         "TCP",
		PublicPort:       "1211",
		PrivatePort:      "1221",
	}
	snatargs2 := DnatRuleArgs{
		RuleName:         "sdk-test-db2",
		PublicIpAddress:  "100.88.0.217",
		PrivateIpAddress: "192.168.1.22",
		Protocol:         "TCP",
		PublicPort:       "1212",
		PrivatePort:      "1222",
	}

	args := &BatchCreateNatGatewayDnatRuleArgs{
		Rules: []DnatRuleArgs{snatargs1, snatargs2},
	}

	result, err := VPC_CLIENT.BatchCreateNatGatewayDnatRule("nat-b1jb3b5e34tc", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestUpdateNatGatewayDnatRule(t *testing.T) {
	args := &UpdateNatGatewayDnatRuleArgs{
		RuleName:         "sdk-test-3",
		PrivateIpAddress: "192.168.1.5",
	}
	result := VPC_CLIENT.UpdateNatGatewayDnatRule("nat-b1jb3b5e34tc", "rule-8gee5abqins0", args)
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
}

func TestDeleteNatGatewayDnatRule(t *testing.T) {
	result := VPC_CLIENT.DeleteNatGatewayDnatRule("nat-b1jb3b5e34tc", "rule-8gee5abqins0", getClientToken())
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
}

func TestListNatGatewayDnatRules(t *testing.T) {
	args := &ListNatGatewaDnatRuleArgs{
		MaxKeys: 2,
		Marker:  "rule-29n3en0z8tku",
	}
	result, err := VPC_CLIENT.ListNatGatewayDnatRules("nat-b1jb3b5e34tc", args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestCreatePeerConn(t *testing.T) {
	// create another VPC
	createVPCArgs := &CreateVPCArgs{
		Name:        "TestSDK-VPC-Peer",
		Description: "vpc test",
		Cidr:        "172.17.0.0/16",
		ClientToken: getClientToken(),
	}
	vpcResult, err := VPC_CLIENT.CreateVPC(createVPCArgs)
	ExpectEqual(t.Errorf, nil, err)
	PeerVPCID = vpcResult.VPCID

	args := &CreatePeerConnArgs{
		ClientToken:     getClientToken(),
		BandwidthInMbps: 10,
		Description:     "test peer conn",
		LocalIfName:     "local-interface",
		LocalVpcId:      VPCID,
		PeerVpcId:       PeerVPCID,
		PeerRegion:      region,
		PeerIfName:      "peer-interface",
		Billing: &Billing{
			PaymentTiming: PAYMENT_TIMING_POSTPAID,
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",
				TagValue: "tagValue",
			},
		},
	}
	result, err := VPC_CLIENT.CreatePeerConn(args)
	ExpectEqual(t.Errorf, nil, err)

	PeerConnID = result.PeerConnId

	// wait until peerconn active
	err = waitStateForPeerConn(PeerConnID, PEERCONN_STATUS_ACTIVE)
	ExpectEqual(t.Errorf, nil, err)
}

func TestListPeerConn(t *testing.T) {
	args := &ListPeerConnsArgs{
		VpcId: VPCID,
	}
	result, err := VPC_CLIENT.ListPeerConn(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(result.PeerConns))
}

func TestGetPeerConnDetail(t *testing.T) {
	result, err := VPC_CLIENT.GetPeerConnDetail(PeerConnID, PEERCONN_ROLE_INITIATOR)
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
	tags := []model.TagModel{
		{
			TagKey:   "tagKey",
			TagValue: "tagValue",
		},
	}
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 10, result.BandwidthInMbps)
	ExpectEqual(t.Errorf, "test peer conn", result.Description)
	ExpectEqual(t.Errorf, "local-interface", result.LocalIfName)
	ExpectEqual(t.Errorf, VPCID, result.LocalVpcId)
	ExpectEqual(t.Errorf, PAYMENT_TIMING_POSTPAID, result.PaymentTiming)
	ExpectEqual(t.Errorf, tags, result.Tags)

	LocalIfID = result.LocalIfId
}

func TestUpdatePeerConn(t *testing.T) {
	args := &UpdatePeerConnArgs{
		LocalIfId:   LocalIfID,
		LocalIfName: "local-interface-update",
		Description: "peer conn update",
	}
	err := VPC_CLIENT.UpdatePeerConn(PeerConnID, args)
	ExpectEqual(t.Errorf, nil, err)

	result, err := VPC_CLIENT.GetPeerConnDetail(PeerConnID, PEERCONN_ROLE_INITIATOR)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "peer conn update", result.Description)
	ExpectEqual(t.Errorf, "local-interface-update", result.LocalIfName)
	ExpectEqual(t.Errorf, VPCID, result.LocalVpcId)
	ExpectEqual(t.Errorf, PAYMENT_TIMING_POSTPAID, result.PaymentTiming)
}

func TestAcceptPeerConnApply(t *testing.T) {
	err := VPC_CLIENT.AcceptPeerConnApply(PeerConnID, getClientToken())
	if err == nil {
		t.Errorf("Test AcceptPeerConnApply failed.")
	}
}

func TestRejectPeerConnReject(t *testing.T) {
	err := VPC_CLIENT.RejectPeerConnApply(PeerConnID, getClientToken())
	if err == nil {
		t.Errorf("Test RejectPeerConnApply failed.")
	}
}

func TestResizePeerConn(t *testing.T) {
	args := &ResizePeerConnArgs{
		NewBandwidthInMbps: 20,
		ClientToken:        getClientToken(),
	}
	err := VPC_CLIENT.ResizePeerConn(PeerConnID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestRenewPeerConn(t *testing.T) {
	args := &RenewPeerConnArgs{
		ClientToken: getClientToken(),
		Billing: &Billing{
			PaymentTiming: PAYMENT_TIMING_POSTPAID,
		},
	}
	err := VPC_CLIENT.RenewPeerConn(PeerConnID, args)
	if err == nil {
		t.Errorf("Test RenewPeerConn failed.")
	}
}

func TestOpenPeerConnSyncDNS(t *testing.T) {
	args := &PeerConnSyncDNSArgs{
		Role:        PEERCONN_ROLE_INITIATOR,
		ClientToken: getClientToken(),
	}
	err := VPC_CLIENT.OpenPeerConnSyncDNS(PeerConnID, args)
	ExpectEqual(t.Errorf, nil, err)

	// wait until dns sync completed
	err = waitStateForPeerConn(PeerConnID, DNS_STATUS_OPEN)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClosePeerConnSyncDNS(t *testing.T) {
	args := &PeerConnSyncDNSArgs{
		Role:        PEERCONN_ROLE_INITIATOR,
		ClientToken: getClientToken(),
	}
	err := VPC_CLIENT.ClosePeerConnSyncDNS(PeerConnID, args)
	ExpectEqual(t.Errorf, nil, err)

	// wait until dns sync completed
	err = waitStateForPeerConn(PeerConnID, DNS_STATUS_CLOSE)
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeletePeerConn(t *testing.T) {
	err := VPC_CLIENT.DeletePeerConn(PeerConnID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)

	err = waitStateForPeerConn(PeerConnID, PEERCONN_STATUS_DELETED)
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteNatGateway(t *testing.T) {
	err := VPC_CLIENT.DeleteNatGateway(NatID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)

	err = waitStateForNatGateway(NatID, NAT_STATUS_DELETED)
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteAclRule(t *testing.T) {
	err := VPC_CLIENT.DeleteAclRule(AclRuleID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteRouteRule(t *testing.T) {
	err := VPC_CLIENT.DeleteRouteRule(RouteRuleID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteSubnet(t *testing.T) {
	err := VPC_CLIENT.DeleteSubnet(SubnetID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteVPC(t *testing.T) {
	// The vpc will be SUBNET_INUSE status after the subnet deleted in a period of time.
	// So delete vpc by waitStateForVPC currently.
	err := waitStateForVPC(VPCID)
	ExpectEqual(t.Errorf, nil, err)

	err = waitStateForVPC(PeerVPCID)
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}

func waitStateForNatGateway(natID string, status NatStatusType) error {
	ticker := time.NewTicker(2 * time.Second)
	errChan := make(chan error, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				nat, err := VPC_CLIENT.GetNatGatewayDetail(natID)
				if err != nil {
					if !strings.Contains(err.Error(), "NoSuchNat") {
						errChan <- err
						return
					}
					nat = &NAT{Status: NAT_STATUS_DELETED}
				}
				if nat.Status == status {
					errChan <- nil
					return
				}
			}
		}
	}()

	select {
	case <-time.After(10 * time.Minute):
		return fmt.Errorf("Test nat gateway %s timeout.", natID)
	case err := <-errChan:
		return err
	}
}

func waitCuNumForNatGateway(natID string, cuNum int) error {
	ticker := time.NewTicker(2 * time.Second)
	errChan := make(chan error, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				nat, err := VPC_CLIENT.GetNatGatewayDetail(natID)
				r, _ := json.Marshal(nat)
				fmt.Println(string(r))
				fmt.Println(cuNum)
				if err != nil {
					if !strings.Contains(err.Error(), "NoSuchNat") {
						errChan <- err
						return
					}
				}
				if nat.CuNum == cuNum {
					errChan <- nil
					return
				}
			}
		}
	}()

	select {
	case <-time.After(10 * time.Minute):
		return fmt.Errorf("Test nat gateway %s timeout.", natID)
	case err := <-errChan:
		return err
	}
}

func waitStateForEIP(eipAddress, status string) error {
	ticker := time.NewTicker(2 * time.Second)
	errChan := make(chan error, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				args := &eip.ListEipArgs{Eip: EIPAddress}
				eips, err := EIP_CLIENT.ListEip(args)
				if err != nil {
					errChan <- err
					return
				}
				if len(eips.EipList) == 1 && eips.EipList[0].Status == status {
					errChan <- nil
					return
				}
			}
		}
	}()

	select {
	case <-time.After(10 * time.Minute):
		return fmt.Errorf("Test eip %s timeout.", eipAddress)
	case err := <-errChan:
		return err
	}
}

func waitStateForVPC(vpcID string) error {
	ticker := time.NewTicker(2 * time.Second)
	errChan := make(chan error, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				err := VPC_CLIENT.DeleteVPC(vpcID, getClientToken())
				if err != nil && strings.Contains(err.Error(), "SUBNET_INUSE") {
					continue
				}
				errChan <- err
				return
			}
		}
	}()

	select {
	case <-time.After(10 * time.Minute):
		return fmt.Errorf("Test VPC %s timeout.", vpcID)
	case err := <-errChan:
		return err
	}
}

func waitStateForPeerConn(peerConnID string, status interface{}) error {
	ticker := time.NewTicker(2 * time.Second)
	errChan := make(chan error, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				result, err := VPC_CLIENT.GetPeerConnDetail(peerConnID, PEERCONN_ROLE_INITIATOR)
				if err != nil {
					// not found error
					if !strings.Contains(err.Error(), "EOF") {
						errChan <- err
						return
					}
					result = &PeerConn{Status: PEERCONN_STATUS_DELETED}
				}

				switch status.(type) {
				case PeerConnStatusType:
					if result.Status == status {
						errChan <- nil
						return
					}
				case DnsStatusType:
					if result.DnsStatus == status {
						errChan <- nil
						return
					}
				default:
					errChan <- fmt.Errorf("The status %s type is not supported.", status)
					return
				}
			}
		}
	}()

	select {
	case <-time.After(10 * time.Minute):
		return fmt.Errorf("Test peer conn %s timeout.", peerConnID)
	case err := <-errChan:
		return err
	}
}

func TestGetNetworkTopologyInfo(t *testing.T) {
	args := &GetNetworkTopologyArgs{
		HostId: "",
		HostIp: "",
	}
	result, err := VPC_CLIENT.GetNetworkTopologyInfo(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestClient_BindDnatEips(t *testing.T) {
	args := &BindDnatEipsArgs{
		ClientToken: getClientToken(),
		DnatEips:    []string{"100.88.14.243"},
	}
	err := VPC_CLIENT.BindDnatEips("nat-bc39ugw5ry9z", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnBindDnatEips(t *testing.T) {
	args := &UnBindDnatEipsArgs{
		ClientToken: getClientToken(),
		DnatEips:    []string{"100.88.14.243"},
	}
	err := VPC_CLIENT.UnBindDnatEips("nat-bc39ugw5ry9z", args)
	ExpectEqual(t.Errorf, nil, err)
}
