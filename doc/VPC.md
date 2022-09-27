# VPC服务

# 概述

本文档主要介绍VPC GO SDK的使用。在使用本文档前，您需要先了解VPC的一些基本知识，并已开通了VPC服务。若您还不了解VPC，可以参考[产品描述](https://cloud.baidu.com/doc/VPC/s/Vjwvytu2v)和[操作指南](https://cloud.baidu.com/doc/VPC/s/qjwvyu0at)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[VPC访问域名](https://cloud.baidu.com/doc/VPC/s/xjwvyuhpw)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

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

要使用百度云VPC，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问VPC做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建VPC Client

VPC Client是VPC服务的客户端，为开发者与VPC服务进行交互提供了一系列的方法。

### 使用AK/SK新建VPC Client

通过AK/SK方式访问VPC，用户可以参考如下代码新建一个VPC Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个VpcClient
	vpcClient, err := vpc.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为VPC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`http://bcc.bj.baidubce.com`。

### 使用STS创建VPC Client

**申请STS token**

VPC可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问VPC，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建VPC Client**

申请好STS后，可将STS Token配置到VPC Client中，从而实现通过STS Token创建VPC Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建VPC Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/vpc" //导入VPC服务模块
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

	// 使用申请的临时STS创建VPC服务的Client对象，Endpoint使用默认值
	vpcClient, err := vpc.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Println("create vpc client failed:", err)
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
	vpcClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置VPC Client时，无论对应VPC服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问VPC

VPC支持HTTPS传输协议，您可以通过在创建VPC Client对象时指定的Endpoint中指明HTTPS的方式，在VPC GO SDK中使用HTTPS访问VPC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/vpc"

ENDPOINT := "https://bcc.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
vpcClient, _ := vpc.NewClient(AK, SK, ENDPOINT)
```

## 配置VPC Client

如果用户需要配置VPC Client的一些细节的参数，可以在创建VPC Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问VPC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/vpc"

//创建VPC Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := vpc.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/vpc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := vpc.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/vpc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := vpc.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问VPC时，创建的VPC Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建VPC Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# VPC管理

私有网络(Virtual private Cloud，VPC) 是一个用户能够自定义的虚拟网络，能够帮助用户构建属于自己的网络环境。通过指定IP地址范围和子网等配置，即可快速创建一个VPC，不同的VPC之间完全隔离，用户可以在VPC内创建和管理BCC实例。

## 创建VPC

通过以下代码可以创建VPC实例:
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.CreateVPCArgs{
	// 设置创建vpc使用的名称
    Name:        "test-vpc",
    // 设置创建vpc使用的描述信息
    Description: "test-vpc-description",
    // 设置创建vpc使用的cidr
    Cidr:        "102.168.0.0/24",
    // 设置创建vpc使用的标签键值对列表
    Tags: []model.TagModel{
        {
            TagKey:   "tagK",
            TagValue: "tagV",
        },
    },
}
if result, err := client.CreateVPC(args); err != nil {
    fmt.Println("create vpc failed: ", err)
    return
} 

fmt.Println("create vpc success, vpc id: ", result.VPCID)
```

> 注意: 对请求参数的内容解释如下
> - Name: 表示VPC名称,不能取值"default",长度不超过65个字符，可由数字，字符，下划线组成;
> - ClientToken: 表示幂等性Token，是一个长度不超过64位的ASCII字符串，详见[ClientToken幂等性](https://cloud.baidu.com/doc/VPC/s/gjwvyu77i/#%E5%B9%82%E7%AD%89%E6%80%A7)
> - Description: VPC描述，不超过200字符
> - Cidr: VPC的cidr
> - Tags: 待创建的标签键值对列表

## 查询VPC列表

使用以下代码查询VPC列表信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.ListVPCArgs{
	// 设置每页包含的最大数量，最大数量通常不超过1000，缺省值为1000
    MaxKeys: 2,
    // 设置批量获取列表的查询的起始位置，是一个由系统生成的字符串
    Marker:  marker,
    // 设置是否为默认VPC，可选值：true、false；当不填写此参数时返回所有VPC
    IsDefault: "false",
}
result, err := client.ListVPC(listArgs)
if err != nil {
	fmt.Println("list vpc error: ", err)
    return
}

// 返回标记查询的起始位置
fmt.Println("vpc list marker: ", result.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("vpc list isTruncated: ", result.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("vpc list nextMarker: ", result.NextMarker)
// 每页包含的最大数量
fmt.Println("vpc list maxKeys: ", result.MaxKeys)
// 获取vpc的具体信息
for _, v := range result.VPCs {
    fmt.Println("vpc id: ", v.VPCID)
    fmt.Println("vpc name: ", v.Name)
    fmt.Println("vpc cidr: ", v.Cidr)
    fmt.Println("vpc description: ", v.Description)
    fmt.Println("vpc isDefault: ", v.IsDefault)
    fmt.Println("vpc secondaryCidr: ", v.SecondaryCidr)
    fmt.Println("vpc tags: ", v.Tags)
}
```

## 查询指定VPC

根据特定的VPC ID可以查看相关VPC的详情信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

result, err := client.GetVPCDetail(vpcId)
if err != nil {
	fmt.Println("get vpc detail error: ", err)
    return 
}

// 查询得到vpc的id
fmt.Println("VPC id: ", result.VPC.VPCId)
// 查询得到vpc的名称
fmt.Println("VPC name: ", result.VPC.Name)
// 查询得到vpc的网段及子网掩码
fmt.Println("VPC cidr: ", result.VPC.Cidr)
// 查询得到vpc的描述
fmt.Println("VPC description: ", result.VPC.Description)
// 查询得到是否为默认vpc
fmt.Println("VPC isDefault: ", result.VPC.IsDefault)
// 查询得到vpc中包含的子网
fmt.Println("VPC subnets: ", result.VPC.Subnets)
// 查询得到vpc的辅助网段cidr列表
fmt.Println("VPC secondaryCidr: ", result.VPC.SecondaryCidr)
```

查询得到的VPC详情信息包括名称、网段、描述等信息。

## 删除VPC

使用以下代码可以删除特定的VPC。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.DeleteVPC(vpcId, clientToken); err != nil {
    fmt.Println("delete vpc error: ", err)
    return 
}

fmt.Printf("delete vpc %s success.", vpcId)
```

> 注意: 参数中的clientToken表示幂等性Token，是一个长度不超过64位的ASCII字符串，详见[ClientToken幂等性](https://cloud.baidu.com/doc/VPC/s/gjwvyu77i/#%E5%B9%82%E7%AD%89%E6%80%A7)

## 更新VPC

使用以下代码可以更新指定VPC的名称和描述信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.UpdateVPCArgs{
	// 设置vpc的新名称
    Name:        "TestVPCUpdate",
    // 设置vpc的新备注
    Description: "Test VPC description update",
    // 设置幂等性Token
    ClientToken: clientToken,
}
if err := client.UpdateVPC(vpcId, args); err != nil {
    fmt.Println("update vpc error: ", err)
    return
}

fmt.Printf("update vpc %s success.", vpcId)
```

> 注意: 更新VPC时，对name和description字段的规范要求参考`创建VPC`一节。


## 查询VPC内内网Ip的信息
使用以下代码可以更新指定VPC的名称和描述信息。
>PrivateIpRange的格式为"192.168.0.1-192.168.0-5"
 参数中PrivateIpAddresses或PrivateIpRange的ip数量大小不能超过100.
 若PrivateIpAddresses和PrivateIpRange同时存在，PrivateIpRange优先。

```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &GetVpcPrivateIpArgs{
		VpcId:              "vpc-2pa2x0bjt26i",
		PrivateIpAddresses: []string{"192.168.0.1,192.168.0.2"},
		PrivateIpRange:     "192.168.0.0-192.168.0.45",
	}

result, err := client.GetVPCDetail(vpcId)
if err != nil {
	fmt.Println("get vpc privateIp address info error: ", err)
    return 
}


fmt.Println("privateIpAddresses size is : ", len(result.VpcPrivateIpAddresses))

```
# 子网管理

子网是 VPC 内的用户可定义的IP地址范围，根据业务需求，通过CIDR(无类域间路由)可以指定不同的地址空间和IP段。未来用户可以将子网作为一个单位，用来定义Internet访问权限、路由规则和安全策略。

## 创建子网

通过以下代码可以在指定VPC中创建子网。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.CreateSubnetArgs{
	// 设置子网的名称
    Name:        "TestSDK-Subnet",
    // 设置子网的可用区名称
    ZoneName:    "cn-bj-a",
    // 设置子网的cidr
    Cidr:        "192.168.1.0/24",
    // 设置子网所属vpc的id
    VpcId:       "vpc-4njbqurm0uag",
    // 设置子网的类型，包括“BCC”、“BCC_NAT”、“BBC”三种
    SubnetType:  vpc.SUBNET_TYPE_BCC,
    // 设置子网的描述
    Description: "test subnet",
    // 设置子网的标签键值对列表
    Tags: []model.TagModel{
        {
            TagKey:   "tagK",
            TagValue: "tagV",
        },
    },
}
result, err := client.CreateSubnet(args)
if err != nil {
    fmt.Println("create subnet error: ", err)
    return
}

fmt.Println("create subnet success, subnet id: ", result.SubnetId)
```

> 注意:
> - 子网名称,不能取值"default",长度不超过65个字符，可由数字，字符，下划线组成
> - 可用区名称, 其查询方式参考[查询可用区列表](https://cloud.baidu.com/doc/BCC/s/ijwvyo9im/)

## 查询子网列表

使用以下代码可以查询符合条件的子网列表。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.ListSubnetArgs{
	// 设置批量获取列表的查询的起始位置，是一个由系统生成的字符串
	Marker:     marker,
	// 设置每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
	MaxKeys:    maxKeys,
	// 设置所属vpc的id
    VpcId:      vpcId,
    // 设置所属可用区的名称
    ZoneName:   zoneName,
    // 设置子网类型
    SubnetType: vpc.SUBNET_TYPE_BCC,
}
result, err := client.ListSubnets(args)
if err != nil {
    fmt.Println("list subnets error: ", err)
    return 
}

// 返回标记查询的起始位置
fmt.Println("subnet list marker: ", result.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("subnet list isTruncated: ", result.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("subnet list nextMarker: ", result.NextMarker)
// 每页包含的最大数量
fmt.Println("subnet list maxKeys: ", result.MaxKeys)
// 获取subnet的具体信息
for _, sub := range result.Subnets {
    fmt.Println("subnet id: ", sub.SubnetId)
    fmt.Println("subnet name: ", sub.Name)
    fmt.Println("subnet zoneName: ", sub.ZoneName)
    fmt.Println("subnet cidr: ", sub.Cidr)
    fmt.Println("subnet vpcId: ", sub.VPCId)
    fmt.Println("subnet subnetType: ", sub.SubnetType)
    fmt.Println("subnet description: ", sub.Description)
    fmt.Println("subnet availableIp: ", sub.AvailableIp)
    fmt.Println("subnet tags: ", sub.Tags)
}
```

根据该API，可以根据vpcId、zoneName、subnetType等条件查询符合要求的子网列表。

## 查询指定子网

根据以下代码可以查询指定子网的详细信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

result, err := client.GetSubnetDetail(subnetId)
if err != nil {
    fmt.Println("get subnet detail error: ", err)
    return 
}

// 查询得到子网的id
fmt.Println("subnet id: ", result.Subnet.SubnetId)
// 查询得到子网的名称
fmt.Println("subnet name: ", result.Subnet.Name)
// 查询得到子网所属可用区的名称
fmt.Println("subnet zoneName: ", result.Subnet.ZoneName)
// 查询得到子网的cidr
fmt.Println("subnet cidr: ", result.Subnet.Cidr)
// 查询得到子网所属vpc的id
fmt.Println("subnet vpcId: ", result.Subnet.VPCId)
// 查询得到子网的类型
fmt.Println("subnet subnetType: ", result.Subnet.SubnetType)
// 查询得到子网的描述
fmt.Println("subnet description: ", result.Subnet.Description)
// 查询得到子网内可用ip数
fmt.Println("subnet availableIp: ", result.Subnet.AvailableIp)
// 查询得到子网绑定的标签列表
fmt.Println("subnet tags: ", result.Subnet.Tags)
```

通过该接口可以得到子网的名称、可用区、cidr、类型、描述、可用ip数、标签列表等信息。

> 注意: 子网类型包括"BCC”、"BCC_NAT”、”BBC”三种。

## 删除子网

通过以下代码可以删除指定子网。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.DeleteSubnet(subnetId, clientToken); err != nil {
    fmt.Println("delete subnet error: ", err)
    return
}

fmt.Printf("delete subnet %s success.", subnetId)
```

> 注意: 参数中的clientToken表示幂等性Token，是一个长度不超过64位的ASCII字符串，详见[ClientToken幂等性](https://cloud.baidu.com/doc/VPC/s/gjwvyu77i/#%E5%B9%82%E7%AD%89%E6%80%A7)

## 更新子网

使用以下代码可以更新子网信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.UpdateSubnetArgs{
	// 设置更新操作使用的幂等性token
    ClientToken: clientToken,
    // 设置更新后的子网名称
    Name:        "TestSDK-Subnet-update",
    // 设置更新后的子网描述
    Description: "subnet update",
}
if err := client.UpdateSubnet(subnetId, args); err != nil {
    fmt.Println("update subnet error: ", err)
    return
}

fmt.Printf("update subnet %s success.", subnetId)
```

使用该接口可以实现对子网名称和描述信息的更新操作。

# 路由表管理

路由表是指路由器上管理路由条目的列表。

## 查询所有路由表

使用以下代码可以完成对路由表的查询。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

// 方式1: 通过路由表id进行查询
result, err := client.GetRouteTableDetail(routeTableId, "")
if err != nil {
    fmt.Println("get route table error: ", err)
    return 
}

// 方式2: 通过vpc id进行查询
result, err := client.GetRouteTableDetail("", vpcId)
if err != nil {
    fmt.Println("get route table error: ", err)
    return 
}

// 查询得到路由表id
fmt.Println("result of route table id: ", result.RouteTableId)
// 查询得到vpc id
fmt.Println("result of vpc id: ", result.VpcId)
// 查询得到所有的路由规则列表
for _, route := range result.RouteRules {
    fmt.Println("route rule id: ", route.RouteRuleId)
    fmt.Println("route rule routeTableId: ", route.RouteTableId)
    fmt.Println("route rule sourceAddress: ", route.SourceAddress)
    fmt.Println("route rule destinationAddress: ", route.DestinationAddress)
    fmt.Println("route rule nexthopId: ", route.NexthopId)
    fmt.Println("route rule nexthopType: ", route.NexthopType)
    fmt.Println("route rule description: ", route.Description)
}
```

> 注意:
> - 请求参数routeTableId和vpcId不可以同时为空
> - 使用该接口可以查询得到所有相关的路由规则列表

## 创建路由规则

使用以下代码可以创建路由规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.CreateRouteRuleArgs{
	// 设置路由表id，必选
    RouteTableId:       RouteTableID,
    // 设置源网段，必选
    SourceAddress:      "192.168.1.0/24",
    // 设置目标网段，必选
    DestinationAddress: "172.17.0.0/16",
    // 设置下一跳类型，必选
    NexthopType:        vpc.NEXTHOP_TYPE_NAT,
    // 设置下一跳id，必选
    NexthopId:          NatID,
    // 设置路由规则的描述信息，可选
    Description:        "test route rule",
}
result, err := client.CreateRouteRule(args)
if err != nil {
    fmt.Println("create route rule error: ", err)
    return
}
fmt.Println("create route rule success, route rule id: ", result.RouteRuleId)
```

创建路由表规则，有以下几点需要注意：
- 源网段选择自定义时，自定义网段需在已有子网范围内,0.0.0.0/0除外；
- 目标网段不能与当前所在VPC cidr重叠（目标网段或本VPC cidr为0.0.0.0/0时例外）；
- 新增路由条目的源网段和目标网段，不能与路由表中已有条目源网段和目标网段完全一致。
- 针对下一跳的类型，目前支持三种: Bcc类型是"custom"；VPN类型是"vpn"；NAT类型是"nat"

## 删除路由规则

使用以下代码可以删除特定的路由规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.DeleteRouteRule(routeRuleId, clientToken); err != nil {
    fmt.Println("delete route rule error: ", err)
    return 
}

fmt.Printf("delete route rule %s success.", routeRuleId)
```

> 注意: 参数中的clientToken表示幂等性Token，是一个长度不超过64位的ASCII字符串，详见[ClientToken幂等性](https://cloud.baidu.com/doc/VPC/s/gjwvyu77i/#%E5%B9%82%E7%AD%89%E6%80%A7)


# ACL管理

访问控制列表（Access Control List，ACL）作为应用在子网上的防火墙组件帮助用户实现子网级别的安全访问控制。

## 查询ACL

使用以下代码可以完成acl信息的查询。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

result, err := client.ListAclEntrys(vpcId)
if err != nil {
    fmt.Println("list acl entrys error: ", err)
    return 
}

// 查询得到acl所属的vpc id
fmt.Println("acl entrys of vpcId: ", result.VpcId)
// 查询得到acl所属的vpc名称
fmt.Println("acl entrys of vpcName: ", result.VpcName)
// 查询得到acl所属的vpc网段
fmt.Println("acl entrys of vpcCidr: ", result.VpcCidr)
// 查询得到acl的详细信息
for _, acl := range result.AclEntrys {
    fmt.Println("subnetId: ", acl.SubnetId)
    fmt.Println("subnetName: ", acl.SubnetName)
    fmt.Println("subnetCidr: ", acl.SubnetCidr)
    fmt.Println("aclRules: ", acl.AclRules)
}
```

根据该接口得到的AclEntry列表，包括subnetId、subnetName、subnetCidr、aclRules。

## 添加ACL规则

根据以下代码可以创建acl规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

requests := []vpc.AclRuleRequest{
    {
    	// 设置acl规则所属的子网id
        SubnetId:             "sbn-e4cg8e8zkizs",
        // 设置acl规则的协议
        Protocol:             vpc.ACL_RULE_PROTOCOL_TCP,
        // 设置acl规则的源ip
        SourceIpAddress:      "192.168.2.0",
        // 设置acl规则的目的ip
        DestinationIpAddress: "192.168.0.0/24",
        // 设置acl规则的源端口
        SourcePort:           "8888",
        // 设置acl规则的目的端口
        DestinationPort:      "9999",
        // 设置acl规则的优先级
        Position:             12,
        // 设置acl规则的方向
        Direction:            vpc.ACL_RULE_DIRECTION_INGRESS,
        // 设置acl规则的策略
        Action:               vpc.ACL_RULE_ACTION_ALLOW,
        // 设置acl规则的描述信息
        Description:          "test",
    },
}
args := &vpc.CreateAclRuleArgs{
    AclRules: requests,
}

if err := client.CreateAclRule(args); err != nil {
    fmt.Println("create acl rule error: ", err)
    return
}

fmt.Println("create acl rule success.")
```

使用该接口可以一次创建多条acl规则，对规则参数中的注意事项描述如下:
- protocol: 支持的协议包括all tcp udp icmp
- sourcePort: 源端口，例如1-65535，或8080
- destinationPort: 目的端口，例如1-65535，或8080
- position: 优先级 1-5000且不能与已有条目重复。数值越小，优先级越高，规则匹配顺序为按优先级由高到低匹配
- direction: 规则的入站ingress, 规则的出站egress
- action: 支持的策略包括allow和deny

## 查询ACL规则

使用以下代码可以查询acl规则信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.ListAclRulesArgs{
	// 设置acl所属子网的id
    SubnetId: subnetId,
    // 设置批量获取列表的查询的起始位置
    Marker: marker,
    // 设置每页包含的最大数量
    MaxKeys: maxKeys,
}

result, err := client.ListAclRules(args)
if err != nil {
    fmt.Println("list acl rules error: ", err)
    return
}

// 返回标记查询的起始位置
fmt.Println("acl list marker: ", result.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("acl list isTruncated: ", result.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("acl list nextMarker: ", result.NextMarker)
// 每页包含的最大数量
fmt.Println("acl list maxKeys: ", result.MaxKeys)
// 获取acl的列表信息
for _, acl := range result.AclRules {
    fmt.Println("acl rule id: ", acl.Id)
    fmt.Println("acl rule subnetId: ", acl.SubnetId)
    fmt.Println("acl rule description: ", acl.Description)
    fmt.Println("acl rule protocol: ", acl.Protocol)
    fmt.Println("acl rule sourceIpAddress: ", acl.SourceIpAddress)
    fmt.Println("acl rule destinationIpAddress: ", acl.DestinationIpAddress)
    fmt.Println("acl rule sourcePort: ", acl.SourcePort)
    fmt.Println("acl rule destinationPort: ", acl.DestinationPort)
    fmt.Println("acl rule position: ", acl.Position)
    fmt.Println("acl rule direction: ", acl.Direction)
    fmt.Println("acl rule action: ", acl.Action)
}
```

> 注意: 
> - 使用该接口时，必需提供subnetId参数，以获取特定子网的acl规则列表信息。
> - 系统为用户创建了2条默认ACL规则(无id)。其中入站和出站各一条，规则内容均为全入全出。默认规则,不支持更改和删除。

## 更新ACL规则

使用以下代码可以实现对特定acl规则的更新操作。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.UpdateAclRuleArgs{
	// 设置acl的最新协议
    Protocol:             vpc.ACL_RULE_PROTOCOL_TCP,
    // 设置acl的源ip
    SourceIpAddress:      "192.168.2.0",
    // 设置acl的目的ip
    DestinationIpAddress: "192.168.0.0/24",
    // 设置acl的源端口
    SourcePort:           "3333",
    // 设置acl的目的端口
    DestinationPort:      "4444",
    // 设置acl的优先级
    Position:             12,
    // 设置acl的策略
    Action:               vpc.ACL_RULE_ACTION_ALLOW,
    // 设置acl最新的描述信息
    Description:          "test",
}

if err := client.UpdateAclRule(aclRuleId, args); err != nil {
    fmt.Println("update acl rule error: ", err)
    return 
}

fmt.Printf("update acl rule %s success.", aclRuleId)
```

以上接口可用于对acl规则各个字段的更新过程。

## 删除ACL规则

使用以下代码可以删除指定的acl规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.DeleteAclRule(aclRuleId, clientToken); err != nil {
    fmt.Println("delete acl rule error: ", err)
    return 
}

fmt.Printf("delete acl rule %s success.", aclRuleId)
```

> 注意: 参数中的clientToken表示幂等性Token，是一个长度不超过64位的ASCII字符串，详见[ClientToken幂等性](https://cloud.baidu.com/doc/VPC/s/gjwvyu77i/#%E5%B9%82%E7%AD%89%E6%80%A7)


# NAT网关管理

NAT（Network Address Translation）网关为私有网络提供访问Internet服务，支持SNAT和DNAT，可以使多台云服务器共享公网IP资源访问Internet，也可以使云服务器能够提供Internet服务。NAT网关可以绑定EIP实例及共享带宽，为云服务器实现从内网IP到公网IP的多对一或多对多的地址转换服务。

## 创建NAT网关

使用以下代码可以创建nat网关。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.CreateNatGatewayArgs{
	// 设置nat网关的名称
    Name:  name,
    // 设置nat网关所属的vpc id
    VpcId: vpcId,
    // 设置nat网关的规格
    Spec:  vpc.NAT_GATEWAY_SPEC_SMALL,
    // 设置nat网关的eip列表
    Eips:  []string{eip},
    // 设置nat网关的计费信息
    Billing: &vpc.Billing{
        PaymentTiming: vpc.PAYMENT_TIMING_POSTPAID,
    },
}
result, err := client.CreateNatGateway(args)
if err != nil {
    fmt.Println("create nat gateway error: ", err)
    return 
}

fmt.Println("create nat gateway success, nat gateway id: ", result.NatId)
```

> 注意: 创建过程中，应注意以下事项:
> - NAT网关的名称，由大小写字母、数字以及-_ /.特殊字符组成，必须以字母开头，长度1-65
> - NAT网关的大小，有small(最多支持绑定5个公网IP)、medium(最多支持绑定10个公网IP)、large(最多支持绑定15个公网IP)三种
> - NAT网关可以关联一个公网EIP或者共享带宽中的一个或多个EIP
> - 付款方式支持预支付（Prepaid）和后支付（Postpaid）两种，预支付当前仅支持按月，时长取值范围: [1,2,3,4,5,6,7,8,9,12,24,36]

## 查询NAT网关列表

使用以下代码可以查询符合条件的nat网关列表。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.ListNatGatewayArgs{
	// 设置nat网关所属的vpc id，必选
    VpcId: vpcId,
    // 指定查询的NAT的Id
    NatId: natId,
    // 指定查询的NAT的名称
    Name: name,
    // 指定查询的NAT绑定的EIP
    Ip: ip,
    // 设置nat网关批量获取列表的查询的起始位置
    Marker: marker,
    // 设置nat网关每页包含的最大数量，最大数量不超过1000。缺省值为1000
    MaxKeys: maxKeys,
}
result, err := client.ListNatGateway(args)
if err != nil {
    fmt.Println("list nat gateway error: ", err)
    return 
}

// 返回标记查询的起始位置
fmt.Println("nat list marker: ", result.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("nat list isTruncated: ", result.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("nat list nextMarker: ", result.NextMarker)
// 每页包含的最大数量
fmt.Println("nat list maxKeys: ", result.MaxKeys)
// 获取nat的列表信息
for _, nat := range result.Nats {
    fmt.Println("nat id: ", nat.Id)
    fmt.Println("nat name: ", nat.Name)
    fmt.Println("nat vpcId: ", nat.VpcId)
    fmt.Println("nat spec: ", nat.Spec)
    fmt.Println("nat eips: ", nat.Eips)
    fmt.Println("nat status: ", nat.Status)
    fmt.Println("nat paymentTiming: ", nat.PaymentTiming)
    fmt.Println("nat expireTime: ", nat.ExpiredTime)
}
```

> 注意: 
> - 可根据NAT网关ID、NAT网关的name、NAT网关绑定的EIP来查询。
> - 若不提供查询条件，则默认查询覆盖所有NAT网关
> - vpcId为必选参数

## 查询NAT网关详情

使用以下代码可以查询特定nat网关的详细信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

result, err := client.GetNatGatewayDetail(natId)
if err != nil {
    fmt.Println("get nat gateway details error: ", err)
    return 
}

// 查询得到nat网关的id
fmt.Println("nat id: ", result.Id)
// 查询得到nat网关的名称
fmt.Println("nat name: ", result.Name)
// 查询得到nat网关所属的vpc id
fmt.Println("nat vpcId: ", result.VpcId)
// 查询得到nat网关的大小
fmt.Println("nat spec: ", result.Spec)
// 查询得到nat网关绑定的EIP的IP地址列表
fmt.Println("nat eips: ", result.Eips)
// 查询得到nat网关的状态
fmt.Println("nat status: ", result.Status)
// 查询得到nat网关的付费方式
fmt.Println("nat paymentTiming: ", result.PaymentTiming)
// 查询得到nat网关的过期时间
fmt.Println("nat expireTime: ", result.ExpiredTime)
```

## 更新NAT网关名称

使用以下代码可以对nat网关的名称进行更改。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.UpdateNatGatewayArgs{
	// 设置nat网关的最新名称
    Name: "TestNatUpdate",
}

if err := client.UpdateNatGateway(natId, args); err != nil {
    fmt.Println("update nat gateway error: ", err)
    return 
}

fmt.Printf("update nat gateway %s success.", natId)
```

> 注意: 目前该接口仅支持对网关名称属性的更改。

## 绑定EIP

使用以下代码可以为nat网关绑定eip。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.BindEipsArgs{
	// 设置绑定的EIP ID列表
    Eips: []string{eip},
}
if err := client.BindEips(natId, args); err != nil {
    fmt.Println("bind eips error: ", err)
    return 
}

fmt.Println("bind eips success.")
```

注意: 
- 若该NAT已经绑定EIP，必须解绑后才可绑定。
- 若该NAT已经绑定共享带宽，可以继续绑定该共享带宽中的其他IP。

## 解绑EIP

使用以下代码可以为nat网关解绑eip。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.UnBindEipsArgs{
	// 设置解绑的EIP ID列表
    Eips: []string{eip},
}
if err := client.UnBindEips(natId, args); err != nil {
    fmt.Println("unbind eips error: ", err)
    return 
}

fmt.Println("unbind eips success.")
```

## 绑定DNAT EIP

使用以下代码可以为nat网关绑定DNAT EIP。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.BindDnatEipsArgs{
	// 设置绑定的DNAT EIP ID列表
    DnatEips: []string{dnatEips},
}
if err := client.BindDnatEips(natId, args); err != nil {
    fmt.Println("bind DNAT Eips error: ", err)
    return 
}

fmt.Println("bind DNAT Eips success.")
```

注意:
- 若该NAT DNAT已经绑定EIP，必须解绑后才可绑定。
- 若该NAT DNAT已经绑定共享带宽，可以继续绑定该共享带宽中的其他IP。

## 解绑DNAT EIP

使用以下代码可以为nat网关解绑DNAT EIP。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.UnBindDnatEipsArgs{
	// 设置解绑的DNAT EIP ID列表
    DnatEips: []string{dnatEips},
}
if err := client.UnBindDnatEips(natId, args); err != nil {
    fmt.Println("unbind DNAT Eips error: ", err)
    return 
}

fmt.Println("unbind DNAT Eips success.")
```

## 释放NAT网关

使用以下代码释放特定的nat网关。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.DeleteNatGateway(natId, clientToken); err != nil {
    fmt.Println("delete nat gateway error: ", err)
    return 
}

fmt.Printf("delete nat gateway %s success.", natId)
```

> 注意: 预付费未到期的NAT网关不能释放。

## NAT网关续费

使用以下接口完成nat网关的续费操作，延长过期时间。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.RenewNatGatewayArgs{
	// 设置nat网关续费的订单信息
    Billing: &vpc.Billing{
        Reservation: &vpc.Reservation{
            ReservationLength:   1,
            ReservationTimeUnit: "month",
        },
    },
}
if err := client.RenewNatGateway(natId, args); err != nil {
    fmt.Println("renew nat gateway error: ", err)
    return 
}

fmt.Printf("renew nat gateway %s success.", natId)
```

> 注意:
- 后付费的NAT网关不能续费

## 创建SNAT规则
使用以下代码可以创建nat网关的snat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &CreateNatGatewaySnatRuleArgs{
RuleName:         "sdk-test",
PublicIpsAddress: []string{"100.88.10.84"},
SourceCIDR:       "192.168.3.3",
}
result, err := VPC_CLIENT.CreateNatGatewaySnatRule("nat-b1jb3b5e34tc", args)
ExpectEqual(t.Errorf, nil, err)
r, err := json.Marshal(result)
fmt.Println(string(r))
```

## 删除SNAT规则
使用以下代码可以删除nat网关的snat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

VPC_CLIENT.DeleteNatGatewaySnatRule("nat-b1jb3b5e34tc", "rule-hprz7sv9zvcx", getClientToken())
```

## 修改SNAT规则
使用以下代码可以修改nat网关的snat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &UpdateNatGatewaySnatRuleArgs{
RuleName:   "sdk-test-1",
SourceCIDR: "192.168.3.6",
}
VPC_CLIENT.UpdateNatGatewaySnatRule("nat-b1jb3b5e34tc", "rule-hprz7sv9zvcx", args)
```

## 查询SNAT规则
使用以下代码可以查询nat网关的snat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &ListNatGatewaySnatRuleArgs{
NatId: "nat-b1jb3b5e34tc",
}
result, err := VPC_CLIENT.ListNatGatewaySnatRules(args)
ExpectEqual(t.Errorf, nil, err)
r, err := json.Marshal(result)
fmt.Println(string(r))
```

## 创建DNAT规则
使用以下代码可以创建nat网关的dnat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &ListNatGatewaySnatRuleArgs{
NatId: "nat-b1jb3b5e34tc",
}
result, err := VPC_CLIENT.ListNatGatewaySnatRules(args)
ExpectEqual(t.Errorf, nil, err)
r, err := json.Marshal(result)
fmt.Println(string(r))
```

## 删除DNAT规则
使用以下代码可以删除nat网关的dnat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

VPC_CLIENT.DeleteNatGatewayDnatRule("nat-b1jb3b5e34tc", "rule-8gee5abqins0", getClientToken())
```

## 修改DNAT规则
使用以下代码可以修改nat网关的dnat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &UpdateNatGatewayDnatRuleArgs{
RuleName:         "sdk-test-3",
PrivateIpAddress: "192.168.1.5",
}
VPC_CLIENT.UpdateNatGatewayDnatRule("nat-b1jb3b5e34tc", "rule-8gee5abqins0", args)
```

## 查询DNAT规则
使用以下代码可以查询nat网关的dnat规则。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &ListNatGatewaDnatRuleArgs{}
result, err := VPC_CLIENT.ListNatGatewayDnatRules("nat-b1jb3b5e34tc", args)
ExpectEqual(t.Errorf, nil, err)
r, err := json.Marshal(result)
fmt.Println(string(r))
```

# 对等连接管理

对等连接（Peer Connection）为用户提供了VPC级别的网络互联服务，使用户实现在不同虚拟网络之间的流量互通，实现同区域/跨区域，同用户/不同用户之间稳定高速的虚拟网络互联。

## 创建对等连接

使用以下代码创建对等连接。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.CreatePeerConnArgs{
	// 设置对等连接的带宽
    BandwidthInMbps: 10,
    // 设置对等连接的描述信息
    Description:     "test peer conn",
    // 设置对等连接的本端端口名称
    LocalIfName:     "local-interface",
    // 设置对等连接的本端vpc的id
    LocalVpcId:      vpcId,
    // 设置对等连接的对端账户ID，只有在建立跨账号的对等连接时需要该字段
    peerAccountId:   peerAccountId,
    // 设置对等连接的对端vpc的id
    PeerVpcId:       peerVpcId,
    // 设置对等连接的对端区域
    PeerRegion:      region,
    // 设置对等连接的对端接口名称，只有本账号的对等连接才允许设置该字段
    PeerIfName:      "peer-interface",
    // 设置对等连接的计费信息
    Billing: &vpc.Billing{
        PaymentTiming: vpc.PAYMENT_TIMING_POSTPAID,
    },
}
result, err := client.CreatePeerConn(args)
if err != nil {
    fmt.Println("create peerconn error: ", err)
    return 
}

fmt.Println("create peerconn success, peerconn id: ", result.PeerConnId)
```

> 注意: 
> - 对于本端区域和对端区域相同的对等连接，只支持后付费。
> - 跨账号的对等连接，必须接受端接受后对等连接才可用。
> - 对于同账号的对等连接，系统会触发对端自动接受。
> - 任意两个VPC之间最多只能存在一条对等连接。
> - 发起端和接收端的VPC不能是同一个。
> - 如果本端vpc和对端vpc均为中继vpc,则不可以建立对等连接。

## 查询对等连接列表

使用以下代码可以查询对等连接的列表信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.ListPeerConnsArgs{
	// 指定对等连接所属的vpc id
	VpcId: vpcId,
	// 指定批量获取列表的查询的起始位置
	Marker: marker,
	// 指定每页包含的最大数量，最大数量不超过1000。缺省值为1000
	MaxKeys: maxKeys,
}
result, err := client.ListPeerConn(args)
if err != nil {
    fmt.Println("list peer conns error: ", err)
    return 
}

// 返回标记查询的起始位置
fmt.Println("peerconn list marker: ", result.Marker)
// true表示后面还有数据，false表示已经是最后一页
fmt.Println("peerconn list isTruncated: ", result.IsTruncated)
// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
fmt.Println("peerconn list nextMarker: ", result.NextMarker)
// 每页包含的最大数量
fmt.Println("peerconn list maxKeys: ", result.MaxKeys)
// 获取对等连接的列表信息
for _, pc := range result.PeerConns {
    fmt.Println("peerconn id: ", pc.PeerConnId)
    fmt.Println("peerconn role: ", pc.Role)
    fmt.Println("peerconn status: ", pc.Status)
    fmt.Println("peerconn bandwithInMbp: ", pc.BandwidthInMbps)
    fmt.Println("peerconn description: ", pc.Description)
    fmt.Println("peerconn localIfId: ", pc.LocalIfId)
    fmt.Println("peerconn localIfName: ", pc.LocalIfName)
    fmt.Println("peerconn localVpcId: ", pc.LocalVpcId)
    fmt.Println("peerconn localRegion: ", pc.LocalRegion)
    fmt.Println("peerconn peerVpcId: ", pc.PeerVpcId)
    fmt.Println("peerconn peerRegion: ", pc.PeerRegion)
    fmt.Println("peerconn peerAccountId: ", pc.PeerAccountId)
    fmt.Println("peerconn paymentTiming: ", pc.PaymentTiming)
    fmt.Println("peerconn dnsStatus: ", pc.DnsStatus)
    fmt.Println("peerconn createdTime: ", pc.CreatedTime)
    fmt.Println("peerconn expiredTime: ", pc.ExpiredTime)
}
```

使用该接口可以查询得到所有符合条件的对等连接信息，其中，vpcId是可选参数。

## 查看对等连接详情

通过以下代码可以查询特定对等连接的详细信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

result, err := client.GetPeerConnDetail(peerConnId, vpc.PEERCONN_ROLE_INITIATOR)
if err != nil {
    fmt.Println("get peer conn detail error: ", err)
    return 
}

// 查询得到对等连接的id
fmt.Println("peerconn id: ", result.PeerConnId)
// 查询得到对等连接的角色, "initiator"表示发起端"acceptor"表示接受端
fmt.Println("peerconn role: ", result.Role)
// 查询得到对等连接的状态
fmt.Println("peerconn status: ", result.Status)
// 查询得到对等连接的带宽
fmt.Println("peerconn bandwithInMbp: ", result.BandwidthInMbps)
// 查询得到对等连接的描述
fmt.Println("peerconn description: ", result.Description)
// 查询得到对等连接的本端接口ID
fmt.Println("peerconn localIfId: ", result.LocalIfId)
// 查询得到对等连接的本端接口名称
fmt.Println("peerconn localIfName: ", result.LocalIfName)
// 查询得到对等连接的本端VPC ID
fmt.Println("peerconn localVpcId: ", result.LocalVpcId)
// 查询得到对等连接的本端区域
fmt.Println("peerconn localRegion: ", result.LocalRegion)
// 查询得到对等连接的对端VPC ID
fmt.Println("peerconn peerVpcId: ", result.PeerVpcId)
// 查询得到对等连接的对端区域
fmt.Println("peerconn peerRegion: ", result.PeerRegion)
// 查询得到对等连接的对端账户ID
fmt.Println("peerconn peerAccountId: ", result.PeerAccountId)
// 查询得到对等连接的计费方式
fmt.Println("peerconn paymentTiming: ", result.PaymentTiming)
// 查询得到对等连接的dns状态
fmt.Println("peerconn dnsStatus: ", result.DnsStatus)
// 查询得到对等连接的创建时间
fmt.Println("peerconn createdTime: ", result.CreatedTime)
// 查询得到对等连接的过期时间
fmt.Println("peerconn expiredTime: ", result.ExpiredTime)
```

> 注意: "initiator"表示发起端"acceptor"表示接受端，同region的对等连接可以据此进行详情查询，若不设置该参数，同region则随机返回一端信息。

## 更新对等连接本端接口名称和备注

使用以下代码可以更新对等连接本端接口名称和备注。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.UpdatePeerConnArgs{
	// 设置对等连接的接口ID 不可更改，必选
    LocalIfId:   localIfId,
    // 设置对等连接的本端端口名称
    LocalIfName: "test-update",
    // 设置对等连接的本端端口描述
    Description: "test-description",
}
if err := client.UpdatePeerConn(peerConnId, args); err != nil {
    fmt.Println("update peer conn error: ", err)
    return 
}

fmt.Printf("update peer conn %s success", peerConnId)
```

## 接受对等连接申请

使用以下代码可以接受对等连接的申请信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.AcceptPeerConnApply(peerConnId, clientToken); err != nil {
    fmt.Println("accept peer conn error: ", err)
    return 
}

fmt.Printf("accept peer conn %s success.", peerConnId)
```

> 注意: 
> - 发起端发出的连接请求超时时间为7天，超时后发起端对等连接的状态为协商失败。
> - 接收端拒绝后，发起端对等连接状态为协商失败。

## 拒绝对等连接申请

使用以下代码可以接受对等连接的申请信息。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.RejectPeerConnApply(peerConnId, clientToken); err != nil {
    fmt.Println("reject peer conn error: ", err)
    return 
}

fmt.Printf("reject peer conn %s success.", peerConnId)
```

## 释放对等连接

使用以下代码可以释放特定的对等连接。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

if err := client.DeletePeerConn(peerConnId, clientToken); err != nil {
    fmt.Println("delete peer conn error: ", err)
    return 
}

fmt.Printf("delete peer conn %s success", peerConnId)
```

> 注意: 
> - 跨账号只有发起端可以释放。
> - 预付费可用且未到期的对等连接不能释放。
> - 预付费协商失败的可以释放。

## 对等连接带宽升降级

使用以下代码可以为指定的对等连接进行带宽升级操作。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.ResizePeerConnArgs{
	// 指定对等连接升降级的带宽
    NewBandwidthInMbps: 20,
}

if err := client.ResizePeerConn(peerConnId, args); err != nil {
    fmt.Println("resize peer conn error: ", err)
    return 
}

fmt.Printf("resize peer conn %s success.", peerConnId)
```

> 注意:
> - 跨账号只有发起端才可以进行带宽的升降级操作。
> - 预付费的对等连接只能进行带宽升级不能降级。
> - 后付费的对等连接可以进行带宽的升级和降级。

## 对等连接续费

使用以下代码可以为对等连接进行续费操作，延长过期时间。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.RenewPeerConnArgs{
	// 指定对等连接的续费信息
    Billing: &vpc.Billing{
        Reservation: &vpc.Reservation{
            ReservationLength:   1,
            ReservationTimeUnit: "month",
        },
    },
}

if err := client.RenewPeerConn(peerConnId, args); err != nil {
    fmt.Println("renew peer conn error: ", err)
    return 
}

fmt.Printf("renew peer conn %s success.", peerConnId)
```

> 注意:
> - 后付费的对等连接不能续费。
> - 跨账号续费操作只能由发起端来操作。

## 开启对等连接同步DNS

使用以下代码可以开启对等连接同步DNS记录。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.PeerConnSyncDNSArgs{
	// 指定对等连接的角色，发起端"initiator" 接收端"acceptor"
    Role: vpc.PEERCONN_ROLE_INITIATOR,
}

if err := client.OpenPeerConnSyncDNS(peerConnId, args); err != nil {
    fmt.Println("open peer conn sync dns error: ", err)
    return 
}

fmt.Printf("open peer conn %s sync dns success.", peerConnId)
```

> 注意:
> - 对等连接的状态为可用的时候才能开启DNS。
> - 对等连接的DNS状态为同步中或同步关闭中不可开启同步DNS。

## 关闭对等连接同步DNS

使用以下代码可以关闭对等连接同步DNS记录。
```go
//import "github.com/baidubce/bce-sdk-go/services/vpc"

args := &vpc.PeerConnSyncDNSArgs{
	// 指定对等连接的角色，发起端"initiator" 接收端"acceptor"
    Role: vpc.PEERCONN_ROLE_INITIATOR,
}

if err := client.ClosePeerConnSyncDNS(peerConnId, args); err != nil {
    fmt.Println("close peer conn sync dns error: ", err)
    return 
}

fmt.Printf("close peer conn %s sync dns success.", peerConnId)
```

> 注意:
> - 对等连接的状态为可用的时候才能关闭DNS。
> - 对等连接的DNS状态为同步中或同步关闭中不可关闭同步DNS。


# 错误处理

GO语言以error类型标识错误，VPC支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | VPC服务返回的错误

用户使用SDK调用VPC相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// vpcClient 为已创建的VPC Client对象
args := &vpc.ListVPCArgs{}
result, err := client.ListVPC(args)
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

客户端异常表示客户端尝试向VPC发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当VPC服务端出现异常时，VPC服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[VPC错误信息格式](https://cloud.baidu.com/doc/VPC/s/sjwvyuhe7)

## SDK日志

VPC GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

VPC GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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
log.Debugf("%s", "logging message using the log package in the VPC go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the VPC go sdk")
```


# 版本变更记录
## v0.9.6 [2020-12-27]
- 增加vpc查询PrivateIpAddress信息接口
## v0.9.5 [2019-09-24]

首次发布:

 - 支持创建VPC、查询VPC列表、查询指定VPC、删除VPC、更新VPC接口;
 - 支持创建子网、查询子网列表、查询指定子网、删除子网、更新子网接口;
 - 支持查询路由表、创建路由规则、删除路由规则接口;
 - 支持查询ACL、添加ACL规则、查询ACL规则、更新ACL规则、删除ACL规则接口;
 - 支持创建NAT网关、查询NAT网关列表、查询NAT网关详情、更新NAT网关名称、绑定EIP、解绑EIP、释放NAT网关、NAT网关续费接口;
 - 支持创建对等连接、查询对等连接列表、查看对等连接详情、更新对等连接本端接口名称和备注、接受对等连接申请、拒绝对等连接申请、释放对等连接、对等连接带宽升降级、对等连接续费、开启对等连接同步DNS、关闭对等连接同步DNS接口。
 