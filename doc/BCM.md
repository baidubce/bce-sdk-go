# BCM服务

# 概述

本文档主要介绍BCM GO SDK的使用。在使用本文档前，您需要先了解
[BCM的基本知识](https://cloud.baidu.com/doc/BCM/index.html)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BCM域名](https://cloud.baidu.com/doc/BCM/s/5jwvym49g)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云BCM，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BCM做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建BCM Client

BCM Client是BCM控制面服务的客户端，为开发者与BCM控制面服务进行交互提供了一系列的方法。

### 使用AK/SK新建BCM Client

通过AK/SK方式访问BCM，用户可以参考如下代码新建一个BCM Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bcm"   //导入BCM服务模块
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BcmClient
	bcmClient, err := bcm.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/BCM/index.html)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BCM的控制面服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`bcm.bj.baidubce.com`。

### 使用STS创建BCM Client

**申请STS token**

BCM可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问BCM，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建BCM Client**

申请好STS后，可将STS Token配置到BCM Client中，从而实现通过STS Token创建BCM Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建BCM Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/bcm"    //导入BCM服务模块
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

	// 使用申请的临时STS创建BCM控制面服务的Client对象，Endpoint使用默认值
	bcmClient, err := bcm.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create bcm client failed:", err)
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
	bcmClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置BCM Client时，无论对应BCM服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问BCM

BCM支持HTTPS传输协议，您可以通过在创建BCM Client对象时指定的Endpoint中指明HTTPS的方式，在BCM GO SDK中使用HTTPS访问BCM服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

ENDPOINT := "https://bcm.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
bcmClient, _ := bcm.NewClient(AK, SK, ENDPOINT)
```

## 配置BCM Client

如果用户需要配置BCM Client的一些细节的参数，可以在创建BCM Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问BCM服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

//创建BCM Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcm.bj.baidubce.com
client, _ := bcm.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcm.bj.baidubce.com"
client, _ := bcm.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcm.bj.baidubce.com"
client, _ := bcm.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问BCM时，创建的BCM Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建BCM Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

云监控BCM（Cloud Monitor），是针对于云平台的监控服务，包括云产品监控、站点监控、自定义监控、应用监控、报警管理等产品功能，实时监控您云平台中的各种资源。

## 查询数据接口

### 接口描述
获取指定指标的一个或多个统计数据的时间序列数据。可获取云产品监控数据、站点监控数据或您推送的自定义监控数据。

### 接口限制
一次返回的数据点数目不能超过1440个。

### 请求示例

```go
dimensions := map[string]string{
    "InstanceId": "xxx",
}
req := &bcm.GetMetricDataRequest{
    UserId:         "xxx",
    Scope:          "BCE_BCC",
    MetricName:     "vCPUUsagePercent",
    Dimensions:     dimensions,
    Statistics:     strings.Split(bcm.Average, ","),
    PeriodInSecond: 60,
    StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
    EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
}
resp, err := bcmClient.GetMetricData(req)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BCM API 文档[查询数据接口](https://cloud.baidu.com/doc/BCM/s/9jwvym3kb)

### 批量查询数据接口

### 接口描述
可获取批量实例数据的接口，支持多维度多指标查找。可获取云产品监控数据、站点监控数据或您推送的自定义监控数据。

### 接口限制
一个实例的任意一个指标一次返回的数据点数目不能超过1440个。

### 请求示例
```go
dimensions := map[string]string{
    "InstanceId": "xxx",
}
req := &bcm.BatchGetMetricDataRequest{
    UserId:         "xxx",
    Scope:          "BCE_BCC",
    MetricNames:    []string{"vCPUUsagePercent", "CpuIdlePercent"},
    Dimensions:     dimensions,
    Statistics:     strings.Split(bcm.Average + "," + bcm.Sum, ","),
    PeriodInSecond: 60,
    StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
    EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
}
resp, err := bcmClient.BatchGetMetricData(req)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BCM API 文档[查询数据接口](https://cloud.baidu.com/doc/BCM/s/9jwvym3kb)



## 客户端异常

客户端异常表示客户端尝试向BCM发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当BCM服务端出现异常时，BCM服务端会返回给用户相应的错误信息，以便定位问题

