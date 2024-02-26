package bcm

import (
	"errors"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// 创建警报配置
func (c *Client) CreateSiteOnce(request *SiteOnceRequest) (*SiteOnceBaseResponse, error) {
	if request == nil || request.ProtocolType == "" {
		return nil, errors.New("SiteOnceRequest and ProtocolType is must not empty")
	}
	url := fmt.Sprintf(CreateSiteOnce, request.ProtocolType)
	var response SiteOnceBaseResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(request).
		WithResult(&response).
		Do()
	return &response, err
}

// 历史列表
func (c *Client) ListSiteOnceTasks(request *SiteOnceTaskRequest) (*SiteOnceTaskListResponse, error) {
	if request == nil {
		return nil, errors.New("SiteOnceTaskRequest is must not empty")
	}
	url := ListSiteOnceTasks
	var response SiteOnceTaskListResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(request).
		WithResult(&response).
		Do()
	return &response, err
}

// 删除任务
func (c *Client) DeleteSiteOnceTask(request *SiteOnceTaskRequest) (*SiteOnceBaseResponse, error) {
	if request == nil {
		return nil, errors.New("SiteOnceTaskRequest is must not empty")
	}
	url := DeleteSiteOnceTasks
	var response SiteOnceBaseResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(request).
		WithResult(&response).
		Do()
	return &response, err
}

// 获取探测结果
func (c *Client) LoadData(request *SiteOnceTaskRequest) (*LoadDataResponse, error) {
	if request == nil {
		return nil, errors.New("SiteOnceTaskRequest is must not empty")
	}
	url := LoadSiteOnceData
	var response LoadDataResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(request).
		WithResult(&response).
		Do()
	return &response, err
}

//  获取探测详情
func (c *Client) DetailTask(request *SiteOnceTaskRequest) (*LoadDataResponse, error) {
	if request == nil {
		return nil, errors.New("SiteOnceTaskRequest is must not empty")
	}
	url := DetailSiteOnceTask
	var response LoadDataResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(request).
		WithResult(&response).
		Do()
	return &response, err
}

// 再次探测
func (c *Client) AgainExec(request *SiteOnceTaskRequest) (*SiteOnceBaseResponse, error) {
	if request == nil {
		return nil, errors.New("SiteOnceTaskRequest is must not empty")
	}
	url := AgainSiteOnce
	var response SiteOnceBaseResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(request).
		WithResult(&response).
		Do()
	return &response, err
}

// 历史列表
func (c *Client) ListHistoryTasks(request *SiteOnceTaskRequest) (*SiteOnceTaskListResponse, error) {
	if request == nil {
		return nil, errors.New("SiteOnceTaskRequest is must not empty")
	}
	url := ListSiteOnceHistory
	var response SiteOnceTaskListResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(request).
		WithResult(&response).
		Do()
	return &response, err
}

//获取当前探测点
func (c *Client) GetSiteAgent(userId, ipType string) (*SiteAgentResponseWrapper, error) {
	parmas := map[string]string{
		"userId": userId,
	}
	if ipType != "" {
		parmas["ipType"] = ipType
	}
	url := GetSiteAgent
	var response SiteAgentResponseWrapper
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithQueryParams(parmas).
		WithResult(&response).
		Do()
	return &response, err
}
