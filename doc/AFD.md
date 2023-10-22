# AFD服务

# 概述

本文档主要介绍昊天镜业务风控系统AFD GO SDK的使用。关于昊天镜业务风控系统的相关信息请查看[产品描述](https://cloud.baidu.com/doc/AFD/index.html)。

# 接口服务使用

## 确认Endpoint

如果无特殊要求，可使用全局Endpoint：afd.baidubce.com

## 获取密钥

您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问AFD做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/#/iam/accesslist)

## 服务调用

### 使用AK/SK新建AFD Client

通过AK/SK方式访问AFD，用户可以参考如下代码新建一个AFD Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/afd"
)

func main() {
	AK := ""
	SK := ""

	// 初始化一个AFD Client
	afdClient, err := afd.NewClient(AK, SK)

	// 指定ENDPOINT
	// ENDPOINT := "afd.baidubce.com"
	// afdClient, err := afd.NewClient(AK, SK, ENDPOINT)

	if err != nil {
		panic(err)
	}
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果未指定，则会使用AFD默认域名的服务地址。


### 调用活动防刷

**代码示例**

```go
fmt.Println(afdClient.Sync(&afd.SyncArgs{
	SC:    "bce_activity",
	TS:    "1658832547594",
	M:     "91945a133f8b61348d16e9dc7c9644c5acead8a0",
	IP:    "4.4.4.4",
	App:   "ios",
	AppID: "1111",
	AID:   "2222",
	EV:    "topic",
}))

```

### 调用风险设备查询

**代码示例**

```go
fmt.Println(afdClient.Factor(&afd.FactorArgs{
	App: "android",
	Z:   "HnHqkHiVEF4B3RAKNZp1m2Cx36KRL4dH2UBmxjB-ha6Lt8RcreMAp93mgCosAOEb5R-5HVk7r1GxXcqKoBJFLew",
}))

```

### 接口参数信息

更多详见[各类API接口参数](https://cloud.baidu.com/doc/AFD/s/nkjmi532w)
