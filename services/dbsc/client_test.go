package dbsc

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	DBSC_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {

	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 6; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "../config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	DBSC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
}

func TestCreateVolumeCluster(t *testing.T) {
	args := &CreateVolumeClusterArgs{
		PurchaseCount:   1,
		ClusterSizeInGB: 97280,
		ClusterName:     "dbsc",
		StorageType:     StorageTypeHdd,
		Billing: &Billing{
			PaymentTiming: PaymentTimingPrePaid,
			Reservation: &Reservation{
				ReservationLength:   6,
				ReservationTimeUnit: "MONTH",
			},
		},
		RenewTimeUnit: "month",
		RenewTime:     6,
	}
	result, err := DBSC_CLIENT.CreateVolumeCluster(args)
	if err != nil {
		fmt.Println(err)
	}
	clusterId := result.ClusterIds[0]
	fmt.Println(clusterId)
	if result.ClusterUuids != nil {
		clusterUuid := result.ClusterUuids[0]
		fmt.Println(clusterUuid)
	}
}

func TestListVolumeCluster(t *testing.T) {
	args := &ListVolumeClusterArgs{}
	result, err := DBSC_CLIENT.ListVolumeCluster(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestGetVolumeClusterDetail(t *testing.T) {
	result, err := DBSC_CLIENT.GetVolumeClusterDetail("DC-xxxxxx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestResizeVolumeCluster(t *testing.T) {
	args := &ResizeVolumeClusterArgs{
		NewClusterSizeInGB: 107520,
	}
	err := DBSC_CLIENT.ResizeVolumeCluster("DC-yWfhpUbN", args)
	if err != nil {
		fmt.Println(err)
	}
}

func TestPurchaseReservedVolumeCluster(t *testing.T) {
	args := &PurchaseReservedVolumeClusterArgs{
		Billing: &Billing{
			PaymentTiming: PaymentTimingPrePaid,
			Reservation: &Reservation{
				ReservationLength:   6,
				ReservationTimeUnit: "month",
			},
		},
	}
	err := DBSC_CLIENT.PurchaseReservedVolumeCluster("DC-yWfhpUbN", args)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAutoRenewVolumeCluster(t *testing.T) {
	args := &AutoRenewVolumeClusterArgs{
		ClusterId:     "DC-yWfhpUbN",
		RenewTime:     6,
		RenewTimeUnit: "month",
	}
	err := DBSC_CLIENT.AutoRenewVolumeCluster(args)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCancelAutoRenewVolumeCluster(t *testing.T) {
	args := &CancelAutoRenewVolumeClusterArgs{
		ClusterId: "DC-yWfhpUbN",
	}
	err := DBSC_CLIENT.CancelAutoRenewVolumeCluster(args)
	if err != nil {
		fmt.Println(err)
	}
}
