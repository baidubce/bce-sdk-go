package api

type DescribeInstanceInventoryQuantityReq struct {
	AppType string `json:"appType"`
	AppPerformanceLevel string `json:"appPerformanceLevel"`
	ZoneName string `json:"zoneName"`
	EhcClusterId string `json:"ehcClusterId,omitempty"`
}
