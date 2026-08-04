# MONGODB服务

# 概述

本文档主要介绍MONGODB GO SDK的使用。在使用本文档前，您需要先了解MONGODB的一些基本知识，并已开通了MONGODB服务。

# 初始化

## 确认Endpoint

目前支持"华北-北京"区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
北京 | mongodb.bj.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云MONGODB，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问MONGODB做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建MONGODB Client

MONGODB Client是MONGODB服务的客户端，为开发者与MONGODB服务进行交互提供了一系列的方法。

### 使用AK/SK新建MONGODB Client

通过AK/SK方式访问MONGODB，用户可以参考如下代码新建一个MONGODB Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/mongodb"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个MONGODBClient
	mongodbClient, err := mongodb.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的"Access Key ID"，`SECRET_ACCESS_KEY`对应控制台中的"Access Key Secret"，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为MONGODB的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`mongodb.bj.baidubce.com`。

### 使用STS创建MONGODB Client

**申请STS token**

MONGODB可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问MONGODB，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建MONGODB Client**

申请好STS后，可将STS Token配置到MONGODB Client中，从而实现通过STS Token创建MONGODB Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建MONGODB Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/mongodb" //导入MONGODB服务模块
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

	// 使用申请的临时STS创建MONGODB服务的Client对象，Endpoint使用默认值
	mongodbClient, err := mongodb.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "mongodb.bj.baidubce.com")
	if err != nil {
		fmt.Println("create mongodb client failed:", err)
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
	mongodbClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置MONGODB Client时，无论对应MONGODB服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问MONGODB

MONGODB支持HTTPS传输协议，您可以通过在创建MONGODB Client对象时指定的Endpoint中指明HTTPS的方式，在MONGODB GO SDK中使用HTTPS访问MONGODB服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

ENDPOINT := "https://mongodb.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
mongodbClient, _ := mongodb.NewClient(AK, SK, ENDPOINT)
```

## 配置MONGODB Client

如果用户需要配置MONGODB Client的一些细节的参数，可以在创建MONGODB Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问MONGODB服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

//创建MONGODB Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "mongodb.bj.baidubce.com"
client, _ := mongodb.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "mongodb.bj.baidubce.com"
client, _ := mongodb.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "mongodb.bj.baidubce.com"
client, _ := mongodb.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = headersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问MONGODB时，创建的MONGODB Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见"使用STS创建MONGODB Client"小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 实例管理

## 实例列表

使用以下代码可以查询MONGODB实例列表
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ListMongodbArgs{
    // 分页标记，从第一页开始
    Marker:  "",
    // 每页最大数量，最大1000，默认1000
    MaxKeys: 100,
    // 引擎版本，可选
    EngineVersion: "",
    // 存储引擎，可选
    StorageEngine: "",
    // 实例类型，可选，replica/sharding
    DbInstanceType: "",
}
result, err := client.ListMongodb(args)
if err != nil {
    fmt.Printf("list mongodb error: %+v\n", err)
    return
}
fmt.Println("list mongodb success: ", result)
```

## 查询实例详情

使用以下代码可以查询指定MONGODB实例的详情
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

result, err := client.GetInstanceDetail("instanceId")
if err != nil {
    fmt.Printf("get instance detail error: %+v\n", err)
    return
}
fmt.Println("get instance detail success: ", result)
```

## 创建副本集实例

使用以下代码可以创建一个副本集MONGODB实例
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.CreateReplicaArgs{
    // 客户端Token，幂等性保证，可选
    ClientToken: "your-client-token",
    // 计费信息，必选
    Billing: mongodb.BillingModel{
        PaymentTiming: "Postpaid", // Postpaid后付费，Prepaid预付费
        // 预付费时设置
        //Reservation: mongodb.Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
        //AutoRenew: mongodb.AutoRenewModel{AutoRenewLength: 1, AutoRenewTimeUnit: "Month"},
    },
    // 批量创建实例个数，可选
    PurchaseCount: 1,
    // 实例名称，可选
    DbInstanceName: "my-mongodb-replica",
    // 存储引擎，必选
    StorageEngine: "WiredTiger",
    // 引擎版本，必选
    EngineVersion: "4.0",
    // CPU核数，必选
    DbInstanceCpuCount: 1,
    // 内存大小，单位GB，必选
    DbInstanceMemoryCapacity: 2,
    // 存储大小，单位GB，必选
    DbInstanceStorage: 10,
    // 投票节点数量，必选
    VotingMemberNum: 3,
    // 只读节点数量，可选
    ReadonlyNodeNum: 0,
    // root账号密码，可选
    AccountPassword: "your-password",
    // VPC ID，可选
    VpcId: "vpc-xxxxx",
    // 子网列表，可选
    Subnets: []mongodb.SubnetMap{
        {
            ZoneName: "cn-bj-a",
            SubnetId: "sbn-xxxxx",
        },
    },
    // 标签列表，可选
    Tags: []mongodb.TagModel{
        {
            TagKey:   "tagK",
            TagValue: "tagV",
        },
    },
}
result, err := client.CreateReplica(args)
if err != nil {
    fmt.Printf("create replica error: %+v\n", err)
    return
}
fmt.Println("create replica success: ", result)
```

## 创建分片集实例

使用以下代码可以创建一个分片集MONGODB实例
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.CreateShardingArgs{
    // 客户端Token，幂等性保证，可选
    ClientToken: "your-client-token",
    // 计费信息，必选
    Billing: mongodb.BillingModel{
        PaymentTiming: "Postpaid",
    },
    // 批量创建实例个数，可选
    PurchaseCount: 1,
    // 实例名称，可选
    DbInstanceName: "my-mongodb-sharding",
    // 存储引擎，必选
    StorageEngine: "WiredTiger",
    // 引擎版本，必选
    EngineVersion: "4.0",
    // Mongos数量，可选
    MongosCount: 2,
    // Mongos CPU核数，必选
    MongosCpuCount: 1,
    // Mongos内存大小，可选
    MongosMemoryCapacity: 2,
    // Shard数量，可选
    ShardCount: 2,
    // Shard CPU核数，必选
    ShardCpuCount: 1,
    // Shard内存大小，单位GB，必选
    ShardMemoryCapacity: 2,
    // Shard存储大小，单位GB，必选
    ShardStorage: 10,
    // root账号密码，可选
    AccountPassword: "your-password",
    // VPC ID，可选
    VpcId: "vpc-xxxxx",
    // 子网列表，可选
    Subnets: []mongodb.SubnetMap{
        {
            ZoneName: "cn-bj-a",
            SubnetId: "sbn-xxxxx",
        },
    },
}
result, err := client.CreateSharding(args)
if err != nil {
    fmt.Printf("create sharding error: %+v\n", err)
    return
}
fmt.Println("create sharding success: ", result)
```

## 释放实例（放入回收站）

使用以下代码可以将MONGODB实例放入回收站
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

// 释放单个实例
err := client.ReleaseMongodb("instanceId")

// 批量释放实例
err := client.ReleaseMongodbs([]string{"instanceId1", "instanceId2"})
```

## 从回收站删除实例

使用以下代码可以从回收站删除MONGODB实例
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

// 删除单个实例
err := client.DeleteMongodb("instanceId")

// 批量删除实例
err := client.DeleteMongodbs([]string{"instanceId1", "instanceId2"})
```

## 从回收站恢复实例

使用以下代码可以从回收站恢复MONGODB实例
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

err := client.RecoverMongodbs([]string{"instanceId1", "instanceId2"})
```

## 重启实例

使用以下代码可以重启MONGODB实例
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

// 重启单个实例
err := client.RestartMongodb("instanceId")

// 批量重启实例
err := client.RestartMongodbs([]string{"instanceId1", "instanceId2"})

// 重启分片集组件
err := client.RestartShardingComponent("instanceId", "nodeId")
```

## 修改实例名称

使用以下代码可以修改MONGODB实例名称
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.UpdateInstanceNameArgs{
    DbInstanceName: "new-name",
}
err := client.UpdateInstanceName("instanceId", args)
```

## 修改分片集组件名称

使用以下代码可以修改分片集组件名称
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.UpdateComponentNameArgs{
    NodeName: "new-node-name",
}
err := client.UpdateShardingComponentName("instanceId", "nodeId", args)
```

## 修改密码

使用以下代码可以修改MONGODB实例账号密码
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.UpdatePasswordArgs{
    AccountPassword: "new-password",
}
err := client.UpdateAccountPassword("instanceId", args)
```

## 副本集主从切换

使用以下代码可以对副本集实例进行主从切换
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

err := client.ReplicaSwitch("instanceId")
```

## 分片集组件主从切换

使用以下代码可以对分片集实例组件进行主从切换
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

err := client.ShardingComponentSwitch("instanceId", "nodeId")
```

## 退款实例

使用以下代码可以退款MONGODB实例
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.RefundModel{
    // 退款类型，"ALL"表示全生命周期退订
    RefundType: "ALL",
    // 退款原因，最大4096字符
    RefundReason: "不再使用",
}
result, err := client.RefundInstance("instanceId", args)
```

## 迁移可用区

使用以下代码可以迁移MONGODB实例的可用区
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.MigrateAzoneArgs{
    Subnets: []mongodb.SubnetMap{
        {
            ZoneName: "cn-bj-b",
            SubnetId: "sbn-xxxxx",
        },
    },
    Members: []mongodb.MemberRoleModel{
        {
            SubnetId: "sbn-xxxxx",
            Role:     "Primary",
        },
    },
}
err := client.MigrateAzone("instanceId", args)
```

# 实例变配

## 副本集实例改配

使用以下代码可以对副本集实例进行变配
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ReplicaResizeArgs{
    ClientToken:              "your-client-token",
    DbInstanceCpuCount:       2,
    DbInstanceMemoryCapacity: 4,
    DbInstanceStorage:        20,
}
err := client.ReplicaResize("instanceId", args)
```

## 分片集组件改配

使用以下代码可以对分片集实例的组件进行变配
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ShardingComponentResizeArgs{
    ClientToken:        "your-client-token",
    NodeCpuCount:       2,
    NodeMemoryCapacity: 4,
    NodeStorage:        20,
}
err := client.ShardingComponentResize("instanceId", "nodeId", args)
```

## 分片集新增组件

使用以下代码可以为分片集实例新增组件
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ShardingAddComponentArgs{
    ClientToken:        "your-client-token",
    // 组件类型：sharding/mongos
    NodeType:           "sharding",
    NodeCpuCount:       1,
    NodeMemoryCapacity: 2,
    NodeStorage:        10,
    PurchaseCount:      1,
}
result, err := client.ShardingAddComponent("instanceId", args)
if err != nil {
    fmt.Printf("add component error: %+v\n", err)
    return
}
fmt.Println("add component success, nodeIds: ", result.NodeIds)
```

## 副本集添加只读节点

使用以下代码可以为副本集实例添加只读节点
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ReplicaAddReadonlyNodesArgs{
    ClientToken:     "your-client-token",
    ReadonlyNodeNum: 2,
    Subnet: mongodb.SubnetMap{
        ZoneName: "cn-bj-a",
        SubnetId: "sbn-xxxxx",
    },
}
result, err := client.ReplicaAddReadonlyNodes("instanceId", args)
if err != nil {
    fmt.Printf("add readonly nodes error: %+v\n", err)
    return
}
fmt.Println("add readonly nodes success: ", result.ReadonlyMemberIds)
```

## 获取只读节点列表

使用以下代码可以获取副本集实例的只读节点列表
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

result, err := client.GetReadonlyNodes("instanceId")
if err != nil {
    fmt.Printf("get readonly nodes error: %+v\n", err)
    return
}
fmt.Println("get readonly nodes success: ", result)
```

# 备份管理

## 创建备份

使用以下代码可以创建MONGODB实例的备份
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

result, err := client.CreateBackup("instanceId", "physical", "备份描述")
if err != nil {
    fmt.Printf("create backup error: %+v\n", err)
    return
}
fmt.Println("create backup success, backupId: ", result.BackupId)
```

## 备份列表

使用以下代码可以查询MONGODB实例的备份列表
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ListBackupArgs{
    Marker:  "",
    MaxKeys: 100,
}
result, err := client.ListBackup("instanceId", args)
if err != nil {
    fmt.Printf("list backup error: %+v\n", err)
    return
}
fmt.Println("list backup success: ", result)
```

## 查询备份详情

使用以下代码可以查询指定备份的详情
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

result, err := client.GetBackupDetail("instanceId", "backupId")
if err != nil {
    fmt.Printf("get backup detail error: %+v\n", err)
    return
}
fmt.Println("get backup detail success: ", result)
```

## 修改备份描述

使用以下代码可以修改备份的描述信息
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ModifyBackupDescriptionArgs{
    BackupDescription: "新的备份描述",
}
err := client.ModifyBackupDescription("instanceId", "backupId", args)
```

## 删除备份

使用以下代码可以删除指定的备份
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

err := client.DeleteBackup("instanceId", "backupId")
```

## 查询备份策略

使用以下代码可以查询MONGODB实例的备份策略
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

result, err := client.GetBackupPolicy("instanceId")
if err != nil {
    fmt.Printf("get backup policy error: %+v\n", err)
    return
}
fmt.Println("get backup policy success: ", result)
```

## 修改备份策略

使用以下代码可以修改MONGODB实例的备份策略
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.BackupPolicy{
    AutoBackupEnable:      "on",
    PreferredBackupPeriod: "Monday,Wednesday,Friday",
    PreferredBackupTime:   "02:00:00Z",
    BackupRetentionPeriod: 7,
    EnableIncrementBackup: 1,
    BackupMethod:          "physical",
}
err := client.ModifyBackupPolicy("instanceId", args)
```

# 安全管理

## 查询安全IP列表

使用以下代码可以查询MONGODB实例的安全IP白名单
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

result, err := client.GetSecurityIps("instanceId")
if err != nil {
    fmt.Printf("get security ips error: %+v\n", err)
    return
}
fmt.Println("get security ips success: ", result.SecurityIps)
```

## 添加安全IP

使用以下代码可以为MONGODB实例添加安全IP
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.SecurityIpModel{
    SecurityIps: []string{"192.168.1.1", "192.168.1.2"},
}
err := client.AddSecurityIps("instanceId", args)
```

## 删除安全IP

使用以下代码可以删除MONGODB实例的安全IP
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.SecurityIpModel{
    SecurityIps: []string{"192.168.1.1"},
}
err := client.DeleteSecurityIps("instanceId", args)
```

# 标签管理

## 全量更新标签

使用以下代码可以全量更新（覆盖）MONGODB实例绑定的标签
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

tags := []mongodb.TagModel{
    {TagKey: "env", TagValue: "prod"},
    {TagKey: "team", TagValue: "backend"},
}
err := client.InstanceAssignTags("instanceId", tags)
```

## 追加标签

使用以下代码可以为MONGODB实例追加标签
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

tags := []mongodb.TagModel{
    {TagKey: "app", TagValue: "myapp"},
}
err := client.InstanceBindTags("instanceId", tags)
```

## 解绑标签

使用以下代码可以解绑MONGODB实例的标签
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

tags := []mongodb.TagModel{
    {TagKey: "app", TagValue: "myapp"},
}
err := client.InstanceUnbindTags("instanceId", tags)
```

# 日志管理

## 开启日志采集

使用以下代码可以开启MONGODB实例的日志采集
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.StartLoggingArgs{
    Type: "slowlog", // 日志类型
}
err := client.StartLogging("instanceId", args)
```

## 查询日志文件列表

使用以下代码可以查询MONGODB实例的日志文件列表
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ListLogFilesArgs{
    MemberId:  "memberId",
    Type:      "slowlog",
    StartTime: "2024-01-01T00:00:00Z",
    EndTime:   "2024-01-02T00:00:00Z",
}
result, err := client.ListLogFiles("instanceId", args)
if err != nil {
    fmt.Printf("list log files error: %+v\n", err)
    return
}
fmt.Println("list log files success: ", result)
```

# 权限管理

## 查询用户列表

使用以下代码可以查询MONGODB实例的用户列表
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ListUsersArgs{
    Marker:  "",
    MaxKeys: 100,
}
result, err := client.ListUsers("instanceId", args)
if err != nil {
    fmt.Printf("list users error: %+v\n", err)
    return
}
fmt.Println("list users success: ", result)
```

## 创建用户

使用以下代码可以为MONGODB实例创建用户
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.CreateUserArgs{
    Name:        "myuser",
    Password:    "your-password",
    Description: "用户描述",
    Roles: []mongodb.RoleInfo{
        {
            DbName: "mydb",
            Role:   "readWrite",
        },
    },
}
err := client.CreateUser("instanceId", args)
if err != nil {
    fmt.Printf("create user error: %+v\n", err)
    return
}
fmt.Println("create user success")
```

## 更新用户权限

使用以下代码可以更新MONGODB实例用户的权限
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.UpdateUserRolesArgs{
    Name: "myuser",
    Roles: []mongodb.RoleInfo{
        {
            DbName: "mydb",
            Role:   "read",
        },
    },
}
err := client.UpdateUserRoles("instanceId", args)
if err != nil {
    fmt.Printf("update user roles error: %+v\n", err)
    return
}
fmt.Println("update user roles success")
```

## 删除用户

使用以下代码可以删除MONGODB实例的用户
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.DropUserArgs{
    Name: "myuser",
}
err := client.DropUser("instanceId", args)
if err != nil {
    fmt.Printf("drop user error: %+v\n", err)
    return
}
fmt.Println("drop user success")
```

# 数据库管理

## 查询用户数据库列表

使用以下代码可以查询MONGODB实例的用户数据库列表
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.ListDatabasesArgs{
    Marker:  "",
    MaxKeys: 100,
}
result, err := client.ListDatabases("instanceId", args)
if err != nil {
    fmt.Printf("list databases error: %+v\n", err)
    return
}
fmt.Println("list databases success: ", result)
```

## 创建用户数据库

使用以下代码可以为MONGODB实例创建用户数据库
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.CreateDatabaseArgs{
    Name:           "mydb",
    CollectionName: "mycollection",
    Description:    "数据库描述",
}
err := client.CreateDatabase("instanceId", args)
if err != nil {
    fmt.Printf("create database error: %+v\n", err)
    return
}
fmt.Println("create database success")
```

## 删除用户数据库

使用以下代码可以删除MONGODB实例的用户数据库
```go
// import "github.com/baidubce/bce-sdk-go/services/mongodb"

args := &mongodb.DropDatabaseArgs{
    Name: "mydb",
}
err := client.DropDatabase("instanceId", args)
if err != nil {
    fmt.Printf("drop database error: %+v\n", err)
    return
}
fmt.Println("drop database success")
```


