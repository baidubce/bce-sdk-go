# SCS服务

# 概述

本文档主要介绍SCS GO SDK的使用。在使用本文档前，您需要先了解SCS的一些基本知识。若您还不了解SCS，可以参考[产品描述](https://cloud.baidu.com/doc/SCS/s/Rjxbm160f)和[入门指南](https://cloud.baidu.com/doc/SCS/s/6jxyfy2ly)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[SCS访问域名](https://cloud.baidu.com/doc/SCS/s/fjwvxtrd9)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云SCS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问SCS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建SCS Client

SCS Client是SCS服务的客户端，为开发者与SCS服务进行交互提供了一系列的方法。

### 使用AK/SK新建SCS Client

通过AK/SK方式访问SCS，用户可以参考如下代码新建一个SCS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/scs"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个SCSClient
	scsClient, err := scs.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/SCS/s/ojwvynrqn)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为SCS的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`redis.bj.baidubce.com`。

### 使用STS创建SCS Client

**申请STS token**

SCS可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问SCS，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建SCS Client**

申请好STS后，可将STS Token配置到SCS Client中，从而实现通过STS Token创建SCS Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建SCS Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/scs" //导入SCS服务模块
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

	// 使用申请的临时STS创建SCS服务的Client对象，Endpoint使用默认值
	scsClient, err := scs.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create scs client failed:", err)
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
	scsClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置SCS Client时，无论对应SCS服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问SCS

SCS支持HTTPS传输协议，您可以通过在创建SCS Client对象时指定的Endpoint中指明HTTPS的方式，在SCS GO SDK中使用HTTPS访问SCS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/scs"

ENDPOINT := "https://redis.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
scsClient, _ := scs.NewClient(AK, SK, ENDPOINT)
```

## 配置SCS Client

如果用户需要配置SCS Client的一些细节的参数，可以在创建SCS Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问SCS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/scs"

//创建SCS Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "redis.bj.baidubce.com"
client, _ := scs.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/scs"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "redis.bj.baidubce.com"
client, _ := scs.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/scs"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "redis.bj.baidubce.com"
client, _ := scs.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问SCS时，创建的SCS Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建SCS Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

云数据库 SCS（Simple Cache Service）提供稳定、高效以及高可扩展性的分布式缓存服务。云数据库 SCS 兼容 Redis/Memcached 协议，基于 Redis 提供标准版和集群版的架构模式，并支持自定义副本数量，为您提供多样化的数据结构支持。

## 实例管理

### 创建实例

使用以下代码可以创建SCS实例，用于创建一个或多个redis实例
```go
args := &scs.CreateInstanceArgs{
    // 选择付款方式，可以选择预付费或后付费 
    Billing: api.Billing{
        PaymentTiming: api.PaymentTimingPostPaid,
    },
    // 购买个数，最大不超过10，默认1
	PurchaseCount: 1,
    // 实例名
    // 要求：1）支持大小写字母、数字以及-_ /.等特殊字符，必须以字母开头；2）长度限制为1-64；
	InstanceName:  "sdk-scs",
    // 端口号 1025<port<22222，22222<port<65535
	Port:          6379,
    // 引擎
    Engine: 2, // redis: 2   PegaDB: 3
    // 引擎版本，集群：3.2、4.0、5.0 主从：2.8、3.2、4.0、5.0  当Engine为3时，可不传
	EngineVersion: "3.2",
    // 存储空间 当Engine为3时添加
    // DiskFlavor:     60,
    // 磁盘类型 当Engine为3时添加目前支持cds  essd两个类型
    // DiskType:     "essd",
    // 节点规格
	NodeType:      "cache.n1.micro",
    // 集群类型，集群版："cluster"，主从版："master_slave"
	ClusterType:   "master_slave",
    // 副本个数，单副本为1，双副本为2，多副本依此类推
	ReplicationNum:  1,
    // 分片个数
	ShardNum:      1,
    // 代理节点数，主从版：0，集群版：代理节点数=分片个数
	ProxyNum:      0,
    // 存储类型 0：高性能内存 1：ssd本地磁盘 3:容量型存储（PegaDB专用）
    StoreType: 0,
    // 副本只读 1打开 2关闭 （PegaDB专用）
    // EnableReadOnly: 1,
}

// 若要生成预付费实例，可以按以下配置生成一个月的预付费实例
args.Billing = api.Billing{
		PaymentTiming: api.PaymentTimingPrePaid,
		Reservation: &api.Reservation{
            ReservationLength:   1,
            ReservationTimeUnit: "month",
        }
	}

// 若要设置自动续费，可以按以下参数设置一年内自动续费
// 按月付费或者按年付费 月是"month"，年是"year"
// 自动续费的时间 按月是1-9 按年是 1-3
args.AutoRenewTimeUnit = "year"
args.AutoRenewTime = 1

result, err := client.CreateInstance(args)
if err != nil {
    fmt.Println("create instance failed:", err)
} else {
    fmt.Println("create instance success: ", result)
}
```

> **提示：**
> 1.  创建SCS请求是一个异步请求，返回200表明订单生成，后续可以通过实例id查询实例创建进度。
> 2.  本接口用于创建一个或多个同配置SCS实例。
> 3.  创建实例需要实名认证，没有通过实名认证的可以前往百度开放云官网控制台中的安全认证下的实名认证中进行认证。
> 4.  创建计费方式为后付费的实例需要账户现金余额+通用代金券大于100；预付费方式的实例则需要账户现金余额大于等于实例费用。
> 5.  支支持批量创建，且如果创建过程中有一个实例创建失败，所有实例将全部回滚。
> 6.  创建请求详细使用请参考SCS API 文档[创建实例](https://cloud.baidu.com/doc/SCS/s/hk0qhwxom)
> 7.  创建SCS需要使用指定规格, 详细请参考SCS API 文档[实例规格](https://cloud.baidu.com/doc/SCS/s/1jwvxtsh0#%E5%AE%9E%E4%BE%8B%E8%A7%84%E6%A0%BC)


### 查询实例列表

以下代码可以查询SCS实例列表
```go
args := &scs.ListInstancesArgs{}

result, err := client.ListInstances(args)
if err != nil {
    fmt.Println("list instance failed:", err)
} else {
    fmt.Println("list instance success: ", result)
}
```

### 查询指定实例详情

使用以下代码可以查询指定SCS虚机的详细信息
```go
result, err := client.GetInstanceDetail(instanceId)
if err != nil {
    fmt.Println("get instance detail failed:", err)
} else 
    fmt.Println("get instance detail success ", result)
}
```

### 修改实例名称

如下代码可以修改实例名称
```go
args := &scs.UpdateInstanceNameArgs{
    InstanceName: "newInstanceName",
}
err := client.UpdateInstanceName(instanceId, args)
if err != nil {
    fmt.Println("update instance name failed:", err)
} else {
    fmt.Println("update instance name success")
}
```

> **提示：**
>
> - 只有实例Running状态时可以修改实例名称

### 释放实例

如下代码可以释放实例,实例将自动进入回收站，保留7天后删除
```go
err := client.DeleteInstance(instanceId)
if err != nil {
    fmt.Println("delete instance failed:", err)
} else {
    fmt.Println("delete instance success")
}
```
> **提示：**
> -   释放单个SCS实例，释放后实例将进入回收站，保留7天后自动删除所有物理资源。
> -   可以从回收站内恢复实例或者彻底删除实例,预付费实例需要通过续费实例接口进行恢复。

## 重启实例

使用以下代码可以重启实例。支持用户决定是否延迟到维护窗口内重启。

```go
// 立即重启
args := &scs.RestartInstanceArgs{
    // 需要延迟到维护窗口重启时需传入true
    IsDefer: false,
}
err := SCS_CLIENT.RestartInstance(instanceId, args)
if err != nil {
    fmt.Println("restart instance failed:", err)
} else {
    fmt.Println("restart instance success")
}
```

## 续费实例
使用以下代码可以对已有的预付费实例进行续费,如果实例在回收站中,续费后会从回收站中恢复。
```go

// 要恢复的实例Id列表
instanceIds := []string{
    instancId,
}
args := &scs.RenewInstanceArgs{
    // 实例Id列表
    InstanceIds: instanceIds,
    // 续费周期，单位为月
    Duration: 1,
}
result, err := client.RenewInstances(args)
if err != nil {
    fmt.Printf("renew instances error: %+v\n", err)
    return
}
// 异步任务，返回订单Id
fmt.Println("renew instances success. orderId:" + result.OrderId)
```

## 获取回收站中的实例列表
使用以下代码可以获取回收站中的实例列表。
```go

// marker分页参数
marker := &scs.Marker{MaxKeys: 10}
instances, err := client.ListRecycleInstances(marker)
if err != nil {
    fmt.Printf("list recycler instances error: %+v\n", err)
    return
}
for _, instance := range instances.Result {
    fmt.Println("instanceId: ", instance.InstanceID)
    fmt.Println("instanceName: ", instance.InstanceName)
    fmt.Println("engine: ", instance.Engine)
    fmt.Println("engineVersion: ", instance.EngineVersion)
    fmt.Println("instanceStatus: ", instance.InstanceStatus)
    // 进入回收站后isolatedStatus为Isolated,表示实例已隔离
    fmt.Println("isolatedStatus: ", instance.IsolatedStatus)
    fmt.Println("paymentTiming: ", instance.PaymentTiming)
    fmt.Println("clusterType: ", instance.ClusterType)
    fmt.Println("domain: ", instance.Domain)
    fmt.Println("port: ", instance.Port)
    fmt.Println("vnetIP: ", instance.VnetIP)
    fmt.Println("instanceCreateTime: ", instance.InstanceCreateTime)
    fmt.Println("usedCapacity: ", instance.UsedCapacity)
    fmt.Println("zoneNames: ", instance.ZoneNames)
    fmt.Println("tags: ", instance.Tags)
}
```

## 从回收站中批量恢复实例
使用以下代码可以从回收站中批量恢复实例,批量恢复仅支持后付费实例,回收站中的预付费实例请通过实例续费接口恢复。
```go

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
使用以下代码可以从回收站中批量删除实例,实例将被彻底删除。
```go

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

### 添加节点

```go
args := &scs.ReplicationArgs{
    ResizeType: "add", // add 添加从节点  add_readonly 添加只读节点 
    ReplicationInfo: []Replication{
        {AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 1},
        {AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
        {AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
    },
    ClientToken: getClientToken(),
}
result, err := client.AddReplication(instanceId, args)

if err != nil {
    fmt.Println("resize instance failed:", err)
} else {
    fmt.Println("resize instance success")
```

### 删除节点

```go
args := &scs.ReplicationArgs{
    ResizeType: "delete", // delete 删除节点  delete_readonly 删除只读节点 
    ReplicationInfo: []Replication{
        {AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 1},
        {AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
        {AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
    },
    ClientToken: getClientToken(),
}
result, err := client.DeleteReplication(instanceId, args)

if err != nil {
    fmt.Println("resize instance failed:", err)
} else {
    fmt.Println("resize instance success")
```

### 变更配置

```go
args := &scs.ResizeInstanceArgs{
    NodeType:"cache.n1.small",
    ShardNum:2,
    // 需要延迟到维护窗口变配时需传入true
    IsDefer: false,
}
err := client.ResizeInstance(instanceId, args)
if err != nil {
    fmt.Println("resize instance failed:", err)
} else {
    fmt.Println("resize instance success")
```

> **提示：**
> - 实例计费方式为预付费时，不能进行缩容操作
> - 实例计费方式为后付费时，可弹性扩缩容
> - 只有实例正常运行状态时才可以进行扩缩容操作，变更接口为异步变更，可通过查询指定实例详情接口查询实例状态

### 查询实例套餐规格

如下代码可以查询当前可以创建的实例的套餐的规格
```go
result, err := client.GetNodeTypeList()
if err != nil { 
    fmt.Println("list node type failed: ", err)
} else {
    fmt.Println("list node type success: ", result)
}
```

### 获取子网列表

如下代码可以获取子网列表

```go
args := &scs.ListSubnetsArgs{}
_, err := client.ListSubnets(args)
if err != nil {
    fmt.Println("get subnet list failed:", err)
} else {
    fmt.Println("get subnet list success")
}
```

> **提示：**
>
> - 请求参数 vpcId 和 zoneName 不是必须的

### 修改实例域名

如下代码可以修改实例域名

```go
args := &scs.UpdateInstanceDomainNameArgs{
    Domain:  "newDomainName",
}
err := client.UpdateInstanceDomainName(instanceId, args)
if err != nil {
    fmt.Println("update instance domain name failed:", err)
} else {
    fmt.Println("update instance domain name success")
}
```

> **提示：**
>
> - 只有实例Running状态时可以修改实例域名

### 获取可用区列表

使用以下代码可以获取可用区列表

```go
result, err := client.GetZoneList()
if err != nil {
    fmt.Println("get zone list failed:", err)
} else 
    fmt.Println("get zone list success ", result)
}
```

### 修改访问密码

如下代码可以修改访问密码

```go
args := &scs.ModifyPasswordArgs{
    Password:  "newPassword",
}
result, err := client.ModifyPassword(instanceId, args)
if err != nil {
    fmt.Println("modify password failed:", err)
} else 
    fmt.Println("modify password success ", result)
}
```

> **提示：**
>
> - 密码长度8～16位，至少包含字母、数字和特殊字符中两种。允许的特殊字符包括 $^*()_+-=

### 清空实例

如下代码可以清空实例

```go
args := &scs.FlushInstanceArgs{
    Password:  "Password",
}
result, err := client.FlushInstance(instanceId, args)
if err != nil {
    fmt.Println("flush instance failed:", err)
} else 
    fmt.Println("flush instance success ", result)
}
```

> **提示：**
>
> - 对redis实例，清空后表现为数据占用内存下降，数据被清空；对memcache实例，清空后表现为占用内存不会下降，但是数据被清空
> - 如果没有设置密码，传递空字符串

### 绑定标签

如下代码可以绑定标签

```go
args := &scs.BindingTagArgs{
    ChangeTags:  []model.TagModel{
				{
					TagKey:   "tag1",	
					TagValue: "var1",
				},
	},
}
result, err := client.BindingTag(instanceId, args)
if err != nil {
    fmt.Println("bind tags failed:", err)
} else 
    fmt.Println("bind tags success ", result)
}
```

### 解绑标签

如下代码可以解绑标签

```go
args := &scs.BindingTagArgs{
    ChangeTags:  []model.TagModel{
				{
					TagKey:   "tag1",	
					TagValue: "var1",
				},
	},
}
result, err := client.UnBindingTag(instanceId, args)
if err != nil {
    fmt.Println("unbind tags failed:", err)
} else 
    fmt.Println("unbind tags success ", result)
}
```

> **提示：**
>
> - 解绑实例上定义的标签
> - 可以同时解绑多个标签

### 查询IP白名单

如下代码可以查询允许访问实例的IP白名单

```go
result, err := client.GetSecurityIp(instanceId)
if err != nil {
    fmt.Println("get security IP failed:", err)
} else 
    fmt.Println("get security IP success ", result)
}
```

> **提示：**
>
> - 返回参数 IP白名单列表, 包括常规地址: 如192.168.0.1，CIDR地址: 如192.168.1.0/24，0.0.0.0/0代表允许所有地址

### 增加IP白名单

如下代码可以增加访问实例的IP白名单

```go
args := &scs.SecurityIpArgs{
    SecurityIps:  []string{
					"192.0.0.1",
				},	
}
result, err := client.AddSecurityIp(instanceId, args)
if err != nil {
    fmt.Println("add security IP failed:", err)
} else 
    fmt.Println("add security IP success ", result)
}
```

### 删除IP白名单

如下代码可以删除访问实例的IP白名单

```go
args := &scs.SecurityIpArgs{
    SecurityIps:  []string{
					"192.0.0.1",
				},	
}
result, err := client.DeleteSecurityIp(instanceId, args)
if err != nil {
    fmt.Println("delete security IP failed:", err)
} else 
    fmt.Println("delete security IP success ", result)
}
```

## 获取VPC下的安全组
使用以下代码可以获取指定VPC下的安全组列表。
```go

// vpcId := "vpc-j1vaxw1cx2mw"
securityGroups, err := client.ListSecurityGroupByVpcId(vpcId)
if err != nil {
    fmt.Printf("list security group by vpcId error: %+v\n", err)
    return
}
for _, group := range securityGroups.Groups {
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
使用以下代码可以获取指定实例已绑定的安全组列表。
```go

// instanceId := "scs-m1h4mma5"
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
使用以下代码可以批量将指定的安全组绑定到实例上。
```go

// instanceIds := []string{
//     "scs-mjafcdu0",
// }
// securityGroupIds := []string{
//     "g-iutg5rtcydsk",
// }
args := &scs.SecurityGroupArgs{
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
使用以下代码可以从实例上批量解绑指定的安全组。
```go
// securityGroupIds := []string{
//     "g-iutg5rtcydsk",
// }
args := &scs.UnbindSecurityGroupArgs{
    InstanceId:      "scs-su-bbjhgxyqyddd",
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

### 获取参数列表

使用以下代码可以获取Redis实例的配置参数和运行参数

```go
result, err := client.GetParameters(instanceId)
if err != nil {
    fmt.Println("get parameter list failed:", err)
} else 
    fmt.Println("get parameter list success ", result)
}
```

### 修改参数

如下代码可以修改redis实例参数值

```go
args := &scs.ModifyParametersArgs{
				Parameter: InstanceParam{
					Name: "parameter name",
					Value: "new value",
				},
}
result, err := client.ModifyParameters(instanceId, args)
if err != nil {
    fmt.Println("modify parameters failed:", err)
} else 
    fmt.Println("modify parameters success ", result)
}
```

### 查看备份列表

使用以下代码可以查询某个实例备份列表

```go
result, err := client.GetBackupList(instanceId)
if err != nil {
    fmt.Println("get backup list failed:", err)
} else 
    fmt.Println("get backup list success ", result)
}
```

### 修改备份策略

如下代码可以修改redis实例自动备份策略

```go
args := &scs.ModifyBackupPolicyArgs{
				BackupDays: "Sun,Mon,Tue,Wed,Thu,Fri,Sta",
				BackupTime: "01:05:00",
				ExpireDay: 7,
}
result, err := client.ModifyBackupPolicy(instanceId, args)
if err != nil {
    fmt.Println("modify backup policy failed:", err)
} else 
    fmt.Println("modify backup policy success ", result)
}
```

> **提示：**
>
> - BackupDays: 标识一周中哪几天进行备份备份周期：Mon（周一）Tue（周二）Wed（周三）Thu（周四）Fri（周五）Sat（周六）Sun（周日）逗号分隔，取值如：Sun,Wed,Thu,Fri,Sta
> - BackupTime: 标识一天中何时进行备份，UTC时间（+8为北京时间）取值如：01:05:00
> - ExpireDay: 备份文件过期时间，取值如：3

# 日志管理

## 日志列表
使用以下代码可以获取一个实例下的运行日志或者慢日志列表。
```go
// import "time"

// 一天前
date := time.Now().
    AddDate(0, 0, -1).
    Format("2006-01-02 03:04:05")
fmt.Println(date)
args := &scs.ListLogArgs{
    // 运行日志 runlog 慢日志 slowlog
    FileType:  "runlog",
    // 开始时间格式 "yyyy-MM-dd hh:mm:ss"
    StartTime: date,
    // 结束时间,可选,默认返回开始时间+24小时内的日志
    // EndTime: date,
}
listLogResult, err := client.ListLogByInstanceId(instanceId, args)
if err != nil {
    fmt.Printf("list logs of instance error: %+v\n", err)
    return
}
fmt.Println("list logs of instance success.")
for _, shardLog := range listLogResult.LogList {
    fmt.Println("shard id: ", shardLog.ShardID)
    fmt.Println("logs size: ", shardLog.TotalNum)
    for _, log := range shardLog.LogItem {
        fmt.Println("log id: ", log.LogID)
        fmt.Println("size: ", log.LogSizeInBytes)
        fmt.Println("start time: ", log.LogStartTime)
        fmt.Println("end time: ", log.LogEndTime)
        fmt.Println("download url: ", log.DownloadURL)
        fmt.Println("download url expires: ", log.DownloadExpires)
    }
}
```

## 日志详情
使用以下代码可以查询日志的详细信息，包括该日志文件有效的下载链接。
```go
args := &GetLogArgs{
    // 下载链接有效时间，单位为秒，可选，默认为1800秒
    ValidSeconds: 60,
}
logId := "scs-bj-mktaypucksot_8742_slowlog_202104160330"
log, err := client.GetLogById(instanceId, logId, args)
if err != nil {
    fmt.Printf("get log detail of instance error: %+v\n", err)
    return
}
fmt.Println("get log detail success.")
fmt.Println("id: ", log.LogID)
// 日志文件下载链接
fmt.Println("download url: ", log.DownloadURL)
// 下载链接截止该时间有效
fmt.Println("download url expires: ", log.DownloadExpires)
```

# 维护窗口
## 查询实例的维护时间窗口
使用以下代码可以查询实例的维护时间窗口。
```go
resp, err := client.GetMaintainTime(instanceId)
if err != nil {
    fmt.Printf("get maintainTime of instance error: %+v\n", err)
    return
}
fmt.Println("get maintainTime success.")
fmt.Println("start time: ", resp.MaintainTime.StartTime)
fmt.Println("dutation: ", resp.MaintainTime.Duration)
fmt.Println("period: ", resp.MaintainTime.Period)
```

> 注意:
>
> - startTime 为北京时间24小时制，例如14:00。

## 修改实例的维护时间窗口
使用以下代码可以修改实例的维护时间窗口。
```go
newMaintainTime := &scs.MaintainTime{
    StartTime: "16:00",
    Duration:  1,
    Period:    []int{1, 2, 3},
}
err := client.ModifyMaintainTime(instanceId, newMaintainTime)
if err != nil {
    fmt.Printf("modify maintainTime of instance error: %+v\n", err)
    return
}
fmt.Println("modify maintainTime success.")
```

# 错误处理

GO语言以error类型标识错误，SCS支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | SCS服务返回的错误

用户使用SDK调用SCS相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// scsClient 为已创建的SCS Client对象
instanceDetail, err := scsClient.GetInstanceDetail(instanceId)
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
	fmt.Println("get instance detail success: ", instanceDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向SCS发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当SCS服务端出现异常时，SCS服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[SCS错误返回](https://cloud.baidu.com/doc/SCS/s/Yjwvxtsti)

