CDN服务（海外版，百度智能云与其它公司合作建设的CDN产品）

# 概述

本文档主要介绍CDN GO SDK的使用。在使用本文档前，您需要先了解CDN的一些基本知识，并已开通了百度智能云海外CDN服务。若您还不了解CDN，可以参考[产品简介](https://cloud.baidu.com/doc/CDN-ABROAD/s/Fjwvxiui7) 和 [产品功能](https://cloud.baidu.com/doc/CDN-ABROAD/s/jjwvxiu3u) 。

注意：当前百度智能云海外版CDN为白名单开通，如有需求请工单联系百度侧人员。

# 初始化

## 确认Endpoint

目前使用CDN服务时，CDN的 Endpoint 统一使用`https://cdn.baidubce.com`，这也是默认值。

## 获取密钥

要使用百度云CDN，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问CDN做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 使用AK/SK新建CDN Client

通过AK/SK方式访问CDN，用户可以参考如下代码新建一个CDN Client：

```go
package main

import (
	"github.com/baidubce/bce-sdk-go/services/cdn/abroad"
)

func main() {
	ak := "your_access_key_id"
	sk := "your_secret_key_id"
	endpoint := "https://cdn.baidubce.com"
	client, err := abroad.NewClient(ak, sk, endpoint)
	if err != nil {
		panic("create BCE abroad CDN client failed: " + err.Error())
	}

	// TODO: Now you can use this client object to do something.
	_ = client
}
```

在上面代码中，变量`ak`对应控制台中的“Access Key ID”，变量`sk`对应控制台中的“Access Key Secret”，获取方式请参考文档：[如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/) 。变量`endpoint`必须为`https://cdn.baidubce.com`，也是默认值，为空表示使用默认值，设置为其他会导致 SDK 无法工作。

在下面的示例中，会频繁使用到GetDefaultClient函数，它的定义为：

```go
func GetDefaultClient() *abroad.Client {
	ak := "your_access_key_id"
	sk := "your_secret_key_id"
	endpoint := "https://cdn.baidubce.com"

	// Here ignore error in test, but you should handle error in your production environment.
	client, _ := abroad.NewClient(ak, sk, endpoint)
	return client
}
```

## 域名操作

### 域名列表查询 ListDomains

**> 查询该用户账户下拥有的加速域名**

```go
testCli := client.GetDefaultClient()
marker := ""
domains, nextMarker, err := testCli.ListDomains(marker)

fmt.Printf("domains:%+v\n", domains)
fmt.Printf("nextMarker:%s\n", nextMarker)
fmt.Printf("err:%+v\n", err)
```

`marker`参数表示从这个字符串所代表的索引开始查询域名列表，空表示从头开始。目前服务器处理分页size是一个很大的数，所以使用的时候将marker赋值为空即可。ListDomains返回的`nextMarker`表示下一个域名列表索引 ，空表示从`marker`开始之后的域名列表全部被获取，非空可用于传递到ListDomains参数。`domains`是一个string slice，表示域名列表数据，如: `["1.baidu.com", "2.baidu.com"]`

### 创建加速域名接口 CreateDomain

**> 添加一个加速域名到用户名下。创建加速域名必须设置源站。**

```go
domainCreatedInfo, err := testCli.CreateDomain(testDomain, []api.OriginPeer{
		{
			Type:   "IP",
			Backup: false,
			Addr:   "1.1.1.1",
		},
	})
fmt.Printf("domainCreatedInfo:%+v\n", domainCreatedInfo)
fmt.Printf("err:%+v\n", err)
```

源站信息**api.OriginPeer**的结构如下：

| 字段      | 类型   | 说明                                                           |
| --------- | ------ | ------------------------------------------------------------ |
| type      | string | 源站的类型，取值为IP或DOMAIN，表示为ip类型源站或域名类型源站。        |
| backup    | bool   | 表示是否为备用源站，备用源站在主源站都不可用的情况下才会使用，false表示这个源站是主源站。需要注意的是在golang中bool对象默认值为false，如果没有显式设置Backup的值，那么默认就是false。注意：至少要有一个主源站，否则会报错。       |
| addr      | string   | 源站地址                                                     |


另外，您也可以使用更高级的 **CreateDomainWithOptions** 方法创建域名，目前支持的Option为：
- CreateDomainWithTags：绑定域名到标签

下面通过代码展示如何调用 CreateDomainWithOptions。

```go
domainCreatedInfo, err := testCli.CreateDomainWithOptions(testDomain, []api.OriginPeer{
		{
			Type:   "IP",
			Backup: false,
			Addr:   "1.1.1.1",
		},
	}, CreateDomainWithTags([]model.TagModel{
		{
			TagKey:   "service",
			TagValue: "web",
		},
		{
			TagKey:   "域名类型",
			TagValue: "网站服务",
		},
	}))

fmt.Printf("domainCreatedInfo:%+v\n", domainCreatedInfo)
fmt.Printf("err:%+v\n", err)
```


### 启用/停止/删除加速域名 EnableDomain/DisableDomain/DeleteDomain

```go
// 启用加速域名
err := testCli.EnableDomain(testDomain)
fmt.Printf("err:%+v\n", err)

// 停用加速域名
err = testCli.DisableDomain(testDomain)
fmt.Printf("err:%+v\n", err)

// 删除该用户名下的加速域名
err = testCli.DeleteDomain(testDomain)
fmt.Printf("err:%+v\n", err)
```

## 域名配置

### 查询加速域名详情 GetDomainConfig

**> 查询某个特定域名的全部配置信息**

```go
domainConfig, err := testCli.GetDomainConfig(testDomain)
if err != nil {
	fmt.Printf("GetDomainConfig for %s failed: %s", testDomain, err)
	return
}
fmt.Printf("GetDomainConfig for %s got: %+v", testDomain, domainConfig)
```

`domainConfig`是`*api.DomainConfig`类型的对象，他的结构相对比较复杂，如下所示：

| 字段             | 类型               | 说明                                   |
|----------------|------------------|--------------------------------------|
| Domain         | string           | 域名信息，如：`test.baidu.com`。             |
| Origin         | []api.OriginPeer | 源站信息，在创建加速域名接口有详细说明。                 |
| CacheTTL       | []api.CacheTTL   | 缓存过期规则，在`设置缓存过期规则`有详细说明。             |
| CacheFullUrl   | bool             | 参数参与缓存规则，在`设置参数参与缓存规则`有详细说明。         |
| OriginHost     | *string          | 回源host，默认为nil，代表着未设置过该值。             |
| RefererACL     | *api.RefererACL  | referer黑白名单，在`设置 referer 黑白名单`有详细说明。 |
| IpAcl          | *api.IpACL       | 客户端IP黑白名单，在`设置客户端 IP 黑白名单`有详细说明。     |
| HTTPSConfig    | *api.HTTPSConfig | HTTPS 配置，在`设置 HTTPS 配置`   有详细说明。     |
| OriginProtocol | string           | 回源协议，`http`或`https`。                 |
| Tags           | []model.TagModel | 域名所属的标签。                             |


### 更新加速域名回源地址 SetDomainOrigin

**> 设置 DOMAIN 类型源站**

```go
originPeers := []api.OriginPeer{
	{
		Type:   "DOMAIN",
		Addr:   "test.badiu.com",
		Backup: false,
	},
}
err := testCli.SetDomainOrigin(testDomain, originPeers)
if err != nil {
	fmt.Printf("SetDomainOrigin for %s failed: %s\n", testDomain, err)
	return
}
fmt.Printf("SetDomainOrigin successfully.\n")
```

**> 设置 IP 类型源站**

```go
originPeers := []api.OriginPeer{
	{
		Type:   "IP",
		Addr:   "220.181.38.251",
		Backup: false,
	},
	{
		Type:   "IP",
		Addr:   "220.181.38.148",
		Backup: false,
	},
}
err := testCli.SetDomainOrigin(testDomain, originPeers)
if err != nil {
	fmt.Printf("SetDomainOrigin for %s failed: %s\n", testDomain, err)
	return
}
fmt.Printf("SetDomainOrigin successfully.\n")
```

`api.OriginPeer`类型的详细说明在**创建加速域名一节**已经有说明。


### 设置缓存过期规则 SetCacheTTL

**> 设置缓存规则为所有资源都缓存3600s，并且覆盖源站的缓存规则**

```go
cacheTTls := []api.CacheTTL{
	{
		Type:           "path",
		Value:          "/",
		TTL:            3600,
		Weight:         100,
		OverrideOrigin: true,
	},
}
err := testCli.SetCacheTTL(testDomain, cacheTTls)
if err != nil {
	fmt.Printf("SetCacheTTL for %s failed: %s\n", testDomain, err)
    return
}
fmt.Printf("SetCacheTTL successfully.\n")
```

`cacheTTls`包含全部的缓存策略，每一个缓存策略用`api.CacheTTL`类型的对象表示，如下关于缓存策略结构的说明。

| 字段   | 类型   | 说明                                                                                                                                                                              |
| ------ | ------ |---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Type   | string | 缓存策略的类型。合法的类型有：`suffix`、`path` 和`exactPath`。"suffix"表示文件名后缀，"path"表示url中的目录 ，“exactPath”表示路径完全匹配。 |
| Value  | string | Type所指定类型的配置规则。                                                                                                                                                                 |
| Weight | int    | 权重，0-100的整数，权重越高优先级越高，默认为0，权重越大，优先级越高，规则优先生效。不推荐两条缓存策略配置相同的权重，如果权重相同，会随机选择其中一条策略生效。                                                                     |
| TTL    | int    | 缓存时间，单位为秒。                                                                                                                                                                      |


### 设置参数参与缓存规则 SetCacheFullUrl

**> 设置URL参数参与缓存key计算**

```go
err := testCli.SetCacheFullUrl(testDomain, true)
if err != nil {
	fmt.Printf("SetCacheFullUrl for %s failed: %s\n", testDomain, err)
	return
}
fmt.Printf("SetCacheFullUrl successfully.\n")
```

### 设置回源域名 SetHostToOrigin

**> 设置域名的回源host为 test.baidu.com**

```go
err := testCli.SetHostToOrigin(testDomain, "test.baidu.com")
if err != nil {
	fmt.Printf("SetHostToOrigin for %s failed: %s\n", testDomain, err)
	return
}
fmt.Printf("SetHostToOrigin successfully.\n")
```

### 设置 referer 黑白名单 SetRefererACL

**> 设置黑名单referer为bad.baidu.com，且允许空referer**

```go
err := testCli.SetRefererACL(testDomain, &api.RefererACL{
	BlackList:  []string{"bad.baidu.com"},
	AllowEmpty: true,
})
if err != nil {
	fmt.Printf("SetRefererACL for %s failed: %s\n", testDomain, err)
	return
}
fmt.Printf("SetRefererACL successfully.\n")
```

`AllowEmpty`表示是否允许空Referer访问，true表示允许空Referer访问，即如果Referer为空，那么不管是设置了黑名单还是白名单都不会生效，大多数情况下这个值都设置为**true**；false表示不允许空Referer访问，当设置为false时，如果访问的HTTP报文中不存在Referer那么CDN系统将返回403。`refererACL`是`api.RefererACL`类型的对象，他的详细说明如下：

| 字段       | 类型     | 说明                                                         |
| ---------- | -------- | ------------------------------------------------------------ |
| BlackList  | []string | 可选项，表示referer黑名单列表，支持使用通配符*，**不需要加protocol**，如设置某个黑名单域名，设置为"www.xxx.com"形式即可，而不是"http://www.xxx.com"。 |
| WhiteList  | []string | *可选项，list类型，表示referer白名单列表，支持通配符*，同样不需要加protocol。 |
| AllowEmpty | bool     | 必选项，bool类型，表示是否允许空referer访问，为true即允许空referer访问。 |

### 设置客户端 IP 黑白名单 SetIpACL

**> 设置客户端IP黑名单为 220.181.38.148**

```go
err := testCli.SetIpACL(testDomain, &api.IpACL{
	BlackList: []string{"220.181.38.148"},
})
if err != nil {
	fmt.Printf("SetIpACL for %s failed: %s\n", testDomain, err)
	return
}
fmt.Printf("SetIpACL successfully.\n")
```

`ipACL`是`api.IpACL`类型对象，详细说明如下：

| 字段      | 类型     | 说明                                                         |
| --------- | -------- | ------------------------------------------------------------ |
| BlackList | []string | IP黑名单列表，当设置黑名单生效时，当客户端的IP属于BlackList，CDN系统返回403。BlackList不可与WhiteList同时设置。 |
| WhiteList | []string | IP白名单列表，当设置白名单生效时，当WhiteList为空时，没有白名单效果。当WhiteList非空时，只有客户端的IP属于WhiteList才允许访问。同样不可与BlackList同时设置。 |


### 设置 HTTPS 配置 SetHTTPSConfigWithOptions/SetHTTPSConfig

**> 使用 SetHTTPSConfig 开启HTTPS**

```go
var err error

// 开启 HTTPS
var certId = "cert-4xkhw3m73hxs"
err = testCli.SetHTTPSConfig(testDomain, &api.HTTPSConfig{
	Enabled:      true,
	CertId:       certId,
	HttpRedirect: true,
	Http2Enabled: true,
})
if err != nil {
	fmt.Printf("SetHTTPSConfig enable HTTPS failed: %s\n", err)
	return
}
fmt.Printf("SetHTTPSConfig enable HTTPS successfully\n")
```

**> 使用 SetHTTPSConfig 关闭HTTPS**

```go
var err error

// 关闭 HTTPS
err = testCli.SetHTTPSConfig(testDomain, &api.HTTPSConfig{
	Enabled: false,
})
if err != nil {
	fmt.Printf("SetHTTPSConfig disable HTTPS failed: %s\n", err)
	return
}
fmt.Printf("SetHTTPSConfig disable HTTPS successfully\n")
```


上述例子中，`api.HTTPSConfig`的详细说明如下：

| 字段      | 类型     | 说明                                                                                                                         |
|---------|--------|----------------------------------------------------------------------------------------------------------------------------|
| Enabled | bool   | true 表示开启HTTPS，certId必须有效且真实存在                                                                                             |                                                                                         |
| CertId  | string | 证书ID，百度智能云统一通过证书中心进行证书管理，您必须先通过 [证书管理中心](https://console.bce.baidu.com/cas/#/cas/purchased/common/list)  进行证书上传/购买，才能获得证书ID。 |
|    HttpRedirect     | bool   | 为true时将HTTP请求重定向到HTTPS（重定向状态码为301)。                                                                                        |
|         Http2Enabled            | bool   | true 表示开启HTTP2。                                                                                                            |



**> 使用 SetHTTPSConfigWithOptions 开启HTTPS**

```go
var err error

// 开启 HTTPS
var certId = "cert-4xkhw3m73hxs"
err = testCli.SetHTTPSConfigWithOptions(testDomain, true,
	api.HTTPSConfigCertID(certId),    // 必选
	api.HTTPSConfigRedirectWith301(), // 可选
	api.HTTPSConfigEnableH2(),        // 可选
)
if err != nil {
	fmt.Printf("SetHTTPSConfigWithOptions enable HTTPS failed: %s\n", err)
	return
}
fmt.Printf("SetHTTPSConfigWithOptions enable HTTPS successfully\n")
```

上述例子展示了开启 HTTPS，使用`api.HTTPSConfigCertID`进行证书传递，通过`api.HTTPSConfigRedirectWith301`开启将HTTP请求重定向到HTTPS，通过`api.HTTPSConfigEnableH2`开启HTTP2。


**> 使用 SetHTTPSConfigWithOptions 关闭HTTPS**

```go
var err error

// 关闭 HTTPS
err = testCli.SetHTTPSConfigWithOptions(testDomain, false)
if err != nil {
	fmt.Printf("SetHTTPSConfigWithOptions disable HTTPS failed: %s\n", err)
	return
}
fmt.Printf("SetHTTPSConfigWithOptions disable HTTPS successfully\n")
```

### 设置回源协议 SetOriginProtocol

**> 设置回源协议为HTTPS**

```go
err := testCli.SetOriginProtocol(testDomain, api.HTTPSOrigin)
if err != nil {
	fmt.Printf("SetOriginProtocol for %s failed: %s\n", testDomain, err)
	return
}
fmt.Printf("SetOriginProtocol successfully.\n")
```

### 管理域名标签 SetTags/GetTags

> 标签是一组标识，您可以将具备某些共同点的域名进行归类，即打上标签，以此对资源进行组织。

一组标签定义为tags，类型如下：

| 字段  | 类型              | 说明  |
| tags | model.TagModel[] | 目标关联到的标签集合，需传入资源的全量标签信息，如原有3个标签，现删除1个标签的情况下新增2个标签，则应传入全量4个标签。当 tags 为空时，表示将域名关联的所有标签清空。标签个数最多为 10 个。 |


model.TagModel 类型如下：

| 参数             | 类型   | 说明                                                          |
| ----------------| ------ | ----------------------------------------------------------  |
| tagKey          | string |  标签的键，可包含大小写字母、数字、中文以及-_ /.特殊字符，长度 1-65  |
| tagValue        | string | 标签的值，可包含大小写字母、数字、中文以及-_ /.特殊字符，长度 0-65     |


下面展示通过SDK方法配置标签与查询标签。

```go
// 配置标签
err := testCli.SetTags(testDomain, []model.TagModel{
		{
			TagKey:   "service",
			TagValue: "download",
		},
	})
fmt.Printf("err: %+v\n", err)

// 查询标签
tags, err := testCli.GetTags(testDomain)
fmt.Printf("tags: %+v\n", tags)

// 移除所有标签
err := testCli.SetTags(testDomain, []model.TagModel{}) // 等同于 testCli.SetTags(testDomain, nil)
fmt.Printf("err: %+v\n", err)
```

## 缓存管理

`api.CStatusQueryData` 是查询条件所抽象出来的结构体，它的详细说明如下：

| 字段        | 类型     | 说明                                               |
|-----------|--------|--------------------------------------------------|
| Id        | string | 通过任务ID查询任务状态，当设置此字段时，其他条件不生效                     |  
| Url       | string | 查询特定资源的刷新/预热状态                                   |  
| StartTime | string | 过滤从此时开始的任务，为 ISO8601 格式时间，如：2006-01-02T15:04:05Z |  
| EndTime   | string | 过滤到此时的任务，时间类型同 StartTime                         |  
| Marker    | string | 分页信息，上一次未完成获取全部任务状态，此时带上 Marker 进一步获取            |

### 刷新缓存/查询刷新状态 Purge/GetPurgedStatus

```go
rawurl := fmt.Sprintf("http://%s/test/index.html", testDomain)
purgedId, err := testCli.Purge([]api.PurgeTask{
	{
		Url: rawurl,
	},
})
if err != nil {
	fmt.Printf("Purge for %s failed: %s\n", rawurl, err)
	return
}
fmt.Printf("Purge successfully: %+v\n", purgedId)

details, err := testCli.GetPurgedStatus(&api.CStatusQueryData{
	Id: string(purgedId),
})
if err != nil {
	fmt.Printf("GetPurgedStatus for %s failed: %s\n", purgedId, err)
	return
}
fmt.Printf("GetPurgedStatus successfully: %+v\n", details)
```

`api.PurgeTask` 的详细说明如下：

| 字段   | 类型     | 说明                                  |
|------|--------|-------------------------------------|
| Url  | string | 预刷新的资源                              |                                                                                         |
| Type | string | 合法值为 **file** 和 **directory**，分别表示文件刷新和目录刷新 |



### 预热资源/查询预热状态 Prefetch/GetPrefetchStatus

```go
rawurl := fmt.Sprintf("http://%s/test/index.html", testDomain)
prefetchId, err := testCli.Prefetch([]api.PrefetchTask{
	{
		Url: rawurl,
	},
})
if err != nil {
	fmt.Printf("Prefetch for %s failed: %s\n", rawurl, err)
	return
}
fmt.Printf("Prefetch successfully: %+v\n", prefetchId)

details, err := testCli.GetPrefetchStatus(&api.CStatusQueryData{
	Id: string(prefetchId),
})
if err != nil {
	fmt.Printf("GetPrefetchStatus for %s failed: %s\n", prefetchId, err)
	return
}
fmt.Printf("GetPrefetchStatus successfully: %+v\n", details)
```

### 查询刷新/预热限额 GetQuota

**> 查询账户的刷新/预热quota以及余额信息**

```go
quota, err := testCli.GetQuota()
	if err != nil {
		fmt.Printf("GetQuota failed: %s\n", err)
		return
	}
	fmt.Printf("GetQuota successfully: %+v\n", quota)
```

其中，GetQuota() 返回值为 QuotaDetail 类型，它的详细说明如下：

| 字段             | 类型      | 说明                  |
|----------------|---------|---------------------|
| UrlPurgeQuota  | integer | 天级别的file类型刷新额度      |
| UrlPurgeRemain |    integer     | 当天的file类型刷新余额       |
| DirPurgeQuota  |   integer      | 天级别的directory类型刷新额度 |
| DirPurgeRemain |  integer       | 当天的directory类型刷新余额  |
| PrefetchQuota  |   integer      | 天级别的预热额度            |
| PrefetchRemain |   integer      |      当天的预热余额               |

## 统计查询

### 查询流量 GetFlow

```go
var details []api.FlowDetail
var err error

// 查询账户整体1小时纬度的整体带宽
details, err = testCli.GetFlow(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetFlow failed: %s\n", err)
	return
}
fmt.Printf("GetFlow successfully: %+v\n", details)

// 查询特定域名在巴西的带宽
details, err = testCli.GetFlow(
	api.QueryStatByCountry("BR"), // BR 是巴西的 GEC 代码
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByDomains([]string{testDomain}),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetFlow failed: %s\n", err)
	return
}
fmt.Printf("GetFlow successfully: %+v\n", details)
```

### 查询PV GetPv

```go
var details []api.PvDetail
var err error

// 查询账户整体1小时纬度的整体PV
details, err = testCli.GetPv(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetPv failed: %s\n", err)
	return
}
fmt.Printf("GetPv successfully: %+v\n", details)

// 查询特定域名在巴西的PV
details, err = testCli.GetPv(
	api.QueryStatByCountry("BR"), // BR 是巴西的 GEC 代码
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByDomains([]string{testDomain}),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetPv failed: %s\n", err)
	return
}
fmt.Printf("GetPv successfully: %+v\n", details)
```

### 查询回源流量 GetSrcFlow

```go
var details []api.FlowDetail
var err error

// 查询账户整体1小时纬度的整体回源带宽
details, err = testCli.GetSrcFlow(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetSrcFlow failed: %s\n", err)
	return
}
fmt.Printf("GetSrcFlow successfully: %+v\n", details)

// 查询特定域名的回源带宽
details, err = testCli.GetSrcFlow(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByDomains([]string{testDomain}),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetSrcFlow failed: %s\n", err)
	return
}
fmt.Printf("GetSrcFlow successfully: %+v\n", details)
```

### 查询HTTP状态码详情 GetHttpCode

```go
var details []api.HttpCodeDetail
var err error

// 查询账户整体1小时纬度的整体状态码分布详情
details, err = testCli.GetHttpCode(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetHttpCode failed: %s\n", err)
	return
}
fmt.Printf("GetHttpCode successfully: %+v\n", details)

// 查询特定域名的状态码分布详情
details, err = testCli.GetHttpCode(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByDomains([]string{testDomain}),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetHttpCode failed: %s\n", err)
	return
}
fmt.Printf("GetHttpCode successfully: %+v\n", details)
```

### 查询流量命中率 GetRealHit

```go
var details []api.HitDetail
var err error

// 查询账户整体1小时纬度的整体流量命中率详情
details, err = testCli.GetRealHit(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetRealHit failed: %s\n", err)
	return
}
fmt.Printf("GetRealHit successfully: %+v\n", details)

// 查询特定域名的流量命中率详情
details, err = testCli.GetRealHit(
	api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
	api.QueryStatByDomains([]string{testDomain}),
	api.QueryStatByPeriod(api.Period3600))
if err != nil {
	fmt.Printf("GetRealHit failed: %s\n", err)
	return
}
fmt.Printf("GetRealHit successfully: %+v\n", details)
```

## 日志查询

**> 查询特定域名在某个时间区间的日志文件下载链接**

```go
endTs := time.Now().Unix()
startTs := endTs - 24*60*60
endTime := util.FormatISO8601Date(endTs)
startTime := util.FormatISO8601Date(startTs)

logEntries, err := testCli.GetDomainLog(testDomain, api.TimeInterval{
	StartTime: startTime,
	EndTime:   endTime,
})
if err != nil {
	fmt.Printf("GetDomainLog failed: %s\n", err)
	return
}
fmt.Printf("GetDomainLog successfully: %+v\n", logEntries)
```
