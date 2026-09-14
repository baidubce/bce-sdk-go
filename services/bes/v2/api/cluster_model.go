package api

import "encoding/json"

// ModuleInfo describes a node template used when creating a cluster.
type ModuleInfo struct {
	Type         string        `json:"type"`
	InstanceNum  int           `json:"instanceNum"`
	SlotType     string        `json:"slotType"`
	DiskSlotInfo *DiskSlotInfo `json:"diskSlotInfo,omitempty"`
}

// DiskSlotInfo describes the disk configuration for a cluster node.
type DiskSlotInfo struct {
	Size       int    `json:"size"`
	Type       string `json:"type"`
	CdsExtraIo int    `json:"cdsExtraIo,omitempty"`
}

// AutoRenewInfo describes the auto-renew configuration used during cluster creation.
type AutoRenewInfo struct {
	RenewTimeUnit string  `json:"renewTimeUnit,omitempty"`
	RenewTime     float64 `json:"renewTime,omitempty"`
}

// Billing describes the billing configuration used during cluster creation.
type Billing struct {
	PaymentType     string         `json:"paymentType"`
	Time            int            `json:"time,omitempty"`
	EnableAutoRenew bool           `json:"enableAutoRenew,omitempty"`
	AutoRenewInfo   *AutoRenewInfo `json:"autoRenewInfo,omitempty"`
	Coupon          string         `json:"coupon,omitempty"`
}

// Tag is a key/value tag attached to a cluster.
type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

// CreateClusterRequest is the request of creating a BES cluster.
type CreateClusterRequest struct {
	Name            string       `json:"name"`
	Password        string       `json:"password"`
	SecurityGroupId string       `json:"securityGroupId"`
	SubnetUuid      string       `json:"subnetUuid"`
	AvailableZone   string       `json:"availableZone"`
	VpcId           string       `json:"vpcId"`
	IsOldPackage    bool         `json:"isOldPackage,omitempty"`
	Version         string       `json:"version"`
	Modules         []ModuleInfo `json:"modules"`
	Billing         *Billing     `json:"billing"`
	Tags            []Tag        `json:"tags,omitempty"`
	EnableSSL       bool         `json:"enableSSL,omitempty"`
	ResGroupId      string       `json:"resGroupId,omitempty"`
}

// CreateClusterResponse is the response of creating a BES cluster.
type CreateClusterResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// StartClusterRequest is the request of starting a BES cluster.
type StartClusterRequest struct {
	ClusterId string `json:"clusterId"`
}

// StartClusterResponse is the response of starting a BES cluster.
type StartClusterResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// ListClustersRequest is the request of listing BES clusters.
type ListClustersRequest struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

// ListClustersResponse is the response of listing BES clusters.
type ListClustersResponse struct {
	Success bool         `json:"success"`
	Status  int          `json:"status"`
	Page    *ClusterPage `json:"page"`
}

// ClusterPage is the paged cluster list carried by ListClustersResponse.
type ClusterPage struct {
	PageNo     int               `json:"pageNo"`
	PageSize   int               `json:"pageSize"`
	TotalCount int               `json:"totalCount"`
	Result     []ClusterListItem `json:"result"`
}

// ClusterListItem is a single cluster entry in the cluster list.
type ClusterListItem struct {
	ClusterId     string                 `json:"clusterId"`
	ClusterName   string                 `json:"clusterName"`
	CreateTime    string                 `json:"createTime"`
	ActualStatus  string                 `json:"actualStatus"`
	RunningTime   string                 `json:"runningTime"`
	Region        string                 `json:"region"`
	Version       string                 `json:"version"`
	ClusterHealth *ClusterHealthSummary  `json:"clusterHealth,omitempty"`
	Billing       *BillingSummary        `json:"billing,omitempty"`
	Tags          []Tag                  `json:"tags,omitempty"`
	ResGroupList  []ResGroup             `json:"resGroupList,omitempty"`
	InnerAccessIp string                 `json:"innerAccessIp,omitempty"`
	AccessEip     string                 `json:"accessEip,omitempty"`
	VpcId         string                 `json:"vpcId,omitempty"`
	Subnet        []ClusterSubnetSummary `json:"subnet,omitempty"`
	Network       []NetworkSummary       `json:"network,omitempty"`
}

// ClusterHealthSummary is the simplified cluster health used by the list response.
type ClusterHealthSummary struct {
	Status string `json:"status"`
}

// BillingSummary is the simplified billing info used by the list/detail responses.
type BillingSummary struct {
	PaymentType string `json:"paymentType"`
}

// ResGroup describes a resource group associated with a cluster.
type ResGroup struct {
	GroupId   string `json:"groupId"`
	GroupName string `json:"groupName"`
}

// ClusterSubnetSummary is the simplified subnet info used by the list response.
type ClusterSubnetSummary struct {
	AvailableZone string `json:"availableZone"`
	SubnetName    string `json:"subnetName"`
}

// NetworkSummary is the simplified network info used by the list response.
type NetworkSummary struct {
	SubnetId string `json:"subnetId"`
}

// NetworkInfo is the full network info used by the detail response.
type NetworkInfo struct {
	SubnetId      string `json:"subnetId"`
	Subnet        string `json:"subnet,omitempty"`
	AvailableZone string `json:"availableZone,omitempty"`
}

// GetClusterDetailRequest is the request of getting a BES cluster detail.
type GetClusterDetailRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetClusterDetailResponse is the response of getting a BES cluster detail.
type GetClusterDetailResponse struct {
	Success bool           `json:"success"`
	Status  int            `json:"status"`
	Result  *ClusterDetail `json:"result"`
}

// ClusterDetail is the detailed cluster information.
type ClusterDetail struct {
	ExpireTime        string                `json:"expireTime,omitempty"`
	ClusterId         string                `json:"clusterId"`
	ClusterName       string                `json:"clusterName"`
	AdminUsername     string                `json:"adminUsername"`
	ActualStatus      string                `json:"actualStatus"`
	DesireStatus      string                `json:"desireStatus"`
	EsUrl             string                `json:"esUrl"`
	KibanaUrl         string                `json:"kibanaUrl"`
	ExporterUrl       string                `json:"exporterUrl"`
	EnableEsSSL       bool                  `json:"enableEsSSL"`
	EnableKibanaSSL   bool                  `json:"enableKibanaSSL"`
	EsEip             string                `json:"esEip,omitempty"`
	KibanaEip         string                `json:"kibanaEip,omitempty"`
	Modules           []ClusterModuleDetail `json:"modules,omitempty"`
	Instances         []ClusterInstance     `json:"instances,omitempty"`
	Region            string                `json:"region"`
	InspectAuthorized SuccessFlag           `json:"inspectAuthorized,omitempty"`
	Vpc               string                `json:"vpc,omitempty"`
	VpcId             string                `json:"vpcId,omitempty"`
	Subnet            string                `json:"subnet,omitempty"`
	AvailableZone     string                `json:"availableZone,omitempty"`
	SecurityGroup     string                `json:"securityGroup,omitempty"`
	ClusterHealth     *ClusterHealthDetail  `json:"clusterHealth,omitempty"`
	Billing           *BillingSummary       `json:"billing,omitempty"`
	Network           []NetworkInfo         `json:"network,omitempty"`
	Tags              []Tag                 `json:"tags,omitempty"`
	Log               *ClusterLogSetting    `json:"log,omitempty"`
	ResGroupList      []ResGroup            `json:"resGroupList,omitempty"`
	BdVersion         string                `json:"bdVersion,omitempty"`
}

// ClusterModuleDetail describes the runtime module information returned by the detail API.
type ClusterModuleDetail struct {
	Type              string `json:"type"`
	Version           string `json:"version"`
	SlotType          string `json:"slotType"`
	SlotDescription   string `json:"slotDescription,omitempty"`
	ActualInstanceNum int    `json:"actualInstanceNum"`
}

// ClusterInstance describes a single instance returned by the detail API.
type ClusterInstance struct {
	InstanceId    string `json:"instanceId"`
	Status        string `json:"status"`
	ModuleType    string `json:"moduleType"`
	ModuleVersion string `json:"moduleVersion"`
	HostIp        string `json:"hostIp"`
}

// ClusterHealthDetail is the full cluster health information returned by the detail API.
type ClusterHealthDetail struct {
	Status              string               `json:"status"`
	TotalDiskSize       string               `json:"totalDiskSize,omitempty"`
	UsedDiskSize        string               `json:"usedDiskSize,omitempty"`
	NodeDiskMaxSizeInGB *NodeDiskMaxSizeInGB `json:"nodeDiskMaxSizeInGB,omitempty"`
	IndexDiskSize       string               `json:"indexDiskSize,omitempty"`
}

// NodeDiskMaxSizeInGB describes the per node type disk size limits.
type NodeDiskMaxSizeInGB struct {
	EsNode         json.Number `json:"es_node,omitempty"`
	EsColdTierNode json.Number `json:"es_cold_tier_node,omitempty"`
}

// ClusterLogSetting describes which cluster logs are enabled.
type ClusterLogSetting struct {
	MainLog       bool `json:"mainLog"`
	GcLog         bool `json:"gcLog"`
	IndexSlowLog  bool `json:"indexSlowLog"`
	SearchSlowLog bool `json:"searchSlowLog"`
	AuditLog      bool `json:"auditLog"`
}

// DeleteClusterRequest is the request of deleting a BES cluster.
type DeleteClusterRequest struct {
	ClusterId    string `json:"clusterId"`
	RefundReason string `json:"refundReason,omitempty"`
}

// DeleteClusterResponse is the response of deleting a BES cluster.
type DeleteClusterResponse struct {
	Success bool         `json:"success"`
	Status  int          `json:"status"`
	Result  *OrderResult `json:"result"`
}

// OrderResult is the common {orderId} result shared by resize/add-module/delete cluster APIs.
type OrderResult struct {
	OrderId string `json:"orderId"`
}

// StopClusterRequest is the request of stopping a BES cluster.
type StopClusterRequest struct {
	ClusterId string `json:"clusterId"`
}

// StopClusterResponse is the response of stopping a BES cluster.
type StopClusterResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// RestartClusterRequest is the request of restarting a BES cluster.
type RestartClusterRequest struct {
	ClusterId string `json:"clusterId"`
	Mode      string `json:"mode"`
}

// RestartClusterResponse is the response of restarting a BES cluster.
type RestartClusterResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// ResizeModuleInfo describes a partial node update used by cluster resize.
type ResizeModuleInfo struct {
	Type              string              `json:"type,omitempty"`
	Version           string              `json:"version,omitempty"`
	DesireInstanceNum int                 `json:"desireInstanceNum,omitempty"`
	SlotType          string              `json:"slotType,omitempty"`
	DiskSlotInfo      *ResizeDiskSlotInfo `json:"diskSlotInfo,omitempty"`
}

// ResizeDiskSlotInfo describes the disk change used by cluster resize.
type ResizeDiskSlotInfo struct {
	Type       string  `json:"type,omitempty"`
	Size       float64 `json:"size,omitempty"`
	CdsExtraIo float64 `json:"cdsExtraIo,omitempty"`
}

// ResizeClusterRequest is the request of resizing a BES cluster.
type ResizeClusterRequest struct {
	ClusterId   string             `json:"clusterId"`
	Modules     []ResizeModuleInfo `json:"modules"`
	PaymentType string             `json:"paymentType"`
	Coupon      string             `json:"coupon,omitempty"`
	ResizeMode  string             `json:"resizeMode"`
	IsShrink    bool               `json:"isShrink,omitempty"`
}

// ResizeClusterResponse is the response of resizing a BES cluster.
type ResizeClusterResponse struct {
	Success bool         `json:"success"`
	Status  int          `json:"status"`
	Result  *OrderResult `json:"result"`
}

// AddModuleInfo describes a new node type added to a cluster.
type AddModuleInfo struct {
	Type              string                 `json:"type"`
	Version           string                 `json:"version"`
	DesireInstanceNum int                    `json:"desireInstanceNum"`
	SlotType          string                 `json:"slotType"`
	DiskSlotInfo      *AddModuleDiskSlotInfo `json:"diskSlotInfo"`
}

// AddModuleDiskSlotInfo describes the disk configuration used by add-module.
type AddModuleDiskSlotInfo struct {
	Type string  `json:"type"`
	Size float64 `json:"size"`
}

// AddClusterModuleRequest is the request of adding a new node type to a BES cluster.
type AddClusterModuleRequest struct {
	ClusterId   string          `json:"clusterId"`
	Modules     []AddModuleInfo `json:"modules"`
	ResizeMode  string          `json:"resizeMode"`
	PaymentType string          `json:"paymentType"`
}

// AddClusterModuleResponse is the response of adding a new node type to a BES cluster.
type AddClusterModuleResponse struct {
	Success bool         `json:"success"`
	Status  int          `json:"status"`
	Result  *OrderResult `json:"result"`
}

// ResetClusterPasswordRequest is the request of resetting a BES cluster password.
type ResetClusterPasswordRequest struct {
	ClusterId       string `json:"clusterId"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

// ResetClusterPasswordResponse is the response of resetting a BES cluster password.
type ResetClusterPasswordResponse struct {
	Success SuccessFlag `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// ToggleClusterHTTPSRequest is the request of enabling/disabling HTTPS for a BES cluster.
type ToggleClusterHTTPSRequest struct {
	ClusterId string `json:"clusterId"`
	EnableSSL bool   `json:"enableSSL"`
}

// ToggleClusterHTTPSResponse is the response of enabling/disabling HTTPS for a BES cluster.
type ToggleClusterHTTPSResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// BindClusterEIPRequest is the request of binding an EIP to a cluster module.
type BindClusterEIPRequest struct {
	InstanceId   string `json:"instanceId"`
	InstanceType string `json:"instanceType"`
	Eip          string `json:"eip"`
}

// BindClusterEIPResponse is the response of binding an EIP to a cluster module.
type BindClusterEIPResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// UnbindClusterEIPRequest is the request of unbinding an EIP from a cluster module.
type UnbindClusterEIPRequest struct {
	DeployId           string `json:"deployId"`
	ModuleTemplateName string `json:"moduleTemplateName"`
}

// UnbindClusterEIPResponse is the response of unbinding an EIP from a cluster module.
type UnbindClusterEIPResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// ToggleClusterMonitorRequest is the request of enabling/disabling Grafana monitor for a cluster.
type ToggleClusterMonitorRequest struct {
	ClusterId     string `json:"clusterId"`
	EnableMonitor *bool  `json:"enableMonitor"`
}

// ToggleClusterMonitorResponse is the response of enabling/disabling Grafana monitor for a cluster.
type ToggleClusterMonitorResponse struct {
	Success SuccessFlag `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// GetClusterTasksRequest is the request of getting a BES cluster's operation history.
type GetClusterTasksRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetClusterTasksResponse is the response of getting a BES cluster's operation history.
type GetClusterTasksResponse struct {
	Success SuccessFlag         `json:"success"`
	Status  int                 `json:"status"`
	Result  *ClusterTasksResult `json:"result"`
}

// ClusterTasksResult carries the operation history of a cluster.
type ClusterTasksResult struct {
	ClusterId json.Number `json:"clusterId"`
	AppTasks  []AppTask   `json:"appTasks"`
}

// AppTask is a top-level cluster operation task.
type AppTask struct {
	SubTasks       []SubTask              `json:"subTasks"`
	Type           string                 `json:"type"`
	Detail         map[string]interface{} `json:"detail,omitempty"`
	Progress       int                    `json:"progress"`
	Status         string                 `json:"status"`
	StartTimestamp int64                  `json:"start_timestamp"`
	EndTimestamp   int64                  `json:"end_timestamp"`
}

// SubTask is a single step under an AppTask.
type SubTask struct {
	Type           string                 `json:"type"`
	Detail         map[string]interface{} `json:"detail,omitempty"`
	Progress       float64                `json:"progress"`
	RealProgress   float64                `json:"realProgress"`
	Status         string                 `json:"status"`
	TaskList       []TaskCheckItem        `json:"taskList,omitempty"`
	StartTimestamp int64                  `json:"start_timestamp"`
	EndTimestamp   int64                  `json:"end_timestamp"`
}

// TaskCheckItem is a single check performed within a SubTask.
type TaskCheckItem struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Level         string   `json:"level"`
	Status        string   `json:"status"`
	FinishTime    string   `json:"finishTime"`
	Message       string   `json:"message"`
	FailedIndices []string `json:"failedIndices,omitempty"`
}

// GetClusterDataSizeTendencyRequest is the request of getting a cluster's data size tendency.
type GetClusterDataSizeTendencyRequest struct {
	ClusterId   string `json:"clusterId"`
	IndexPrefix string `json:"indexPrefix,omitempty"`
	DatePattern string `json:"datePattern"`
	Times       int    `json:"times"`
	TimeUnit    string `json:"timeUnit"`
}

// GetClusterDataSizeTendencyResponse is the response of getting a cluster's data size tendency.
type GetClusterDataSizeTendencyResponse struct {
	Success SuccessFlag             `json:"success"`
	Status  int                     `json:"status"`
	Result  *DataSizeTendencyResult `json:"result"`
}

// DataSizeTendencyResult carries the data size tendency series of a cluster.
type DataSizeTendencyResult struct {
	ClusterId json.Number    `json:"clusterId"`
	Data      []DataSizeItem `json:"data"`
}

// DataSizeItem is a single data point in the data size tendency series.
type DataSizeItem struct {
	ByteSize  string          `json:"byteSize"`
	Indices   []DataSizeIndex `json:"indices"`
	Size      string          `json:"size"`
	Time      string          `json:"time"`
	TimeMills string          `json:"timeMills"`
}

// DataSizeIndex is a single index's data size within a DataSizeItem.
type DataSizeIndex struct {
	ByteSize string `json:"byteSize"`
	Index    string `json:"index"`
	Size     string `json:"size"`
}

// ListAvailableCouponsResponse is the response of listing the user's available coupons.
type ListAvailableCouponsResponse struct {
	Success bool           `json:"success"`
	Status  int            `json:"status"`
	Result  *CouponsResult `json:"result"`
}

// CouponsResult carries the list of available coupons.
type CouponsResult struct {
	Coupons []Coupon `json:"coupons"`
}

// Coupon describes a single available coupon.
type Coupon struct {
	Id                         float64                `json:"id"`
	Name                       string                 `json:"name"`
	CouponType                 string                 `json:"couponType"`
	ProductType                string                 `json:"productType"`
	ProductTypes               []string               `json:"productTypes,omitempty"`
	ProductRuleDescription     string                 `json:"productRuleDescription,omitempty"`
	TotalAmount                float64                `json:"totalAmount"`
	Balance                    string                 `json:"balance"`
	AmountOrDiscount           string                 `json:"amountOrDiscount"`
	UsedAmount                 string                 `json:"usedAmount"`
	BeginTime                  float64                `json:"beginTime"`
	EndTime                    float64                `json:"endTime"`
	CouponStatus               string                 `json:"couponStatus"`
	Status                     string                 `json:"status"`
	Region                     string                 `json:"region"`
	ConditionArgsMap           map[string]interface{} `json:"conditionArgsMap,omitempty"`
	EffectArgsMap              map[string]interface{} `json:"effectArgsMap,omitempty"`
	ConditionEffectDescription string                 `json:"conditionEffectDescription,omitempty"`
}

// AssessClusterSourceRequest is the request of running an intelligent capacity assessment.
type AssessClusterSourceRequest struct {
	UseType            string `json:"useType"`
	OriginDataSize     int    `json:"originDataSize"`
	OriginDataSizeUnit string `json:"originDataSizeUnit"`
	DataAdd            int    `json:"dataAdd"`
	DataAddUnit        string `json:"dataAddUnit"`
	StorageDays        int    `json:"storageDays"`
	WriteThroughput    int    `json:"writeThroughput"`
	ReadThroughput     int    `json:"readThroughput,omitempty"`
	Replica            string `json:"replica"`
	NeedBos            bool   `json:"needBos,omitempty"`
	BosStorageDays     int    `json:"bosStorageDays,omitempty"`
	NeedVector         bool   `json:"needVector,omitempty"`
	VectorDims         int    `json:"vectorDims,omitempty"`
	VectorType         string `json:"vectorType,omitempty"`
}

// AssessClusterSourceResponse is the response of running an intelligent capacity assessment.
type AssessClusterSourceResponse struct {
	Success SuccessFlag   `json:"success"`
	Status  int           `json:"status"`
	Result  *AssessResult `json:"result"`
}

// AssessResult carries the capacity assessment result.
type AssessResult struct {
	ZoneList          []AssessZone          `json:"zoneList"`
	PackageStatusList []AssessPackageStatus `json:"packageStatusList"`
	OtherConfig       *AssessOtherConfig    `json:"otherConfig,omitempty"`
}

// AssessZone describes the sell-out status of an available zone.
type AssessZone struct {
	Name    string      `json:"name"`
	SellOut SuccessFlag `json:"sellOut"`
}

// AssessPackageStatus describes a recommended node package.
type AssessPackageStatus struct {
	ModuleType     string      `json:"moduleType"`
	PackageVersion string      `json:"packageVersion"`
	DiskType       string      `json:"diskType"`
	DiskSize       json.Number `json:"diskSize"`
	ModuleNum      json.Number `json:"moduleNum"`
}

// AssessOtherConfig carries extra assessment configuration.
type AssessOtherConfig struct {
	BosDataSize json.Number `json:"bosDataSize"`
}
