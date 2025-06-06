package dev

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
)

type Client struct {
	client.Client
}

// NewClient make the aihc dev client with default configuration.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		return nil, fmt.Errorf("endpoint can not be empty")
	}
	aihcClient, err := client.NewClient(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	newClient := Client{*aihcClient}
	return &newClient, nil
}

func (c *Client) SetBceClient(client *bce.BceClient) {
	c.DefaultClient = client
}

func (c *Client) GetBceClient() *bce.BceClient {
	return c.DefaultClient
}

// NewClientWithSTS make the aihc dev client with STS configuration.
func NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		return nil, fmt.Errorf("endpoint can not be empty")
	}
	aihcClient, err := client.NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint)
	if err != nil {
		return nil, err
	}
	newClient := Client{*aihcClient}
	return &newClient, nil
}

// CreateDevInstance 创建开发机实例
func (c *Client) CreateDevInstance(args *CreateDevInstanceArgs) (*CreateDevInstanceResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return CreateDevInstance(c.DefaultClient, body)
}

// ListDevInstance 查询开发机列表
func (c *Client) ListDevInstance(args *ListDevInstanceArgs) (*ListDevInstanceResult, error) {
	return ListDevInstance(c.DefaultClient, args)
}

// QueryDevInstanceDetail 查询开发机详情
func (c *Client) QueryDevInstanceDetail(args *QueryDevInstanceDetailArgs) (*QueryDevInstanceDetailResult, error) {
	return QueryDevInstanceDetail(c.DefaultClient, args)
}

// UpdateDevInstance 更新开发机配置
func (c *Client) UpdateDevInstance(args *CreateDevInstanceArgs) (*CreateDevInstanceResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return UpdateDevInstance(c.DefaultClient, body)
}

// StartDevInstance 启动开发机实例
func (c *Client) StartDevInstance(args *StartDevInstanceArgs) (*StartDevInstanceResult, error) {
	return StartDevInstance(c.DefaultClient, args)
}

// StopDevInstance 停止开发机实例
func (c *Client) StopDevInstance(args *StopDevInstanceArgs) (*StopDevInstanceResult, error) {
	return StopDevInstance(c.DefaultClient, args)
}

// DeleteDevInstance 删除开发机
func (c *Client) DeleteDevInstance(args *DeleteDevInstanceArgs) (*DeleteDevInstanceResult, error) {
	return DeleteDevInstance(c.DefaultClient, args)
}

// TimedStopDevInstance 定时停止开发机
func (c *Client) TimedStopDevInstance(args *TimedStopDevInstanceArgs) (*TimedStopDevInstanceResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return TimedStopDevInstance(c.DefaultClient, body)
}

// ListDevInstanceEvent 查询开发机事件
func (c *Client) ListDevInstanceEvent(args *ListDevInstanceEventArgs) (*ListDevInstanceEventResult, error) {
	return ListDevInstanceEvent(c.DefaultClient, args)
}

// CreateDevInstanceImagePackJob 创建镜像任务
func (c *Client) CreateDevInstanceImagePackJob(args *CreateDevInstanceImagePackJobArgs) (*CreateDevInstanceImagePackJobResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return CreateDevInstanceImagePackJob(c.DefaultClient, body)
}

// DevInstanceImagePackJobDetail 查询镜像任务详情
func (c *Client) DevInstanceImagePackJobDetail(args *DevInstanceImagePackJobDetailArgs) (*DevInstanceImagePackJobDetailResult, error) {
	return DevInstanceImagePackJobDetail(c.DefaultClient, args)
}
