package api

// OrderIdResponse contains the order id returned by billing order operations
// (convert to prepay/postpay, cancel convert to postpay, renew cluster).
type OrderIdResponse struct {
	OrderId string `json:"orderId,omitempty"`
}

// ConvertToPrepayRequest contains the request body for POST /v3/clusters/orders/to-prepay.
type ConvertToPrepayRequest struct {
	Request
	ClusterId  string `json:"clusterId"`
	TimeLength int    `json:"timeLength"`
	TimeUnit   string `json:"timeUnit,omitempty"`
}

// ConvertToPostpayRequest contains the request body for POST /v3/clusters/orders/to-postpay.
type ConvertToPostpayRequest struct {
	Request
	ClusterId string `json:"clusterId"`
}

// CancelConvertToPostpayRequest contains the request body for POST /v3/clusters/orders/to-postpay/cancel.
type CancelConvertToPostpayRequest struct {
	Request
	ClusterId string `json:"clusterId"`
}

// RenewClusterRequest contains the request body for POST /v3/order/renew/confirm.
type RenewClusterRequest struct {
	Request
	ClusterId  string `json:"clusterId"`
	TimeLength int    `json:"timeLength"`
	TimeUnit   string `json:"timeUnit,omitempty"`
}

// BillingConfig defines billing settings referenced by billing and resize requests.
type BillingConfig struct {
	Payment        string `json:"payment,omitempty"`
	TimeLength     *int   `json:"timeLength,omitempty"`
	TimeUnit       string `json:"timeUnit,omitempty"`
	ExpirationTime *int64 `json:"expirationTime,omitempty"`
	IsAutoPay      *bool  `json:"isAutoPay,omitempty"`
}

// ComponentPrice describes the price of a billing component.
type ComponentPrice struct {
	Price        float64 `json:"price,omitempty"`
	CatalogPrice float64 `json:"catalogPrice,omitempty"`
	DiscountRate float64 `json:"discountRate,omitempty"`
}

// QueryNodeTypePrice describes the price detail of a node role.
type QueryNodeTypePrice struct {
	Type             string  `json:"type,omitempty"`
	NodePrice        float64 `json:"nodePrice,omitempty"`
	NodeCatalogPrice float64 `json:"nodeCatalogPrice,omitempty"`
	DiskPrice        float64 `json:"diskPrice,omitempty"`
	DiskCatalogPrice float64 `json:"diskCatalogPrice,omitempty"`
}

// QueryConfigPriceRequest contains the request body for POST /v3/orders/prices.
type QueryConfigPriceRequest struct {
	Request
	Payment    string     `json:"payment"`
	TimeLength int        `json:"timeLength,omitempty"`
	TimeUnit   string     `json:"timeUnit,omitempty"`
	NodeSpecs  []NodeSpec `json:"nodeSpecs"`
}

// QueryConfigPriceResponse contains the response body for POST /v3/orders/prices.
type QueryConfigPriceResponse struct {
	Price          float64              `json:"price,omitempty"`
	CatalogPrice   float64              `json:"catalogPrice,omitempty"`
	NodeTypePrices []QueryNodeTypePrice `json:"nodeTypePrices,omitempty"`
}

// QueryClusterMarginPriceRequest contains the request body for POST /v3/clusters/margin-prices.
type QueryClusterMarginPriceRequest struct {
	Request
	ClusterId       string         `json:"clusterId"`
	ResizeType      *int           `json:"resizeType,omitempty"`
	BillingConfig   *BillingConfig `json:"billingConfig,omitempty"`
	LogicalZones    []string       `json:"logicalZones"`
	VpcId           string         `json:"vpcId"`
	SubnetId        string         `json:"subnetId"`
	EnableDeploySet *bool          `json:"enableDeploySet"`
	NodeSpecs       []NodeSpec     `json:"nodeSpecs"`
	Mode            string         `json:"mode"`
}

// QueryClusterMarginPriceResponse contains the response body for POST /v3/clusters/margin-prices.
type QueryClusterMarginPriceResponse struct {
	CurrentComponentPrices map[string]ComponentPrice `json:"currentComponentPrices,omitempty"`
	UpdateComponentPrices  map[string]ComponentPrice `json:"updateComponentPrices,omitempty"`
	MarginComponentPrices  map[string]ComponentPrice `json:"marginComponentPrices,omitempty"`
	ExpirationTime         string                    `json:"expirationTime,omitempty"`
}
