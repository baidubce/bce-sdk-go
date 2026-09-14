package api

type IndexItem struct {
	Index        string `json:"index,omitempty"`
	IndexSize    string `json:"indexSize,omitempty"`
	IndexPriSize string `json:"indexPriSize,omitempty"`
}

type RestoreIndexInfo struct {
	Index        string `json:"index,omitempty"`
	SourceIndex  string `json:"sourceIndex,omitempty"`
	IndexSize    string `json:"indexSize,omitempty"`
	IndexPriSize string `json:"indexPriSize,omitempty"`
}

type RestoreDetailItem struct {
	RestoreShardsTotal        string             `json:"restoreShardsTotal,omitempty"`
	RestoreShardsUnassigned   string             `json:"restoreShardsUnassigned,omitempty"`
	RestoreShardsInitializing string             `json:"restoreShardsInitializing,omitempty"`
	RestoreShardsStarted      string             `json:"restoreShardsStarted,omitempty"`
	StartTimestamp            string             `json:"startTimestamp,omitempty"`
	EndTimestamp              string             `json:"endTimestamp,omitempty"`
	RestoreIndices            string             `json:"restoreIndices,omitempty"`
	RestoreIndexPattern       string             `json:"restoreIndexPattern,omitempty"`
	RenamePattern             string             `json:"renamePattern,omitempty"`
	RenameReplacement         string             `json:"renameReplacement,omitempty"`
	RestoreIndexInfoList      []RestoreIndexInfo `json:"restoreIndexInfoList,omitempty"`
}

type RestoreItem struct {
	HistoryId         string             `json:"historyId,omitempty"`
	RestoreId         string             `json:"restoreId,omitempty"`
	DestClusterId     string             `json:"destClusterId,omitempty"`
	DestClusterName   string             `json:"destClusterName,omitempty"`
	DestEngineVersion string             `json:"destEngineVersion,omitempty"`
	DestKernelVersion string             `json:"destKernelVersion,omitempty"`
	RestoreStatus     string             `json:"restoreStatus,omitempty"`
	RestoreDetail     *RestoreDetailItem `json:"restoreDetail,omitempty"`
}

type RestoreClusterItem struct {
	ClusterId         string   `json:"clusterId,omitempty"`
	Name              string   `json:"name,omitempty"`
	Engine            string   `json:"engine,omitempty"`
	EngineType        string   `json:"engineType,omitempty"`
	Version           string   `json:"version,omitempty"`
	KernelVersion     string   `json:"kernelVersion,omitempty"`
	Status            string   `json:"status,omitempty"`
	ClusterHealth     string   `json:"clusterHealth,omitempty"`
	LogicalZones      []string `json:"logicalZones,omitempty"`
	Available         bool     `json:"available,omitempty"`
	UnavailableReason string   `json:"unavailableReason,omitempty"`
}

type SnapshotRuleIndexItem struct {
	IndexName string `json:"indexName,omitempty"`
	TotalSize string `json:"totalSize,omitempty"`
	IncreSize string `json:"increSize,omitempty"`
}

type SnapshotDetailItem struct {
	SnapshotName   string `json:"snapshotName,omitempty"`
	StartTimestamp string `json:"startTimestamp,omitempty"`
	EndTimestamp   string `json:"endTimestamp,omitempty"`
	Total          string `json:"total,omitempty"`
	Successful     string `json:"successful,omitempty"`
	Failed         string `json:"failed,omitempty"`
}

type SnapshotItem struct {
	HistoryId             string              `json:"historyId,omitempty"`
	ConfigId              string              `json:"configId,omitempty"`
	Type                  string              `json:"type,omitempty"`
	SnapshotStatus        string              `json:"snapshotStatus,omitempty"`
	RestoreId             string              `json:"restoreId,omitempty"`
	RestoreStatus         string              `json:"restoreStatus,omitempty"`
	RestoreStartTimestamp string              `json:"restoreStartTimestamp,omitempty"`
	InvalidTime           string              `json:"invalidTime,omitempty"`
	InvalidReason         string              `json:"invalidReason,omitempty"`
	CleanupTime           string              `json:"cleanupTime,omitempty"`
	CleanupReason         string              `json:"cleanupReason,omitempty"`
	TotalSize             string              `json:"totalSize,omitempty"`
	IncreSize             string              `json:"increSize,omitempty"`
	Remark                string              `json:"remark,omitempty"`
	SnapshotDetail        *SnapshotDetailItem `json:"snapshotDetail,omitempty"`
	RestoreDetail         *RestoreDetailItem  `json:"restoreDetail,omitempty"`
}

type AutoSnapshotConfigItem struct {
	ConfigId               string `json:"configId,omitempty"`
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	SnapshotType           string `json:"snapshotType,omitempty"`
	Period                 string `json:"period,omitempty"`
	ExpireAfter            string `json:"expireAfter,omitempty"`
	Remark                 string `json:"remark,omitempty"`
	Enabled                bool   `json:"enabled,omitempty"`
	AllIndex               bool   `json:"allIndex,omitempty"`
	SnapshotIndices        string `json:"snapshotIndices,omitempty"`
	SnapshotIndexPattern   string `json:"snapshotIndexPattern,omitempty"`
	CreateTimestamp        string `json:"createTimestamp,omitempty"`
	UpdateTimestamp        string `json:"updateTimestamp,omitempty"`
	SnapshotStatus         string `json:"snapshotStatus,omitempty"`
	SnapshotStartTimestamp string `json:"snapshotStartTimestamp,omitempty"`
	RestoreStatus          string `json:"restoreStatus,omitempty"`
	RestoreStartTimestamp  string `json:"restoreStartTimestamp,omitempty"`
}

// 1. 按模式查询快照索引
type ListIndicesByPatternRequest struct {
	Request
	ClusterId string `json:"-"`
	Pattern   string `json:"pattern,omitempty"`
}

type ListIndicesByPatternResponse struct {
	Indices []IndexItem `json:"indices,omitempty"`
}

// 2. 查询恢复任务列表
type ListRestoresRequest struct {
	Request
	ClusterId     string `json:"-"`
	SnapshotId    string `json:"-"`
	PageNo        int    `json:"-"`
	PageSize      int    `json:"-"`
	Order         string `json:"-"`
	OrderBy       string `json:"-"`
	RestoreStatus string `json:"-"`
}

type ListRestoresResponse struct {
	Restores   []RestoreItem `json:"restores,omitempty"`
	PageNo     int           `json:"pageNo,omitempty"`
	PageSize   int           `json:"pageSize,omitempty"`
	TotalCount int           `json:"totalCount,omitempty"`
	OrderBy    string        `json:"orderBy,omitempty"`
	Order      string        `json:"order,omitempty"`
}

// 3. 查询可恢复集群
type ListRestoreClustersRequest struct {
	Request
	ClusterId   string `json:"-"`
	SnapshotId  string `json:"-"`
	IncludeSelf *bool  `json:"-"`
	Keyword     string `json:"-"`
}

type ListRestoreClustersResponse struct {
	Clusters []RestoreClusterItem `json:"clusters,omitempty"`
}

// 4. 查询快照规则
type GetSnapshotRuleRequest struct {
	Request
	ClusterId           string `json:"-"`
	SnapshotId          string `json:"-"`
	PageNo              int    `json:"-"`
	PageSize            int    `json:"-"`
	OrderBy             string `json:"-"`
	Order               string `json:"-"`
	IndexName           string `json:"-"`
	RestoreIndexPattern string `json:"-"`
}

type GetSnapshotRuleResponse struct {
	SnapshotId      string                  `json:"snapshotId,omitempty"`
	ConfigId        string                  `json:"configId,omitempty"`
	SnapshotName    string                  `json:"snapshotName,omitempty"`
	SnapshotStatus  string                  `json:"snapshotStatus,omitempty"`
	ExpireAfter     string                  `json:"expireAfter,omitempty"`
	ExpireAfterText string                  `json:"expireAfterText,omitempty"`
	IndexDataReady  bool                    `json:"indexDataReady,omitempty"`
	PageNo          int                     `json:"pageNo,omitempty"`
	PageSize        int                     `json:"pageSize,omitempty"`
	TotalCount      int                     `json:"totalCount,omitempty"`
	Indices         []SnapshotRuleIndexItem `json:"indices,omitempty"`
}

// 5. 查询快照列表
type ListSnapshotsRequest struct {
	Request
	ClusterId      string `json:"-"`
	PageNo         int    `json:"-"`
	PageSize       int    `json:"-"`
	Order          string `json:"-"`
	OrderBy        string `json:"-"`
	SnapshotName   string `json:"-"`
	Type           string `json:"-"`
	StartTimestamp int64  `json:"-"`
	EndTimestamp   int64  `json:"-"`
	SnapshotStatus string `json:"-"`
	RestoreStatus  string `json:"-"`
}

type ListSnapshotsResponse struct {
	Snapshots  []SnapshotItem `json:"snapshots,omitempty"`
	PageNo     int            `json:"pageNo,omitempty"`
	PageSize   int            `json:"pageSize,omitempty"`
	TotalCount int            `json:"totalCount,omitempty"`
	OrderBy    string         `json:"orderBy,omitempty"`
	Order      string         `json:"order,omitempty"`
}

// 6. 查询快照配置详情
type GetSnapshotConfigRequest struct {
	Request
	ClusterId string `json:"-"`
	ConfigId  string `json:"-"`
}

type GetSnapshotConfigResponse struct {
	ConfigId               string `json:"configId,omitempty"`
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	SnapshotType           string `json:"snapshotType,omitempty"`
	Period                 string `json:"period,omitempty"`
	ExpireAfter            string `json:"expireAfter,omitempty"`
	Remark                 string `json:"remark,omitempty"`
	Repository             string `json:"repository,omitempty"`
	RepositoryBucket       string `json:"repositoryBucket,omitempty"`
	RepositoryBasePath     string `json:"repositoryBasePath,omitempty"`
	Enabled                bool   `json:"enabled,omitempty"`
	CreateTimestamp        string `json:"createTimestamp,omitempty"`
	UpdateTimestamp        string `json:"updateTimestamp,omitempty"`
	EnableTimestamp        string `json:"enableTimestamp,omitempty"`
	AllIndex               bool   `json:"allIndex,omitempty"`
	SnapshotIndices        string `json:"snapshotIndices,omitempty"`
	SnapshotIndexPattern   string `json:"snapshotIndexPattern,omitempty"`
	SnapshotStatus         string `json:"snapshotStatus,omitempty"`
	SnapshotStartTimestamp string `json:"snapshotStartTimestamp,omitempty"`
	RestoreStatus          string `json:"restoreStatus,omitempty"`
	RestoreStartTimestamp  string `json:"restoreStartTimestamp,omitempty"`
}

// 7. 查询快照详情
type GetSnapshotRequest struct {
	Request
	ClusterId  string `json:"-"`
	SnapshotId string `json:"-"`
}

type GetSnapshotResponse struct {
	HistoryId                string `json:"historyId,omitempty"`
	ConfigId                 string `json:"configId,omitempty"`
	Type                     string `json:"type,omitempty"`
	SnapshotName             string `json:"snapshotName,omitempty"`
	SnapshotStatus           string `json:"snapshotStatus,omitempty"`
	RestoreId                string `json:"restoreId,omitempty"`
	RestoreStatus            string `json:"restoreStatus,omitempty"`
	RestoreStartTimestamp    string `json:"restoreStartTimestamp,omitempty"`
	TotalSize                string `json:"totalSize,omitempty"`
	IncreSize                string `json:"increSize,omitempty"`
	Remark                   string `json:"remark,omitempty"`
	ExpireAfter              string `json:"expireAfter,omitempty"`
	ExpireAfterText          string `json:"expireAfterText,omitempty"`
	StartTimestamp           string `json:"startTimestamp,omitempty"`
	EndTimestamp             string `json:"endTimestamp,omitempty"`
	InvalidTime              string `json:"invalidTime,omitempty"`
	InvalidReason            string `json:"invalidReason,omitempty"`
	CleanupTime              string `json:"cleanupTime,omitempty"`
	CleanupReason            string `json:"cleanupReason,omitempty"`
	EsVersion                string `json:"esVersion,omitempty"`
	SnapshotShardsTotal      string `json:"snapshotShardsTotal,omitempty"`
	SnapshotShardsSuccessful string `json:"snapshotShardsSuccessful,omitempty"`
	SnapshotShardsFailed     string `json:"snapshotShardsFailed,omitempty"`
	SnapshotIndexData        string `json:"snapshotIndexData,omitempty"`
}

// 9. 查询自动快照配置列表
type ListAutoSnapshotConfigsRequest struct {
	Request
	ClusterId string `json:"-"`
}

type ListAutoSnapshotConfigsResponse struct {
	ConfigList []AutoSnapshotConfigItem `json:"configList,omitempty"`
}

// 10. 创建快照恢复
type CreateRestoreRequest struct {
	Request
	ClusterId           string `json:"-"`
	SnapshotId          string `json:"-"`
	DestClusterId       string `json:"destClusterId"`
	AllIndex            *bool  `json:"allIndex,omitempty"`
	RestoreIndices      string `json:"restoreIndices,omitempty"`
	RestoreIndexPattern string `json:"restoreIndexPattern,omitempty"`
	RenamePattern       string `json:"renamePattern,omitempty"`
	RenameReplacement   string `json:"renameReplacement,omitempty"`
}

type CreateRestoreResponse struct {
	RestoreId int64 `json:"restoreId,omitempty"`
}

// 12. 创建手工快照
type CreateManualSnapshotRequest struct {
	Request
	ClusterId            string `json:"-"`
	Name                 string `json:"name"`
	SnapshotType         string `json:"snapshotType,omitempty"`
	ExpireAfter          string `json:"expireAfter,omitempty"`
	Remark               string `json:"remark,omitempty"`
	AllIndex             *bool  `json:"allIndex,omitempty"`
	SnapshotIndices      string `json:"snapshotIndices,omitempty"`
	SnapshotIndexPattern string `json:"snapshotIndexPattern,omitempty"`
}

type CreateManualSnapshotResponse struct {
	ConfigId int64 `json:"configId,omitempty"`
}

// 13. 创建自动快照配置
type CreateAutoSnapshotConfigRequest struct {
	Request
	ClusterId            string `json:"-"`
	Name                 string `json:"name"`
	SnapshotType         string `json:"snapshotType,omitempty"`
	Period               string `json:"period"`
	ExpireAfter          string `json:"expireAfter,omitempty"`
	Remark               string `json:"remark,omitempty"`
	AllIndex             *bool  `json:"allIndex,omitempty"`
	SnapshotIndices      string `json:"snapshotIndices,omitempty"`
	SnapshotIndexPattern string `json:"snapshotIndexPattern,omitempty"`
	Enabled              *bool  `json:"enabled,omitempty"`
}

type CreateAutoSnapshotConfigResponse struct {
	ConfigId int64 `json:"configId,omitempty"`
}

// 14. 更新自动快照配置
type UpdateAutoSnapshotConfigRequest struct {
	Request
	ClusterId            string `json:"-"`
	ConfigId             string `json:"-"`
	Name                 string `json:"name"`
	SnapshotType         string `json:"snapshotType,omitempty"`
	Period               string `json:"period"`
	ExpireAfter          string `json:"expireAfter,omitempty"`
	Remark               string `json:"remark,omitempty"`
	AllIndex             *bool  `json:"allIndex,omitempty"`
	SnapshotIndices      string `json:"snapshotIndices,omitempty"`
	SnapshotIndexPattern string `json:"snapshotIndexPattern,omitempty"`
	Enabled              *bool  `json:"enabled,omitempty"`
}

type UpdateAutoSnapshotConfigResponse struct {
	ConfigId int64 `json:"configId,omitempty"`
}

// 15. 删除快照
type DeleteSnapshotRequest struct {
	Request
	ClusterId  string `json:"-"`
	SnapshotId string `json:"-"`
}

type DeleteSnapshotResponse struct {
	SnapshotId int64 `json:"snapshotId,omitempty"`
}

// 16. 删除自动快照配置
type DeleteAutoSnapshotConfigRequest struct {
	Request
	ClusterId string `json:"-"`
	ConfigId  string `json:"-"`
}

type DeleteAutoSnapshotConfigResponse struct {
	ConfigId int64 `json:"configId,omitempty"`
}

// 17. 批量启用禁用自动快照
type BatchEnableAutoSnapshotConfigsRequest struct {
	Request
	ClusterId    string  `json:"-"`
	Enabled      *bool   `json:"enabled"`
	ConfigIdList []int64 `json:"configIdList"`
}

type BatchEnableAutoSnapshotConfigsResponse struct {
	Enabled      bool    `json:"enabled,omitempty"`
	ConfigIdList []int64 `json:"configIdList,omitempty"`
}
