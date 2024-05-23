# CAS服务

# 概述

本文档主要介绍CAS（SSL证书服务Certificate Authority Service） GO SDK的使用。在使用本文档前，您需要先了解
[CAS的基本知识](https://cloud.baidu.com/doc/CAS/s/bjwvxiksz)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[CAS域名](https://cloud.baidu.com/doc/CAS/s/Tk9meyjjb)的部分，理解Endpoint相关的概念。

## 获取密钥

要使用百度云CAS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问CAS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建CAS Client

CAS Client是CAS控制面服务的客户端，为开发者与CAS控制面服务进行交互提供了一系列的方法。

### 使用AK/SK新建CAS Client

通过AK/SK方式访问CAS，用户可以参考如下代码新建一个CAS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/cas"   //导入CAS服务模块
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个casClient
	casClient, err := cas.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/CAS/index.html)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为CAS的控制面服务地址。

### 使用STS创建CAS Client

**申请STS token**

CAS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问CAS，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建CAS Client**

申请好STS后，可将STS Token配置到CAS Client中，从而实现通过STS Token创建CAS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建CAS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/cas"    //导入CAS服务模块
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

	// 使用申请的临时STS创建CAS控制面服务的Client对象，Endpoint使用默认值
	casClient, err := cas.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create cas client failed:", err)
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
	casClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置CAS Client时，无论对应CAS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问CAS

CAS支持HTTPS传输协议，您可以通过在创建CAS Client对象时指定的Endpoint中指明HTTPS的方式，在CAS GO SDK中使用HTTPS访问CAS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cas"

ENDPOINT := "https://cas.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
casClient, _ := cas.NewClient(AK, SK, ENDPOINT)
```

## 配置CAS Client

如果用户需要配置CAS Client的一些细节的参数，可以在创建CAS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问CAS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cas"

//创建CAS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cas.baidubce.com"
client, _ := cas.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/cas"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cas.baidubce.com"
client, _ := cas.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/cas"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cas.baidubce.com"
client, _ := cas.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问CAS时，创建的CAS Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建CAS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

SSL证书服务CAS（Certificate Authority Service），是针对于云平台的证书服务，包括证书查询、下单、询价、证书管理等功能。

## 获取已购证书列表

### 接口描述
本接口用于查询用户已购买的证书，使用分页

### 接口限制
无

### 请求示例

```go
req := &GetSslListReq{
PageSize: 10,
PageNo:   1,
}
res, err := casClient.GetSslList(req)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CAS API 文档[获取已购证书列表](https://cloud.baidu.com/doc/CAS/s/Fk9mf2s9p#%E8%8E%B7%E5%8F%96%E5%B7%B2%E8%B4%AD%E8%AF%81%E4%B9%A6%E5%88%97%E8%A1%A8)

## 免费证书购买限制检查

### 接口描述
本接口用于查询用户是否可以继续购买免费证书

### 接口限制
无

### 请求示例
```go
res, err := casClient.CheckFreeSsl()
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CAS API 文档[免费证书购买限制检查](https://cloud.baidu.com/doc/CAS/s/Fk9mf2s9p#%E5%85%8D%E8%B4%B9%E8%AF%81%E4%B9%A6%E8%B4%AD%E4%B9%B0%E9%99%90%E5%88%B6%E6%A3%80%E6%9F%A5)

## 计算证书价格

### 接口描述
本接口用于查询用户选择的证书价格

### 请求示例
```go
req := &QuerySslPriceReq{
    OrderType:      "NEW",
    CertType:       "DV",
    ProductType:    "SINGLE",
    Brand:          "TRUSTASIA",
    DomainNumber:   1,
    WildcardNumber: 0,
    PurchaseLength: 1,
}
res, err := casClient.QuerySslPrice(req)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CAS API 文档[计算证书价格](https://cloud.baidu.com/doc/CAS/s/1k9mfs44r#%E8%AE%A1%E7%AE%97%E8%AF%81%E4%B9%A6%E4%BB%B7%E6%A0%BC)


## SSL新购订单并自动支付

### 接口描述
本接口用于购买SSL证书，并且自动支付，自动支付可以使用 代金券，账号返点，账号余额

### 请求示例
```go
req := &CreateNewOrderReq{
    AutoApply:      true,
    OrderType:      "NEW",
    CertType:       "DV",
    ProductType:    "SINGLE",
    Brand:          "TRUSTASIA",
    DomainNumber:   1,
    WildcardNumber: 0,
    PurchaseLength: 1,
}
res, err := casClient.CreateNewOrder(req)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[SSL新购订单并自动支付](https://cloud.baidu.com/doc/CAS/s/Vk9mfzdz8#ssl%E6%96%B0%E8%B4%AD%E8%AE%A2%E5%8D%95%E5%B9%B6%E8%87%AA%E5%8A%A8%E6%94%AF%E4%BB%98)

## 申请证书

### 接口描述
本接口用于对已购买的证书进行申请


### 请求示例
```go
req := &ApplyCertReq{
    Company:    "com",
    Address:    "addr",
    PostalCode: "111000",
    Region: Region{
    Province: "110000",
    City:     "110100",
    Country:  "中国",
    },
    Algorithm:       "RSA",
    Strength:        "RSA_4096",
    Domain:          "domain.com",
    Password:        "abc",
    VerifyMode:      "DNS",
    MultiDomain:     []string{},
    Department:      "cas team",
    CompanyPhone:    "18500000000",
    OrderGivenName:  "hua",
    OrderFamilyName: "li",
    OrderPosition:   "director",
    OrderEmail:      "test@test.com",
    OrderPhone:      "18500000002",
    TechGivenName:   "an",
    TechFamilyName:  "li",
    TechPosition:    "tech director",
    TechEmail:       "tech@test.com",
    TechPhone:       "18500000001",
}
err := casClient.ApplyCert(req, certId)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[申请证书](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E7%94%B3%E8%AF%B7%E8%AF%81%E4%B9%A6)

## 下载确认函模板

### 接口描述
本接口用于对需要提交确认函的证书，提供下载确认函模板的功能


### 请求示例
```go
body, err := casClient.DownloadLetterTemplate(certId)
defer body.close()
// download certificate to local dir
// Zip decompression password is req.FilePassword
file, err := os.Create("data/confirmation.doc")
written, err := io.Copy(file, body)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[下载确认函模板](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E4%B8%8B%E8%BD%BD%E7%A1%AE%E8%AE%A4%E5%87%BD%E6%A8%A1%E6%9D%BF)

## 上传确认函

### 接口描述
本接口用于对需要提交确认函的证书，提供填写确认函后上传的功能


### 请求示例
```go
body, _ = bce.NewBodyFromFile("data/confirmation.doc")
err := casClient.UploadLetter(body, certId)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[上传确认函](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E4%B8%8A%E4%BC%A0%E7%A1%AE%E8%AE%A4%E5%87%BD)

## 取消申请

### 接口描述
本接口用于取消申请流程中的证书


### 请求示例
```go
certId := ""
err := casClient.CancelCertApplication(certId)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[取消申请](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E5%8F%96%E6%B6%88%E7%94%B3%E8%AF%B7)

## 删除申请

### 接口描述
本接口用于删除失败或者到期的证书，释放免费DV证书配额。


### 请求示例
```go
certId := ""
err := casClient.DeleteCertApplication(certId)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[删除申请](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E5%88%A0%E9%99%A4%E7%94%B3%E8%AF%B7)

## 下载证书

### 接口描述
本接口用于下载已经颁发成功的证书。不支持下载重新颁发的证书



### 请求示例
```go
req := &DownloadCertReq{
    Format:       "PEM",
    FilePassword: "aaa",
}
body, err := casClient.DownloadCert(req, certId)

defer func(body io.ReadCloser) {
err := body.Close()
if err != nil {
    fmt.println(err)
}
}(body)
// download certificate to local dir
// Zip decompression password is req.FilePassword
file, err := os.Create("data/downloaded_cert.zip")
written, err := io.Copy(file, body)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[下载证书](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E4%B8%8B%E8%BD%BD%E8%AF%81%E4%B9%A6)

## 证书详情

### 接口描述
本接口用于查看证书详情。不支持下载重新颁发的证书



### 请求示例
```go
certId := ""
detail, err := casClient.GetCertDetail(certId)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[证书详情](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E8%AF%81%E4%B9%A6%E8%AF%A6%E6%83%85)

## PKI信息

### 接口描述
本接口用于查看证书pki信息详情。不支持下载重新颁发的证书


### 请求示例
```go
certId := ""
pki, err := casClient.GetCertPki(certId)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[PKI信息](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#pki%E4%BF%A1%E6%81%AF)

## 公司及联系人信息

### 接口描述
本接口用于查看证书公司及联系人信息。


### 请求示例
```go
certId := ""
contact, err := casClient.GetCertContact(certId)
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[公司及联系人信息](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E5%85%AC%E5%8F%B8%E5%8F%8A%E8%81%94%E7%B3%BB%E4%BA%BA%E4%BF%A1%E6%81%AF)

## 证书过户

### 接口描述
本接口用于对已购买的证书进行转移账号


### 请求示例
```go
req := &ChangeCertUserReq{
    Params: []string{
        "",
    },
}
err := casClient.ChangeCertUser(req, "newUserId")
```
> **提示：**
> - 详细的参数配置及使用，可以参考CAS API 文档[证书过户](https://cloud.baidu.com/doc/CAS/s/Kk9mg74kf#%E8%AF%81%E4%B9%A6%E8%BF%87%E6%88%B7)