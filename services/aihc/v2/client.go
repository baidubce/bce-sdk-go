package v2

import (
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/v2/api"
	resourceQueue "github.com/baidubce/bce-sdk-go/services/aihc/v2/queue"
	resourcePool "github.com/baidubce/bce-sdk-go/services/aihc/v2/resourcepool"
)

type Interface interface {
	resourcePool.Interface
	resourceQueue.Interface
}

type Client struct {
	resourcePoolClient  *resourcePool.Client
	resourceQueueClient *resourceQueue.Client
}

func NewClient(ak, sk, endpoint string) (Interface, error) {
	clientset := &Client{}
	resourcePoolClient, err := resourcePool.NewClient(ak, sk, endpoint)
	if err != nil {
		return nil, err
	}
	clientset.resourcePoolClient = resourcePoolClient

	resourceQueueClient, err := resourceQueue.NewClient(ak, sk, endpoint)
	if err != nil {
		return nil, err
	}
	clientset.resourceQueueClient = resourceQueueClient

	return clientset, nil
}

func (c *Client) DescribeResourcePool(resourcePoolId string) (*v2.DescribeResourcePoolResponse, error) {
	return c.resourcePoolClient.DescribeResourcePool(resourcePoolId)
}

func (c *Client) DescribeResourcePools(describeResourcePoolsRequest v2.DescribeResourcePoolsRequest) (*v2.DescribeResourcePoolsResponse, error) {
	return c.resourcePoolClient.DescribeResourcePools(&describeResourcePoolsRequest)
}

func (c *Client) DescribeResourcePoolOverview() (*v2.DescribeResourcePoolOverviewResponse, error) {
	return c.resourcePoolClient.DescribeResourcePoolOverview()
}

func (c *Client) DescribeResourceQueue(queueID string) (*v2.DescribeResourceQueueResponse, error) {
	return c.resourceQueueClient.DescribeResourceQueue(queueID)
}

func (c *Client) DescribeResourceQueues(describeResourceQueuesRequest *v2.DescribeResourceQueuesRequest) (*v2.DescribeResourceQueuesResponse, error) {
	return c.resourceQueueClient.DescribeResourceQueues(describeResourceQueuesRequest)
}
