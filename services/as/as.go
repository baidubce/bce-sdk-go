package as

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

const (
	GROUP_NAME      = "groupName"
	GROUP_ID        = "groupId"
	MARKER          = "marker"
	MAX_KEYS        = "maxKeys"
	MANNER          = "manner"
	AS_SCALING_DOWN = "scalingDown"
	AS_SCALING_UP   = "scalingUp"
	AS_ADJUST_NODE  = "adjustNode"
)

// ListAsGroup 方法用于获取指定用户下的As组列表
func (c *Client) ListAsGroup(req *ListAsGroupRequest) (*ListAsGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ListAsGroupRequest is nil")
	}
	if req.MaxKeys <= 0 || req.MaxKeys > 1000 {
		req.MaxKeys = 1000
	}
	result := &ListAsGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupListUri()).
		WithMethod(http.GET).
		WithQueryParamFilter(MANNER, MARKER).
		WithQueryParamFilter(MARKER, req.Marker).
		WithQueryParamFilter(MAX_KEYS, strconv.Itoa(req.MaxKeys)).
		WithQueryParamFilter(GROUP_NAME, req.GroupName).
		WithResult(result).
		Do()
	return result, err
}

// GetAsGroup 根据groupId获取AsGroup信息
func (c *Client) GetAsGroup(req *GetAsGroupRequest) (*GetAsGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetAsGroupRequest is nil")
	}
	result := &GetAsGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupDetailUri(req.GroupId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// IncreaseAsGroup 伸缩组扩容，用于在指定伸缩组下添加节点
func (c *Client) IncreaseAsGroup(req *IncreaseAsGroupRequest) error {
	if req == nil {
		return fmt.Errorf("IncreaseAsGroupRequest is nil")
	}
	if len(req.GroupId) == 0 {
		return fmt.Errorf("IncreaseAsGroupRequest GroupId is empty")
	}
	if req.NodeCount <= 0 {
		return fmt.Errorf("IncreaseAsGroupRequest NodeCount is invalid")
	}
	if req.Zone == nil || len(req.Zone) == 0 {
		return fmt.Errorf("IncreaseAsGroupRequest Zone is nil")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupUri(req.GroupId)).
		WithQueryParam(AS_SCALING_UP, "").
		WithBody(req).
		WithMethod(http.POST).
		Do()
	if err != nil {
		return err
	}
	return nil
}

// DecreaseAsGroup 伸缩组缩容，用于伸缩组下节点的缩容
func (c *Client) DecreaseAsGroup(req *DecreaseAsGroupRequest) error {
	if req == nil {
		return fmt.Errorf("DecreaseAsGroupRequest is nil")
	}
	if len(req.GroupId) == 0 {
		return fmt.Errorf("IncreaseAsGroupRequest GroupId is empty")
	}
	if req.Nodes == nil || len(req.Nodes) == 0 {
		return fmt.Errorf("DecreaseAsGroupRequest Nodes is nil")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupUri(req.GroupId)).
		WithQueryParam(AS_SCALING_DOWN, "").
		WithBody(req).
		WithMethod(http.POST).
		Do()
	if err != nil {
		return err
	}
	return nil
}

// ListAsNode 方法用于获取节点列表
func (c *Client) ListAsNode(req *ListAsNodeRequest) (*ListAsNodeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ListAsNodeRequest is nil")
	}
	if req.MaxKeys <= 0 || req.MaxKeys > 1000 {
		req.MaxKeys = 1000
	}
	result := &ListAsNodeResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsNodeUri()).
		WithMethod(http.GET).
		WithQueryParamFilter(MANNER, MARKER).
		WithQueryParamFilter(MARKER, req.Marker).
		WithQueryParamFilter(MAX_KEYS, strconv.Itoa(req.MaxKeys)).
		WithQueryParamFilter(GROUP_ID, req.GroupId).
		WithResult(result).
		Do()
	return result, err
}

// AdjustAsGroup 将伸缩组节点调整到指定值。
func (c *Client) AdjustAsGroup(req *AdjustAsGroupRequest) error {
	if req == nil {
		return fmt.Errorf("AdjustAsGroupByNodeIdRequest is nil")
	}
	if len(req.GroupId) == 0 {
		return fmt.Errorf("AdjustAsGroupByNodeIdRequest GroupId is empty")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupUri(req.GroupId)).
		WithQueryParam(AS_ADJUST_NODE, "").
		WithMethod(http.POST).
		WithBody(req).
		Do()
	return err
}
