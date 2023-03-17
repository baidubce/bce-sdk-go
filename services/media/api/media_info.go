package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func GetMediaInfoOfFile(cli bce.Client, bucket, key string) (*GetMediaInfoOfFileResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMediaInfoUrl())
	req.SetMethod(http.GET)
	paramMap := make(map[string]string)
	paramMap["bucket"] = bucket
	paramMap["key"] = key
	req.SetParams(paramMap)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetMediaInfoOfFileResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
