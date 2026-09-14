package api

import (
	"bytes"
	"encoding/json"

	bcehttp "github.com/baidubce/bce-sdk-go/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// SuccessFlag is a boolean flag that can be represented as a string.
type SuccessFlag string

const (
	SuccessFlagTrue  SuccessFlag = "true"
	SuccessFlagFalse SuccessFlag = "false"
)

// UnmarshalJSON accepts a JSON boolean, a quoted string, or null.
func (s *SuccessFlag) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		*s = ""
		return nil
	}

	var asBool bool
	if err := json.Unmarshal(trimmed, &asBool); err == nil {
		if asBool {
			*s = SuccessFlagTrue
		} else {
			*s = SuccessFlagFalse
		}
		return nil
	}

	var asString string
	if err := json.Unmarshal(trimmed, &asString); err != nil {
		return err
	}
	*s = SuccessFlag(asString)
	return nil
}

// Bool reports whether the flag means success.
func (s SuccessFlag) Bool() bool {
	return s == SuccessFlagTrue
}

func isEmptyJSONResult(data []byte) bool {
	trimmed := bytes.TrimSpace(data)
	return len(trimmed) == 0 ||
		bytes.Equal(trimmed, []byte("null")) ||
		bytes.Equal(trimmed, []byte(`""`))
}

func createJSONRequest(cli bce.Client, region, method, uri string, payload interface{}, resp interface{}) error {
	builder := bce.NewRequestBuilder(cli).
		WithMethod(method).
		WithURL(uri).
		WithHeader(bcehttp.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithHeader(HEADER_REGION, region)

	if payload != nil {
		builder = builder.WithBody(payload)
	}
	if resp != nil {
		builder = builder.WithResult(resp)
	}
	return builder.Do()
}

// createMultipartRequest sends a multipart/form-data request. RequestBuilder cannot be reused
// here because it always JSON-marshals the body; this path sets the body explicitly instead.
func createMultipartRequest(cli bce.Client, region, method, uri string, body *bce.Body,
	contentType string, resp interface{}) error {
	req := &bce.BceRequest{}
	req.SetMethod(method)
	req.SetUri(uri)
	req.SetHeader(bcehttp.CONTENT_TYPE, contentType)
	req.SetHeader(HEADER_REGION, region)
	req.SetBody(body)

	bceResp := &bce.BceResponse{}
	if err := cli.SendRequest(req, bceResp); err != nil {
		return err
	}
	if bceResp.IsFail() {
		return bceResp.ServiceError()
	}
	defer bceResp.Body().Close()

	if resp == nil {
		return nil
	}
	return bceResp.ParseJsonBody(resp)
}
