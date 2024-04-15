# 证书管理服务

# 概述

本文档主要介绍CERT GO SDK的使用。在使用本文档前，您需要先了解CERT的一些基本知识。若您还不了解CERT，可以参考[证书管理](https://cloud.baidu.com/doc/Reference/s/8jwvz26si)。

# 初始化

## 确认Endpoint

目前使用CERT服务时，Endpoint统一使用`certificate.baidubce.com`, 支持http和https两种协议。

## 获取密钥

要使用百度云CERT，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问CERT做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建CERT Client

CERT Client是CERT服务的客户端，为开发者与CERT服务进行交互提供了一系列的方法。

### 使用AK/SK新建CERT Client

通过AK/SK方式访问CERT，用户可以参考如下代码新建一个CERT Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/cert"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个CERTClient
	certClient, err := cert.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为VPC的服务地址。

### 使用STS创建CERT Client

**申请STS token**

CERT可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问CERT，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建CERT Client**

申请好STS后，可将STS Token配置到CERT Client中，从而实现通过STS Token创建CERT Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建CERT Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/cert" //导入CERT服务模块
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

	// 使用申请的临时STS创建CERT服务的Client对象，Endpoint使用默认值
	certClient, err := cert.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "certificate.baidubce.com")
	if err != nil {
		fmt.Println("create cert client failed:", err)
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
	certClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置CERT Client时，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问CERT

CERT支持HTTPS传输协议，您可以通过在创建CERT Client对象时指定的Endpoint中指明HTTPS的方式，在CERT GO SDK中使用HTTPS访问CERT服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

ENDPOINT := "https://certificate.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
certClient, _ := cert.NewClient(AK, SK, ENDPOINT)
```

## 配置CERT Client

如果用户需要配置CERT Client的一些细节的参数，可以在创建CERT Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问CERT服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

//创建CERT Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "certificate.baidubce.com"
client, _ := cert.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "certificate.baidubce.com"
client, _ := cert.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "certificate.baidubce.com"
client, _ := cert.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问CERT时，创建的CERT Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建CERT Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# CERT管理

证书管理模块主要用于管理用户的SSL证书，方便用户录入以及查看SSL证书。

## 创建证书

使用以下代码可以创建证书。
```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

args := &cert.CreateCertArgs{
	// 指定证书名称, 必选
    CertName:        "sdkcreateTest",
    // 指定服务器证书的数据内容 (Base64编码), 必选
    CertServerData:  testCertServerData,
    // 指定证书的私钥数据内容 (Base64编码), 必选
    CertPrivateData: testCertPrivateData,
    // 指定证书链数据内容 (Base64编码), 可选
    CertLinkData: certLinkData,
}
result, err := client.CreateCert(args)
if err != nil {
    fmt.Printf("create cert error: %+v\n", err)
    return
}

fmt.Printf("create cert success: %+v\n", result)
```

> 注意:
> - 证书的名称: 长度限制为1-65个字符，以字母开头，只允许包含字母、数字、’-‘、’/’、’.’、’’，Java正则表达式` ^[a-zA-Z]a-zA-Z0-9\-/\.]{2,64}$`

该请求可能存在的异常描述如下。

异常code | 说明 
--- | --- 
CertExceedLimit (409) |	超过用户最大证书数
UnmatchedPairParameterInvalidException (400) | 证书有效时间不包含当前时间
PrivateKeyParameterInvalid (400) | 私钥解析异常
CertificateParameterInvalid (400) |	证书解析异常
CertChainParameterInvalid (400) | 证书链解析异常
UnmatchedPairParameterInvalid (400) | 公钥私钥不匹配

## 修改证书名称

使用以下代码可以修改证书名称。
```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

args := &cert.UpdateCertNameArgs{
    CertName: "test-sdk-cert",
}
err = client.UpdateCertName(certId, args)
if err != nil {
    fmt.Printf("update cert error: %+v\n", err)
    return
}
fmt.Printf("update cert success\n")
```

该请求可能存在的异常描述如下。

异常code	| 说明
--- | ---
AccessDeniedException |	无权限访问
ResourceNotFoundException |	证书不存在

## 查看证书列表

使用以下代码可以查看用户的证书列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

result, err := client.ListCerts()
if err != nil {
    fmt.Printf("list certs error: %+v\n", err)
    return
}

// 查看证书列表的详细信息
for _, c := range listResult.Certs {
    fmt.Println("cert id: ", c.CertId)
    fmt.Println("cert name: ", c.CertName)
    fmt.Println("cert common name: ", c.CertCommonName)
    fmt.Println("cert start time: ", c.CertStartTime)
    fmt.Println("cert stop time: ", c.CertStopTime)
    fmt.Println("cert create time: ", c.CertCreateTime)
    fmt.Println("cert update time: ", c.CertUpdateTime)
    fmt.Println("cert type: ", c.CertType)
}
```

## 获取证书信息(无证书公钥私钥)

使用以下代码可以获取指定的证书信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

result, err := client.GetCertMeta(certId)
if err != nil {
    fmt.Printf("get certs meta error: %+v\n", err)
    return
}

// 获取得到证书id
fmt.Println("cert id: ", result.CertId)
// 获取得到证书名称
fmt.Println("cert name: ", result.CertName)
// 获取得到证书通用名称
fmt.Println("cert common name: ", result.CertCommonName)
// 获取得到证书指纹
fmt.Println("cert fingerprint: ", result.CertFingerprint)
// 获取得到证书生效时间
fmt.Println("cert start time: ", result.CertStartTime)
// 获取得到证书到期时间
fmt.Println("cert stop time: ", result.CertStopTime)
// 获取得到证书创建时间
fmt.Println("cert create time: ", result.CertCreateTime)
// 获取得到证书更新时间
fmt.Println("cert update time: ", result.CertUpdateTime)
// 获取得到证书类型
fmt.Println("cert type: ", result.CertType)
```

该请求可能存在的异常描述如下。

异常code | 说明
--- | ---
AccessDeniedException |	无权限访问
ResourceNotFoundException |	证书不存在

## 删除证书

使用以下代码可以删除指定的证书。
```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

if err := client.DeleteCert(certId); err != nil {
    fmt.Printf("delete certs error: %+v\n", err)
    return
}
fmt.Printf("delete certs success\n")
```

该请求可能存在的异常描述如下。

异常code | 说明
--- | ---
OperationNotAllowedException |	证书使用中
AccessDeniedException |	无权限访问
ResourceNotFoundException |	证书不存在

## 替换证书

使用以下代码可以替换过期且不再使用中的证书。
```go
// import "github.com/baidubce/bce-sdk-go/services/cert"

args := &cert.UpdateCertDataArgs{
	// 指定要替换的证书名称
    CertName:        "test-sdk-cert",
    // 指定服务器证书的数据内容 (Base64编码)
    CertServerData:  testUpdateCertServerData,
    // 指定证书的私钥数据内容 (Base64编码)
    CertPrivateData: testUpdateCertPrivateData,
    // 指定证书链数据内容 (Base64编码)
    CertLinkData: certLinkData,
}
if err := client.UpdateCertData(createResult.CertId, args); err != nil {
    fmt.Printf("update cert data error: %+v\n", err)
    return
}
fmt.Printf("update cert data success\n")
```

> 注意: 使用该api替换证书后，证书的id保持不变。

该请求可能存在的异常描述如下。

异常code | 说明
--- | ---
OperationNotAllowedException(409) |	证书使用中或者证书
AccessDeniedException（403） |	证书非本用户或子用户无该证书运维权限
ResourceNotFoundException（404）| 无证书
CertExceedLimit (409) | 超过用户最大证书数
UnmatchedPairParameterInvalidException (400) |	证书有效时间不包含当前时间
PrivateKeyParameterInvalid (400) |	私钥解析异常
CertificateParameterInvalid (400) |	证书解析异常
CertChainParameterInvalid (400) | 证书链解析异常
UnmatchedPairParameterInvalid (400) | 公钥私钥不匹配


# 错误处理

GO语言以error类型标识错误，CERT支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | CERT服务返回的错误

用户使用SDK调用CERT相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// certClient 为已创建的CERT Client对象
result, err := client.ListCerts()
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

客户端异常表示客户端尝试向CERT发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当CERT服务端出现异常时，CERT服务端会返回给用户相应的错误信息，以便定位问题。

## SDK日志

CERT GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

CERT GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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
log.Debugf("%s", "logging message using the log package in the CERT go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the CERT go sdk")
```


# 版本变更记录

## v0.9.5 [2019-09-24]

首次发布:

 - 支持创建证书、修改证书名称、查看证书列表、获取证书信息(无证书公钥私钥)、删除证书、替换证书接口。

## v0.9.175 [2024-04-10]
- 查询证书接口增加证书指纹信息CertFingerprint