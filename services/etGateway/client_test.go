package etGateway

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	EtGateway_CLIENT *Client
	region           string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK          string `json:"AK"`
	SK          string `json:"SK"`
	VPCEndpoint string `json:"VPC"`
	EIPEndpoint string `json:"EIP"`
}

func init() {
	log.SetLogHandler(log.STDERR)
	log.SetLogLevel(log.DEBUG)
	_, f, _, _ := runtime.Caller(0)
	// Get the directory of GOPATH, the config file should be under the directory.
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	EtGateway_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.VPCEndpoint)

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

func TestClient_CreateEtGateway(t *testing.T) {
	args := &CreateEtGatewayArgs{
		Name:        "TestSDK-VPN",
		Description: "vpn test",
		VpcId:       "vpc-2pa2x0bjt26i",
		Speed:       100,
		ClientToken: getClientToken(),
	}
	result, err := EtGateway_CLIENT.CreateEtGateway(args)
	ExpectEqual(t.Errorf, nil, err)
	EtGatewayId := result.EtGatewayId
	log.Debug(EtGatewayId)
}

func TestClient_ListEtGateway(t *testing.T) {
	args := &ListEtGatewayArgs{
		VpcId: "vpc-xsd65rcsp5ue",
	}
	result := &ListEtGatewayResult{}
	result, err := EtGateway_CLIENT.ListEtGateway(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))
}

func TestClient_GetEtGatewayDetail(t *testing.T) {
	res := &EtGatewayDetail{}
	res, err := EtGateway_CLIENT.GetEtGatewayDetail("dcgw-vs1rvp9qy79f")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(res)
	fmt.Println(string(r))
}

func TestClient_UpdateEtGateway(t *testing.T) {
	args := &UpdateEtGatewayArgs{
		ClientToken: getClientToken(),
		EtGatewayId: "dcgw-mx3annmentbu",
		Name:        "aaa",
		Description: "test",
	}
	err := EtGateway_CLIENT.UpdateEtGateway(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteEtGateway(t *testing.T) {
	err := EtGateway_CLIENT.DeleteEtGateway("dcgw-iiyc0ers2qx4", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindEt(t *testing.T) {
	args := &BindEtArgs{
		ClientToken: getClientToken(),
		EtGatewayId: "dcgw-iiyc0ers2qx4",
		EtId:        "et-aaccd",
		ChannelId:   "sdxs",
		LocalCidrs:  []string{"192.168.0.1"},
	}
	err := EtGateway_CLIENT.BindEt(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnBindEt(t *testing.T) {
	err := EtGateway_CLIENT.UnBindEt("dcgw-iiyc0ers2qx4", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateHealthCheck(t *testing.T) {
	auto := true
	args := &CreateHealthCheckArgs{
		ClientToken:           getClientToken(),
		EtGatewayId:           "dcgw-iiyc0ers2qx4",
		HealthCheckSourceIp:   "1.2.3.4",
		HealthCheckType:       HEALTH_CHECK_ICMP,
		HealthCheckPort:       80,
		HealthCheckInterval:   60,
		HealthThreshold:       3,
		UnhealthThreshold:     4,
		AutoGenerateRouteRule: &auto,
	}
	err := EtGateway_CLIENT.CreateHealthCheck(args)
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}
