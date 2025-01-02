package v2

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateRemedyRule 创建自愈规则
func (c *Client) CreateRemedyRule(args *model.RemedyRule) (*model.CreateRemedyRuleResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &model.CreateRemedyRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(remedyRuleURI(args.ClusterID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListRemedyRules 获取自愈规则列表
func (c *Client) ListRemedyRules(args *model.ListRemedyRulesOptions) (*model.ListRemedyRulesResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &model.ListRemedyRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(remedyRuleURI(args.ClusterID)).
		WithQueryParamFilter("orderBy", args.OrderBy).
		WithQueryParamFilter("order", args.Order).
		WithQueryParamFilter("pageNo", fmt.Sprintf("%d", args.PageNo)).
		WithQueryParamFilter("pageSize", fmt.Sprintf("%d", args.PageSize)).
		WithResult(result).
		Do()

	return result, err
}

// GetRemedyRule 获取自愈规则
func (c *Client) GetRemedyRule(args *model.RemedyRuleOptions) (*model.GetRemedyRuleResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.RemedyRuleID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.GetRemedyRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRemedyRuleURI(args.ClusterID, args.RemedyRuleID)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateRemedyRule(args *model.RemedyRule) (*model.UpdateRemedyRuleResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.RemedyRuleID == "" {
		return nil, fmt.Errorf("RemedyRule is required")
	}

	result := &model.UpdateRemedyRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRemedyRuleURI(args.ClusterID, args.RemedyRuleID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteRemedyRule(args *model.RemedyRuleOptions) (*model.DeleteRemedyRuleResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.RemedyRuleID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.DeleteRemedyRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemedyRuleURI(args.ClusterID, args.RemedyRuleID)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) CheckWebhookAddress(args *model.CheckWebhookAddressRequest) (*model.CheckWebhookAddressResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	result := &model.CheckWebhookAddressResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCheckWebhookAddressURI(args.ClusterID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateInstanceGroupRemediation(args *model.BindingOrUnBindingRequest) (*model.BindingOrUnBindingResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	result := &model.BindingOrUnBindingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceGroupRemediationURI(args.ClusterID, args.InstanceGroupID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// BindInstanceGroupRemedyRule 绑定节点组自愈规则
func (c *Client) BindInstanceGroupRemedyRule(args *model.RemedyRuleBinding) (*model.BindInstanceGroupResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	if args.RemedyRuleID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.BindInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(instanceGroupRemedyRuleURI(args.ClusterID, args.InstanceGroupID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UnbindInstanceGroupRemedyRule 解绑节点组自愈规则
func (c *Client) UnbindInstanceGroupRemedyRule(args *model.RemedyRuleBinding) (*model.UnbindInstanceGroupResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	if args.RemedyRuleID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.UnbindInstanceGroupResponse{}
	fmt.Println(getInstanceGroupRemedyRuleURI(args.ClusterID, args.InstanceGroupID, args.RemedyRuleID))
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getInstanceGroupRemedyRuleURI(args.ClusterID, args.InstanceGroupID, args.RemedyRuleID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListInstanceGroupRemedyTasks 获取节点组维修任务详情列表
func (c *Client) ListInstanceGroupRemedyTasks(args *model.ListRemedyTaskOptions) (*model.ListRemedyTasksResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	result := &model.ListRemedyTasksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(instanceGroupRemedyTaskURI(args.ClusterID, args.InstanceGroupID)).
		WithQueryParamFilter("pageNo", fmt.Sprintf("%d", args.PageNo)).
		WithQueryParamFilter("pageSize", fmt.Sprintf("%d", args.PageSize)).
		WithQueryParam("startTime", args.StartTime).
		WithQueryParam("stopTime", args.StopTime).
		WithResult(result).
		Do()

	return result, err
}

// GetInstanceGroupRemedyTask 获取节点组维修任务详情
func (c *Client) GetInstanceGroupRemedyTask(args *model.GetRemedyTaskOptions) (*model.GetRemedyTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	if args.RemedyTaskID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.GetRemedyTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceGroupRemedyTaskURI(args.ClusterID, args.InstanceGroupID, args.RemedyTaskID)).
		WithResult(result).
		Do()

	return result, err
}

// AuthRepairInstanceGroupRemedyTask 授权维修任务
func (c *Client) AuthRepairInstanceGroupRemedyTask(args *model.AuthorizeTaskArgs) (*model.AuthorizeTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	if args.RemedyTaskID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.AuthorizeTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceGroupRemedyTaskAuthRepairURI(args.ClusterID, args.InstanceGroupID, args.RemedyTaskID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ConfirmRemedyTask 确认恢复
func (c *Client) ConfirmRemedyTask(args *model.ConfirmRemedyTask) (*model.ConfirmRemedyTaskResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	if args.RemedyTaskID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.ConfirmRemedyTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceGroupRemedyTaskConfirmURI(args.ClusterID, args.InstanceGroupID, args.RemedyTaskID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// RequireRepairAuth 请求授权
func (c *Client) RequireRepairAuth(args *model.RequestAuthorizeArgs) (*model.RequestAuthorizeResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	if args.ClusterID == "" {
		return nil, fmt.Errorf("ClusterID is required")
	}

	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("InstanceGroupID is required")
	}

	if args.RemedyTaskID == "" {
		return nil, fmt.Errorf("RemedyRuleID is required")
	}

	result := &model.RequestAuthorizeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceGroupRemedyTaskRequestAuthURI(args.ClusterID, args.InstanceGroupID, args.RemedyTaskID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}
