package resourcePool

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

func (c *Client) DescribeResourcePool(resourcePoolId string) (*v2.DescribeResourcePoolResponse, error) {
	if resourcePoolId == "" {
		return nil, fmt.Errorf("resource pool id is empty")
	}

	result := &v2.DescribeResourcePoolResponse{}
	err := bce.NewRequestBuilder(c.DefaultClient).
		WithMethod(http.GET).
		WithURL("/").
		WithQueryParamFilter("action", string(v2.ResourcePoolActionDescribeResourcePool)).
		WithQueryParamFilter("resourcePoolId", resourcePoolId).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DescribeResourcePools(describeResourcePoolsRequest *v2.DescribeResourcePoolsRequest) (*v2.DescribeResourcePoolsResponse, error) {
	if describeResourcePoolsRequest == nil {
		return nil, fmt.Errorf("describe resource pools request is nil")
	}

	if describeResourcePoolsRequest.ResourcePoolType == "" {
		return nil, fmt.Errorf("resource pool type is empty")
	}

	if describeResourcePoolsRequest.ResourcePoolType != v2.ResourcePoolTypeDedicatedV2 &&
		describeResourcePoolsRequest.ResourcePoolType != v2.ResourcePoolTypeCommon &&
		describeResourcePoolsRequest.ResourcePoolType != v2.ResourcePoolTypeBHCMP {
		return nil, fmt.Errorf("resource pool type: %+v is invalid", describeResourcePoolsRequest.ResourcePoolType)
	}

	if describeResourcePoolsRequest.PageNumber <= 0 {
		describeResourcePoolsRequest.PageNumber = 1
	}
	if describeResourcePoolsRequest.PageSize <= 0 {
		describeResourcePoolsRequest.PageSize = 10
	}

	result := &v2.DescribeResourcePoolsResponse{}
	err := bce.NewRequestBuilder(c.DefaultClient).
		WithMethod(http.GET).
		WithURL("/").
		WithQueryParamFilter("action", string(v2.ResourcePoolActionDescribeResourcePools)).
		WithQueryParamFilter("resourcePoolType", string(describeResourcePoolsRequest.ResourcePoolType)).
		WithQueryParamFilter("keywordType", string(describeResourcePoolsRequest.KeywordType)).
		WithQueryParamFilter("keyword", describeResourcePoolsRequest.Keyword).
		WithQueryParamFilter("orderBy", string(describeResourcePoolsRequest.OrderBy)).
		WithQueryParamFilter("order", string(describeResourcePoolsRequest.Order)).
		WithQueryParamFilter("pageNumber", strconv.Itoa(describeResourcePoolsRequest.PageNumber)).
		WithQueryParamFilter("pageSize", strconv.Itoa(describeResourcePoolsRequest.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DescribeResourcePoolOverview() (*v2.DescribeResourcePoolOverviewResponse, error) {
	result := &v2.DescribeResourcePoolOverviewResponse{}
	err := bce.NewRequestBuilder(c.DefaultClient).
		WithMethod(http.GET).
		WithURL("/").
		WithQueryParamFilter("action", string(v2.ResourcePoolActionDescribeResourcePoolOverview)).
		WithResult(result).
		Do()
	return result, err
}
