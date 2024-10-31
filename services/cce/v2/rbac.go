package v2

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateRBAC 为用户创建 RBAC
func (c *Client) CreateRBAC(args *model.RBACRequest) (*model.CreateRBACResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.ClusterID == "" {
		return nil, fmt.Errorf("clusterID is required")
	}
	if args.ClusterID == model.AllCluster && (args.Namespace != "" && args.Namespace != model.AllNamespace) {
		return nil, fmt.Errorf("namespace cannot be set when clusterID='all'")
	}

	result := &model.CreateRBACResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRBACURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// RenewRBAC 重置 RBAC kubeconfig
func (c *Client) RenewRBAC(args *model.RBACRequest) (*model.RBACResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.ClusterID == "" || args.ClusterID == model.AllCluster {
		return nil, fmt.Errorf("clusterID is required and cannot be 'all'")
	}

	result := &model.RBACResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRBACURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// DeleteRBAC 删除 RBAC 权限
func (c *Client) DeleteRBAC(args *model.RBACRequest) (*model.RBACResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.ClusterID == "" || args.ClusterID == model.AllCluster {
		return nil, fmt.Errorf("clusterID is required and cannot be 'all'")
	}

	result := &model.RBACResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRBACURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListRBAC 查询 RBAC 权限
func (c *Client) ListRBAC(args *model.RBACRequest) (*model.GetRBACResponse, error) {
	// 参数校验
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.UserID == "" {
		return nil, fmt.Errorf("userID is required")
	}
	result := &model.GetRBACResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRBACURI()).
		WithQueryParamFilter("userID", args.UserID).
		WithResult(result).
		Do()

	return result, err
}
