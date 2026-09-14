EO服务

# 概述

本文档主要介绍EO GO SDK的使用。在使用本文档前，您需要先了解EO的一些基本知识，并已开通了EO服务。若您还不了解CDN，可以参考[产品介绍](https://cloud.baidu.com/doc/GEO/s/lmj18vwxu)和[快速入门](https://cloud.baidu.com/doc/GEO/s/ymocqbtaz)。

# 初始化

## 确认Endpoint

目前使用EO服务时，EO的 Endpoint 统一使用`https://geo.baidubce.com`，这也是默认值。

## 获取密钥

要使用百度云EO，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问EO做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 使用AK/SK新建EO Client

通过AK/SK方式访问EO，用户可以参考如下代码新建一个EO Client：

```go
ak := "your_access_key_id"
sk := "your_secret_key_id"
endpoint := "geo.baidubce.com"

cli, err := eo.NewClient(ak, sk, endpoint)
```

在上面代码中，变量`ak`对应控制台中的“Access Key ID”，变量`sk`对应控制台中的“Access Key Secret”，获取方式请参考《 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。变量`endpoint`必须为`https://geo.baidubce.com`，也是默认值，为空表示使用默认值，设置为其他则SDK无法工作。

在下面的示例中，会频繁使用到GetDefaultClient函数，它的定义为：

```go
func GetDefaultClient() *eo.Client {
	ak := "your_access_key_id"
	sk := "your_secret_key_id"
	endpoint := "https://geo.baidubce.com"

	// ignore error in test, but you should handle error in dev
	client, _ := eo.NewClient(ak, sk, endpoint)
	return client
}
```

## 缓存管理接口

### 刷新缓存/查询刷新状态 Purge/GetPurgedStatus

> 缓存清除方式有URL刷新、目录刷新。URL刷新是以文件或一个资源为单位进行缓存刷新。目录刷新是以目录为单位，将目录下的所有文件进行缓存清除。提交刷新任务时需要指定站点（site）。

```go
// 刷除
cli := GetDefaultClient()
purgedId, err := cli.Purge("your_site.com", []api.PurgeTask{
	{
		Url:  "http://your_site.com/path/to/purge/1.data",
		Type: "file",
	},
	{
		Url:  "http://your_site.com/path/to/purge/html/",
		Type: "directory",
	},
})
fmt.Printf("purgedId:%+v\n", purgedId)
fmt.Printf("err:%+v\n", err)

// 方式一：根据站点和任务ID查询刷除状态
purgedStatus, err := cli.GetPurgedStatus(&api.PurgeStatusQueryData{
	Site: "your_site.com",
	Id:   string(purgedId),
})
fmt.Printf("purgedStatus:%+v\n", purgedStatus)
fmt.Printf("err:%+v\n", err)

// 方式二：根据站点和时间范围查询刷除状态（可选按刷新类型过滤）
purgedStatus, err := cli.GetPurgedStatus(&api.PurgeStatusQueryData{
	Site:      "your_site.com",
	StartTime: "2026-05-01T00:00:00Z",
	EndTime:   "2026-05-31T23:59:59Z",
	Type:      "file", // 可选，按刷新类型过滤，可选值为 "file" 或 "directory"
})
fmt.Printf("purgedStatus:%+v\n", purgedStatus)
fmt.Printf("err:%+v\n", err)
```

接口更多细节可以参考缓存管理文档：https://cloud.baidu.com/doc/GEO/s/4mhsrv9ry  、 https://cloud.baidu.com/doc/GEO/s/mmhssw91q

### 预热资源/查询预热状态 Prefetch/GetPrefetchStatus

> URL预热是以文件为单位进行资源预热。

```go
// 预热
cli := GetDefaultClient()    
prefetchId, err := cli.Prefetch("your_site.com", []api.PrefetchTask{
	{
		Url: "http://your_site.com/path/to/prefetch/1.data",
	},
	{
		Url: "http://your_site.com/path/to/prefetch/2.data",
	},
})
fmt.Printf("prefetchId:%+v\n", prefetchId)
fmt.Printf("err:%+v\n", err)

// 方式一：根据站点和任务ID查询预热状态
prefetchStatus, err := cli.GetPrefetchStatus(&api.PrefetchStatusQueryData{
	Site: "your_site.com",
	Id:   string(prefetchId),
})
fmt.Printf("prefetchStatus:%+v\n", prefetchStatus)
fmt.Printf("err:%+v\n", err)

// 方式二：根据站点和时间范围查询预热状态
prefetchStatus, err := cli.GetPrefetchStatus(&api.PrefetchStatusQueryData{
	Site:      "your_site.com",
	StartTime: "2026-05-01T00:00:00Z",
	EndTime:   "2026-05-31T23:59:59Z",
})
fmt.Printf("prefetchStatus:%+v\n", prefetchStatus)
fmt.Printf("err:%+v\n", err)
```

接口更多细节可以参考缓存管理文档：https://cloud.baidu.com/doc/GEO/s/5mhsuituv 、 https://cloud.baidu.com/doc/GEO/s/Bmhsv5i9u

## 离线日志接口

### 获取离线日志下载地址 GetOfflineLog

> 获取用户某个站点下单个域名或多个域名某一指定时间段内的日志下载地址。日志的保存时间为 180 天。

```go
cli := GetDefaultClient()
logResult, err := cli.GetOfflineLog(&api.LogQueryData{
	Site:       "your_site.com",
	StartTime:  "2026-05-01T00:00:00Z",
	EndTime:    "2026-05-31T23:59:59Z",
	DomainList: []string{"your.domain.com", "your.domain.com"},
	PageNo:     1,
	PageSize:   20,
})
fmt.Printf("logResult:%+v\n", logResult)
fmt.Printf("err:%+v\n", err)
```

`api.LogQueryData`字段说明：

| 字段       | 类型     | 是否必选 | 说明                                                                  |
| ---------- | -------- | -------- | --------------------------------------------------------------------- |
| Site       | String   | 是       | 站点名称。                                                            |
| StartTime  | String   | 否       | 查询时间范围起始值，UTC 时间。默认为 `EndTime` 前推 8 小时。 |
| EndTime    | String   | 否       | 查询时间范围结束值，UTC 时间。默认为当前时间。          |
| DomainList | []String | 否       | 查询的域名列表。                                                      |
| PageNo     | int      | 否       | 分页编号，默认值为 1。                                                |
| PageSize   | int      | 否       | 每页返回日志数目，默认值为 20。                                       |

`logResult`是`*api.LogQueryResult`类型的对象，详细说明如下：

| 字段          | 类型        | 说明                 |
| ------------- | ----------- | -------------------- |
| LogEntryList  | []LogEntry  | 离线日志列表。       |
| TotalCount    | String      | 离线日志列表总数。   |

`LogEntry`类型说明：

| 字段          | 类型   | 说明                                  |
| ------------- | ------ | ------------------------------------- |
| Domain        | String | 域名。                                |
| Url           | String | 可下载离线日志的 URL。                |
| Name          | String | 离线日志文件名称。                    |
| Size          | int64  | 离线日志文件大小，单位为 B。          |
| LogTimeBegin  | String | 文件中日志开始时间，UTC 时间。        |
| LogTimeEnd    | String | 文件中日志结束时间，UTC 时间。        |

## 统计接口

> 查询用户站点维度或域名维度的统计指标信息，支持总带宽 / 上下行带宽、总流量 / 上下行流量、PV 等指标。
> 按接口规范，请求被划分为 5 种类型，分别对应不同的 `(group, showType)` 组合，SDK 提供 5 个独立方法，每个方法对应一种类型，调用方无需关心 `showType`：

| 类型 | 方法                  | 是否带 group | 内置 showType | 说明              |
| ---- | --------------------- | ------------ | ------------- | ----------------- |
| 1    | `GetStatTime`         | 否           | `time`        | 时间打点          |
| 2    | `GetStatPeak`         | 否           | `peak`        | 峰值              |
| 3    | `GetStatTimeByGroup`  | 是           | `time`        | 时间打点带分组    |
| 4    | `GetStatSumByGroup`   | 是           | `sum`         | 聚合带分组        |
| 5    | `GetStatTopByGroup`   | 是           | `top`         | TOP 带分组        |

### 类型 1：GetStatTime — 时间打点（无 group）

```go
cli := GetDefaultClient()
result, err := cli.GetStatTime(&api.StatTimeQueryData{
	Metric:    []string{"sum_bps"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T09:06:29Z",
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})
b, _ := json.Marshal(result)
fmt.Printf("result:%s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 2：GetStatPeak — 峰值（无 group）

```go
cli := GetDefaultClient()
result, err := cli.GetStatPeak(&api.StatPeakQueryData{
	Metric:    []string{"sum_bps", "upstream_bps", "download_bps"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T09:06:29Z",
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})
b, _ := json.Marshal(result)
fmt.Printf("result:%s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 3：GetStatTimeByGroup — 时间打点带 group

```go
cli := GetDefaultClient()
result, err := cli.GetStatTimeByGroup(&api.StatTimeByGroupQueryData{
	Metric:    []string{"pv"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T11:05:47Z",
	Group:     []string{"code"},
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})
b, _ := json.Marshal(result)
fmt.Printf("result:%s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 4：GetStatSumByGroup — 聚合带 group

```go
cli := GetDefaultClient()
result, err := cli.GetStatSumByGroup(&api.StatSumByGroupQueryData{
	Metric:    []string{"pv"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T11:05:47Z",
	Group:     []string{"code"},
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
})

b, _ := json.Marshal(result)
fmt.Printf("stat sum by group: %s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 类型 5：GetStatTopByGroup — TOP 带 group

```go
cli := GetDefaultClient()
result, err := cli.GetStatTopByGroup(&api.StatTopByGroupQueryData{
	Metric:    []string{"pv"},
	StartTime: "2026-05-10T16:00:00Z",
	EndTime:   "2026-05-11T11:05:47Z",
	Group:     []string{"host"},
	Filter: []api.StatFilterItem{
		{
			Key:       "site",
			Operation: "equal",
			Value:     []string{"your_site.com"},
		},
	},
	Limit: &api.StatLimit{PageSize: 100}, // 可选，限制 TOP 返回条数
})
b, _ := json.Marshal(result)
fmt.Printf("result: %s\n", b)
fmt.Printf("err:%+v\n", err)
```

### 请求字段说明

5 种请求类型共享如下基础字段：

| 字段       | 类型              | 是否必选 | 说明                                                                                                            |
| ---------- | ----------------- | -------- |---------------------------------------------------------------------------------------------------------------|
| Metric     | []String          | 是       | 指标类型，可选值：`sum_bps` / `upstream_bps` / `download_bps` / `sum_flow` / `upstream_flow` / `download_flow` / `pv`。 |
| StartTime  | String            | 否       | 查询时间范围起始值，UTC 时间。默认 `EndTime` 前推 24 小时。最长可查近 31 天。                                                            |
| EndTime    | String            | 否       | 查询时间范围结束值，UTC 时间。默认当前时间。                                                                                      |
| Filter     | []StatFilterItem  | 否       | 过滤条件列表。                                                                                                       |

类型 3、4、5 在以上字段基础上额外要求 `Group` 字段：

| 字段   | 类型     | 是否必选 | 说明                                  |
| ------ | -------- | -------- | ------------------------------------- |
| Group  | []String | 是       | 聚合字段，例如 `code`、`host`。       |

类型 5（TOP 查询）在以上字段基础上还可选 `Limit` 字段：

| 字段   | 类型         | 是否必选 | 说明                                   |
| ------ | ------------ | -------- | -------------------------------------- |
| Limit  | *StatLimit   | 否       | TOP 返回条数限制，例如 `{PageSize:100}`。 |

`api.StatFilterItem` 字段说明：

| 字段       | 类型     | 是否必选 | 说明                                          |
| ---------- | -------- | -------- | --------------------------------------------- |
| Key        | String   | 是       | 过滤的 key，支持 `site`、`host`。             |
| Value      | []String | 是       | key 对应的值，支持多个。                      |
| Operation  | String   | 是       | 操作类型，支持 `equal`、`notequal`。          |

### 返回值说明

`result`是`api.StatResult`类型（即 `map[string][]api.StatDataPoint`），key 为请求中的 metric 名称（例如 `sum_bps`、`pv`），value 是该指标的数据点列表。

`api.StatDataPoint` 字段说明：

| 字段        | 类型             | 说明                                                                                  |
| ----------- | ---------------- | ------------------------------------------------------------------------------------- |
| Timestamp   | *int64           | 时间戳。类型 1 / 3（`time`）时为有效值；类型 2 / 4 / 5（`peak` / `sum` / `top`）时为 `nil`。 |
| Value       | json.Number      | 指标值。服务端可能返回数字或字符串，统一用 `json.Number` 承载，可通过 `Value.Int64()` / `Value.Float64()` / `Value.String()` 取值。 |
| GroupParams | json.RawMessage  | 聚合参数原始 JSON。无 group（类型 1 / 2）时为空数组 `[]`；带 group（类型 3 / 4 / 5）时为对象（如 `{"code":"2xx"}`）。 |


更多详细说明可以参考API文档：https://cloud.baidu.com/doc/GEO/s/4mp0unbx3

# 站点配置接口

EO 站点的全局配置功能共用同一接口：

- `PUT /v2/geo/site/{site}/config`：设置配置
- `GET /v2/geo/site/{site}/config`：查询配置（返回站点全部配置项）

`api.SiteConfig` 是统一的配置容器，所有字段均为指针类型 + `omitempty`：未赋值的字段不会出现在请求 body 中，未来新增配置项时按相同方式扩展即可。如需将某个配置显式设置为空（例如 `cacheTtl: []` 表示遵循源站-默认缓存策略），传一个非 nil 的空值（如 `&[]api.CacheTtl{}`）。

相关配置接口说明可参考对应配置的API文档：https://cloud.baidu.com/doc/GEO/s/Pmiigxbf0

## 设置节点缓存配置

```go
cli := GetDefaultClient()
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	CacheTtl: &[]api.CacheTtl{
		{
			Value:          "/",
			Weight:         100,
			OverrideOrigin: true,
			Ttl:            2592000,
			Type:           "path",
		},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.CacheTtl` 字段说明：

| 字段           | 类型   | 是否必选 | 说明                                                             |
| -------------- | ------ | -------- |----------------------------------------------------------------|
| Type           | String | 是       | 其合法值为“path”。表示缓存目录的路径。                                         |
| Value          | String | 是       | 其合法值为“/”。表示根目录。                                                |
| Weight         | Int    | 是       | 权重，合法值 `100`。                                                  |
| OverrideOrigin | Bool   | 是       | 表示缓存是否遵循源站。值为 `true` 时，表示不遵循源站，按照该条配置规则缓存。值为 `false` 时，表示遵循源站。 |
| Ttl            | Int    | 是       | 缓存时间，单位秒；`0` 表示不缓存。                                            |

## 设置查询字符串配置

```go
cli := GetDefaultClient()
query := false
cacheKeyIgnoreCase := false
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	CacheKey: &api.CacheKey{
		Query:       &query,
		IncludeArgs: &[]string{"test1"},
		IgnoreCase:  &cacheKeyIgnoreCase,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.CacheKey` 字段说明（`IncludeArgs` 与 `ExcludeArgs` 不可同时设置；二者仅在 `Query=false` 时生效）：

| 字段        | 类型      | 是否必选 | 说明                                    |
| ----------- | --------- | -------- |---------------------------------------|
| Query       | *Bool     | 否       | `true` 保留全部参数参与缓存；`false` 忽略全部参数参与缓存。 |
| IgnoreCase  | *Bool     | 否       | `true` 开启忽略大小写；`false` 关闭忽略大小写。       |
| IncludeArgs | *[]String | 否       | 保留指定参数参与缓存（仅 `Query=false` 时有效）。      |
| ExcludeArgs | *[]String | 否       | 忽略指定参数参与缓存（仅 `Query=false` 时有效）。      |

## 设置离线模式配置

```go
cli := GetDefaultClient()
offlineMode := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	OfflineMode: &offlineMode,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`OfflineMode` 字段说明：

| 字段        | 类型    | 是否必选 | 说明                                            |
| ----------- | ------- |------| ----------------------------------------------- |
| OfflineMode | *String | 是    | `ON` 开启离线模式；`OFF` 关闭离线模式。         |

## 设置强制 HTTPS 配置

```go
cli := GetDefaultClient()
httpToHttpsEnabled := "ON"
httpToHttpsCode := "302"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	HttpToHttpsEnabled: &httpToHttpsEnabled,
	HttpToHttpsCode:    &httpToHttpsCode,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`HttpToHttps*` 字段说明：

| 字段               | 类型    | 是否必选 | 说明                                                              |
| ------------------ | ------- |------| ----------------------------------------------------------------- |
| HttpToHttpsEnabled | *String | 是    | `ON` 开启强制 HTTPS 重定向；`OFF` 关闭。                          |
| HttpToHttpsCode    | *String | 否    | 重定向状态码，`301` 或 `302`；`HttpToHttpsEnabled=OFF` 时无效。   |

## 设置 HSTS 配置

```go
cli := GetDefaultClient()
maxAge := -1
includeSubDomains := false
preload := false
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Hsts: &api.HSTS{
		MaxAge:            &maxAge,
		IncludeSubDomains: &includeSubDomains,
		Preload:           &preload,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.HSTS` 字段说明：

| 字段              | 类型  | 是否必选 | 说明                                             |
| ----------------- | ----- | -------- |------------------------------------------------|
| MaxAge            | *Int  | 是       | 配置保存时间，单位为天, 用户输入值为 0 ~ 730 或者 -1，为 -1 时表示关闭该配置项。 |
| IncludeSubDomains | *Bool | 是       | 是否包含子域名。                                       |
| Preload           | *Bool | 是       | 是否支持预加载。                                       |

## 设置 HTTP2 配置

```go
cli := GetDefaultClient()
http2Disable := "OFF"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Http2Disable: &http2Disable,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Http2Disable` 字段说明：

| 字段         | 类型    | 是否必选 | 说明                                          |
| ------------ | ------- |------| --------------------------------------------- |
| Http2Disable | *String | 是    | `OFF` 开启 HTTP2；`ON` 关闭 HTTP2。           |

## 设置 HTTP3 配置

```go
cli := GetDefaultClient()
http3Enable := true
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Http3: &api.HTTP3{
		Enable: &http3Enable,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.HTTP3` 字段说明：

| 字段   | 类型  | 是否必选 | 说明                                |
| ------ | ----- | -------- |-----------------------------------|
| Enable | *Bool | 是       | `true` 开启 HTTP3；`false` 关闭 HTTP3。 |

## 设置最大上传大小配置

```go
cli := GetDefaultClient()
clientMaxBodySize := "500m"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	ClientMaxBodySize: &clientMaxBodySize,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`ClientMaxBodySize` 字段说明：

| 字段              | 类型    | 是否必选 | 说明                                                                                          |
| ----------------- | ------- |------| --------------------------------------------------------------------------------------------- |
| ClientMaxBodySize | *String | 是    | 单位支持 `b`、`k`、`m`（忽略大小写，`b` 可省略）。最大 `500m`。示例：`100`、`100k`、`100M`。 |

## 设置页面压缩配置

```go
cli := GetDefaultClient()
compress := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Compress:            &compress,
	CompressMethodArray: &[]string{"gzip", "br"},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Compress*` 字段说明：

| 字段                | 类型      | 是否必选 | 说明                          |
| ------------------- | --------- |------|-----------------------------|
| Compress            | *String   | 是    | `ON` 开启页面压缩；`OFF` 关闭页面压缩。   |
| CompressMethodArray | *[]String | 是    | 压缩方式，合法值：`gzip`、`br`，可同时启用。 |

## 设置智能加速配置

```go
cli := GetDefaultClient()
isa := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Isa: &isa,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Isa` 字段说明：

| 字段 | 类型    | 是否必选 | 说明                        |
| ---- | ------- |------|---------------------------|
| Isa  | *String | 是    | `ON` 开启智能加速；`OFF` 关闭智能加速。 |

## 设置状态码缓存配置

```go
cli := GetDefaultClient()
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	CacheCodeTtl: &[]api.CacheCodeTtl{
		{Value: "404", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
		{Value: "400", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.CacheCodeTtl` 字段说明：

| 字段           | 类型   | 是否必选 | 说明                                                                                  |
| -------------- | ------ | -------- | ------------------------------------------------------------------------------------- |
| Type           | String | 是       | 合法值 `code`，表示异常状态码缓存。                                                   |
| Value          | String | 是       | 4xx：`400`/`401`/`403`/`404`/`405`/`407`/`414`/`451`；<br/>5xx：`500`/`501`/`502`/`503`/`504`/`509`/`514`。 |
| Ttl            | Int    | 是       | 缓存时间，单位秒，取值范围 `0~315360000`。                                            |
| OverrideOrigin | Bool   | 是       | 合法值 `true`。                                                                       |
| Weight         | Int    | 是       | 权重，合法值 `100`。                                                                  |

注：最多 15 条状态码缓存规则，且不可重复。

## 设置 gRPC 回源配置

```go
cli := GetDefaultClient()
grpcOrigin := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	GrpcOrigin: &grpcOrigin,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`GrpcOrigin` 字段说明：

| 字段       | 类型    | 是否必选 | 说明                                  |
| ---------- | ------- |------| ------------------------------------- |
| GrpcOrigin | *String | 是    | `ON` 开启 gRPC 回源；`OFF` 关闭 gRPC 回源。     |

## 设置 HTTP2 回源配置

```go
cli := GetDefaultClient()
http2Origin := "ON"
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Http2Origin: &http2Origin,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Http2Origin` 字段说明：

| 字段        | 类型    | 是否必选 | 说明                                       |
| ----------- | ------- |------| ------------------------------------------ |
| Http2Origin | *String | 是    | `ON` 开启 HTTP2 回源；`OFF` 关闭 HTTP2 回源。         |

## 设置规则引擎配置

```go
cli := GetDefaultClient()

// 忽略大小写
ignoreCase := true

// 各配置项取值
cacheKeyQuery := false
cacheKeyIgnoreCase := true
offlineMode := "OFF"
refreshRevalidateEnabled := true
httpToHttpsEnabled := "ON"
httpToHttpsCode := "302"
hstsMaxAge := -1
hstsIncludeSubDomains := false
hstsPreload := false
isa := "ON"
http2Disable := "OFF"
http3Enable := true
webSocketEnabled := true
webSocketTimeout := 10
http2Origin := "ON"
clientMaxBodySize := "500m"
compress := "ON"
originLoadTimeout := 30
originConnectTimeout := 5
enableRedirectFollow := "ON"
maxRedirectFollowCount := 2
originRange := "force_all"
originPartSize := "512k"
realIpEnabled := true
realIpName := "True-Client-Ip"
errorPageCode400 := 400
errorPageUrl400 := "http://test.eo.com/400.html"
errorPageCode403 := 403
errorPageUrl403 := "http://test.eo.com/403.html"
trafficLimitEnable := true
trafficLimitRate := 1
trafficLimitStartHour := 10
trafficLimitEndHour := 19
trafficLimitRateUnit := "k"
antiType := "typeA"
antiSecretKey := "your_secret_key"
antiNewsecretKey := "your_new_secret_key"
antiTimeout := 1800
antiTimestampFormat := "dec"
antiAuthArg := "auth_key"
originArgIgnore := true
urlRuleScheme := "http"
urlRuleHost := "test.eo.com"
urlRuleDstPath := "/test/1.txt"
urlRuleQuery := "OFF"
urlRuleStatus := 302
sslMode := "custom"
sslCipherList := "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"
sslModePtr := &sslMode
sslCipherListPtr := &sslCipherList
ocsp := "ON"
clientCacheTtlMode := "custom"
clientCacheTtlValue := 150
originHttpPort := 80
originHttpsPort := 443
originHost := "test.eo.com"

resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	PageRules: &[]api.PageRule{
		{
			Name:   "rule",
			Status: "ON",
			Rules: [][]api.Rule{
				// 同一内层切片内的条件为「与」关系
				{
					{MatchFrom: "path", Operator: "inValues", Values: []string{"/test"}, IgnoreCase: &ignoreCase},
					{MatchFrom: "directory", Operator: "inValues", Values: []string{"/test/"}},
					{MatchFrom: "arg", Operator: "inValues", MatchKey: "test", Values: []string{"abc"}},
					{MatchFrom: "suffix", Operator: "inValues", Values: []string{"jpg"}},
					{MatchFrom: "header", Operator: "inValues", MatchKey: "user-agent", Values: []string{"test"}},
					{MatchFrom: "fullUrl", Operator: "inValues", Values: []string{"/test/test.txt"}},
					{MatchFrom: "basename", Operator: "inValues", Values: []string{"test.mp4"}},
					// 中国内地用 Pro，其他国家用 Cty
					{MatchFrom: "remoteGeos", Operator: "inValues", Values: []api.RemoteGeo{
						{Pro: "beijing"},
						{Cty: "bt"},
					}},
					// 支持 IP 地址和 IP 段
					{MatchFrom: "remoteAddrs", Operator: "inValues", Values: []string{"1.2.3.4", "172.16.0.0/12"}},
					// 运营商用单个字符串，多个以 | 分隔
					{MatchFrom: "remoteIsp", Operator: "inValues", Values: "cm|un|ct"},
					{MatchFrom: "method", Operator: "inValues", Values: []string{"GET", "POST"}},
					{MatchFrom: "cookie", Operator: "inValues", MatchKey: "test", Values: []string{"abc"}},
					{MatchFrom: "host", Operator: "inValues", Values: []string{"1.test.com"}},
				},
				// 不同内层切片之间为「或」关系
				{
					// 正则匹配，Values 传单个字符串
					{MatchFrom: "path", Operator: "regex", Values: "^/example/test[123]/$"},
					// 存在性判断，不需要 Values
					{MatchFrom: "arg", Operator: "exists", MatchKey: "test"},
				},
			},
			Config: &api.RuleConfig{
				// 节点缓存
				CacheTtl: &[]api.CacheTtl{
					{Value: "/", Weight: 100, OverrideOrigin: false, Ttl: 0, Type: "path"},
				},

				// 自定义 cacheKey
				CacheKey: &api.RuleCacheKey{
					Query:       &cacheKeyQuery,
					IgnoreCase:  &cacheKeyIgnoreCase,
					IncludeArgs: &[]string{"a", "b"},
					Headers:     &[]string{"Accept-Language"},
					Cookie:      &[]string{"uid"},
				},

				// 离线模式
				OfflineMode: &offlineMode,

				// 忽略客户端刷新
				RefreshRevalidate: &api.RefreshRevalidate{Enabled: &refreshRevalidateEnabled},

				// 强制 HTTPS
				HttpToHttpsEnabled: &httpToHttpsEnabled,
				HttpToHttpsCode:    &httpToHttpsCode,

				// HSTS
				Hsts: &api.HSTS{
					MaxAge:            &hstsMaxAge,
					IncludeSubDomains: &hstsIncludeSubDomains,
					Preload:           &hstsPreload,
				},

				// 智能加速
				Isa: &isa,

				// HTTP2 / HTTP3 / WebSocket
				Http2Disable: &http2Disable,
				Http3:        &api.HTTP3{Enable: &http3Enable},
				WebSocket: &api.WebSocket{
					Enabled: &webSocketEnabled,
					Timeout: &webSocketTimeout,
				},

				// HTTP2 回源、最大上传大小
				Http2Origin:       &http2Origin,
				ClientMaxBodySize: &clientMaxBodySize,

				// 页面压缩
				Compress:            &compress,
				CompressMethodArray: &[]string{"gzip", "br"},

				// 回源超时时间
				OriginTimeout: &api.OriginTimeout{
					LoadTimeout:    &originLoadTimeout,
					ConnectTimeout: &originConnectTimeout,
				},

				// 回源 301/302 跟随
				OriginRedirectOptions: &api.OriginRedirectOptions{
					EnableRedirectFollow:   &enableRedirectFollow,
					MaxRedirectFollowCount: &maxRedirectFollowCount,
				},

				// 回源 range
				OriginOptions: &api.OriginOptions{
					Range:    &originRange,
					PartSize: &originPartSize,
				},

				// 用户 IP 获取
				RealIp: &api.RealIp{
					Enabled: &realIpEnabled,
					Name:    &realIpName,
				},

				// 自定义错误页面
				ErrorPage: &[]api.ErrorPage{
					{Code: &errorPageCode400, Url: &errorPageUrl400},
					{Code: &errorPageCode403, Url: &errorPageUrl403},
				},

				// 单链接限速
				TrafficLimit: &api.TrafficLimit{
					Enable:         &trafficLimitEnable,
					LimitRate:      &trafficLimitRate,
					LimitStartHour: &trafficLimitStartHour,
					LimitEndHour:   &trafficLimitEndHour,
					LimitRateUnit:  &trafficLimitRateUnit,
				},

				// URL 鉴权
				AntiHotLink: &api.AntiHotLink{
					AntiType:        &antiType,
					SecretKey:       &antiSecretKey,
					NewsecretKey:    &antiNewsecretKey,
					Timeout:         &antiTimeout,
					TimestampFormat: &antiTimestampFormat,
					AuthArg:         &antiAuthArg,
				},

				// 回源请求参数
				OriginArg: &api.OriginArg{
					Ignore: &originArgIgnore,
					Args:   &[]string{"test"},
				},

				// 访问 URL 重定向
				UrlRules: &[]api.UrlRules{
					{
						Scheme:  &urlRuleScheme,
						Host:    &urlRuleHost,
						DstPath: &urlRuleDstPath,
						Query:   &urlRuleQuery,
						Status:  &urlRuleStatus,
					},
				},

				// SSL/TLS 安全配置
				SslProtocols:  &[]string{"TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3"},
				SslMode:       &sslModePtr,
				SslCipherList: &sslCipherListPtr,

				// OCSP Stapling
				Ocsp: &ocsp,

				// 状态码缓存
				CacheCodeTtl: &[]api.CacheCodeTtl{
					{Value: "404", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
					{Value: "400", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
				},

				// 浏览器缓存 TTL
				ClientCacheTtl: &api.ClientCacheTtl{
					Mode: &clientCacheTtlMode,
					Ttl:  &clientCacheTtlValue,
				},

				// HTTP 回源请求头
				OriginRequest: &api.OriginRequest{
					AddHeaders:    &api.HeaderMap{"test": "${uri}", "test2": "11"},
					RemoveHeaders: &api.HeaderMap{"Content-Type": ""},
				},

				// HTTP 节点响应头
				ClientResponse: &api.ClientResponse{
					AddHeaders:    &api.HeaderMap{"test": "${uri}", "test2": "11"},
					RemoveHeaders: &api.HeaderMap{"Content-Type": ""},
				},

				// 源站修改（与 OriginPool 二选一）
				OriginConfig: &[]api.OriginItem{
					{
						Addr:             "1.1.1.1",
						Type:             "IP",
						UpstreamProtocol: "*",
						HttpPort:         &originHttpPort,
						HttpsPort:        &originHttpsPort,
						Host:             &originHost,
					},
				},

				// 源站池（与 OriginConfig 二选一，二者不可同时设置）
				//OriginPool: &api.OriginPool{
				//	OriginPoolId:     "your_origin_pool_id",
				//	UpstreamProtocol: "*",
				//	HttpPort:         &originHttpPort,
				//	HttpsPort:        &originHttpsPort,
				//},
			},
		},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```
规则引擎的详细参数说明可参考API文档：https://cloud.baidu.com/doc/GEO/s/dmimubjhg

`api.PageRule` 字段说明：

| 字段   | 类型             | 是否必选 | 说明           |
| ------ | ---------------- | -------- |--------------|
| Name   | String           | 是       | 规则名称。        |
| Status | String           | 是       | 规则状态。        |
| Rules  | [][]api.Rule     | 是       | 请求的匹配规则。     |
| Config | *api.RuleConfig  | 是       | 命中后生效的配置项集合。 |

注：`SiteConfig.PageRules` 为指针 + `omitempty`，未设置时不下发；如需清空所有规则，可显式传 `&[]api.PageRule{}`。

`api.Rule` 字段说明：

| 字段       | 类型                | 是否必选 | 说明                                                                                        |
| ---------- |-------------------| -------- |-------------------------------------------------------------------------------------------|
| MatchFrom  | String            | 是       | 匹配类型。 |
| Operator   | String            | 是       | 操作符。                                     |
| MatchKey   | String            | 否       | 匹配类型的 key，`arg`、`header`、`cookie` 等需要指定 key 的匹配类型使用。                                |
| Values     | interface{}       | 否       | 匹配值。 |
| IgnoreCase | *Bool             | 否       | 是否忽略大小写。                                                                                  |

`api.RemoteGeo` 字段说明（`matchFrom=remoteGeos` 时作为 `Values` 的元素）：

| 字段 | 类型   | 是否必选 | 说明                       |
| ---- | ------ | -------- |--------------------------|
| Pro  | String | 否       | 中国内地省份，如 `beijing`。      |
| Cty  | String | 否       | 其他国家，如 `bt`。             |

`api.RuleConfig` 主要字段说明：

| 字段                  | 类型                    | 说明              |
| --------------------- | ----------------------- |-----------------|
| CacheTtl              | *[]api.CacheTtl         | 节点缓存。           |
| CacheKey              | *api.RuleCacheKey       | 自定义cacheKey。    |
| OfflineMode           | *String                 | 离线模式。           |
| RefreshRevalidate     | *api.RefreshRevalidate  | 忽略客户端刷新。        |
| HttpToHttpsEnabled    | *String                 | 强制 HTTPS 开关。    |
| HttpToHttpsCode       | *String                 | 强制 HTTPS 跳转状态码。 |
| Hsts                  | *api.HSTS               | HSTS。           |
| Isa                   | *String                 | 智能加速。           |
| Http2Disable          | *String                 | HTTP2。          |
| Http3                 | *api.HTTP3              | HTTP3。          |
| WebSocket             | *api.WebSocket          | WebSocket。      |
| Http2Origin           | *String                 | HTTP2 回源。       |
| ClientMaxBodySize     | *String                 | 最大上传大小。         |
| Compress              | *String                 | 页面压缩开关。         |
| CompressMethodArray   | *[]String               | 页面压缩方式。         |
| OriginTimeout         | *api.OriginTimeout      | 回源超时时间。         |
| OriginRedirectOptions | *api.OriginRedirectOptions | 回源301/302跟随。    |
| OriginOptions         | *api.OriginOptions      | 回源range。        |
| RealIp                | *api.RealIp             | 用户IP获取。         |
| ErrorPage             | *[]api.ErrorPage        | 自定义错误页面。        |
| TrafficLimit          | *api.TrafficLimit       | 单链接限速。          |
| AntiHotLink           | *api.AntiHotLink        | URL鉴权。          |
| OriginArg             | *api.OriginArg          | 回源请求参数。         |
| UrlRules              | *[]api.UrlRules         | 访问URL重定向。       |
| SslProtocols          | *[]String               | SSL/TLS 版本。      |
| SslMode               | **String                | 加密算法套件模式。       |
| SslCipherList         | **String                | 加密算法套件列表。       |
| Ocsp                  | *String                 | OCSP。            |
| CacheCodeTtl          | *[]api.CacheCodeTtl     | 状态码缓存。          |
| ClientCacheTtl        | *api.ClientCacheTtl     | 浏览器缓存 TTL。      |
| OriginConfig          | *[]api.OriginItem       | 源站配置。            |
| OriginPool            | *api.OriginPool         | 源站池。             |
| OriginRequest         | *api.OriginRequest      | HTTP 回源请求头。      |
| ClientResponse        | *api.ClientResponse     | HTTP 节点响应头。      |

`api.RuleCacheKey` 字段说明（`IncludeArgs` 与 `ExcludeArgs` 不可同时设置；二者仅在 `Query=false` 时生效）：

| 字段        | 类型      | 是否必选 | 说明                                                            |
| ----------- | --------- | -------- |---------------------------------------------------------------|
| Query       | *Bool     | 否       | `true` 保留全部参数参与缓存；`false` 忽略全部参数参与缓存。                        |
| IgnoreCase  | *Bool     | 否       | `true` 开启忽略大小写；`false` 关闭忽略大小写。                              |
| IncludeArgs | *[]String | 否       | 保留指定参数参与缓存（仅 `Query=false` 时有效）。                             |
| ExcludeArgs | *[]String | 否       | 忽略指定参数参与缓存（仅 `Query=false` 时有效）。                             |
| Headers     | *[]String | 否       | 参与缓存键的 HTTP 请求头名称，最多 10 个，单个名称最长 255 字符。                     |
| Cookie      | *[]String | 否       | 参与缓存键的 cookie 参数名，最多 10 个，单个参数最长 255 字符。                     |

`api.OriginItem` 字段说明：

| 字段             | 类型                     | 是否必选 | 说明                                                                     |
| ---------------- | ------------------------ | -------- |------------------------------------------------------------------------|
| Addr             | String                   | 是       | 源站地址。支持 IPv4/IPv6 地址或域名，也支持 BOS 的 bucket 地址；不能重复。               |
| Type             | String                   | 是       | 源站类型。合法值：`IP`、`DOMAIN`、`BUCKET`。                                   |
| UpstreamProtocol | String                   | 是       | 回源协议。合法值：`http`、`https`、`*`（协议跟随）。                                |
| HttpPort         | *Int                     | 否       | HTTP 回源端口号。                                                          |
| HttpsPort        | *Int                     | 否       | HTTPS 回源端口号。                                                         |
| Host             | *String                  | 否       | 回源时使用的 host。                                                        |
| ThirdBucketAuth  | *api.ThirdBucketAuth     | 否       | 对象存储源站鉴权配置。                                                |

`api.ThirdBucketAuth` 字段说明（`AuthType` 为 `bos` 时，`Ak`/`Sk`/`Bucket`/`Region`/`Service` 均无需传值）：

| 字段     | 类型    | 是否必选 | 说明                                                                                                          |
| -------- | ------- | -------- |-------------------------------------------------------------------------------------------------------------|
| AuthType | String  | 是       | 对象存储来源类型。合法值：`aws_v2`、`aws_v4`（AWS S3）、`bos`。取 `aws_v2`/`aws_v4` 时，`OriginItem.Type` 必须为 `DOMAIN`；取 `bos` 时，`OriginItem.Type` 必须为 `BUCKET`。 |
| Enabled  | *Bool   | 否       | 是否启用私有 bucket 鉴权，合法值 `true`、`false`，默认 `false`。                                              |
| Ak       | *String | 是       | 对象存储的 Access Key。关闭私有 bucket 鉴权时传空即可。                                                       |
| Sk       | *String | 是       | 对象存储的 Secret Access Key。关闭私有 bucket 鉴权时传空即可。                                                |
| Bucket   | *String | 否       | 对象存储的 bucket。`aws_v2` 必须设置；`aws_v4` 无需设置。                                                     |
| Region   | *String | 否       | 对象存储的区域。`aws_v4` 选填（默认 `us-east-1`）；`aws_v2` 无需设置。                                         |
| Service  | *String | 否       | 对象存储的服务。`aws_v4` 选填（默认 `s3`）；`aws_v2` 无需设置。                                                |

注：「无需设置」表示该字段对当前来源类型无效，即使传值也不会生效。`Ak`/`Sk` 属于敏感凭证，建议从环境变量或密钥管理服务读取，不要硬编码在代码中。

`api.OriginPool` 字段说明：

| 字段             | 类型      | 是否必选 | 说明             |
| ---------------- | --------- | -------- |----------------|
| OriginPoolId     | String    | 是       | 源站池 ID。        |
| UpstreamProtocol | String    | 是       | 回源协议。          |
| HttpPort         | *Int      | 否       | HTTP 回源端口号。    |
| HttpsPort        | *Int      | 否       | HTTPS 回源端口号。   |

注：`OriginConfig` 与 `OriginPool` 不可同时设置，同一条规则中只能选择其中一种回源方式。

## 设置回源请求头配置

> 注意：本接口为全量更新，每次设置需带上希望保留的全部回源请求头，否则原有配置会被覆盖。

```go
cli := GetDefaultClient()
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	OriginRequest: &api.OriginRequest{
		AddHeaders: &api.HeaderMap{
			"test":  "${uri}",
			"test2": "11",
		},
		RemoveHeaders: &api.HeaderMap{
			"Content-Type": "",
		},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

清空回源请求头配置：显式传入非 nil 的空 `HeaderMap`（序列化为 `[]`）。

```go
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	OriginRequest: &api.OriginRequest{
		AddHeaders:    &api.HeaderMap{},
		RemoveHeaders: &api.HeaderMap{},
	},
})
```

`api.OriginRequest` 字段说明：

| 字段          | 类型                 | 是否必选 | 说明                                                         |
| ------------- | -------------------- | -------- | ------------------------------------------------------------ |
| AddHeaders    | *api.HeaderMap       | 否       | 添加回源请求头。key 为 HTTP 头字段，一般为 HTTP 标准 Header，长度限制 128，也可以是自定义 Header；value 为该 header 的值，长度限制 1000，可以是常量也可以是变量。<br>**变量约束**：以 `$` 开始的子串必须符合 `${x}` 模式，合法变量为：<br>• `${uri}` 客户端请求的 URL 路径部分（不含查询参数）<br>• `${host}` 客户端请求 host 头部值<br>• `${scheme}` 客户端请求协议（`http` 或 `https`）<br>• `${request_uri}` 客户端请求路径和参数（含查询参数）<br>• `${jvip}` 节点 IP<br>• `${remote_addr}` 客户端 IP（存在代理时不准确）<br>• `${request_id}` 请求的唯一标识符<br>**典型非法值**：<br>• 变量不符合限制，如 `X-REQ-${url}`<br>• 含 `$` 但不符合 `${x}` 模式，如 `X-REQ-$uri`<br>注意：value 不支持 `$` 纯字符透传，如 `X-$` 非法。 |
| RemoveHeaders | *api.HeaderMap       | 否       | 删除回源请求头。key、value 的合法取值同 `AddHeaders`。       |

`api.HeaderMap` 本质是 `map[string]string`。服务端在「未配置任何头」时返回的是空数组 `[]` 而非空对象 `{}`，因此 SDK 为该类型定制了 JSON 编解码：解码时同时兼容 `[]` 和 `{}`，编码时空 map 会序列化为 `[]`。

注：最多设置 20 条 HTTP 回源请求头规则；不支持删除以 `ohc`、`baidu` 开头的回源请求头。

## 设置节点响应头配置

> 注意：本接口为全量更新，每次设置需带上希望保留的全部节点响应头，否则原有配置会被覆盖。

```go
cli := GetDefaultClient()
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	ClientResponse: &api.ClientResponse{
		AddHeaders: &api.HeaderMap{
			"test":  "${uri}",
			"test2": "11",
		},
		RemoveHeaders: &api.HeaderMap{
			"Content-Type": "",
		},
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

清空节点响应头配置：显式传入非 nil 的空 `HeaderMap`（序列化为 `[]`）。

```go
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	ClientResponse: &api.ClientResponse{
		AddHeaders:    &api.HeaderMap{},
		RemoveHeaders: &api.HeaderMap{},
	},
})
```

`api.ClientResponse` 字段说明：

| 字段          | 类型                 | 是否必选 | 说明                                                         |
| ------------- | -------------------- | -------- | ------------------------------------------------------------ |
| AddHeaders    | *api.HeaderMap       | 否       | 添加节点响应头。key 为 HTTP 头字段，一般为 HTTP 标准 Header，长度限制 128，也可以是自定义 Header；value 为该 header 的值，长度限制 1000，可以是常量也可以是变量。<br>**变量约束**：以 `$` 开始的子串必须符合 `${x}` 模式，合法变量为：<br>• `${uri}` 客户端请求的 URL 路径部分（不含查询参数）<br>• `${host}` 客户端请求 host 头部值<br>• `${scheme}` 客户端请求协议（`http` 或 `https`）<br>• `${request_uri}` 客户端请求路径和参数（含查询参数）<br>• `${jvip}` 节点 IP<br>• `${remote_addr}` 客户端 IP（存在代理时不准确）<br>• `${request_id}` 请求的唯一标识符<br>**典型非法值**：<br>• 变量不符合限制，如 `X-REQ-${url}`<br>• 含 `$` 但不符合 `${x}` 模式，如 `X-REQ-$uri`<br>注意：value 不支持 `$` 纯字符透传，如 `X-$` 非法。 |
| RemoveHeaders | *api.HeaderMap       | 否       | 删除节点响应头。key、value 的合法取值同 `AddHeaders`。       |

注：最多设置 20 条 HTTP 节点响应头规则。

## 设置浏览器缓存 TTL 配置

```go
cli := GetDefaultClient()
clientCacheTtlMode := "custom"
clientCacheTtl := 150

resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	ClientCacheTtl: &api.ClientCacheTtl{
		Mode: &clientCacheTtlMode,
		Ttl:  &clientCacheTtl,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.ClientCacheTtl` 字段说明：

| 字段 | 类型    | 是否必选 | 说明                                                                                          |
| ---- | ------- | -------- | --------------------------------------------------------------------------------------------- |
| Mode | *String | 是       | 工作模式。合法值：`follow`（遵循源站）、`no_cache`（不缓存）、`custom`（自定义缓存时间）。     |
| Ttl  | *Int    | 否       | 浏览器端自定义缓存时长，单位秒；仅当 `Mode` 为 `custom` 时有效。最小值 1，最大值 315360000（10 年）。 |

## 设置 webSocket 配置

```go
cli := GetDefaultClient()
webSocketEnabled := true
webSocketTimeout := 10

resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	WebSocket: &api.WebSocket{
		Enabled: &webSocketEnabled,
		Timeout: &webSocketTimeout,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`api.WebSocket` 字段说明：

| 字段    | 类型  | 是否必选 | 说明                                                       |
| ------- | ----- | -------- | ---------------------------------------------------------- |
| Enabled | *Bool | 是       | `true` 表示开启，`false` 表示关闭。                        |
| Timeout | *Int  | 否       | 最大连接超时时长，取值范围 1-300 秒；开启时必传，关闭时不传。 |

## 设置 SSL/TLS 安全配置

```go
cli := GetDefaultClient()
sslMode := "strong"
sslCipherList := "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256:TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:" +
	"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:" +
	"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
sslModePtr := &sslMode
sslCipherListPtr := &sslCipherList

resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	SslProtocols:  &[]string{"TLSv1.2", "TLSv1.3"},
	SslMode:       &sslModePtr,
	SslCipherList: &sslCipherListPtr,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

选择「全部加密算法套件」时无需传 `SslCipherList`：

```go
cli := GetDefaultClient()
sslMode := "all"
sslModePtr := &sslMode
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	SslProtocols: &[]string{"TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3"},
	SslMode:      &sslModePtr,
})
```

关闭 SSL/TLS 配置：三个字段都需要下发 `null`。`SslProtocols` 传指向 nil 切片的指针，`SslMode`/`SslCipherList` 传指向 nil 指针的指针。

```go
cli := GetDefaultClient()
var nilProtocols []string
var nilStr *string
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	SslProtocols:  &nilProtocols,
	SslMode:       &nilStr,
	SslCipherList: &nilStr,
})
```

SSL/TLS 相关字段说明：

| 字段          | 类型       | 是否必选 | 说明                                                                                                       |
| ------------- | ---------- | -------- | ---------------------------------------------------------------------------------------------------------- |
| SslProtocols  | *[]String  | 是       | SSL/TLS 版本。合法值：`TLSv1.0`、`TLSv1.1`、`TLSv1.2`、`TLSv1.3`，可传一个或多个，且**必须是连续版本**。      |
| SslMode       | **String   | 是       | 加密算法套件模式。合法值：`all`（全部套件）、`strong`（强套件，必须同时设置 TLSv1.2 与 TLSv1.3）、`custom`（自定义套件，必须同时设置 TLSv1.2 与 TLSv1.3）。 |
| SslCipherList | **String   | 否       | 加密算法套件列表，多个套件以英文冒号 `:` 分隔。`all` 模式无需传；`strong` 模式必须传齐 TLSv1.2 的 6 项强套件；`custom` 模式可选传 TLSv1.2 支持的套件。 |

注：`SslMode` 与 `SslCipherList` 是**双指针**类型。外层指针为 nil 时该字段不下发；外层非 nil 而内层为 nil 时下发 `null`（用于关闭配置）；内层非 nil 时下发其字符串值。这是因为关闭 SSL/TLS 配置要求显式下发 `null`，单层指针无法与「不设置」区分。

`SslCipherList` 可选值：

| TLS 协议 | 支持的加密算法套件 |
| -------- | ------------------ |
| TLSv1.0  | `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA`、`TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA`、`TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA`、`TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA`、`TLS_RSA_WITH_AES_128_CBC_SHA`、`TLS_RSA_WITH_AES_256_CBC_SHA` |
| TLSv1.2  | **强加密算法套件（`strong` 模式下 6 项缺一不可）**：`TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256`、`TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256`、`TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256`、`TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`、`TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`、`TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`<br>**其他加密算法套件**：`TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256`、`TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256`、`TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384`、`TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384`、`TLS_RSA_WITH_AES_128_GCM_SHA256`、`TLS_RSA_WITH_AES_256_GCM_SHA384`、`TLS_RSA_WITH_AES_128_CBC_SHA256`、`TLS_RSA_WITH_AES_256_CBC_SHA256`<br>`custom` 模式下强套件与其他套件均可自由选择。 |
| TLSv1.3  | `TLS_AES_256_GCM_SHA384`、`TLS_CHACHA20_POLY1305_SHA256`、`TLS_AES_128_GCM_SHA256`。`strong` 与 `custom` 模式下这 3 项默认支持，无需重复传参。 |

## 设置 OCSP 配置

```go
cli := GetDefaultClient()
ocsp := "ON"

resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	Ocsp: &ocsp,
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

`Ocsp` 字段说明：

| 字段 | 类型    | 是否必选 | 说明                                    |
| ---- | ------- | -------- | --------------------------------------- |
| Ocsp | *String | 是       | `ON` 开启 OCSP；`OFF` 关闭 OCSP。       |

## 设置用户 ip 获取配置

```go
cli := GetDefaultClient()
realIpEnabled := true
realIpName := "True-Client-Ip"

resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	RealIp: &api.RealIp{
		Enabled: &realIpEnabled,
		Name:    &realIpName,
	},
})
result, _ := json.Marshal(resp)
fmt.Printf("result: %s\n", result)
fmt.Printf("err:%+v\n", err)
```

关闭用户 ip 获取时不传 `Name`：

```go
realIpEnabled := false
resp, err := cli.SetSiteConfig("your_site.com", &api.SiteConfig{
	RealIp: &api.RealIp{
		Enabled: &realIpEnabled,
	},
})
```

`api.RealIp` 字段说明：

| 字段    | 类型    | 是否必选 | 说明                                                                       |
| ------- | ------- | -------- | -------------------------------------------------------------------------- |
| Enabled | *Bool   | 是       | `true` 表示开启，`false` 表示关闭。                                        |
| Name    | *String | 否       | 承载用户 IP 的头部名称，合法值：`True-Client-Ip`、`X-Real-IP`；开启时必传，关闭时不传。 |

## 查询站点配置

```go
cli := GetDefaultClient()
cfg, err := cli.GetSiteConfig("your_site.com")
if err != nil {
    fmt.Printf("err:%+v\n", err)
    return
}

// 只查询某一个配置项
if cfg.CacheTtl != nil {
    cacheTtlData, _ := json.Marshal(cfg.CacheTtl)
    fmt.Printf("cacheTtl: %s\n", cacheTtlData)
}

// 查询由多个字段组成的配置项（如「页面压缩」由 Compress + CompressMethodArray 共同组成）
if cfg.Compress != nil || cfg.CompressMethodArray != nil {
    compressData, _ := json.Marshal(struct {
        Compress            *string   `json:"compress,omitempty"`
        CompressMethodArray *[]string `json:"compressMethodArray,omitempty"`
    }{cfg.Compress, cfg.CompressMethodArray})
    fmt.Printf("compress: %s\n", compressData)
}
// 同理「强制 HTTPS」由 HttpToHttpsEnabled 和 HttpToHttpsCode 组成，可按相同方式聚合查询。
// 注意：当 HttpToHttpsEnabled = "OFF" 时，服务端不返回 HttpToHttpsCode
if cfg.HttpToHttpsEnabled != nil {
    httptohttpsData, _ := json.Marshal(struct {
        HttpToHttpsEnabled *string `json:"httpToHttpsEnabled,omitempty"`
        HttpToHttpsCode    *string `json:"httpToHttpsCode,omitempty"`
    }{cfg.HttpToHttpsEnabled, cfg.HttpToHttpsCode})
    fmt.Printf("httpToHttps: %s\n", httptohttpsData)
}

// 查询全部配置
allData, _ := json.Marshal(cfg)
fmt.Printf("SiteConfig: %s\n", allData)

// cfg.CacheTtl 为该站点当前的节点缓存配置；
// cfg.CacheKey 为该站点当前的查询字符串配置;
// cfg.OfflineMode 为该站点当前的离线模式配置;
// cfg.HttpToHttpsEnabled 和 cfg.HttpToHttpsCode 为该站点当前的强制 HTTPS 配置;
// cfg.Hsts 为该站点当前的 HSTS 配置;
// cfg.Http2Disable 为该站点当前的 HTTP2 配置;
// cfg.Http3 为该站点当前的 HTTP3 配置;
// cfg.ClientMaxBodySize 为该站点当前的最大上传大小配置;
// cfg.Compress 和 cfg.CompressMethodArray 为该站点当前的页面压缩配置;
// cfg.Isa 为该站点当前的智能加速配置;
// cfg.CacheCodeTtl 为该站点当前的状态码缓存配置;
// cfg.GrpcOrigin 为该站点当前的 gRPC 回源配置;
// cfg.Http2Origin 为该站点当前的 HTTP2 回源配置;
// cfg.PageRules 为该站点当前的规则引擎配置;
// cfg.OriginRequest 为该站点当前的回源请求头配置;
// cfg.ClientResponse 为该站点当前的节点响应头配置;
// cfg.ClientCacheTtl 为该站点当前的浏览器缓存 TTL 配置;
// cfg.WebSocket 为该站点当前的 webSocket 配置;
// cfg.SslProtocols、cfg.SslMode 和 cfg.SslCipherList 为该站点当前的 SSL/TLS 安全配置;
// cfg.Ocsp 为该站点当前的 OCSP 配置;
// cfg.RealIp 为该站点当前的用户 ip 获取配置;
```


