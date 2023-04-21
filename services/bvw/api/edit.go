package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

func CreateDraft(cli bce.Client, request *CreateDraftRequest) (*MatlibTaskResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDraftURL())
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
	result := &MatlibTaskResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetSingleDraft(cli bce.Client, id int) (*GetDraftResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDraftURL() + "/" + strconv.Itoa(id) + "/" + "draftAndTimeline")
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetDraftResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetDraftList(cli bce.Client, request *DraftListRequest) (*ListByPageResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDraftURL())
	req.SetMethod(http.GET)

	paramsMap := make(map[string]string)
	paramsMap["status"] = request.Status
	paramsMap["beginTime"] = request.BeginTime
	paramsMap["endTime"] = request.EndTime
	paramsMap["pageNo"] = strconv.Itoa(request.PageNo)
	paramsMap["pageSize"] = strconv.Itoa(request.PageSize)

	req.SetParams(paramsMap)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListByPageResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func UpdateDraft(cli bce.Client, id int, request *MatlibTaskRequest) error {
	req := &bce.BceRequest{}
	req.SetUri(getDraftURL() + "/" + strconv.Itoa(id) + "/" + "draftAndTimeline")
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

func PollingVideoEdit(cli bce.Client, id int) (*VideoEditPollingResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPollingVideoURL() + strconv.Itoa(id))
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &VideoEditPollingResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func CreateVideoEdit(cli bce.Client, request *VideoEditCreateRequest) (*VideoEditCreateResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getCreateVideoURL())
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
	result := &VideoEditCreateResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}
