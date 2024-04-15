package eccr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CCR_CLIENT      *Client
	CCR_INSTANCE_ID string
	CCR_REGISTRY_ID string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fmt.Println(conf)
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	if err := decoder.Decode(confObj); err != nil {
		log.Fatal("decode config obj err:", err)
		os.Exit(1)
	}

	log.SetLogLevel(log.WARN)

	CCR_CLIENT, err = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Setup Complete")
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

func TestClient_ListInstances(t *testing.T) {
	args := &ListInstancesArgs{
		KeywordType: "clusterName",
		Keyword:     "",
		PageNo:      1,
		PageSize:    10,
	}
	resp, err := CCR_CLIENT.ListInstances(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetInstanceDetail(t *testing.T) {

	resp, err := CCR_CLIENT.GetInstanceDetail(CCR_INSTANCE_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_CreateInstance(t *testing.T) {

	billing := Billing{
		ReservationTimeUnit: "MONTH",
		ReservationTime:     1,
		AutoRenew:           false,
		AutoRenewTimeUnit:   "MONTH",
		AutoRenewTime:       1,
	}

	args := &CreateInstanceArgs{
		Type:          "BASIC",
		Name:          "instanceName",
		Bucket:        "",
		PaymentTiming: "prepay",
		Billing:       billing,
		PaymentMethod: []PaymentMethod{},
		Tags:          []Tag{},
	}

	resp, err := CCR_CLIENT.CreateInstance(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_RenewInstance(t *testing.T) {

	orderType := "RENEW"
	args := &RenewInstanceArgs{
		Items:         []Item{},
		PaymentMethod: []PaymentMethod{},
	}

	resp, err := CCR_CLIENT.RenewInstance(orderType, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_UpdateInstance(t *testing.T) {

	args := &UpdateInstanceArgs{
		Name: "InstanceName",
	}

	resp, err := CCR_CLIENT.UpdateInstance(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")

	fmt.Println("Response:", string(s))
}

func TestClient_UpgradeInstance(t *testing.T) {

	args := &UpgradeInstanceArgs{
		Type:          "STANDARD",
		PaymentMethod: []PaymentMethod{},
	}

	resp, err := CCR_CLIENT.UpgradeInstance(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_ListPrivateNetworks(t *testing.T) {

	resp, err := CCR_CLIENT.ListPrivateNetworks(CCR_INSTANCE_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_CreatePrivateNetwork(t *testing.T) {

	args := &CreatePrivateNetworkArgs{
		VpcID:          "VpcID",
		SubnetID:       "SubnetID",
		IPAddress:      "",
		IPType:         "",
		AutoDNSResolve: false,
	}

	resp, err := CCR_CLIENT.CreatePrivateNetwork(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_DeletePrivateNetwork(t *testing.T) {

	args := &DeletePrivateNetworkArgs{
		VpcID:    "VpcID",
		SubnetID: "SubnetID",
	}

	err := CCR_CLIENT.DeletePrivateNetwork(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	fmt.Println("Delete Private Network Test Passed")
}

func TestClient_ListPublicNetworks(t *testing.T) {

	resp, err := CCR_CLIENT.ListPublicNetworks(CCR_INSTANCE_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_UpdatePublicNetwork(t *testing.T) {
	args := &UpdatePublicNetworkArgs{
		Action: "open",
	}

	err := CCR_CLIENT.UpdatePublicNetwork(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	fmt.Println("Update Public Network Test Passed")

}

func TestClient_DeletePublicNetworkWhitelist(t *testing.T) {
	args := &DeletePublicNetworkWhitelistArgs{
		Items: []string{"PublicNetworkWhiteListArgs"},
	}

	err := CCR_CLIENT.DeletePublicNetworkWhitelist(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	fmt.Println("Delete Public Network Whitelist Test Passed")
}

func TestClient_AddPublicNetworkWhitelist(t *testing.T) {
	args := &AddPublicNetworkWhitelistArgs{
		IPAddr:      "cidrv4",
		Description: "",
	}

	err := CCR_CLIENT.AddPublicNetworkWhitelist(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	fmt.Println("Add Public Network Whitelist Test Passed")
}

func TestClient_ResetPassword(t *testing.T) {
	args := &ResetPasswordArgs{
		Password: "Password",
	}

	resp, err := CCR_CLIENT.ResetPassword(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_CreateTemporaryToken(t *testing.T) {
	args := &CreateTemporaryTokenArgs{
		Duration: 10,
	}

	resp, err := CCR_CLIENT.CreateTemporaryToken(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_CreateRegistry(t *testing.T) {
	args := &CreateRegistryArgs{
		Credential:  &RegistryCredential{},
		Description: "",
		Insecure:    false,
		Name:        "",
		Type:        "harbor",
		URL:         "https://baidu.com",
	}

	resp, err := CCR_CLIENT.CreateRegistry(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_GetRegistryDetail(t *testing.T) {

	resp, err := CCR_CLIENT.GetRegistryDetail(CCR_INSTANCE_ID, CCR_REGISTRY_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_ListRegistries(t *testing.T) {
	args := &ListRegistriesArgs{
		RegistryName: "",
		RegistryType: "",
		PageNo:       1,
		PageSize:     10,
	}

	resp, err := CCR_CLIENT.ListRegistries(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.Marshal(resp)
	fmt.Println("Response:" + string(s))
}

func TestClient_CheckHealthRegistry(t *testing.T) {
	args := &RegistryRequestArgs{
		Credential:  &RegistryCredential{},
		Description: "",
		Insecure:    false,
		Name:        "",
		Type:        "harbor",
		URL:         "https://registry.baidubce.com",
	}

	err := CCR_CLIENT.CheckHealthRegistry(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("Check Health Registry Test Passed")
}

func TestClient_UpdateRegistry(t *testing.T) {
	args := &RegistryRequestArgs{
		Credential:  &RegistryCredential{},
		Description: "",
		Insecure:    true,
		Name:        "",
		Type:        "harbor",
		URL:         "https://registry.baidubce.com",
	}

	resp, err := CCR_CLIENT.UpdateRegistry(CCR_INSTANCE_ID, CCR_REGISTRY_ID, args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_DeleteRegistry(t *testing.T) {

	err := CCR_CLIENT.DeleteRegistry(CCR_INSTANCE_ID, CCR_REGISTRY_ID)

	ExpectEqual(t.Errorf, nil, err)

	fmt.Println("Delete Registry Test Passed")
}

func TestClient_CreateBuildRepositoryTask(t *testing.T) {
	args := &BuildRepositoryTaskArgs{
		TagName:    "v1.0",
		IsLatest:   false,
		Dockerfile: "from busybox \n yum install pip",
		FromType:   "dcokerfile",
	}

	resp, err := CCR_CLIENT.CreateBuildRepositoryTask(CCR_INSTANCE_ID, "test", "pip", args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_ListBuildRepositoryTaskArgs(t *testing.T) {
	args := &ListBuildRepositoryTaskArgs{
		KeywordType: "tag",
		Keyword:     "v1.0",
		PageNo:      1,
		PageSize:    10,
	}

	resp, err := CCR_CLIENT.ListBuildRepositoryTask(CCR_INSTANCE_ID, "test", "pip", args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_GetBuildRepositoryTask(t *testing.T) {
	resp, err := CCR_CLIENT.GetBuildRepositoryTask(CCR_INSTANCE_ID, "test", "pip", "1")

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:", string(s))
}

func TestClient_BatchDeleteBuildRepositoryTask(t *testing.T) {
	args := &BatchDeleteBuildRepositoryTaskArgs{
		Items: []string{"1", "2"},
	}
	err := CCR_CLIENT.BatchDeleteBuildRepositoryTask(CCR_INSTANCE_ID, "test", "pip", args)

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AssignInstanceTag(t *testing.T) {
	args := &AssignTagsRequest{
		Tags: []Tag{
			{
				TagKey:   "key1",
				TagValue: "value1",
			},
		},
	}
	err := CCR_CLIENT.AssignInstanceTag(CCR_INSTANCE_ID, args)

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteBuildRepositoryTask(t *testing.T) {
	err := CCR_CLIENT.DeleteBuildRepositoryTask(CCR_INSTANCE_ID, "test", "pip", "1")

	ExpectEqual(t.Errorf, nil, err)
}

func Test_getImageBuildURI(t *testing.T) {
	type args struct {
		instanceID     string
		projectName    string
		repositoryName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "image build uri",
			args: args{
				instanceID:     "test_instance",
				projectName:    "test_project",
				repositoryName: "test_repository",
			},
			want: fmt.Sprintf(`/v1/projects/%s/repositories/%s/imageBuilds/%s`, "test_instance", "test_project", "test_repository"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImageBuildURI(tt.args.instanceID, tt.args.projectName, tt.args.repositoryName); got != tt.want {
				t.Errorf("getImageBuildURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
