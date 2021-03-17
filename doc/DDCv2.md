# DDC-RDS服务

# 概述

本文档主要介绍DDC-RDS GO SDK的使用。在使用本文档前，您需要先了解DDC和RDS的一些基本知识，并已开通了DDC服务和RDS服务。

# 初始化

## 确认Endpoint


目前支持“华东-苏州”区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
SU | ddc.su.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云DDC-RDS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问DDC-RDS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建DDC-RDS Client

DDC-RDS Client是同时支持DDC和RDS服务的客户端，为开发者与DDC和RDS服务进行交互提供了一系列的方法。

### 使用AK/SK新建DDC-RDS Client

通过AK/SK方式访问DDC-RDS，用户可以参考如下代码新建一个DDC-RDS Client：

```go
import (
	ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个DDC-RDSClient
	client, err := ddcrds.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。
第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名`ddc.su.baidubce.com`作为DDC-RDS的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为苏州，则为`ddc.su.baidubce.com`。

### 使用STS创建DDC-RDS Client

**申请STS token**

DDC-RDS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问DDC-RDS，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建DDC-RDS Client**

申请好STS后，可将STS Token配置到DDC-RDS Client中，从而实现通过STS Token创建DDC-RDS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建DDC-RDS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/sts" //导入STS服务模块
    ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2" //导入DDC-RDS服务模块
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

	// 使用申请的临时STS创建DDC-RDS服务的Client对象，Endpoint使用默认值
	client, err := ddcrds.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "ddc.su.baidubce.com")
	if err != nil {
		fmt.Println("create v2 client failed:", err)
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
	client.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置DDC-RDS Client时，无论对应DDC-RDS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问DDC-RDS

DDC-RDS支持HTTPS传输协议，您可以通过在创建DDC-RDS Client对象时指定的Endpoint中指明HTTPS的方式，在DDC-RDS GO SDK中使用HTTPS访问DDC-RDS服务：

```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

ENDPOINT := "https://ddc.su.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := ddcrds.NewClient(AK, SK, ENDPOINT)
```

## 配置DDC-RDS Client

如果用户需要配置DDC-RDS Client的一些细节的参数，可以在创建DDC-RDS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。
### 切换地域

下面一段代码可以让客户端使用切换地域:
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

//创建DDC-RDS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ddc.su.baidubce.com"
client, _ := ddcrds.NewClient(AK, SK, ENDPOINT)

// 切换到北京地域
client.ConfigRegion("BJ")
```


### 使用代理

下面一段代码可以让客户端使用代理访问DDC-RDS服务：

```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

//创建DDC-RDS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ddc.su.baidubce.com"
client, _ := ddcrds.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.ConfigProxyUrl("127.0.0.1:8080")
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ddc.su.baidubce.com"
client, _ := ddcrds.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.ConfigRetry(bce.NewNoRetryPolicy())

// 配置连接超时时间为30秒
client.ConfigConnectionTimeoutInMillis(30 * 1000)
```

### 配置生成签名字符串选项

```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ddc.su.baidubce.com"
client, _ := ddcrds.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.ConfigSignOptionHeadersToSign(headersToSign)

// 配置签名的有效期为30秒
client.ConfigSignOptionExpireSeconds(30)
```

**参数说明**

用户使用GO SDK访问DDC-RDS时，创建的DDC-RDS Client对象的`Config`字段支持的所有参数如下表所示,可以通过
`Config{Field}()`等方法设置配置项,例如:`ConfigRegion("BJ")`

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建DDC-RDS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# DDC-RDS管理
DDC-RDS服务可以管理以下两种服务
- 云数据库专属集群 DDC （Dedicated Database Cluster）是专业、高性能、高可靠的云数据库服务。云数据库 DDC 提供 Web 界面进行配置、操作数据库实例，还为您提供可靠的数据备份和恢复、完备的安全管理、完善的监控、轻松扩展等功能支持。相对于自建数据库，云数据库 DDC 具有更经济、更专业、更高效、更可靠、简单易用等特点，使您能更专注于核心业务。
- 云数据库 RDS （Relational Database Service）是专业、高性能、高可靠的云数据库服务。云数据库 RDS 提供 Web 界面进行配置、操作数据库实例，还为您提供可靠的数据备份和恢复、完备的安全管理、完善的监控、轻松扩展等功能支持。相对于自建数据库，云数据库 RDS 具有更经济、更专业、更高效、更可靠、简单易用等特点，使您能更专注于核心业务。

# 资源池管理

## 查询资源池列表
使用以下代码可以查询当前用户的资源池列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

result, err := client.ListPool(nil, "ddc")
if err != nil {
    fmt.Printf("list pool error: %+v\n", err)
    return
}

for i := range result.Result {
    pool := result.Result[i]
    fmt.Println("ddc pool id: ", pool.PoolID)
    fmt.Println("ddc pool vpc id: ", pool.VpcID)
    fmt.Println("ddc pool engine: ", pool.Engine)
    fmt.Println("ddc pool create time: ", pool.CreateTime)
    fmt.Println("ddc pool hosts: ", pool.Hosts)
}
```

# 部署集管理

## 创建部署集
使用以下代码可以在指定资源池下创建一个新的部署集(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
args := &ddcrds.CreateDeployRequest{
    // 幂等 Token
    ClientToken: "xxxyyyzzz",
    // 部署集名称
    DeployName: "api-from-go",
    // 部署策略 支持集中部署(centralized)/完全打散(distributed)
    Strategy:   "distributed",
    // 亲和度阈值 取值范围【0-96】
    CentralizeThreshold: 23,
}
deploy,err := client.CreateDeploySet(poolId, args)
if err != nil {
    fmt.Printf("create deploy set error: %+v\n", err)
    return
}

fmt.Println("create deploy set success.")
fmt.Println("returned deploy id: ", deploy.DeployID)
```

## 查询部署集列表
使用以下代码可以查询指定资源池下的部署集列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

result, err := client.ListDeploySets(poolId, nil)
if err != nil {
    fmt.Printf("list deploy set error: %+v\n", err)
    return
}

for i := range result.Result {
    deploy := result.Result[i]
	fmt.Println("ddc deploy id: ", deploy.DeployID)
    fmt.Println("ddc deploy name: ", deploy.DeployName)
    fmt.Println("ddc deploy strategy: ", deploy.Strategy)
    fmt.Println("ddc deploy create time: ", deploy.CreateTime)
    fmt.Println("ddc deploy centralizeThreshold: ", deploy.CentralizeThreshold)
    fmt.Println("ddc instance ids: ", deploy.Instances)
}
```

## 查询特定部署集信息
使用以下代码可以查询特定部署集的详细信息(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
deploy, err := client.GetDeploySet(poolId, deployId)
if err != nil {
    fmt.Printf("get deploy set error: %+v\n", err)
    return
}

// 获取部署集的详细信息
fmt.Println("ddc deploy id: ", deploy.DeployID)
fmt.Println("ddc deploy name: ", deploy.DeployName)
fmt.Println("ddc deploy strategy: ", deploy.Strategy)
fmt.Println("ddc deploy create time: ", deploy.CreateTime)
fmt.Println("ddc deploy centralizeThreshold: ", deploy.CentralizeThreshold)
fmt.Println("ddc instance ids: ", deploy.Instances)
```
## 更新部署集
使用以下代码更新部署集的部署策略和亲和度(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
args := &ddcrds.UpdateDeployRequest{
    // 幂等 Token
    ClientToken: "xxxyyyzzz",
    // 部署策略 支持集中部署(centralized)/完全打散(distributed)
    Strategy:   "distributed",
    // 亲和度阈值 取值范围【0-96】，必须大于原亲和度阈值
	CentralizeThreshold: 30,
}
err := client.UpdateDeploySet(poolId, deployId, args)
if err != nil {
    fmt.Printf("update deploy set error: %+v\n", err)
    return
}

fmt.Println("update deploy set success.")
```

## 删除部署集
使用以下代码可以删除某个资源池下特定的部署集(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

err := DDC_CLIENT.DeleteDeploySet(poolId, deployId)
if err != nil {
    fmt.Printf("delete deploy set error: %+v\n", err)
    return
}
fmt.Printf("delete deploy set success\n")
```

# 实例管理

## 创建实例
使用以下代码可以创建主实例。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
args := &ddcrds.CreateRdsArgs{
    // 指定ddc的数据库引擎，取值mysql,必选
    Engine:            "mysql",
    // 指定ddc的数据库版本，必选
    EngineVersion:  "5.6",
    // 计费相关参数，PaymentTiming取值为 预付费：Prepaid，后付费：Postpaid；Reservation：支付方式为后支付时不需要设置，预支付时必须设置；必选
    Billing: ddcrds.Billing{
        PaymentTiming: "Postpaid",
        //Reservation: ddcrds.Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
    },
    // 预付费时可指定自动续费参数 AutoRenewTime 和 AutoRenewTimeUnit
    // 自动续费时长（续费单位为year 不大于3，续费单位为mouth 不大于9）
    // AutoRenewTime: 1,
    // 自动续费单位（"year";"mouth"）
    // AutoRenewTimeUnit: "year",
    // CPU核数，必选
    CpuCount: 1,
    //套餐内存大小，单位GB，必选
    MemoryCapacity: 1,
    //套餐磁盘大小，单位GB，每5G递增，必选
    VolumeCapacity: 5,
    //批量创建云数据库 ddc 实例个数, 最大不超过10，默认1，可选
    PurchaseCount: 1,
    //ddc实例名称，允许小写字母、数字，长度限制为1~32，默认命名规则:{engine} + {engineVersion}，可选
    InstanceName: "instanceName",
    //所属系列，Singleton:单机版，Basic：单机基础版，Standard：双机高可用版。仅SQLServer 2012sp3 支持单机基础版。默认Standard，可选
    Category: "Standard",
    //指定zone信息，默认为空，由系统自动选择，可选
    //zoneName命名规范是小写的“国家-region-可用区序列"，例如北京可用区A为"cn-bj-a"。
    ZoneNames:[]string{"cn-bj-a"},
    //vpc ShortId 如果不提供则属于默认vpc，可选
    VpcId: "vpc-IyrqYIQ7",
    //是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
    IsDirectPay: false,
    //vpc内，每个可用区的subnetId；如果不是默认vpc则必须指定 subnetId，可选
    Subnets: []ddcrds.SubnetMap{
    {
        ZoneName: "cn-bj-a",
        // 指定子网的ShortId
        SubnetId: "subnet-short-id",
    },
    },
    // 实例绑定的标签信息，可选
    Tags: []ddcrds.TagModel{
        {
            TagKey:   "tagK",
            TagValue: "tagV",
        },
    },
    // 部署集id 产品类型为DDC时必选,产品类型为RDS忽略该参数
    DeployId:"xxxyyy-123",
    // 资源池id 产品类型为DDC时必选,产品类型为RDS忽略该参数
    PoolId:"xxxyzzzyy-123",
}
// 创建DDC数据库专属集群产品，需要传入产品类型参数ddc
result, err := client.CreateRds(args,"ddc")
if err != nil {
    fmt.Printf("create ddc error: %+v\n", err)
    return
}

for _, e := range result.InstanceIds {
    fmt.Println("create ddc success, instanceId: ", e)
}

// RDS
// 创建RDS数据库，需要传入产品类型参数rds
result, err = client.CreateRds(args,"rds")
if err != nil {
    fmt.Printf("create ddc error: %+v\n", err)
    return
}

for _, e := range result.InstanceIds {
    fmt.Println("create ddc success, instanceId: ", e)
}
```

## 创建只读实例
使用以下代码可以创建只读实例。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// DDC
args := &ddcrds.CreateReadReplicaArgs{
    //主实例ID，必选
    SourceInstanceId: "sourceInstanceId",
    // 计费相关参数，只读实例只支持后付费Postpaid，必选
    Billing: ddcrds.Billing{
        PaymentTiming: "Postpaid",
    },
    // CPU核数，必选
    CpuCount: 1,
    //套餐内存大小，单位GB，必选
    MemoryCapacity: 1,
    //套餐磁盘大小，单位GB，每5G递增，必选
    VolumeCapacity: 5,
    //批量创建云数据库 只读实例个数, 目前rds产品只支持一次创建一个,必选
    PurchaseCount: 1,
    //实例名称，允许小写字母、数字，长度限制为1~32，默认命名规则:{engine} + {engineVersion}，可选
    InstanceName: "instanceName",
    //指定zone信息，默认为空，由系统自动选择，可选
    //zoneName命名规范是小写的“国家-region-可用区序列"，例如北京可用区A为"cn-bj-a"。
    ZoneNames: []string{"cn-bj-a"},
    //与主实例 vpcId 相同，可选
    VpcId: "vpc-IyrqYIQ7",
    //是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
    IsDirectPay: false,
    //vpc内，每个可用区的subnetId；如果不是默认vpc则必须指定 subnetId，可选
    Subnets: []ddcrds.SubnetMap{
    {
        ZoneName: "cn-bj-a",
        SubnetId: "sbn-IyWRnII7",
    },
    },
    // 实例绑定的标签信息，可选
    Tags: []ddcrds.TagModel{
    {
        TagKey:   "tagK",
        TagValue: "tagV",
    },
    },
    // 部署集id 产品类型为DDC时必选,产品类型为RDS忽略该参数
    DeployId:"xxxyyy-123",
    // 资源池id 产品类型为DDC时必选,产品类型为RDS忽略该参数 与主实例保持一致
    PoolId:"xxxyzzzyy-123",
    // RO组ID。(创建只读实例时) 可选
    // 如果不传，默认会创建一个RO组，并将该只读加入RO组中
    RoGroupId:"yyzzcc",
    // RO组是否启用延迟剔除，默认为0。可选值0和1,0代表不开启,1代表开启（创建只读实例时）可选
    EnableDelayOff:"1",
    // 延迟阈值，单位为秒。（创建只读实例时）可选
    DelayThreshold: "1",
    // RO组最少保留实例数目。默认为1. （创建只读实例时）可选
    LeastInstanceAmount: "1",
    // 只读实例在RO组中的读流量权重，取值范围为[0-100]。默认为50（创建只读实例时）可选
    RoGroupWeight: "50",
}
result, err := client.CreateReadReplica(args)
if err != nil {
    fmt.Printf("create ddc readReplica error: %+v\n", err)
    return
}

for _, e := range result.InstanceIds {
    fmt.Println("create ddc readReplica success, instanceId: ", e)
}

// RDS
result, err = client.CreateReadReplica(args)
if err != nil {
    fmt.Printf("create rds readReplica error: %+v\n", err)
    return
}

for _, e := range result.InstanceIds {
    fmt.Println("create rds readReplica success, instanceId: ", e)
}
```

## 实例详情

使用以下代码可以查询指定实例的详情。
```go
package main
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

result, err := client.GetDetail(instanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
// 获取实例详情信息
fmt.Println("ddc instanceId: ", result.InstanceId)
fmt.Println("ddc instanceName: ", result.InstanceName)
fmt.Println("ddc engine: ", result.Engine)
fmt.Println("ddc engineVersion: ", result.EngineVersion)
fmt.Println("ddc instanceStatus: ", result.InstanceStatus)
fmt.Println("ddc cpuCount: ", result.CpuCount)
fmt.Println("ddc memoryCapacity: ", result.MemoryCapacity)
fmt.Println("ddc volumeCapacity: ", result.VolumeCapacity)
fmt.Println("ddc usedStorage: ", result.UsedStorage)
fmt.Println("ddc paymentTiming: ", result.PaymentTiming)
fmt.Println("ddc instanceType: ", result.InstanceType)
fmt.Println("ddc instanceCreateTime: ", result.InstanceCreateTime)
fmt.Println("ddc instanceExpireTime: ", result.InstanceExpireTime)
fmt.Println("ddc publicAccessStatus: ", result.PublicAccessStatus)
fmt.Println("ddc vpcId: ", result.VpcId)
fmt.Println("ddc Subnets: ", result.Subnets)
fmt.Println("ddc BackupPolicy: ", result.BackupPolicy)
fmt.Println("ddc RoGroupList: ", result.RoGroupList)
fmt.Println("ddc NodeMaster: ", result.NodeMaster)
fmt.Println("ddc NodeSlave: ", result.NodeSlave)
fmt.Println("ddc NodeReadReplica: ", result.NodeReadReplica)
fmt.Println("ddc DeployId: ", result.DeployId)
fmt.Println("ddc SyncMode: ", result.SyncMode)
fmt.Println("ddc Category: ", result.Category)
fmt.Println("ddc ZoneNames: ", result.ZoneNames)
fmt.Println("ddc Endpoint: ", result.Endpoint)
fmt.Println("ddc vnetIp: ", result.Endpoint.VnetIp)
fmt.Println("ddc vnetIpBackup: ", result.Endpoint.VnetIpBackup)

// RDS
result, err := client.GetDetail(instanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
// 获取实例详情信息
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
fmt.Println("rds Subnets: ", result.Subnets)
fmt.Println("rds BackupPolicy: ", result.BackupPolicy)
fmt.Println("rds RoGroupList: ", result.RoGroupList)
fmt.Println("rds NodeMaster: ", result.NodeMaster)
fmt.Println("rds NodeSlave: ", result.NodeSlave)
fmt.Println("rds NodeReadReplica: ", result.NodeReadReplica)
fmt.Println("rds DeployId: ", result.DeployId)
fmt.Println("rds SyncMode: ", result.SyncMode)
fmt.Println("rds Category: ", result.Category)
fmt.Println("rds ZoneNames: ", result.ZoneNames)
fmt.Println("rds Endpoint: ", result.Endpoint)
```

## 实例列表
使用以下代码可以查询实例列表信息,优先获取DDC实例列表,获取完成后再获取RDS实例列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC&RDS
args := &ddcrds.ListRdsArgs{
    // 批量获取列表的查询的起始位置，实例列表中Marker需要指定实例Id，可选
    // Marker: "marker",
    // 指定每页包含的最大数量(主实例)，最大数量不超过1000，缺省值为1000，可选
    MaxKeys: 1,
}
resp, err := client.ListRds(args)

if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}

// 返回标记查询的起始位置
fmt.Println("list marker: ", resp.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("list isTruncated: ", resp.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("list nextMarker: ", resp.NextMarker)
// 每页包含的最大数量
fmt.Println("list maxKeys: ", resp.MaxKeys)

// 获取instance的列表信息
for _, e := range resp.Instances {
    fmt.Println("instance productType: ", e.ProductType())
    fmt.Println("instanceId: ", e.InstanceId)
    fmt.Println("instanceName: ", e.InstanceName)
    fmt.Println("engine: ", e.Engine)
    fmt.Println("engineVersion: ", e.EngineVersion)
    fmt.Println("instanceStatus: ", e.InstanceStatus)
    fmt.Println("cpuCount: ", e.CpuCount)
    fmt.Println("memoryCapacity: ", e.MemoryCapacity)
    fmt.Println("volumeCapacity: ", e.VolumeCapacity)
    fmt.Println("usedStorage: ", e.UsedStorage)
    fmt.Println("paymentTiming: ", e.PaymentTiming)
    fmt.Println("instanceType: ", e.InstanceType)
    fmt.Println("instanceCreateTime: ", e.InstanceCreateTime)
    fmt.Println("instanceExpireTime: ", e.InstanceExpireTime)
    fmt.Println("publicAccessStatus: ", e.PublicAccessStatus)
    fmt.Println("vpcId: ", e.VpcId)
    fmt.Println("endpoint: ", e.Endpoint)
    fmt.Println("vnetIp: ", e.Endpoint.VnetIp)
    fmt.Println("vnetIpBackup: ", e.Endpoint.VnetIpBackup)
}
```

## 删除实例
使用以下代码可以批量删除实例,RDS产品将删除实例,DDC产品会将实例放入回收站。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// 多个实例间用英文半角逗号","隔开，最多可输入10个
// DDC
err := client.DeleteRds(instanceIds)
if err != nil {
    fmt.Printf("delete instance error: %+v\n", err)
    return
}
fmt.Printf("delete instance success\n")

// RDS
err := client.DeleteRds(instanceIds)
if err != nil {
    fmt.Printf("delete instance error: %+v\n", err)
    return
}
fmt.Printf("delete instance success\n")

```

## 修改实例名称
使用以下代码可以修改实例名称。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.UpdateInstanceNameArgs{
	// DDC实例名称，允许小写字母、数字，中文，长度限制为1~64
	InstanceName: "instanceName",
}
// DDC
err := client.UpdateInstanceName(instanceId, args)
if err != nil {
    fmt.Printf("update instance name error: %+v\n", err)
    return
}
fmt.Printf("update instance name success\n")

// RDS
err := client.UpdateInstanceName(instanceId, args)
if err != nil {
    fmt.Printf("update instance name error: %+v\n", err)
    return
}
fmt.Printf("update instance name success\n")
```

## 主备切换
使用以下代码可以进行主备切换,支持立即切换或者下一个操作窗口内切换(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.SwitchArgs{
	// 立即切换
    IsSwitchNow: true,
}
err := client.SwitchInstance(instanceId, args)
if err != nil {
    fmt.Printf(" main standby switching of the instance error: %+v\n", err)
    return
}
fmt.Printf(" main standby switching of the instance success\n")
```

## 只读组列表
使用以下代码可以查询只读组列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

resp, err := client.ListRoGroup(instanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
// 获取只读组信息
for _, e := range resp.RoGroups {
	fmt.Println("ddc roGroupId: ", e.RoGroupId)
    fmt.Println("ddc vnetIp: ", e.VnetIp)
}
```
> 注意:
> - 实例状态为Available，实例必须是主实例。


## 更新只读组并分配权重
用于更新只读组的信息，包括设置实例延迟剔除，设置只读实例权重等(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// DDC
replicaWeight := ddcrds.ReplicaWeight{
    // 只读实例ID
    InstanceId: replicaId,
    // 新权重
    Weight: 20,
}
args := &ddcrds.UpdateRoGroupWeightArgs{
    // 只读实例组名称
    RoGroupName: "testRo",
    // RO组是否启用延迟剔除。需传入数字,取值为0或1,0代表不开启,可选
    EnableDelayOff: "0",
    // 延迟阈值,单位为秒。需传入数字,可选，启用延迟剔除时必选
    DelayThreshold: "0",
    // RO组最少保留实例数目。需传入数字,可选,启用延迟剔除时必选
    LeastInstanceAmount: "1",
    // 是否重新进行负载均衡,需传入数字,取值为0或1,可选
    IsBalanceRoLoad: "0",
    // 只读副本新权重数组
    ReplicaList: []ddcrds.ReplicaWeight{replicaWeight},
}
err := client.UpdateRoGroupReplicaWeight(roGroupId, args, "ddc")
if err != nil {
    fmt.Printf("update ddc roGroup replica weight error: %+v\n", err)
    return
}
fmt.Println("update ddc roGroup replica weight success")
```

## 只读组发起重新负载均衡
使用以下代码可以对只读组发起一次重新负载均衡(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// DDC
err := client.ReBalanceRoGroup(roGroupId, "ddc")
if err != nil {
    fmt.Printf("reBalance ddc roGroup error: %+v\n", err)
    return
}
fmt.Println("reBalance ddc roGroup success")
```

## 创建RDS代理实例

使用以下代码可以创建一个RDS代理实例(仅支持RDS)
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.CreateRdsProxyArgs{
    //主实例ID，必选
    SourceInstanceId: "sourceInstanceId",
    // 计费相关参数，代理实例只支持后付费Postpaid，必选
    Billing: ddcrds.Billing{
        PaymentTiming: "Postpaid",
    },
    // 代理实例节点数。取值范围2，4，6，8，16，必选
    NodeAmount: 2,
    //实例名称，允许小写字母、数字，长度限制为1~32，默认命名规则:{engine} + {engineVersion}，可选
    InstanceName: "instanceName",
    //指定zone信息，默认为空，由系统自动选择，可选
    //zoneName命名规范是小写的“国家-region-可用区序列"，例如北京可用区A为"cn-bj-a"，建议与主实例的可用区保持一致
    ZoneNames: []string{"cn-bj-a"},
    //与主实例 vpcId 相同，可选
    VpcId: "vpc-IyrqYIQ7",
    //是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
    IsDirectPay: false,
    //vpc内，每个可用区的subnetId；如果不是默认vpc则必须指定 subnetId，可选
    Subnets: []ddcrds.SubnetMap{
        {
            ZoneName: "cn-bj-a",
            SubnetId: "sbn-IyWRnII7",
        },   
    },
    // 实例绑定的标签信息，可选
    Tags: []ddcrds.TagModel{
        {
            TagKey:   "tagK",
            TagValue: "tagV",
        },
    },
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

## RDS实例扩缩容

使用以下代码可以对RDS实例扩缩容操作。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.ResizeRdsArgs{
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
    // 是否立即变配 RDS只支持立即变配
    IsResizeNow: true,
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
DDC实例支持用户决定是否延迟重启(下一个操作窗口重启)

```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// 立即重启:实例ID为RDS实例ID或DDC实例ID
err := client.RebootInstance(instanceId)
if err != nil {
    fmt.Printf("reboot error: %+v\n", err)
    return
}

// 延迟重启(仅支持DDC)
args := &ddcrds.RebootArgs{
    IsRebootNow: false,
}
err = client.RebootInstanceWithArgs(instanceId, args)
if err != nil {
    fmt.Printf("reboot ddc error: %+v\n", err)
    return
}
```

## 修改同步模式

使用以下代码可以修改RDS实例同步模式(仅支持RDS)。

```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.ModifySyncModeArgs{
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

使用以下代码可以修改RDS域名前缀(仅支持RDS)。

```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.ModifyEndpointArgs{
    Address: "newAddress",
}
err := client.ModifyEndpoint(instanceId, args)
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
使用以下代码可以修改RDS公网访问状态(仅支持RDS)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.ModifyPublicAccessArgs{
    // true or false
    PublicAccess: true,
}
err := client.ModifyPublicAccess(instanceId, args)
if err != nil {
    fmt.Printf("modify public access error: %+v\n", err)
    return
}
fmt.Printf("modify public access success\n")
```

> 注意:
>
> - true：开启公网访问； false：关闭公网访问。

## 已创建实例自动续费(仅支持RDS)

使用以下代码可以为已创建的预付费实例创建自动续费
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.AutoRenewArgs{
    // 自动续费时长（续费单位为year 不大于3，续费单位为month 不大于9）必选
	AutoRenewTime: 1,
    // 自动续费单位（"year";"month"）必选
	AutoRenewTimeUnit: "month",
    // 实例id集合,必须为预付费实例的Id列表 必选
    InstanceIds: []string{
        "rds-y9dJu77d", 
        "rds-aQFOoncr",
    },
}
err := client.AutoRenew(args, "rds")
if err != nil {
    fmt.Printf("create auto renew error: %+v\n", err)
    return
}
```
> 注意:
>
> - 用于已创建的实例开启自动续费。
> - 可以传入多个实例id，多个实例需保证在同一地域。

## 查询实例的维护时间窗口
使用以下代码可以查询实例的维护时间窗口(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

maintainTime,err := client.GetMaintainTime(instanceId)
if err != nil {
    fmt.Printf("get maintain time error: %+v\n", err)
    return
}
fmt.Println("maintainTime duration", maintainTime.Duration)
fmt.Println("maintainTime period", maintainTime.Period)
fmt.Println("maintainTime startTime", maintainTime.StartTime)
fmt.Printf("get maintain time success\n")
```

> 注意:
>
> - startTime 为北京时间24小时制，例如14:00。

## 修改实例的维护时间窗口
使用以下代码可以修改实例的维护时间窗口(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.MaintainTime{
    // 时长间隔,单位为小时
    Duration: 3,
    // 1-0分别代表周一到周日
    Period: "1,2,3,4,5,6,0",
    // 所有涉及的时间皆为北京时间24小时制
    StartTime: "14:00",
}
err := client.UpdateMaintainTime(instanceId, args)
if err != nil {
    fmt.Printf("update maintain time error: %+v\n", err)
}
```

> 注意:
>
> - startTime 为北京时间24小时制，例如14:00。

## 获取回收站中的实例列表
使用以下代码可以获取回收站中的实例列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// marker分页参数
marker := &ddcrds.Marker{MaxKeys: 10}
instances, err := client.ListRecycleInstances(marker, "ddc")
if err != nil {
    fmt.Printf("list recycler instances error: %+v\n", err)
    return
}
for _, instance := range instances.Result {
    fmt.Println("instanceId: ", instance.InstanceId)
    fmt.Println("instanceName: ", instance.InstanceName)
    fmt.Println("engine: ", instance.Engine)
    fmt.Println("engineVersion: ", instance.EngineVersion)
    fmt.Println("instanceStatus: ", instance.InstanceStatus)
    fmt.Println("cpuCount: ", instance.CpuCount)
    fmt.Println("memoryCapacity: ", instance.MemoryCapacity)
    fmt.Println("volumeCapacity: ", instance.VolumeCapacity)
    fmt.Println("usedStorage: ", instance.UsedStorage)
    fmt.Println("instanceType: ", instance.InstanceType)
    fmt.Println("instanceCreateTime: ", instance.InstanceCreateTime)
    fmt.Println("instanceExpireTime: ", instance.InstanceExpireTime)
    fmt.Println("publicAccessStatus: ", instance.PublicAccessStatus)
    fmt.Println("vpcId: ", instance.VpcId)
}
```

## 从回收站中批量恢复实例
使用以下代码可以从回收站中批量恢复实例(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// 要恢复的实例Id列表
instanceIds := []string{
    instanceId_1,
	instanceId_2,
}
err := client.RecoverRecyclerInstances(instanceIds)
if err != nil {
    fmt.Printf("recover recycler instances error: %+v\n", err)
    return
}
fmt.Println("recover recycler instances success.")
```

## 从回收站中批量删除实例
使用以下代码可以从回收站中批量删除实例,实例将被彻底删除(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// 要删除的实例Id列表
instanceIds := []string{
    instanceId_1,
    instanceId_2,
}
err := client.DeleteRecyclerInstances(instanceIds)
if err != nil {
    fmt.Printf("delete recycler instances error: %+v\n", err)
    return
}
fmt.Println("delete recycler instances success.")
```


# 数据库管理

## 创建数据库

使用以下代码可以在某个主实例下创建一个新的数据库(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
args := &ddcrds.CreateDatabaseArgs{
    // 幂等性Token，使用 uuid 生成一个长度不超过64位的ASCII字符串，可选参数
    ClientToken: "xxxyyyzzz",
    // 数据库名称（由大小写字母、数字、下划线组成、字母开头，字母或数字结尾，最长64个字符）
    DbName: "dbName",
    // 数据库字符集（取值范围：utf8、gbk、latin1、utf8mb4）
    CharacterSetName: "utf8",
    // 数据库备注，最多256个字符（一个汉字等于三个字符）
    Remark: "remark",
}
err := client.CreateDatabase(instanceId, args)
if err != nil {
    fmt.Printf("create database error: %+v\n", err)
    return
}

fmt.Println("create database success.")
```

> 注意:
> - 实例状态为Available，实例必须是主实例。

## 更新数据库备注

使用以下代码可以更新数据库的备注(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
args := &ddcrds.UpdateDatabaseRemarkArgs{
    // 数据库备注，最多256个字符（一个汉字等于三个字符）
    Remark: "remark",
}
err = client.UpdateDatabaseRemark(instanceId, dbName, args)
if err != nil {
    fmt.Printf("update database remark error: %+v\n", err)
    return
}

fmt.Println("update database remark success.")
```

## 查询特定数据库信息

使用以下代码可以查询特定数据库信息。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
result, err := client.GetDatabase(instanceId,dbName)
if err != nil {
    fmt.Printf("get database error: %+v\n", err)
    return
}

// 获取account的列表信息
fmt.Println("ddc dbName: ", result.DbName)
fmt.Println("ddc characterSetName: ", result.CharacterSetName)
// 数据库状态（创建中：Creating；可用中：Available；删除中：Deleting；已删除：Deleted）
fmt.Println("ddc dbStatus: ", result.DbStatus)
fmt.Println("ddc remark: ", result.Remark)
fmt.Println("ddc accountPrivileges: ", result.AccountPrivileges)
```

## 查询数据库列表

使用以下代码可以查询指定实例的数据库列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
result, err := client.ListDatabase(instanceId)
if err != nil {
    fmt.Printf("list database error: %+v\n", err)
    return
}

// 获取account的列表信息
for _, database := range result.Databases {
    fmt.Println("ddc dbName: ", database.DbName)
    fmt.Println("ddc characterSetName: ", database.CharacterSetName)
    // 数据库状态（创建中：Creating；可用中：Available；删除中：Deleting；已删除：Deleted）
    fmt.Println("ddc dbStatus: ", database.DbStatus)
    fmt.Println("ddc remark: ", database.Remark)
    fmt.Println("ddc accountPrivileges: ", database.AccountPrivileges)
}
```

## 删除特定数据库

使用以下代码可以删除特定数据库信息。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
err := client.DeleteDatabase(instanceId,dbName)
if err != nil {
    fmt.Printf("delete database error: %+v\n", err)
    return
}
fmt.Printf("delete database success\n")
```

## 延迟删除数据表:创建硬链接

使用以下代码可以为数据库下的指定数据表创建硬链接(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
err := client.LazyDropCreateHardLink(instanceId, dbName, tableName)
if err != nil {
    fmt.Printf("[lazy drop] create hard link error: %+v\n", err)
    return
}
fmt.Println("[lazy drop] create hard link success.")
```


## 延迟删除数据表:删除硬链接

使用以下代码可以删除数据库下特定数据表创建的硬链接(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
err := client.LazyDropDeleteHardLink(instanceId, dbName, tableName)
if err != nil {
    fmt.Printf("[lazy drop] delete hard link error: %+v\n", err)
    return
}
fmt.Println("[lazy drop] delete hard link success.")
```

# 账号管理

## 创建账号

使用以下代码可以在某个主实例下创建一个新的账号。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.CreateAccountArgs{
    // 幂等性Token，使用 uuid 生成一个长度不超过64位的ASCII字符串，可选参数
    ClientToken: "xxxyyyzzz",
	// 账号名称，由小写字母、数字、下划线组成、字母开头，字母或数字结尾，最长16个字符，不能为保留关键字，必选
    AccountName: "accountName",
    // 账号的密码，由字母、数字和特殊字符（!@#%^_）中的至少两种组成，长度8-32位，必选
    Password: "password",
    // 账号权限类型，Common：普通帐号,Super：super账号。可选，默认为 Common
    AccountType: "Common",
    // 权限设置，可选
    DatabasePrivileges: []ddcrds.DatabasePrivilege{
            {
                // 数据库名称
                DbName: "user_photo_001",
                // 授权类型。ReadOnly：只读，ReadWrite：读写
                AuthType: "ReadOnly",
            },   
        },
    // 帐号备注，最多256个字符（一个汉字等于三个字符），可选
    Desc: "账号user1",
}
// DDC产品
err := client.CreateAccount(instanceId, args)
if err != nil {
    fmt.Printf("create account error: %+v\n", err)
    return
}
// RDS产品
err = client.CreateAccount(instanceId, args)
if err != nil {
    fmt.Printf("create account error: %+v\n", err)
    return
}

fmt.Println("create account success.")
```
## 更新账号密码

使用以下代码可以更新账号的密码(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.UpdateAccountPasswordArgs{
    // 密码，由字母、数字和特殊字符（!@#%^_）中的至少两种组成，长度8-32位，必选
    Password: "password",
}
// DDC
err := client.UpdateAccountPassword(instanceId, accountName, args)
if err != nil {
    fmt.Printf("update account password error: %+v\n", err)
    return
}

fmt.Println("update account password success.")
```

## 更新账号备注

使用以下代码可以更新账号的备注(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.UpdateAccountDescArgs{
    // 帐号备注，最多256个字符（一个汉字等于三个字符），可选
    Desc: "desc",
}
err := client.UpdateAccountDesc(instanceId, accountName, args)
if err != nil {
    fmt.Printf("update account desc error: %+v\n", err)
    return
}

fmt.Println("update account desc success.")
```

## 更新账号权限

使用以下代码可以更新账号的权限(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

databasePrivileges := []ddcrds.DatabasePrivilege{
    {
        DbName:   "hello",
		// 授权类型。ReadOnly：只读，ReadWrite：读写
        AuthType: "ReadOnly",
    },
}

args := &ddcrds.UpdateAccountPrivilegesArgs{
    DatabasePrivileges: databasePrivileges,
}
// DDC
err = client.UpdateAccountPrivileges(instanceId, accountName, args)
if err != nil {
    fmt.Printf("update account privileges error: %+v\n", err)
    return
}

fmt.Println("update account privileges success.")
```

## 查询特定账号信息

使用以下代码可以查询特定账号信息。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// DDC产品
result, err := client.GetAccount(instanceId,accountName)
if err != nil {
    fmt.Printf("get account error: %+v\n", err)
    return
}

// 获取account的信息
fmt.Println("ddc accountName: ", result.AccountName)
fmt.Println("ddc desc: ", result.Desc)
// 账号状态（创建中：Creating；可用中：Available；更新中：Updating；删除中：Deleting；已删除：Deleted）
fmt.Println("ddc accountStatus: ", result.Status)
// 账号类型（super账号：Super；普通账号：Common）
fmt.Println("ddc accountType: ", result.AccountType)
fmt.Println("ddc databasePrivileges: ", result.DatabasePrivileges)

// RDS产品
result, err = client.GetAccount(instanceId,accountName)
    if err != nil {
    fmt.Printf("get account error: %+v\n", err)
    return
}

// 获取account的信息
fmt.Println("ddc accountName: ", result.AccountName)
fmt.Println("ddc desc: ", result.Desc)
// 账号状态（创建中：Creating；可用中：Available；更新中：Updating；删除中：Deleting；已删除：Deleted）
fmt.Println("ddc accountStatus: ", result.Status)
// 账号类型（super账号：Super；普通账号：Common）
fmt.Println("ddc accountType: ", result.AccountType)
fmt.Println("ddc databasePrivileges: ", result.DatabasePrivileges)
```

## 查询账号列表

使用以下代码可以查询指定实例的账号列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
result, err := client.ListAccount(instanceId)
    if err != nil {
    fmt.Printf("list account error: %+v\n", err)
    return
}

// 获取account的列表信息
for _, account := range result.Accounts {
    fmt.Println("ddc accountName: ", account.AccountName)
    fmt.Println("ddc desc: ", account.Desc)
    // 账号状态（创建中：Creating；可用中：Available；更新中：Updating；删除中：Deleting；已删除：Deleted）
    fmt.Println("ddc accountStatus: ", account.Status)
    // 账号类型（super账号：Super；普通账号：Common）
    fmt.Println("ddc accountType: ", account.AccountType)
    fmt.Println("ddc databasePrivileges: ", account.DatabasePrivileges)
}

// RDS
result, err = client.ListAccount(instanceId)
    if err != nil {
    fmt.Printf("list account error: %+v\n", err)
    return
}

// 获取account的列表信息
for _, account := range result.Accounts {
    fmt.Println("ddc accountName: ", account.AccountName)
    fmt.Println("ddc desc: ", account.Desc)
    // 账号状态（创建中：Creating；可用中：Available；更新中：Updating；删除中：Deleting；已删除：Deleted）
    fmt.Println("ddc accountStatus: ", account.Status)
    // 账号类型（super账号：Super；普通账号：Common）
    fmt.Println("ddc accountType: ", account.AccountType)
    fmt.Println("ddc databasePrivileges: ", account.DatabasePrivileges)
}
```

## 删除特定账号信息

使用以下代码可以删除特定账号信息。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
err := client.DeleteAccount(instanceId,accountName)
    if err != nil {
    fmt.Printf("delete account error: %+v\n", err)
    return
}
fmt.Printf("delete account success\n")

// RDS
err = client.DeleteAccount(instanceId,accountName)
    if err != nil {
    fmt.Printf("delete account error: %+v\n", err)
    return
}
fmt.Printf("delete account success\n")
```

# 参数管理

## 实例参数列表
使用以下代码可以查询实例参数列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

resp, err := client.ListParameters(instanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
// 获取参数列表信息
for _, e := range resp.Parameters {
	fmt.Println("ddc name: ", e.Name)
    fmt.Println("ddc defaultValue: ", e.DefaultValue)
    fmt.Println("ddc value: ", e.Value)
    fmt.Println("ddc pendingValue: ", e.PendingValue)
    fmt.Println("ddc type: ", e.Type)
    fmt.Println("ddc dynamic: ", e.Dynamic)
    fmt.Println("ddc modifiable: ", e.Modifiable)
    fmt.Println("ddc allowedValues: ", e.AllowedValues)
    fmt.Println("ddc desc: ", e.Desc)
}

// RDS
resp, err := client.ListParameters(instanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
// 获取参数列表信息
for _, e := range resp.Parameters {
    fmt.Println("rds name: ", e.Name)
    fmt.Println("rds defaultValue: ", e.DefaultValue)
    fmt.Println("rds value: ", e.Value)
    fmt.Println("rds pendingValue: ", e.PendingValue)
    fmt.Println("rds type: ", e.Type)
    fmt.Println("rds dynamic: ", e.Dynamic)
    fmt.Println("rds modifiable: ", e.Modifiable)
    fmt.Println("rds allowedValues: ", e.AllowedValues)
    fmt.Println("rds desc: ", e.Desc)
}
```

## 修改实例参数
使用以下代码可以修改云数据库 DDC 的参数配置。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.UpdateParameterArgs{
	Parameters:  []ddcrds.KVParameter{
		{		
			Name: "connect_timeout",
			Value: "15",
		},
	},
}
// DDC
// Etag传入空字符串即可
err := client.UpdateParameter(instanceId, "", args)
if err != nil {
	fmt.Printf("update parameter: %+v\n", err)
	return
}
fmt.Printf("update parameter success\n")

// RDS
// RDS修改参数时需要匹配Etag
res, err := client.ListParameters(e.InstanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
err = client.UpdateParameter(instanceId, res.Etag, args)
if err != nil {
    fmt.Printf("update parameter: %+v\n", err)
    return
}
fmt.Printf("update parameter success\n")
```

# 安全管理

## 白名单列表
使用以下代码可以查询实例白名单列表。
```go
import (
    "encoding/json"
    ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
)

// DDC
result, err := client.GetSecurityIps(instanceId)
if err != nil {
    fmt.Printf("get securityIp list error: %+v\n", err)
    return
}
data, _ := json.Marshal(result)
fmt.Println(string(data))
fmt.Printf("get securityIp list success\n")

// RDS
result, err := client.GetSecurityIps(instanceId)
if err != nil {
    fmt.Printf("get securityIp list error: %+v\n", err)
    return
}
data, _ := json.Marshal(result)
fmt.Println(string(data))
fmt.Printf("get securityIp list success\n")
```

## 更新白名单
使用以下代码可以更新一个实例下的白名单列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

args := &ddcrds.UpdateSecurityIpsArgs{
SecurityIps:  []string{
		"%",
		"192.0.0.1",
		"192.0.0.2",
	},
}
// DDC
// Etag传入空字符串即可
er := client.UpdateSecurityIps(instanceId, "",args)
if er != nil {
	fmt.Printf("update securityIp list error: %+v\n", er)
	return
}
fmt.Printf("update securityIp list success\n")

// RDS
// Etag传入v0
er = client.UpdateSecurityIps(instanceId, "v0",args)
if er != nil {
    fmt.Printf("update securityIp list error: %+v\n", er)
    return
}
fmt.Printf("update securityIp list success\n")
```

## 获取VPC下的安全组
使用以下代码可以获取指定VPC下的安全组列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// vpcId := "vpc-j1vaxw1cx2mw"
securityGroups, err := client.ListSecurityGroupByVpcId(vpcId)
if err != nil {
    fmt.Printf("list security group by vpcId error: %+v\n", err)
    return
}
for _, group := range *securityGroups {
    fmt.Println("securityGroup id: ", group.SecurityGroupID)
    fmt.Println("name: ", group.Name)
    fmt.Println("description: ", group.Description)
    fmt.Println("associateNum: ", group.AssociateNum)
    fmt.Println("createdTime: ", group.CreatedTime)
    fmt.Println("version: ", group.Version)
    fmt.Println("defaultSecurityGroup: ", group.DefaultSecurityGroup)
    fmt.Println("vpc name: ", group.VpcName)
    fmt.Println("vpc id: ", group.VpcShortID)
    fmt.Println("tenantId: ", group.TenantID)
}
fmt.Println("list security group by vpcId success.")
```

## 获取实例已绑定安全组
使用以下代码可以获取指定实例已绑定的安全组列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// instanceId := "ddc-m1h4mma5"
result, err := client.ListSecurityGroupByInstanceId(instanceId)
if err != nil {
    fmt.Printf("list security group by instanceId error: %+v\n", err)
    return
}
for _, group := range result.Groups {
    fmt.Println("securityGroupId: ", group.SecurityGroupID)
    fmt.Println("securityGroupName: ", group.SecurityGroupName)
    fmt.Println("securityGroupRemark: ", group.SecurityGroupRemark)
    fmt.Println("projectId: ", group.ProjectID)
    fmt.Println("vpcId: ", group.VpcID)
    fmt.Println("vpcName: ", group.VpcName)
    fmt.Println("inbound: ", group.Inbound)
    fmt.Println("outbound: ", group.Outbound)
}
fmt.Println("list security group by instanceId success.")
```

## 绑定安全组
使用以下代码可以批量将指定的安全组绑定到实例上(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// instanceIds := []string{
//     "ddc-mjafcdu0",
// }
// securityGroupIds := []string{
//     "g-iutg5rtcydsk",
// }
args := &ddcrds.SecurityGroupArgs{
    InstanceIds: instanceIds,
    SecurityGroupIds: securityGroupIds,
}

err := client.BindSecurityGroups(args)
if err != nil {
    fmt.Printf("bind security groups to instances error: %+v\n", err)
    return
}
fmt.Println("bind security groups to instances success.")
```
> 注意:
> - 实例状态必须为Available。
> - 实例ID最多可以传入10个。
> - 安全组ID最多可以传入10个。
> - 每个实例最多可以绑定10个安全组。

## 解绑安全组
使用以下代码可以从实例上批量解绑指定的安全组(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// instanceIds := []string{
//     "ddc-mjafcdu0",
// }
// securityGroupIds := []string{
//     "g-iutg5rtcydsk",
// }
args := &ddcrds.SecurityGroupArgs{
    InstanceIds: instanceIds,
    SecurityGroupIds: securityGroupIds,
}

err := client.UnBindSecurityGroups(args)
if err != nil {
    fmt.Printf("unbind security groups to instances error: %+v\n", err)
    return
}
fmt.Println("unbind security groups to instances success.")
```
> 注意:
> - 实例状态必须为Available。
> - 当前版本实例ID最多可以传入1个。
> - 安全组ID最多可以传入10个。

# 备份管理
## 	获取备份列表
使用以下代码可以获取一个实例下的备份列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// DDC
args := &ddcrds.GetBackupListArgs{}
resp, err := client.GetBackupList(instanceId, args)
if err != nil {
    fmt.Printf("get backup list error: %+v\n", err)
    return
}
// 返回标记查询的起始位置
fmt.Println("ddc usedSpaceInMB: ", resp.UsedSpaceInMB)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("ddc freeSpaceInMB: ", resp.FreeSpaceInMB)
// 获取参数列表信息
for _, e := range resp.Backups {
    fmt.Println("ddc snapshotId: ", e.SnapshotId)
    fmt.Println("ddc snapshotSizeInBytes: ", e.SnapshotSizeInBytes)
    fmt.Println("ddc snapshotType: ", e.SnapshotType)
    fmt.Println("ddc snapshotStatus: ", e.SnapshotStatus)
    fmt.Println("ddc snapshotStartTime: ", e.SnapshotStartTime)
    fmt.Println("ddc snapshotEndTime: ", e.SnapshotEndTime)
}

// RDS
resp, err = client.GetBackupList(instanceId, args)
if err != nil {
    fmt.Printf("get backup list error: %+v\n", err)
    return
}
// 返回标记查询的起始位置
fmt.Println("rds usedSpaceInMB: ", resp.UsedSpaceInMB)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("rds freeSpaceInMB: ", resp.FreeSpaceInMB)
// 获取参数列表信息
for _, e := range resp.Backups {
    fmt.Println("rds snapshotId: ", e.SnapshotId)
    fmt.Println("rds snapshotSizeInBytes: ", e.SnapshotSizeInBytes)
    fmt.Println("rds snapshotType: ", e.SnapshotType)
    fmt.Println("rds snapshotStatus: ", e.SnapshotStatus)
    fmt.Println("rds snapshotStartTime: ", e.SnapshotStartTime)
    fmt.Println("rds snapshotEndTime: ", e.SnapshotEndTime)
}
```

## 创建备份
使用以下代码创建实例备份(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
err := client.CreateBackup(instanceId)
if err != nil {
    fmt.Printf("create backup error: %+v\n", err)
    return
}
fmt.Printf("create backup success\n")
```

## 备份详情
使用以下代码可以查询一个备份的详情(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
resp, err := client.GetBackupDetail(instanceId, snapshotId)
if err != nil {
    fmt.Printf("get backup detail error: %+v\n", err)
    return
}

fmt.Println("ddc snapshotId: ", resp.Snapshot.SnapshotId)
fmt.Println("ddc snapshotSizeInBytes: ", resp.Snapshot.SnapshotSizeInBytes)
fmt.Println("ddc snapshotType: ", resp.Snapshot.SnapshotType)
fmt.Println("ddc snapshotStatus: ", resp.Snapshot.SnapshotStatus)
fmt.Println("ddc snapshotStartTime: ", resp.Snapshot.SnapshotStartTime)
fmt.Println("ddc snapshotEndTime: ", resp.Snapshot.SnapshotEndTime)
fmt.Println("ddc downloadUrl: ", resp.Snapshot.DownloadUrl)
fmt.Println("ddc downloadExpires: ", resp.Snapshot.DownloadExpires)
```

## 设置备份策略
使用以下代码设置实例的备份策略(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// DDC
args := &ddcrds.BackupPolicy{
	// 以英文半角逗号分隔的备份时日间，周日为第一天，取值0
	BackupDays: "0,1,2,3,5,6",
	// 备份开始时间，使用UTC时间
	BackupTime: "17:00:00Z",
	// 是否启用备份数据持久化
	Persistent:	true,
	// 持久化天数，范围7-730天；未启用则为0或不填
	ExpireInDays: 7,
}
err := client.ModifyBackupPolicy(instanceId, args)
if err != nil {
    fmt.Printf("modify instance's backupPolicy error: %+v\n", err)
    return
}
fmt.Printf("modify instance's backupPolicy success\n")
```

## binlog列表
使用以下代码可以获取一个实例下的binlog列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// DDC
// datetime UTC时间
// 获取了两天前的日志备份
datetime := time.Now().
    AddDate(0, 0, -2).
    Format("2006-01-02T15:04:05Z")
resp, err := client.GetBinlogList(instanceId, datetime)
if err != nil {
    fmt.Printf("get binlog list error: %+v\n", err)
    return
}
// 获取binlog列表信息
for _, e := range resp.Binlogs {
    fmt.Println("ddc binlogId: ", e.BinlogId)
    fmt.Println("ddc binlogSizeInBytes: ", e.BinlogSizeInBytes)
    fmt.Println("ddc binlogStatus: ", e.BinlogStatus)
    fmt.Println("ddc binlogStartTime: ", e.BinlogStartTime)
    fmt.Println("ddc binlogEndTime: ", e.BinlogEndTime)
}
```

## binlog 详情
使用以下代码可以查询一个binlog详情(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
resp, err := client.GetBinlogDetail(instanceId, binlogId)
if err != nil {
    fmt.Printf("get binlog detail error: %+v\n", err)
    return
}

fmt.Println("ddc binlogId: ", resp.Binlog.BinlogId)
fmt.Println("ddc binlogSizeInBytes: ", resp.Binlog.BinlogSizeInBytes)
fmt.Println("ddc binlogStatus: ", resp.Binlog.BinlogStatus)
fmt.Println("ddc binlogStartTime: ", resp.Binlog.BinlogStartTime)
fmt.Println("ddc binlogEndTime: ", resp.Binlog.BinlogEndTime)
fmt.Println("ddc downloadUrl: ", resp.Binlog.DownloadUrl)
fmt.Println("ddc downloadExpires: ", resp.Binlog.DownloadExpires)
```

# 日志管理

## 日志列表
使用以下代码可以获取一个实例下的错误日志或者慢日志列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
// import "time"

// datetime UTC时间
// 获取两天前的错误日志,传入日期即可
date := time.Now().
AddDate(0, 0, -2).
Format("2006-01-02")
args := &ddcrds.ListLogArgs{
    // 日志类型 错误日志为error 慢日志为slow
    LogType:  "error",
    Datetime: date,
}
logs, err := client.ListLogByInstanceId(instanceId, args)
if err != nil {
    fmt.Printf("list logs of instance error: %+v\n", err)
    return
}
fmt.Println("list logs of instance success.")
for _, log := range *logs {
    fmt.Println("id: ", log.LogID)
    fmt.Println("size: ", log.LogSizeInBytes)
    fmt.Println("start time: ", log.LogStartTime)
    fmt.Println("end time: ", log.LogEndTime)
}
```

## 日志详情
使用以下代码可以查询日志的详细信息，包括该日志文件有效的下载链接(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
args := &ddcrds.GetLogArgs{
    // 下载链接有效时间，单位为秒
    ValidSeconds: 20,
}
logId := "errlog.202103091300"
log, err := client.GetLogById(instanceId, logId, args)
if err != nil {
    fmt.Printf("get log detail of instance error: %+v\n", err)
    return
}
fmt.Println("list logs of instances success.")
fmt.Println("id: ", log.LogID)
fmt.Println("size: ", log.LogSizeInBytes)
fmt.Println("start time: ", log.LogStartTime)
fmt.Println("end time: ", log.LogEndTime)
// 日志文件下载链接
fmt.Println("download url: ", log.DownloadURL)
// 下载链接截止该时间有效
fmt.Println("download url expires: ", log.DownloadExpires)
```

# 其他
## VPC列表
使用以下代码可以查询vpc列表(仅支持DDC)。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
resp, err := client.ListVpc("ddc")
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
// 获取vpc列表信息
for _, e := range* resp {
    fmt.Println("ddc vpcId: ", e.VpcId)
    fmt.Println("ddc shortId: ", e.ShortId)
    fmt.Println("ddc name: ", e.Name)
    fmt.Println("ddc cidr: ", e.Cidr)
    fmt.Println("ddc status: ", e.Status)
    fmt.Println("ddc createTime: ", e.CreateTime)
    fmt.Println("ddc description: ", e.Description)
    fmt.Println("ddc defaultVpc: ", e.DefaultVpc)
    fmt.Println("ddc ipv6Cidr: ", e.Ipv6Cidr)
    fmt.Println("ddc auxiliaryCidr: ", e.AuxiliaryCidr)
    fmt.Println("ddc relay: ", e.Relay)
}
```

## 可用区列表
使用以下代码可以获取可用区列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"
resp, err := client.GetZoneList("ddc")
if err != nil {
	fmt.Printf("get zone list error: %+v\n", err)
	return
}
for _, e := range resp.Zones {
    fmt.Println("ddc zoneNames: ", e.ZoneNames)
    fmt.Println("ddc apiZoneNames: ", e.ApiZoneNames)
    fmt.Println("ddc available: ", e.Available)
    fmt.Println("ddc defaultSubnetId: ", e.DefaultSubnetId)
}

// RDS
resp, err = client.GetZoneList("rds")
if err != nil {
    fmt.Printf("get zone list error: %+v\n", err)
    return
}
for _, e := range resp.Zones {
    fmt.Println("rds zoneNames: ", e.ZoneNames)
    fmt.Println("rds apiZoneNames: ", e.ApiZoneNames)
    fmt.Println("rds available: ", e.Available)
    fmt.Println("rds defaultSubnetId: ", e.DefaultSubnetId)
}
```

## 子网列表
使用以下代码可以获取一个实例下的子网列表。
```go
// import ddcrds "github.com/baidubce/bce-sdk-go/services/ddc/v2"

// DDC
args := &ddcrds.ListSubnetsArgs{}
resp, err := client.ListSubnets(args, "ddc")
if err != nil {
    fmt.Printf("get subnet list error: %+v\n", err)
    return
}
for _, e := range resp.Subnets {
    fmt.Println("ddc name: ", e.Name)
    fmt.Println("ddc longId: ", e.LongId)
    fmt.Println("ddc zoneName: ", e.ZoneName)
    fmt.Println("ddc shortId: ", e.ShortId)
    fmt.Println("ddc vpcId: ", e.VpcId)
    fmt.Println("ddc vpcShortId: ", e.VpcShortId)
    fmt.Println("ddc az: ", e.Az)
    fmt.Println("ddc cidr: ", e.Cidr)
    fmt.Println("ddc createdTime: ", e.CreatedTime)
    fmt.Println("ddc updatedTime: ", e.UpdatedTime)
}

// RDS
args = &ddcrds.ListSubnetsArgs{}
resp, err := client.ListSubnets(args, "rds")
if err != nil {
    fmt.Printf("get subnet list error: %+v\n", err)
    return
}
for _, e := range resp.Subnets {
    fmt.Println("rds name: ", e.Name)
    fmt.Println("rds longId: ", e.LongId)
    fmt.Println("rds zoneName: ", e.ZoneName)
    fmt.Println("rds shortId: ", e.ShortId)
    fmt.Println("rds vpcId: ", e.VpcId)
    fmt.Println("rds vpcShortId: ", e.VpcShortId)
    fmt.Println("rds az: ", e.Az)
    fmt.Println("rds cidr: ", e.Cidr)
    fmt.Println("rds createdTime: ", e.CreatedTime)
    fmt.Println("rds updatedTime: ", e.UpdatedTime)
}
```

# 错误处理

GO语言以error类型标识错误，DDC支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | DDC服务返回的错误

用户使用SDK调用DDC相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// ddcClient 为已创建的DDC Client对象
result, err := client.ListDdcInstance()
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

客户端异常表示客户端尝试向DDC发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当DDC服务端出现异常时，DDC服务端会返回给用户相应的错误信息，以便定位问题。

## SDK日志

DDC GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

DDC GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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
log.Debugf("%s", "logging message using the log package in the DDC go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the DDC go sdk")
```

首次发布:

- 支持创建账号、更新账号密码、更新账号备注、更新账号权限、查询账号列表、查询特定账号信息、删除特定账号信息
- 支持创建数据库、更新数据库备注、查询数据库列表、查询特定数据库信息、删除特定数据库信息
- 兼容RDS SDK中现有功能