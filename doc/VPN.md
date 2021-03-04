# VPN服务

# 概述

本文档主要介绍VPN GO SDK的使用。在使用本文档前，您需要先了解VPN的一些基本知识，并已开通了VPN服务。若您还不了解VPN，可以参考[产品描述](https://cloud.baidu.com/doc/VPC/s/sjwvytvh0)和[操作指南](https://cloud.baidu.com/doc/VPC/s/9jwvytzz8)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[VPN服务域名](https://cloud.baidu.com/doc/VPC/s/xjwvyuhpw)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

目前支持“华北-北京”、“华南-广州”、“华东-苏州”、“香港”、“金融华中-武汉”和“华北-保定”六个区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
BJ | bcc.bj.baidubce.com | HTTP and HTTPS
GZ | bcc.gz.baidubce.com | HTTP and HTTPS
SU | bcc.su.baidubce.com | HTTP and HTTPS
HKG| bcc.hkg.baidubce.com| HTTP and HTTPS
FWH| bcc.fwh.baidubce.com| HTTP and HTTPS
BD | bcc.bd.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云VPN，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问VPN做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建VPN Client

VPN Client是VPN服务的客户端，为开发者与VPN服务进行交互提供了一系列的方法。

### 使用AK/SK新建VPN Client

通过AK/SK方式访问VPN，用户可以参考如下代码新建一个VPN Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个VPNClient
	vpnClient, err := vpn.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为VPC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`bcc.bj.baidubce.com`。

### 使用STS创建VPN Client

**申请STS token**

VPN可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问VPN，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建VPN Client**

申请好STS后，可将STS Token配置到VPN Client中，从而实现通过STS Token创建VPN Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建VPN Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/vpn" //导入VPN服务模块
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

	// 使用申请的临时STS创建VPN服务的Client对象，Endpoint使用默认值
	vpnClient, err := vpn.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Println("create vpn client failed:", err)
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
	vpnClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置VPN Client时，无论对应VPN服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问VPN

VPN支持HTTPS传输协议，您可以通过在创建VPN Client对象时指定的Endpoint中指明HTTPS的方式，在VPN GO SDK中使用HTTPS访问VPN服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

ENDPOINT := "https://bcc.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
vpnClient, _ := vpn.NewClient(AK, SK, ENDPOINT)
```

## 配置VPN Client

如果用户需要配置VPN Client的一些细节的参数，可以在创建VPN Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问VPN服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

//创建VPN Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := vpn.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := vpn.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := vpn.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问VPN时，创建的VPN Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建VPN Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。



## 创建VPN

使用以下代码可以申请一个VPN。
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

args := &vpn.CreateVpnGatewayArgs{
	        VpnName:     "TestSDK-VPN",
    		Description: "vpn test",
    		VpcId:       "vpcId",
    		Billing: &Billing{
    			PaymentTiming: PAYMENT_TIMING_PREPAID,
    			Reservation: &Reservation{
    				ReservationLength:   1,
    				ReservationTimeUnit: "month",
    			},
    		},
    		ClientToken: getClientToken(),
}
result, err := client.CreateVpnGateway(args)
if err != nil {
    fmt.Printf("create vpn error: %+v\n", err)
    return
}

fmt.Println("create vpn success, vpn: ", result.VpnId)
```


## 查询VPN 列表

使用以下代码可以查询VPN列表。
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

args := &vpn.ListVpnGatewayArgs{
		MaxKeys: 1000,
		VpcId:   "vpcId",
	}
	result, err := client.ListVpnGateway(args)
    if err != nil {
        fmt.Printf("list vpn error: %+v\n", err)
        return
    }
   // 返回标记查询的起始位置
   fmt.Println("vpn list marker: ", result.Marker)
   // true表示后面还有数据，false表示已经是最后一页
   fmt.Println("vpn list isTruncated: ", result.IsTruncated)
   // 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
   fmt.Println("vpn list nextMarker: ", result.NextMarker)
   // 每页包含的最大数量
   fmt.Println("vpn list maxKeys: ", result.MaxKeys)
   // 获取vpn的列表信息
   for _, e := range res.Vpns {
   		fmt.Println("vpn id: ", e.VpnId)
   		fmt.Println("vpn eip: ", e.Eip)
   		fmt.Println("vpn status: ", e.Status)
   		fmt.Println("vpn vpcId: ", e.VpcId)
   		fmt.Println("vpn description: ", e.Description)
   		fmt.Println("vpn expiredTime: ", e.ExpiredTime)
   		fmt.Println("vpn paymentTiming: ", e.ProductType)
   		fmt.Println("vpn vpnConnNum: ", e.VpnConnNum)
   		fmt.Println("vpn bandwidthInMbps: ", e.BandwidthInMbps)
   		fmt.Println("vpn vpnName: ", e.Name)
   		fmt.Println("vpn expireTime: ", e.ExpiredTime)
   	}
```

## 查询VPN详情

使用以下代码可以实现查询VPN的详情信息。
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"
result,err:=vpn.GetVpnGatewayDetail("vpnId")
        
        if  err != nil {
            fmt.Printf("get vpn detail error: %+v\n", err)
            return
        }
        fmt.Println("vpn id: ", result.VpnId)
   		fmt.Println("vpn eip: ", result.Eip)
   		fmt.Println("vpn status: ", result.Status)
   		fmt.Println("vpn vpcId: ", result.VpcId)
   		fmt.Println("vpn description: ", result.Description)
   		fmt.Println("vpn expiredTime: ", result.ExpiredTime)
   		fmt.Println("vpn paymentTiming: ", result.ProductType)
   		fmt.Println("vpn vpnConnNum: ", result.VpnConnNum)
   		fmt.Println("vpn bandwidthInMbps: ", result.BandwidthInMbps)
   		fmt.Println("vpn vpnName: ", result.Name)
   		fmt.Println("vpn expireTime: ", result.ExpiredTime)
}
```

## 更新VPN网关

使用以下代码可以实现VPN网关的更新。
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"
args := &vpn.UpdateVpnGatewayArgs{
		ClientToken: getClientToken(),
		Name:        "vpnTest",
	}
	err := client.UpdateVpnGateway("vpnId", args)
if  err != nil {
    fmt.Printf("update vpn error: %+v\n", err)
    return
}
fmt.Printf("update vpn success\n")
```


## 释放VPN

使用以下代码可以释放指定的VPN。
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

err = client.DeleteVpn(vpnId, clientToken)
if err != nil {
    fmt.Printf("delete vpn error: %+v\n", err)
    return
}

fmt.Printf("delete vpn success\n")
```

> 注意:
> - 释放指定VPN，被释放的VPN无法找回
> - 预付费购买的VPN如需提前释放，请通过工单进行

## 绑定EIP

使用以下代码可以将EIP绑定到VPN。
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

args := &vpn.BindEipArgs{
	ClientToken: ClientToken(),
	Eip:         Eip,
	}
err := client.BindEip(vpnId, args)
if err != nil {
	fmt.Printf("bind eip error: %+v\n", err)
	return
}
```

> 注意:
> - 绑定EIP是一个异步过程，可以通过查询VPN的状态判断绑定是否成功


## 解绑EIP

使用以下代码可以将VPN的EIP进行解绑。
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"
if err := client.UnBindEip(vpnId, clientToken); err != nil {
    fmt.Printf("unbind eip error: %+v\n", err)
    return
}

fmt.Printf("unbind eip success.")
```

> 注意: 解绑EIP是一个异步过程，可以通过查询VPN的状态判断绑定是否成功

## VPN网关续费
使用以下代码可以延长VPN的到期时间
 ```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"
args := &vpn.RenewVpnGatewayArgs{
	ClientToken: ClientToken,
	Billing: &Billing{
		Reservation: &Reservation{
			ReservationLength:   1,
			ReservationTimeUnit: "month",
		},
	},
}
if err := client.RenewVpnGateway(vpnId, args); err != nil {
    fmt.Printf(" renew vpn error: %+v\n", err)
    return
}

fmt.Printf("renew vpn success.")
```
 
> 注意:
> - 仅预付费资源可以进行续费操作。

## 创建VPN隧道
使用以下代码可以为指定的VPN创建隧道
 ```go
 // import "github.com/baidubce/bce-sdk-go/services/vpn"
 args := &vpn.CreateVpnConnArgs{
 		VpnId:         VpnId,
 		VpnConnName:   VpnConnName,
 		LocalIp:      LocalIp,
 		SecretKey:     SecretKey,
 		LocalSubnets:  []string{"subnet"},
 		RemoteIp:      RemoteIp,
 		RemoteSubnets: []string{"RemoteSubnet"},
 		CreateIkeConfig: &CreateIkeConfig{
 			IkeVersion:  "v1",
 			IkeMode:     "main",
 			IkeEncAlg:   "aes",
 			IkeAuthAlg:  "sha1",
 			IkePfs:      "group2",
 			IkeLifeTime: 25500,
 		},
 		CreateIpsecConfig: &CreateIpsecConfig{
 			IpsecEncAlg:   "aes",
 			IpsecAuthAlg:  "sha1",
 			IpsecPfs:      "group2",
 			IpsecLifetime: 25500,
 		},
 	}
 	res, err := client.CreateVpnConn(args)
    if  err != nil {
        fmt.Printf(" create vpnconn error: %+v\n", err)
        return
    }
    fmt.Printf(" create vpnconn success,connId is: %+v",res.VpnConnId)
 ```
## 查询VPN隧道
使用一下代码可以查询指定VPN的隧道信息
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"
    result,err:=vpn.ListVpnConn("vpnId")
      if  err != nil {
          fmt.Printf("get vpn detail error: %+v\n", err)
          return
      }
      for _, e := range result.VpnConns {
      		fmt.Println("vpnconn Id: ", e.VpnConnId)
      		fmt.Println("vpnconn LocalIp: ",e.LocalIp)
      		fmt.Println("vpnconn Description: ",e.Description)
      		fmt.Println("vpnconn CreatedTime: ",e.CreatedTime)
      		fmt.Println("vpnconn HealthStatus: ",e.HealthStatus)
      		fmt.Println("vpnconn LocalIp: ",e.LocalIp)
      		fmt.Println("vpnconn LocalSubnets: ",e.LocalSubnets)
      		fmt.Println("vpnconn RemoteSubnets: ",e.RemoteSubnets)
      		fmt.Println("vpnconn IkeConfig: ",e.IkeConfig)
      		fmt.Println("vpnconn IpsecConfig: ",e.IpsecConfig)
      	}
```
## 更新VPN隧道
使用以下代码可以修改指定的VPN隧道
 ```go
 // import "github.com/baidubce/bce-sdk-go/services/vpn"
 args := &vpn.UpdateVpnConnArgs{
 		vpnConnId: vpnConnId,
 		updateVpnconn: &CreateVpnConnArgs{
 			VpnId:         VpnId,
 			VpnConnName:   VpnConnName,
 			LocalIp:       LocalIp,
 			SecretKey:     SecretKey,
 			LocalSubnets:  []string{"LocalSubnets"},
 			RemoteIp:      RemoteIp,
 			RemoteSubnets: []string{"RemoteSubnets"},
 			CreateIkeConfig: &CreateIkeConfig{
 				IkeVersion:  "v1",
 				IkeMode:     "main",
 				IkeEncAlg:   "aes",
 				IkeAuthAlg:  "sha1",
 				IkePfs:      "group2",
 				IkeLifeTime: 25500,
 			},
 			CreateIpsecConfig: &CreateIpsecConfig{
 				IpsecEncAlg:   "aes",
 				IpsecAuthAlg:  "sha1",
 				IpsecPfs:      "group2",
 				IpsecLifetime: 25500,
 			},
 		},
 	}
 	err := client.UpdateVpnConn(args)
    if  err != nil {
        fmt.Printf(" uodate vpnconn error: %+v\n", err)
        return
    }
    fmt.Printf(" update vpnconn success,connId is: %+v",res.VpnConnId)
 ```
## 删除VPN隧道
```go
// import "github.com/baidubce/bce-sdk-go/services/vpn"

err = client.DeleteVpnConn(vpnconnId, clientToken)
if err != nil {
    fmt.Printf("delete vpnconn error: %+v\n", err)
    return
}

fmt.Printf("delete vpnconn success\n")
```
# 错误处理

GO语言以error类型标识错误，VPN支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | VPN服务返回的错误

用户使用SDK调用VPN相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// vpnClient 为已创建的VPN Client对象
args := &vpn.ListVpnGatewayArgs{}
result, err := client.ListVpnGateway(args)
if err != nil {
	switch realErr := err.(type) {
	case *bce.BceClientError:
		fmt.Println("client occurs error:", realErr.Error())
	case *bce.BceServiceError:
		fmt.Println("service occurs error:", realErr.Error())
	default:
		fmt.Println("unknown error:", err)
	}
} 
```

## 客户端异常

客户端异常表示客户端尝试向VPN发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当VPN服务端出现异常时，VPN服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[VPN错误码](https://cloud.baidu.com/doc/VPC/s/sjwvyuhe7)

## SDK日志

VPN GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

VPN GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

```
// import "github.com/baidubce/bce-sdk-go/util/log"

// 指定输出到标准错误，输出INFO及以上级别
log.SetLogHandler(log.STDERR)
log.SetLogLevel(log.INFO)

// 指定输出到标准错误和文件，DEBUG及以上级别，以1GB文件大小进行滚动
log.SetLogHandler(log.STDERR | log.FILE)
log.SetLogDir("/tmp/gosdk-log")
log.SetRotateType(log.ROTATE_SIZE)
log.SetRotateSize(1 << 30)

// 输出到标准输出，仅输出级别和日志消息
log.SetLogHandler(log.STDOUT)
log.SetLogFormat([]string{log.FMT_LEVEL, log.FMT_MSG})
```

说明：
  1. 日志默认输出级别为`DEBUG`
  2. 如果设置为输出到文件，默认日志输出目录为`/tmp`，默认按小时滚动
  3. 如果设置为输出到文件且按大小滚动，默认滚动大小为1GB
  4. 默认的日志输出格式为：`FMT_LEVEL, FMT_LTIME, FMT_LOCATION, FMT_MSG`

### 项目使用

该日志模块无任何外部依赖，用户使用GO SDK开发项目，可以直接引用该日志模块自行在项目中使用，用户可以继续使用GO SDK使用的包级别的日志对象，也可创建新的日志对象，详见如下示例：

```
// 直接使用包级别全局日志对象（会和GO SDK自身日志一并输出）
log.SetLogHandler(log.STDERR)
log.Debugf("%s", "logging message using the log package in the VPN go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the VPN go sdk")
```


# 版本变更记录

## v0.9.5  [2020-05-26]

首次发布:

 - 支持创建VPN、查询VPN列表、查询VPN详情、更新VPN、释放VPN、绑定/解绑EIP、创建VPN隧道、查询VPN隧道、更新VPN隧道、删除VPN隧道。