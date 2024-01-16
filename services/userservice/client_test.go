package userservice

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	USER_SERVICE_CLIENT *Client

	// set these values before start test
	VPC_TEST_ID           = ""
	SUBNET_TEST_ID        = ""
	INSTANCE_ID           = ""
	CERT_ID               = ""
	CLUSTER_ID            = ""
	CLUSTER_PROPERTY_TEST = ""
	TEST_BLB_ID           = "lb-xxxxxxxx"
	SERVICE               = "xxxxxxxxxx.beijing.baidubce.com"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

// init 函数用于初始化全局变量
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
	err = decoder.Decode(confObj)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	USER_SERVICE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

// getClientToken 函数返回一个随机的UUID字符串
func getClientToken() string {
	return util.NewUUID()
}

// TestClient_CreateUserService 用于测试创建用户服务的函数
func TestClient_CreateUserService(t *testing.T) {
	args := &CreateUserServiceArgs{
		ClientToken: getClientToken(),
		InstanceId:  TEST_BLB_ID,
		Name:        "test_name",
		Description: "test_user_service_description",
		ServiceName: "testUserServiceName",
	}

	createResult, err := USER_SERVICE_CLIENT.CreateUserService(args)
	ExpectEqual(t.Errorf, nil, err)
	SERVICE = createResult.Service
}

// TestClient_UpdateUserService 是一个用于测试客户端函数UpdateUserService的测试函数
func TestClient_UpdateUserService(t *testing.T) {

	args := &UpdateServiceArgs{
		ClientToken: getClientToken(),
		Name:        "update_test_name",
		Description: "update_test_user_service_description",
	}

	err := USER_SERVICE_CLIENT.UpdateUserService(SERVICE, args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_UnBindInstance 测试客户端解绑实例函数
func TestClient_UnBindInstance(t *testing.T) {

	args := &UserServiceUnBindArgs{
		ClientToken: getClientToken(),
	}

	err := USER_SERVICE_CLIENT.UserServiceUnBindInstance(SERVICE, args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_BindInstance 测试绑定实例函数
func TestClient_BindInstance(t *testing.T) {

	args := &UserServiceBindArgs{
		ClientToken: getClientToken(),
		InstanceId:  TEST_BLB_ID,
	}

	err := USER_SERVICE_CLIENT.UserServiceBindInstance(SERVICE, args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_UserServiceRemoveAuth 测试用户服务移除授权
func TestClient_UserServiceRemoveAuth(t *testing.T) {

	args := &UserServiceRemoveAuthArgs{
		ClientToken: getClientToken(),
		UidList:     []string{"7cc5aff841ff4b648028d80000000000"},
	}

	err := USER_SERVICE_CLIENT.UserServiceRemoveAuth(SERVICE, args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_UserServiceAddAuth 用于测试用户服务添加权限的方法
func TestClient_UserServiceAddAuth(t *testing.T) {

	args := &UserServiceAuthArgs{
		ClientToken: getClientToken(),
		AuthList: []UserServiceAuthModel{
			{
				Uid:  "7cc5aff841ff4b648028d80000000000",
				Auth: ServiceAuthDeny,
			}},
	}

	err := USER_SERVICE_CLIENT.UserServiceAddAuth(SERVICE, args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_UserServiceEditAuth 测试UserServiceEditAuth函数
func TestClient_UserServiceEditAuth(t *testing.T) {

	args := &UserServiceAuthArgs{
		ClientToken: getClientToken(),
		AuthList: []UserServiceAuthModel{
			{
				Uid:  "7cc5aff841ff4b648028d80000000000",
				Auth: ServiceAuthAllow,
			}},
	}

	err := USER_SERVICE_CLIENT.UserServiceEditAuth(SERVICE, args)
	ExpectEqual(t.Errorf, nil, err)
}

// TestClient_DescribeUserServices 是一个测试函数，用于测试客户端的DescribeUserServices方法
func TestClient_DescribeUserServices(t *testing.T) {

	args := &DescribeUserServicesArgs{}

	result, err := USER_SERVICE_CLIENT.DescribeUserServices(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(result.Services))
}

// TestClient_DescribeUserServiceDetail 测试函数，用于测试客户端的DescribeUserServiceDetail方法
func TestClient_DescribeUserServiceDetail(t *testing.T) {

	result, err := USER_SERVICE_CLIENT.DescribeUserServiceDetail(SERVICE)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "available", len(result.Status))
}

// TestClient_DeleteUserService 是一个测试函数，用于测试 DeleteUserService 方法是否能够成功删除用户服务。
func TestClient_DeleteUserService(t *testing.T) {

	err := USER_SERVICE_CLIENT.DeleteUserService(SERVICE)
	ExpectEqual(t.Errorf, nil, err)
}
