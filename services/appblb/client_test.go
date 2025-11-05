package appblb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	APPBLB_CLIENT                    *Client
	APPBLB_ID                        string
	APPBLB_SERVERGROUP_ID            string
	APPBLB_SERVERGROUPPORT_ID        string
	APPBLB_POLICY_ID                 string
	APPBLB_IPGROUP_ID                string
	IPGROUP_MEMBER_ID                string
	APPBLB_IPGROUPP_BACKENDPOLICY_ID string

	// set these values before start test
	VPC_TEST_ID           = ""
	SUBNET_TEST_ID        = ""
	INSTANCE_ID           = ""
	CERT_ID               = ""
	IPGROUP_MEMBER_IP     = ""
	CLUSTER_PROPERTY_TEST = ""
	TEST_APPBLB_ID        = ""
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

// In order to more conveniently represent some bool types
var (
	trueVal  = true
	falseVal = false
	True     = &trueVal
	False    = &falseVal
)

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	APPBLB_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
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

func TestClient_CreateLoadBalancer(t *testing.T) {
	AllowDelete := true
	AllocateIpv6 := true
	createArgs := &CreateLoadBalancerArgs{
		ClientToken:  getClientToken(),
		Name:         "sdkBlb",
		VpcId:        VPC_TEST_ID,
		SubnetId:     SUBNET_TEST_ID,
		AllowDelete:  &AllowDelete,
		AllocateIpv6: &AllocateIpv6,
	}

	createResult, err := APPBLB_CLIENT.CreateLoadBalancer(createArgs)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_ID = createResult.BlbId
}

func TestClient_UpdateLoadBalancer(t *testing.T) {
	updateArgs := &UpdateLoadBalancerArgs{
		Name:        "testSdk",
		Description: "test desc",
	}
	err := APPBLB_CLIENT.UpdateLoadBalancer(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeLoadBalancers(t *testing.T) {
	describeArgs := &DescribeLoadBalancersArgs{}
	res, err := APPBLB_CLIENT.DescribeLoadBalancers(describeArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeLoadBalancerDetail(t *testing.T) {
	res, err := APPBLB_CLIENT.DescribeLoadBalancerDetail(APPBLB_ID)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppServerGroup(t *testing.T) {
	createArgs := &CreateAppServerGroupArgs{
		ClientToken: getClientToken(),
		Name:        "sdkTest",
	}
	createResult, err := APPBLB_CLIENT.CreateAppServerGroup(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_SERVERGROUP_ID = createResult.Id
}

func TestClient_UpdateAppServerGroup(t *testing.T) {
	updateArgs := &UpdateAppServerGroupArgs{
		SgId:        APPBLB_SERVERGROUP_ID,
		Name:        "testSdk",
		Description: "test desc",
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.UpdateAppServerGroup(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppServerGroup(t *testing.T) {
	describeArgs := &DescribeAppServerGroupArgs{}
	_, err := APPBLB_CLIENT.DescribeAppServerGroup(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppServerGroupPort(t *testing.T) {
	createArgs := &CreateAppServerGroupPortArgs{
		ClientToken: getClientToken(),
		SgId:        APPBLB_SERVERGROUP_ID,
		Port:        80,
		Type:        "TCP",
	}
	createResult, err := APPBLB_CLIENT.CreateAppServerGroupPort(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_SERVERGROUPPORT_ID = createResult.Id
}

func TestClient_UpdateAppServerGroupPort(t *testing.T) {
	updateArgs := &UpdateAppServerGroupPortArgs{
		ClientToken:                 getClientToken(),
		SgId:                        APPBLB_SERVERGROUP_ID,
		PortId:                      APPBLB_SERVERGROUPPORT_ID,
		HealthCheck:                 "TCP",
		HealthCheckPort:             30,
		HealthCheckIntervalInSecond: 10,
		HealthCheckTimeoutInSecond:  10,
	}
	err := APPBLB_CLIENT.UpdateAppServerGroupPort(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppServerGroupPort(t *testing.T) {
	deleteArgs := &DeleteAppServerGroupPortArgs{
		SgId:        APPBLB_SERVERGROUP_ID,
		PortIdList:  []string{APPBLB_SERVERGROUPPORT_ID},
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppServerGroupPort(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateBlbRs(t *testing.T) {
	weigth := 30
	createArgs := &CreateBlbRsArgs{
		BlbRsWriteOpArgs: BlbRsWriteOpArgs{
			ClientToken: getClientToken(),
			SgId:        APPBLB_SERVERGROUP_ID,
			BackendServerList: []AppBackendServer{
				{InstanceId: INSTANCE_ID, Weight: &weigth},
			},
		},
	}
	err := APPBLB_CLIENT.CreateBlbRs(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateBlbRs(t *testing.T) {
	weigth := 50
	updateArgs := &UpdateBlbRsArgs{
		BlbRsWriteOpArgs: BlbRsWriteOpArgs{
			ClientToken: getClientToken(),
			SgId:        APPBLB_SERVERGROUP_ID,
			BackendServerList: []AppBackendServer{
				{InstanceId: INSTANCE_ID, Weight: &weigth},
			},
		},
	}
	err := APPBLB_CLIENT.UpdateBlbRs(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeBlbRs(t *testing.T) {
	describeArgs := &DescribeBlbRsArgs{
		SgId: APPBLB_SERVERGROUP_ID,
	}
	_, err := APPBLB_CLIENT.DescribeBlbRs(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteBlbRs(t *testing.T) {
	deleteArgs := &DeleteBlbRsArgs{
		SgId:                APPBLB_SERVERGROUP_ID,
		BackendServerIdList: []string{INSTANCE_ID},
		ClientToken:         getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteBlbRs(APPBLB_ID, deleteArgs)

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeRsMount(t *testing.T) {
	_, err := APPBLB_CLIENT.DescribeRsMount(APPBLB_ID, APPBLB_SERVERGROUP_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeRsUnMount(t *testing.T) {
	_, err := APPBLB_CLIENT.DescribeRsUnMount(APPBLB_ID, APPBLB_SERVERGROUP_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppServerGroup(t *testing.T) {
	deleteArgs := &DeleteAppServerGroupArgs{
		SgId:        APPBLB_SERVERGROUP_ID,
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppServerGroup(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppIpGroup(t *testing.T) {
	createArgs := &CreateAppIpGroupArgs{
		ClientToken: getClientToken(),
		Name:        "sdkTest",
	}
	createResult, err := APPBLB_CLIENT.CreateAppIpGroup(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_IPGROUP_ID = createResult.Id
}

func TestClient_UpdateAppIpGroup(t *testing.T) {
	updateArgs := &UpdateAppIpGroupArgs{
		IpGroupId:   APPBLB_IPGROUP_ID,
		Name:        "testSdk",
		Desc:        "test desc",
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.UpdateAppIpGroup(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppIpGroup(t *testing.T) {
	describeArgs := &DescribeAppIpGroupArgs{}
	res, err := APPBLB_CLIENT.DescribeAppIpGroup(APPBLB_ID, describeArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppIpGroupBackendPolicy(t *testing.T) {
	createArgs := &CreateAppIpGroupBackendPolicyArgs{
		ClientToken: getClientToken(),
		IpGroupId:   APPBLB_IPGROUP_ID,
		Type:        "TCP",
	}
	err := APPBLB_CLIENT.CreateAppIpGroupBackendPolicy(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAppIpGroupBackendPolicy(t *testing.T) {
	updateArgs := &UpdateAppIpGroupBackendPolicyArgs{
		ClientToken:                 getClientToken(),
		IpGroupId:                   APPBLB_IPGROUP_ID,
		Id:                          APPBLB_IPGROUPP_BACKENDPOLICY_ID,
		HealthCheckIntervalInSecond: 10,
		HealthCheckTimeoutInSecond:  10,
	}
	err := APPBLB_CLIENT.UpdateAppIpGroupBackendPolicy(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppIpGroupBackendPolicy(t *testing.T) {
	deleteArgs := &DeleteAppIpGroupBackendPolicyArgs{
		IpGroupId:           APPBLB_IPGROUP_ID,
		BackendPolicyIdList: []string{APPBLB_IPGROUPP_BACKENDPOLICY_ID},
		ClientToken:         getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppIpGroupBackendPolicy(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppIpGroupMember(t *testing.T) {
	createArgs := &CreateAppIpGroupMemberArgs{
		AppIpGroupMemberWriteOpArgs: AppIpGroupMemberWriteOpArgs{
			ClientToken: getClientToken(),
			IpGroupId:   APPBLB_IPGROUP_ID,
			MemberList: []AppIpGroupMember{
				{Ip: IPGROUP_MEMBER_IP, Port: 30},
			},
		},
	}
	err := APPBLB_CLIENT.CreateAppIpGroupMember(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAppIpGroupMember(t *testing.T) {
	updateArgs := &UpdateAppIpGroupMemberArgs{
		AppIpGroupMemberWriteOpArgs: AppIpGroupMemberWriteOpArgs{
			ClientToken: getClientToken(),
			IpGroupId:   APPBLB_IPGROUP_ID,
			MemberList: []AppIpGroupMember{
				{Ip: IPGROUP_MEMBER_IP, Port: 50, MemberId: IPGROUP_MEMBER_ID},
			},
		},
	}
	err := APPBLB_CLIENT.UpdateAppIpGroupMember(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppIpGroupMember(t *testing.T) {
	describeArgs := &DescribeAppIpGroupMemberArgs{
		IpGroupId: APPBLB_IPGROUP_ID,
	}
	res, err := APPBLB_CLIENT.DescribeAppIpGroupMember(APPBLB_ID, describeArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppIpGroupMember(t *testing.T) {
	deleteArgs := &DeleteAppIpGroupMemberArgs{
		IpGroupId:    APPBLB_IPGROUP_ID,
		MemberIdList: []string{IPGROUP_MEMBER_ID},
		ClientToken:  getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppIpGroupMember(APPBLB_ID, deleteArgs)

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppTCPListener(t *testing.T) {
	createArgs := &CreateAppTCPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 90,
		Scheduler:    "RoundRobin",
	}
	err := APPBLB_CLIENT.CreateAppTCPListener(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAppTCPListener(t *testing.T) {
	updateArgs := &UpdateAppTCPListenerArgs{
		UpdateAppListenerArgs: UpdateAppListenerArgs{
			ListenerPort: 90,
			Scheduler:    "Hash",
		},
	}
	err := APPBLB_CLIENT.UpdateAppTCPListener(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppTCPListeners(t *testing.T) {
	describeArgs := &DescribeAppListenerArgs{
		ListenerPort: 90,
	}
	_, err := APPBLB_CLIENT.DescribeAppTCPListeners(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppUDPListener(t *testing.T) {
	createArgs := &CreateAppUDPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 91,
		Scheduler:    "RoundRobin",
	}
	err := APPBLB_CLIENT.CreateAppUDPListener(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAppUDPListener(t *testing.T) {
	updateArgs := &UpdateAppUDPListenerArgs{
		UpdateAppListenerArgs: UpdateAppListenerArgs{
			ListenerPort: 91,
			Scheduler:    "Hash",
		},
	}
	err := APPBLB_CLIENT.UpdateAppUDPListener(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppUDPListeners(t *testing.T) {
	describeArgs := &DescribeAppListenerArgs{
		ListenerPort: 53,
	}
	result, err := APPBLB_CLIENT.DescribeAppUDPListeners(APPBLB_ID, describeArgs)
	if err != nil {
		fmt.Println("get udp listener failed:", err)
	} else {
		fmt.Println("get udp listener success: ", result)
	}
}

func TestClient_CreateAppHTTPListener(t *testing.T) {
	createArgs := &CreateAppHTTPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 92,
		Scheduler:    "RoundRobin",
	}
	err := APPBLB_CLIENT.CreateAppHTTPListener(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAppHTTPListener(t *testing.T) {
	updateArgs := &UpdateAppHTTPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 92,
		Scheduler:    "LeastConnection",
		KeepSession:  True,
	}
	err := APPBLB_CLIENT.UpdateAppHTTPListener(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreatePolicys(t *testing.T) {
	createArgs := &CreatePolicysArgs{
		ListenerPort: 92,
		ClientToken:  getClientToken(),
		AppPolicyVos: []AppPolicy{
			{
				Description:      "test policy",
				AppServerGroupId: "",
				BackendPort:      92,
				Priority:         300,
				RuleList: []AppRule{
					{
						Key:   "*",
						Value: "*",
					},
				},
			},
		},
	}
	err := APPBLB_CLIENT.CreatePolicys(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreatePolicysIpGroup(t *testing.T) {
	createArgs := &CreatePolicysArgs{
		ListenerPort: 80,
		ClientToken:  getClientToken(),
		AppPolicyVos: []AppPolicy{
			{
				Description:  "test policy",
				AppIpGroupId: APPBLB_IPGROUP_ID,
				Priority:     100,
				RuleList: []AppRule{
					{
						Key:   "*",
						Value: "*",
					},
				},
			},
		},
	}
	err := APPBLB_CLIENT.CreatePolicys(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribePolicys(t *testing.T) {
	describeArgs := &DescribePolicysArgs{
		Port: 80,
	}
	res, err := APPBLB_CLIENT.DescribePolicys(APPBLB_ID, describeArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_POLICY_ID = res.PolicyList[0].Id
}

func TestClient_DeletePolicys(t *testing.T) {
	deleteArgs := &DeletePolicysArgs{
		Port:         80,
		PolicyIdList: []string{APPBLB_POLICY_ID},
		ClientToken:  getClientToken(),
	}
	err := APPBLB_CLIENT.DeletePolicys(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppHTTPListeners(t *testing.T) {
	describeArgs := &DescribeAppListenerArgs{
		ListenerPort: 92,
	}
	_, err := APPBLB_CLIENT.DescribeAppHTTPListeners(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppHTTPSListener(t *testing.T) {
	createArgs := &CreateAppHTTPSListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 93,
		Scheduler:    "RoundRobin",
		CertIds:      []string{CERT_ID},
	}
	err := APPBLB_CLIENT.CreateAppHTTPSListener(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAppHTTPSListener(t *testing.T) {
	updateArgs := &UpdateAppHTTPSListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 93,
		Scheduler:    "LeastConnection",
		KeepSession:  True,
		CertIds:      []string{CERT_ID},
	}
	err := APPBLB_CLIENT.UpdateAppHTTPSListener(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppHTTPSListeners(t *testing.T) {
	describeArgs := &DescribeAppListenerArgs{
		ListenerPort: 93,
	}
	_, err := APPBLB_CLIENT.DescribeAppHTTPSListeners(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppSSLListener(t *testing.T) {
	createArgs := &CreateAppSSLListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 94,
		Scheduler:    "RoundRobin",
		CertIds:      []string{CERT_ID},
	}
	err := APPBLB_CLIENT.CreateAppSSLListener(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAppSSLListener(t *testing.T) {
	updateArgs := &UpdateAppSSLListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 94,
		Scheduler:    "LeastConnection",
		CertIds:      []string{CERT_ID},
	}
	err := APPBLB_CLIENT.UpdateAppSSLListener(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppSSLListeners(t *testing.T) {
	describeArgs := &DescribeAppListenerArgs{
		ListenerPort: 94,
	}
	_, err := APPBLB_CLIENT.DescribeAppSSLListeners(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppAllListeners(t *testing.T) {
	describeArgs := &DescribeAppListenerArgs{}
	result, err := APPBLB_CLIENT.DescribeAppAllListeners(APPBLB_ID, describeArgs)
	if err != nil {
		fmt.Println("get all listener failed:", err)
	} else {
		fmt.Println("get all listener success: ", result)
	}
}

func TestClient_BindSecurityGroups(t *testing.T) {
	updateArgs := &UpdateSecurityGroupsArgs{
		ClientToken:      getClientToken(),
		SecurityGroupIds: []string{"sg-id"},
	}
	err := APPBLB_CLIENT.BindSecurityGroups(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnbindSecurityGroups(t *testing.T) {
	updateArgs := &UpdateSecurityGroupsArgs{
		ClientToken:      getClientToken(),
		SecurityGroupIds: []string{"sg-id"},
	}
	err := APPBLB_CLIENT.UnbindSecurityGroups(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeSecurityGroups(t *testing.T) {
	res, err := APPBLB_CLIENT.DescribeSecurityGroups(APPBLB_ID)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindEnterpriseSecurityGroups(t *testing.T) {
	updateArgs := &UpdateEnterpriseSecurityGroupsArgs{
		ClientToken:                getClientToken(),
		EnterpriseSecurityGroupIds: []string{"esg-id"},
	}
	err := APPBLB_CLIENT.BindEnterpriseSecurityGroups(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnbindEnterpriseSecurityGroups(t *testing.T) {
	updateArgs := &UpdateEnterpriseSecurityGroupsArgs{
		ClientToken:                getClientToken(),
		EnterpriseSecurityGroupIds: []string{"esg-id"},
	}
	err := APPBLB_CLIENT.UnbindEnterpriseSecurityGroups(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeEnterpriseSecurityGroups(t *testing.T) {
	res, err := APPBLB_CLIENT.DescribeEnterpriseSecurityGroups(APPBLB_ID)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppListeners(t *testing.T) {
	deleteArgs := &DeleteAppListenersArgs{
		PortList:    []uint16{90, 91, 92, 93, 94},
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppListeners(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppPortTypeListeners(t *testing.T) {
	deleteArgs := &DeleteAppListenersArgs{
		PortTypeList: []PortTypeModel{
			{
				Port: 80,
				Type: "UDP",
			},
			{
				Port: 80,
				Type: "HTTP",
			},
		},
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppListeners(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteLoadBalancer(t *testing.T) {
	err := APPBLB_CLIENT.DeleteLoadBalancer(APPBLB_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}
