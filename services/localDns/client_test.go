package localDns

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
	LD_CLIENT *Client

	// set these values before start test
	VPC_ID    = "vpc-0n1hhh8759b0"
	Region    = "bj"
	RR        = "rr"
	TYPE      = "A"
	VALUE     = "1.2.3.5"
	ZONE_ID   = "zone-ky94ixins803"
	RECORD_ID = "rc-9mfacvjk6weq"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

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

	LD_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreatePrivateZone(t *testing.T) {
	createArgs := &CreatePrivateZoneRequest{
		ClientToken: getClientToken(),
		ZoneName:    "sdkLd.com",
	}
	createResult, err := LD_CLIENT.CreatePrivateZone(createArgs)
	ExpectEqual(t.Errorf, nil, err)

	ZONE_ID = createResult.ZoneId
}

func TestClient_DeletePrivateZone(t *testing.T) {
	err := LD_CLIENT.DeletePrivateZone(ZONE_ID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListPrivateZone(t *testing.T) {
	listPrivateZoneRequest := &ListPrivateZoneRequest{}
	res, err := LD_CLIENT.ListPrivateZone(listPrivateZoneRequest)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetPrivateZone(t *testing.T) {
	res, err := LD_CLIENT.GetPrivateZone(ZONE_ID)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindVpc(t *testing.T) {
	bindVpcRequest := &BindVpcRequest{
		Region: Region,
		VpcIds: []string{VPC_ID},
	}
	err := LD_CLIENT.BindVpc(ZONE_ID, bindVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnbindVpc(t *testing.T) {
	unbindVpcRequest := &UnbindVpcRequest{
		Region: Region,
		VpcIds: []string{VPC_ID},
	}
	err := LD_CLIENT.UnbindVpc(ZONE_ID, unbindVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AddRecord(t *testing.T) {
	addRecordRequest := &AddRecordRequest{
		Rr:    RR,
		Type:  TYPE,
		Value: VALUE,
	}
	res, err := LD_CLIENT.AddRecord(ZONE_ID, addRecordRequest)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateRecord(t *testing.T) {
	updateRecordRequest := &UpdateRecordRequest{
		Rr:    RR,
		Type:  TYPE,
		Value: VALUE,
	}
	err := LD_CLIENT.UpdateRecord(RECORD_ID, updateRecordRequest)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteRecord(t *testing.T) {
	err := LD_CLIENT.DeleteRecord(RECORD_ID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListRecord(t *testing.T) {
	listRecordRequest := &ListRecordRequest{
		Marker:  "rc-14ifmguzv5j2",
		MaxKeys: 1000,
	}
	res, err := LD_CLIENT.ListRecord(ZONE_ID, listRecordRequest)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_EnableRecord(t *testing.T) {
	err := LD_CLIENT.EnableRecord(RECORD_ID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DisableRecord(t *testing.T) {
	err := LD_CLIENT.DisableRecord(RECORD_ID, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}
