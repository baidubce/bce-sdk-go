package bill

import "math/big"

type GetResourceMonthBillArgs struct {
	Month          string `json:"month"`
	ProductType    string `json:"ProductType"`
	ServiceType    string `json:"serviceType"`
	QueryAccountId string `json:"queryAccountId"`
	PageNo         int    `json:"pageNo"`
	PageSize       int    `json:"pageSize"`
}

type GetResourceMonthBillResult struct {
	BillMonth    string
	AccountId    string
	LoginName    string
	SubAccountId string
	SubLoginName string
	OuName       string
	PageNo       int
	PageSize     int
	TotalCount   int
	Bills        []ResourceMonthInstanceBill
}
type ResourceMonthInstanceBill struct {
	Vendor              string
	AccountId           string
	ServiceType         string
	ServiceTypeName     string
	ProductType         string
	Region              string
	InstanceId          string
	OrderId             string
	OrderType           string
	OrderTypeDesc       string
	OrderPurchaseTime   string
	StartTime           string
	EndTime             string
	ConfigurationCH     string
	Tag                 string
	Duration            string
	ChargeItem          string
	ChargeItemDesc      string
	Amount              string
	AmountUnit          string
	UnitPrice           string
	PricingUnit         string
	DiscountUnit        string
	Tex                 big.Float
	OriginPrice         big.Float
	FinancePrice        big.Float
	Cash                big.Float
	Rebate              big.Float
	CreditCost          big.Float
	CreditRefund        big.Float
	Debt                big.Float
	NoPaidPrice         big.Float
	CouponPrice         big.Float
	DiscountCouponPrice big.Float
	DiscountPrice       big.Float
	SysGold             big.Float
}
