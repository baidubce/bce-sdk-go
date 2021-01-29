# DDC服务

# 概述

本文档主要介绍DDC GO SDK的使用。在使用本文档前，您需要先了解DDC的一些基本知识，并已开通了DDC服务。若您还不了解DDC，可以参考[产品描述](https://cloud.baidu.com/doc/RDS/s/ujwvyzdzg)和[操作指南](https://cloud.baidu.com/doc/RDS/s/Qjwvz0ikk)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[RDS服务域名](https://cloud.baidu.com/doc/RDS/s/Ejwvz0uoq)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

目前支持“华东-苏州”区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
SU | ddc.su.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云DDC，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问DDC做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建DDC Client

DDC Client是DDC服务的客户端，为开发者与DDC服务进行交互提供了一系列的方法。

### 使用AK/SK新建DDC Client

通过AK/SK方式访问DDC，用户可以参考如下代码新建一个DDC Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/ddc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个DDCClient
	ddcClient, err := ddc.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为DDC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为苏州，则为`ddc.su.baidubce.com`。

### 使用STS创建DDC Client

**申请STS token**

DDC可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问DDC，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建DDC Client**

申请好STS后，可将STS Token配置到DDC Client中，从而实现通过STS Token创建DDC Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建DDC Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/ddc" //导入DDC服务模块
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
	ddcClient, err := ddc.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "ddc.su.baidubce.com")
	if err != nil {
		fmt.Println("create ddc client failed:", err)
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
	ddcClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置DDC Client时，无论对应DDC服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问DDC

DDC支持HTTPS传输协议，您可以通过在创建DDC Client对象时指定的Endpoint中指明HTTPS的方式，在DDC GO SDK中使用HTTPS访问DDC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

ENDPOINT := "https://ddc.su.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
ddcClient, _ := ddc.NewClient(AK, SK, ENDPOINT)
```

## 配置DDC Client

如果用户需要配置DDC Client的一些细节的参数，可以在创建DDC Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问RDS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

//创建DDC Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ddc.su.baidubce.com"
client, _ := ddc.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ddc.su.baidubce.com"
client, _ := ddc.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ddc.su.baidubce.com"
client, _ := ddc.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问DDC时，创建的DDC Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建DDC Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# DDC管理

云数据库专属集群 DDC （Dedicated Database Cluster）是专业、高性能、高可靠的云数据库服务。云数据库 DDC 提供 Web 界面进行配置、操作数据库实例，还为您提供可靠的数据备份和恢复、完备的安全管理、完善的监控、轻松扩展等功能支持。相对于自建数据库，云数据库 DDC 具有更经济、更专业、更高效、更可靠、简单易用等特点，使您能更专注于核心业务。

# 部署集管理

## 创建部署集
使用以下代码可以在指定资源池下创建一个新的部署集。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.CreateDeployRequest{
    // 幂等 Token
    ClientToken: "xxxyyyzzz",
    // 部署集名称
    DeployName: "api-from-go",
    // 部署策略 支持集中部署(centralized)/完全打散(distributed)
    Strategy:   "distributed",
}
err = client.CreateDeploySet(poolId, args)
if err != nil {
    fmt.Printf("create deploy set error: %+v\n", err)
    return
}

fmt.Println("create deploy set success.")
```

## 查询部署集列表
使用以下代码可以查询指定资源池下的部署集列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

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
    fmt.Println("ddc instance ids: ", deploy.Instances)
}
```

## 查询特定部署集信息
使用以下代码可以查询特定部署集的详细信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"
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
fmt.Println("ddc instance ids: ", deploy.Instances)
```

## 删除部署集
使用以下代码可以删除某个资源池下特定的部署集。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

err := DDC_CLIENT.DeleteDeploySet(poolId, deployId)
if err != nil {
    fmt.Printf("delete deploy set error: %+v\n", err)
    return
}
fmt.Printf("delete deploy set success\n")
```

# 账号管理

## 创建账号

使用以下代码可以在某个主实例下创建一个新的账号。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.CreateAccountArgs{
    // 幂等性Token，使用 uuid 生成一个长度不超过64位的ASCII字符串，可选参数
    ClientToken: "xxxyyyzzz",
	// 账号名称，由小写字母、数字、下划线组成、字母开头，字母或数字结尾，最长16个字符，不能为保留关键字，必选
    AccountName: "accountName",
    // 账号的密码，由字母、数字和特殊字符（!@#%^_）中的至少两种组成，长度8-32位，必选
    Password: "password",
    // 账号权限类型，AccountType_Common：普通帐号,AccountType_Super：super账号。必选
    Type: ddc.AccountType_Common,
    // 权限设置，可选
    DatabasePrivileges: []ddc.DatabasePrivilege{
            {
                // 数据库名称
                DbName: "user_photo_001",
                // 授权类型。AuthType_ReadOnly：只读，AuthType_ReadWrite：读写
                AuthType: ddc.AuthType_ReadOnly,
            },   
        },
    // 帐号备注，最多256个字符（一个汉字等于三个字符），可选
    Remark: "账号user1",
}
err = client.CreateAccount(instanceId, args)
if err != nil {
    fmt.Printf("create account error: %+v\n", err)
    return
}

fmt.Println("create account success.")
```

> 注意:
> - 实例状态为available，实例必须是主实例。
> - 没有超出实例最大账号数量。

## 更新账号密码

使用以下代码可以更新账号的密码。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.UpdateAccountPasswordArgs{
    // 密码，由字母、数字和特殊字符（!@#%^_）中的至少两种组成，长度8-32位，必选
    Password: "password",
}
err = client.UpdateAccountPassword(instanceId, accountName, args)
if err != nil {
    fmt.Printf("update account password error: %+v\n", err)
    return
}

fmt.Println("update account password success.")
```

## 更新账号备注

使用以下代码可以更新账号的备注。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.UpdateAccountRemarkArgs{
    // 帐号备注，最多256个字符（一个汉字等于三个字符），可选
    Remark: "remark",
}
err = client.UpdateAccountRemark(instanceId, accountName, args)
if err != nil {
    fmt.Printf("update account remark error: %+v\n", err)
    return
}

fmt.Println("update account remark success.")
```

## 更新账号权限

使用以下代码可以更新账号的权限。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

databasePrivileges := []ddc.DatabasePrivilege{
    {
        DbName:   "hello",
		// 授权类型。AuthType_ReadOnly：只读，AuthType_ReadWrite：读写
        AuthType: ddc.AuthType_ReadWrite,
    },
}

args := &ddc.UpdateAccountPrivilegesArgs{
    DatabasePrivileges: databasePrivileges,
}
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
// import "github.com/baidubce/bce-sdk-go/services/ddc"

result, err := client.GetAccount(instanceId,accountName)
if err != nil {
    fmt.Printf("get account error: %+v\n", err)
    return
}

// 获取account的信息
fmt.Println("ddc accountName: ", result.AccountName)
fmt.Println("ddc remark: ", result.Remark)
// 账号状态（创建中：creating；可用中：available；更新中：updating；删除中：deleting；已删除：deleted）
fmt.Println("ddc accountStatus: ", result.AccountStatus)
// 账号类型（super账号：rdssuper；普通账号：common）
fmt.Println("ddc accountType: ", result.AccountType)
fmt.Println("ddc databasePrivileges: ", result.DatabasePrivileges)
```

## 查询账号列表

使用以下代码可以查询指定实例的账号列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

result, err := client.ListAccount(instanceId)
if err != nil {
    fmt.Printf("list account error: %+v\n", err)
    return
}

// 获取account的列表信息
for _, account := range result.Accounts {
	fmt.Println("ddc accountName: ", account.AccountName)
    fmt.Println("ddc remark: ", account.Remark)
    // 账号状态（创建中：creating；可用中：available；更新中：updating；删除中：deleting；已删除：deleted）
    fmt.Println("ddc accountStatus: ", account.AccountStatus)
    // 账号类型（super账号：rdssuper；普通账号：common）
    fmt.Println("ddc accountType: ", account.AccountType)
    fmt.Println("ddc databasePrivileges: ", account.DatabasePrivileges)
}
```

## 删除特定账号信息

使用以下代码可以删除特定账号信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

err := client.DeleteAccount(instanceId,accountName)
if err != nil {
    fmt.Printf("delete account error: %+v\n", err)
    return
}
fmt.Printf("delete account success\n")
```

# 数据库管理

## 创建数据库

使用以下代码可以在某个主实例下创建一个新的数据库。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.CreateDatabaseArgs{
    // 幂等性Token，使用 uuid 生成一个长度不超过64位的ASCII字符串，可选参数
    ClientToken: "xxxyyyzzz",
    // 数据库名称（由大小写字母、数字、下划线组成、字母开头，字母或数字结尾，最长64个字符）
    DbName: "dbName",
    // 数据库字符集（取值范围：utf8、gbk、latin1、utf8mb4）
    CharacterSetName: "utf8",
    // 数据库备注，最多256个字符（一个汉字等于三个字符）
    Remark: "remark",
    // 权限设置，可选
    AccountPrivileges: []ddc.AccountPrivilege{
            {
                AccountName: "user_photo_001",
                //授权类型。AuthType_ReadOnly：只读，AuthType_ReadWrite：读写
                AuthType: ddc.AuthType_ReadOnly,
            },   
        },
}
err = client.CreateDatabase(instanceId, args)
if err != nil {
    fmt.Printf("create database error: %+v\n", err)
    return
}

fmt.Println("create database success.")
```

> 注意:
> - 实例状态为available，实例必须是主实例。

## 更新数据库备注

使用以下代码可以更新数据库的备注。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.UpdateDatabaseRemarkArgs{
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
// import "github.com/baidubce/bce-sdk-go/services/ddc"

result, err := client.GetDatabase(instanceId,dbName)
if err != nil {
    fmt.Printf("get database error: %+v\n", err)
    return
}

// 获取account的列表信息
fmt.Println("ddc dbName: ", result.DbName)
fmt.Println("ddc characterSetName: ", result.CharacterSetName)
// 数据库状态（创建中：creating；可用中：available；删除中：deleting；已删除：deleted）
fmt.Println("ddc dbStatus: ", result.DbStatus)
fmt.Println("ddc remark: ", result.Remark)
fmt.Println("ddc accountPrivileges: ", result.AccountPrivileges)
```

## 查询数据库列表

使用以下代码可以查询指定实例的账号列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

result, err := client.ListDatabase(instanceId)
if err != nil {
    fmt.Printf("list database error: %+v\n", err)
    return
}

// 获取account的列表信息
for _, database := range result.Databases {
	fmt.Println("ddc dbName: ", database.DbName)
    fmt.Println("ddc characterSetName: ", database.CharacterSetName)
	// 数据库状态（创建中：creating；可用中：available；删除中：deleting；已删除：deleted）
    fmt.Println("ddc dbStatus: ", database.DbStatus)
    fmt.Println("ddc remark: ", database.Remark)
    fmt.Println("ddc accountPrivileges: ", database.AccountPrivileges)
}
```

## 删除特定数据库信息

使用以下代码可以删除特定数据库信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

err := client.DeleteDatabase(instanceId,dbName)
if err != nil {
    fmt.Printf("delete database error: %+v\n", err)
    return
}
fmt.Printf("delete database success\n")
```

# 实例管理

## 创建实例
使用以下代码可以创建一个实例。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.CreateInstanceArgs{
	// 幂等 Token
    ClientToken: "xxxyyyzzz",
    // 实例类型 必选
	InstanceType: "RDS",
	// 创建数量 必选
	Number: 1,
    // 权限设置，可选
    Instance: ddc.Instance{
            // rds实例名称，允许小写字母、数字，中文，长度限制为1~64，可选	
            InstanceName 		 :"mysql_5.7",	
			// 主实例id （创建只读实例时需要）
			SourceInstanceId	 :"sourceInstanceId",
			Engine               :"mysql",
			// 数据库版本 必选
			EngineVersion        :"engineVersion",
			// CPU核数 必选
			CpuCount             :1,
			// 套餐内存大小，单位GB，必选
			AllocatedMemoryInGB  :8,
			//套餐磁盘大小，单位GB，每5G递增，必选
			AllocatedStorageInGB :10,
			// 可用区 必选 调用可用区列表获取
			AZone                :"zoneA",
			// vpc，必选 （只读实例必须和主实例一致）
			VpcId                :"vpcId",
			// vpc内，每个可用区的subnetId；必选
			SubnetId             :"zoneA:71cfc441-fd54-484e-8ea7-192c6649dc18",
			// 磁盘类型 必选
			DiskIoType           :"ssd",
			// 部署集id 必选
			DeployId             :"deployId",
			// 资源池id 必选
			PoolId             :"poolId",
            // RO组ID。(创建只读实例时) 可选
			// 如果不传，默认会创建一个RO组，并将该只读加入RO组中
			RoGroupId             :"roGroupId",
            // RO组是否启用延迟剔除，默认不启动。（创建只读实例时）可选
			EnableDelayOff        :false,
            // 延迟阈值。（创建只读实例时）可选
			DelayThreshold         : 1,
            // RO组最少保留实例数目。默认为1. （创建只读实例时）可选
			LeastInstanceAmount    : 1,
            // 只读实例在RO组中的读流量权重。默认为1（创建只读实例时）可选
			RoGroupWeight          : 1,
        },
}
resp,err := client.CreateInstance(args)
if err != nil {
    fmt.Printf("create instance error: %+v\n", err)
    return
}
fmt.Println(resp)
fmt.Println("create instance success.")
```

## 实例详情

使用以下代码可以查询指定实例的详情。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

resp, err := client.GetDetail(instanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}

// 获取实例详情信息
fmt.Println("ddc dbName: ", resp.Instance.InstanceId)
fmt.Println("ddc instanceName: ", resp.Instance.InstanceName)
fmt.Println("ddc engine: ", resp.Instance.Engine)
fmt.Println("ddc engineVersion: ", resp.Instance.EngineVersion)
fmt.Println("ddc instanceStatus: ", resp.Instance.InstanceStatus)
fmt.Println("ddc cpuCount: ", resp.Instance.CpuCount)
fmt.Println("ddc allocatedMemoryInGB: ", resp.Instance.AllocatedMemoryInGB)
fmt.Println("ddc allocatedStorageInGB: ", resp.Instance.AllocatedStorageInGB)
fmt.Println("ddc nodeAmount: ", resp.Instance.NodeAmount)
fmt.Println("ddc usedStorageInGB: ", resp.Instance.UsedStorageInGB)
fmt.Println("ddc publicAccessStatus: ", resp.Instance.PublicAccessStatus)
fmt.Println("ddc endpoint: ", resp.Instance.Endpoint)
fmt.Println("ddc backupPolicy: ", resp.Instance.BackupPolicy)
fmt.Println("ddc region: ", resp.Instance.Region)
fmt.Println("ddc instanceType: ", resp.Instance.InstanceType)
fmt.Println("ddc sourceInstanceId: ", resp.Instance.SourceInstanceId)
fmt.Println("ddc zoneNames: ", resp.Instance.ZoneNames)
fmt.Println("ddc vpcId: ", resp.Instance.VpcId)
fmt.Println("ddc subnets: ", resp.Instance.Subnets)
fmt.Println("ddc topology: ", resp.Instance.Topology)
fmt.Println("ddc diskType: ", resp.Instance.DiskType)
fmt.Println("ddc type: ", resp.Instance.Type)
fmt.Println("ddc applicationType: ", resp.Instance.ApplicationType)
fmt.Println("ddc roGroupList: ", resp.Instance.RoGroupList)
fmt.Println("ddc nodeMaster: ", resp.Instance.NodeMaster)
fmt.Println("ddc nodeSlave: ", resp.Instance.NodeSlave)
```

## 实例列表
使用以下代码可以查询实例列表信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

resp, err := client.ListDdcInstance(instanceId)

args := &ddc.ListRdsArgs{
    // 批量获取列表的查询的起始位置，是一个由系统生成的字符串，可选
    Marker: "marker",
    // 指定每页包含的最大数量(主实例)，最大数量不超过1000，缺省值为1000，可选
    MaxKeys: 1,
}

if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}

// 返回标记查询的起始位置
fmt.Println("rds list marker: ", resp.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("rds list isTruncated: ", resp.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("rds list nextMarker: ", resp.NextMarker)
// 每页包含的最大数量
fmt.Println("rds list maxKeys: ", resp.MaxKeys)

// 获取account的列表信息
for _, e := range resp.Result {
	fmt.Println("ddc instanceId: ", e.InstanceId)
    fmt.Println("ddc instanceName: ", e.InstanceName)
    fmt.Println("ddc engine: ", e.Engine)
    fmt.Println("ddc engineVersion: ", e.EngineVersion)
    fmt.Println("ddc instanceStatus: ", e.InstanceStatus)
    fmt.Println("ddc cpuCount: ", e.CpuCount)
    fmt.Println("ddc allocatedMemoryInGB: ", e.AllocatedMemoryInGB)
    fmt.Println("ddc allocatedStorageInGB: ", e.AllocatedStorageInGB)
    fmt.Println("ddc nodeAmount: ", e.NodeAmount)
    fmt.Println("ddc usedStorageInGB: ", e.UsedStorageInGB)
    fmt.Println("ddc publicAccessStatus: ", e.PublicAccessStatus)
    fmt.Println("ddc endpoint: ", e.Endpoint)
    fmt.Println("ddc backupPolicy: ", e.BackupPolicy)
    fmt.Println("ddc region: ", e.Region)
    fmt.Println("ddc instanceType: ", e.InstanceType)
    fmt.Println("ddc sourceInstanceId: ", e.SourceInstanceId)
    fmt.Println("ddc zoneNames: ", e.ZoneNames)
    fmt.Println("ddc vpcId: ", e.VpcId)
    fmt.Println("ddc subnets: ", e.Subnets)
    fmt.Println("ddc topology: ", e.Topology)
    fmt.Println("ddc diskType: ", e.DiskType)
    fmt.Println("ddc type: ", e.Type)
    fmt.Println("ddc applicationType: ", e.ApplicationType)
}
```

## 删除实例
使用以下代码可以批量删除实例。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

// 多个实例间用英文半角逗号","隔开，最多可输入10个
err := client.DeleteDdcInstance(instanceIds)
if err != nil {
    fmt.Printf("delete instance error: %+v\n", err)
    return
}
fmt.Printf("delete instance success\n")
```

## 修改实例名称
使用以下代码可以修改实例名称。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.UpdateInstanceNameArgs{
	// rds实例名称，允许小写字母、数字，中文，长度限制为1~64
	InstanceName: "instanceName",
}
err := client.UpdateInstanceName(instanceId, args)
if err != nil {
    fmt.Printf("update instance name error: %+v\n", err)
    return
}
fmt.Printf("update instance name success\n")
```

## 主备切换
使用以下代码主备切换。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

err := client.SwitchInstance(instanceId)
if err != nil {
    fmt.Printf(" main standby switching of the instance error: %+v\n", err)
    return
}
fmt.Printf(" main standby switching of the instance success\n")
```
## 只读组列表
使用以下代码可以查询只读组列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

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

## VPC列表
使用以下代码可以查询vpc列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

resp, err := client.ListVpc()
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
// import "github.com/baidubce/bce-sdk-go/services/rds"
resp, err = client.GetZoneList()
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
```

## 子网列表
使用以下代码可以获取一个实例下的子网列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
resp, err := client.ListSubnets()
if err != nil {
	fmt.Printf("get subnet list error: %+v\n", err)
	return
}
for _, e := range resp.Subnets {
    fmt.Println("ddc name: ", e.Name)
    fmt.Println("ddc subnetId: ", e.SubnetId)
    fmt.Println("ddc zoneName: ", e.ZoneName)
    fmt.Println("ddc shortId: ", e.ShortId)
    fmt.Println("ddc vpcId: ", e.VpcId)
    fmt.Println("ddc vpcShortId: ", e.VpcShortId)
    fmt.Println("ddc az: ", e.Az)
    fmt.Println("ddc cidr: ", e.Cidr)
    fmt.Println("ddc createdTime: ", e.CreatedTime)
    fmt.Println("ddc updatedTime: ", e.UpdatedTime)
}
```
# 参数管理

## 实例参数列表
使用以下代码可以查询实例参数列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

resp, err := client.ListParameters(instanceId)
if err != nil {
    fmt.Printf("get instance error: %+v\n", err)
    return
}
// 获取参数列表信息
for _, e := range resp.Items {
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
```

## 修改实例参数
使用以下代码可以云数据库 RDS for MySQL 的参数配置。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.UpdateParameterArgs{
	Parameters:  []ddc.KVParameter{
		{		
			Name: "connect_timeout",
			Value: "15",
		},
	},
}
err := client.UpdateParameter(instanceId, args)
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
// import "github.com/baidubce/bce-sdk-go/services/ddc"

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
// import "github.com/baidubce/bce-sdk-go/services/rds"

args := &ddc.UpdateSecurityIpsArgs{
SecurityIps:  []string{
		"%",
		"192.0.0.1",
		"192.0.0.2",
	},
}
er := client.UpdateSecurityIps(instanceId, args)
if er != nil {
	fmt.Printf("update securityIp list error: %+v\n", er)
	return
}
fmt.Printf("update securityIp list success\n")
```

# 备份管理
## 	获取备份列表
使用以下代码可以获取一个实例下的备份列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
resp, err := client.GetBackupList(instanceId)
if err != nil {
	fmt.Printf("get backup list error: %+v\n", err)
	return
}
// 获取参数列表信息
for _, e := range resp.Snapshots {
    fmt.Println("ddc snapshotId: ", e.SnapshotId)
    fmt.Println("ddc snapshotSizeInBytes: ", e.SnapshotSizeInBytes)
    fmt.Println("ddc snapshotType: ", e.SnapshotType)
    fmt.Println("ddc snapshotStatus: ", e.SnapshotStatus)
    fmt.Println("ddc snapshotStartTime: ", e.SnapshotStartTime)
    fmt.Println("ddc snapshotEndTime: ", e.SnapshotEndTime)
}
```

## 创建备份
使用以下代码创建实例备份。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

err := client.CreateBackup(instanceId)
if err != nil {
    fmt.Printf("create backup error: %+v\n", err)
    return
}
fmt.Printf("create backup success\n")
```

## 备份详情
使用以下代码可以查询一个备份的详情。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

resp, err := client.GetBackupDetail(instanceId, snapshotId)
if err != nil {
    fmt.Printf("get backup detail error: %+v\n", err)
    return
}

fmt.Println("ddc dbName: ", resp.Snapshot.SnapshotId)
fmt.Println("ddc instanceName: ", resp.Snapshot.SnapshotSizeInBytes)
fmt.Println("ddc engine: ", resp.Snapshot.SnapshotType)
fmt.Println("ddc engineVersion: ", resp.Snapshot.SnapshotStatus)
fmt.Println("ddc instanceStatus: ", resp.Snapshot.SnapshotStartTime)
fmt.Println("ddc cpuCount: ", resp.Snapshot.SnapshotEndTime)
fmt.Println("ddc downloadUrl: ", resp.Snapshot.DownloadUrl)
fmt.Println("ddc downloadExpires: ", resp.Snapshot.DownloadExpires)
```

## 设置备份
使用以下代码设置实例的备份策略。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

args := &ddc.BackupPolicy{
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
使用以下代码可以获取一个实例下的binlog列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/rds"
resp, err := client.GetBinlogList(instanceId)
if err != nil {
	fmt.Printf("get binlog list error: %+v\n", err)
	return
}
// 获取参数列表信息
for _, e := range resp.Binlogs {
    fmt.Println("ddc binlogId: ", e.BinlogId)
    fmt.Println("ddc binlogSizeInBytes: ", e.BinlogSizeInBytes)
    fmt.Println("ddc binlogStatus: ", e.BinlogStatus)
    fmt.Println("ddc binlogStartTime: ", e.BinlogStartTime)
    fmt.Println("ddc binlogEndTime: ", e.BinlogEndTime)
}
```

## binlog 详情
使用以下代码可以查询一个binlog详情。
```go
// import "github.com/baidubce/bce-sdk-go/services/ddc"

resp, err := client.GetBinlogDetail(instanceId, binlog)
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


# 错误处理

GO语言以error类型标识错误，DDC支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | RDS服务返回的错误

用户使用SDK调用DDC相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// ddcClient 为已创建的DDC Client对象
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
log.Debugf("%s", "logging message using the log package in the RDS go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the RDS go sdk")
```



首次发布:

 - 支持创建账号、更新账号密码、更新账号备注、更新账号权限、查询账号列表、查询特定账号信息、删除特定账号信息
- 支持创建数据库、更新数据库备注、查询数据库列表、查询特定数据库信息、删除特定数据库信息