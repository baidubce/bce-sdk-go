package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateImagedwmJob(cli bce.Client, digitalWm *Imagedwm) (*CreateJobResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getImagedwmUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, err := json.Marshal(digitalWm)
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
	defer func() { resp.Body().Close() }()
	return result, nil
}

func GetImagedwmResult(cli bce.Client, jobId string) (*GetImagedwmResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getImagedwmUrl() + "/" + jobId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetImagedwmResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
