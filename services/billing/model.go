package billing

type ResourceChargeItemBillRequest struct {
	BillMonth              string `json:"billMonth,omitempty"`
	ServiceType            string `json:"serviceType,omitempty"`
	QueryAccountId         string `json:"queryAccountId,omitempty"`
	PageNo                 int    `json:"pageNo"`
	PageSize               int    `json:"pageSize"`
	NeedSplitConfiguration bool   `json:"needSplitConfiguration"`
}

type ShareBillRequest struct {
	Month                  string `json:"month,omitempty"`
	StartDay               string `json:"startDay,omitempty"`
	EndDay                 string `json:"endDay,omitempty"`
	ProductType            string `json:"productType,omitempty"`
	ServiceType            string `json:"serviceType,omitempty"`
	QueryAccountId         string `json:"queryAccountId,omitempty"`
	ShowDeductPrice        bool   `json:"showDeductPrice"`
	ShowControversial      bool   `json:"showControversial"`
	ShowTags               bool   `json:"showTags"`
	DisplaySystemUnit      string `json:"displaySystemUnit,omitempty"`
	PageNo                 int    `json:"pageNo"`
	PageSize               int    `json:"pageSize"`
	NeedSplitConfiguration bool   `json:"needSplitConfiguration"`
}

type CostSplitBillRequest struct {
	Month                  string `json:"month,omitempty"`
	StartDay               string `json:"startDay,omitempty"`
	EndDay                 string `json:"endDay,omitempty"`
	ServiceType            string `json:"serviceType,omitempty"`
	QueryAccountId         string `json:"queryAccountId,omitempty"`
	InstanceId             string `json:"instanceId,omitempty"`
	ShowTags               bool   `json:"showTags"`
	PageNo                 int    `json:"pageNo"`
	PageSize               int    `json:"pageSize"`
	NeedSplitConfiguration bool   `json:"needSplitConfiguration"`
}

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
	Tex                             float64 `json:"tex"`
	OriginPrice                     float64 `json:"originPrice"`
	CatalogPrice                    float64 `json:"catalogPrice"`
	PromotionPrice                  float64 `json:"promotionPrice"`
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
	DeductPromotionPrice            float64 `json:"deductPromotionPrice"`
	DeductCash                      float64 `json:"deductCash"`
	DeductRebate                    float64 `json:"deductRebate"`
	DeductCreditCost                float64 `json:"deductCreditCost"`
	DeductCouponPrice               float64 `json:"deductCouponPrice"`
	DeductDiscountCouponPrice       float64 `json:"deductDiscountCouponPrice"`
	DeductDiscountPrice             float64 `json:"deductDiscountPrice"`
	DeductCashEquivalentCouponPrice float64 `json:"deductCashEquivalentCouponPrice"`
}

type ResourceChargeItemBillResponse struct {
	BeginTime    string                   `json:"beginTime"`
	EndTime      string                   `json:"endTime"`
	AccountId    string                   `json:"accountId"`
	LoginName    string                   `json:"loginName"`
	SubAccountId string                   `json:"subAccountId"`
	SubLoginName string                   `json:"subLoginName"`
	OuName       string                   `json:"ouName"`
	PageNo       int                      `json:"pageNo"`
	PageSize     int                      `json:"pageSize"`
	TotalCount   int                      `json:"totalCount"`
	Bills        []ResourceChargeItemBill `json:"bills"`
}

type ResourceChargeItemBill struct {
	Vendor                    string            `json:"vendor"`
	BillId                    string            `json:"billId"`
	ServiceType               string            `json:"serviceType"`
	ServiceTypeName           string            `json:"serviceTypeName"`
	ProductType               string            `json:"productType"`
	Region                    string            `json:"region"`
	InstanceId                string            `json:"instanceId"`
	ShortId                   string            `json:"shortId"`
	StartTime                 string            `json:"startTime"`
	EndTime                   string            `json:"endTime"`
	ConfigurationCH           string            `json:"configurationCH"`
	UnitPrice                 string            `json:"unitPrice"`
	PricingUnit               string            `json:"pricingUnit"`
	CatalogUnitPrice          string            `json:"catalogUnitPrice"`
	CatalogPricingUnit        string            `json:"catalogPricingUnit"`
	ChargeItem                string            `json:"chargeItem"`
	ChargeItemDesc            string            `json:"chargeItemDesc"`
	Amount                    string            `json:"amount"`
	AmountUnit                string            `json:"amountUnit"`
	DiscountAmount            string            `json:"discountAmount"`
	PromotionPrice            float64           `json:"promotionPrice"`
	OriginPrice               float64           `json:"originPrice"`
	CatalogPrice              float64           `json:"catalogPrice"`
	FinancePrice              float64           `json:"financePrice"`
	CouponPrice               float64           `json:"couponPrice"`
	DiscountCouponPrice       float64           `json:"discountCouponPrice"`
	CashEquivalentCouponPrice float64           `json:"cashEquivalentCouponPrice"`
	DiscountPrice             float64           `json:"discountPrice"`
	SysGold                   float64           `json:"sysGold"`
	DeductPrice               float64           `json:"deductPrice"`
	SplitConfiguration        map[string]string `json:"splitConfiguration"`
}

type ShareBillResponse struct {
	BillMonth    string      `json:"billMonth"`
	BeginTime    string      `json:"beginTime"`
	EndTime      string      `json:"endTime"`
	AccountId    string      `json:"accountId"`
	LoginName    string      `json:"loginName"`
	SubAccountId string      `json:"subAccountId"`
	SubLoginName string      `json:"subLoginName"`
	OuName       string      `json:"ouName"`
	PageNo       int         `json:"pageNo"`
	PageSize     int         `json:"pageSize"`
	TotalCount   int         `json:"totalCount"`
	Bills        []ShareBill `json:"bills"`
}

type ShareBill struct {
	BillMonth                            string            `json:"billMonth"`
	BillDay                              string            `json:"billDay"`
	ProductType                          string            `json:"productType"`
	ProductTypeName                      string            `json:"productTypeName"`
	ServiceType                          string            `json:"serviceType"`
	ServiceTypeName                      string            `json:"serviceTypeName"`
	SourceServiceType                    string            `json:"sourceServiceType"`
	SourceServiceTypeName                string            `json:"sourceServiceTypeName"`
	SourceInstanceId                     string            `json:"sourceInstanceId"`
	SourceShortId                        string            `json:"sourceShortId"`
	InstanceId                           string            `json:"instanceId"`
	ShortId                              string            `json:"shortId"`
	InstanceName                         string            `json:"instanceName"`
	ConfigurationCH                      string            `json:"configurationCH"`
	SourceConfigurationCH                string            `json:"sourceConfigurationCH"`
	ChargeItem                           string            `json:"chargeItem"`
	ChargeItemDesc                       string            `json:"chargeItemDesc"`
	Tag                                  string            `json:"tag"`
	Region                               string            `json:"region"`
	RegionName                           string            `json:"regionName"`
	OrderId                              string            `json:"orderId"`
	OrderType                            string            `json:"orderType"`
	OrderTypeDesc                        string            `json:"orderTypeDesc"`
	CreateTime                           string            `json:"createTime"`
	ServiceStartTime                     string            `json:"serviceStartTime"`
	ServiceEndTime                       string            `json:"serviceEndTime"`
	ServiceTimeSpan                      string            `json:"serviceTimeSpan"`
	Capacity                             string            `json:"capacity"`
	Price                                float64           `json:"price"`
	FinancePrice                         float64           `json:"financePrice"`
	CatalogPrice                         float64           `json:"catalogPrice"`
	CouponPrice                          float64           `json:"couponPrice"`
	DiscountCouponPrice                  float64           `json:"discountCouponPrice"`
	CashEquivalentCouponPrice            float64           `json:"cashEquivalentCouponPrice"`
	DiscountPrice                        float64           `json:"discountPrice"`
	SysGoldPrice                         float64           `json:"sysGoldPrice"`
	SharePrice                           float64           `json:"sharePrice"`
	ShareFinancePrice                    float64           `json:"shareFinancePrice"`
	SharePromotionPrice                  float64           `json:"sharePromotionPrice"`
	ShareCatalogPrice                    float64           `json:"shareCatalogPrice"`
	ShareCouponPrice                     float64           `json:"shareCouponPrice"`
	ShareDiscountCouponPrice             float64           `json:"shareDiscountCouponPrice"`
	ShareCashEquivalentCouponPrice       float64           `json:"shareCashEquivalentCouponPrice"`
	ShareDiscountPrice                   float64           `json:"shareDiscountPrice"`
	ShareSysGoldPrice                    float64           `json:"shareSysGoldPrice"`
	DeductShareFinancePrice              float64           `json:"deductShareFinancePrice"`
	DeductSharePromotionPrice            float64           `json:"deductSharePromotionPrice"`
	DeductShareCouponPrice               float64           `json:"deductShareCouponPrice"`
	DeductShareDiscountCouponPrice       float64           `json:"deductShareDiscountCouponPrice"`
	DeductShareCashEquivalentCouponPrice float64           `json:"deductShareCashEquivalentCouponPrice"`
	DeductShareDiscountPrice             float64           `json:"deductShareDiscountPrice"`
	DeductPackageFinancePrice            float64           `json:"deductPackageFinancePrice"`
	ControversialItem                    int               `json:"controversialItem"`
	OpTime                               string            `json:"opTime"`
	ConfirmTime                          string            `json:"confirmTime"`
	ShareAmount                          float64           `json:"shareAmount"`
	AmountUnit                           string            `json:"amountUnit"`
	ShareDays                            int               `json:"shareDays"`
	JsonFlavor                           string            `json:"jsonFlavor"`
	SkuId                                string            `json:"skuId"`
	PsiCode                              string            `json:"psiCode"`
	ProductCategory                      string            `json:"productCategory"`
	SplitConfiguration                   map[string]string `json:"splitConfiguration"`
}
