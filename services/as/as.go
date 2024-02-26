package as

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	GROUP_NAME      = "groupName"
	GROUP_ID        = "groupId"
	GROUP_ID_V2     = "groupid"
	ORDER           = "order"
	ORDER_BY        = "orderBy"
	PAGE_NO         = "pageNo"
	PAGE_SIZE       = "pageSize"
	KEW_WORD        = "keyWord"
	KEW_WORD_TYPE   = "keyWordType"
	AS_SCALING_DOWN = "scalingDown"
	AS_SCALING_UP   = "scalingUp"
	AS_ADJUST_NODE  = "adjustNode"
)

// CreateAsGroup 用于创建 AS 组
//
// 参数：
// req：CreateAsGroupRequest类型，包含创建 AS 组的请求参数
//
// 返回值：
// CreateAsGroupResponse：创建 AS 组的响应结果
// error：错误信息，如果请求成功则为 nil
func (c *Client) CreateAsGroup(req *CreateAsGroupRequest) (*CreateAsGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("CreateAsGroupRequest is nil")
	}
	result := &CreateAsGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupListUri()).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// DeleteAsGroup 用于删除一个弹性伸缩组
//
// req: 删除弹性伸缩组的请求参数
//
// 返回值: 错误信息，如果成功则返回nil
func (c *Client) DeleteAsGroup(req *DeleteAsGroupRequest) error {
	if req == nil {
		return fmt.Errorf("DeleteAsGroupRequest is nil")
	}
	if len(req.GroupIds) == 0 {
		return fmt.Errorf("DeleteAsGroupRequest GroupIds is empty")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupDeleteUri()).
		WithBody(req).
		WithMethod(http.POST).
		Do()
	if err != nil {
		return err
	}
	return nil
}

// ListAsGroup 方法用于获取指定用户下的As组列表
func (c *Client) ListAsGroup(req *ListAsGroupRequest) (*ListAsGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ListAsGroupRequest is nil")
	}
	if req.PageNo <= 0 {
		req.PageNo = 1
	}
	result := &ListAsGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsGroupListUri()).
		WithMethod(http.GET).
		WithQueryParamFilter(PAGE_NO, strconv.Itoa(req.PageNo)).
		WithQueryParamFilter(PAGE_SIZE, strconv.Itoa(req.PageNo)).
		WithQueryParamFilter(KEW_WORD_TYPE, req.KeyWordType).
		WithQueryParamFilter(KEW_WORD, req.KeyWord).
		WithQueryParamFilter(ORDER, req.Order).
		WithQueryParamFilter(ORDER_BY, req.OrderBy).
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
func (c *Client) ListAsNode(req *ListAsGroupRequest) (*ListAsNodeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ListAsNodeRequest is nil")
	}
	if req.PageNo <= 0 {
		req.PageNo = 1
	}
	if len(req.GroupId) <= 0 {
		return nil, fmt.Errorf("ListAsNodeRequest GroupId is empty")
	}
	result := &ListAsNodeResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(getAsNodeUri()).
		WithMethod(http.GET).
		WithQueryParamFilter(PAGE_NO, strconv.Itoa(req.PageNo)).
		WithQueryParamFilter(PAGE_SIZE, strconv.Itoa(req.PageNo)).
		WithQueryParamFilter(ORDER, req.Order).
		WithQueryParamFilter(ORDER_BY, req.OrderBy).
		WithQueryParamFilter(GROUP_ID_V2, req.GroupId).
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

// ListRecords 查询伸缩活动
func (c *Client) ListRecords(req *ListRecordsRequest) (*ListRecordsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ListRecordsRequest is nil")
	}
	var res ListRecordsResponse
	err := bce.NewRequestBuilder(c).
		WithURL(ListRecordsUrl).
		WithQueryParams(req.GetBceQueryParams()).
		WithMethod(http.GET).
		WithResult(&res).
		Do()
	return &res, err
}

// ExecRule 执行伸缩规则
func (c *Client) ExecRule(groupId, ruleId string) (*CreateDagResponse, error) {

	if groupId == "" || ruleId == "" {
		return nil, fmt.Errorf("groupId or ruleId is empty")
	}
	body := struct {
		RuleId string `json:"ruleId"`
	}{
		RuleId: ruleId,
	}
	url := fmt.Sprintf(BaseURL, groupId)
	var res CreateDagResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("execRule", "").
		WithMethod(http.POST).
		WithBody(body).
		WithResult(&res).
		Do()
	return &res, err
}

// ScalingUp  扩容
func (c *Client) ScalingUp(groupId string, req *ScalingUpRequest) (*CreateDagResponse, error) {
	if groupId == "" || req == nil {
		return nil, fmt.Errorf("groupId or ScalingUpRequest is empty")
	}
	url := fmt.Sprintf(BaseURL, groupId)
	var res CreateDagResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("scalingUp", "").
		WithMethod(http.POST).
		WithBody(req).
		WithResult(&res).
		Do()
	return &res, err
}

// ScalingDown 缩容
func (c *Client) ScalingDown(groupId string, req *ScalingDownRequest) error {
	if groupId == "" || req == nil {
		return fmt.Errorf("groupId or ScalingDownRequest is empty")
	}
	url := fmt.Sprintf(BaseURL, groupId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("scalingDown", "").
		WithMethod(http.POST).
		WithBody(req).
		Do()
	return err
}

// AdjustNode  伸缩组手动调整节点数
func (c *Client) AdjustNode(groupId string, adjustNum int) (*CreateDagResponse, error) {
	if groupId == "" {
		return nil, fmt.Errorf("groupId is empty")
	}
	req := struct {
		AdjustNum int `json:"adjustNum,omitempty"`
	}{
		AdjustNum: adjustNum,
	}
	url := fmt.Sprintf(BaseURL, groupId)
	var res CreateDagResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("adjustNode", "").
		WithMethod(http.POST).
		WithBody(req).
		WithResult(&res).
		Do()
	return &res, err
}

// AttachNode  手动添加节点
func (c *Client) AttachNode(groupId string, req *NodeRequest) error {
	if groupId == "" {
		return fmt.Errorf("groupId is empty")
	}
	url := fmt.Sprintf(BaseURL, groupId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("attachNode", "").
		WithMethod(http.POST).
		WithBody(req).
		Do()
	return err
}

// DetachNode  手动移除节点
func (c *Client) DetachNode(groupId string, req *NodeRequest) error {
	if groupId == "" {
		return fmt.Errorf("groupId is empty")
	}
	url := fmt.Sprintf(BaseURL, groupId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("detachNode", "").
		WithMethod(http.POST).
		WithBody(req).
		Do()
	return err
}

// CreateRule  创建伸缩规则
func (c *Client) CreateRule(req *RuleRequest) (*CreateRuleResult, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.RuleName) <= 0 {
		return nil, errors.New("ruleName should not be empty")
	}
	if len(req.GroupID) <= 0 {
		return nil, errors.New("groupId should not be empty")
	}
	if len(req.State) <= 0 {
		return nil, errors.New("state should not be empty")
	}
	if len(req.Type) <= 0 {
		return nil, errors.New("type should not be empty")
	}
	if len(req.ActionType) <= 0 {
		return nil, errors.New("actionType should not be empty")
	}
	if req.ActionNum <= 0 {
		return nil, errors.New("actionNum should not be empty")
	}
	if req.CooldownInSec <= 0 {
		return nil, errors.New("cooldownInSec should not be empty")
	}
	if len(req.PeriodStartTime) > 0 && !isUtcTime(req.PeriodStartTime) {
		return nil, errors.New("periodStartTime should be UTC format")
	}
	if len(req.PeriodEndTime) > 0 && !isUtcTime(req.PeriodEndTime) {
		return nil, errors.New("periodEndTime should be UTC format")
	}

	url := fmt.Sprintf(RuleURL)
	res := &CreateRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(req).
		WithResult(res).
		Do()
	return res, err
}

// UpdateRule  更新伸缩规则
func (c *Client) UpdateRule(ruleId string, req *RuleRequest) error {
	if len(ruleId) <= 0 {
		return errors.New("ruleId should not be empty")
	}
	if req == nil {
		return errors.New("req should not be empty")
	}
	if len(req.RuleName) <= 0 {
		return errors.New("ruleName should not be empty")
	}
	if len(req.GroupID) <= 0 {
		return errors.New("groupId should not be empty")
	}
	if len(req.State) <= 0 {
		return errors.New("state should not be empty")
	}
	if len(req.Type) <= 0 {
		return errors.New("type should not be empty")
	}
	if len(req.ActionType) <= 0 {
		return errors.New("actionType should not be empty")
	}
	if req.ActionNum <= 0 {
		return errors.New("actionNum should not be empty")
	}
	if req.CooldownInSec <= 0 {
		return errors.New("cooldownInSec should not be empty")
	}
	if len(req.PeriodStartTime) > 0 && !isUtcTime(req.PeriodStartTime) {
		return errors.New("periodStartTime should be UTC format")
	}
	if len(req.PeriodEndTime) > 0 && !isUtcTime(req.PeriodEndTime) {
		return errors.New("periodEndTime should be UTC format")
	}

	url := fmt.Sprintf(RuleIdURL, ruleId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.PUT).
		WithBody(req).
		Do()
	return err
}

// GetRuleList  查询伸缩规则
func (c *Client) GetRuleList(req *RuleListQuery) (*RuleVOListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.GroupID) <= 0 {
		return nil, errors.New("groupId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}

	url := fmt.Sprintf(RuleURL)
	res := &RuleVOListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithQueryParam("groupid", req.GroupID).
		WithQueryParam("keyword", req.Keyword).
		WithQueryParam("keywordType", req.KeywordType).
		WithQueryParam("order", req.Order).
		WithQueryParam("orderBy", req.OrderBy).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithResult(res).
		Do()
	return res, err
}

// GetRule  查询伸缩规则
func (c *Client) GetRule(ruleId string) (*RuleVO, error) {
	if len(ruleId) <= 0 {
		return nil, errors.New("ruleId should not be empty")
	}

	url := fmt.Sprintf(RuleIdURL, ruleId)
	res := &RuleVO{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithResult(res).
		Do()
	return res, err
}

// DeleteRule  删除伸缩规则
func (c *Client) DeleteRule(req *RuleDelRequest) error {
	if req == nil {
		return errors.New("req should not be empty")
	}
	if len(req.RuleIds) <= 0 && len(req.GroupIds) <= 0 {
		return errors.New("ruleIds and groupIds cannot both be empty\n ")
	}

	url := fmt.Sprintf(RuleURL)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("delete", "").
		WithMethod(http.POST).
		WithBody(req).
		Do()
	return err
}

// 判断传入字符串是否为utc时间，格式为：2023-12-12T00:00:00Z
func isUtcTime(str string) bool {
	_, err := time.Parse(time.RFC3339, str)
	return err == nil
}
