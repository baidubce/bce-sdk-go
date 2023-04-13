# BCI服务

# 概述

本文档主要介绍BCI GO SDK的使用。在使用本文档前，您需要先了解BCI的一些基本知识，并已开通了BCI服务。若您还不了解BCI，可以参考[产品描述](https://cloud.baidu.com/doc/BCI/s/Ujxessavb)和[操作指南](https://cloud.baidu.com/doc/BCI/s/elc8ri0af)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BCI服务域名](https://cloud.baidu.com/doc/BCI/s/Qlf0qoqw1)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

目前支持“华北-北京”、“华南-广州”、“华东-苏州”和“华北-保定”四个区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
BD | bci.bd.baidubce.com | HTTP and HTTPS
GZ | bci.gz.baidubce.com | HTTP and HTTPS
SU | bci.su.baidubce.com | HTTP and HTTPS
BD | bci.bd.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云BCI，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BCI做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建BCI Client

BCI Client是BCI服务的客户端，为开发者与BCI服务进行交互提供了一系列的方法。

### 使用AK/SK新建BCI Client

通过AK/SK方式访问BCI，用户可以参考如下代码新建一个BCI Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bci"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BCIClient
	bciClient, err := bci.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BCI的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`bci.bj.baidubce.com`。

### 使用STS创建BCI Client

**申请STS token**

BCI可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问BCI，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建BCI Client**

申请好STS后，可将STS Token配置到BCI Client中，从而实现通过STS Token创建BCI Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建BCI Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/bci" //导入BCI服务模块
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

	// 使用申请的临时STS创建BCI服务的Client对象，Endpoint使用默认值
	bciClient, err := bci.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "bci.bj.baidubce.com")
	if err != nil {
		fmt.Println("create bci client failed:", err)
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
	bciClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置BCI Client时，无论对应BCI服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问BCI

BCI支持HTTPS传输协议，您可以通过在创建BCI Client对象时指定的Endpoint中指明HTTPS的方式，在BCI GO SDK中使用HTTPS访问BCI服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

ENDPOINT := "https://bci.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
bciClient, _ := bci.NewClient(AK, SK, ENDPOINT)
```

## 配置BCI Client

如果用户需要配置BCI Client的一些细节的参数，可以在创建BCI Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问BCI服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

//创建BCI Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bci.bj.baidubce.com"
client, _ := bci.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bci.bj.baidubce.com"
client, _ := bci.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bci.bj.baidubce.com"
client, _ := bci.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问BCI时，创建的BCI Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建BCI Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

百度智能云容器实例（Baidu Container Instance，即BCI）提供无服务器化的容器资源。您只需提供容器镜像及启动容器所需的配置参数，即可运行容器，而无需关心这些容器如何被调度部署到底层的物理服务器资源中。BCI服务将为您完成IaaS层资源的调度和运维工作，从而简化您对容器的使用流程，降低部署和维护成本。同时BCI只会对您创建容器时申请的资源计费，因此实现真正的按需付费。结合容器本身秒级启动的特性，BCI可以帮助您在百度智能云上构建灵活弹性且易于维护的大规模容器集群。

> 注意:
> - BCI使用限制可以参考[BCI使用限制](https://cloud.baidu.com/doc/BCI/s/jjxv2cma2)。
> - 创建BCI需要实名认证，若未通过实名认证可以前往[百度智能云官网控制台](https://console.bce.baidu.com/qualify/#/qualify/result)中的安全认证下的实名认证中进行认证。

## 创建实例

使用以下代码可以创建一个BCI实例。
```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

args := &CreateInstanceArgs{
    // 保证请求幂等性
    ClientToken: "random-uuid",
    // BCI实例名称
    Name: "instanceName",
    // 可用区名称
    ZoneName: "zoneC",
    // 实例所属于的安全组Id
    SecurityGroupIds: []string{"g-59gf44p4jmwe"},
    // 实例所属的子网Id
    SubnetIds: []string{"sbn-g463qx0aqu7q"},
    // 实例重启策略
    RestartPolicy: "Always",
    // 弹性公网IP
    EipIp: "106.13.234.xx",
    // 自动创建EIP
    AutoCreateEip: false,
    // 弹性公网名称
    EipName: "zwj-test-eip",
    // EIP线路类型
    EipRouteType: "BGP",
    // 公网带宽，单位为Mbps
    EipBandwidthInMbps: 1,
    // 付款时间
    EipPaymentTiming: "Postpaid",
    // 计费方式
    EipBillingMethod: "ByTraffic",
    // 支持创建 EIP同时开通自动续费单位
    EipAutoRenewTimeUnit: "month",
    // 支持创建 EIP同时开通自动续费时间
    EipAutoRenewTime: 1,
    // 实例所需的 CPU 资源型号
    CPUType: "intel",
    // 实例所需的 GPU 资源型号
    GPUType: "Nvidia A10 PCIE",
    // 程序的缓冲时间，用于处理关闭之前的操作
    TerminationGracePeriodSeconds: 0,
    // 主机名称
    HostName: "zwj-go-sdktest",
    // 用户标签列表
    Tags: []Tag{
        {
            TagKey:   "appName",
            TagValue: "zwj-test",
        },
    },
    // 镜像仓库凭证信息
    ImageRegistryCredentials: []ImageRegistryCredential{
        {
            Server:   "docker.io/wywcoder",
            UserName: "wywcoder",
            Password: "Qaz123456",
        },
    },
    // 业务容器组
    Containers: []Container{
        {
            Name:            "container01",
            Image:           "registry.baidubce.com/bci-zjm-public/ubuntu:18.04",
            Memory:          0.5,
            CPU:             0.25,
            GPU:             0,
            WorkingDir:      "",
            ImagePullPolicy: "IfNotPresent",
            Commands:        []string{"/bin/sh"},
            Args:            []string{"-c", "sleep 30000 && exit 0"},
            VolumeMounts: []VolumeMount{
                {
                    MountPath: "/usr/local/nfs",
                    ReadOnly:  false,
                    Name:      "nfs",
                    Type:      "NFS",
                },
                {
                    MountPath: "/usr/local/share",
                    ReadOnly:  false,
                    Name:      "emptydir",
                    Type:      "EmptyDir",
                },
                {
                    MountPath: "/config",
                    ReadOnly:  false,
                    Name:      "configfile",
                    Type:      "ConfigFile",
                },
            },
            Ports: []Port{
                {
                    Port:     8099,
                    Protocol: "TCP",
                },
            },
            EnvironmentVars: []Environment{
                {
                    Key:   "java",
                    Value: "/usr/local/jre",
                },
            },
            LivenessProbe: &Probe{
                InitialDelaySeconds:           0,
                TimeoutSeconds:                0,
                PeriodSeconds:                 0,
                SuccessThreshold:              0,
                FailureThreshold:              0,
                TerminationGracePeriodSeconds: 0,
                Exec: &ExecAction{
                    Command: []string{"echo 0"},
                },
            },
            Stdin:           false,
            StdinOnce:       false,
            Tty:             false,
            SecurityContext: &ContainerSecurityContext{},
        },
    },
    // Init 容器
    InitContainers: []Container{},
    // 数据卷信息
    Volume: &Volume{
        Nfs: []NfsVolume{
            {
                Name:   "nfs",
                Server: "xxx.cfs.gz.baidubce.com",
                Path:   "/",
            },
        },
        EmptyDir: []EmptyDirVolume{
            {
                Name: "emptydir",
            },
        },
        ConfigFile: []ConfigFileVolume{
            {
                Name: "configfile",
                ConfigFiles: []ConfigFileDetail{
                    {
                        Path: "podconfig",
                        File: "filenxx",
                    },
                },
            },
        },
    },
}
result, err := client.CreateInstance(args)
if err != nil {
    fmt.Printf("CreateInstance error: %+v \n", err)
    return
}
fmt.Printf("CreateInstance success, bci instance id: %+v \n", result.InstanceId)
```

> 注意: 
> - 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。只支持ASCII字符，且不能超过64个字符。
> - BCI实例名称，即容器组名称；支持长度为2~253个英文小写字母、数字或者连字符（-），不能以连接字符开始或结尾。如果填写大写字母，后台会自动转为小写。
> - 实例重启策略。取值范围：Always：总是重启，Never：从不重启，OnFailure：失败时重启。默认值：Always。
> - 自动创建EIP，并绑定到BCI实例上。只有当eipIp为空的情况下，此字段才生效。默认值为：false。
> - EIP线路类型，包含标准BGP（BGP）和增强BGP（BGP_S），默认标准BGP。当autoCreateEip为true时，此字段才生效。默认值为：BGP。
> - 公网带宽，单位为Mbps。对于预付费以及按使用带宽计费的后付费EIP，标准型BGP限制为1~500之间的整数，增强型BGP限制为100~5000之间的整数（代表带宽上限）；对于按使用流量计费的后付费EIP，标准型BGP限制为1~200之间的整数（代表允许的带宽流量峰值）。如果填写浮点数会向下取整。当autoCreateEip为true时，此字段才生效。默认值为100。
> - 付款时间，预支付（Prepaid）和后支付（Postpaid）当autoCreateEip为true时，此字段才生效。默认值为Postpaid。
> - 计费方式，按流量（ByTraffic）、按带宽（ByBandwidth）、按增强95（ByPeak95）（只有共享带宽后付费支持）。当autoCreateEip为true时，此字段才生效。增强型BGP_S不支持按流量计费（ByTraffic），需要按带宽计费（ByBandwidth）。默认值为ByTraffic。
> - 支持创建 EIP同时开通自动续费单位，取值为 month 获 year （默认 month）。当autoCreateEip为true时，此字段才生效。默认值为month。
> - 支持创建 EIP同时开通自动续费时间。根据autoRenewTimeUnit的取值有不同的范围，month 为1到9，year 为1到3。当autoCreateEip为true时，此字段才生效。默认值为1。
> - 实例所需的 CPU 资源型号，如果不填写则默认不强制指定 CPU 类型。
> - 实例所需的 GPU 资源型号。目前仅支持：Nvidia A10 PCIE。

## 查询实例列表

使用以下代码可以查询BCI实例列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

args := &ListInstanceArgs{
    // 查询关键字名称
    KeywordType:    "podId",
    // 查询关键字值
    keyword:        "p-xxx",
    // 表示下一个查询开始的marker，marker为空表示没有下一个
    Marker:         "",
    // 每页包含的最大数量
    MaxKeys:        5,
}
result, err := client.ListInstances(args)
fmt.Printf("ListInstances result: %+v, err: %+v \n", result, err)
```

> 注意: 
> - 查询关键字名称，取值范围：name、podId。
> - 表示下一个查询开始的marker，marker为空表示没有下一个。说明：首次查询时无需设置该参数，后续查询的marker从返回结果中获取。
> - 每页包含的最大数量，最大数量通常不超过1000，缺省值为10。maxKeys的取值范围：[1, 1000]之间的正整数。

## 查询实例详情

使用以下代码可以查询BCI实例详情。
```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

args := &GetInstanceArgs{
    // BCI实例ID
    InstanceId: "p-xxx",
}
result, err := client.GetInstance(args)
fmt.Printf("ListInstances result: %+v, err: %+v \n", result, err)
```

## 删除实例

使用以下代码可以删除BCI实例。
```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

args := &DeleteInstanceArgs{
    // 待删除的BCI实例ID
    InstanceId:         "p-xxxx",
    // 释放关联资源
    RelatedReleaseFlag: true,
}
err := client.DeleteInstance(args)
fmt.Printf("DeleteInstance err: %+v\n", err)
```

> 注意: 
> - 释放关联资源，目前只有EIP资源，默认值为false。

## 批量删除实例

使用以下代码可以批量删除BCI实例。
```go
// import "github.com/baidubce/bce-sdk-go/services/bci"

args := &BatchDeleteInstanceArgs{
    // 待删除的BCI实例ID列表
    InstanceIds:        []string{"p-axxx", "p-bxxx"},
    // 释放关联资源
    RelatedReleaseFlag: true,
}
err := client.BatchDeleteInstance(args)
fmt.Printf("BatchDeleteInstance err: %+v\n", err)
```

> 注意: 
> - 释放关联资源，目前只有EIP资源，默认值为false。

# 错误处理

GO语言以error类型标识错误，BCI支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | BCI服务返回的错误

用户使用SDK调用BCI相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

``` go
// client 为已创建的BCI Client对象
args := &bci.ListInstanceArgs{}
result, err := client.ListInstances(args)
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

客户端异常表示客户端尝试向BCI发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当BCI服务端出现异常时，BCI服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[BCI错误码](https://cloud.baidu.com/doc/BCI/s/Vlf18flyw)

## SDK日志

BCI GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

BCI GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

```go
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

```go
// import "github.com/baidubce/bce-sdk-go/util/log"

// 直接使用包级别全局日志对象（会和GO SDK自身日志一并输出）
log.SetLogHandler(log.STDERR)
log.Debugf("%s", "logging message using the log package in the BCI go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the BCI go sdk")
```


# 版本变更记录

首次发布:

 - 支持创建实例、查询实例列表、查询实例详情、删除实例、批量删除实例接口。