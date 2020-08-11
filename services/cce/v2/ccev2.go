package v2

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// 创建集群
func (c *Client) CreateCluster(args *CreateClusterArgs) (*CreateClusterResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &CreateClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterURI()).
		WithBody(args.CreateClusterRequest).
		WithResult(result).
		Do()

	return result, err
}

//删除集群
func (c *Client) DeleteCluster(args *DeleteClusterArgs) (*DeleteClusterResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &DeleteClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getClusterUriWithIDURI(args.ClusterID)).
		WithQueryParamFilter("deleteResource", strconv.FormatBool(args.DeleteResource)).
		WithQueryParamFilter("deleteCDSSnapshot", strconv.FormatBool(args.DeleteCDSSnapshot)).
		WithResult(result).
		Do()

	return result, err
}

//获得集群详情
func (c *Client) GetCluster(clusterID string) (*GetClusterResponse, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}

	result := &GetClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterUriWithIDURI(clusterID)).
		WithResult(result).
		Do()

	return result, err
}

//集群列表
func (c *Client) ListClusters(args *ListClustersArgs) (*ListClustersResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.PageNum <= 0 || args.PageSize <= 0 {
		return nil, fmt.Errorf("invlaid pageNo or pageSize")
	}

	result := &ListClustersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterListURI()).
		WithQueryParamFilter("keywordType", string(args.KeywordType)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNum)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()

	return result, err
}

//创建节点（扩容）
func (c *Client) CreateInstances(args *CreateInstancesArgs) (*CreateInstancesResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &CreateInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterInstanceListURI(args.ClusterID)).
		WithBody(args.Instances).
		WithResult(result).
		Do()

	return result, err
}

//查询节点
func (c *Client) GetInstance(args *GetInstanceArgs) (*GetInstanceResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &GetInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterInstanceURI(args.ClusterID, args.InstanceID)).
		WithResult(result).
		Do()

	return result, err
}

//删除节点（缩容）
func (c *Client) DeleteInstances(args *DeleteInstancesArgs) (*DeleteInstancesResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &DeleteInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClusterInstanceListURI(args.ClusterID)).
		WithBody(args.DeleteInstancesRequest).
		WithResult(result).
		Do()

	return result, err
}

//集群内节点列表
func (c *Client) ListInstancesByPage(args *ListInstancesByPageArgs) (*ListInstancesResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &ListInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterInstanceListURI(args.ClusterID)).
		WithQueryParamFilter("keywordType", string(args.Params.KeywordType)).
		WithQueryParamFilter("keyword", args.Params.Keyword).
		WithQueryParamFilter("orderBy", string(args.Params.OrderBy)).
		WithQueryParamFilter("order", string(args.Params.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.Params.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.Params.PageSize)).
		WithResult(result).
		Do()

	return result, err
}

//检查容器网络网段
func (c *Client) CheckContainerNetworkCIDR(args *CheckContainerNetworkCIDRArgs) (*CheckContainerNetworkCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("CheckContainerNetworkCIDRRequest is nil")
	}

	result := &CheckContainerNetworkCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetCheckContainerNetworkCIDRURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//检查集群网络网段
func (c *Client) CheckClusterIPCIDR(args *CheckClusterIPCIDRArgs) (*CheckClusterIPCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &CheckClusterIPCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetCheckClusterIPCIDRURL()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//推荐容器CIDR
func (c *Client) RecommendContainerCIDR(args *RecommendContainerCIDRArgs) (*RecommendContainerCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &RecommendContainerCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetRecommendContainerCidrURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//推荐集群CIDR
func (c *Client) RecommendClusterIPCIDR(args *RecommendClusterIPCIDRArgs) (*RecommendClusterIPCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &RecommendClusterIPCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetRecommendClusterIpCidrURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//用户集群 Quota
func (c *Client) GetClusterQuota() (*GetQuotaResponse, error) {
	result := &GetQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuotaURI()).
		WithResult(result).
		Do()

	return result, err
}

//用户集群 Node Quota
func (c *Client) GetClusterNodeQuota(clusterID string) (*GetQuotaResponse, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}

	result := &GetQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuotaNodeURI(clusterID)).
		WithResult(result).
		Do()

	return result, err
}