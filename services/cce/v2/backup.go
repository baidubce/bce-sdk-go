package v2

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"

	backuoModel "github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateBackupRepositorys 创建备份仓库
func (c *Client) CreateBackupRepositorys(args *backuoModel.CreateBackupRequest) (*backuoModel.CreateBackupRepositoryResp, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	fmt.Println(getBackupRepoURL(""))
	result := &backuoModel.CreateBackupRepositoryResp{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBackupRepoURL("")).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListBackupRepositorys 获取备份仓库列表
func (c *Client) ListBackupRepositorys(args *backuoModel.ListTasksRequest) (*backuoModel.ListBackupRepositoryResp, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &backuoModel.ListBackupRepositoryResp{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackupRepoURL("")).
		WithQueryParamFilter("orderBy", args.OrderBy).
		WithQueryParamFilter("order", args.Order).
		WithQueryParamFilter("pageNo", fmt.Sprintf("%d", args.PageNo)).
		WithQueryParamFilter("pageSize", fmt.Sprintf("%d", args.PageSize)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("keywordType", args.KeywordType).
		WithResult(result).
		Do()

	return result, err
}

// DeleteBackupRepositorys 删除备份仓库
func (c *Client) DeleteBackupRepositorys(args *backuoModel.DeleteBackupRepositoryReq) (*backuoModel.DeleteBackupRepositoryResp, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.RepositoryID == "" {
		return nil, fmt.Errorf("RepositoryID is required")
	}

	result := &backuoModel.DeleteBackupRepositoryResp{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getBackupRepoURL(args.RepositoryID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateBackupTasks 创建备份任务
func (c *Client) CreateBackupTasks(args *backuoModel.CreateBackupTaskRequest) (*backuoModel.CreateBackupTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &backuoModel.CreateBackupTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBackupTaskURL(args.ClusterID, "")).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListBackupTasks 获取备份任务列表
func (c *Client) ListBackupTasks(args *backuoModel.ListTasksRequest) (*backuoModel.ListBackupTasksResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &backuoModel.ListBackupTasksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackupTaskURL(args.ClusterID, "")).
		WithQueryParamFilter("orderBy", args.OrderBy).
		WithQueryParamFilter("order", args.Order).
		WithQueryParamFilter("pageNo", fmt.Sprintf("%d", args.PageNo)).
		WithQueryParamFilter("pageSize", fmt.Sprintf("%d", args.PageSize)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("keywordType", args.KeywordType).
		WithQueryParamFilter("crossCluster", args.CrossCluster).
		WithResult(result).
		Do()

	return result, err
}

// DeleteBackupTasks 删除备份任务
func (c *Client) DeleteBackupTasks(args *backuoModel.DeleteBackupTaskRequest) (*backuoModel.DeleteBackupTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.BackupTaskID == "" {
		return nil, fmt.Errorf("RepositoryID is required")
	}

	result := &backuoModel.DeleteBackupTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getBackupTaskURL(args.ClusterID, args.BackupTaskID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateBackupScheduleRules 创建定时备份任务
func (c *Client) CreateBackupScheduleRules(args *backuoModel.CreateScheduleRulesRequest) (*backuoModel.CreateScheduleRulesResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &backuoModel.CreateScheduleRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getScheduleBackupTaskURL(args.ClusterID, "")).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListBackupScheduleRules 获取定时备份任务列表
func (c *Client) ListBackupScheduleRules(args *backuoModel.ListTasksRequest) (*backuoModel.ListScheduleTasksResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &backuoModel.ListScheduleTasksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getScheduleBackupTaskURL(args.ClusterID, "")).
		WithQueryParamFilter("orderBy", args.OrderBy).
		WithQueryParamFilter("order", args.Order).
		WithQueryParamFilter("pageNo", fmt.Sprintf("%d", args.PageNo)).
		WithQueryParamFilter("pageSize", fmt.Sprintf("%d", args.PageSize)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("keywordType", args.KeywordType).
		WithResult(result).
		Do()

	return result, err
}

// DeleteBackupScheduleRules 删除定时备份任务
func (c *Client) DeleteBackupScheduleRules(args *backuoModel.DeleteScheduleTaskRequest) (*backuoModel.DeleteBackupTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.ScheduleTaskID == "" {
		return nil, fmt.Errorf("ScheduleTaskID is required")
	}

	result := &backuoModel.DeleteBackupTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getScheduleBackupTaskURL(args.ClusterID, args.ScheduleTaskID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateRestoreTasks 创建恢复任务
func (c *Client) CreateRestoreTasks(args *backuoModel.CreateRestoreTaskRequest) (*backuoModel.CreateCCERestoreTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &backuoModel.CreateCCERestoreTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRestoreBackupTaskURL(args.ClusterID, "")).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListRestoreTasks 获取恢复任务列表
func (c *Client) ListRestoreTasks(args *backuoModel.ListTasksRequest) (*backuoModel.ListCCERestoreTasksResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &backuoModel.ListCCERestoreTasksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRestoreBackupTaskURL(args.ClusterID, "")).
		WithQueryParamFilter("orderBy", args.OrderBy).
		WithQueryParamFilter("order", args.Order).
		WithQueryParamFilter("pageNo", fmt.Sprintf("%d", args.PageNo)).
		WithQueryParamFilter("pageSize", fmt.Sprintf("%d", args.PageSize)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("keywordType", args.KeywordType).
		WithResult(result).
		Do()

	return result, err
}

// DeleteRestoreTasks 删除恢复任务
func (c *Client) DeleteRestoreTasks(args *backuoModel.DeleteRestoreTaskRequest) (*backuoModel.DeleteRestoreTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.RestoreTaskID == "" {
		return nil, fmt.Errorf("RestoreTaskID is required")
	}

	result := &backuoModel.DeleteRestoreTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRestoreBackupTaskURL(args.ClusterID, args.RestoreTaskID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}
