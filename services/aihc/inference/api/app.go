package api

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strings"
)

func CreateApp(cli bce.Client, region string, reqBody *bce.Body, extraInfo map[string]string) (*CreateAppResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(createAppUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	req.SetHeader("X-Region", region)

	if extraInfo != nil {
		if source, ok := extraInfo["source"]; ok {
			req.SetParam("source", source)
		}
		if authToken, ok := extraInfo["authToken"]; ok {
			req.SetParam("authToken", authToken)
		}
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateAppResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ListApp(cli bce.Client, region string, args *ListAppArgs, extraInfo map[string]string) (*ListAppResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(listAppUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

	if extraInfo != nil {
		if source, ok := extraInfo["source"]; ok {
			req.SetParam("source", source)
		}
	}

	if args.PageNo != 0 {
		req.SetParam("pageNo", fmt.Sprintf("%d", args.PageNo))
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
	if len(args.Keyword) != 0 {
		req.SetParam("keyword", args.Keyword)
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListAppResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ListAppStats(cli bce.Client, region string, args *ListAppStatsArgs) (*ListAppStatsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(listAppStatUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

	if len(args.AppIds) != 0 {
		req.SetParam("appIds", strings.Join(args.AppIds, "#"))
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListAppStatsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func AppDetails(cli bce.Client, region string, args *AppDetailsArgs) (*AppDetailsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(appDetailsUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &AppDetailsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func UpdateApp(cli bce.Client, region string, reqBody *bce.Body, args *UpdateAppArgs) (*UpdateAppResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(updateAppUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)
	if len(args.ShortDesc) != 0 {
		req.SetParam("shortDesc", args.ShortDesc)
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &UpdateAppResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ScaleApp(cli bce.Client, region string, args *ScaleAppArgs) (*ScaleAppResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(scaleAppUri())
	req.SetMethod(http.POST)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)
	req.SetParam("insCount", fmt.Sprintf("%d", args.InsCount))

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ScaleAppResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func PubAccess(cli bce.Client, region string, args *PubAccessArgs) (*PubAccessResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(pubAccessUri())
	req.SetMethod(http.POST)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)
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

func ListChange(cli bce.Client, region string, args *ListChangeArgs) (*ListChangeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(listChangeUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)
	if args.PageNo != 0 {
		req.SetParam("pageNo", fmt.Sprintf("%d", args.PageNo))
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

func ChangeDetail(cli bce.Client, region string, args *ChangeDetailArgs) (*ChangeDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(changeDetailUri())
	req.SetMethod(http.GET)
	req.SetHeader("X-Region", region)

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

func DeleteApp(cli bce.Client, region string, args *DeleteAppArgs) (*DeleteAppResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(deleteAppUri())
	req.SetMethod(http.POST)
	req.SetHeader("X-Region", region)

	req.SetParam("appId", args.AppId)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeleteAppResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
