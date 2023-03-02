package api

import "time"

type BalanceResponse struct {
	CashBalance float64 `json:"cashBalance"`
}
type BillingParams struct {
	Month          string
	BeginTime      string
	EndTime        string
	ProductType    string
	ServiceType    string
	QueryAccountId string
	PageNo         int
	PageSize       int
}

type BillingResponse struct {
	BillMonth    string  `json:"billMonth"`
	BeginTime    string  `json:"beginTime"`
	EndTime      string  `json:"endTime"`
	AccountID    string  `json:"accountId"`
	LoginName    string  `json:"loginName"`
	SubAccountID string  `json:"subAccountId"`
	SubLoginName string  `json:"subLoginName"`
	OuName       string  `json:"ouName"`
	PageNo       int     `json:"pageNo"`
	PageSize     int     `json:"pageSize"`
	TotalCount   int     `json:"totalCount"`
	Bills        []Bills `json:"bills"`
}
type Bills struct {
	Vendor              string    `json:"vendor"`
	AccountID           string    `json:"accountId"`
	ServiceType         string    `json:"serviceType"`
	ServiceTypeName     string    `json:"serviceTypeName"`
	ProductType         string    `json:"productType"`
	Region              string    `json:"region"`
	InstanceID          string    `json:"instanceId"`
	OrderID             string    `json:"orderId"`
	OrderType           string    `json:"orderType"`
	OrderTypeDesc       string    `json:"orderTypeDesc"`
	OrderPurchaseTime   time.Time `json:"orderPurchaseTime"`
	StartTime           time.Time `json:"startTime"`
	EndTime             time.Time `json:"endTime"`
	ConfigurationCH     string    `json:"configurationCH"`
	Tag                 string    `json:"tag"`
	Duration            string    `json:"duration"`
	ChargeItem          string    `json:"chargeItem"`
	ChargeItemDesc      string    `json:"chargeItemDesc"`
	Amount              string    `json:"amount"`
	AmountUnit          string    `json:"amountUnit"`
	DiscountAmount      string    `json:"discountAmount"`
	UnitPrice           string    `json:"unitPrice"`
	PricingUnit         string    `json:"pricingUnit"`
	DiscountUnit        string    `json:"discountUnit"`
	Tex                 float64   `json:"tex"`
	CatalogPrice        float64   `json:"catalogPrice"`
	OriginPrice         float64   `json:"originPrice"`
	FinancePrice        float64   `json:"financePrice"`
	Cash                float64   `json:"cash"`
	Rebate              float64   `json:"rebate"`
	CreditCost          float64   `json:"creditCost"`
	CreditRefund        float64   `json:"creditRefund"`
	Debt                float64   `json:"debt"`
	NoPaidPrice         float64   `json:"noPaidPrice"`
	CouponPrice         float64   `json:"couponPrice"`
	DiscountCouponPrice float64   `json:"discountCouponPrice"`
	DiscountPrice       float64   `json:"discountPrice"`
	SysGold             float64   `json:"sysGold"`
}
