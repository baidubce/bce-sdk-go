# CCR服务 企业版

# 概述

本文档主要介绍CCR企业版 GO SDK的使用。在使用本文档前，您需要先了解CCR的一些基本知识，并已开通了CCR服务。若您还不了解CCR，可以参考[产品描述](https://cloud.baidu.com/doc/CCR/s/qk8gwqs4a)和[操作指南](https://cloud.baidu.com/doc/CCR/s/skw63yms7)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[CCR服务域名](https://cloud.baidu.com/doc/CCR/s/Fjwvy1fl4)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/CCR/s/Fjwvy1fl4)。

目前支持“华北-北京”、“华南-广州”、“华东-苏州”、“香港”、“金融华中-武汉”和“华北-保定”六个区域。对应信息为：

访问区域 | 对应Endpoint           | 协议
---|----------------------|---
BJ | ccr.bj.baidubce.com  | HTTP and HTTPS
GZ | ccr.gz.baidubce.com  | HTTP and HTTPS
SU | ccr.su.baidubce.com  | HTTP and HTTPS
HKG| ccr.hkg.baidubce.com | HTTP and HTTPS
FWH| ccr.fwh.baidubce.com | HTTP and HTTPS
BD | ccr.bd.baidubce.com  | HTTP and HTTPS

## 获取密钥

要使用百度云CCR，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问CCR做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建CCR Client

CCR企业版 Client是CCR服务的客户端，为开发者与CCR服务进行交互提供了一系列的方法。

### 使用AK/SK新建CCR Client

通过AK/SK方式访问CCR，用户可以参考如下代码新建一个CCR Client：
```go
import (
	"github.com/baidubce/bce-sdk-go/services/eccr"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := <your-access-key-id>, <your-secret-access-key>

	//用户指定的endpoint 
	ENDPOINT := "endpoint"
    
	// 初始化一个CCRClient
	ccrClient, err := eccr.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为CCR的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`ccr.bj.baidubce.com`。

### 使用STS创建CCR Client

**申请STS token**

CCR可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问CCR，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建CCR Client**

申请好STS后，可将STS Token配置到CCR Client中，从而实现通过STS Token创建CCR Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建CCR Client对象：
```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/eccr" //导入CCR服务模块
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

	// 使用申请的临时STS创建CCR服务的Client对象，Endpoint使用默认值
	ccrClient, err := eccr.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "ccr.bj.baidubce.com")
	if err != nil {
		fmt.Println("create ccr client failed:", err)
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
	ccrClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置CCR Client时，无论对应CCR服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问CCR

CCR支持HTTPS传输协议，您可以通过在创建CCR Client对象时指定的Endpoint中指明HTTPS的方式，在CCR GO SDK中使用HTTPS访问CCR企业版服务：
```go
// import "github.com/baidubce/bce-sdk-go/services/ccr"
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "https://ccr.bj.baidubce.com" //指明使用HTTPS协议

ccrClient, _ := eccr.NewClient(AK, SK, ENDPOINT)
```

## 配置CCR Client

如果用户需要配置CCR Client的一些细节的参数，可以在创建CCR Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问CCR服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/ccr"

//创建CCR Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ccr.bj.baidubce.com"

ccrClient, _ := eccr.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
ccrClient.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/ccr"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ccr.bj.baidubce.com"

ccrClient, _ := eccr.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
ccrClient.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
ccrClient.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/ccr"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ccr.bj.baidubce.com"

ccrClient, _ := eccr.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
ccrClient.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
ccrClient.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问CCR时，创建的CCR Client对象的`Config`字段支持的所有参数如下表所示：

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

1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建CCR Client”小节。
2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# CCR管理

百度智能云容器镜像服务（Cloud Container Registry，简称CCR）是面向容器镜像、Helm Chart等符合OCI规范的云原生制品安全托管以及高效分发平台。CCR支持在多个地域创建独享托管服务，具备多种安全保障；支持同步容器镜像等云原生制品，与容器引擎CCE等服务无缝集成，助力企业提升云原生容器应用交付效率。

> 注意:
> - 企业版实例托管的云原生应用制品（如容器镜像、Helm Chart）存储在您的 BOS Bucket 中，根据实际使用情况将产生存储和流量费用。

## 列举CCR实例
使用以下代码可以列举CCR实例列表。
```go
args := &ListInstancesArgs{
		KeywordType: "clusterName",
		Keyword:     "",
		PageNo:      1,
		PageSize:    10,
	}

resp, err := ccrClient.ListInstances(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 查询单个CCR实例详情
列举CCR实例列表。
```go
instanceID := "instance-id"
resp, err := ccrClient.GetInstanceDetail(instanceID)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 获取私有网络列表
查询私有网络列表。
```go
instanceID := "instance-id"
resp, err := ccrClient.ListPrivateNetworks(instanceID)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```