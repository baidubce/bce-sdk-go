package eccr

import (
	"time"
)

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
	Tags         []Tag     `json:"tags"`
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

type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

type AssignTagsRequest struct {
	Tags []Tag `json:"tags"`
}

type CreateInstanceArgs struct {
	Type   string `json:"type" binding:"required,oneof=BASIC STANDARD ADVANCED"`
	Name   string `json:"name" binding:"required,max=256,min=1"`
	Bucket string `json:"bucket"`

	PaymentTiming string  `json:"paymentTiming" binding:"required,oneof=prepay"`
	Billing       Billing `json:"billing"`

	PaymentMethod []PaymentMethod `json:"paymentMethod"`

	Tags []Tag `json:"tags"`
}

type PaymentMethod struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type Billing struct {
	ReservationTimeUnit string `json:"reservationTimeUnit" binding:"required,oneof=MONTH YEAR"`
	ReservationTime     int    `json:"reservationTime" binding:"required,oneof=1 2 3 4 5 6 7 8 9"`
	AutoRenew           bool   `json:"autoRenew"`
	AutoRenewTimeUnit   string `json:"autoRenewTimeUnit" binding:"oneof=MONTH YEAR"`
	AutoRenewTime       int    `json:"autoRenewTime" binding:"required,oneof=1 2 3 4 5 6 7 8 9"`
}

type CreateInstanceResponse struct {
	InstanceID string `json:"instanceID"`
	OrderID    string `json:"orderID"`
}

type RenewInstanceArgs struct {
	Items []Item `json:"items"`
	// Payment information
	PaymentMethod []PaymentMethod `json:"paymentMethod"`
}

type RenewInstanceResponse struct {
	OrderId string `json:"orderId"`
}

type Item struct {
	Config Config `json:"config"`
	// Payment information
	PaymentMethod []PaymentMethod `json:"paymentMethod"`
}
type Config struct {
	// renewal time
	Duration   int    `json:"duration"`
	InstanceId string `json:"instanceId"`
	// Product name
	ServiceType string `json:"serviceType"`
	// renew time unit 'year' | 'month' | 'day'; default `month`
	TimeUnit string `json:"timeUnit"`
	// Whether it is an order that expires uniformly. default(no): no parameter, yes: 1
	UnionExpireOrderFlag string `json:"unionExpireOrderFlag"`
	// UUID
	UUID string `json:"uuid"`
}

type UpdateInstanceArgs struct {
	Name string `json:"name,omitempty" binding:"required,max=256,min=1"`
}

type UpdateInstanceResponse struct {
	ID           string    `json:"id"`
	InstanceType string    `json:"instanceType"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	CreateTime   time.Time `json:"createTime"`
	Region       string    `json:"region"`
	PublicURL    string    `json:"publicURL"`
	ExpireTime   time.Time `json:"expireTime"`
}

type UpgradeInstanceArgs struct {
	Type string `json:"type" binding:"required,oneof=BASIC STANDARD ADVANCED"`

	PaymentMethod []PaymentMethod `json:"paymentMethod"`
}

type UpgradeInstanceResponse struct {
	OrderID string `json:"orderID"`
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

type CreatePrivateNetworkArgs struct {
	VpcID          string `json:"vpcID,omitempty" binding:"required"`
	SubnetID       string `json:"subnetID,omitempty" binding:"required"`
	IPAddress      string `json:"ipAddress,omitempty"`
	IPType         string `json:"ipType,omitempty"`
	AutoDNSResolve bool   `json:"autoDnsResolve,omitempty"`
}

type DeletePrivateNetworkArgs struct {
	VpcID    string `json:"vpcID,omitempty"`
	SubnetID string `json:"subnetID,omitempty"`
}

type PublicNetworkInfoWhitelist struct {
	IPAddr      string `json:"ipAddr,omitempty" binding:"required,cidrv4|ipv4"`
	Description string `json:"description"`
}

type ListPublicNetworksResponse struct {
	Domain    string                       `json:"domain"`
	Status    string                       `json:"status"`
	Whitelist []PublicNetworkInfoWhitelist `json:"whitelist"`
}

type UpdatePublicNetworkArgs struct {
	Action string `json:"action,omitempty" enums:"open,close" binding:"required,oneof=open close"`
}

type DeletePublicNetworkWhitelistArgs struct {
	Items []string `json:"items,omitempty" binding:"required"`
}

type AddPublicNetworkWhitelistArgs struct {
	IPAddr      string `json:"ipAddr,omitempty" binding:"required,cidrv4|ipv4"`
	Description string `json:"description"`
}

type ResetPasswordArgs struct {
	Password string `json:"password,omitempty" binding:"required,password"`
}

type CreateTemporaryTokenArgs struct {
	Duration int `json:"duration,omitempty" binding:"required,min=1,max=24"`
}

type CreateTemporaryTokenResponse struct {
	Password string `json:"password,omitempty"`
}

type RegistryCredential struct {

	// Access key, e.g. user name when credential type is 'basic'.
	AccessKey string `json:"accessKey"`

	// Access secret, e.g. password when credential type is 'basic'.
	AccessSecret string `json:"accessSecret,omitempty"`

	// Credential type, such as 'basic', 'oauth'.
	Type string `json:"type"`
}

type CreateRegistryArgs struct {
	// credential
	Credential *RegistryCredential `json:"credential"`

	// Description of the registry.
	Description string `json:"description"`

	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure"`

	// The registry name.
	Name string `json:"name"`

	// Type of the registry, e.g. 'harbor'.
	Type string `json:"type" binding:"required,oneof=harbor docker-hub docker-registry baidu-ccr"`

	// The registry URL string.
	URL string `json:"url"`
}

type CreateRegistryResponse struct {
	// The create time of the policy.
	CreationTime string `json:"creation_time,omitempty"`

	// credential
	Credential *RegistryCredential `json:"credential,omitempty"`

	// Description of the registry.
	Description string `json:"description,omitempty"`

	// The registry ID.
	ID int64 `json:"id,omitempty"`

	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure,omitempty"`

	// The registry name.
	Name string `json:"name,omitempty"`

	// Health status of the registry.
	Status string `json:"status,omitempty"`

	// Type of the registry, e.g. 'harbor'.
	Type string `json:"type,omitempty"`

	// The update time of the policy.
	UpdateTime string `json:"update_time,omitempty"`

	// The registry URL string.
	URL string `json:"url,omitempty"`
}

type RegistryResponse struct {
	// id
	ID int64 `json:"id"`

	// creation time
	CreationTime string `json:"creationTime"`

	// credential
	Credential *RegistryCredential `json:"credential"`

	// description
	Description string `json:"description"`

	// insecure
	Insecure bool `json:"insecure"`

	// name
	Name string `json:"name"`

	// status
	Status string `json:"status"`

	// type
	Type string `json:"type"`

	// update time
	UpdateTime string `json:"updateTime"`

	// url
	URL string `json:"url"`
}

type ListRegistriesArgs struct {
	RegistryName string `json:"registryName"`
	RegistryType string `json:"registryType"`
	PageNo       int    `json:"pageNo"`
	PageSize     int    `json:"pageSize"`
}

type ListRegistriesResponse struct {
	PageInfo `json:",inline"`
	Items    []*RegistryResponse `json:"items"`
}

type RegistryRequestArgs struct {
	// credential
	Credential *RegistryCredential `json:"credential"`

	// Description of the registry.
	Description string `json:"description"`

	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure"`

	// The registry name.
	Name string `json:"name"`

	// Type of the registry, e.g. 'harbor'.
	Type string `json:"type" binding:"required,oneof=harbor docker-hub docker-registry baidu-ccr"`

	// The registry URL string.
	URL string `json:"url"`
}

type ListBuildRepositoryTaskArgs struct {
	KeywordType string `json:"keywordType"`
	Keyword     string `json:"keyword"`
	PageNo      int    `json:"pageNo"`
	PageSize    int    `json:"pageSize"`
}

// ListBuildRepositoryTaskResponse list repository task response
type ListBuildRepositoryTaskResponse struct {
	PageInfo `json:",inline"`
	Items    []*BuildRepositoryTaskResult `json:"items"`
}

// BuildRepositoryTaskResult build repository task result
type BuildRepositoryTaskResult struct {
	ID         string    `json:"id"`
	TagName    string    `json:"tagName"`
	IsLatest   bool      `json:"isLatest"`
	Status     string    `json:"status"`
	FromType   string    `json:"fromType"`
	Dockerfile string    `json:"dockerfile"`
	CreateBy   string    `json:"createBy"`
	CreateAt   time.Time `json:"createAt"`
	Image      string    `json:"image"`
}

// BuildRepositoryTaskArgs build repository task request
type BuildRepositoryTaskArgs struct {
	TagName    string `json:"tagName"`
	IsLatest   bool   `json:"isLatest"`
	FromType   string `json:"fromType"`
	Dockerfile string `json:"dockerfile"`
}

// BuildRepositoryTaskResponse build repository task response
type BuildRepositoryTaskResponse struct {
	ID string `json:"id"`
}

// BatchDeleteBuildRepositoryTaskArgs delete task request
type BatchDeleteBuildRepositoryTaskArgs struct {
	Items []string `json:"items"`
}

type ListProjectsArgs struct {
	ProjectName string `json:"projectName"`
	PageNo      int    `json:"pageNo"`
	PageSize    int    `json:"pageSize"`
}

type ProjectResponse struct {
	ChartCount   int64  `json:"chartCount"`
	CreationTime string `json:"creationTime"`
	ProjectName  string `json:"projectName"`
	ProjectID    int64  `json:"projectId"`
	RepoCount    int64  `json:"repoCount"`
	UpdateTime   string `json:"updateTime"`
	AutoScan     string `json:"autoScan"`
	Public       string `json:"public"`
}

type ListProjectsResponse struct {
	PageInfo `json:",inline"`
	Projects []*ProjectResponse `json:"projects"`
}

type CreateProjectArgs struct {
	ProjectName string `json:"projectName"`
	Public      string `json:"public,omitempty"`
}

type BatchDeleteProjectsArgs struct {
	Items []string `json:"items"`
}

type ListRepositoriesArgs struct {
	RepositoryName string `json:"repositoryName"`
	PageNo         int    `json:"pageNo"`
	PageSize       int    `json:"pageSize"`
}

type RepositoryResponse struct {
	TagCount              int64  `json:"tagCount"`
	CreationTime          string `json:"creationTime"`
	Description           string `json:"description"`
	RepositoryName        string `json:"repositoryName"`
	PullCount             int64  `json:"pullCount"`
	UpdateTime            string `json:"updateTime"`
	RepositoryPath        string `json:"repositoryPath"`
	PrivateRepositoryPath string `json:"privateRepositoryPath"`
	ProjectName           string `json:"projectName"`
}

type ListRepositoriesResponse struct {
	PageInfo `json:",inline"`
	Items    []*RepositoryResponse `json:"items"`
}

type UpdateRepositoryArgs struct {
	Description string `json:"description"`
}

type BatchDeleteRepositoriesArgs struct {
	Items []string `json:"items"`
}

type ListTagsArgs struct {
	TagName  string `json:"tagName"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

type TagScanOverview struct {
	ScanStatus string `json:"scanStatus"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	ReportID   string `json:"reportId"`
	Severity   string `json:"severity"`
}

type ImageTag struct {
	TagName           string           `json:"tagName"`
	Digest            string           `json:"digest"`
	ProjectID         int64            `json:"projectId"`
	PullTime          string           `json:"pullTime"`
	PushTime          string           `json:"pushTime"`
	RepositoryID      int64            `json:"repositoryId"`
	Architecture      string           `json:"architecture"`
	OS                string           `json:"os"`
	Author            string           `json:"author"`
	Type              string           `json:"type"`
	Size              int64            `json:"size"`
	ScanOverview      *TagScanOverview `json:"scanOverview"`
	AcceleratorStatus string           `json:"acceleratorStatus"`
	Immutabled        bool             `json:"immutabled"`
	PreventStatus     bool             `json:"preventStatus"`
}

type ListTagsResponse struct {
	PageInfo `json:",inline"`
	Items    []*ImageTag `json:"items"`
}

type BatchDeleteTagsArgs struct {
	Items []string `json:"items"`
}

type ListRobotsArgs struct {
	Status   string `json:"status"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

type RobotAccess struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

type RobotPermission struct {
	Kind      string        `json:"kind"`
	Namespace string        `json:"namespace"`
	Access    []RobotAccess `json:"access"`
}

type Robot struct {
	ID           int64             `json:"id"`
	Name         string            `json:"name"`
	Level        string            `json:"level"`
	Disable      bool              `json:"disable"`
	Description  string            `json:"description"`
	Duration     int               `json:"duration"`
	ExpiresAt    int64             `json:"expiresAt"`
	CreationTime string            `json:"creationTime"`
	UpdateTime   string            `json:"updateTime"`
	Editable     bool              `json:"editable"`
	Permissions  []RobotPermission `json:"permissions"`
}

type ListRobotsResponse struct {
	PageInfo `json:",inline"`
	Robots   []*Robot `json:"robots"`
}

type CreateRobotArgs struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Duration    int               `json:"duration"`
	Disable     bool              `json:"disable"`
	Permissions []RobotPermission `json:"permissions"`
	Secret      string            `json:"secret,omitempty"`
}

type UpdateRobotArgs struct {
	Description string            `json:"description,omitempty"`
	Duration    int               `json:"duration,omitempty"`
	Disable     bool              `json:"disable,omitempty"`
	Permissions []RobotPermission `json:"permissions,omitempty"`
}

type RobotResponse struct {
	ID           int64             `json:"id"`
	Name         string            `json:"name"`
	Level        string            `json:"level"`
	Disable      bool              `json:"disable"`
	Description  string            `json:"description"`
	Duration     int               `json:"duration"`
	ExpiresAt    int64             `json:"expiresAt"`
	CreationTime string            `json:"creationTime"`
	UpdateTime   string            `json:"updateTime"`
	Editable     bool              `json:"editable"`
	Permissions  []RobotPermission `json:"permissions"`
	Secret       string            `json:"secret,omitempty"`
}

type RobotTokenResponse struct {
	Secret string `json:"secret"`
}

type RefreshRobotSecretArgs struct {
	Secret string `json:"secret,omitempty"`
}
