package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
)

func TestPutUserQuota(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	args := &UserQuotaArgs{
		MaxBucketCount:       100,
		MaxCapacityMegaBytes: 999999999,
	}
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutUserQuota(client, args, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = PutUserQuota(client, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutUserQuota(client, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "", "userQuota", "Put", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = PutUserQuota(client, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
}

func TestGetUserQuota(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	//case1: ok
	respBody1 := `{
		"maxBucketCount" : 50,  
		"maxCapacityMegaBytes" : 12334424 
	}`
	AttachMockHttpClientOk(t, client, &respBody1)
	res, err := GetUserQuota(client, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 50, res.MaxBucketCount)
	ExpectEqual(t, 12334424, res.MaxCapacityMegaBytes)
	//case2: handle options error
	res, err = GetUserQuota(client, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetUserQuota(client, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "", "userQuota", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = GetUserQuota(client, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetUserQuota(client, nil)
	result5 := &BucketQuotaArgs{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)
}

func TestDeleteUserQuota(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	//case1: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteUserQuota(client, nil)
	ExpectEqual(t, nil, err)
	//case2: handle options error
	err = DeleteUserQuota(client, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = DeleteUserQuota(client, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "", "userQuota", "Delete", err3, err)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	err = DeleteUserQuota(client, nil)
	ExpectEqual(t, bceServiceErro404, err)
}
