package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// ListTags lists all tags used across the user's clusters.
func ListTags(cli bce.Client, region string) (*TagListResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}

	uri := URI_PREFIX_V2 + "/tagList"

	result := &TagListResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, struct{}{}, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterTags updates the tags of a single cluster.
func UpdateClusterTags(cli bce.Client, region string, request *UpdateClusterTagsRequest) (*UpdateClusterTagsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster tags request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster tags request clusterId should not be empty")
	}
	if request.Tags == nil {
		return nil, errors.New("update cluster tags request tags should not be nil")
	}

	uri := URI_PREFIX_V2 + "/updateTags"

	result := &UpdateClusterTagsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// BatchInsertTags inserts the same set of tags into multiple clusters at once.
func BatchInsertTags(cli bce.Client, region string, request *BatchInsertTagsRequest) (*BatchInsertTagsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("batch insert tags request should not be nil")
	}
	if len(request.ClusterIdList) == 0 {
		return nil, errors.New("batch insert tags request clusterIdList should not be empty")
	}
	if len(request.InsertTags) == 0 {
		return nil, errors.New("batch insert tags request insertTags should not be empty")
	}

	uri := URI_PREFIX_V2 + "/batchInsertTags"

	result := &BatchInsertTagsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
