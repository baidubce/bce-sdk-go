package dataset

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/api/v2"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
)

const (
	CreateDatasetAction           = "CreateDataset"
	DeleteDatasetAction           = "DeleteDataset"
	ModifyDatasetAction           = "ModifyDataset"
	DescribeDatasetAction         = "DescribeDataset"
	DescribeDatasetsAction        = "DescribeDatasets"
	CreateDatasetVersionAction    = "CreateDatasetVersion"
	DeleteDatasetVersionAction    = "DeleteDatasetVersion"
	DescribeDatasetVersionAction  = "DescribeDatasetVersion"
	DescribeDatasetVersionsAction = "DescribeDatasetVersions"

	VERSION = "v2"
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

	newClient := Client{
		*aihcClient,
	}
	return &newClient, nil
}

// CreateDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) CreateDataset(req *v2.CreateDatasetRequest) (*v2.CreateDatasetResponse, error) {
	result := &v2.CreateDatasetResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.POST).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", CreateDatasetAction).
		WithBody(req).
		WithResult(result).
		Do()
	return result, err
}

// DeleteDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) DeleteDataset(datasetId string) (*v2.DeleteDatasetResponse, error) {
	result := &v2.DeleteDatasetResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.POST).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", DeleteDatasetAction).
		WithQueryParamFilter("datasetId", datasetId).
		WithResult(result).
		Do()
	return result, err
}

// ModifyDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) ModifyDataset(datasetId string, req *v2.ModifyDatasetRequest) (*v2.ModifyDatasetResponse, error) {
	result := &v2.ModifyDatasetResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.POST).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", ModifyDatasetAction).
		WithQueryParamFilter("datasetId", datasetId).
		WithBody(req).
		WithResult(result).
		Do()
	return result, err
}

// DescribeDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) DescribeDataset(datasetId string) (*v2.DescribeDatasetResponse, error) {
	result := &v2.DescribeDatasetResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", DescribeDatasetAction).
		WithQueryParamFilter("datasetId", datasetId).
		WithResult(result).
		Do()
	return result, err
}

// DescribeDatasets
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap

func (c *Client) DescribeDatasets(options *v2.DescribeDatasetsOptions) (*v2.DescribeDatasetsResponse, error) {
	result := &v2.DescribeDatasetsResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", DescribeDatasetsAction).
		WithQueryParamFilter("keyword", options.Keyword).
		WithQueryParamFilter("storageType", options.StorageType).
		WithQueryParamFilter("storageInstances", options.StorageInstances).
		WithQueryParamFilter("importFormat", options.ImportFormat).
		WithQueryParamFilter("pageNumber", strconv.Itoa(options.PageNumber)).
		WithQueryParamFilter("pageSize", strconv.Itoa(options.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

// CreateDatasetVersion
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) CreateDatasetVersion(datasetId string, req *v2.CreateDatasetVersionRequest) (*v2.CreateDatasetVersionResponse, error) {
	result := &v2.CreateDatasetVersionResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.POST).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", CreateDatasetVersionAction).
		WithQueryParamFilter("datasetId", datasetId).
		WithBody(req).
		WithResult(result).
		Do()
	return result, err
}

// DeleteDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) DeleteDatasetVersion(datasetId, versionId string) (*v2.DeleteDatasetVersionResponse, error) {
	result := &v2.DeleteDatasetVersionResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.POST).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", DeleteDatasetVersionAction).
		WithQueryParamFilter("datasetId", datasetId).
		WithQueryParamFilter("versionId", versionId).
		WithResult(result).
		Do()
	return result, err
}

// DescribeDatasetVersion
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) DescribeDatasetVersion(datasetId, versionId string) (*v2.DescribeDatasetVersionResponse, error) {
	result := &v2.DescribeDatasetVersionResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", DescribeDatasetVersionAction).
		WithQueryParamFilter("datasetId", datasetId).
		WithQueryParamFilter("versionId", versionId).
		WithResult(result).
		Do()
	return result, err
}

// DescribeDatasetVersions
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
func (c *Client) DescribeDatasetVersions(datasetId string, options *v2.DescribeDatasetVersionsOptions) (*v2.DescribeDatasetVersionsResponse, error) {
	result := &v2.DescribeDatasetVersionsResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL("/").
		WithHeader("version", VERSION).
		WithQueryParamFilter("action", DescribeDatasetVersionsAction).
		WithQueryParamFilter("datasetId", datasetId).
		WithQueryParamFilter("pageNumber", strconv.Itoa(options.PageNumber)).
		WithQueryParamFilter("pageSize", strconv.Itoa(options.PageSize)).
		WithResult(result).
		Do()
	return result, err
}
