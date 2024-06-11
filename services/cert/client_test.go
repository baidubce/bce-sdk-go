package cert

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CERT_CLIENT *Client
	CERT_ID     string

	// set these values before start test
	testCertServerData  = ``
	testCertPrivateData = ``

	testUpdateCertServerData  = ``
	testUpdateCertPrivateData = ``
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
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

	CERT_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateCert(t *testing.T) {
	args := &CreateCertArgs{
		CertName:        "sdkCreateTest",
		CertServerData:  testCertServerData,
		CertPrivateData: testCertPrivateData,
	}
	createResult, err := CERT_CLIENT.CreateCert(args)
	ExpectEqual(t.Errorf, nil, err)

	CERT_ID = createResult.CertId
}

func TestClient_ListCerts(t *testing.T) {
	_, err := CERT_CLIENT.ListCerts()
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListCertDetail(t *testing.T) {
	_, err := CERT_CLIENT.ListCertDetail()
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetCertMeta(t *testing.T) {
	result, err := CERT_CLIENT.GetCertMeta(CERT_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "sdkCreateTest", result.CertName)
	ExpectEqual(t.Errorf, 1, result.CertType)
}

func TestClient_GetCertDetail(t *testing.T) {
	result, err := CERT_CLIENT.GetCertDetail(CERT_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, result.CertType)
}

func TestClient_UpdateCertName(t *testing.T) {
	args := &UpdateCertNameArgs{
		CertName: "sdkUpdateNameTest",
	}
	err := CERT_CLIENT.UpdateCertName(CERT_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateCertData(t *testing.T) {
	args := &UpdateCertDataArgs{
		CertName:        "sdkTest",
		CertServerData:  testUpdateCertServerData,
		CertPrivateData: testUpdateCertPrivateData,
	}
	err := CERT_CLIENT.UpdateCertData(CERT_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteCert(t *testing.T) {
	err := CERT_CLIENT.DeleteCert(CERT_ID)
	ExpectEqual(t.Errorf, nil, err)
}
