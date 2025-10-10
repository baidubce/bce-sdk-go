# BCT服务

# 概述

本文档主要介绍BCT GO SDK的使用。在使用本文档前，您需要先了解BCT的一些基本知识。若您还不了解BCT，可以参考[产品描述](https://cloud.baidu.com/doc/BCT/s/1jwvxn6wt)。

# 初始化

## 获取密钥

要使用百度云BCT，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BCT做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/#/iam/accesslist)

## 新建BCT Client

BCT Client是BCT服务的客户端，为开发者与BCT服务进行交互提供了一系列的方法。

### 使用AK/SK新建BCT Client

通过AK/SK方式访问BCT，用户可以参考如下代码新建一个BCT Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bct"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 初始化一个BCTClient
	bctClient, err := bct.NewClient(AK, SK)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/IAM/s/ojwvynrqn)》。

# 主要接口

## 事件查询

### 事件查询
通过以下代码可以查询BCT事件

```go

    start,_ := time.Parse(time.RFC3339, "2025-09-02T16:00:00Z")
    end,_ := time.Parse(time.RFC3339, "2025-09-03T16:00:00Z")
    args := &api.QueryEventsV2Request{
        StartTime: start,
        EndTime:   end,
        PageSize:  10,
	}
    res, err := bctClient.QueryEventsV2(args)
    if err != nil {
        log.Fatal(err)
	}
    log.Info(json.Marshal(res))	
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考BCT API 文档[事件查询接口](https://cloud.baidu.com/doc/BCT/s/Kkdsj7q07#%E6%9F%A5%E8%AF%A2v2%E6%8E%A5%E5%8F%A3)
