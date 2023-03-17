# DBSC服务

# 概述

本文档主要介绍DBSC GO SDK的使用。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BCC访问域名](https://cloud.baidu.com/doc/BCC/s/0jwvyo603)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云DBSC，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问DBSC做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建DBSC Client

DBSC Client是BCC服务的客户端，为开发者与DBSC服务进行交互提供了一系列的方法。

### 使用AK/SK新建DBSC Client

通过AK/SK方式使用DBSC，用户可以参考如下代码新建一个dbsc Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/dbsc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BCCClient
	dbscClient, err := dbsc.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/BCC/s/ojwvynrqn)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为dbsc的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`bcc.bj.baidubce.com`。


# 主要接口


### 创建磁盘专属集群
以下代码可以根据实例ID批量查询实例列表
```go
args := &CreateVolumeClusterArgs{
    // 创建一个磁盘磁盘专属集群，若要同时创建多个，可以修改此参数
	PurchaseCount:   1,
	// 集群大小,支持最小容量:85TB（87040GB）,支持最大容量：1015TB（1039360GB）,购买步长：10TB
	ClusterSizeInGB: 97280,
    // 集群名称
	ClusterName:     "dbsc",
	// 集群磁盘类型：通用型HDD，通用型SSD
	StorageType:     StorageTypeHdd, 
	Billing: &Billing{
        // 只支持预付费
		Reservation: &Reservation{
            // 购买时长
			ReservationLength:   6,
			ReservationTimeUnit: "MONTH",
		},
	},
    // 自动续费时长
	RenewTimeUnit: "MONTH",
	RenewTime:     6,
}
result, err := DBSC_CLIENT.CreateVolumeCluster(args)
if err != nil {
	fmt.Println(err)
}
clusterId := result.ClusterIds[0]
fmt.Print(clusterId)
```

### 磁盘专属集群列表
以下代码可以根据实例ID批量查询实例列表
```go
args := &ListVolumeClusterArgs{
}
result, err := DBSC_CLIENT.ListVolumeCluster(args)
if err != nil {
	fmt.Println(err)
}
fmt.Println(result)
```

### 磁盘专属集群详情
以下代码可以根据实例ID批量查询实例列表
```go
clusterId := "clusterId"
result, err := DBSC_CLIENT.GetVolumeClusterDetail(clusterId)
if err != nil {
	fmt.Println(err)
}
fmt.Println(result)
```

### 磁盘专属集群扩容
以下代码可以根据实例ID批量查询实例列表
```go
clusterId := "clusterId"
args := &ResizeVolumeClusterArgs{
	NewClusterSizeInGB int  `json:"newClusterSizeInGB"`
}
err := DBSC_CLIENT.ResizeVolumeCluster(clusterId, args)
if err != nil {
	fmt.Println(err)
}
```

### 磁盘专属集群续费
以下代码可以根据实例ID批量查询实例列表
```go
args := &PurchaseReservedVolumeClusterArgs{
	Billing: &Billing{
		Reservation: &Reservation{
            // 续费时长
			ReservationLength:   6,
			ReservationTimeUnit: "MONTH",
		},
	},
}
clusterId := "clusterId"
err := DBSC_CLIENT.PurchaseReservedVolumeCluster(clusterId, args)
if err != nil {
	fmt.Println(err)
}
```

### 磁盘专属集群自动续费
以下代码可以根据实例ID批量查询实例列表
```go
clusterId := "clusterId"
args := &AutoRenewVolumeClusterArgs{
	ClusterId:     clusterId,
	RenewTime:     6,
	RenewTimeUnit: "month",
}
err := DBSC_CLIENT.AutoRenewVolumeCluster(args)
if err != nil {
	fmt.Println(err)
}
```

### 磁盘专属集群取消自动续费
以下代码可以根据实例ID批量查询实例列表
```go
clusterId := "clusterId"
args := &CancelAutoRenewVolumeClusterArgs{
	ClusterId: clusterId,
}
err := DBSC_CLIENT.CancelAutoRenewVolumeCluster(args)
if err != nil {
	fmt.Println(err)
}
```

首次发布：

 - 创建磁盘专属集群、磁盘专属集群列表、磁盘专属集群详情、磁盘专属集群扩容、磁盘专属集群续费、磁盘专属集群自动续费、磁盘专属集群取消自动续费
