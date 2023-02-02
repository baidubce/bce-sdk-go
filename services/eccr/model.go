package eccr

import "time"

type PagedListOption struct {
	PageNo      int
	PageSize    int
	KeywordType string
	Keyword     string
}

type PageInfo struct {
	Total    int `json:"total"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type ListInstancesArgs struct {
	PageNo       int
	PageSize     int
	KeywordType  string
	Keyword      string
	Acrossregion string
}

type ListInstancesResponse struct {
	PageInfo  `json:",inline"`
	Instances []*InstanceInfo `json:"instances"`
}

type InstanceInfo struct {
	ID           string    `json:"id"`
	InstanceType string    `json:"instanceType"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	CreateTime   time.Time `json:"createTime"`
	Region       string    `json:"region"`
	PublicURL    string    `json:"publicURL"`
	ExpireTime   time.Time `json:"expireTime"`
}

type InstanceStatistic struct {
	Repo      int64 `json:"repo"`
	Chart     int64 `json:"chart"`
	Namespace int64 `json:"namespace"`
	Storage   int64 `json:"storage"`
}

type UserQuota struct {
	Repo      int64 `json:"repo"`
	Chart     int64 `json:"chart"`
	Namespace int64 `json:"namespace"`
}

type GetInstanceDetailResponse struct {
	Info      *InstanceInfo     `json:"info,omitempty"`
	Statistic InstanceStatistic `json:"statistic,omitempty"`
	Quota     UserQuota         `json:"quota,omitempty"`
	Bucket    string            `json:"bucket,omitempty"`
	Region    string            `json:"region,omitempty"`
}

type ListPrivateNetworksResponse struct {
	Domain string                 `json:"domain"`
	Items  []PrivateNetworksItems `json:"items"`
}

type PrivateNetworksItems struct {
	VpcID          string `json:"vpcID,omitempty"`
	SubnetID       string `json:"subnetID,omitempty"`
	ServiceNetID   string `json:"serviceNetID,omitempty"`
	Status         string `json:"status,omitempty"`
	IPAddress      string `json:"ipAddress,omitempty"`
	ResourceSource string `json:"resourceSource,omitempty"`
}
