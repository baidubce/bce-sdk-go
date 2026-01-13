CDN服务

# 概述

本文档主要介绍CDN GO SDK的使用。在使用本文档前，您需要先了解CDN的一些基本知识，并已开通了CDN服务。若您还不了解CDN，可以参考[产品描述](https://cloud.baidu.com/doc/CDN/index.html)和[快速入门](https://cloud.baidu.com/doc/CDN/s/ujwvye8ao)。

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
ak := "your_access_key_id"
sk := "your_secret_key_id"
endpoint := "https://cdn.baidubce.com"
client, err := cdn.NewClient(ak, sk, endpoint)
```

在上面代码中，变量`ak`对应控制台中的“Access Key ID”，变量`sk`对应控制台中的“Access Key Secret”，获取方式请参考《 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。变量`endpoint`必须为`https://cdn.baidubce.com`，也是默认值，为空表示使用默认值，设置为其他则SDK无法工作。

在下面的示例中，会频繁使用到GetDefaultClient函数，它的定义为：

```go
func GetDefaultClient() *cdn.Client {
	ak := "your_access_key_id"
	sk := "your_secret_key_id"
	endpoint := "https://cdn.baidubce.com"

	// ignore error in test, but you should handle error in dev
	client, _ := cdn.NewClient(ak, sk, endpoint)
	return client
}
```

## 自定义HTTP请求

> 支持对任何开放接口自定义HTTP请求

示例展示了不通过SDK的方法，而是自行根据[文档](https://cloud.baidu.com/doc/CDN/s/qjwvyexh6)构造请求，查询某个域名是否可以添加到百度云CDN系统。

```go
cli := GetDefaultClient()
method := "GET"
urlPath := fmt.Sprintf("/v2/domain/%s/valid", testDomain)
var params map[string]string
// 此请求头非必须，仅测试发送请求头
reqHeaders := map[string]string{
	"X-Test": "go-sdk-test",
}
var bodyObj interface{}
var respObj interface{}
err := cli.SendCustomRequest(method, urlPath, params, reqHeaders, bodyObj, &respObj)
fmt.Printf("err:%+v\n", err)
fmt.Printf("respObj details:\n\ttype:%T\n\tvalue:%+v", respObj, respObj)
```

## 域名操作接口

### 域名列表查询 ListDomains

> 查询该用户账户下拥有的加速域名

```go
marker := ""
cli := GetDefaultClient()
domains, nextMarker, err := cli.ListDomains(marker)
fmt.Printf("domains:%+v\n", domains)
fmt.Printf("nextMarker:%+v\n", nextMarker)
fmt.Printf("err:%+v\n", err)
```

`marker`参数表示从这个字符串所代表的索引开始查询域名列表，空表示从头开始。目前服务器处理分页size是一个很大的数，所以使用的时候将marker赋值为空即可。ListDomains返回的`nextMarker`表示下一个域名列表索引 ，空表示从`marker`开始之后的域名列表全部被获取，非空可用于传递到ListDomains参数。`domains`是一个string slice，表示域名列表数据，如: `["1.baidu.com", "2.baidu.com"]`

### 查询用户名下所有域名 GetDomainStatus

> 查询用户名下所有域名和域名状态，可以根据特定状态查询属于这个状态的域名。

```go
cli := client.GetDefaultClient()
status := "ALL"
domainStatus, err := cli.GetDomainStatus(status, "")
fmt.Printf("domainStatus:%+v\n", domainStatus)
fmt.Printf("err:%+v\n", err)
```

`status`表示域名的状态，如果为`ALL`，表示查询所有状态的域名，如果为`RUNNING`查询已经加速的域名，`STOPPED`查询停止加速的域名，`OPERATING`查询操作中的域名。
`domainStatus`是**DomainStatus**类型对象，如下所示：

| 字段   | 类型   | 说明                                                |
| ------ | ------ | --------------------------------------------------- |
| Domain | string | 域名，如`www.baidu.com`                             |
| Status | string | 域名状态，合法值为`RUNNING`、`STOPPED`和`OPERATING` |

`domainStatus`的示例数据如：`[{"Domain":"1.baidu.com", "Status":"RUNNING"}, {"Domain":"2.baidu.com", "Status":"STOPPED"}]`

### 查询域名是否可添加 IsValidDomain

> 查询特定域名是否可以使用CDN加速，一个可以被添加的域名必须是合法的域名，不可重复添加，必须通过ICP备案。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"
domainValidInfo, err := cli.IsValidDomain(testDomain)
fmt.Printf("domainValidInfo:%+v\n", domainValidInfo)
fmt.Printf("err:%+v\n", err)
```
`testDomain`是要测试的域名，`domainValidInfo`是**DomainValidInfo**类型对象，表示的是关于是否可添加的详细信息

| 字段   | 类型   | 说明                                                |
| ------ | ------ | --------------------------------------------------- |
| Domain | string | 域名，如`www.baidu.com`                             |
| IsValid | bool | true表示该域名可添加，false表示该域名不可被添加 |
| Message | string | 当IsValid为false，表示域名不可被添加的原因；当IsValid为true时，为空 |

### 创建加速域名接口 CreateDomain

> 添加一个加速域名到用户名下，该域名在调用IsValidDomain的时候返回必须为true。创建加速域名必须设置源站。

```go
cli := client.GetDefaultClient()
domainCreatedInfo, err := cli.CreateDomain("test_go_sdk.baidu.com", &api.OriginInit{
	Origin: []api.OriginPeer{
		{
			Peer:      "1.1.1.1",
			Host:      "www.baidu.com",
			Backup:    true,
			Follow302: true,
		},
		{
			Peer:      "http://2.2.2.2",
			Host:      "www.baidu.com",
			Backup:    false,
			Follow302: true,
		},
	},
	DefaultHost: "www.baidu.com",
	Form:        "default",
})
fmt.Printf("domainCreatedInfo:%+v\n", domainCreatedInfo)
fmt.Printf("err:%+v\n", err)
```

`api.OriginPeer`表示一个源站，`api.OriginInit`表示对于这个加速域名的源站配置，包括一个或多个源站，和一个默认的回源host，表示为`DefaultHost`。`Form`表示加速服务类型，默认为`default`，其他可选的服务类型有`image`表示图片小文件，`download`表示大文件下载，`media`表示流媒体点播，`dynamic`表示动静态加速。
源站信息**api.OriginPeer**的结构如下：

| 字段      | 类型   | 说明                                                         |
| --------- | ------ | ------------------------------------------------------------ |
| Peer      | string | 源站的endpoint，如`http://2.2.2.2:9090`或     `https://test.baidu.com:9090`，源站类型可以为域名形式、IP形式和bucket，示例代码展示的是IP形式的源站类型，需要注意的是一个域名所有的源站必须具有相同的源站类型。**如果使用的是默认端口，即HTTP的80和HTTPS的443，CDN系统会将源站视为HTTP和HTTPS两种协议都支持**。 |
| Host      | string | 表示回源时在HTTP请求头中携带的Host，如果没有设置则使用`DefaultHost` |
| Backup    | bool   | 表示是否为备用源站，备用源站在主源站都不可用的情况下才会使用，false表示这个源站是主源站。需要注意的是在golang中bool对象默认值为false，如果没有显式设置Backup的值，那么默认就是false。 |
| Follow302 | bool   | true表示当回源到源站时源站返回302时，CDN系统会处理Location重定向，主动获取资源返回给客户端。false表示当源站返回302时，CDN系统透传302给客户端。**需要注意的是，目前Follow302已经修改为所有源站级别的配置，所以要求所有的源站Follow302必须一致，否则可能会出现不可预料的结果**。 |


另外，您也可以使用更高级的 **CreateDomainWithOptions** 方法创建域名，目前支持的 Option 有以下几项：
- CreateDomainWithForm：配置域名的form（类型、类别）
- CreateDomainWithOriginDefaultHost：配置域名的域名级别回源Host，优先级低于源站级别回源Host
- CreateDomainWithTags：绑定域名到标签
- CreateDomainAsDrcdnType：将域名创建为 DRCDN 类型，必须要先开通 DRCDN 产品，关于如何开通 DRCDN 请参考教程：https://cloud.baidu.com/doc/DRCDN/s/gkh08rmki。当配置此项时， dsa 配置不能为 nil。
- CreateDomainWithCacheTTL：创建域名的同时设置缓存过期规则。

下面通过代码展示如何调用 CreateDomainWithOptions。

```go
cli := client.GetDefaultClient()
domainCreatedInfo, err := testCli.CreateDomainWithOptions("test_go_sdk.baidu.com", []api.OriginPeer{
		{
			Peer: "1.2.3.4",
			Host: "www.baidu.com",
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
	}), CreateDomainWithForm("image"), CreateDomainWithOriginDefaultHost("origin.baidu.com"))

fmt.Printf("domainCreatedInfo:%+v\n", domainCreatedInfo)
fmt.Printf("err:%+v\n", err)
```


### 启用/停止/删除加速域名 EnableDomain/DisableDomain/DeleteDomain

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 启用加速域名
err := cli.EnableDomain(testDomain)
fmt.Printf("err:%+v\n", err)

// 停用加速域名
err = cli.DisableDomain(testDomain)
fmt.Printf("err:%+v\n", err)

// 删除该用户名下的加速域名
err = cli.DeleteDomain(testDomain)
fmt.Printf("err:%+v\n", err)
```

## 域名配置接口

### 查询加速域名详情 GetDomainConfig

> 查询某个特定域名的全部配置信息

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"
domainConfig, err := cli.GetDomainConfig(testDomain)
fmt.Printf("domainConfig:%+v\n", domainConfig)
fmt.Printf("err:%+v\n", err)
```

`domainConfig`是`*api.DomainConfig`类型的对象，他的结构相对比较复杂，如下所示：

| 字段           | 类型             | 说明                                                         |
| -------------- | ---------------- | ------------------------------------------------------------ |
| Domain         | string           | 域名信息，如：`test.baidu.com`。                             |
| Cname          | string           | 域名的CNAME，如：`test.baidu.com.a.bdydns.com`。             |
| Status         | string           | 域名状态，合法值为`RUNNING`、`STOPPED`和`OPERATING`。        |
| CreateTime     | string           | 域名创建的时间，UTC时间，如：`2019-09-02T10:08:38Z。`        |
| LastModifyTime | string           | 上次修改域名配置的时间，UTC时间，如：`2019-09-06T15:00:21Z`。 |
| IsBan          | string           | 是否被禁用，禁用的意思是欠费或者其他违规操作被系统封禁，被封禁的域名不拥有加速服务。`ON`表示为被封禁，`YES`表示被封禁了。 |
| Origin         | []api.OriginPeer | 源站信息，在创建加速域名接口有详细说明。                     |
| DefaultHost    | string           | 默认的回源host，在创建加速域名接口有详细说明。               |
| CacheTTL       | []api.CacheTTL   | 文件类型与路径的缓存策略，在`设置/查询缓存过期规则`小段有详细说明。     |
| LimitRate      | int              | 下载限速，单位Byte/s。                                       |
| RequestAuth    | api.RequestAuth  | 访问鉴权配置，在设置访问鉴权有详细说明。                     |
| Https          | api.HTTPSConfig  | HTTPS加速配置，在设置HTTPS加速有详细说明。                   |
| FollowProtocol | bool             | 是否开启了协议跟随回源，true表示开启了协议跟随回源，即访问资源是https://xxx或http://xxx之类的url，回源也使用相对应的scheme，即分别为HTTPS和HTTP。 |
| SeoSwitch      | api.SeoSwitch    | seo 开关配置，在设置SEO开关属性有详细说明。                  |

### 设置/查询缓存过期规则 SetCacheTTL/GetCacheTTL

> 设置针对文件、目录和错误码等相关的缓存策略，合理设置缓存策略可以提高命中率。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置缓存策略使用源站规则
err := cli.SetCacheTTL(testDomain, []api.CacheTTL{
	{
		Type:  "origin",
		Value: "-",
		TTL:   0,
	},
})
fmt.Printf("err:%+v\n", err)

// 设置缓存策略，分别设置后缀规则、目录规则、路径完全匹配和错误码
err = cli.SetCacheTTL(testDomain, []api.CacheTTL{
	{
		Type:   "suffix",
		Value:  ".jpg",
		TTL:    420000,
		Weight: 30,
	},
	{
		Type:  "path",
		Value: "/js",
		TTL:   10000,
	},
	{
		Type:   "exactPath",
		Value:  "index.html",
		TTL:    600,
		Weight: 100,
	},
	{
		Type:   "code",
		Value:  "404",
		TTL:    600,
		Weight: 100,
	},
})
fmt.Printf("err:%+v\n", err)

// 查询缓存策略
cacheTTls, err := cli.GetCacheTTL(testDomain)
fmt.Printf("cacheTTls:%+v\n", cacheTTls)
fmt.Printf("err:%+v\n", err)
```

`cacheTTls`包含全部的缓存策略，每一个缓存策略用`api.CacheTTL`类型的对象表示，如下关于缓存策略结构的说明。

| 字段   | 类型   | 说明                                                         |
| ------ | ------ | ------------------------------------------------------------ |
| Type   | string | 缓存策略的类型。合法的类型有：`suffix`、`path`、`origin`、`code`和`exactPath`。"suffix"表示文件名后缀，"path"表示url中的目录，"origin"表示源站规则，此规则只有一条，只表示出权重即可，value为"-", ttl为 0，"code"表示异常码缓存，如可以配置404缓存100s ，“exactPath”表示路径完全匹配。 |
| Value  | string | Type所指定类型的配置规则。                                   |
| Weight | int    | 权重，0-100的整数，权重越高优先级越高，默认为0，优先级在为code类型下是没有作用的，可以忽略。权重越大，优先级越高，规则优先生效。不推荐两条缓存策略配置相同的权重，如果权重相同，会随机选择其中一条策略生效。 |
| TTL    | int    | 缓存时间，单位为秒。                                         |


### 设置/查询域名条件源站配置 SetPageRulesOriginConfig/GetPageRulesOriginConfig

> 按请求特征（`matchRules`）匹配后，命中规则的请求使用指定的源站配置（`originConfig`）。

**限制与注意事项：**
- 最多允许设置 **10** 个 PageRule；每个 PageRule 最多 **5** 个 MatchRule。
- `matchRules` 为数组，**n 条规则全部匹配成功（AND）** 才视为该 PageRule 匹配成功。
- 通配符仅支持 `*`。
- `MatchRule.target` 支持：`url` / `suffix` / `directory` / `header`（源站配置支持 4 类匹配目标）。
- 当 `target=header` 时，`headerValues` 必须有效；否则 `pathValues` 必须有效。
  “有效”指字段存在且元素个数 > 0。
- `headerValues` 最多只能包含 **1** 个 header 键值对；key/value 的长度最大均为 **200** 字符。
- `pathValues` 取值格式要求：
  - `url`：以 `/` 开头，不含 `http(s)://` 头和域名，支持 `*` 通配符
  - `suffix`：以 `.` 开头，例如 `.jpg`
  - `directory`：以 `/` 开头（前缀匹配），且不支持用 `/` 匹配全部文件
- `MatchRule.modifier`（匹配规则）：
  - 默认值 `""`：目标等于或包含匹配值
  - `!`：不等于或不包含匹配值
  - `~`：忽略匹配值大小写
  - `!~`：不等于或不包含匹配值，且忽略匹配值大小写
- 调用成功后响应体 `status` 为 `"RUNNING"`（异步生效），建议在查询前等待几秒或重试。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置条件源站配置（示例参考官方文档“请求示例1: 设置条件源站配置”）
pageRules := []api.PageRuleOriginConfig{
    {
        MatchRules: []api.MatchRule{
            {
                Target: "suffix",
                Modifier: "",
                PathValues: []string{".mp4", ".flv"},
            },
            {
                Target: "header",
                Modifier: "!",
                HeaderValues: map[string][]string{
                    "user-agent": {"android"},
                },
            },
        },
        Config: api.OriginConfigWrapper{
            OriginConfig: []api.OriginItem{
                {
                    Type: "DOMAIN",
                    Addr: "video.test.com",
                    HttpPort: 8088,
                },
            },
        },
    },
    {
        MatchRules: []api.MatchRule{
            {
                Target: "directory",
                Modifier: "~",
                PathValues: []string{"/autio", "/music"},
            },
            {
                Target: "header",
                Modifier: "!~",
                HeaderValues: map[string][]string{
                    "user-agent": {"android"},
                },
            },
        },
        Config: api.OriginConfigWrapper{
            OriginConfig: []api.OriginItem{
                {
                    Type: "DOMAIN",
                    Addr: "audio.test.com",
                    HttpPort: 8000,
                },
            },
        },
    },
}

err := cli.SetPageRulesOriginConfig(testDomain, pageRules)
fmt.Printf("err:%+v\n", err)

// 查询条件源站配置（建议在 Set 后等待/重试）
rules, err := cli.GetPageRulesOriginConfig(testDomain)
fmt.Printf("rules:%+v, err:%+v\n", rules, err)
```

`pageRules` 是 `[]api.PageRuleOriginConfig` 类型的对象，下面是详细说明：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| MatchRules | []api.MatchRule | 匹配规则集合（数组）。当 n 条规则全部匹配成功时视为该 PageRule 匹配成功。最多 5 条 MatchRule。通配符仅支持 `*`。 |
| Config | api.OriginConfigWrapper | 命中后生效的源站配置（字段为 `originConfig`）。具体源站条目结构参考“设置回源地址”文档。 |

`matchRules` 中 `api.MatchRule` 字段说明：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| Target | string | 匹配目标：`url`（路径）、`suffix`（扩展名）、`directory`（目录前缀）、`header`（请求头）。 |
| Modifier | string | 匹配修饰符：默认 `""`；可选 `!`、`~`、`!~`。 |
| HeaderValues | map[string][]string | 仅当 `Target=header` 时使用且必传（有效）。最多 1 个 header 键值对，key/value 最大 200 字符。 |
| PathValues | []string | 当 `Target=url/suffix/directory` 时使用且必传（有效）。`url` 以 `/` 开头且不含协议域名；`suffix` 以 `.` 开头；`directory` 以 `/` 开头且不支持 `/` 匹配全部文件。 |

---

### 设置/查询域名条件缓存参数过滤规则 SetPageRulesCacheFullUrl/GetPageRulesCacheFullUrl

> 按请求特征（`matchRules`）匹配后，命中规则的请求使用指定的缓存参数过滤规则（`cacheFullUrl / cacheUrlArgs / ignoreUrlArgs`）。

**限制与注意事项：**
- 最多允许设置 **10** 个 PageRule；每个 PageRule 最多 **5** 个 MatchRule。
- `matchRules` 为数组，**n 条规则全部匹配成功（AND）** 才视为该 PageRule 匹配成功。
- 通配符仅支持 `*`。
- 缓存参数过滤规则下的 `MatchRule.target` 仅支持：`url` / `suffix` / `directory`（不支持 `header`）。
- 缓存参数过滤规则下的 `MatchRule.modifier` 仅支持：`""` 和 `"~"`（不支持 `!`、`!~`）。
- 缓存参数过滤规则下不支持 `headerValues`；`target=url/suffix/directory` 时必须传 `pathValues` 且有效。
- `pathValues` 取值格式要求同上：`url` 以 `/` 开头且不含协议域名；`suffix` 以 `.` 开头；`directory` 以 `/` 开头且不支持 `/` 匹配全部文件。
- `Config` 中 `cacheUrlArgs` 与 `ignoreUrlArgs` 为两种不同的过滤方式，建议二选一使用：
  - `cacheUrlArgs`：保留（参与缓存 key 生成）的参数列表
  - `ignoreUrlArgs`：忽略（不参与缓存 key 生成）的参数列表
- 调用成功后响应体 `status` 为 `"RUNNING"`（异步生效），建议在查询前等待几秒或重试。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置条件缓存参数过滤规则（示例参考官方文档“请求示例2: 设置条件缓存参数规则配置”）
pageRules := []api.PageRuleCacheFullUrl{
    {
        MatchRules: []api.MatchRule{
            {
                Target: "url",
                Modifier: "",
                PathValues: []string{"/test"},
            },
        },
        Config: api.CacheUrlArgs{
            CacheFullUrl: false,
            CacheUrlArgs: []string{"a"},
        },
    },
    {
        MatchRules: []api.MatchRule{
            {
                Target: "url",
                Modifier: "",
                PathValues: []string{"/test"},
            },
        },
        Config: api.CacheUrlArgs{
            CacheFullUrl: false,
            IgnoreUrlArgs: []string{"b"},
        },
    },
}

err := cli.SetPageRulesCacheFullUrl(testDomain, pageRules)
fmt.Printf("err:%+v\n", err)

// 查询条件缓存参数过滤规则（建议在 Set 后等待/重试）
rules, err := cli.GetPageRulesCacheFullUrl(testDomain)
fmt.Printf("rules:%+v, err:%+v\n", rules, err)
```

`pageRules` 是 `[]api.PageRuleCacheFullUrl` 类型的对象，下面是详细说明：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| MatchRules | []api.MatchRule | 匹配规则集合（数组）。当 n 条规则全部匹配成功时视为该 PageRule 匹配成功。最多 5 条 MatchRule。通配符仅支持 `*`。 |
| Config | api.CacheUrlArgs | 命中后生效的缓存参数过滤规则。 |

`Config`（`api.CacheUrlArgs`）字段说明：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| CacheFullUrl | bool | `true` 表示全 URL 缓存；`false` 表示按参数过滤规则生成缓存 key。 |
| CacheUrlArgs | []string | `CacheFullUrl=false` 时有效：表示“保留的参数列表”。为空表示忽略所有参数。 |
| IgnoreUrlArgs | []string | `CacheFullUrl=false` 时有效：表示“忽略的参数列表”（建议与 `CacheUrlArgs` 二选一）。 |

`matchRules` 在本配置项下的额外限制：
- `Target` 仅支持 `url/suffix/directory`
- `Modifier` 仅支持 `""` 和 `"~"`
- 不支持 `HeaderValues`


### 设置/查询源站配置 SetOriginConfig/GetOriginConfig

> 设置或查询域名的源站配置，该配置支持IP类型、域名类型、BUCKET类型和第三方对象存储类型。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置IP类型源站配置
err := testCli.SetOriginConfig(testAuthorityDomain, []api.OriginItem{
  {
      Addr:             "test1.baidu.com",
      Type:             "DOMAIN",
      Host:             "test1.baidu.com",
      UpstreamProtocol: "http",
      Weight:           10,
      Backup:           false,
  },
})
fmt.Printf("err:%+v\n", err)

// 设置域名类型源站配置
err := testCli.SetOriginConfig(testAuthorityDomain, []api.OriginItem{
  {
      Addr:             "test2.baidu.com",
      Type:             "DOMAIN",
      Host:             "test2.baidu.com",
      UpstreamProtocol: "http",
      Weight:           10,
      Backup:           false,
  },
})
fmt.Printf("err:%+v\n", err)

// 设置BUCKET类型源站配置
err := testCli.SetOriginConfig(testAuthorityDomain, []api.OriginItem{
  {
      Type:             "BUCKET",
      Host:             "test.bj.bcebos.com",
      UpstreamProtocol: "https",
      Addr:             "test.bj.bcebos.com",
      Backup:           false,
  },
})
fmt.Printf("err:%+v\n", err)

// 设置第三方对象存储类型源站配置
err := testCli.SetOriginConfig(testAuthorityDomain, []api.OriginItem{
  {
      Type:             "DOMAIN",
      Addr:             "test1.baidu.com",
      Host:             "test1.baidu.com",
      UpstreamProtocol: "https",
      Backup:           false,
      ThirdBucketAuth: &api.ThirdBucketAuth{
          AuthType: "aws_v4",
          Enabled:  true,
          Ak:       "xxx",
          Sk:       "xxx",
          Bucket:   "mybucket",
          Region:   "us-east-1",
          Service:  "s3",
      },
  },
})
fmt.Printf("err:%+v\n", err)

// 查询源站配置
originConfig, err := testCli.GetOriginConfig(testAuthorityDomain)
fmt.Printf("originConfig:%+v\n", originConfig)
```

`originConfig`是`[]api.OriginItem`类型的对象，下面是详细说明：

| <span style="display:inline-block;width: 100px">参数</span>    | <span style="display:inline-block;width: 30px">类型</span>  | <span style="display:inline-block;width: 30px">是否必选</span>  | 说明
| ---- | ------ | ---- | ------- |
| addr     | String      | 是      | 源站地址。支持 IPv4、IPv6 形式的 IP 地址，或者域名，不能重复。  |
| type     | String      | 是      | 源站类型。可选值为 IP、DOMAIN、BUCKET，值为 DOMAIN 时，会忽略 isp 配置；值为 BUCKET 时，addr 要填写 BUCKET 的完整地址，并且会忽略 weight、isp 的配置。  |
| httpPort     | Int      | 否      | http 回源端口。默认80。  |
| httpsPort     | Int      | 否      | https 回源端口。默认443。  |
| host     | String      | 否      | 回源时使用的 host 值。  |
| upstreamProtocol     | String      | 否      | 回源协议。可选值为 http、https、\*， 其中 \* 表示协议跟随。  |
| weight     | Int      | 否      | 源站权重，值为1-100之间的整数。举例：按照权重分配回源的流量，假设某加速域名有两个源站，一个源站 A 权重是80，另一个 B 是20，总的回源量是1G，那么其中A源站大约会有800M的回源，B大约会有200M的回源。  |
| backup     | Bool      | 否      | 是否为备源站。true 表示备源站，false 表示主源站，默认为 false。  |
| isp     | String      | 否      | 源站所属的运营商。默认无，可选值为un（联通）、ct（电信）、cm（移动）。  |
| thirdBucketAuth     | ThirdBucketAuth      | 否      |第三方对象存储的源站配置。   |
| probeUrl     | String      | 否      | 探测地址。配置了 probeUrl，表明配置对源站进行应用层探测，百度智能云 CDN 会定期给源站发送 GET /{probeUrl} 请求进行探测，如果源站响应的 HTTP 状态码小于500，那么认为源站正常；否则，认为源站异常，在探测恢复正常前不会选择异常源站进行回源。如果你希望探测的资源为 scheme://$addr:$http(s)Port/1.gif，那么此处的 probeUrl 应设置为"1.gif"，**而不是"/1.gif"**。需要注意，probeUrl 设置为空字符串表示不开启源站探测。  |

ThirdBucketAuth 类型说明：

| <span style="display:inline-block;width: 100px">参数</span>   | <span style="display:inline-block;width: 30px">类型</span>   | <span style="display:inline-block;width: 30px">是否必选</span> | 说明
| ---- | ------ |------------------------------------------------------------| -------
| authType     | String      | 是                                                          | 表示第三方的对象存储来源类型，当前支持 AWS S3、腾讯云 COS、阿里云 OSS、华为云 OBS。其合法值为："aws_v2"、"aws_v4"、"cos"、"oss"、"obs"。
| enabled     | Bool      | 否                                                          | 启用第三方私有bucket鉴权的开关，默认值为 flase。
| ak     | String      | 是                                                          | 第三方对象存储的访问密钥 Access Key。
| sk     | String      | 是                                                          | 第三方对象存储的私有访问密钥 Secret Access Key。
| bucket     | String      | 否                                                          | 第三方对象存储的 bucket。<br><br>注: <br>• 当对象存储来源类型为 "aws_v2" 时，**必须设置** bucket。<br> • 当对象存储来源类型为 "aws_v4" 时，**无需设置** bucket。<br> • 当对象存储来源类型为 "cos" 时，**无需设置** bucket。<br> • 当对象存储来源类型为 "oss" 时，**必须设置** bucket。<br> • 当对象存储来源类型为 "obs" 时，**无需设置** bucket。
| region     | String      | 否                                                          | 第三方对象存储的区域。<br><br>注: <br>• 当对象存储来源类型为 "aws_v2" 时，**无需设置**区域。<br> • 当对象存储来源类型为 "aws_v4" 时，选填，默认区域为 "us-east-1"。<br> • 当对象存储来源类型为 "cos" 时，**无需设置**区域。<br> • 当对象存储来源类型为 "oss" 时，选填，默认区域为 "cn-hangzhou"。<br> • 当对象存储来源类型为 "obs" 时，**无需设置**区域。
| service     | String      | 否                                                          | 第三方对象存储的服务。<br><br>注: <br>• 当对象存储来源类型为 "aws_v2" 时，**无需设置**服务。<br> • 当对象存储来源类型为 "aws_v4" 时，选填，默认服务为 "s3"。<br> • 当对象存储来源类型为 "cos" 时，**无需设置**服务。<br> • 当对象存储来源类型为 "oss" 时，**无需设置**服务。<br> • 当对象存储来源类型为 "obs" 时，**无需设置**服务。

注：当 “无需设置” 用于描述某字段时，指的是该字段无效，即使在第三方对象存储配置中设置了值，也不会生效。

---

### 设置/查询缓存参数过滤规则 SetCacheUrlArgs/GetCacheUrlArgs

> 设置缓存key，缓存key为CDN系统对某个资源的进行缓存的时候所采用的key，一个url可能带有很多参数，那么时候所有的参数都需要放在缓存key中呢？其实是不必的，下面展示了不同的设置。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置全URL缓存
err := cli.SetCacheUrlArgs(testDomain, &api.CacheUrlArgs{
	CacheFullUrl: true,
})
fmt.Printf("err:%+v\n", err)

// 设置缓存带有部分参数
err = cli.SetCacheUrlArgs(testDomain, &api.CacheUrlArgs{
	CacheFullUrl: false,
	CacheUrlArgs: []string{"name", "id"},
})
fmt.Printf("err:%+v\n", err)

// 设置忽略所有参数
err = cli.SetCacheUrlArgs(testDomain, &api.CacheUrlArgs{
	CacheFullUrl: false,
	CacheUrlArgs: []string{"name", "id"},
})
fmt.Printf("err:%+v\n", err)

// 查询关于URL参数缓存设置
cacheUrlArgs, err := cli.GetCacheUrlArgs(testDomain)
fmt.Printf("cacheUrlArgs:%+v\n", cacheUrlArgs)
```

`cacheUrlArgs`是`api.CacheUrlArgs`类型的对象，下面是详细说明：

| 字段         | 类型     | 说明                                                         |
| ------------ | -------- | ------------------------------------------------------------ |
| CacheFullUrl | bool     | true或false，true表示支持全URL缓存，false表示忽略参数缓存(可保留部分参数)。注意golang中如果不显式赋值CacheFullUrl为true或false，那么取零值false。 |
| CacheUrlArgs | []string | CacheFullUrl为true时，此项不起作用；CacheFullUrl为false时，此项表示保留的参数列表，如果为空，表示忽略所有参数。 |

### 设置/查询自定义错误码页面 SetErrorPage/GetErrorPage

> 用户可以定义当访问源站或CDN系统内部出现错误时（通常是返回4xx或5xx错误码），有CDN系统返回给用户的重定向页面，用户可以设置这个重定向页面的链接。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置错误出现时的重定向页面
err := cli.SetErrorPage(testDomain, []api.ErrorPage{
	{
		Code:         510,
		RedirectCode: 301,
		Url:          "/customer_404.html",
	},
	{
		Code: 403,
    // 这里没有设置RedirectCode，表示使用默认的302重定向
		Url:  "/custom_403.html",
	},
})
fmt.Printf("err:%+v\n", err)

// 取消设置错误出现时的重定向页面
err = cli.SetErrorPage(testDomain, []api.ErrorPage{})
fmt.Printf("err:%+v\n", err)

// 设置错误出现时的重定向页面
err = cli.SetErrorPage(testDomain, []api.ErrorPage{})
fmt.Printf("err:%+v\n", err)

// 查询设置的错误重定向页面
errorPages, err := cli.GetErrorPage(testDomain)
fmt.Printf("errorPages:%+v\n", errorPages)
```

`errorPages`是`[]api.ErrorPage`类型的对象，`api.ErrorPage`结构的详细说明如下：

| 字段         | 类型   | 说明                                                         |
| ------------ | ------ | ------------------------------------------------------------ |
| Code         | int    | 特定的状态码，要求必须为HTTP的标准错误码，且不能是408、444、499等客户端异常/提前断开这类特殊状态码。 |
| RedirectCode | int    | 重定向状态码，当出现code错误码时，重定向的类型。支持301和302，默认302。 |
| Url          | string | 重定向目标地址 ，当出现code错误码是，重定向到这个用户自定义的url，即301或302重定向中HTTP报文中的Location的值为了Url。 |

### 设置/查询和其他域名共享缓存配置 SetCacheShared/GetCacheShared

> 用户可以设置域名与另外一个域名共享缓存，提升命中率，前提是两个域名必须属于同一账户。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置test_go_sdk.baidu.com和www.baidu.com共享缓存
err := cli.SetCacheShared(testDomain, &api.CacheShared{
		Enabled:    true,
		SharedWith: "www.baidu.com",
})
fmt.Printf("err:%+v\n", err)

// 取消test_go_sdk.baidu.com和任何域名共享缓存
err := cli.SetCacheShared(testDomain, &api.CacheShared{
		Enabled:    false,
})
fmt.Printf("err:%+v\n", err)

// 查询test_go_sdk.baidu.com的共享缓存配置
cacheSharedConfig, err := cli.GetCacheShared(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("cacheSharedConfig:%+v\n", cacheSharedConfig)
```

### 设置/查询访问Referer控制 SetRefererACL/GetRefererACL

> 设置Referer的访问控制规则，当通过浏览器或者其他形式跳转到资源是，浏览器通常会在请求头中加入Referer信息。CDN系统提供对Referer进行过滤，可以设置Referer黑名单或者白名单。当Referer出现在黑名单中的请求到达时响应403，当Referer没有出现在白名单中的请求到达时响应403。黑名单和白名单设置只能选其一，要么设置黑名单要么设置白名单。当设置了白名单并且过滤规则为空时，白名单不生效。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"
isAllowEmpty := true

// 设置白名单ACL
err := cli.SetRefererACL(testDomain, nil, []string{
	"a.bbbbbb.c",
	"*.baidu.com.*",
}, isAllowEmpty)
fmt.Printf("err:%+v\n", err)

// 设置白名单为空，白名单规则不生效
err = cli.SetRefererACL(testDomain, nil, []string{}, isAllowEmpty)
fmt.Printf("err:%+v\n", err)

// 设置黑名单ACL
err = cli.SetRefererACL(testDomain, []string{
	"a.b.c",
	"*.xxxxx.com.*",
}, nil, isAllowEmpty)
fmt.Printf("err:%+v\n", err)

// 查询referer ACL设置
refererACL, err := cli.GetRefererACL(testDomain)
fmt.Printf("refererACL:%+v\n", refererACL)
fmt.Printf("err:%+v\n", err)
```

`isAllowEmpty`表示是否允许空Referer访问，true表示允许空Referer访问，即如果Referer为空，那么不管是设置了黑名单还是白名单都不会生效，大多数情况下这个值都设置为**true**；false表示不允许空Referer访问，当设置为false时，如果访问的HTTP报文中不存在Referer那么CDN系统将返回403。`refererACL`是`api.RefererACL`类型的对象，他的详细说明如下：

| 字段       | 类型     | 说明                                                         |
| ---------- | -------- | ------------------------------------------------------------ |
| BlackList  | []string | 可选项，表示referer黑名单列表，支持使用通配符*，不需要加protocol，如设置某个黑名单域名，设置为"www.xxx.com"形式即可，而不是"http://www.xxx.com"。* |
| WhiteList  | []string | *可选项，list类型，表示referer白名单列表，支持通配符*，同样不需要加protocol。 |
| AllowEmpty | bool     | 必选项，bool类型，表示是否允许空referer访问，为true即允许空referer访问。 |

### 设置/查询访问IP控制 SetIpACL/GetIpACL

> CDN获取客户端IP，同配置中的IP黑/白名单进行匹配，对匹配上的客户端请求进行拒绝/放过。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置IP白名单
err := cli.SetIpACL(testDomain, []string{
	"1.1.1.1",
	"2.2.2.2",
}, nil)
fmt.Printf("err:%+v\n", err)

// 设置IP黑名单，CIDR格式的IP
err = cli.SetIpACL(testDomain, nil, []string{
	"1.2.3.4/24",
})
fmt.Printf("err:%+v\n", err)

// 查询IP黑白设置
ipACL, err := cli.GetIpACL(testDomain)
fmt.Printf("ipACL:%+v\n", ipACL)
fmt.Printf("err:%+v\n", err)
```

`ipACL`是`api.IpACL`类型对象，详细说明如下：

| 字段      | 类型     | 说明                                                         |
| --------- | -------- | ------------------------------------------------------------ |
| BlackList | []string | IP黑名单列表，当设置黑名单生效时，当客户端的IP属于BlackList，CDN系统返回403。BlackList不可与WhiteList同时设置。 |
| WhiteList | []string | IP白名单列表，当设置白名单生效时，当WhiteList为空时，没有白名单效果。当WhiteList非空时，只有客户端的IP属于WhiteList才允许访问。同样不可与BlackList同时设置。 |

### 设置/查询访问UA控制 SetUaACL/GetUaACL

> CDN获取HTTP请求头中的User-Agent，同配置中的黑/白名单进行匹配，对匹配上的请求进行拒绝/放过。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置UA白名单
err = cli.SetUaACL(testDomain, nil, []string{
		"curl/7.73.0",
})
fmt.Printf("err:%+v\n", err)

// 设置UA黑名单
err := cli.SetUaACL(testDomain, []string{
		"Test-Bad-UA",
}, nil)
fmt.Printf("err:%+v\n", err)

// 查询UA黑白设置
uaACL, err := cli.GetUaACL(testDomain)
fmt.Printf("uaACL:%+v\n", uaACL)
fmt.Printf("err:%+v\n", err)
```

`ipACL`是`api.IpACL`类型对象，详细说明如下：

| 字段      | 类型     | 说明                                                         |
| --------- | -------- | ------------------------------------------------------------ |
| BlackList | []string | IP黑名单列表，当设置黑名单生效时，当客户端的IP属于BlackList，CDN系统返回403。BlackList不可与WhiteList同时设置。 |
| WhiteList | []string | IP白名单列表，当设置白名单生效时，当WhiteList为空时，没有白名单效果。当WhiteList非空时，只有客户端的IP属于WhiteList才允许访问。同样不可与BlackList同时设置。 |

### 设置访问鉴权 SetDomainRequestAuth

> 高级鉴权也是为了防止客户源站内容被盗用，比Referer黑白名单和IP黑白名单更加安全。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"
err := cli.SetDomainRequestAuth(testDomain, &api.RequestAuth{
	Type:    "c",
	Key1:    "secretekey1",
	Key2:    "secretekey2",
	Timeout: 300,
	WhiteList: []string{
		"/crossdomain.xml",
	},
	SignArg: "sign",
	TimeArg: "t",
})

fmt.Printf("err:%+v\n", err)
```

示例代码设置一个C类鉴权方式，对应的字段在[高级鉴权](https://cloud.baidu.com/doc/CDN/s/ujwvyeo0t)。有非常消息的说明。

### 设置域名限速 SetLimitRate (废弃，请使用SetTrafficLimit)

> 限定此域名下向客户端传输的每份请求的最大响应速率。该速率是针对单个请求的，多请求自动翻倍。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置单请求限速1024Bytes/s
err := cli.SetLimitRate(testDomain, 1024)
fmt.Printf("err:%+v\n", err)

// 不做任何限速
err = cli.SetLimitRate(testDomain, 0)
fmt.Printf("err:%+v\n", err)
```

### 设置单连接限速 SetTrafficLimit/GetTrafficLimit

```go
cli := client.GetDefaultClient()
// 开启单连接限速，开启时间为北京时间上午10～12点之间，单链接限速值为1000KB左右。
trafficLimit := &api.TrafficLimit{
		Enable:           true,
		LimitRate:        1000,
		LimitStartHour:   10,
		LimitEndHour:     12,
		TrafficLimitUnit: "k",
	}
// 设置单连接限速
err := cli.SetTrafficLimit(testDomain, trafficLimit)

// 查询单连接限速配置
trafficLimit, err := cli.GetTrafficLimit(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("trafficLimit:%+v\n", trafficLimit)
```

`trafficLimit`是`api.TrafficLimit`类型的对象，详细说明如下：


| 字段    | 类型 | 说明                                                         |
| ------- | ---- | ------------------------------------------------------------ |
| Enabled | bool | true表示开启单链接限速，false表示关闭。当enable为false时，下面的配置均无效 |
| limitRate   | int  | 限速值，建议显示设置，否则效果未定义。单位为limitRateUnit设置的值，默认为Byte/s。 |
| limitStartHour | int | 限速开始时间，请输入0 - 24范围的数字，小于限速结束时间，默认值为0(可选) |
| limitEndHour | int | 限速结束时间，请输入0 - 24范围的数字，大于限速开始时间，默认值为24(可选) |
| limitRateAfter | int | 在发送了多少数据之后限速，单位Byte(可选) |
| trafficLimitArg | string | 限速参数名称，根据url中提取的arg进行限速(可选) |
| trafficLimitUnit | string | 限速参数单位，支持m、k、g，默认为Byte(可选) |

### 设置/查询Cors跨域 SetCors/GetCors

> 跨域访问是指发起请求的资源所在域不同于该请求所指向的资源所在域，出于安全考虑，浏览器会限制这种非同源的访问。开启此功能，用户可以自己进行清除缓存及跨域访问配置，当源站（BOS）对象更新后，CDN所有对应的缓存可进行同步自动更新。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置允许的跨域域名
err := cli.SetCors(testDomain, true, []string{
	"http://www.baidu.com",
	"http://*.bce.com",
})
fmt.Printf("err:%+v\n", err)

// 取消跨域设置
err = cli.SetCors(testDomain, false, nil)
fmt.Printf("err:%+v\n", err)

// 查询跨域设置
cors, err := cli.GetCors(testDomain)
fmt.Printf("cors:%+v\n", cors)
fmt.Printf("err:%+v\n", err)
```

### 设置/查询IP访问限频 SetAccessLimit/GetAccessLimit

> 限制IP单节点的每秒访问次数，针对所有的访问路径。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置单IP访问限频为200/s
err := cli.SetAccessLimit(testDomain, &api.AccessLimit{
	Enabled: true,
	Limit:   200,
})
fmt.Printf("err:%+v\n", err)

// 取消IP访问限频
err = cli.SetAccessLimit(testDomain, &api.AccessLimit{
	Enabled: false,
	Limit:   0,
})
fmt.Printf("err:%+v\n", err)

// 查询IP访问限频设置
accessLimit, err := cli.GetAccessLimit(testDomain)
fmt.Printf("accessLimit:%+v\n", accessLimit)
fmt.Printf("err:%+v\n", err)
```

`accessLimit`是`api.AccessLimit`类型的对象，详细说明如下：

| 字段    | 类型 | 说明                                                         |
| ------- | ---- | ------------------------------------------------------------ |
| Enabled | bool | true表示开启IP单节点访问限频，false表示取消限频。这里要注意golang的bool对象的零值为false，设置Limit的值必须要设置Enabled为true。 |
| Limit   | int  | 1秒内单个IP节点请求次数上限，enabled为true时此项默认为1000，enabled为false此项无意义。 |

### 设置/查询获取真实用户IP SetClientIp/GetClientIp

> 用户在使用CDN加速的同时可获取访问源的真实IP地址或客户端IP地址.。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置Client IP，源站可以获取到访问源的客户端IP地址，携带True-Client-Ip
err := cli.SetClientIp(testDomain, &api.ClientIp{
	Enabled: true,
	Name:    "True-Client-IP",
})
fmt.Printf("err:%+v\n", err)

// 设置Real IP：源站可以获取到访问源的真实IP地址，携带X-Real-IP。
err = cli.SetClientIp(testDomain, &api.ClientIp{
	Enabled: true,
	Name:    "X-Real-IP",
})
fmt.Printf("err:%+v\n", err)

// 关闭设置Client IP和Real IP
err = cli.SetClientIp(testDomain, &api.ClientIp{
	Enabled: false,
})
fmt.Printf("err:%+v\n", err)

// 查询关于客户端IP的设置
clientIp, err := cli.GetClientIp(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("clientIp:%+v\n", clientIp)
```

`clientIp`是`api.ClientIp`类型的对象，详细说明如下：

| 字段    | 类型   | 说明                                                         |
| ------- | ------ | ------------------------------------------------------------ |
| Enabled | bool   | true表示开启，false表示关闭。                                |
| Name    | string | 只能设置为"True-Client-Ip"或"X-Real-IP"两种之一，默认为"True-Client-Ip"，enabled为false时此项无意义。 |

### 设置/查询回源重试条件

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置当回源遇到429、500、502或者503错误码时进行重试
err := client.SetRetryOrigin(testDomain, &api.RetryOrigin{
		Codes: []int{429, 500, 502},
})
fmt.Printf("err:%+v\n", err)

// 查询回源重试策略
retryOrigin, err := client.GetRetryOrigin(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("retryOrigin:%+v\n", retryOrigin)
```


### 更新加速域名回源地址 SetDomainOrigin

> 加速域名的回源地址可以在创建域名的时候设置好，也可以在创建完成后进行更新。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

err := cli.SetDomainOrigin(testDomain, []api.OriginPeer{
	{
		Peer:      "1.1.1.1",
		Host:      "www.baidu.com",
		Backup:    true,
		Follow302: true,
	},
	{
		Peer:      "http://2.2.2.2",
		Host:      "www.baidu.com",
		Backup:    false,
		Follow302: true,
	},
}, "www.baidu.com")
fmt.Printf("err:%+v\n", err)
```

`api.OriginPeer`类型的详细说明在**创建加速域名一节**已经有说明。

### 设置/查询回源协议 SetOriginProtocol/GetOriginProtocol

> 设置回源协议指的是设置CDN侧接受请求miss的时候请求源站的协议，可以设置的回源协议有"http"、"https"或者"*"，"*"表示协议跟随回源(请求是HTTP/HTTPS协议那么请求源站也是HTTP/HTTPS协议)。默认是HTTP回源。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置回源协议为HTTP
err := cli.SetOriginProtocol(testDomain, "http")
fmt.Printf("err:%+v\n", err)

// 设置回源协议为HTTPS（域名必须HTTPS，否则以下请求失败）
err = cli.SetOriginProtocol(testDomain, "https")
fmt.Printf("err:%+v\n", err)

// 设置回源跟随协议
err = cli.SetOriginProtocol(testDomain, "*")
fmt.Printf("err:%+v\n", err)

// 查询回源协议
originProtocol, err := cli.GetOriginProtocol(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("originProtocol:%s\n", originProtocol)
```

### 设置协议跟随回源 SetFollowProtocol(废弃，设置协议跟随请使用SetOriginProtocol)

> 设置协议跟随回源，表示CDN节点回源协议与客户端访问协议保持一致。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置协议跟随回源
err := cli.SetFollowProtocol(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 取消设置协议跟随回源
err = cli.SetFollowProtocol(testDomain, false)
fmt.Printf("err:%+v\n", err)
```

### 设置/查询Range回源 SetRangeSwitch/GetRangeSwitch

> 设置Range回源，有助于减少大文件分发时回源消耗并缩短响应时间。此功能需源站支持Range请求。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置range回源
err := cli.SetRangeSwitch(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 取消设置range回源
err = cli.SetRangeSwitch(testDomain, false)
fmt.Printf("err:%+v\n", err)

// 查询range回源设置
rangeSwitch, err := cli.GetRangeSwitch(testDomain)
fmt.Printf("rangeSwitch:%+v\n", rangeSwitch)
```

### 设置/查询移动访问控制 SetMobileAccess/GetMobileAccess

> 开启移动访问控制，源站可有针对性地进行移动端／PC端的资源内容分发，暂不支持自定义进行移动端配置。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置移动访问
err := cli.SetMobileAccess(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 取消设置移动访问
err = cli.SetMobileAccess(testDomain, false)
fmt.Printf("err:%+v\n", err)

// 查询移动访问设置
mobileAccess, err := cli.GetMobileAccess(testDomain)
fmt.Printf("mobileAccess:%+v\n", mobileAccess)
fmt.Printf("err:%+v\n", err)
```

### 设置/查询HttpHeader SetHttpHeader/GetHttpHeader

> CDN支持CDN节点到客户端的response（HTTP响应头）、CDN节点到源站的request（HTTP请求头）中的header信息修改。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置CDN系统工作时增删的HTTP请求头
err := cli.SetHttpHeader(testDomain, []api.HttpHeader{
	{
		Type:   "origin",
		Header: "x-auth-cn",
		Value:  "xxxxxxxxx",
		Action: "remove",
	},
	{
		Type:   "response",
		Header: "content-type",
		Value:  "application/octet-stream",
		Action: "add",
	},
})
fmt.Printf("err:%+v\n", err)

// 取消CDN系统工作时增删的HTTP请求头
err = cli.SetHttpHeader(testDomain, []api.HttpHeader{})
fmt.Printf("err:%+v\n", err)

// 查询CDN系统工作时增删的HTTP请求头
headers, err := cli.GetHttpHeader(testDomain)
fmt.Printf("headers:%+v\n", headers)
fmt.Printf("err:%+v\n", err)
```

`headers`是`api.HttpHeader`类型的对象，详细说明如下：

| 字段     | 类型   | 说明                                                         |
| -------- | ------ | ------------------------------------------------------------ |
| Type     | string | "origin"表示此header 回源生效，"response"表示给用户响应时生效。 |
| Header   | string | header为http头字段，一般为HTTP的标准Header，也可以是用户自定义的；如x-bce-authoriztion。 |
| Value    | string | 指定Header的值。                                             |
| Action   | string | 表示是删除还是添加，可选remove/add，默认是add；目前console只支持add action; API做后端remove配置的兼容。 |
| Describe | string | 描述，可选，可以是中文，统一使用Unicode统码；长度不能超过100个字符。 |

### 设置/查询SEO开关属性 SetDomainSeo/GetDomainSeo

> SEO（Search Engine Optimization）优化是一种利用搜索引擎的规则提高网站在有关搜索引擎内的自然排名的方式。目前CDN系统支持两项优化配置：（1）搜索引擎开启自动回源；（2）数据与百度搜索链接。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置SEO优化
err := cli.SetDomainSeo(testDomain, &api.SeoSwitch{
	DirectlyOrigin: "ON",
	PushRecord:     "OFF",
})
fmt.Printf("err:%+v\n", err)

// 查询SEO优化设置
seoSwitch, err := cli.GetDomainSeo(testDomain)
fmt.Printf("seoSwitch:%+v\n", seoSwitch)
fmt.Printf("err:%+v\n", err)
```

`seoSwitch`是`api.SeoSwitch`类型的对象，详细说明如下：

| 字段           | 类型   | 说明                                  |
| -------------- | ------ | ------------------------------------- |
| DirectlyOrigin | string | ON表示设置直接回源，OFF则相反。       |
| PushRecord     | string | ON表示给大搜推送访问记录，OFF则相反。 |

### 设置/查询页面优化 SetFileTrim/GetFileTrim

> 用户开启页面优化功能，将自动删除 html中的注释以及重复的空白符，这样可以有效地去除页面的冗余内容，减小文件体积，提高加速分发效率。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置页面优化
err := cli.SetFileTrim(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 取消页面优化
err = cli.SetFileTrim(testDomain, false)
fmt.Printf("err:%+v\n", err)

// 查询页面优化设置
fileTrim, err := cli.GetFileTrim(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("fileTrim:%+v\n", fileTrim)
```

### 设置/查询IPv6开关 SetIPv6/GetIPv6

> 开启后，IPv6的客户端请求将支持以IPv6协议访问CDN，CDN也将携带IPv6的客户端IP信息访问您的源站。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 开启IPv6
err := cli.SetIPv6(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 关闭IPv6
err = cli.SetIPv6(testDomain, false)
fmt.Printf("err:%+v\n", err)

// 查询IPv6开关
ipv6Switch, err := cli.GetIPv6(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("ipv6Switch:%+v\n", ipv6Switch)
```

### 设置/查询QUIC开关 SetQUIC/GetQUIC

> Quick UDP Internet Connection(QUIC)协议是Google公司提出基于UDP的高效可靠的互联网传输层协议。 本接口用于查询特定域名的QUIC启用状态。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 开启QUIC
err := cli.SetQUIC(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 关闭QUIC
err = cli.SetQUIC(testDomain, false)
fmt.Printf("err:%+v\n", err)

// 查询QUIC开关
quicSwitch, err := cli.GetQUIC(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("quicSwitch:%+v\n", quicSwitch)
```

### 设置/查询离线模式 SetOfflineMode/GetOfflineMode

> 离线模式指的是在资源过期回源的时候，如果源站异常，那么CDN会响应之前缓存的资源，响应客户端200，但是回源日志中还是显示5xx。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 开启离线模式
err := cli.SetOfflineMode(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 关闭离线模式
err = cli.SetOfflineMode(testDomain, false)
fmt.Printf("err:%+v\n", err)

// 查询离线模式
offlineMode, err := cli.GetOfflineMode(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("offlineMode:%+v\n", offlineMode)
```

### 设置/查询视屏拖拽 SetMediaDrag/GetMediaDrag

> CDN支持flv与mp4视频类型的拖拽，开启拖拽可降低回源率，提升速度。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置视频拖拽设置
err := cli.SetMediaDrag(testDomain, &api.MediaDragConf{
	Mp4: &api.MediaCfg{
		DragMode: "second",
		FileSuffix: []string{
			"mp4",
			"m4a",
			"m4z",
		},
		StartArgName: "startIndex",
	},
	Flv: &api.MediaCfg{
		DragMode:   "byteAV",
		FileSuffix: []string{},
	},
})
fmt.Printf("err:%+v\n", err)

// 查询视频拖拽设置
mediaDragConf, err := cli.GetMediaDrag(testDomain)
fmt.Printf("mediaDragConf:%+v\n", mediaDragConf)
fmt.Printf("err:%+v\n", err)
```

`mediaDragConf`是`api.MediaDragConf`类型的对象，定义如下：

```go
type MediaDragConf struct {
	Mp4 *MediaCfg
	Flv *MediaCfg
}
```

可以设置Mp4或Flv类型视频流相关的拖拽，`MediaCfg`的详细说明如下：

| 字段         | 类型     | 说明                                                         |
| ------------ | -------- | ------------------------------------------------------------ |
| FileSuffix   | []string | CDN系统支持MP4文件的伪流（pseudo-streaming）播放，通常这些文件拓展名为.mp4，.m4v，.m4a，因此这个fileSuffix值为文件拓展名集合，如： ["mp4", "m4v", "m4a"]，type为mp4，fileSuffix默认值为["mp4"]；type为flv，fileSuffix默认值为["flv"] |
| StartArgName | string   | start参数名称，默认为“start”，您可以自定义参数名称，但是要求不能和`endArgName`相同 |
| EndArgName   | string   | end参数名称，默认为“end”，您可以自定义参数名称，但是要求不能和`startArgName`相同 |
| DragMode     | string   | mp4类型按秒进行拖拽，flv类型按字节进行拖拽。type为flv可选择的模式为“byteAV”或”byte”；type为mp4只能是"second"模式 |

### 设置/查询页面压缩 SetContentEncoding/GetContentEncoding

> 开启页面压缩功能后，您可以对大多数静态文件进行压缩，有效减少用户传输内容大小，加速分发效果。目前页面压缩支持Brotli压缩和Gzip压缩两种方式。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 设置页面压缩算法为gzip
err := cli.SetContentEncoding(testDomain, true, "gzip")
fmt.Printf("err:%+v\n", err)

// 设置页面压缩算法为br
err = cli.SetContentEncoding(testDomain, true, "br")
fmt.Printf("err:%+v\n", err)

// 关闭页面压缩
err = cli.SetContentEncoding(testDomain, false, "br")
fmt.Printf("err:%+v\n", err)

// 查询页面压缩算法，当关闭页面压缩时contentEncoding为空
contentEncoding, err := cli.GetContentEncoding(testDomain)
fmt.Printf("contentEncoding:%+v\n", contentEncoding)
```

### 设置HTTPS加速 SetDomainHttps

> 配置HTTPS的一个加速域名，必须要上传证书，了解证书详情请参考[证书管理](https://cloud.baidu.com/doc/Reference/s/8jwvz26si/)。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

err := cli.SetDomainHttps(testDomain, &api.HTTPSConfig{
	Enabled:          false,
	CertId:           "ssl-xxxxxx",
	Http2Enabled:     true,
	HttpRedirect:     true,
	HttpRedirectCode: 301,
})
fmt.Printf("err:%+v\n", err)

err = cli.SetDomainHttps(testDomain, &api.HTTPSConfig{
	Enabled: false,
})
fmt.Printf("err:%+v\n", err)
```

`api.HTTPSConfig`的结构比较复杂，详细说明如下：

| 字段              | 类型   | 说明                                                         |
| ----------------- | ------ | ------------------------------------------------------------ |
| Enabled           | bool   | 开启HTTPS加速，默认为false，当enabled=false，以下几列字段设置无效。 |
| CertId            | string | 当enabled=true时此项为必选，为SSL证书服务返回的证书ID，当enabled=False此项无效。 |
| HttpRedirect      | bool   | 为true时将HTTP请求重定向到HTTPS（重定向状态码为httpRedirectCode所配置），默认为false，当enabled=false此项无效，不可与httpsRedirect同时为true。 |
| HttpRedirectCode  | int    | 重定向状态码，可选值301/302，默认302，当enabled=false此项无效，httpRedirect=false此项无效。 |
| HttpsRedirect     | bool   | 为true时将HTTPS请求重定向到HTTP重定向状态码为httpsRedirectCode所配置），默认为false，当enabled=false此项无效，不可与httpRedirect同时为true。 |
| HttpsRedirectCode | int    | 重定向状态码，可选值301/302，默认302，当enabled=false此项无效，httpsRedirect=false此项无效。 |
| Http2Enabled      | bool   | 开启HTTP2特性，当enabled=false此项无效。必须要注意go的bool对象零值为false。 |
| ~~HttpOrigin~~        | bool   | **已弃用**，设置回源协议请参考**SetOriginProtocol**。 |
| SslVersion        | string | 设置TLS版本，默认为支持从TLSv1.0到TLSv1.3的版本，也可以设置为以下四个之一，SSLV3,TLSV1,TLSV11,TLSV12，当enabled=false时此项无效，此项一般取默认值，无需设置。 |

### 设置/查询OCSP SetOCSP/GetOCSP

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 开启OCSP
err := cli.SetOCSP(testDomain, true)
fmt.Printf("err:%+v\n", err)

// 关闭OCSP
err = cli.SetOCSP(testDomain, false)
fmt.Printf("err:%+v\n", err)

// 查询OCSP
ocspSwitch, err := cli.GetOCSP(testDomain)
fmt.Printf("err:%+v\n", err)
fmt.Printf("ocspSwitch:%+v\n", ocspSwitch)
```


### 设置/查询标签

> 标签是一组标识，您可以将具备某些共同点的域名进行归类，即打上标签，以此对资源进行组织。

一组标签定义为tags，他的类型如下：

| 字段  | 类型   | 说明  |
| :----------| :----- | :------- |
| tags | model.TagModel[] | 目标关联到的标签集合，需传入资源的全量标签信息，如原有3个标签，现删除1个标签的情况下新增2个标签，则应传入全量4个标签。当 tags 为空时，表示将域名关联的所有标签清空。标签个数最多为 10 个。 |


model.TagModel 类型如下：

| 参数             | 类型   | 说明                                                       |
| ---------------- | ------ | ---------------------------------------------------------- |
| tagKey          | string   |  标签的键，可包含大小写字母、数字、中文以及-_ /.特殊字符，长度 1-65                             |
| tagValue           | string | 标签的值，可包含大小写字母、数字、中文以及-_ /.特殊字符，长度 0-65         |


下面展示通过SDK方法配置标签与查询标签。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

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


## 证书管理接口

### 添加/修改域名证书 PutCert

> 给某个域名添加或修改证书，如果该域名已经绑定了一个证书，则该方法为修改（将用户新上传的证书替换掉证书库中的老证书，且给域名绑定新证书）。 如果域名之前没有绑定证书，则该方法为上传新证书。可自定义是否开启HTTPS。

```go
cli := client.GetDefaultClient()
// certData要带有证书链信息
certData := "-----BEGIN CERTIFICATE-----\nMIIFPTCCBCWgAwIBAgISBGRxBpT9H0OPrHDku006DU9SMA0GCSqGSIb3DQEBCwUA\nMDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD\nEwJSMzAeFw0yMTAzMTgwNzE0MzNaFw0yMTA2MTYwNzE0MzNaMB0xGzAZBgNVBAMM\nEiouY29kaW5nMzY1eDI0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBALM1g+BnRWl26auFUgAEU6L/2OWixH6K1BOp7JtlHW1RYM1I72pohuXyu0Hc\n20vOYpIQHpN1UY1e7bTKHsWiE/qUO/MeORbuslkYLBi3rZDE6i3aiYIFmaxjeYL9\nVQEJFrJI8X1AiIoP3MyiA4QsPotBH++3FyBb6sl5HtgqmCqItbsh6NV2FDN9+bm0\nz4hGSlJ1Wg9N03pXTaE8FT4AsCJzd/m+z+5u02aO5tEHWHiwrL4Yj/Y5H39x8/Ax\nHyaUjK62bgFySYH/XpU89dAKKo+uwOx5iBLXR8ni7yUj3y5NJ5A+AZUa61WbNNh0\n3jmgXff+s1zSwSiK+GP8q+Fz9+sCAwEAAaOCAmAwggJcMA4GA1UdDwEB/wQEAwIF\noDAdBgNVHSUEFjAUwYBBQUHAwIwDAYDVR0TAQH/BAIwADAd\nBgNVHQ4EFgQUArIIp83mJ+O7zRPwVr5ehX6JVJ4wHwYDVR0jBBgwFoAUFC6zF7dY\nVsuuUAlA5h+vnYsUwsYwVQYIKwYBBQUHAQEESTBHMCEGCCsGAQUFBzABhhVodHRw\nOi8vcjMuby5sZW5jci5vcmcwIgYIKwYBBQUHMAKGFmh0dHA6Ly9yMy5pLmxlbmNy\nLm9yZy8wLwYDVR0RBCgwJoISKi5jb2RpbmczNjV4MjQuY29tghBjb2RpbmczNjV4\nMjQuY29tMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYI\nKwYBBQUHAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBBQYKKwYBBAHW\neQIEAgSB9gSB8wDxAHcAXNxDkv7mq0VEsV6a1FbmEDf71fpH3KFzlLJe5vbHDsoA\nAAF4RGaGEA1Hh4zntDdVfmKVzQT5p/mQdczLsoQp2hmkHrKiTw\nl8cCIQDlceLBxn2RWzl+LD00gvTZDlqL/iWI/pAJ5qtTKzMH1QB2APZclC/RdzAi\nFFQYCDCUVo7jTRMZM7/fDC8gC8xO8WTjAAABeERmhrYAAAQDAEcwRQIhAKCycKu6\nNch4dkzO9gfQdjwhyCsaKi8nxNDgS199gp+eAiBkePK1AEwf+fvWGV+mXWDXcjjS\n6QCjL7w5lKi7CrVJLzANBgkqhkiG9w0BAQsFA0MKbiwTqtJdEb\nPaaATAtN/NXHoESO/KHFGjJT9ua1PByM0Qqn6mvonck+Tu9fWxM6ZOvXheEdcCbS\n5zLI4AJTdp5yySPRQe2v9UwVO6keDsk0Ux1JWFWgV7otwV5P22pDxaKOKqUQ11ZM\neu5Pk9dKFjlhT+oP88acHKVBJ0CJk/D72jWlDUhn4LwiEJ/+mfGW0oPu5ht+z0Zb\nQLMt7xTX6fATALSCELFPwVrBwleUpCYeafxAJC6XAMF1xeifFfZduORqPAUqyj7U\nYp6h3Wy+rJRKo0bN1roPxgu4aexXnOth6Qeyvi7zq7IO+jMY1/VaEBQJGS/hqity\nKQ==\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIEZTCCA02gAwIBAgIQQAF1BIMUpMghjISpDBbN3zANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTIwMTAwNzE5MjE0MFoXDTIxMDkyOTE5MjE0MFow\nMjELMAkGA1UEBhMCVVMxFjAUBgQxCzAJBgNVBAMT\nAlIzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuwIVKMz2oJTTDxLs\njVWSw/iC8ZmmekKIp10mqrUrucVMsa+Oa/l1yKPXD0eUFFU1V4yeqKI5GfWCPEKp\nTm71O8Mu243AsFzzWTjn7c9p8FoLG77AlCQlh/o3cbMT5xys4Zvv2+Q7RVJFlqnB\nU840yFLuta7tj95gcOKlVKu2bQ6XpUA0ayvTvGbrZjR8+muLj1cpmfgwF126cm/7\ngcWt0oZYPRfH5wm78Sv3htzB2nFd1EbjzK0lwYi8YGd1ZrPxGPeiXOZT/zqItkel\n/xMY6pgJdz+dU/nPAeX1pnAXFK9jpP+Zs5Od3FOnBv5IhR2haa4ldbsTzFID9e1R\noYvbFQIDAQABo4IBaDCCAWQwEgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8E\nBAMCAYYwSwYIKwYBBQUHAQEEPzA9ModHRwOi8vYXBwcy5p\nZGVudHJ1c3QuY29tL3Jvb3RzL2RzdHJvb3RjYXgzLnA3YzAfBgNVHSMEGDAWgBTE\np7Gkeyxx+tvhS5B1/8QVYIWJEDBUBgNVHSAETTBLMAgGBmeBDAECATA/BgsrBgEE\nAYLfEwEBATAwMC4GCCsGAQUFBwIBFiJodHRwOi8vY3BzLnJvb3QteDEubGV0c2Vu\nY3J5cHQub3JnMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly9jcmwuaWRlbnRydXN0\nLmNvbS9EU1RST09UQ0FYM0NSTC5jcmwwHQYDVR0OBBYEFBQusxe3WFbLrlAJQOYf\nr52LFMLGMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjANBgkqhkiG9w0B\nAQsFAAOCAQEA2UzgyfWEiDcx27sT4rP8i2tiEmxYt0l+PAK3qB8oYevO4C5z70kH\nejWEHx2taPDY/laBL21/WKZuNTYQHHPVvCadTQsvd8\nS8MXjohyc9z9/G2948kLjmE6Flh9dDYrVYA9x2O+hEPGOaEOa1eePynBgPayvUfL\nqjBstzLhWVQLGAkXXmNs+5ZnPBxzDJOLxhF2JIbeQAcH5H0tZrUlo5ZYyOqA7s9p\nO5b85o3AM/OJ+CktFBQtfvBhcJVd9wvlwPsk+uyOy2HI7mNxKKgsBTt375teA2Tw\nUdHkhVNcsAKX1H7GNNLOEADksd86wuoXvg==\n-----END CERTIFICATE-----"

privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAszWD4GdFaXbpq4VSAARTov/Y5aLEforUE6nsm2UdbVFgzUjv\namiG5fK7QdzbS85ikhAek3VRjV7ttMoexaIT+pQ78x45Fu6yWRgsGLetkMTqLdqJ\nggWZrGN5gv1VAQkWskjxfUCIig/czKIDhCw+i0Ef77cXIFvqyXke2CqYKoi1uyHo\n1XYUM335ubTPiEZKUnVaD03Te7P7m7TZo7m0QdYeLCsvhiP\n9jkff3Hz8DEfJpSMrrZuAXJJgf9elTz10Aoqj67A7HmIEtdHyeLvJSPfLk0nkD4B\nlRrrVZs02HTeOaBd9/6zXNLBKIr4Y/yr4XP36wIDAQABAoIBAHnKZcyNApxRJy7d\nFURTrG97RvGRM874FHckpVtaVaxkgNAiwCrlzL/bva1eJl8XbN/tOopmUb0tBYk3\nT8BqjP9f3Ho2UQAnymdISTenJLrdSHVPLuKBYdXJaNw/xJRGk/koH45K3EBP1XPw\nq0kZNIw4/zZPjNT+AstXmEGApnLPFIoiEuB8grbfqQfr\nvY5ywxkXZdKQzeUr2ZVEPBp27q6lT13L0VBCCtOxZNNXPca+Lv\nJKlfvXBckuoyqHaTythwnw1PFPjq+uVMWfaa6gttfGGr11NVPnucRHwCMWUO0Cba\nfPw5ShECgYEA6fmMLOK/xRvaax7Hkr/In7IXUwSR1yNYupSeo+OF7C/+lkvvdwJP\nj/An9So2wVWY+DK3gKB7GCrDjYHB7Lv0m0dw7b+vEXv39z86mwUZsD8rMHy+HuqH\nZ4NieBkARWDaP/iE3HbrsWUHog7YvOZLr529byTRRQZdga5R1Vmf5RkCgYEAxBQx\nTPD6PypvVpQ4fkg58EJsFKwkGYTgcfS9dfwsjWbOTNW5gM/hCKgI/LTjsU7a0YYH\nvEMCTim81kuUKZ9rYlfXqB1Xzk9k6mUFeoP4t5KCe79YFi1A\nBUxqHUArJiLyO1g8cZruKK37DQKpT0sYHw7AAaMCgYBuSNUc1yiDVTSn51M0xbdg\nJsa9t9qyaJPLJoB8SaN3h8vdth9CnlE4TH/ZHLPAf4NiAi3isEI1Svrv+WiaGKIc\nixkcx4xSlnd0EFakeUv5elz2NuY6lluKnDBO4aHyEcvt+UtOy7Me47ssVQkuSPMF\n7Tk8aUNG4NA0byFdiihHCQKBgQCBXOkh4CLaJb8LGgMjnbdMAiaYhPHUPExwIo4V\nF2i1acxV+PPIPl4zfdlgEF/gjSvk7E6SMIuG0haaM4bu5xTL7zSC38kcflkQI9Ja\nWy6WYJFh6kjsdfguDn+7bVfVGOqexv/j/wRLhBhzsrxvZOO\negbHjQKBgEe3ghxGKW8dl3+/PDZ3KF8YBb5xUxe9BO8ufM0Pe+tCn8iWrTeAHG16\ncl9JgGlN9eIgx9VOh7suKlb9SsZLbAN60IO9nIx23g2nLqH+HyEZmE6zK5onSxua\n9vgcKjhNmg5WLvmQwz0ECw050HtDpptawvfNPinUY1LsjvkWqwiX\n-----END RSA PRIVATE KEY-----"
certId, err := testCli.PutCert("my.domain.com", &api.UserCertificate{
		CertName:    "test",
		ServerData:  certData,
		PrivateData: privateKey,
	}, "ON")
fmt.Printf("certId:%s\n", certId)
fmt.Printf("err:%+v\n", err)
```

### 查询域名证书 GetCert

> 查询某个域名的证书信息，如果证书不存在此方法会返回404错误。

```go
cli := client.GetDefaultClient()
detail, err := testCli.GetCert("my.domain.com")
fmt.Printf("detail:%v\n", detail)
// {"certId":"cert-8j774s9y3ww2","certName":"test","status":"IN_USE","certCommonName":"*.domain.com","certDNSNames":"*.domain.com,domain.com","certStartTime":"2021-03-18T07:14:33Z","certStopTime":"2021-06-16T07:14:33Z","certCreateTime":"2021-04-25T06:50:55Z","certUpdateTime":"2021-04-25T14:50:54Z"}
fmt.Printf("err:%+v\n", err)
```

### 删除域名证书 DeleteCert

> 删除某个域名的证书，且关闭HTTPS。如果该域名原来没有证书，那么什么都不会做。

```go
cli := client.GetDefaultClient()
err := testCli.DeleteCert("my.domain.com")
fmt.Printf("err:%+v\n", err)
```

## 缓存管理接口

### 刷新缓存/查询刷新状态 Purge/GetPurgedStatus

> 缓存清除方式有URL刷新、目录刷新除。URL刷新除是以文件或一个资源为单位进行缓存刷新。目录刷新除是以目录为单位，将目录下的所有文件进行缓存清除。

```go
cli := client.GetDefaultClient()

// 刷除
purgedId, err := cli.Purge([]api.PurgeTask{
	{
		Url: "http://my.domain.com/path/to/purge/2.data",
	},
	{
		Url:  "http://my.domain.com/path/to/purege/html/",
		Type: "directory",
	},
})
fmt.Printf("purgedId:%+v\n", purgedId)
fmt.Printf("err:%+v\n", err)

// 根据任务ID查询刷除状态
purgedStatus, err := cli.GetPurgedStatus(&api.CStatusQueryData{
	Id: string(purgedId),
})

fmt.Printf("purgedStatus:%+v\n", purgedStatus)
fmt.Printf("err:%+v\n", err)
```

示例中刷除了两类资源，第一种是刷除一个文件，第二种是刷除某个目录的所有的文件。根据Purge返回的task id去查询任务进度。`api.CStatusQueryData`是一个相对较复杂的结构，可以根据不同的条件查询，具体可以查看定义。

### 预热资源/查询预热状态 Prefetch/GetPrefetchStatus

> URL预热是以文件为单位进行资源预热。

```go
cli := client.GetDefaultClient()

prefetchId, err := cli.Prefetch([]api.PrefetchTask{
	{
		Url: "http://my.domain.com/path/to/purge/1.data",
	},
})
fmt.Printf("prefetchId:%+v\n", prefetchId)
fmt.Printf("err:%+v\n", err)


prefetchStatus, err := cli.GetPrefetchStatus(&api.CStatusQueryData{
	Id: string(prefetchId),
})
fmt.Printf("prefetchStatus:%+v\n", prefetchStatus)
fmt.Printf("err:%+v\n", err)
```

### 查询刷新/预热限额 GetQuota

```go
cli := client.GetDefaultClient()
quotaDetail, err := cli.GetQuota()
fmt.Printf("quotaDetail:%+v\n", quotaDetail)
fmt.Printf("err:%+v\n", err)
```

`quotaDetail`是`api.QuotaDetail`类型的对象，详细说明如下：

| 字段      | 类型 | 说明                            |
| --------- | ---- | ------------------------------- |
| DirRemain | int  | 当日刷新目录限额余量。          |
| UrlRemain | int  | 当日刷新（含预热）URL限额余量。 |
| DirQuota  | int  | 刷新目录限额总量。              |
| UrlQuota  | int  | 刷新（含预热）URL限额总量。     |

## 动态加速接口

### 配置动态加速服务 EnableDsa/DisableDsa

> 开启/关闭DSA是针对用户级别的开启关闭。

```go
cli := client.GetDefaultClient()

// 开启DSA服务
err := cli.EnableDsa()
fmt.Printf("err:%+v\n", err)

// 关闭DSA服务
err = cli.DisableDsa()
fmt.Printf("err:%+v\n", err)
```

### 查询动态加速域名列表 ListDsaDomains

> 查询某个用户配置了DSA加速规则的域名列表。

```go
cli := client.GetDefaultClient()
dsaDomains, err := cli.ListDsaDomains()
fmt.Printf("dsaDomains:%+v\n", dsaDomains)
fmt.Printf("err:%+v\n", err)
```

`dsaDomains`是string数组，代表配置了DSA加速规则的域名。

### 配置域名动态加速规则 SetDsaConfig

> 配置某个域名的DSA加速规则。

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"

// 配置DSA规则
err := cli.SetDsaConfig(testDomain, &api.DSAConfig{
	Enabled: true,
	Rules: []api.DSARule{
		{
			Type:  "suffix",
			Value: ".mp4;.jpg;.php",
		},
		{
			Type:  "path",
			Value: "/path",
		},
		{
			Type:  "exactPath",
			Value: "/path/to/file.mp4",
		},
	},
	Comment: "test",
})
fmt.Printf("err:%+v\n", err)

// 取消DSA规则
err = cli.SetDsaConfig(testDomain, &api.DSAConfig{
	Enabled: false,
})
fmt.Printf("err:%+v\n", err)
```

`api.DSAConfig`的详细说明如下：

| 字段  | 类型   | 说明                                                         |
| ----- | ------ | ------------------------------------------------------------ |
| Type  | string | "suffix"表示文件类型，"path"表示动态路径，“exactPath“表示动态URL。 |
| Value | string | Type所指定类型的配置规则，多条规则使用";"分割。              |


## 日志接口

### 获取单个域名日志 GetDomainLog

```go
cli := client.GetDefaultClient()
testDomain := "test_go_sdk.baidu.com"
endTime := "2019-09-01T07:12:00Z"
startTime := "2019-09-09T07:18:00Z"
domainLogs, err := cli.GetDomainLog(testDomain, api.TimeInterval{
	StartTime: startTime,
	EndTime:   endTime,
})

fmt.Printf("domainLogs:%+v\n", domainLogs)
fmt.Printf("err:%+v\n", err)
```

示例查询了单个域名在2019-09-01T07:12:00Z～2019-09-09T07:18:00Z之间的日志，`domainLogs`是`api.LogEntry`数组类型，LogEntry包含日志名，所属域名，下载路径，起始时间等信息。

### 获取多个域名日志 GetMultiDomainLog

```go
cli := client.GetDefaultClient()
endTime := "2019-09-01T07:12:00Z"
startTime := "2019-09-09T07:18:00Z"

domainLogs, err := cli.GetMultiDomainLog(&api.LogQueryData{
	TimeInterval: api.TimeInterval{
		StartTime: startTime,
		EndTime:   endTime,
	},
	Type:    1,
	Domains: []string{"1.baidu.com", "2.baidu.com"},
})

fmt.Printf("domainLogs:%+v\n", domainLogs)
fmt.Printf("err:%+v\n", err)
```

示例查询["1.baidu.com", "2.baidu.com"]这些域名的日志，`domainLogs`和上一节GetDomainLog返回格式一致。

## 工具接口

### IP检测 GetIpInfo

> 验证指定的IP是否属于百度开放云CDN服务节点。

```go
cli := client.GetDefaultClient()
ipStr := "1.2.3.4"
ipInfo, err := cli.GetIpInfo(ipStr, "describeIp")

fmt.Printf("ipInfo:%+v\n", ipInfo)
fmt.Printf("err:%+v\n", err)
```

其中GetIpInfo的第二个参数只能为**"describeIp"**，ipInfo包含IP的详细信息，包括区域和ISP，如果不属于百度开放云CDN的节点，区域和ISP都为空。

### 批量IP检测接口 GetIpListInfo

> 验证多个IP是否属于百度开放云CDN服务节点。

```go
cli := client.GetDefaultClient()
ipStr := "1.2.3.4"
ipsInfo, err := cli.GetIpListInfo([]string{"116.114.98.35", "59.24.3.174"}, "describeIp")

fmt.Printf("ipsInfo:%+v\n", ipInfo)
fmt.Printf("err:%+v\n", err)
```

第二个参数只能为**"describeIp"**

### 获取百度云CDN的回源节点信息 GetBackOriginNodes

```go
cli := client.GetDefaultClient()
backOriginNodes, err := testCli.GetBackOriginNodes()

fmt.Printf("backOriginNodes:%+v\n", backOriginNodes)
fmt.Printf("err:%+v\n", err)
```

## 统计查询

### 通用统计接口

> 通用统计接口在文档[统计接口](https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn)中有详细说明。

`api.QueryCondition`结构包含了最基本的查询条件，如下：

| **参数**  | **类型** | **说明**                                                     |
| --------- | -------- | ------------------------------------------------------------ |
| StartTime | string   | 查询的时间范围起始值，默认为endTime前推24小时。格式为UTC时间字符串，如："2019-09-01T07:12:00Z"。 |
| EndTime   | string   | 查询的时间范围结束值，默认为当前时间。时间跨度最长90天，时间格式和StartTime一样。 |
| Period    | int      | 查询结果的粒度，单位秒，可选值为60,300,3600,86400；默认为300，uv 默认3600（选60s的时间粒度时建议StartTime和EndTime区间跨度建议选择0.5到1h，否则可能会因为数据量太大无法正常返回） |
| KeyType   | int      | 标识key的内容，0=>域名，1=>用户id，2=>tag，默认0。           |
| Key       | []string | 域名、用户Id或Tag。                                          |
| GroupBy   | string   | 返回结果聚合粒度，key => 根据key聚合，最后的key是`total`， 空 => 返回整体结果，每个key的每个时间段都对应一组数据。 |

`metric`表示查询的统计数据类型，如下，接口具体返回结果的格式在对应的函数中可以看到。

| metric           | 函数                | 接口类型                           | 额外参数                                                     |
| ---------------- | ------------------- | ---------------------------------- | ------------------------------------------------------------ |
| avg_speed        | GetAvgSpeed         | 查询平均速率                       | 无。                                                         |
| avg_speed_region | GetAvgSpeedByRegion | 客户端访问分布查询平均速率         | prov和isp。prov是查询的省份全拼，默认为空，查询全国数据。isp是查询的运营商代码，默认为空，查询所有运营商数据。 |
| pv               | GetPv               | pv/qps查询                         | level，查询边缘节点或者中心节点pv。可填写"all"或"edge"或者"internal"，默认为“all”。 |
| pv_src           | GetSrcPv            | 回源pv/qps查询                     | 无。                                                         |
| pv_region        | GetPvByRegion       | 查询pv/qps(分客户端访问分布)       | prov和isp。prov是查询的省份全拼，默认为空，查询全国数据。isp是查询的运营商代码，默认为空，查询所有运营商数据。 |
| uv               | GetUv               | uv查询                             | 无。                                                         |
| flow             | GetFlow             | 查询流量、带宽                     | level，查询边缘节点或者中心节点带宽。可填写"all"或"edge"或"internal"，默认为"all"。 |
| flow_protocol    | GetFlowByProtocol   | 查询流量、带宽(分协议)             | protocol，查询http或https的流量、带宽, 取值"http", "https"或者 "all"，默认"all"。 |
| flow_region      | GetFlowByRegion     | 查询流量、带宽（分客户端访问分布） | prov和isp。prov是查询的省份全拼，默认为空，查询全国数据。isp是查询的运营商代码，默认为空，查询所有运营商数据。 |
| src_flow         | GetSrcFlow          | 查询回源流量、回源带宽             | 无。                                                         |
| real_hit         | GetRealHit          | 字节命中率查询                     | 无。                                                         |
| pv_hit           | GetPvHit            | 请求命中率查询                     | 无。                                                         |
| httpcode         | GetHttpCode         | 状态码统计查询                     | 无。                                                         |
| src_httpcode     | GetSrcHttpCode      | 回源状态码查询                     | 无。                                                         |
| httpcode_region  | GetHttpCodeByRegion | 状态码统计查询（分客户端访问分布） | prov和isp。prov是查询的省份全拼，默认为空，查询全国数据。isp是查询的运营商代码，默认为空，查询所有运营商数据。 |
| top_urls         | GetTopNUrls         | TopN urls                          | extra，查询指定http状态码的记录，默认值： ""。               |
| top_referers     | GetTopNReferers     | TopN referers                      | extra，查询指定http状态码的记录，默认值： ""。               |
| top_domains      | GetTopNDomains      | TopN domains                       | extra，查询指定http状态码的记录，默认值： ""。               |
| error            | GetError            | cdn错误码分类统计查询              | 无。                                                         |

### 计费统计接口

#### 查询域名或者tag的95带宽 GetPeak95Bandwidth

查询条件：
| **参数**  | **类型** | **说明**                                                     |
| --------- | -------- | ------------------------------------------------------------ |
| StartTime | string   | 查询的时间范围起始值，默认为endTime前推24小时。格式为UTC时间字符串，如："2019-09-01T07:12:00Z"。 |
| EndTime   | string   | 查询的时间范围结束值，默认为当前时间。时间跨度最长90天，时间格式和StartTime一样。 |
| domains   | []string | 域名集合，和tags互斥存在，设置了domains请设置tags为nil |
| tags      | []string | tag集合，和domains互斥存在，设置了tags请设置domains为nil |

请求示例：

```go
cli := client.GetDefaultClient()
peak95Time, peak95Band, err := cli.GetPeak95Bandwidth(
		"2020-05-01T00:00:00Z", "2020-05-10T00:00:00Z", nil, []string{"www.test.com"})

fmt.Printf("peak95Time:%s\n", peak95Time)
fmt.Printf("peak95Band:%d\n", peak95Band)
fmt.Printf("err:%+v\n", err)
```