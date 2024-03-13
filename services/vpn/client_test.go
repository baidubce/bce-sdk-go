package vpn

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
	VPN_CLIENT *Client

	region string

	VPNID string
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

	VPN_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.VPCEndpoint)

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

func TestClient_CreateVpnGateway(t *testing.T) {
	args := &CreateVpnGatewayArgs{
		VpnName:     "TestSDK-VPN",
		Description: "vpn test",
		VpcId:       "vpc-2pa2x0bjt26i",
		Billing: &Billing{
			PaymentTiming: PAYMENT_TIMING_PREPAID,
			Reservation: &Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
		ClientToken: getClientToken(),
	}
	result, err := VPN_CLIENT.CreateVpnGateway(args)
	ExpectEqual(t.Errorf, nil, err)
	VPNID := result.VpnId
	log.Debug(VPNID)
}

func TestClient_ListVpnGateway(t *testing.T) {
	args := &ListVpnGatewayArgs{
		MaxKeys: 1000,
		VpcId:   "vpc-2pa2x0bjt26i",
	}
	res, err := VPN_CLIENT.ListVpnGateway(args)
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(res)
	fmt.Println(string(r))

}

func TestClient_BindEip(t *testing.T) {
	args := &BindEipArgs{
		ClientToken: getClientToken(),
		Eip:         "100.88.4.213",
	}
	err := VPN_CLIENT.BindEip("vpn-sd3vxkwisgux", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnBindEip(t *testing.T) {
	err := VPN_CLIENT.UnBindEip("vpn-sd3vxkwisgux", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateVpnGateway(t *testing.T) {
	args := &UpdateVpnGatewayArgs{
		ClientToken: getClientToken(),
		Name:        "vpnTest",
	}
	err := VPN_CLIENT.UpdateVpnGateway("vpn-sd3vxkwisgux", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetVpnGatewayDetail(t *testing.T) {
	result, err := VPN_CLIENT.GetVpnGatewayDetail("vpn-shr6dtz1jjbk")
	ExpectEqual(t.Errorf, nil, err)
	r, err := json.Marshal(result)
	fmt.Println(string(r))

}

func TestClient_RenewVpnGateway(t *testing.T) {
	args := &RenewVpnGatewayArgs{
		ClientToken: getClientToken(),
		Billing: &Billing{
			Reservation: &Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
	}
	err := VPN_CLIENT.RenewVpnGateway("vpn-smt119kcvqb1", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListVpnConn(t *testing.T) {
	res, err := VPN_CLIENT.ListVpnConn("vpn-shr6dtz1jjbk")
	ExpectEqual(t.Errorf, nil, err)
	log.Debug("%+v", res)
	r, _ := json.Marshal(*res)
	fmt.Println(string(r))

}
func TestClient_UpdateVpnConn(t *testing.T) {
	args := &UpdateVpnConnArgs{
		vpnConnId: "vpnconn-mpfwkca8zsuv",
		updateVpnconn: &CreateVpnConnArgs{
			VpnId:         "vpn-shr6dtz1jjbk",
			VpnConnName:   "vpnconntest",
			LocalIp:       "0.1.2.3",
			SecretKey:     "!sdse154d",
			LocalSubnets:  []string{"192.168.0.0/20"},
			RemoteIp:      "3.4.5.6",
			RemoteSubnets: []string{"192.168.100.0/24"},
			CreateIkeConfig: &CreateIkeConfig{
				IkeVersion:  "v1",
				IkeMode:     "main",
				IkeEncAlg:   "aes",
				IkeAuthAlg:  "sha1",
				IkePfs:      "group2",
				IkeLifeTime: 25500,
			},
			CreateIpsecConfig: &CreateIpsecConfig{
				IpsecEncAlg:   "aes",
				IpsecAuthAlg:  "sha1",
				IpsecPfs:      "group2",
				IpsecLifetime: 25500,
			},
		},
	}
	err := VPN_CLIENT.UpdateVpnConn(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateVpnConn(t *testing.T) {
	args := &CreateVpnConnArgs{
		VpnId:         "vpn-shr6dtz1jjbk",
		VpnConnName:   "vpnconntest111",
		LocalIp:       "0.1.2.3",
		SecretKey:     "!sdse154d",
		LocalSubnets:  []string{"192.168.0.0/20"},
		RemoteIp:      "3.4.5.6",
		RemoteSubnets: []string{"192.168.100.0/24"},
		CreateIkeConfig: &CreateIkeConfig{
			IkeVersion:  "v1",
			IkeMode:     "main",
			IkeEncAlg:   "aes",
			IkeAuthAlg:  "sha1",
			IkePfs:      "group2",
			IkeLifeTime: 25500,
		},
		CreateIpsecConfig: &CreateIpsecConfig{
			IpsecEncAlg:   "aes",
			IpsecAuthAlg:  "sha1",
			IpsecPfs:      "group2",
			IpsecLifetime: 25500,
		},
	}
	res, err := VPN_CLIENT.CreateVpnConn(args)
	ExpectEqual(t.Errorf, nil, err)
	log.Debug("%+v", res)
}

func TestClient_DeleteVpnConn(t *testing.T) {
	err := VPN_CLIENT.DeleteVpnConn("vpnconn-mpfwkca8zsuv", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteVpn(t *testing.T) {
	err := VPN_CLIENT.DeleteVpn("vpn-sd3vxkwisgux", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateSslVpnServer(t *testing.T) {
	interfaceTypeStr := "tun"
	ClientDnsStr := ""
	args := &CreateSslVpnServerArgs{
		ClientToken:      getClientToken(),
		VpnId:            "vpn-s5bi05y3yyds", // ssl vpn
		SslVpnServerName: "server_1",
		InterfaceType:    &interfaceTypeStr,
		LocalSubnets:     []string{"192.168.0.0/20"},
		RemoteSubnet:     "192.168.100.0/24",
		ClientDns:        &ClientDnsStr,
	}
	log.Info(args)
	log.Info("args.InterfaceType:", args.InterfaceType)
	log.Info("args.ClientDns:", args.ClientDns)
	if args.InterfaceType == nil {
		log.Info("args.InterfaceType is unassigned")
	}
	res, err := VPN_CLIENT.CreateSslVpnServer(args)
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(res)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}

func TestClient_GetSslVpnServer(t *testing.T) {
	res, err := VPN_CLIENT.GetSslVpnServer("vpn-s5bi05y3yyds", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(res)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}

func TestClient_UpdateSslVpnServer(t *testing.T) {
	clientDnsStr := "100.88.0.83"
	args := &UpdateSslVpnServerArgs{
		ClientToken:    getClientToken(),
		VpnId:          "vpn-s5bi05y3yyds",
		SslVpnServerId: "sslvpn-s8aqk5iw6ki1",
		UpdateSslVpnServer: &UpdateSslVpnServer{
			SslVpnServerName: "SslVpnServer1test twice",
			LocalSubnets:     []string{"192.168.0.0/20"},
			RemoteSubnet:     "192.168.100.0/24",
			ClientDns:        &clientDnsStr,
		},
	}
	err := VPN_CLIENT.UpdateSslVpnServer(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteSslVpnServer(t *testing.T) {
	err := VPN_CLIENT.DeleteSslVpnServer("vpn-s5bi05y3yyds", "sslvpn-1swvxqrzn1we", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BatchCreateSslVpnUser(t *testing.T) {
	desc1 := "user1 description"
	args := &BatchCreateSslVpnUserArgs{
		ClientToken: getClientToken(),
		VpnId:       "vpn-s5bi05y3yyds", // ssl vpn
		SslVpnUsers: []SslVpnUser{
			SslVpnUser{
				UserName:    "user1test",
				Password:    "psd123456!",
				Description: &desc1,
			},
			SslVpnUser{
				UserName: "user2test",
				Password: "psd123456!",
			},
		},
	}
	res, err := VPN_CLIENT.BatchCreateSslVpnUser(args)
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(res)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}

func TestClient_ListSslVpnUser(t *testing.T) {
	args := &ListSslVpnUserArgs{
		MaxKeys: 1000,
		VpnId:   "vpn-s5bi05y3yyds",
	}
	result, err := VPN_CLIENT.ListSslVpnUser(args)
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(result)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}

func TestClient_UpdateSslVpnUser(t *testing.T) {
	psdStr := "123abcd!!!"
	desc := "333"
	args := &UpdateSslVpnUserArgs{
		ClientToken: getClientToken(),
		VpnId:       "vpn-s5bi05y3yyds",
		UserId:      "vpn-ssl-user-ebfysi53dye3",
		SslVpnUser: &UpdateSslVpnUser{
			Password:    psdStr,
			Description: &desc,
		},
	}
	err := VPN_CLIENT.UpdateSslVpnUser(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteSslVpnUser(t *testing.T) {
	err := VPN_CLIENT.DeleteSslVpnUser("vpn-s5bi05y3yyds", "vpn-ssl-user-ggngktunui0k", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}
