package queue

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/v2/api"
	"github.com/baidubce/bce-sdk-go/services/aihc/v2/client"
)

type Client struct {
	client.Client
}

// NewClient make the aihc inference service client with default configuration.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	aihcClient, err := client.NewClient(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	newClient := Client{*aihcClient}
	return &newClient, nil
}

func (c *Client) DescribeResourceQueue(queueID string) (*v2.DescribeResourceQueueResponse, error) {
	if queueID == "" {
		return nil, fmt.Errorf("queueID is empty")
	}

	result := &v2.DescribeResourceQueueResponse{}
	err := bce.NewRequestBuilder(c.DefaultClient).
		WithMethod(http.GET).
		WithURL("/").
		WithQueryParamFilter("action", string(v2.ResourceQueueActionDescribeResourceQueue)).
		WithQueryParamFilter("queueId", queueID).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DescribeResourceQueues(describeResourceQueuesRequest *v2.DescribeResourceQueuesRequest) (*v2.DescribeResourceQueuesResponse, error) {
	if describeResourceQueuesRequest == nil {
		return nil, fmt.Errorf("describe resource queues request is nil")
	}
	if describeResourceQueuesRequest.ResourcePoolID == "" {
		return nil, fmt.Errorf("describe resource queues resourcePoolID is empty")
	}

	if describeResourceQueuesRequest.PageNumber <= 0 {
		describeResourceQueuesRequest.PageNumber = 1
	}
	if describeResourceQueuesRequest.PageSize <= 0 {
		describeResourceQueuesRequest.PageSize = 10
	}

	result := &v2.DescribeResourceQueuesResponse{}
	err := bce.NewRequestBuilder(c.DefaultClient).
		WithMethod(http.GET).
		WithURL("/").
		WithQueryParamFilter("action", string(v2.ResourceQueueActionDescribeResourceQueues)).
		WithQueryParamFilter("resourcePoolId", string(describeResourceQueuesRequest.ResourcePoolID)).
		WithQueryParamFilter("keywordType", string(describeResourceQueuesRequest.KeywordType)).
		WithQueryParamFilter("keyword", string(describeResourceQueuesRequest.Keyword)).
		WithQueryParamFilter("PageNumber", strconv.Itoa(describeResourceQueuesRequest.PageNumber)).
		WithQueryParamFilter("PageSize", strconv.Itoa(describeResourceQueuesRequest.PageSize)).
		WithQueryParamFilter("orderBy", string(describeResourceQueuesRequest.OrderBy)).
		WithQueryParamFilter("order", string(describeResourceQueuesRequest.Order)).
		WithResult(result).
		Do()

	return result, err
}
