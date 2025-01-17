package resource

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
)

const (
	URI_PREFIX                = bce.URI_PREFIX + "api" + bce.URI_PREFIX + "v1"
	REQUEST_RESOURCE_POOL_URL = "/resourcepools"
	REQUEST_NODE_URL          = "/nodes"
	REQUEST_QUEUE_URL         = "/queue"
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

func (c *Client) GetBceClient() *bce.BceClient {
	return c.DefaultClient
}

func (c *Client) SetBceClient(client *bce.BceClient) {
	c.DefaultClient = client
}

// NewClientWithSTS make the aihc inference service client with STS configuration.
func NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint string) (*Client, error) {
	aihcClient, err := client.NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint)
	if err != nil {
		return nil, err
	}
	newClient := Client{*aihcClient}
	return &newClient, nil
}

func (c *Client) GetResourcePool(resourcePoolID string) (result *v1.GetResourcePoolResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	result = &v1.GetResourcePoolResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getResourcePoolUriWithID(resourcePoolID)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListResourcePool(args *v1.ListResourcePoolRequest) (result *v1.ListResourcePoolResponse, err error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.PageNo <= 0 || args.PageSize <= 0 {
		return nil, fmt.Errorf("invlaid pageNo or pageSize")
	}

	result = &v1.ListResourcePoolResponse{}
	err = bce.NewRequestBuilder(c.DefaultClient).
		WithMethod(http.GET).
		WithURL(listResourcePoolUri()).
		WithQueryParamFilter("keywordType", string(args.KeywordType)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) ListNodeByResourcePoolID(resourcePoolID string, args *v1.ListResourcePoolNodeRequest) (result *v1.ListNodeByResourcePoolResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	if args != nil && (args.PageNo <= 0 || args.PageSize <= 0) {
		return nil, fmt.Errorf("invlaid pageNo or pageSize")
	}

	result = &v1.ListNodeByResourcePoolResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(listResourcePoolNodesUri(resourcePoolID)).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetQueue(resourcePoolID, queueName string) (result *v1.GetQueuesResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	if queueName == "" {
		return nil, fmt.Errorf("queueName is empty")
	}
	result = &v1.GetQueuesResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getQueueUri(resourcePoolID, queueName)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListQueue(resourcePoolID string, args *v1.ListQueueRequest) (result *v1.ListQueuesResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	if args != nil && (args.PageNo <= 0 || args.PageSize <= 0) {
		return nil, fmt.Errorf("invlaid pageNo or pageSize")
	}
	result = &v1.ListQueuesResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(listQueueUri(resourcePoolID)).
		WithQueryParamFilter("keywordType", string(args.KeywordType)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

func listResourcePoolUri() string {
	return URI_PREFIX + REQUEST_RESOURCE_POOL_URL
}

func getResourcePoolUriWithID(resourcePoolID string) string {
	return listResourcePoolUri() + "/" + resourcePoolID
}

func listResourcePoolNodesUri(resourcePoolID string) string {
	return getResourcePoolUriWithID(resourcePoolID) + REQUEST_NODE_URL
}

func listQueueUri(resourcePoolID string) string {
	return getResourcePoolUriWithID(resourcePoolID) + REQUEST_QUEUE_URL
}

func getQueueUri(resourcePoolID, queueName string) string {
	return getResourcePoolUriWithID(resourcePoolID) + REQUEST_QUEUE_URL + "/" + queueName
}
