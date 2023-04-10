# IAM服务

# 概述

本文档主要介绍普通型IAM GO SDK的使用。在使用本文档前，您需要先了解IAM的一些基本知识。若您还不了解IAM，可以参考[产品描述](https://cloud.baidu.com/doc/IAM/s/xjwvybxhv)和[应用场景](https://cloud.baidu.com/doc/IAM/s/Djwvybxus)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[IAM访问域名](https://cloud.baidu.com/doc/IAM/s/cjwvxnzix)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云IAM，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问IAM做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/#/iam/accesslist)

## 新建IAM Client

普通型IAM Client是IAM服务的客户端，为开发者与IAM服务进行交互提供了一系列的方法。

### 使用AK/SK新建IAM Client

通过AK/SK方式访问IAM，用户可以参考如下代码新建一个IAM Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/iam"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 初始化一个IAMClient
	iamClient, err := iam.NewClient(AK, SK)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/IAM/s/ojwvynrqn)》。

### 使用STS创建IAM Client

**申请STS token**

IAM可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问IAM，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建IAM Client**

申请好STS后，可将STS Token配置到IAM Client中，从而实现通过STS Token创建IAM Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建IAM Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"            //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/iam"    //导入IAM服务模块
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

	// 使用申请的临时STS创建IAM服务的Client对象，Endpoint使用默认值
	iamClient, err := iam.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create iam client failed:", err)
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
	iamClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置IAM Client时，无论对应IAM服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问IAM

IAM支持HTTPS传输协议，您可以通过在创建IAM Client对象时指定的Endpoint中指明HTTPS的方式，在IAM GO SDK中使用HTTPS访问IAM服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/iam"

ENDPOINT := "https://iam.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
iamClient, _ := iam.NewClientWithEndpoint(AK, SK, ENDPOINT)
```

## 配置IAM Client

如果用户需要配置IAM Client的一些细节的参数，可以在创建IAM Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问IAM服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/iam"

//创建IAM Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "iam.bj.baidubce.com
client, _ := iam.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/iam"

AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := iam.NewClient(AK, SK)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/iam"

AK, SK := <your-access-key-id>, <your-secret-access-key>
client, _ := iam.NewClient(AK, SK)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问IAM时，创建的IAM Client对象的`Config`字段支持的所有参数如下表所示：

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

1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建IAM Client”小节。
2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。

# 主要接口

## 用户管理

### 创建用户
通过以下代码可以创建子用户

```go

    name := "test-user-sdk-go"
    args := &api.CreateUserArgs{
		Name:        name,
        Description: "description",
    }
    
    result, err := client.CreateUser(args)
    if err != nil {
        fmt.Println("Create iam user failed", err)
    } else {
        fmt.Println("Create iam user success", result)
    }	
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[CreateUser创建用户](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%9B%E5%BB%BA%E7%94%A8%E6%88%B7)

### 查询用户
通过以下代码可以查询单个子用户
```go
	name := "test-user-sdk-go"
	result, err := client.GetUser(name)
	if err != nil {
		fmt.Println("Get iam user failed", err)
	} else {
		fmt.Println("Get iam user success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[GetUser查询用户](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7)

### 更新用户
通过以下代码可以更新子用户

```go
    name := "test-user-sdk-go"
	args := &api.UpdateUserArgs{
		Description: "newDescription",
	}

	result, err := client.UpdateUser(name, args)
	if err != nil {
		fmt.Println("Update iam user failed", err)
	} else {
		fmt.Println("Update iam user success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[UpdateUser更新用户](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%9B%E5%BB%BA%E7%94%A8%E6%88%B7)

### 列举用户
通过以下代码可以列举子用户
```go
	result, err := client.ListUser()
	if err != nil {
		fmt.Println("List iam user failed", err)
	} else {
		fmt.Println("List iam user success", result)
	}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[ListUser创建用户](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%97%E4%B8%BE%E7%94%A8%E6%88%B7)

### 删除用户
通过以下代码可以更新子用户

```go
	name := "test-user-sdk-go"
	err = client.DeleteUser(name)
	if err != nil {
		fmt.Println("Delete iam user failed", err)
	} else {
		fmt.Println("Delete iam user success", name)
	}
```

> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[DeleteUser删除用户](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%A0%E9%99%A4%E7%94%A8%E6%88%B7)


### 配置用户控制台登录
通过以下代码可以配置用户的控制台登录，为其配置登录密码、开启登录MFA、配置第三方账号绑定等

```go
	name := "test-user-sdk-go-login-profile"
	args := &api.UpdateUserLoginProfileArgs{
		Password:        "1@3Qwe4f",
		EnabledLoginMfa: false,
		LoginMfaType:    "PHONE",
	}

	result, err := client.UpdateUserLoginProfile(name, args)
	if err != nil {
		fmt.Println("Update iam user login profile failed", err)
	} else {
		fmt.Println("Update iam user login profile success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[UpdateUserLoginProfile配置用户控制台登录](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E9%85%8D%E7%BD%AE%E7%94%A8%E6%88%B7%E7%9A%84%E6%8E%A7%E5%88%B6%E5%8F%B0%E7%99%BB%E5%BD%95)

### 查询控制台登录配置
通过以下代码可以查询用户的控制台登录配置
```go
	name := "test-user-sdk-go-login-profile"
	result, err := client.GetUserLoginProfile(name)
	if err != nil {
		fmt.Println("Get iam user login profile failed", err)
	} else {
		fmt.Println("Get iam user login profile success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[GetUserLoginProfile查询用户控制台登录](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E6%9F%A5%E8%AF%A2%E6%8E%A7%E5%88%B6%E5%8F%B0%E7%99%BB%E5%BD%95%E9%85%8D%E7%BD%AE)

### 关闭控制台登录配置
关闭用户的控制台登录配置，即关闭用户的控制台登录
```go
	name := "test-user-sdk-go-login-profile"
	err = client.DeleteUserLoginProfile(name)
	if err != nil {
		fmt.Println("Delete iam user login profile failed", err)
	} else {
		fmt.Println("Delete iam user login profile success", name)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[DeleteUserLoginProfile关闭用户控制台登录](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%85%B3%E9%97%AD%E6%8E%A7%E5%88%B6%E5%8F%B0%E7%99%BB%E5%BD%95%E9%85%8D%E7%BD%AE)

### 创建用户的AccessKey
通过以下代码为用户创建一组AccessKey访问密钥

```go
	name := "test-user-sdk-go-accessKey"
	result, err := client.CreateAccessKey(name)
	if err != nil {
		fmt.Println("Create accessKey failed", err)
	} else {
		fmt.Println("Create accessKey success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[创建用户的AccessKey](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%9B%E5%BB%BA%E7%94%A8%E6%88%B7%E7%9A%84accesskey)

### 禁用用户的AccessKey
通过以下代码为禁用用户的AccessKey
```go
	name := "test-user-sdk-go-accessKey"
	accessKeyId := "<your-access-key-id>"
	result, err := client.DisableAccessKey(name, accessKeyId)
	if err != nil {
		fmt.Println("Disable accessKey failed", err)
	} else {
		fmt.Println("Disable accessKey success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[禁用用户的AccessKey](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%9B%E5%BB%BA%E7%94%A8%E6%88%B7%E7%9A%84accesskey)

### 启用用户的AccessKey
通过以下代码为启用用户的AccessKey
```go
	name := "test-user-sdk-go-accessKey"
	accessKeyId := "<your-access-key-id>"
	result, err := client.EnableAccessKey(name, accessKeyId)
	if err != nil {
		fmt.Println("Enable accessKey failed", err)
	} else {
		fmt.Println("Enable accessKey success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[启用用户的AccessKey](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%90%AF%E7%94%A8%E7%94%A8%E6%88%B7%E7%9A%84accesskey)

### 删除用户的AccessKey
删除用户的指定一组AccessKey访问密钥
```go
	name := "test-user-sdk-go-accessKey"
	accessKeyId := "<your-access-key-id>"
	err = client.DeleteAccessKey(name, accessKeyId)
	if err != nil {
		fmt.Println("Delete accessKey failed", err)
	} else {
		fmt.Println("Delete accessKey success", accessKeyId)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[删除用户的AccessKey](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%A0%E9%99%A4%E7%94%A8%E6%88%B7%E7%9A%84accesskey)

### 列举用户的AccessKey
列举用户的全部AccessKey访问密钥
```go
	name := "test-user-sdk-go-accessKey"
	result, err := client.ListAccessKey(name)
	if err != nil {
		fmt.Println("List accessKey failed", err)
	} else {
		fmt.Println("List accessKey success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举用户的AccessKey](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E5%88%97%E4%B8%BE%E7%94%A8%E6%88%B7%E7%9A%84accesskey)


### 创建组
通过以下代码创建组
```go
	name := "test_group_sdk_go"
	args := &api.CreateGroupArgs{
		Name:        name,
		Description: name,
	}

	result, err := client.CreateGroup(args)
	if err != nil {
		fmt.Println("Create group failed", err)
	} else {
		fmt.Println("Create group success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[创建组](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E5%88%9B%E5%BB%BA%E7%BB%84)


### 查询组
通过以下代码查询组
```go
	name := "test_group_sdk_go"
	result, err := client.GetGroup(name)
	if err != nil {
		fmt.Println("Get group failed", err)
	} else {
		fmt.Println("Get group success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[查询组](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E6%9F%A5%E8%AF%A2%E7%BB%84)


### 更新组
通过以下代码更新组
```go
	name := "test_group_sdk_go"
	args := &api.UpdateGroupArgs{
		Description: "newDes",
	}
	result, err := client.UpdateGroup(name, args)
	if err != nil {
		fmt.Println("Update group failed", err)
	} else {
		fmt.Println("Update group success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[更新组](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E6%9B%B4%E6%96%B0%E7%BB%84)

### 删除组
通过以下代码删除组
```go
	name := "test_group_sdk_go"
	err = client.DeleteGroup(name)
	if err != nil {
		fmt.Println("Delete group failed", err)
	} else {
		fmt.Println("Delete group success", name)
	}

```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[删除组](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E5%88%A0%E9%99%A4%E7%BB%84)

### 列举组
通过以下代码删除组
```go
	result, err := client.ListGroup()
	if err != nil {
		fmt.Println("Delete group failed", err)
	} else {
		fmt.Println("Delete group success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举组](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E5%88%97%E4%B8%BE%E7%BB%84)

### 添加用户到组
通过以下代码添加用户到组
```go
	userName := "test_user_sdk_go"
	groupName := "test_user_sdk_go"
	err = client.AddUserToGroup(userName, groupName)
	if err != nil {
		fmt.Println("Add user to group failed", err)
	} else {
		fmt.Println("Add user to group success", userName)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[添加用户到组](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E6%B7%BB%E5%8A%A0%E7%94%A8%E6%88%B7%E5%88%B0%E7%BB%84)


### 从组内移除用户
通过以下代码从组内移除用户
```go
	userName := "test_user_sdk_go"
    groupName := "test_user_sdk_go"
    err = client.DeleteUserFromGroup(userName, groupName)
    if err != nil {
        fmt.Println("Add user to group failed", err)
    } else {
        fmt.Println("Add user to group success", userName)
    }
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[从组内移除用户](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E4%BB%8E%E7%BB%84%E5%86%85%E7%A7%BB%E9%99%A4%E7%94%A8%E6%88%B7)

### 列举用户的组
通过以下代码列举用户的组
```go
    userName := "test_user_sdk_go"
	result, err := client.ListGroupsForUser(userName)
	if err != nil {
		fmt.Println("List groups for user failed", err)
	} else {
		fmt.Println("List groups for user success", result)
	}
	
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举用户的组](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E5%88%97%E4%B8%BE%E7%94%A8%E6%88%B7%E7%9A%84%E7%BB%84)

### 列举组内用户
通过以下代码列举组内用户
```go
	groupName := "test_user_sdk_go"
	result, err := client.ListUsersInGroup(groupName)
	if err != nil {
		fmt.Println("List user in group failed", err)
	} else {
		fmt.Println("List user in group success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举组内用户](https://cloud.baidu.com/doc/IAM/s/ljx35h8lx#%E5%88%97%E4%B8%BE%E7%BB%84%E5%86%85%E7%94%A8%E6%88%B7)

### 创建角色
通过以下代码创建角色
```go
    roleName := "test_role_sdk_go"
	args := &api.CreateRoleArgs{
		Name:        roleName,
		Description: "description",
		AssumeRolePolicyDocument: "{\"version\":\"v1\",\"accessControlList\":[{\"service\":\"bce:iam\",\"permission\"" +
			":[\"AssumeRole\"],\"region\":\"*\",\"grantee\":[{\"id\":\"grantee-id\"}],\"effect\":\"Allow\"}]}",
	}

	result, err := client.CreateRole(args)
	if err != nil {
		fmt.Println("Create role failed", err)
	} else {
		fmt.Println("Create role success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[创建角色](https://cloud.baidu.com/doc/IAM/s/ek5eq1zp1)

### 查询角色
通过以下代码查询角色
```go
	roleName := "test_role_sdk_go"
	result, err := client.GetRole(roleName)
	if err != nil {
		fmt.Println("Get role failed", err)
	} else {
		fmt.Println("Get role success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[查询角色](https://cloud.baidu.com/doc/IAM/s/ek5eq1zp1#%E6%9F%A5%E8%AF%A2%E8%A7%92%E8%89%B2)

### 更新角色
通过以下代码查询角色
```go
	args := &api.UpdateRoleArgs{
		Description: "newDescription",
		AssumeRolePolicyDocument: "{\"version\":\"v1\",\"accessControlList\":[{\"service\":\"bce:iam\",\"permission\"" +
			":[\"AssumeRole\"],\"region\":\"*\",\"grantee\":[{\"id\":\"grantee-id\"}],\"effect\":\"Allow\"}]}",
	}

	roleName := "test_role_sdk_go"
	result, err := client.UpdateRole(roleName, args)
	if err != nil {
		fmt.Println("Update role failed", err)
	} else {
		fmt.Println("Update role success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[更新角色](https://cloud.baidu.com/doc/IAM/s/ek5eq1zp1#%E6%9B%B4%E6%96%B0%E8%A7%92%E8%89%B2)

### 删除角色
通过以下代码查询角色
```go
	roleName := "test_role_sdk_go"
	err = client.DeleteRole(roleName)
	if err != nil {
		fmt.Println("Delete role failed", err)
	} else {
		fmt.Println("Delete role success", roleName)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[删除角色](https://cloud.baidu.com/doc/IAM/s/ek5eq1zp1#%E5%88%A0%E9%99%A4%E8%A7%92%E8%89%B2)

### 列举角色
通过以下代码列举角色
```go
	result, err := client.ListRole()
	if err != nil {
		fmt.Println("List role failed", err)
	} else {
		fmt.Println("List role success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举角色](https://cloud.baidu.com/doc/IAM/s/ek5eq1zp1#%E5%88%97%E4%B8%BE%E8%A7%92%E8%89%B2)

### 创建策略
通过以下代码创建策略
```go

name := "test_sdk_go_policy"
args := &api.CreatePolicyArgs{
    Name:        name,
    Description: "description",
    Document:    "{\"accessControlList\": [{\"region\":\"bj\",\"service\":\"bcc\"," +
"\"resource\":[\"*\"],\"permission\":[\"*\"],\"effect\":\"Allow\"}]}",
}

result, err := client.CreatePolicy(args)
if err != nil {
    fmt.Println("Update policy failed", err)
} else {
    fmt.Println("Update policy success", result)
}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[创建策略](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%88%9B%E5%BB%BA%E7%AD%96%E7%95%A5)

### 查询策略
通过以下代码查询策略
```go
    name := "test_sdk_go_policy"
	policyType := "Custom"
	result, err := client.GetPolicy(name, policyType)
	if err != nil {
		fmt.Println("Update policy failed", err)
	} else {
		fmt.Println("Update policy success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[创建策略](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%88%9B%E5%BB%BA%E7%AD%96%E7%95%A5)

### 删除策略
通过以下代码删除策略
```go
	name := "test_sdk_go_policy"
	err = client.DeletePolicy(name)
	if err != nil {
		fmt.Println("List policy failed", err)
	} else {
		fmt.Println("List policy success", name)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[删除策略](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%88%A0%E9%99%A4%E7%AD%96%E7%95%A5)


### 列举策略
通过以下代码列举策略
```go
	name := "test_sdk_go_policy"
	policyType := "Custom"
	result, err := client.ListPolicy(name, policyType)
	if err != nil {
		fmt.Println("List policy failed", err)
	} else {
		fmt.Println("List policy success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举策略](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%88%97%E4%B8%BE%E7%AD%96%E7%95%A5)

### 关联用户权限
通过以下代码关联用户权限
```go
	userName := "test_sdk_go_user"
	policyName := "test_sdk_go_policy"
	args := &api.AttachPolicyToUserArgs{
		UserName:   userName,
		PolicyName: policyName,
	}
	err = client.AttachPolicyToUser(args)
	if err != nil {
		fmt.Println("Attach policy to user failed", err)
	} else {
		fmt.Println("Attach policy to user success", args)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[关联用户权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%85%B3%E8%81%94%E7%94%A8%E6%88%B7%E6%9D%83%E9%99%90)

### 解除用户权限
通过以下代码解除用户权限
```go
	userName := "test_sdk_go_user"
	policyName := "test_sdk_go_policy"
	args := &api.DetachPolicyFromUserArgs{
		UserName:   userName,
		PolicyName: policyName,
	}
	err = client.DetachPolicyFromUser(args)
	if err != nil {
		fmt.Println("Detach policy to user failed", err)
	} else {
		fmt.Println("Detach policy to user success", args)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[解除用户权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E8%A7%A3%E9%99%A4%E7%94%A8%E6%88%B7%E6%9D%83%E9%99%90)

### 列举用户的权限
通过以下代码列举用户的权限
```go
	userName := "test_sdk_go_user"
	result, err := client.ListUserAttachedPolicies(userName)
	if err != nil {
		fmt.Println("List user attached policy failed", err)
	} else {
		fmt.Println("List user attached policy success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举用户的权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%88%97%E4%B8%BE%E7%94%A8%E6%88%B7%E7%9A%84%E6%9D%83%E9%99%90)

### 关联组权限
通过以下代码关联组权限
```go
	groupName := "test_sdk_go_group"
	policyName := "test_sdk_go_policy"
	args := &api.AttachPolicyToGroupArgs{
		GroupName:  groupName,
		PolicyName: policyName,
	}
	err = client.AttachPolicyToGroup(args)
	if err != nil {
		fmt.Println("Attach policy to group failed", err)
	} else {
		fmt.Println("Attach policy to group success", args)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[关联组权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%85%B3%E8%81%94%E7%BB%84%E6%9D%83%E9%99%90)

### 解除组权限
通过以下代码解除组权限
```go
	groupName := "test_sdk_go_group"
	policyName := "test_sdk_go_policy"
	args := &api.DetachPolicyFromGroupArgs{
		GroupName:  groupName,
		PolicyName: policyName,
	}
	err = client.DetachPolicyFromGroup(args)
	if err != nil {
		fmt.Println("Detach policy to group failed", err)
	} else {
		fmt.Println("Detach policy to group success", args)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[解除组权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E8%A7%A3%E9%99%A4%E7%BB%84%E6%9D%83%E9%99%90)

### 列举组权限
通过以下代码列举组权限
```go
	groupName := "test_sdk_go_group"
	result, err := client.ListGroupAttachedPolicies(groupName)
	if err != nil {
		fmt.Println("List group attached policy failed", err)
	} else {
		fmt.Println("List group attached policy success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举组权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%88%97%E4%B8%BE%E7%BB%84%E7%9A%84%E6%9D%83%E9%99%90)

### 关联角色权限
通过以下代码关联角色权限
```go
	roleName := "test_sdk_go_group"
	policyName := "test_sdk_go_policy"
	args := &api.AttachPolicyToRoleArgs{
		RoleName:   roleName,
		PolicyName: policyName,
	}
	err = client.AttachPolicyToRole(args)
	if err != nil {
		fmt.Println("Attach policy to role failed", err)
	} else {
		fmt.Println("Attach policy to role success", args)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[关联角色权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%85%B3%E8%81%94%E8%A7%92%E8%89%B2%E6%9D%83%E9%99%90)

### 解除角色权限
通过以下代码关联角色权限
```go
	roleName := "test_sdk_go_group"
	policyName := "test_sdk_go_policy"
	args := &api.DetachPolicyToRoleArgs{
		RoleName:   roleName,
		PolicyName: policyName,
	}
	err = client.DetachPolicyFromRole(args)
	if err != nil {
		fmt.Println("Detach policy to role failed", err)
	} else {
		fmt.Println("Detach policy to role success", args)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[解除角色权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E8%A7%A3%E9%99%A4%E8%A7%92%E8%89%B2%E6%9D%83%E9%99%90)

### 列举角色的权限
通过以下代码列举角色权限
```go
	roleName := "test_sdk_go_group"
	result, err := client.ListRoleAttachedPolicies(roleName)
	if err != nil {
		fmt.Println("List role attached policy failed", err)
	} else {
		fmt.Println("List role attached policy success", result)
	}
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[列举角色的权限](https://cloud.baidu.com/doc/IAM/s/Wjx35jxes#%E5%88%97%E4%B8%BE%E8%A7%92%E8%89%B2%E7%9A%84%E6%9D%83%E9%99%90)

### 修改子用户操作保护
通过以下代码修改子用户操作保护
```go
    userName := "test-user-sdk-go-switch-operation-mfa"
    enableMfa := true
    mfaType := "PHONE,TOTP"
    args := &api.UserSwitchMfaArgs{
        UserName:   userName,
        EnabledMfa: enableMfa,
        MfaType:    mfaType,
    }
    err := IAM_CLIENT.UserOperationMfaSwitch(args)
    if err != nil {
        fmt.Println("switch user mfa failed", err)
    } else {
        fmt.Println("switch user mfa success", result)
    }
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[UserOperationMfaSwitch修改子用户操作保护](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E4%BF%AE%E6%94%B9%E5%AD%90%E7%94%A8%E6%88%B7%E6%93%8D%E4%BD%9C%E4%BF%9D%E6%8A%A4)


### 修改子用户密码
通过以下代码修改子用户密码
```go
    userName := "test-user-name-sdk-go-sub-update"
    Password := "Baidu@123"
    args := &api.UpdateSubUserArgs{
        Password: Password,
    }
    res, err := IAM_CLIENT.SubUserUpdate(userName, args)
    if err != nil {
        fmt.Println("update sub user failed", err)
    } else {
        fmt.Println("update sub user success", result)
    }
```
> **提示：**
> - 详细的参数配置及限制条件，可以参考IAM API 文档[SubUserUpdate修改子用户密码](https://cloud.baidu.com/doc/IAM/s/mjx35fixq#%E4%BF%AE%E6%94%B9%E5%AD%90%E7%94%A8%E6%88%B7%E5%AF%86%E7%A0%81)

# 错误处理

GO语言以error类型标识错误，IAM支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | IAM服务返回的错误

## 客户端异常

客户端异常表示客户端尝试向IAM发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当IAM服务端出现异常时，IAM服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[IAM错误返回](https://cloud.baidu.com/doc/IAM/s/Rjx4d0rxo)

# 版本变更记录

## v0.9.11 [2022-10-13]

首次发布：

- 创建、查看、列表、更新、删除IAM用户
- 配置、查询、关闭用户控制台配置
- 创建、查看、列表、删除、启用、禁用AccessKey
- 创建、查看、列表、更新、删除IAM用户组
- 创建、列表、列表、更新、删除角色
- 创建、查看、列表、更新、删除、关联用户权限、解除用户权限、列举用户权限、关联组权限、解除组权限、列举组权限、关联角色权限、解除角色权限、列举角色的权限

## v0.9.12 [2023-02-13]

- 修改子用户的操作保护
- 修改子用户密码
