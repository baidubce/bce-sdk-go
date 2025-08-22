package resourcePool

import (
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/v2/api"
)

type Interface interface {
	DescribeResourcePool(resourcePoolId string) (*v2.DescribeResourcePoolResponse, error)
	DescribeResourcePools(describeResourcePoolsRequest v2.DescribeResourcePoolsRequest) (*v2.DescribeResourcePoolsResponse, error)
	DescribeResourcePoolOverview() (*v2.DescribeResourcePoolOverviewResponse, error)
}
