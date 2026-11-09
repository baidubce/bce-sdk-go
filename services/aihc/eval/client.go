package eval

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
)

// Client 具身评测 OpenAPI client
type Client struct {
	client.Client
}

// NewClient make the aihc embodied eval client with default configuration.
// endPoint 为空时使用默认的 aihc.bj.baidubce.com。
func NewClient(ak, sk, endPoint string) (*Client, error) {
	aihcClient, err := client.NewClient(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{*aihcClient}, nil
}

// NewClientWithSTS make the aihc embodied eval client with STS configuration.
func NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint string) (*Client, error) {
	aihcClient, err := client.NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{*aihcClient}, nil
}

func (c *Client) SetBceClient(client *bce.BceClient) {
	c.DefaultClient = client
}

func (c *Client) GetBceClient() *bce.BceClient {
	return c.DefaultClient
}

// CreateEmbodiedEval 创建评测任务
func (c *Client) CreateEmbodiedEval(args *CreateEmbodiedEvalArgs) (*CreateEmbodiedEvalResult, error) {
	return CreateEmbodiedEval(c.DefaultClient, args)
}

// DescribeEmbodiedEvals 查询评测任务列表
func (c *Client) DescribeEmbodiedEvals(args *DescribeEmbodiedEvalsArgs) (*DescribeEmbodiedEvalsResult, error) {
	return DescribeEmbodiedEvals(c.DefaultClient, args)
}

// DescribeEmbodiedEval 查询评测任务详情
func (c *Client) DescribeEmbodiedEval(evalId string) (*DescribeEmbodiedEvalDetailResult, error) {
	return DescribeEmbodiedEval(c.DefaultClient, evalId)
}

// StopEmbodiedEval 停止评测任务，已停止再停返回成功
func (c *Client) StopEmbodiedEval(args *StopEmbodiedEvalArgs) (*EmbodiedEvalIdResult, error) {
	return StopEmbodiedEval(c.DefaultClient, args)
}

// DeleteEmbodiedEval 删除评测任务，删了再删返回 404
func (c *Client) DeleteEmbodiedEval(args *DeleteEmbodiedEvalArgs) (*EmbodiedEvalIdResult, error) {
	return DeleteEmbodiedEval(c.DefaultClient, args)
}

// DescribeEmbodiedEvalResult 查询评测结果
func (c *Client) DescribeEmbodiedEvalResult(evalId string) (*DescribeEmbodiedEvalResultResult, error) {
	return DescribeEmbodiedEvalResult(c.DefaultClient, evalId)
}

// DescribeEmbodiedEvalVideo 查询回放视频地址
func (c *Client) DescribeEmbodiedEvalVideo(args *DescribeEmbodiedEvalVideoArgs) (*DescribeEmbodiedEvalVideoResult, error) {
	return DescribeEmbodiedEvalVideo(c.DefaultClient, args)
}

// DescribeEmbodiedEvalModels 查询可评测模型及其评测集
func (c *Client) DescribeEmbodiedEvalModels(args *DescribeEmbodiedEvalModelsArgs) (*DescribeEmbodiedEvalModelsResult, error) {
	return DescribeEmbodiedEvalModels(c.DefaultClient, args)
}

// DescribeEmbodiedEvalModel 查询单个模型元数据
func (c *Client) DescribeEmbodiedEvalModel(args *DescribeEmbodiedEvalModelArgs) (*DescribeEmbodiedEvalModelResult, error) {
	return DescribeEmbodiedEvalModel(c.DefaultClient, args)
}

// DescribeEmbodiedEvalBenchmarks 查询评测集元数据
func (c *Client) DescribeEmbodiedEvalBenchmarks(args *DescribeEmbodiedEvalBenchmarksArgs) (*DescribeEmbodiedEvalBenchmarksResult, error) {
	return DescribeEmbodiedEvalBenchmarks(c.DefaultClient, args)
}
