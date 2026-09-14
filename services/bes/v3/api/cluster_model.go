package api

import "io"

// Request carries request-scoped metadata for BES v3 operations.
// It is embedded in concrete request types and excluded from JSON payloads.
type Request struct {
	Region string `json:"-"`
}

// NodeSpec defines a cluster node specification used by create and update payloads.
type NodeSpec struct {
	Type      string `json:"type,omitempty"`
	NodeType  string `json:"nodeType,omitempty"`
	NodeCount *int   `json:"nodeCount,omitempty"`
	DiskType  string `json:"diskType,omitempty"`
	DiskSize  *int   `json:"diskSize,omitempty"`
	DiskCount *int   `json:"diskCount,omitempty"`
}

// Tag defines a cluster tag.
type Tag struct {
	TagKey   string `json:"tagKey,omitempty"`
	TagValue string `json:"tagValue,omitempty"`
}

// CreateClusterRequest contains the request body for POST /v3/clusters.
type CreateClusterRequest struct {
	Request
	Payment            string                 `json:"payment"`
	PaymentInfo        map[string]interface{} `json:"paymentInfo,omitempty"`
	CouponIds          []string               `json:"couponIds,omitempty"`
	Name               string                 `json:"name"`
	Engine             string                 `json:"engine"`
	EngineType         string                 `json:"engineType"`
	Version            string                 `json:"version"`
	AdminPassword      string                 `json:"adminPassword"`
	EnableHttps        *bool                  `json:"enableHttps"`
	DeletionProtection *bool                  `json:"deletionProtection,omitempty"`
	EnableDeploySet    *bool                  `json:"enableDeploySet,omitempty"`
	LogicalZones       []string               `json:"logicalZones"`
	VpcId              string                 `json:"vpcId"`
	SubnetId           string                 `json:"subnetId"`
	NodeSpecs          []NodeSpec             `json:"nodeSpecs"`
	Tags               []Tag                  `json:"tags,omitempty"`
}

// CreateClusterResponse contains the response body for POST /v3/clusters.
type CreateClusterResponse struct {
	ClusterId string `json:"clusterId,omitempty"`
	OrderId   string `json:"orderId,omitempty"`
	ActionId  string `json:"actionId,omitempty"`
}

// ListClustersRequest contains the optional filters for GET /v3/clusters.
type ListClustersRequest struct {
	Request
	PageNo        int    `json:"-"`
	PageSize      int    `json:"-"`
	Order         string `json:"-"`
	OrderBy       string `json:"-"`
	Payment       string `json:"-"`
	Engine        string `json:"-"`
	Status        string `json:"-"`
	ClusterHealth string `json:"-"`
	LogicalZone   string `json:"-"`
	ClusterName   string `json:"-"`
	ClusterId     string `json:"-"`
	TagKey        string `json:"-"`
	TagValue      string `json:"-"`
}

// ClusterHealth models the cluster health summary.
type ClusterHealth struct {
	Status        string `json:"status,omitempty"`
	Message       string `json:"message,omitempty"`
	TotalDiskSize string `json:"totalDiskSize,omitempty"`
	UsedDiskSize  string `json:"usedDiskSize,omitempty"`
}

// ClusterDetailNodeSpec defines node spec details in cluster detail responses.
type ClusterDetailNodeSpec struct {
	Type                string `json:"type,omitempty"`
	NodeType            string `json:"nodeType,omitempty"`
	NodeTypeDisplayName string `json:"nodeTypeDisplayName,omitempty"`
	CPU                 *int   `json:"cpu,omitempty"`
	Memory              *int   `json:"memory,omitempty"`
	NodeCount           *int   `json:"nodeCount,omitempty"`
	DiskType            string `json:"diskType,omitempty"`
	DiskTypeDisplayName string `json:"diskTypeDisplayName,omitempty"`
	DiskSize            *int   `json:"diskSize,omitempty"`
	DiskCount           *int   `json:"diskCount,omitempty"`
	CdsExtraIo          *int   `json:"cdsExtraIo,omitempty"`
}

// NodeListNodeSpec is the node spec structure returned by the node list API.
type NodeListNodeSpec = ClusterDetailNodeSpec

// ClusterSummary describes a BES cluster as returned by the cluster list API.
type ClusterSummary struct {
	ClusterId          string        `json:"clusterId,omitempty"`
	Name               string        `json:"name,omitempty"`
	CreateTime         string        `json:"createTime,omitempty"`
	Engine             string        `json:"engine,omitempty"`
	EngineType         string        `json:"engineType,omitempty"`
	Version            string        `json:"version,omitempty"`
	KernelVersion      string        `json:"kernelVersion,omitempty"`
	Status             string        `json:"status,omitempty"`
	ClusterHealth      ClusterHealth `json:"clusterHealth,omitempty"`
	Region             string        `json:"region,omitempty"`
	LogicalZonesNumber *int          `json:"logicalZonesNumber,omitempty"`
	LogicalZones       []string      `json:"logicalZones,omitempty"`
	RunningTime        *RunningTime  `json:"runningTime,omitempty"`
	DeletionProtection *bool         `json:"deletionProtection,omitempty"`
	Payment            string        `json:"payment,omitempty"`
	IsPrepayToPostpay  *bool         `json:"isPrepayToPostpay,omitempty"`
	ExpirationTime     string        `json:"expirationTime,omitempty"`
	Tags               []Tag         `json:"tags,omitempty"`
	InRecycleBin       *bool         `json:"inRecycleBin,omitempty"`
	DeleteExpireTime   string        `json:"deleteExpireTime,omitempty"`
	RecycleCount       *int          `json:"recycleCount,omitempty"`
}

// RunningTime models the elapsed running duration of a cluster.
type RunningTime struct {
	Day           *int   `json:"day,omitempty"`
	Hour          *int   `json:"hour,omitempty"`
	Minute        *int   `json:"minute,omitempty"`
	Second        *int   `json:"second,omitempty"`
	RunningTimeMs *int64 `json:"runningTimeMs,omitempty"`
}

// ClusterBilling models the billing information of a cluster.
type ClusterBilling struct {
	ProductType  string `json:"productType,omitempty"`
	TimeLength   *int   `json:"timeLength,omitempty"`
	TimeUnit     string `json:"timeUnit,omitempty"`
	InstanceId   string `json:"instanceId,omitempty"`
	ShortId      string `json:"shortId,omitempty"`
	SubAccountId string `json:"subAccountId,omitempty"`
}

// ClusterDetail describes a BES cluster as returned by the cluster detail API.
type ClusterDetail struct {
	ClusterId                string                  `json:"clusterId,omitempty"`
	Name                     string                  `json:"name,omitempty"`
	CreateTime               string                  `json:"createTime,omitempty"`
	Engine                   string                  `json:"engine,omitempty"`
	EngineType               string                  `json:"engineType,omitempty"`
	Version                  string                  `json:"version,omitempty"`
	KernelVersion            string                  `json:"kernelVersion,omitempty"`
	Status                   string                  `json:"status,omitempty"`
	ClusterHealth            ClusterHealth           `json:"clusterHealth,omitempty"`
	Region                   string                  `json:"region,omitempty"`
	LogicalZonesCount        *int                    `json:"logicalZonesCount,omitempty"`
	LogicalZones             []string                `json:"logicalZones,omitempty"`
	RunningTime              *RunningTime            `json:"runningTime,omitempty"`
	DeletionProtection       *bool                   `json:"deletionProtection,omitempty"`
	MaintenanceTimeZone      string                  `json:"maintenanceTimeZone,omitempty"`
	MaintenancePeriods       []string                `json:"maintenancePeriods,omitempty"`
	MaintenanceStartTime     string                  `json:"maintenanceStartTime,omitempty"`
	MaintenanceEndTime       string                  `json:"maintenanceEndTime,omitempty"`
	Payment                  string                  `json:"payment,omitempty"`
	IsPrepayToPostpay        *bool                   `json:"isPrepayToPostpay,omitempty"`
	ExpirationTime           string                  `json:"expirationTime,omitempty"`
	IsAutoRenew              *bool                   `json:"isAutoRenew,omitempty"`
	AutoRenew                *bool                   `json:"autoRenew,omitempty"`
	Billing                  *ClusterBilling         `json:"billing,omitempty"`
	AdminUserName            string                  `json:"adminUserName,omitempty"`
	EnablePublicAccess       *bool                   `json:"enablePublicAccess,omitempty"`
	PrivateEndpoint          string                  `json:"privateEndpoint,omitempty"`
	PublicEndpoint           string                  `json:"publicEndpoint,omitempty"`
	EnableHttps              *bool                   `json:"enableHttps,omitempty"`
	EnableVisualPublicAccess *bool                   `json:"enableVisualPublicAccess,omitempty"`
	VisualPrivateEndpoint    string                  `json:"visualPrivateEndpoint,omitempty"`
	VisualPublicEndpoint     string                  `json:"visualPublicEndpoint,omitempty"`
	EnableVisualHttps        *bool                   `json:"enableVisualHttps,omitempty"`
	PublicIpWhitelist        []string                `json:"publicIpWhitelist,omitempty"`
	PrivateIpWhitelist       []string                `json:"privateIpWhitelist,omitempty"`
	VisualPublicIpWhitelist  []string                `json:"visualPublicIpWhitelist,omitempty"`
	VisualPrivateIpWhitelist []string                `json:"visualPrivateIpWhitelist,omitempty"`
	EnableCerebro            *bool                   `json:"enableCerebro,omitempty"`
	CerebroPublicEndpoint    string                  `json:"cerebroPublicEndpoint,omitempty"`
	CerebroPrivateEndpoint   string                  `json:"cerebroPrivateEndpoint,omitempty"`
	VpcId                    string                  `json:"vpcId,omitempty"`
	SubnetId                 string                  `json:"subnetId,omitempty"`
	NodeSpecs                []ClusterDetailNodeSpec `json:"nodeSpecs,omitempty"`
	Tags                     []Tag                   `json:"tags,omitempty"`
	InRecycleBin             *bool                   `json:"inRecycleBin,omitempty"`
	DeleteExpireTime         string                  `json:"deleteExpireTime,omitempty"`
	RecycleCount             *int                    `json:"recycleCount,omitempty"`
}

// ListClustersResponse contains the response body for GET /v3/clusters.
type ListClustersResponse struct {
	Clusters   []ClusterSummary `json:"clusters,omitempty"`
	PageNo     int              `json:"pageNo,omitempty"`
	PageSize   int              `json:"pageSize,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
}

// GetClusterRequest contains the path variables for GET /v3/clusters/{clusterId}.
type GetClusterRequest struct {
	Request
	ClusterId string `json:"-"`
}

// GetClusterResponse is returned by the cluster detail API.
type GetClusterResponse = ClusterDetail

// ClusterNode describes a single node returned by the cluster node list API.
type ClusterNode struct {
	NodeId            string   `json:"nodeId,omitempty"`
	NodeName          string   `json:"nodeName,omitempty"`
	Type              string   `json:"type,omitempty"`
	Status            string   `json:"status,omitempty"`
	NodeType          string   `json:"nodeType,omitempty"`
	DiskType          string   `json:"diskType,omitempty"`
	LogicalZone       string   `json:"logicalZone,omitempty"`
	CPU               *int     `json:"cpu,omitempty"`
	CPUPercent        *float64 `json:"cpuPercent,omitempty"`
	Memory            *int     `json:"memory,omitempty"`
	MemoryPercent     *float64 `json:"memoryPercent,omitempty"`
	HeapMemory        *int     `json:"heapMemory,omitempty"`
	HeapMemoryPercent *float64 `json:"heapMemoryPercent,omitempty"`
	DiskCount         *int     `json:"diskCount,omitempty"`
	DiskSize          *int     `json:"diskSize,omitempty"`
	DiskUsedPercent   *float64 `json:"diskUsedPercent,omitempty"`
	CdsExtraIo        *int     `json:"cdsExtraIo,omitempty"`
}

// ListClusterNodesRequest contains the path variables for GET /v3/clusters/{clusterId}/nodes.
type ListClusterNodesRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListClusterNodesResponse is returned by the cluster node list API.
type ListClusterNodesResponse struct {
	Nodes     []ClusterNode      `json:"nodes,omitempty"`
	NodeSpecs []NodeListNodeSpec `json:"nodeSpecs,omitempty"`
}

// NodeAvailableSpec models the node spec availability returned by the spec API.
type NodeAvailableSpec struct {
	Type            string         `json:"type,omitempty"`
	NodeMinSize     *int           `json:"nodeMinSize,omitempty"`
	NodeMaxSize     *int           `json:"nodeMaxSize,omitempty"`
	NodeTypes       []string       `json:"nodeTypes,omitempty"`
	DefaultNodeType string         `json:"defaultNodeType,omitempty"`
	DiskTypes       []string       `json:"diskTypes,omitempty"`
	DefaultDiskType string         `json:"defaultDiskType,omitempty"`
	DiskMinSize     map[string]int `json:"diskMinSize,omitempty"`
	DiskMaxSize     map[string]int `json:"diskMaxSize,omitempty"`
}

// QueryNodeSpec models a node spec entry returned by the spec API.
type QueryNodeSpec struct {
	NodeType    string `json:"nodeType,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	CPU         *int   `json:"cpu,omitempty"`
	Memory      *int   `json:"memory,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

// DiskSpecInfo models the disk spec availability returned by the spec API.
type DiskSpecInfo struct {
	DiskType    string `json:"diskType,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
}

// QueryAvailableSpecsRequest contains the filters for POST /v3/clusters/specs.
type QueryAvailableSpecsRequest struct {
	Request
	Engine       string   `json:"engine"`
	EngineType   string   `json:"engineType"`
	Version      string   `json:"version"`
	Features     []string `json:"features,omitempty"`
	LogicalZones []string `json:"logicalZones"`
}

// QueryAvailableSpecsResponse contains the response body for POST /v3/clusters/specs.
type QueryAvailableSpecsResponse struct {
	NodeAvailableSpecs []NodeAvailableSpec `json:"nodeAvailableSpecs,omitempty"`
	NodeSpecs          []QueryNodeSpec     `json:"nodeSpecs,omitempty"`
	DiskSpecs          []DiskSpecInfo      `json:"diskSpecs,omitempty"`
}

// KernelVersion models a kernel version entry.
type KernelVersion struct {
	Engine     string   `json:"engine,omitempty"`
	EngineType string   `json:"engineType,omitempty"`
	Version    string   `json:"version,omitempty"`
	Status     string   `json:"status,omitempty"`
	Features   []string `json:"features,omitempty"`
	Scenes     []string `json:"scenes,omitempty"`
}

// KernelUpgrade models a kernel upgrade entry.
type KernelUpgrade struct {
	Kernel       string   `json:"kernel,omitempty"`
	UpgradeModes []string `json:"upgradeModes,omitempty"`
}

// QueryAvailableKernelsRequest contains the filters for GET /v3/clusters/kernels.
type QueryAvailableKernelsRequest struct {
	Request
}

// QueryAvailableKernelsResponse contains the response body for GET /v3/clusters/kernels.
type QueryAvailableKernelsResponse struct {
	Versions []KernelVersion `json:"versions,omitempty"`
}

// DownloadClusterCertRequest contains the path variables for GET /v3/clusters/{clusterId}/certs.
type DownloadClusterCertRequest struct {
	Request
	ClusterId string `json:"-"`
}

// DownloadClusterCertResponse contains the certificate zip stream and its metadata.
// Body is the raw response stream; the caller is responsible for reading and closing it.
type DownloadClusterCertResponse struct {
	Body                io.ReadCloser
	ContentType         string
	ContentDisposition  string
	ContentLength       int64
	CacheControl        string
	Pragma              string
	Expires             string
	XContentTypeOptions string
}

// UpdateClusterNameRequest contains the body for PUT /v3/clusters/{clusterId}/name.
type UpdateClusterNameRequest struct {
	Request
	ClusterId string `json:"-"`
	NewName   string `json:"newName,omitempty"`
}

// UpdateClusterMaintenanceRequest contains the body for PUT /v3/clusters/{clusterId}/maintenance-duration.
type UpdateClusterMaintenanceRequest struct {
	Request
	ClusterId            string   `json:"-"`
	MaintenanceTimeZone  string   `json:"maintenanceTimeZone,omitempty"`
	MaintenancePeriods   []string `json:"maintenancePeriods,omitempty"`
	MaintenanceStartTime string   `json:"maintenanceStartTime,omitempty"`
	MaintenanceEndTime   string   `json:"maintenanceEndTime,omitempty"`
}

// UpdateClusterDeletionProtectionRequest contains the body for PUT /v3/clusters/{clusterId}/deletion-protection.
type UpdateClusterDeletionProtectionRequest struct {
	Request
	ClusterId          string `json:"-"`
	DeletionProtection *bool  `json:"deletionProtection,omitempty"`
}

// UpdateClusterDeletionProtectionResponse is returned by the deletion protection update API.
type UpdateClusterDeletionProtectionResponse struct {
	ClusterId          string `json:"clusterId,omitempty"`
	DeletionProtection *bool  `json:"deletionProtection,omitempty"`
}

// UpdateClusterNameResponse is returned by the name update API.
type UpdateClusterNameResponse struct {
	ClusterId string `json:"clusterId,omitempty"`
	Name      string `json:"name,omitempty"`
}

// UpdateClusterMaintenanceResponse is returned by the maintenance update API.
type UpdateClusterMaintenanceResponse struct {
	Success *bool `json:"success,omitempty"`
}

// ClusterActionResponse is returned by action-style APIs.
type ClusterActionResponse struct {
	ClusterId string `json:"clusterId,omitempty"`
	ActionId  string `json:"actionId,omitempty"`
}

// DeleteClusterRequest contains the filters for DELETE /v3/clusters/{clusterId}.
type DeleteClusterRequest struct {
	Request
	ClusterId string `json:"-"`
	Recycle   *bool  `json:"-"`
}

// DeleteClusterResponse is returned by the cluster deletion API.
type DeleteClusterResponse struct {
	ClusterId string `json:"clusterId,omitempty"`
}

// RecoverClusterRequest contains the path variables for PUT /v3/clusters/{clusterId}/recovery.
type RecoverClusterRequest struct {
	Request
	ClusterId string `json:"-"`
}

// RecoverClusterResponse is returned by the cluster recovery API.
type RecoverClusterResponse = ClusterActionResponse

// ResizeClusterRequest contains the body for PUT /v3/clusters/{clusterId} (partial resize).
type ResizeClusterRequest struct {
	Request
	ClusterId       string     `json:"-"`
	Mode            string     `json:"mode"`
	LogicalZones    []string   `json:"logicalZones,omitempty"`
	VpcId           string     `json:"vpcId,omitempty"`
	SubnetId        string     `json:"subnetId,omitempty"`
	EnableDeploySet *bool      `json:"enableDeploySet,omitempty"`
	NodeSpecs       []NodeSpec `json:"nodeSpecs,omitempty"`
	DeleteNodeTypes []string   `json:"deleteNodeTypes,omitempty"`
}

// ResizeClusterResponse is returned by the cluster resize API.
type ResizeClusterResponse struct {
	ClusterId string `json:"clusterId,omitempty"`
	OrderId   string `json:"orderId,omitempty"`
}

// MigrationNode describes a node returned by the migratable/suggested node APIs.
type MigrationNode struct {
	NodeId      string `json:"nodeId,omitempty"`
	NodeName    string `json:"nodeName,omitempty"`
	Status      string `json:"status,omitempty"`
	LogicalZone string `json:"logicalZone,omitempty"`
}

// ListSuggestedMigrationNodesRequest contains the filters for
// GET /v3/clusters/{clusterId}/migrations/suggested-nodes.
type ListSuggestedMigrationNodesRequest struct {
	Request
	ClusterId    string `json:"-"`
	Type         string `json:"-"`
	MigrateCount int    `json:"-"`
}

// ListMigratableNodesRequest contains the filters for GET /v3/clusters/{clusterId}/migrations/nodes.
type ListMigratableNodesRequest struct {
	Request
	ClusterId string `json:"-"`
	Type      string `json:"-"`
}

// ListMigrationNodesResponse is returned by the migratable/suggested node APIs.
type ListMigrationNodesResponse struct {
	Nodes []MigrationNode `json:"nodes,omitempty"`
}

// VersionUpgrade models a version upgrade entry.
type VersionUpgrade struct {
	Version      string   `json:"version,omitempty"`
	UpgradeModes []string `json:"upgradeModes,omitempty"`
}

// QueryAvailableUpgradesRequest contains the path variables for
// GET /v3/clusters/{clusterId}/upgrades/versions.
type QueryAvailableUpgradesRequest struct {
	Request
	ClusterId string `json:"-"`
}

// QueryAvailableUpgradesResponse is returned by GET /v3/clusters/{clusterId}/upgrades/versions.
type QueryAvailableUpgradesResponse struct {
	ClusterId       string           `json:"clusterId,omitempty"`
	Name            string           `json:"name,omitempty"`
	Engine          string           `json:"engine,omitempty"`
	CurrentVersion  string           `json:"currentVersion,omitempty"`
	CurrentKernel   string           `json:"currentKernel,omitempty"`
	EnableDeploySet *bool            `json:"enableDeploySet,omitempty"`
	VersionUpgrades []VersionUpgrade `json:"versionUpgrades,omitempty"`
	KernelUpgrades  []KernelUpgrade  `json:"kernelUpgrades,omitempty"`
}

// UpdateClusterPublicAccessRequest contains the body for PUT /v3/clusters/{clusterId}/public-access.
type UpdateClusterPublicAccessRequest struct {
	Request
	ClusterId  string `json:"-"`
	Enabled    *bool  `json:"enabled"`
	PublicIp   string `json:"publicIp,omitempty"`
	AccessType string `json:"accessType,omitempty"`
}

// UpdateClusterCerebroRequest contains the body for PUT /v3/clusters/{clusterId}/cerebro.
// The API only supports enabling Cerebro; Enabled must be true.
type UpdateClusterCerebroRequest struct {
	Request
	ClusterId string `json:"-"`
	Enabled   *bool  `json:"enabled"`
}

// UpdateClusterCerebroResponse is returned by the Cerebro enable API.
type UpdateClusterCerebroResponse struct {
	ClusterId string `json:"clusterId,omitempty"`
	Enabled   *bool  `json:"enabled,omitempty"`
}

// StartClusterRequest contains the path variables for PUT /v3/clusters/{clusterId}/start.
type StartClusterRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ClusterNodesRequest contains the body for node-scoped mutation APIs
// (PUT /v3/clusters/{clusterId}/nodes/start, PUT /v3/clusters/{clusterId}/nodes/stop).
type ClusterNodesRequest struct {
	Request
	ClusterId string   `json:"-"`
	Nodes     []string `json:"nodes"`
}

// MigrateClusterNodeDataRequest contains the body for POST /v3/clusters/{clusterId}/migrations.
type MigrateClusterNodeDataRequest struct {
	Request
	ClusterId string   `json:"-"`
	Nodes     []string `json:"nodes"`
	Type      string   `json:"type"`
}

// UpdateClusterAccessWhitelistRequest contains the body for
// POST /v3/clusters/{clusterId}/access-white-ips.
type UpdateClusterAccessWhitelistRequest struct {
	Request
	ClusterId   string   `json:"-"`
	AccessType  string   `json:"accessType"`
	NetworkType string   `json:"networkType"`
	IpWhitelist []string `json:"ipWhitelist"`
}

// UpdateClusterAccessWhitelistResponse is returned by the access whitelist update API.
type UpdateClusterAccessWhitelistResponse struct {
	ClusterId                string   `json:"clusterId,omitempty"`
	PublicIpWhitelist        []string `json:"publicIpWhitelist,omitempty"`
	VisualPublicIpWhitelist  []string `json:"visualPublicIpWhitelist,omitempty"`
	PrivateIpWhitelist       []string `json:"privateIpWhitelist,omitempty"`
	VisualPrivateIpWhitelist []string `json:"visualPrivateIpWhitelist,omitempty"`
}

// UpgradeClusterRequest contains the body for POST /v3/clusters/{clusterId}/upgrades.
type UpgradeClusterRequest struct {
	Request
	ClusterId            string `json:"-"`
	UpgradeType          string `json:"upgradeType"`
	TargetVersion        string `json:"targetVersion,omitempty"`
	TargetKernel         string `json:"targetKernel,omitempty"`
	Operation            string `json:"operation"`
	UpgradeMode          string `json:"upgradeMode,omitempty"`
	SkipHealthCheck      *bool  `json:"skipHealthCheck,omitempty"`
	SkipDeprecationCheck *bool  `json:"skipDeprecationCheck,omitempty"`
	UpgradeVisual        *bool  `json:"upgradeVisual,omitempty"`
	EnableDeploySet      *bool  `json:"enableDeploySet,omitempty"`
}

// UpdateClusterProtocolRequest contains the body for PUT /v3/clusters/{clusterId}/protocol.
type UpdateClusterProtocolRequest struct {
	Request
	ClusterId   string `json:"-"`
	EnableHttps *bool  `json:"enableHttps"`
}

// StopClusterRequest contains the path variables for PUT /v3/clusters/{clusterId}/stop.
type StopClusterRequest struct {
	Request
	ClusterId string `json:"-"`
}

// RestartClusterRequest contains the body for PUT /v3/clusters/{clusterId}/restart.
type RestartClusterRequest struct {
	Request
	ClusterId   string   `json:"-"`
	RestartType string   `json:"restartType"`
	Mode        string   `json:"mode"`
	Nodes       []string `json:"nodes,omitempty"`
	BatchCount  *float64 `json:"batchCount,omitempty"`
	BatchUnit   string   `json:"batchUnit,omitempty"`
}
