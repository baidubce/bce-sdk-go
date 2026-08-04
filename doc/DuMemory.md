# DuMemory服务

# 概述

本文档主要介绍 DuMemory（云端记忆）GO SDK 的使用。DuMemory 是基于向量数据库 VectorDB 提供的长期记忆服务，可用于为大模型应用提供持久化的记忆读写、心智模型、指令、文档与实体关系等能力。对外暴露与 [接口速查表](https://cloud.baidu.com/doc/VDB/s/cmpl4uayz) 完全对齐的 API。

> 注意：DuMemory 的鉴权方式不同于其它 BCE 服务，**不使用 AK/SK 签名**，而是通过 HTTP Bearer Token（API Key）进行认证。

# 初始化

## 确认 Endpoint

DuMemory 的服务端点由用户在控制台或自建部署中获取，通常形如：

- 公有云端点：`https://cloud.memory.<region>.baidubce.com/api`
- 自建/本地部署：`http://127.0.0.1:8888`

API 支持 HTTP 和 HTTPS 两种调用方式。生产环境建议使用 HTTPS。

## 获取 API Key

通过 [百度智能云控制台](https://console.bce.baidu.com) 创建并获取 DuMemory 的 API Key。所有请求会通过下面的请求头传递：


## 云记忆 DuMemory API 参考
本文档参照百度智能云接口文档结构，对 services/DuMemory/api 包封装的全部接口逐一说明。每个接口均包括：接口说明、请求 URI、请求头域、请求参数、返回头域、返回参数、请求示例、返回示例，以及对应的 SDK 调用方法。
服务端实现基于DuMemory HTTP API（OpenAPI 0.6.x / 0.8.x），SDK 是其上的 Go 封装。所有 /v1/default/banks/{bankId}/... 路径中的 default 表示当前默认命名空间，bankId 为记忆库 ID。
---

## 公共说明

### 公共请求头域

| 参数名称 | 参数类型 | 是否必须 | 示例值 | 描述 |
| --- | --- | --- | --- | --- |
| Host | String | 是 | cloud.memory.bj.baidubce.com | 服务域名 |
| Authorization | String | 是 | Bearer bce-v3/ALTAK-xxx/xxx | Bearer Token；通过 `DuMemory.New(baseURL, apiKey)` 自动注入 |
| Content-Type | String | 否 | application/json | 携带请求体的接口必填 |
| Accept | String | 否 | application/json | 期望返回 JSON |

### 公共返回头域

| 参数名称 | 参数类型 | 示例值 | 描述 |
| --- | --- | --- | --- |
| Content-Type | String | application/json | 返回体类型 |
| Date | String | Mon, 09 Jun 2026 02:30:00 GMT | 服务端时间 |

### SDK 客户端构造

```go
import (
    "context"
    "time"

    DuMemory "github.com/baidubce/bce-sdk-go/services/DuMemory/api"
)

client := DuMemory.New("https://cloud.memory.bj.baidubce.com/api", "<API_KEY>")
// 或带超时
client = DuMemory.NewWithTimeout("https://cloud.memory.bj.baidubce.com/api", "<API_KEY>", 2*time.Minute)
```

### 接口总览

| 模块 | 方法 | 路径 | SDK 方法 |
| --- | --- | --- | --- |
| 监控 | GET | /health | `Health` |
| 监控 | GET | /version | `Version` |
| 记忆库 | GET | /v1/default/banks | `ListBanks` |
| 记忆库 | POST | /v1/default/banks/{bankId} | `CreateBank` |
| 记忆库 | GET | /v1/default/banks/{bankId}/profile | `GetBank` |
| 记忆库 | DELETE | /v1/default/banks/{bankId} | `DeleteBank` |
| 记忆库 | GET | /v1/default/banks/{bankId}/config | `GetBankConfig` |
| 记忆库 | PATCH | /v1/default/banks/{bankId}/config | `UpdateBankConfig` |
| 记忆库 | GET | /v1/default/banks/{bankId}/stats | `GetBankStats` |
| 记忆库 | POST | /v1/default/banks/{bankId}/consolidate | `ConsolidateBank` |
| 记忆 | POST | /v1/default/banks/{bankId}/memories | `Retain` / `RetainAsync` |
| 记忆 | POST | /v1/default/banks/{bankId}/memories/recall | `Recall` |
| 记忆 | POST | /v1/default/banks/{bankId}/reflect | `Reflect` |
| 记忆 | GET | /v1/default/banks/{bankId}/memories/list | `ListMemories` |
| 实体 | GET | /v1/default/banks/{bankId}/entities | `ListEntities` |
| 实体 | GET | /v1/default/banks/{bankId}/entities/graph | `EntityGraph` |
| 文档 | GET | /v1/default/banks/{bankId}/documents | `ListDocuments` |
| 文档 | GET | /v1/default/banks/{bankId}/documents/{documentId} | `GetDocument` |
| 文档 | PATCH | /v1/default/banks/{bankId}/documents/{documentId} | `UpdateDocument` |
| 文档 | DELETE | /v1/default/banks/{bankId}/documents/{documentId} | `DeleteDocument` |
| 文档 | GET | /v1/default/banks/{bankId}/documents/{documentId}/chunks | `ListDocumentChunks` |
| 心智模型 | GET | /v1/default/banks/{bankId}/mental-models | `ListMentalModels` |
| 心智模型 | POST | /v1/default/banks/{bankId}/mental-models | `CreateMentalModel` |
| 心智模型 | GET | /v1/default/banks/{bankId}/mental-models/{modelId} | `GetMentalModel` |
| 心智模型 | PATCH | /v1/default/banks/{bankId}/mental-models/{modelId} | `UpdateMentalModel` |
| 心智模型 | DELETE | /v1/default/banks/{bankId}/mental-models/{modelId} | `DeleteMentalModel` |
| 心智模型 | POST | /v1/default/banks/{bankId}/mental-models/{modelId}/refresh | `RefreshMentalModel` |
| 指令 | GET | /v1/default/banks/{bankId}/directives | `ListDirectives` |
| 指令 | POST | /v1/default/banks/{bankId}/directives | `CreateDirective` |
| 指令 | GET | /v1/default/banks/{bankId}/directives/{directiveId} | `GetDirective` |
| 指令 | PATCH | /v1/default/banks/{bankId}/directives/{directiveId} | `UpdateDirective` |
| 指令 | DELETE | /v1/default/banks/{bankId}/directives/{directiveId} | `DeleteDirective` |
| 操作 | GET | /v1/default/banks/{bankId}/operations | `ListOperations` |
| 操作 | GET | /v1/default/banks/{bankId}/operations/{operationId} | `GetOperationStatus` |
| 操作 | DELETE | /v1/default/banks/{bankId}/operations/{operationId} | `CancelOperation` |
| 文件 | POST | /v1/default/banks/{bankId}/files/retain | `FilesRetain` |
| 标签作用域扩展 | POST | /v1/default/banks/{bankId}/memories | `RetainWithScope` |
| 标签作用域扩展 | POST | /v1/default/banks/{bankId}/memories/recall | `RecallWithScope` |
| 标签作用域扩展 | POST | /v1/default/banks/{bankId}/reflect | `ReflectWithScope` |
| 标签作用域扩展 | GET | /v1/default/banks/{bankId}/memories/{memoryId} | `GetMemoryWithScope` |
| 标签作用域扩展 | GET | /v1/default/banks/{bankId}/tags | `ListTagsWithScope` |
| 标签作用域扩展 | GET | /v1/default/banks/{bankId}/memories/list | `ListMemoriesWithScope` |
| 标签作用域扩展 | GET | /v1/default/banks/{bankId}/documents | `ListDocumentsWithScope` |
| 标签作用域扩展 | PATCH | /v1/default/banks/{bankId}/documents/{documentId} | `UpdateDocumentTagsWithScope` |
| 标签作用域扩展 | GET | /v1/default/banks/{bankId}/directives | `ListDirectivesWithScope` |
| 标签作用域扩展 | POST | /v1/default/banks/{bankId}/directives | `CreateDirectiveWithScope` |
| 标签作用域扩展 | PATCH | /v1/default/banks/{bankId}/directives/{directiveId} | `UpdateDirectiveWithScope` |
| 标签作用域扩展 | GET | /v1/default/banks/{bankId}/mental-models | `ListMentalModelsWithScope` |
| 标签作用域扩展 | POST | /v1/default/banks/{bankId}/mental-models | `CreateMentalModelWithScope` |
| 标签作用域扩展 | PATCH | /v1/default/banks/{bankId}/mental-models/{modelId} | `UpdateMentalModelWithScope` |

---

## 一、监控

### 1.1 健康检查 Health

**接口说明**

检查服务是否存活；不需要鉴权。

**请求 URI**

```
GET /health
Host: cloud.memory.bj.baidubce.com
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

无。

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| status | String | 健康状态，例如 `healthy` |
| database | String | 数据库连接状态，例如 `connected` |

**请求示例**

```
GET /health
Host: cloud.memory.bj.baidubce.com
```

**返回示例**

```json
{ "status": "healthy", "database": "connected" }
```

**SDK 方法**

```go
out, err := client.Health(context.Background())
```

---

### 1.2 服务版本 Version

**接口说明**

返回服务端 API 版本与已启用的特性开关，需要鉴权。

**请求 URI**

```
GET /version
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

无。

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| api_version | String | 服务端版本号，如 `0.6.2` |
| features | Object | 特性开关键值对，键为特性名，值为布尔；具体键随服务端版本变化 |

**请求示例**

```
GET /version
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "api_version": "0.6.2",
  "features": {
    "observations": true,
    "mcp": true,
    "worker": true,
    "bank_config_api": true,
    "file_upload_api": true
  }
}
```

**SDK 方法**

```go
info, err := client.Version(context.Background())
// info.APIVersion / info.Features["observations"]
```

---

## 二、记忆库 Banks

### 2.1 列出记忆库 ListBanks

**接口说明**

列出当前账号下所有记忆库（bank）。

**请求 URI**

```
GET /v1/default/banks
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

无。

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| banks | Array<BankListItem> | 记忆库列表，每项包含 `bank_id`、`name` 等基础信息 |

**请求示例**

```
GET /v1/default/banks
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "banks": [
    { "bank_id": "demo-bank", "name": "Demo" }
  ]
}
```

**SDK 方法**

```go
out, err := client.ListBanks(context.Background())
```

---

### 2.2 创建/更新记忆库 CreateBank

**接口说明**

按 `bankId` 创建记忆库；若已存在则按请求体更新元信息（名字、使命、人格特质等）。

**请求 URI**

```
POST /v1/default/banks/{bankId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| name | String | 否 | RequestBody | 显示名称 |
| disposition | Object | 否 | RequestBody | 综合人格特质 |
| disposition_skepticism | Integer | 否 | RequestBody | 怀疑性 0-100 |
| disposition_literalism | Integer | 否 | RequestBody | 字面性 0-100 |
| disposition_empathy | Integer | 否 | RequestBody | 共情性 0-100 |
| mission | String | 否 | RequestBody | Agent 使命 |
| background | String | 否 | RequestBody | 背景描述 |
| reflect_mission | String | 否 | RequestBody | reflect 自定义使命 |
| retain_mission | String | 否 | RequestBody | retain 自定义使命 |
| retain_extraction_mode | String | 否 | RequestBody | 抽取模式 |
| retain_custom_instructions | String | 否 | RequestBody | retain 自定义指令 |
| retain_chunk_size | Integer | 否 | RequestBody | 切片大小 |
| enable_observations | Boolean | 否 | RequestBody | 是否启用观察 |
| observations_mission | String | 否 | RequestBody | 观察使命 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| bank_id | String | 记忆库 ID |
| name | String | 名称 |
| disposition | Object | 人格特质 |
| mission | String | 使命 |
| background | String | 背景，可空 |

**请求示例**

```
POST /v1/default/banks/demo-bank
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "name": "Demo Bank",
  "mission": "Help me remember everything"
}
```

**返回示例**

```json
{
  "bank_id": "demo-bank",
  "name": "Demo Bank",
  "disposition": {},
  "mission": "Help me remember everything"
}
```

**SDK 方法**

```go
req := DuMemory.CreateBankRequest{}
req.SetName("Demo Bank")
req.SetMission("Help me remember everything")
out, err := client.CreateBank(ctx, "demo-bank", req)
```

---

### 2.3 查看记忆库 GetBank

**接口说明**

按 `bankId` 获取记忆库的 profile（名称、人格、使命、背景）。

**请求 URI**

```
GET /v1/default/banks/{bankId}/profile
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

参见 2.2 返回参数。

**请求示例**

```
GET /v1/default/banks/demo-bank/profile
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "bank_id": "demo-bank",
  "name": "Demo Bank",
  "disposition": {},
  "mission": "Help me remember everything",
  "background": null
}
```

**SDK 方法**

```go
out, err := client.GetBank(ctx, "demo-bank")
```

---

### 2.4 删除记忆库 DeleteBank

**接口说明**

删除指定记忆库及其下全部数据，操作不可恢复。

**请求 URI**

```
DELETE /v1/default/banks/{bankId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| success | Boolean | 是否成功 |
| message | String | 信息描述 |

**请求示例**

```
DELETE /v1/default/banks/demo-bank
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "success": true, "message": "deleted" }
```

**SDK 方法**

```go
out, err := client.DeleteBank(ctx, "demo-bank")
```

---

### 2.5 获取记忆库配置 GetBankConfig

**接口说明**

获取记忆库的完整配置（合并所有继承层级）和当前 bank 自身的覆盖项。

**请求 URI**

```
GET /v1/default/banks/{bankId}/config
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| bank_id | String | 记忆库 ID |
| config | Object | 合并所有继承后的最终配置 |
| overrides | Object | 当前 bank 自有的覆盖项 |

**请求示例**

```
GET /v1/default/banks/demo-bank/config
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "bank_id": "demo-bank",
  "config": { "llm_provider": "openai" },
  "overrides": {}
}
```

**SDK 方法**

```go
out, err := client.GetBankConfig(ctx, "demo-bank")
```

---

### 2.6 更新记忆库配置 UpdateBankConfig

**接口说明**

按需覆盖 bank 级配置。Key 支持 Python 字段名（如 `llm_provider`）或环境变量名（如 `HINDSIGHT_API_LLM_PROVIDER`），仅可覆盖支持层级化的字段。

**请求 URI**

```
PATCH /v1/default/banks/{bankId}/config
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| updates | Object | 是 | RequestBody | 要覆盖的配置键值对 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

参见 2.5 返回参数。

**请求示例**

```
PATCH /v1/default/banks/demo-bank/config
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "updates": { "llm_provider": "openai" } }
```

**返回示例**

```json
{
  "bank_id": "demo-bank",
  "config": { "llm_provider": "openai" },
  "overrides": { "llm_provider": "openai" }
}
```

**SDK 方法**

```go
upd := DuMemory.NewBankConfigUpdate(map[string]interface{}{
    "llm_provider": "openai",
})
out, err := client.UpdateBankConfig(ctx, "demo-bank", *upd)
```

---

### 2.7 记忆库统计 GetBankStats

**接口说明**

返回 bank 级的节点/链接/文档/操作等统计指标。

**请求 URI**

```
GET /v1/default/banks/{bankId}/stats
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| bank_id | String | 记忆库 ID |
| total_nodes | Integer | 节点数 |
| total_links | Integer | 关系数 |
| total_documents | Integer | 文档数 |
| nodes_by_fact_type | Object<String,Integer> | 各类节点计数 |
| links_by_link_type | Object<String,Integer> | 各类关系计数 |
| links_by_fact_type | Object<String,Integer> | 关系按事实类型计数 |
| links_breakdown | Object | 关系明细 |
| pending_operations | Integer | 排队操作数 |
| failed_operations | Integer | 失败操作数 |
| operations_by_status | Object<String,Integer> | 异步操作分状态计数（可选） |
| last_consolidated_at | String | 上次整合时间，可空 |
| pending_consolidation | Integer | 待整合的源记忆数（可选） |
| failed_consolidation | Integer | 整合失败的源记忆数（可选） |
| total_observations | Integer | 总观察数（可选） |

**请求示例**

```
GET /v1/default/banks/demo-bank/stats
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "bank_id": "demo-bank",
  "total_nodes": 124,
  "total_links": 56,
  "total_documents": 8,
  "nodes_by_fact_type": { "Person": 12 },
  "links_by_link_type": { "knows": 6 },
  "links_by_fact_type": {},
  "links_breakdown": {},
  "pending_operations": 0,
  "failed_operations": 0
}
```

**SDK 方法**

```go
out, err := client.GetBankStats(ctx, "demo-bank")
```

---

### 2.8 触发整合 ConsolidateBank

**接口说明**

手动触发一次记忆整合（observations）。整合是异步任务，接口返回任务 ID。

**请求 URI**

```
POST /v1/default/banks/{bankId}/consolidate
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| body | Object | 否 | RequestBody | 服务端 0.6.x 版本可省略请求体；0.8.x 起支持可选过滤参数 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| operation_id | String | 异步任务 ID，可用 `ListOperations` 跟踪 |
| deduplicated | Boolean | 是否复用了同一个等待中的任务（可选） |

**请求示例**

```
POST /v1/default/banks/demo-bank/consolidate
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "operation_id": "op-abc-123", "deduplicated": false }
```

**SDK 方法**

```go
out, err := client.ConsolidateBank(ctx, "demo-bank", nil)
```

---

## 三、记忆 Memory

### 3.1 写入记忆 Retain / RetainAsync

**接口说明**

向指定 bank 写入一组记忆条目，可选同步或异步处理。`RetainAsync` 会在 SDK 内强制设置请求体 `async=true`。

**请求 URI**

```
POST /v1/default/banks/{bankId}/memories
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| items | Array<MemoryItem> | 是 | RequestBody | 记忆条目数组 |
| items[].content | String | 是 | RequestBody | 记忆正文 |
| items[].timestamp | String | 否 | RequestBody | ISO 时间戳 |
| items[].context | String | 否 | RequestBody | 上下文 |
| items[].metadata | Object<String,String> | 否 | RequestBody | 元数据 |
| items[].document_id | String | 否 | RequestBody | 关联的文档 ID |
| items[].tags | Array<String> | 否 | RequestBody | 标签 |
| items[].strategy | String | 否 | RequestBody | 写入策略 |
| items[].update_mode | String | 否 | RequestBody | 更新模式 |
| async | Boolean | 否 | RequestBody | 异步处理标记 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| success | Boolean | 是否成功 |
| bank_id | String | 记忆库 ID |
| items_count | Integer | 写入条目数 |
| async | Boolean | 是否异步处理 |
| operation_id | String | 异步主任务 ID（可空） |
| operation_ids | Array<String> | 异步任务 ID 列表（可空） |
| usage | Object | Token 用量（可空） |

**请求示例**

```
POST /v1/default/banks/demo-bank/memories
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "items": [
    { "content": "今天和小明在楼下喝了咖啡。", "tags": ["life"] }
  ]
}
```

**返回示例**

```json
{
  "success": true,
  "bank_id": "demo-bank",
  "items_count": 1,
  "async": false
}
```

**SDK 方法**

```go
items := []DuMemory.MemoryItem{*DuMemory.NewMemoryItem("今天和小明在楼下喝了咖啡。")}
out, err := client.Retain(ctx, "demo-bank", *DuMemory.NewRetainRequest(items))
// 异步:
out, err = client.RetainAsync(ctx, "demo-bank", *DuMemory.NewRetainRequest(items))
```

---

### 3.2 召回记忆 Recall

**接口说明**

按自然语言 query 检索记忆，可附带类型、标签、tag_groups 等过滤条件。

**请求 URI**

```
POST /v1/default/banks/{bankId}/memories/recall
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| query | String | 是 | RequestBody | 查询语句 |
| types | Array<String> | 否 | RequestBody | 节点/事实类型过滤 |
| budget | Object | 否 | RequestBody | 检索预算 |
| max_tokens | Integer | 否 | RequestBody | 返回 token 上限 |
| trace | Boolean | 否 | RequestBody | 是否返回检索 trace |
| query_timestamp | String | 否 | RequestBody | 时间相关查询的基准时间 |
| include | Object | 否 | RequestBody | 控制是否包含实体、源记忆等 |
| tags | Array<String> | 否 | RequestBody | 标签过滤 |
| tags_match | String | 否 | RequestBody | 标签匹配模式：`any` `all` `any_strict` `all_strict` |
| tag_groups | Array<Object> | 否 | RequestBody | 多组标签条件 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| results | Array<RecallResult> | 命中结果 |
| trace | Object | 检索 trace（可选） |
| entities | Object<String,EntityState> | 关联实体（可选） |
| chunks | Object<String,ChunkData> | 关联文档片段（可选） |
| source_facts | Object<String,RecallResult> | 源事实（可选） |

**请求示例**

```
POST /v1/default/banks/demo-bank/memories/recall
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "query": "我和小明喝过几次咖啡？" }
```

**返回示例**

```json
{
  "results": [
    { "id": "m-1", "content": "今天和小明在楼下喝了咖啡。", "score": 0.92 }
  ]
}
```

**SDK 方法**

```go
out, err := client.Recall(ctx, "demo-bank",
    *DuMemory.NewRecallRequest("我和小明喝过几次咖啡？"))
```

---

### 3.3 反思 Reflect

**接口说明**

基于召回记忆生成结构化或 markdown 形式的回答；支持指定 fact_types、tag_groups 等过滤项。

**请求 URI**

```
POST /v1/default/banks/{bankId}/reflect
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| query | String | 是 | RequestBody | 反思问题 |
| budget | Object | 否 | RequestBody | 检索预算 |
| context | String | 否 | RequestBody | 附加上下文 |
| max_tokens | Integer | 否 | RequestBody | 返回 token 上限 |
| include | Object | 否 | RequestBody | 是否包含 trace、引用等 |
| response_schema | Object | 否 | RequestBody | 期望的 JSON 输出 schema |
| tags | Array<String> | 否 | RequestBody | 标签过滤 |
| tags_match | String | 否 | RequestBody | 标签匹配模式 |
| tag_groups | Array<Object> | 否 | RequestBody | 多组标签条件 |
| fact_types | Array<String> | 否 | RequestBody | 事实类型过滤 |
| exclude_mental_models | Boolean | 否 | RequestBody | 是否排除全部心智模型 |
| exclude_mental_model_ids | Array<String> | 否 | RequestBody | 指定要排除的心智模型 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| text | String | Markdown 文本 |
| based_on | Object | 引用的记忆（可选） |
| structured_output | Object | 结构化输出（按 response_schema） |
| usage | Object | Token 用量（可选） |
| trace | Object | 推理 trace（可选） |

**请求示例**

```
POST /v1/default/banks/demo-bank/reflect
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "query": "总结我和小明的关系" }
```

**返回示例**

```json
{ "text": "你和小明经常一起喝咖啡，关系亲近。" }
```

**SDK 方法**

```go
out, err := client.Reflect(ctx, "demo-bank",
    *DuMemory.NewReflectRequest("总结我和小明的关系"))
```

---

### 3.4 列出记忆 ListMemories

**接口说明**

分页列出 bank 下的记忆单元，支持按类型、关键词、整合状态、记忆状态、来源文档、关联实体和 tags 过滤。

**请求 URI**

```
GET /v1/default/banks/{bankId}/memories/list?type=&q=&consolidation_state=&state=&document_id=&entity_id=&tags=&tags_match=&limit=&offset=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| type | String | 否 | Query 参数 | 节点/事实类型 |
| q | String | 否 | Query 参数 | 关键词模糊匹配 |
| consolidation_state | String | 否 | Query 参数 | 整合状态过滤 |
| state | String | 否 | Query 参数 | 记忆状态过滤 |
| document_id | String | 否 | Query 参数 | 来源文档 ID |
| entity_id | String | 否 | Query 参数 | 关联实体 ID |
| tags | Array<String> | 否 | Query 参数 | 标签过滤，可传多个 |
| tags_match | String | 否 | Query 参数 | 标签匹配模式：`any`、`all`、`any_strict`、`all_strict`、`exact` |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| items | Array<Object> | 记忆单元列表 |
| total | Integer | 总条数 |
| limit | Integer | 分页大小 |
| offset | Integer | 偏移量 |

**请求示例**

```
GET /v1/default/banks/demo-bank/memories/list?state=active&document_id=doc-1&entity_id=entity-1&tags=topic:coffee&tags_match=all_strict&limit=20&offset=0
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [], "total": 0, "limit": 20, "offset": 0 }
```

**SDK 方法**

```go
out, err := client.ListMemories(ctx, "demo-bank", DuMemory.ListMemoriesOptions{
    State:      "active",
    DocumentID: "doc-1",
    EntityID:   "entity-1",
    Tags:       []string{"topic:coffee"},
    TagsMatch:  "all_strict",
    Limit:      20,
})
```

---

## 四、实体 Entities

### 4.1 列出实体 ListEntities

**接口说明**

分页列出 bank 内的实体（人/物/概念）。

**请求 URI**

```
GET /v1/default/banks/{bankId}/entities?limit=&offset=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| items | Array<EntityListItem> | 实体列表 |
| total | Integer | 总数 |
| limit | Integer | 分页大小 |
| offset | Integer | 偏移量 |

**请求示例**

```
GET /v1/default/banks/demo-bank/entities?limit=10
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [], "total": 0, "limit": 10, "offset": 0 }
```

**SDK 方法**

```go
out, err := client.ListEntities(ctx, "demo-bank", 10, 0)
```

---

### 4.2 实体关系图 EntityGraph

**接口说明**

返回 bank 内实体的图结构（节点 + 边），便于可视化关系网络。

**请求 URI**

```
GET /v1/default/banks/{bankId}/entities/graph?limit=&min_count=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| limit | Integer | 否 | Query 参数 | 节点数量上限 |
| min_count | Integer | 否 | Query 参数 | 入度下限 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| nodes | Array<Object> | 节点 |
| edges | Array<Object> | 边 |
| total_entities | Integer | 实体总数 |
| total_edges | Integer | 边总数 |
| limit | Integer | 节点数上限 |

**请求示例**

```
GET /v1/default/banks/demo-bank/entities/graph?limit=50&min_count=2
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "nodes": [], "edges": [], "total_entities": 0, "total_edges": 0, "limit": 50 }
```

**SDK 方法**

```go
out, err := client.EntityGraph(ctx, "demo-bank", 50, 2)
```

---

## 五、文档 Documents

### 5.1 列出文档 ListDocuments

**接口说明**

分页列出 bank 下的文档，支持关键词、标签过滤。

**请求 URI**

```
GET /v1/default/banks/{bankId}/documents?q=&tags=&tags_match=&limit=&offset=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| q | String | 否 | Query 参数 | 关键词 |
| tags | Array<String> | 否 | Query 参数 | 标签过滤（多次出现） |
| tags_match | String | 否 | Query 参数 | 标签匹配模式 |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| items | Array<Object> | 文档列表 |
| total | Integer | 总数 |
| limit | Integer | 分页大小 |
| offset | Integer | 偏移量 |

**请求示例**

```
GET /v1/default/banks/demo-bank/documents?limit=10
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [], "total": 0, "limit": 10, "offset": 0 }
```

**SDK 方法**

```go
out, err := client.ListDocuments(ctx, "demo-bank",
    DuMemory.ListDocumentsOptions{Limit: 10})
```

---

### 5.2 获取文档 GetDocument

**接口说明**

按文档 ID 获取详情，包括原文、tags、关联节点统计、retain 参数等。

**请求 URI**

```
GET /v1/default/banks/{bankId}/documents/{documentId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| documentId | String | 是 | URL 参数 | 文档 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| id | String | 文档 ID |
| bank_id | String | 记忆库 ID |
| original_text | String | 原文 |
| content_hash | String | 内容哈希 |
| created_at | String | 创建时间 |
| updated_at | String | 更新时间 |
| memory_unit_count | Integer | 派生的记忆单元数 |
| nodes_by_fact_type | Object<String,Integer> | 节点按事实类型计数 |
| tags | Array<String> | 标签 |
| document_metadata | Object | 文档元数据 |
| retain_params | Object | retain 时使用的参数 |

**请求示例**

```
GET /v1/default/banks/demo-bank/documents/doc-123
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "id": "doc-123",
  "bank_id": "demo-bank",
  "original_text": "...",
  "content_hash": "abc",
  "created_at": "2026-06-01T10:00:00Z",
  "updated_at": "2026-06-01T10:00:00Z",
  "memory_unit_count": 5
}
```

**SDK 方法**

```go
out, err := client.GetDocument(ctx, "demo-bank", "doc-123")
```

---

### 5.3 更新文档 UpdateDocument

**接口说明**

仅支持修改文档的 tags。

**请求 URI**

```
PATCH /v1/default/banks/{bankId}/documents/{documentId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| documentId | String | 是 | URL 参数 | 文档 ID |
| tags | Array<String> | 否 | RequestBody | 全量替换 tags |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| success | Boolean | 是否成功 |

**请求示例**

```
PATCH /v1/default/banks/demo-bank/documents/doc-123
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "tags": ["personal", "coffee"] }
```

**返回示例**

```json
{ "success": true }
```

**SDK 方法**

```go
req := *DuMemory.NewUpdateDocumentRequest()
req.Tags = []string{"personal", "coffee"}
out, err := client.UpdateDocument(ctx, "demo-bank", "doc-123", req)
```

---

### 5.4 删除文档 DeleteDocument

**接口说明**

删除文档及其派生的记忆单元。

**请求 URI**

```
DELETE /v1/default/banks/{bankId}/documents/{documentId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| documentId | String | 是 | URL 参数 | 文档 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| success | Boolean | 是否成功 |
| message | String | 信息描述 |
| document_id | String | 已删文档 ID |
| memory_units_deleted | Integer | 同时删除的记忆单元数 |

**请求示例**

```
DELETE /v1/default/banks/demo-bank/documents/doc-123
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "success": true, "message": "deleted", "document_id": "doc-123", "memory_units_deleted": 5 }
```

**SDK 方法**

```go
out, err := client.DeleteDocument(ctx, "demo-bank", "doc-123")
```

---

### 5.5 列出文档分块 ListDocumentChunks

**接口说明**

分页列出文档的分块（chunk）。

**请求 URI**

```
GET /v1/default/banks/{bankId}/documents/{documentId}/chunks?limit=&offset=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| documentId | String | 是 | URL 参数 | 文档 ID |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| items | Array<ChunkResponse> | 分块列表（chunk_id / chunk_index / chunk_text / created_at） |
| total | Integer | 总数 |
| limit | Integer | 分页大小 |
| offset | Integer | 偏移量 |

**请求示例**

```
GET /v1/default/banks/demo-bank/documents/doc-123/chunks?limit=20
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [], "total": 0, "limit": 20, "offset": 0 }
```

**SDK 方法**

```go
out, err := client.ListDocumentChunks(ctx, "demo-bank", "doc-123", 20, 0)
```

---

## 六、心智模型 Mental Models

### 6.1 列出心智模型 ListMentalModels

**接口说明**

列出 bank 下的心智模型。

**请求 URI**

```
GET /v1/default/banks/{bankId}/mental-models?tags=&tags_match=&detail=&limit=&offset=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| tags | Array<String> | 否 | Query 参数 | 标签过滤 |
| tags_match | String | 否 | Query 参数 | 标签匹配模式 |
| detail | String | 否 | Query 参数 | 详情级别（如 `summary`/`full`） |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| items | Array<MentalModelResponse> | 心智模型列表 |

**请求示例**

```
GET /v1/default/banks/demo-bank/mental-models?limit=20
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [] }
```

**SDK 方法**

```go
out, err := client.ListMentalModels(ctx, "demo-bank",
    DuMemory.ListMentalModelsOptions{Limit: 20})
```

---

### 6.2 创建心智模型 CreateMentalModel

**接口说明**

创建一个由 `source_query` 驱动生成的心智模型。

**请求 URI**

```
POST /v1/default/banks/{bankId}/mental-models
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| id | String | 否 | RequestBody | 自定义 ID |
| name | String | 是 | RequestBody | 名称 |
| source_query | String | 是 | RequestBody | 生成内容用的查询语句 |
| tags | Array<String> | 否 | RequestBody | 标签 |
| max_tokens | Integer | 否 | RequestBody | 生成上限 |
| trigger | Object | 否 | RequestBody | 触发条件 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| 同 GetMentalModel 返回 | - | 返回创建后的模型详情 |

**请求示例**

```
POST /v1/default/banks/demo-bank/mental-models
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "name": "User Profile", "source_query": "Who is the user?" }
```

**返回示例**

```json
{ "id": "mm-1", "bank_id": "demo-bank", "name": "User Profile" }
```

**SDK 方法**

```go
out, err := client.CreateMentalModel(ctx, "demo-bank",
    *DuMemory.NewCreateMentalModelRequest("User Profile", "Who is the user?"))
```

---

### 6.3 查看心智模型 GetMentalModel

**接口说明**

按 ID 获取心智模型详情，包括内容、最近一次刷新时间等。

**请求 URI**

```
GET /v1/default/banks/{bankId}/mental-models/{modelId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| modelId | String | 是 | URL 参数 | 心智模型 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| id | String | 模型 ID |
| bank_id | String | 记忆库 ID |
| name | String | 名称 |
| source_query | String | 源查询 |
| content | String | 已生成的内容 |
| tags | Array<String> | 标签 |
| max_tokens | Integer | 生成上限 |
| trigger | Object | 触发器 |
| last_refreshed_at | String | 上次刷新时间 |
| created_at | String | 创建时间 |
| reflect_response | Object | 最近一次 reflect 结果 |
| is_stale | Boolean | 是否过期 |

**请求示例**

```
GET /v1/default/banks/demo-bank/mental-models/mm-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "id": "mm-1", "bank_id": "demo-bank", "name": "User Profile" }
```

**SDK 方法**

```go
out, err := client.GetMentalModel(ctx, "demo-bank", "mm-1")
```

---

### 6.4 更新心智模型 UpdateMentalModel

**接口说明**

修改心智模型的名称、查询、触发条件等。

**请求 URI**

```
PATCH /v1/default/banks/{bankId}/mental-models/{modelId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| modelId | String | 是 | URL 参数 | 心智模型 ID |
| name | String | 否 | RequestBody | 新名称 |
| source_query | String | 否 | RequestBody | 新查询 |
| max_tokens | Integer | 否 | RequestBody | 生成上限 |
| tags | Array<String> | 否 | RequestBody | 标签 |
| trigger | Object | 否 | RequestBody | 触发条件 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

参见 6.3。

**请求示例**

```
PATCH /v1/default/banks/demo-bank/mental-models/mm-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "name": "User Profile v2" }
```

**返回示例**

```json
{ "id": "mm-1", "bank_id": "demo-bank", "name": "User Profile v2" }
```

**SDK 方法**

```go
req := DuMemory.NewUpdateMentalModelRequest()
req.SetName("User Profile v2")
out, err := client.UpdateMentalModel(ctx, "demo-bank", "mm-1", *req)
```

---

### 6.5 删除心智模型 DeleteMentalModel

**接口说明**

按 ID 删除心智模型。

**请求 URI**

```
DELETE /v1/default/banks/{bankId}/mental-models/{modelId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| modelId | String | 是 | URL 参数 | 心智模型 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

由服务端返回任意 JSON 对象（一般含 `success`/`message`）。

**请求示例**

```
DELETE /v1/default/banks/demo-bank/mental-models/mm-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "success": true }
```

**SDK 方法**

```go
out, err := client.DeleteMentalModel(ctx, "demo-bank", "mm-1")
```

---

### 6.6 刷新心智模型 RefreshMentalModel

**接口说明**

异步触发一次心智模型重新生成。

**请求 URI**

```
POST /v1/default/banks/{bankId}/mental-models/{modelId}/refresh
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| modelId | String | 是 | URL 参数 | 心智模型 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| operation_id | String | 异步任务 ID |
| status | String | 任务状态 |

**请求示例**

```
POST /v1/default/banks/demo-bank/mental-models/mm-1/refresh
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "operation_id": "op-1", "status": "pending" }
```

**SDK 方法**

```go
out, err := client.RefreshMentalModel(ctx, "demo-bank", "mm-1")
```

---

## 七、指令 Directives

### 7.1 列出指令 ListDirectives

**接口说明**

列出 bank 下的提示指令。

**请求 URI**

```
GET /v1/default/banks/{bankId}/directives?tags=&tags_match=&active_only=&limit=&offset=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| tags | Array<String> | 否 | Query 参数 | 标签过滤 |
| tags_match | String | 否 | Query 参数 | 标签匹配模式 |
| active_only | Boolean | 否 | Query 参数 | 仅返回激活的 |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| items | Array<DirectiveResponse> | 指令列表 |

**请求示例**

```
GET /v1/default/banks/demo-bank/directives?active_only=true
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [] }
```

**SDK 方法**

```go
out, err := client.ListDirectives(ctx, "demo-bank",
    DuMemory.ListDirectivesOptions{ActiveOnly: true})
```

---

### 7.2 创建指令 CreateDirective

**接口说明**

为 bank 添加一条提示指令，会被注入到 reflect 等流程的提示词中。

**请求 URI**

```
POST /v1/default/banks/{bankId}/directives
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| name | String | 是 | RequestBody | 指令名称 |
| content | String | 是 | RequestBody | 指令内容 |
| priority | Integer | 否 | RequestBody | 优先级，越大越先注入 |
| is_active | Boolean | 否 | RequestBody | 是否生效，默认 true |
| tags | Array<String> | 否 | RequestBody | 标签 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| id | String | 指令 ID |
| bank_id | String | 记忆库 ID |
| name | String | 名称 |
| content | String | 内容 |
| priority | Integer | 优先级 |
| is_active | Boolean | 是否生效 |
| tags | Array<String> | 标签 |
| created_at | String | 创建时间 |
| updated_at | String | 更新时间 |

**请求示例**

```
POST /v1/default/banks/demo-bank/directives
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "name": "tone", "content": "用中文回答" }
```

**返回示例**

```json
{ "id": "dir-1", "bank_id": "demo-bank", "name": "tone", "content": "用中文回答" }
```

**SDK 方法**

```go
out, err := client.CreateDirective(ctx, "demo-bank",
    *DuMemory.NewCreateDirectiveRequest("tone", "用中文回答"))
```

---

### 7.3 查看指令 GetDirective

**接口说明**

按 ID 查询指令详情。

**请求 URI**

```
GET /v1/default/banks/{bankId}/directives/{directiveId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| directiveId | String | 是 | URL 参数 | 指令 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

参见 7.2 返回参数。

**请求示例**

```
GET /v1/default/banks/demo-bank/directives/dir-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "id": "dir-1", "bank_id": "demo-bank", "name": "tone", "content": "用中文回答" }
```

**SDK 方法**

```go
out, err := client.GetDirective(ctx, "demo-bank", "dir-1")
```

---

### 7.4 更新指令 UpdateDirective

**接口说明**

按 ID 更新指令的字段（任何字段均为可选）。

**请求 URI**

```
PATCH /v1/default/banks/{bankId}/directives/{directiveId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| directiveId | String | 是 | URL 参数 | 指令 ID |
| name | String | 否 | RequestBody | 新名称 |
| content | String | 否 | RequestBody | 新内容 |
| priority | Integer | 否 | RequestBody | 优先级 |
| is_active | Boolean | 否 | RequestBody | 是否生效 |
| tags | Array<String> | 否 | RequestBody | 全量替换标签 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

参见 7.2。

**请求示例**

```
PATCH /v1/default/banks/demo-bank/directives/dir-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{ "is_active": false }
```

**返回示例**

```json
{ "id": "dir-1", "bank_id": "demo-bank", "is_active": false }
```

**SDK 方法**

```go
req := DuMemory.NewUpdateDirectiveRequest()
req.SetIsActive(false)
out, err := client.UpdateDirective(ctx, "demo-bank", "dir-1", *req)
```

---

### 7.5 删除指令 DeleteDirective

**接口说明**

按 ID 删除指令。

**请求 URI**

```
DELETE /v1/default/banks/{bankId}/directives/{directiveId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| directiveId | String | 是 | URL 参数 | 指令 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

服务端返回 `{ "success": true, "message": "..." }`。

**请求示例**

```
DELETE /v1/default/banks/demo-bank/directives/dir-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "success": true, "message": "deleted" }
```

**SDK 方法**

```go
out, err := client.DeleteDirective(ctx, "demo-bank", "dir-1")
```

---

## 八、操作 Operations

### 8.1 列出后台操作 ListOperations

**接口说明**

分页列出 bank 下的异步操作（retain、consolidation、refresh 等）。

**请求 URI**

```
GET /v1/default/banks/{bankId}/operations?status=&type=&limit=&offset=&exclude_parents=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| status | String | 否 | Query 参数 | 状态过滤：`pending` `processing` `completed` `failed` `cancelled` |
| type | String | 否 | Query 参数 | 操作类型过滤 |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |
| exclude_parents | Boolean | 否 | Query 参数 | 是否排除父任务 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| bank_id | String | 记忆库 ID |
| total | Integer | 总数 |
| limit | Integer | 分页大小 |
| offset | Integer | 偏移量 |
| operations | Array<OperationResponse> | 操作列表（id / task_type / status / created_at / progress 等） |

**请求示例**

```
GET /v1/default/banks/demo-bank/operations?status=pending&limit=20
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "bank_id": "demo-bank",
  "total": 0,
  "limit": 20,
  "offset": 0,
  "operations": []
}
```

**SDK 方法**

```go
out, err := client.ListOperations(ctx, "demo-bank",
    DuMemory.ListOperationsOptions{Status: "pending", Limit: 20})
```

---

### 8.2 查询后台操作状态 GetOperationStatus

**接口说明**

根据 operation ID 查询单个异步操作的当前状态、进度、错误信息、结果元数据及子操作状态。

**请求 URI**

```
GET /v1/default/banks/{bankId}/operations/{operationId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| operationId | String | 是 | URL 参数 | 操作 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| operation_id | String | 操作 ID |
| status | String | 操作状态，如 `pending`、`processing`、`completed`、`failed`、`cancelled` |
| operation_type | String | 操作类型，可空 |
| created_at | String | 创建时间，可空 |
| updated_at | String | 更新时间，可空 |
| completed_at | String | 完成时间，可空 |
| error_message | String | 错误信息，可空 |
| retry_count | Integer | 已重试次数，可空 |
| next_retry_at | String | 下次重试时间，可空 |
| progress | Object | 操作进度信息，可空 |
| result_metadata | Object | 操作结果元数据，可空 |
| child_operations | Array<ChildOperationStatus> | 子操作状态列表，可空 |
| task_payload | Object | 任务请求载荷，可空 |

**请求示例**

```
GET /v1/default/banks/demo-bank/operations/op-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "operation_id": "op-1",
  "status": "processing",
  "operation_type": "retain",
  "created_at": "2026-06-09T10:00:00Z",
  "updated_at": "2026-06-09T10:00:05Z",
  "retry_count": 0,
  "progress": {
    "current": 1,
    "total": 3
  }
}
```

**SDK 方法**

```go
out, err := client.GetOperationStatus(ctx, "demo-bank", "op-1")
```

---

### 8.3 取消后台操作 CancelOperation

**接口说明**

取消单个异步操作。

**请求 URI**

```
DELETE /v1/default/banks/{bankId}/operations/{operationId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| operationId | String | 是 | URL 参数 | 操作 ID |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| success | Boolean | 是否成功 |
| message | String | 信息描述 |
| operation_id | String | 操作 ID |

**请求示例**

```
DELETE /v1/default/banks/demo-bank/operations/op-1
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "success": true, "message": "cancelled", "operation_id": "op-1" }
```

**SDK 方法**

```go
out, err := client.CancelOperation(ctx, "demo-bank", "op-1")
```

---

## 九、文件 Files

### 9.1 文件上传写入 FilesRetain

**接口说明**

通过 multipart/form-data 上传文件并自动转写为记忆，返回每个文件对应的异步任务 ID。

**请求 URI**

```
POST /v1/default/banks/{bankId}/files/retain
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: multipart/form-data; boundary=...
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| files | File[] | 是 | Form 字段 | 一个或多个待上传文件 |
| request | String | 否 | Form 字段 | 序列化后的 RetainRequest JSON |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| operation_ids | Array<String> | 每个文件对应的异步任务 ID，可用 `ListOperations` 跟踪 |

**请求示例**

```
POST /v1/default/banks/demo-bank/files/retain
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: multipart/form-data; boundary=----X

------X
Content-Disposition: form-data; name="files"; filename="note.pdf"
Content-Type: application/pdf

<binary>
------X--
```

**返回示例**

```json
{ "operation_ids": ["op-1"] }
```

**SDK 方法**

```go
f, _ := os.Open("note.pdf")
defer f.Close()
out, err := client.FilesRetain(ctx, "demo-bank", []*os.File{f}, "")
```

---

## 十、标签作用域扩展接口

本节描述本次新增的 `WithScope` 系列 SDK 扩展接口。该系列接口复用服务端已有 REST API，不新增服务端路径；SDK 会在请求中的 `tags`、`document_tags` 或查询参数 `tags` 中自动追加实体作用域标签，用于把记忆、文档、指令、心智模型隔离到指定用户、Agent、App 或 Run。

作用域由 `EntityScope` 表示：

```go
type EntityScope struct {
    UserID  string
    AgentID string
    AppID   string
    RunID   string
}
```

生成的标签格式固定为：

| 字段 | 生成 tag | 示例 |
| --- | --- | --- |
| UserID | `user_id:{UserID}` | `user_id:u-001` |
| AgentID | `agent_id:{AgentID}` | `agent_id:agent-001` |
| AppID | `app_id:{AppID}` | `app_id:app-001` |
| RunID | `run_id:{RunID}` | `run_id:run-001` |

至少需要设置一个实体 ID。如果 `UserID`、`AgentID`、`AppID`、`RunID` 全部为空，SDK 返回错误：

```text
At least one entity ID (user_id, agent_id, app_id, or run_id) is required.
```

对于过滤类接口，如果调用方未显式设置 `tags_match`，或仍保持上游构造器默认值 `any`，SDK 会自动使用 `all_strict`。这样可以确保所有 scope tags 都必须命中，并排除无 tag 的 global memory / document / directive / mental model。如果调用方显式设置 `exact`、`all`、`any_strict` 或 `all_strict`，SDK 会保留调用方设置。

---

### 10.1 作用域写入记忆 RetainWithScope

**接口说明**

向指定 bank 写入记忆，并自动把实体作用域 tags 追加到每条 `items[].tags`。当前 Go SDK 生成模型已支持 `document_tags`，因此 SDK 也会把实体作用域 tags 追加到批次级 `document_tags`，使文档级标签与 memory units 保持一致。调用方已有业务 tags 会保留并去重。

**请求 URI**

```
POST /v1/default/banks/{bankId}/memories
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope.UserID | String | 否 | SDK 参数 | 用户 ID，生成 `user_id:{UserID}` |
| scope.AgentID | String | 否 | SDK 参数 | Agent ID，生成 `agent_id:{AgentID}` |
| scope.AppID | String | 否 | SDK 参数 | App ID，生成 `app_id:{AppID}` |
| scope.RunID | String | 否 | SDK 参数 | Run ID，生成 `run_id:{RunID}` |
| items | Array<MemoryItem> | 是 | RequestBody | 待写入的记忆条目 |
| items[].tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |
| document_tags | Array<String> | 否 | RequestBody | 批次级文档 tags；SDK 自动追加 scope tags |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `Retain`，包含 `success`、`bank_id`、`items_count`、`async`、`operation_id`、`operation_ids`、`usage` 等字段。

**请求示例**

```http
POST /v1/default/banks/demo-bank/memories
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "items": [
    {
      "content": "User likes concise technical explanations.",
      "tags": ["topic:preference", "user_id:u-001", "app_id:app-001"]
    }
  ],
  "document_tags": ["source:example", "user_id:u-001", "app_id:app-001"]
}
```

**返回示例**

```json
{
  "success": true,
  "bank_id": "demo-bank",
  "items_count": 1,
  "async": false
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AppID: "app-001"}
item := *dumemory.NewMemoryItem("User likes concise technical explanations.")
item.Tags = []string{"topic:preference"}
req := dumemory.NewRetainRequest([]dumemory.MemoryItem{item})
req.DocumentTags = []string{"source:example"}
out, err := client.RetainWithScope(ctx, "demo-bank", scope, *req)
```

---

### 10.2 作用域召回记忆 RecallWithScope

**接口说明**

在指定实体作用域内召回记忆。SDK 会把 scope tags 追加到请求体 `tags`，并在未显式设置匹配模式时使用 `all_strict`，避免召回无 tag 的 global memory。

**请求 URI**

```
POST /v1/default/banks/{bankId}/memories/recall
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| query | String | 是 | RequestBody | 查询语句 |
| tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |
| tags_match | String | 否 | RequestBody | 标签匹配模式，默认使用 `all_strict` |
| tag_groups | Array<Object> | 否 | RequestBody | 多组标签条件，SDK 保留原值 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `Recall`，包含 `results`、`trace`、`entities`、`chunks`、`source_facts` 等字段，返回结果中可包含 tags。

**请求示例**

```http
POST /v1/default/banks/demo-bank/memories/recall
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "query": "What response style does the user prefer?",
  "tags": ["topic:preference", "user_id:u-001", "app_id:app-001"],
  "tags_match": "all_strict"
}
```

**返回示例**

```json
{
  "results": [
    {
      "id": "m-001",
      "content": "User likes concise technical explanations.",
      "tags": ["topic:preference", "user_id:u-001", "app_id:app-001"]
    }
  ]
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AppID: "app-001"}
req := dumemory.NewRecallRequest("What response style does the user prefer?")
req.Tags = []string{"topic:preference"}
out, err := client.RecallWithScope(ctx, "demo-bank", scope, *req)
```

---

### 10.3 作用域综合记忆 ReflectWithScope

**接口说明**

在指定实体作用域内综合记忆生成回答。SDK 会把 scope tags 追加到请求体 `tags`，默认使用 `all_strict` 过滤参与综合的 memory。

**请求 URI**

```
POST /v1/default/banks/{bankId}/reflect
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| query | String | 是 | RequestBody | 综合问题 |
| tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |
| tags_match | String | 否 | RequestBody | 标签匹配模式，默认使用 `all_strict` |
| response_schema | Object | 否 | RequestBody | 结构化输出 schema |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `Reflect`，包含 `text`、`based_on`、`structured_output`、`usage`、`trace` 等字段。

**请求示例**

```http
POST /v1/default/banks/demo-bank/reflect
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "query": "Summarize the user's working preferences.",
  "tags": ["topic:preference", "user_id:u-001", "agent_id:agent-001"],
  "tags_match": "all_strict"
}
```

**返回示例**

```json
{
  "text": "The user prefers concise and practical technical explanations."
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AgentID: "agent-001"}
req := dumemory.NewReflectRequest("Summarize the user's working preferences.")
req.Tags = []string{"topic:preference"}
out, err := client.ReflectWithScope(ctx, "demo-bank", scope, *req)
```

---

### 10.4 作用域获取记忆 GetMemoryWithScope

**接口说明**

获取单条 memory，并要求调用方显式传入实体作用域。服务端 `GET /memories/{memoryId}` 不支持 tag 过滤，因此该方法主要用于在 SDK 层统一校验调用方必须提供 scope；返回体中的 tags 可由业务侧进一步校验。

**请求 URI**

```
GET /v1/default/banks/{bankId}/memories/{memoryId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| memoryId | String | 是 | URL 参数 | Memory ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID，用于强制调用方声明作用域 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

服务端返回 memory 详情，通常包含 `id`、`content`、`metadata`、`tags` 等字段。

**请求示例**

```http
GET /v1/default/banks/demo-bank/memories/m-001
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "id": "m-001",
  "content": "User likes concise technical explanations.",
  "tags": ["topic:preference", "user_id:u-001"]
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001"}
out, err := client.GetMemoryWithScope(ctx, "demo-bank", "m-001", scope)
```

---

### 10.5 作用域列出标签 ListTagsWithScope

**接口说明**

列出指定实体作用域下的 tag。SDK 会先校验 scope；如果 `ListTagsOptions.Q` 为空，则默认使用第一个 scope tag 加通配符作为查询条件，例如 `user_id:u-001*`。如果调用方显式设置 `Q`，SDK 保留该查询条件。

**请求 URI**

```
GET /v1/default/banks/{bankId}/tags?q={q}&source={source}&limit={limit}&offset={offset}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| q | String | 否 | Query 参数 | 通配查询，如 `user_id:u-001*` |
| source | String | 否 | Query 参数 | 标签来源，如 `memories`、`mental_models` |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

| 参数名称 | 参数类型 | 描述 |
| --- | --- | --- |
| items | Array<TagItem> | 标签列表，每项包含 `tag` 与 `count` |
| total | Integer | 总数 |
| limit | Integer | 分页大小 |
| offset | Integer | 偏移量 |

**请求示例**

```http
GET /v1/default/banks/demo-bank/tags?q=user_id:u-001*&source=memories&limit=20
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{
  "items": [
    { "tag": "user_id:u-001", "count": 10 },
    { "tag": "topic:preference", "count": 3 }
  ],
  "total": 2,
  "limit": 20,
  "offset": 0
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001"}
out, err := client.ListTagsWithScope(ctx, "demo-bank", scope, dumemory.ListTagsOptions{
    Source: "memories",
    Limit:  20,
})
```

---

### 10.6 作用域列出记忆 ListMemoriesWithScope

**接口说明**

按实体作用域分页列出记忆单元。SDK 会将 `EntityScope` 生成的 tags 与 `ListMemoriesOptions.Tags` 合并并去重；`TagsMatch` 为空或为上游默认值 `any` 时自动使用 `all_strict`，确保所有作用域 tags 都命中并排除无 tag 的 global memory。其它 `ListMemoriesOptions` 过滤条件保持不变。

**请求 URI**

```
GET /v1/default/banks/{bankId}/memories/list?type=&q=&consolidation_state=&state=&document_id=&entity_id=&tags=&tags_match=&limit=&offset=
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| type | String | 否 | Query 参数 | 节点/事实类型 |
| q | String | 否 | Query 参数 | 关键词模糊匹配 |
| consolidation_state | String | 否 | Query 参数 | 整合状态过滤 |
| state | String | 否 | Query 参数 | 记忆状态过滤 |
| document_id | String | 否 | Query 参数 | 来源文档 ID |
| entity_id | String | 否 | Query 参数 | 关联实体 ID |
| tags | Array<String> | 否 | Query 参数 | 业务 tags；SDK 自动追加 scope tags |
| tags_match | String | 否 | Query 参数 | 默认使用 `all_strict` |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `ListMemories`，包含 `items`、`total`、`limit`、`offset`。

**请求示例**

```
GET /v1/default/banks/demo-bank/memories/list?state=active&document_id=doc-1&entity_id=entity-1&tags=topic:coffee&tags=user_id:u-001&tags=app_id:app-001&tags_match=all_strict&limit=20
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [], "total": 0, "limit": 20, "offset": 0 }
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AppID: "app-001"}
out, err := client.ListMemoriesWithScope(ctx, "demo-bank", scope, dumemory.ListMemoriesOptions{
    Q:          "coffee",
    State:      "active",
    DocumentID: "doc-1",
    EntityID:   "entity-1",
    Tags:       []string{"topic:coffee"},
    Limit:      20,
})
```

---

### 10.7 作用域列出文档 ListDocumentsWithScope

**接口说明**

按实体作用域列出文档。SDK 会把 scope tags 追加到 `ListDocumentsOptions.Tags`，并默认使用 `all_strict`，避免返回无 tag 的 global document。

**请求 URI**

```
GET /v1/default/banks/{bankId}/documents?tags={tag}&tags_match={tagsMatch}&limit={limit}&offset={offset}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| tags | Array<String> | 否 | Query 参数 | 业务 tags；SDK 自动追加 scope tags |
| tags_match | String | 否 | Query 参数 | 默认使用 `all_strict` |
| q | String | 否 | Query 参数 | 关键词 |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `ListDocuments`，包含 `items`、`total`、`limit`、`offset`。

**请求示例**

```http
GET /v1/default/banks/demo-bank/documents?tags=source:example&tags=user_id:u-001&tags=run_id:run-001&tags_match=all_strict&limit=20
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [], "total": 0, "limit": 20, "offset": 0 }
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", RunID: "run-001"}
out, err := client.ListDocumentsWithScope(ctx, "demo-bank", scope, dumemory.ListDocumentsOptions{
    Tags:  []string{"source:example"},
    Limit: 20,
})
```

---

### 10.8 作用域更新文档标签 UpdateDocumentTagsWithScope

**接口说明**

更新文档 tags，并自动追加实体作用域 tags。服务端会把文档 tags 传播到关联 memory units，并使衍生 observation 失效，等待后续重算。

**请求 URI**

```
PATCH /v1/default/banks/{bankId}/documents/{documentId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| documentId | String | 是 | URL 参数 | 文档 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `UpdateDocument`，通常包含 `success`。

**请求示例**

```http
PATCH /v1/default/banks/demo-bank/documents/doc-001
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "tags": ["source:example", "topic:preference", "user_id:u-001", "app_id:app-001"]
}
```

**返回示例**

```json
{ "success": true }
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AppID: "app-001"}
req := dumemory.NewUpdateDocumentRequest()
req.Tags = []string{"source:example", "topic:preference"}
out, err := client.UpdateDocumentTagsWithScope(ctx, "demo-bank", "doc-001", scope, *req)
```

---

### 10.9 作用域列出指令 ListDirectivesWithScope

**接口说明**

按实体作用域列出指令。SDK 会把 scope tags 追加到 `ListDirectivesOptions.Tags`，并默认使用 `all_strict`。

**请求 URI**

```
GET /v1/default/banks/{bankId}/directives?tags={tag}&tags_match={tagsMatch}&active_only={activeOnly}&limit={limit}&offset={offset}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| tags | Array<String> | 否 | Query 参数 | 业务 tags；SDK 自动追加 scope tags |
| tags_match | String | 否 | Query 参数 | 默认使用 `all_strict` |
| active_only | Boolean | 否 | Query 参数 | 是否只返回启用指令 |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `ListDirectives`，包含 `items`。

**请求示例**

```http
GET /v1/default/banks/demo-bank/directives?tags=topic:style&tags=user_id:u-001&tags=agent_id:agent-001&tags_match=all_strict&active_only=true
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [] }
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AgentID: "agent-001"}
out, err := client.ListDirectivesWithScope(ctx, "demo-bank", scope, dumemory.ListDirectivesOptions{
    Tags:       []string{"topic:style"},
    ActiveOnly: true,
    Limit:      20,
})
```

---

### 10.10 作用域创建指令 CreateDirectiveWithScope

**接口说明**

创建指令，并自动把实体作用域 tags 追加到请求体 `tags`。

**请求 URI**

```
POST /v1/default/banks/{bankId}/directives
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| name | String | 是 | RequestBody | 指令名称 |
| content | String | 是 | RequestBody | 指令内容 |
| priority | Integer | 否 | RequestBody | 优先级 |
| is_active | Boolean | 否 | RequestBody | 是否启用 |
| tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `CreateDirective`，包含 `id`、`bank_id`、`name`、`content`、`priority`、`is_active`、`tags`、`created_at`、`updated_at`。

**请求示例**

```http
POST /v1/default/banks/demo-bank/directives
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "name": "scoped-tone",
  "content": "Use a concise and practical tone.",
  "tags": ["topic:style", "user_id:u-001", "agent_id:agent-001"]
}
```

**返回示例**

```json
{
  "id": "dir-001",
  "bank_id": "demo-bank",
  "name": "scoped-tone",
  "content": "Use a concise and practical tone.",
  "tags": ["topic:style", "user_id:u-001", "agent_id:agent-001"]
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AgentID: "agent-001"}
req := dumemory.NewCreateDirectiveRequest("scoped-tone", "Use a concise and practical tone.")
req.Tags = []string{"topic:style"}
out, err := client.CreateDirectiveWithScope(ctx, "demo-bank", scope, *req)
```

---

### 10.11 作用域更新指令 UpdateDirectiveWithScope

**接口说明**

更新指令，并自动把实体作用域 tags 追加到请求体 `tags`。

**请求 URI**

```
PATCH /v1/default/banks/{bankId}/directives/{directiveId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| directiveId | String | 是 | URL 参数 | 指令 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| name | String | 否 | RequestBody | 新名称 |
| content | String | 否 | RequestBody | 新内容 |
| priority | Integer | 否 | RequestBody | 优先级 |
| is_active | Boolean | 否 | RequestBody | 是否启用 |
| tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `UpdateDirective`，返回更新后的指令对象。

**请求示例**

```http
PATCH /v1/default/banks/demo-bank/directives/dir-001
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "tags": ["topic:style", "state:active", "user_id:u-001", "agent_id:agent-001"]
}
```

**返回示例**

```json
{
  "id": "dir-001",
  "bank_id": "demo-bank",
  "name": "scoped-tone",
  "content": "Use a concise and practical tone.",
  "tags": ["topic:style", "state:active", "user_id:u-001", "agent_id:agent-001"]
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AgentID: "agent-001"}
req := dumemory.NewUpdateDirectiveRequest()
req.Tags = []string{"topic:style", "state:active"}
out, err := client.UpdateDirectiveWithScope(ctx, "demo-bank", "dir-001", scope, *req)
```

---

### 10.12 作用域列出心智模型 ListMentalModelsWithScope

**接口说明**

按实体作用域列出心智模型。SDK 会把 scope tags 追加到 `ListMentalModelsOptions.Tags`，并默认使用 `all_strict`。

**请求 URI**

```
GET /v1/default/banks/{bankId}/mental-models?tags={tag}&tags_match={tagsMatch}&detail={detail}&limit={limit}&offset={offset}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| tags | Array<String> | 否 | Query 参数 | 业务 tags；SDK 自动追加 scope tags |
| tags_match | String | 否 | Query 参数 | 默认使用 `all_strict` |
| detail | String | 否 | Query 参数 | 详情级别 |
| limit | Integer | 否 | Query 参数 | 分页大小 |
| offset | Integer | 否 | Query 参数 | 偏移量 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `ListMentalModels`，包含 `items`。

**请求示例**

```http
GET /v1/default/banks/demo-bank/mental-models?tags=topic:preference&tags=user_id:u-001&tags=app_id:app-001&tags_match=all_strict&limit=20
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
```

**返回示例**

```json
{ "items": [] }
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AppID: "app-001"}
out, err := client.ListMentalModelsWithScope(ctx, "demo-bank", scope, dumemory.ListMentalModelsOptions{
    Tags:  []string{"topic:preference"},
    Limit: 20,
})
```

---

### 10.13 作用域创建心智模型 CreateMentalModelWithScope

**接口说明**

创建心智模型，并自动把实体作用域 tags 追加到请求体 `tags`。

**请求 URI**

```
POST /v1/default/banks/{bankId}/mental-models
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| name | String | 是 | RequestBody | 心智模型名称 |
| source_query | String | 是 | RequestBody | 用于生成模型内容的查询语句 |
| tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |
| max_tokens | Integer | 否 | RequestBody | 生成上限 |
| trigger | Object | 否 | RequestBody | 触发条件 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `CreateMentalModel`，包含 `mental_model_id`、`operation_id`。

**请求示例**

```http
POST /v1/default/banks/demo-bank/mental-models
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "name": "User Preference Model",
  "source_query": "What does this user prefer while working?",
  "tags": ["topic:preference", "user_id:u-001", "app_id:app-001"]
}
```

**返回示例**

```json
{
  "mental_model_id": "mm-001",
  "operation_id": "op-001"
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AppID: "app-001"}
req := dumemory.NewCreateMentalModelRequest("User Preference Model", "What does this user prefer while working?")
req.Tags = []string{"topic:preference"}
out, err := client.CreateMentalModelWithScope(ctx, "demo-bank", scope, *req)
```

---

### 10.14 作用域更新心智模型 UpdateMentalModelWithScope

**接口说明**

更新心智模型，并自动把实体作用域 tags 追加到请求体 `tags`。

**请求 URI**

```
PATCH /v1/default/banks/{bankId}/mental-models/{modelId}
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**请求头域**

除公共头域外，无其它特殊头域。

**请求参数**

| 参数名称 | 参数类型 | 是否必须 | 参数位置 | 描述 |
| --- | --- | --- | --- | --- |
| bankId | String | 是 | URL 参数 | 记忆库 ID |
| modelId | String | 是 | URL 参数 | 心智模型 ID |
| scope | EntityScope | 是 | SDK 参数 | 至少包含一个实体 ID |
| name | String | 否 | RequestBody | 新名称 |
| source_query | String | 否 | RequestBody | 新查询 |
| max_tokens | Integer | 否 | RequestBody | 生成上限 |
| tags | Array<String> | 否 | RequestBody | 业务 tags；SDK 自动追加 scope tags |
| trigger | Object | 否 | RequestBody | 触发条件 |

**返回头域**

除公共头域外，无其它特殊头域。

**返回参数**

同 `UpdateMentalModel`，返回更新后的心智模型对象。

**请求示例**

```http
PATCH /v1/default/banks/demo-bank/mental-models/mm-001
Host: cloud.memory.bj.baidubce.com
Authorization: Bearer bce-v3/ALTAK-xxx/xxx
Content-Type: application/json

{
  "tags": ["topic:preference", "state:active", "user_id:u-001", "app_id:app-001"]
}
```

**返回示例**

```json
{
  "id": "mm-001",
  "bank_id": "demo-bank",
  "name": "User Preference Model",
  "tags": ["topic:preference", "state:active", "user_id:u-001", "app_id:app-001"]
}
```

**SDK 方法**

```go
scope := dumemory.EntityScope{UserID: "u-001", AppID: "app-001"}
req := dumemory.NewUpdateMentalModelRequest()
req.Tags = []string{"topic:preference", "state:active"}
out, err := client.UpdateMentalModelWithScope(ctx, "demo-bank", "mm-001", scope, *req)
```

---

## 附录 A：错误处理

错误统一通过 DuMemory 客户端封装的 *GenericOpenAPIError 抛出，可使用 errors.As 提取响应体：

```go
var apiErr *DuMemory.GenericOpenAPIError
if errors.As(err, &apiErr) {
    log.Printf("status=%d body=%s", 0, apiErr.Body())
}
```

## 附录 B：版本兼容

- 服务端 0.6.x 与 0.8.x 在 `/version` 返回的 `features` 字段不一致；SDK 的 `Version()` 用宽松解码兼容两端，返回 `*VersionInfo` 而非 hindsight 自带的强类型。
- `ConsolidateBank` 在 0.6.x 不支持请求体，0.8.x 起支持。SDK 中 `req` 参数允许为 `nil` 以保持兼容。
