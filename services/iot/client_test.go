package iot

import (
	"encoding/json"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	client *Client
	config *Conf
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK         string `json:"ak"`
	SK         string `json:"sk"`
	CoreId     string `json:"iotCoreId"`
	TemplateId string `json:"templateId"`
}

func init() {
	log.SetLogHandler(log.STDERR)
	log.SetLogLevel(log.DEBUG)

	configJson := "config.json"
	fp, err := os.Open(configJson)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", configJson)
		os.Exit(1)
	}

	config = &Conf{}
	decoder := json.NewDecoder(fp)
	decoder.Decode(config)

	log.Debugln(config)
	client, _ = NewClient(config.AK, config.SK)
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

func TestClientDevice(t *testing.T) {
	err := client.DeleteDevice(config.CoreId, t.Name())
	ExpectEqual(t.Errorf, nil, err)

	result, err := client.CreateDevice(config.CoreId, config.TemplateId, t.Name(), SIGNATURE, "iot core test")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, t.Name(), result.Name)

	result, err = client.CreateDevice(config.CoreId, config.TemplateId, t.Name(), SIGNATURE, "iot core test")
	ExpectEqual(t.Errorf, true, strings.Contains(err.Error(), "Device name has been occupied"))
	ExpectEqual(t.Errorf, nil, result)

	err = client.DeleteDevice(config.CoreId, t.Name())
	ExpectEqual(t.Errorf, nil, err)
}

func TestClientDeviceException(t *testing.T) {
	result, err := client.CreateDevice("", "", "", "", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)
	ExpectEqual(t.Errorf, nil, result)

	result, err = client.CreateDevice("1", "", "", "", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)
	ExpectEqual(t.Errorf, nil, result)

	result, err = client.CreateDevice("1", "2", "", "", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)
	ExpectEqual(t.Errorf, nil, result)

	result, err = client.CreateDevice("1", "", "", "", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)
	ExpectEqual(t.Errorf, nil, result)

	result, err = client.CreateDevice("1", "2", "", "", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)
	ExpectEqual(t.Errorf, nil, result)

	result, err = client.CreateDevice("1", "2", "3", "", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)
	ExpectEqual(t.Errorf, nil, result)

	result, err = client.CreateDevice("1", "2", "3", "4", "")
	ExpectEqual(t.Errorf, true, strings.Contains(err.Error(), "Bad IotCoreId"))
	ExpectEqual(t.Errorf, nil, result)

	result, err = client.CreateDevice("1", "2", "3", "4", "5")
	ExpectEqual(t.Errorf, true, strings.Contains(err.Error(), "Bad IotCoreId"))
	ExpectEqual(t.Errorf, nil, result)


	err = client.DeleteDevice("", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)

	err = client.DeleteDevice("1", "")
	ExpectEqual(t.Errorf, os.ErrInvalid, err)

	err = client.DeleteDevice("1", "2")
	ExpectEqual(t.Errorf, true, strings.Contains(err.Error(), "Bad IotCoreId"))
}
