# DTS服务

# 概述

本文档主要介绍DTS GO SDK的使用。在使用本文档前，您需要先了解DTS的一些基本知识，并已开通了DTS服务。若您还不了解DTS，可以参考[产品描述](https://cloud.baidu.com/doc/DTS/s/ujwvyzdzg)和[操作指南](https://cloud.baidu.com/doc/DTS/s/Qjwvz0ikk)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

目前支持“华北-北京”区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
BJ | rds.bj.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云DTS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问DTS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建DTS Client

DTS Client是DTS服务的客户端，为开发者与DTS服务进行交互提供了一系列的方法。

### 使用AK/SK新建DTS Client

通过AK/SK方式访问DTS，用户可以参考如下代码新建一个DTS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/dts"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个DTSClient
	dtsClient, err := dts.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为VPC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`dts.bj.baidubce.com`。

### 使用STS创建DTS Client

**申请STS token**

DTS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问DTS，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建DTS Client**

申请好STS后，可将STS Token配置到DTS Client中，从而实现通过STS Token创建DTS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建DTS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/dts" //导入DTS服务模块
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

	// 使用申请的临时STS创建DTS服务的Client对象，Endpoint使用默认值
	dtsClient, err := dts.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "dts.bj.baidubce.com")
	if err != nil {
		fmt.Println("create dts client failed:", err)
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
	dtsClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置DTS Client时，无论对应DTS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问DTS

DTS支持HTTPS传输协议，您可以通过在创建DTS Client对象时指定的Endpoint中指明HTTPS的方式，在DTS GO SDK中使用HTTPS访问DTS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

ENDPOINT := "https://dts.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
dtsClient, _ := dts.NewClient(AK, SK, ENDPOINT)
```

## 配置DTS Client

如果用户需要配置DTS Client的一些细节的参数，可以在创建DTS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问DTS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

//创建DTS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "dts.bj.baidubce.com"
client, _ := dts.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "dts.bj.baidubce.com"
client, _ := dts.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "dts.bj.baidubce.com"
client, _ := dts.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问DTS时，创建的DTS Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建DTS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# DTS管理

DTS（Data Transmission Service）提供数据迁移、数据同步、数据订阅于一体的数据库数据传输服务，帮助您在业务不停服的前提下轻松完成数据库迁移，利用实时同步通道轻松构建异地容灾的高可用数据库架构。

## 创建DTS任务

使用以下代码可以创建一个DTS任务
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

args := &dts.CreateDtsArgs{
    ProductType:        "postpay",
    Type:               "migration",
    Standard:           "Large",
    SourceInstanceType: "bcerds",
    TargetInstanceType: "bcerds",
    CrossRegionTag:     1,
}
result, err := client.CreateDts(args)
if err != nil {
    fmt.Printf("create dts error: %+v\n", err)
    return
}

for _, e := range result.DtsTasks {
	fmt.Println("create dts success, task id: ", e.DtsId)
}
```

## 删除DTS任务

使用以下代码可以删除一个DTS任务
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

result, err := client.DeleteDts(taskId)
if err != nil {
    fmt.Printf("delete dts error: %+v\n", err)
    return
}

fmt.Println("delete dts success\n")
```

## 查看DTS任务详情

使用以下代码可以查看一个DTS任务详情
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

result, err := client.GetDetail(taskId)
if err != nil {
    fmt.Printf("get dts detail error: %+v\n", err)
    return
}

fmt.Println("dts taskName: ",result.TaskName)
fmt.Println("dts status: ",result.Status)
fmt.Println("dts region: ",result.Region)
fmt.Println("dts createTime: ",result.CreateTime)
```

## 查看DTS任务列表

使用以下代码可以查看DTS任务列表
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

args := &dts.ListDtsArgs{
    Type:               "migration",
}
result, err := client.ListDts(args)
if err != nil {
    fmt.Printf("get dts list error: %+v\n", err)
    return
}

for _, e := range result.Task {
	fmt.Println("dtsId: ", e.DtsId)
    fmt.Println("taskName: ", e.TaskName)
	fmt.Println("status: ", e.Status)
}
```

## 配置DTS任务

使用以下代码可以配置一个DTS任务
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

dataType :=[]string{"schema","base"}
	srcConnection := Connection{
		Region:       "public",
		DbType:       "mysql",
		DbUser:       "usrname",
		DbPass:       "password",
		DbPort:       3306,
		DbHost:       "180.76.120.207",
		InstanceId:   "rds-TOzVOznv",
		InstanceType: "public",
	}
	dstConnection := Connection{
		Region:       "public",
		DbType:       "mysql",
		DbUser:       "usrname",
		DbPass:       "password",
		DbPort:       3306,
		DbHost:       "180.76.120.207",
		InstanceId:   "rds-TOzVOznv",
		InstanceType: "public",
	}
	schema := Schema{
		Type:  "db",
		Src:   "db1",
		Dst:   "db2",
		Where: "",
	}
	schemaMapping := []Schema{schema}
	args := &ConfigArgs{
		TaskName:      "test",
		DataType:      dataType,
		Type:          "migration",
		SrcConnection: srcConnection,
		DstConnection: dstConnection,
		SchemaMapping: schemaMapping,
	}
result, err := client.CreateDts(args)
if err != nil {
    fmt.Printf("config dts error: %+v\n", err)
    return
}

	fmt.Println("config dts success, task id: ", result.DtsId)
```

## 前置检查

使用以下代码可以对一个DTS任务前置检查
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

result, err := client.PreCheck(taskId)
if err != nil {
    fmt.Printf("preCheck dts error: %+v\n", err)
    return
}

fmt.Println("preCheck dts: ",result.Success)
```

## 查看前置检查结果

使用以下代码可以查看前置检查结果
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

result, err := client.GetPreCheck(taskId)
if err != nil {
    fmt.Printf("get dts preCheck result error: %+v\n", err)
    return
}

fmt.Println("get dts preCheck result: ",result.GetPreCheckResult)
for _, e := range result.Result {
    fmt.Println("check result name: ",e.Name)
    fmt.Println("check result status: ",e.Status)
    fmt.Println("check result message: ",e.Message)
    fmt.Println("check result subscription: ",e.Subscription)
}
```

## 启动DTS任务

使用以下代码可以启动一个DTS任务
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

result, err := client.StartDts(taskId)
if err != nil {
    fmt.Printf("start dts error: %+v\n", err)
    return
}

fmt.Println("start dts success\n")
```

## 暂停DTS任务

使用以下代码可以暂停一个DTS任务
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

result, err := client.PauseDts(taskId)
if err != nil {
    fmt.Printf("pause dts error: %+v\n", err)
    return
}

fmt.Println("pause dts success\n")
```

## 结束DTS任务

使用以下代码可以结束一个DTS任务
```go
// import "github.com/baidubce/bce-sdk-go/services/dts"

result, err := client.ShutdownDts(taskId)
if err != nil {
    fmt.Printf("shut down dts error: %+v\n", err)
    return
}

fmt.Println("shut down dts success\n")
```

# 错误处理

GO语言以error类型标识错误，DTS支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | DTS服务返回的错误

用户使用SDK调用DTS相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// dtsClient 为已创建的DTS Client对象
result, err := client.ListDts()
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

客户端异常表示客户端尝试向DTS发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当DTS服务端出现异常时，DTS服务端会返回给用户相应的错误信息，以便定位问题。

## SDK日志

DTS GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

DTS GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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
log.Debugf("%s", "logging message using the log package in the DTS go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the DTS go sdk")
```



首次发布:

 - 支持创建DTS任务、删除DTS任务、配置DTS任务、查看DTS任务详情、查看DTS任务列表、前置检查、查看前置检查结果、启动DTS任务、暂停DTS任务、结束DTS任务