# CCE服务 v2版本

# 概述

本文档主要介绍CCE GO SDK的使用。在使用本文档前，您需要先了解CCE的一些基本知识，并已开通了CCE服务。若您还不了解CCE，可以参考[产品描述](https://cloud.baidu.com/doc/CCE/s/Bjwvy0x5g)和[操作指南](https://cloud.baidu.com/doc/CCE/s/zjxpoqohb)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[CCE服务域名](https://cloud.baidu.com/doc/CCE/s/Fjwvy1fl4)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/CCE/s/Fjwvy1fl4)。

目前支持“华北-北京”、“华南-广州”、“华东-苏州”、“香港”、“金融华中-武汉”和“华北-保定”六个区域。对应信息为：

访问区域 | 对应Endpoint | 协议
---|---|---
BJ | cce.bj.baidubce.com | HTTP and HTTPS
GZ | cce.gz.baidubce.com | HTTP and HTTPS
SU | cce.su.baidubce.com | HTTP and HTTPS
HKG| cce.hkg.baidubce.com| HTTP and HTTPS
FWH| cce.fwh.baidubce.com| HTTP and HTTPS
BD | cce.bd.baidubce.com | HTTP and HTTPS

## 获取密钥

要使用百度云CCE，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问CCE做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建CCE Client

CCE Client是CCE服务的客户端，为开发者与CCE服务进行交互提供了一系列的方法。

### 使用AK/SK新建CCE Client

通过AK/SK方式访问CCE，用户可以参考如下代码新建一个CCE Client：
```go
import (
	"github.com/baidubce/bce-sdk-go/services/ccev2"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := <your-access-key-id>, <your-secret-access-key>

	//用户指定的endpoint 
	ENDPOINT := "endpoint"
    
	// 初始化一个CCEClient
	ccev2Client, err := ccev2.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`AK`对应控制台中的“Access Key ID”，`SK`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb/)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为CCE的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`cce.bj.baidubce.com`。

### 使用STS创建CCE Client

**申请STS token**

CCE可以通过STS机制实现第三方的临时授权访问。STS（Security Token Service）是百度云提供的临时授权服务。通过STS，您可以为第三方用户颁发一个自定义时效和权限的访问凭证。第三方用户可以使用该访问凭证直接调用百度云的API或SDK访问百度云资源。

通过STS方式访问CCE，用户需要先通过STS的client申请一个认证字符串。

**用STS token新建CCE Client**

申请好STS后，可将STS Token配置到CCE Client中，从而实现通过STS Token创建CCE Client。

**代码示例**

GO SDK实现了STS服务的接口，用户可以参考如下完整代码，实现申请STS Token和创建CCE Client对象：
```go
import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/cce" //导入CCE服务模块
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

	// 使用申请的临时STS创建CCE服务的Client对象，Endpoint使用默认值
	ccev2Client, err := ccev2.NewClient(stsObj.AccessKeyId, stsObj.SecretAccessKey, "cce.bj.baidubce.com")
	if err != nil {
		fmt.Println("create cce client failed:", err)
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
	ccev2Client.Config.Credentials = stsCredential
}
```

> 注意：
> 目前使用STS配置CCE Client时，无论对应CCE服务的Endpoint在哪里，STS的Endpoint都需配置为http://sts.bj.baidubce.com。上述代码中创建STS对象时使用此默认值。

# 配置HTTPS协议访问CCE

CCE支持HTTPS传输协议，您可以通过在创建CCE Client对象时指定的Endpoint中指明HTTPS的方式，在CCE GO SDK中使用HTTPS访问CCE服务：
```go
// import "github.com/baidubce/bce-sdk-go/services/cce"
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "https://cce.bj.baidubce.com" //指明使用HTTPS协议

ccev2Client, _ := ccev2.NewClient(AK, SK, ENDPOINT)
```

## 配置CCE Client

如果用户需要配置CCE Client的一些细节的参数，可以在创建CCE Client对象之后，使用该对象的导出字段`Config`进行自定义配置，可以为客户端配置代理，最大连接数等参数。

### 使用代理

下面一段代码可以让客户端使用代理访问CCE服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

//创建CCE Client对象
AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cce.bj.baidubce.com"

ccev2Client, _ := ccev2.NewClient(AK, SK, ENDPOINT)

//代理使用本地的8080端口
ccev2Client.Config.ProxyUrl = "127.0.0.1:8080"
```

### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "cce.bj.baidubce.com"

ccev2Client, _ := ccev2.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
ccev2Client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
ccev2Client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/cce"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "ccev2.bj.baidubce.com"

ccev2Client, _ := ccev2.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
ccev2Client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
ccev2Client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问CCE时，创建的CCE Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者，后者为使用STS鉴权时使用，详见“使用STS创建CCE Client”小节。
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。


# CCE管理

百度智能云容器引擎(Cloud Container Engine，即CCE)是高度可扩展的高性能容器管理服务，您可以在托管的云服务器实例集群上轻松运行应用程序。

> 注意:
> - 百度智能云容器引擎免费为用户提供服务，只会对其使用的资源例如BCC、BLB和EIP等资源收费

## 创建集群
使用以下代码可以创建一个CCE Cluster。
```go
args := &ccev2.CreateClusterArgs{
    CreateClusterRequest: &ccev2.CreateClusterRequest{
        ClusterSpec: &types.ClusterSpec{
            ClusterName: "your-cluster-name",
            K8SVersion: types.K8S_1_16_8,
            RuntimeType: types.RuntimeTypeDocker,
            VPCID: "vpc-id",
            MasterConfig: types.MasterConfig {
                MasterType: types.MasterTypeManaged,
                ClusterHA: 1,
                ExposedPublic: false,
                ClusterBLBVPCSubnetID: "cluster-blb-vpc-subnet-id",
                ManagedClusterMasterOption: types.ManagedClusterMasterOption{
                    MasterVPCSubnetZone: types.AvailableZoneA,
                    MasterFlavor: types.MasterFlavorSmall,
                },
            },
            ContainerNetworkConfig: types.ContainerNetworkConfig{
                Mode: types.ContainerNetworkModeKubenet,
                LBServiceVPCSubnetID: "lb-service-vpc-subnet-id",
                ClusterPodCIDR: "172.28.0.0/16",
                ClusterIPServiceCIDR: "172.31.0.0/16",
            },
        },
        NodeSpecs: []*ccev2.InstanceSet{
            &ccev2.InstanceSet{
                Count: 1,
                InstanceSpec: types.InstanceSpec{
                    InstanceName: "",
                    ClusterRole: types.ClusterRoleNode,
                    Existed: false,
                    MachineType: types.MachineTypeBCC,
                    InstanceType: bccapi.InstanceTypeN3,
                    VPCConfig: types.VPCConfig{
                        VPCID: "vpc-id",
                        VPCSubnetID: "vpc-subnet-id",
                        SecurityGroupID: "security-group-id",
                        AvailableZone: types.AvailableZoneA,
                    },
                    InstanceResource:types.InstanceResource{
                        CPU: 4,
                        MEM: 8,
                        RootDiskSize: 40,
                        LocalDiskSize: 0,
                        CDSList: []types.CDSConfig{},
                    },
                    ImageID: "image-id",
                    InstanceOS: types.InstanceOS{
                        ImageType: bccapi.ImageTypeSystem,
                    },
                    NeedEIP: false,
                    AdminPassword: "admin-password",
                    SSHKeyID: "ssh-key-id",
                    InstanceChargingType: bccapi.PaymentTimingPostPaid,
                    RuntimeType: types.RuntimeTypeDocker,
                },
            },
        },
    },
}

resp, err := ccev2Client.CreateCluster(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 查询集群详情
使用以下代码可以查询一个集群的详细信息。
```go
clusterID := "cluster-id"
resp, err := ccev2Client.GetCluster(clusterID)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 获取集群列表
使用以下代码可以按关键字获取排序后的集群列表
```go
args := &ccev2.ListClustersArgs{
    KeywordType: "clusterName",
    Keyword: "",
    OrderBy: "clusterID",
    Order: ccev2.OrderASC,
    PageSize: 10,
    PageNum: 1,
}

resp, err := ccev2Client.ListClusters(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 删除集群
使用以下代码可以删除一个集群
```go
args := &ccev2.DeleteClusterArgs{
    ClusterID: "cluster-id",
    DeleteResource: true,
    DeleteCDSSnapshot: true,
}

resp, err := ccev2Client.DeleteCluster(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 创建节点（集群扩容）
使用以下代码可以新创建节点并对集群进行扩容
```go
args := &ccev2.CreateInstancesArgs{
    ClusterID: "cluster-id",
    Instances: []*ccev2.InstanceSet{
        {
            Count: 1,
            InstanceSpec: types.InstanceSpec{
                ClusterRole: types.ClusterRoleNode,
                Existed: false,
                MachineType: types.MachineTypeBCC,
                InstanceType: bccapi.InstanceTypeN3,
                VPCConfig: types.VPCConfig{
                    VPCID: "vpc-id",
                    VPCSubnetID: "vpc-subnet-id",
                    SecurityGroupID: "security-group-id",
                    AvailableZone: types.AvailableZoneA,
                },
                InstanceResource:types.InstanceResource{
                    CPU: 1,
                    MEM: 4,
                    RootDiskSize: 40,
                    LocalDiskSize: 0,
                },
                ImageID: "image-id",
                InstanceOS: types.InstanceOS{
                    ImageType: bccapi.ImageTypeSystem,
                },
                NeedEIP: false,
                AdminPassword: "admin-password",
                SSHKeyID: "ssh-key-id",
                InstanceChargingType: bccapi.PaymentTimingPostPaid,
                RuntimeType: types.RuntimeTypeDocker,
            },
        },
    },
}

resp, err := ccev2Client.CreateInstances(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```


## 获取节点详情
使用以下代码可以获取一个节点的详细信息
```go
args := &ccev2.GetInstanceArgs{
    ClusterID: "cluster-id",
    InstanceID: "instance-id",
}

resp, err := ccev2Client.GetInstance(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 获取集群的节点列表
使用以下代码可以获取一个集群的节点列表
```go
args :=  &ccev2.ListInstancesByPageArgs {
    ClusterID: "cluster-id",
    Params: &ccev2.ListInstancesByPageParams{
        KeywordType: ccev2.InstanceKeywordTypeInstanceName,
        Keyword: "",
        OrderBy: "createdAt",
        Order: ccev2.OrderASC,
        PageNo: 1,
        PageSize: 10,
    },
}

resp, err := ccev2Client.ListInstancesByPage(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 获取节点组的节点列表
使用以下代码可以获取节点组的节点列表
```go
args := &ListInstanceByInstanceGroupIDArgs{
    ClusterID: "your-cluster-id",
    InstanceGroupID: "your-instance-group-id",
    PageSize: 0,
    {ageNo: 0,
}
resp, err := ccev2Client.ListInstancesByInstanceGroupID(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 更新节点配置
使用以下代码可以更新节点的配置信息。需要注意不是所有配置信息都是可更改的。
```go
args := &UpdateInstanceArgs{
	ClusterID:  "your-cluster-id",
	InstanceID: "your-instance-id",
	InstanceSpec: YOUR_NEW_INSTANCE_SPEC,
}
respUpdate, err := CCE_CLIENT.UpdateInstance(args)

s, _ := json.MarshalIndent(respUpdate, "", "\t")
fmt.Println("Response:" + string(s))
```

## 删除节点(集群缩容)
使用以下代码可以删除集群内的一个节点
```go
args := &ccev2.DeleteInstancesArgs{
    ClusterID:   "cluster-id",
    DeleteInstancesRequest: &ccev2.DeleteInstancesRequest{
        InstanceIDs: []string{ "instance-id" },
        DeleteOption: &types.DeleteOption{
            MoveOut: false,
            DeleteCDSSnapshot: true,
            DeleteResource: true,
        },
    },
}

resp, err := ccev2Client.DeleteInstances(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 检查集群网络网段
使用以下代码可以检查集群网络网段是否冲突
```go
args := &ccev2.CheckClusterIPCIDArgs{
    VPCID: "vpc-id",
    VPCCIDR: "192.168.0.0/16",
    ClusterIPCIDR: "172.31.0.0/16",
    IPVersion: "ipv4",
}

resp, err := ccev2Client.CheckClusterIPCIDR(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 检查容器网络网段
使用以下代码可以检测容器网络网段是否冲突
```go
args := &ccev2.CheckContainerNetworkCIDRArgs{
    VPCID: "vpc-id",
    VPCCIDR: "192.168.0.0/16",
    ContainerCIDR: "172.28.0.0/16",
    ClusterIPCIDR: "172.31.0.0/16",
    MaxPodsPerNode: 256,
}

resp, err := ccev2Client.CheckContainerNetworkCIDR(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 推荐集群网络网段
使用以下代码可以推荐可使用的集群网络网段
```go
args := &ccev2.RecommendClusterIPCIDRArgs{
    ClusterMaxServiceNum: 2,
    ContainerCIDR: "172.28.0.0/16",
    IPVersion: "ipv4",
    PrivateNetCIDRs: []ccev2.PrivateNetString{ccev2.PrivateIPv4Net172},
    VPCCIDR: "192.168.0.0/16",
}

resp, err := ccev2Client.RecommendClusterIPCIDR(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 推荐容器网络网段
使用以下代码可以推荐可使用的容器网络网段
```go
args := &ccev2.RecommendContainerCIDRArgs{
    ClusterMaxNodeNum: 2,
    IPVersion: "ipv4",
    K8SVersion: types.K8S_1_16_8,
    MaxPodsPerNode: 32,
    PrivateNetCIDRs: []ccev2.PrivateNetString{ccev2.PrivateIPv4Net172},
    VPCCIDR: "192.168.0.0/16",
    VPCID: "vpc-id",
}

resp, err := ccev2Client.RecommendContainerCIDR(args)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 获取集群配额
使用以下代码可以获取集群配额
```go
resp, err := ccev2Client.GetClusterQuota()
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 获取集群节点配额
使用以下代码可以获取集群的节点配额
```go	
clusterID := "cluster-id"

resp, err := ccev2Client.GetClusterNodeQuota(clusterID)
if err != nil {
    fmt.Println(err.Error())
    return
}

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:"+ string(s))
```

## 创建节点组
使用以下代码可以创建节点组
```go	
args := &CreateInstanceGroupArgs{
    ClusterID: CCE_CLUSTER_ID,
    Request: &CreateInstanceGroupRequest{
        types.InstanceGroupSpec{
            InstanceGroupName: "your-instance-group-name",
            CleanPolicy: types.DeleteCleanPolicy,
            Replicas: 3,
            InstanceTemplate: types.InstanceTemplate{
                InstanceSpec: types.InstanceSpec{
                    ClusterRole:  types.ClusterRoleNode,
                    Existed:      false,
                    MachineType:  types.MachineTypeBCC,
                    InstanceType: bccapi.InstanceTypeN3,
                    VPCConfig: types.VPCConfig{
                        VPCID:           "your-vpc-id",
                        VPCSubnetID:     "your-vpc-subnet-id",
                        SecurityGroupID: "your-secuirity-group-id",
                        AvailableZone:   types.AvailableZoneA,
                    },
                    DeployCustomConfig: types.DeployCustomConfig{
                        PreUserScript: "your-script",
                        PostUserScript:"your-script",
                    },
                    InstanceResource: types.InstanceResource{
                        CPU:           1,
                        MEM:           4,
                        RootDiskSize:  40,
                        LocalDiskSize: 0,
                    },
                    ImageID: IMAGE_TEST_ID,
                    InstanceOS: types.InstanceOS{
                        ImageType: bccapi.ImageTypeSystem,
                    },
                    NeedEIP:              false,
                    AdminPassword:        "your-admin-password",
                    SSHKeyID:             "your-ssh-key-id",
                    InstanceChargingType: bccapi.PaymentTimingPostPaid,
                    RuntimeType:          types.RuntimeTypeDocker,
                },
            },
        },
    },
}
resp, err := ccev2Client.CreateInstanceGroup(args)

s, _ = json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```


## 获取节点组列表
使用以下代码可以获取节点组列表
```go	
args := &ListInstanceGroupsArgs{
    ClusterID: "your-cluster-id",
    ListOption: &InstanceGroupListOption{
        PageNo: 0,
        PageSize: 0,
    },
}
resp, err := ccev2Client.ListInstanceGroups(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```


## 查询节点组详情
使用以下代码可以查询节点组详情
```go	
args := &GetInstanceGroupArgs{
    ClusterID: "your-cluster-id",
    InstanceGroupID: "your-instance-group-id",
}
resp, err := ccev2Client.GetInstanceGroup(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```


## 修改节点组内节点副本数
使用以下代码可以修改节点组内节点副本数
```go	
args := &UpdateInstanceGroupReplicasArgs{
    ClusterID: "your-cluster-id",
    InstanceGroupID: "your-instance-group-id",
    Request:  &UpdateInstanceGroupReplicasRequest{
        Replicas: 1,
        DeleteInstance: true,
        DeleteOption: &types.DeleteOption{
            MoveOut: false,
            DeleteCDSSnapshot: true,
            DeleteResource: true,
        },
    },
}
resp, err := ccev2Client.UpdateInstanceGroupReplicas(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 修改节点组Autoscaler配置
使用以下代码可以修改节点Autoscaler配置
```go
args := &UpdateInstanceGroupClusterAutoscalerSpecArgs{
    ClusterID: "your-cluster0id",
    InstanceGroupID: "cce-instance-group-id",
    Request: &ClusterAutoscalerSpec{
        Enabled: true,
        MinReplicas: 2,
        MaxReplicas: 5,
        ScalingGroupPriority: 1,
    },
}

resp, err := ccev2Client.UpdateInstanceGroupClusterAutoscalerSpec(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 删除节点组
使用以下代码可以删除节点组
```go	
args := &DeleteInstanceGroupArgs{
    ClusterID: "your-cluster-id",
    InstanceGroupID: "your-instance-group-id",
    DeleteInstances: true,
}
resp, err := ccev2Client.DeleteInstanceGroup(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 创建Autoscaler
使用以下代码为集群创建Autoscaler
```go	
args := &CreateAutoscalerArgs{
	ClusterID: "your-cce-cluster-id",
}

resp, err := ccev2Client.CreateAutoscaler(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 查询Autoscaler
使用以下代码可以查询集群的Autoscaler
```go	
args := &GetAutoscalerArgs{
	ClusterID: "your-cce-cluster-id",
}

resp, err := ccev2Client.GetAutoscaler(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 更新Autoscaler
使用以下代码可以更新集群的Autoscaler
```go	
args := &UpdateAutoscalerArgs{
	ClusterID: "your-cluster-id",
	AutoscalerConfig: ClusterAutoscalerConfig{
		ReplicaCount: 5,
		ScaleDownEnabled: true,
		Expander: "random",
	},
}

resp, err := ccev2Client.UpdateAutoscaler(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

## 获取Kubeconfig
使用以下代码可以获取集群的Kubeconfig
```go	
args := &GetKubeConfigArgs{
	ClusterID: "your-cluster-id",
	KubeConfigType: "kubeconfig-type-you-need",
}

resp, err := ccev2Client.GetAdminKubeConfig(args)

s, _ := json.MarshalIndent(resp, "", "\t")
fmt.Println("Response:" + string(s))
```

# 错误处理

GO语言以error类型标识错误，CCE支持两种错误见下表：

错误类型        |  说明
----------------|-------------------
BceClientError  | 用户操作产生的错误
BceServiceError | CCE服务返回的错误

用户使用SDK调用CCE相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```go
// client 为已创建的cce Client对象
args := &ccev2.ListClusterArgs{}
result, err := ccev2Client.ListClusters(args)
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

客户端异常表示客户端尝试向CCE发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当CCE服务端出现异常时，CCE服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[CCE错误码](https://cloud.baidu.com/doc/CCE/s/4jwvy1evj)

## SDK日志

CCE GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

CCE GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

```go
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

```go
// 直接使用包级别全局日志对象（会和GO SDK自身日志一并输出）
log.SetLogHandler(log.STDERR)
log.Debugf("%s", "logging message using the log package in the CCE go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the CCE go sdk")
```


# 版本变更记录

## v1.0.0 [2020-08-07]

首次发布:
 - 支持创建集群、获取集群列表、获取集群详情、删除集群。
 - 支持创建节点（集群扩容）、获取集群的节点列表、获取节点详情、删除节点（集群缩容）
 - 支持检查集群网络网段、检查容器网络网段、推荐集群网络网段、推荐容器网络网段
 - 支持查询集群配额、集群节点配额
 
 
## v1.1.0 [2020-08-20]
 
 增加节点组相关接口:
  - 支持节点组创建、获取节点组列表、查询节点组详情、修改节点组内节点副本数、删除节点组
  - 获取节点组的节点列表
  
  
## v1.2.0 [2020-09-18]
   
增加Autoscaler相关接口:
  - 支持集群Autoscaler的初始化、查询与更新
增加Kubeconfig相关接口:
  - 支持获取集群的kubeconfig
增加Instance 相关接口：
  - 支持对Instance部分属性的更新