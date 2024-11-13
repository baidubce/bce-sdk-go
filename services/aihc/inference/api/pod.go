package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func ListPod(cli bce.Client, region string, args *ListPodArgs) (*ListPodResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(listPodUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListPodResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func BlockPod(cli bce.Client, region string, args *BlockPodArgs) (*BlockPodResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(blockPodUri())
	req.SetMethod(http.POST)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)
	req.SetParam("insID", args.InsID)
	if args.Block {
		req.SetParam("block", "true")
	} else {
		req.SetParam("block", "false")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &BlockPodResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func DeletePod(cli bce.Client, region string, args *DeletePodArgs) (*DeletePodResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(deletePodUri())
	req.SetMethod(http.POST)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)
	req.SetParam("insID", args.InsID)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeletePodResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
