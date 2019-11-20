package dcc

import (
	 "strconv"
	"github.com/baidubce/bce-sdk-go/services/bbc"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// Client used for client
type Client struct {
	*bce.BceClient
}

// NewClient return a client
func NewClient(ak, sk, ep string) (ret *Client, err error) {
	tmp, err := bbc.NewClient(ak, sk, ep)
	return &Client{tmp.BceClient},err
}

// GetDedicatedHost xx
func (c *Client) GetDedicatedHost(args *DedicatedHostArgs) (list *DedicatedHostResult, err error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(bce.URI_PREFIX+"v1/dedicatedHost")
	req.SetMethod(http.GET)

	// Optional arguments settings
	if args != nil {
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(args.MaxKeys))
		}
		if len(args.ZoneName) != 0 {
			req.SetParam("zoneName", args.ZoneName)
		}
	}
	if args == nil || args.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DedicatedHostResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
