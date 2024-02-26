package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

type DigitalWmType string

const (
	DigitalWmTypeImage DigitalWmType = "image"
	DigitalWmTypeText  DigitalWmType = "text"
)

func CreateDigitalWmPreset(cli bce.Client, preset *DigitalWmPreset) (*CreateDigitalWmPresetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, err := json.Marshal(preset)
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
	result := &CreateDigitalWmPresetResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return result, nil
}

func CreateDigitalWmImagePreset(cli bce.Client, digitalWmId string, description string, bucket string, key string,
) (*CreateDigitalWmPresetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	preset := &DigitalWmPreset{}
	preset.DigitalWmId = digitalWmId
	preset.Description = description
	preset.Bucket = bucket
	preset.Key = key
	preset.DigitalWmType = string(DigitalWmTypeImage)
	jsonBytes, err := json.Marshal(preset)
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
	result := &CreateDigitalWmPresetResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return result, nil
}

func CreateDigitalWmTextPreset(cli bce.Client, digitalWmId string, description string,
	textContent string) (*CreateDigitalWmPresetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	preset := &DigitalWmPreset{}
	preset.DigitalWmId = digitalWmId
	preset.Description = description
	preset.TextContent = textContent
	preset.DigitalWmType = string(DigitalWmTypeText)
	jsonBytes, err := json.Marshal(preset)
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
	result := &CreateDigitalWmPresetResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return result, nil
}

func GetDigitalWmPreset(cli bce.Client, digitalWmId string) (*GetDigitalWmPresetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkUrl() + "/" + digitalWmId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetDigitalWmPresetResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListDigitalWmPreset(cli bce.Client) (*ListDigitalWmPresetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkUrl())
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListDigitalWmPresetResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DeleteDigitalWmPreset(cli bce.Client, digitalWmId string) error {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkUrl() + "/" + digitalWmId)
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

func CreateDwmSecretkeyPreset(cli bce.Client, preset *DwmSecretkeyPreset) (*CreateDwmSecretkeyPresetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkSecretkeyUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, err := json.Marshal(preset)
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
	result := &CreateDwmSecretkeyPresetResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return result, nil
}

func GetDwmSecretkeyPreset(cli bce.Client, digitalWmSecretKeyId string) (*GetDwmSecretkeyPresetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkSecretkeyUrl() + "/" + digitalWmSecretKeyId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetDwmSecretkeyPresetResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListDwmSecretkeyPresets(cli bce.Client) (*ListDwmPresetSecretkeyResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkSecretkeyUrl())
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListDwmPresetSecretkeyResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DeleteDwmSecretkeyPreset(cli bce.Client, digitalWmSecretKeyId string) error {
	req := &bce.BceRequest{}
	req.SetUri(getDigitalwatermarkSecretkeyUrl() + "/" + digitalWmSecretKeyId)
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
