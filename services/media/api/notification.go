package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateNotification(cli bce.Client, name, endpoint string) error {
	notify := &CreateNotificationArgs{}
	notify.Name = name
	notify.Endpoint = endpoint

	req := &bce.BceRequest{}
	req.SetUri(getNotification())
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	jsonBytes, err := json.Marshal(notify)
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

func GetNotification(cli bce.Client, name string) (*CreateNotificationArgs, error) {
	req := &bce.BceRequest{}
	req.SetUri(getNotification() + "/" + name)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreateNotificationArgs{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListNotification(cli bce.Client) (*ListNotificationsResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getNotification())
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListNotificationsResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DeleteNotification(cli bce.Client, name string) error {
	req := &bce.BceRequest{}
	req.SetUri(getNotification() + "/" + name)
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
