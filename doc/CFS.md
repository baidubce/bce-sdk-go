# CFS服务

# 概述

本文档主要介绍CFS GO SDK的使用。在使用本文档前，您需要先了解普通型CFS的一些基本知识。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[CFS访问域名](https://cloud.baidu.com/doc/CFS/s/pjwvy1siw)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云CFS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问CFS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建CFS Client

CFS Client是CFS控制面服务的客户端，为开发者与CFS控制面服务进行交互提供了一系列的方法。

### 使用AK/SK新建CFS Client

通过AK/SK方式访问CFS，用户可以参考如下代码新建一个CFS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/cfs"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个CFSClient
	cfsClient, err := cfs.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/CFS/index.html)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为CFS的控制面服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`cfs.bj.baidubce.com`。

### 使用STS创建CFS Client

**申请STS token**

CFS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问CFS，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建CFS Client**

申请好STS后，可将STS Token配置到CFS Client中，从而实现通过STS Token创建CFS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建CFS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/cfs" //导入CFS服务模块
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

	// 使用申请的临时STS创建CFS控制面服务的Client对象，Endpoint使用默认值
	cfsClient, err := cfs.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create cfs client failed:", err)
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
	cfsClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置CFS Client时，无论对应CFS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问CFS

CFS支持HTTPS传输协议，您可以通过在创建CFS Client对象时指定的Endpoint中指明HTTPS的方式，在CFS GO SDK中使用HTTPS访问CFS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cfs"

ENDPOINT := "https://cfs.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
cfsClient, _ := cfs.NewClient(AK, SK, ENDPOINT)
```

## 配置CFS Client

如果用户需要配置CFS Client的一些细节的参数，可以在创建CFS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问CFS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cfs"

//创建CFS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cfs.bj.baidubce.com
client, _ := cfs.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/cfs"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cfs.bj.baidubce.com"
client, _ := cfs.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/cfs"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cfs.bj.baidubce.com"
client, _ := cfs.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问CFS时，创建的CFS Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建CFS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

CFS通过对标准NFS协议的支持，兼容POSIX接口，为云上的虚机、容器资源提供了跨操作系统的文件存储及共享能力。同时，百度智能云CFS提供简单、易操作的对外接口，并支持按实际使用量计费（公测期间免费），免去部署、维护费用的同时，最大化提升您的业务效率

## 文件系统管理

### 创建文件系统实例

通过以下代码，可以创建一个CFS文件系统实例，返回对应的文件实例ID

```go
args := &cfs.CreateFSArgs{
    ClientToken: "be31b98c-5e41-4838-9830-9be700de5a20",
    // 设置实例名称
    Name:        "sdkCFS",
    // 设置实例所属vpc
    VpcId:       vpcId,
    // 设置实例所属协议类型：1.nfs 2.smb
    Protocol:       protocol,
    // 设置实例所属可用区 
    Zone:    zone,
}


result, err := client.CreateFS(args)
if err != nil {
    fmt.Println("create cfs failed:", err)
} else {
    fmt.Println("create cfs success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CFS API 文档[CreateFileSystem创建文件系统](https://cloud.baidu.com/doc/CFS/s/mjwvy1reo#createfilesystem%E5%88%9B%E5%BB%BA%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F)

### 更新实例

通过以下代码，可以更新一个CFS实例的配置信息，如实例名称

```go
args := &cfs.UpdateFSArgs{
	// 实例ID
	FSID:       "cfs-xxxxx"
   	Name:        "testSdk", 
}
err := client.UpdateFS(args)
if err != nil {
    fmt.Println("update cfs failed:", err)
} else {
    fmt.Println("update cfs success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CFS API 文档[UpdateFileSystem更新文件系统](https://cloud.baidu.com/doc/CFS/s/mjwvy1reo#updatefilesystem%E6%9B%B4%E6%96%B0%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F)

### 查询已有的实例

通过以下代码，可以查询用户账户下所有CFS的信息

```go
args := &cfs.DescribeFSArgs{}

// 支持按fsId、userId，匹配规则支持部分包含（不支持正则）
args.FSID = cfsId
args.UserId = userId


result, err := client.DescribeFS(args)
if err != nil {
    fmt.Println("list all cfs failed:", err)
} else {
    fmt.Println("list all cfs success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CFS API 文档[DescribeFileSystem查询文件系统](https://cloud.baidu.com/doc/CFS/s/mjwvy1reo#describefilesystem%E6%9F%A5%E8%AF%A2%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F)


### 释放实例

通过以下代码，可以释放指定CFS实例，被释放的CFS无法找回

```go
args := &cfs.DropFSArgs{}
args.FSID = cfsId

err := client.DropFS(args)
if err != nil {
    fmt.Println("delete cfs failed:", err)
} else {
    fmt.Println("delete cfs success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CFS API 文档[DropFileSystem释放文件系统实例](https://cloud.baidu.com/doc/CFS/s/mjwvy1reo#dropfilesystem%E9%87%8A%E6%94%BE%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%E5%AE%9E%E4%BE%8B)


## 挂载点管理

### 创建挂载点

通过以下代码，在指定CFS实例下，创建一个文件系统的挂载点，返回domain

```go
args := &cfs.CreateMountTargetArgs{
    // 所属文件系统实例ID
    FSID: cfsId,
    // 所属子网ID
    SubnetId: subnetId,
    // 所属vpc短ID
    VpcID: vpcId,
}
err := client.CreateMountTarget(args)
if err != nil {
    fmt.Println("create Mount Target failed:", err)
} else {
    fmt.Println("create Mount Target success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CFS API 文档[CreateMountTarget创建挂载点](https://cloud.baidu.com/doc/CFS/s/mjwvy1reo#createmounttarget%E5%88%9B%E5%BB%BA%E6%8C%82%E8%BD%BD%E7%82%B9)


### 查询挂载点

通过以下代码，查询指定CFS实例下下所有挂载点信息，支持按挂载点匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定

```go
args := &cfs.DescribeMountTargetArgs{
    // 要查询的文件系统实例id
    FSID: cfsid,
}
result, err := client.DescribeMountTarget(args)
if err != nil {
    fmt.Println("describe Mount Target failed:", err)
} else {
    fmt.Println("describe Mount Target success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CFS API 文档[DescribeMountTarget描述挂载点](https://cloud.baidu.com/doc/CFS/s/mjwvy1reo#describemounttarget%E6%8F%8F%E8%BF%B0%E6%8C%82%E8%BD%BD%E7%82%B9)


### 删除挂载点

通过以下代码，释放指定CFS实例下的挂载点

```go
args := &cfs.DropMountTargetArgs{
    // 要删除的文件系统实例ID
    FSID:    cfsId,
    MountId: mountId,
}
err := client.DropMountTarget(args)
if err != nil {
    fmt.Println("delete Mount Target failed:", err)
} else {
    fmt.Println("delete Mount Target success: ")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考CFS API 文档[DeleteMountTarget删除挂载点](https://cloud.baidu.com/doc/CFS/s/mjwvy1reo#deletemounttarget%E5%88%A0%E9%99%A4%E6%8C%82%E8%BD%BD%E7%82%B9)


# 错误处理

GO语言以error类型标识错误，CFS支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | CFS服务返回的错误

用户使用SDK调用CFS相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```go
// cfsClient 为已创建的CFS Client对象
describeArgs := &DescribeFSArgs {
     FSID: CFS_ID,
}
cfsDetail, err := cfsClient.DescribeFS(describeArgs)
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
	fmt.Println("get cfs detail success: ", cfsDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向CFS发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当CFS服务端出现异常时，CFS服务端会返回给用户相应的错误信息，以便定位问题

# 版本变更记录

## v0.9.107 [2022-02-28]

首次发布：

 - 创建、查看、列表、更新、删除CFS实例
 - 创建、查看、列表、更新、删除挂载点
