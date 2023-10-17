# GO SDK 文档

# 概述

本文档主要介绍百度云Go语言版的开发者工具包（SDK），用户可基于该SDK使用Go语言接入百度云的各项产品（详见支持产品列表）。SDK封装了便捷的调用接口，保持了多种编程语言版的使用方式、调用接口相似，提供了统一的错误码和返回格式，方便开发者调试。

# 安装SDK工具包

## 运行环境

GO SDK可以在go1.3及以上环境下运行。

## 安装SDK

**直接从github下载**

使用`go get`工具从github进行下载：

```shell
go get github.com/baidubce/bce-sdk-go
```

**SDK目录结构**

```text
bce-sdk-go
|--auth                   //BCE签名和权限认证
|--bce                    //BCE公用基础组件
|--http                   //BCE的http通信模块
|--services               //BCE相关服务目录
|  |--appblb              //应用型负载均衡服务目录
|  |--as                  //弹性伸缩
|  |--bbc                 //物理服务器
|  |--bcc                 //云服务器
|  |--bcm                 //云监控
|  |--bec                 //百度边缘计算
|  |--bie                 //百度智能边缘
|  |--bls                 //日志服务
|  |--bos                 //BOS服务目录
|  |  |--bos_client.go    //BOS客户端入口
|  |  |--api              //BOS相关API目录
|  |     |--bucket.go     //BOS的Bucket相关API实现
|  |     |--object.go     //BOS的Object相关API实现
|  |     |--multipart.go  //BOS的Multipart相关API实现
|  |     |--module.go     //BOS相关API的数据模型
|  |     |--util.go       //BOS相关API实现使用的工具
|  |--bvw                 //BVW服务目录
|  |--cce                 //容器引擎
|  |--cdn                 //内容分布网络
|  |--cert                //SSL证书服务
|  |--cfc                 //函数计算
|  |--cfs                 //CFS文件存储服务
|  |--cfw                 //云防火墙
|  |--csn                 //云智能网
|  |--ddc                 //DDC数据库专属集群
|  |--dts                 //数据传输服务
|  |--eccr                //容器镜像服务企业版
|  |--eip                 //弹性公网IP
|  |--endpoint            //SNIC服务网卡
|  |--eni                 //ENIC弹性网卡
|  |--esg                 //企业安全组
|  |--etGateway           //专线网关
|  |--havip               //高可用虚拟IP
|  |--iam                 //身份管理
|  |--media               //音视频处理MCP
|  |--rds                 //云数据库
|  |--scs                 //SCS服务目录
|  |--sms                 //SMS服务目录
|  |--sts                 //STS服务目录
|  |--vca                 //VCA服务目录
|  |--vcr                 //VCR服务目录
|  |--vpc                 //私有网络
|  |--vpn                 //VPN网关
|  |--mms                 //多模态媒资检索目录
|--util                   //BCE公用的工具实现
```

## 卸载SDK

预期卸载SDK时，删除下载的源码即可。

# 使用步骤

## 确认Endpoint

在使用SDK之前，需确认您将接入的百度云产品的Endpoint（服务域名）。以百度对象存储产品为例，可阅读[BOS访问域名](https://cloud.baidu.com/doc/BOS/DevRef.html#BOS.E8.AE.BF.E9.97.AE.E5.9F.9F.E5.90.8D)的部分，理解Endpoint相关的概念。其他服务类似，需理解并确认对应服务的Endpoint。

## 创建Client对象

每种具体的服务都有一个`Client`对象，为开发者与对应的服务进行交互封装了一系列易用的方法。开发者可参考SDK中具体服务对应的目录下的说明文档使用相应的服务。

## 调用功能接口

开发者基于创建的对应服务的`Client`对象，即可调用相应的功能接口，使用百度云产品的功能。

## 示例

下面以百度云对象存储服务（BOS）为例，给出一个基本的使用示例，详细使用说明请参考各服务的详细说明文档。

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bos"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := "<your-access-key-id>", "<your-secret-access-key>"

	// BOS服务的Endpoint
	ENDPOINT := "<domain-name>"

	// 创建BOS服务的Client
	bosClient, err := bos.NewClient(AK, SK, ENDPOINT)

	// 创建Bucket
	location, err := bosClient.PutBucket("<your-bucket-name>"); err != nil {
	if err != nil {
		fmt.Println("create bucket failed:", err)
	}
	fmt.Println("create bucket success in", location)

	// 上传对象
	etag, err := bosClient.PutObjectFromFile("<your-bucket-name>", "<your-object-name>", "<file-name>", nil)
	if err != nil {
		fmt.Println("upload file to BOS failed:", err)
	}
	fmt.Println("upload file to BOS success, etag = ", etag)
}
```

# 配置

## 使用HTTPS协议

该SDK支持使用HTTPS协议访问百度云的服务产品。要使用HTTPS协议，只需在您创建对应服务的`Client`对象时指定的Endpoint中指明使用https协议的域名即可，SDK会自动识别并使用HTTPS协议访问。

## 详细配置

开发者使用GO SDK时，创建的对应服务的`Client`对象，其导出字段`Config`提供如下参数以便支持详细配置：

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


开发者可据此进行详细参数的配置，下面给出部分配置示例：

```go
// client为某一种具体服务的`Client`对象

// 配置请求代理地址
client.Config.ProxyUrl = "127.0.0.1:8080"

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000

// 配置签名使用的HTTP请求头为`Host`
client.Config.SignOption.HeadersToSign = map[string]struct{}{"Host": struct{}{}}

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

# 错误处理

GO语言以error类型标识错误，定义了如下两种错误类型：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | BOS服务返回的错误

用户使用SDK调用各服务的相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误的详细信息进行处理。实例如下：

```go
// bosClient 为已创建的BOS服务的Client对象
bucketLocation, err := bosClient.PutBucket("test-bucket")
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
fmt.Println("create bucket success, bucket location:", bucketLocation)
```

## 客户端异常

客户端异常表示客户端尝试向百度云服务发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当服务端出现异常时，百度云服务端会返回给用户相应的错误信息，以便定位问题。每种服务端的异常需参考各服务的官网文档。

## SDK日志

GO SDK自行实现了支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

该日志模块无任何外部依赖，开发者使用GO SDK开发项目，可以直接引用该日志模块自行在项目中使用。可使用GO SDK使用的包级别的日志对象，也可创建新的日志对象，详见如下示例：

```go
// 直接使用包级别全局日志对象（会和GO SDK自身日志一并输出）
log.SetLogHandler(log.STDERR)
log.Debugf("%s", "logging message using the log package in the sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the sdk")
```

# 支持产品列表

产品名称   | 产品缩写      | 导入路径                                              | 说明文档
-----------|-----------|---------------------------------------------------|----------
应用型负载均衡 | APPBLB    | github.com/baidubce/bce-sdk-go/services/appblb    | [APPBLB.md](./doc/APPBLB.md)
物理服务器 | BBC       | github.com/baidubce/bce-sdk-go/services/bbc       | [BBC.md](./doc/BBC.md)
云服务器 | BCC       | github.com/baidubce/bce-sdk-go/services/bcc       | [BCC.md](./doc/BCC.md)
云监控 | BCM       | github.com/baidubce/bce-sdk-go/services/bcm       | [BCM.md](./doc/BCM.md)
边缘计算节点 | BEC       | github.com/baidubce/bce-sdk-go/services/bec       |[BEC.md](./doc/BEC.md)
百度智能边缘 | BIE       | github.com/baidubce/bce-sdk-go/services/bie       |
负载均衡 | BLB       | github.com/baidubce/bce-sdk-go/services/blb       | [BLB.md](./doc/BLB.md)
日志服务 | BLS       | github.com/baidubce/bce-sdk-go/services/bls       | [BLS.md](./doc/BLS.md)
百度对象存储 | BOS       | github.com/baidubce/bce-sdk-go/services/bos       | [BOS.md](./doc/BOS.md)
百度视频创作分发平台 | BVW       | github.com/baidubce/bce-sdk-go/services/bvw       | [BVW.md](./doc/BVW.md)
容器引擎 | CCE       | github.com/baidubce/bce-sdk-go/services/cce       | [CCE.md](./doc/CCE.md)
内容分布网络 | CDN       | github.com/baidubce/bce-sdk-go/services/cdn       | [CDN.md](./doc/CDN.md)
SSL证书服务 | CERT      | github.com/baidubce/bce-sdk-go/services/cert      | [CERT.md](./doc/CERT.md)
函数计算 | CFC       | github.com/baidubce/bce-sdk-go/services/cfc       | [CFC.md](./doc/CFC.md)
文件存储服务 | CFS       | github.com/baidubce/bce-sdk-go/services/cfs       | [CFS.md](./doc/CFS.md)
云防火墙 | CFW       | github.com/baidubce/bce-sdk-go/services/cfw       | [CFW.md](./doc/CFW.md)
云智能网 | CSN       | github.com/baidubce/bce-sdk-go/services/csn       | [CSN.md](./doc/CSN.md)
文档服务 | DOC       | github.com/baidubce/bce-sdk-go/services/doc       | [DOC.md](./doc/DOC.md)
数据传输服务 | DTS       | github.com/baidubce/bce-sdk-go/services/dts       | [DTS.md](./doc/DTS.md)
容器镜像服务 企业版 | ECCR      | github.com/baidubce/bce-sdk-go/services/eccr      | [ECCR.md](./doc/ECCR.md)
弹性公网IP | EIP       | github.com/baidubce/bce-sdk-go/services/eip       | [EIP.md](./doc/EIP.md)
ENIC弹性网卡 | ENIC      | github.com/baidubce/bce-sdk-go/services/eni       | [ENIC.md](./doc/ENIC.md)
企业安全组 | ESG       | github.com/baidubce/bce-sdk-go/services/esg       | [ESG.md](./doc/ESG.md)
高可用虚拟IP | HAVIP       | github.com/baidubce/bce-sdk-go/services/havip       | [HAVIP.md](./doc/HAVIP.md)
专线网关 | EtGateway | github.com/baidubce/bce-sdk-go/services/etGateway | [ETGateway.md](./doc/ETGateway.md)
身份管理 | IAM       | github.com/baidubce/bce-sdk-go/services/iam       | [IAM.md](./doc/IAM.md)
内网DNS | LOCALDNS  | github.com/baidubce/bce-sdk-go/services/localDns  | [LOCALDNS.md](./doc/LOCALDNS.md)
音视频处理MCP | MCP       | github.com/baidubce/bce-sdk-go/services/media     | [MCP.md](./doc/MCP.md)
云数据库 | RDS       | github.com/baidubce/bce-sdk-go/services/rds       | [RDS.md](./doc/RDS.md)
分布式缓存服务 | SCS       | github.com/baidubce/bce-sdk-go/services/scs       | [SCS.md](./doc/SCS.md)
SMS简单消息服务 | SMS       | github.com/baidubce/bce-sdk-go/services/sms       | [SMS.md](./doc/SMS.md)
SNIC服务网卡 | SNIC      | github.com/baidubce/bce-sdk-go/services/endpoint  | [SNIC.md](./doc/SNIC.md)
安全Token服务| STS       | github.com/baidubce/bce-sdk-go/services/sts       | [STS.md](./doc/STS.md)
视频内容分析 | VCA       | github.com/baidubce/bce-sdk-go/services/vca       |
视频内容审核 | VCR       | github.com/baidubce/bce-sdk-go/services/vcr       |
私有网络 | VPC       | github.com/baidubce/bce-sdk-go/services/vpc       | [VPC.md](./doc/VPC.md)
VPN网关 | VPN       | github.com/baidubce/bce-sdk-go/services/vpn       | [VPN.md](./doc/VPN.md)
多模态媒资检索 | MMS       | github.com/baidubce/bce-sdk-go/services/mms       | [MMS.md](./doc/MMS.md)
数据库专属集群 | DDC       | github.com/baidubce/bce-sdk-go/services/ddc       | [DDC.md](./doc/DDC.md)
