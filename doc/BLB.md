# BLB服务

# 概述

本文档主要介绍普通型BLB GO SDK的使用。在使用本文档前，您需要先了解普通型BLB的一些基本知识。若您还不了解普通型BLB，可以参考[产品描述](https://cloud.baidu.com/doc/BLB/s/Ajwvxno34)和[入门指南](https://cloud.baidu.com/doc/BLB/s/cjwvxnr91)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BLB访问域名](https://cloud.baidu.com/doc/BLB/s/cjwvxnzix)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云BLB，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BLB做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建普通型BLB Client

普通型BLB Client是普通型BLB服务的客户端，为开发者与BLB服务进行交互提供了一系列的方法。

### 使用AK/SK新建普通型BLB Client

通过AK/SK方式访问BLB，用户可以参考如下代码新建一个BLB Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/blb"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BLBClient
	blbClient, err := blb.NewClient(AK, SK, ENDPOINT)
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
	"github.com/baidubce/bce-sdk-go/services/blb" //导入BLB服务模块
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
	blbClient, err := blb.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create blb client failed:", err)
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
// import "github.com/baidubce/bce-sdk-go/services/blb"

ENDPOINT := "https://blb.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
blbClient, _ := blb.NewClient(AK, SK, ENDPOINT)
```

## 配置BLB Client

如果用户需要配置BLB Client的一些细节的参数，可以在创建BLB Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问BLB服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/blb"

//创建BLB Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com
client, _ := blb.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/blb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com"
client, _ := blb.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/blb"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "blb.bj.baidubce.com"
client, _ := blb.NewClient(AK, SK, ENDPOINT)

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

普通型blb实例针对用户复杂应用部署架构，特别是大型网站架构。使用基于策略的网络管理框架构建，实现业务驱动的流量负载均衡。

## 实例管理

### 创建实例

通过以下代码，可以创建一个普通型LoadBalancer，返回分配的服务地址及实例ID
```go
args := &blb.CreateLoadBalancerArgs{
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
    fmt.Println("create blb failed:", err)
} else {
    fmt.Println("create blb success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[CreateLoadBalancer创建实例](https://cloud.baidu.com/doc/BLB/s/njwvxnv79#createloadbalancer%E5%88%9B%E5%BB%BA%E5%AE%9E%E4%BE%8B)

### 更新实例

通过以下代码，可以更新一个LoadBalancer的配置信息，如实例名称和描述
```go
args := &blb.UpdateLoadBalancerArgs{
    Name:        "testSdk", 
    Description: "test desc",
}
err := client.UpdateLoadBalancer(blbId, args)
if err != nil {
    fmt.Println("update blb failed:", err)
} else {
    fmt.Println("update blb success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[UpdateLoadBalancer更新实例](https://cloud.baidu.com/doc/BLB/s/njwvxnv79#updateloadbalancer%E6%9B%B4%E6%96%B0%E5%AE%9E%E4%BE%8B)

### 查询已有的实例

通过以下代码，可以查询用户账户下所有LoadBalancer的信息
```go
args := &blb.DescribeLoadBalancersArgs{}

// 支持按LoadBalancer的id、name、address进行查询，匹配规则支持部分包含（不支持正则）
args.BlbId = blbId
args.Name = blbName
args.Address = blbAddress
args.ExactlyMatch = true

// 支持查找绑定指定BLB的LoadBalancer，通过blbId参数指定
args.BlbId = blbId

result, err := client.DescribeLoadBalancers(args)
if err != nil {
    fmt.Println("list all blb failed:", err)
} else {
    fmt.Println("list all blb success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeLoadBalancers查询已有的BLB实例](https://cloud.baidu.com/doc/BLB/s/njwvxnv79#describeloadbalancers%E6%9F%A5%E8%AF%A2%E5%B7%B2%E6%9C%89%E7%9A%84blb%E5%AE%9E%E4%BE%8B)

### 查询实例详情

通过以下代码，可以按id查询用户账户下特定的普通型LoadBalancer的详细信息，包含LoadBalancer所有的监听器端口信息
```go
result, err := client.DescribeLoadBalancerDetail(blbId)
if err != nil {
    fmt.Println("get blb detail failed:", err)
} else {
    fmt.Println("get blb detail success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeLoadBalancerDetail查询BLB实例详情](https://cloud.baidu.com/doc/BLB/s/njwvxnv79#describeloadbalancerdetail%E6%9F%A5%E8%AF%A2blb%E5%AE%9E%E4%BE%8B%E8%AF%A6%E6%83%85)

### 释放实例

通过以下代码，可以释放指定LoadBalancer，被释放的LoadBalancer无法找回
```go
err := client.DeleteLoadBalancer(blbId)
if err != nil {
    fmt.Println("delete blb failed:", err)
} else {
    fmt.Println("delete blb success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DeleteLoadBalancer释放BLB实例](https://cloud.baidu.com/doc/BLB/s/njwvxnv79#deleteloadbalancer%E9%87%8A%E6%94%BEblb%E5%AE%9E%E4%BE%8B)

### 添加普通型BLB后端服务器

通过以下代码，在指定普通型BLB下绑定后端服务器
```go
args := &blb.AddBackendServersArgs{
    // 配置后端服务器的列表及权重
    BackendServerList: []blb.BackendServerModel{  
        {InstanceId: instanceId, Weight: 100},
    },
}
err := client.AddBackendServers(blbId, args)
if err != nil {
    fmt.Println("add backend servers failed:", err)
} else {
    fmt.Println("add backend servers success: ")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[AddBackendServers添加后端服务器](https://cloud.baidu.com/doc/BLB/s/Ujwvxnvxe#addbackendservers%E6%B7%BB%E5%8A%A0%E5%90%8E%E7%AB%AF%E6%9C%8D%E5%8A%A1%E5%99%A8)

### 更新后端服务器权重

通过以下代码，更新指定blb下的信息
```go
args := &blb.UpdateBackendServersArgs{
    // 配置后端服务器的列表及权重
    BackendServerList: []blb.BackendServerModel{  
        {InstanceId: instanceId, Weight: 30},
    },
}
err := client.UpdateBackendServers(blbId, args)
if err != nil {
    fmt.Println("update backend servers failed:", err)
} else {
    fmt.Println("update backend servers success: ")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[UpdateBackendServers更新后端服务器](https://cloud.baidu.com/doc/BLB/s/Ujwvxnvxe#updatebackendservers%E6%9B%B4%E6%96%B0%E5%90%8E%E7%AB%AF%E6%9C%8D%E5%8A%A1%E5%99%A8)

### 查询后端服务器列表信息

通过以下代码，查询指定LoadBalancer下所有服务器的信息
```go
args := &blb.DescribeBackendServersArgs{
	
}
result, err := client.DescribeBackendServers(blbId, args)
if err != nil {
    fmt.Println("describe backend servers failed:", err)
} else {
    fmt.Println("describe backend servers success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeBackendServers查询后端服务器列表](https://cloud.baidu.com/doc/BLB/s/Ujwvxnvxe#describebackendservers%E6%9F%A5%E8%AF%A2%E5%90%8E%E7%AB%AF%E6%9C%8D%E5%8A%A1%E5%99%A8%E5%88%97%E8%A1%A8)

### 查询后端服务器健康状态

通过以下代码，查询指定监听端口下后端服务器的健康状态信息
```go
args := &blb.DescribeHealthStatusArgs{
	ListenerPort: 80,
}
result, err := client.DescribeHealthStatus(blbId, args)
if err != nil {
    fmt.Println("describe health status failed:", err)
} else {
    fmt.Println("describe health status success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeHealthStatus查询后端服务器健康状态](https://cloud.baidu.com/doc/BLB/s/Ujwvxnvxe#describehealthstatus%E6%9F%A5%E8%AF%A2%E5%90%8E%E7%AB%AF%E6%9C%8D%E5%8A%A1%E5%99%A8%E5%81%A5%E5%BA%B7%E7%8A%B6%E6%80%81)

### 释放后端服务器

通过以下代码，释放后端服务器
```go
args := &blb.RemoveBackendServersArgs{
    // 要从后端服务器列表中释放的实例列表
    BackendServerList: []string{instanceId},
}
err := client.RemoveBackendServers(blbId, args)
if err != nil {
    fmt.Println("remove backend servers failed:", err)
} else {
    fmt.Println("remove backend servers success: ")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[RemoveBackendServers释放后端服务器](https://cloud.baidu.com/doc/BLB/s/Ujwvxnvxe#removebackendservers%E9%87%8A%E6%94%BE%E5%90%8E%E7%AB%AF%E6%9C%8D%E5%8A%A1%E5%99%A8)

## 监听器管理

### 创建TCP监听器

通过以下代码，在指定LoadBalancer下，创建一个基于TCP协议的普通型blb监听器，监听一个前端端口，将发往该端口的所有TCP流量，根据策略进行转发
```go
args := &blb.CreateTCPListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 80,
    // 后端服务器的监听端口，需要在1-65535之间
    BackendPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection/Hash
    Scheduler: "RoundRobin",
    // TCP设置链接超时时间，默认900秒，需要为10-4000之间的整数
    TcpSessionTimeout: 1000,
}
err := client.CreateTCPListener(BLBID, args)
if err != nil {
    fmt.Println("create TCP Listener failed:", err)
} else {
    fmt.Println("create TCP Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[CreateTCPListener创建TCP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#createtcplistener%E5%88%9B%E5%BB%BAtcp%E7%9B%91%E5%90%AC%E5%99%A8)

### 更新TCP监听器

通过以下代码，更新指定LoadBalancer下的TCP监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &blb.UpdateTCPListenerArgs{
    // 要更新的监听器端口号
    ListenerPort: 80,
    // 更新负载均衡的算法
    Scheduler:    "Hash", 
    // 更新tcp链接超时时间
    TcpSessionTimeout: 2000,
}
err := client.UpdateTCPListener(BLBID, args)
if err != nil {
    fmt.Println("update TCP Listener failed:", err)
} else {
    fmt.Println("update TCP Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[UpdateTCPListener更新TCP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#updatetcplistener%E6%9B%B4%E6%96%B0tcp%E7%9B%91%E5%90%AC%E5%99%A8)

### 查询TCP监听器

通过以下代码，查询指定LoadBalancer下所有TCP监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &blb.DescribeListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 80,
}
result, err := client.DescribeTCPListeners(BLBID, args)
if err != nil {
    fmt.Println("describe TCP Listener failed:", err)
} else {
    fmt.Println("describe TCP Listener success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeTCPListeners查询TCP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#describetcplisteners%E6%9F%A5%E8%AF%A2tcp%E7%9B%91%E5%90%AC%E5%99%A8)

### 创建UDP监听器

通过以下代码，在指定LoadBalancer下，创建一个基于UDP协议的监听器，监听一个前端端口，将发往该端口的所有UDP流量，根据策略进行转发
```go
args := &blb.CreateUDPListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 53,
    // 后端服务器的监听端口，需要在1-65535之间
    BackendPort: 53,
    // 负载均衡算法，支持RoundRobin/LeastConnection/Hash
    Scheduler:    "RoundRobin",
    // 健康检查字符串 健康发送的请求字符串，后端服务器收到后需要进行应答，支持16进制\00-\FF和标准ASCII字符串，最大长度1299
    HealthCheckString: "\00\01\01\00\00\01\00\00\00\00\00\00\05baidu\03com\00\00\01\00\01"
}
err := client.CreateUDPListener(BLBID, args)
if err != nil {
    fmt.Println("create UDP Listener failed:", err)
} else {
    fmt.Println("create UDP Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[CreateUDPListener创建UDP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#createudplistener%E5%88%9B%E5%BB%BAudp%E7%9B%91%E5%90%AC%E5%99%A8)

### 更新UDP监听器

通过以下代码，更新指定LoadBalancer下的UDP监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &blb.UpdateUDPListenerArgs{
    // 要更新的监听器端口号
    ListenerPort: 53, 
    // 更新负载均衡的算法
    Scheduler:    "Hash",
}
err := client.UpdateUDPListener(BLBID, args)
if err != nil {
    fmt.Println("update UDP Listener failed:", err)
} else {
    fmt.Println("update UDP Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[UpdateUDPListener更新UDP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#updateudplistener%E6%9B%B4%E6%96%B0udp%E7%9B%91%E5%90%AC%E5%99%A8)

### 查询UDP监听器

通过以下代码，查询指定LoadBalancer下所有UDP监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &blb.DescribeListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 53,
}
result, err := client.DescribeUDPListeners(BLBID, args)
if err != nil {
    fmt.Println("describe UDP Listener failed:", err)
} else {
    fmt.Println("describe UDP Listener success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeUDPListeners查询UDP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#describeudplisteners%E6%9F%A5%E8%AF%A2udp%E7%9B%91%E5%90%AC%E5%99%A8)

### 创建HTTP监听器

通过以下代码，在指定LoadBalancer下，创建一个基于HTTP协议的监听器，监听一个前端端口，将发往该端口的所有HTTP请求，根据策略转发到后端服务器监听的后端端口上
```go
args := &blb.CreateHTTPListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 80,
    // 后端服务器的监听端口，需要在1-65535之间
    BackendPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection
    Scheduler:    "RoundRobin",
}
err := client.CreateHTTPListener(BLBID, args)
if err != nil {
    fmt.Println("create HTTP Listener failed:", err)
} else {
    fmt.Println("create HTTP Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[CreateHTTPListener创建HTTP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#createhttplistener%E5%88%9B%E5%BB%BAhttp%E7%9B%91%E5%90%AC%E5%99%A8)

### 更新HTTP监听器

通过以下代码，更新指定LoadBalancer下的HTTP监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &blb.UpdateHTTPListenerArgs{
	// 要更新的监听器端口号
    ListenerPort: 80,
    // 更新负载均衡的算法
    Scheduler:    "LeastConnection", 
    // 开启会话保持功能
    KeepSession:  true,
}
err := client.UpdateHTTPListener(BLBID, args)
if err != nil {
    fmt.Println("update HTTP Listener failed:", err)
} else {
    fmt.Println("update HTTP Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[更新HTTP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#updatehttplistener%E6%9B%B4%E6%96%B0http%E7%9B%91%E5%90%AC%E5%99%A8)

### 查询HTTP监听器

通过以下代码，查询指定LoadBalancer下所有HTTP监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &blb.DescribeListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 80,
}
result, err := client.DescribeHTTPListeners(BLBID, args)
if err != nil {
    fmt.Println("describe HTTP Listener failed:", err)
} else {
    fmt.Println("describe HTTP Listener success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeHTTPListeners查询HTTP监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#describehttplisteners%E6%9F%A5%E8%AF%A2http%E7%9B%91%E5%90%AC%E5%99%A8)

### 创建HTTPS监听器

通过以下代码，在指定LoadBalancer下，创建一个基于HTTPS协议的监听器，监听一个前端端口，将发往该端口的所有HTTPS请求，先通过SSL卸载转换为HTTP请求后，再根据策略转发到后端服务器监听的后端端口上
```go
args := &blb.CreateHTTPSListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 443,
    // 后端服务器的监听端口，需要在1-65535之间
    BackendPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection
    Scheduler:    "RoundRobin", 
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.CreateHTTPSListener(BLBID, args)
if err != nil {
    fmt.Println("create HTTPS Listener failed:", err)
} else {
    fmt.Println("create HTTPS Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[CreateHTTPSListener创建HTTPS监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#createhttpslistener%E5%88%9B%E5%BB%BAhttps%E7%9B%91%E5%90%AC%E5%99%A8)

### 更新HTTPS监听器

通过以下代码，更新指定LoadBalancer下的HTTPS监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &blb.UpdateHTTPSListenerArgs{
	// 要更新的监听器端口号
    ListenerPort: 443,
    // 更新负载均衡的算法
    Scheduler:    "LeastConnection",
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.UpdateHTTPSListener(BLBID, args)
if err != nil {
    fmt.Println("update HTTPS Listener failed:", err)
} else {
    fmt.Println("update HTTPS Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[UpdateHTTPSListener更新HTTPS监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#updatehttpslistener%E6%9B%B4%E6%96%B0https%E7%9B%91%E5%90%AC%E5%99%A8)


### 查询HTTPS监听器

通过以下代码，查询指定LoadBalancer下所有HTTPS监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &blb.DescribeListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 443,
}
result, err := client.DescribeHTTPSListeners(BLBID, args)
if err != nil {
    fmt.Println("describe HTTPS Listener failed:", err)
} else {
    fmt.Println("describe HTTPS Listener success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeHTTPSListeners查询HTTPS监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#describehttpslisteners%E6%9F%A5%E8%AF%A2https%E7%9B%91%E5%90%AC%E5%99%A8)



### 创建SSL监听器

通过以下代码，在指定LoadBalancer下，创建一个基于SSL协议的blb监听器，监听一个前端端口，将发往该端口的所有SSL流量，根据策略进行转发
```go
args := &blb.CreateSSLListenerArgs{
    // 监听器监听的端口，需要在1-65535之间
    ListenerPort: 443,
    // 后端服务器的监听端口，需要在1-65535之间
    BackendPort: 80,
    // 负载均衡算法，支持RoundRobin/LeastConnection/Hash
    Scheduler:    "RoundRobin", 
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.CreateSSLListener(BLBID, args)
if err != nil {
    fmt.Println("create SSL Listener failed:", err)
} else {
    fmt.Println("create SSL Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[CreateSSLListener创建SSL监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#createssllistener%E5%88%9B%E5%BB%BAssl%E7%9B%91%E5%90%AC%E5%99%A8)



### 更新SSL监听器

通过以下代码，更新指定LoadBalancer下的SSL监听器参数，所有请求参数中指定的域都会被更新，未指定的域保持不变，监听器通过端口指定
```go
args := &blb.UpdateSSLListenerArgs{
	// 要更新的监听器端口号
    ListenerPort: 443,
    // 更新负载均衡的算法
    Scheduler:    "LeastConnection", 
    // 配置证书列表
    CertIds:      []string{certId},
}
err := client.UpdateSSLListener(BLBID, args)
if err != nil {
    fmt.Println("update SSL Listener failed:", err)
} else {
    fmt.Println("update SSL Listener success")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[UpdateSSLListener更新SSL监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#updatessllistener%E6%9B%B4%E6%96%B0ssl%E7%9B%91%E5%90%AC%E5%99%A8)


### 查询SSL监听器

通过以下代码，查询指定LoadBalancer下所有SSL监听器的信息，支持按监听器端口进行匹配查询，结果支持marker分页，分页大小默认为1000，可通过maxKeys参数指定
```go
args := &blb.DescribeListenerArgs{
    // 要查询的监听器端口
    ListenerPort: 443,
}
result, err := client.DescribeSSLListeners(BLBID, args)
if err != nil {
    fmt.Println("describe SSL Listener failed:", err)
} else {
    fmt.Println("describe SSL Listener success: ", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DescribeSSLListeners查询SSL监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#describessllisteners%E6%9F%A5%E8%AF%A2ssl%E7%9B%91%E5%90%AC%E5%99%A8)


### 删除监听器

通过以下代码，释放指定LoadBalancer下的监听器，监听器通过监听端口来指定，支持批量释放
```go
args := &blb.DeleteListenersArgs{
    ClientToken: "be31b98c-5e41-4838-9830-9be700de5a20",
    // 要删除的监听器监听的端口
    PortList:    []uint16{80, 443}, 
}
err := client.DeleteListeners(BLBID, args)
if err != nil {
    fmt.Println("delete Listeners failed:", err)
} else {
    fmt.Println("delete Listeners success: ")
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BLB API 文档[DeleteListeners释放监听器](https://cloud.baidu.com/doc/BLB/s/yjwvxnvl6#deletelisteners%E9%87%8A%E6%94%BE%E7%9B%91%E5%90%AC%E5%99%A8)


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
	fmt.Println("get blb detail success: ", blbDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向BLB发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当BLB服务端出现异常时，BLB服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[BLB错误返回](https://cloud.baidu.com/doc/BLB/s/Djwvxnzw6)

# 版本变更记录

## v0.9.11 [2020-05-20]

首次发布：

 - 创建、查看、列表、更新、删除普通型BLB实例
 - 创建、列表、更新、删除后端RS，并支持查询后端服务器健康检查状态
 - 创建、查看、更新、删除监听器端口，支持TCP/UDP/HTTP/HTTPS/SSL协议
