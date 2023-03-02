package api

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX          = bce.URI_PREFIX + "v1"
	REQUEST_BALANCE_URL = "/finance/cash/balance"
	REQUEST_BILLING_URL = "/bill/resource/month"
)

func getBalanceUri() string {
	return URI_PREFIX + REQUEST_BALANCE_URL
}

func getBillingUri() string {
	return URI_PREFIX + REQUEST_BILLING_URL
}
