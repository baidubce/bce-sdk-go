package eip

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
	EIP_CLIENT  *Client
	EIP_ADDRESS string
	CLUSTER_ID  string
	// set this value before start test
	BCC_TEST_ID = ""
	EIP_TP_ID   = ""
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

	EIP_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateEip(t *testing.T) {
	args := &CreateEipArgs{
		Name:            "sdk-eip",
		BandWidthInMbps: 1,
		Billing: &Billing{
			PaymentTiming: "Postpaid",
			BillingMethod: "ByTraffic",
		},
		ClientToken: getClientToken(),
	}
	result, err := EIP_CLIENT.CreateEip(args)
	ExpectEqual(t.Errorf, nil, err)

	EIP_ADDRESS = result.Eip
}

func TestClient_BatchCreateEip(t *testing.T) {
	args := &BatchCreateEipArgs{
		Name:            "sdk-eip",
		BandWidthInMbps: 1,
		Billing: &Billing{
			PaymentTiming: "Postpaid",
			BillingMethod: "ByBandwidth",
		},
		Count:     254,
		RouteType: "BGP",
		//Idc:         "bjdd_cu",
		Continuous:  true,
		ClientToken: getClientToken(),
	}
	result, err := EIP_CLIENT.BatchCreateEip(args)
	if err != nil {
		fmt.Println(err)
	} else {
		r, _ := json.Marshal(result)
		fmt.Println(string(r))
	}
}

func TestClient_ResizeEip(t *testing.T) {
	args := &ResizeEipArgs{
		NewBandWidthInMbps: 2,
		ClientToken:        getClientToken(),
	}
	err := EIP_CLIENT.ResizeEip(EIP_ADDRESS, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindEip(t *testing.T) {
	args := &BindEipArgs{
		InstanceType: "BCC",
		InstanceId:   BCC_TEST_ID,
		ClientToken:  getClientToken(),
	}
	err := EIP_CLIENT.BindEip(EIP_ADDRESS, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UnBindEip(t *testing.T) {
	err := EIP_CLIENT.UnBindEip(EIP_ADDRESS, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_PurchaseReservedEip(t *testing.T) {
	args := &PurchaseReservedEipArgs{
		Billing: &Billing{
			Reservation: &Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
	}
	err := EIP_CLIENT.PurchaseReservedEip(EIP_ADDRESS, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListEip(t *testing.T) {
	args := &ListEipArgs{}
	EIP_CLIENT.ListEip(args)
}

func TestClient_DeleteEip(t *testing.T) {
	err := EIP_CLIENT.DeleteEip(EIP_ADDRESS, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_OptionalDeleteEip(t *testing.T) {
	err := EIP_CLIENT.OptionalDeleteEip(EIP_ADDRESS, getClientToken(), false)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListEipCluster(t *testing.T) {
	args := &ListEipArgs{}
	result, err := EIP_CLIENT.ListEipCluster(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.ClusterList {
		ExpectEqual(t.Errorf, "zone-A|zone-C", e.ClusterAz)
		ExpectEqual(t.Errorf, "c-76a34e7b", e.ClusterId)
		ExpectEqual(t.Errorf, "su-eip-pdd", e.ClusterName)
		ExpectEqual(t.Errorf, "su", e.ClusterRegion)
	}
}

func TestClient_GetEipCluster(t *testing.T) {
	result, err := EIP_CLIENT.GetEipCluster(CLUSTER_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "zone-A|zone-C", result.ClusterAz)
	ExpectEqual(t.Errorf, "c-76a34e7b", result.ClusterId)
	ExpectEqual(t.Errorf, "su-eip-pdd", result.ClusterName)
	ExpectEqual(t.Errorf, "su", result.ClusterRegion)
	ExpectEqual(t.Errorf, 240000000000, result.NetworkInBps)
	ExpectEqual(t.Errorf, 240000000000, result.NetworkOutBps)
	ExpectEqual(t.Errorf, 48000000, result.NetworkInPps)
	ExpectEqual(t.Errorf, 48000000, result.NetworkOutPps)
}

func TestClient_DirectEip(t *testing.T) {
	EIP_CLIENT.DirectEip(EIP_ADDRESS, getClientToken())
}

func TestClient_UnDirectEip(t *testing.T) {
	EIP_CLIENT.UnDirectEip(EIP_ADDRESS, getClientToken())
}

func TestClient_ListRecycleEip(t *testing.T) {
	args := &ListRecycleEipArgs{
		Marker:  "",
		MaxKeys: 1,
	}
	result, err := EIP_CLIENT.ListRecycleEip(args)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(result)
}

func TestClient_RestoreRecycleEip(t *testing.T) {
	err := EIP_CLIENT.RestoreRecycleEip(EIP_ADDRESS, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteRecycleEip(t *testing.T) {
	err := EIP_CLIENT.DeleteRecycleEip(EIP_ADDRESS, getClientToken())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateEipTp(t *testing.T) {
	args := &CreateEipTpArgs{
		ReservationLength: 1,
		Capacity:          "10G",
		DeductPolicy:      "TimeDurationPackage",
		PackageType:       "WebOutBytes",
		ClientToken:       getClientToken(),
	}
	result, err := EIP_CLIENT.CreateEipTp(args)
	ExpectEqual(t.Errorf, nil, err)
	EIP_TP_ID = result.Id
}

func TestClient_ListEipTp(t *testing.T) {
	args := &ListEipTpArgs{
		Id:           "tp-R4E8cWijbf",
		DeductPolicy: "FullTimeDurationPackage",
		Status:       "RUNNING",
	}
	fmt.Println(EIP_CLIENT.ListEipTp(args))
}

func TestClient_GetEipTp(t *testing.T) {
	fmt.Println(EIP_CLIENT.GetEipTp("tp-hfI9pKIL58"))
}

func TestClient_CreateEipGroup(t *testing.T) {
	args := &CreateEipGroupArgs{
		Name:            "sdk-eipGroup",
		BandWidthInMbps: 500,
		Billing: &Billing{
			PaymentTiming: "Postpaid",
			BillingMethod: "ByBandwidth",
		},
		EipCount:    254,
		RouteType:   "BGP",
		Continuous:  true,
		ClientToken: getClientToken(),
	}
	result, err := EIP_CLIENT.CreateEipGroup(args)
	if err != nil {
		fmt.Println(err)
	} else {
		r, _ := json.Marshal(result)
		fmt.Println(string(r))
	}
}

func TestClient_ResizeEipGroupBandWidth(t *testing.T) {
	args := &ResizeEipGroupArgs{
		BandWidthInMbps: 22,
		ClientToken:     getClientToken(),
	}
	EIP_CLIENT.ResizeEipGroupBandWidth("eg-2b1ef0db", args)
}

func TestClient_EipGroupAddEipCount(t *testing.T) {
	args := &GroupAddEipCountArgs{
		EipAddCount: 1,
		ClientToken: getClientToken(),
	}
	EIP_CLIENT.EipGroupAddEipCount("eg-2b1ef0db", args)
}

func TestClient_ReleaseEipGroupIps(t *testing.T) {
	Eips := []string{"100.88.3.216"}
	args := &ReleaseEipGroupIpsArgs{
		ReleaseIps:  Eips,
		ClientToken: getClientToken(),
	}
	EIP_CLIENT.ReleaseEipGroupIps("eg-2b1ef0db", args)
}

func TestClient_RenameEipGroup(t *testing.T) {
	args := &RenameEipGroupArgs{
		Name:        "new_SDK_111",
		ClientToken: getClientToken(),
	}
	EIP_CLIENT.RenameEipGroup("eg-2b1ef0db", args)
}

func TestClient_ListEipGroup(t *testing.T) {
	result, err := EIP_CLIENT.ListEipGroup(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		r, _ := json.Marshal(result)
		fmt.Println(string(r))
	}
}

func TestClient_EipGroupDetail(t *testing.T) {
	result, err := EIP_CLIENT.EipGroupDetail("eg-2b1ef0db")
	if err != nil {
		fmt.Println(err)
	} else {
		r, _ := json.Marshal(result)
		fmt.Println(string(r))
	}
}

func TestClient_EipGroupMoveIn(t *testing.T) {
	Ips := []string{"100.88.1.231"}
	args := &EipGroupMoveInArgs{
		Eips:        Ips,
		ClientToken: getClientToken(),
	}
	EIP_CLIENT.EipGroupMoveIn("eg-2b1ef0db", args)
}

func TestClient_EipGroupMoveOut(t *testing.T) {
	args := &EipGroupMoveOutArgs{
		MoveOutEips: []MoveOutEip{
			{
				Eip:             "100.88.11.128",
				BandWidthInMbps: 11,
				Billing: &Billing{
					PaymentTiming: "Postpaid",
					BillingMethod: "ByBandwidth",
				},
			},
		},
		ClientToken: getClientToken(),
	}
	EIP_CLIENT.EipGroupMoveOut("eg-2b1ef0db", args)
}

func getClientToken() string {
	return util.NewUUID()
}
