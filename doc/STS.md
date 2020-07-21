# STS服务

# 概述

本文档主要介绍STS服务的使用。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。
若您还不了解STS，可以参考[产品描述](https://cloud.baidu.com/doc/IAM/s/xjwvybxhv)和[操作指南](https://cloud.baidu.com/doc/IAM/s/njwvyc2zd)。

# 使用方法

## 确认Endpoint

目前使用STS服务时，STS的Endpoint都统一使用`http://sts.bj.baidubce.com`，为默认值。

## 获取密钥

要使用百度云STS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问服务做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 创建Client对象并获取临时token

STS Client是STS服务的客户端，为开发者与STS服务进行交互提供了获取临时token的方法。示例代码如下：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/sts" //导入STS服务模块
)

func main() {
	// 创建STS服务的Client对象，Endpoint使用默认值
	AK, SK := "<your-access-key-id>", "<your-secret-access-key>"
	stsClient, err := sts.NewClient(AK, SK)
	if err != nil {
		fmt.Println("create sts client object :", err)
		return
	}

	// 获取临时认证token，有效期为60秒，ACL为空
	obj, err := stsClient.GetSessionToken(60, "")
	if err != nil {
		fmt.Println("get session token failed:", err)
		return
	}
	fmt.Println("GetSessionToken result:")
	fmt.Println("  accessKeyId:", obj.AccessKeyId)
	fmt.Println("  secretAccessKey:", obj.SecretAccessKey)
	fmt.Println("  sessionToken:", obj.SessionToken)
	fmt.Println("  createTime:", obj.CreateTime)
	fmt.Println("  expiration:", obj.Expiration)
	fmt.Println("  userId:", obj.UserId)
}
```

## 获取关联指定角色的临时身份凭证

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/sts" //导入STS服务模块 
	"github.com/baidubce/bce-sdk-go/services/sts/api"
)

func main() {
	// 创建STS服务的Client对象，Endpoint使用默认值
	AK, SK := "<your-access-key-id>", "<your-secret-access-key>"
	stsClient, err := sts.NewClient(AK, SK)
	if err != nil {
		fmt.Println("create sts client object :", err)
		return
	}

	// 获取临时认证token，有效期为60秒，ACL为空
    AccountId, RoleName := "<your-account-id>", "<your-assume-role-name>"
	args := &api.AssumeRoleArgs{
		AccountId: AccountId,
		RoleName:  RoleName,
    }
    obj, err := client.AssumeRole(args)
    if err != nil {
    	return
    }
	fmt.Println("GetSessionToken result:")
	fmt.Println("  accessKeyId:", obj.AccessKeyId)
	fmt.Println("  secretAccessKey:", obj.SecretAccessKey)
	fmt.Println("  sessionToken:", obj.SessionToken)
	fmt.Println("  createTime:", obj.CreateTime)
	fmt.Println("  expiration:", obj.Expiration)
	fmt.Println("  userId:", obj.UserId)
	fmt.Println("  roleId:", obj.RoleId)
}
```