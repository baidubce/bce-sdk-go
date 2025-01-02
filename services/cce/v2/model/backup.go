package model

import (
	"time"

	backupv1 "github.com/baidubce/bce-sdk-go/services/cce/v2/model/backup/api/v1"
)

type BackupRepositoryStatus string
type BackupTaskType string
type RestoreResourcePolicy string
type BackupScope string

const (
	BackupRepositoryStatusavailable   BackupRepositoryStatus = "Available"
	BackupRepositoryStatusunavailable BackupRepositoryStatus = "Unavailable"

	BackupTaskTypeScheduled BackupTaskType = "Scheduled"
	BackupTaskTypeManual    BackupTaskType = "Manual"

	RestoreResourcePolicyNONE   RestoreResourcePolicy = "None"
	RestoreResourcePolicyUPDATE RestoreResourcePolicy = "Update"

	BackupScopeSpecified BackupScope = "Specified"
	BackupScopeAll       BackupScope = "All"
)

// BackupRepository
type BackupRepository struct {
	Name          string                 `json:"name"`
	RepositoryID  string                 `json:"repositoryID"`
	Region        string                 `json:"region"`
	BucketName    string                 `json:"bucketName"`
	BucketSubPath string                 `json:"bucketSubPath"`
	Status        BackupRepositoryStatus `json:"status"`
	CreateTime    time.Time              `json:"createTime"`
	ErrMsg        string                 `json:"errMsg"`
}

type CreateBackupRequest struct {
	BackupRepository *BackupRepository `json:"backupRepository"`
}

type GetBackupRepositoryResp struct {
	RequestID        string            `json:"requestId"`
	BackupRepository *BackupRepository `json:"backupRepository"`
}

type BackupRepositoryPage struct {
	PageNo             int                 `json:"pageNo"`
	PageSize           int                 `json:"pageSize"`
	TotalCount         int                 `json:"totalCount"`
	BackupRepositories []*BackupRepository `json:"backupRepositorys"`
}

type ListBackupRepositoryResp struct {
	RequestID            string               `json:"requestId"`
	BackupRepositoryPage BackupRepositoryPage `json:"backupRepositoryPage"`
}

type ListBackupNameFromRepoResp struct {
	RequestID   string      `json:"requestId"`
	BackupInfos BackupInfos `json:"backupInfos"`
}

type BackupInfos struct {
	BackupNames []string `json:"backupNames"`
	Prefix      string   `json:"prefix"`
	BucketName  string   `json:"bucketName"`
}

type GetBackupInfoByNameResp struct {
	RequestID     string         `json:"requestId"`
	BackupDetails *BackupDetails `json:"backupDetails"`
}

type BackupDetails struct {
	TaskName           string            `json:"taskName"`
	BackupClusterID    string            `json:"backupClusterID"`
	LabelSelector      map[string]string `json:"labelSelector,omitempty"`
	IncludedNamespaces []string          `json:"includedNamespaces,omitempty"`
	ExcludedNamespaces []string          `json:"excludedNamespaces,omitempty"`
	IncludedResources  []string          `json:"includedResources,omitempty"`
	ExcludedResources  []string          `json:"excludedResources,omitempty"`

	Task *backupv1.Backup `json:"task"`
}

type CreateBackupRepositoryResp struct {
	RequestID        string            `json:"requestId"`
	BackupRepository *BackupRepository `json:"backupRepository"`
}

type DeleteBackupRepositoryReq struct {
	RepositoryID string `json:"repositoryID"`
}

type DeleteBackupRepositoryResp struct {
	RequestID string `json:"requestId"`
}

// BackupTasks
type BackupTaskArgs struct {
	Region   string `json:"region"`
	TaskName string `json:"taskName"`
	//BucketName           string            `json:"bucketName"`
	RepositoryID string `json:"repositoryID"`
	//RepositoryName       string            `json:"repositoryName"`
	BackupScope          BackupScope       `json:"backupScope"`
	IncludedNamespaces   []string          `json:"includedNamespaces"`
	ExcludedNamespaces   []string          `json:"excludedNamespaces"`
	IncludedResources    []string          `json:"includedResources"`
	ExcludedResources    []string          `json:"excludedResources"`
	LabelSelector        map[string]string `json:"labelSelector"`
	BackupExpirationDays int               `json:"backupExpirationDays"`

	Schedule string `json:"schedule"`
}

type CreateBackupTaskRequest struct {
	ClusterID      string          `json:"clusterID"`
	BackupTaskArgs *BackupTaskArgs `json:"backupTaskArgs"`
}

type CreateBackupTaskResponse struct {
	RequestID string           `json:"requestId"`
	Task      *backupv1.Backup `json:"task"`
}

type CreateCCEV1BackupTaskResponse struct {
	RequestID string           `json:"requestId"`
	Task      *backupv1.Backup `json:"task"`
}

type DeleteBackupTaskRequest struct {
	ClusterID    string `json:"clusterId"`
	BackupTaskID string `json:"backupTaskID"`
}

type DeleteBackupTaskResponse struct {
	RequestID string `json:"requestId"`
}

type ListTasksRequest struct {
	ClusterID    string `json:"clusterId"`
	PageNo       int    `json:"pageNo"`
	PageSize     int    `json:"pageSize"`
	OrderBy      string `json:"orderBy"`
	Order        string `json:"order"`
	Phases       string `json:"phases"`
	Keyword      string `json:"keyword"`
	KeywordType  string `json:"keywordType"`
	CrossCluster string `json:"crossCluster"`
}

type ListBackupTasksResponse struct {
	RequestID       string           `json:"requestId"`
	BackupTasksPage *BackupTasksPage `json:"backupTasksPage"`
}

type CCEV1ListBackupTasksResponse struct {
	RequestID       string                `json:"requestId"`
	BackupTasksPage *CCEV1BackupTasksPage `json:"backupTasksPage"`
}

type BackupTasksPage struct {
	PageNo      int           `json:"pageNo"`
	PageSize    int           `json:"pageSize"`
	TotalCount  int           `json:"totalCount"`
	BackupTasks []*BackupTask `json:"backupTasks"`
	//NextPageMark string        `json:"nextPageMark"`
}

type CCEV1BackupTasksPage struct {
	PageNo      int                      `json:"pageNo"`
	PageSize    int                      `json:"pageSize"`
	TotalCount  int                      `json:"totalCount"`
	BackupTasks []*CCEV1BackupBackupTask `json:"backupTasks"`
	//NextPageMark string        `json:"nextPageMark"`
}

type BackupTask struct {
	TaskName       string          `json:"taskName"`
	TaskID         string          `json:"taskID"`
	BackupType     BackupTaskType  `json:"backupType"`
	RepositoryName string          `json:"repositoryName"`
	RepositoryID   string          `json:"repositoryID"`
	ClusterID      string          `json:"clusterID"`
	BackupScope    BackupScope     `json:"backupScope"`
	Task           backupv1.Backup `json:"task"`

	ErrMsg               string            `json:"errMsg"`
	BackupExpirationDays int               `json:"backupExpirationDays"`
	LabelSelector        map[string]string `json:"labelSelector,omitempty"`
	IncludedNamespaces   []string          `json:"includedNamespaces,omitempty"`
	ExcludedNamespaces   []string          `json:"excludedNamespaces,omitempty"`
	IncludedResources    []string          `json:"includedResources,omitempty"`
	ExcludedResources    []string          `json:"excludedResources,omitempty"`
	//Schedule         string         `json:"schedule"`
	//ScheduleDescribe string         `json:"scheduleDescribe "`
}

type CCEV1BackupBackupTask struct {
	TaskName       string          `json:"taskName"`
	TaskID         string          `json:"taskID"`
	BackupType     BackupTaskType  `json:"backupType"`
	RepositoryName string          `json:"repositoryName"`
	RepositoryID   string          `json:"repositoryID"`
	ClusterID      string          `json:"clusterID"`
	BackupScope    BackupScope     `json:"backupScope"`
	Task           backupv1.Backup `json:"task"`

	ErrMsg               string            `json:"errMsg"`
	BackupExpirationDays int               `json:"backupExpirationDays"`
	LabelSelector        map[string]string `json:"labelSelector,omitempty"`
	IncludedNamespaces   []string          `json:"includedNamespaces,omitempty"`
	ExcludedNamespaces   []string          `json:"excludedNamespaces,omitempty"`
	IncludedResources    []string          `json:"includedResources,omitempty"`
	ExcludedResources    []string          `json:"excludedResources,omitempty"`
	//Schedule         string         `json:"schedule"`
	//ScheduleDescribe string         `json:"scheduleDescribe "`
}

type GetBackupTaskRequest struct {
	ClusterID    string `json:"clusterId"`
	BackupTaskID string `json:"backupTaskID"`
}

type GetBackupTaskResponse struct {
	RequestID  string      `json:"requestId"`
	BackupTask *BackupTask `json:"backupTask"`
}

type GetCCEV1BackupTaskResponse struct {
	RequestID  string                 `json:"requestId"`
	BackupTask *CCEV1BackupBackupTask `json:"backupTask"`
}

// 定时备份
type CreateScheduleRulesRequest struct {
	//Schedule       string          `json:"schedule"`
	ClusterID      string          `json:"clusterID"`
	BackupTaskArgs *BackupTaskArgs `json:"backupTaskArgs"`
}

type CreateScheduleRulesResponse struct {
	RequestID string             `json:"requestId"`
	Task      *backupv1.Schedule `json:"task"`
}

type CreateCCEScheduleRulesResponse struct {
	RequestID string             `json:"requestId"`
	Task      *backupv1.Schedule `json:"task"`
}

type DeleteScheduleTaskRequest struct {
	ClusterID      string `json:"clusterId"`
	ScheduleTaskID string `json:"scheduleTaskID"`
}

type DeleteScheduleTaskResponse struct {
	RequestID string `json:"requestId"`
}

type GetScheduleTaskRequest struct {
	ClusterID      string `json:"clusterId"`
	ScheduleTaskID string `json:"scheduleTaskID"`
}

type GetScheduleTaskResponse struct {
	RequestID    string        `json:"requestId"`
	ScheduleTask *ScheduleTask `json:"scheduleTask"`
}
type GetCCEScheduleTaskResponse struct {
	RequestID    string             `json:"requestId"`
	ScheduleTask *CCEV1ScheduleTask `json:"scheduleTask"`
}

type GetScheduleDescribeResponse struct {
	RequestID        string `json:"requestId"`
	ScheduleDescribe string `json:"scheduleDescribe"`
}

type ListScheduleTasksResponse struct {
	RequestID         string             `json:"requestId"`
	ScheduleTasksPage *ScheduleTasksPage `json:"scheduleTasksPage"`
}

type ListCCEScheduleTasksResponse struct {
	RequestID         string                  `json:"requestId"`
	ScheduleTasksPage *CCEV1ScheduleTasksPage `json:"scheduleTasksPage"`
}

type ScheduleTasksPage struct {
	PageNo        int             `json:"pageNo"`
	PageSize      int             `json:"pageSize"`
	TotalCount    int             `json:"totalCount"`
	ScheduleTasks []*ScheduleTask `json:"scheduleTasks"`
	//NextPageMark string        `json:"nextPageMark"`
}

type CCEV1ScheduleTasksPage struct {
	PageNo        int                  `json:"pageNo"`
	PageSize      int                  `json:"pageSize"`
	TotalCount    int                  `json:"totalCount"`
	ScheduleTasks []*CCEV1ScheduleTask `json:"scheduleTasks"`
	//NextPageMark string        `json:"nextPageMark"`
}

type ScheduleTask struct {
	TaskName         string            `json:"taskName"`
	TaskID           string            `json:"taskID"`
	Schedule         string            `json:"schedule"`
	ScheduleDescribe string            `json:"scheduleDescribe"`
	RepositoryName   string            `json:"repositoryName"`
	RepositoryID     string            `json:"repositoryID"`
	ClusterID        string            `json:"clusterID"`
	ErrMsg           string            `json:"errMsg"`
	Task             backupv1.Schedule `json:"task"`

	BackupExpirationDays int               `json:"backupExpirationDays"`
	LabelSelector        map[string]string `json:"labelSelector,omitempty"`
	IncludedNamespaces   []string          `json:"includedNamespaces,omitempty"`
	ExcludedNamespaces   []string          `json:"excludedNamespaces,omitempty"`
	IncludedResources    []string          `json:"includedResources,omitempty"`
	ExcludedResources    []string          `json:"excludedResources,omitempty"`
}

type CCEV1ScheduleTask struct {
	TaskName         string            `json:"taskName"`
	TaskID           string            `json:"taskID"`
	Schedule         string            `json:"schedule"`
	ScheduleDescribe string            `json:"scheduleDescribe"`
	RepositoryName   string            `json:"repositoryName"`
	RepositoryID     string            `json:"repositoryID"`
	ClusterID        string            `json:"clusterID"`
	ErrMsg           string            `json:"errMsg"`
	Task             backupv1.Schedule `json:"task"`

	BackupExpirationDays int               `json:"backupExpirationDays"`
	LabelSelector        map[string]string `json:"labelSelector,omitempty"`
	IncludedNamespaces   []string          `json:"includedNamespaces,omitempty"`
	ExcludedNamespaces   []string          `json:"excludedNamespaces,omitempty"`
	IncludedResources    []string          `json:"includedResources,omitempty"`
	ExcludedResources    []string          `json:"excludedResources,omitempty"`
}

// 备份恢复
type RestoreTaskArgs struct {
	//Region                 string                `json:"region"`
	BackupClusterID string `json:"backupClusterID"`

	TaskName               string                `json:"taskName"`
	RepositoryID           string                `json:"repositoryID"`
	RepositoryName         string                `json:"repositoryName"`
	BackupTaskID           string                `json:"backupTaskID"`
	BackupTaskName         string                `json:"backupTaskName"`
	BackupScope            BackupScope           `json:"backupScope"`
	IncludedNamespaces     []string              `json:"includedNamespaces,omitempty"`
	ExcludedNamespaces     []string              `json:"excludedNamespaces,omitempty"`
	ExistingResourcePolicy RestoreResourcePolicy `json:"existingResourcePolicy"`
}

type CreateRestoreTaskRequest struct {
	ClusterID       string           `json:"clusterID"`
	RestoreTaskArgs *RestoreTaskArgs `json:"restoreTaskArgs"`
}

type CreateRestoreTaskResponse struct {
	RequestID string            `json:"requestId"`
	Task      *backupv1.Restore `json:"task"`
}

type CreateCCERestoreTaskResponse struct {
	RequestID string            `json:"requestId"`
	Task      *backupv1.Restore `json:"task"`
}

type DeleteRestoreTaskRequest struct {
	ClusterID     string `json:"clusterId"`
	RestoreTaskID string `json:"restoreTaskID"`
}

type DeleteRestoreTaskResponse struct {
	RequestID string `json:"requestId"`
}

type GetRestoreTaskRequest struct {
	ClusterID     string `json:"clusterId"`
	RestoreTaskID string `json:"restoreTaskID"`
}

type GetRestoreTaskResponse struct {
	RequestID   string       `json:"requestId"`
	RestoreTask *RestoreTask `json:"restoreTask"`
}

type GetCCERestoreTaskResponse struct {
	RequestID   string          `json:"requestId"`
	RestoreTask *CCERestoreTask `json:"restoreTask"`
}

type ListRestoreTasksResponse struct {
	RequestID        string            `json:"requestId"`
	RestoreTasksPage *RestoreTasksPage `json:"restoreTasksPage"`
}

type ListCCERestoreTasksResponse struct {
	RequestID        string               `json:"requestId"`
	RestoreTasksPage *CCERestoreTasksPage `json:"restoreTasksPage"`
}

type RestoreTasksPage struct {
	PageNo       int            `json:"pageNo"`
	PageSize     int            `json:"pageSize"`
	TotalCount   int            `json:"totalCount"`
	RestoreTasks []*RestoreTask `json:"restoreTasks"`
}

type CCERestoreTasksPage struct {
	PageNo       int               `json:"pageNo"`
	PageSize     int               `json:"pageSize"`
	TotalCount   int               `json:"totalCount"`
	RestoreTasks []*CCERestoreTask `json:"restoreTasks"`
}

type RestoreTask struct {
	ClusterID string `json:"clusterID"`
	TaskName  string `json:"taskName"`
	TaskID    string `json:"taskID"`
	//RepositoryID           string                `json:"repositoryID"`
	RepositoryName string `json:"repositoryName"`
	BackupTaskName string `json:"backupTaskName"`
	ErrMsg         string `json:"errMsg"`

	Task backupv1.Restore `json:"task"`
}

type CCERestoreTask struct {
	ClusterID string `json:"clusterID"`
	TaskName  string `json:"taskName"`
	TaskID    string `json:"taskID"`
	//RepositoryID           string                `json:"repositoryID"`
	RepositoryName string `json:"repositoryName"`
	BackupTaskName string `json:"backupTaskName"`
	ErrMsg         string `json:"errMsg"`

	Task backupv1.Restore `json:"task"`
}
