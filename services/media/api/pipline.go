package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreatePipeline(cli bce.Client, pipelineName, sourceBucket, targetBucket string, capacity int) error {
	req := &bce.BceRequest{}
	args := &CreatePiplineArgs{}

	args.PipelineName = pipelineName
	args.SourceBucket = sourceBucket
	args.TargetBucket = targetBucket

	config := &CreatePiplineConfig{}
	config.Capacity = capacity

	args.Config = config
	req.SetUri(getPipLineUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

func CreatePipelineCustomize() {

}

func ListPipelines(cli bce.Client) (*ListPipelinesResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPipLineUrl())
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListPipelinesResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetPipeline(cli bce.Client, pipelineName string) (*PipelineStatus, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPipLineUrl() + "/" + pipelineName)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &PipelineStatus{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetPipelineUpdate(cli bce.Client, pipelineName string) (*UpdatePipelineArgs, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPipLineUrl() + "/" + pipelineName)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &UpdatePipelineArgs{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DeletePipeline(cli bce.Client, pipelineName string) error {
	req := &bce.BceRequest{}
	req.SetUri(getPipLineUrl() + "/" + pipelineName)
	req.SetMethod(http.DELETE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

func UpdatePipeline(cli bce.Client, pipelineName string, updatePipelineArgs *UpdatePipelineArgs) error {
	req := &bce.BceRequest{}
	jsonBytes, err := json.Marshal(updatePipelineArgs)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetUri(getPipLineUrl() + "/" + pipelineName)
	req.SetMethod(http.PUT)
	req.SetBody(body)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}
