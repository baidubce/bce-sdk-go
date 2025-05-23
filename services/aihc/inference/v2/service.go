package api

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateService(cli bce.Client, reqBody *bce.Body, clientToken string) (*CreateServiceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiCreateService)
	req.SetParam("clientToken", clientToken)
	req.SetBody(reqBody)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreateServiceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListService(cli bce.Client, args *ListServiceArgs) (*ListServiceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeServices)

	if args.PageNumber != 0 {
		req.SetParam("pageNumber", fmt.Sprintf("%d", args.PageNumber))
	}
	if args.PageSize != 0 {
		req.SetParam("pageSize", fmt.Sprintf("%d", args.PageSize))
	}
	if len(args.OrderBy) != 0 {
		req.SetParam("orderBy", args.OrderBy)
	}
	if len(args.Order) != 0 {
		req.SetParam("order", args.Order)
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListServiceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ListServiceStats(cli bce.Client, args *ListServiceStatsArgs) (*ListServiceStatsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeServiceStatus)

	if len(args.ServiceId) != 0 {
		req.SetParam("serviceId", args.ServiceId)
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListServiceStatsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ServiceDetails(cli bce.Client, args *ServiceDetailsArgs) (*ServiceDetailsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeService)

	req.SetParam("serviceId", args.ServiceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ServiceDetailsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func UpdateService(cli bce.Client, reqBody *bce.Body, args *UpdateServiceArgs) (*UpdateServiceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiModifyService)

	req.SetParam("serviceId", args.ServiceId)
	if len(args.Description) != 0 {
		req.SetParam("description", args.Description)
	}
	req.SetBody(reqBody)
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &UpdateServiceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ScaleService(cli bce.Client, args *ScaleServiceArgs) (*ScaleServiceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiModifyServiceReplicas)

	req.SetParam("serviceId", args.ServiceId)
	req.SetParam("instanceCount", fmt.Sprintf("%d", args.InstanceCount))

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ScaleServiceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func PubAccess(cli bce.Client, args *PubAccessArgs) (*PubAccessResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiModifyServiceNetConfig)

	req.SetParam("serviceId", args.ServiceId)
	if args.PublicAccess {
		req.SetParam("publicAccess", "true")
	} else {
		req.SetParam("publicAccess", "false")
	}

	if len(args.Eip) != 0 {
		req.SetParam("eip", args.Eip)
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &PubAccessResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ListChange(cli bce.Client, args *ListChangeArgs) (*ListChangeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeServiceChangelogs)

	req.SetParam("serviceId", args.ServiceId)
	if args.PageNumber != 0 {
		req.SetParam("pageNumber", fmt.Sprintf("%d", args.PageNumber))
	}
	if args.PageSize != 0 {
		req.SetParam("pageSize", fmt.Sprintf("%d", args.PageSize))
	}
	if len(args.OrderBy) != 0 {
		req.SetParam("orderBy", args.OrderBy)
	}
	if len(args.Order) != 0 {
		req.SetParam("order", args.Order)
	}
	if args.ChangeType != 0 {
		req.SetParam("changeType", fmt.Sprintf("%d", args.ChangeType))
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListChangeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ChangeDetail(cli bce.Client, args *ChangeDetailArgs) (*ChangeDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.GET)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDescribeServiceChangelog)

	req.SetParam("changeId", args.ChangeId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ChangeDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func DeleteService(cli bce.Client, args *DeleteServiceArgs) (*DeleteServiceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(http.POST)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", ApiDeleteService)

	req.SetParam("serviceId", args.ServiceId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeleteServiceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
