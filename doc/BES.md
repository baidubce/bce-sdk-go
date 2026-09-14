# 百度Elasticsearch(BES)服务

## 概述

BES Go SDK 封装了百度智能云 Elasticsearch 服务的集群、实例、插件、日志、配置、智能巡检、自动续费、定时调度和标签等接口。

使用 SDK 前，需要准备有效的访问凭证，并确认资源所在区域和服务 Endpoint。

SDK 按 OpenAPI 版本拆成两个独立包，二者的路径风格、字段命名和响应结构都不同，不能混用：

```go
import bes "github.com/baidubce/bce-sdk-go/services/bes/v2"
```

```go
import bes "github.com/baidubce/bce-sdk-go/services/bes/v3"
```

v2 与 v3 的主要差异：

| 差异点 | v2 | v3 |
| --- | --- | --- |
| 集群 ID 形态 | 纯数字串，如 `1210468255077109760` | `search-` 前缀串 |
| 请求风格 | 全部 POST，固定路径，参数走请求体 | RESTful，路径含变量 |
| 响应结构 | 统一包一层 `success` / `status` / `result` | 业务字段在顶层 |

## 初始化客户端

### Endpoint

SDK 会从 Endpoint 的主机名中解析区域，并作为 `x-Region` 请求头发送。解析失败时使用 SDK 默认区域 `bj`。服务端对该请求头大小写敏感，SDK 已固定为 `x-Region`，调用方不需要自己设置。

### 使用 AK/SK

```go
package main

import (
	"fmt"

	bes "github.com/baidubce/bce-sdk-go/services/bes/v3"
)

func main() {
	client, err := bes.NewClient(
		"<your-access-key-id>",
		"<your-secret-access-key>",
		"https://bes.bj.baidubce.com",
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.ListClusters(&bes.ListClustersRequest{
		PageNo:   1,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}
	if resp.Page == nil {
		return
	}

	for _, cluster := range resp.Page.Result {
		fmt.Printf("clusterId=%s name=%s status=%s\n",
			cluster.ClusterId, cluster.ClusterName, cluster.ActualStatus)
	}
}
```

请注意不要在代码仓库、日志或命令行历史中保存真实 AK/SK。建议通过环境变量或密钥管理服务注入凭证。

客户端创建后可以复用，不需要为每个请求重复创建。

### 使用 STS 凭证

```go
client, err := bes.NewClientWithSTS(
	"<temporary-access-key-id>",
	"<temporary-secret-access-key>",
	"<session-token>",
	"https://bes.bj.baidubce.com",
)
if err != nil {
	panic(err)
}
```

STS 凭证过期后，需要使用新凭证重新创建客户端。

## 使用示例

### 查询集群列表

`ListClusters` 是使用其他接口的入口，绝大多数方法都需要它返回的 `ClusterId`。其中 `PageNo` 和 `PageSize` 在V3版本必填。

```go
resp, err := client.ListClusters(&bes.ListClustersRequest{
	PageNo:   1,
	PageSize: 10,
})
if err != nil {
	panic(err)
}
if resp.Page == nil {
	return
}

fmt.Printf("totalCount=%d\n", resp.Page.TotalCount)
for _, cluster := range resp.Page.Result {
	fmt.Printf("clusterId=%s name=%s status=%s health=%s\n",
		cluster.ClusterId, cluster.ClusterName, cluster.ActualStatus,
		cluster.ClusterHealth.Status)
}
```

### 查看定时调度任务

`ListSchedules` 只接受 `ClusterId`，**没有分页参数**。返回的 `ScheduleName` 是更新和删除定时调度任务时的定位键。

```go
resp, err := client.ListSchedules(&bes.ListSchedulesRequest{
	ClusterId: "<cluster-id>",
})
if err != nil {
	panic(err)
}
if resp.Result == nil {
	return
}

for _, schedule := range resp.Result.Schedules {
	fmt.Printf("name=%s taskType=%s schedule=%q status=%s\n",
		schedule.ScheduleName, schedule.TaskType, schedule.Schedule, schedule.Status)
}
```

### 列举可选巡检项

`ListInspectItems` 是唯一一个不接受请求参数的方法，也是智能巡检域里唯一不要求集群先授权的接口。返回的 14 个巡检项 ID 就是 `UpdateManualInspectConfig` 的 `Items` 取值范围。

```go
resp, err := client.ListInspectItems()
if err != nil {
	panic(err)
}
if !resp.Success.Bool() || resp.Result == nil {
	return
}

for _, item := range resp.Result.Items {
	fmt.Printf("id=%s type=%s\n", item.Id, item.Type)
}
```

### 创建定时调度任务

`Task` 的结构随 `TaskType` 变化。其中V2版本 `TaskType` 支持 `CREATE_INDEX`、`DELETE`、`INDEX_STORE_MANAGER`、`MIGRATE_COLD`、`SNAPSHOT`、`CLUSTER_SETTINGS`、`ROLLOVER`、`FORCEMERGE` 共 8 种；V3版本仅支持 `DELETE`、`ROLLOVER`、`FORCEMERGE` 3 种。

`Schedule` 是 Quartz 风格的 6 段 cron 表达式，日和周互斥位写 `?`。`ScheduleName` 只支持字母和数字，长度 6 到 24 个字符，同集群内唯一。

```go
resp, err := client.CreateSchedule(&bes.CreateScheduleRequest{
	ClusterId:    "<cluster-id>",
	Schedule:     "0 2 0 * * ?",
	ScheduleName: "<schedule-name>",
	TaskType:     "DELETE",
	Task: map[string]interface{}{
		"indexPatterns": []string{"log-*"},
		"minIndexAge":   "24h",
	},
})
if err != nil {
	panic(err)
}
fmt.Printf("success=%v status=%d\n", resp.Success, resp.Status)
```

`ClusterId`、`Schedule`、`ScheduleName`、`TaskType` 和 `Task` 全部必填，`Task` 为空 map 会被本地校验拦下。

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

resp, err := client.GetClusterDetail(&bes.GetClusterDetailRequest{
	ClusterId: "<cluster-id>",
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

**业务失败不一定返回 error。** 部分接口在业务校验失败时返回 HTTP 200，响应体里 `success` 为 `false`、`status` 为 4xx，而错误消息在 SDK 未建模的字段里，因此既拿不到 `error` 也拿不到消息。例如未授权集群上调用智能巡检接口，得到的是 `Success` 为 `"false"`、`Status` 为 400、`Result` 为 `nil`。判断调用是否成功应同时检查 `err` 和 `Success`：

```go
resp, err := client.CheckAutoInspect(&bes.CheckAutoInspectRequest{
	ClusterId: "<cluster-id>",
})
if err != nil {
	panic(err)
}
if !resp.Success.Bool() {
	fmt.Printf("业务失败，status=%d\n", resp.Status)
	return
}
```

排查服务端错误时，应保留 `RequestId`，但不要在日志中输出 AK/SK 或集群密码。

## 接口列表（V2版本）

### 集群接口

| 方法 | 说明 |
| --- | --- |
| `CreateCluster` | 创建集群 |
| `ListClusters` | 查询集群列表 |
| `GetClusterDetail` | 查询集群详情 |
| `DeleteCluster` | 删除集群 |
| `StartCluster` | 启动集群 |
| `StopCluster` | 停止集群 |
| `RestartCluster` | 重启集群 |
| `ResizeCluster` | 扩容、升配或降配集群 |
| `AddClusterModule` | 新增节点类型 |
| `ResetClusterPassword` | 重置集群密码 |
| `ToggleClusterHTTPS` | 开启或关闭 HTTPS |
| `BindClusterEIP` | 绑定 EIP |
| `UnbindClusterEIP` | 解绑 EIP |
| `ToggleClusterMonitor` | 开启或关闭 Grafana 监控 |
| `GetClusterTasks` | 查询操作历史 |
| `GetClusterDataSizeTendency` | 查询数据量观测数据 |
| `ListAvailableCoupons` | 查询可用代金券 |
| `AssessClusterSource` | 智能评估集群配置 |

### 实例接口

| 方法 | 说明 |
| --- | --- |
| `StartInstance` | 启动实例 |
| `StopInstance` | 停止实例 |
| `BatchStartInstances` | 批量启动实例 |
| `BatchStopInstances` | 批量停止实例 |
| `DeleteInstances` | 删除实例，即缩容 |
| `ListScaleInInstances` | 查询可缩容节点列表 |
| `ConfirmDataMigration` | 发起数据迁移 |
| `RollbackDataMigration` | 回滚数据迁移 |
| `ListDataMigrationInstances` | 查询可迁移节点列表 |
| `SuggestDataMigrationInstances` | 查询系统建议的迁移节点 |

### 智能巡检接口

| 方法 | 说明 |
| --- | --- |
| `AuthorizeInspect` | 集群巡检授权 |
| `SwitchAutoInspect` | 开启或关闭自动巡检 |
| `CheckAutoInspect` | 查询是否开启自动巡检 |
| `CreateManualInspectTask` | 提交手动巡检任务 |
| `CheckInspectBusy` | 查询是否可提交巡检任务 |
| `GetManualInspectCount` | 查询今日手动巡检次数 |
| `GetManualInspectConfig` | 查询手动巡检配置 |
| `UpdateManualInspectConfig` | 修改手动巡检配置 |
| `ListInspectItems` | 列举所有可选巡检项 |
| `GetInspectTask` | 查询巡检任务状态和结果 |
| `ListInspectTasks` | 查询近 7 天已完成的巡检任务 |
| `GetLatestInspectOverview` | 查询最新一次巡检概况 |
| `GetWeeklyInspectOverview` | 查询近七天巡检概况 |

除 `ListInspectItems` 外，其余接口都要求集群已通过 `AuthorizeInspect` 授权，未授权时返回业务错误。

### 定时调度接口

| 方法 | 说明 |
| --- | --- |
| `CreateSchedule` | 创建定时调度任务 |
| `UpdateSchedule` | 更新定时调度任务 |
| `ListSchedules` | 查询定时调度任务 |
| `DeleteSchedule` | 删除定时调度任务 |

### 自动续费接口

| 方法 | 说明 |
| --- | --- |
| `CreateAutoRenewRule` | 创建自动续费规则 |
| `GetAutoRenewRuleDetail` | 查询自动续费规则详情 |
| `ListAutoRenewRules` | 查询自动续费规则列表 |
| `UpdateAutoRenewRule` | 修改自动续费规则 |
| `DeleteAutoRenewRule` | 删除自动续费规则 |
| `RenewCluster` | 集群续费 |
| `ListRenewals` | 查询即将到期的集群 |

`RenewTimeUnit` 取值为小写的 `month` 或 `year`。`RenewCluster` 的 `Time` 单位固定为月。

### 插件接口

| 方法 | 说明 |
| --- | --- |
| `InstallDefaultPlugin` | 安装系统预置插件 |
| `UninstallDefaultPlugin` | 卸载系统预置插件 |
| `InstallCustomPlugin` | 安装自定义插件 |
| `UninstallCustomPlugin` | 卸载自定义插件 |
| `DeleteCustomPlugin` | 删除自定义插件 |
| `UploadCustomPlugin` | 上传自定义插件 |
| `GetPluginInfo` | 查询插件信息 |
| `GetNLPDict` | 查询 NLP 分词词典 |
| `UpdateNLPDict` | 上传 NLP 分词词典 |

### 配置接口

| 方法 | 说明 |
| --- | --- |
| `GetClusterConfig` | 查询集群额外配置 |
| `UpdateClusterConfig` | 修改集群额外配置 |
| `ListSynonymDicts` | 查询同义词文件列表 |
| `UploadSynonymDict` | 上传同义词配置文件 |
| `DeleteSynonymDict` | 删除同义词配置文件 |

### 日志接口

| 方法 | 说明 |
| --- | --- |
| `UpdateLogSettings` | 修改日志采集开关 |
| `SearchLog` | 检索日志 |
| `CreateLogExportTask` | 创建日志导出任务 |
| `GetLogExportRecord` | 查询日志导出记录 |

### 标签接口

| 方法 | 说明 |
| --- | --- |
| `ListTags` | 查询标签列表 |
| `UpdateClusterTags` | 更新集群标签 |
| `BatchInsertTags` | 批量添加标签 |

## 使用注意事项

- 集群 ID 是纯数字串。文档说明超过 18 位会返回 500 错误，不要把 v3 的 `search-` 前缀 ID 传入 v2 接口。
- 以下接口会真实下单并可能产生费用，调用前应确认账号和环境：`CreateCluster`、`ResizeCluster`、`AddClusterModule`、`DeleteInstances`、`DeleteCluster`、`CreateAutoRenewRule`、`RenewCluster`。它们的响应中包含 `OrderId`。
- 以下接口不可逆：`DeleteCluster`（数据全部丢失）、`DeleteInstances`（缩容）。缩容后集群至少要保留两个节点。
- 以下接口影响集群可用性：集群和实例的启停与重启、`ToggleClusterHTTPS`（会改变客户端访问协议）、`ResetClusterPassword`（旧密码立即失效）。
- 不要并发修改同一集群的配置、规格或状态，避免操作冲突。
- 客户端可在多个请求间复用。生产环境应记录错误码、HTTP 状态码和 `RequestId`，同时过滤凭证和密码。