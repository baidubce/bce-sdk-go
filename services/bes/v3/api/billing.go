package api

import (
	"errors"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// ConvertToPrepay converts a postpaid cluster to prepaid billing.
func ConvertToPrepay(cli bce.Client, region string, request *ConvertToPrepayRequest) (*OrderIdResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("convert to prepay request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("convert to prepay request clusterId should not be empty")
	}
	if request.TimeLength <= 0 {
		return nil, errors.New("convert to prepay request timeLength should be greater than 0")
	}
	region = setRegion(region, request.Region)

	uri := URI_CLUSTERS + "/" + URI_ORDERS + "/" + URI_TO_PREPAY

	result := &OrderIdResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ConvertToPostpay converts a prepaid cluster to postpaid billing.
func ConvertToPostpay(cli bce.Client, region string, request *ConvertToPostpayRequest) (*OrderIdResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("convert to postpay request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("convert to postpay request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := URI_CLUSTERS + "/" + URI_ORDERS + "/" + URI_TO_POSTPAY

	result := &OrderIdResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CancelConvertToPostpay cancels a pending convert-to-postpay order.
func CancelConvertToPostpay(cli bce.Client, region string, request *CancelConvertToPostpayRequest) (*OrderIdResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("cancel convert to postpay request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("cancel convert to postpay request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := URI_CLUSTERS + "/" + URI_ORDERS + "/" + URI_TO_POSTPAY + "/" + URI_CANCEL

	result := &OrderIdResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RenewCluster creates a renewal order for a prepaid cluster.
func RenewCluster(cli bce.Client, region string, request *RenewClusterRequest) (*OrderIdResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("renew cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("renew cluster request clusterId should not be empty")
	}
	if request.TimeLength <= 0 {
		return nil, errors.New("renew cluster request timeLength should be greater than 0")
	}
	region = setRegion(region, request.Region)

	result := &OrderIdResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, URI_ORDER_RENEW_CONFIRM, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryConfigPrice queries the price of a hypothetical cluster configuration before purchase.
func QueryConfigPrice(cli bce.Client, region string, request *QueryConfigPriceRequest) (*QueryConfigPriceResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("query config price request should not be nil")
	}
	if request.Payment == "" {
		return nil, errors.New("query config price request payment should not be empty")
	}
	if len(request.NodeSpecs) == 0 {
		return nil, errors.New("query config price request nodeSpecs should not be empty")
	}
	region = setRegion(region, request.Region)

	result := &QueryConfigPriceResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, URI_ORDER_PRICES, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryClusterMarginPrice queries the price delta for resizing an existing cluster.
func QueryClusterMarginPrice(cli bce.Client, region string, request *QueryClusterMarginPriceRequest) (*QueryClusterMarginPriceResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("query cluster margin price request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("query cluster margin price request clusterId should not be empty")
	}
	if len(request.NodeSpecs) == 0 {
		return nil, errors.New("query cluster margin price request nodeSpecs should not be empty")
	}
	if request.Mode == "" {
		return nil, errors.New("query cluster margin price request mode should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := URI_CLUSTERS + "/" + URI_MARGIN_PRICES

	result := &QueryClusterMarginPriceResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
