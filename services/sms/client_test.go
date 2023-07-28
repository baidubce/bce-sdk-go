package sms

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/sms/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	SMS_CLIENT        *Client
	TEST_SIGNATURE_ID = ""
	TEST_TEMPLATE_ID  = ""
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
	conf := filepath.Join(f, "/config.json")
	fmt.Println(conf)
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	SMS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
	//log.SetLogLevel(log.DEBUG)
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

func TestSendSms(t *testing.T) {
	contentMap := make(map[string]interface{})
	contentMap["code"] = "123"
	contentMap["minute"] = "1"
	sendSmsArgs := &api.SendSmsArgs{
		Mobile:      "13800138000",
		Template:    "your template id",
		SignatureId: "your signature id",
		ContentVar:  contentMap,
	}
	result, err := SMS_CLIENT.SendSms(sendSmsArgs)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", result)
}

func TestCreateSignature(t *testing.T) {
	result, err := SMS_CLIENT.CreateSignature(&api.CreateSignatureArgs{
		Content:     "测试",
		ContentType: "Enterprise",
		Description: "This is a test",
		CountryType: "DOMESTIC",
	})
	ExpectEqual(t.Errorf, err, nil)
	TEST_SIGNATURE_ID = result.SignatureId
	t.Logf("%v", result)
}

func TestGetSignature(t *testing.T) {
	_, err := SMS_CLIENT.GetSignature(&api.GetSignatureArgs{SignatureId: TEST_SIGNATURE_ID})
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifySignature(t *testing.T) {
	err := SMS_CLIENT.ModifySignature(&api.ModifySignatureArgs{
		SignatureId:         TEST_SIGNATURE_ID,
		Content:             "测试变更",
		ContentType:         "Enterprise",
		Description:         "This is a test",
		CountryType:         "INTERNATIONAL",
		SignatureFileBase64: "",
		SignatureFileFormat: "",
	})
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteSignature(t *testing.T) {
	err := SMS_CLIENT.DeleteSignature(&api.DeleteSignatureArgs{SignatureId: TEST_SIGNATURE_ID})
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateTemplate(t *testing.T) {
	result, err := SMS_CLIENT.CreateTemplate(&api.CreateTemplateArgs{
		Name:        "测试",
		Content:     "${content}",
		SmsType:     "CommonNotice",
		CountryType: "DOMESTIC",
		Description: "gogogo",
	})
	ExpectEqual(t.Errorf, err, nil)
	TEST_TEMPLATE_ID = result.TemplateId
}

func TestGetTemplate(t *testing.T) {
	_, err := SMS_CLIENT.GetTemplate(&api.GetTemplateArgs{TemplateId: TEST_TEMPLATE_ID})
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyTemplate(t *testing.T) {
	err := SMS_CLIENT.ModifyTemplate(&api.ModifyTemplateArgs{
		TemplateId:  TEST_TEMPLATE_ID,
		Name:        "测试变更模板",
		Content:     "${code}",
		SmsType:     "CommonVcode",
		CountryType: "GLOBAL",
		Description: "this is a test",
	})
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteTemplate(t *testing.T) {
	err := SMS_CLIENT.DeleteTemplate(&api.DeleteTemplateArgs{
		TemplateId: TEST_TEMPLATE_ID,
	})
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateMobileBlack(t *testing.T) {
	err := SMS_CLIENT.CreateMobileBlack(&api.CreateMobileBlackArgs{
		Type:           "SignatureBlack",
		Phone:          "17600000000",
		SmsType:        "CommonNotice",
		SignatureIdStr: "1234",
	})
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetMobileBlack(t *testing.T) {
	res, err := SMS_CLIENT.GetMobileBlack(&api.GetMobileBlackArgs{
		Phone:          "17600000000",
		SmsType:        "CommonNotice",
		SignatureIdStr: "1234",
		StartTime:      "2023-07-10",
		EndTime:        "2023-07-20",
		PageSize:       "10",
		PageNo:         "1",
	})
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
}

func TestDeleteMobileBlack(t *testing.T) {
	err := SMS_CLIENT.DeleteMobileBlack(&api.DeleteMobileBlackArgs{
		Phones: "17600000000,17600000001",
	})
	ExpectEqual(t.Errorf, err, nil)
}
