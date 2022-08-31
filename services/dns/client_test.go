package dns

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
	DNS_CLIENT *Client

	// set these values before start test
	Region         = "bj"
	RR             = "rr"
	TYPE           = "A"
	VALUE          = "1.2.3.5"
	ZONE_NAME      = "ccq.com"
	ZONE_NAME1     = "sdkdns.com"
	PREPAID        = "Prepaid"
	ProductVersion = "discount"
	RECORD_ID      = "48526"
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

	DNS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateZone(t *testing.T) {
	createArgs := &CreateZoneRequest{
		Name: ZONE_NAME1,
	}
	err := DNS_CLIENT.CreateZone(createArgs, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListZone(t *testing.T) {
	listZoneRequest := &ListZoneRequest{
		Name: ZONE_NAME,
	}
	res, err := DNS_CLIENT.ListZone(listZoneRequest)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteZone(t *testing.T) {
	err := DNS_CLIENT.DeleteZone(ZONE_NAME, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreatePaidZone(t *testing.T) {
	createArgs := &CreatePaidZoneRequest{
		Names:          []string{ZONE_NAME},
		ProductVersion: ProductVersion,
		Billing: Billing{
			PaymentTiming: PREPAID,
			Reservation: Reservation{
				ReservationLength: 1,
			},
		},
	}
	err := DNS_CLIENT.CreatePaidZone(createArgs, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpgradeZone(t *testing.T) {
	createArgs := &UpgradeZoneRequest{
		Names: []string{ZONE_NAME1},
		Billing: Billing{
			PaymentTiming: PREPAID,
			Reservation: Reservation{
				ReservationLength: 1,
			},
		},
	}
	err := DNS_CLIENT.UpgradeZone(createArgs, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RenewZone(t *testing.T) {
	createArgs := &RenewZoneRequest{
		Billing: Billing{
			Reservation: Reservation{
				ReservationLength: 1,
			},
		},
	}
	err := DNS_CLIENT.RenewZone(ZONE_NAME1, createArgs, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateRecord(t *testing.T) {
	createRecordRequest := &CreateRecordRequest{
		Rr:    RR,
		Type:  TYPE,
		Value: VALUE,
	}
	err := DNS_CLIENT.CreateRecord(ZONE_NAME, createRecordRequest, "")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListRecord(t *testing.T) {
	listRecordRequest := &ListRecordRequest{}
	res, err := DNS_CLIENT.ListRecord(ZONE_NAME, listRecordRequest)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateRecord(t *testing.T) {
	updateRecordRequest := &UpdateRecordRequest{
		Rr:    RR,
		Type:  TYPE,
		Value: "1.1.1.1",
	}
	err := DNS_CLIENT.UpdateRecord(ZONE_NAME, RECORD_ID, updateRecordRequest, "")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateRecordEnable(t *testing.T) {
	err := DNS_CLIENT.UpdateRecordEnable(ZONE_NAME, "48540", "")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateRecordDisable(t *testing.T) {
	err := DNS_CLIENT.UpdateRecordDisable(ZONE_NAME, "48540", "")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteRecord(t *testing.T) {
	err := DNS_CLIENT.DeleteRecord(ZONE_NAME, "48540", "")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateLineGroup(t *testing.T) {
	addLineGroupRequest := &AddLineGroupRequest{
		Name:  "ccq0826",
		Lines: []string{"yunnan.ct"},
	}
	err := DNS_CLIENT.AddLineGroup(addLineGroupRequest, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateLineGroup(t *testing.T) {
	updateLineGroupRequest := &UpdateLineGroupRequest{
		Name:  "ccq0826_1",
		Lines: []string{"india.any"},
	}
	err := DNS_CLIENT.UpdateLineGroup("6165", updateLineGroupRequest, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListLineGroup(t *testing.T) {
	listLineGroupRequest := &ListLineGroupRequest{}
	res, err := DNS_CLIENT.ListLineGroup(listLineGroupRequest)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteLineGroup(t *testing.T) {
	err := DNS_CLIENT.DeleteLineGroup("6039", getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}
