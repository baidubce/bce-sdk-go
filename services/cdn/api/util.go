package api

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
)

// NewBodyFromJsonObj - transfer a goland object to a bce.Body object
//
// PARAMS:
//     - obj: the goland object
// RETURNS:
//     - *bce.Body: the transferred object, nil if error occurred
//     - error: nil if success otherwise the specific error
func NewBodyFromJsonObj(obj interface{}) (*bce.Body, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return bce.NewBodyFromBytes(data)
}

// httpRequest - do a HTTP request, and response data or error, it use the default times for retrying
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - method: the HTTP requested method, e.g. "GET", "POST", "PUT" ...
//     - params: the query params, which will be append to the query path, and separate by "&"
//         e.g. http://www.baidu.com?query_param1=value1&query_param2=value2
//     - bodyObj: the HTTP requested body content transferred to a goland object
//     - respObj: the HTTP response content transferred to a goland object
// RETURNS:
//     - error: nil if success otherwise the specific error
func httpRequest(cli bce.Client, method string, urlPath string, params map[string]string, bodyObj interface{}, respObj interface{}) error {
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		return errors.New("invalid http method")
	}

	req := &bce.BceRequest{}
	req.SetUri(urlPath)
	req.SetMethod(method)
	req.SetParams(params)

	if bodyObj != nil {
		bodyBytes, err := NewBodyFromJsonObj(bodyObj)
		if err != nil {
			return err
		}
		req.SetBody(bodyBytes)
	}

	resp := &bce.BceResponse{}
	defer func() { _ = resp.Body().Close() }()

	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body())
	respBodyBytes := buf.Bytes()

	if respObj != nil {
		err := json.Unmarshal(respBodyBytes, respObj)
		if err != nil {
			return err
		}
	}

	return nil
}
