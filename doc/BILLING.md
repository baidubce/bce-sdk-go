# Billing服务

# 概述

本文档主要介绍Billing GO SDK的使用。在使用本文档前，您需要先了解Billing/Finance的一些基本知识。若您还不了解Billing/Finance，可以参考[产品描述](https://cloud.baidu.com/doc/Finance/s/ekps2o2mj)和[常见问题](https://cloud.baidu.com/doc/Finance/s/Vk3i5qy0p)。

# 初始化

## 确认Endpoint

目前使用Billing服务时，CDN的Endpoint都统一使用`https://billing.baidubce.com`，这也是默认值。

## 获取密钥

要使用百度云Billing服务，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问Billing做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/#/iam/accesslist)

## 新建Billing Client

Billing Client是Billing服务的客户端，为开发者与Billing服务进行交互提供了一系列的方法。

### 使用AK/SK新建Billing Client

通过AK/SK方式访问Billing，用户可以参考如下代码新建一个Billing Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/billing"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := <your-access-key-id>, <your-secret-access-key>

	// 初始化一个IAMClient
	billingClient, err := billing.NewClient(AK, SK)
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/IAM/s/ojwvynrqn)》。

### 使用STS创建Billing Client

**申请STS token**

Billing可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问Billing，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建Billing Client**

申请好STS后，可将STS Token配置到Billing Client中，从而实现通过STS Token创建Billing Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建IAM Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/billing"    //导入billing服务模块
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

	// 使用申请的临时STS创建IAM服务的Client对象，Endpoint使用默认值
	billingClient, err := billing.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create billing client failed:", err)
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
	billingClient.Config.Credentials = stsCredential
}
```

## 配置Billing Client

如果用户需要配置Billing Client的一些细节的参数，可以在创建Billing Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/billing"

AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := billing.NewClient(AK, SK)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/billing"

AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := billing.NewClient(AK, SK)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问Billing时，创建的Billing Client对象的`Config`字段支持的所有参数如下表所示：

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

1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建IAM Client”小节。
2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。

# 账单接口

## 资源月账单接口 ResourceMonthBill

### 请求参数说明

名称  | 类型    |   参数位置    | 描述                                                                                                                               |   是否必须
--------|-------|----------|----------------------------------------------------------------------------------------------------------------------------------|--------------
month| String |Query参数| 按月查询账单所属月份，格式yyyy-MM，例：2019-02                                                                                                   |可选
beginTime| String |Query参数| 按天查询账单所属时间区间最早时间，格式yyyy-MM-dd，例：2019-02-01                                                                                       |可选
endTime| String |Query参数| 按天查询账单所属时间区间最晚时间，格式yyyy-MM-dd，例：2019-02-18。若按天查询参数非空，优先返回按天查询的结果；按天查询时间区间不支持跨月；按天查询可以查询本月今天之前的账单；若需查询昨日的账单，为保证数据准确处理完毕，请当日十点后查询。 |可选
productType| String |Query参数| 计费类型：prepay/ postpay，分别表示预付费/后付费                                                                                                 |必须
serviceType| String | Query参数 | 产品类型，例：BCC，BOS等                                                                                                                  |可选
queryAccountId|String|Query参数| 查询账户ID，只有企业组织的主账户可以查询加入财务圈的子账户账单，其他查询场景会提示AccessDenied                                                                           |可选
pageNo|int|Query参数| 分页查询的页数，默认为1                                                                                                                     |可选
pageSize|int|Query参数| 每页包含的最大数量，最大数量不超过100，缺省值为20。	                                                                                                    |可选

### 请求示例
```go
// import "github.com/baidubce/bce-sdk-go/services/billing"
AK, SK := <your-access-key-id>, <your-secret-access-key>
	client, err := billing.NewClient(AK, SK)
	if err != nil {
		fmt.Println("create billing client failed:", err)
		return
	}
	bill, err := client.ResourceMonthBill("2024-02", "", "", "postpay", "", "", 1, 10)
```
