package abroad

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cdn/abroad/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "cdn.baidubce.com"
)

// Client of CDN service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the CDN service client with default configuration
// Use `cli.Config.xxx` to access the config or change it to non-default value
func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// SendCustomRequest - send a HTTP request, and response data or error, it use the default times for retrying
//
// PARAMS:
//     - method: the HTTP requested method, e.g. "GET", "POST", "PUT" ...
//     - urlPath: a path component, consisting of a sequence of path segments separated by a slash ( / ).
//     - params: the query params, which will be append to the query path, and separate by "&"
//         e.g. http://www.baidu.com?query_param1=value1&query_param2=value2
//     - reqHeaders: the request http headers
//     - bodyObj: the HTTP requested body content transferred to a goland object
//     - respObj: the HTTP response content transferred to a goland object
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SendCustomRequest(method string, urlPath string, params, reqHeaders map[string]string, bodyObj interface{}, respObj interface{}) error {
	return api.SendCustomRequest(cli, method, urlPath, params, reqHeaders, bodyObj, respObj)
}
