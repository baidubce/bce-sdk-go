package bes

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BES_CLIENT              *Client
	BES_TestBbcId           string
	BES_TestImageId         string
	BES_TestFlavorId        string
	BES_TestRaidId          string
	BES_TestZoneName        string
	BES_TestSubnetId        string
	BES_TestName            string
	BES_TestAdminPass       string
	BES_TestDeploySetId     string
	BES_TestClientToken     string
	BES_TestSecurityGroupId string
	BES_TestTaskId          string
	BES_TestErrResult       string
	BES_TestRuleId          string
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
	conf := filepath.Join(f, "config.json")
	fmt.Println(conf)
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Println("config json file of ak/sk not given: ", conf)
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	BES_TestFlavorId = "flavor-id"
	BES_TestImageId = "image-id"
	BES_TestRaidId = "raid-id"
	BES_TestZoneName = "zone-name"
	BES_TestSubnetId = "subnet-id"
	BES_TestName = "sdkTest"
	BES_TestAdminPass = "123@adminPass"
	BES_TestDeploySetId = "deployset-id"
	BES_TestBbcId = "bbc_id"
	BES_TestSecurityGroupId = "bbc-security-group-id"
	BES_TestTaskId = "task-id"
	BES_TestErrResult = "err-result"
	BES_TestRuleId = "rule-id"
	BES_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
	//log.SetLogLevel(log.DEBUG)
}

func TestCreateInstance(t *testing.T) {
	InternalIps := []string{"ip"}
	createInstanceArgs := &CreateInstanceArgs{
		FlavorId:         BES_TestFlavorId,
		ImageId:          BES_TestImageId,
		RaidId:           BES_TestRaidId,
		RootDiskSizeInGb: 40,
		PurchaseCount:    1,
		AdminPass:        "AdminPass",
		ZoneName:         BES_TestZoneName,
		SubnetId:         BES_TestSubnetId,
		SecurityGroupId:  BES_TestSecurityGroupId,
		ClientToken:      BES_TestClientToken,
		Billing: Billing{
			PaymentTiming: PaymentTimingPostPaid,
		},
		DeploySetId: BES_TestDeploySetId,
		Name:        BES_TestName,
		EnableNuma:  false,
		InternalIps: InternalIps,
		Tags: []model.TagModel{
			{
				TagKey:   "tag1",
				TagValue: "var1",
			},
		},
	}
	res, err := BES_CLIENT.CreateInstance(createInstanceArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}
