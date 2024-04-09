package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func BindInstanceToTagsByResourceType(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(GetServiceTypeTagsUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	req.SetParam("action", "AttachTags")

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func UnbindInstanceToTagsByResourceType(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(GetServiceTypeTagsUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	req.SetParam("action", "DetachTags")

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}
