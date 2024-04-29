# 标签管理(TAG)服务

# 概述

本文档主要介绍标签管理(TAG) GO SDK的使用。在使用本文档前，您需要先了解标签管理(TAG)的一些基本知识。若您还不了解标签管理(TAG)，可以参考[产品描述](https://cloud.baidu.com/doc/TAG/s/ukboeqze7)和[操作指南](https://cloud.baidu.com/doc/TAG/s/Tkbp242pi)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，标签管理(TAG)服务不分区域，使用全局域名，可先阅读开发人员指南中关于[API参考](https://cloud.baidu.com/doc/TAG/s/Hkbrb3rw5)的部分。

## 获取密钥

要使用百度云标签管理(TAG)，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问标签管理(TAG)做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建标签管理(TAG) Client

标签管理(TAG) Client是标签管理(TAG)服务的客户端，为开发者与标签管理(TAG)服务进行交互提供了一系列的方法。

### 使用AK/SK新建tag Client

通过AK/SK方式访问标签管理(TAG)服务，用户可以参考如下代码新建一个tag Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/tag"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个tagClient
    tagClient, err := tag.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Access Key Secret”。标签管理(TAG)为全局服务，第三个参数`ENDPOINT`，直接设置为空字符串，会使用默认域名作为标签管理(TAG)的服务地址。

### 使用STS创建标签管理(TAG) Client

**申请STS token**

标签管理(TAG)可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问标签管理(TAG)，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建tag Client**

申请好STS后，可将STS Token配置到tag Client中，从而实现通过STS Token创建tag Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建tag Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/tag" //导入标签管理(TAG)服务模块
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

	// 使用申请的临时STS创建标签管理(TAG)服务的Client对象，Endpoint使用默认值 
	tagClient, err := tag.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create tag client failed:", err)
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
    tagClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置标签管理(TAG) Client时，无论对应标签管理(TAG)服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问标签管理(TAG)

标签管理(TAG)支持HTTPS传输协议，您可以通过在创建tag Client对象时指定的Endpoint中指明HTTPS的方式，在标签管理(TAG) GO SDK中使用HTTPS访问标签管理(TAG)服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/tag"

ENDPOINT := "https://tag.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
tagClient, _ := tag.NewClient(AK, SK, ENDPOINT)
```

## 配置标签管理(TAG) Client

如果用户需要配置tag Client的一些细节的参数，可以在创建tag Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问标签管理(TAG)服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/tag"

//创建标签管理(TAG) Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "tag.baidubce.com"
client, _ := tag.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/tag"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "tag.baidubce.com"
client, _ := tag.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/tag"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "tag.baidubce.com"
client, _ := tag.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问标签管理(TAG)时，创建的标签管理(TAG) Client对象的`Config`字段支持的所有参数如下表所示：

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

1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建标签管理(TAG) Client”小节。
2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

标签管理(TAG)API对于每个HTTP请求，认证签名放在Authorization头域中，后端统一认证。且后端在Response头域中会添加x-bce-request-id，作为请求唯一标识，方便追踪定位问题。

目前支持的标签管理(TAG)API接口如下：

- 创建标签 
- 删除标签 
- 标签列表 
- 查看标签下绑定的资源 
- 标签列表 
- 批量绑定资源 
- 批量解绑资源 
- 删除标签关联关系 
- 全量更新资源的绑定关系


### 创建标签

使用以下代码可以创建标签
```go
args := &TagsRequest{
    Tags: []Tag{
    {
        TagKey:   "key",
        TagValue: "val",
    }, {
        TagKey:   "key",
        TagValue: "val1",
    },
    },
}

result, err := client.CreateTags(args)
if err != nil {
    fmt.Println("CreateTags failed:", err)
} else {
    fmt.Println("CreateTags success: ", result)
}
```

> **提示：**
> 1.  接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/Okbrb3ral)。


### 删除标签

使用以下代码可以删除标签
```go
args := &TagsRequest{
    Tags: []Tag{
    {
        TagKey:   "key",
        TagValue: "val",
    }, {
        TagKey:   "key",
        TagValue: "val1",
    },
    },
}

result, err := client.DeleteTags(args)
if err != nil {
    fmt.Println("DeleteTags failed:", err)
} else {
    fmt.Println("DeleteTags success: ")
}
```
> **提示：**
> 1.  接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/Xkbrb3rhr)。


### 查看标签下绑定的资源

使用以下代码可以查看标签下绑定的资源
```go
tagKey := "key"
tagValue := "val"
region := "bj"
resourceType := "BCC"

res, err := tagClient.TagsResources(tagKey, tagValue, region, resourceType)
if err != nil {
    fmt.Println("TagsResources failed:", err)
} else {
    fmt.Println("TagsResources success: ", result)
}
```
> **提示：**
> 1.  接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/Bkbrb3roy)。


### 标签列表

使用以下代码可以标签列表
```go
tagKey := "go_sdk_key"
tagValue := "go_sdk_val"
res, err := tagClient.UserTagList(tagKey, tagValue)
if err != nil {
    fmt.Println("UserTagList failed:", err)
} else {
    fmt.Println("UserTagList success: ", result)
}
```
> **提示：**
> 1. 接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/Ukbrb3r3d)。


### 批量绑定资源

使用以下代码可以批量绑定资源
```go
args := &CreateAssociationsByTagRequest{
    TagKey:      "key",
    TagValue:    "val",
    ServiceType: "BCC",
    Resource: []Resource{
        {
            ResourceId:  "58426",
            ServiceType: "BCC",
            Region:      "bj",
        },
        {
            ResourceId:  "58428",
            ServiceType: "BCC",
            Region:      "bj",
        },
    },
}
err := tagClient.CreateAssociationsByTag(args)
if err != nil {
    fmt.Println("CreateAssociationsByTag failed:", err)
} else {
    fmt.Println("CreateAssociationsByTag success: ", result)
}
```
> **提示：**
> 1. 接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/rkm1yqvhz)。


### 批量解绑资源

使用以下代码可以批量解绑资源
```go
args := &DeleteAssociationsByTagRequest{
    TagKey:      "key",
    TagValue:    "val",
    ServiceType: "BCC",
    Resource: []Resource{
        {
            ResourceId:  "58426",
            ServiceType: "BCC",
            Region:      "bj",
        },
        {
            ResourceId:  "58428",
            ServiceType: "BCC",
            Region:      "bj",
        },
    },
}
err := tagClient.DeleteAssociationsByTag(args)
if err != nil {
    fmt.Println("DeleteAssociationsByTag failed:", err)
} else {
    fmt.Println("DeleteAssociationsByTag success: ", result)
}
```
> **提示：**
> 1. 接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/7km1zf2j2)。


### 删除标签关联关系

使用以下代码可以删除标签关联关系
```go
args := &DeleteTagAssociationRequest{
    Resource: &Resource{
        {
            ResourceId:  "58426",
            ServiceType: "BCC",
            Region:      "bj",
        },
        {
            ResourceId:  "58428",
            ServiceType: "BCC",
            Region:      "bj",
        },
    },
}
err := tagClient.DeleteTagAssociation(args)
if err != nil {
    fmt.Println("DeleteTagAssociation failed:", err)
} else {
    fmt.Println("DeleteTagAssociation success: ", result)
}
```
> **提示：**
> 1. 接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/wkm1sfady)。

### 查询用户的标签信息列表

使用以下代码可以查询用户的标签信息列表
```go
var strongAssociation = true
args := &FullTagListRequest{
    TagKey:   "key",
    TagValue: "val",
    Regions: []string{
        "bj",
    },
    ServiceTypes: []string{
        "BCC",
    },
    ResourceIds:   []string{
		"58426",
    },
}
err := tagClient.QueryFullList(args)
if err != nil {
    fmt.Println("QueryFullList failed:", err)
} else {
    fmt.Println("QueryFullList success: ", result)
}
```
> **提示：**
> 1. 接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/blvc6j3w1)。

### 全量更新资源的绑定关系

使用以下代码可以全量更新资源的绑定关系
```go
args := &CreateAssociationsByTagRequest{
    TagKey:      "key",
    TagValue:    "val",
    ServiceType: "BCC",
    Resource: []Resource{
        {
            ResourceId:  "58426",
            ServiceType: "BCC",
            Region:      "bj",
        },
        {
            ResourceId:  "58428",
            ServiceType: "BCC",
            Region:      "bj",
        },
    },
}
err := tagClient.CreateAssociationsByTag(args)
if err != nil {
    fmt.Println("CreateAssociationsByTag failed:", err)
} else {
    fmt.Println("CreateAssociationsByTag success: ", result)
}
```
> **提示：**
> 1. 接口详细描述请参考标签管理(TAG)[API文档](https://cloud.baidu.com/doc/TAG/s/rkm1yqvhz)。


# 错误处理

GO语言以error类型标识错误，标签管理(TAG)支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | 标签管理(TAG)服务返回的错误

用户使用SDK调用标签管理(TAG)相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// tag Client 为已创建的tag Client对象
tagKey := "key"
tagValue := "val"
region := "bj"
resourceType := "BCC"
res, err := tagClient.TagsResources(tagKey, tagValue, region, resourceType)
if err != nil {
	switch realErr := err.(type) {
	case *bce.BceClientError:
		fmt.Println("client occurs error:", realErr.Error())
	case *bce.BceServiceError:
		fmt.Println("service occurs error:", realErr.Error())
	default:
		fmt.Println("unknown error:", err)
	}
} else {
	fmt.Println("TagsResources success: ", result)
}
```

## 客户端异常

客户端异常表示客户端尝试向标签管理(TAG)服务发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

# 版本变更记录

## v0.9.175 [2024-04-10]

首次发布：
创建标签、删除标签、标签列表、查看标签下绑定的资源 、标签列表、批量绑定资源、批量解绑资源、删除标签关联关系、全量更新资源的绑定关系
