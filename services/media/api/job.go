package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateJob(cli bce.Client, pipelineName, sourceKey, targetKey, presetName string) (*CreateJobResponse, error) {
	args := &CreateJobArgs{}
	args.PipelineName = pipelineName

	source := &Source{}
	source.SourceKey = sourceKey

	target := &Target{}
	target.TargetKey = targetKey
	target.PresetName = presetName

	args.Source = source
	args.Target = target

	req := &bce.BceRequest{}
	req.SetUri(getTrandCodingJobUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, jsonErr := bce.NewBodyFromBytes(jsonBytes)
	if jsonErr != nil {
		return nil, jsonErr
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
	defer func() { resp.Body().Close() }()
	return result, nil

}

func CreateJobCustomize(cli bce.Client, args *CreateJobArgs) (*CreateJobResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getTrandCodingJobUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, jsonErr := bce.NewBodyFromBytes(jsonBytes)
	if jsonErr != nil {
		return nil, jsonErr
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
	defer func() { resp.Body().Close() }()
	return result, nil
}

func ListTranscodingJobs(cli bce.Client, pipelineName string) (*ListTranscodingJobsResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getTrandCodingJobUrl())
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListTranscodingJobsResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetTranscodingJob(cli bce.Client, jobId string) (*GetTranscodingJobResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getTrandCodingJobUrl() + "/" + jobId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetTranscodingJobResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
