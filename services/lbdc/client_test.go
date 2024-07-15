package lbdc

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	LDBC_CLIENT *Client

	region string

	LDBCID string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Endpoint string `json:"Endpoint"`
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	// Get the directory of GOPATH, the config file should be under the directory.
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

	LDBC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	region = confObj.Endpoint[4:6]

	log.SetLogHandler(log.FILE)
	log.SetLogDir("/Users/xxx/go/log/baidu/bce/go-sdk")
	log.SetRotateType(log.ROTATE_SIZE)
	log.Info("this is my own logger from the sdk")
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

func getClientToken() string {
	return util.NewUUID()
}

func TestClient_CreateLbdc(t *testing.T) {
	description := ""
	args := &CreateLbdcArgs{
		ClientToken: getClientToken(),
		Name:        "abc",
		Type:        "7Layer",
		CcuCount:    2,
		Description: &description,
		Billing: &Billing{
			PaymentTiming: "Prepaid",
			Reservation: &Reservation{
				ReservationLength: 1,
			},
		},
		RenewReservation: &Reservation{
			ReservationLength: 1,
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",
				TagValue: "tagValue",
			},
		},
	}
	res, err := LDBC_CLIENT.CreateLbdc(args)
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(res)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}

func TestClient_UpgradeLbdc(t *testing.T) {
	args := &UpgradeLbdcArgs{
		ClientToken: getClientToken(),
		Id:          "bgw_group-81196491",
		CcuCount:    4,
	}
	err := LDBC_CLIENT.UpgradeLbdc(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RenewLbdc(t *testing.T) {
	args := &RenewLbdcArgs{
		ClientToken: getClientToken(),
		Id:          "bgw_group-81196491",
		Billing: &BillingForRenew{
			Reservation: &Reservation{
				ReservationLength: 1,
			},
		},
	}
	err := LDBC_CLIENT.RenewLbdc(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListLbdc(t *testing.T) {
	args := &ListLbdcArgs{
		//Id: "bgw_group-a6dd5dc2",
		//Name: "abc",
	}
	res, err := LDBC_CLIENT.ListLbdc(args)
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(res)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}

func TestClient_GetLbdcDetail(t *testing.T) {
	lbdcId := "bgw_group-81196491"
	//lbdcId := "nginx_group-39d9d255"
	res, err := LDBC_CLIENT.GetLbdcDetail(lbdcId)
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(res)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}

func TestClient_UpdateLbdc(t *testing.T) {
	name := ""
	desc := "test"
	args := &UpdateLbdcArgs{
		ClientToken: getClientToken(),
		Id:          "bgw_group-81196491",
		UpdateLbdcBody: &UpdateLbdcBody{
			Name:        &name,
			Description: &desc,
		},
	}
	err := LDBC_CLIENT.UpdateLbdc(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetBoundBlBListOfLbdc(t *testing.T) {
	lbdcId := "bgw_group-81196491"
	res, err := LDBC_CLIENT.GetBoundBlBListOfLbdc(lbdcId)
	ExpectEqual(t.Errorf, nil, err)
	e, err1 := json.Marshal(res)
	if err1 != nil {
		log.Error("json format result error")
	}
	log.Info(string(e))
}
