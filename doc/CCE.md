# CCE服务

# 概述

本文档主要介绍CCE GO SDK的使用。在使用本文档前，您需要先了解CCE的一些基本知识，并已开通了CCE服务。若您还不了解CCE，可以参考[产品描述](https://cloud.baidu.com/doc/CCE/s/Bjwvy0x5g)和[操作指南](https://cloud.baidu.com/doc/CCE/s/zjxpoqohb)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[CCE服务域名](https://cloud.baidu.com/doc/CCE/s/Fjwvy1fl4)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/CCE/s/Fjwvy1fl4)。

目前支持“华北-北京”、“华南-广州”、“华东-苏州”、“香港”、“金融华中-武汉”和“华北-保定”六个区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
BJ | cce.bj.baidubce.com | HTTP and HTTPS
GZ | cce.gz.baidubce.com | HTTP and HTTPS
SU | cce.su.baidubce.com | HTTP and HTTPS
HKG| cce.hkg.baidubce.com| HTTP and HTTPS
FWH| cce.fwh.baidubce.com| HTTP and HTTPS
BD | cce.bd.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云CCE，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问CCE做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建CCE Client

CCE Client是CCE服务的客户端，为开发者与CCE服务进行交互提供了一系列的方法。

### 使用AK/SK新建CCE Client

通过AK/SK方式访问CCE，用户可以参考如下代码新建一个CCE Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/cce"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个CCEClient
	cceClient, err := cce.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为CCE的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`cce.bj.baidubce.com`。

### 使用STS创建CCE Client

**申请STS token**

CCE可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问CCE，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建CCE Client**

申请好STS后，可将STS Token配置到CCE Client中，从而实现通过STS Token创建CCE Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建CCE Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/cce" //导入CCE服务模块
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

	// 使用申请的临时STS创建CCE服务的Client对象，Endpoint使用默认值
	cceClient, err := cce.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "cce.bj.baidubce.com")
	if err != nil {
		fmt.Println("create cce client failed:", err)
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
	cceClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置CCE Client时，无论对应CCE服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问CCE

CCE支持HTTPS传输协议，您可以通过在创建CCE Client对象时指定的Endpoint中指明HTTPS的方式，在CCE GO SDK中使用HTTPS访问CCE服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

ENDPOINT := "https://cce.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
cceClient, _ := cce.NewClient(AK, SK, ENDPOINT)
```

## 配置CCE Client

如果用户需要配置CCE Client的一些细节的参数，可以在创建CCE Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问CCE服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

//创建CCE Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cce.bj.baidubce.com"
client, _ := cce.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cce.bj.baidubce.com"
client, _ := cce.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cce.bj.baidubce.com"
client, _ := cce.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问CCE时，创建的CCE Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建CCE Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# CCE管理

百度智能云容器引擎(Cloud Container Engine，即CCE)是高度可扩展的高性能容器管理服务，您可以在托管的云服务器实例集群上轻松运行应用程序。

> 注意:
> - 百度智能云容器引擎免费为用户提供服务，只会对其使用的资源例如BCC、BLB和EIP等资源收费

## 获取容器网络
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"
args := &cce.GetContainerNetArgs{
	// 容器所在VPC ID
	VpcShortId: "vpc-xxxxxx",
	VpcCidr:    "192.168.0.0/24",
}
result, err := client.GetContainerNet(args)
if err != nil {
    fmt.Printf("get container net error: %+v\n", err)
	return
}
fmt.Printf("get cce container net success: %s\n", result.ContainerNet)
```

## 获取支持版本列表
```go
result, err := client.ListVersions()
if err != nil {
	fmt.Printf("list version error: %+v\n", err)
	return
}
fmt.Printf("list version success with result: %+v\n", result)
```

## 创建CCE Cluster
使用以下代码可以创建一个CCE Cluster。
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.CreateClusterArgs{
    // 指定CCE Cluster名称
	ClusterName:       "sdk-test",
	Version:           listVersionResult.Data[0],
	MainAvailableZone: "zoneA",
	ContainerNet:      getContainernetResult.ContainerNet,
    // CCE 集群高级配置
	//AdvancedOptions:    cce.AdvancedOptions{},
    // CCE CDS盘预挂载信息
	//CdsPreMountInfo:    cce.CdsPreMountInfo{},
	Comment:           "sdk create",
	DeployMode:        cce.DeployModeBcc,
	OrderContent: &cce.BaseCreateOrderRequestVo{Items: []cce.Item{
		{
			Config: cce.BccConfig{
                // BCC实例名称，若不指定，将随机生成
				Name:            "sdk-create",
				ProductType:     cce.ProductTypePostpay,
				InstanceType:    cce.InstanceTypeG3,
				Cpu:             1,
				Memory:          2,
				ImageType:       cce.ImageTypeCommon,
				SubnetUuid:      SubnetId,
				SecurityGroupId: Security,
				AdminPass:       AdminPass,
				PurchaseNum:     2,
				ImageId:         "m-Nlv9C0tF",
				ServiceType:     cce.ServiceTypeBCC,
			},
		},
     }},
}
result, err := client.CreateCluster(args)
if err != nil {
	fmt.Printf("create cce cluster error: %+v\n", err)
	return
}
fmt.Printf("create cce cluster success with result: %+v\n", result)
```

## 获取CCE Cluster列表
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.ListClusterArgs{}

// 如果想获取某些状态下的集群，如RUNNING状态下的，可以配置参数
args.Status = cce.ClusterStatusRunning

result, err := client.ListClusters(args)
if err != nil {
	fmt.Printf("list cce cluster error: %+v\n", err)
	return
}
fmt.Printf("list cce cluster success with result: %+v\n", result)
```

## 获取CCE Cluster详情
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

result, err := client.GetCluster("cluster_uuid")
if err != nil {
	fmt.Printf("get cce cluster error: %+v\n", err)
	return
}
fmt.Printf("get cce cluster success with result: %+v\n", result)
```

## 获取CCE Cluster Node信息
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.ListNodeArgs{
	ClusterUuid: "cluster_uuid",
}

result, err := client.ListNodes(args)
if err != nil {
	fmt.Printf("list cce cluster nodes error: %+v\n", err)
	return
}
fmt.Printf("list cce cluster nodes success with result: %+v\n", result)
```

## CCE Cluster扩容
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.ScalingUpArgs{
	ClusterUuid: "cluster_uuid",
	// CCE CDS盘预挂载信息
	//CdsPreMountInfo: cce.CdsPreMountInfo{},
	OrderContent: &cce.BaseCreateOrderRequestVo{Items: []cce.Item{
		{
			Config: cce.BccConfig{
                // BCC实例名称，若不指定，将随机生成
				Name:            "sdk-create",
				ProductType:     cce.ProductTypePostpay,
				InstanceType:    cce.InstanceTypeG3,
				Cpu:             1,
				Memory:          2,
				ImageType:       cce.ImageTypeCommon,
				SubnetUuid:      SubnetId,
				SecurityGroupId: Security,
				AdminPass:       AdminPass,
				PurchaseNum:     2,
				ImageId:         "m-Nlv9C0tF",
				ServiceType:     cce.ServiceTypeBCC,
			},
		},
	}},
}

result, err := client.ScalingUp(args)
if err != nil {
	fmt.Printf("scaling up cce cluster error: %+v\n", err)
	return
}
fmt.Printf("scaling up cce cluster success with result: %+v\n", result)
```

## CCE Cluster缩容
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.ScalingDownArgs{
	ClusterUuid:  "cluster_uuid",
    // 可选择是否连带删除EIP和CDS
	DeleteEipCds: true,
    // 可选择是否连带删除快照
	DeleteSnap:   true,
	NodeInfo: []cce.NodeInfo{
		{
			InstanceId: "instance_id",
		},
	},
}

err = client.ScalingDown(args)
if err != nil {
	fmt.Printf("scaling down cce cluster error: %+v\n", err)
	return
}
```

## 移除CCE Cluster节点
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.ShiftOutNodeArgs{
		ClusterUuid: "cluster_uuid",
		NodeInfoList: []cce.CceNodeInfo{
			{InstanceId: "instance_id"},
		},
}

err = client.ShiftOutNode(args)
if err != nil {
	fmt.Printf("shift node out from cce cluster error: %+v\n", err)
	return
}
```

## 移入CCE Cluster节点
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.ShiftInNodeArgs{
	ClusterUuid: "cluster_uuid",
    // 是否要重装系统
	NeedRebuild: false,
    // 若重装系统，则选择系统镜像ID
	//ImageId:      "",
    // 若重装系统，则配置新的密码
	//AdminPass:    AdminPass,
	InstanceType: cce.ShiftInstanceTypeBcc,
	NodeInfoList: []cce.CceNodeInfo{
		{
			InstanceId: "instance_id",
            // 若选择不重装系统，则需要输入该节点密码
			AdminPass:  AdminPass,
		},
	},
}

err = client.ShiftInNode(args)
if err != nil {
	fmt.Printf("shift node into cce cluster error: %+v\n", err)
	return
}
```

## 获取CCE Cluster可移入节点列表
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.ListExistedNodeArgs{
	ClusterUuid: "cluster_uuid",
}

result, err := client.ListExistedBccNode(args)
if err != nil {
	fmt.Printf("list cce cluster exist node error: %+v\n", err)
	return
}
fmt.Printf("list cce cluster exist node success with result: %+v\n", result)
```

## 获取CCE Cluster kubeconfig配置
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.GetKubeConfigArgs{
	ClusterUuid: "cluster_uuid",
}

// 若要获取公网访问配置，可以选择默认default
args.Type = cce.KubeConfigTypeDefault

// 若要获取VPC内部访问配置，可以选择
args.Type = cce.KubeConfigTypeIntraVpc

// 若要获取B区内部访问配置，可以选择
args.Type = cce.KubeConfigTypeInternal

configResult, err := client.GetKubeConfig(configArgs)
if err != nil {
	fmt.Printf("create cce cluster error: %+v\n", err)
	return
}
fmt.Printf("create cce cluster success with result: %+v\n", result)
```

## 删除CCE 集群
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

args := &cce.DeleteClusterArgs{
	ClusterUuid:  "cluster_uuid",
    // 是否需要删除EIP和CDS
	DeleteEipCds: true,
    // 是否需要删除快照
	DeleteSnap:   true,
}
err = client.DeleteCluster(args)
if err != nil {
	fmt.Printf("delete cce cluster error: %+v\n", err)
	return
}
```

# 错误处理

GO语言以error类型标识错误，CCE支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | CCE服务返回的错误

用户使用SDK调用CCE相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// client 为已创建的cce Client对象
args := &cce.ListClusterArgs{}
result, err := client.ListClusters(args)
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

客户端异常表示客户端尝试向CCE发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当CCE服务端出现异常时，CCE服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[CCE错误码](https://cloud.baidu.com/doc/CCE/s/4jwvy1evj)

## SDK日志

CCE GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

CCE GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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
log.Debugf("%s", "logging message using the log package in the CCE go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the CCE go sdk")
```


# 版本变更记录

## v0.9.5 [2020-07-02]

首次发布:

 - 获取CCE支持版本列表，获取CCE容器网络。
 - 支持创建CCE，获取CCE详情，获取CCE列表，CCE扩容，CCE缩容，获取CCE节点信息，删除CCE。
 - 节点移入CCE集群，节点移处CCE集群，获取可移入节点列表。