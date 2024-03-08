# RDS服务

# 概述

本文档主要介绍RDS GO SDK的使用。在使用本文档前，您需要先了解RDS的一些基本知识，并已开通了RDS服务。若您还不了解RDS，可以参考[产品描述](https://cloud.baidu.com/doc/RDS/s/ujwvyzdzg)和[操作指南](https://cloud.baidu.com/doc/RDS/s/Qjwvz0ikk)。

本SDK基于官方开放API进行封装，详细参数解释可参考相关API说明文档[API参考](https://cloud.baidu.com/doc/RDS/s/ajwvz0x1m)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[RDS服务域名](https://cloud.baidu.com/doc/RDS/s/Ejwvz0uoq)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

目前支持“华北-北京”、“华北-保定”、“华南-广州”、“华东-苏州”、“金融华中-武汉”、“华东-上海”、“中国香港”、“新加坡”区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
北京 | rds.bj.baidubce.com | HTTP and HTTPS
保定 | rds.bj.baidubce.com | HTTP and HTTPS
广州 | rds.gz.baidubce.com | HTTP and HTTPS
苏州 | rds.su.baidubce.com | HTTP and HTTPS
武汉 | rds.fwh.baidubce.com| HTTP and HTTPS
上海 | rds.fsh.baidubce.com| HTTP and HTTPS
香港 | rds.hkg.baidubce.com| HTTP and HTTPS
新加坡 | rds.sin.baidubce.com| HTTP and HTTPS
## 获取密钥

要使用百度云RDS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问RDS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建RDS Client

RDS Client是RDS服务的客户端，为开发者与RDS服务进行交互提供了一系列的方法。

### 使用AK/SK新建RDS Client

通过AK/SK方式访问RDS，用户可以参考如下代码新建一个RDS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/rds"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个RDSClient
	rdsClient, err := rds.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为VPC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`rds.bj.baidubce.com`。

### 使用STS创建RDS Client

**申请STS token**

RDS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问RDS，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建RDS Client**

申请好STS后，可将STS Token配置到RDS Client中，从而实现通过STS Token创建RDS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建RDS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/rds" //导入RDS服务模块
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

	// 使用申请的临时STS创建RDS服务的Client对象，Endpoint使用默认值
	rdsClient, err := rds.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "rds.bj.baidubce.com")
	if err != nil {
		fmt.Println("create rds client failed:", err)
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
	rdsClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置RDS Client时，无论对应RDS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问RDS

RDS支持HTTPS传输协议，您可以通过在创建RDS Client对象时指定的Endpoint中指明HTTPS的方式，在RDS GO SDK中使用HTTPS访问RDS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

ENDPOINT := "https://rds.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
rdsClient, _ := rds.NewClient(AK, SK, ENDPOINT)
```

## 配置RDS Client

如果用户需要配置RDS Client的一些细节的参数，可以在创建RDS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问RDS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

//创建RDS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "rds.bj.baidubce.com"
client, _ := rds.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "rds.bj.baidubce.com"
client, _ := rds.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "rds.bj.baidubce.com"
client, _ := rds.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问RDS时，创建的RDS Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建RDS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# RDS管理

云数据库 RDS （Relational Database Service）是专业、高性能、高可靠的云数据库服务。云数据库 RDS 提供 Web 界面进行配置、操作数据库实例，还为您提供可靠的数据备份和恢复、完备的安全管理、完善的监控、轻松扩展等功能支持。相对于自建数据库，云数据库 RDS 具有更经济、更专业、更高效、更可靠、简单易用等特点，使您能更专注于核心业务。

## 创建RDS主实例

使用以下代码可以创建一个RDS主实例
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.CreateRdsArgs{
	// 指定rds的数据库引擎，取值mysql，sqlserver，postgresql，必选
    Engine:            "mysql",
    // 指定rds的数据库版本，必选
    EngineVersion:  "5.6",
    // 计费相关参数，PaymentTiming取值为 预付费：Prepaid，后付费：Postpaid；Reservation：支付方式为后支付时不需要设置，预支付时必须设置；必选
    Billing: rds.Billing{
        PaymentTiming: "Postpaid",
        //Reservation: rds.Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
    },
    // 预付费时可指定自动续费参数 AutoRenewTime 和 AutoRenewTimeUnit
    // 自动续费时长（续费单位为year 不大于3，续费单位为month 不大于9）
    // AutoRenewTime: 1,
    // 自动续费单位（"year";"month"）
    // AutoRenewTimeUnit: "year",
    // CPU核数，必选
    CpuCount: 1,
    //套餐内存大小，单位GB，必选
    MemoryCapacity: 1,
    //套餐磁盘大小，单位GB，每5G递增，必选
    VolumeCapacity: 5,
    //磁盘类型, normal_io:本地盘ssd磁盘, cloud_high:高性能云磁盘, cloud_nor:通用型SSD, cloud_enha:增强型SSD, 必选
    DiskIoType: "normal_io",
    //批量创建云数据库 RDS 实例个数, 最大不超过10，默认1，可选
    PurchaseCount: 1,
    //rds实例名称，允许小写字母、数字，长度限制为1~32，默认命名规则:{engine} + {engineVersion}，可选
    InstanceName: "instanceName",
    //所属系列，Basic：单机基础版，Standard：双机高可用版。仅SQLServer 2012sp3 支持单机基础版。默认Standard，可选
    Category: "Standard",
    //指定zone信息，默认为空，由系统自动选择，可选
    //zoneName命名规范是小写的“国家-region-可用区序列"，例如北京可用区A为"cn-bj-a"。
    ZoneNames: []string{"cn-bj-d"},
    //vpc，如果不提供则属于默认vpc，可选
    VpcId: "vpc-IyrqYIQ7",
    //是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
    IsDirectPay: false,
    //vpc内，每个可用区的subnetId；如果不是默认vpc则必须指定 subnetId，可选
    Subnets: []rds.SubnetMap{
        {
            ZoneName: "cn-bj-a",
            SubnetId: "sbn-IyWRnII7",
        },   
    },
    // 实例绑定的标签信息，可选
    Tags: []model.TagModel{
        {
            TagKey:   "tagK",
            TagValue: "tagV",
        },
    },
    // 实例绑定的资源分组，可选
    ResourceGroupId: "RESG-gNjxxxxxx",
}
result, err := client.CreateRds(args)
if err != nil {
    fmt.Printf("create rds error: %+v\n", err)
    return
}

for _, e := range result.InstanceIds {
	fmt.Println("create rds success, instanceId: ", e)
}
```

> 注意: 
> - 实例可选套餐详见(https://cloud.baidu.com/doc/RDS/s/9jwvz0wd3)
> - 创建计费方式为后付费的实例需要账户现金余额+通用代金券大于100；预付费方式的实例则需要账户现金余额大于等于实例费用。
> - 支持批量创建，且如果创建过程中有一个实例创建失败，所有实例将全部回滚，均创建失败。
> - 创建接口为异步创建，可通过查询指定实例详情接口查询实例状态。

## 创建RDS只读实例

使用以下代码可以创建一个RDS只读实例
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.CreateReadReplicaArgs{
    //主实例ID，必选
    SourceInstanceId: "sourceInstanceId"
    // 计费相关参数，只读实例只支持后付费Postpaid，必选
    Billing: rds.Billing{
        PaymentTiming: "Postpaid",
    },
    // CPU核数，必选
    CpuCount: 1,
    //套餐内存大小，单位GB，必选
    MemoryCapacity: 1,
    //套餐磁盘大小，单位GB，每5G递增，必选
    VolumeCapacity: 5,
    //批量创建云数据库 RDS 只读实例个数, 目前只支持一次创建一个,可选
    PurchaseCount: 1,
    //实例名称，允许小写字母、数字，长度限制为1~32，默认命名规则:{engine} + {engineVersion}，可选
    InstanceName: "instanceName",
    //指定zone信息，默认为空，由系统自动选择，可选
    //zoneName命名规范是小写的“国家-region-可用区序列"，例如北京可用区A为"cn-bj-a"。
    ZoneNames: []string{"cn-bj-d"},
    //与主实例 vpcId 相同，可选
    VpcId: "vpc-IyrqYIQ7",
    //是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
    IsDirectPay: false,
    //vpc内，每个可用区的subnetId；如果不是默认vpc则必须指定 subnetId，可选
    Subnets: []rds.SubnetMap{
        {
            ZoneName: "cn-bj-a",
            SubnetId: "sbn-IyWRnII7",
        },   
    },
    // 实例绑定的资源分组，可选
    ResourceGroupId: "RESG-gNjxxxxxx",
}
result, err := client.CreateReadReplica(args)
if err != nil {
    fmt.Printf("create rds readReplica error: %+v\n", err)
    return
}

for _, e := range result.InstanceIds {
	fmt.Println("create rds readReplica success, instanceId: ", e)
}
```

> 注意: 
> - 需要在云数据库 RDS 主实例的基础上进行创建
> - 实例可选套餐详见(https://cloud.baidu.com/doc/RDS/s/9jwvz0wd3)
> - 仅数据库类型为 MySQL 的主实例支持创建只读实例
> - 只读实例的数据库引擎和数据库版本与主实例相同，无需设置，主实例版本最低是 MySQL 5.6
> - 只读实例的磁盘容量不能小于主实例的磁盘容量
> - 只读实例的 vpcId 需跟主实例一致
> - 一个云数据库 RDS 实例，最多只能有 5 个只读实例，且一次只能创建一个
> - 只读实例只支持后付费方式购买

## 创建RDS代理实例

使用以下代码可以创建一个RDS代理实例
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.CreateRdsProxyArgs{
    //主实例ID，必选
    SourceInstanceId: "sourceInstanceId"
    // 计费相关参数，代理实例只支持后付费Postpaid，必选
    Billing: rds.Billing{
        PaymentTiming: "Postpaid",
    },
    // 代理实例节点数。取值范围2，4，6，8，16，必选
    NodeAmount: 2,
    //实例名称，允许小写字母、数字，长度限制为1~32，默认命名规则:{engine} + {engineVersion}，可选
    InstanceName: "instanceName",
    //指定zone信息，默认为空，由系统自动选择，可选
    //zoneName命名规范是小写的“国家-region-可用区序列"，例如北京可用区A为"cn-bj-a"，建议与主实例的可用区保持一致
    ZoneNames: []string{"cn-bj-d"},
    //与主实例 vpcId 相同，可选
    VpcId: "vpc-IyrqYIQ7",
    //是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
    IsDirectPay: false,
    //vpc内，每个可用区的subnetId；如果不是默认vpc则必须指定 subnetId，可选
    Subnets: []rds.SubnetMap{
        {
            ZoneName: "cn-bj-a",
            SubnetId: "sbn-IyWRnII7",
        },   
    },
    // 实例绑定的资源分组，可选
    ResourceGroupId: "RESG-gNjxxxxxx",
}
result, err := client.CreateRdsProxy(args)
if err != nil {
    fmt.Printf("create rds proxy error: %+v\n", err)
    return
}

for _, e := range result.InstanceIds {
	fmt.Println("create rds proxy success, instanceId: ", e)
}
```

> 注意: 
> - 需要在云数据库 RDS 主实例的基础上进行创建
> - 仅数据库类型为 MySQL 的主实例支持创建只读实例
> - 代理实例套餐和主实例套餐绑定，主实例版本最低是MySQL 5.6
> - 每个主实例最多可以创建1个代理实例
> - 需与主实例在同一vpc中
> - 代理实例只支持后付费方式购买

## 查询RDS列表

使用以下代码可以查询RDS列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.ListRdsArgs{
    // 批量获取列表的查询的起始位置，是一个由系统生成的字符串，可选
    Marker: "marker",
    // 指定每页包含的最大数量(主实例)，最大数量不超过1000，缺省值为1000，可选
    MaxKeys: 1,
}
result, err := client.ListRds(args)
if err != nil {
    fmt.Printf("list rds error: %+v\n", err)
    return
}

// 返回标记查询的起始位置
fmt.Println("rds list marker: ", result.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("rds list isTruncated: ", result.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("rds list nextMarker: ", result.NextMarker)
// 每页包含的最大数量
fmt.Println("rds list maxKeys: ", result.MaxKeys)
// 获取rds的列表信息
for _, e := range result.Instances {
    fmt.Println("rds instanceId: ", e.InstanceId)
    fmt.Println("rds instanceName: ", e.InstanceName)
    fmt.Println("rds engine: ", e.Engine)
    fmt.Println("rds engineVersion: ", e.EngineVersion)
    fmt.Println("rds instanceStatus: ", e.InstanceStatus)
    fmt.Println("rds cpuCount: ", e.CpuCount)
    fmt.Println("rds memoryCapacity: ", e.MemoryCapacity)
    fmt.Println("rds volumeCapacity: ", e.VolumeCapacity)
    fmt.Println("rds usedStorage: ", e.UsedStorage)
    fmt.Println("rds paymentTiming: ", e.PaymentTiming)
    fmt.Println("rds instanceType: ", e.InstanceType)
    fmt.Println("rds instanceCreateTime: ", e.InstanceCreateTime)
    fmt.Println("rds instanceExpireTime: ", e.InstanceExpireTime)
    fmt.Println("rds publicAccessStatus: ", e.PublicAccessStatus)
    fmt.Println("rds task: ", e.Task)
    fmt.Println("rds vpcId: ", e.VpcId)
}
```

> 注意:
> - 只能查看属于自己账号的实例列表。
> - 接口将每个主实例和其只读、代理实例分成一组，参数maxKeys代表分组数，也就是主实例的个数.

## 查询指定RDS实例信息

使用以下代码可以查询指定RDS实例信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err := client.GetDetail(instanceId)
if err != nil {
    fmt.Printf("get rds detail error: %+v\n", err)
    return
}

fmt.Println("rds instanceId: ", result.InstanceId)
fmt.Println("rds instanceName: ", result.InstanceName)
fmt.Println("rds engine: ", result.Engine)
fmt.Println("rds engineVersion: ", result.EngineVersion)
fmt.Println("rds instanceStatus: ", result.InstanceStatus)
fmt.Println("rds cpuCount: ", result.CpuCount)
fmt.Println("rds memoryCapacity: ", result.MemoryCapacity)
fmt.Println("rds volumeCapacity: ", result.VolumeCapacity)
fmt.Println("rds usedStorage: ", result.UsedStorage)
fmt.Println("rds paymentTiming: ", result.PaymentTiming)
fmt.Println("rds instanceType: ", result.InstanceType)
fmt.Println("rds instanceCreateTime: ", result.InstanceCreateTime)
fmt.Println("rds instanceExpireTime: ", result.InstanceExpireTime)
fmt.Println("rds publicAccessStatus: ", result.PublicAccessStatus)
fmt.Println("rds vpcId: ", result.VpcId)

```

## 删除RDS实例

使用以下代码可以删除RDS实例。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

//多个实例间用英文半角逗号","隔开，最多可输入10个
if err := client.DeleteRds(instanceIds); err != nil {
    fmt.Printf("delete rds error: %+v\n", err)
    return
}
fmt.Printf("delete rds success\n")
```

> 注意:
> - 只有付费类型为Postpaid或者付费类型为Prepaid且已过期的实例才可以释放。
> - 如果主实例被释放，那么和主实例关联的只读实例和代理实例也会被释放。

## RDS实例扩缩容

使用以下代码可以对RDS实例扩缩容操作。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.ResizeRdsArgs{
	// cpu核数
    CpuCount: 2,
    // 内存大小，单位GB
    MemoryCapacity: 8,
    // 磁盘大小，单位GB，每5G递增
    VolumeCapacity: 20,
    // 代理实例节点数，代理实例变配时此项必填
    NodeAmount: 2,
    // 是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
    IsDirectPay: false,
}
err = client.ResizeRds(instanceId, args)
if err != nil {
    fmt.Printf("resize rds error: %+v\n", err)
    return
}

fmt.Println("resize rds success.")
```

> 注意:
> - 实例可选套餐详见(https://cloud.baidu.com/doc/RDS/s/9jwvz0wd3)
> - 主实例或只读实例变配时至少填写cpuCount、memoryCapacity、volumeCapacity其中的一个。
> - 实例计费方式采用后付费时，可弹性扩缩容；采用预付费方式，不能进行缩容操作。
> - 只有实例available状态时才可以进行扩缩容操作。
> - 实例扩缩容之后会重启一次。
> - 为异步接口，可通过查询实例详情接口查看instanceStatus是否恢复。

## 重启实例

使用以下代码可以重启实例。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

err := client.RebootInstance(instanceId)
if err != nil {
    fmt.Printf("reboot rds error: %+v\n", err)
    return
}
```

## 修改实例名称

使用以下代码可以修改RDS实例名称。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.UpdateInstanceNameArgs{
    InstanceName: "instanceName",
}
err = client.UpdateInstanceName(instanceId, args)
if err != nil {
    fmt.Printf("update instance name error: %+v\n", err)
    return
}
fmt.Printf("update instance name success\n")
```
> 注意:
>
> - 实例名称支持大小写字母、数字以及-_ /.等特殊字符，必须以字母开头，长度1-64。

## 已创建实例自动续费

使用以下代码可以为已创建的预付费实例创建自动续费
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.AutoRenewArgs{
    // 自动续费时长（续费单位为year 不大于3，续费单位为month 不大于9）必选
	AutoRenewTime: 1,
    // 自动续费单位（"year";"month"）必选
	AutoRenewTimeUnit: "year",
    // 实例id集合 必选
    InstanceIds: []string{
        "rds-y9dJu77d", 
        "rds-aQFOoncr",
    },
}
err := client.AutoRenew(args)
if err != nil {
    fmt.Printf("create auto renew error: %+v\n", err)
    return
}
```
> 注意:
>
> - 用于已创建的实例开启自动续费。
> - 可以传入多个实例id，多个实例需保证在同一地域。


## 修改同步模式

使用以下代码可以修改RDS实例同步模式。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.ModifySyncModeArgs{
    //"Async"异步复制，"Semi_sync"半同步复制。
    SyncMode: "Async",
}
err = client.ModifySyncMode(instanceId, args)
if err != nil {
    fmt.Printf("modify syncMode error: %+v\n", err)
    return
}
fmt.Printf("modify syncMode success\n")
```

## 修改连接信息

使用以下代码可以修改RDS域名前缀。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.ModifyEndpointArgs{
    Address: "newAddress",
}
err = client.ModifyEndpoint(instanceId, args)
if err != nil {
    fmt.Printf("modify endpoint error: %+v\n", err)
    return
}
fmt.Printf("modify endpoint success\n")
```

> 注意:
>
> - 只传输域名前缀即可。域名前缀由小写字母和数字组成，以小写字母开头，长度在3-30之间。

## 开关公网访问

使用以下代码可以修改RDS域名前缀。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.ModifyPublicAccessArgs{
    // true or false
    PublicAccess: true,
}
err = client.ModifyPublicAccess(instanceId, args)
if err != nil {
    fmt.Printf("modify public access error: %+v\n", err)
    return
}
fmt.Printf("modify public access success\n")
```


> 注意:
>
> - true：开启公网访问； false：关闭公网访问。
## 修改时间窗口

使用以下代码可以修改操作时间窗口

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.MaintainTimeArgs{
    MaintainStartTime: "14:00:00",
    MaintainDuration:  2,
}
err = client.UpdateMaintainTime(instanceId, args)
if err != nil {
    fmt.Printf("update maintain time error: %+v\n", err)
    return
}
fmt.Printf("update maintain time success\n")
```
## 实例开启关闭修改存储自动扩容配置

使用以下代码可以开启关闭修改存储自动扩容配置

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.DiskAutoResizeArgs{
    FreeSpaceThreshold: 10,
    DiskMaxLimit:       2000,
}
err = client.ConfigDiskAutoResize(instanceId,"open", args)
if err != nil {
    fmt.Printf("config disk auto resize error: %+v\n", err)
    return
}
fmt.Printf("config disk auto resize success\n")
```

## 获取指定实例的自动扩容配置信息

使用以下代码可以获取指定实例的自动扩容配置信息

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err = client.GetAutoResizeConfig(instanceId)
if err != nil {
    fmt.Printf("get config error: %+v\n", err)
    return
}
fmt.Printf("get config success\n")
```

## 实例是否支持启用自动扩容

使用以下代码可以实例是否支持启用自动扩容

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err = client.EnableAutoExpansion(instanceId)
if err != nil {
    fmt.Printf("get enable auto expansion error: %+v\n", err)
    return
}
fmt.Printf("get enable auto expansion success\n")
```

## 可用区迁移

使用以下代码可以操作实例可用区迁移

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.AzoneMigration{
    MasterAzone: "cn-bj-d",
    BackupAzone: "cn-bj-e",
    ZoneNames:   []string{"cn-bj-d", "cn-bj-e"},
    Subnets: []SubnetMap{
        {
            ZoneName: "cn-bj-d",
            SubnetId: "sbn-nedt51qre6r2",
        },
        {
            ZoneName: "cn-bj-e",
            SubnetId: "sbn-hc20wss3idai",
        },
    },
    EffectiveTime: "timewindow",
}
result, err = client.AzoneMigration(instanceId, args)
if err != nil {
    fmt.Printf("azone migration error: %+v\n", err)
    return
}
fmt.Printf("azone migration success\n")
```
# 账号管理

## 创建账号

使用以下代码可以在某个主实例下创建一个新的账号。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &rds.CreateAccountArgs{
	// 账号名称，不能为保留关键字，必选
    AccountName: "accountName",
    // 账号的密码，由字母、数字或下划线组成，长度6～32位，密码需要加密传输，禁止明文传输，必选
    Password: "password",
    // 账号权限类型，Common：普通帐号,Super：super账号。默认为普通账号，可选
    AccountType: "Common",
    // MySQL和SQL Server实例可设置此项，可选
    DatabasePrivileges: []rds.DatabasePrivilege{
            {
                //数据库名称
                DbName: "user_photo_001",
                //授权类型。ReadOnly：只读，ReadWrite：读写
                AuthType: "ReadOnly",
            },   
        },
    // 帐号的描述信息，可选
    Desc: "账号user1", 
    // 帐号归属类型，OnlyMaster：主实例上使用的帐号,RdsProxy：该主实例对应的代理实例上使用的帐号。默认为OnlyMaster账号，可选
    Type: "OnlyMaster",
}
err = client.CreateAccount(instanceId, args)
if err != nil {
    fmt.Printf("create account error: %+v\n", err)
    return
}

fmt.Println("create account success.")
```

> 注意:
> - 实例状态为Available，实例必须是主实例。
> - 没有超出实例最大账号数量。
> - 若实例的数据库引擎为PostgreSQL，则只允许创建Super账号。其它账号和数据库操作通过这个Super账号来管理。
> - 若实例的数据库引擎为MySQL，则允许创建任意类型的账号。
> - 若实例的数据库引擎为SQLServer，则只允许创建Common账号。

## 查询账号列表

使用以下代码可以查询指定实例的账号列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err := client.ListAccount(instanceId)
if err != nil {
    fmt.Printf("list account error: %+v\n", err)
    return
}

// 获取account的列表信息
for _, e := range result.Accounts {
    fmt.Println("rds accountName: ", e.AccountName)
    fmt.Println("rds desc: ", e.Desc)
    fmt.Println("rds status: ", e.Status)
    fmt.Println("rds type: ", e.Type)
    fmt.Println("rds accountType: ", e.AccountType)
}
```

## 查询特定账号信息

使用以下代码可以查询特定账号信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err := client.GetAccount(instanceId,accountName)
if err != nil {
    fmt.Printf("get account error: %+v\n", err)
    return
}

// 获取account的列表信息
fmt.Println("rds accountName: ", result.AccountName)
fmt.Println("rds desc: ", result.Desc)
fmt.Println("rds status: ", result.Status)
fmt.Println("rds type: ", result.Type)
fmt.Println("rds accountType: ", result.AccountType)
```
## 更新账号描述信息

使用以下代码可以更新账号描述信息
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &ModifyAccountDesc{
    Remark: "test",
}
err := client.ModifyAccountDesc(instanceId,accountName,args)
if err != nil {
    fmt.Printf("modify account desc error: %+v\n", err)
    return
}
fmt.Printf("modify account desc success\n")
```

## 删除特定账号信息

使用以下代码可以删除特定账号信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err := client.DeleteAccount(instanceId,accountName)
if err != nil {
    fmt.Printf("delete account error: %+v\n", err)
    return
}
fmt.Printf("delete account success\n")
```
## 更新账号密码

使用以下代码可以更新账号密码
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &UpdatePasswordArgs{
    Password: "test",
}
err := client.UpdateAccountPassword(instanceId,accountName,args)
if err != nil {
    fmt.Printf("update account password error: %+v\n", err)
    return
}
fmt.Printf("update account password success\n")
```
## 更新账号权限
使用以下代码可以更新账号密码
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &UpdateAccountPrivileges{
    DatabasePrivileges: []DatabasePrivilege{{
        DbName:   "test_db",
        AuthType: "ReadOnly",
    }},
}
err := client.UpdateAccountPrivileges(instanceId,accountName,args)
if err != nil {
    fmt.Printf("update account privilege error: %+v\n", err)
    return
}
fmt.Printf("update account privilege success\n")
```
# 参数管理

## 获取参数列表

使用以下代码可以获取一个实例下的数据库参数列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.ListParameters(instanceId)
if err != nil {
    fmt.Printf("get parameter list error: %+v\n", err)
    return
}
data, _ := json.Marshal(result)
fmt.Println(string(data))
fmt.Printf("get parameter list success\n")
fmt.Println(result.Etag)
```

> 注意:
>
> - 在修改配置参数时需要通过该接口获取Etag。

## 修改配置参数

使用以下代码可以云数据库 RDS for MySQL 的参数配置。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.ListParameters(instanceId)
if err != nil {
    fmt.Printf("get parameter list error: %+v\n", err)
    return
}
fmt.Printf("get parameter list success\n")
fmt.Println(result.Etag)

args := &rds.UpdateParameterArgs{
				Parameters:  []rds.KVParameter{
					{
						Name: "connect_timeout",
						Value: "15",
					},
				},
			}
er := client.UpdateParameter(instanceId, result.Etag, args)
if er != nil {
    fmt.Printf("update parameter error: %+v\n", er)
    return
}
fmt.Printf("update parameter success\n")
```

> 注意:
>
> - 在修改配置参数时需要通过获取参数列表接口获取最新的Etag。

## 参数修改历史

使用以下代码可以云数据库 RDS for MySQL 的参数配置。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.ParameterHistory(instanceId)
if err != nil {
    fmt.Printf("get parameter history error: %+v\n", err)
    return
}
fmt.Printf("get parameter history success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

# 数据库管理

## 修改数据库端口

使用以下代码可以修改数据库端口
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.UpdateDatabasePortArgs{
    EntryPort: 3309,
})
err := client.UpdateDatabasePort(instanceId, args)
if err != nil {
    fmt.Printf("update database port error: %+v\n", err)
    return
}
fmt.Printf("update database port success\n")

```

## 获取数据库列表

使用以下代码可以获取数据库列表
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.ListDatabases(instanceId)
if err != nil {
    fmt.Printf("get database list error: %+v\n", err)
    return
}
fmt.Printf("get database list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))

```

## 修改数据库描述

使用以下代码可以修改数据库描述
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.ModifyDatabaseDesc{
    Remark: "test",
}
err := client.ModifyDatabaseDesc(instanceId, "test_db", args)
if err != nil {
    fmt.Printf("modify database discription error: %+v\n", err)
    return
}
fmt.Printf("modify database discriptio success\n")

```

## 删除数据库

使用以下代码可以删除数据库
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

err := client.DeleteDatabase(instanceId, "test_db")
if err != nil {
    fmt.Printf("delete database error: %+v\n", err)
    return
}
fmt.Printf("delete database success\n")

```
## 创建数据库

使用以下代码可以创建数据库
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.CreateDatabaseArgs{
    CharacterSetName: "utf8",
    DbName:           "test_db",
    Remark:           "test_db",
    AccountPrivileges: []AccountPrivilege{
        {
            AccountName: "baidu",
            AuthType:    "ReadOnly",
        },
    },
}
err := client.CreateDatabase(instanceId, args)
if err != nil {
    fmt.Printf("create database error: %+v\n", err)
    return
}
fmt.Printf("create database success\n")

```

# 任务管理

## 获取任务列表

使用以下代码可以获取任务列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.TaskListArgs{
    InstanceId: instanceId,
}
result, err := client.TaskList(args)
if err != nil {
    fmt.Printf("get task list error: %+v\n", err)
    return
}
fmt.Printf("get task list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

# 回收站管理

## 查询回收站实例列表

使用以下代码可以获取回收站实例列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.ListRdsArgs{}
result, err := client.ListRecyclerInstance(args)
if err != nil {
    fmt.Printf("get recycle list error: %+v\n", err)
    return
}
fmt.Printf("get recycle list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 实例开机

使用以下代码可以开机。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.RecyclerRecoverArgs{
    InstanceIds: []string{instanceId},
}
err := client.RecyclerRecover(args)
if err != nil {
    fmt.Printf("recycler recover error: %+v\n", err)
    return
}
fmt.Printf("recycler recover success\n")
```
## 删除单个回收站实例

使用以下代码可以删除单个回收站实例

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

err := client.DeleteRecyclerInstance(instanceId)
if err != nil {
    fmt.Printf("delete error: %+v\n", err)
    return
}
fmt.Printf("delete success\n")
```

# 实例组管理

## 创建实例组
使用以下代码可以创建实例组。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.InstanceGroupArgs{
    Name:     "test_group",
    LeaderId: instanceId,
}
err := client.CreateInstanceGroup(args)
if err != nil {
    fmt.Printf("create instance group error: %+v\n", err)
    return
}
fmt.Printf("create instance group success\n")
```

## 实例组列表
使用以下代码可以获取实例组列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.ListInstanceGroupArgs{
    Manner: "page",
}
result, err := client.ListInstanceGroup(args)
if err != nil {
    fmt.Printf("get instance group list error: %+v\n", err)
    return
}
fmt.Printf("get instance group list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 实例组详情
使用以下代码可以获取实例组详情

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.InstanceGroupDetail(groupId)
if err != nil {
    fmt.Printf("get instance group detail error: %+v\n", err)
    return
}
fmt.Printf("get instance group detail success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 实例组前置检查-GTID检查
使用以下代码可以进行GTID检查

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.CheckGtidArgs{
    InstanceId: instanceId,
}
result, err := client.InstanceGroupCheckGtid(args)
if err != nil {
    fmt.Printf("GTID check error: %+v\n", err)
    return
}
fmt.Printf("GTID check success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 实例组前置检查-连通性检查
使用以下代码可以进行连通性检查

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.CheckPingArgs{
    SourceId: instanceId,
    TargetId: instanceId,
}
result, err := client.InstanceGroupCheckPing(args)
if err != nil {
    fmt.Printf("check ping error: %+v\n", err)
    return
}
fmt.Printf("check ping success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 实例组前置检查-数据检查
使用以下代码可以进行数据检查

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.CheckDataArgs{
    InstanceId: instanceId,
}
result, err := client.InstanceGroupCheckData(args)
if err != nil {
    fmt.Printf("check data error: %+v\n", err)
    return
}
fmt.Printf("check data success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 小版本前置检查
使用以下代码可以进行小版本前置检查

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.CheckVersionArgs{
    LeaderId:   instanceId,
    FollowerId: instanceId,
}
result, err := client.InstanceGroupCheckVersion(args)
if err != nil {
    fmt.Printf("check data error: %+v\n", err)
    return
}
fmt.Printf("check data success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 修改热活实例组的名称
使用以下代码可以修改热活实例组的名称

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.InstanceGroupNameArgs{
    Name: "test_group_name",
}
err := client.UpdateInstanceGroupName(groupId, args)
if err != nil {
    fmt.Printf("update instance group name error: %+v\n", err)
    return
}
fmt.Printf("update instance group name success\n")
```

## 加入热活实例组
使用以下代码可以加入热活实例组

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.InstanceGroupAddArgs{
    FollowerId: instanceId,
}
err := client.InstanceGroupAdd(groupId, args)
if err != nil {
    fmt.Printf("add instance group error: %+v\n", err)
    return
}
fmt.Printf("add instance group success\n")
```

## 批量加入热活实例组
使用以下代码可以批量加入热活实例组

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.InstanceGroupBatchAddArgs{
    FollowerIds: []string{instanceId},
    Name:        "test_group_name",
    LeaderId:    instanceId,
}
err := client.InstanceGroupBatchAdd(args)
if err != nil {
    fmt.Printf("batch add instance group error: %+v\n", err)
    return
}
fmt.Printf("batch add instance group success\n")
```
## 强制切换热活实例组
使用以下代码可以强制切换热活实例组

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.ForceChangeArgs{
    LeaderId: instanceId,
    Force:    0,
}
err := client.InstanceGroupForceChange(groupId, args)
if err != nil {
    fmt.Printf("instance group force change error: %+v\n", err)
    return
}
fmt.Printf("instance group force change success\n")
```
## 主角色变更
使用以下代码可以主角色变更

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GroupLeaderChangeArgs{
    LeaderId: instanceId,
}
err := client.InstanceGroupLeaderChange(groupId, args)
if err != nil {
    fmt.Printf("instance group leader change error: %+v\n", err)
    return
}
fmt.Printf("instance group leader change success\n")
```

## 退出热活实例组
使用以下代码可以退出热活实例组

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
err := client.InstanceGroupRemove(groupId, instanceId)
if err != nil {
    fmt.Printf("instance group remove error: %+v\n", err)
    return
}
fmt.Printf("instance group remove success\n")
```

## 删除热活实例组
使用以下代码可以删除热活实例组

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
err := client.DeleteInstanceGroup(groupId)
if err != nil {
    fmt.Printf("delete instance group error: %+v\n", err)
    return
}
fmt.Printf("delete instance group success\n")
```

# 版本管理
## 查看实例允许升级的小版本列表

使用以下代码可以查看实例允许升级的小版本列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.InstanceMinorVersionList(instanceId)
if err != nil {
    fmt.Printf("get instance minor version list error: %+v\n", err)
    return
}
fmt.Printf("get instance minor version list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))

```
## 实例升级小版本

使用以下代码可以升级小版本

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.UpgradeMinorVersionArgs{
    TargetMinorVersion: "5.7.38",
    EffectiveTime:      "immediate",
}
err := client.InstanceUpgradeMinorVersion(instanceId, args)
if err != nil {
    fmt.Printf("update instance minor version list error: %+v\n", err)
    return
}
fmt.Printf("update instance minor version list success\n")

```
# SmartDBA
## 查询慢SQL诊断开通状态
使用以下代码可以查询慢SQL诊断开通状态

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err := client.SlowSqlFlowStatus(instanceId)
if err != nil {
    fmt.Printf("get slow sql flow status error: %+v\n", err)
    return
}
fmt.Printf("get slow sql flow status success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 开通慢SQL诊断

使用以下代码可以开通慢SQL诊断

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
err := client.EnableSlowSqlFlow(instanceId)
if err != nil {
    fmt.Printf("enable slow sql flow error: %+v\n", err)
    return
}
fmt.Printf("enable slow sql flow success\n")

```
## 关闭慢SQL诊断

使用以下代码可以关闭慢SQL诊断

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
err := client.DisableSlowSqlFlow(instanceId)
if err != nil {
    fmt.Printf("disable slow sql flow error: %+v\n", err)
    return
}
fmt.Printf("disable slow sql flow success\n")

```

## 获取慢SQL诊断列表

使用以下代码可以获取慢SQL诊断列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &GetSlowSqlArgs{}
result, err := client.GetSlowSqlList(instanceId, args)
if err != nil {
    fmt.Printf("get slow sql flow list error: %+v\n", err)
    return
}
fmt.Printf("get slow sql flow list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 根据SQLID获取慢SQL

使用以下代码可以根据SQLID获取慢SQL

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSlowSqlBySqlId(instanceId, sqlId)
if err != nil {
    fmt.Printf("get slow sql detail by sqlid error: %+v\n", err)
    return
}
fmt.Printf("get slow sql detail by sqlid success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取慢SQL说明

使用以下代码可以获取慢SQL说明

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSlowSqlExplain(instanceId, sqlId, db)
if err != nil {
    fmt.Printf("get slow sql explain error: %+v\n", err)
    return
}
fmt.Printf("get slow sql explain success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 获取SQL模板维度的统计信息

使用以下代码可以获取SQL模板维度的统计信息

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GetSlowSqlArgs{}
result, err := client.GetSlowSqlStatsDigest(instanceId, args)
if err != nil {
    fmt.Printf("get slow sql stats digest error: %+v\n", err)
    return
}
fmt.Printf("get slow sql stats digest success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 获取慢SQL耗时分布

使用以下代码可以获取慢SQL耗时分布

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GetSlowSqlDurationArgs{}
result, err := client.GetSlowSqlDuration(instanceId, args)
if err != nil {
    fmt.Printf("get slow sql duration error: %+v\n", err)
    return
}
fmt.Printf("get slow sql duration success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取慢SQL来源IP分布

使用以下代码可以获取慢SQL来源IP分布

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GetSlowSqlSourceArgs{}
result, err := client.GetSlowSqlSource(instanceId, args)
if err != nil {
    fmt.Printf("get slow sql source error: %+v\n", err)
    return
}
fmt.Printf("get slow sql source success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取慢SQL中的表

使用以下代码可以获取慢SQL中的表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSlowSqlSchema(instanceId, sqlId, db)
if err != nil {
    fmt.Printf("get slow sql schema error: %+v\n", err)
    return
}
fmt.Printf("get slow sql schema success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取慢SQL中的列

使用以下代码可以获取慢SQL中的列

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSlowSqlTable(instanceId, sqlId, db, table)
if err != nil {
    fmt.Printf("get slow sql table error: %+v\n", err)
    return
}
fmt.Printf("get slow sql table success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 获取慢SQL表中的索引

使用以下代码可以获取慢SQL表中的索引

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GetSlowSqlIndexArgs{
    SqlId:  "e9fa9802-0d0e-41b4-b3ba-6496466b6cad",
    Schema: "db1",
    Table:  "table1",
}
result, err := client.GetSlowSqlIndex(instanceId, args)
if err != nil {
    fmt.Printf("get slow sql index error: %+v\n", err)
    return
}
fmt.Printf("get slow sql index success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取慢SQL趋势

使用以下代码可以获取慢SQL趋势

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GetSlowSqlTrendArgs{
    Start: "2023-05-05T05:30:13.000Z",
    End:   "2023-05-06T05:30:13.000Z",
}
result, err := client.GetSlowSqlTrend(instanceId, args)
if err != nil {
    fmt.Printf("get slow sql trend error: %+v\n", err)
    return
}
fmt.Printf("get slow sql trend success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取慢SQL调优建议

使用以下代码可以获取慢SQL调优建议

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSlowSqlAdvice(instanceId, sqlId, db)
if err != nil {
    fmt.Printf("get slow sql advice error: %+v\n", err)
    return
}
fmt.Printf("get slow sql advice success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取库表空间概况

使用以下代码可以获取库表空间概况

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetDiskInfo(instanceId)
if err != nil {
    fmt.Printf("get disk info error: %+v\n", err)
    return
}
fmt.Printf("get disk info success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取数据空间的数据库列表

使用以下代码可以获取数据空间的数据库列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetDbListSize(instanceId)
if err != nil {
    fmt.Printf("get db list size info error: %+v\n", err)
    return
}
fmt.Printf("get db list size info success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取数据空间表的详情

使用以下代码可以获取数据空间表的详情

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GetTableListArgs{
    DbName: "db1",
}
result, err := client.GetTableListInfo(instanceId, args)
if err != nil {
    fmt.Printf("get table list error: %+v\n", err)
    return
}
fmt.Printf("get table list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取指定会话kill类型的相关参数

使用以下代码可以获取指定会话kill类型的相关参数

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetKillSessionTypes(instanceId)
if err != nil {
    fmt.Printf("get kill session types error: %+v\n", err)
    return
}
fmt.Printf("get kill session types success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取指定实例的会话概览

使用以下代码可以获取指定实例的会话概览

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSessionSummary(instanceId)
if err != nil {
    fmt.Printf("get kill session summary error: %+v\n", err)
    return
}
fmt.Printf("get kill session summary success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取指定实例的实时会话

使用以下代码可以获取指定实例的实时会话

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.SessionDetailArgs{}
result, err := client.GetSessionDetail(instanceId, args)
if err != nil {
    fmt.Printf("get session detail error: %+v\n", err)
    return
}
fmt.Printf("get session detail success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 校验执行kill操作的数据库用户及密码是否正确

使用以下代码可以校验执行kill操作的数据库用户及密码是否正确

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.KillSessionAuthArgs{}
result, err := client.CheckKillSessionAuth(instanceId, args)
if err != nil {
    fmt.Printf("check kill session auth error: %+v\n", err)
    return
}
fmt.Printf("check kill session auth success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取指定实例的会话kill记录

使用以下代码可以获取指定实例的会话kill记录

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.KillSessionHistory{}
result, err := client.GetKillSessionHistory(instanceId, args)
if err != nil {
    fmt.Printf("get kill session history error: %+v\n", err)
    return
}
fmt.Printf("get kill session history success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 根据传入的kill类型及类型所对应的值执行kill会话的操作

使用以下代码可以根据传入的kill类型及类型所对应的值执行kill会话的操作

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.KillSessionArgs{}
result, err := client.KillSession(instanceId, args)
if err != nil {
    fmt.Printf("kill session error: %+v\n", err)
    return
}
fmt.Printf("kill session success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取指定实例的会话统计

使用以下代码可以获取指定实例的会话统计

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSessionStatistics(instanceId)
if err != nil {
    fmt.Printf("get session statistics error: %+v\n", err)
    return
}
fmt.Printf("get session statistics success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 查询错误日志服务是否开启

使用以下代码可以查询错误日志服务是否开启

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetErrorLogStatus(instanceId)
if err != nil {
    fmt.Printf("get error log status error: %+v\n", err)
    return
}
fmt.Printf("get error log status success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 开启错误日志服务

使用以下代码可以开启错误日志服务

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.EnableErrorLog(instanceId)
if err != nil {
    fmt.Printf("enable error log status error: %+v\n", err)
    return
}
fmt.Printf("enable error log status success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 关闭错误日志服务

使用以下代码可以关闭错误日志服务

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.DisableErrorLog(instanceId)
if err != nil {
    fmt.Printf("disable error log status error: %+v\n", err)
    return
}
fmt.Printf("disable error log status success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取错误日志列表

使用以下代码可以获取错误日志列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.ErrorLogListArgs{}
result, err := client.GetErrorLogList(instanceId, args)
if err != nil {
    fmt.Printf("get error log list error: %+v\n", err)
    return
}
fmt.Printf("get error log list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 获取实例限流规则列表

使用以下代码可以获取实例限流规则列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSqlFilterList(instanceId)
if err != nil {
    fmt.Printf("get sql filter list error: %+v\n", err)
    return
}
fmt.Printf("get sql filter list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 获取某个限流规则详情

使用以下代码可以获取某个限流规则详情

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSqlFilterDetail(instanceId, filterId)
if err != nil {
    fmt.Printf("get sql filter detail error: %+v\n", err)
    return
}
fmt.Printf("get sql filter detail success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 添加一条限流规则

使用以下代码可以添加一条限流规则

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.SqlFilterArgs{
    FilterType:  "SELECT",
    FilterKey:   "123",
    FilterLimit: 0,
}
result, err := client.AddSqlFilter(instanceId, args)
if err != nil {
    fmt.Printf("add sql filter error: %+v\n", err)
    return
}
fmt.Printf("add sql filter success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 更新一条限流规则

使用以下代码可以更新一条限流规则

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.SqlFilterArgs{
    FilterType:  "SELECT",
    FilterKey:   "123",
    FilterLimit: 0,
}
result, err := client.UpdateSqlFilter(instanceId, args)
if err != nil {
    fmt.Printf("update sql filter error: %+v\n", err)
    return
}
fmt.Printf("update sql filter success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 开启关闭某个限流规则

使用以下代码可以开启关闭某个限流规则

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.StartOrStopSqlFilterArgs{
    Action: "OFF",
}
result, err := client.StartOrStopSqlFilter(instanceId, filterId, args)
if err != nil {
    fmt.Printf("start or stop sql filter error: %+v\n", err)
    return
}
fmt.Printf("start or stop sql filter success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 删除某个限流规则

使用以下代码可以删除某个限流规则

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err := client.DeleteSqlFilter(instanceId, filterId)
if err != nil {
    fmt.Printf("delete sql filter error: %+v\n", err)
    return
}
fmt.Printf("delete sql filter success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 实例是否支持限流

使用以下代码可以实例是否支持限流

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"

result, err := client.IsAllowedSqlFilter(instanceId)
if err != nil {
    fmt.Printf("is allowed sql filter error: %+v\n", err)
    return
}
fmt.Printf("is allowed sql filter success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

# Performance
## Kill会话

使用以下代码可以Kill会话

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &ProcessArgs{
    Ids: []int64{123},
}

err := client.ProcessKill(instanceId, args)
if err != nil {
    fmt.Printf("process kill error: %+v\n", err)
    return
}
fmt.Printf("process kill success\n")
```
## 查询innodbstatus快照数据

使用以下代码可以查询innodbstatus快照数据

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.InnodbStatus(instanceId)
if err != nil {
    fmt.Printf("get innodb status error: %+v\n", err)
    return
}
fmt.Printf("get innodb status success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```

## 查询processlist快照数据

使用以下代码可以查询processlist快照数据

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.ProcessList(instanceId)
if err != nil {
    fmt.Printf("get process list error: %+v\n", err)
    return
}
fmt.Printf("get process list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 查询事务列表

使用以下代码可以查询事务列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.TransactionList(instanceId)
if err != nil {
    fmt.Printf("get transaction list error: %+v\n", err)
    return
}
fmt.Printf("get transaction list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
## 查询连接列表

使用以下代码可以查询连接列表

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.ConnectionList(instanceId)
if err != nil {
    fmt.Printf("get ConnectionList list error: %+v\n", err)
    return
}
fmt.Printf("get ConnectionList list success\n")
jsonData, _ := json.Marshal(result)
fmt.Println(string(jsonData))
```
# 备份管理

## 获取备份列表

使用以下代码可以获取一个实例下的备份列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.GetBackupListArgs{}
_, err := client.GetBackupList(instanceId, args)
if err != nil {
    fmt.Printf("get backup list error: %+v\n", err)
    return
}
fmt.Printf("get backup list success\n")
```

> 注意:
>
> - 请求参数 marker 和 maxKeys 不是必须的。

## 获取备份详情

使用以下代码可以获取一个实例备份的详情信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetBackupDetail(instanceId, backupId)
if err != nil {
    fmt.Printf("get backup detail error: %+v\n", err)
    return
}
fmt.Printf("get backup detail success\n")
```

## 删除备份

使用以下代码可以删除手动备份
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
_, err := client.DeleteBackup(instanceId, backupId)
if err != nil {
    fmt.Printf("delete backup detail error: %+v\n", err)
    return
}
fmt.Printf("delete backup detail success\n")
```
## 更新备份策略

使用以下代码可以更新一个实例的备份策略。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.ModifyBackupPolicyArgs{
    BackupDays: "1,3",
    BackupTime: "10:00:00Z",
    Persistent: true,
    ExpireInDays: 20,
}
err := client.ModifyBackupPolicy(instanceId, args)
if err != nil {
    fmt.Printf("modify backup policy error: %+v\n", err)
    return
}
fmt.Printf("modify backup policy success\n")
```
## 获取binlog列表

使用以下代码可以获取binlog列表
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetBinlogList(instanceId, detaTime)
if err != nil {
    fmt.Printf("get binlog list error: %+v\n", err)
    return
}
fmt.Printf("get binlog list success\n")
```

## 获取binlog信息

使用以下代码可以获取binlog信息
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetBinlogInfo(instanceId, binlogId, downloadValidTimeInSec)
if err != nil {
    fmt.Printf("get binlog detail error: %+v\n", err)
    return
}
fmt.Printf("get binlog detail success\n")
```

## 按时间点进行库表恢复

使用以下代码可以按时间点进行库表恢复
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
recoveryByDatetimeArgs := &RecoveryByDatetimeArgs{
    Datetime: "2022-01-11T16:05:52Z",
    Data: []RecoveryData{
        {
            DbName:      "test_db",
            NewDbname:   "new_test_db",
            RestoreMode: "database",
            Tables: []TableData{
                {
                    TableName:    "table_name",
                    NewTablename: "new_table_name",
                },
            },
        },
    },
}
err := client.RecoveryToSourceInstanceByDatetime(instanceId, recoveryByDatetimeArgs)
if err != nil {
    fmt.Printf("recovery by datetime error: %+v\n", err)
    return
}
fmt.Printf("recovery by datetime success\n")
```
## 按备份集进行库表恢复

使用以下代码可以按备份集进行库表恢复
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
recoveryBySnapshotArgs := &RecoveryBySnapshotArgs{
    SnapshotId: "1691734023130272802",
    Data: []RecoveryData{
        {
            DbName:      "test_db",
            NewDbname:   "new_test_db",
            RestoreMode: "database",
            Tables: []TableData{
                {
                    TableName:    "table_name",
                    NewTablename: "new_table_name",
                },
            },
        },
    },
}
err := client.RecoveryToSourceInstanceBySnapshot(instanceId, recoveryBySnapshotArgs)
if err != nil {
    fmt.Printf("recovery by snapshot error: %+v\n", err)
    return
}
fmt.Printf("recovery by snapshot success\n")
```
# 慢日志下载任务

## 慢日志下载任务列表
``` go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSlowLogDownloadTaskList(instanceId, datetime)
if err != nil {
    fmt.Printf("get slowlog download task list error: %+v\n", err)
    return
}
fmt.Printf("get slowlog download task list success\n")
fmt.Printf(result)
```
## 慢日志下载详情 
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSlowLogDownloadDetail(instanceId, logId, downloadValidTimeInSec)
if err != nil {
    fmt.Printf("get slowlog download detail error: %+v\n", err)
    return
}
fmt.Printf("get slowlog download detail success\n")
fmt.Printf(result)
```
# 其它

## 获取可用区列表

使用以下代码可以获取可用区列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
err = client.GetZoneList()
if err != nil {
    fmt.Printf("get zone list error: %+v\n", err)
    return
}
fmt.Printf("get zone list success\n")
fmt.Println("rds instanceId: ", result.InstanceId)
```

## 获取子网列表

使用以下代码可以获取一个实例下的子网列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
args := &rds.ListSubnetsArgs{}
_, err := client.ListSubnets(args)
if err != nil {
    fmt.Printf("get subnet list error: %+v\n", err)
    return
}
fmt.Printf("get subnet list success\n")
```

> 注意:
>
> - 请求参数 vpcId 和 zoneName 不是必须的。

## 查看白名单

使用以下代码可以获取一个实例下的白名单列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSecurityIps(instanceId)
if err != nil {
    fmt.Printf("get securityIp list error: %+v\n", err)
    return
}
data, _ := json.Marshal(result)
fmt.Println(string(data))
fmt.Println(result.Etag)
fmt.Printf("get securityIp list success\n")
```

> 注意:
>
> - 在更新白名单时需要通过该接口获取最新的Etag。

## 更新白名单

使用以下代码可以更新一个实例下的白名单列表。

```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
result, err := client.GetSecurityIps(instanceId)
if err != nil {
    fmt.Printf("get securityIp list error: %+v\n", err)
    return
}
fmt.Println(result.Etag)
fmt.Printf("get securityIp list success\n")

args := &rds.UpdateSecurityIpsArgs{
				SecurityIps:  []string{
					"%",
					"192.0.0.1",
					"192.0.0.2",
				},
			}
er := client.UpdateSecurityIps(instanceId, result.Etag, args)
if er != nil {
    fmt.Printf("update securityIp list error: %+v\n", er)
    return
}
fmt.Printf("update securityIp list success\n")
```

> 注意:
>
> - 在更新白名单时需要通过查看白名单接口获取最新的Etag。
> - 白名单需要全量更新，每次更新需要把全部白名单列表都添加上。

# 错误处理

GO语言以error类型标识错误，RDS支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | RDS服务返回的错误

用户使用SDK调用RDS相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// rdsClient 为已创建的RDS Client对象
result, err := client.ListRds()
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

客户端异常表示客户端尝试向RDS发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当RDS服务端出现异常时，RDS服务端会返回给用户相应的错误信息，以便定位问题。

## SDK日志

RDS GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

RDS GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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
log.Debugf("%s", "logging message using the log package in the RDS go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the RDS go sdk")
```



首次发布:

 - 支持创建RDS主实例、创建RDS只读实例、创建RDS代理实例、查询RDS列表、查询指定RDS实例信息、删除RDS实例、RDS实例扩缩容、创建账号、查询账号列表、查询特定账号信息、删除特定账号信息。