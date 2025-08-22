package api

type CreateCombinedCouponReq struct {
	Name                       string `json:"name,omitempty"`
	PurchaseNum                int    `json:"purchaseNum,omitempty"`
	AutoRenew                  bool   `json:"autoRenew,omitempty"`
	AutoRenewPeriod            int    `json:"autoRenewPeriod,omitempty"`
	AutoRenewPeriodUnit        string `json:"autoRenewPeriodUnit,omitempty"`
	ReservedInstancePeriod     int    `json:"reservedInstancePeriod,omitempty"`
	ReservedInstancePeriodUnit string `json:"reservedInstancePeriodUnit,omitempty"`
}
