package api

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func ListBriefResPool(cli bce.Client, region string, args *ListBriefResPoolArgs) (*ListBriefResPoolResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(listBriefResPoolUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

	if args.PageNo != 0 {
		req.SetParam("pageNo", fmt.Sprintf("%d", args.PageNo))
	}
	if args.PageSize != 0 {
		req.SetParam("pageSize", fmt.Sprintf("%d", args.PageSize))
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListBriefResPoolResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ResPoolDetail(cli bce.Client, region string, args *ResPoolDetailArgs) (*ResPoolDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(resPoolDetailUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

	req.SetParam("resPoolId", args.ResPoolId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ResPoolDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
