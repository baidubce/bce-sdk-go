package dataset

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/baidubce/bce-sdk-go/services/aihc"
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/api/v2"
)

const (
	ak_test       = ""
	sk_test       = ""
	endpoint_test = ""
)

var (
	DatasetID        string
	DatasetVersionID string
)

// 创建数据集
func CreateDataset() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)

	req := &v2.CreateDatasetRequest{
		Name:            "test-dataset-" + time.Now().Format("20060102150405"),
		StorageType:     "BOS",
		StorageInstance: "agilecloud-image",
		ImportFormat:    "FOLDER",
		Description:     "test dataset",
		VisibilityScope: "ONLY_OWNER",
		InitVersionEntry: v2.DatasetVersionEntry{
			Description: "test dataset version",
			StoragePath: "/",
			MountPath:   "/",
		},
	}

	result, err := client.CreateDataset(req)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 删除数据集
func DeleteDataset() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.DeleteDataset(DatasetID)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 修改数据集
func ModifyDataset() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.ModifyDataset(DatasetID, &v2.ModifyDatasetRequest{
		Name: "test-dataset-modify-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 查询数据集
func DescribeDataset() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.DescribeDataset(DatasetID)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 查询数据集列表
func DescribeDatasets() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.DescribeDatasets(&v2.DescribeDatasetsOptions{
		PageNumber: 1,
		PageSize:   10,
	})
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 创建数据集版本
func CreateDatasetVersion() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.CreateDatasetVersion(DatasetID, &v2.CreateDatasetVersionRequest{
		StoragePath: "/test-dataset-version",
		MountPath:   "/test-dataset-version",
		Description: "test dataset version",
	})
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 删除数据集版本
func DeleteDatasetVersion() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.DeleteDatasetVersion(DatasetID, DatasetVersionID)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 查询数据集版本
func DescribeDatasetVersion() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.DescribeDatasetVersion(DatasetID, DatasetVersionID)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 查询数据集版本列表
func DescribeDatasetVersions() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.DescribeDatasetVersions(DatasetID, &v2.DescribeDatasetVersionsOptions{
		PageNumber: 1,
		PageSize:   10,
	})
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

// 主函数
func main() {
	// fmt.Println("CreateDataset")
	// CreateDataset()

	fmt.Println("DescribeDatasets")
	DescribeDatasets()

	fmt.Println("DescribeDataset")
	DatasetID = "ds-CgyRs2EE"
	DescribeDataset()

	fmt.Println("DescribeDatasetVersions")
	DescribeDatasetVersions()

	// fmt.Println("ModifyDataset")
	// ModifyDataset()

	// DatasetVersionID = "ds-Iif8U49n"
	// fmt.Println("DescribeDatasetVersion")
	// DescribeDatasetVersion()

	// fmt.Println("DeleteDatasetVersion")
	// DeleteDatasetVersion()

	// fmt.Println("DeleteDataset")
	// DeleteDataset()

	fmt.Println("CreateDatasetVersion")
	CreateDatasetVersion()
}
