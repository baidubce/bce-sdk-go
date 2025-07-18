package bos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BOS_CLIENT    *Client
	EXISTS_BUCKET = "gosdk-unittest-bucket"
	EXISTS_OBJECT = "gosdk-unittest-object"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	ENDPOINT string
}

// init client with your ak/sk written in config.json
func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given: %+v\n", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)
	BOS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.ENDPOINT)
	//log.SetLogHandler(log.STDERR | log.FILE)
	//log.SetRotateType(log.ROTATE_SIZE)
	log.SetLogLevel(log.WARN)
	//log.SetLogHandler(log.STDERR)
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
func TestListBuckets(t *testing.T) {
	res, err := BOS_CLIENT.ListBuckets()
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestListObjects(t *testing.T) {
	args := &api.ListObjectsArgs{Prefix: "test", MaxKeys: 10}
	res, err := BOS_CLIENT.ListObjects(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestSimpleListObjects(t *testing.T) {
	res, err := BOS_CLIENT.SimpleListObjects(EXISTS_BUCKET, "test", 10, "", "")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestHeadBucket(t *testing.T) {
	err := BOS_CLIENT.HeadBucket(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
func TestDoesBucketExist(t *testing.T) {
	exist, err := BOS_CLIENT.DoesBucketExist(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, exist, true)
	ExpectEqual(t.Errorf, err, nil)
	exist, _ = BOS_CLIENT.DoesBucketExist("xxx")
	ExpectEqual(t.Errorf, exist, false)
}
func TestPutBucket(t *testing.T) {
	res, err := BOS_CLIENT.PutBucket("test-put-bucket")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
}
func TestDeleteBucket(t *testing.T) {
	err := BOS_CLIENT.DeleteBucket("test-put-bucket")
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketLocation(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketLocation(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
}
func TestPutBucketAcl(t *testing.T) {
	acl := `{
    "accessControlList":[
        {
            "grantee":[{
                "id":"e13b12d0131b4c8bae959df4969387b8"
            }],
            "permission":["FULL_CONTROL"]
        }
    ]
}`
	body, _ := bce.NewBodyFromString(acl)
	err := BOS_CLIENT.PutBucketAcl(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[0].Id,
		"e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "FULL_CONTROL")
}
func TestPutBucketAclFromCanned(t *testing.T) {
	err := BOS_CLIENT.PutBucketAclFromCanned(EXISTS_BUCKET, api.CANNED_ACL_PUBLIC_READ)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[0].Id, "*")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "READ")
}
func TestPutBucketAclFromFile(t *testing.T) {
	acl := `{
    "accessControlList":[
        {
            "grantee":[
                {"id":"e13b12d0131b4c8bae959df4969387b8"},
                {"id":"a13b12d0131b4c8bae959df4969387b8"}
            ],
            "permission":["FULL_CONTROL"]
        }
    ]
}`
	fname := "/tmp/test-put-bucket-acl-by-file"
	f, _ := os.Create(fname)
	f.WriteString(acl)
	f.Close()
	err := BOS_CLIENT.PutBucketAclFromFile(EXISTS_BUCKET, fname)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	os.Remove(fname)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[0].Id,
		"e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[1].Id,
		"a13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "FULL_CONTROL")
}
func TestPutBucketAclFromString(t *testing.T) {
	acl := `{
    "accessControlList":[
        {
            "grantee":[
                {"id":"e13b12d0131b4c8bae959df4969387b8"},
                {"id":"a13b12d0131b4c8bae959df4969387b8"}
            ],
            "permission":["FULL_CONTROL"]
        }
    ]
}`
	err := BOS_CLIENT.PutBucketAclFromString(EXISTS_BUCKET, acl)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[0].Id,
		"e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[1].Id,
		"a13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "FULL_CONTROL")
}
func TestPutBucketAclFromStruct(t *testing.T) {
	args := &api.PutBucketAclArgs{
		AccessControlList: []api.GrantType{
			{
				Grantee: []api.GranteeType{
					{Id: "e13b12d0131b4c8bae959df4969387b8"},
				},
				Permission: []string{
					"FULL_CONTROL",
				},
			},
		},
	}
	err := BOS_CLIENT.PutBucketAclFromStruct(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[0].Id,
		"e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "FULL_CONTROL")
}
func TestPutBucketLogging(t *testing.T) {
	body, _ := bce.NewBodyFromString(
		`{"targetBucket": "gosdk-unittest-bucket", "targetPrefix": "my-log/"}`)
	err := BOS_CLIENT.PutBucketLogging(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.TargetBucket, "gosdk-unittest-bucket")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.TargetPrefix, "my-log/")
}
func TestPutBucketLoggingFromString(t *testing.T) {
	logging := `{"targetBucket": "gosdk-unittest-bucket", "targetPrefix": "my-log2/"}`
	err := BOS_CLIENT.PutBucketLoggingFromString(EXISTS_BUCKET, logging)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.TargetBucket, "gosdk-unittest-bucket")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.TargetPrefix, "my-log2/")
}
func TestPutBucketLoggingFromStruct(t *testing.T) {
	obj := &api.PutBucketLoggingArgs{
		TargetBucket: "gosdk-unittest-bucket",
		TargetPrefix: "my-log3/",
	}
	err := BOS_CLIENT.PutBucketLoggingFromStruct(EXISTS_BUCKET, obj)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.TargetBucket, "gosdk-unittest-bucket")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.TargetPrefix, "my-log3/")
}
func TestDeleteBucketLogging(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Status, "disabled")
}
func TestPutBucketLifecycle(t *testing.T) {
	str := `{
    "rule": [
        {
            "id": "transition-to-cold",
            "status": "enabled",
            "resource": ["gosdk-unittest-bucket/test*"],
            "condition": {
                "time": {
                    "dateGreaterThan": "2018-09-07T00:00:00Z"
                }
            },
            "action": {
                "name": "DeleteObject"
            }
        }
    ]
}`
	body, _ := bce.NewBodyFromString(str)
	err := BOS_CLIENT.PutBucketLifecycle(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketLifecycle(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Rule[0].Id, "transition-to-cold")
	ExpectEqual(t.Errorf, res.Rule[0].Status, "enabled")
	ExpectEqual(t.Errorf, res.Rule[0].Resource[0], "gosdk-unittest-bucket/test*")
	ExpectEqual(t.Errorf, res.Rule[0].Condition.Time.DateGreaterThan, "2018-09-07T00:00:00Z")
	ExpectEqual(t.Errorf, res.Rule[0].Action.Name, "DeleteObject")
}
func TestPutBucketLifecycleFromString(t *testing.T) {
	obj := `{
    "rule": [
        {
            "id": "transition-to-cold",
            "status": "enabled",
            "resource": ["gosdk-unittest-bucket/test*"],
            "condition": {
                "time": {
                    "dateGreaterThan": "2018-09-07T00:00:00Z"
                }
            },
            "action": {
                "name": "DeleteObject"
            }
        }
    ]
}`
	err := BOS_CLIENT.PutBucketLifecycleFromString(EXISTS_BUCKET, obj)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketLifecycle(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Rule[0].Id, "transition-to-cold")
	ExpectEqual(t.Errorf, res.Rule[0].Status, "enabled")
	ExpectEqual(t.Errorf, res.Rule[0].Resource[0], "gosdk-unittest-bucket/test*")
	ExpectEqual(t.Errorf, res.Rule[0].Condition.Time.DateGreaterThan, "2018-09-07T00:00:00Z")
	ExpectEqual(t.Errorf, res.Rule[0].Action.Name, "DeleteObject")
}
func TestDeleteBucketLifecycle(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketLifecycle(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	res, _ := BOS_CLIENT.GetBucketLifecycle(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, res, nil)
}
func TestPutBucketStorageClass(t *testing.T) {
	err := BOS_CLIENT.PutBucketStorageclass(EXISTS_BUCKET, api.STORAGE_CLASS_STANDARD_IA)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketStorageclass(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, api.STORAGE_CLASS_STANDARD_IA)
}
func TestGetBucketStorageClass(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketStorageclass(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestPutBucketReplication(t *testing.T) {
	BOS_CLIENT.DeleteBucketReplication(EXISTS_BUCKET, "")
	str := `{
		"id": "abc",
		"status":"enabled",
		"resource": ["gosdk-unittest-bucket/films"],
		"destination": {
			"bucket": "bos-rd-su-test",
			"storageClass": "COLD"
		},
		"replicateDeletes": "disabled"
}`
	body, _ := bce.NewBodyFromString(str)
	err := BOS_CLIENT.PutBucketReplication(EXISTS_BUCKET, body, "")
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Id, "abc")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.Resource[0], "gosdk-unittest-bucket/films")
	ExpectEqual(t.Errorf, res.Destination.Bucket, "bos-rd-su-test")
	ExpectEqual(t.Errorf, res.ReplicateDeletes, "disabled")
}
func TestPutBucketReplicationFromFile(t *testing.T) {
	BOS_CLIENT.DeleteBucketReplication(EXISTS_BUCKET, "")
	str := `{
		"id": "abc",
		"status":"enabled",
		"resource": ["gosdk-unittest-bucket/films"],
		"destination": {
			"bucket": "bos-rd-su-test",
			"storageClass": "COLD"
		},
		"replicateDeletes": "disabled"
}`
	fname := "/tmp/test-put-bucket-replication-by-file"
	f, _ := os.Create(fname)
	f.WriteString(str)
	f.Close()
	err := BOS_CLIENT.PutBucketReplicationFromFile(EXISTS_BUCKET, fname, "")
	os.Remove(fname)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Id, "abc")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.Resource[0], "gosdk-unittest-bucket/films")
	ExpectEqual(t.Errorf, res.Destination.Bucket, "bos-rd-su-test")
	ExpectEqual(t.Errorf, res.ReplicateDeletes, "disabled")
}
func TestPutBucketReplicationFromString(t *testing.T) {
	BOS_CLIENT.DeleteBucketReplication(EXISTS_BUCKET, "")
	str := `{
		"id": "abc",
		"status":"enabled",
		"resource": ["gosdk-unittest-bucket/films"],
		"destination": {
			"bucket": "bos-rd-su-test",
			"storageClass": "COLD"
		},
		"replicateDeletes": "disabled"
}`
	err := BOS_CLIENT.PutBucketReplicationFromString(EXISTS_BUCKET, str, "")
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Id, "abc")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.Resource[0], "gosdk-unittest-bucket/films")
	ExpectEqual(t.Errorf, res.Destination.Bucket, "bos-rd-su-test")
	ExpectEqual(t.Errorf, res.ReplicateDeletes, "disabled")
}
func TestPutBucketReplicationFromStruct(t *testing.T) {
	BOS_CLIENT.DeleteBucketReplication(EXISTS_BUCKET, "")
	args := &api.PutBucketReplicationArgs{
		Id:               "abc",
		Status:           "enabled",
		Resource:         []string{"gosdk-unittest-bucket/films"},
		Destination:      &api.BucketReplicationDescriptor{Bucket: "bos-rd-su-test", StorageClass: "COLD"},
		ReplicateDeletes: "disabled",
	}
	err := BOS_CLIENT.PutBucketReplicationFromStruct(EXISTS_BUCKET, args, "")
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Id, "abc")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.Resource[0], "gosdk-unittest-bucket/films")
	ExpectEqual(t.Errorf, res.Destination.Bucket, "bos-rd-su-test")
	ExpectEqual(t.Errorf, res.ReplicateDeletes, "disabled")
}
func TestGetBucketReplication(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Id, "abc")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.Resource[0], "gosdk-unittest-bucket/films")
	ExpectEqual(t.Errorf, res.Destination.Bucket, "bos-rd-su-test")
	ExpectEqual(t.Errorf, res.ReplicateDeletes, "disabled")
}
func TestGetBucketReplicationProcess(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketReplicationProgress(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
}
func TestDeleteBucketReplication(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutBucketEncryption(t *testing.T) {
	err := BOS_CLIENT.PutBucketEncryption(EXISTS_BUCKET, api.ENCRYPTION_AES256)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketEncryption(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, api.ENCRYPTION_AES256)
}
func TestGetBucketEncryption(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketEncryption(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestDeleteBucketEncryption(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketEncryption(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketEncryption(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestPutBucketStaticWebsite(t *testing.T) {
	BOS_CLIENT.DeleteBucketStaticWebsite(EXISTS_BUCKET)
	body, _ := bce.NewBodyFromString(`{"index": "index.html", "notFound":"blank.html"}`)
	err := BOS_CLIENT.PutBucketStaticWebsite(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Index, "index.html")
	ExpectEqual(t.Errorf, res.NotFound, "blank.html")
}
func TestPutBucketStaticWebsiteFromString(t *testing.T) {
	BOS_CLIENT.DeleteBucketStaticWebsite(EXISTS_BUCKET)
	jsonConf := `{"index": "index.html", "notFound":"blank.html"}`
	err := BOS_CLIENT.PutBucketStaticWebsiteFromString(EXISTS_BUCKET, jsonConf)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Index, "index.html")
	ExpectEqual(t.Errorf, res.NotFound, "blank.html")
}
func TestPutBucketStaticWebsiteFromStruct(t *testing.T) {
	BOS_CLIENT.DeleteBucketStaticWebsite(EXISTS_BUCKET)
	obj := &api.PutBucketStaticWebsiteArgs{Index: "index.html", NotFound: "blank.html"}
	err := BOS_CLIENT.PutBucketStaticWebsiteFromStruct(EXISTS_BUCKET, obj)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Index, "index.html")
	ExpectEqual(t.Errorf, res.NotFound, "blank.html")
}
func TestSimplePutBucketStaticWebsite(t *testing.T) {
	BOS_CLIENT.DeleteBucketStaticWebsite(EXISTS_BUCKET)
	err := BOS_CLIENT.SimplePutBucketStaticWebsite(EXISTS_BUCKET, "index.html", "blank.html")
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Index, "index.html")
	ExpectEqual(t.Errorf, res.NotFound, "blank.html")
}
func TestGetBucketStaticWebsite(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
}
func TestDeleteBucketStaticWebsite(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v", res)
}
func TestPutBucketCors(t *testing.T) {
	body, _ := bce.NewBodyFromString(`
	{
		"corsConfiguration": [
			{
				"allowedOrigins": ["https://www.baidu.com"],
				"allowedMethods": ["GET"],
				"maxAgeSeconds": 1800
			}
		]
	}
	`)
	err := BOS_CLIENT.PutBucketCors(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedOrigins[0], "https://www.baidu.com")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedMethods[0], "GET")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].MaxAgeSeconds, 1800)
}
func TestPutBucketCorsFromFile(t *testing.T) {
	str := `{
		"corsConfiguration": [
			{
				"allowedOrigins": ["https://www.baidu.com"],
				"allowedMethods": ["GET"],
				"maxAgeSeconds": 1800
			}
		]
	}`
	fname := "/tmp/test-put-bucket-cors-by-file"
	f, _ := os.Create(fname)
	f.WriteString(str)
	f.Close()
	err := BOS_CLIENT.PutBucketCorsFromFile(EXISTS_BUCKET, fname)
	os.Remove(fname)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedOrigins[0], "https://www.baidu.com")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedMethods[0], "GET")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].MaxAgeSeconds, 1800)
	err = BOS_CLIENT.PutBucketCorsFromFile(EXISTS_BUCKET, "/tmp/not-exist")
	ExpectEqual(t.Errorf, err != nil, true)
}
func TestPutBucketCorsFromString(t *testing.T) {
	str := `{
		"corsConfiguration": [
			{
				"allowedOrigins": ["https://www.baidu.com"],
				"allowedMethods": ["GET"],
				"maxAgeSeconds": 1800
			}
		]
	}`
	err := BOS_CLIENT.PutBucketCorsFromString(EXISTS_BUCKET, str)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedOrigins[0], "https://www.baidu.com")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedMethods[0], "GET")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].MaxAgeSeconds, 1800)
	err = BOS_CLIENT.PutBucketCorsFromString(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err != nil, true)
}
func TestPutBucketCorsFromStruct(t *testing.T) {
	obj := &api.PutBucketCorsArgs{
		CorsConfiguration: []api.BucketCORSType{
			{
				AllowedOrigins: []string{"https://www.baidu.com"},
				AllowedMethods: []string{"GET"},
				MaxAgeSeconds:  1200,
			},
		},
	}
	err := BOS_CLIENT.PutBucketCorsFromStruct(EXISTS_BUCKET, obj)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedOrigins[0], "https://www.baidu.com")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedMethods[0], "GET")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].MaxAgeSeconds, 1200)
	err = BOS_CLIENT.PutBucketCorsFromStruct(EXISTS_BUCKET, nil)
	ExpectEqual(t.Errorf, err != nil, true)
}
func TestGetBucketCors(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedOrigins[0], "https://www.baidu.com")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedMethods[0], "GET")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].MaxAgeSeconds, 1200)
}
func TestDeleteBucketCors(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v, %v", res, err)
}
func TestPutBucketCopyrightProtection(t *testing.T) {
	err := BOS_CLIENT.PutBucketCopyrightProtection(EXISTS_BUCKET,
		"gosdk-unittest-bucket/test-put-object", "gosdk-unittest-bucket/films/*")
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketCopyrightProtection(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res[0], "gosdk-unittest-bucket/test-put-object")
	ExpectEqual(t.Errorf, res[1], "gosdk-unittest-bucket/films/*")
}
func TestGetBucketCopyrightProtection(t *testing.T) {
	res, err := BOS_CLIENT.GetBucketCopyrightProtection(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v, %v", res, err)
}
func TestDeleteBucketCopyrightProtection(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketCopyrightProtection(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketCopyrightProtection(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v, %v", res, err)
}

func TestPutObject(t *testing.T) {
	args := &api.PutObjectArgs{StorageClass: api.STORAGE_CLASS_COLD}
	body, _ := bce.NewBodyFromString("12345")
	etag, err := BOS_CLIENT.PutObject(EXISTS_BUCKET, "test-put-object", body, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
}

func TestBasicPutObject(t *testing.T) {
	body, _ := bce.NewBodyFromString("12345")
	etag, err := BOS_CLIENT.BasicPutObject(EXISTS_BUCKET, "test-put-object", body)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
}

func TestPutObjectFromBytes(t *testing.T) {
	arr := []byte("12345")
	etag, err := BOS_CLIENT.PutObjectFromBytes(EXISTS_BUCKET, "test-put-object", arr, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
}

func TestPutObjectFromString(t *testing.T) {
	etag, err := BOS_CLIENT.PutObjectFromString(EXISTS_BUCKET, "test-put-object", "12345", nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
}

func TestPutObjectFromFile(t *testing.T) {
	fname := "/tmp/test-put-file"
	f, _ := os.Create(fname)
	f.WriteString("12345")
	f.Close()
	etag, err := BOS_CLIENT.PutObjectFromFile(EXISTS_BUCKET, "test-put-object", fname, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	args := &api.PutObjectArgs{ContentLength: 6}
	etag, err = BOS_CLIENT.PutObjectFromFile(EXISTS_BUCKET, "test-put-object", fname, args)
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, "", etag)
	os.Remove(fname)
}

func TestPutObjectFromStream(t *testing.T) {
	fname := "/tmp/test-put-file"
	fw, _ := os.Create(fname)
	defer os.Remove(fname)
	fw.WriteString("12345")
	fw.Close()
	fr, _ := os.Open(fname)
	defer fr.Close()
	etag, err := BOS_CLIENT.PutObjectFromStream(EXISTS_BUCKET, "test-put-object", fr, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
}

func TestCopyObject(t *testing.T) {
	args := new(api.CopyObjectArgs)
	args.StorageClass = api.STORAGE_CLASS_COLD
	res, err := BOS_CLIENT.CopyObject(EXISTS_BUCKET, "test-copy-object",
		EXISTS_BUCKET, "test-put-object", args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("copy result: %+v", res)
}
func TestBasicCopyObject(t *testing.T) {
	res, err := BOS_CLIENT.BasicCopyObject(EXISTS_BUCKET, "test-copy-object",
		EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("copy result: %+v", res)
}
func TestGetObject(t *testing.T) {
	res, err := BOS_CLIENT.GetObject(EXISTS_BUCKET, "test-put-object", nil)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
	t.Logf("%v", res.ContentLength)
	buf := make([]byte, 1024)
	n, _ := res.Body.Read(buf)
	t.Logf("%s", buf[0:n])
	res.Body.Close()
	/*
		respHeaders := map[string]string{"ContentEncoding": "qqqqqqqqqqqqq"}

			res, err = BOS_CLIENT.GetObject(EXISTS_BUCKET, "test-put-object", respHeaders)
			ExpectEqual(t.Errorf, err, nil)
			t.Logf("%+v", res)
			t.Logf("%v", res.ContentLength)
			n, _ = res.Body.Read(buf)
			t.Logf("%s", buf[0:n])
			res.Body.Close()
			res, err = BOS_CLIENT.GetObject(EXISTS_BUCKET, "test-put-object", respHeaders, 2)
			ExpectEqual(t.Errorf, err, nil)
			t.Logf("%+v", res)
			t.Logf("%v", res.ContentLength)
			n, _ = res.Body.Read(buf)
			t.Logf("%s", buf[0:n])
			res, err = BOS_CLIENT.GetObject(EXISTS_BUCKET, "test-put-object", respHeaders, 2, 4)
			ExpectEqual(t.Errorf, err, nil)
			t.Logf("%+v", res)
			t.Logf("%v", res.ContentLength)
			n, _ = res.Body.Read(buf)
			t.Logf("%s", buf[0:n])
			res.Body.Close()
	*/
}
func TestBasicGetObject(t *testing.T) {
	res, err := BOS_CLIENT.BasicGetObject(EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
	defer res.Body.Close()
	t.Logf("%v", res.ContentLength)
	for {
		buf := make([]byte, 1024)
		n, e := res.Body.Read(buf)
		t.Logf("%s", buf[0:n])
		if e != nil {
			break
		}
	}
}
func TestBasicGetObjectToFile(t *testing.T) {
	fname := "/tmp/test-get-object"
	err := BOS_CLIENT.BasicGetObjectToFile(EXISTS_BUCKET, "test-put-object", fname)
	ExpectEqual(t.Errorf, err, nil)
	os.Remove(fname)
	fname = "/bin/test-get-object"
	err = BOS_CLIENT.BasicGetObjectToFile(EXISTS_BUCKET, "test-put-object", fname)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v", err)
	err = BOS_CLIENT.BasicGetObjectToFile(EXISTS_BUCKET, "not-exist-object-name", fname)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v", err)
}

/*
	func TestGetObjectMeta(t *testing.T) {
		res, err := BOS_CLIENT.GetObjectMeta(EXISTS_BUCKET, "test-put-object")
		ExpectEqual(t.Errorf, err, nil)
		t.Logf("get object meta result: %+v", res)
	}
*/
func TestFetchObject(t *testing.T) {
	args := &api.FetchObjectArgs{
		FetchMode:    api.FETCH_MODE_ASYNC,
		StorageClass: api.STORAGE_CLASS_COLD,
	}
	res, err := BOS_CLIENT.FetchObject(EXISTS_BUCKET, "test-fetch-object",
		"https://cloud.baidu.com/doc/BOS/API.html", args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("result: %+v", res)
}
func TestBasicFetchObject(t *testing.T) {
	res, err := BOS_CLIENT.BasicFetchObject(EXISTS_BUCKET, "test-fetch-object",
		"https://bj.bcebos.com/gosdk-unittest-bucket/testsumlink")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("result: %+v", res)
	res1, err1 := BOS_CLIENT.GetObjectMeta(EXISTS_BUCKET, "test-fetch-object")
	ExpectEqual(t.Errorf, err1, nil)
	t.Logf("meta: %+v", res1)
}
func TestSimpleFetchObject(t *testing.T) {
	res, err := BOS_CLIENT.SimpleFetchObject(EXISTS_BUCKET, "test-fetch-object",
		"https://bj.bcebos.com/gosdk-unittest-bucket/testsumlink",
		api.FETCH_MODE_ASYNC, api.STORAGE_CLASS_COLD)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("result: %+v", res)
}
func TestAppendObject(t *testing.T) {
	args := &api.AppendObjectArgs{}
	body, _ := bce.NewBodyFromString("aaaaaaaaaaa")
	res, err := BOS_CLIENT.AppendObject(EXISTS_BUCKET, "test-append-object", body, args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestSimpleAppendObject(t *testing.T) {
	body, _ := bce.NewBodyFromString("bbbbbbbbbbb")
	res, err := BOS_CLIENT.SimpleAppendObject(EXISTS_BUCKET, "test-append-object", body, 11)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestSimpleAppendObjectFromString(t *testing.T) {
	res, err := BOS_CLIENT.SimpleAppendObjectFromString(
		EXISTS_BUCKET, "test-append-object", "123", 22)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestSimpleAppendObjectFromFile(t *testing.T) {
	fname := "/tmp/test-append-file"
	f, _ := os.Create(fname)
	f.WriteString("12345")
	f.Close()
	res, err := BOS_CLIENT.SimpleAppendObjectFromFile(EXISTS_BUCKET, "test-append-object", fname, 25)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
	os.Remove(fname)
}
func TestDeleteObject(t *testing.T) {
	err := BOS_CLIENT.DeleteObject(EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, err, nil)
}
func TestDeleteMultipleObjectsFromString(t *testing.T) {
	multiDeleteStr := `{
    "objects":[
        {"key": "aaaa"},
        {"key": "test-copy-object"},
        {"key": "test-append-object"},
        {"key": "cccc"}
    ]
}`
	res, err := BOS_CLIENT.DeleteMultipleObjectsFromString(EXISTS_BUCKET, multiDeleteStr)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestDeleteMultipleObjectsFromStruct(t *testing.T) {
	multiDeleteObj := &api.DeleteMultipleObjectsArgs{
		Objects: []api.DeleteObjectArgs{
			{Key: "1"}, {Key: "test-fetch-object"},
		},
	}
	res, err := BOS_CLIENT.DeleteMultipleObjectsFromStruct(EXISTS_BUCKET, multiDeleteObj)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestDeleteMultipleObjectsFromKeyList(t *testing.T) {
	keyList := []string{"aaaa", "test-copy-object", "test-append-object", "cccc"}
	res, err := BOS_CLIENT.DeleteMultipleObjectsFromKeyList(EXISTS_BUCKET, keyList)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestInitiateMultipartUpload(t *testing.T) {
	args := &api.InitiateMultipartUploadArgs{Expires: "aaaaaaa"}
	res, err := BOS_CLIENT.InitiateMultipartUpload(EXISTS_BUCKET, "test-multipart-upload", "", args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
	err1 := BOS_CLIENT.AbortMultipartUpload(EXISTS_BUCKET,
		"test-multipart-upload", res.UploadId)
	ExpectEqual(t.Errorf, err1, nil)
}
func TestBasicInitiateMultipartUpload(t *testing.T) {
	res, err := BOS_CLIENT.BasicInitiateMultipartUpload(EXISTS_BUCKET, "test-multipart-upload")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
	err1 := BOS_CLIENT.AbortMultipartUpload(EXISTS_BUCKET,
		"test-multipart-upload", res.UploadId)
	ExpectEqual(t.Errorf, err1, nil)
}
func TestUploadPart(t *testing.T) {
	res, err := BOS_CLIENT.UploadPart(EXISTS_BUCKET, "a", "b", 1, nil, nil)
	t.Logf("%+v, %+v", res, err)
}
func TestUploadPartCopy(t *testing.T) {
	res, err := BOS_CLIENT.UploadPartCopy(EXISTS_BUCKET, "test-multipart-upload",
		EXISTS_BUCKET, "test-multipart-copy", "12345", 1, nil)
	t.Logf("%+v, %+v", res, err)
}
func TestBasicUploadPartCopy(t *testing.T) {
	res, err := BOS_CLIENT.BasicUploadPartCopy(EXISTS_BUCKET, "test-multipart-upload",
		EXISTS_BUCKET, "test-multipart-copy", "12345", 1)
	t.Logf("%+v, %+v", res, err)
}
func TestListMultipartUploads(t *testing.T) {
	args := &api.ListMultipartUploadsArgs{MaxUploads: 10}
	res, err := BOS_CLIENT.ListMultipartUploads(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestBasicListMultipartUploads(t *testing.T) {
	res, err := BOS_CLIENT.BasicListMultipartUploads(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}
func TestUploadSuperFile(t *testing.T) {
	err := BOS_CLIENT.UploadSuperFile(EXISTS_BUCKET, "super-object", "test-object", "")
	ExpectEqual(t.Errorf, err, nil)
	err = BOS_CLIENT.UploadSuperFile(EXISTS_BUCKET, "not-exist", "not-exist", "")
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%+v", err)
}
func TestDownloadSuperFile(t *testing.T) {
	err := BOS_CLIENT.DownloadSuperFile(EXISTS_BUCKET, "super-object", "/dev/null")
	ExpectEqual(t.Errorf, err, nil)
	err = BOS_CLIENT.DownloadSuperFile(EXISTS_BUCKET, "not-exist", "/tmp/not-exist")
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%+v", err)
}

func TestGeneratePresignedUrl(t *testing.T) {
	url := BOS_CLIENT.BasicGeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 100)
	resp, err := http.Get(url)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, resp.StatusCode, 200)
	params := map[string]string{"responseContentType": "text"}
	url = BOS_CLIENT.GeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 1000, "HEAD", nil, params)
	resp, err = http.Head(url)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, resp.StatusCode, 200)
	BOS_CLIENT.Config.Endpoint = "10.180.112.31"
	url = BOS_CLIENT.GeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 100, "HEAD", nil, params)
	resp, err = http.Head(url)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, resp.StatusCode, 200)
	BOS_CLIENT.Config.Endpoint = "10.180.112.31:80"
	url = BOS_CLIENT.GeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 100, "HEAD", nil, params)
	resp, err = http.Head(url)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, resp.StatusCode, 200)
}

func TestGeneratePresignedUrl1(t *testing.T) {

	url := BOS_CLIENT.GeneratePresignedUrlPathStyle(EXISTS_BUCKET, EXISTS_OBJECT, 1000, "", nil, nil)
	t.Logf("res:%v\n", url)
}

func TestPutObjectAcl(t *testing.T) {
	acl := `{
    "accessControlList":[
        {
            "grantee":[{
                "id":"e13b12d0131b4c8bae959df4969387b8"
            }],
            "permission":["READ"]
        }
    ]
}`
	body, _ := bce.NewBodyFromString(acl)
	err := BOS_CLIENT.PutObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT, body)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "READ")
}
func TestPutObjectAclFromCanned(t *testing.T) {
	err := BOS_CLIENT.PutObjectAclFromCanned(EXISTS_BUCKET, EXISTS_OBJECT, api.CANNED_ACL_PUBLIC_READ)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
}
func TestPutObjectAclGrantRead(t *testing.T) {
	err := BOS_CLIENT.PutObjectAclGrantRead(EXISTS_BUCKET,
		EXISTS_OBJECT, "e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "READ")
}
func TestPutObjectAclGrantFullControl(t *testing.T) {
	err := BOS_CLIENT.PutObjectAclGrantFullControl(EXISTS_BUCKET,
		EXISTS_OBJECT, "e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "FULL_CONTROL")
}
func TestPutObjectAclFromFile(t *testing.T) {
	acl := `{
    "accessControlList":[
        {
            "grantee":[{
                "id":"e13b12d0131b4c8bae959df4969387b8"
            }],
            "permission":["READ"]
        }
    ]
}`
	fname := "/tmp/test-put-object-acl-by-file"
	f, _ := os.Create(fname)
	f.WriteString(acl)
	f.Close()
	err := BOS_CLIENT.PutObjectAclFromFile(EXISTS_BUCKET, EXISTS_OBJECT, fname)
	os.Remove(fname)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "READ")
}
func TestPutObjectAclFromString(t *testing.T) {
	acl := `{
    "accessControlList":[
        {
            "grantee":[{
                "id":"e13b12d0131b4c8bae959df4969387b8"
            }],
            "permission":["FULL_CONTROL"]
        }
    ]
}`
	err := BOS_CLIENT.PutObjectAclFromString(EXISTS_BUCKET, EXISTS_OBJECT, acl)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "FULL_CONTROL")
}
func TestPutObjectAclFromStruct(t *testing.T) {
	aclObj := &api.PutObjectAclArgs{
		AccessControlList: []api.GrantType{
			{
				Grantee: []api.GranteeType{
					{Id: "e13b12d0131b4c8bae959df4969387b8"},
				},
				Permission: []string{
					"READ",
				},
			},
		},
	}
	err := BOS_CLIENT.PutObjectAclFromStruct(EXISTS_BUCKET, EXISTS_OBJECT, aclObj)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v", res)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "READ")
}
func TestDeleteObjectAcl(t *testing.T) {
	err := BOS_CLIENT.DeleteObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v, %v", res, err)
}

func TestPutBucketTrash(t *testing.T) {
	args := api.PutBucketTrashReq{
		TrashDir: ".trash",
	}

	err := BOS_CLIENT.PutBucketTrash(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)

	res, err := BOS_CLIENT.GetBucketTrash(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err != nil, false)
	t.Logf("%v, %v", res, err)
}

func TestDeleteBucketTrash(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketTrash(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)

	res, err := BOS_CLIENT.GetBucketTrash(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v, %v", res, err)
}

func TestPutBucketNotification(t *testing.T) {
	notificationReq := api.PutBucketNotificationReq{
		Notifications: []api.PutBucketNotificationSt{
			{
				Id:        "water-test-1",
				Name:      "water-rule-1",
				AppId:     "water-app-id-1",
				Status:    "enabled",
				Resources: []string{EXISTS_BUCKET + "/path1", "/path2"},
				Events:    []string{"PutObject"},
				Apps: []api.PutBucketNotificationAppsSt{
					{
						Id:       "app-id-1",
						EventUrl: "http://xxx.com/event",
						XVars:    "",
					},
				},
			},
		},
	}

	err := BOS_CLIENT.PutBucketNotification(EXISTS_BUCKET, notificationReq)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketNotification(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err != nil, false)
	t.Logf("%v, %v", res, err)
}

func TestDeleteBucketNotification(t *testing.T) {
	err := BOS_CLIENT.DeleteBucketNotification(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BOS_CLIENT.GetBucketNotification(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err != nil, true)
	t.Logf("%v, %v", res, err)
}

func TestParallelUpload(t *testing.T) {
	res, err := BOS_CLIENT.ParallelUpload(EXISTS_BUCKET, "test_multiupload", "test_object", "", nil)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v,%v", res, err)
}

func TestParallelCopy(t *testing.T) {
	args := api.MultiCopyObjectArgs{
		StorageClass: api.STORAGE_CLASS_COLD,
	}
	res, err := BOS_CLIENT.ParallelCopy(EXISTS_BUCKET, "test_multiupload", EXISTS_BUCKET, "test_multiupload_copy", &args, nil)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%v,%v", res, err)
}

func TestBucketTag(t *testing.T) {
	args := &api.PutBucketTagArgs{
		Tags: []api.Tag{
			{
				TagKey:   "key1",
				TagValue: "value1",
			},
		},
	}
	err := BOS_CLIENT.PutBucketTag(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	_, err = BOS_CLIENT.GetBucketTag(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	err = BOS_CLIENT.DeleteBucketTag(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}

func TestSymLink(t *testing.T) {
	args := &api.PutSymlinkArgs{}

	err := BOS_CLIENT.PutSymlink(EXISTS_BUCKET, "test-symlink", EXISTS_OBJECT, args)
	ExpectEqual(t.Errorf, err, nil)

	res, err := BOS_CLIENT.GetSymlink(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "test-symlink")
	t.Logf("%v, %v", res, err)
}

func TestBucketMirror(t *testing.T) {
	args := &api.PutBucketMirrorArgs{
		BucketMirroringConfiguration: []api.MirrorConfigurationRule{
			{
				SourceUrl:    "http://gosdk-unittest-bucket.bj.bcebos.com",
				Prefix:       "testprefix",
				StorageClass: api.STORAGE_CLASS_STANDARD,
				Mode:         "fetch",
			},
		},
	}

	err := BOS_CLIENT.PutBucketMirror(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	MirroConfig, err := BOS_CLIENT.GetBucketMirror(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, MirroConfig.BucketMirroringConfiguration[0].SourceUrl, "http://gosdk-unittest-bucket.bj.bcebos.com")
	ExpectEqual(t.Errorf, MirroConfig.BucketMirroringConfiguration[0].Prefix, "testprefix")
	err = BOS_CLIENT.DeleteBucketMirror(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
