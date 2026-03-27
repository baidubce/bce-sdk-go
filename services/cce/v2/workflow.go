package v2

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// GetInstanceGroupUpgradeComponentVersions queries upgradable component versions for an instance group.
func (c *Client) GetInstanceGroupUpgradeComponentVersions(args *GetInstanceGroupUpgradeComponentVersionsArgs) (*GetInstanceGroupUpgradeComponentVersionsResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.ClusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}
	if args.InstanceGroupID == "" {
		return nil, fmt.Errorf("instanceGroupID is empty")
	}

	result := &GetInstanceGroupUpgradeComponentVersionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceGroupUpgradeComponentVersionsURI(args.ClusterID, args.InstanceGroupID)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) CreateWorkflow(args *CreateWorkflowArgs) (*CreateWorkflowResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.ClusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}
	if args.Request == nil {
		return nil, fmt.Errorf("request is nil")
	}

	result := &CreateWorkflowResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterWorkflowURI(args.ClusterID)).
		WithBody(args.Request).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetWorkflow(args *GetWorkflowArgs) (*GetWorkflowResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.ClusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}
	if args.WorkflowID == "" {
		return nil, fmt.Errorf("workflowID is empty")
	}

	result := &GetWorkflowResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterWorkflowWithIDURI(args.ClusterID, args.WorkflowID)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) UpdateWorkflow(args *UpdateWorkflowArgs) (*UpdateWorkflowResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.ClusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}
	if args.WorkflowID == "" {
		return nil, fmt.Errorf("workflowID is empty")
	}
	if args.Request == nil {
		return nil, fmt.Errorf("request is nil")
	}

	result := &UpdateWorkflowResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClusterWorkflowWithIDURI(args.ClusterID, args.WorkflowID)).
		WithBody(args.Request).
		WithResult(result).
		Do()
	return result, err
}
