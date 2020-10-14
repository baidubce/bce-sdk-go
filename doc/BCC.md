# BCC服务

# 概述

本文档主要介绍BCC GO SDK的使用。在使用本文档前，您需要先了解BCC的一些基本知识。若您还不了解BCC，可以参考[产品描述](https://cloud.baidu.com/doc/BCC/s/Jjwvymo32)和[入门指南](https://cloud.baidu.com/doc/BCC/s/ojwvymvfe)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BCC访问域名](https://cloud.baidu.com/doc/BCC/s/0jwvyo603)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云BCC，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BCC做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建BCC Client

BCC Client是BCC服务的客户端，为开发者与BCC服务进行交互提供了一系列的方法。

### 使用AK/SK新建BCC Client

通过AK/SK方式访问BCC，用户可以参考如下代码新建一个Bcc Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bcc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BCCClient
	bccClient, err := bcc.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/BCC/s/ojwvynrqn)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BCC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`bcc.bj.baidubce.com`。

### 使用STS创建BCC Client

**申请STS token**

BCC可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问BCC，用户需要先通过STS的client申请一个认证字符串，申请方式可参见[百度云STS使用介绍](https://cloud.baidu.com/doc/IAM/s/gjwvyc7n7)。

**用STS token新建BCC Client**

申请好STS后，可将STS Token配置到BCC Client中，从而实现通过STS Token创建BCC Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建BCC Client对象：

```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/bcc" //导入BCC服务模块
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

	// 使用申请的临时STS创建BCC服务的Client对象，Endpoint使用默认值
	bccClient, err := bcc.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create bcc client failed:", err)
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
	bccClient.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置BCC Client时，无论对应BCC服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

## 配置HTTPS协议访问BCC

BCC支持HTTPS传输协议，您可以通过在创建BCC Client对象时指定的Endpoint中指明HTTPS的方式，在BCC GO SDK中使用HTTPS访问BCC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcc"

ENDPOINT := "https://bcc.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
bccClient, _ := bcc.NewClient(AK, SK, ENDPOINT)
```

## 配置BCC Client

如果用户需要配置BCC Client的一些细节的参数，可以在创建BCC Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问BCC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcc"

//创建BCC Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := bcc.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/bcc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := bcc.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/bcc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bcc.bj.baidubce.com"
client, _ := bcc.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问BCC时，创建的BCC Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建BCC Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# 主要接口

BCC实例是一个虚拟的计算环境，包含CPU、内存等最基础的计算组件，是云服务器呈献给每个用户的实际操作实体。BCC实例是云服务器最为核心的概念，支持IP绑定，镜像和快照等功能，诸如CDS磁盘、SCS简单缓存服务只有挂载在BCC实例后才可使用。

- 每个用户最多可同时拥有20个BCC实例，若需要更多的配额，请发工单申请。

## 实例管理

### 创建实例

使用以下代码可以创建BCC实例，包括专属实例、普通型Ⅰ 型实例、普通型Ⅱ型实例、存储优化型BCC、计算优化型BCC
```go
args := &api.CreateInstanceArgs{
    // 选择实例创建镜像ID 
    ImageId: "m-DpgNg8lO",
	// 选择付款方式，可以选择预付费或后付费 
	Billing: api.Billing{
        PaymentTiming: api.PaymentTimingPostPaid,
	},
    // 选择实例类型，可以选择N1, N2, N3等 
    InstanceType:        api.InstanceTypeN1,
    // 选择1核CPU 
    CpuCount:            1,
    // 选择1GB内存 
    MemoryCapacityInGB:  1,
    // 选择40GB磁盘空间 
    RootDiskSizeInGb:    40,
    // 选择待创建的实例系统盘介质为HP1 
    RootDiskStorageType: api.StorageTypeCloudHP1,
    // 临时盘数据盘大小
    EphemeralDisks        []EphemeralDisk  "ephemeralDisks"
    // 选择创建100GB大小SSD类型CDS磁盘并挂载到实例上 
    CreateCdsList: []api.CreateCdsModel{
		{
			StorageType: api.StorageTypeSSD,
			CdsSizeInGB: 100,
		},
	},
    // 设置管理员密码 
    AdminPass: "123qaz!@#",
	// 设置实例名称 
	Name:      "terraform_sdkTest",
	// 设置创建BCC使用的网络带宽大小
    NetWorkCapacityInMbps int              networkCapacityInMbps
    // 设置需要创建BCC使用的DCC服务器id
    DedicateHostId        string           "dedicatedHostId"
    // 设置待查询的竞价实例的购买个数
    PurchaseCount         int              purchaseCount
    // 设置可用区
    ZoneName              string           "zoneName"
    // 指定子网和安全组创建，要求子网和安全组必须同时指定或同时不指定，
    // 同时指定的子网和安全组必须同属于一个VPC，都不指定会使用默认子网和默认安全组。
    // 设置创建BCC使用的子网
    SubnetId              string           "subnetId"
    // // 设置创建BCC使用的安全组
    SecurityGroupId       string           "securityGroupId"
    // 设置需要创建GPU卡信息
    GpuCard               string           "gpuCard"
    // 设置需要创建FPGA卡信息
    FpgaCard              string           "fpgaCard"
    // 设置GPU卡或FPGA卡数量
    CardCount             string           "cardCount"
    // 设置按月付费或者按年付费 月是"month",年是"year"
    AutoRenewTimeUnit     string           "autoRenewTimeUnit"
    // 设置自动续费的时间 按月是1-9 按年是 1-3
    AutoRenewTime         int              autoRenewTime
    // cds是否自动续费 是:true 否:false
    CdsAutoRenew          bool             cdsAutoRenew
    // 待创建实例指定的标签是否需要和已有标签键进行关联，默认为false。注意值为true时要保证该标签键已存在
    RelationTag           bool             relationTag
    // 待创建的标签列表
    Tags                  []model.TagModel tags
    // 指定实例所在的部署集id
    DeployId              string           "deployId"
    // 设置要绑定的密钥对ID
    KeypairId             string           "keypairId"
    // 设置要绑定的自动快照策略ID
    AspId                 string           "aspId"
    // 公网带宽计费方式，若不指定internetChargeType，默认付费方式同BCC，预付费默认为包年包月按带宽，后付费默认为按使用带宽计费。（BANDWIDTH_PREPAID：预付费按带宽结算；TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费；BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费）
    InternetChargeType    string           "internetChargeType"
    // 设置内网IP
    InternalIps           []string         internalIps
    // 使用 uuid 生成一个长度不超过64位的ASCII字符串
    ClientToken           string           "random-uuid"
    //创建实例支持幂等的token
    RequestToken     string  "requestToken"
}

// 若要生成预付费实例，可以按以下配置生成一个月的预付费实例
args.Billing = api.Billing{
		PaymentTiming: api.PaymentTimingPrePaid,
		Reservation: &api.Reservation{
            ReservationLength:   1,
            ReservationTimeUnit: "month",
        }
	}

// 若要设置自动续费，可以按以下参数设置一年内自动续费
args.AutoRenewTimeUnit = "year"
args.AutoRenewTime = 1

// 若要创建公网EIP，可以设置以下参数
args.networkCapacityInMbps = 1

// 若要一次创建多个相同配置的实例，可以设置以下参数
args.PurchaseCount = 2

result, err := client.CreateInstance(args)
if err != nil {
    fmt.Println("create instance failed:", err)
} else {
    fmt.Println("create instance success: ", result)
}
```

> **提示：**
> 1.  创建BCC请求是一个异步请求，返回200表明订单生成，后续可以通过查询返回的实例id信息了解BCC虚机的创建进度。
> 2.  本接口用于创建一个或多个同配虚拟机实例。
> 3.  创建实例需要实名认证，没有通过实名认证的可以前往百度开放云官网控制台中的安全认证下的实名认证中进行认证。
> 4.  创建计费方式为后付费的实例需要账户现金余额+通用代金券大于100；预付费方式的实例则需要账户现金余额大于等于实例费用。
> 5.  支持批量创建，且如果创建过程中有一个实例创建失败，所有实例将全部回滚，均创建失败，如果创建时包含CDS，CDS也会回滚。
> 6.  缺省情形下，一个实例最多只能挂载5个云磁盘。
> 7.  创建CDS磁盘和临时数据盘时，磁盘容量大小限制为5的倍数。
> 8.  创建实例支持创建和添加临时数据盘，但不支持单独创建或添加临时数据盘。
> 9.  临时数据盘不支持挂载、卸载、删除。
> 10. 普通实例的临时数据盘最大不能超过500G。
> 11. 指定子网和安全组创建，要求子网和安全组必须同时指定或同时不指定，同时指定的子网和安全组必须同属于一个VPC，都不指定会使用默认子网和默认安全组。
> 12. 指定公网IP带宽创建，计费方式为按照带宽计费。
> 13. 创建接口为异步创建，可通过查询实例详情接口查询实例状态
> 14. 可通过该接口指定专属服务器创建实例，专属实例不参与计费。专属实例只能通过ephemeralDisks创建临时盘并指定磁盘类型。
> 15. 每个实例最多只能购买一块临时数据盘。
> 16. 实例的临时数据盘默认只有hp1类型。
> 17. 通过instanceType字段指定需要创建的虚机类型，目前API支持创建的虚机类型参见下述InstanceType。参数(instanceType，cpuCount，memoryCapacityInGB)可以确定需要的机型以及配置。
> 18. 创建存储优化型实例必须购买临时数据盘，通过ephemeralDisks指定临时盘数据盘大小，默认nvme类型数据盘，无需指定。
> 19. 创建请求详细使用请参考BCC API 文档[创建实例](https://cloud.baidu.com/doc/BCC/s/yjwvyoe0s)
> 20. 创建FPGA BCC虚机需要使用指定的(CPU、内存、本地数据盘、FPGA卡类型以及专用镜像), 详细请参考BCC API 文档[FPGA型BCC可选规格配置](https://cloud.baidu.com/doc/BCC/s/6jwvyo0q2#fpga%E5%9E%8Bbcc%E5%8F%AF%E9%80%89%E8%A7%84%E6%A0%BC%E9%85%8D%E7%BD%AE)
> 21. 创建GPU BCC虚机需要使用指定的(CPU、内存、本地数据盘、GPU卡类型), 详细请参考BCC API 文档[GPU型BCC可选规格配置](https://cloud.baidu.com/doc/BCC/s/6jwvyo0q2#gpu%E5%9E%8Bbcc%E5%8F%AF%E9%80%89%E8%A7%84%E6%A0%BC%E9%85%8D%E7%BD%AE)

### 创建实例（通过指定实例套餐规格）
使用以下代码可以创建BCC实例，包括普通型BCC、存储优化型BCC、计算优化型BCC、大数据机型BCC、GPU机型BCC、FPGA机型BCC：

```go
createInstanceBySpecArgs := &api.CreateInstanceBySpecArgs{
    // 选择实例创建镜像ID 
	ImageId:               "m-1PyVLtic",
    // 选择创建BCC的套餐规格
	Spec:                  "bcc.g2.c2m8",
    // 选择40GB磁盘空间 
    RootDiskSizeInGb:      40,
    // 选择待创建的实例系统盘介质为HP1 
    RootDiskStorageType:   api.StorageTypeCloudHP1,
    // 选择创建100GB大小SSD类型CDS磁盘并挂载到实例上 
    CreateCdsList: []api.CreateCdsModel{
		{
			StorageType: api.StorageTypeSSD,
			CdsSizeInGB: 100,
    	},
    },
    // 选择付款方式，可以选择预付费或后付费 
    Billing: api.Billing{
		PaymentTiming: api.PaymentTimingPostPaid,
    },
    Name:                  "sdkTest",
    AdminPass:             "123qaz!@#",
    // 临时盘数据盘大小
    EphemeralDisks        []EphemeralDisk  "ephemeralDisks"
    // 设置创建BCC使用的网络带宽大小
    NetWorkCapacityInMbps int              networkCapacityInMbps
    // 设置待查询的竞价实例的购买个数
    PurchaseCount         int              purchaseCount
    // 设置可用区
    ZoneName              string           "zoneName"
    // 指定子网和安全组创建，要求子网和安全组必须同时指定或同时不指定，
    // 同时指定的子网和安全组必须同属于一个VPC，都不指定会使用默认子网和默认安全组。
    // 设置创建BCC使用的子网
    SubnetId              string           "subnetId"
    // 设置创建BCC使用的安全组
    SecurityGroupId       string           "securityGroupId"
    // 设置按月付费或者按年付费 月是"month",年是"year"
    AutoRenewTimeUnit     string           "autoRenewTimeUnit"
    // 设置自动续费的时间 按月是1-9 按年是 1-3
    AutoRenewTime         int              autoRenewTime
    // cds是否自动续费 是:true 否:false
    CdsAutoRenew          bool             cdsAutoRenew
    // 待创建实例指定的标签是否需要和已有标签键进行关联，默认为false。注意值为true时要保证该标签键已存在
    RelationTag           bool             relationTag
    // 待创建的标签列表
    Tags                  []model.TagModel tags
    // 设置要绑定的密钥对ID
    KeypairId             string           "keypairId"
    // 设置要绑定的自动快照策略ID
    AspId                 string           "aspId"
    // 公网带宽计费方式，若不指定internetChargeType，默认付费方式同BCC，预付费默认为包年包月按带宽，后付费默认为按使用带宽计费。（BANDWIDTH_PREPAID：预付费按带宽结算；TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费；BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费）
    InternetChargeType    string           "internetChargeType"
    // 设置内网IP
    InternalIps           []string         internalIps
    // 使用 uuid 生成一个长度不超过64位的ASCII字符串
    ClientToken           string           "random-uuid"
	// 创建实例支持幂等的token,成功后永久有效
    RequestToken          string           "requestToken"

}
result, err := client.TestCreateInstanceBySpec(args)
if err != nil {
    fmt.Println("create instance failed:", err)
} else {
    fmt.Println("create instance success: ", result)
}
```
> **提示：**
> 1.  创建BCC请求是一个异步请求，返回200表明订单生成，后续可以通过查询返回的实例id信息了解BCC虚机的创建进度。
> 2.  本接口用于创建一个或多个同配虚拟机实例。
> 3.  创建实例需要实名认证，没有通过实名认证的可以前往百度开放云官网控制台中的安全认证下的实名认证中进行认证。
> 4.  创建计费方式为后付费的实例需要账户现金余额+通用代金券大于100；预付费方式的实例则需要账户现金余额大于等于实例费用。
> 5.  支持批量创建，且如果创建过程中有一个实例创建失败，所有实例将全部回滚，均创建失败，如果创建时包含CDS，CDS也会回滚。
> 6.  缺省情形下，一个实例最多只能挂载5个云磁盘。
> 7.  创建CDS磁盘和临时数据盘时，磁盘容量大小限制为5的倍数。
> 8.  创建实例支持创建和添加临时数据盘，但不支持单独创建或添加临时数据盘。
> 9.  临时数据盘不支持挂载、卸载、删除。
> 10. 普通实例的临时数据盘最大不能超过500G。
> 11. 指定子网和安全组创建，要求子网和安全组必须同时指定或同时不指定，同时指定的子网和安全组必须同属于一个VPC，都不指定会使用默认子网和默认安全组。
> 12. 指定公网IP带宽创建，计费方式为按照带宽计费。
> 13. 创建接口为异步创建，可通过查询实例详情接口查询实例状态
> 14. 每个实例最多只能购买一块临时数据盘。
> 15. 实例的临时数据盘默认只有hp1类型。
> 16. 创建存储优化型实例必须购买临时数据盘，通过ephemeralDisks指定临时盘数据盘大小，默认nvme类型数据盘，无需指定。

## 创建竞价实例
使用以下代码可以创建BCC实例：
```go
createInstanceArgs := &CreateInstanceArgs{
    // 输入你要创建instance使用的镜像ID
    ImageId:          "your-choose-image-id",
    // BCC实例类型
    InstanceType: "InstanceType"
    // BCC核数
    CpuCount: CpuCount
    // BCC的内存大小GB
    MemoryCapacityInGB: MemoryCapacityInGB
    // 系统盘大小GB
	RootDiskSizeInGb      int              "rootDiskSizeInGb"
	// 设置待查询的竞价实例的系统盘介质
    RootDiskStorageType   StorageType      "rootDiskStorageType"
	// 临时盘数据盘大小
    EphemeralDisks        []EphemeralDisk  "ephemeralDisks"
	// 创建需要创建的CDS磁盘列表
    CreateCdsList         []CreateCdsModel "createCdsList"
	// 设置创建BCC使用的网络带宽大小
    NetWorkCapacityInMbps int              networkCapacityInMbps
	// 设置需要创建BCC使用的DCC服务器id
    DedicateHostId        string           "dedicatedHostId"
	// 设置待查询的竞价实例的购买个数
    PurchaseCount         int              purchaseCount
	// 实例名称
    Name                  string           "name"
	// 设置BCC虚机密码
    AdminPass             string           "adminPass"
	// 设置可用区
    ZoneName              string           "zoneName"
	// 指定子网和安全组创建，要求子网和安全组必须同时指定或同时不指定，
    // 同时指定的子网和安全组必须同属于一个VPC，都不指定会使用默认子网和默认安全组。
    // 设置创建BCC使用的子网
    SubnetId              string           "subnetId"
	// // 设置创建BCC使用的安全组
    SecurityGroupId       string           "securityGroupId"
	// 设置需要创建GPU卡信息
    GpuCard               string           "gpuCard"
	// 设置需要创建FPGA卡信息
    FpgaCard              string           "fpgaCard"
	// 设置GPU卡或FPGA卡数量
    CardCount             string           "cardCount"
	// 设置按月付费或者按年付费 月是"month",年是"year"
    AutoRenewTimeUnit     string           "autoRenewTimeUnit"
	// 设置自动续费的时间 按月是1-9 按年是 1-3
    AutoRenewTime         int              autoRenewTime
	// cds是否自动续费 是:true 否:false
    CdsAutoRenew          bool             cdsAutoRenew
	// 待创建实例指定的标签是否需要和已有标签键进行关联，默认为false。注意值为true时要保证该标签键已存在
    RelationTag           bool             relationTag
	// 待创建的标签列表
    Tags                  []model.TagModel tags
	// 指定实例所在的部署集id
    DeployId              string           "deployId"
	// 设置创建BCC虚机使用的竞价模式：market 或者 custom
    BidModel              string           "bidModel"
	// 设置创建BCC虚机使用的竞价金额,只有当bidModel为custom时有效
    BidPrice              string           "bidPrice"
	// 设置要绑定的密钥对ID
    KeypairId             string           "keypairId"
	// 设置要绑定的自动快照策略ID
    AspId                 string           "aspId"
	// 公网带宽计费方式，若不指定internetChargeType，默认付费方式同BCC，预付费默认为包年包月按带宽，后付费默认为按使用带宽计费。（BANDWIDTH_PREPAID：预付费按带宽结算；TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费；BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费）
    InternetChargeType    string           "internetChargeType"
    // 设置内网IP
    InternalIps           []string         internalIps
    // 设置创建BCC虚机使用的竞价模式：market 或者 custom
    BidModel              string           "bidModel"
	// 设置创建BCC虚机使用的竞价金额,只有当bidModel为custom时有效
    BidPrice              string           "bidPrice"
	// 使用 uuid 生成一个长度不超过64位的ASCII字符串
    ClientToken           string           "random-uuid"
    // 创建实例支持幂等的token,成功后永久有效
    RequestToken          string           "requestToken"
}
if res, err := bccClient.CreateBidInstance(createInstanceBySpecArgs); err != nil {
    fmt.Println("create instance failed: ", err)
} else {
    fmt.Println("create instance success, instanceId: ", res.InstanceIds[0])
}
```
## 取消竞价实例订单
通过以下代码可以取消竞价实例订单
```go
cancelBidOrderRequest := &CancelBidOrderRequest{
    // 订单ID
    OrderId     string "orderId"
}
if err := bccClient.CancelBidOrder(cancelBidOrderRequest); err != nil {
    fmt.Println("CancelBidOrderRequest failed: ", err)
} else {
    fmt.Println("CancelBidOrderRequest success.")
}
```
## 查询竞价实例套餐
通过以下代码可以查询竞价实例套餐
```go
if res, err := bccClient.ListBidFlavor(); err != nil {
    fmt.Println("List bidding instance flavors failed: ", err)
} else {
    fmt.Println("List bidding instance flavors success, result: ", res)
}
```
## 查询竞价实例市场价
通过以下代码可以查询竞价实例市场价
```go
createCdsList := []api.CreateCdsModel{{
    // 设置CDS磁盘容量，必须为大于0的整数，单位为GB，大小为0~5120G，可选参数
    CdsSizeInGB: 40,
    // 设置CDS磁盘存储类别，默认为高性能型，可选参数
    StorageType: api.StorageTypeHP1,
    // 设置快照ID
    SnapshotId:  "your-snapshot-id",
    // 设置加密密钥，true或false
    EncryptKey:  true,
}}
tagList := []model.TagModel{{
    // 设置要查询的tagKey
    TagKey:   "your-tag-key",
    // 设置要查询的tagValue
    TagValue: "your-tag-value",
}}
args := &api.GetBidInstancePriceArgs{
    // 使用 uuid 生成一个长度不超过64位的ASCII字符串
    ClientToken:           "random-uuid",
     // 设置待查询的虚拟机实例类型，具体可选类型参见InstanceType
    InstanceType:          "your-choose-instance-type",
    // 设置待查询虚拟机实例的CPU核数
    CpuCount:              1,
    // 设置待查询虚拟机实例的内存容量
    MemoryCapacityInGB:    2,
    // 设置待查询虚拟机实例的系统盘大小，单位GB，默认是40GB，范围为[40, 100]GB，超过40GB按照云磁盘价格收费。注意指定的系统盘大小需要满足所使用镜像最小磁盘空间限制。
    RootDiskSizeInGb:      45,
    // 设置待查询虚拟机实例系统盘介质，默认使用SSD型云磁盘，可指定系统盘磁盘类型可参见StorageType
    RootDiskStorageType:   api.StorageTypeCloudHP1,
    // 设置待查询的CDS磁盘列表
    CreateCdsList:           createCdsList,
    // 设置批量查询（购买）的虚拟机实例个数，必须为大于0的整数，可选参数，缺省为1
    PurchaseCount:         1,
    // 设置虚拟机名字，可选参数
    Name:                  "your-choose-instance-name",
    // 设置实例管理员密码(8-16位字符，英文，数字和符号必须同时存在，符号仅限!@#$%^*()),可选参数
    AdminPass:             "your-admin-pass",
    // 设置待查询实例所要绑定的密钥对ID,可选参数
    KeypairId:             "your-keypair-id",
    // 设置自动快照策略ID,可选参数
    AspId:                 "your-asp-id",
    // 设置待查询虚拟机实例的镜像ID,可选参数
    ImageId:               "your-image-id",
    // 设置竞价实例出价模型， 市场价: "market" 自定义："custom",可选参数
    BidModel:              "your-bid-model",
    // 设置竞价实例出价金额，若是自定义出价，且出价金额小于市场价，则不允许创建。当bidModel='custom'时才有效,可选参数
    BidPrice:              "your-bid-price",
    // 设置公网带宽，单位为Mbps。必须为0~200之间的整数，为0表示不分配公网IP，默认为0Mbps,可选参数
    NetWorkCapacityInMbps: 20,
    // 设置待查询实例指定的标签是否需要和已有标签键进行关联，默认为false。注意值为true时要保证该标签键已存在,可选参数
    RelationTag:           false,
    // 设置待查询的标签列表,可选参数
    Tags:                    tagList,
    // 设置securityGroup信息，为空时将使用默认安全组,可选参数
    SecurityGroupId:       "your-security-group-id",
    // 设置subnet信息，为空时将使用默认子网,可选参数
    SubnetId:              "your-subnet-id",
    // 设置指定zone信息，默认为空，由系统自动选择,可选参数
    ZoneName:              "your-zone-name",
    // 设置公网带宽计费方式,可选参数
    InternetChargeType:    "your-internet-charge-type",
}
if res, err := bccClient.GetBidInstancePrice(args); err != nil {
    fmt.Println("Get bidding instance price failed: ", err)
} else {
    fmt.Println("Get bidding instance price success, result: ", res)
}
```

##查询实例套餐库存

查询实例资源套餐规格对应的库存。
```go
// 实例类型
instanceType := "instanceType"
// CPU核数
cpuCount := cpuCount
// 内存容量（GB）
memoryCapacityInGB := memoryCapacityInGB
// 可用区名称
zoneName := "zoneNamen"
// GPU卡类型，GPU和VGPU可填
gpuCard := "gpuCard"
// GPU卡数量，GPU和VGPU可填
cardCount := "cardCount"
//本地盘信息
ephemeralDisks := []EphemeralDisks{{
      "storageType": "ssd",
      "sizeInGB": sizeInGB,
}}

args := &api.CreateInstanceStockArgs{
    InstanceType:     instanceType,
    CpuCount: cpuCount,
    MemoryCapacityInGB:  memoryCapacityInGB,
    ZoneName:  zoneName,
    GpuCard:  gpuCard,
    CardCount:  cardCount,
    EphemeralDisks:  ephemeralDisks,
}
if res, err := bccClient.GetInstanceCreateStock(args); err != nil {
    fmt.Println("GetInstanceCreateStock failed: ", err)
} else {
    fmt.Println("GetInstanceCreateStock success: ", res)
}

```
###实例扩缩容库存查询

实例变配余量查询
```go
// 实例id
instanceId := "instanceId"
// CPU核数
cpuCount := cpuCount
// 内存容量（GB）
memoryCapacityInGB := memoryCapacityInGB

args := &api.CreateInstanceStockArgs{
    InstanceId:     instanceId,
    CpuCount: cpuCount,
    MemoryCapacityInGB:  memoryCapacityInGB,
}
if res, err := bccClient.ResizeInstanceStockArgs(args); err != nil {
    fmt.Println("ResizeInstanceStockArgs failed: ", err)
} else {
    fmt.Println("ResizeInstanceStockArgs success: ", res)
}

```
###查询带部署集相关字段实例详情
   查询带部署集相关字段实例详情：
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
//设置是否展示部署集字段,true or false
isDeploySet := "your-isDeploySet"

if res, err := bccClient.GetInstanceDetailWithDeploySet(instanceId,isDeploySet); err != nil {
fmt.Println("get instance detail failed: ", err)
} else {
fmt.Println("get instance detail success, result: ", res)
}

```

## 查询可关机不计费的实例列表
查询可关机不计费的BCC实例列表：
```go
listInstanceArgs := &ListInstanceArgs{
    // 批量获取列表的查询起始位置，是一个由系统产生的字符串
    Marker          string
	// 设置返回数据大小，缺省为1000
    MaxKeys         int
	// 通过internal Ip过滤
    InternalIp      string
	// 通过DedicatedHostId过滤
    DedicatedHostId string
	// 通过ZoneName过滤
    ZoneName        string
	// 通过KeypairId过滤
    KeypairId       string
}

if res, err := bccClient.GetInstanceNoChargeList(listInstanceArgs); err != nil {
    fmt.Println("GetInstanceNoChargeList failed: ", err)
} else {
    fmt.Println("GetInstanceNoChargeList success, result: ", res)
}
```

### 查询实例列表

以下代码可以查询BCC虚机实例列表,支持通过内网ip、专属服务器id、可用区名称进行筛选
```go
args := &api.ListInstanceArgs{}

// 若要查询某个内网IP对应的实例列表，可以配置以下参数
args.InternalIp = "1.1.1.1"

result, err := client.ListInstances(args)
if err != nil {
    fmt.Println("list instance failed:", err)
} else {
    fmt.Println("list instance success: ", result)
}
```

### 查询指定实例详情

使用以下代码可以查询指定BCC虚机的详细信息
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
//设置是否展示部署集字段
isDeploySet := "your-isDeploySet"

result, err := client.GetInstanceDetail(instanceId,isDeploySet)
if err != nil {
    fmt.Println("get instance detail failed:", err)
} else 
    fmt.Println("get instance detail success ", result)
}
```

### 启动实例

如下代码可以启动一个实例
```go
err := client.StartInstance(instanceId)
if err != nil {
    fmt.Println("start instance failed:", err)
} else {
    fmt.Println("start instance success")
}
```

> **提示：**
> - 实例状态必须为Stopped，调用此接口才能成功返回，否则返回409错误
> - 该接口调用后，实例会进入Starting状态

### 停止实例

如下代码可以停止一个实例
```go
// 以下代码可以强制停止一个实例
err := client.StopInstance(instanceId, true)
if err != nil {
    fmt.Println("stop instance failed:", err)
} else {
    fmt.Println("stop instance success")
}
```

## 停止实例(支持强制停止&关机不计费)
使用以下代码停止实例:

```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
// 设置是否强制停止，强制停止等同于断电处理，可能丢失实例操作系统中未写入磁盘的数据
forceStop ：= false
// 设置是否关机不计费，TRUE为关机不计费，FALSE为关机计费。注意：只有白名单用户才可以实行关机不计费
stopWithNoCharge := false

if err := bccClient.StopInstanceWithNoCharge(instanceId, forceStop, stopWithNoCharge); err != nil {
    fmt.Println("Stop instance failed: ", err)
} else {
    fmt.Println("Stop instance success.")
}
```

> **提示：**
> -   系统后台会在实例实际 Stop 成功后进入“已停止”状态。
> -   只有状态为 Running 的实例才可以进行此操作，否则提示 409 错误。
> -   实例支持强制停止，强制停止等同于断电处理，可能丢失实例操作系统中未写入磁盘的数据。

### 重启实例

如下代码可以重启实例
```go
// 以下代码可以强制重启一个实例
err := client.RebootInstance(instanceId, true)
if err != nil {
    fmt.Println("reboot instance failed:", err)
} else {
    fmt.Println("reboot instance success")
}
```

> **提示：**
> -   只有状态为 Running 的实例才可以进行此操作，否则提示 409 错误。
> -   接口调用成功后实例进入 Starting 状态。
> -   支持强制重启，强制重启等同于传统服务器的断电重启，可能丢失实例操作系统中未写入磁盘的数据。

### 修改实例密码

如下代码可以修改实例密码
```go
args := &api.ChangeInstancePassArgs{
	AdminPass: "321zaq#@!",
}
err := client.ChangeInstancePass(instanceId, args)
if err != nil {
    fmt.Println("change instance password failed:", err)
} else {
    fmt.Println("change instance password success")
}
```

> **提示：**
> 只有 Running 和 Stopped 状态的实例才可以用调用接口，否则提示 409 错误。

### 修改实例属性

如下代码可以修改实例属性
```go
args := &api.ModifyInstanceAttributeArgs{
    Name: "newInstanceName",
}
err := client.ModifyInstanceAttribute(instanceId, args)
if err != nil {
    fmt.Println("modify instance failed:", err)
} else {
    fmt.Println("modify instance success")
}
```

> **提示：**
> - 目前该接口仅支持修改实例名称

### 修改实例描述

如下代码可以修改实例描述
```go
args := &api.ModifyInstanceDescArgs{
    Description: "new Instance description",
}
err := client.ModifyInstanceDesc(instanceId, args)
if err != nil {
    fmt.Println("modify instance failed:", err)
} else {
    fmt.Println("modify instance success")
}
```

### 重装实例

如下代码可以重装实例
```go
args := &api.RebuildInstanceArgs{
	ImageId:   "m-DpgNg8lO",
	AdminPass: "123qaz!@#",
	// 设置要绑定的密钥对ID
    KeypairId string "keypairId"
}
err := client.RebuildInstance(instanceId, args)
if err != nil {
    fmt.Println("rebuild instance failed:", err)
} else {
    fmt.Println("rebuild instance success")
}
```

## 重装实例（批量）
使用以下代码重装实例:

```go
rebuildBatchInstanceArgs := &RebuildBatchInstanceArgs{
    // 输入你要重装instance使用的镜像ID
    ImageId   string "imageId"
	// 设置BCC虚机密码
    AdminPass string "adminPass"
	// 设置要绑定的密钥对ID
    KeypairId string "keypairId"
    // 实例ID集合
    InstanceIds []string "instanceIds"
}

if err := bccClient.BatchRebuildInstances(rebuildBatchInstanceArgs); err != nil {
    fmt.Println("rebuild instance failed: ", err)
} else {
    fmt.Println("rebuild instance success.")
}
```

> **提示：**
> - 实例重装后，基于原系统盘的快照会自动删除，基于原系统盘的自定义镜像会保留。

### 释放实例

如下代码可以释放实例
```go
err := client.DeleteInstance(instanceId)
if err != nil {
    fmt.Println("delete instance failed:", err)
} else {
    fmt.Println("delete instance success")
}
```

## 释放实例（POST）
使用以下代码修改实例描述:
```go
deleteInstanceWithRelateResourceArgs := &DeleteInstanceWithRelateResourceArgs{
    // 设置释放的时候是否关联释放当前时刻，实例挂载的eip+数据盘 false代表否 true代表是
    // (只有该字段为true时 deleteCdsSnapshotFlag字段才会有效，若该字段为false,deleteCdsSnapshotFlag字段的值无效）
    RelatedReleaseFlag    bool "relatedReleaseFlag"
	// 设置是否释放云磁盘快照 false代表否 true代表是
    DeleteCdsSnapshotFlag bool "deleteCdsSnapshotFlag"
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"

if err := bccClient.DeleteInstanceWithRelateResource(instanceId, deleteInstanceWithRelateResourceArgs); err != nil {
    fmt.Println("release instance failed: ", err)
} else {
    fmt.Println("release instance success.")
}
```
## 释放实例（包含预付费实例）
不区分后付费还是预付费实例，释放bcc以及关联的资源，可以使用以下代码将其释放:
```go
	args := &api.DeleteInstanceIngorePaymentArgs{
		InstanceId:      "instanceid",
		//设置是否释放eip和cds false代表否 true代表是，默认false
		RelatedReleaseFlag: true,
		//设置是否释放 false代表否 true代表是，默认false
		DeleteRelatedEnisFlag:true,
		//设置是否释放云磁盘快照 false代表否 true代表是，默认false，释放预付费bcc时DeleteCdsSnapshotFlag和RelatedReleaseFlag不存在绑定关系，
		// 选择DeleteCdsSnapshotFlag=true即会释放虚机绑定的各种快照
		// 释放后付费bcc时，DeleteCdsSnapshotFlag和RelatedReleaseFlag之间逻辑和之前逻辑保持一致
		DeleteCdsSnapshotFlag:true,
	}
	if res, err := BCC_CLIENT.DeleteInstanceIngorePayment(args); err != nil {
		fmt.Println("delete  instance failed: ", err)
	} else {
		fmt.Printf("delete  instance success, result: %s", res.String())
	}

```

> **提示：**
> -   释放单个云服务器实例，释放后实例所使用的物理资源都被收回，相关数据全部丢失且不可恢复。
> -   只有付费类型为Postpaid或者付费类型为Prepaid且已过期的实例才可以释放。
> -   实例释放后，已挂载的CDS磁盘自动卸载，，基于此CDS磁盘的快照会保留。
> -   实例释放后，基于原系统盘的快照会自动删除，基于原系统盘的自定义镜像会保留。

## 变配实例
使用以下代码可以选择CPU,MemoryCapacityInGB,EphemeralDisks变配指定BCC实例:

```go
resizeInstanceArgs := &ResizeInstanceArgs{
    // BCC核数
    CpuCount: CpuCount
    // BCC的内存大小GB
    MemoryCapacityInGB: MemoryCapacityInGB
	// 临时盘数据盘大小
    EphemeralDisks        []EphemeralDisk  "ephemeralDisks"
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"

if err := bccClient.ResizeInstance(instanceId, resizeInstanceArgs); err != nil {
    fmt.Println("resize instance failed: ", err)
} else {
    fmt.Println("resize instance success.")
}
```
## 变配实例（通过实例套餐规格）
使用以下代码可以选择CPU,MemoryCapacityInGB,EphemeralDisks变配指定BCC实例:

```go
resizeInstanceArgs := &ResizeInstanceArgs{
    // 实例套餐规格
    Spec               string          "spec"
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"

if err := bccClient.ResizeInstanceBySpec(instanceId, resizeInstanceArgs); err != nil {
    fmt.Println("resize instance failed: ", err)
} else {
    fmt.Println("resize instance success.")
}
```

## 绑定安全组
使用以下代码绑定安全组:

```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
// 设置BCC绑定的安全组
SecurityGroupId := "SecurityGroupId"

if err := bccClient.BindSecurityGroup(instanceId, SecurityGroupId); err != nil {
    fmt.Println("Bind Security Group failed: ", err)
} else {
    fmt.Println("Bind Security Group success.")
}
```
## 解绑安全组
使用以下代码解绑安全组:

```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
// 设置BCC解绑的安全组
SecurityGroupId := "SecurityGroupId"

if err := bccClient.UnBindSecurityGroup(instanceId, SecurityGroupId); err != nil {
    fmt.Println("UnBind Security Group failed: ", err)
} else {
    fmt.Println("UnBind Security Group success.")
}
```

> **提示：**
> - 每个实例至少关联一个安全组，默认关联默认安全组。
> - 如果实例仅属于一个安全组，尝试移出时，请求会报 403 错。

### 实例扩缩容

```go
args := &api.ResizeInstanceArgs{
	CpuCount:           2,
	MemoryCapacityInGB: 4,
}
err := client.ResizeInstance(instanceId, args)
if err != nil {
    fmt.Println("resize instance failed:", err)
} else {
    fmt.Println("resize instance success")
```

> **提示：**
> - 实例计费方式为预付费时，不能进行缩容操作
> - 实例计费方式为后付费时，可弹性扩缩容
> - 只有实例Running或Stopped状态时可以进行扩缩容操作
> - 实例扩缩容之后会重启一次
> - 异步接口，可通过查询实例详情接口查看扩缩容状态是否完成
> - 专属实例可以通过指定的cpu、内存以及临时盘大小，专属实例临时盘大小只支持扩容而不支持缩容，具体请参考API文档 [实例扩缩容](https://cloud.baidu.com/doc/BCC/s/1jwvyoc9l)

### 查询实例VNC地址

如下代码可以查询实例的VNC地址
```go
result, err := client.GetInstanceVNC(instanceId)
if err != nil {
    fmt.Println("get instance VNC url failed:", err)
} else {
    fmt.Println("get instance VNC url success: ", result)
}
```

> **提示：**
> -   VNC地址一次使用后即失效
> -   URL地址有效期为10分钟

### 实例续费

对BCC虚机的续费操作，可以延长过期时长，以下代码可以对磁盘进行续费
```go
args := &api.PurchaseReservedArgs{
    Billing: &api.Billing{
        PaymentTiming: api.PaymentTimingPrePaid,
        Reservation: &api.Reservation{
            ReservationLength:   1,
            ReservationTimeUnit: "month",
        }
    }
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
// 设置实例关联续费标识，默认为空字符串。
relatedRenewFlag := "CDS"

if err := bccClient.InstancePurchaseReserved(instanceId, ModifyInstanceDescArgs); err != nil {
    fmt.Println("Modify Instance Attribute failed: ", err)
} else {
    fmt.Println("Modify Instance Attribute success.")
}
```
> **注意：** 
>
> 关联续费产品(relatedRenewFlag)可选：
>
> -   CDS	        只对BCC实例关联的预付费CDS进行续费
> -   EIP	        只对BCC实例关联的预付费EIP进行续费
> -   MKT	        只对BCC实例关联的预付费MKT进行续费
> -   CDS_EIP	    只对BCC实例关联的预付费CDS、EIP进行续费
> -   CDS_MKT	    只对BCC实例关联的预付费CDS、MKT进行续费
> -   EIP_MKT	    只对BCC实例关联的预付费EIP、MKT进行续费
> -   CDS_EIP_MKT	只对BCC实例关联的预付费CDS、EIP、MKT进行续费

> **提示：**
> - BCC虚机实例扩缩容期间不能进行续费操作。
> - 续费时若实例已欠费停机，续费成功后有个BCC虚机实例启动的过程。
> - 该接口是一个异步接口。
> - 专属实例不支持续费。

### 释放实例（POST请求的释放）

如下代码可以释放实例及相关联的资源，如EIP, CDS等
```go
args := &api.DeleteInstanceWithRelateResourceArgs{
    RelatedReleaseFlag:    true,
    DeleteCdsSnapshotFlag: true,
}
err := client.DeleteInstanceWithRelateResource(instanceId, args)
if err != nil {
    fmt.Println("delete instance failed:", err)
} else {
    fmt.Println("delete instance success")
}
```

> **提示：**
> - 释放后实例所使用的物理资源都被收回，相关数据全部丢失且不可恢复。
> - 释放的时候默认只释放实例和系统盘，用户可以选择是否关联释放当前实例挂载的eip+数据盘（只能统一释放完或者不释放。而不是挂载的数据盘释放一个或者多个）是否一起释放。

### 实例变更子网

如下代码可以变更实例的子网
```go
args := &api.InstanceChangeSubnetArgs{
	InstanceId: instanceId,
	SubnetId:   subnetId,
	Reboot:     false,
}
err := client.InstanceChangeSubnet(args)
if err != nil {
    fmt.Println("change instance subnet failed:", err)
} else {
    fmt.Println("change instance subnet success")
}
```

> **提示：**
> - 变更子网后默认自动重启，用户选择是否执行该操作。
> - 变更子网的范围目前仅支持在同AZ下变更子网，不支持跨AZ或跨VPC变更子网，如果从普通子网变更至NAT专属子网请先手动解绑EIP。

### 向指定实例批量添加指定ip

```go
batchAddIpArgs := &BatchAddIpArgs{
    // 实例ID
    InstanceId string "instanceId"
    // 辅助IP，和SecondaryPrivateIpAddressCount不可同时使用
	PrivateIps []string "privateIps"
    // 自动分配IP数量，和PrivateIps不可同时使用
    SecondaryPrivateIpAddressCount int 1
}

if res, err := bccClient.BatchAddIP(batchAddIpArgs); err != nil {
    fmt.Println("BatchAddIP failed: ", err)
} else {
    fmt.Println("BatchAddIP success, result: ", res)
}
```

### 批量删除指定实例的ip

```go
privateIps := []string{"192.168.1.25"}
instanceId := "your-choose-instance-id"
batchDelIpArgs := &api.BatchDelIpArgs{
	InstanceId: instanceId,
	PrivateIps: privateIps,
}
if err := client.BatchDelIP(batchDelIpArgs); err != nil {
    fmt.Println("delete ips failed: ", err)
} else {
    fmt.Println("delete ips success.")
}
```

### 开通自动续费（包含关联产品）
自动续费仅限预付费产品

```go
bccCreateAutoRenewArgs := &api.BccCreateAutoRenewArgs{
    // 实例ID
    InstanceId: instanceId,
    // 续费单位，month，year
    RenewTimeUnit: "month",
    // 续费时长，单位：month，支持1, 2, 3, 4, 5, 6, 7, 8, 9；单位：year，支持1, 2, 3
    RenewTime: 1,
}

if err := bccClient.BatchCreateAutoRenewRules(bccCreateAutoRenewArgs); err != nil {
    fmt.Println("BatchCreateAutoRenewRules failed: ", err)
} else {
    fmt.Println("BatchCreateAutoRenewRules success.")
}
```

### 关闭自动续费（包含关联产品）
自动续费仅限预付费产品

```go
bccDeleteAutoRenewArgs := &api.BccDeleteAutoRenewArgs{
	// 实例ID
	InstanceId: instanceId,
}

if err := bccClient.BatchDeleteAutoRenewRules(bccDeleteAutoRenewArgs); err != nil {
    fmt.Println("BatchDeleteAutoRenewRules failed: ", err)
} else {
    fmt.Println("BatchDeleteAutoRenewRules success.")
}
```

## 计费变更-转预付费
使用以下代码对实例计费变更-转预付费:

```go
changeToPrepaidRequest := &ChangeToPrepaidRequest{
    // 设置将要变更预付费的时长，单位为月
    Duration    int  "duration"
	// 设置是否变更关联的数据盘，TRUE表示变更，FLASE表示不变更
    RelationCds bool "relationCds"
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"

if err := bccClient.ChangeToPrepaid(instanceId, changeToPrepaidRequest); err != nil {
    fmt.Println("ChangeToPrepaid failed: ", err)
} else {
    fmt.Println("ChangeToPrepaid success.")
}
```

## 实例绑定标签
使用以下代码对实例绑定标签:

```go
bindTagsRequest := &BindTagsRequest{
    // 设置想要绑定的标签
    ChangeTags: []model.TagModel{
        {
            TagKey:   "TagKey",
            TagValue: "TagValue",
        },
    },
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"

if err := bccClient.BindInstanceToTags(instanceId, bindTagsRequest); err != nil {
    fmt.Println("BindInstanceToTags failed: ", err)
} else {
    fmt.Println("BindInstanceToTags success.")
}
```
## 实例解绑标签
使用以下代码对实例解绑标签:

```go
unBindTagsRequest := &UnBindTagsRequest{
    // 设置想要解绑的标签
    ChangeTags: []model.TagModel{
        {
            TagKey:   "TagKey",
            TagValue: "TagValue",
        },
    },
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"

if err := bccClient.UnBindInstanceToTags(instanceId, unBindTagsRequest); err != nil {
    fmt.Println("UnBindInstanceToTags failed: ", err)
} else {
    fmt.Println("UnBindInstanceToTags success.")
}
```

### 查询可用区的磁盘信息
使用以下代码可以查询指定可用区的磁盘信息
```GO
// 设置你要操作的zoneName
zoneName := "cn-bj-a"
if res, err := bccClient.GetAvailableDiskInfo(zoneName); err != nil {
    fmt.Println("Get the specific zone flavor failed: ", err)
} else {
    fmt.Println("Get the specific zone flavor success, result: ", res)
}
```

## 磁盘管理

### 创建CDS磁盘

支持新建空白CDS磁盘或者从CDS数据盘快照创建CDS磁盘，参考以下代码可以创建CDS磁盘：

```go
// 新建CDS磁盘
args := &api.CreateCDSVolumeArgs{
    // 创建一个CDS磁盘，若要同时创建多个相同配置的磁盘，可以修改此参数
	PurchaseCount: 1, 
    // 磁盘空间大小
    CdsSizeInGB:   40,
    // 设置磁盘存储介质 
    StorageType:   api.StorageTypeSSD, 
    // 设置磁盘付费模式为后付费
    Billing: &api.Billing{
        PaymentTiming: api.PaymentTimingPostPaid,
    }, 
    // 设置磁盘名称
    Name:        "sdkCreate", 
    // 设置磁盘描述
    Description: "sdk test",
    // 快照ID
    SnapshotId    string      "snapshotId"
    // 可用区
    ZoneName      string      "zoneName"
	// 设置磁盘加密密钥
	EncryptKey    string      "encryptKey"
}
result, err := client.CreateCDSVolume(args)
if err != nil {
    fmt.Println("create CDS volume failed:", err)
} else {
    fmt.Println("create CDS volume success: ", result)
}
```

> **提示：**
> - 创建CDS磁盘接口为异步接口，可通过[查询磁盘详情](#查询磁盘详情)接口查询磁盘状态，详细接口使用请参考BCC API 文档[查询磁盘详情](https://cloud.baidu.com/doc/BCC/s/1jwvyo4ly)

### 查询磁盘列表

以下代码可以查询所有的磁盘列表，不包含临时数据盘，支持分页查询以及通过次磁盘所挂载的BCC实例id进行过滤筛选:

```go
args := &api.ListCDSVolumeArgs{}

// 设置查询绑定了特定实例的CDS磁盘
args.InstanceId = instanceId

result, err := client.ListCDSVolume(args)
if err != nil {
    fmt.Println("list CDS volume failed:", err)
} else {
    fmt.Println("list CDS volume success: ", result)
}
```

### 查询磁盘详情

通过磁盘id可以获取对应磁盘的详细信息，以下代码可以查询磁盘详情：

```go
result, err := client.GetCDSVolumeDetail(volumeId)
if err != nil {
    fmt.Println("get CDS volume detail failed:", err)
} else {
    fmt.Println("get CDS volume detail success: ", result)
}
```

### 挂载CDS磁盘

可以将未挂载的磁盘挂载在对应的BCC虚机下，以下代码将一个CDS挂载在对应的BCC虚机下:

```go
args := &api.AttachVolumeArgs{
	InstanceId: instanceId,
}
result, err := client.AttachCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("attach CDS volume to instance failed:", err)
} else {
    fmt.Println("attach CDS volume to instance success: ", result)
}
```

> **提示：**
> - CDS磁盘需要挂载在与其处在相同zone下的虚机实例上，否则将返回403错误。
> - 只有磁盘状态为 Available 且实例状态为 Running 或 Stopped 时才允许挂载，否则调用此接口将返回 409 错误。

### 卸载CDS磁盘

可以将已挂载的磁盘从对应的BCC虚机上卸载下来，以下代码卸载CDS磁盘:

```go
args := &api.DetachVolumeArgs{
	InstanceId: instanceId,
}
err := client.DetachCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("detach CDS volume from instance failed:", err)
} else {
    fmt.Println("detach CDS volume from instance success")
}
```

> **提示：**
> - 只有实例状态为 Running 或 Stopped 时，磁盘才可以执行此操作，否则将提示 409 错误。
> - 如果 volumeId 的磁盘不挂载在 instanceId 的实例上，该操作失败，提示 404 错误。

### 释放CDS磁盘

用于释放未挂载的CDS磁盘，可指定是否删除磁盘关联的快照，缺省情况下，该磁盘的所有快照将保留，但会删除与磁盘的关联关系:

```go
err = client.DeleteCDSVolume(volumeId)
if err != nil {
    fmt.Println("delete CDS volume failed:", err)
} else {
    fmt.Println("delete CDS volume success")
}
```

> **提示：**
> - 已挂载的CDS磁盘不能释放，系统盘不能释放。
> - 磁盘释放后不可恢复。缺省情况下，该磁盘的所有快照将保留，但会删除与磁盘的关联关系。
> - 只有磁盘状态为 Available 或 Expired 或 Error 时才可以执行此操作，否则将提示 409 错误。

## 释放磁盘（POST）
使用以下代码可以释放磁盘：
```go
deleteCDSVolumeArgs := &DeleteCDSVolumeArgs{
	// 设置关联释放手动快照，取值为"on"时，会删除磁盘关联的手动快照
    ManualSnapshot string "on"
	// 设置关联释放自动快照，取值为"on"时，会删除磁盘关联的自动快照
    AutoSnapshot   string "on"
}
// 设置你要操作的volumeId
volumeId := "your-choose-volume-id"
if err := bccClient.DeleteCDSVolumeNew(volumeId, deleteCDSVolumeArgs); err != nil {
    fmt.Println("DeleteCDSVolumeNew failed: ", err)
} else {
    fmt.Println("DeleteCDSVolumeNew success")
}
```

### 磁盘重命名

如下代码可以给一个CDS磁盘重命名
```go
args := &api.RenameCSDVolumeArgs{
	Name: "testVolume",
}
err := client.RenameCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("rename CDS volume failed", err)
} else {
    fmt.Println("rename CDS volume success")
}
```

### 修改磁盘属性

可以使用以下代码修改指定磁盘名称、描述信息：

```go
args := &api.ModifyCSDVolumeArgs{
	CdsName: "aaa",
	Desc:    "desc",
}
err := client.ModifyCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("modify CDS volume failed: ", err)
} else {
    fmt.Println("modify CDS volume success")
}
```

### 磁盘计费变更

可以使用以下代码变更磁盘计费方式，仅支持后付费转预付费、预付费转后付费两种方式。变更为预付费需要指定购买时长。

```go
args := &api.ModifyChargeTypeCSDVolumeArgs{
    Billing: &api.Billing{
        PaymentTiming: api.PaymentTimingPrePaid,
        Reservation: &api.Reservation{
            ReservationLength:   1,
            ReservationTimeUnit: "month",
        }
    }
}
err := client.ModifyChargeTypeCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("modify CDS volume charge type failed:", err)
} else {
    fmt.Println("modify CDS volume charge type success")
}
```

## 磁盘开通自动续费
通过以下代码可以开通磁盘自动续费
```go
// 设置要开通自动续费的磁盘id
volumeId := "your-volume-id"
// 设置按月付费或者按年付费 月是"month",年是"year"
renewTimeUnit := "your-renew-time-unit"
// 设置自动续费的时间 按月是1-9 按年是 1-3
renewTime := 3
args := &api.AutoRenewCDSVolumeArgs{
    VolumeId:       volumeId,
    RenewTimeUnit: renewTimeUnit,
    RenewTime:       renewTime,
}
if err := bccClient.AutoRenewCDSVolume(args); err != nil {
    fmt.Println("Auto renew cds volume failed: ", err)
} else {
    fmt.Println("Auto renew cds volume success")
}
```
>- 注意数据盘才能进行自动续费操作
>- 注意过期状态磁盘不能进行自动续费操作
>- 异步接口，可通过查询磁盘详情接口查询磁盘到期时间

## 磁盘取消自动续费
通过以下代码可以取消磁盘自动续费
```go
// 设置要取消自动续费的磁盘id
volumeId := "your-volume-id"
args := &api.CancelAutoRenewCDSVolumeArgs{
    VolumeId:       volumeId,
}
if err := bccClient.CancelAutoRenewCDSVolume(args); err != nil {
    fmt.Println("Cancel auto renew cds volume failed: ", err)
} else {
    fmt.Println("Cancel auto renew cds volume success")
}
```
>- 注意数据盘才能取消进行自动续费操作
>- 异步接口，可通过查询磁盘详情接口查询磁盘到期时间

### 磁盘扩容

使用以下代码可以对磁盘进行扩大容量操作：

```go
args := &api.ResizeCSDVolumeArgs{
	// 磁盘容量GB
    NewCdsSizeInGB int         `json:"newCdsSizeInGB"`
    // 磁盘类型
    NewVolumeType  StorageType `json:"newVolumeType"`
}
err := client.ResizeCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("resize CDS volume failed:", err)
} else {
    fmt.Println("resize CDS volume success")
}
```

> **提示：**
> - 磁盘只能进行扩大容量，不支持缩小容量。
> - 只有Available状态的磁盘，才能进行扩容操作
> - 磁盘扩容是一个异步接口，可通过[查询磁盘详情](#查询磁盘详情)接口查询磁盘扩容状态。

### 回滚磁盘

可以使用指定磁盘自身的快照回滚磁盘内容，使用以下代码可以对磁盘进行回滚：

```go
args := &api.RollbackCSDVolumeArgs{
    SnapshotId: snapshotId
}
err := client.RollbackCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("rollback CDS volume failed:", err)
} else {
    fmt.Println("rollback CDS volume success")
}
```

> **提示：**
> - 磁盘状态必须为 Available 才可以执行回滚磁盘操作。
> - 指定快照id必须是由该磁盘id创建的快照。
> - 若是回滚系统盘，实例状态必须为 Running 或 Stopped 才可以执行此操作。
> - 回滚系统盘快照，自本次快照以来的系统盘数据将全部丢失，不可恢复。

### 磁盘续费

对磁盘的续费操作，可以延长过期时长，以下代码可以对磁盘进行续费：

```go
args := &api.PurchaseReservedCSDVolumeArgs{
    Billing: &api.Billing{
        PaymentTiming: api.PaymentTimingPrePaid,
        Reservation: &api.Reservation{
            ReservationLength:   1,
            ReservationTimeUnit: "month",
        }
    }
}
err := client.PurchaseReservedCDSVolume(volumeId, args)
if err != nil {
    fmt.Println("purchase reserve CDS volume failed:", err)
} else {
    fmt.Println("purchase reserve CDS volume success")
}
```

### 释放CDS磁盘（新）

如下代码可以释放一个CDS磁盘及相关联的快照
```go
args := &api.DeleteCDSVolumeArgs{
    // 删除与磁盘关联的手动快照
    ManualSnapshot: "on",
    // 删除与磁盘关联的自动快照
    AutoSnapshot:   "on",
}
err := client.DeleteCDSVolumeNew(volumeId, args)
if err != nil {
    fmt.Println("create instance failed:", err)
} else {
    fmt.Println("create instance ", )
}
```

> **提示：**
> - 该接口用于释放未挂载的CDS磁盘，系统盘不能释放。
> - 与老接口功能上的区别在于，可以控制是否删除与磁盘关联的快照。

### 查询可用区的磁盘信息
使用以下代码可以查询指定可用区的磁盘信息
```GO
// 设置你要操作的zoneName
zoneName := "cn-bj-a"
if res, err := bccClient.GetAvailableDiskInfo(zoneName); err != nil {
    fmt.Println("Get the specific zone flavor failed: ", err)
} else {
    fmt.Println("Get the specific zone flavor success, result: ", res)
}
```

## 镜像管理

### 创建自定义镜像
- 该接口用于创建自定义镜像，默认配额20个每账号，创建后的镜像可用于创建实例。
- 支持通过实例创建和通过快照创建两种方式。
- 当通过实例创建时，只有 Running 或 Stopped 状态的实例才可以执行成功，否则会提示 409 错误。
- 仅限通过系统盘快照创建自定义镜像。
- 当通过快照创建时，只有 Available 状态的快照才可以执行成功，否则会提示 409 错误
```go
// 待创建的自定义镜像名称，支持大小写字母、数字、中文以及-_ /.特殊字符，必须以字母开头，长度1-65。
imageName := "your-imageName"
// 当从实例创建镜像时，此参数是指用于创建镜像的实例ID。可选参数，与 snapshotId 不同时存在，同时存在时只取instanceId
instanceId := "your-instanceId"
// 当从快照创建镜像时，此参数是指用于创建镜像的快照ID。可选参数，与 instanceId 不同时存在，同时存在时只取instanceId
snapshotId := "your-snapshotId"
// 是否关联CDS，默认为false
relateCds := true or false
args := &api.CreateImageArgs{
    ImageName:      imageName,
    InstanceId:     instanceId,
    SnapshotId:     snapshotId,
    RelateCds:      relateCds,
}
if res, err := bccClient.CreateImage(args); err != nil {
    fmt.Println("create image failed: ", err)
} else {
    fmt.Println("create image success:", res)
}
```

> 注意，创建自定义镜像，默认配额20个每账号。

### 查询镜像列表
- 使用以下代码可以查询有权限的镜像列表。
- 查询的镜像信息中包括系统镜像、自定义镜像和服务集成镜像。
- 支持按 imageType 来过滤查询，此参数非必需，缺省为 All,即查询所有类型的镜像。
```go
// 批量获取列表的查询的起始位置，是一个由系统生成的字符串,可选参数
marker := "your-marker"
// 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000,可选参数
maxKeys := your-maxKeys
// 指定要查询何种类型的镜像，包括All(所有)，System(系统镜像/公共镜像)，Custom(自定义镜像)，Integration(服务集成镜像)，Sharing(共享镜像)，GpuBccSystem(GPU专用公共镜像)，GpuBccCustom(GPU专用自定义镜像)，FpgaBccSystem(FPGA专用公共镜像)，FpgaBccCustom(FPGA专用自定义镜像)，缺省值为All
imageType := "your-imageType"

args := &api.ListImageArgs{
	Marker:         marker,
    MaxKeys:        maxKeys,
    ImageType:      imageType,
}
if res, err := bccClient.ListImage(args); err != nil {
    fmt.Println("get image list failed: ", err)
} else {
    fmt.Println("get image list success,res: ", res)
}
```

> 具体的镜像类型可详细参考BCC API文档[查询镜像列表](https://cloud.baidu.com/doc/BCC/s/Ajwvynu5r)

### 查询镜像详情

以下代码可以查询镜像详细信息：

```go
result, err := client.GetImageDetail(imageId)
if err != nil {
    fmt.Println("get image detail failed:", err)
} else {
    fmt.Println("get image detail success: ", result.Image)
}
```

### 删除自定义镜像
- 该接口用于删除用户自己的指定的自定义镜像，仅限自定义镜像，系统镜像和服务集成镜像不能删除。
- imageId 所指定的镜像不存在，提示404错误。
- 镜像删除后无法恢复，不能再用于创建、重置实例。

以下代码可以删除一个自定义镜像：

```go
err := client.DeleteImage(imageId)
if err != nil {
    fmt.Println("delete image failed:", err)
} else {
    fmt.Println("delete image success")
}
```

### 跨区域复制自定义镜像
- 用于用户跨区域复制自定义镜像，仅限自定义镜像，系统镜像和服务集成镜像不能复制
- regions如北京"bj",广州"gz",苏州"su"，可多选：
```go
 args := &api.RemoteCopyImageArgs{
    Name:       "test2",
    DestRegion: []string{"gz"},
 }
 err := client.RemoteCopyImage(imageId, args)
 if err != nil {
     fmt.Println("remote copy image failed:", err)
 } else {
     fmt.Println("remote copy image success")
 }
```

### 取消跨区域复制自定义镜像

用于取消跨区域复制自定义镜像，仅限自定义镜像，系统镜像和服务集成镜像不能复制：

```go
err := client.CancelRemoteCopyImage(imageId)
if err != nil {
    fmt.Println("cancel remote copy image failed:", err)
} else {
    fmt.Println("cancel remote copy image success")
}
```

### 共享自定义镜像

- 该接口用于共享用户自己的指定的自定义镜像，仅限自定义镜像，系统镜像和服务集成镜像不能共享。
- imageId 所指定的镜像不存在，提示404错误。
- 镜像共享后，被共享的用户可以使用此镜像创建、重置实例。
- 请求参数中的account和accountId均为可选参数，但不能全部为空，当两个参数同时出现时，服务端会自动去重。

```go
// 待共享的用户id
accountId := "your-accountId"
//待共享的用户名
account := "your-account"
// 待共享的镜像ID
imageId := "your-imageId"

args := &api.SharedUser{
	AccountId:  accountId,
    Account:    account,
}
if  err := bccClient.ShareImage(imageId,args); err != nil {
    fmt.Println("ShareImage failed: ", err)
} else {
    fmt.Println("ShareImage success")
}
```

### 取消共享自定义镜像

- 该接口用于取消共享用户自己的指定的自定义镜像，仅限自定义镜像，系统镜像和服务集成镜像不能共享。
- imageId 所指定的镜像不存在，提示404错误。
- 镜像取消共享后，被取消共享的用户不能再使用此镜像创建、重置实例。
- 请求参数中的account和accountId均为可选参数，但不能全部为空，当两个参数同时出现时，服务端会自动去重。

```go
// 待共享的用户id
accountId := "your-accountId"
//待共享的用户名
account := "your-account"
// 待共享的镜像ID
imageId := "your-imageId"

args := &api.SharedUser{
	AccountId:  accountId,
    Account:    account,
}
if  err := bccClient.UnShareImage(imageId,args); err != nil {
    fmt.Println("UnShareImage failed: ", err)
} else {
    fmt.Println("UnShareImage success")
}
```

### 查询镜像已共享用户列表
- mageId 所指定的镜像不存在，提示404错误。
用于查询镜像已共享的用户列表：

```go
result, err := client.GetImageSharedUser(imageId)
if err != nil {
    fmt.Println("get image shared user list failed: ", err)
} else {
    fmt.Println("get image shared user list success: ", result)
}
```

### 根据实例ID批量查询OS信息

如下代码可以根据实例的ID来查询相应OS的信息
```go
args := &api.GetImageOsArgs{
    InstanceIds: []string{instanceId},
}
result, err := client.GetImageOS(args)
if err != nil {
    fmt.Println("get image os failed:", err)
} else {
    fmt.Println("get image os success: ", result)
}
```

## 快照管理

### 创建快照

如下代码可以创建一个快照
```go
args := &api.CreateSnapshotArgs{
	VolumeId:     volumeId,
	SnapshotName: "sdk",
	Description:  "create by sdk",
}
result, err := client.CreateSnapshot(args)
if err != nil {
    fmt.Println("create snapshot failed:", err)
} else {
    fmt.Println("create snapshot success: ", result)
}
```

### 查询快照列表

如下代码可以查询当前账户下所有快照的列表
```go
args := &api.ListSnapshotArgs{}
result, err := client.ListSnapshot(args)
if err != nil {
    fmt.Println("list all snapshot failed:", err)
} else {
    fmt.Println("list all snapshot success: ", result)
}
```

### 查询快照详情

如下代码可以查询特定快照的详细信息
```go
result, err := client.GetSnapshotDetail(snapshotId)
if err != nil {
    fmt.Println("get snapshot detail failed:", err)
} else {
    fmt.Println("get snapshot detail success: ", result)
}
```

### 删除快照

如下代码可以删除一个快照
```go
err := client.DeleteSnapshot(snapshotId)
if err != nil {
    fmt.Println("delete snapshot failed:", err)
} else {
    fmt.Println("delete snapshot success")
}
```
### 查询快照链列表
通过以下代码可以查询快照链列表
```go
// 设置排序属性：chainId(快照链id,默认值),chainSize(快照链大小),volumeSize(磁盘大小)。可选参数
orderBy := "your-order-by"
// 设置排序方式:asc(正序,默认值), desc(倒序)。可选参数
order := "your-order"
// 设置每页容量，默认值为1000，可选参数
pageSize := 100
// 设置页数，默认为1，可选参数
pageNo := 1
// 设置磁盘ID，该字段非空则只返回这个磁盘的快照链信息,可选
volumeId := "your-volume-id"
args := &api.ListSnapshotChainArgs{
    OrderBy:    orderBy,
    Order:      order,
    PageSize:   pageSize,
    PageNo:     pageNo,
    VolumeId:   volumeId,
}
if res, err := bccClient.ListSnapshotChain(args); err != nil {
    fmt.Println("get snapshot chain list failed: ", err)
} else {
    fmt.Println("get snapshot chain list success, SnapshotId: ", res.Snapshot.Id)
}
```

## 自动快照策略管理

### 创建自动快照策略

如下代码可以创建一个自动快照策略
```go
args := &api.CreateASPArgs{
    Name:           "sdkCreate", 
    // 设置一天中做快照的时间点，取值为0~23，0为午夜12点
    // 例如设置做快照的时间点为下午两点：
    TimePoints:     []string{"14"}, 
    // 设置一周中做快照的时间，取值为0~6，0代表周日，1~6代表周一到周六
    // 例如设置做快照的时间为礼拜五：
    RepeatWeekdays: []string{"5"}, 
    // 设置自动快照保留天数，取-1则永久保留
    RetentionDays:  "7",
}
result, err := client.CreateAutoSnapshotPolicy(args)
if err != nil {
    fmt.Println("create auto snapshot policy failed:", err)
} else {
    fmt.Println("ceate auto snapshot policy success: ", result)
}
```

### 绑定自动快照策略

如下代码可以将自动快照策略绑定到某个CDS磁盘
```go
args := &api.AttachASPArgs{
    // 设置需要绑定的磁盘ID列表
	VolumeIds: []string{volumeId},
}
err := client.AttachAutoSnapshotPolicy(aspId, args)
if err != nil {
    fmt.Println("attach auto snapshot policy with CDS volume failed:", err)
} else {
    fmt.Println("attach auto snapshot policy with CDS volume success")
}
```

### 解绑自动快照策略

如下代码可以将自动快照策略与特定CDS磁盘解除绑定
```go
args := &api.DetachASPArgs{
    // 设置需要解绑的磁盘ID列表
	VolumeIds: []string{volumeId},
}
err := client.DetachAutoSnapshotPolicy(aspId, args)
if err != nil {
    fmt.Println("detach auto snapshot policy from CDS volume failed:", err)
} else {
    fmt.Println("detach auto snapshot policy from CDS volume success")
}
```

### 删除自动快照策略

如下代码可以删除自动快照策略
```go
err := client.DeleteAutoSnapshotPolicy(aspId)
if err != nil {
    fmt.Println("delete auto snapshot policy failed:", err)
} else {
    fmt.Println("delete auto snapshot policy success")
}
```

### 查询自动快照策略列表

如下代码可以查询到当前账户下当前区域所有自动快照策略的列表
```go
args := &api.ListASPArgs{}
result, err := client.ListAutoSnapshotPolicy(args)
if err != nil {
    fmt.Println("list all auto snapshot policy failed:", err)
} else {
    fmt.Println("list all auto snapshot policy success: ", result)
}
```

### 查询自动快照策略详情

如下代码可以查询到特定自动快照策略的详细信息
```go
result, err := client.GetAutoSnapshotPolicy(aspId)
if err != nil {
    fmt.Println("get auto snapshot policy detail failed:", err)
} else {
    fmt.Println("get auto snapshot policy detail success", result)
}
```

### 自动快照策略变更

如下代码可以更新一个自动快照策略
```go
args := &api.UpdateASPArgs{
	Name:           "testUpdate",
	TimePoints:     []string{"10"},
	RepeatWeekdays: []string{"0", "1", "7"},
	RetentionDays:  "2",
	AspId:          aspId,
}
err := client.UpdateAutoSnapshotPolicy(args)
if err != nil {
    fmt.Println("update auto snapshot policy failed:", err)
} else {
    fmt.Println("update auto snapshot policy success")
}
```


## 安全组管理

## 查询安全组列表

 以下代码可以查询安全组列表：

```go
args := &api.ListSecurityGroupArgs{}

// 设置筛选的实例Bcc实例id
args.InstanceId = instanceId

// 设置筛选的安全组绑定的VPC实例ID
args.VpcId = vpcId

result, err := client.ListSecurityGroup(args)
if err != nil {
    fmt.Println("list all security group failed:", err)
} else {
    fmt.Println("list all security group success: ", result)
}
```

## 创建安全组

以下代码可以创建一个安全组：

```go
args := &api.CreateSecurityGroupArgs{
    // 设置安全组名称
	Name: "sdk-create",
    // 设置安全组规则
	Rules: []api.SecurityGroupRuleModel{
		{
            // 设置安全组规则备注
			Remark:        "备注",
            // 设置协议类型
			Protocol:      "tcp",
            // 设置端口范围，默认空时为1-65535，可以指定80等单个端口
			PortRange:     "1-65535",
            // 设置入站出站，取值ingress/egress
			Direction:     "ingress",
            // 设置源IP地址，与sourceGroupId不能同时设置
			SourceIp:      "",
		},
	},
}
result, err := client.CreateSecurityGroup(args)
if err != nil {
    fmt.Println("create security group failed:", err)
} else {
    fmt.Println("create security group success: ", result)
}
```

> 同一安全组中的规则以remark、protocol、direction、portRange、sourceIp|destIp、sourceGroupId|destGroupId唯一性索引，重复记录报409错误。
>   protocol的取值（tcp|udp|icmp），默认值为空，代表all。
>   具体的创建安全组规则接口描述BCC API 文档[创建安全组](https://cloud.baidu.com/doc/BCC/s/0jwvynwij)。

## 删除安全组

以下代码可以删除指定的安全组:

```go
err := client.DeleteSecurityGroup(securityGroupId)
if err != nil {
    fmt.Println("delete security group failed:", err)
} else {
    fmt.Println("delete security group success")
}
```

## 授权安全组规则

使用以下代码可以在指定安全组中添加授权安全组规则:

```go
args := &api.AuthorizeSecurityGroupArgs{
	Rule: &api.SecurityGroupRuleModel{
		Remark:    "备注",
		Protocol:  "udp",
		PortRange: "1-65535",
		Direction: "ingress",
	},
}
err := client.AuthorizeSecurityGroupRule(securityGroupId, args)
if err != nil {
    fmt.Println("authorize security group new rule failed:", err)
} else {
    fmt.Println("authorize security group new rule success")
}
```

> -   同一安全组中的规则以remark、protocol、direction、portRange、sourceIp|destIp、sourceGroupId|destGroupId六元组作为唯一性索引，若安全组中已存在相同的规则将报409错误。
> -   具体的接口描述BCC API 文档[授权安全组规则](https://cloud.baidu.com/doc/BCC/s/pjwvynxvl)。

## 撤销安全组规则

使用以下代码可以在指定安全组中撤销指定安全组规则授权:

```go
args := &api.RevokeSecurityGroupArgs{
    Rule: &api.SecurityGroupRuleModel{
	    Remark:        "备注",
	    Protocol:      "udp",
	    PortRange:     "1-65535",
	    Direction:     "ingress",
	    SourceIp:      "",
	},
}
err := client.RevokeSecurityGroupRule(securityGroupId, args)
if err != nil {
    fmt.Println("revoke security group rule failed:", err)
} else {
    fmt.Println("revoke security group rule success")
}
```

> -   同一安全组中的规则以remark、protocol、direction、portRange、sourceIp|destIp、sourceGroupId|destGroupId六元组作为唯一性索引，若安全组中不存在对应的规则将报404错误。
> -   具体的接口描述BCC API 文档[撤销安全组规则](https://cloud.baidu.com/doc/BCC/s/yjwvynxk0)。

## 部署集
### 创建部署集

通过以下代码可以根据指定的部署集策略创建部署集

```go
// 设置创建部署集的名称
deploySetName := "your-deploy-set-name"
// 设置创建的部署集的描述信息
deployDesc := "your-deploy-set-desc"
// 设置创建部署集的策略，BBC实例策略只支持："tor_ha"
strategy := "tor_ha"
queryArgs := &CreateDeploySetArgs{
    Strategy:    strategy,
    Name:        deploySetName,
    Desc:        deployDesc,
}
if res, err := bccClient.CreateDeploySet(queryArgs); err != nil {
    fmt.Println("Create deploy set failed: ", err)
} else {
    fmt.Println("Create deploy set success, result: ", res)
}
```

### 查询部署集列表

通过以下代码可以查询所有部署集实例的列表及详情信息

```go
if res, err := bccClient.ListDeploySets(); err != nil {
    fmt.Println("List deploy sets failed: ", err)
} else {
    fmt.Println("List deploy sets success, result: ", res)
}
```
### 修改部署集属性

使用以下代码可以修改指定部署集的属性值

```go
    // 设置创建部署集的名称
	testDeploySetName := "testName"
    // 设置创建的部署集的描述信息
	testDeployDesc := "goDesc"
	queryArgs := &api.ModifyDeploySetArgs{
		Name: testDeploySetName,
		Desc: testDeployDesc,
	}
	BCC_TestDeploySetId = "DeploySetId"
	rep, err := BCC_CLIENT.ModifyDeploySet(BCC_TestDeploySetId, queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
```

### 删除指定的部署集

使用以下代码删除用户自己的指定的部署集

```go
// 设置你要删除的deploySetID
deploySetID := "your-choose-deploy-set-id"
if err := bccClient.DeleteDeploySet(deploySetID); err != nil {
    fmt.Println("Delete deploy set failed: ", err)
}
```

## 密钥对接口
### 创建密钥对

该接口用于创建密钥对，创建后的密钥对可植入实例以实现远程登录虚机。

```go
// 待创建的密钥对名称
name := "your-keypair-name"
// 待创建的密钥对的描述,可选参数
description := "your-keypair-desc"
args := &api.CreateKeypairArgs{
    Name:     name,
    Description:  description,
}
if res, err := bccClient.CreateKeypair(args); err != nil {
    fmt.Println("create keypair failed: ", err)
} else {
    fmt.Println("create keypair success,res: ", res)
}
```

### 导入密钥对

该接口用于用户自行导入创建密钥对。

```go
// 待创建的密钥对名称
name := "your-keypair-name"
// 待创建的密钥对的描述,可选参数
description := "your-keypair-desc"
publicKey := "your-publickey"
args := &api.ImportKeypairArgs{
    Name:     name,
    Description:  description,
    PublicKey:  publicKey,
}
if res, err := bccClient.ImportKeypair(args); err != nil {
    fmt.Println("create keypair failed: ", err)
} else {
    fmt.Println("create keypair success,res: ", res)
}
```

### 绑定密钥对

该接口用于将所选密钥对（限单个）绑定到所选虚机（支持多台）。 目前一台虚机只能绑定一个密钥对，若操作的虚机已经绑定密钥对，则此操作将替换该虚机原有的密钥对。此操作仅适用于linux系统的虚机，且所选虚机必须处于运行中或关机状态。

```go
instanceIds := []string{"your-instanceId"}
keypairId:= "your-keypair-id"
args := &api.AttackKeypairArgs{
    InstanceIds:     instanceIds,
    KeypairId:  keypairId,
}
if err := bccClient.AttachKeypair(args); err != nil {
    fmt.Println("attach keypair failed: ", err)
} else {
    fmt.Println("attach keypair success")
}
```

### 解绑密钥对

该接口用于将所选虚机与它们各自绑定的密钥对解绑。 目前一台虚机只能绑定一个密钥对，此操作将使所选虚机与其通过百度云控制台操作所绑定的密钥对解绑，若该虚机不存在这种密钥对，则不进行处理。 此操作仅适用于linux系统的虚机，且所选虚机必须处于运行中或关机状态。 
注： 1）用户所选虚机可能并未绑定任何密钥对，则此操作对该虚机无任何作用； 2）用户可能对所选虚机手动绑定了密钥对，则此操作对其手动绑定的密钥对不产生任何影响； 3）用户如若此前通过百度云控制台操作，为所选虚机绑定过密钥对，且该密钥对状态正常，则此操作将从该虚机中删除该密钥对
```go
instanceIds := []string{"your-instanceId"}
keypairId:= "your-keypair-id"
args := &api.DetachKeypairArgs{
    InstanceIds:     instanceIds,
    KeypairId:  keypairId,
}
if err := bccClient.DetachKeypair(args); err != nil {
    fmt.Println("detach keypair failed: ", err)
} else {
    fmt.Println("detach keypair success")
}
```

### 删除密钥对

该接口用于删除密钥对，已经绑定虚机的密钥对无法被删除。
```go
keypairId:= "your-keypair-id"
args := &api.DetachKeypairArgs{
    KeypairId:  keypairId,
}
if err := bccClient.DeleteKeypair(args); err != nil {
    fmt.Println("detach keypair failed: ", err)
} else {
    fmt.Println("detach keypair success")
}
```

### 查询密钥对详情

该接口用于查询单个密钥对的详细信息。
```go
keypairId:= "your-keypair-id"
if res, err := bccClient.GetKeypairDetail(keypairId); err != nil {
    fmt.Println("get keypair failed: ", err)
} else {
    fmt.Println("get keypair success,res: ", res)
}
```

### 查询密钥对列表

该接口用于查询密钥对列表。
```go
// 批量获取列表的查询的起始位置，是一个由系统生成的字符串,可选参数
marker := "your-marker"
// 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000,可选参数
maxKeys := your-maxKeys

args := &api.ListKeypairArgs{
    Marker     marker,
    MaxKeys    maxKeys,
}
if res, err := bccClient.ListKeypair(args); err != nil {
    fmt.Println("get keypair list failed: ", err)
} else {
    fmt.Println("get keypair list success,res: ", res)
}
```

### 重命名密钥对

该接口用于重命名密钥对。
```go
name := "your-keypair-name"
keypairId:= "your-keypair-id"
args := &api.RenameKeypairArgs{
    Name:     name,
    KeypairId:  keypairId,
}
if err := bccClient.RenameKeypair(args); err != nil {
    fmt.Println("update keypair name failed: ", err)
} else {
    fmt.Println("update keypair name  success")
}
```

### 更改密钥对描述

该接口用于更改密钥对描述，若用户提供的新的描述内容为空，则删除所操作虚机的描述。
```go
description := "your-keypair-desc"
keypairId:= "your-keypair-id"
args := &api.KeypairUpdateDescArgs{
    Description:  description,
    KeypairId:  keypairId,
}
if  err := bccClient.UpdateKeypairDescription(args); err != nil {
    fmt.Println("create keypair failed: ", err)
} else {
    fmt.Println("update keypair desc success")
}
```

## 其他接口
### 查询实例套餐规格

如下代码可以查询当前可以创建的实例的套餐的规格
```go
result, err := client.ListSpec()
if err != nil { 
    fmt.Println("list specs failed: ", err)
```

### 查询实例套餐规格(新)

通过以下代码可以查询实例资源套餐规格列表信息
```go
// 设置可用区名称，可选
zoneName := "your-zone-name"
args := &api.ListFlavorSpecArgs{
    ZoneName:   zoneName,
}
if res, err := bccClient.ListFlavorSpec(args); err != nil {
    fmt.Println("Get specified flavor list failed: ", err)
} else {
    fmt.Println("Get specified flavor list success, result: ", res)
}
```
>- 创建虚机时建议使用参数(spec)指定需要的机型以及配置。

### 查询实例套餐价格
通过以下代码可以查询实例资源套餐规格对应的价格
```go
// 设置实例规格族
specId := "your-spec-id"
// 设置实例套餐规格
spec := "your-spec"
// 设置付费方式，包括Postpaid(后付费)，Prepaid(预付费)两种
paymentTiming := "your-payment-timing"
// 设置可用区名称
zoneName := "your-zone-name"
// 设置所要购买的实例数量，缺省值为1，可选参数
purchaseCount := 1
// 设置时长，[1,2,3,4,5,6,7,8,9,12,24,36]，单位：月
reservationLength := 9
args := &api.GetPriceBySpecArgs{
    SpecId:         specId,
    Spec:           spec,
    PaymentTiming:  paymentTiming,
    ZoneName:       zoneName,
    PurchaseCount:  purchaseCount,
    PurchaseLength: reservationLength,
}
if res, err := bccClient.GetPriceBySpec(args); err != nil {
    fmt.Println("Get specified instance's price failed: ", err)
} else {
    fmt.Println("Get specified intstance's price success, result: ", res)
}
```

### 查询机型的可用区

使用以下代码可以查询指定机型支持的可用区列表
```go
args := &api.ListTypeZonesArgs{
    InstanceType: "",
    ProductType:  "",
    Spec:         "bcc.g3.c2m12",
    SpecId:       "",
}
if res, err := BCC_CLIENT.ListTypeZones(args); err != nil {
    fmt.Println("Get the specific zone flavor failed: ", err)
} else {
    fmt.Println("Get the specific zone flavor success, result: ", res)
}
```

### 查询可用区列表

如下代码可以所有的可用区的列表
```go
result, err := client.ListZone()
if err != nil {
    fmt.Println("list zone failed: ", err)
} else {
    fmt.Println("list zone success: ", result)
}
```

### 查询机型的可用区

使用以下代码可以查询指定机型支持的可用区列表
```go
args := &api.ListTypeZonesArgs{
    // 选择实例类型，可以选择N1, N2, N3等
    InstanceType: "",
    // 产品类型, 可以选择Prepaid，Postpaid
    ProductType:  "",
    // 实例套餐规格，可以选择bcc.g3.c2m12等
    Spec:         "",
    // 实例套餐规格族,可以选择g3,ic4等
    SpecId:       "",
}
if res, err := BCC_CLIENT.ListTypeZones(args); err != nil {
    fmt.Println("Get the specific zone flavor failed: ", err)
} else {
    fmt.Println("Get the specific zone flavor success, result: ", res)
}
```
### 查询实例套餐库存
查询实例资源套餐规格对应的库存。
```go
// 实例类型
instanceType := "instanceType"
// CPU核数
cpuCount := cpuCount
// 内存容量（GB）
memoryCapacityInGB := memoryCapacityInGB
// 可用区名称
zoneName := "zoneNamen"
// GPU卡类型，GPU和VGPU可填
gpuCard := "gpuCard"
// GPU卡数量，GPU和VGPU可填
cardCount := "cardCount"
//本地盘信息
ephemeralDisks := []EphemeralDisks{{
      "storageType": "ssd",
      "sizeInGB": sizeInGB,
}}

args := &api.CreateInstanceStockArgs{
    InstanceType:     instanceType,
    CpuCount: cpuCount,
    MemoryCapacityInGB:  memoryCapacityInGB,
    ZoneName:  zoneName,
    GpuCard:  gpuCard,
    CardCount:  cardCount,
    EphemeralDisks:  ephemeralDisks,
}
if res, err := bccClient.GetInstanceCreateStock(args); err != nil {
    fmt.Println("GetInstanceCreateStock failed: ", err)
} else {
    fmt.Println("GetInstanceCreateStock success: ", res)
}
```

  
### 实例扩缩容库存查询
实例变配余量查询
```go
// 实例id
instanceId := "instanceId"
// CPU核数
cpuCount := cpuCount
// 内存容量（GB）
memoryCapacityInGB := memoryCapacityInGB

args := &api.CreateInstanceStockArgs{
    InstanceId:     instanceId,
    CpuCount: cpuCount,
    MemoryCapacityInGB:  memoryCapacityInGB,
}
if res, err := bccClient.ResizeInstanceStockArgs(args); err != nil {
    fmt.Println("ResizeInstanceStockArgs failed: ", err)
} else {
    fmt.Println("ResizeInstanceStockArgs success: ", res)
}
```

### 查询实例绑定的弹性网卡列表

使用以下代码可以查询实例绑定的弹性网卡列表

```go
// 设置你要操作的instanceId
instanceId := "instanceId"
if res, err := BCC_CLIENT.ListInstanceEnis(instanceId); err != nil {
	fmt.Println("Get specific instance eni failed: ", err)
} else {
	fmt.Println("Get specific instance eni success, result: ", res)
}
```

# 错误处理

GO语言以error类型标识错误，BCC支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | BCC服务返回的错误

用户使用SDK调用BCC相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```
// bccClient 为已创建的BCC Client对象
instanceDetail, err := bccClient.GetInstanceDetail(instanceId)
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
	fmt.Println("get instance detail success: ", instanceDetail)
}
```

## 客户端异常

客户端异常表示客户端尝试向BCC发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当BCC服务端出现异常时，BCC服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[BCC错误返回](https://cloud.baidu.com/doc/BCC/s/Ojwvyo6nc)

# 版本变更记录

## v0.9.1 [2019-09-26]

首次发布：

 - 创建、查看、列表、启动、停止、重启、重装、删除实例，修改实例密码、安全组、属性、子网等，为实例续费
 - 创建、查看、列表、挂载、卸载、扩容、回滚、重命名、删除CDS磁盘，为磁盘续费，修改磁盘属性和计费方式等
 - 创建、查看、列表、删除、跨区域复制、取消跨区域复制、共享、取消共享镜像，删除自定义镜像，查询已共享的用户列表及根据实例ID查询OS信息等
 - 创建、查看、列表、删除快照
 - 创建、查看、列表、绑定、解绑、变更、删除自动快照策略
 - 创建、查询、删除安全组，为安全组授权和撤销规则
 - 查询实例套餐规格和查询可用区列表
