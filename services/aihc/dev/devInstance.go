package dev

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	DefaultVersion = "v2"
	DefaultUrl     = "/"

	ApiCreateDevInstance               = "CreateDevInstance"
	ApiDescribeDevInstance             = "DescribeDevInstance"
	ApiDescribeDevInstances            = "DescribeDevInstances"
	ApiModifyDevInstance               = "ModifyDevInstance"
	ApiStartDevInstance                = "StartDevInstance"
	ApiStopDevInstance                 = "StopDevInstance"
	ApiDeleteDevInstance               = "DeleteDevInstance"
	ApiTimedStopDevInstance            = "TimedStopDevInstance"
	ApiDescribeDevInstanceEvents       = "DescribeDevInstanceEvents"
	ApiCreateDevInstanceImagePackJob   = "CreateDevInstanceImagePackJob"
	ApiDescribeDevInstanceImagePackJob = "DescribeDevInstanceImagePackJob"
)

func CreateDevInstance(cli bce.Client, reqBody *bce.Body) (*CreateDevInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiCreateDevInstance)
	req.SetBody(reqBody)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreateDevInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListDevInstance(cli bce.Client, args *ListDevInstanceArgs) (*ListDevInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeDevInstances)

	if args.PageNumber != 0 {
		req.SetParam("pageNumber", fmt.Sprintf("%d", args.PageNumber))
	}
	if args.PageSize != 0 {
		req.SetParam("pageSize", fmt.Sprintf("%d", args.PageSize))
	}
	if args.QueryKey != "" {
		req.SetParam("queryKey", args.QueryKey)
	}
	if args.QueryVal != "" {
		req.SetParam("queryVal", args.QueryVal)
	}
	if args.QueueName != "" {
		req.SetParam("queueName", args.QueueName)
	}
	if args.ResourcePoolId != "" {
		req.SetParam("resourcePoolId", args.ResourcePoolId)
	}
	if args.OnlyMyDevs {
		req.SetParam("onlyMyDevs", "true")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListDevInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func QueryDevInstanceDetail(cli bce.Client, args *QueryDevInstanceDetailArgs) (*QueryDevInstanceDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeDevInstance)

	req.SetParam("devInstanceId", args.DevInstanceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &QueryDevInstanceDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func UpdateDevInstance(cli bce.Client, reqBody *bce.Body) (*CreateDevInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiModifyDevInstance)
	req.SetBody(reqBody)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreateDevInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func StartDevInstance(cli bce.Client, args *StartDevInstanceArgs) (*StartDevInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiStartDevInstance)

	req.SetParam("devInstanceId", args.DevInstanceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &StartDevInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func StopDevInstance(cli bce.Client, args *StopDevInstanceArgs) (*StopDevInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiStopDevInstance)

	req.SetParam("devInstanceId", args.DevInstanceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &StopDevInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func DeleteDevInstance(cli bce.Client, args *DeleteDevInstanceArgs) (*DeleteDevInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDeleteDevInstance)

	req.SetParam("devInstanceId", args.DevInstanceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeleteDevInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func TimedStopDevInstance(cli bce.Client, reqBody *bce.Body) (*TimedStopDevInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiTimedStopDevInstance)
	req.SetBody(reqBody)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &TimedStopDevInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListDevInstanceEvent(cli bce.Client, args *ListDevInstanceEventArgs) (*ListDevInstanceEventResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeDevInstanceEvents)

	if args.PageNumber != 0 {
		req.SetParam("pageNumber", fmt.Sprintf("%d", args.PageNumber))
	}
	if args.PageSize != 0 {
		req.SetParam("pageSize", fmt.Sprintf("%d", args.PageSize))
	}
	if args.DevInstanceId != "" {
		req.SetParam("devInstanceId", args.DevInstanceId)
	}
	if args.StartTime != "" {
		req.SetParam("startTime", args.StartTime)
	}
	if args.EndTime != "" {
		req.SetParam("endTime", args.EndTime)
	}
	if args.EventType != "" {
		req.SetParam("eventType", args.EventType)
	}
	if args.Message != "" {
		req.SetParam("message", args.Message)
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListDevInstanceEventResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func CreateDevInstanceImagePackJob(cli bce.Client, reqBody *bce.Body) (*CreateDevInstanceImagePackJobResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiCreateDevInstanceImagePackJob)
	req.SetBody(reqBody)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreateDevInstanceImagePackJobResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DevInstanceImagePackJobDetail(cli bce.Client, args *DevInstanceImagePackJobDetailArgs) (*DevInstanceImagePackJobDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeDevInstanceImagePackJob)
	req.SetParam("imagePackJobId", args.ImagePackJobId)
	req.SetParam("devInstanceId", args.DevInstanceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DevInstanceImagePackJobDetailResult{}

	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
