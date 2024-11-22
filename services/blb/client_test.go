package blb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BLB_CLIENT *Client
	BLB_ID     string

	// set these values before start test
	VPC_TEST_ID    = ""
	SUBNET_TEST_ID = ""
	INSTANCE_ID    = ""
	CERT_ID        = ""
	CLUSTER_ID     = ""
	TEST_BLB_ID    = ""
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

	BLB_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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
	AllowDelete := false
	AllocateIpv6 := true
	Layer4ClusterExclusive := false
	createArgs := &CreateLoadBalancerArgs{
		ClientToken:            getClientToken(),
		Name:                   "sdkBlb051103",
		VpcId:                  "vpc-rfujg1q3zqqf",
		SubnetId:               "sbn-sx0zf4qbcbmp",
		AllowDelete:            &AllowDelete,
		AllocateIpv6:           &AllocateIpv6,
		Layer4ClusterExclusive: &Layer4ClusterExclusive,
	}

	createResult, err := BLB_CLIENT.CreateLoadBalancer(createArgs)
	ExpectEqual(t.Errorf, nil, err)

	BLB_ID = createResult.BlbId
}

func TestClient_CreatePrePayLoadBalancer(t *testing.T) {
	AllowDelete := false
	AllocateIpv6 := true
	createArgs := &CreateLoadBalancerArgs{
		ClientToken:     getClientToken(),
		Name:            "sdkBlb051101",
		VpcId:           "vpc-rfujg1q3zqqf",
		SubnetId:        "sbn-sx0zf4qbcbmp",
		AutoRenewLength: 1,
		Billing: &Billing{
			PaymentTiming: "Prepaid",
			BillingMethod: "BySpec",
			Reservation: &Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "Month",
			},
		},
		PerformanceLevel: "small1",
		AllowDelete:      &AllowDelete,
		AllocateIpv6:     &AllocateIpv6,
	}

	createResult, err := BLB_CLIENT.CreateLoadBalancer(createArgs)
	ExpectEqual(t.Errorf, nil, err)

	BLB_ID = createResult.BlbId
}

func TestClient_UpdateLoadBalancer(t *testing.T) {
	updateArgs := &UpdateLoadBalancerArgs{
		Name:        "testSdk",
		Description: "test desc",
	}
	err := BLB_CLIENT.UpdateLoadBalancer(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeLoadBalancers(t *testing.T) {
	describeArgs := &DescribeLoadBalancersArgs{}
	res, err := BLB_CLIENT.DescribeLoadBalancers(describeArgs)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(time.Duration(1) * time.Second)
}

func TestClient_DescribeLoadBalancerDetail(t *testing.T) {
	res, err := BLB_CLIENT.DescribeLoadBalancerDetail(BLB_ID)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateTCPListener(t *testing.T) {
	createArgs := &CreateTCPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 90,
		BackendPort:  90,
		Scheduler:    "RoundRobin",
	}
	err := BLB_CLIENT.CreateTCPListener(BLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateTCPListener(t *testing.T) {
	updateArgs := &UpdateTCPListenerArgs{
		ListenerPort: 90,
		Scheduler:    "Hash",
	}
	err := BLB_CLIENT.UpdateTCPListener(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeTCPListeners(t *testing.T) {
	describeArgs := &DescribeListenerArgs{
		ListenerPort: 90,
	}
	_, err := BLB_CLIENT.DescribeTCPListeners(BLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateUDPListener(t *testing.T) {
	createArgs := &CreateUDPListenerArgs{
		ClientToken:       getClientToken(),
		ListenerPort:      91,
		BackendPort:       91,
		Scheduler:         "RoundRobin",
		HealthCheckString: "a",
	}
	err := BLB_CLIENT.CreateUDPListener(BLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateUDPListener(t *testing.T) {
	updateArgs := &UpdateUDPListenerArgs{
		ListenerPort: 91,
		Scheduler:    "Hash",
	}
	err := BLB_CLIENT.UpdateUDPListener(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeUDPListeners(t *testing.T) {
	describeArgs := &DescribeListenerArgs{
		ListenerPort: 91,
	}
	_, err := BLB_CLIENT.DescribeUDPListeners(BLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateHTTPListener(t *testing.T) {
	createArgs := &CreateHTTPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 92,
		BackendPort:  92,
		Scheduler:    "RoundRobin",
	}
	err := BLB_CLIENT.CreateHTTPListener(BLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateHTTPListener(t *testing.T) {
	updateArgs := &UpdateHTTPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 92,
		Scheduler:    "LeastConnection",
		KeepSession:  True,
	}
	err := BLB_CLIENT.UpdateHTTPListener(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeHTTPListeners(t *testing.T) {
	describeArgs := &DescribeListenerArgs{
		ListenerPort: 92,
	}
	_, err := BLB_CLIENT.DescribeHTTPListeners(BLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateHTTPSListener(t *testing.T) {
	createArgs := &CreateHTTPSListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 93,
		BackendPort:  93,
		Scheduler:    "RoundRobin",
		CertIds:      []string{CERT_ID},
	}
	err := BLB_CLIENT.CreateHTTPSListener(BLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateHTTPSListener(t *testing.T) {
	updateArgs := &UpdateHTTPSListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 93,
		Scheduler:    "LeastConnection",
		KeepSession:  True,
		CertIds:      []string{CERT_ID},
	}
	err := BLB_CLIENT.UpdateHTTPSListener(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeHTTPSListeners(t *testing.T) {
	describeArgs := &DescribeListenerArgs{
		ListenerPort: 93,
	}
	_, err := BLB_CLIENT.DescribeHTTPSListeners(BLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateSSLListener(t *testing.T) {
	createArgs := &CreateSSLListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 94,
		BackendPort:  94,
		Scheduler:    "RoundRobin",
		CertIds:      []string{CERT_ID},
	}
	err := BLB_CLIENT.CreateSSLListener(BLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateSSLListener(t *testing.T) {
	updateArgs := &UpdateSSLListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 94,
		Scheduler:    "LeastConnection",
		CertIds:      []string{CERT_ID},
	}
	err := BLB_CLIENT.UpdateSSLListener(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeSSLListeners(t *testing.T) {
	describeArgs := &DescribeListenerArgs{
		ListenerPort: 94,
	}
	_, err := BLB_CLIENT.DescribeSSLListeners(BLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAllListeners(t *testing.T) {
	describeArgs := &DescribeListenerArgs{}
	result, err := BLB_CLIENT.DescribeAllListeners(BLB_ID, describeArgs)
	if err != nil {
		fmt.Println("get all listener failed:", err)
	} else {
		fmt.Println("get all listener success: ", result)
	}
}

func TestClient_AddBackendServers(t *testing.T) {
	createArgs := &AddBackendServersArgs{
		ClientToken: getClientToken(),
		BackendServerList: []BackendServerModel{
			{InstanceId: INSTANCE_ID, Weight: 30},
		},
	}
	err := BLB_CLIENT.AddBackendServers(BLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateBackendServers(t *testing.T) {
	updateArgs := &UpdateBackendServersArgs{
		ClientToken: getClientToken(),
		BackendServerList: []BackendServerModel{
			{InstanceId: INSTANCE_ID, Weight: 50},
		},
	}
	err := BLB_CLIENT.UpdateBackendServers(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeBackendServers(t *testing.T) {
	describeArgs := &DescribeBackendServersArgs{}
	_, err := BLB_CLIENT.DescribeBackendServers(BLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeHealthStatus(t *testing.T) {
	describeArgs := &DescribeHealthStatusArgs{
		ListenerPort: 90,
	}
	_, err := BLB_CLIENT.DescribeHealthStatus(BLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RemoveBackendServers(t *testing.T) {
	deleteArgs := &RemoveBackendServersArgs{
		BackendServerList: []string{INSTANCE_ID},
		ClientToken:       getClientToken(),
	}
	err := BLB_CLIENT.RemoveBackendServers(BLB_ID, deleteArgs)

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteListeners(t *testing.T) {
	deleteArgs := &DeleteListenersArgs{
		PortList:    []uint16{90, 91, 92, 93, 94},
		ClientToken: getClientToken(),
	}
	err := BLB_CLIENT.DeleteListeners(BLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeletePortTypeListeners(t *testing.T) {
	deleteArgs := &DeleteListenersArgs{
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
	err := BLB_CLIENT.DeleteListeners(BLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteLoadBalancer(t *testing.T) {
	err := BLB_CLIENT.DeleteLoadBalancer(BLB_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeLbClusters(t *testing.T) {
	describeArgs := &DescribeLbClustersArgs{}
	res, err := BLB_CLIENT.DescribeLbClusters(describeArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeLbClusterDetail(t *testing.T) {
	res, err := BLB_CLIENT.DescribeLbClusterDetail(CLUSTER_ID)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateLoadBalancerAcl(t *testing.T) {
	supportAcl := new(bool)
	*supportAcl = true
	updateArgs := &UpdateLoadBalancerAclArgs{
		ClientToken: getClientToken(),
		SupportAcl:  supportAcl,
	}
	err := BLB_CLIENT.UpdateLoadBalancerAcl(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindSecurityGroups(t *testing.T) {
	updateArgs := &UpdateSecurityGroupsArgs{
		ClientToken:      getClientToken(),
		SecurityGroupIds: []string{"sg-id"},
	}
	err := BLB_CLIENT.BindSecurityGroups(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnbindSecurityGroups(t *testing.T) {
	updateArgs := &UpdateSecurityGroupsArgs{
		ClientToken:      getClientToken(),
		SecurityGroupIds: []string{"sg-id"},
	}
	err := BLB_CLIENT.UnbindSecurityGroups(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeSecurityGroups(t *testing.T) {
	res, err := BLB_CLIENT.DescribeSecurityGroups(BLB_ID)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindEnterpriseSecurityGroups(t *testing.T) {
	updateArgs := &UpdateEnterpriseSecurityGroupsArgs{
		ClientToken:                getClientToken(),
		EnterpriseSecurityGroupIds: []string{"esg-id"},
	}
	err := BLB_CLIENT.BindEnterpriseSecurityGroups(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnbindEnterpriseSecurityGroups(t *testing.T) {
	updateArgs := &UpdateEnterpriseSecurityGroupsArgs{
		ClientToken:                getClientToken(),
		EnterpriseSecurityGroupIds: []string{"esg-id"},
	}
	err := BLB_CLIENT.UnbindEnterpriseSecurityGroups(BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeEnterpriseSecurityGroups(t *testing.T) {
	res, err := BLB_CLIENT.DescribeEnterpriseSecurityGroups(BLB_ID)
	fmt.Println(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_StartLoadBalancerAutoRenew(t *testing.T) {
	updateArgs := &StartLoadBalancerAutoRenewArgs{
		ClientToken:     getClientToken(),
		AutoRenewLength: 1,
	}
	err := BLB_CLIENT.StartLoadBalancerAutoRenew(TEST_BLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RefundLoadBalancer(t *testing.T) {
	updateArgs := &RefundLoadBalancerArgs{
		ClientToken: getClientToken(),
	}
	err := BLB_CLIENT.RefundLoadBalancer("lb-3d3d4a58", updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}
