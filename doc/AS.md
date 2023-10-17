# AS服务

# 概述

本文档主要介绍AS GO SDK的使用。在使用本文档前，您需要先了解

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[AS域名](https://cloud.baidu.com/doc/AS/s/fk3imcz57)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云AS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问AS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建AS Client

AS Client是AS控制面服务的客户端，为开发者与AS控制面服务进行交互提供了一系列的方法。

### 使用AK/SK新建AS Client

通过AK/SK方式访问AS，用户可以参考如下代码新建一个AS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/as"   //导入AS服务模块
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个ASClient
	asClient, err := as.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/as/index.html)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为AS的控制面服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`as.bj.baidubce.com`。

### 使用STS创建AS Client

**申请STS token**

AS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问AS，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建AS Client**

申请好STS后，可将STS Token配置到AS Client中，从而实现通过STS Token创建AS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建AS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/as"    //导入AS服务模块
	"github.com/baidubce/bce-sdk-go/services/sts"    //导入STS服务模块
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

	// 使用申请的临时STS创建AS控制面服务的Client对象，Endpoint使用默认值
	asClient, err := as.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create as client failed:", err)
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
	asClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置AS Client时，无论对应AS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问AS

AS支持HTTPS传输协议，您可以通过在创建AS Client对象时指定的Endpoint中指明HTTPS的方式，在AS GO SDK中使用HTTPS访问AS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/as"

ENDPOINT := "https://as.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
asClient, _ := as.NewClient(AK, SK, ENDPOINT)
```

## 配置AS Client

如果用户需要配置AS Client的一些细节的参数，可以在创建AS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问AS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/as"

//创建AS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "as.bj.baidubce.com
client, _ := as.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/AS"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "as.bj.baidubce.com"
client, _ := as.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/AS"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "as.bj.baidubce.com"
client, _ := as.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问AS时，创建的AS Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建AS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

百度智能云弹性伸缩(Auto Scaling)是自动化扩缩容用户云资源的管理服务，当您业务所需的云资源用量经常性变化时，弹性伸缩会是您使用云资源的理想方式。

## 查询伸缩组列表接口

### 接口描述
可查询所有伸缩组的详细信息。

### 请求示例

```go
req := &as.ListAsGroupRequest{
	// 可选，伸缩组名称
        GroupName: "as-Group-Name",
        // 可选，批量获取列表的查询的起始位置，是一个由系统生成的字符串
        Marker:    "marker",
        // 可选，每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
        MaxKeys:   100,
}
resp, err := asClient.ListAsGroup(req)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考AS API 文档[查询伸缩组列表](https://cloud.baidu.com/doc/AS/s/hk3imj0oq#%E6%9F%A5%E8%AF%A2%E4%BC%B8%E7%BC%A9%E7%BB%84%E5%88%97%E8%A1%A8)

## 查询伸缩组详情接口

### 接口描述
可查询单个伸缩组的详细信息。

### 请求示例
```go
req := &as.GetAsGroupRequest{
        // 必填，待查询的伸缩组ID 
	GroupId: "asg-wqksXo95",
}
resp, err := asClient.GetAsGroup(req)
```

## 查询伸缩组下节点列表

### 接口描述
可查询指定伸缩组下节点的详细信息。

### 请求示例
```go
req := &as.ListAsNodeRequest{
        // 必填，伸缩组ID
        GroupId: "asg-wqksXo95",
	// 可选，批量获取列表的查询的起始位置，是一个由系统生成的字符串
	Marker:    "marker",
	// 可选，每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
	MaxKeys:   100,
}
resp, err := asClient.ListAsNode(req)
```

## 伸缩组扩容

### 接口描述
在指定伸缩组下添加节点。

### 请求示例
```go
req := &as.IncreaseAsGroupRequest{
    // 必填，伸缩组ID
    GroupId:   "asg-Hhm2ucIK",
    // 必填，扩容可指定可用区（扩容时会与伸缩组配置的可用区取交集）
    Zone:      []string{"zoneB"},
    // 扩容节点数量
    NodeCount: 1,
    // 扩容时的可用区选择策略
    // Priority - 以单独可用区进行创建
    // Balanced - 在选定可用区中均衡创建
    ExpansionStrategy:"Priority"
}
err := asClient.IncreaseAsGroup(req)
```

## 伸缩组缩容

### 接口描述
用于伸缩组下节点的缩容。

### 请求示例
```go
req := &as.DecreaseAsGroupRequest{
    // 必填，伸缩组ID
    GroupId: "asg-Hhm2ucIK",
	// 必填，手动缩容指定的实例短Id
    Nodes:   []string{"i-z0PXqFD3"},
}
err := asClient.DecreaseAsGroup(req)
```


## 客户端异常

客户端异常表示客户端尝试向AS发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当AS服务端出现异常时，AS服务端会返回给用户相应的错误信息，以便定位问题

