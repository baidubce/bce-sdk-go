package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

func UploadMaterial(client bce.Client, matlibUploadRequest *MatlibUploadRequest) (*MatlibUploadResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getUploadMaterial())
	req.SetParam("upload", "")
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, jsonErr := json.Marshal(matlibUploadRequest)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	resp := &bce.BceResponse{}
	if err := client.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &MatlibUploadResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

func SearchMaterial(cli bce.Client, materialSearchRequest *MaterialSearchRequest) (*MaterialSearchResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMaterialLibrary())
	req.SetMethod(http.GET)

	paramsMap := make(map[string]string)
	if materialSearchRequest.End != "" {
		paramsMap["end"] = materialSearchRequest.End
	}
	if materialSearchRequest.Begin != "" {
		paramsMap["begin"] = materialSearchRequest.Begin
	}
	if materialSearchRequest.InfoType != "" {
		paramsMap["infoType"] = materialSearchRequest.InfoType
	}
	if materialSearchRequest.MediaType != "" {
		paramsMap["mediaType"] = materialSearchRequest.MediaType
	}
	if materialSearchRequest.Status != "" {
		paramsMap["status"] = materialSearchRequest.Status
	}
	if materialSearchRequest.SourceType != "" {
		paramsMap["begin"] = materialSearchRequest.Begin
	}
	if materialSearchRequest.TitleKeyword != "" {
		paramsMap["titleKeyword"] = materialSearchRequest.TitleKeyword
	}
	if materialSearchRequest.PageNo >= 1 {
		paramsMap["pageNo"] = strconv.Itoa(materialSearchRequest.PageNo)
	}
	if materialSearchRequest.Size >= 1 {
		paramsMap["size"] = strconv.Itoa(materialSearchRequest.Size)
	}

	req.SetParams(paramsMap)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &MaterialSearchResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetMaterial(cli bce.Client, id string) (*MaterialGetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMaterialLibrary() + "/" + id)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &MaterialGetResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DeleteMaterial(cli bce.Client, id string) error {
	req := &bce.BceRequest{}
	req.SetUri(getMaterialLibrary() + "/" + id)
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

func UploadMaterialPreset(cli bce.Client, fileType string, request *MatlibUploadRequest) (*MaterialPresetUploadResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMateriaPresrtURL() + fileType)
	req.SetParam("upload", "")
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return nil, jsonErr
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
	result := &MaterialPresetUploadResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

func SearchMaterialPreset(cli bce.Client, request *MaterialPresetSearchRequest) (*MaterialPresetSearchResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMateriaPresrtURL())
	req.SetMethod(http.GET)

	paramsMap := make(map[string]string)

	if request.MediaType != "" {
		paramsMap["type"] = request.MediaType
	}
	if request.Status != "" {
		paramsMap["status"] = request.Status
	}
	if request.SourceType != "" {
		paramsMap["sourceType"] = request.SourceType
	}
	if request.PageNo != "" {
		paramsMap["pageNo"] = request.PageNo
	} else {
		paramsMap["pageNo"] = "1"
	}
	if request.PageSize != "" {
		paramsMap["pageSize"] = request.PageSize
	} else {
		paramsMap["pageSize"] = "10"
	}

	req.SetParams(paramsMap)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &MaterialPresetSearchResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetMaterialPreset(cli bce.Client, id string) (*MaterialPresetGetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMateriaPresrtURL() + id)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &MaterialPresetGetResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DeleteMaterialPreset(cli bce.Client, id string) error {
	req := &bce.BceRequest{}
	req.SetUri(getMateriaPresrtURL() + id)
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

func CreateMatlibConfig(cli bce.Client, request *MatlibConfigBaseRequest) (*bce.BceResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMatlibConfigURL())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return nil, jsonErr
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
	return resp, nil
}

func GetMatlibConfig(cli bce.Client) (*MatlibConfigGetResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getMatlibConfigURL())
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &MatlibConfigGetResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func UpdateMatlibConfig(cli bce.Client, request *MatlibConfigUpdateRequest) error {
	req := &bce.BceRequest{}
	req.SetUri(getMatlibConfigURL())
	req.SetMethod(http.PUT)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, jsonErr := json.Marshal(request)
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
