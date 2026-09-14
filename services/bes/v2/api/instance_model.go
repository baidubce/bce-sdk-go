package api

// InstanceOperationRequest is the request shared by start/stop instance APIs.
type InstanceOperationRequest struct {
	ClusterId  string `json:"clusterId"`
	InstanceId string `json:"instanceId"`
}

// InstanceOperationResponse is the response shared by start/stop instance APIs.
type InstanceOperationResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// BatchInstanceOperationRequest is the request shared by batch start/stop instance APIs.
type BatchInstanceOperationRequest struct {
	ClusterId      string   `json:"clusterId"`
	InstanceIdList []string `json:"instanceIdList"`
}

// BatchInstanceOperationResponse is the response shared by batch start/stop instance APIs.
type BatchInstanceOperationResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// DeleteInstancesRequest is the request of deleting (scaling in) cluster instances.
type DeleteInstancesRequest struct {
	ClusterId      string   `json:"clusterId"`
	InstanceIdList []string `json:"instanceIdList"`
	ModuleType     string   `json:"moduleType"`
}

// DeleteInstancesResponse is the response of deleting (scaling in) cluster instances.
type DeleteInstancesResponse struct {
	Success bool         `json:"success"`
	Status  int          `json:"status"`
	Result  *OrderResult `json:"result"`
}

// ListScaleInInstancesRequest is the request of listing scale-in candidate instances.
type ListScaleInInstancesRequest struct {
	ClusterId  string `json:"clusterId"`
	ModuleType string `json:"moduleType"`
}

// ListScaleInInstancesResponse is the response of listing scale-in candidate instances.
type ListScaleInInstancesResponse struct {
	Success bool                    `json:"success"`
	Status  int                     `json:"status"`
	Result  *ScaleInInstancesResult `json:"result"`
}

// ScaleInInstancesResult carries the scale-in candidate instances grouped by zone.
type ScaleInInstancesResult struct {
	InstanceInfoList []ScaleInInstanceInfo `json:"instanceInfoList"`
}

// ScaleInInstanceInfo groups scale-in candidate instances by available zone.
type ScaleInInstanceInfo struct {
	AvailableZone string                `json:"availableZone"`
	InstanceList  []ScaleInInstanceItem `json:"instanceList"`
}

// ScaleInInstanceItem describes a single scale-in candidate instance.
type ScaleInInstanceItem struct {
	InstanceId string `json:"instanceId"`
	HostIp     string `json:"hostIp"`
	IsMaster   bool   `json:"isMaster"`
	IsEmpty    bool   `json:"isEmpty"`
}

// ConfirmDataMigrationRequest is the request of confirming a data migration.
type ConfirmDataMigrationRequest struct {
	ClusterId      string   `json:"clusterId"`
	ModuleType     string   `json:"moduleType"`
	InstanceIdList []string `json:"instanceIdList"`
}

// ConfirmDataMigrationResponse is the response of confirming a data migration.
type ConfirmDataMigrationResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// RollbackDataMigrationRequest is the request of rolling back a data migration.
type RollbackDataMigrationRequest struct {
	ClusterId string `json:"clusterId"`
	RequestId string `json:"requestId"`
}

// RollbackDataMigrationResponse is the response of rolling back a data migration.
type RollbackDataMigrationResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// ListDataMigrationInstancesRequest is the request of listing data migration candidate instances.
type ListDataMigrationInstancesRequest struct {
	ClusterId  string `json:"clusterId"`
	ModuleType string `json:"moduleType"`
}

// ListDataMigrationInstancesResponse is the response of listing data migration candidate instances.
type ListDataMigrationInstancesResponse struct {
	Success bool                      `json:"success"`
	Status  int                       `json:"status"`
	Result  *MigrationInstancesResult `json:"result"`
}

// MigrationInstancesResult carries the data migration candidate instances grouped by zone.
type MigrationInstancesResult struct {
	InstanceInfoList []MigrationInstanceInfo `json:"instanceInfoList"`
}

// MigrationInstanceInfo groups data migration candidate instances by available zone.
type MigrationInstanceInfo struct {
	AvailableZone string                  `json:"availableZone"`
	InstanceList  []MigrationInstanceItem `json:"instanceList"`
}

// MigrationInstanceItem describes a single data migration candidate instance.
type MigrationInstanceItem struct {
	InstanceId string `json:"instanceId"`
	HostIp     string `json:"hostIp"`
	IsMaster   bool   `json:"isMaster"`
}

// SuggestDataMigrationInstancesRequest is the request of getting system-suggested migration instances.
type SuggestDataMigrationInstancesRequest struct {
	ClusterId    string `json:"clusterId"`
	ModuleType   string `json:"moduleType"`
	MigrateCount int    `json:"migrateCount"`
}

// SuggestDataMigrationInstancesResponse is the response of getting system-suggested migration instances.
type SuggestDataMigrationInstancesResponse struct {
	Success bool                    `json:"success"`
	Status  int                     `json:"status"`
	Result  *MigrationSuggestResult `json:"result"`
}

// MigrationSuggestResult carries the system-suggested migration instances grouped by zone.
type MigrationSuggestResult struct {
	SuggestList []MigrationSuggestInfo `json:"suggestList"`
}

// MigrationSuggestInfo groups suggested migration instances by available zone.
type MigrationSuggestInfo struct {
	AvailableZone string                 `json:"availableZone"`
	InstanceList  []MigrationSuggestItem `json:"instanceList"`
}

// MigrationSuggestItem describes a single suggested migration instance.
type MigrationSuggestItem struct {
	InstanceId string `json:"instanceId"`
	HostIp     string `json:"hostIp"`
}
