# 百舸自定义部署服务

# 概述

本文档主要介绍百舸自定义部署服务 GO SDK的使用。在使用本文档前，您需要先了解百舸自定义部署服务的一些基本知识，并已开通了相关服务。若您还不了解百舸自定义部署服务，可以参考[产品描述](https://cloud.baidu.com/doc/AIHC/s/Dm1j2lqo6)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[服务域名](https://cloud.baidu.com/doc/AIHC/s/qly5ja12q)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

对于百舸自定义服务来说，endpoint都为 aihc.baidubce.com，需要指定区域参数region来对应不同区域的访问，对应信息为：

访问区域	 | 对应Endpoint 
---|---
北京 | bj
广州 | gz
苏州 | su
香港 | hkg
武汉 | fwh
保定 | bd

## 获取密钥

要使用百度云百舸自定义部署服务，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问百舸自定义部署服务做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建 AihcInference Client

AihcInferenceClient是百舸自定义部署服务的客户端，为开发者访问接口提供了一系列的方法。

```go
import (
	"github.com/baidubce/bce-sdk-go/services/aihc/inference"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个AihcInferenceClient
	client, err := inference.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`为用户自己填写指定域名，如果设置为空字符串，会使用默认域名作为自定义部署服务的服务地址。

### 使用STS创建 AihcInference Client

**申请STS token**

自定义部署服务可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问自定义部署服务，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建 AihcInference Client**

申请好STS后，可将STS Token配置到 AihcInference Client中，从而实现通过STS Token创建 AihcInference Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建 AihcInference Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"                    //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/aihc/inference" //导入自定义部署服务模块
	"github.com/baidubce/bce-sdk-go/services/sts"            //导入STS服务模块
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

	// 使用申请的临时STS创建百舸自定义部署服务服务的Client对象，Endpoint使用默认值
	aihcInferenceClient, err := inference.NewClientWithSTS(stsObj.AccessKeyId, stsObj.SecretAccessKey, stsObj.SessionToken, "aihc.baidubce.com")
	if err != nil {
		fmt.Println("create aihc inference client failed:", err)
		return
	}
}
```

> 注意：
> 目前使用STS配置 aihcInference Client时，无论对应的自定义部署服务的Endpoint在哪个区域，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问自定义部署服务

自定义部署服务支持HTTPS传输协议，您可以通过在创建 AihcInference Client对象时指定的Endpoint中指明HTTPS的方式，在AihcInference GO SDK中使用HTTPS访问自定义部署服务：

```go
// import "github.com/baidubce/bce-sdk-go/aihc/inference"

ENDPOINT := "https://aihc.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
AihcInferenceClient, _ := inference.NewClient(AK, SK, ENDPOINT)
```

## 配置AihcInference Client

如果用户需要配置AihcInference Client的一些细节的参数，可以在创建AihcInference Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问百舸自定义部署服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"

//创建AihcInference Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "aihc.baidubce.com"
client, _ := inference.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "aihc.baidubce.com"
client, _ := inference.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "aihc.baidubce.com"
client, _ := inference.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问百舸自定义部署服务时，创建的 AihcInference Client对象的`Config`字段支持的所有参数如下表所示：

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

1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用。
2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。

# 接口文档
## 创建服务
使用以下代码可以创建服务。
```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.CreateApp(&api.CreateAppArgs{
    AppName: "", //这里参数自己按照接口文档填写即可
}, region, nil)

if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[创建服务](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E5%88%9B%E5%BB%BA%E6%9C%8D%E5%8A%A1)

## 查询服务列表
使用以下代码可以查询服务列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ListApp(&api.ListAppArgs{
    PageSize: 10,
    PageNo:   1,
}, region, nil)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询服务列表](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E6%9C%8D%E5%8A%A1%E5%88%97%E8%A1%A8)

## 查询服务状态
使用以下代码可以查询服务状态。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ListAppStats(&api.ListAppStatsArgs{
    AppIds: []string{"ap-test"},
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询服务状态](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E6%9C%8D%E5%8A%A1%E7%8A%B6%E6%80%81)

## 查询服务详情
使用以下代码可以查询服务详情。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.AppDetails(&api.AppDetailsArgs{
    AppId: "ap-test",
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询服务详情](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E6%9C%8D%E5%8A%A1%E8%AF%A6%E6%83%85)

## 更新服务
使用以下代码可以更新服务。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.UpdateApp(&api.UpdateAppArgs{}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[更新服务](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9B%B4%E6%96%B0%E6%9C%8D%E5%8A%A1)

## 扩缩容实例
使用以下代码可以扩缩容实例。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ScaleApp(&api.ScaleAppArgs{
    AppId:    "ap-test",
    InsCount: 1,
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[扩缩容实例](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%89%A9%E7%BC%A9%E5%AE%B9%E5%AE%9E%E4%BE%8B)

## 配置公网访问
使用以下代码可以配置公网访问。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.PubAccess(&api.PubAccessArgs{
    AppId:        "ap-test",
    PublicAccess: false,
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[配置公网访问](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E9%85%8D%E7%BD%AE%E5%85%AC%E7%BD%91%E8%AE%BF%E9%97%AE)

## 查询服务变更记录
使用以下代码可以查询服务变更记录。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ListChange(&api.ListChangeArgs{
    AppId: "ap-test",
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询服务变更记录](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E6%9C%8D%E5%8A%A1%E5%8F%98%E6%9B%B4%E8%AE%B0%E5%BD%95)

## 查询服务变更记录详情
使用以下代码可以查询服务变更记录详情。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ChangeDetail(&api.ChangeDetailArgs{
    ChangeId: "ch-test",
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询服务变更记录详情](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E6%9C%8D%E5%8A%A1%E5%8F%98%E6%9B%B4%E8%AE%B0%E5%BD%95%E8%AF%A6%E6%83%85)

## 删除服务
使用以下代码可以删除服务。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.DeleteApp(&api.DeleteAppArgs{
    AppId: "ap-test",
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[删除服务](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E5%88%A0%E9%99%A4%E6%9C%8D%E5%8A%A1)

## 查询实例列表
使用以下代码可以查询实例列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ListPod(&api.ListPodArgs{
    AppId: "ap-test",
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询实例列表](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E5%AE%9E%E4%BE%8B%E5%88%97%E8%A1%A8)

## 实例摘流
使用以下代码可以实例摘流。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.BlockPod(&api.BlockPodArgs{
    AppId: "ap-test",
    InsID: "ins-test",
    Block: true,
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[实例摘流](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E5%AE%9E%E4%BE%8B%E6%91%98%E6%B5%81)

## 删除实例重建
使用以下代码可以删除实例重建。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.DeletePod(&api.DeletePodArgs{
    AppId: "ap-test",
    InsID: "ins-test",
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[删除实例重建](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E5%88%A0%E9%99%A4%E5%AE%9E%E4%BE%8B%E9%87%8D%E5%BB%BA)

## 查询资源池列表
使用以下代码可以查询资源池列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ListBriefResPool(&api.ListBriefResPoolArgs{
    PageNo:   1,
    PageSize: 10,
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询资源池列表](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E8%B5%84%E6%BA%90%E6%B1%A0%E5%88%97%E8%A1%A8)

## 查询资源池详情
使用以下代码可以查询资源池详情。

```go
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference"
// import "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
region := "bj"

client, _ := inference.NewClient(ak, sk, endpoint)
result, err := client.ResPoolDetail(&api.ResPoolDetailArgs{
    ResPoolId: "id-test",
}, region)
if err != nil {
    panic(err)
}
```

> 注意:
> - 根据接口文档去填写具体的访问参数，接口链接为[查询资源池详情](https://cloud.baidu.com/doc/AIHC/s/Rm2eg9dtl#%E6%9F%A5%E8%AF%A2%E8%B5%84%E6%BA%90%E6%B1%A0%E8%AF%A6%E6%83%85)

# 错误处理

GO语言以error类型标识错误，百舸自定义部署服务支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | 百舸自定义部署服务返回的错误

用户使用SDK调用百舸自定义部署相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// AihcInferenceClient 为已创建的AihcInference Client对象
args := &inference.ListAppArgs{}
result, err := client.ListApp(args)
if err != nil {
	switch realErr := err.(type) {
	case *bce.BceClientError:
		fmt.Println("client occurs error:", realErr.Error())
	case *bce.BceServiceError:
		fmt.Println("service occurs error:", realErr.Error())
	default:
		fmt.Println("unknown error:", err)
	}
} 
```

## 客户端异常

客户端异常表示客户端尝试向服务端发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当服务端出现异常时，服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[错误码](https://cloud.baidu.com/doc/AIHC/s/Ply5jdh5m)

## SDK日志

百舸自定义部署服务 GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

百舸自定义部署服务 GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

```
// import "github.com/baidubce/bce-sdk-go/util/log"

// 指定输出到标准错误，输出INFO及以上级别
log.SetLogHandler(log.STDERR)
log.SetLogLevel(log.INFO)

// 指定输出到标准错误和文件，DEBUG及以上级别，以1GB文件大小进行滚动
log.SetLogHandler(log.STDERR | log.FILE)
log.SetLogDir("/tmp/gosdk-log")
log.SetRotateType(log.ROTATE_SIZE)
log.SetRotateSize(1 << 30)

// 输出到标准输出，仅输出级别和日志消息
log.SetLogHandler(log.STDOUT)
log.SetLogFormat([]string{log.FMT_LEVEL, log.FMT_MSG})
```

说明：
1. 日志默认输出级别为`DEBUG`
2. 如果设置为输出到文件，默认日志输出目录为`/tmp`，默认按小时滚动
3. 如果设置为输出到文件且按大小滚动，默认滚动大小为1GB
4. 默认的日志输出格式为：`FMT_LEVEL, FMT_LTIME, FMT_LOCATION, FMT_MSG`

### 项目使用

该日志模块无任何外部依赖，用户使用GO SDK开发项目，可以直接引用该日志模块自行在项目中使用，用户可以继续使用GO SDK使用的包级别的日志对象，也可创建新的日志对象，详见如下示例：

```
// 直接使用包级别全局日志对象（会和GO SDK自身日志一并输出）
log.SetLogHandler(log.STDERR)
log.Debugf("%s", "logging message using the log package in the AIHC go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the AIHC go sdk")
```

# 版本变更记录

## v0.9.200 [2024-11-13]

首次发布:

- 支持创建服务、获取服务列表、获取服务状态、查询服务详情信息、更新服务、扩缩容服务、配置访问公网、获取服务变更记录、获取变更详情、删除服务、获取实例列表、摘除实例流量、删除实例触发重建、获取资源池列表、获取指定资源池的详情等接口。