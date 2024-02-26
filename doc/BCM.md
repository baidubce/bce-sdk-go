# BCM服务

# 概述

本文档主要介绍BCM GO SDK的使用。在使用本文档前，您需要先了解
[BCM的基本知识](https://cloud.baidu.com/doc/BCM/index.html)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BCM域名](https://cloud.baidu.com/doc/BCM/s/5jwvym49g)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云BCM，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BCM做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建BCM Client

BCM Client是BCM控制面服务的客户端，为开发者与BCM控制面服务进行交互提供了一系列的方法。

### 使用AK/SK新建BCM Client

通过AK/SK方式访问BCM，用户可以参考如下代码新建一个BCM Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bcm"   //导入BCM服务模块
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BcmClient
	bcmClient, err := bcm.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/BCM/index.html)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BCM的控制面服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`bcm.bj.baidubce.com`。

### 使用STS创建BCM Client

**申请STS token**

BCM可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问BCM，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建BCM Client**

申请好STS后，可将STS Token配置到BCM Client中，从而实现通过STS Token创建BCM Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建BCM Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/bcm"    //导入BCM服务模块
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

	// 使用申请的临时STS创建BCM控制面服务的Client对象，Endpoint使用默认值
	bcmClient, err := bcm.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create bcm client failed:", err)
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
	bcmClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置BCM Client时，无论对应BCM服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问BCM

BCM支持HTTPS传输协议，您可以通过在创建BCM Client对象时指定的Endpoint中指明HTTPS的方式，在BCM GO SDK中使用HTTPS访问BCM服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

ENDPOINT := "https://bcm.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
bcmClient, _ := bcm.NewClient(AK, SK, ENDPOINT)
```

## 配置BCM Client

如果用户需要配置BCM Client的一些细节的参数，可以在创建BCM Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问BCM服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

//创建BCM Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcm.bj.baidubce.com
client, _ := bcm.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcm.bj.baidubce.com"
client, _ := bcm.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/bcm"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcm.bj.baidubce.com"
client, _ := bcm.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问BCM时，创建的BCM Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建BCM Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

云监控BCM（Cloud Monitor），是针对于云平台的监控服务，包括云产品监控、站点监控、自定义监控、应用监控、报警管理等产品功能，实时监控您云平台中的各种资源。

## 查询数据接口

### 接口描述
获取指定指标的一个或多个统计数据的时间序列数据。可获取云产品监控数据、站点监控数据或您推送的自定义监控数据。

### 接口限制
一次返回的数据点数目不能超过1440个。

### 请求示例

```go
dimensions := map[string]string{
    "InstanceId": "xxx",
}
req := &bcm.GetMetricDataRequest{
    UserId:         "xxx",
    Scope:          "BCE_BCC",
    MetricName:     "vCPUUsagePercent",
    Dimensions:     dimensions,
    Statistics:     strings.Split(bcm.Average, ","),
    PeriodInSecond: 60,
    StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
    EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
}
resp, err := bcmClient.GetMetricData(req)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BCM API 文档[查询数据接口](https://cloud.baidu.com/doc/BCM/s/9jwvym3kb)

## 批量查询数据接口

### 接口描述
可获取批量实例数据的接口，支持多维度多指标查找。可获取云产品监控数据、站点监控数据或您推送的自定义监控数据。

### 接口限制
一个实例的任意一个指标一次返回的数据点数目不能超过1440个。

### 请求示例
```go
dimensions := map[string]string{
    "InstanceId": "xxx",
}
req := &bcm.BatchGetMetricDataRequest{
    UserId:         "xxx",
    Scope:          "BCE_BCC",
    MetricNames:    []string{"vCPUUsagePercent", "CpuIdlePercent"},
    Dimensions:     dimensions,
    Statistics:     strings.Split(bcm.Average + "," + bcm.Sum, ","),
    PeriodInSecond: 60,
    StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
    EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
}
resp, err := bcmClient.BatchGetMetricData(req)
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BCM API 文档[查询数据接口](https://cloud.baidu.com/doc/BCM/s/9jwvym3kb)

## 创建名字空间

### 接口描述
创建一个自定义的名字空间，在该空间下可以创建自定义监控。

### 请求示例
```go
ns := &Namespace{
    Name:           "Test01",
    NamespaceAlias: "test",
    Comment:        "test",
    UserId:         "453bf9588c9e488f9ba2c9841290xxxx",
}
err := bcmClient.CreateNamespace(ns)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 删除名字空间

### 接口描述
删除自定义的名字空间

### 请求示例
```go
cns := &CustomBatchNames{
    Names:  []string{"Test01"},
    UserId: 453bf9588c9e488f9ba2c9841290xxxx,
}
err := bcmClient.BatchDeleteNamespaces(cns)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 编辑名字空间

### 接口描述
编辑某个自定义的名字空间

### 请求示例
```go
ns := &Namespace{
    Name:           "Test01",
    NamespaceAlias: "test1",
    Comment:        "test1",
    UserId:         453bf9588c9e488f9ba2c9841290xxxx,
}
err := bcmClient.UpdateNamespace(ns)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 获取（搜索）名字空间

### 接口描述
获取或者搜索自定义的名字空间列表

### 请求示例
```go
req := &ListNamespacesRequest{
    UserId:   453bf9588c9e488f9ba2c9841290xxxx,
    PageNo:   1,
    PageSize: 10,
}
res, err := bcmClient.ListNamespaces(req)
fmt.Println(res)
```
```go
req = &ListNamespacesRequest{
    UserId:   453bf9588c9e488f9ba2c9841290xxxx,
    Name:     "Test01",
    PageNo:   1,
    PageSize: 10,
}
res, err = bcmClient.ListNamespaces(req)
fmt.Println(res)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控创建指标

### 接口描述
在某个自定义空间创建自定义的指标

### 请求示例
```go
nm := &NamespaceMetric{
    UserId:      453bf9588c9e488f9ba2c9841290xxxx,
    Namespace:   "Test01",
    MetricName:  "TestMetric01",
    MetricAlias: "test",
    Unit:        "sec",
    Cycle:       60,
}
err := bcmClient.CreateNamespaceMetric(nm)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控删除(批量删除)指标

### 接口描述
在某个自定义空间删除自定义指标

### 请求示例
```go
cns := &CustomBatchIds{
    UserId:    453bf9588c9e488f9ba2c9841290xxxx,
    Namespace: "Test01",
    Ids:       []int64{1710},
}
err := bcmClient.BatchDeleteNamespaceMetric(cns)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控编辑指标

### 接口描述
在某个自定义空间编辑自定义的指标

### 请求示例
```go
nm := &NamespaceMetric{
    UserId:      453bf9588c9e488f9ba2c9841290xxxx,
    Namespace:   "Test01",
    MetricName:  "TestMetric01",
    MetricAlias: "test01",
    Unit:        "sec",
    Cycle:       60,
}
err := bcmClient.UpdateNamespaceMetric(nm)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控获取（搜索）指标

### 接口描述
获取或者搜索某个自定义的名字空间的指标列表

### 请求示例
```go
req := &ListNamespaceMetricsRequest{
    UserId:    453bf9588c9e488f9ba2c9841290xxxx,
    Namespace: "Test01",
    PageNo:    1,
    PageSize:  10,
}
res, err := bcmClient.ListNamespaceMetrics(req)
fmt.Println(res)
```
```go
req = &ListNamespaceMetricsRequest{
    UserId:      453bf9588c9e488f9ba2c9841290xxxx,
    Namespace:   "Test01",
    MetricName:  "TestMetric01",
    MetricAlias: "test",
    PageNo:      1,
    PageSize:    10,
}
res, err = bcmClient.ListNamespaceMetrics(req)
fmt.Println(res)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控指标详情

### 接口描述
获取某个自定义名字空间的指标详情

### 请求示例
```go
userId := "453bf9588c9e488f9ba2c9841290xxxx"
namespace := "Test01"
metricName := "TestMetric01"
nm, err := bcmClient.GetCustomMetric(userId, namespace, metricName)
fmt.Println(nm)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控创建事件

### 接口描述
在某个自定义空间创建自定义的事件

### 请求示例
```go
userId := "453bf9588c9e488f9ba2c9841290xxxx"
namespace := "Test01"
metricName := "TestMetric01"
nm, err := bcmClient.GetCustomMetric(userId, namespace, metricName)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控删除（批量删除）事件

### 接口描述
在某个自定义空间删除自定义事件

### 请求示例
```go
ces := &CustomBatchEventNames{
    UserId:    453bf9588c9e488f9ba2c9841290xxxx,
    Namespace: "Test01",
    Names:     []string{"TestEvent01"},
}
err := bcmClient.BatchDeleteNamespaceEvent(ces)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控编辑事件

### 接口描述
在某个自定义空间编辑自定义的事件

### 请求示例
```go
ne := &NamespaceEvent{
    UserId:         453bf9588c9e488f9ba2c9841290xxxx,
    Namespace:      "Test01",
    EventName:      "TestEvent01",
    EventNameAlias: "test01",
    EventLevel:     WarningEventLevel,
    Comment:        "event01",
}
err := bcmClient.UpdateNamespaceEvent(ne)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监获取（搜索）事件

### 接口描述
获取或者搜索某个自定义名字空间的事件列表

### 请求示例
```go
req := &ListNamespaceEventsRequest{
    UserId:    453bf9588c9e488f9ba2c9841290xxxx,
    Namespace: "Test01",
    PageNo:    1,
    PageSize:  10,
}
res, err := bcmClient.ListNamespaceEvents(req)
fmt.Println(res)
```
```go
req = &ListNamespaceEventsRequest{
    UserId:     453bf9588c9e488f9ba2c9841290xxxx,
    Namespace:  "Test01",
    Name:       "TestEvent01",
    EventLevel: WarningEventLevel,
    PageNo:     1,
    PageSize:   10,
}
res, err = bcmClient.ListNamespaceEvents(req)
fmt.Println(res)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 自定义监控事件详情

### 接口描述
获取某个自定义名字空间的事件详情

### 请求示例
```go
userId := "453bf9588c9e488f9ba2c9841290xxxx"
namespace := "Test01"
eventName := "TestEvent01"
ne, err := bcmClient.GetCustomEvent(userId, namespace, eventName)
fmt.Println(res)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[自定义监控接口](https://cloud.baidu.com/doc/BCM/s/cktnhszhv)

## 创建应用

### 接口描述
创建一个应用监控，在该应用下可以添加应用示例，创建监控任务，维度映射表以及查看相关监控数据。

### 请求示例
```go
req := &ApplicationInfoRequest{
Alias:       "test",
Name:        "test_1206",
Type:        "BCC",
UserId:      "453bf9588c9e488f9ba2c9841290****",
Description: "test",
}
res, err := bcmClient.CreateApplicationData(req)
fmt.Println(res)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 获取应用列表
### 接口描述
获取应用列表，支持分页查询以及名称模糊查询。

### 请求示例
```go
userId := "456cfa48356845b6bac1d29abc85****"
pageSize := 10
pageNo := 1
searchName := "test"
res, err := bcmClient.GetApplicationDataList(userId, searchName, pageSize, pageNo)
fmt.Println(res)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 更新应用信息
### 接口描述
更新应用信息。

### 请求示例
```go
res := &ApplicationInfoUpdateRequest{
    Id:          5336,
    Alias:       "test1206",
    Name:        "test_1206",
    Type:        "BCC",
    UserId:      "453bf9588c9e488f9ba2c9841290****",
    Description: "testD",
}
req, err := bcmClient.UpdateApplicationData(res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 删除应用信息
### 接口描述
删除应用信息。

### 请求示例
```go
res := &ApplicationInfoDeleteRequest{
    Name: "test_1206",
}
userId := "453bf9588c9e488f9ba2c9841290****"
err := bcmClient.DeleteApplicationData(userId, res)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 创建实例时实例列表查询
### 接口描述
创建实例时实例列表查询。支持分页查询以及名称模糊字段查询。searchName支持name，instanceId，internalIp字段查询。

### 请求示例
```go
res := &ApplicationInstanceListRequest{
    PageNo:      1,
    PageSize:    10,
    AppName:     "test-1130",
    SearchName:  "name",
    SearchValue: "",
    Region:      "bj",
}
userId := "456cfa48356845b6bac1d29abc85****"
req, err := bcmClient.GetApplicationInstanceList(userId, res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 创建实例
### 接口描述
创建实例，绑定到实例到应用下面。

### 请求示例
```go
infos := []*HostInstanceInfo{
    {
    InstanceId: "5cd74fe6-f508-4238-b20e-bddb62243a**",
    Region:     "bj",
    },
    {
    InstanceId: "2427ad4f-ac45-48b2-9a22-b92109b3fd**",
    Region:     "bj",
    },
}
res := &ApplicationInstanceCreateRequest{
    AppName:  "test-1130",
    UserId:   "456cfa48356845b6bac1d29abc85****",
    HostList: infos,
}
err := bcmClient.CreateApplicationInstance(res)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 查询已创建的实例列表
### 接口描述
查询已创建的实例

### 请求示例
```go
res := &ApplicationInstanceCreatedListRequest{
    UserId:  "456cfa48356845b6bac1d29abc85****",
    AppName: "test-1130",
    Region:  "bj",
}
req, err := bcmClient.GetApplicationInstanceCreatedList(res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 删除实例
### 接口描述
删除应用下面的实例

### 请求示例
```go
res := &ApplicationInstanceDeleteRequest{
    Id:      "6980",
    AppName: "test-1130",
}
userId := "456cfa48356845b6bac1d29abc85****"
err := bcmClient.DeleteApplicationInstance(userId, res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 创建监控任务
### 接口描述
创建监控任务

### 请求示例
```go
res := &ApplicationMonitorTaskInfoRequest{
    AppName:     "test-1130",
    Type:        0,
    AliasName:   "test-proc",
    Cycle:       60,
    Target:      "/proc",
    Description: "test-1207",
}
userId := "456cfa48356845b6bac1d29abc85****"
req, err := bcmClient.CreateApplicationInstanceTask(userId, res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 创建日志监控任务
### 接口描述
创建日志类型监控任务

### 请求示例
```go
extractResult := []*LogExtractResult{
{
    ExtractFieldName:  "namespace",
    ExtractFieldValue: "04b91096-a294-477d-bd11-1a7bcfb5a921",
    DimensionMapTable: "namespaceTable",
},
}
tags := []*AggTag{
{
    Range: "App",
    Tags:  "",
    },
    {
    Range: "App",
    Tags:  "namespace",
    },
}
metrics := []*Metric{
{
    MetricName:       "space",
    SaveInstanceData: 1,
    ValueFieldType:   0,
    MetricAlias:      "",
    MetricUnit:       "",
    ValueFieldName:   "",
    AggrTags:         tags,
},
}
res := &ApplicationMonitorTaskInfoLogRequest{
    AppName:       "test-1130",
    Type:          2,
    AliasName:     "test-log-1207",
    Cycle:         60,
    Target:        "/opt/bcm-agent/log/bcm-agent.INFO",
    Description:   "test-LOG-1207",
    Rate:          5,
    ExtractResult: extractResult,
    LogExample:    "namespace:04b91096-a294-477d-bd11-1a7bcfb5a921\n",
    MatchRule:     "namespace:(?P<namespace>[0-9a-fA-F-]+)",
    UserId:        "456cfa48356845b6bac1d29abc85****",
    Metrics:       metrics,
}
userId := "456cfa48356845b6bac1d29abc85****"
req, err := bcmClient.CreateApplicationMonitorLogTask(userId, res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 获取监控任务详情
### 接口描述
获取监控任务详情

### 请求示例
```go
res := &ApplicationMonitorTaskDetailRequest{
    AppName:  "test-1130",
    UserId:   "456cfa48356845b6bac1d29abc85****",
    TaskName: "d917e9963d6349909e3793b101e90333",
}
req, err := bcmClient.GetApplicationMonitorTaskDetail(res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 获取监控任务列表
### 接口描述
获取监控任务列表

### 请求示例
```go
res := &ApplicationMonitorTaskListRequest{
    AppName: "test-1130",
    UserId:  "456cfa48356845b6bac1d29abc85****",
}
req, err := bcmClient.GetApplicationMonitorTaskList(res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 更新监控任务信息
### 接口描述
更新监控任务信息，支持进程，端口，脚本类型任务

### 请求示例
```go
res := &ApplicationMonitorTaskInfoUpdateRequest{
    Name:        "456cfa48356845b6bac1d29abc85****",
    AppName:     "test-1130",
    Type:        0,
    AliasName:   "test-proc-update02",
    Cycle:       60,
    Target:      "/proc/exec",
    Description: "test-1207",
}
userId := "456cfa48356845b6bac1d29abc85****"
req, err := bcmClient.UpdateApplicationMonitorTask(userId, res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 更新日志监控任务信息
### 接口描述
更新日志类型监控任务信息
### 请求示例
```go
extractResult := []*LogExtractResult{
{
    ExtractFieldName:  "namespace",
    ExtractFieldValue: "04b91096-a294-477d-bd11-1a7bcfb5a921",
    DimensionMapTable: "namespaceTable",
},
}
tags := []*AggTag{
{
    Range: "App",
    Tags:  "",
    },
    {
    Range: "App",
    Tags:  "namespace",
    },
}
metrics := []*Metric{
{
    MetricName:       "space",
    SaveInstanceData: 1,
    ValueFieldType:   0,
    MetricAlias:      "",
    MetricUnit:       "",
    ValueFieldName:   "",
    AggrTags:         tags,
},
}
res := &ApplicationMonitorTaskInfoUpdateRequest{
    Id:            "3922",
    Name:          "d917e9963d6349909e3793b101e90333",
    AppName:       "test-1130",
    Type:          2,
    AliasName:     "test-log-1207",
    Cycle:         60,
    Target:        "/opt/bcm-agent/log/bcm-agent.INFO",
    Description:   "test-LOG1207",
    Rate:          5,
    ExtractResult: extractResult,
    LogExample:    "namespace:04b91096-a294-477d-bd11-1a7bcfb5a921\n",
    MatchRule:     "namespace:(?P<namespace>[0-9a-fA-F-]+)",
    UserId:        456cfa48356845b6bac1d29abc85****,
    Metrics:       metrics,
}
userId := "456cfa48356845b6bac1d29abc85****"
req, err := bcmClient.UpdateApplicationMonitorLogTask(userId, res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 删除监控任务
### 接口描述
删除监控任务

### 请求示例
```go
    res := &ApplicationMonitorTaskDeleteRequest{
        Name:    "d917e9963d6349909e3793b101e90333",
        AppName: "test-1130",
        UserId:  456cfa48356845b6bac1d29abc85****,
}
err := bcmClient.DeleteApplicationMonitorTask(res)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 创建维度映射表
### 接口描述
创建维度映射表

### 请求示例
```go
 res := &ApplicationDimensionTableInfoRequest{
        UserId:         456cfa48356845b6bac1d29abc85****,
        AppName:        "test-1130",
        TableName:      "test-table",
        MapContentJson: "a=>1\\nb=>2",
}
req, err := bcmClient.CreateApplicationDimensionTable(res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 查询维度映射表
### 接口描述
查询维度映射表列表详细，支持名称字段模糊查询

### 请求示例
```go
 res := &ApplicationDimensionTableListRequest{
    UserId:     456cfa48356845b6bac1d29abc85****,
    AppName:    "test-1130",
    SearchName: "test",
}
req, err := bcmClient.GetApplicationDimensionTableList(res)
fmt.Println(req)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 更新维度映射表
### 接口描述
更新维度映射表信息

### 请求示例
```go
res := &ApplicationDimensionTableInfoRequest{
    UserId:         456cfa48356845b6bac1d29abc85****,
    AppName:        "test-1130",
    TableName:      "test-table",
    MapContentJson: "a=>1",
}
err := bcmClient.UpdateApplicationDimensionTable(res)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)

## 删除维度映射表
### 接口描述
删除维度映射表

### 请求示例
```go
res := &ApplicationDimensionTableDeleteRequest{
    UserId:    456cfa48356845b6bac1d29abc85****,
    AppName:   "test-1130",
    TableName: "test-table",
}
err := bcmClient.DeleteApplicationDimensionTable(res)
```
> **提示：**
> 详细的参数配置及使用，可以参考BCM API 文档[应用监控接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto)



# 通知模版

## 查询用户组列表
### 接口描述
查询用户组列表
### 请求示例
```go
req := &model.ListNotifyGroupsRequest{
	Name:     "test",
	PageNo:   1,
	PageSize: 5,
}
res, err := bcmClient.ListNotifyGroups(req)
fmt.Println(res)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[查询用户组列表](https://cloud.baidu.com/doc/BCM/s/elmiysvfo#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E7%BB%84%E5%88%97%E8%A1%A8)

## 查询用户列表
### 接口描述
查询用户列表
### 请求示例
```go
req := &model.ListNotifyPartiesRequest{
	Name:     "test",
	PageNo:   1,
	PageSize: 5,
}
res, err := bcmClient.ListNotifyParty(req)
fmt.Println(res)
fmt.Println(err)
```

> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[查询用户列表](https://cloud.baidu.com/doc/BCM/s/elmiysvfo#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8)

## 新建通知模版
### 接口描述
新建通知模版
### 请求示例
```go
notification := model.ActionNotification{
	Type:     model.ActionNotificationTypeEmail,
	Receiver: "",
}
member := model.ActionMember{
	Type: "notifyParty",
	Id:   "56c9e0e2138c4f",
	Name: "lzs",
}
req := &model.CreateActionRequest{
	UserId:          "453bf9588c9e488f9ba2c98412******",
	Notifications:   []model.ActionNotification{notification},
	Members:         []model.ActionMember{member},
	Alias:           "test_wjr",
	DisableTimes:    nil,
	ActionCallBacks: nil,
}
err := bcmClient.CreateAction(req)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[新建通知模版](https://cloud.baidu.com/doc/BCM/s/elmiysvfo#%E6%96%B0%E5%BB%BA%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88)


## 删除通知模版
### 接口描述
删除通知模版
### 请求示例
```go
req := &model.DeleteActionRequest{
	UserId: "453bf9588c9e488f9ba2c98412******",
	Name:   "b90d86da-e3a0-4c63-9bd2-a7e210d2027f",
}
err := bcmClient.DeleteAction(req)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[删除通知模版](https://cloud.baidu.com/doc/BCM/s/elmiysvfo#%E5%88%A0%E9%99%A4%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88)


## 查询通知模版列表
### 接口描述
查询通知模版列表
### 请求示例
```go
req := &model.ListActionsRequest{
	UserId:   "453bf9588c9e488f9ba2c98412******",
	PageNo:   1,
	PageSize: 10,
}
resp, err := bcmClient.ListActions(req)
fmt.Println(*resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[查询通知模版列表](https://cloud.baidu.com/doc/BCM/s/elmiysvfo#%E6%9F%A5%E8%AF%A2%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88%E5%88%97%E8%A1%A8)

## 编辑通知模版
### 接口描述
编辑通知模版
### 请求示例
```go
notification := model.ActionNotification{
	Type:     model.ActionNotificationTypeEmail,
	Receiver: "",
}
member := model.ActionMember{
	Type: "notifyParty",
	Id:   "453bf9588c9e488f9ba2c98412******",
	Name: "lzs",
}
req := &model.UpdateActionRequest{
	UserId:          "453bf9588c9e488f9ba2c98412******",
	Name:            "4e9630d6-6348-450d-aab8-ea5f40******",
	Notifications:   []model.ActionNotification{notification},
	Members:         []model.ActionMember{member},
	Alias:           "test_wjr",
	DisableTimes:    nil,
	ActionCallBacks: nil,
}
err := bcmClient.UpdateAction(req)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[编辑通知模版](https://cloud.baidu.com/doc/BCM/s/elmiysvfo#%E7%BC%96%E8%BE%91%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88)

# 应用监控日志接口
## 日志提取接口
### 接口描述
日志提取接口
### 请求示例
```go
req := &model.LogExtractRequest{
	UserId:      "453bf9588c9e488f9ba2c98412******",
	ExtractRule: "800] \"(?<method>(GET|POST|PUT|DELETE)) .*/v1/dashboard/metric/(?<widget>(cycle|trend|report|billboard|gaugechart)) HTTP/1.1\".* (?<resTime>[0-9]+)ms",
	LogExample:  "10.157.16.207 - - [09/Apr/2020:20:45:33 +0800] \"POST /v1/dashboard/metric/gaugechart HTTP/1.1\" 200 117 109ms\n10.157.16.207 - - [09/Apr/2020:20:45:33 +0800] \"GET /v1/dashboard/metric/report HTTP/1.1\" 200 117 19ms",
}
resp, err := bcmClient.LogExtract(req)
fmt.Println(resp)
fmt.Println(err)
```

> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[日志提取接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E6%97%A5%E5%BF%97%E6%8F%90%E5%8F%96%E6%8E%A5%E5%8F%A3)

# 应用监控监控数据查询接口
## 维度值查询接口
### 接口描述
维度值查询接口
### 请求示例
```go
req := &model.GetMetricMetaForApplicationRequest{
	UserId:        "453bf9588c9e488f9ba2c98412******",
	AppName:       "test14",
	TaskName:      "task13",
	MetricName:    "log.responseTime",
	Instances:     []string{"0.test14"},
	DimensionKeys: []string{"method"},
}
resp, err := bcmClient.GetMetricMetaForApplication(req)
fmt.Println(resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[维度值查询接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E7%BB%B4%E5%BA%A6%E5%80%BC%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3)

## 多监控对象-单指标查询接口
### 接口描述
多监控对象-单指标查询接口
### 请求示例
```go
req := &model.GetMetricDataForApplicationRequest{
	UserId:     "453bf9588c9e488f9ba2c98412******",
	AppName:    "gjm-test",
	TaskName:   "bbceac2807014fce920e92e31debf092",
	MetricName: "port.err_code",
	Instances:  []string{"8.gjm-test"},
	StartTime:  "2023-12-07T01:10:48Z",
	EndTime:    "2023-12-07T02:10:48Z",
	Cycle:      0,
	Statistics: []string{"average"},
	Dimensions: nil,
	AggrData:   false,
}
resp, err := bcmClient.GetMetricDataForApplication(req)
fmt.Println(resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[多监控对象-单指标查询接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E5%A4%9A%E7%9B%91%E6%8E%A7%E5%AF%B9%E8%B1%A1-%E5%8D%95%E6%8C%87%E6%A0%87%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3)

# 应用监控报警相关接口

## 报警策略创建接口
### 接口描述
创建应用监控报警策略
### 请求示例
```go
req := &model.AppMonitorAlarmConfig{
	AlarmDescription:  "",
	AlarmName:         "test_wjr",
	UserId:            "453bf9588c9e488f9ba2c98412******",
	AppName:           "zmq-log-1115",
	MonitorObjectType: "APP",
	MonitorObject: model.MonitorObject{
		Id: 4030,
		MonitorObjectView: []model.MonitorObjectViewModel{{
			MonitorObjectName: "ab3b543f41974e26ab984d94fc7b9b92",
		}},
		MonitorObjectType: "APP",
	},
	SrcName:       "ab3b543f41974e26ab984d94fc7b9b92",
	SrcType:       "LOG",
	Type:          "INSTANCE",
	Level:         model.AlarmLevelMajor,
	ActionEnabled: true,
	PolicyEnabled: true,
	Rules: [][]model.AppMonitorAlarmRule{[]model.AppMonitorAlarmRule{{
		Metric:             "log.log_metric2",
		MetricAlias:        "log_metric2",
		Cycle:              60,
		Statistics:         "average",
		Threshold:          5,
		ComparisonOperator: ">",
		Count:              1,
		Function:           "THRESHOLD",
		Sequence:           0,
	}}},
	IncidentActions:     []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
	ResumeAction:        []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
	InsufficientActions: []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
	RepeatAlarmCycle:    300,
	MaxRepeatCount:      1,
}
resp, err := bcmClient.CreateAppMonitorAlarmConfig(req)
fmt.Println(*resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[报警策略创建接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E6%8A%A5%E8%AD%A6%E7%AD%96%E7%95%A5%E5%88%9B%E5%BB%BA%E6%8E%A5%E5%8F%A3)

## 报警策略更新接口
### 接口描述
更新应用监控报警策略
### 请求示例
```go
req := &model.AppMonitorAlarmConfig{
	AlarmDescription:  "",
	AlarmName:         "test_wjr",
	UserId:            "453bf9588c9e488f9ba2c98412******",
	AppName:           "zmq-log-1115",
	MonitorObjectType: "APP",
	MonitorObject: model.MonitorObject{
		Id: 4030,
		MonitorObjectView: []model.MonitorObjectViewModel{{
			MonitorObjectName: "ab3b543f41974e26ab984d94fc7b9b92",
		}},
		MonitorObjectType: "APP",
	},
	SrcName:       "ab3b543f41974e26ab984d94fc7b9b92",
	SrcType:       "LOG",
	Type:          "INSTANCE",
	Level:         model.AlarmLevelMajor,
	ActionEnabled: true,
	PolicyEnabled: true,
	Rules: [][]model.AppMonitorAlarmRule{[]model.AppMonitorAlarmRule{{
		Metric:             "log.log_metric2",
		MetricAlias:        "log_metric2",
		Cycle:              60,
		Statistics:         "average",
		Threshold:          5,
		ComparisonOperator: ">",
		Count:              1,
		Function:           "THRESHOLD",
		Sequence:           0,
	}}},
	IncidentActions:     []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
	ResumeAction:        []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
	InsufficientActions: []string{"624c99b5-5436-478c-8326-0efc8163c7d5"},
	RepeatAlarmCycle:    300,
	MaxRepeatCount:      1,
}
resp, err := bcmClient.UpdateAppMonitorAlarmConfig(req)
fmt.Println(*resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[报警策略更新接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E6%8A%A5%E8%AD%A6%E7%AD%96%E7%95%A5%E6%9B%B4%E6%96%B0%E6%8E%A5%E5%8F%A3)

## 报警策略列表接口
### 接口描述
请求应用监控报警策略列表
### 请求示例
```go
req := &model.ListAppMonitorAlarmConfigsRequest{
	UserId:    "453bf9588c9e488f9ba2c98412******",
	AlarmName: "test",
	SrcType:   model.SrcTypeProc,
	PageNo:    1,
	PageSize:  10,
}
resp, err := bcmClient.ListAppMonitorAlarmConfigs(req)
fmt.Println(*resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[报警策略列表接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E6%8A%A5%E8%AD%A6%E7%AD%96%E7%95%A5%E5%88%97%E8%A1%A8%E6%8E%A5%E5%8F%A3)

## 报警策略删除接口
### 接口描述
删除某一应用监控报警策略
### 请求示例
```go
req := &model.DeleteAppMonitorAlarmConfigRequest{
	UserId:    "453bf9588c9e488f9ba2c98412******",
	AppName:   "zmq-log-1115",
	AlarmName: "test_wjr",
}
err := bcmClient.DeleteAppMonitorAlarmConfig(req)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[报警策略删除接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E6%8A%A5%E8%AD%A6%E7%AD%96%E7%95%A5%E5%88%A0%E9%99%A4%E6%8E%A5%E5%8F%A3)

## 报警策略详情接口
### 接口描述
查询某一应用监控报警策略的详情
### 请求示例
```go
req := &model.GetAppMonitorAlarmConfigDetailRequest{
	UserId:    "453bf9588c9e488f9ba2c98412******",
	AlarmName: "config-yyy",
	AppName:   "yyy-test",
}
resp, err := bcmClient.GetAppMonitorAlarmConfig(req)
fmt.Println(*resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[报警策略详情接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E6%8A%A5%E8%AD%A6%E7%AD%96%E7%95%A5%E8%AF%A6%E6%83%85%E6%8E%A5%E5%8F%A3)

## 报警指标列表接口
### 接口描述
查询某一应用监控某监控任务的指标列表
### 请求示例
```go
req := &model.GetAppMonitorAlarmMetricsRequest{
	UserId:   bcmConf.UserId,
	AppName:  "gjm-app-test",
	TaskName: "eecb448f083447498012cff4473c7ea1",
}
resp, err := bcmClient.GetAlarmMetricsForApplication(req)
fmt.Println(resp)
fmt.Println(err)
```
> **提示：**
> - 详细的参数配置及使用，可以参考BCM API 文档[报警指标列表接口](https://cloud.baidu.com/doc/BCM/s/Sku56poto#%E6%8A%A5%E8%AD%A6%E6%8C%87%E6%A0%87%E5%88%97%E8%A1%A8%E6%8E%A5%E5%8F%A3)


## 客户端异常

客户端异常表示客户端尝试向BCM发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当BCM服务端出现异常时，BCM服务端会返回给用户相应的错误信息，以便定位问题

