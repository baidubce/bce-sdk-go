package quotacenter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	QUOTA_CENTER_CLIENT *Client
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

	QUOTA_CENTER_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
}

func TestClient_ListProducts(t *testing.T) {
	args := &ProductQueryArgs{
		ProductType: "EIP",
	}
	result, err := QUOTA_CENTER_CLIENT.ListProducts(args)
	fmt.Println(result)
	fmt.Println(err)

}

func TestClient_ListRegions(t *testing.T) {
	args := &RegionQueryArgs{
		ProductType: "EIP",
		ServiceType: "EIP_BP",
		Type:        "QUOTA",
	}
	result, _ := QUOTA_CENTER_CLIENT.ListRegions(args)
	fmt.Println(result)
}

func TestClient_QuotaCenterQuery(t *testing.T) {
	args := &QuotaCenterQueryArgs{
		ServiceType: "EIP",
		Type:        "QUOTA",
		Region:      "su",
	}
	result, err := QUOTA_CENTER_CLIENT.QuotaCenterQuery(args)
	fmt.Println(result)
	fmt.Println(err)
}

func TestClient_InfoQuery(t *testing.T) {
	args := &InfoQueryArgs{
		Region: "sin",
	}
	result, err := QUOTA_CENTER_CLIENT.InfoQuery(args)
	fmt.Println(result)
	fmt.Println(err)
}
