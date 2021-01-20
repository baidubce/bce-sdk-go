# BLS服务

# 概述

本文档主要介绍BLS GO SDK的使用。在使用本文档前，您需要先了解BLS的一些基本知识，并已开通了BLS服务。若您还不了解BLS，可以参考[产品描述](https://cloud.baidu.com/doc/BLS/index.html)和[入门指南](https://cloud.baidu.com/doc/BLS/s/Gjwvyjbvg)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BLS访问域名](https://cloud.baidu.com/doc/BLS/s/4k8qysj2z)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx)。

目前支持“华北-北京”和“华南-广州”两个区域。北京区域：`bls-log.bj.baidubce.com`广州区域：`bls-log.gz.baidubce.com`。对应信息为：

| 访问区域  | 对应Endpoint            | Protocol   |
| --------- | ----------------------- | ---------- |
| 华北-北京 | bls-log.bj.baidubce.com | HTTP/HTTPS |
| 华南-广州 | bls-log.gz.baidubce.com | HTTP/HTTPS |

## 获取密钥

要使用百度云BLS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BLS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建BLS Client

BLS Client是BLS服务的客户端，为开发者与BLS服务进行交互提供了一系列的方法。

### 使用AK/SK新建BLS Client

通过AK/SK方式访问BLS，用户可以参考如下代码新建一个BLS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bls"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BLSClient
	blsClient, err := bls.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Secret Access Key”，获取方式请参考《通用参考 [获取AKSK](https://cloud.baidu.com/doc/Reference/s/jjwvz2e3p)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BLS的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`http://bls-log.bj.baidubce.com`。

### 使用STS创建BLS Client

**申请STS token**

BLS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问BLS，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建BLS Client**

申请好STS后，可将STS Token配置到BLS Client中，从而实现通过STS Token创建BLS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建BLS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/bls" //导入BLS服务模块
	"github.com/baidubce/bce-sdk-go/services/sts" //导入STS服务模块
)

func main() {
	// 创建STS服务的Client对象，Endpoint使用默认值
	AK, SK := <your-access-key-id>, <your-secret-access-key>
	stsClient, err := sts.NewClient(AK, SK)
	if err != nil {
		fmt.Println("create sts client object :", err)
		return
	}

	// 获取临时认证token，有效期为60秒，ACL为空
	stsObj, err := stsClient.GetSessionToken(60, "")
	if err != nil {
		fmt.Println("get session token failed:", err)
		return
    }
	fmt.Println("GetSessionToken result:")
	fmt.Println("  accessKeyId:", stsObj.AccessKeyId)
	fmt.Println("  secretAccessKey:", stsObj.SecretAccessKey)
	fmt.Println("  sessionToken:", stsObj.SessionToken)
	fmt.Println("  createTime:", stsObj.CreateTime)
	fmt.Println("  expiration:", stsObj.Expiration)
	fmt.Println("  userId:", stsObj.UserId)

	// 使用申请的临时STS创建BLS服务的Client对象，Endpoint使用默认值
	blsClient, err := bls.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create bls client failed:", err)
		return
	}
	stsCredential, err := auth.NewSessionBceCredentials(
		stsObj.AccessKeyId,
		stsObj.SecretAccessKey,
		stsObj.SessionToken)
	if err != nil {
		fmt.Println("create sts credential object failed:", err)
		return
	}
	blsClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置BLS Client时，无论对应BLS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问BLS

BLS支持HTTPS传输协议，您可以通过在创建BLS Client对象时指定的Endpoint中指明HTTPS的方式，在BLS GO SDK中使用HTTPS访问BLS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bls"

ENDPOINT := "https://bls-log.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
blsClient, _ := bls.NewClient(AK, SK, ENDPOINT)
```

## 配置BLS Client

如果用户需要配置BLS Client的一些细节的参数，可以在创建BLS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问BLS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bls"

//创建BLS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bls-log.bj.baidubce.com"
blsClient, _ := bls.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
blsClient.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/bls"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bls-log.bj.baidubce.com"
blsClient, _ := bls.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
blsClient.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
blsClient.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/bls"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bls-log.bj.baidubce.com"
blsClient, _ := bls.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
blsClient.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
blsClient.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问BLS时，创建的BLS Client对象的`Config`字段支持的所有参数如下表所示：

| 配置项名称                | 类型                  | 含义                                   |
| ------------------------- | --------------------- | -------------------------------------- |
| Endpoint                  | string                | 请求服务的域名                         |
| ProxyUrl                  | string                | 客户端请求的代理地址                   |
| Region                    | string                | 请求资源的区域                         |
| UserAgent                 | string                | 用户名称，HTTP请求的User-Agent头       |
| Credentials               | \*auth.BceCredentials | 请求的鉴权对象，分为普通AK/SK与STS两种 |
| SignOption                | \*auth.SignOptions    | 认证字符串签名选项                     |
| Retry                     | RetryPolicy           | 连接重试策略                           |
| ConnectionTimeoutInMillis | int                   | 连接超时时间，单位毫秒，默认20分钟     |

>  说明：
>
>     1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建BLS Client”小节。
>     2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：
>
> | 名称          | 类型                | 含义                                                   |
> | ------------- | ------------------- | ------------------------------------------------------ |
> | HeadersToSign | map[string]struct{} | 生成签名字符串时使用的HTTP头                           |
> | Timestamp     | int64               | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值 |
> | ExpireSeconds | int                 | 签名字符串的有效期                                     |
>
>      其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
>
>     3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。

# 主要接口

用户可以通过 API 的方式管理 BLS 日志集、写入、下载、查询和分析日志数据等操作。

## LogStore操作

### 创建LogStore

创建日志集，命名日志组时，需遵循以下准则：

- 每个账户每个区域日志集名称不能重复
- 日志集名称长度不能超过 128 个字符
- 日志集名称包含的字符仅限于： `a-z, A-Z, 0-9, '_', '-', '.'`

通过以下代码，创建一个LogStore并指定其存储期限。

```go
err := blsClient.CreateLogStore("demo", 3)
if err != nil {
  fmt.Println("Create logStore failed: ", err)
} else {
  fmt.Println("Create logStore success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[CreateLogStore](https://cloud.baidu.com/doc/BLS/s/pk8to0k59)

### 更新指定LogStore

通过以下代码，更新指定的日志集，目前仅支持更改与日志集关联的存储期限。

```go
err := blsClient.UpdateLogStore("demo", 5)
if err != nil {
  fmt.Println("Update logStore failed: ", err)
} else {
  fmt.Println("Update logStore success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[UpdateLogStore](https://cloud.baidu.com/doc/BLS/s/ok8to0kla)

### 查询指定LogStore

通过以下代码，获取指定日志集的详情信息。

```go
res, err := blsClient.DescribeLogStore("demo")
if err != nil {
  fmt.Println("Get logStore failed: ", err)
} else {
  fmt.Println("LogStore info: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[DescribeLogStore](https://cloud.baidu.com/doc/BLS/s/Bk8to0jp3)

### 获取LogStore列表

通过以下代码，获取当前用户的日志集列表。

```go
// 可选参数列表
args := &api.QueryConditions{
  NamePattern: "bls-log",
  Order:       "asc",
  OrderBy:     "creationDateTime",
  PageNo:      1,
  PageSize:    10}
res, err := blsClient.ListLogStore(args)
if err != nil {
  fmt.Println("Get logStore list failed: ", err)
} else {
  fmt.Println("List logStore success: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[ListLogStore](https://cloud.baidu.com/doc/BLS/s/Hk8to0kda)

### 删除LogStore

通过以下代码，删除指定的日志集，并且会永久删除与其关联的所有已存储日志记录。

```go
err := blsClient.DeleteLogStore("demo")
if err != nil {
  fmt.Println("Delete logStore failed: ", err)
} else {
  fmt.Println("Delete logStore success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[DeleteLogStore](https://cloud.baidu.com/doc/BLS/s/ak8to0jx8)

## LogStream操作

LogStream会随着LogStore的创建自动创建，目前暂不支持对LogStream的删除操作。

### 获取LogStream列表

通过以下代码，获取指定日志集的日志流列表。

```go
// 可选参数列表
args := &api.QueryConditions{
  NamePattern: "bls-log",
  Order:       "desc",
  OrderBy:     "creationDateTime",
  PageNo:      1,
  PageSize:    20,
}
res, err := blsClient.ListLogStore(args)
if err != nil {
  fmt.Println("Get logStream list failed: ", err)
} else {
  fmt.Println("List logStore success: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[ListLogStream](https://cloud.baidu.com/doc/BLS/s/dk8to0lhy)

## LogRecord操作

### 推送LogRecord

支持批量推送日志记录到 BLS 平台，日志记录的格式可以是 TEXT，也可以是 JSON 格式。如果是 TEXT，则不对日志进行解析；如果是 JSON 格式，可以自动发现 JSON 字段（仅支持首层字段发现，暂不支持嵌套类型字段的自动发现）。

如果既想上传日志原文，又想上传解析出的具体字段，可以使用 JSON 格式进行上传，并在 JSON 中包含日志原文(使用 @raw 作为key，日志原文作为 value)。 BLS 解析到 @raw 的时候，会将其内容作为日志原文处理。

通过以下代码，可以批量推送JSON日志记录到指定日志集的指定日志流中。

```go
// 推送JSON格式日志记录
jsonRecords := []api.LogRecord{
  {
    Message:   "{\"body_bytes_sent\":184,\"bytes_sent\":398,\"client_ip\":\"120.193.204.39\"}",
    Timestamp: time.Now().UnixNano() / 1e6,
    Sequence:  1,
  },
  {
    Message:   "{\"body_bytes_sent\":14,\"bytes_sent\":408,\"client_ip\":\"120.193.222.39\"}",
    Timestamp: time.Now().UnixNano() / 1e6,
    Sequence:  2,
  },
}
// 指定logRecord类型为JSON，并将日志记录推送到日志集demo中的日志流json-logStream中
// 若没有该日志流，则自动创建
err := blsClient.PushLogRecord("demo", "json-logStream", "JSON", jsonRecords)
if err != nil {
  fmt.Println("Push logRecords failed: ", err)
} else {
  fmt.Println("Push logRecords success")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[PushLogRecord](https://cloud.baidu.com/doc/BLS/s/dk8to0ktn)

### 查看指定LogRecord

通过以下代码，查看指定日志流中的日志记录，您可以获取最近的日志记录或使用时间范围进行过滤。

```go
args := &api.PullLogRecordArgs{
  // 必须指定日志流名称
  LogStreamName: "json-logStream",
  // 可选参数
  StartDateTime: "2021-01-01T10:11:44Z",
  EndDateTime:   "2021-12-10T16:11:44Z",
  Limit:         500, // 返回最大条目数
  Marker:        "",  // 指定查看的位置标记
}
res, err := blsClient.PullLogRecord("demo", args)
if err != nil {
  fmt.Println("Pull logRecord failed: ", err)
} else {
  fmt.Println("LogRecords result: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[PullLogRecord](https://cloud.baidu.com/doc/BLS/s/Ek8to0l1o)

### 查询指定LogRecord

用户通过提交 Query 检索或分析指定日志集中的数据，每次只能查询一个日志集的内容。

Query 语句格式支持 检索 + SQL分析，通过竖线分隔，即在检索的结果集上执行 SQL，形如：`Search | SQL`。

例如 `method:GET and status >= 400 | select host, count(*) group by host`

注：

- 如果只需要检索原日志，不需要执行 SQL 分析，可以省略竖线和 SQL 语句。
- 如果不需要检索，只需要 SQL 分析，那么检索语句可以写为 `*`，表示匹配所有记录。即 `* | SQL`。如果查询的日志集没有开启索引，也可以省略检索语句和竖线，只写 SQL 语句。

查询相关限制如下：

- 每个账户支持最多的查询并发数是 15 个
- 限制返回的结果集大小不超过 1MB 或 1000 条记录。

检索语法请参考 [检索语法](https://cloud.baidu.com/doc/BLS/s/Okbta3asp)

SQL 语句中可以不包括 from 子句，语法详情可以参考 [SQL 语法](https://cloud.baidu.com/doc/BLS/s/xk5cc9piu)

通过以下代码，您可以在指定日志集中查询满足条件的日志记录。

```go
args := &api.QueryLogRecordArgs{
  // 必选参数
  Query:         "select count(*)",      // 查询SQL
  StartDateTime: "2021-01-01T10:11:44Z", 
  EndDateTime:   "2021-12-10T16:11:44Z",
  // 可选参数
  LogStreamName: "json-logStream",       // 不指定则在日志集所有日志流中查询
  Limit:         100，
}
res, err := blsClient.QueryLogRecord("demo", args)
if err != nil {
  fmt.Println("Query logRecord failed: ", err)
} else {
  fmt.Println("LogRecords result: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[QueryLogRecord](https://cloud.baidu.com/doc/BLS/s/hk8to0l9o)

## FastQuery操作

### 创建FastQuery

创建快速查询的实例名称必须遵循以下准则：

- 每个账户每个区域快速查询名称不能相同
- 快速查询名称长度不能超过128个字符
- 快速查询名称包含的字符仅限于：`a-z, A-Z, 0-9, '_', '-', '.'`

通过以下代码，可以创建一个快速查询。

```go
args := &api.CreateFastQueryBody{
  FastQueryName: "macro",
  Query:         "select count(*)",
  LogStoreName:  "demo",
  // 可选参数
  Description:   "calculate record number",
  LogStreamName: "json-logStream",
}
err := blsClient.CreateFastQuery(args)
if err != nil {
  fmt.Println("Create fastQuery failed: ", err)
} else {
  fmt.Println('Create fastQuery success.')
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[CreateFastQuery](https://cloud.baidu.com/doc/BLS/s/kk8to0m6g)

### 获取指定FastQuery

通过以下代码，获取指定名称的快速查询的详细信息。

```go
res, err := blsClient.DescribeFastQuery("macro")
if err != nil {
  fmt.Println("Get fastQuery failed: ", err)
} else {
  fmt.Println("Fastquery info: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[DescribeFastQuery](https://cloud.baidu.com/doc/BLS/s/Zk8to0mmn)

### 更新指定FastQuery

通过以下代码，更新指定名称的快速查询实例信息。

```go
args := &api.UpdateFastQueryBody{
  LogStoreName:  "demo",
  // 可选参数
  Query:         "select * limit 3",
  Description:   "Top 3",
  LogStreamName: "",
}
err := blsClient.UpdateFastQuery("macro", args)
if err != nil {
  fmt.Println("Update fastQuery failed: ", err)
} else {
  fmt.Println("Update fastQuery success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[UpdateFastQuery](https://cloud.baidu.com/doc/BLS/s/Ik8to0mei)

### 获取FastQuery列表

通过以下代码，获取当前用户保存的快速查询列表。

```go
// 可选参数列表
args := &api.QueryConditions{
  NamePattern: "m",
  Order:       "desc",
  OrderBy:     "",
  PageNo:      1,
  PageSize:    20,
}
res, err := blsClient.ListFastQuery(args)
if err != nil {
  fmt.Println("List fastQuery failed: ", err)
} else {
  fmt.Println("FastQuery list: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[ListFastQuery](https://cloud.baidu.com/doc/BLS/s/0k8to0lyf)

### 删除指定FastQuery

通过以下代码，删除指定名称的快速查询示例。

```go
err := blsClient.DeleteFastQuery("macro")
if err != nil {
  fmt.Println("Delete fastQuery failed: ", err)
} else {
  fmt.Println("Delete fastQuery success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[DeleteFastQuery](https://cloud.baidu.com/doc/BLS/s/Uk8to0lqd)

## Index操作

### 创建Index

通过以下代码，为指定的日志集创建索引。

```go
indexMappings := map[string]api.LogField{
  "age": {
    Type: "long",
  },
  "salary": {
    Type: "text",
  },
  "name": {
    Type: "text",
  },
}
err := blsClient.CreateIndex("demo", true, indexMappings) // true表示索引开启状态
if err != nil {
  fmt.Println("Create index failed: ", err)
} else {
  fmt.Println("Create index success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[CreateIndex](https://cloud.baidu.com/doc/BLS/s/dkbt4q6ps)

### 更新指定Index

通过以下代码，更新指定日志集的索引结构。

```go
indexMappings := map[string]api.LogField{
  "age": {
    Type: "long",
  },
  "wage": {
    Type: "float",
  },
  "name": {
    Type: "text",
  },
}
err := blsClient.UdpateIndex("demo", true, indexMappings)
if err != nil {
  fmt.Println("Update index failed: ", err)
} else {
  fmt.Println("Update index success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[UpdateIndex](https://cloud.baidu.com/doc/BLS/s/bkbt4w0fe)

### 获取指定Index

通过以下代码，获取指定日志集的索引结构。

```go
res, err := blsClient.DescribeIndex("demo")
if err != nil {
  fmt.Println("Get index failed: ", err)
} else {
  fmt.Println("Index info: ", res)
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[DescribeIndex](https://cloud.baidu.com/doc/BLS/s/1kbt4yem2)

### 删除指定Index

通过以下代码，删除指定日志集的索引，该操作会将索引数据清空。

```go
res, err := blsClient.DeleteIndex("demo")
if err != nil {
  fmt.Println("Delete index failed: ", err)
} else {
  fmt.Println("Delete index success.")
}
```

> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BLS API 文档[DeleteIndex](https://cloud.baidu.com/doc/BLS/s/bkbt56uvu)

