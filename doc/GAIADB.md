# GAIADB服务

# 概述

本文档主要介绍GAIADB GO SDK的使用。在使用本文档前，您需要先了解GAIADB的一些基本知识。若您还不了解GAIADB，可以参考[产品描述](https://cloud.baidu.com/doc/GaiaDB/s/mkd45c3ap)和[入门指南](https://cloud.baidu.com/doc/GaiaDB/s/vkgkfur5q)。

相关参数说明可参考官方API文档[API参考](https://cloud.baidu.com/doc/GaiaDB/s/Zl82ton8n)

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[GAIADB访问域名](https://cloud.baidu.com/doc/GaiaDB/s/al84889rz)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/vkgkfur5q/)。

## 获取密钥

要使用百度云SCS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问SCS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建GAIADB Client

GAIADB Client是GAIADB服务的客户端，为开发者与GAIADB服务进行交互提供了一系列的方法。

### 使用AK/SK新建GAIADB Client

通过AK/SK方式访问GAIADB，用户可以参考如下代码新建一个GAIADB Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/gaiadb"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个GAIADBClient
	gaiadbClient, err := scs.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”。第三个参数`ENDPOINT`支持用户自己指定域名.

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`gaiadb.bj.baidubce.com`。

### 使用STS创建GAIADB Client

**申请STS token**

GAIADB可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问GAIADB，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建GAIADB Client**

申请好STS后，可将STS Token配置到GAIADB Client中，从而实现通过STS Token创建GAIADB Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建GAIADB Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/gaiadb" //导入GAIADB服务模块
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

	// 使用申请的临时STS创建GAIADB服务的Client对象，Endpoint使用默认值
	gaiadbClient, err := gaiadb.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create GAIADB client failed:", err)
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
	gaiadbClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置GAIADB Client时，无论对应GAIADB服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问GAIADB

GAIADB支持HTTPS传输协议，您可以通过在创建GAIADB Client对象时指定的Endpoint中指明HTTPS的方式，在GAIADB GO SDK中使用HTTPS访问SCS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/gaiadb"

ENDPOINT := "https://gaiadb.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
gaiadbClient, _ := gaiadb.NewClient(AK, SK, ENDPOINT)
```

## 配置GAIADB Client

如果用户需要配置GAIADB Client的一些细节的参数，可以在创建GAIADB Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问SCS服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/gaiadb"

//创建GAIADB Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "gaiadb.bj.baidubce.com"
client, _ := gaiadb.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/gaiadb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "gaiadb.bj.baidubce.com"
client, _ := gaiadb.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/gaiadb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "gaiadb.bj.baidubce.com"
client, _ := gaiadb.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问GAIADB时，创建的GAIADB Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建GAIADB Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口


## 集群管理

### 创建集群
使用以下代码可以创建Gaiadb集群
```go
args := &gaiadb.CreateClusterArgs{
    ClientToken: getClientToken(),
    Number:      1,
    ProductType: "postpay",
    InstanceParam: InstanceParam{
        EngineVersion:        "8.0",
        SubnetId:             "sbn-na4tmg4v11hs",
        AllocatedCpuInCore:   2,
        AllocatedMemoryInMB:  8192,
        AllocatedStorageInGB: 5120,
        Engine:               "MySQL",
        VpcId:                "vpc-it3v6qt3jhvj",
        InstanceAmount:       2,
        ProxyAmount:          2,
    },
}
result, err := GAIADB_CLIENT.CreateCluster(args)

if err != nil {
    fmt.Println("create cluster failed:", err)
} else {
    fmt.Println("create cluster success: ", result)
}
```
### 删除集群
使用以下代码可以删除Gaiadb集群
```go
err := GAIADB_CLIENT.DeleteCluster(clusterId)

if err != nil {
    fmt.Println("delete cluster failed:", err)
} else {
    fmt.Println("delete cluster success")
}
```
### 修改集群名称
使用以下代码可以修改集群名称
```go
args := &gaiadb.ClusterName{
	ClusterName: "cluster_test",
}
err := GAIADB_CLIENT.RenameCluster(clusterId, args)

if err != nil {
    fmt.Println("rename cluster failed:", err)
} else {
    fmt.Println("rename cluster success")
}
```

### 变配集群
使用以下代码可以变配Gaiadb集群
```go
args := &gaiadb.ResizeClusterArgs{
	ResizeType:          "resizeSlave",
	AllocatedCpuInCore:  2,
	AllocatedMemoryInMB: 8192,
}
result, err := GAIADB_CLIENT.ResizeCluster(clusterId, args)

if err != nil {
    fmt.Println("resize cluster failed:", err)
} else {
    fmt.Println("resize cluster success: ", result)
}
```

### 获取集群列表
使用以下代码可以获取集群列表
```go
args := &gaiadb.Marker{
	Marker:  "-1",
	MaxKeys: 1000,
}
result, err := GAIADB_CLIENT.GetClusterList(args)
re, _ := json.Marshal(result)
fmt.Println(string(re))

if err != nil {
    fmt.Println("get cluster list failed:", err)
} else {
    fmt.Println("get cluster list success: ", result)
}
```
### 获取集群详情
使用以下代码可以获取集群详情
```go
result, err := GAIADB_CLIENT.GetClusterDetail(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))

if err != nil {
    fmt.Println("get cluster detail failed:", err)
} else {
    fmt.Println("get cluster detail success: ", result)
}
```
### 查询集群存储容量
使用以下代码可以查询集群存储容量
```go
result, err := GAIADB_CLIENT.GetClusterCapacity(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))

if err != nil {
    fmt.Println("get cluster capacity failed:", err)
} else {
    fmt.Println("get cluster capacity success: ", result)
}
```

### 查询新购集群价格
使用以下代码可以查询新购集群价格
```go
args := &gaiadb.QueryPriceArgs{
	Number: 1,
	InstanceParam: InstanceInfo{
		ReleaseVersion:       "5.7",
		AllocatedCpuInCore:   2,
		AllocatedMemoryInMB:  8192,
		AllocatedStorageInGB: 5120,
		InstanceAmount:       2,
		ProxyAmount:          2,
	},
	ProductType: "postpay",
}
result, err := GAIADB_CLIENT.QueryClusterPrice(args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get cluster price failed:", err)
} else {
    fmt.Println("get cluster price success: ", result)
}
```

### 查询变配集群价格
使用以下代码可以查询变配集群价格
```go
args := &gaiadb.QueryResizePriceArgs{
	ClusterId:           clusterId,
	ResizeType:          "resizeSlave",
	AllocatedCpuInCore:  2,
	AllocatedMemoryInMB: 8192,
}
result, err := GAIADB_CLIENT.QueryResizeClusterPrice(args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get resize cluster price failed:", err)
} else {
    fmt.Println("get resize cluster price success: ", result)
}
```

### 重启计算节点
使用以下代码可以重启计算节点
```go
args := &gaiadb.RebootInstanceArgs{
	ExecuteAction: "executeNow",
}
err := GAIADB_CLIENT.RebootInstance(clusterId, instanceId, args)
if err != nil {
    fmt.Println("reboot instance failed:", err)
} else {
    fmt.Println("reboot instance success ")
}
```
### 绑定标签
使用以下代码可以绑定标签
```go
args := &gaiadb.BindTagsArgs{
	Resources: []Resource{
		{
			ResourceId: clusterId,
			Tags: []Tag{
				{
					TagKey:   "testTagKey",
					TagValue: "testTagValue",
				},
			},
		},
	},
}
err := GAIADB_CLIENT.BindTags(args)
if err != nil {
    fmt.Println("bind tag failed:", err)
} else {
    fmt.Println("bind tag success ")
}
```

### 主从切换
使用以下代码可以主从切换
```go
args := &gaiadb.ClusterSwitchArgs{
	ExecuteAction:       "executeNow",
	SecondaryInstanceId: instanceId,
}
result, err := GAIADB_CLIENT.ClusterSwitch(clusterId, args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("cluster switch failed:", err)
} else {
    fmt.Println("cluster switch success: ", result)
}
```
## 入口管理
### 查询入口列表
使用以下代码查询入口列表
```go
result, err := GAIADB_CLIENT.GetInterfaceList(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get interface list failed:", err)
} else {
    fmt.Println("get interface list success: ", result)
}
```
### 更新入口域名
使用以下代码更新入口域名
```go
args := &gaiadb.UpdateDnsNameArgs{
	InterfaceId: "gaiadbm5h6ys_interface0000",
	DnsName:     "my.gaiadb.bj.baidubce.com",
}
err := GAIADB_CLIENT.UpdateDnsName(clusterId, args)
if err != nil {
    fmt.Println("update dns name failed:", err)
} else {
    fmt.Println("update dns name success")
}
```
### 更新入口配置
使用以下代码更新入口配置
```go
args := &gaiadb.UpdateInterfaceArgs{
	InterfaceId: "gaiadbm5h6ys_interface0000",
	Interface: InterfaceInfo{
		MasterReadable: 1,
		AddressName:    "addressname",
		InstanceBinding: []string{
			"gaiadbymbrc8-primary-6f1cc3a2",
			"gaiadbymbrc8-secondary-ec909467",
		},
	},
}
err := GAIADB_CLIENT.UpdateInterface(clusterId, args)
if err != nil {
    fmt.Println("update interface failed:", err)
} else {
    fmt.Println("update interface success")
}
```
### 更新入口新节点自动加入配置
使用以下代码更新入口新节点自动加入配置
```go
args := &gaiadb.NewInstanceAutoJoinArgs{
	AutoJoinRequestItems: []AutoJoinRequestItem{
		{
			NewInstanceAutoJoin: "off",
			InterfaceId:         "gaiadbymbrc8-primary-6f1cc3a2",
		},
	},
}
err := GAIADB_CLIENT.NewInstanceAutoJoin(clusterId, args)
if err != nil {
    fmt.Println("new instance audo join failed:", err)
} else {
    fmt.Println("new instance audo join success")
}
```
## 账号管理
### 创建账号
使用以下代码创建账号
```go
args := &gaiadb.CreateAccountArgs{
	AccountName: "testaccount",
	Password:    "baidu@123",
	AccountType: "common",
	Remark:      "testRemark",
}
err := GAIADB_CLIENT.CreateAccount(clusterId, args)
if err != nil {
    fmt.Println("create account failed:", err)
} else {
    fmt.Println("create account success")
}
```
### 删除账号
使用以下代码删除账号
```go
err := GAIADB_CLIENT.DeleteAccount(clusterId, "testaccount")
if err != nil {
    fmt.Println("delete account failed:", err)
} else {
    fmt.Println("delete account success")
}
```
### 查询账号详情
使用以下代码查询账号详情
```go
result, err := GAIADB_CLIENT.GetAccountDetail(clusterId, "testaccount")
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get account detail failed:", err)
} else {
    fmt.Println("get account detail success")
}
```
### 查询账号列表
使用以下代码查询账号列表
```go
result, err := GAIADB_CLIENT.GetAccountList(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get account list failed:", err)
} else {
    fmt.Println("get account list success")
}
```


### 更新账号备注
使用以下代码更新账号备注
```go
args := &gaiadb.RemarkArgs{
	Remark: "remark",
	Etag:   "v0",
}
result, err := GAIADB_CLIENT.UpdateAccountRemark(clusterId, accountName, args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("update account remark failed:", err)
} else {
    fmt.Println("update account remark success")
}
```

### 更新账号白名单
使用以下代码更新账号白名单
```go
args := &gaiadb.AuthIpArgs{
	Action: "ipAdd",
	Value: AuthIp{
		Authip:  []string{"10.10.10.10"},
		Authbns: []string{},
	},
}
err := GAIADB_CLIENT.UpdateAccountAuthIp(clusterId, "testaccount", args)
if err != nil {
    fmt.Println("update account auth ip failed:", err)
} else {
    fmt.Println("update account auth ip success")
}
```

### 更新账号权限
使用以下代码更新账号权限
```go
args := &gaiadb.PrivilegesArgs{
	DatabasePrivileges: []DatabasePrivilege{
		{
			DbName:     "testdb",
			AuthType:   "definePrivilege",
			Privileges: []string{"UPDATE"},
		},
	},
	Etag: "v0",
}
err := GAIADB_CLIENT.UpdateAccountPrivileges(clusterId, "testaccount", args)
if err != nil {
    fmt.Println("update account privileges failed:", err)
} else {
    fmt.Println("update account privileges success")
}
```
### 更新账号密码
使用以下代码更新账号密码
```go
args := &gaiadb.PasswordArgs{
	Password: "testpassword",
	Etag:     "v0",
}
err := GAIADB_CLIENT.UpdateAccountPassword(clusterId, "testaccount", args)
if err != nil {
    fmt.Println("update account password failed:", err)
} else {
    fmt.Println("update account password success")
}
```
## 数据库管理
### 创建数据库
使用以下代码创建数据库
```go
args := &gaiadb.CreateDatabaseArgs{
	DbName:           "test_db",
	CharacterSetName: "utf8",
	Remark:           "sdk test",
}
err := GAIADB_CLIENT.CreateDatabase(clusterId, args)
if err != nil {
    fmt.Println("create database failed:", err)
} else {
    fmt.Println("create database success")
}
```
### 删除数据库
使用以下代码删除数据库
```go
err := GAIADB_CLIENT.DeleteDatabase(clusterId, "test_db")
if err != nil {
    fmt.Println("delete database failed:", err)
} else {
    fmt.Println("delete database success")
}
```
### 查看数据库列表
使用以下代码查看数据库列表
```go
result, err := GAIADB_CLIENT.ListDatabase(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("list database failed:", err)
} else {
    fmt.Println("list database success")
}
```
## 备份管理
### 创建备份
使用以下代码创建备份
```go
err := GAIADB_CLIENT.CreateSnapshot(clusterId)
if err != nil {
    fmt.Println("create snapshot failed:", err)
} else {
    fmt.Println("create snapshot success")
}
```

### 查询备份列表
使用以下代码查询备份列表
```go
result, err := GAIADB_CLIENT.ListSnapshot(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get snapshot list failed:", err)
} else {
    fmt.Println("get snapshot list success")
}
```
### 更新备份策略
使用以下代码更新备份策略
```go
args := &gaiadb.UpdateSnapshotPolicyArgs{
	DataBackupWeekDay: []string{"Monday"},
	DataBackupRetainStrategys: []DataBackupRetainStrategy{{
		StartSeconds: 0,
		RetainCount:  "8",
		Precision:    86400,
		EndSeconds:   -691200,
	}},
	DataBackupTime: "02:00:00Z",
}
err := GAIADB_CLIENT.UpdateSnapshotPolicy(clusterId, args)
if err != nil {
    fmt.Println("update snapshot policy failed:", err)
} else {
    fmt.Println("update snapshot policy success")
}
```
## 安全管理
### 更新 IP 白名单
使用以下代码更新 IP 白名单
```go
args := &gaiadb.UpdateWhiteListArgs{
	AuthIps: []string{"192.168.1.2"},
	Etag:    "v0",
}
err := GAIADB_CLIENT.UpdateWhiteList(clusterId, args)
if err != nil {
    fmt.Println("update white list failed:", err)
} else {
    fmt.Println("update white list success")
}
```
### 查询 IP 白名单
使用以下代码查询 IP 白名单
```go
result, err := GAIADB_CLIENT.GetWhiteList(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get white list failed:", err)
} else {
    fmt.Println("get white list success: ", result)
}
```
## 热活集群组管理
### 创建热活集群组
使用以下代码创建热活集群组
```go
args := &gaiadb.CreateMultiactiveGroupArgs{
	LeaderClusterId:      clusterId,
	MultiActiveGroupName: "test_multiactive_group",
}
result, err := GAIADB_CLIENT.CreateMultiactiveGroup(args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("create multiactive group failed:", err)
} else {
    fmt.Println("create multiactive group success: ", result)
}
```
### 删除热活集群组
使用以下代码删除热活集群组
```go
err := GAIADB_CLIENT.DeleteMultiactiveGroup(groupId)
if err != nil {
    fmt.Println("delete multiactive group failed:", err)
} else {
    fmt.Println("delete multiactive group success: ", result)
}
```
### 更新热活集群组名称
使用以下代码更新热活集群组名称
```go
args := &gaiadb.RenameMultiactiveGroupArgs{
	MultiActiveGroupName: "test_multiactive_group",
}
err := GAIADB_CLIENT.RenameMultiactiveGroup(groupId, args)
if err != nil {
    fmt.Println("rename multiactive group failed:", err)
} else {
    fmt.Println("rename multiactive group success: ", result)
}
```
### 查询热活集群组列表
使用以下代码查询热活集群组列表
```go
result, err := GAIADB_CLIENT.MultiactiveGroupList()
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("list multiactive group failed:", err)
} else {
    fmt.Println("list multiactive group success: ", result)
}
```
### 查询热活集群组详情
使用以下代码查询热活集群组详情
```go
result, err := GAIADB_CLIENT.MultiactiveGroupDetail("gaiagroup-0luzwo")
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get multiactive group detail failed:", err)
} else {
    fmt.Println("get multiactive group detail success: ", result)
}
```
### 查询从集群延迟信息
使用以下代码查询从集群延迟信息
```go
result, err := GAIADB_CLIENT.GetSyncStatus("gaiagroup-0luzwo", clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get sync status failed:", err)
} else {
    fmt.Println("get sync status success: ", result)
}
```
### 主从切换
使用以下代码主从切换
```go
args := &gaiadb.ExchangeArgs{
	ExecuteAction:      "executeNow",
	NewLeaderClusterId: clusterId,
}
err := GAIADB_CLIENT.GroupExchange("gaiagroup-0luzwo", args)
if err != nil {
    fmt.Println("exchange failed:", err)
} else {
    fmt.Println("exchange success: ", result)
}
```
## 参数模板管理
### 查询参数列表
使用以下代码查询参数列表
```go
result, err := GAIADB_CLIENT.GetParamsList(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get params list failed:", err)
} else {
    fmt.Println("get params list success: ", result)
}
```
### 查询参数更新历史
使用以下代码查询参数更新历史
```go
result, err := GAIADB_CLIENT.GetParamsHistory(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get params history failed:", err)
} else {
    fmt.Println("get params history success: ", result)
}
```
### 更新参数
使用以下代码更新参数
```go
args := &gaiadb.UpdateParamsArgs{
	Params: map[string]interface{}{
		"auto_increment_increment": "5",
	},
	Timing: "now",
}
err := GAIADB_CLIENT.UpdateParams(clusterId, args)
if err != nil {
    fmt.Println("update params failed:", err)
} else {
    fmt.Println("update params success: ", result)
}
```
## 参数模板管理
### 查询参数模板列表
使用以下代码查询参数模板列表
```go
args := &gaiadb.ListParamTempArgs{
	Detail:   0,
	Type:     "mysql",
	PageNo:   1,
	PageSize: 10,
}
result, err := GAIADB_CLIENT.ListParamTemplate(args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("list param template failed:", err)
} else {
    fmt.Println("list param template success: ", result)
}
```
### 保存为参数模板
使用以下代码保存为参数模板
```go
args := &gaiadb.ParamTempArgs{
	Type:        "mysql",
	Version:     "8.0",
	Description: "create by sdk",
	Name:        "sdk_test",
	Source:      clusterId,
}
err := GAIADB_CLIENT.SaveAsParamTemplate(args)
if err != nil {
    fmt.Println("save as template failed:", err)
} else {
    fmt.Println("save as template success: ", result)
}
```
### 查询参数模板应用记录
使用以下代码查询参数模板应用记录
```go
result, err := GAIADB_CLIENT.GetTemplateApplyRecords(templateId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get apply record failed:", err)
} else {
    fmt.Println("get apply record success: ", result)
}
```
### 删除参数模板中的参数
使用以下代码删除参数模板中的参数
```go
args := &gaiadb.Params{
	Params: []string{"long_query_time"},
}
err := GAIADB_CLIENT.DeleteParamsFromTemp(templateId, args)
if err != nil {
    fmt.Println("delete params from template failed:", err)
} else {
    fmt.Println("delete params from template success: ", result)
}
```
### 更新参数模板
使用以下代码更新参数模板
```go
args := &gaiadb.UpdateParamTplArgs{
	Name:        "test_template",
	Description: "test_template_description",
}
err := GAIADB_CLIENT.UpdateParamTemplate(templateId, args)
	
if err != nil {
    fmt.Println("update params template failed:", err)
} else {
    fmt.Println("update params template success: ", result)
}
```
### 修改参数模板中的参数
使用以下代码修改参数模板中的参数
```go
args := &ModifyParamsArgs{
	Params: map[string]interface{}{
		"auto_increment_increment": "5",
		"long_query_time":          "6.6",
	},
}
err := GAIADB_CLIENT.ModifyParams(templateId, args)
	
if err != nil {
    fmt.Println("modify params failed:", err)
} else {
    fmt.Println("modify params success: ", result)
}
```
### 删除参数模板
使用以下代码删除参数模板
```go
err := GAIADB_CLIENT.DeleteParamTemplate(templateId)	
if err != nil {
    fmt.Println("delete param template failed:", err)
} else {
    fmt.Println("delete param template success: ", result)
}
```
### 创建参数模板
使用以下代码创建参数模板
```go
args := &gaiadb.CreateParamTemplateArgs{
	Name:        "test_template",
	Type:        "mysql",
	Version:     "8.0",
	Description: "test_template_description",
}
err := GAIADB_CLIENT.CreateParamTemplate(args)	
if err != nil {
    fmt.Println("create param template failed:", err)
} else {
    fmt.Println("create param template success: ", result)
}
```
### 查询参数模板详情
使用以下代码查询参数模板详情
```go
result, err := GAIADB_CLIENT.GetParamTemplateDetail(templateId, "0")
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get param template detail failed:", err)
} else {
    fmt.Println("get param template detail success: ", result)
}
```
### 查询参数模板更新历史
使用以下代码查询参数模板更新历史
```go
result, err := GAIADB_CLIENT.GetParamTemplateHistory(templateId, "addParam")
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get param template records failed:", err)
} else {
    fmt.Println("get param template records success: ", result)
}
```
### 应用参数模板
使用以下代码应用参数模板
```go
args := &gaiadb.ApplyParamTemplateArgs{
	Timing: "now",
	Clusters: map[string]interface{}{
		"gaiadbk3pyxv": []interface{}{},
	},
}
err := GAIADB_CLIENT.ApplyParamTemplate(templateId, args)
if err != nil {
    fmt.Println("apply param template failed:", err)
} else {
    fmt.Println("apply param template success: ", result)
}
```
## 维护时间窗口管理
### 更新时间窗口
使用以下代码更新时间窗口
```go
args := &gaiadb.UpdateMaintenTimeArgs{
	Period:    "1,2,3",
	StartTime: "03:00",
	Duration:  1,
}
err := GAIADB_CLIENT.UpdateMaintenTime(clusterId, args)
if err != nil {
    fmt.Println("update mainten time failed:", err)
} else {
    fmt.Println("update mainten time success: ", result)
}
```
### 查询时间窗口详情
使用以下代码查询时间窗口详情
```go
result, err := GAIADB_CLIENT.GetMaintenTime(clusterId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get mainten time detail failed:", err)
} else {
    fmt.Println("get mainten time detail success: ", result)
}
```
## 慢日志管理
### 查询慢sql详情
使用以下代码查询慢sql详情
```go
args := &gaiadb.GetSlowSqlArgs{
	Page:     "1",
	PageSize: "10",
}
result, err := GAIADB_CLIENT.GetSlowSqlDetail(clusterId, args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get slow sql detail failed:", err)
} else {
    fmt.Println("get slow sql detail success: ", result)
}
```
### 查询慢sql优化建议
使用以下代码查询慢sql优化建议
```go
result, err := GAIADB_CLIENT.SlowSqlAdvice(clusterId, sqlId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get slow sql advice failed:", err)
} else {
    fmt.Println("get slow sql advice success: ", result)
}
```
## Binlog管理
### 查询 binlog 备份详情
使用以下代码查询 binlog 备份详情
```go
args := &gaiadb.GetBinlogArgs{
	AppId:         clusterId,
	LogBackupType: "logical",
}
result, err := GAIADB_CLIENT.GetBinlogDetail(logId, args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get binlog detail failed:", err)
} else {
    fmt.Println("get binlog detail success: ", result)
}
```
### 查询 binlog 备份列表
使用以下代码查询 binlog 备份列表
```go
args := &gaiadb.GetBinlogListArgs{
	AppID:         clusterId,
	LogBackupType: "logical",
	PageNo:        1,
	PageSize:      10,
}
result, err := GAIADB_CLIENT.GetBinlogList(args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get binlog list failed:", err)
} else {
    fmt.Println("get binlog list success: ", result)
}
```
## 任务管理
### 立即执行任务
使用以下代码立即执行任务
```go
err := GAIADB_CLIENT.ExecuteTaskNow(taskId)
if err != nil {
    fmt.Println("execute task now failed:", err)
} else {
    fmt.Println("execute task now success: ", result)
}
```
### 取消任务
使用以下代码取消任务
```go
err := GAIADB_CLIENT.CancelTask(taskId)
if err != nil {
    fmt.Println("cancel task failed:", err)
} else {
    fmt.Println("cancel task success: ", result)
}
```
### 查询任务列表
使用以下代码查询任务列表
```go
args := &gaiadb.TaskListArgs{
	Region:    "bj",
	StartTime: "2023-09-11 16:00:00",
}
result, err := GAIADB_CLIENT.GetTaskList(args)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get task list failed:", err)
} else {
    fmt.Println("get task list success: ", result)
}
```
## 安全组管理
### 通过VPC ID查询GaiaDB集群
使用以下代码通过VPC ID查询GaiaDB集群
```go
result, err := GAIADB_CLIENT.GetClusterByVpcId(vpcId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get cluster by vpcid failed:", err)
} else {
    fmt.Println("get cluster by vpcid success: ", result)
}
```
### 通过Lb ID查询GaiaDB集群
使用以下代码通过Lb ID查询GaiaDB集群
```go
result, err := GAIADB_CLIENT.GetClusterByLbId(lbIds)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get cluster by lbids failed:", err)
} else {
    fmt.Println("get cluster by lbids success: ", result)
}
```
## 其他
### 查询订单信息
使用以下代码查询订单信息
```go
result, err := GAIADB_CLIENT.GetOrderInfo(orderId)
re, _ := json.Marshal(result)
fmt.Println(string(re))
if err != nil {
    fmt.Println("get order info failed:", err)
} else {
    fmt.Println("get order info success: ", result)
}
```
# 错误处理

GO语言以error类型标识错误，SCS支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | GAIADB服务返回的错误

用户使用SDK调用GAIADB相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// gaiadbClient 为已创建的GAIADB Client对象
clusterDetail, err := gaiadbClient.GetClusterDetail(clusterId)
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
	fmt.Println("get instance detail success: ", clusterDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向GAIADB发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当GAIADB服务端出现异常时，GAIADB服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[GAIADB错误返回](https://cloud.baidu.com/doc/GaiaDB/s/al84889rz)

