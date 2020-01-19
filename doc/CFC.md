# CFC GO SDK文档

# 初始化

## 确认Endpoint

目前支持“华北-北京”、“华南-广州” 两个区域。北京区域：`http://cfc.bj.baidubce.com`，广州区域：`http://cfc.gz.baidubce.com` 对应信息为：

访问区域 | 对应Endpoint
---|---
BJ | cfc.bj.baidubce.com
GZ | cfc.gz.baidubce.com

## 获取密钥

要使用百度云BOS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BOS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## CFC Client

CFC Client是CFC服务的客户端，为开发者与CFC服务进行交互提供了一系列的方法。

### 使用AK/SK新建CFC Client

通过AK/SK方式访问CFC，用户可以参考如下代码新建一个CFC Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/cfc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个cfcClient
	cfcClient, err := cfc.NewClient(AK, SK, ENDPOINT)
}
```

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`http://cfc.bj.baidubce.com`。

### 使用STS创建CFC Client

**申请STS token**

CFC可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问CFC，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建CFC Client**

申请好STS后，可将STS Token配置到CFC Client中，从而实现通过STS Token创建CFC Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建CFC Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/cfc" //导入CFC服务模块
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

	// 使用申请的临时STS创建CFC服务的Client对象，Endpoint使用默认值
	cfcClient, err := cfc.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create cfc client failed:", err)
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
	cfcClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置CFC Client时，无论对应CFC服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问CFC

CFC支持HTTPS传输协议，您可以通过在创建CFC Client对象时指定的Endpoint中指明HTTPS的方式，在CFC GO SDK中使用HTTPS访问CFC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cfc"

ENDPOINT := "https://cfc.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
cfcClient, _ := cfc.NewClient(AK, SK, ENDPOINT)
```

## 配置cfc Client

如果用户需要配置CFC Client的一些细节的参数，可以在创建CFC Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问CFC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cfc"

//创建CFC Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cfc.bj.baidubce.com"
client, _ := cfc.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/cfc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cfc.bj.baidubce.com"
client, _ := cfc.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/cfc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cfc.bj.baidubce.com"
client, _ := cfc.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问CFC时，创建的CFC Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建CFC Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

## 函数调用

使用以下代码可以调用执行一个指定的CFC函数
```go
args := &api.InvocationsArgs{
	FunctionName:   "sdk-create",
	InvocationType: api.InvocationTypeRequestResponse,
	Payload:        nil,
}

// 若想执行特定版本的函数，可以设置
args.Qualifier = "1"

result, err := client.Invocations(args)
if err != nil {
    fmt.Println("invocation function failed:", err)
} else {
    fmt.Println("invocation function success: ", result)
}
```

## 函数操作

### 创建函数

使用以下代码可以创建一个CFC函数
```go
arge := &api.CreateFunctionArgs{
    // 配置函数的代码，需要上传代码的zip压缩包
	Code:         &api.CodeFile{ZipFile: zipFile},
    // 函数名称，每个用户的函数名称不可重复，不可修改
	FunctionName: "sdk-create",
    // 函数调用的入口函数
	Handler:      "index.handler",
    // 函数的runtime
	Runtime:      "nodejs8.5",
    // 函数运行的内存大小，单位mb，必须是128的整数倍，最大可选1024
	MemorySize:   256,
    // 函数执行超时时间，可选1-300s
	Timeout:      3,
    // 函数描述信息
	Description:  "sdk create",
    // 函数日志存放方式，可选bos，表示函数执行日志存放在bos中
	LogType:      "bos",
    // 若LogType配置为bos，此参数设置函数执行日志在bos中的存储地址
	LogBosDir:    "bos://ashjfdklsfhlk/",
})

// 若要配置从bos bucket中上传函数代码，可以如下设置
// 这两个参数不能和args.Code.ZipFile同时设置
args.Code.BosBucket = "bucketName"
args.Code.BosObject = "objectKey"

// 若要直接发布函数，可以设置
args.Code.Publish = true

// 若要配置函数访问VPC网络，可以如下设置
args.VpcConfig = &api.VpcConfig{
    SubnetIds:        []string{"subnet_id1"},
    SecurityGroupIds: []string{"security_group_id1"},
}

// 若要配置环境变量，可以如下设置
args.Environment = &api.Environment{
    Variables: map[string]string{
        "key": "value",
    },
},

result, err := client.CreateFunction(args)
if err != nil {
    fmt.Println("create function failed:", err)
} else {
    fmt.Println("create function success: ", result)
}
```

### 函数列表

使用以下代码可以获取CFC函数的列表
```go
args := &api.ListFunctionArgs{}

// 若想查询指定版本1的函数，可以如下设置
args.FunctionVersion = "1"

result, err := client.ListFunctions(args)
if err != nil {
    fmt.Println("list function failed:", err)
} else {
    fmt.Println("list function success: ", result)
}
```

### 函数信息

使用以下代码可以获取特定函数的信息
```go
args := &api.GetFunctionArgs{
    FunctionName: "functionName"
}
result, err := client.GetFunction(args)
if err != nil {
    fmt.Println("get function failed:", err)
} else {
    fmt.Println("get function success: ", result)
}
```

### 删除函数

使用以下代码可以删除一个特定的CFC函数
```go
args := &api.DeleteFunctionArgs{
    FunctionName: "sdk-create",
}

// 若想删除函数的某个版本，可以设置
args.Qualifier = "1"

err := client.DeleteFunction(args)
if err != nil {
    fmt.Println("delete function failed:", err)
} 
```

### 更新函数代码

使用以下代码可以更新特定CFC函数的代码
```go
args := &api.UpdateFunctionCodeArgs{
	FunctionName: "sdk-creat"
	ZipFile:      []byte(functionZipCode)
}

// 若要配置从bos bucket中上传函数代码，可以如下设置
// 这两个参数不能和args.ZipFile同时设置
args.BosBucket = "bucketName"
args.BosObject = "objectKey"

// 若要直接发布函数，可以设置
args.Publish = true

result, err := client.UpdateFunctionCode(args)
if err != nil {
    fmt.Println("update function code failed:", err)
} else {
    fmt.Println("update function code success: ", result)
}
```

### 获取函数配置

使用以下代码可以获取特定CFC函数的配置
```go
args := &api.GetFunctionConfigurationArgs{
    FunctionName: "sdk-create",
}

// 若想查询特定版本的函数的配置，可以设置
args.Qualifier = functionBrn

if err != nil {
    fmt.Println("get function configure failed:", err)
} else {
    fmt.Println("get function configure success: ", result)
}
```

### 更新函数配置

使用以下代码可以更新特定CFC函数的配置
```go
args := &api.UpdateFunctionConfigurationArgs{
	FunctionName: "sdk-create",
	Timeout:      20,
	Description:  "sdk update",
	Runtime:      "nodejs8.5",
	MemorySize:   &memorySize,
	Environment:  &api.Environment{
		Variables: map[string]string{
			"name": "Test",
		},
	},
})

result, err := client.UpdateFunctionConfiguration(args)
if err != nil {
    fmt.Println("update function configure failed:", err)
} else {
    fmt.Println("update function configure success: ", result)
}
```

### 设置函数预留并发度

使用以下代码可以设置和更新特定CFC函数的预留并发度
```go
args := &api.ReservedConcurrentExecutionsArgs{
    FunctionName: "sdk-create",
    // 预留并发度会由本函数的所有版本共享，最高能设置90
	ReservedConcurrentExecutions: 10,
})

err := client.SetReservedConcurrentExecutions(args)
if err != nil {
    fmt.Println("set function reserved concurrent executions failed:", err)
}
```

### 删除函数预留并发度设置

使用以下代码可以删除特定CFC函数的预留并发度设置
```go
args := &api.DeleteReservedConcurrentExecutionsArgs{
    FunctionName: "sdk-create",
})

err := client.DeleteReservedConcurrentExecutions(args)
if err != nil {
    fmt.Println("delete function reserved concurrent executions failed:", err)
}
```

## 版本操作

### 获取函数版本列表

使用以下代码可以获取函数版本列表
```go
args := &api.ListVersionsByFunctionArgs{
    FunctionName: "sdk-create",
}

result, err := client.ListVersionsByFunction(args)
if err != nil {
    fmt.Println("get function version failed:", err)
} else {
    fmt.Println("get function version success: ", result)
}
```

### 发布版本

使用以下代码可以为函数发布一个版本
```go
args := &api.PublishVersionArgs{
    FunctionName: "sdk-create",
}

// 若想添加版本描述，可以设置
args.Descirption = "publish description"

// 若想对版本的部署包进行sha256验证，可以设置
args.CodeSha256 = "codeSha256"

result, err := client.PublishVersion(args)
if err != nil {
    fmt.Println("publish function version failed:", err)
} else {
    fmt.Println("publish function version success: ", result)
}
```

## 别名操作

### 获取别名列表

使用以下代码可以获取函数的别名列表
```go
args := &api.ListAliasesArgs{
    FunctionName: "sdk-create",
}

// 若想获取特定函数版本的别名，可以设置
args.FunctionVersion = "1"

result, err := client.ListAliases(args)
if err != nil {
    fmt.Println("list function alias failed:", err)
} else {
    fmt.Println("list function alias success: ", result)
}
```


### 创建别名

使用以下代码可以为特定函数版本创建一个别名
```go
args := &api.CreateAliasArgs{
    FunctionName: "sdk-create",
    Name:         "alias-create",
}

// 若要将别名绑定到特定函数版本，可以设置
args.FunctionVersion = "1"

// 若要设置别名标书，可以设置
args.Description = "alias description"


result, err := client.CreateAlias(args)
if err != nil {
    fmt.Println("create function alias failed:", err)
} else {
    fmt.Println("create function alias success: ", result)
}
```

### 获取别名信息

使用以下代码可以获取一个特定函数的别名的信息
```go
args := &api.GetAliasArgs{
    FunctionName: "sdk-create",
    AliasName:    "alias-create",
}

result, err := client.GetAlias(args)
if err != nil {
    fmt.Println("get function alias failed:", err)
} else {
    fmt.Println("get function alias success: ", result)
}
```

### 更新别名

使用以下代码可以更新一个函数的别名
```go
args := &api.UpdateAliasArgs{
    FunctionName:    "sdk-create",
    AliasName:       "alias-create",
    Description:     "test alias",
}

// 若要修改别名绑定的函数版本，可以设置
args.FunctionVersion = "$LATEST"

result, err := client.UpdateAlias(args)
if err != nil {
    fmt.Println("update function alias failed:", err)
} else {
    fmt.Println("update function alias success: ", result)
}
```

### 删除别名

使用以下代码可以删除一个函数的别名
```go
args := &api.DeleteAliasArgs{
    FunctionName: "sdk-create",
    AliasName:    "alias-create",
}

err := client.DeleteAlias(args)
if err != nil {
    fmt.Println("delete function alias failed:", err)
}
```

## 触发器操作

### 获取触发器列表

使用以下代码可以获取一个触发器的列表
```go
args := &api.ListTriggersArgs{
    FunctionBrn: "functionBrn",
}

// 默认不返回cfc-edge触发器，若想查询所有的触发器，可以设置
args.ScopeType = "all"

result, err := client.ListTriggers(args)
if err != nil {
    fmt.Println("get function trigger failed:", err)
} else {
    fmt.Println("get function trigger success: ", result)
}
```

### 创建触发器

使用以下代码可以创建一个特定的触发器并绑定
```go
args := &api.CreateTriggerArgs{
    Target: "functionBrn",
	Source: api.SourceTypeCrontab,
    // 创建crontab触发器所需的数据
	Data: &api.CrontabTriggerData{
		Name:               "sdkName",
		Brn:                "functionBrn",
		ScheduleExpression: "cron(0 10 * * ?)",
		Enabled:            "Disabled",
	},
}

result, err := client.CreateTrigger(args)
if err != nil {
    fmt.Println("create function trigger failed:", err)
} else {
    fmt.Println("create function trigger success: ", result)
}
```

> **提示：**
> 1.  不同类型的触发器，其Data字段所需内容不同，具体可以参考文档[触发器配置](https://cloud.baidu.com/doc/CFC/s/Kjwvz47o9#relationconfiguration)

### 更新触发器

使用以下代码可以更新一个函数的触发器
```go
args := &api.UpdateTriggerArgs{
    RelationId: RelationId,
	Target:     functionBRN,
	Source:     api.SourceTypeHTTP,
	Data: &api.HttpTriggerData{
		ResourcePath: fmt.Sprintf("tr99-%s", time.Now().Format("2006-01-02T150405")),
		Method:       "GET",
		AuthType:     "anonymous",
	},
}

result, err := client.UpdateTrigger(args)
if err != nil {
    fmt.Println("update function trigger failed:", err)
} else {
    fmt.Println("update function trigger success: ", result)
}
```

### 删除触发器

使用以下代码可以删除一个触发器
```go
args := &api.DeleteTriggerArgs{
    RelationId: RelationId,
	Target:     functionBRN,
	Source:     api.SourceTypeHTTP,
}

err := client.DeleteTrigger(args)
if err != nil {
    fmt.Println("delete function trigger failed:", err)
}
```


# 错误处理

GO语言以error类型标识错误，CFC支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
CFCClientError  | 用户操作产生的错误
BceServiceError | CFC服务返回的错误

用户使用SDK调用CFC相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// cfcClient 为已创建的CFC Client对象
args := &api.GetFunctionArgs{
    FunctionName: "functionName"
}
result, err := cfcClient.GetFunction(args)
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
	fmt.Println("get function detail success: ", result)
}
```

## 客户端异常

客户端异常表示客户端尝试向CFC发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当CFC服务端出现异常时，CFC服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[CFC错误返回](https://cloud.baidu.com/doc/CFC/s/Djwvz4cwc)

# 版本变更记录

## v0.9.1 [2019-09-26]

首次发布：

 - 执行函数
 - 创建、查看、列表、删除函数，更新函数代码，更新、获取函数配置
 - 设置、删除函数预留并发度
 - 列表、创建函数版本
 - 列表、创建、获取、更新、删除别名
 - 获取、创建、更新、删除触发器
