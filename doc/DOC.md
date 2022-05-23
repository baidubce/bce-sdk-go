# DOC 服务

# 概述

本文档主要介绍 DOC GO SDK 的使用。在使用本文档前，您需要先了解 DOC 的一些基本知识，并已开通了 DOC 服务。若您还不了解 DOC，可以参考[产品描述](https://cloud.baidu.com/doc/DOC/s/Djwvypqoi)和 [API 参考](https://cloud.baidu.com/doc/DOC/s/Cjwvypy6e)。

# 初始化

## 确认Endpoint

DOC 为全局服务，服务域名是 `doc.bj.baidubce.com`。 DOC API 支持 HTTP 和 HTTPS 两种协议。为了提升数据的安全性，建议使用HTTPS协议。SDK 默认使用 HTTPS 协议。

## 获取密钥

要使用百度云 DOC，您需要拥有一个有效的 AK(Access Key ID) 和 SK(Secret Access Key) 用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问 DOC 做签名验证。

可以通过如下步骤获得并了解您的 AK/SK 信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建DOC Client

DOC Client 是 DOC 服务的客户端，为开发者与 DOC 服务进行交互提供了一系列的方法。

### 使用 AK/SK 新建 DOC Client

通过 AK/SK 方式访问 DOC，用户可以参考如下代码新建一个 DOC Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/doc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 初始化一个DocClient
	docClient, err := doc.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《[如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb)》。

## 配置 DOC Client

如果用户需要配置 DOC Client 的一些细节的参数，可以在创建 DOC Client 对象之后，使用该对象的导出字段 `Config` 进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问 DOC 服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/doc"

//创建 DOC Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := doc.NewClient(AK, SK)

//代理使用本地的 8080 端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/doc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := doc.NewClient(AK, SK)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/doc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := doc.NewClient(AK, SK)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用 GO SDK 访问 DOC 时，创建的 DOC Client 对象的 `Config` 字段支持的所有参数如下表所示：

配置项名称 |  类型   | 含义
-----------|---------|--------
Endpoint   |  string | 请求服务的域名
ProxyUrl   |  string | 客户端请求的代理地址
Region     |  string | 请求资源的区域
UserAgent  |  string | 用户名称，HTTP请求的User-Agent头
Credentials| \*auth.BceCredentials | 请求的鉴权对象，分为普通AK/SK与STS两种
SignOption | \*auth.SignOptions    | 认证字符串签名选项
Retry      | RetryPolicy | 连接重试策略
ConnectionTimeoutInMillis| int     | 连接超时时间，单位毫秒，默认20分钟

说明：

  1. `Credentials` 字段使用 `auth.NewBceCredentials` 与 `auth.NewSessionBceCredentials` 函数创建，默认使用前者。
  2. `SignOption` 字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry` 字段指定重试策略，目前支持两种：`NoRetryPolicy` 和 `BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 文档服务
文档接口流程如下
![百度云文档接口](https://doc.bce.baidu.com/bce-documentation/DOC/wendangjiekou_1.png)

## 注册文档
注册文档接口用于生成文档的唯一标识documentId、用于存储源文档文件的 BOS Bucket 相关信息。注册成功后，对应文档状态为 `UPLOADING`，对应 Bucket 对用户开放写权限，对于用户的 BOS 空间不可见。

注册文档是文档三步创建法（注册文档、上传BOS、发布文档）的第一步。

如下代码可以注册一个文档：

```go
regParam := &api.RegDocumentParam{
    Title:  <your-doc-title>,
    Format: <your-doc-format>,
}

if res, err := docClient.RegisterDocument(regParam); err != nil {
	fmt.Println("failed to register document:", err)
} else {
	fmt.Println("register document success, id:", res.DocumentId)
}
```

## 发布文档
用于对已完成注册和 BOS 上传的文档进行发布处理。仅对状态为 `UPLOADING` 的文档有效。处理过程中，文档状态为 `PROCESSING`；处理完成后，状态转为 `PUBLISHED`。

发布文档是文档三步创建法（注册文档、上传BOS、发布文档）的第三步。

```go
err = docClient.PublishDocument(<your-doc-id>)
```

## 查询文档
通过文档的唯一标识 documentId 查询指定文档的详细信息。

```go
qRes, err := docClient.QueryDocument(<your-doc-id>, &api.QueryDocumentParam{Https: false})
```

## 文档列表
查询所有文档，以列表形式返回，支持用文档状态作为筛选条件进行筛选。

```go
listParam := &api.ListDocumentsParam{
		Status:  api.DOC_STATUS_PUBLISHED,
		MaxSize: 2,
	}
res, err := docClient.ListDocuments(listParam)
```

## 阅读文档
通过文档的唯一标识 documentId 获取指定文档的阅读信息，以便在 PC/Android/iOS 设备上阅读。仅对状态为 `PUBLISHED` 的文档有效。
```go
rRes, err := docClient.ReadDocument(<your-doc-id>, &api.ReadDocumentParam{ExpireInSeconds: 3600})
```

## 查询文档转码结果图片列表
对于转码结果类型为图片的文档，通过本接口可以在文档转码完成后，获取转码结果图片的URL列表。

```go
if res, err := docClient.GetImages(<your-doc-id>); err != nil {
	fmt.Println("failed to get images list:", err)
} else {
	fmt.Println("get images list success, images count:", len(res.Images))
}
```

## 删除文档
删除文档，仅对状态 status 不是 `PROCESSING` 时的文档有效，清除文档占用的存储空间。

```go
err = docClient.DeleteDocument(<your-doc-id>)
```

> **注意：**
> 文档一经删除，无法通过查询文档/文档列表等接口获取，并且无法阅读、下载，请谨慎操作。