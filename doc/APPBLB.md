# APPBLB服务

# 概述

本文档主要介绍应用型BLB GO SDK的使用。在使用本文档前，您需要先了解应用型BLB的一些基本知识。若您还不了解应用型BLB，可以参考[产品描述](https://cloud.baidu.com/doc/BLB/s/Ajwvxno34)和[入门指南](https://cloud.baidu.com/doc/BLB/s/cjwvxnr91)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BLB访问域名](https://cloud.baidu.com/doc/BLB/s/cjwvxnzix)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云BLB，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BLB做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建应用型BLB Client

应用型BLB Client是应用型BLB服务的客户端，为开发者与BLB服务进行交互提供了一系列的方法。

### 使用AK/SK新建应用型BLB Client

通过AK/SK方式访问BLB，用户可以参考如下代码新建一个BLB Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/appblb"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BLBClient
	blbClient, err := appblb.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/BLB/s/ojwvynrqn)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BLB的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`blb.bj.baidubce.com`。

### 使用STS创建BLB Client

**申请STS token**

BLB可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问BLB，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建BLB Client**

申请好STS后，可将STS Token配置到BLB Client中，从而实现通过STS Token创建BLB Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建BLB Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/appblb" //导入APPBLB服务模块
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

	// 使用申请的临时STS创建BLB服务的Client对象，Endpoint使用默认值
	blbClient, err := appblb.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create appblb client failed:", err)
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
	blbClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置BLB Client时，无论对应BLB服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问BLB

BLB支持HTTPS传输协议，您可以通过在创建BLB Client对象时指定的Endpoint中指明HTTPS的方式，在BLB GO SDK中使用HTTPS访问BLB服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/appblb"

ENDPOINT := "https://blb.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
blbClient, _ := appblb.NewClient(AK, SK, ENDPOINT)
```

## 配置BLB Client

如果用户需要配置BLB Client的一些细节的参数，可以在创建BLB Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问BLB服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/appblb"

//创建BLB Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com
client, _ := appblb.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/appblb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com"
client, _ := appblb.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/appblb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com"
client, _ := appblb.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问BLB时，创建的BLB Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建BLB Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

应用型blb实例针对用户复杂应用部署架构，特别是大型网站架构。使用基于策略的网络管理框架构建，实现业务驱动的流量负载均衡。

## 实例管理

### 创建实例

通过以下代码，可以创建一个应用型LoadBalancer，返回分配的服务地址及实例ID
```go
args := &appblb.CreateLoadBalancerArgs{
    // 设置实例名称
    Name:        "sdkBlb",
    // 设置实例描述
    Description: "sdk create",
    // 设置实例所属vpc
    VpcId:       vpcId,
    // 设置实例所属子网 
    SubnetId:    subnetId,
}

// 若要为实例设置标签，可以按照以下代码，标签设置之后，不可修改和删除
args.Tags = []model.TagModel{
    {
        TagKey:   "key", 
        TagValue: "value",
    },
}
result, err := client.CreateLoadBalancer(args)
if err != nil {
    fmt.Println("create appblb failed:", err)
} else {
    fmt.Println("create appblb success: ", result)
}
```

### 更新实例

通过以下代码，可以更新一个LoadBalancer的配置信息，如实例名称和描述
```go
args := &appblb.UpdateLoadBalancerArgs{
    Name:        "testSdk", 
    Description: "test desc",
}
err := client.UpdateLoadBalancer(blbId, args)
if err != nil {
    fmt.Println("update appblb failed:", err)
} else {
    fmt.Println("update appblb success")
}
```
### 查询已有的实例

通过以下代码，可以查询用户账户下所有LoadBalancer的信息
```go
args := &appblb.DescribeLoadBalancersArgs{}

// 支持按LoadBalancer的id、name、address进行查询，匹配规则支持部分包含（不支持正则）
args.BlbId = blbId
args.Name = blbName
args.Address = blbAddress
args.ExactlyMatch = true

// 支持查找绑定指定BLB的LoadBalancer，通过blbId参数指定
args.BlbId = blbId

result, err := client.DescribeLoadBalancers(args)
if err != nil {
    fmt.Println("list all appblb failed:", err)
} else {
    fmt.Println("list all appblb success: ", result)
}
```

### 查询实例详情

通过以下代码，可以按id查询用户账户下特定的应用型LoadBalancer的详细信息，包含LoadBalancer所有的监听器端口信息
```go
result, err := client.DescribeLoadBalancerDetail(blbId)
if err != nil {
    fmt.Println("get appblb detail failed:", err)
} else {
    fmt.Println("get appblb detail success: ", result)
}
```

### 释放实例

通过以下代码，可以释放指定LoadBalancer，被释放的LoadBalancer无法找回
```go
err := client.DeleteLoadBalancer(blbId)
if err != nil {
    fmt.Println("delete appblb failed:", err)
} else {
    fmt.Println("delete appblb success")
}
```

## 服务器组管理

### 创建应用型服务器组

通过以下代码，在指定应用型BLB下，创建一个服务器组，用来绑定后端服务器，以及为监听器开放相应的端口
```go
args := &appblb.CreateAppServerGroupArgs{
    // 设置服务器组名称
    Name: "sdkTest", 
    // 设置服务器组描述
    Description: "sdk test desc",
}

// 若想直接绑定后端服务器，可以设置以下参数
args.BackendServerList = []appblb.AppBackendServer{
    {
        InstanceId: instanceId, 
        Weight:     50,
    },
}

result, err := client.CreateAppServerGroup(blbId, args)
if err != nil {
    fmt.Println("create server group failed:", err)
} else {
    fmt.Println("create server group success: ", result)
}
```

### 更新服务器组

通过以下代码，更新指定LoadBalancer下的TCP监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &appblb.UpdateAppServerGroupArgs{
    // 设置要更新的服务器组ID
    SgId:        serverGroupId, 
    // 设置新的服务器组名称
    Name:        "testSdk",
    // 设置新的服务器组描述
    Description: "test desc",
}
err := client.UpdateAppServerGroup(blbId, args)
if err != nil {
    fmt.Println("update server group failed:", err)
} else {
    fmt.Println("update server group success: ", result)
}
```

### 查询服务器组

通过以下代码，查询指定LoadBalancer下所有服务器组的信息
```go
args := &appblb.DescribeAppServerGroupArgs{}

// 按BLBID、name为条件进行全局查询
args.BlbId = blbId
args.Name = servergroupName
args.ExactlyMatch = true

result, err := client.DescribeAppServerGroup(blbId, args)
if err != nil {
    fmt.Println("describe server group failed:", err)
} else {
    fmt.Println("describe server group success: ", result)
}
```

### 删除服务器组

通过以下代码，删除服务器组，通过服务器组id指定
```go
args := &appblb.DeleteAppServerGroupArgs{
    // 要删除的服务器组ID
    SgId: serverGroupId,
}
err := client.DeleteAppServerGroup(blbId, args)
if err != nil {
    fmt.Println("delete server group failed:", err)
} else {
    fmt.Println("delete server group success: ", result)
}
```

### 创建应用型服务器组端口

通过以下代码，在指定应用型BLB下，创建一个服务器组后端端口，将发往该端口的所有流量按权重轮询分发到其绑定的对应服务器列表中的服务器
```go
args := &appblb.CreateAppServerGroupPortArgs{
    // 端口所属服务器组ID
    SgId: serverGroupId, 
    // 监听的端口号
    Port: 80,
    // 监听的协议类型
    Type: "TCP",
}

// 可以选择设置相应健康检查协议的参数
args.HealthCheck = "TCP"
args.HealthCheckPort = 30
args.HealthCheckIntervalInSecond = 10
args.HealthCheckTimeoutInSecond = 10

result, err := client.CreateAppServerGroupPort(blbId, args)
if err != nil {
    fmt.Println("create server group port failed:", err)
} else {
    fmt.Println("create server group port success: ", result)
}
```

> **提示：**
> - 配置健康检查协议的参数，详细请参考BLB API 文档[创建应用型服务器组端口](https://cloud.baidu.com/doc/BLB/s/Bjwvxny4u#createappservergroupport%E5%88%9B%E5%BB%BA%E5%BA%94%E7%94%A8%E5%9E%8B%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%BB%84%E7%AB%AF%E5%8F%A3)

### 更新服务器组端口

通过以下代码，根据id更新服务器组端口
```go
args := &appblb.UpdateAppServerGroupPortArgs{
	// 端口所属服务器组ID
    SgId: serverGroupId, 
    // 端口Id，由创建创建应用型服务器组端口返回
    PortId: portId,
    // 设置健康检查协议参数
    HealthCheck:                 "TCP", 
    HealthCheckPort:             30,
    HealthCheckIntervalInSecond: 10, 
    HealthCheckTimeoutInSecond:  10,
}
err := client.UpdateAppServerGroupPort(blbId, args)
if err != nil {
    fmt.Println("update server group port failed:", err)
} else {
    fmt.Println("update server group port success: ", result)
}
```

> **提示：**
> - 配置健康检查协议的参数，详细请参考BLB API 文档[更新服务器组端口](https://cloud.baidu.com/doc/BLB/s/Bjwvxny4u#updateappservergroupport%E6%9B%B4%E6%96%B0%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%BB%84%E7%AB%AF%E5%8F%A3)

### 删除服务器组端口

通过以下代码，删除服务器组端口，通过服务器组id指定
```go
args := &appblb.DeleteAppServerGroupPortArgs{
	// 端口所属服务器组ID
    SgId:        serverGroupId,
	// 要删除的端口服务ID列表
    PortIdList:  []string{portId}, 
}
err := client.DeleteAppServerGroupPort(blbId, args)
if err != nil {
    fmt.Println("delete server group port failed:", err)
} else {
    fmt.Println("delete server group port success: ", result)
}
```

### 添加应用型BLB后端RS

通过以下代码，在指定应用型BLB和服务器组下绑定后端服务器RS
```go
args := &appblb.CreateBlbRsArgs{
    BlbRsWriteOpArgs: appblb.BlbRsWriteOpArgs{
        // RS所属服务器组ID
        SgId:        serverGroupId, 
        // 配置后端服务器的列表及权重
        BackendServerList: []appblb.AppBackendServer{  
            {InstanceId: instanceId, Weight: 30},
        },
    },
}
err := client.CreateBlbRs(blbId, args)
if err != nil {
    fmt.Println("create blbRs failed:", err)
} else {
    fmt.Println("create blbRs success: ", result)
}
```

### 更新服务器组下挂载的RS权重

通过以下代码，更新指定服务器组下的RS信息
```go
args := &appblb.UpdateBlbRsArgs{
    BlbRsWriteOpArgs: appblb.BlbRsWriteOpArgs{
        // RS所属服务器组ID
        SgId:        serverGroupId, 
        // 配置要更新的后端服务器的列表及权重
        BackendServerList: []appblb.AppBackendServer{
            {InstanceId: instanceId, Weight: 60},
        },
    },
}
err := client.UpdateBlbRs(blbId, args)
if err != nil {
    fmt.Println("update blbRs failed:", err)
} else {
    fmt.Println("update blbRs success: ", result)
}
```

### 查询服务器组下的RS列表信息

通过以下代码，查询指定LoadBalancer下所有服务器组的信息
```go
args := &appblb.DescribeBlbRsArgs{
	// RS所属服务器组ID
    SgId: serverGroupId,
}
result, err := client.DescribeBlbRs(blbId, args)
if err != nil {
    fmt.Println("describe blbRs failed:", err)
} else {
    fmt.Println("describe blbRs success: ", result)
}
```

### 删除服务器组下挂载的rs

通过以下代码，删除服务器组，通过服务器组id指定
```go
args := &appblb.DeleteBlbRsArgs{
	// RS所属服务器组ID
    SgId: serverGroupId, 
    // 要从RS列表中删除的实例列表
    BackendServerIdList: []string{instanceId},
}
err := client.DeleteBlbRs(blbId, args)
if err != nil {
    fmt.Println("delete blbRs failed:", err)
} else {
    fmt.Println("delete blbRs success: ", result)
}
```

### 查询服务器组下绑定的server

通过以下代码，查询服务器组下绑定的server
```go
result, err := client.DescribeRsMount(blbId, serverGroupId)
if err != nil {
    fmt.Println("describe mount Rs list failed:", err)
} else {
    fmt.Println("describe mount Rs list success: ", result)
}
```

### 查询服务器组下未绑定的RS

通过以下代码，查询服务器组下未绑定的RS
```go
result, err := client.DescribeRsUnMount(blbId, serverGroupId)
if err != nil {
    fmt.Println("describe unmount Rs list failed:", err)
} else {
    fmt.Println("describe unmount Rs list success: ", result)
}
```

## 监听器管理

### 创建TCP监听器

通过以下代码，在指定LoadBalancer下，创建一个基于TCP协议的应用型blb监听器，监听一个前端端口，将发往该端口的所有TCP流量，根据策略进行转发
```go
args := &appblb.CreateAppTCPListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 90,
    // 负载均衡算法，支持RoundRobin/LeastConnection/Hash
    Scheduler: "RoundRobin",
    // TCP设置链接超时时间，默认900秒，需要为10-4000之间的整数
    TcpSessionTimeout: 1000,
}
err := client.CreateAppTCPListener(BLBID, args)
if err != nil {
    fmt.Println("create TCP Listener failed:", err)
} else {
    fmt.Println("create TCP Listener success")
}
```

### 更新TCP监听器

通过以下代码，更新指定LoadBalancer下的TCP监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &appblb.UpdateAppTCPListenerArgs{
    UpdateAppListenerArgs: appblb.UpdateAppListenerArgs{
        // 要更新的监听器端口号
        ListenerPort: 90, 
        // 更新负载均衡的算法
        Scheduler:    "Hash", 
        // 更新tcp链接超时时间
        TcpSessionTimeout: 2000,
    },
}
err := client.UpdateAppTCPListener(BLBID, args)
if err != nil {
    fmt.Println("update TCP Listener failed:", err)
} else {
    fmt.Println("update TCP Listener success")
}
```

### 查询TCP监听器

通过以下代码，查询指定LoadBalancer下所有TCP监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &appblb.DescribeAppListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 90,
}
result, err := client.DescribeAppTCPListeners(BLBID, args)
if err != nil {
    fmt.Println("describe TCP Listener failed:", err)
} else {
    fmt.Println("describe TCP Listener success: ", result)
}
```

### 创建UDP监听器

通过以下代码，在指定LoadBalancer下，创建一个基于UDP协议的应用型监听器，监听一个前端端口，将发往该端口的所有UDP流量，根据策略进行转发
```go
args := &appblb.CreateAppUDPListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection/Hash
    Scheduler:    "RoundRobin",
}
err := client.CreateAppUDPListener(BLBID, args)
if err != nil {
    fmt.Println("create UDP Listener failed:", err)
} else {
    fmt.Println("create UDP Listener success")
}
```

### 更新UDP监听器

通过以下代码，更新指定LoadBalancer下的UDP监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &appblb.UpdateAppUDPListenerArgs{
    UpdateAppListenerArgs: appblb.UpdateAppListenerArgs{
        // 要更新的监听器端口号
        ListenerPort: 80, 
        // 更新负载均衡的算法
        Scheduler:    "Hash",
    },
}
err := client.UpdateAppUDPListener(BLBID, args)
if err != nil {
    fmt.Println("update UDP Listener failed:", err)
} else {
    fmt.Println("update UDP Listener success")
}
```

### 查询UDP监听器

通过以下代码，查询指定LoadBalancer下所有UDP监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &appblb.DescribeAppListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 80,
}
result, err := client.DescribeAppUDPListeners(BLBID, args)
if err != nil {
    fmt.Println("describe UDP Listener failed:", err)
} else {
    fmt.Println("describe UDP Listener success: ", result)
}
```

### 创建HTTP监听器

通过以下代码，在指定LoadBalancer下，创建一个基于HTTP协议的应用型监听器，监听一个前端端口，将发往该端口的所有HTTP请求，根据策略转发到后端服务器监听的后端端口上
```go
args := &appblb.CreateAppHTTPListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection
    Scheduler:    "RoundRobin",
}
err := client.CreateAppHTTPListener(BLBID, args)
if err != nil {
    fmt.Println("create HTTP Listener failed:", err)
} else {
    fmt.Println("create HTTP Listener success")
}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[创建HTTP监听器](https://cloud.baidu.com/doc/BLB/s/ujwvxnyux#createapphttplistener%E5%88%9B%E5%BB%BAhttp%E7%9B%91%E5%90%AC%E5%99%A8)

### 更新HTTP监听器

通过以下代码，更新指定LoadBalancer下的HTTP监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &appblb.UpdateAppHTTPListenerArgs{
	// 要更新的监听器端口号
    ListenerPort: 80,
    // 更新负载均衡的算法
    Scheduler:    "LeastConnection", 
    // 开启会话保持功能
    KeepSession:  true,
}
err := client.UpdateAppHTTPListener(BLBID, args)
if err != nil {
    fmt.Println("update HTTP Listener failed:", err)
} else {
    fmt.Println("update HTTP Listener success")
}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[更新HTTP监听器](https://cloud.baidu.com/doc/BLB/s/ujwvxnyux#updateapphttplistener%E6%9B%B4%E6%96%B0http%E7%9B%91%E5%90%AC%E5%99%A8)

### 查询HTTP监听器

通过以下代码，查询指定LoadBalancer下所有HTTP监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &appblb.DescribeAppListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 80,
}
result, err := client.DescribeAppHTTPListeners(BLBID, args)
if err != nil {
    fmt.Println("describe HTTP Listener failed:", err)
} else {
    fmt.Println("describe HTTP Listener success: ", result)
}
```

### 创建HTTPS监听器

通过以下代码，在指定LoadBalancer下，创建一个基于HTTPS协议的应用型监听器，监听一个前端端口，将发往该端口的所有HTTPS请求，先通过SSL卸载转换为HTTP请求后，再根据策略转发到后端服务器监听的后端端口上
```go
args := &appblb.CreateAppHTTPSListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection
    Scheduler:    "RoundRobin", 
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.CreateAppHTTPSListener(BLBID, args)
if err != nil {
    fmt.Println("create HTTPS Listener failed:", err)
} else {
    fmt.Println("create HTTPS Listener success")
}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[创建HTTPS监听器](https://cloud.baidu.com/doc/BLB/s/ujwvxnyux#createapphttpslistener%E5%88%9B%E5%BB%BAhttps%E7%9B%91%E5%90%AC%E5%99%A8)

### 更新HTTPS监听器

通过以下代码，更新指定LoadBalancer下的HTTPS监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &appblb.UpdateAppHTTPSListenerArgs{
	// 要更新的监听器端口号
    ListenerPort: 80,
    // 更新负载均衡的算法
    Scheduler:    "LeastConnection", 
    // 开启会话保持功能
    KeepSession:  true,
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.UpdateAppHTTPSListener(BLBID, args)
if err != nil {
    fmt.Println("update HTTPS Listener failed:", err)
} else {
    fmt.Println("update HTTPS Listener success")
}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[更新HTTPS监听器](https://cloud.baidu.com/doc/BLB/s/ujwvxnyux#updateapphttpslistener%E6%9B%B4%E6%96%B0https%E7%9B%91%E5%90%AC%E5%99%A8)

### 查询HTTPS监听器

通过以下代码，查询指定LoadBalancer下所有HTTPS监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &appblb.DescribeAppListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 80,
}
result, err := client.DescribeAppHTTPSListeners(BLBID, args)
if err != nil {
    fmt.Println("describe HTTPS Listener failed:", err)
} else {
    fmt.Println("describe HTTPS Listener success: ", result)
}
```

### 创建SSL监听器

通过以下代码，在指定LoadBalancer下，创建一个基于SSL协议的应用型blb监听器，监听一个前端端口，将发往该端口的所有SSL流量，根据策略进行转发
```go
args := &appblb.CreateAppSSLListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection/Hash
    Scheduler:    "RoundRobin", 
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.CreateAppSSLListener(BLBID, args)
if err != nil {
    fmt.Println("create SSL Listener failed:", err)
} else {
    fmt.Println("create SSL Listener success")
}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[创建SSL监听器](https://cloud.baidu.com/doc/BLB/s/ujwvxnyux#createappssllistener%E5%88%9B%E5%BB%BAssl%E7%9B%91%E5%90%AC%E5%99%A8)

### 更新SSL监听器

通过以下代码，更新指定LoadBalancer下的SSL监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &appblb.UpdateAppSSLListenerArgs{
	// 要更新的监听器端口号
    ListenerPort: 80,
    // 更新负载均衡的算法
    Scheduler:    "LeastConnection", 
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.UpdateAppSSLListener(BLBID, args)
if err != nil {
    fmt.Println("update SSL Listener failed:", err)
} else {
    fmt.Println("update SSL Listener success")
}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[更新SSL监听器](https://cloud.baidu.com/doc/BLB/s/ujwvxnyux#updateappssllistener%E6%9B%B4%E6%96%B0ssl%E7%9B%91%E5%90%AC%E5%99%A8)

### 查询SSL监听器

通过以下代码，查询指定LoadBalancer下所有HTTPS监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &appblb.DescribeAppListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 80,
}
result, err := client.DescribeAppSSLListeners(BLBID, args)
if err != nil {
    fmt.Println("describe SSL Listener failed:", err)
} else {
    fmt.Println("describe SSL Listener success: ", result)
}
```

### 删除监听器

通过以下代码，释放指定LoadBalancer下的监听器，监听器通过监听端口来指定，支持批量释放
```go
args := &appblb.DeleteAppListenersArgs{
    // 要删除的监听器监听的端口
    PortList:    []uint16{80}, 
}
err := client.DeleteAppListeners(BLBID, args)
if err != nil {
    fmt.Println("delete Listeners failed:", err)
} else {
    fmt.Println("delete Listeners success: ", result)
}
```

### 创建策略

通过以下代码，在指定应用型BLB监听器端口下创建策略
```go
args := &appblb.CreatePolicysArgs{
    // 需要绑定策略的监听器监听的端口
    ListenerPort: 80,
    // 需要绑定的策略，其中TCP/UDP/SSL仅支持绑定一个策略，HTTP/HTTPS支持绑定多个策略 
    AppPolicyVos: []appblb.AppPolicy{
    {
        // 策略描述
        Description:      "test policy", 
        // 策略绑定的服务器组ID
        AppServerGroupId: servergroupId, 
        // 目标端口号
        BackendPort:      80, 
        // 策略的优先级，有效取值范围为1-32768
        Priority:         301, 
        // 策略的规则列表，TCP/UDP/SSL仅支持{*：*}策略
        RuleList: []appblb.AppRule{
                {
                    Key:   "*",
                    Value: "*",	
                },
            },
        },
    },
}
err := client.CreatePolicys(blbId, args)
if err != nil {
    fmt.Println("create policy failed:", err)
} else {
    fmt.Println("create policy success")
}
```

> **提示：**
> - 策略中backendPort参数，即目标端口号，当listenerPort对应监听器为TCP或SSL时需要传入对应服务器组（appServerGroupId）下开放的TCP端口号；当listenerPort对应监听器为HTTP或HTTPS时需要传入对应服务器组（appServerGroupId）下开放的HTTP端口号；当listenerPort对应监听器为UDP时需要传入对应服务器组（appServerGroupId）下开放的UDP端口号
> - 各个协议可绑定的策略的具体配置参数，可以参考BLB API 文档[创建策略](https://cloud.baidu.com/doc/BLB/s/ujwvxnyux#createpolicys%E5%88%9B%E5%BB%BA%E7%AD%96%E7%95%A5)

### 查询对应BLB端口下策略信息

通过以下代码，查询指定LoadBalancer下所有服务器组的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &appblb.DescribePolicysArgs{
    // 要查询的监听器端口号
    Port: 80,
}
result, err := client.DescribePolicys(blbId, args)
if err != nil {
    fmt.Println("describe policy failed:", err)
} else {
    fmt.Println("describe policy success: ", result)
}
```

### 批量删除策略

通过以下代码，批量删除对应BLB端口下的策略
```go
args := &appblb.DeletePolicysArgs{
    // 要删除策略所在监听器的端口号
    Port: 80, 
    // 要删除的策略ID列表
    PolicyIdList: []string{describeResult.PolicyList[0].Id}, 
}
err := client.DeletePolicys(blbId, args)
if err != nil {
    fmt.Println("delete policy failed:", err)
} else {
    fmt.Println("delete policy success")
}
```

# 错误处理

GO语言以error类型标识错误，BLB支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | BLB服务返回的错误

用户使用SDK调用BLB相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```go
// blbClient 为已创建的BLB Client对象
blbDetail, err := blbClient.DescribeLoadBalancerDetail(blbId)
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
	fmt.Println("get appblb detail success: ", blbDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向BLB发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当BLB服务端出现异常时，BLB服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[BLB错误返回](https://cloud.baidu.com/doc/BLB/s/Djwvxnzw6)

# 版本变更记录

## v0.9.1 [2019-09-26]

首次发布：

 - 创建、查看、列表、更新、删除应用型BLB实例
 - 创建、列表、更新、删除服务器组
 - 创建、更新、删除服务器组端口
 - 创建、列表、更新、删除服务器组后端RS，并支持查询已绑定和未绑定的服务器
 - 创建、查看、更新、删除监听器端口，支持TCP/UDP/HTTP/HTTPS/SSL协议
 - 创建、查看、删除监听器相关策略
