// nolint
package api

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

// sendMockRequest drives the real SendRequest choke point against a mock http client and
// returns the response together with the error, so the tests below exercise exactly the
// code path taken by every BOS api function.
func sendMockRequest(t *testing.T, bucket string, options []Option,
	mockOptions ...util.MockRoundTripperOption) (*BosResponse, error) {
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	client.HTTPClient = util.NewMockHTTPClient(mockOptions...)

	req, resp := &BosRequest{}, &BosResponse{}
	req.SetUri("/v1/" + bucket)
	req.SetMethod(my_http.GET)
	req.SetBucket(bucket)
	ExpectEqual(t, nil, handleOptions(req, options))
	return resp, SendRequest(client, req, resp, newDefaultBosContext())
}

func mockOkOptions(requestId, debugId string) []util.MockRoundTripperOption {
	return []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.SetRespBody("{}"),
		util.AddHeaders(map[string]string{
			my_http.BCE_REQUEST_ID: requestId,
			my_http.BCE_DEBUG_ID:   debugId,
		}),
	}
}

func TestWithResponseCommonNil(t *testing.T) {
	// a nil sink yields a nil option, which every option consumer already skips
	ExpectEqual(t, true, WithResponseCommon(nil) == nil)

	req := &BosRequest{}
	ExpectEqual(t, nil, handleOptions(req, []Option{WithResponseCommon(nil)}))
	ExpectEqual(t, true, req.ResponseCommon == nil)
}

func TestWithResponseCommonHandleOptions(t *testing.T) {
	meta := &ResponseCommon{}

	req := &BosRequest{}
	ExpectEqual(t, nil, handleOptions(req, []Option{WithResponseCommon(meta)}))
	ExpectEqual(t, true, req.ResponseCommon == meta)
	// the metadata option must not leak into the request headers
	ExpectEqual(t, "", req.Header(RESPONSE_METADATA_KEY))

	ctx := newDefaultBosContext()
	ExpectEqual(t, nil, handleBosContextOptions(ctx, []Option{WithResponseCommon(meta)}))
	ExpectEqual(t, true, ctx.ResponseCommon == meta)
}

// TestResponseCommonNoRoundTrip pins the contract that a failure happening before any
// http response is received never overwrites the value held by the caller. It would also
// panic if responseMetadataReady used resp.IsFail() instead of resp.StatusCode().
func TestResponseCommonNoRoundTrip(t *testing.T) {
	meta := &ResponseCommon{RequestId: "preexisting", DebugId: "dbg", StatusCode: 999}
	fillResponseCommon(&BosRequest{ResponseCommon: meta}, newDefaultBosContext(), &BosResponse{})
	ExpectEqual(t, "preexisting", meta.RequestId)
	ExpectEqual(t, "dbg", meta.DebugId)
	ExpectEqual(t, 999, meta.StatusCode)

	// nil sinks must be tolerated everywhere
	fillResponseCommon(nil, nil, &BosResponse{})
	fillResponseCommon(&BosRequest{}, newDefaultBosContext(), nil)
	retrieveResponseFields(&ListBucketsResult{}, &BosResponse{})
	var nilMeta *ResponseCommon
	nilMeta.setResponseCommon(&BosResponse{})
}

func TestFillResponseCommonPrecedence(t *testing.T) {
	resp, err := sendMockRequest(t, "test-bucket", nil, mockOkOptions("req-1", "dbg-1")...)
	ExpectEqual(t, nil, err)

	// the request scoped sink wins over the context scoped one
	reqMeta, ctxMeta := &ResponseCommon{}, &ResponseCommon{}
	ctx := newDefaultBosContext()
	ctx.ResponseCommon = ctxMeta
	fillResponseCommon(&BosRequest{ResponseCommon: reqMeta}, ctx, resp)
	ExpectEqual(t, "req-1", reqMeta.RequestId)
	ExpectEqual(t, "dbg-1", reqMeta.DebugId)
	ExpectEqual(t, 200, reqMeta.StatusCode)
	ExpectEqual(t, "", ctxMeta.RequestId)

	// falling back to the context scoped sink when the request has none
	fillResponseCommon(&BosRequest{}, ctx, resp)
	ExpectEqual(t, "req-1", ctxMeta.RequestId)
}

// TestSendRequestFillsMetadata covers both the success path and the http failure path.
func TestSendRequestFillsMetadata(t *testing.T) {
	cases := []struct {
		statusCode int
		statusMsg  string
		requestId  string
		expectFail bool
	}{
		{200, "200 OK", "req-ok", false},
		{403, "403 Forbidden", "req-denied", true},
	}
	for _, c := range cases {
		meta := &ResponseCommon{}
		mockOptions := []util.MockRoundTripperOption{
			util.SetStatusCode(c.statusCode),
			util.SetStatusMsg(c.statusMsg),
			util.SetRespBody("{}"),
			util.AddHeaders(map[string]string{
				my_http.BCE_REQUEST_ID: c.requestId,
				my_http.BCE_DEBUG_ID:   "dbg",
			}),
		}
		_, err := sendMockRequest(t, "test-bucket",
			[]Option{WithResponseCommon(meta)}, mockOptions...)
		if c.expectFail {
			serviceErr, ok := err.(*bce.BceServiceError)
			ExpectEqual(t, true, ok)
			ExpectEqual(t, serviceErr.RequestId, meta.RequestId)
		} else {
			ExpectEqual(t, nil, err)
		}
		ExpectEqual(t, c.requestId, meta.RequestId)
		ExpectEqual(t, "dbg", meta.DebugId)
		ExpectEqual(t, c.statusCode, meta.StatusCode)
	}
}

// TestSendRequestMetadataRetryKeepsLast pins the "last attempt wins" semantics: the first
// attempt gets a retryable 500, the second one succeeds, and the metadata must hold the
// values of the second response.
func TestSendRequestMetadataRetryKeepsLast(t *testing.T) {
	meta := &ResponseCommon{}
	mockOptions := []util.MockRoundTripperOption{
		util.AppendStatusCode([]int{500, 200}),
		util.AppendStatusMsg([]string{"500 Internal Server Error", "200 OK"}),
		util.SetRespBody("{}"),
		util.AppendHeaders([]map[string]string{
			{my_http.BCE_REQUEST_ID: "req-first", my_http.BCE_DEBUG_ID: "dbg-first"},
			{my_http.BCE_REQUEST_ID: "req-last", my_http.BCE_DEBUG_ID: "dbg-last"},
		}),
	}
	_, err := sendMockRequest(t, "test-bucket",
		[]Option{WithResponseCommon(meta)}, mockOptions...)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "req-last", meta.RequestId)
	ExpectEqual(t, "dbg-last", meta.DebugId)
	ExpectEqual(t, 200, meta.StatusCode)
}

// TestSendRequestMetadataValidationError checks that the validations running before the
// request is sent leave the caller supplied struct untouched.
func TestSendRequestMetadataValidationError(t *testing.T) {
	meta := &ResponseCommon{RequestId: "untouched", StatusCode: 999}
	_, err := sendMockRequest(t, "INVALID_BUCKET",
		[]Option{WithResponseCommon(meta)}, mockOkOptions("req-1", "dbg-1")...)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, "untouched", meta.RequestId)
	ExpectEqual(t, 999, meta.StatusCode)
}

// newMetadataApiClient returns a client answering every request with the given body and
// with the request id / debug id headers set.
func newMetadataApiClient(t *testing.T, respBody string) *bce.BceClient {
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	client.HTTPClient = util.NewMockHTTPClient(append(mockOkOptions("api-req", "api-dbg"),
		util.SetRespBody(respBody))...)
	return client
}

// TestResultStructsCarryMetadata pins Layer 2 on one representative of every embedding
// shape: a plain result, a result also embedding ObjectMeta, a defined type of a struct
// embedding ResponseCommon (where the promotion has to survive the type definition),
// and a request struct reused as a result.
func TestResultStructsCarryMetadata(t *testing.T) {
	cases := []struct {
		name string
		call func(cli bce.Client) (ResponseCommon, error)
	}{
		{"ListBuckets", func(cli bce.Client) (ResponseCommon, error) {
			res, err := ListBuckets(cli, newDefaultBosContext())
			return res.ResponseCommon, err
		}},
		{"GetObjectMeta", func(cli bce.Client) (ResponseCommon, error) {
			res, err := GetObjectMeta(cli, "test-bucket", "test-object", newDefaultBosContext())
			return res.ResponseCommon, err
		}},
		{"GetObjectAcl", func(cli bce.Client) (ResponseCommon, error) {
			res, err := GetObjectAcl(cli, "test-bucket", "test-object", newDefaultBosContext())
			return res.ResponseCommon, err
		}},
		{"GetBucketMirror", func(cli bce.Client) (ResponseCommon, error) {
			res, err := GetBucketMirror(cli, "test-bucket", newDefaultBosContext())
			return res.ResponseCommon, err
		}},
		{"ListParts", func(cli bce.Client) (ResponseCommon, error) {
			res, err := ListParts(cli, "test-bucket", "test-object", "upload-id", nil,
				newDefaultBosContext())
			return res.ResponseCommon, err
		}},
	}
	for _, c := range cases {
		meta, err := c.call(newMetadataApiClient(t, "{}"))
		ExpectEqual(t, nil, err)
		if meta.RequestId != "api-req" || meta.DebugId != "api-dbg" || meta.StatusCode != 200 {
			t.Errorf("%s does not carry the metadata: got %+v", c.name, meta)
		}
	}
}

// TestCopyObjectResultMetadata covers the first shadowing exception: CopyObjectResult
// keeps its own RequestId field, filled from the response header as before, and it still
// serializes it.
func TestCopyObjectResultMetadata(t *testing.T) {
	res, err := CopyObject(newMetadataApiClient(t, "{}"), "test-bucket", "test-object",
		"/src-bucket/src-object", nil, newDefaultBosContext())
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "api-req", res.RequestId)
	ExpectEqual(t, "api-dbg", res.DebugId)
	ExpectEqual(t, 200, res.StatusCode)

	raw, err := json.Marshal(&CopyObjectResult{RequestId: "kept"})
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, strings.Contains(string(raw), `"requestId":"kept"`))
	ExpectEqual(t, false, strings.Contains(string(raw), "statusCode"))
}

// TestFetchObjectResultMetadata covers the second shadowing exception: the RequestId of
// FetchObjectResult comes from the response body, which must keep winning over the header.
func TestFetchObjectResultMetadata(t *testing.T) {
	res, err := FetchObject(newMetadataApiClient(t, `{"requestId":"body-req","jobId":"job"}`),
		"test-bucket", "test-object", "http://example.com/src", nil, newDefaultBosContext())
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "body-req", res.RequestId)
	ExpectEqual(t, "job", res.JobId)
	ExpectEqual(t, "api-dbg", res.DebugId)
	ExpectEqual(t, 200, res.StatusCode)
}

// TestMetadataNotSerialized pins that embedding ResponseCommon into the structs also
// used as request payloads leaves the bytes on the wire untouched.
func TestMetadataNotSerialized(t *testing.T) {
	meta := ResponseCommon{RequestId: "req", DebugId: "dbg", StatusCode: 200}
	payloads := []interface{}{
		&PutBucketNotificationReq{ResponseCommon: meta},
		&PutBucketMirrorArgs{ResponseCommon: meta},
		&BucketVersioningArgs{ResponseCommon: meta, Status: "enabled"},
		&PutBucketInventoryArgs{ResponseCommon: meta},
		&BucketQuotaArgs{ResponseCommon: meta},
		&RequestPaymentArgs{ResponseCommon: meta, RequestPayment: "Requester"},
		&UserQuotaArgs{ResponseCommon: meta},
		&PutBucketReplicationArgs{ResponseCommon: meta},
		&PutBucketStaticWebsiteArgs{ResponseCommon: meta},
		&PutObjectAclArgs{ResponseCommon: meta},
		&PutBucketLifecycleArgs{},
		&CopyObjectArgs{},
	}
	for _, payload := range payloads {
		raw, err := json.Marshal(payload)
		ExpectEqual(t, nil, err)
		lower := strings.ToLower(string(raw))
		for _, leak := range []string{"requestid", "debugid", "statuscode", "responsecommon"} {
			if strings.Contains(lower, leak) {
				t.Errorf("%T leaks %q into the payload: %s", payload, leak, raw)
			}
		}
	}
}
