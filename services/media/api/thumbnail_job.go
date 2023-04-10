package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateThumbnailJob(cli bce.Client, pipelineName, sourceKey string, createThumbnialArgs *CreateThumbnailJobArgs) (
	*CreateJobResponse, error) {

	createThumbnialArgs.PipelineName = pipelineName
	source := &ThumbnailSource{}
	source.Key = sourceKey
	createThumbnialArgs.ThumbnailSource = source
	req := &bce.BceRequest{}
	req.SetUri(getThumbnailUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	jsonBytes, err := json.Marshal(createThumbnialArgs)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &CreateJobResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetThumbanilJob(cli bce.Client, jobId string) (*GetThumbnailJobResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getThumbnailUrl() + "/" + jobId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetThumbnailJobResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListThumbnailJobs(cli bce.Client, pipelineName string) (*ListThumbnailJobsResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getThumbnailUrl())
	req.SetMethod(http.GET)
	paramMap := make(map[string]string)
	paramMap["pipelineName"] = pipelineName
	req.SetParams(paramMap)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListThumbnailJobsResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
