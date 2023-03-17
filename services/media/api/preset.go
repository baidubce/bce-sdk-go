package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func ListPresets(cli bce.Client) (*ListPresetsResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPresetUrl())
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListPresetsResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetPreset(cli bce.Client, presetName string) (*Preset, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPresetUrl() + "/" + presetName)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &Preset{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil

}

func CreatePreset(cli bce.Client, presetName, description, container string) error {
	req := &bce.BceRequest{}
	req.SetUri(getPresetUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	args := &Preset{}
	args.PresetName = presetName
	args.Description = description
	args.Container = container
	args.Transmux = true
	extraCfg := &ExtraCfg{}
	args.ExtraCfg = extraCfg

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
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

func CreatePrestCustomize(cli bce.Client, preset *Preset) error {
	req := &bce.BceRequest{}
	req.SetUri(getPresetUrl())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, err := json.Marshal(preset)
	if err != nil {
		return err
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

func UpdatePreset(cli bce.Client, preset *Preset) error {
	presetName := preset.PresetName
	req := &bce.BceRequest{}
	req.SetUri(getPresetUrl() + "/" + presetName)
	req.SetMethod(http.PUT)
	jsonBytes, err := json.Marshal(preset)
	if err != nil {
		return err
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
