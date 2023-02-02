package eccr

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

type Interface interface {
	ListInstances(args *ListInstancesArgs) (*ListInstancesResponse, error)
	GetInstanceDetail(instanceID string) (*GetInstanceDetailResponse, error)
	ListPrivateNetworks(instanceID string) (*ListPrivateNetworksResponse, error)
}

// ListInstances - list all instance with the specific parameters
//
// PARAMS:
//   - ListInstancesArgs: the arguments to list all instance
//
// RETURNS:
//   - ListInstancesResponse: the result of list Instance
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *ListInstancesArgs) (*ListInstancesResponse, error) {

	result := &ListInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceListURI()).
		WithQueryParamFilter("keywordType", args.KeywordType).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithQueryParamFilter("acrossregion", args.Acrossregion).
		WithResult(result).
		Do()

	return result, err
}

// GetInstanceDetail - get a specific instance detail info
//
// PARAMS:
//   - instanceID: the specific instance ID
//
// RETURNS:
//   - *GetInstanceDetailResponse: the result of get instance detail info
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceID string) (*GetInstanceDetailResponse, error) {

	result := &GetInstanceDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceURI(instanceID)).
		WithResult(result).
		Do()

	return result, err
}

// ListPrivateNetworks - list all Privatelinks in an instance with the specific parameters
//
// PARAMS:
//   - instanceID: the specific instance ID
//
// RETURNS:
//   - *ListPrivateNetworksResponse: the result of list Privatelinks
//   - error: nil if success otherwise the specific error
func (c *Client) ListPrivateNetworks(instanceID string) (*ListPrivateNetworksResponse, error) {

	result := &ListPrivateNetworksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getPrivateNetworkListResponseURI(instanceID)).
		WithResult(result).
		Do()

	return result, err
}
