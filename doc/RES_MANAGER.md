# 资源管理服务

# 概述

本文档主要介绍资源管理 GO SDK的使用。在使用本文档前，您需要先了解资源管理的一些基本知识。若您还不了解资源管理，可以参考[产品描述](https://cloud.baidu.com/doc/ResManagement/s/eklm4lri0)和[操作指南](https://cloud.baidu.com/doc/ResManagement/s/eklq4p17g)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，资源管理服务不分区域，使用全局域名，可先阅读开发人员指南中关于[API参考-域名](https://cloud.baidu.com/doc/ResManagement/s/Llth1oso4)的部分。

## 获取密钥

要使用百度云资源管理，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问资源管理做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建资源管理 Client

资源管理 Client是资源管理服务的客户端，为开发者与资源管理服务进行交互提供了一系列的方法。

### 使用AK/SK新建resmanager Client

通过AK/SK方式访问资源管理服务，用户可以参考如下代码新建一个resmanager Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/resmanager"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个resmanagerClient
    resmanagerClient, err := resmanager.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Access Key Secret”。资源管理为全局服务，第三个参数`ENDPOINT`，直接设置为空字符串，会使用默认域名作为资源管理的服务地址。

### 使用STS创建资源管理 Client

**申请STS token**

资源管理可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问资源管理，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建resmanager Client**

申请好STS后，可将STS Token配置到resmanager Client中，从而实现通过STS Token创建resmanager Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建resmanager Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/resmanager" //导入资源管理服务模块
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

	// 使用申请的临时STS创建资源管理服务的Client对象，Endpoint使用默认值 
	resmanagerClient, err := resmanager.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create resmanager client failed:", err)
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
    resmanagerClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置资源管理 Client时，无论对应资源管理服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问资源管理

资源管理支持HTTPS传输协议，您可以通过在创建resmanager Client对象时指定的Endpoint中指明HTTPS的方式，在资源管理 GO SDK中使用HTTPS访问资源管理服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/resmanager"

ENDPOINT := "https://resourcemanager.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
resmanagerClient, _ := resmanager.NewClient(AK, SK, ENDPOINT)
```

## 配置资源管理 Client

如果用户需要配置resmanager Client的一些细节的参数，可以在创建resmanager Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问资源管理服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/resmanager"

//创建资源管理 Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "resourcemanager.baidubce.com"
client, _ := resmanager.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/resmanager"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "resourcemanager.baidubce.com"
client, _ := resmanager.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/resmanager"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "resourcemanager.baidubce.com"
client, _ := resmanager.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问资源管理时，创建的资源管理 Client对象的`Config`字段支持的所有参数如下表所示：

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

1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建资源管理 Client”小节。
2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

资源管理API对于每个HTTP请求，认证签名放在Authorization头域中，后端统一认证。且后端在Response头域中会添加x-bce-request-id，作为请求唯一标识，方便追踪定位问题。

目前支持的资源管理API接口如下：

- 资源加入资源分组
- 资源从资源分组移除
- 变更资源绑定的资源分组
- 查询资源分组列表
- 资源ID查询资源资源分组
- 创建资源分组

### 资源加入资源分组

使用以下代码可以将资源加入资源分组
```go
args := &BindResourceToGroupArgs{
	ResourceBrief:[]ResourceBrief{
            {
                // 资源ID
                ResourceId:"123.org",
                // 资源类型
                ResourceType:"CDN", 
                // 资源区域
                ResourceRegion: "global", 
                // 资源分组ID
                GroupId:"group-123",
            },
        },
    }

result, err := client.BindResourceToGroup(args)
if err != nil {
    fmt.Println("BindResourceToGroup failed:", err)
} else {
    fmt.Println("BindResourceToGroup success: ", result)
}
```

> **提示：**
> 1.  接口详细描述请参考资源管理[API文档](https://cloud.baidu.com/doc/ResManagement/s/ilth0vmb9)。

### 资源从资源分组移除

使用以下代码可以将资源从资源分组移除
```go
args := &BindResourceToGroupArgs{
	ResourceBrief:[]ResourceBrief{
            {
                // 资源ID
                ResourceId:"123.org",
                // 资源类型
                ResourceType:"CDN", 
                // 资源区域
                ResourceRegion: "global", 
                // 资源分组ID
                GroupId:"group-123",
            },
        },
    }

result, err := client.RemoveResourceFromGroup(args)
if err != nil {
    fmt.Println("RemoveResourceFromGroup failed:", err)
} else {
    fmt.Println("RemoveResourceFromGroup success: ")
}
```

> **提示：**
> 1.  接口详细描述请参考资源管理[API文档](https://cloud.baidu.com/doc/ResManagement/s/ilth0vmb9)。

### 变更资源绑定的资源分组

使用以下代码变更资源绑定的资源分组
```go
args := &ChangeResourceGroupArgs{
            MoveResModels:[]MoveResModels{
            {
                // 目标分组
                TargetGroupId: "group-456",
                // 原有的绑定关系信息
                OldGroupResInfo: OldGroupResInfo{
                // 资源ID 
                ResourceId:     "123.org",
                // 资源类型
                ResourceType:   "CDN",
                // 资源区域
                ResourceRegion: "global",
                // 资源分组ID
                GroupId:        "group-123",
				},
            },
        },
    }

result, err := client.ChangeResourceGroup(args)
if err != nil {
    fmt.Println("ChangeResourceGroup failed:", err)
} else {
    fmt.Println("ChangeResourceGroup success: ", result)
}
```

> **提示：**
> 1.  接口详细描述请参考资源管理[API文档](https://cloud.baidu.com/doc/ResManagement/s/ilth0vmb9)。


### 查询资源分组列表


使用以下代码可以查询资源资源分组
```go
// 资源分组名称
name := "test"

result, err := client.QueryGroupList(name)
if err != nil {
    fmt.Println("QueryGroupList failed:", err)
} else {
    fmt.Println("QueryGroupList success: ", result)
}
```

> **提示：**
> 1. 接口详细描述请参考资源管理[API文档](https://cloud.baidu.com/doc/ResManagement/s/ilth0vmb9)。

### 资源ID查询资源资源分组

使用以下代码可以查询资源资源分组
```go
args := &ResGroupDetailRequest{
	ResourceBrief:[]ResourceBrief{
            {
            // 资源ID
            ResourceId:"123.org",
            // 资源类型
            ResourceType:"CDN", 
            // 资源区域
            ResourceRegion: "global",
            },
        },
    }

result, err := client.GetResGroupBatch(args)
if err != nil {
    fmt.Println("GetResGroupBatch failed:", err)
} else {
    fmt.Println("GetResGroupBatch success: ", result)
}
```

> **提示：**
> 1. 资源ID查询资源分组接口支持批量查询，一次最多支持查询2000个资源。 
> 2. 仅能查询当前账号资源，也就是根据AK, SK签名解析的账号下的资源。
> 3. 接口详细描述请参考资源管理[API文档](https://cloud.baidu.com/doc/ResManagement/s/ilth0vmb9)。

### 创建资源分组

使用以下代码创建资源分组
```go
args := &CreateResourceGroupArgs{
    // 资源组名称
    Name: "资源组名称",
    // 资源组的备注
    Extra: "备注",
}

result, err := client.CreateResourceGroup(args)
if err != nil {
    fmt.Println("CreateResourceGroup failed:", err)
} else {
    fmt.Println("CreateResourceGroup success: ", result)
}
```

> **提示：**
> 1. 接口详细描述请参考资源管理[API文档](https://cloud.baidu.com/doc/ResManagement/s/ilth0vmb9)。
> 2. 资源组的名称，同一用户下不能重复，支持中英文及常见符号-_ /.，1～20个字符，必填。
> 3. 资源组的备注，可以为空。最多存储 256 个字符的文本。

# 错误处理

GO语言以error类型标识错误，资源管理支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | 资源管理服务返回的错误

用户使用SDK调用资源管理相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// resmanager Client 为已创建的resmanager Client对象
result, err := resmanagerClient.QueryGroupList(name)
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
	fmt.Println("QueryGroupList success: ", result)
}
```

## 客户端异常

客户端异常表示客户端尝试向资源管理服务发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

# 版本变更记录

## v0.9.175 [2024-04-10]

首次发布：
- 资源加入资源分组、资源从资源分组移除、变更资源绑定的资源分组、查询资源分组列表
- 资源ID查询资源资源分组

## v0.9.207 [2024-12-16]
- 创建资源分组