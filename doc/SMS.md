# SMS简单消息服务

# 概述
本文档主要介绍普通型SMS GO SDK的使用。在使用本文档前，您需要先了解普通型SMS的一些基本知识。若您还不了解普通型SMS，可以参考[产品描述](https://cloud.baidu.com/doc/SMS/s/0jwvxrjyt)和[入门指南](https://cloud.baidu.com/doc/SMS/s/sk4m6mxke)。

# 初始化
## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[SMS访问域名](https://cloud.baidu.com/doc/SMS/s/pjwvxrw6w)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云SMS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问SMS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建SMS Client

SMS Client是SMS服务的客户端，为开发者与SMS服务进行交互提供了一系列的方法。

### 使用AK/SK新建SMS Client

通过AK/SK方式访问SMS，用户可以参考如下代码新建一个SMS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/sms"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个SmsClient
	smsClient, err := sms.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/SMS/GettingStarted.html#.E7.AE.A1.E7.90.86ACCESSKEY) 》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为SMS的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`http://smsv3.bj.baidubce.com`。

### 使用STS创建SMS Client

**申请STS token**

SMS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问SMS，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7) 。

**用STS token新建SMS Client**

申请好STS后，可将STS Token配置到SMS Client中，从而实现通过STS Token创建SMS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建SMS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/sms" //导入SMS服务模块
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

	// 使用申请的临时STS创建SMS服务的Client对象，Endpoint使用默认值
	smsClient, err := sms.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create sms client failed:", err)
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
	smsClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置SMS Client时，无论对应SMS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com 。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问SMS

SMS支持HTTPS传输协议，您可以通过在创建SMS Client对象时指定的Endpoint中指明HTTPS的方式，在SMS GO SDK中使用HTTPS访问SMS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/sms"

ENDPOINT := "https://smsv3.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
smsClient, _ := sms.NewClient(AK, SK, ENDPOINT)
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/sms"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "smsv3.bj.baidubce.com"
client, _ := sms.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/sms"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "smsv3.bj.baidubce.com"
client, _ := sms.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问SMS时，创建的SMS Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建SMS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。

# 主要接口
## 发送消息

通过以下代码可以发送SMS消息
```go
	contentMap := make(map[string]interface{})
	contentMap["code"] = "123"
	contentMap["minute"] = "1"
	sendSmsArgs := &api.SendSmsArgs{
		Mobile:      "13800138000",
		Template:    "your template id",
		SignatureId: "your signature id",
		ContentVar:  contentMap,
	}
	result, err := client.SendSms(sendSmsArgs)
	if err != nil {
		fmt.Printf("send sms error, %s", err)
		return
	}
	fmt.Printf("send sms success. %s", result)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考SMS API 文档[短信下发](https://cloud.baidu.com/doc/SMS/s/Ykdtz1h96)

## 签名
### 申请签名
通过以下代码，可以申请一个SMS签名
```go
	// Open file on disk.
	f, _ := os.Open("/dir1/dir2/your_sign_pic.png")
	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	result, err := client.CreateSignature(&api.CreateSignatureArgs{
		Content:             "Baidu",
		ContentType:         "Enterprise",
		Description:         "test",
		CountryType:         "DOMESTIC",
		SignatureFileBase64: encoded,
		SignatureFileFormat: "png",
	})
	if err != nil {
		fmt.Printf("create signature error, %s", err)
		return
	}
	fmt.Printf("create signature success. %s", result)
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[申请签名](https://cloud.baidu.com/doc/SMS/s/Pkdtz1os3)

### 获取签名详情
通过以下代码，可以获取一个SMS签名详情
```go
	result, err := client.GetSignature(&api.GetSignatureArgs{
		SignatureId: "your signature id",
	})
	if err != nil {
		fmt.Printf("get signature error, %s", err)
		return
	}
	fmt.Printf("get signature success. %s", result)
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[获取签名详情](https://cloud.baidu.com/doc/SMS/s/Pkdtz1os3)

### 变更签名申请
通过以下代码，可以变更一个SMS签名申请
```go
	// Open file on disk.
	f, _ := os.Open("/dir1/dir2/your_sign_pic.png")
	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	err := client.ModifySignature(&api.ModifySignatureArgs{
		SignatureId:         "your signature id",
		Content:             "Baidu",
		ContentType:         "MobileApp",
		Description:         "this is a test",
		CountryType:         "INTERNATIONAL",
		SignatureFileBase64: encoded,
		SignatureFileFormat: "png",
	})
	if err != nil {
		fmt.Printf("modify signature error, %s", err)
		return
	}
	fmt.Printf("modify signature success.")
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[变更签名申请](https://cloud.baidu.com/doc/SMS/s/Pkdtz1os3)

### 删除签名
通过以下代码，可以删除一个SMS签名
```go
	err := client.DeleteSignature(
		&api.DeleteSignatureArgs{SignatureId: "your signature id"})
	if err != nil {
		fmt.Printf("delete signature error, %s", err)
		return
	}
	fmt.Printf("delete signature success.")
```

> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[删除签名](https://cloud.baidu.com/doc/SMS/s/Pkdtz1os3)

## 模板
### 申请模板
通过以下代码，可以申请一个sms模板
```go
	result, err := client.CreateTemplate(&api.CreateTemplateArgs{
		Name:        "my template",
		Content:     "${content}",
		SmsType:     "CommonNotice",
		CountryType: "DOMESTIC",
		Description: "this is a test",
	})
	if err != nil {
		fmt.Printf("create template error, %s", err)
		return
	}
	fmt.Printf("create template success. %s", result)
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[申请模板](https://cloud.baidu.com/doc/SMS/s/jkdtz1vy6)

### 获取模板详情
通过以下代码，可以获取一个sms模板详情
```go
	result, err := client.GetTemplate(&api.GetTemplateArgs{TemplateId: your template id"})
	if err != nil {
		fmt.Printf("get template error, %s", err)
		return
	}
	fmt.Printf("get template success. %s", result)
```

> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[获取模板详情](https://cloud.baidu.com/doc/SMS/s/jkdtz1vy6)

### 变更模板
通过以下代码，可以变更一个sms模板申请
```go
	err := client.ModifyTemplate(&api.ModifyTemplateArgs{
		TemplateId:  "your template id",
		Name:        "my template",
		Content:     "${code}",
		SmsType:     "CommonVcode",
		CountryType: "GLOBAL",
		Description: "this is a test",
	})
	if err != nil {
		fmt.Printf("modify template error, %s", err)
		return
	}
	fmt.Printf("modify template success.")
```

> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[变更模板](https://cloud.baidu.com/doc/SMS/s/jkdtz1vy6)

### 删除模板
通过以下代码，可以删除一个模板
```go
	err := client.DeleteTemplate(&api.DeleteTemplateArgs{TemplateId: "your template id"})
	if err != nil {
		fmt.Printf("delete template error, %s", err)
		return
	}
	fmt.Printf("delete template success.")
```

> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[删除模板](https://cloud.baidu.com/doc/SMS/s/jkdtz1vy6)

## 配额频控
### 查看配额和频控
通过以下代码，可以查看配额和频控
```go
	result, err := client.QueryQuotaAndRateLimit()
	if err != nil {
		fmt.Printf("query quota or rate limit error, %s", err)
		return
	}
	fmt.Printf("query quota or rate limit success. %s", result)
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[查看配额和频控](https://cloud.baidu.com/doc/SMS/s/Wkdtzv7t9)

### 变更配额或频控
通过以下代码，可以变更配额或频控
```go
	err := client.UpdateQuotaAndRateLimit(&api.UpdateQuotaRateArgs{
		QuotaPerDay:        200,
		QuotaPerMonth:      200,
		RateLimitPerDay:    60,
		RateLimitPerHour:   20,
		RateLimitPerMinute: 10,
	})
	if err != nil {
		fmt.Printf("update quota or rate limit template error, %s", err)
		return
	}
	fmt.Printf("update quota or rate limit success")
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[变更配额和频控](https://cloud.baidu.com/doc/SMS/s/Wkdtzv7t9)

## 手机号黑名单
### 创建手机号黑名单
通过以下代码，可以创建手机号黑名单
```go
	err := client.CreateMobileBlack(&api.CreateMobileBlackArgs{
            Type:                "MerchantBlack",
            SmsType:             "CommonNotice",
            SignatureIdStr:      "sddd",
            Phone:               "12345678901",
	})
	if err != nil {
		fmt.Printf("CreateMobileBlack error, %s", err)
		return
	}
	fmt.Printf("CreateMobileBlack success")
```

### 查询手机号黑名单
通过以下代码，可以查询手机号黑名单
```go
	err := client.GetMobileBlack(&api.GetMobileBlackArgs{
            SmsType:            "CommonNotice",
            SignatureIdStr:     "sddd",
            Phone:              "12345678901",
            StartTime:          "2023-07-18",
            EndTime:            "2023-07-19",
            PageNo:             "1",
            PageSize:           "10",
	})
	if err != nil {
		fmt.Printf("GetMobileBlack error, %s", err)
		return
	}
	fmt.Printf("GetMobileBlack success")
```

### 删除手机号黑名单
通过以下代码，可以删除手机号黑名单
```go
	err := client.DeleteMobileBlack(&api.DeleteMobileBlackArgs{
            Phones:   "12345678901",
	})
	if err != nil {
		fmt.Printf("DeleteMobileBlack error, %s", err)
		return
	}
	fmt.Printf("DeleteMobileBlack success")
```

> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[手机号黑名单](https://cloud.baidu.com/doc/SMS/s/wlk8bli1z)

## 业务统计
### 获取下发短信的业务统计信息
通过以下代码，可以查询短信的业务统计信息
```go
	err := client.ListStatistics(&api.ListStatisticsArgs{
            SmsType:               "normal", 
            StartTime:             "2023-10-01",
            EndTime:               "2023-11-01",
	})
	if err != nil {
		fmt.Printf("ListStatistics error, %s", err)
		return
	}
	fmt.Printf("ListStatistics success")
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[业务统计](https://cloud.baidu.com/doc/SMS/s/8lohz4ey0)


## 量包查询
### 获取商户短信流量包查询
通过以下代码，可以查询商户短信流量包详情
```go
	err := client.GetPrepaidPackages(&api.GetPrepaidPackageArgs{
            UserID:                "your_userid",
	})
	if err != nil {
		fmt.Printf("GetPrepaidPackages error, %s", err)
		return
	}
	fmt.Printf("GetPrepaidPackages success")
```
> **提示：**
> - 详细参数配置及限制条件，可以参考SMS API 文档[量包查询](https://cloud.baidu.com/doc/SMS/s/Yjwvxrwzb)



# 错误处理

GO语言以error类型标识错误，SMS支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | SMS服务返回的错误

用户使用SDK调用SMS相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```go
// smsClient 为已创建的SMS Client对象
smsDetail, err := smsClient.SendSms(args)
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
	fmt.Println("send sms success: ", smsDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向SMS发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当SMS服务端出现异常时，SMS服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[短信发送接口响应码](https://cloud.baidu.com/doc/SMS/s/zjwvxry6e)

# 版本变更记录

## v0.9.19 [2020-08-08]

首次发布：
- 支持短信发送接口
- 支持签名、模板管理接口
- 支持配额、频控查看和变更

## v0.9.32 [2023-07-20]
- 配额频控查询增加申请信息字段
- 增加手机号黑名单增、删、查接口
- 业务统计

## v0.9.255 [2025-12-10]
- 量包查询