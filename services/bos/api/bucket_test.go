package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
	mhttp "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

const errorJsonBody string = `{ "owner":{ "id":"10 , "buckets":[ { "na }`

func AttachMockHttpClientOk(t *testing.T, client *bce.BceClient, respBody *string) {
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
	}
	if respBody != nil {
		options1 = append(options1, util.SetRespBody(*respBody))
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
}

func AttachMockHttpClientError(t *testing.T, client *bce.BceClient, err error) {
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
}

func CheckMockHttpClientError(t *testing.T, host, path, feature, method string, expErr, actErr error) {
	expectUrl := fmt.Sprintf("http://%s%s", host, path)
	if len(feature) > 0 {
		expectUrl += fmt.Sprintf("?%s", feature)
	}
	expectError := &url.Error{
		Op:  method,
		URL: expectUrl,
		Err: expErr,
	}
	ExpectEqual(t, httpClientDoError(3, expectError).Error(), actErr.Error())
}

func AttachMockHttpClientRespFail(t *testing.T, client *bce.BceClient, options []util.MockRoundTripperOption) {
	mockHttpClient4 := util.NewMockHTTPClient(options...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
}

func AttachMockHttpClientJsonBodyError(t *testing.T, client *bce.BceClient) {

	options5 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(errorJsonBody),
	}
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
}

func TestListBuckets(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	//case1: ok
	respBody1 := `{
		"owner":{ "id":"10eb6f5ff6ff4605bf044313e8f3ffa5", "displayName":"BosUser" },
		"buckets":[
			{ "name":"bucket1", "location":"bj", "creationDate":"2016-04-05T10:20:35Z", "enableMultiAz":true },
			{ "name":"bucket2",	"location":"bj", "creationDate":"2016-04-05T16:41:58Z" }
		]
	}`
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody1),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	res, err := ListBuckets(client, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Buckets))
	ExpectEqual(t, "10eb6f5ff6ff4605bf044313e8f3ffa5", res.Owner.Id)
	ExpectEqual(t, "BosUser", res.Owner.DisplayName)
	ExpectEqual(t, "bucket1", res.Buckets[0].Name)
	ExpectEqual(t, "bj", res.Buckets[0].Location)
	ExpectEqual(t, "2016-04-05T10:20:35Z", res.Buckets[0].CreationDate)
	ExpectEqual(t, true, res.Buckets[0].EnableMultiAz)
	ExpectEqual(t, "bucket2", res.Buckets[1].Name)
	ExpectEqual(t, "bj", res.Buckets[1].Location)
	ExpectEqual(t, "2016-04-05T16:41:58Z", res.Buckets[1].CreationDate)
	ExpectEqual(t, false, res.Buckets[1].EnableMultiAz)
	//case2: handle options error
	res, err = ListBuckets(client, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = ListBuckets(client, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "", "", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = ListBuckets(client, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = ListBuckets(client, nil)
	result5 := &ListBucketsResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestListObjects(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test-bucket"
	args := &ListObjectsArgs{
		Delimiter:       "/",
		Marker:          "marker",
		MaxKeys:         100,
		Prefix:          "prefix",
		VersionIdMarker: "version_id_marker",
	}
	//case1: ok
	respBody1 := `{
		"name":"bucket", "prefix":"",
		"delimiter":"/", "marker":"",
		"maxKeys":1000, "isTruncated":false,
		"contents":[
			{
			    "key":"my-image.jpg", "lastModified":"2009-10-12T17:50:30Z",
				"eTag":"fba9dede5f27731c9771645a39863328",
			    "size":434234,  "storageClass":"STANDARD",
			    "owner":{ "id":"168bf6fd8fa74d9789f35a283a1f15e2", "displayName":"mtd" }
			},
			{
			    "key":"my-image1.jpg", "lastModified":"2009-10-12T17:51:30Z",
			    "eTag":"0cce7caecc8309864f663d78d1293f98",
			    "size":124231, "storageClass":"COLD",
			    "owner":{ "id":"168bf6fd8fa74d9789f35a283a1f15e2", "displayName":"mtd" }
			}
		],
		"commonPrefixes":[ {"prefix":"photos/"}, {"prefix":"mtd/"} ]
    }`
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody1),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	res, err := ListObjects(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Contents))
	ExpectEqual(t, 2, len(res.CommonPrefixes))
	ExpectEqual(t, "my-image.jpg", res.Contents[0].Key)
	ExpectEqual(t, "STANDARD", res.Contents[0].StorageClass)
	ExpectEqual(t, "photos/", res.CommonPrefixes[0].Prefix)
	ExpectEqual(t, "2009-10-12T17:51:30Z", res.Contents[1].LastModified)
	ExpectEqual(t, "mtd/", res.CommonPrefixes[1].Prefix)
	ExpectEqual(t, 124231, res.Contents[1].Size)
	res, err = ListObjectsVersions(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Contents))
	ExpectEqual(t, 2, len(res.CommonPrefixes))
	ExpectEqual(t, "my-image.jpg", res.Contents[0].Key)
	ExpectEqual(t, "STANDARD", res.Contents[0].StorageClass)
	ExpectEqual(t, "photos/", res.CommonPrefixes[0].Prefix)
	ExpectEqual(t, "2009-10-12T17:51:30Z", res.Contents[1].LastModified)
	ExpectEqual(t, "mtd/", res.CommonPrefixes[1].Prefix)
	ExpectEqual(t, 124231, res.Contents[1].Size)
	//case2: handle options err
	args.MaxKeys = 0
	res, err = ListObjects(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	res, err = ListObjectsVersions(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request fail
	args.MaxKeys = 0
	err3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = ListObjects(client, bucket, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = ListObjects(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	res, err = ListObjectsVersions(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	respBody5 := `{ "owner":{ "id":"10 , "buckets":[ { "na }`
	options5 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody5),
	}
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	res, err = ListObjects(client, bucket, args, nil)
	result5 := &ListObjectsResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(respBody5))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
	res, err = ListObjectsVersions(client, bucket, args, nil)
	jsonDecoder51 := json.NewDecoder(bytes.NewBufferString(respBody5))
	ExpectEqual(t, jsonDecoder51.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestHeadBucket(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	err, resp := HeadBucket(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, resp != nil)
	//case2: handle options error
	err, _ = HeadBucket(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err, _ = HeadBucket(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "", "Head", err3, err)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	err, _ = HeadBucket(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucket(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	args := &PutBucketArgs{
		TagList:       "key1=value1&key2=value2",
		EnableMultiAz: true,
		LccLocation:   "lcc_location",
	}
	//case1: ok
	options1 := []util.MockRoundTripperOption{
		util.AddHeaders(map[string]string{mhttp.LOCATION: "bj"}),
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	res, err := PutBucket(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "bj", res)
	//case2: handle options error
	res, err = PutBucket(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = PutBucket(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "", "Put", err3, err)
	ExpectEqual(t, "", res)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = PutBucket(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", res)
}

func TestDeleteBucket(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	err = DeleteBucket(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucket(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = DeleteBucket(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "", "Delete", err3, err)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	err = DeleteBucket(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketLocation(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{ "locationConstraint": "bj" }`
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody1),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	res, err := GetBucketLocation(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "bj", res)
	//case2: handle options error
	res, err = GetBucketLocation(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = GetBucketLocation(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "location", "Get", err3, err)
	ExpectEqual(t, "", res)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = GetBucketLocation(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", res)
	//case5: parse json body fail
	respBody5 := `{ "owner":{ "id":"10 , "buckets":[ { "na }`
	options5 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody5),
	}
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	res, err = GetBucketLocation(client, bucket, nil)
	result5 := &LocationType{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(respBody5))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, "", res)
}

func TestPutBucketAcl(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	err = PutBucketAcl(client, bucket, CANNED_ACL_PRIVATE, nil, nil)
	ExpectEqual(t, nil, err)
	//case2: not support cannedAcl and acl file at the same time
	body2, err := bce.NewBodyFromString("acl string")
	ExpectEqual(t, nil, err)
	err = PutBucketAcl(client, bucket, CANNED_ACL_PRIVATE, body2, nil)
	ExpectEqual(t, bce.NewBceClientError("BOS does not support cannedAcl and acl file at the same time"), err)
	//case3: valid canned acl
	err = PutBucketAcl(client, bucket, CANNED_ACL_PRIVATE, nil, nil)
	ExpectEqual(t, nil, err)
	//case4: valid body
	body4, err := bce.NewBodyFromString("acl string")
	ExpectEqual(t, nil, err)
	err = PutBucketAcl(client, bucket, "", body4, nil)
	ExpectEqual(t, nil, err)
	//case5: handle options error
	err = PutBucketAcl(client, bucket, CANNED_ACL_PRIVATE, nil, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case6: send request error
	err6 := fmt.Errorf("IO Error")
	options6 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err6),
	}
	mockHttpClient6 := util.NewMockHTTPClient(options6...)
	ExpectEqual(t, true, mockHttpClient6 != nil)
	client.HTTPClient = mockHttpClient6
	err = PutBucketAcl(client, bucket, CANNED_ACL_PRIVATE, nil, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "acl", "Put", err6, err)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	err = PutBucketAcl(client, bucket, CANNED_ACL_PRIVATE, nil, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketAcl(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"owner" : { "id" : "16df583fe6824d73a5f858f06de0af03" },
		"accessControlList" : [
			{
				"grantee" : [{ "id":"168bf6fd8fa74d9789f35a283a1f15e2" }],
				"permission":[ "FULL_CONTROL" ]
			},
			{ 
				"grantee":[{ "id":"10eb6f5ff6ff4605bf044313e8f3ffa5" }],
				"permission":[ "READ" ]
			}
		]
	}`
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody1),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	res, err := GetBucketAcl(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "16df583fe6824d73a5f858f06de0af03", res.Owner.Id)
	ExpectEqual(t, 2, len(res.AccessControlList))
	ExpectEqual(t, "FULL_CONTROL", res.AccessControlList[0].Permission[0])
	ExpectEqual(t, "READ", res.AccessControlList[1].Permission[0])
	//case2: handle options error
	res, err = GetBucketAcl(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = GetBucketAcl(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "acl", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = GetBucketAcl(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketAcl(client, bucket, nil)
	result5 := &LocationType{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestPutBucketLogging(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	body1, err := bce.NewBodyFromString("logging config string")
	ExpectEqual(t, nil, err)
	err = PutBucketLogging(client, bucket, body1, nil)
	ExpectEqual(t, nil, err)
	//case 1.1: body is nil
	err = PutBucketLogging(client, bucket, nil, nil)
	ExpectEqual(t, bce.NewBceClientError("logging config is nil"), err)
	//case2: handle options error
	body2, err := bce.NewBodyFromString("logging config string")
	ExpectEqual(t, nil, err)
	err = PutBucketLogging(client, bucket, body2, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	body3, err := bce.NewBodyFromString("logging config string")
	ExpectEqual(t, nil, err)
	err = PutBucketLogging(client, bucket, body3, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "logging", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	body4, err := bce.NewBodyFromString("logging config string")
	ExpectEqual(t, nil, err)
	err = PutBucketLogging(client, bucket, body4, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketLogging(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"status" : "enabled",
		"targetBucket" : "dscbucket",
		"targetPrefix" : "mylog-",
		"enableOtherConfig" : "LOGGING_ENABLE_NONE"
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketLogging(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "enabled", res.Status)
	ExpectEqual(t, "dscbucket", res.TargetBucket)
	ExpectEqual(t, "mylog-", res.TargetPrefix)
	//case2: handle options error
	res, err = GetBucketLogging(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketLogging(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "logging", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketLogging(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketLogging(client, bucket, nil)
	result5 := &GetBucketLoggingResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketLogging(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketLogging(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketLogging(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketLogging(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "logging", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketLogging(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketLifecycle(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	body1, err := bce.NewBodyFromString("lifecycle config string")
	ExpectEqual(t, nil, err)
	err = PutBucketLifecycle(client, bucket, body1, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	body2, err := bce.NewBodyFromString("lifecycle config string")
	ExpectEqual(t, nil, err)
	err = PutBucketLifecycle(client, bucket, body2, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	body3, err := bce.NewBodyFromString("lifecycle config string")
	ExpectEqual(t, nil, err)
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketLifecycle(client, bucket, body3, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "lifecycle", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	body4, err := bce.NewBodyFromString("lifecycle config string")
	ExpectEqual(t, nil, err)
	err = PutBucketLogging(client, bucket, body4, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketLifecycle(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"rule": [
			{
				"id" : "sample-rule-delete-prefix", 
				"status" : "enabled", 
				"resource" : [ "bucket/prefix/*" ], 
				"condition": { "time": { "dateGreaterThan": "2016-09-07T00:00:00Z" } }, 
				"action": { "name": "DeleteObject" }
			},
			{
				"id": "sample-rule-transition-prefix",
				"status": "enabled",
				"resource": [ "bucket/prefix/*" ],
				"condition": { "time": { "dateGreaterThan": "$(lastModified)+P7D" } },
				"action": { "name": "Transition", "storageClass": "STANDARD_IA" }
			},
			{
				"id": "sample-rule-abort-multiupload-prefix",
				"status": "enabled",
				"resource": [ "bucket/prefix/*" ],
				"condition": { "time": { "dateGreaterThan": "$(lastModified)+P7D" } },
				"action": { "name": "AbortMultipartUpload" }
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketLifecycle(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 3, len(res.Rule))
	//case2: handle options error
	res, err = GetBucketLifecycle(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketLifecycle(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "lifecycle", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketLifecycle(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketLifecycle(client, bucket, nil)
	result5 := &GetBucketLifecycleResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
	//case6: compex lifecycle
	respBody6 := `{
		"rule": [
			{
				"id": "rule-id", "status": "enabled",
				"ExpiredObjectDeleteMarker" :"true",
				"resource": [ "bucket/prefix/*" ],
				"condition": {
					"time": { "dateGreaterThan": "$(lastModified)+P3D" },
					"objectSize":{ "minSize":1024, "maxSize":2048 },
					"tag": { "key1":"value1", "key2":"value2" }
				},
				"action": { "name": "DeleteObject" },
				"not":{
					"resource": "bucket/prefix/prefix1*",
					"tag": { "key3":"value3" }
				}
			},
			{
				"id": "rule-id2", "status": "enabled",
				"resource": [ "bucket/*" ],
				"condition": {
					"time": { "dateGreaterThan": "$(lastModified)+P5D" },
					"tag": { "key1":"value1",  "key2":"value2" }
				},
				"action": { "name": "NonCurrentVersionDeleteObject" }
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody6)
	res, err = GetBucketLifecycle(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Rule))
	ExpectEqual(t, "true", res.Rule[0].ExpiredObjectDeleteMarker)
	ExpectEqual(t, 1024, res.Rule[0].Condition.ObjectSize.MinSize)
	ExpectEqual(t, 2048, res.Rule[0].Condition.ObjectSize.MaxSize)
	ExpectEqual(t, 2, len(res.Rule[0].Condition.Tag))
	ExpectEqual(t, "value1", res.Rule[0].Condition.Tag["key1"])
	ExpectEqual(t, "bucket/prefix/prefix1*", res.Rule[0].Not.Resource)
	ExpectEqual(t, "value3", res.Rule[0].Not.Tag["key3"])
}

func TestDeleteBucketLifecycle(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketLifecycle(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketLifecycle(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketLifecycle(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "lifecycle", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketLifecycle(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketStorageclass(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketStorageclass(client, bucket, STORAGE_CLASS_STANDARD, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketStorageclass(client, bucket, STORAGE_CLASS_STANDARD, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketStorageclass(client, bucket, STORAGE_CLASS_STANDARD, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "storageClass", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketStorageclass(client, bucket, STORAGE_CLASS_STANDARD, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketStorageclass(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{"storageClass": "COLD"}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketStorageclass(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "COLD", res)
	//case2: handle options error
	res, err = GetBucketStorageclass(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketStorageclass(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "storageClass", "Get", err3, err)
	ExpectEqual(t, "", res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketStorageclass(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketStorageclass(client, bucket, nil)
	result5 := &StorageClassType{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, "", res)
}

func TestPutBucketReplication(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	body1, err := bce.NewBodyFromString("replication config string")
	ExpectEqual(t, nil, err)
	err = PutBucketReplication(client, bucket, body1, "ruleId", nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	body2, err := bce.NewBodyFromString("replication config string")
	ExpectEqual(t, nil, err)
	err = PutBucketReplication(client, bucket, body2, "ruleId", nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	body3, err := bce.NewBodyFromString("replication config string")
	ExpectEqual(t, nil, err)
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketReplication(client, bucket, body3, "ruleId", nil)
	ExpectEqual(t, true, err != nil)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	body4, err := bce.NewBodyFromString("replication config string")
	ExpectEqual(t, nil, err)
	err = PutBucketReplication(client, bucket, body4, "ruleId", nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketReplication(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"status": "enabled",
		"resource": [ "src-bucket-name/abc", "src-bucket-name/cd*" ],
		"destination": { "bucket": "dst-bucket-name", "storageClass": "COLD" },
		"replicateHistory": { "storageClass": "COLD" },
		"replicateDeletes": "enabled",
		"id": "sample-bucket",
		"createTime": 1583060606,
		"destRegion": "bj"
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketReplication(client, bucket, "ruleId", nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "enabled", res.Status)
	//case2: handle options error
	res, err = GetBucketReplication(client, bucket, "ruleId", nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketReplication(client, bucket, "ruleId", nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketReplication(client, bucket, "ruleId", nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketReplication(client, bucket, "ruleId", nil)
	result5 := &GetBucketReplicationResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestListBucketReplication(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"rules": [
			{
				"status": "enabled",
				"resource": [ "src-bucket-name/abc", "src-bucket-name/cd*" ],
				"destination": { "bucket": "dst-bucket-name", "storageClass": "COLD" },
				"replicateHistory": { "storageClass": "COLD" },
				"replicateDeletes": "enabled",
				"id": "sample-bucket-replication-config",
				"createTime": 1583060606,
				"destRegion": "bj"
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := ListBucketReplication(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "enabled", res.Rules[0].Status)
	//case2: handle options error
	res, err = ListBucketReplication(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = ListBucketReplication(client, bucket, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = ListBucketReplication(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = ListBucketReplication(client, bucket, nil)
	result5 := &ListBucketReplicationResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketReplication(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketReplication(client, bucket, "ruleId", nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketReplication(client, bucket, "ruleId", nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketReplication(client, bucket, "ruleId", nil)
	ExpectEqual(t, true, err != nil)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketReplication(client, bucket, "ruleId", nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketReplicationProgress(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"status" : "enabled",
		"historyReplicationPercent" : 5,
		"latestReplicationTime" : "1504448315"
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketReplicationProgress(client, bucket, "ruleId", nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "enabled", res.Status)
	//case2: handle options error
	res, err = GetBucketReplicationProgress(client, bucket, "ruleId", nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketReplicationProgress(client, bucket, "ruleId", nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketReplicationProgress(client, bucket, "ruleId", nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketReplicationProgress(client, bucket, "ruleId", nil)
	result5 := &GetBucketReplicationProgressResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestPutBucketEncryption(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketEncryption(client, bucket, ENCRYPTION_AES256, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketEncryption(client, bucket, ENCRYPTION_AES256, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketEncryption(client, bucket, ENCRYPTION_AES256, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "encryption", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketEncryption(client, bucket, ENCRYPTION_AES256, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketEncryption(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{"encryptionAlgorithm":"SM4"}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketEncryption(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "SM4", res)
	//case2: handle options error
	res, err = GetBucketEncryption(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketEncryption(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "encryption", "Get", err3, err)
	ExpectEqual(t, "", res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketEncryption(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketEncryption(client, bucket, nil)
	result5 := &GetBucketReplicationResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, "", res)
}

func TestDeleteBucketEncryption(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketEncryption(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketEncryption(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketEncryption(client, bucket, nil)
	ExpectEqual(t, true, err != nil)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketEncryption(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketStaticWebsite(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	body1, err := bce.NewBodyFromString("static website config string")
	ExpectEqual(t, nil, err)
	err = PutBucketStaticWebsite(client, bucket, body1, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	body2, err := bce.NewBodyFromString("static website config string")
	ExpectEqual(t, nil, err)
	err = PutBucketStaticWebsite(client, bucket, body2, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	body3, err := bce.NewBodyFromString("static website config string")
	ExpectEqual(t, nil, err)
	err = PutBucketStaticWebsite(client, bucket, body3, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "website", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	body4, err := bce.NewBodyFromString("static website config string")
	ExpectEqual(t, nil, err)
	err = PutBucketStaticWebsite(client, bucket, body4, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketStaticWebsite(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := ` {
		"index": "index.html",
		"notFound": "404.html",
		"notFoundHttpStatus":"404"
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketStaticWebsite(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "index.html", res.Index)
	//case2: handle options error
	res, err = GetBucketStaticWebsite(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketStaticWebsite(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "website", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketStaticWebsite(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketStaticWebsite(client, bucket, nil)
	result5 := &GetBucketStaticWebsiteResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketStaticWebsite(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketStaticWebsite(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketStaticWebsite(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketStaticWebsite(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "website", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketStaticWebsite(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketCors(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	body1, err := bce.NewBodyFromString("cors config string")
	ExpectEqual(t, nil, err)
	err = PutBucketCors(client, bucket, body1, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	body2, err := bce.NewBodyFromString("cors config string")
	ExpectEqual(t, nil, err)
	err = PutBucketCors(client, bucket, body2, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	body3, err := bce.NewBodyFromString("cors config string")
	ExpectEqual(t, nil, err)
	err = PutBucketCors(client, bucket, body3, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "cors", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	body4, err := bce.NewBodyFromString("cors config string")
	ExpectEqual(t, nil, err)
	err = PutBucketCors(client, bucket, body4, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketCors(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"corsConfiguration": [
			{
				"allowedOrigins": [ "http://www.example.com", "www.example2.com" ],
				"allowedMethods": [ "GET", "HEAD", "DELETE" ],
				"allowedHeaders": [ "Authorization", "x-bce-test", "x-bce-test2" ],
				"allowedExposeHeaders": [ "user-custom-expose-header" ],
				"maxAgeSeconds": 3600
			},
			{
				"allowedOrigins": [ "http://www.baidu.com" ],
				"allowedMethods": [ "GET", "HEAD", "DELETE" ],
				"allowedHeaders": [ "*" ],
				"allowedExposeHeaders": [ "user-custom-expose-header" ],
				"maxAgeSeconds": 1800
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketCors(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.CorsConfiguration))
	//case2: handle options error
	res, err = GetBucketCors(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketCors(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "cors", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketCors(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketCors(client, bucket, nil)
	result5 := &GetBucketStaticWebsiteResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketCors(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketCors(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketCors(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketCors(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "cors", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketCors(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketCopyrightProtection(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketCopyrightProtection(client, nil, bucket, "resource")
	ExpectEqual(t, nil, err)
	//case2: resource is nil
	err = PutBucketCopyrightProtection(client, nil, bucket, []string{}...)
	ExpectEqual(t, bce.NewBceClientError("the resource to set copyright protection is empty"), err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketCopyrightProtection(client, nil, bucket, "resource")
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "copyrightProtection", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketCopyrightProtection(client, nil, bucket, "resource")
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketCopyrightProtection(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{	"resource": [ "bucket/prefix/*", "bucket/*/suffix" ] }`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketCopyrightProtection(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res))
	//case2: handle options error
	res, err = GetBucketCopyrightProtection(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketCopyrightProtection(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "copyrightProtection", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketCopyrightProtection(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketCopyrightProtection(client, bucket, nil)
	result5 := &CopyrightProtectionType{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketCopyrightProtection(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketCopyrightProtection(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketCopyrightProtection(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketCopyrightProtection(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "copyrightProtection", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketCopyrightProtection(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketTrash(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketTrash(client, bucket, PutBucketTrashReq{"trashDir"}, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketTrash(client, bucket, PutBucketTrashReq{"trashDir"}, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketTrash(client, bucket, PutBucketTrashReq{"trashDir"}, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "trash", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketTrash(client, bucket, PutBucketTrashReq{"trashDir"}, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketTrash(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"trashDir": ".trash/"
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketTrash(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, ".trash/", res.TrashDir)
	//case2: handle options error
	res, err = GetBucketTrash(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketTrash(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "trash", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketTrash(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketTrash(client, bucket, nil)
	result5 := &GetBucketTrashResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketTrash(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketTrash(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketTrash(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketTrash(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "trash", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketTrash(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutAndGetBucketNotification(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	argsJsonStr := `{
		"notifications": [
			{
				"id": "notify-id-1",
				"name": "rule-name",
				"appId": "app-id-1",
				"status": "enabled",
				"encryption": { "key": "06a62b70f47dc4a0a7da349609f1a1ac"},
				"resources": [ "bucket-a/path1", "/path2", "/path3/*.jpg", "/path4/*" ],
				"events": [ "PutObject" ],
				"apps": [
					{ "id": "app-id-1", "eventUrl": "http://xxx.com/event", "xVars": "" },
					{ "id": "app-id-2", "eventUrl": "brn:bce:cfc:bj:1f1c3e0d5b22:function:hello_bos:$LATEST" },
					{ "id": "app-id-3", "eventUrl": "app:ImageOcr", "xVars": "{\"saveUrl\": \"http://xxx.com/ocr\"}" }
				]
			}
		]
	}`
	args := &PutBucketNotificationReq{}
	ExpectEqual(t, nil, json.Unmarshal([]byte(argsJsonStr), args))
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketNotification(client, bucket, *args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketNotification(client, bucket, *args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketNotification(client, bucket, *args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "notification", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketNotification(client, bucket, *args, nil)
	ExpectEqual(t, bceServiceErro404, err)

	//GetBucketNotification
	//case1: ok
	AttachMockHttpClientOk(t, client, &argsJsonStr)
	res, err := GetBucketNotification(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 1, len(res.Notifications))
	//case2: handle options error
	res, err = GetBucketNotification(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketNotification(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "notification", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketNotification(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketNotification(client, bucket, nil)
	result5 := &PutBucketNotificationReq{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketNotification(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketNotification(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketNotification(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketNotification(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "notification", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketNotification(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutAndGetBucketMirror(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	argsJsonStr := `{
	"bucketMirroringConfiguration":[
		{
			"mode":"fetch", 						
			"sourceUrl":"http://www.baidu.com", 
			"backSourceUrl":"bos://bj.bcebos.com/bucket",   
			"resource" : "folder1/folder2*.jpeg",  
			"prefix": "testprefix",					
			"suffix": ".jpeg",						
			"fixedKey":"folder1/404.jpeg", 
			"version":"v2", 							
			"prefixReplace" : "a/b/c",			
			"passQuerystring":true,					
			"storageClass":"STANDARD",				
			"customHeaders":[
				{ "headerName":"testheader1", "headerValue":"name1" },
				{ "headerName":"TestHeaderName", "headerValue":"TestHeaderValue" }
			],
			"ignoreHeaders" : ["BanHeader1","BanHeader2"],
			"passHeaders" : ["AllowHeader1","AllowHeader2"]
		}
	]}`
	args := &PutBucketMirrorArgs{}
	ExpectEqual(t, nil, json.Unmarshal([]byte(argsJsonStr), args))
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketMirror(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketMirror(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketMirror(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "mirroring", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketMirror(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)

	//GetBucketMirror
	//case1: ok
	AttachMockHttpClientOk(t, client, &argsJsonStr)
	res, err := GetBucketMirror(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 1, len(res.BucketMirroringConfiguration))
	//case2: handle options error
	res, err = GetBucketMirror(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketMirror(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "mirroring", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketMirror(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketMirror(client, bucket, nil)
	result5 := &PutBucketMirrorArgs{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketMirror(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketMirror(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketMirror(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketMirror(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "mirroring", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketMirror(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketTag(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	args := &PutBucketTagArgs{
		Tags: []Tag{
			{TagKey: "key1", TagValue: "value1"},
			{TagKey: "key2", TagValue: "value2"},
		},
	}
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketTag(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketTag(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketTag(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "tagging", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketTag(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketTag(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
        "tag":[
            { "tagKey":"key1", "tagValue":"value123" },
            { "tagKey":"ttt2", "tagValue":"6863gerg" }
        ]
    }`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketTag(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Tags))
	//case2: handle options error
	res, err = GetBucketTag(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketTag(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "tagging", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketTag(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketTag(client, bucket, nil)
	result5 := &GetBucketTagResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketTag(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketTag(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketTag(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketTag(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "tagging", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketTag(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBosShareLink(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: share code invalid
	res, err := GetBosShareLink(client, bucket, "prefix", "111", 180, nil)
	ExpectEqual(t, fmt.Errorf("shareCode length must be 0 or 6"), err)
	ExpectEqual(t, "", res)
	//case2: duration invalid
	res, err = GetBosShareLink(client, bucket, "prefix", "111111", 10, nil)
	ExpectEqual(t, fmt.Errorf("duration must between 1 minute and 18 hours"), err)
	ExpectEqual(t, "", res)
	//case3: ok
	respBody1 := `{"shareUrl":"url","linkExpireTime":180,"shareCode":"111111"}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err = GetBosShareLink(client, bucket, "prefix", "111111", 180, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, respBody1, res)
	//case4: handle options error
	res, err = GetBosShareLink(client, bucket, "prefix", "111111", 180, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", res)
	//case5: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBosShareLink(client, bucket, "prefix", "111111", 180, nil)
	CheckMockHttpClientError(t, "bos-share.baidubce.com", "", "action", "Post", err3, err)
	ExpectEqual(t, "", res)
	//case6: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBosShareLink(client, bucket, "prefix", "111111", 180, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", res)
	//case7: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBosShareLink(client, bucket, "prefix", "111111", 180, nil)
	result5 := &BosShareResBody{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, "", res)
}

func TestPutBucketVersioning(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	args := &BucketVersioningArgs{Status: "enabled"}
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketVersioning(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketVersioning(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketVersioning(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "versioning", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketVersioning(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketVersioning(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{"status": "enabled"}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketVersioning(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "enabled", res.Status)
	//case2: handle options error
	res, err = GetBucketVersioning(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketVersioning(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "versioning", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketVersioning(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketVersioning(client, bucket, nil)
	result5 := &GetBucketTagResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestPutAndGetBucketInventory(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	argsJsonStr := `{
		"inventoryRuleList":[
			{
				"id": "inventory-configuration-ID1", 
				"status": "enabled", 
				"resource": [ "bucket/prefix/*" ], 
				"schedule": "Weekly", 
				"destination": { "targetBucket": "destBucketName", "targetPrefix": "destination-prefix", "format": "CSV" }
			}, 
			{
				"id": "inventory-configuration-ID2", 
				"status": "enabled", 
				"resource": [ "bucket/prefix2/*" ], 
				"schedule": "Monthly", 
				"monthlyDate": 15,
				"destination": { "targetBucket": "destBucketName", "targetPrefix": "destination-prefix-another", "format": "CSV" }
			}
		]
	}`
	listRes := &ListBucketInventoryResult{}
	ExpectEqual(t, nil, json.Unmarshal([]byte(argsJsonStr), listRes))
	ExpectEqual(t, 2, len(listRes.RuleList))
	args := &PutBucketInventoryArgs{Rule: listRes.RuleList[0]}
	jsonStr, err := json.Marshal(args.Rule)
	ExpectEqual(t, nil, err)
	singleRuleStr := string(jsonStr)
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketInventory(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketInventory(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketInventory(client, bucket, args, nil)
	ExpectEqual(t, true, err != nil)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketInventory(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)

	//GetBucketInventory
	//case1: ok
	AttachMockHttpClientOk(t, client, &singleRuleStr)
	res, err := GetBucketInventory(client, bucket, "id", nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "inventory-configuration-ID1", res.Rule.Id)
	//case2: handle options error
	res, err = GetBucketInventory(client, bucket, "id", nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketInventory(client, bucket, "id", nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketInventory(client, bucket, "id", nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketInventory(client, bucket, "id", nil)
	result5 := &PutBucketInventoryArgs{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)

	//ListBucketInventory
	//case1: ok
	AttachMockHttpClientOk(t, client, &argsJsonStr)
	res1, err := ListBucketInventory(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res1.RuleList))
	//case2: handle options error
	res1, err = ListBucketInventory(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res1)
	//case3: send request error
	AttachMockHttpClientError(t, client, err3)
	res1, err = ListBucketInventory(client, bucket, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res1)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res1, err = ListBucketInventory(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res1)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res1, err = ListBucketInventory(client, bucket, nil)
	result51 := &ListBucketInventoryResult{}
	jsonDecoder = json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result51), err)
	ExpectEqual(t, nil, res1)
}

func TestDeleteBucketInventory(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketInventory(client, bucket, "id", nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketInventory(client, bucket, "id", nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketInventory(client, bucket, "id", nil)
	ExpectEqual(t, true, err != nil)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketInventory(client, bucket, "id", nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketQuota(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	args := &BucketQuotaArgs{
		MaxObjectCount:       100,
		MaxCapacityMegaBytes: 99999999,
	}
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketQuota(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketQuota(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketQuota(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "quota", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketQuota(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketQuota(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{ "maxObjectCount": 50,  "maxCapacityMegaBytes"  : 12334424 }`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketQuota(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 50, res.MaxObjectCount)
	ExpectEqual(t, 12334424, res.MaxCapacityMegaBytes)
	//case2: handle options error
	res, err = GetBucketQuota(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketQuota(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "quota", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketQuota(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketQuota(client, bucket, nil)
	result5 := &BucketQuotaArgs{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketQuota(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketQuota(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketQuota(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketQuota(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "quota", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketQuota(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestPutBucketRequestPayment(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	args := &RequestPaymentArgs{RequestPayment: "Requester"}
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutBucketRequestPayment(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutBucketRequestPayment(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutBucketRequestPayment(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "requestPayment", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutBucketRequestPayment(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketRequestPayment(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{"requestPayment": "Requester"}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketRequestPayment(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "Requester", res.RequestPayment)
	//case2: handle options error
	res, err = GetBucketRequestPayment(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketRequestPayment(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "requestPayment", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketRequestPayment(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketRequestPayment(client, bucket, nil)
	result5 := &BucketQuotaArgs{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestInitBucketObjectLock(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	args := &InitBucketObjectLockArgs{
		RetentionDays: 10,
	}
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = InitBucketObjectLock(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = InitBucketObjectLock(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = InitBucketObjectLock(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "objectlock", "Post", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = InitBucketObjectLock(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetBucketObjectLock(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	respBody1 := `{
		"lockStatus": "IN_PROGRESS",
		"createDate": 1569317168,
		"expirationDate": 1569403568,
		"retentionDays": 3
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetBucketObjectLock(client, bucket, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 1569317168, res.CreateDate)
	ExpectEqual(t, 1569403568, res.ExpirationDate)
	ExpectEqual(t, "IN_PROGRESS", res.LockStatus)
	ExpectEqual(t, 3, res.RetentionDays)
	//case2: handle options error
	res, err = GetBucketObjectLock(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetBucketObjectLock(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "objectlock", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetBucketObjectLock(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetBucketObjectLock(client, bucket, nil)
	result5 := &BucketQuotaArgs{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteBucketObjectLock(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteBucketObjectLock(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteBucketObjectLock(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteBucketObjectLock(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "objectlock", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteBucketObjectLock(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestCompleteBucketObjectLock(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = CompleteBucketObjectLock(client, bucket, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = CompleteBucketObjectLock(client, bucket, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = CompleteBucketObjectLock(client, bucket, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "completeobjectlock", "Post", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = CompleteBucketObjectLock(client, bucket, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestExtendBucketObjectLock(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test_bucket"
	args := &ExtendBucketObjectLockArgs{ExtendRetentionDays: 1}
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = ExtendBucketObjectLock(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = ExtendBucketObjectLock(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = ExtendBucketObjectLock(client, bucket, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "extendobjectlock", "Post", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = ExtendBucketObjectLock(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
}
