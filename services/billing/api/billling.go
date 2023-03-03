package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

var (
	checkParamsSet = map[string][]checkParams{
		"must": {
			checkNullOfMonth,
			checkFormatOfMonth,
			checkNullOfProductType,
			checkTypeOfProductType,
		},
		"prefer": {
			checkFormatOfBeginTime,
			checkFormatOfEndTime,
			checkBeginAndEndTimeInMonth,
		},
	}
)

type checkParams func(*BillingParams) error

func GetBilling(cli bce.Client, queryArgs *BillingParams) (*BillingResponse, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBillingUri())
	req.SetMethod(http.GET)

	if queryArgs == nil {
		return nil, errors.New("query parameters must not be null")
	}
	for _, mustCheck := range checkParamsSet["must"] {
		err := mustCheck(queryArgs)
		if err != nil {
			return nil, err
		}
	}
	req.SetParam("month", queryArgs.Month)
	req.SetParam("productType", queryArgs.ProductType)

	if (queryArgs.BeginTime != "") || (queryArgs.EndTime != "") {
		for _, preferCheck := range checkParamsSet["prefer"] {
			err := preferCheck(queryArgs)
			if err != nil {
				return nil, err
			}
		}
		req.SetParam("beginTime", queryArgs.BeginTime)
		req.SetParam("endTime", queryArgs.EndTime)
	}
	if queryArgs.PageNo == 0 {
		queryArgs.PageNo = 1
	}
	if queryArgs.PageSize == 0 {
		queryArgs.PageSize = 20
	}
	req.SetParam("pageNo", strconv.Itoa(queryArgs.PageNo))
	req.SetParam("pageSize", strconv.Itoa(queryArgs.PageSize))

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &BillingResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil

}

// Check begin and end time in the month
func checkBeginAndEndTimeInMonth(bp *BillingParams) error {
	if bp.BeginTime[:7] != bp.Month {
		return errors.New("begin time is not in the month")
	}
	if bp.EndTime[:7] != bp.Month {
		return errors.New("end time is not in the month")
	}
	return nil
}

// Check begin or end time in billing query parameters is legal
func checkFormatOfBeginTime(bp *BillingParams) error {
	_, err := time.Parse("2006-01-02", bp.BeginTime)
	if err != nil {
		return err
	}
	return nil
}

func checkFormatOfEndTime(bp *BillingParams) error {
	_, err := time.Parse("2006-01-02", bp.EndTime)
	if err != nil {
		return err
	}
	return nil
}

// Check month in billing query parameters is legal
func checkNullOfMonth(bp *BillingParams) error {
	if bp.Month == "" {
		return errors.New("month must not be null")
	}
	return nil
}
func checkFormatOfMonth(bp *BillingParams) error {
	_, err := time.Parse("2006-01", bp.Month)
	if err != nil {
		return err
	}
	return nil
}

// Check productType in billing query parameters is legal
func checkNullOfProductType(bp *BillingParams) error {
	if bp.ProductType == "" {
		return errors.New("product type must not be null")
	}
	return nil
}
func checkTypeOfProductType(bp *BillingParams) error {
	if (bp.ProductType == "postpay") || (bp.ProductType == "prepay") {
		return nil
	}
	return errors.New("invalid product type")
}
