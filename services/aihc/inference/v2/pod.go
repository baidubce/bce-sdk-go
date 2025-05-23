package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func ListPod(cli bce.Client, args *ListPodArgs) (*ListPodResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeServicePods)

	req.SetParam("serviceId", args.ServiceId)

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

func BlockPod(cli bce.Client, args *BlockPodArgs) (*BlockPodResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDisableServicePod)

	req.SetParam("serviceId", args.ServiceId)
	req.SetParam("instanceId", args.InstanceId)
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

func DeletePod(cli bce.Client, args *DeletePodArgs) (*DeletePodResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDeleteServicePod)

	req.SetParam("serviceId", args.ServiceId)
	req.SetParam("instanceId", args.InstanceId)

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

func ListPodGroups(cli bce.Client, args *ListPodGroupsArgs) (*ListPodGroupsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeServicePodGroups)

	req.SetParam("serviceId", args.ServiceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListPodGroupsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
