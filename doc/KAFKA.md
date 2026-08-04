# Kafka 服务

## 概述

Kafka Go SDK 封装了百度智能云 Kafka 服务的集群、集群配置、主题、消费组、用户、ACL、任务和 Quota 等接口。

使用 SDK 前，需要开通 Kafka 服务、准备有效的访问凭证，并确认资源所在区域和服务 Endpoint。SDK 包路径如下：

```go
import "github.com/baidubce/bce-sdk-go/services/kafka"
```

## 初始化客户端

### Endpoint

默认 Endpoint 为 `kafka.bj.baidubce.com`。创建客户端时：

- 显式传入 Endpoint 时，SDK 使用该地址。
- Endpoint 为空且指定了 `Region` 时，SDK 使用 `kafka.<region>.baidubce.com`。
- Endpoint 和 `Region` 均为空时，SDK 使用默认区域和默认 Endpoint。

生产环境应使用资源实际所在区域的 Endpoint。

### 使用 AK/SK

```go
package main

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/kafka"
)

func main() {
	client, err := kafka.NewClient(
		"<your-access-key-id>",
		"<your-secret-access-key>",
		"kafka.bj.baidubce.com",
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.ListClusters(&kafka.ListClustersRequest{
		ListRequest: kafka.ListRequest{MaxKeys: 100},
	})
	if err != nil {
		panic(err)
	}

	for _, cluster := range resp.Clusters {
		fmt.Printf("clusterId=%s name=%s state=%s\n", cluster.ClusterID, cluster.Name, cluster.State)
	}
}
```

不要在代码仓库、日志或命令行历史中保存真实 AK/SK。建议通过环境变量或密钥管理服务注入凭证。

### 使用 STS 凭证

```go
client, err := kafka.NewClientWithSTS(
	"<temporary-access-key-id>",
	"<temporary-secret-access-key>",
	"<session-token>",
	"kafka.bj.baidubce.com",
)
if err != nil {
	panic(err)
}
```

STS 凭证过期后，需要使用新凭证重新创建客户端。

### 自定义客户端配置

`KafkaClientConfiguration` 提供 Endpoint、区域、凭证、重试、代理、网络超时、自定义 HTTP Client 和限速等配置。

```go
package main

import (
	"time"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/kafka"
)

func main() {
	credentials, err := auth.NewBceCredentials(
		"<your-access-key-id>",
		"<your-secret-access-key>",
	)
	if err != nil {
		panic(err)
	}

	dialTimeout := 10 * time.Second
	readTimeout := 30 * time.Second
	httpClientTimeout := 60 * time.Second

	client, err := kafka.NewClientWithConfig(&kafka.KafkaClientConfiguration{
		Region:                    "gz",
		Credentials:               credentials,
		ConnectionTimeoutInMillis: 10 * 1000,
		DialTimeout:               &dialTimeout,
		ReadTimeout:               &readTimeout,
		HTTPClientTimeout:         &httpClientTimeout,
		Retry:                     bce.NewNoRetryPolicy(),
	})
	if err != nil {
		panic(err)
	}
	_ = client
}
```

未显式配置时，Kafka Client 使用以下值：

| 配置项 | 默认值 |
| --- | --- |
| `Region` | `bj` |
| `Endpoint` | `kafka.<region>.baidubce.com` |
| `ConnectionTimeoutInMillis` | 50000 毫秒 |
| `DialTimeout` | 与连接超时一致 |
| `ReadTimeout` | 50 秒 |
| `Retry` | SDK 默认退避重试策略 |

常用配置字段：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `Endpoint` | `string` | Kafka 服务地址，优先级高于 `Region` |
| `Region` | `string` | Endpoint 为空时用于生成服务地址 |
| `Credentials` | `*auth.BceCredentials` | AK/SK 或 STS 凭证 |
| `Retry` | `bce.RetryPolicy` | 请求重试策略 |
| `ProxyUrl` | `string` | HTTP 代理地址 |
| `ConnectionTimeoutInMillis` | `int` | 连接超时，单位为毫秒 |
| `DialTimeout` | `*time.Duration` | 建立连接的超时 |
| `ReadTimeout` | `*time.Duration` | 读取连接的超时 |
| `HTTPClientTimeout` | `*time.Duration` | 整个 HTTP 请求的超时 |
| `HTTPClient` | `*http.Client` | 自定义 HTTP Client |
| `UploadRatelimit` | `*int64` | 上传限速，单位为 KB/s |
| `DownloadRatelimit` | `*int64` | 下载限速，单位为 KB/s |

`NewClientWithConfig(nil)` 会返回错误。客户端创建后可以复用，不需要为每个请求重复创建。

## 请求模型约定

### 必填字段

每个方法都会在发送请求前校验必要参数。常见必填字段包括 `ClusterID`、`TopicName`、`GroupName`、`Username`、`ActionID` 和 `ConfigID`。请求对象不能为 `nil`。

### 指针字段

布尔值、数字和部分字符串使用指针表示“未设置”和“显式设置为零值”的区别。例如：

```go
enabled := false
partitionID := 0
emptyKey := ""
value := "message"

accessRequest := &kafka.SwitchClusterAdvertisedIpRequest{
	ClusterID:           "<cluster-id>",
	AdvertisedIPEnabled: &enabled,
}

messageRequest := &kafka.SendTopicMessageRequest{
	ClusterID:   "<cluster-id>",
	TopicName:   "<topic-name>",
	PartitionID: &partitionID,
	Key:         &emptyKey,
	Value:       &value,
}

_ = accessRequest
_ = messageRequest
```

`nil` 指针通常表示不传递该字段，非 `nil` 指针会保留 `false`、`0` 或空字符串等显式值。`SendTopicMessage` 的 `Value` 必须是非空字符串。

### 集合字段

请求中的切片和 map 可以区分：

- `nil`：不发送该字段。
- 显式空切片或空 map：发送空数组或空对象。

只有服务端支持清空的字段才应传递显式空集合。

### 默认模型

以下构造函数会初始化模型默认值：

- `NewBilling()`
- `NewStorageMeta()`
- `NewConfigMeta()`

创建复杂请求时，建议优先使用这些构造函数，再填写业务字段。

## 分页

### Marker 分页

`ListClusters`、`GetClusterNodes`、`ListClusterConfigs` 和 `ListJobs` 使用 `Marker`/`MaxKeys` 分页。`MaxKeys` 的有效范围为 1 到 1000，超出该范围时 SDK 不会发送该参数。

```go
marker := ""
for {
	resp, err := client.ListClusters(&kafka.ListClustersRequest{
		ListRequest: kafka.ListRequest{
			Marker:  marker,
			MaxKeys: 100,
		},
	})
	if err != nil {
		panic(err)
	}

	for _, cluster := range resp.Clusters {
		fmt.Println(cluster.ClusterID, cluster.Name)
	}

	if !resp.IsTruncated || resp.NextMarker == "" {
		break
	}
	marker = resp.NextMarker
}
```

### 页码分页

`ListTopicPartitions` 使用 `PageNo`/`PageSize` 分页。未设置时，默认页码为 1，默认每页 10 条。

```go
pageNo := 1
pageSize := 20

resp, err := client.ListTopicPartitions(&kafka.ListTopicPartitionsRequest{
	PageListRequest: kafka.PageListRequest{
		PageNo:   &pageNo,
		PageSize: &pageSize,
	},
	ClusterID: "<cluster-id>",
	TopicName: "<topic-name>",
})
```

## 使用示例

### 查询集群详情

```go
resp, err := client.GetClusterDetail(&kafka.GetClusterDetailRequest{
	ClusterID: "<cluster-id>",
})
if err != nil {
	panic(err)
}
if resp.Cluster != nil {
	fmt.Printf("name=%s state=%s\n", resp.Cluster.Name, resp.Cluster.State)
}
```

### 创建主题

```go
resp, err := client.CreateTopic(&kafka.CreateTopicRequest{
	ClusterID:         "<cluster-id>",
	TopicName:         "<topic-name>",
	PartitionNum:      3,
	ReplicationFactor: 3,
	OtherConfigs: map[string]string{
		"cleanup.policy": "delete",
	},
})
if err != nil {
	panic(err)
}
fmt.Println(resp.TopicName)
```

### 更新主题

`UpdateTopic` 至少需要设置 `PartitionNum` 或 `OtherConfigs` 中的一项。

```go
resp, err := client.UpdateTopic(&kafka.UpdateTopicRequest{
	ClusterID:   "<cluster-id>",
	TopicName:   "<topic-name>",
	PartitionNum: "6",
})
if err != nil {
	panic(err)
}
fmt.Println(resp.TopicName)
```

### 发送和查询消息

```go
partitionID := 0
key := "order-1001"
value := `{"status":"created"}`

sent, err := client.SendTopicMessage(&kafka.SendTopicMessageRequest{
	ClusterID:   "<cluster-id>",
	TopicName:   "<topic-name>",
	PartitionID: &partitionID,
	Key:         &key,
	Value:       &value,
})
if err != nil {
	panic(err)
}
if sent.Message == nil {
	panic("empty message response")
}

messages, err := client.QueryTopicMessagesByStartOffset(
	&kafka.QueryTopicMessagesByStartOffsetRequest{
		ClusterID:   "<cluster-id>",
		TopicName:   "<topic-name>",
		PartitionID: partitionID,
		StartOffset: sent.Message.Offset,
	},
)
if err != nil {
	panic(err)
}
fmt.Println(len(messages.Messages))
```

调用 `QueryTopicMessagesByStartOffset` 时，`PartitionID` 和 `StartOffset` 不能为负数。调用 `QueryTopicMessagesByStartTime` 时，`StartTime` 必须为正数，单位按服务接口定义传递。

### 重置消费组位点

```go
resp, err := client.ResetConsumerGroup(&kafka.ResetConsumerGroupRequest{
	ClusterID:     "<cluster-id>",
	GroupName:     "<group-name>",
	TopicName:     "<topic-name>",
	Partitions:    []int{0, 1, 2},
	ResetStrategy: "EARLIEST",
})
if err != nil {
	panic(err)
}
fmt.Println(resp.GroupName)
```

`Partitions` 和 `ResetStrategy` 必须设置。重置位点会影响消费进度，执行前应确认目标消费组、主题和分区。

### 管理用户

创建用户和重置密码时，SDK 使用客户端凭证中的 SK 对密码进行 AES-128 加密，然后发送加密后的值。调用方仍传入原始密码，SDK 不会修改传入的请求对象。

```go
created, err := client.CreateUser(&kafka.CreateUserRequest{
	ClusterID: "<cluster-id>",
	Username:  "<username>",
	Password:  "<password>",
	SASLMechanisms: []string{
		"SCRAM-SHA-256",
	},
})
if err != nil {
	panic(err)
}
fmt.Println(created.Username)
```

不要记录原始密码、加密后的密码或完整请求体。

查询用户可以直接传入集群 ID，也可以使用请求对象：

```go
users, err := client.ListUsers("<cluster-id>")

users, err = client.ListUsers(&kafka.ListUsersRequest{
	ClusterID: "<cluster-id>",
})
```

### 管理 ACL

```go
created, err := client.CreateAcl(&kafka.CreateAclRequest{
	ClusterID:    "<cluster-id>",
	Username:     "<username>",
	PatternType:  "LITERAL",
	ResourceType: "TOPIC",
	ResourceName: "<topic-name>",
	Operations:   []string{"READ", "DESCRIBE"},
})
if err != nil {
	panic(err)
}
fmt.Println(created.Username)
```

查询 ACL 同样支持直接传入集群 ID 或使用过滤请求：

```go
acls, err := client.ListAcls("<cluster-id>")

acls, err = client.ListAcls(&kafka.ListAclRequest{
	ClusterID:    "<cluster-id>",
	Username:     "<username>",
	PatternType:  "LITERAL",
	ResourceType: "TOPIC",
	ResourceName: "<topic-name>",
})
```

`ListUsers` 和 `ListAcls` 接受字符串、请求结构体值或请求结构体指针。传入其他类型会返回参数错误。

### 管理 Quota

```go
producerRate := int64(1024 * 1024)
consumerRate := int64(2 * 1024 * 1024)

resp, err := client.CreateQuota(&kafka.CreateQuotaRequest{
	ClusterID:        "<cluster-id>",
	Username:         "<username>",
	ProducerByteRate: &producerRate,
	ConsumerByteRate: &consumerRate,
})
if err != nil {
	panic(err)
}
fmt.Printf("%+v\n", resp.Quota)
```

## 错误处理

SDK 方法返回的错误主要包括：

- 请求参数错误：请求发送前由客户端校验产生。
- `*bce.BceClientError`：构建请求、网络连接或响应解析等客户端错误。
- `*bce.BceServiceError`：服务端返回的错误，包含错误码、消息、Request ID 和 HTTP 状态码。

```go
import (
	"errors"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

resp, err := client.GetClusterDetail(&kafka.GetClusterDetailRequest{
	ClusterID: "<cluster-id>",
})
if err != nil {
	var serviceErr *bce.BceServiceError
	var clientErr *bce.BceClientError

	switch {
	case errors.As(err, &serviceErr):
		fmt.Printf("code=%s status=%d requestId=%s message=%s\n",
			serviceErr.Code,
			serviceErr.StatusCode,
			serviceErr.RequestId,
			serviceErr.Message,
		)
	case errors.As(err, &clientErr):
		fmt.Printf("client error: %s\n", clientErr.Message)
	default:
		fmt.Printf("request error: %v\n", err)
	}
	return
}
_ = resp
```

排查服务端错误时，应保留 `RequestId`，但不要在日志中输出 AK/SK、用户密码或消息中的敏感数据。

## 接口列表

### 集群接口

| 方法 | 说明 |
| --- | --- |
| `CreateCluster` | 创建集群 |
| `DeleteCluster` | 删除集群 |
| `ListClusters` | 查询集群列表 |
| `GetClusterDetail` | 查询集群详情 |
| `GetClusterDeletion` | 查询集群删除信息 |
| `GetClusterAccessEndpoints` | 查询集群访问地址 |
| `GetClusterNodes` | 查询集群节点 |
| `GetClusterConfigurations` | 查询集群当前配置项 |
| `IncreaseBrokerCount` | 增加 Broker 数量 |
| `DecreaseBrokerCount` | 减少 Broker 数量 |
| `MigrateClusterAz` | 迁移集群可用区 |
| `UnifyClusterEndpoint` | 统一集群访问地址 |
| `UpdateBrokerNodeType` | 修改 Broker 节点规格 |
| `ExpandBrokerDiskCapacity` | 扩容 Broker 磁盘 |
| `UpdateAccessConfig` | 修改访问和认证配置 |
| `StartCluster` | 启动集群 |
| `StopCluster` | 停止集群 |
| `ResizeClusterEipBandwidth` | 修改公网带宽 |
| `SwitchClusterEip` | 开启或关闭公网访问 |
| `UpdateStoragePolicy` | 修改存储策略 |
| `UpdateKafkaConfig` | 应用集群配置版本 |
| `UpdateSecurityGroup` | 修改安全组 |
| `UpdateMaintenanceDuration` | 修改维护时间 |
| `SwitchClusterIntranetIp` | 开启或关闭内网 IP 访问 |
| `GetClusterCurrentController` | 查询当前 Controller |
| `GetClusterHistoryController` | 查询历史 Controller |
| `RestartCluster` | 重启集群 |
| `RestartBroker` | 重启指定 Broker |
| `SwitchClusterAdvertisedIp` | 开启或关闭 Advertised IP |
| `SwitchClusterDomain` | 开启或关闭域名访问 |
| `GetZkPassword` | 查询 ZooKeeper 访问凭证 |

### 集群配置接口

| 方法 | 说明 |
| --- | --- |
| `CreateClusterConfig` | 创建集群配置 |
| `ListClusterConfigs` | 查询集群配置列表 |
| `DeleteClusterConfig` | 删除集群配置 |
| `GetClusterConfig` | 查询集群配置详情 |
| `CreateClusterConfigRevision` | 创建配置版本 |
| `ListClusterConfigRevisions` | 查询配置版本列表 |
| `GetClusterConfigRevision` | 查询配置版本详情 |

### 主题接口

| 方法 | 说明 |
| --- | --- |
| `CreateTopic` | 创建主题 |
| `DeleteTopic` | 删除主题 |
| `ListTopic` | 查询主题列表 |
| `GetTopicDetail` | 查询主题详情 |
| `UpdateTopic` | 修改主题分区数或配置 |
| `ListTopicPartitions` | 查询主题分区状态 |
| `GetTopicPartitionDetail` | 查询分区详情 |
| `GetTopicPartitionOverview` | 查询分区概览 |
| `ListSubscribedGroups` | 查询订阅主题的消费组 |
| `GetSubscribedGroupDetail` | 查询消费组订阅详情 |
| `GetSubscribedGroupOverview` | 查询消费组订阅概览 |
| `ListTopicConfigOptions` | 查询主题支持的配置项 |
| `SendTopicMessage` | 向主题发送消息 |
| `QueryTopicMessagesByStartTime` | 按起始时间查询消息 |
| `QueryTopicMessagesByStartOffset` | 按起始 Offset 查询消息 |

### 消费组接口

| 方法 | 说明 |
| --- | --- |
| `ListConsumerGroup` | 查询消费组列表 |
| `DeleteConsumerGroup` | 删除消费组 |
| `ResetConsumerGroup` | 重置消费组位点 |
| `ListSubscribedTopics` | 查询消费组订阅的主题 |
| `GetSubscribedTopicOverview` | 查询消费组订阅主题概览 |

### 用户与 ACL 接口

| 方法 | 说明 |
| --- | --- |
| `CreateUser` | 创建用户 |
| `DeleteUser` | 删除用户 |
| `ResetUserPassword` | 重置用户密码 |
| `ListUsers` | 查询用户列表 |
| `CreateAcl` | 创建 ACL |
| `DeleteAcl` | 删除 ACL |
| `ListAcls` | 查询 ACL 列表 |

### 任务接口

| 方法 | 说明 |
| --- | --- |
| `ListJobs` | 查询任务列表 |
| `GetJob` | 查询任务详情 |
| `GetOperation` | 查询任务操作详情 |
| `StartJob` | 启动任务 |
| `CancelJob` | 取消任务 |
| `SuspendJob` | 暂停任务 |
| `ResumeJob` | 恢复任务 |

### Quota 接口

| 方法 | 说明 |
| --- | --- |
| `ListQuotas` | 查询 Quota 列表 |
| `CreateQuota` | 创建 Quota |
| `UpdateQuota` | 修改 Quota |
| `DeleteQuota` | 删除 Quota |

## 使用注意事项

- 创建、更新、删除、启停、重启、扩缩容和位点重置等接口会直接修改真实资源，调用前应核对请求参数。
- 多数变更接口返回 `ActionID`。需要通过任务接口查询异步操作状态，不能仅根据请求成功判断资源已完成变更。
- 不要并发修改同一资源的配置、规格或状态，避免操作冲突。
- 客户端可在多个请求间复用。自定义 `HTTPClient` 时，应设置合理的连接池和超时。
- 生产环境应记录错误码、HTTP 状态码和 `RequestId`，同时过滤凭证、密码和消息内容。
