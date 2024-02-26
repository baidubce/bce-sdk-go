package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateDwmDetectJob(cli bce.Client, dwmdetect *Dwmdetect) (*CreateJobResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDwmdetectUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, err := json.Marshal(dwmdetect)
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

func GetDwmdetectResult(cli bce.Client, jobId string) (*GetDwmdetectResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDwmdetectUrl() + "/" + jobId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetDwmdetectResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
