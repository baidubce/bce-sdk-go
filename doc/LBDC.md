# LBDC服务

# 概述

本文档主要介绍负载均衡专属集群LBDC GO SDK的使用。在使用本文档前，您需要先了解BLB的一些基本知识。若您还不了解BLB，可以参考[产品描述](https://cloud.baidu.com/doc/BLB/s/Ajwvxno34)和[入门指南](https://cloud.baidu.com/doc/BLB/s/cjwvxnr91)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BLB访问域名](https://cloud.baidu.com/doc/BLB/s/cjwvxnzix)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云LBDC，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问LBDC做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建负载均衡专属集群LBDC Client

LBDC Client是负载均衡专属集群LBDC服务的客户端，为开发者与LBDC服务进行交互提供了一系列的方法。

### 使用AK/SK新建负载均衡专属集群LBDC Client

通过AK/SK方式访问LBDC，用户可以参考如下代码新建一个LBDC Client：
```go
import (
	"github.com/baidubce/bce-sdk-go/services/lbdc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个LBDC Client
	lbdcClient, err := lbdc.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/BLB/s/ojwvynrqn)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为LBDC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`blb.bj.baidubce.com`。

### 使用STS创建LBDC Client

**申请STS token**

LBDC可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问LBDC，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建LBDC Client**

申请好STS后，可将STS Token配置到LBDC Client中，从而实现通过STS Token创建LBDC Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建LBDC Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/lbdc"   //导入LBDC服务模块
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

	// 使用申请的临时STS创建LBDC服务的Client对象，Endpoint使用默认值
	lbdcClient, err := lbdc.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create lbdc client failed:", err)
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
	lbdcClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置LBDC Client时，无论对应LBDC服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问LBDC

LBDC支持HTTPS传输协议，您可以通过在创建LBDC Client对象时指定的Endpoint中指明HTTPS的方式，在LBDC GO SDK中使用HTTPS访问LBDC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

ENDPOINT := "https://lbdc.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
lbdcClient, _ := lbdc.NewClient(AK, SK, ENDPOINT)
```

## 配置LBDC Client

如果用户需要配置LBDC Client的一些细节的参数，可以在创建LBDC Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问LBDC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

//创建LBDC Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com"
client, _ := lbdc.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com"
client, _ := lbdc.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com"
client, _ := lbdc.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问LBDC时，创建的LBDC Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建LBDC Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

负载均衡专属集群提供性能可控、资源独享、物理资源隔离的专属负载均衡服务，满足超高性能和独占资源需求。

## 实例管理

### 创建LBDC

通过以下代码，可以创建LBDC
```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

args := &CreateLbdcArgs{
    ClientToken: ClientToken(),
    Name:        Name,
    Type:        Type,
    CcuCount:    CcuCount,
    Description: &Description,
    Billing: &Billing{
        PaymentTiming: PaymentTiming,
        Reservation: &Reservation{
            ReservationLength: ReservationLength,
        },
    },
    RenewReservation: &Reservation{
        ReservationLength: ReservationLength,
    },
}
res, err := client.CreateLbdc(args)
ExpectEqual(t.Errorf, nil, err)
e, err := json.Marshal(res)
if  err != nil {
    fmt.Printf("create lbdc error: %+v\n", err)
    return
}
fmt.Printf("create lbdc success,lbdcId is: %+v",res.Id)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考LBDC API 文档[创建LBDC](https://cloud.baidu.com/doc/BLB/s/6kszzygx4#%E5%88%9B%E5%BB%BAlbdc)

### 升级LBDC

通过以下代码，可以升级LBDC
```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

args := &UpgradeLbdcArgs{
    ClientToken: ClientToken(),
    Id:          Id,
    CcuCount:    CcuCount,
}
err := client.UpgradeLbdc(args)
if  err != nil {
    fmt.Printf("upgrade lbdc error: %+v\n", err)
    return
}
fmt.Printf("upgrade lbdc success")
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考LBDC API 文档[升级LBDC](https://cloud.baidu.com/doc/BLB/s/6kszzygx4#%E5%8D%87%E7%BA%A7lbdc)

### 续费LBDC

通过以下代码，可以续费LBDC
```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

args := &RenewLbdcArgs{
    ClientToken: ClientToken(),
    Id:          Id,
    Billing: &BillingForRenew{
		Reservation: &Reservation{
			ReservationLength: ReservationLength,
		},
	},
}
err := client.RenewLbdc(args)
if  err != nil {
    fmt.Printf("renew lbdc error: %+v\n", err)
    return
}
fmt.Printf("renew lbdc success")
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考LBDC API 文档[续费LBDC](https://cloud.baidu.com/doc/BLB/s/6kszzygx4#%E7%BB%AD%E8%B4%B9lbdc)

### 查询LBDC列表

通过以下代码，可以查询LBDC列表
```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

args := &ListLbdcArgs{
    Id:          Id,
    Name:        Name,
}
res, err := client.ListLbdc(args)
if  err != nil {
    fmt.Printf("get lbdc list error: %+v\n", err)
    return
}

// 返回标记查询的起始位置
fmt.Println("lbdc list marker: ", res.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("lbdc list isTruncated: ", res.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("lbdc list nextMarker: ", res.NextMarker)
// 每页包含的最大数量
fmt.Println("lbdc list maxKeys: ", res.MaxKeys)
// 获取lbdc的具体信息
for _, v := range result.ClusterList {
    // 集群id
    fmt.Println("Cluster id: ", v.Id)
    // 集群名称
    fmt.Println("Cluster name: ", v.Name)
    // 集群类型
    fmt.Println("Cluster type: ", v.Type)
    // 集群状态
    fmt.Println("Cluster status: ", v.Status)
    // 集群性能容量
    fmt.Println("Cluster ccuCount: ", v.CcuCount)
    // 集群创建时间
    fmt.Println("Cluster createTime: ", v.CreateTime)
    // 集群失效时间
    fmt.Println("Cluster expireTime: ", v.ExpireTime)
    // 描述
    fmt.Println("Cluster description: ", v.Description)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考LBDC API 文档[查询LBDC列表](https://cloud.baidu.com/doc/BLB/s/6kszzygx4#lbdc%E5%88%97%E8%A1%A8)

### 查询LBDC详情

通过以下代码，可以查询LBDC详情
```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

res, err := client.GetLbdcDetail(lbdcId)
if  err != nil {
    fmt.Printf("get lbdc detail error: %+v\n", err)
    return
}

// 集群id
fmt.Println("lbdc id: ", res.VpnId)
// 集群名称
fmt.Println("lbdc name: ", res.Name)
// 集群类型
fmt.Println("lbdc type: ", res.Type)
// 集群状态
fmt.Println("lbdc status: ", res.Status)
// 集群性能容量
fmt.Println("lbdc ccuCount: ", res.CcuCount)
// 集群创建时间
fmt.Println("lbdc createTime: ", res.CreateTime)
// 集群失效时间
fmt.Println("lbdc expireTime: ", res.ExpireTime)
// 集群并发连接数
fmt.Println("lbdc totalConnectCount: ", res.TotalConnectCount)
// 集群网络输入带宽
fmt.Println("lbdc networkInBps: ", res.NetworkInBps)
// 集群网络输出带宽
fmt.Println("lbdc networkOutBps: ", res.NetworkOutBps)

// if 4 layers
// 集群新建连接速度
fmt.Println("lbdc newConnectCps: ", res.NewConnectCps)

// if 7 layers
// 集群https的qps
fmt.Println("lbdc httpsQps: ", res.HttpsQps)
// 集群http的qps
fmt.Println("lbdc httpQps: ", res.HttpQps)
// 集群http新建速度
fmt.Println("lbdc httpNewConnectCps: ", res.HttpNewConnectCps)
// 集群https新建速度
fmt.Println("lbdc httpsNewConnectCps: ", res.HttpsNewConnectCps)
// 集群ssl新建速度
fmt.Println("lbdc sslNewConnectCps: ", res.SslNewConnectCps)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考LBDC API 文档[LBDC详情](https://cloud.baidu.com/doc/BLB/s/6kszzygx4#lbdc%E8%AF%A6%E6%83%85)

### 更新LBDC

通过以下代码，可以更新LBDC名称或者描述
```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

args := &UpdateLbdcArgs{
    ClientToken: ClientToken(),
    Id:          Id,
    UpdateLbdcBody: &UpdateLbdcBody{
		Name:        &Name,
		Description: &Description,
	},
}
err := client.UpdateLbdc(args)
if  err != nil {
    fmt.Printf("update lbdc error: %+v\n", err)
    return
}
fmt.Printf("update lbdc success")
```
> **提示：**
> 1. 名称和描述都可以为空
> 2. 注意名称和描述是指针类型
> 3. 详细的参数配置及限制条件，可以参考LBDC API 文档[更新LBDC](https://cloud.baidu.com/doc/BLB/s/6kszzygx4#%E6%9B%B4%E6%96%B0lbdc)

### 查询LBDC关联的BLB列表

通过以下代码，可以查询LBDC关联的BLB列表
```go
// import "github.com/baidubce/bce-sdk-go/services/lbdc"

res, err := client.GetBoundBlBListOfLbdc(lbdcId)
if  err != nil {
    fmt.Printf("get bound blb list of lbdc error: %+v\n", err)
    return
}

// 获取lbdc关联的BLB的具体信息
for _, v := range result.BlbList {
    // 负载均衡id
    fmt.Println("BLB id: ", v.BlbId)
    // blb名称
    fmt.Println("BLB name: ", v.Name)
    // blb状态
    fmt.Println("BLB status: ", v.Status)
    // blb类型
    fmt.Println("BLB type: ", v.BlbType)
    // 公网ip
    fmt.Println("BLB publicIp: ", v.PublicIp)
    // eip线路类型
    fmt.Println("BLB eipRouteType: ", v.EipRouteType)
    // 带宽
    fmt.Println("BLB bandwidth: ", v.Bandwidth)
    // 内网ip地址
    fmt.Println("BLB address: ", v.Address)
    // ipv6地址
    fmt.Println("BLB ipv6: ", v.Ipv6)
    // vpcId
    fmt.Println("BLB vpcId: ", v.VpcId)
    // 子网id
    fmt.Println("BLB subnetId: ", v.SubnetId)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考LBDC API 文档[查询LBDC关联的BLB列表](https://cloud.baidu.com/doc/BLB/s/6kszzygx4#lbdc%E5%85%B3%E8%81%94%E7%9A%84blb%E5%88%97%E8%A1%A8)


# 错误处理

GO语言以error类型标识错误，LBDC支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | LBDC服务返回的错误

用户使用SDK调用LBDC相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```go
// client 为已创建的LBDC Client对象
lbdcDetail, err := client.GetLbdcDetail(lbdcId)
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
	fmt.Println("get lbdc detail success: ", lbdcDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向LBDC发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当LBDC服务端出现异常时，LBDC服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[BLB错误返回](https://cloud.baidu.com/doc/BLB/s/Djwvxnzw6)

# 版本变更记录

## v1.0.0 [2022-09-21]

首次发布：

 - 创建、升级、续费、查看、列表、更新LBDC实例，查LBDC关联的BLB列表


