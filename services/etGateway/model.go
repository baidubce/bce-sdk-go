package etGateway

type (
	HealthCheckType string
)

const (
	HEALTH_CHECK_ICMP HealthCheckType = "ICMP"
)

type CreateEtGatewayArgs struct {
	Name        string   `json:"name"`
	VpcId       string   `json:"vpcId"`
	Speed       int      `json:"speed"`
	Description string   `json:"description"`
	EtId        string   `json:"etId"`
	ChannelId   string   `json:"channelId"`
	LocalCidrs  []string `json:"localCidrs"`
	ClientToken string   `json:"clientToken,omitempty"`
}

type CreateEtGatewayResult struct {
	EtGatewayId string `json:"etGatewayId"`
}

type ListEtGatewayArgs struct {
	VpcId       string
	EtGatewayId string
	Name        string
	Status      string
	Marker      string
	MaxKeys     int
}

type ListEtGatewayResult struct {
	EtGateways  []EtGateway `json:"etGateways"`
	Marker      string      `json:"marker"`
	IsTruncated bool        `json:"isTruncated"`
	NextMarker  string      `json:"nextMarker"`
	MaxKeys     int         `json:"maxKeys"`
}
type EtGateway struct {
	EtGatewayId string   `json:"etGatewayId"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Speed       int      `json:"speed"`
	CreateTime  string   `json:"createTime"`
	Description string   `json:"description"`
	VpcId       string   `json:"vpcId"`
	EtId        string   `json:"etId"`
	ChannelId   string   `json:"channelId"`
	LocalCidrs  []string `json:"localCidrs"`
}

type EtGatewayDetail struct {
	EtGatewayId         string   `json:"etGatewayId"`
	Name                string   `json:"name"`
	Status              string   `json:"status"`
	Speed               int      `json:"speed"`
	CreateTime          string   `json:"createTime"`
	Description         string   `json:"description"`
	VpcId               string   `json:"vpcId"`
	EtId                string   `json:"etId"`
	ChannelId           string   `json:"channelId"`
	LocalCidrs          []string `json:"localCidrs"`
	HealthCheckSourceIp string   `json:"healthCheckSourceIp"`
	HealthCheckDestIp   string   `json:"healthCheckDestIp"`
	HealthCheckType     string   `json:"healthCheckType"`
	HealthCheckInterval int      `json:"healthCheckInterval"`
	HealthThreshold     int      `json:"healthThreshold"`
	UnhealthThreshold   int      `json:"unhealthThreshold"`
}

//  参数localCidrs只有在专线网关处于running状态时允许更新。
type UpdateEtGatewayArgs struct {
	ClientToken string   `json:"clientToken,omitempty"`
	EtGatewayId string   `json:"etGatewayId"`
	Name        string   `json:"name,omitempty"`
	Speed       int      `json:"speed,omitempty"`
	Description string   `json:"description,omitempty"`
	LocalCidrs  []string `json:"localCidrs,omitempty"`
}
type BindEtArgs struct {
	ClientToken string   `json:"clientToken,omitempty"`
	EtGatewayId string   `json:"etGatewayId"`
	EtId        string   `json:"etId"`
	ChannelId   string   `json:"channelId"`
	LocalCidrs  []string `json:"localCidrs,omitempty"`
}

type CreateHealthCheckArgs struct {
	ClientToken           string          `json:"clientToken,omitempty"`
	EtGatewayId           string          `json:"etGatewayId"`
	HealthCheckSourceIp   string          `json:"healthCheckSourceIp,omitempty"`
	HealthCheckType       HealthCheckType `json:"healthCheckType,omitempty"`
	HealthCheckPort       int             `json:"healthCheckPort,omitempty"`
	HealthCheckInterval   int             `json:"healthCheckInterval"`
	HealthThreshold       int             `json:"healthThreshold"`
	UnhealthThreshold     int             `json:"unhealthThreshold"`
	AutoGenerateRouteRule *bool           `json:"autoGenerateRouteRule,omitempty"`
}
