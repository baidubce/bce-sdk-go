package dataset

import (
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/api/v2"
)

// InterfaceV2 针对 v2 API 的数据集接口
// 方法签名参考 v2/dataset.go 的结构体定义
// 你可以根据实际实现进行调整
type Interface interface {
	// 数据集相关
	CreateDataset(req *v2.CreateDatasetRequest) (*v2.CreateDatasetResponse, error)
	DeleteDataset(datasetId string) (*v2.DeleteDatasetResponse, error)
	ModifyDataset(datasetId string, req *v2.ModifyDatasetRequest) (*v2.ModifyDatasetResponse, error)
	DescribeDataset(datasetId string) (*v2.DescribeDatasetResponse, error)
	DescribeDatasets(options *v2.DescribeDatasetsOptions) (*v2.DescribeDatasetsResponse, error)

	// 数据集版本相关
	CreateDatasetVersion(datasetId string, req *v2.CreateDatasetVersionRequest) (*v2.CreateDatasetVersionResponse, error)
	DeleteDatasetVersion(datasetId, versionId string) (*v2.DeleteDatasetVersionResponse, error)
	DescribeDatasetVersion(datasetId, versionId string) (*v2.DescribeDatasetVersionResponse, error)
	DescribeDatasetVersions(datasetId string, options *v2.DescribeDatasetVersionsOptions) (*v2.DescribeDatasetVersionsResponse, error)
}
