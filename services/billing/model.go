package billing

type ResourceMonthBillResponse struct {
	BillMonth    string              `json:"billMonth"`
	BeginTime    string              `json:"beginTime"`
	EndTime      string              `json:"endTime"`
	AccountId    string              `json:"accountId"`
	LoginName    string              `json:"loginName"`
	SubAccountId string              `json:"subAccountId"`
	SubLoginName string              `json:"subLoginName"`
	OuName       string              `json:"ouName"`
	PageNo       int                 `json:"pageNo"`
	PageSize     int                 `json:"pageSize"`
	TotalCount   int                 `json:"totalCount"`
	Bills        []ResourceMonthBill `json:"bills"`
}

type ResourceMonthBill struct {
	Vendor                          string  `json:"vendor"`
	BillId                          string  `json:"billId"`
	AccountId                       string  `json:"accountId"`
	ServiceType                     string  `json:"serviceType"`
	ServiceTypeName                 string  `json:"serviceTypeName"`
	ProductType                     string  `json:"productType"`
	Region                          string  `json:"region"`
	InstanceId                      string  `json:"instanceId"`
	InstanceName                    string  `json:"instanceName"`
	OrderId                         string  `json:"orderId"`
	OrderType                       string  `json:"orderType"`
	OrderTypeDesc                   string  `json:"orderTypeDesc"`
	OrderPurchaseTime               string  `json:"orderPurchaseTime"`
	StartTime                       string  `json:"startTime"`
	EndTime                         string  `json:"endTime"`
	ServiceStartTime                string  `json:"serviceStartTime"`
	ServiceEndTime                  string  `json:"serviceEndTime"`
	ConfigurationCH                 string  `json:"configurationCH"`
	Tag                             string  `json:"tag"`
	Duration                        string  `json:"duration"`
	ChargeItem                      string  `json:"chargeItem"`
	ChargeItemDesc                  string  `json:"chargeItemDesc"`
	Amount                          string  `json:"amount"`
	AmountUnit                      string  `json:"amountUnit"`
	DiscountAmount                  string  `json:"discountAmount"`
	UnitPrice                       string  `json:"unitPrice"`
	PricingUnit                     string  `json:"pricingUnit"`
	CatalogUnitPrice                string  `json:"catalogUnitPrice"`
	DiscountUnit                    string  `json:"discountUnit"`
	Tex                             string  `json:"tex"`
	OriginPrice                     float64 `json:"originPrice"`
	CatalogPrice                    float64 `json:"catalogPrice"`
	FinancePrice                    float64 `json:"financePrice"`
	Cash                            float64 `json:"cash"`
	Rebate                          float64 `json:"rebate"`
	CreditCost                      float64 `json:"creditCost"`
	CreditRefund                    float64 `json:"creditRefund"`
	CreditRefundDeduct              float64 `json:"creditRefundDeduct"`
	Debt                            float64 `json:"debt"`
	NoPaidPrice                     float64 `json:"noPaidPrice"`
	CouponPrice                     float64 `json:"couponPrice"`
	DiscountCouponPrice             float64 `json:"discountCouponPrice"`
	CashEquivalentCouponPrice       float64 `json:"cashEquivalentCouponPrice"`
	DiscountPrice                   float64 `json:"discountPrice"`
	SysGold                         float64 `json:"sysGold"`
	DeductPrice                     float64 `json:"deductPrice"`
	DeductCash                      float64 `json:"deductCash"`
	DeductRebate                    float64 `json:"deductRebate"`
	DeductCreditCost                float64 `json:"deductCreditCost"`
	DeductCouponPrice               float64 `json:"deductCouponPrice"`
	DeductDiscountCouponPrice       float64 `json:"deductDiscountCouponPrice"`
	DeductDiscountPrice             float64 `json:"deductDiscountPrice"`
	DeductCashEquivalentCouponPrice float64 `json:"deductCashEquivalentCouponPrice"`
}
