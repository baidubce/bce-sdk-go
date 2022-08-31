package bec

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

//////////////////////////////////////////////
// deploy set test
//////////////////////////////////////////////

func TestCreateAppBlb(t *testing.T) {
	getReq := &api.CreateAppBlbRequest{
		Name:         "wcw_test_applb",
		Desc:         "wcw-test",
		RegionId:     "cn-hangzhou-cm",
		NeedPublicIp: true,
		SubnetId:     "sbn-tafnx9dd",
		VpcId:        "vpc-wljmvzmt",
	}
	res, err := CLIENT.CreateAppBlb("testCreateAppBlb", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestUpdateAppBlb(t *testing.T) {
	getReq := &api.ModifyBecBlbRequest{
		Name: "wcw_test_applb",
		Desc: "wcw-test1",
	}
	err := CLIENT.UpdateAppBlb("testUpdateAppBlb", "lb-zo8wibx1", getReq)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetAppBlbList(t *testing.T) {
	getReq := &api.MarkerRequest{}
	res, err := CLIENT.GetAppBlbList(getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestGetAppBlbDetails(t *testing.T) {

	res, err := CLIENT.GetAppBlbDetails("applb-cn-hangzhou-cm-h9nh3vpe")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestDeleteAppBlbInstance(t *testing.T) {

	err := CLIENT.DeleteAppBlbInstance("applb-cn-hangzhou-cm-h9nh3vpe", "")
	ExpectEqual(t.Errorf, nil, err)
}
func TestCreateTcpListener(t *testing.T) {
	getReq := &api.CreateBecAppBlbTcpListenerRequest{
		ListenerPort:      80,
		Scheduler:         "RoundRobin",
		TcpSessionTimeout: 1000,
	}
	err := CLIENT.CreateTcpListener("testCreateTcpListener", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
}
func TestCreateUdpListener(t *testing.T) {
	getReq := &api.CreateBecAppBlbUdpListenerRequest{
		ListenerPort:      80,
		Scheduler:         "RoundRobin",
		UdpSessionTimeout: 1000,
	}
	err := CLIENT.CreateUdpListener("testCreateTcpListener", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
}
func TestUpdateTcpListener(t *testing.T) {
	getReq := &api.UpdateBecAppBlbTcpListenerRequest{
		Scheduler:         "RoundRobin",
		TcpSessionTimeout: 800,
	}
	err := CLIENT.UpdateTcpListener("testUpdateTcpListener", "applb-cn-hangzhou-cm-h9nh3vpe", "80", getReq)
	ExpectEqual(t.Errorf, nil, err)
}

func TestUpdateUdpListener(t *testing.T) {
	getReq := &api.UpdateBecAppBlbUdpListenerRequest{
		Scheduler:         "RoundRobin",
		UdpSessionTimeout: 800,
	}
	err := CLIENT.UpdateUdpListener("testUpdateUdpListener", "applb-cn-hangzhou-cm-h9nh3vpe", "80", getReq)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetTcpListener(t *testing.T) {
	getReq := &api.GetBecAppBlbListenerRequest{
		ListenerPort: 80,
	}
	res, err := CLIENT.GetTcpListener("applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGetUdpListener(t *testing.T) {
	getReq := &api.GetBecAppBlbListenerRequest{
		ListenerPort: 80,
	}
	res, err := CLIENT.GetUdpListener("applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestDeleteAppBlbListener(t *testing.T) {
	getReq := &api.DeleteBlbListenerRequest{
		PortTypeList: []api.PortTypeList{
			{
				Port: 80,
				Type: "TCP",
			},
			{
				Port: 80,
				Type: "UDP",
			},
		},
	}
	err := CLIENT.DeleteAppBlbListener("applb-cn-hangzhou-cm-h9nh3vpe", "deleteApplbInstance", getReq)
	ExpectEqual(t.Errorf, nil, err)
}
func TestCreateIpGroup(t *testing.T) {
	getReq := &api.CreateBlbIpGroupRequest{
		Name: "wcw-testIpGroup",
		Desc: "wcw-test",
	}
	res, err := CLIENT.CreateIpGroup("testIpGroup", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

// bec_ip_group-ukadxdrq
func TestUpdateIpGroup(t *testing.T) {
	getReq := &api.UpdateBlbIpGroupRequest{
		Name:      "wcw-testIpGroupupdate",
		Desc:      "wcw-testupdate",
		IpGroupId: "bec_ip_group-ukadxdrq",
	}
	err := CLIENT.UpdateIpGroup("testIpGroup", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
}
func TestGetIpGroup(t *testing.T) {
	getReq := &api.GetBlbIpGroupListRequest{}
	res, err := CLIENT.GetIpGroup("applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestDeleteIpGroup(t *testing.T) {
	getReq := &api.DeleteBlbIpGroupRequest{
		IpGroupId: "bec_ip_group-ukadxdrq",
	}
	err := CLIENT.DeleteIpGroup("testDeleteIpGroup", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
}
func TestCreateIpGroupPolicy(t *testing.T) {
	getReq := &api.CreateBlbIpGroupBackendPolicyRequest{
		IpGroupId:                   "bec_ip_group-ukadxdrq",
		Type:                        "TCP",
		HealthCheck:                 "TCP",
		HealthCheckPort:             80,
		HealthCheckTimeoutInSecond:  10,
		HealthCheckIntervalInSecond: 3,
		HealthCheckDownRetry:        4,
		HealthCheckUpRetry:          5,
	}
	res, err := CLIENT.CreateIpGroupPolicy("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestUpdateIpGroupPolicy(t *testing.T) {
	getReq := &api.UpdateBlbIpGroupBackendPolicyRequest{
		IpGroupId:       "bec_ip_group-ukadxdrq",
		Id:              "bec_ip_group_policy-yodpsqqr",
		HealthCheckPort: 80,
	}
	err := CLIENT.UpdateIpGroupPolicy("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
}
func TestGetIpGroupPolicy(t *testing.T) {
	getReq := &api.GetBlbIpGroupPolicyListRequest{
		IpGroupId: "bec_ip_group-ukadxdrq",
	}
	res, err := CLIENT.GetIpGroupPolicyList("applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestDeleteIpGroupPolicy(t *testing.T) {
	getReq := &api.DeleteBlbIpGroupBackendPolicyRequest{
		IpGroupId:           "bec_ip_group-ukadxdrq",
		BackendPolicyIdList: []string{"bec_ip_group_policy-yodpsqqr"},
	}
	err := CLIENT.DeleteIpGroupPolicy("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)

}
func TestCreateIpGroupMember(t *testing.T) {
	getReq := &api.CreateBlbIpGroupMemberRequest{
		IpGroupId: "bec_ip_group-ukadxdrq",
		MemberList: []api.BlbIpGroupMember{
			{
				Ip:     "172.16.240.25",
				Port:   90,
				Weight: 100,
			},
		},
	}
	res, err := CLIENT.CreateIpGroupMember("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestUpdateIpGroupMember(t *testing.T) {
	getReq := &api.UpdateBlbIpGroupMemberRequest{
		IpGroupId: "bec_ip_group-ukadxdrq",
		MemberList: []api.UpdateBlbIpGroupMember{
			{
				MemberId: "bec_ip_member-ouiinabp",
				Port:     8080,
				Weight:   100,
			},
		},
	}
	err := CLIENT.UpdateIpGroupMember("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
}
func TestGetIpGroupMemberList(t *testing.T) {
	getReq := &api.GetBlbIpGroupMemberListRequest{
		IpGroupId: "bec_ip_group-ukadxdrq",
	}
	res, err := CLIENT.GetIpGroupMemberList("applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestDeleteIpGroupMember(t *testing.T) {
	getReq := &api.DeleteBlbIpGroupBackendMemberRequest{
		IpGroupId:    "bec_ip_group-ukadxdrq",
		MemberIdList: []string{"bec_ip_member-ouiinabp"},
	}
	err := CLIENT.DeleteIpGroupMember("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)

}
func TestCreateListenerPolicy(t *testing.T) {
	getReq := &api.CreateAppBlbPoliciesRequest{
		ListenerPort: 80,
		AppPolicyVos: []api.AppPolicyVo{
			{
				AppIpGroupId: "bec_ip_group-ukadxdrq",
				Priority:     1,
				Desc:         "wcw-test",
			},
		},
	}
	err := CLIENT.CreateListenerPolicy("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetListenerPolicy(t *testing.T) {
	getReq := &api.GetBlbListenerPolicyRequest{
		Port: 80,
	}
	res, err := CLIENT.GetListenerPolicy("applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
func TestDeleteListenerPolicy(t *testing.T) {
	getReq := &api.DeleteAppBlbPoliciesRequest{
		Port: 80,
		PolicyIdList: []string{
			"bec_policy-scr9cwtk",
		},
	}
	err := CLIENT.DeleteListenerPolicy("", "applb-cn-hangzhou-cm-h9nh3vpe", getReq)
	ExpectEqual(t.Errorf, nil, err)

}
