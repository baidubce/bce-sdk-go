package api

type DailySchedule struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type MigrationConfigCommon struct {
	Provider string `json:"provider,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Region   string `json:"region,omitempty"`
	Bucket   string `json:"bucket,omitempty"`
	Ak       string `json:"ak,omitempty"`
	Sk       string `json:"sk,omitempty"`
}

type MigrationPrefixSourceConfig struct {
	MigrationConfigCommon
	Prefixes        []string `json:"prefixes,omitempty"`
	Prefix          string   `json:"prefix,omitempty"`
	ObjectBeginTime int64    `json:"objectBeginTime,omitempty"` // default: -1
	ObjectEndTime   int64    `json:"objectEndTime,omitempty"`   // default: -1
}

type MigrationListSourceConfig struct {
	MigrationConfigCommon
	ListFileURL []string `json:"listFileURL,omitempty"`
}

type MigrationDestinationConfig struct {
	MigrationConfigCommon
	Prefix       string `json:"prefix,omitempty"`
	StorageClass string `json:"storageClass,omitempty"`
	Acl          string `json:"acl,omitempty"`
}

type MigrationType struct {
	Type                        string `json:"type,omitempty"`
	IncreaseScanIntervalInHours int32  `json:"increaseScanIntervalInHours,omitempty"`
	IncreaseTimes               int32  `json:"increaseTimes,omitempty"` // default: -1
}

type PerformanceConfig struct {
	StartTime     string `json:"startTime,omitempty"`
	EndTime       string `json:"endTime,omitempty"`
	BandWidthInMB int32  `json:"bandWidthInMB,omitempty"`
}

type ValidationMethodConfig struct {
	EnableCrc64ECMAValidation bool `json:"enableCRC64ECMAValidation"`
	EnableMD5Validation       bool `json:"enableMD5Validation"`
}

type PostMigrationArgsCommon struct {
	Name              string                     `json:"name,omitempty"`
	ScheduleStartTime int64                      `json:"scheduleStartTime,omitempty"` // unix timestamp
	DailySchedule     DailySchedule              `json:"dailySchedule,omitempty"`
	DestinationConfig MigrationDestinationConfig `json:"destinationConfig,omitempty"`
	Strategy          string                     `json:"strategy,omitempty"`
	MigrationType     MigrationType              `json:"migrationType,omitempty"`
	MigrationMode     string                     `json:"migrationMode,omitempty"`
	Qps               int64                      `json:"qps,omitempty"`
	ValidationConfig  ValidationMethodConfig     `json:"validationMethodConfig,omitempty"`
}

type PostMigrationArgs struct {
	PostMigrationArgsCommon
	SourceConfig       MigrationPrefixSourceConfig `json:"sourceConfig,omitempty"`
	PerformanceSetting []PerformanceConfig         `json:"performanceSetting,omitempty"`
	NotIncludeContent  []string                    `json:"notIncludeContent,omitempty"`
}

type PostMigrationFromListArgs struct {
	PostMigrationArgsCommon
	SourceConfig MigrationListSourceConfig `json:"sourceConfig,omitempty"`
}

type TaskIdList struct {
	TaskId []string `json:"taskID,omitempty"`
}

type MigrationResultCommon struct {
	Success   bool   `json:"success,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestId string `json:"requestId,omitempty"`
}

type PostMigrationResult struct {
	MigrationResultCommon
	Result TaskIdList `json:"result,omitempty"`
}

type ValidationResult struct {
	FinishedCount int64 `json:"finishedCount,omitempty"`
	FinishedBytes int64 `json:"finishedBytes,omitempty"`
	FailedCount   int64 `json:"failedCount,omitempty"`
}

type MigrationTaskInfo struct {
	MigrationBasicInfo
	CreateTime         int64                       `json:"createTime,omitempty"`
	TaskStartTime      int64                       `json:"taskStartTime,omitempty"`
	UpdateTime         int64                       `json:"updateTime,omitempty"`
	ScheduleStartTime  int64                       `json:"scheduleStartTime,omitempty"`
	DailySchedule      DailySchedule               `json:"dailySchedule,omitempty"`
	SourceConfig       MigrationPrefixSourceConfig `json:"sourceConfig,omitempty"`
	DestinationConfig  MigrationDestinationConfig  `json:"destinationConfig,omitempty"`
	Strategy           string                      `json:"strategy,omitempty"`
	MigrationType      MigrationType               `json:"migrationType,omitempty"`
	MigrationMode      string                      `json:"migrationMode,omitempty"`
	Qps                int64                       `json:"qps,omitempty"`
	PerformanceSetting []PerformanceConfig         `json:"performanceSetting"`
	TotalCount         int64                       `json:"totalCount,omitempty"`
	FinishedCount      int64                       `json:"finishedCount,omitempty"`
	FailedCount        int64                       `json:"failedCount,omitempty"`
	TotalBytes         int64                       `json:"totalBytes,omitempty"`
	FinishedBytes      int64                       `json:"finishedBytes,omitempty"`
	ValidationResult   ValidationResult            `json:"validationResult,omitempty"`
	ValidationConfig   ValidationMethodConfig      `json:"validationMethodConfig,omitempty"`
}

type GetMigrationInfo struct {
	MigrationResultCommon
	TaskInfos []MigrationTaskInfo `json:"result,omitempty"`
}

type MigrationResInfo struct {
	FailObjectListurl []string `json:"failObjectListURLs,omitempty"`
}

type MigrationResult struct {
	MigrationResultCommon
	Result MigrationResInfo `json:"result,omitempty"`
}

type MigrationBasicInfo struct {
	TaskId        string `json:"taskID,omitempty"`
	Name          string `json:"name,omitempty"`
	RunningStatus string `json:"runningStatus,omitempty"`
}

type ListMigrationInfo struct {
	MigrationResultCommon
	TaskInfos []MigrationBasicInfo `json:"result,omitempty"`
}
