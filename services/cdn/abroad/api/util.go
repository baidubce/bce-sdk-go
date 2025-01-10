package api

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
)

// SendCustomRequest - send a HTTP request, and response data or error, it use the default times for retrying
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - method: the HTTP requested method, e.g. "GET", "POST", "PUT" ...
//   - urlPath: a path component, consisting of a sequence of path segments separated by a slash ( / ).
//   - params: the query params, which will be append to the query path, and separate by "&"
//     e.g. http://www.baidu.com?query_param1=value1&query_param2=value2
//   - reqHeaders: the request http headers
//   - bodyObj: the HTTP requested body content transferred to a goland object
//   - respObj: the HTTP response content transferred to a goland object
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SendCustomRequest(cli bce.Client, method string, urlPath string, params, reqHeaders map[string]string, bodyObj interface{}, respObj interface{}) error {
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		return errors.New("invalid http method")
	}

	req := &bce.BceRequest{}
	req.SetUri(urlPath)
	req.SetMethod(method)
	req.SetParams(params)
	req.SetHeaders(reqHeaders)

	if bodyObj != nil {
		bodyBytes, err := newBodyFromJsonObj(bodyObj)
		if err != nil {
			return err
		}
		req.SetBody(bodyBytes)
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	defer func() {
		reader := resp.Body()
		if reader != nil {
			_ = reader.Close()
		}
	}()

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

func httpRequest(cli bce.Client, method string, urlPath string, params map[string]string, bodyObj interface{}, respObj interface{}) error {
	return SendCustomRequest(cli, method, urlPath, params, nil, bodyObj, respObj)
}

func newBodyFromJsonObj(obj interface{}) (*bce.Body, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return bce.NewBodyFromBytes(data)
}
