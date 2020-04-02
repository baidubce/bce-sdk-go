# BBC服务

# 概述

本文档主要介绍BBC GO SDK的使用。在使用本文档前，您需要先了解BBC的一些基本知识。若您还不了解BBC，可以
参考[产品描述](https://cloud.baidu.com/doc/BBC/s/ojwvxu4di)和
[操作指南](https://cloud.baidu.com/doc/BBC/s/fjwvxu86k)。

# 初始化

## 确认Endpoint
在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于
[BBC访问域名](https://cloud.baidu.com/doc/BBC/s/3jwvxu9iz#%E6%9C%8D%E5%8A%A1%E5%9F%9F%E5%90%8D)的部分，
理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

## 获取密钥

要使用百度云BBC，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。
AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BBC做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建BBC Client

BBC Client是BBC服务的客户端，为开发者与BBC服务进行交互提供了一系列的方法。

### 使用AK/SK新建BBC Client

通过AK/SK方式访问BBC，用户可以参考如下代码新建一个Bbc Client：

```go
import "github.com/baidubce/bce-sdk-go/services/bbc"

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BBClient
	bbcClient, err := bbc.NewClient(AK, SK, ENDPOINT)
}
```
在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，
获取方式请参考[获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb)。
第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BBC的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`bbc.bj.baidubce.com`。

## 配置HTTPS协议访问BBC

BBC支持HTTPS传输协议，您可以通过在创建BBC Client对象时指定的Endpoint中指明HTTPS的方式，在BBC GO SDK中使用HTTPS访问BBC服务：

```go
import "github.com/baidubce/bce-sdk-go/services/bbc"

ENDPOINT := "https://bbc.bj.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
bbcClient, _ := bbc.NewClient(AK, SK, ENDPOINT)
```
### 设置网络参数

用户可以通过如下的示例代码进行网络参数的设置：

```go
import "github.com/baidubce/bce-sdk-go/services/bbc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bbc.bj.baidubce.com"
client, _ := bbc.NewClient(AK, SK, ENDPOINT)

// 配置不进行重试，默认为Back Off重试
client.Config.Retry = bce.NewNoRetryPolicy()

// 配置连接超时时间为30秒
client.Config.ConnectionTimeoutInMillis = 30 * 1000
```

### 配置生成签名字符串选项

```go
// import "github.com/baidubce/bce-sdk-go/services/bbc"

AK, SK := <your-access-key-id>, <your-secret-access-key>
ENDPOINT := "bbc.bj.baidubce.com"
client, _ := bbc.NewClient(AK, SK, ENDPOINT)

// 配置签名使用的HTTP请求头为`Host`
headersToSign := map[string]struct{}{"Host": struct{}{}}
client.Config.SignOption.HeadersToSign = HeadersToSign

// 配置签名的有效期为30秒
client.Config.SignOption.ExpireSeconds = 30
```

**参数说明**

用户使用GO SDK访问BBC时，创建的BBC Client对象的`Config`字段支持的所有参数如下表所示：

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

  1. `Credentials`字段使用`auth.NewBceCredentials`与`auth.NewSessionBceCredentials`函数创建，默认使用前者. 后者为使用STS鉴权时使用
  2. `SignOption`字段为生成签名字符串时的选项，详见下表说明：

名称          | 类型  | 含义
--------------|-------|-----------
HeadersToSign |map[string]struct{} | 生成签名字符串时使用的HTTP头
Timestamp     | int64 | 生成的签名字符串中使用的时间戳，默认使用请求发送时的值
ExpireSeconds | int   | 签名字符串的有效期

     其中，HeadersToSign默认为`Host`，`Content-Type`，`Content-Length`，`Content-MD5`；TimeStamp一般为零值，表示使用调用生成认证字符串时的时间戳，用户一般不应该明确指定该字段的值；ExpireSeconds默认为1800秒即30分钟。
  3. `Retry`字段指定重试策略，目前支持两种：`NoRetryPolicy`和`BackOffRetryPolicy`。默认使用后者，该重试策略是指定最大重试次数、最长重试时间和重试基数，按照重试基数乘以2的指数级增长的方式进行重试，直到达到最大重试测试或者最长重试时间为止。

# 主要接口

云服务器BCC（Baidu Cloud Compute）是处理能力可弹性伸缩的计算服务。
管理方式比物理服务器更简单高效，可根据您的业务需要创建、释放任意多台云服务器实例，提升运维效率。

## 实例
   
### 创建实例

使用以下代码可以创建一个物理机实例：
```go
createInstanceArgs := &CreateInstanceArgs{
    // 输入你选择的flavor（套餐）ID，通过SDK获取可用flavor id的方法详见套餐章节
    FlavorId:         "your-choose-flavor-id",
    // 输入你要创建instance使用的镜像ID，通过SDK获取可用镜像ID的方法详见镜像章节
    ImageId:          "your-choose-image-id",
    // 输入你要创建instance使用的raid ID，通过SDK获取可用raid id的方法详见套餐章节
    RaidId:           "your-choose-raid-id",
    // 输入待创建物理磁盘的大小，单位为GB，缺省为20
    RootDiskSizeInGb: 20,
    // 批量创建（购买）的虚拟机实例个数，必须为大于0的整数，可选参数，缺省为1
    PurchaseCount:    1,
    // 可用区，格式为：国家-区域-可用区，如'中国-北京-可用区A'就是'cn-bj-a'
    ZoneName:         "cn-bj-a",
    // 指定子网 ID，必填参数
    SubnetId:         "your-choose-subnet-id",
    // 指定安全组id，可选参数
    SecurityGroupId:  "your-choose-security-group-id"
    // 使用 uuid 生成一个长度不超过64位的ASCII字符串
    ClientToken:      "random-uuid",
    // 选择付费方式
    Billing: Billing{
        PaymentTiming: PaymentTimingPostPaid,
        Reservation: Reservation{
            Length: 1,
            TimeUnit: "Month",
        },
    },
    // 指定使用的部署集id，可选参数，通过SDK获取可用部署集id的方法详见部署集章节
    DeploySetId: "your-choose-raid-id",
    // 设置实例管理员密码(8-16位字符，英文，数字和符号必须同时存在，符号仅限!@#$%^*())
    AdminPass:   "your-admin-pass",
    // 实例名称
    Name:        "your-choose-instance-name",
}
if res, err := bbcClient.CreateInstance(createInstanceArgs); err != nil {
    fmt.Println("create instance failed: ", err)
} else {
    fmt.Println("create instance success, instanceId: ", res.InstanceIds[0])
}
```

> **注意：**
>付费方式(PaymentTiming)可选：
>- 后付费: PaymentTimingPostPaid
>- 预付费: PaymentTimingPrePaid

### 查询实例列表
使用以下代码查询所有BBC实例的列表及详情信息：
```go
listArgs := &ListInstancesArgs{
    // 批量获取列表的查询起始位置，是一个由系统产生的字符串
    Marker: "your-marker",
    // 设置返回数据大小，缺省为1000
    MaxKeys: 100,
    // 通过internal Ip过滤BBC列表
    InternalIp: "your-choose-internal-ip",
}
if res, err := bbcClient.ListInstances(listArgs); err != nil {
    fmt.Println("list instances failed: ", err)
} else {
    fmt.Println("list instances success, result: ", res)
}
```
### 查询实例详情
使用以下代码可以查询指定BBC实例的详细信息：
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
if res, err := bbcClient.GetInstanceDetail(instanceId); err != nil {
    fmt.Println("get instance detail failed: ", err)
} else {
    fmt.Println("get instance detail success, result: ", res)
}
```

### 启动实例
使用以下代码可以启动指定BBC实例，实例状态必须为 Stopped，调用此接口才可以成功返回，否则提示409错误：
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
if err := bbcClient.StartInstance(instanceId); err != nil {
    fmt.Println("start instance failed: ", err)
} else {
    fmt.Println("start instance success.")
}
```

### 停止实例
使用以下代码可以停止指定BBC实例，只有状态为 Running 的实例才可以进行此操作，否则提示 409 错误：
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
// 是否强制停止实例，为True代表强制停止
forceStop := true
if err := bbcClient.StopInstance(instanceId, forceStop); err != nil {
    fmt.Println("stop instance failed: ", err)
} else {
    fmt.Println("stop instance success.")
}
```

### 重启实例
使用以下代码可以重启指定BBC实例，只有状态为 Running 的实例才可以进行此操作，否则提示 409 错误：
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
// 是否强制停止实例，为True代表强制停止
forceStop := true
if err := bbcClient.RebootInstance(instanceId, forceStop); err != nil {
    fmt.Println("reboot instance failed: ", err)
} else {
    fmt.Println("reboot instance success.")
}
```

### 修改实例名称
使用以下代码可以修改指定BBC实例的名称：
```go
modifyInstanceNameArgs := &ModifyInstanceNameArgs{
    Name: "new_bbc_name",
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
if err := bbcClient.ModifyInstanceName(instanceId, modifyInstanceNameArgs); err != nil {
    fmt.Println("modify instance name failed: ", err)
} else {
    fmt.Println("modify instance name success.")
}
```

### 修改实例描述
使用以下代码可以修改指定BBC实例的描述：
```go
modifyInstanceDescArgs := &ModifyInstanceDescArgs{
    Description: "new_bbc_description",
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
if err := bbcClient.ModifyInstanceDesc(instanceId, modifyInstanceDescArgs); err != nil {
    fmt.Println("modify instance desc failed: ", err)
} else {
    fmt.Println("modify instance desc success.")
}
```

### 重装实例
使用以下代码可以使用镜像重建指定BBC实例:
```go
rebuildArgs := &RebuildInstanceArgs{
    // 设置使用的镜像id
    ImageId:        "your-choose-image-id",
    // 设置管理员密码
    AdminPass:      "your-new-admin-pass",
    // 是否保留数据。当该值为true时，raidId和sysRootSize字段不生效
    IsPreserveData: false,
    // 此参数在isPreserveData为false时为必填，在isPreserveData为true时不生效
    RaidId:         "your_raid_id",
    // 系统盘根分区大小，默认为20G，取值范围为20-100。此参数在isPreserveData为true时不生效
    SysRootSize: 20,
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
// 设置是否保留数据
isPreserveData = false
if err := bbcClient.RebuildInstance(instanceId, isPreserveData, rebuildArgs); err != nil {
    fmt.Println("rebuild instance failed: ", err)
} else {
    fmt.Println("rebuild instance success.")
}
```
> **注意：**
>IsPreserveData表示是否保留数据：
>- 当IsPreserveData设置为 false 时，RaidId 和 SysRootSize 是必填参数
>- 当IsPreserveData设置为 true 时，RaidId 和 SysRootSize 参数不生效

### 释放实例
对于后付费Postpaid以及预付费Prepaid过期的BBC实例，可以使用以下代码将其释放:
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
if err := bbcClient.DeleteInstance(instanceId); err != nil {
    fmt.Println("release instance failed: ", err)
} else {
    fmt.Println("release instance success.")
}
```

### 修改实例密码
使用以下代码可以修改指定BBC实例的管理员密码：
```go
modifyInstancePasswordArgs := &ModifyInstancePasswordArgs{
    AdminPass: "your_new_password",
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
if err := bbcClient.ModifyInstancePassword(instanceId, modifyInstancePasswordArgs); err != nil {
    fmt.Println("modify instance password failed: ", err)
} else {
    fmt.Println("modify instance password success.")
}
```

> **注意：**
>BBC 实例密码要求：
>- 8-16位字符，英文，数字和符号必须同时存在，符号仅限!@#$%^*()

### 查询实例VPC/Subnet信息
使用以下代码可以通过BBC实例id查询VPC/Subnet信息：
```go
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
getVpcSubnetArgs := &GetVpcSubnetArgs{
    BbcIds: []string{instanceId},
}
if res, err := bbcClient.GetVpcSubnet(getVpcSubnetArgs); err != nil {
    fmt.Println("get vpc subnet failed: ", err)
} else {
    fmt.Println("get vpc subnet success. res: ", res)
}
```

### 向指定实例批量添加指定ip
```go
privateIps := []string{"192.168.1.25"}
instanceId := "your-choose-instance-id"
batchAddIpArgs := &BatchAddIpArgs{
	InstanceId: instanceId,
	PrivateIps: privateIps,
}
if err := bbcClient.BatchAddIP(batchAddIpArgs); err != nil {
    fmt.Println("add ips failed: ", err)
} else {
    fmt.Println("add ips success.")
}
```

### 批量删除指定实例的ip
```go
privateIps := []string{"192.168.1.25"}
instanceId := "your-choose-instance-id"
batchDelIpArgs := &BatchDelIpArgs{
	InstanceId: instanceId,
	PrivateIps: privateIps,
}
if err := bbcClient.BatchDelIP(batchDelIpArgs); err != nil {
    fmt.Println("delete ips failed: ", err)
} else {
    fmt.Println("delete ips success.")
}
```

## 标签
### 实例解绑标签
通过以下代码解绑实例已有的标签
```go
unbindTagsArgs := &UnbindTagsArgs{
    // 设置您要解绑的标签
    ChangeTags: []model.TagModel{
        {
            TagKey:   "tag1",
            TagValue: "var1",
        },
    },
}
// 设置你要操作的instanceId
instanceId := "your-choose-instance-id"
if err := BBC_CLIENT.UnbindTags(instanceId, unbindTagsArgs); err != nil {
    fmt.Println("unbind instance tags failed: ", err)
} else {
    fmt.Println("unbind instance tags success.")
}
```

## 套餐
### 查询套餐列表
使用以下代码查询所有BBC套餐的列表及详情信息
```go
if res, err := bbcClient.ListFlavors(); err != nil {
    fmt.Println("List flavors failed: ", err)
} else {
    fmt.Println("List flavors success, result: ", res)
}
```

### 查询套餐详情
使用以下代码可以查询指定套餐的详细信息

```go
// 设置你要操作的flavorId
flavorId := "your-choose-flavor-id"
if res, err := bbcClient.GetFlavorDetail(testFlavorId); err != nil {
    fmt.Println("Get flavor failed: ", err)
} else {
    fmt.Println("Get flavor success, result: ", res)
}
```
### 查询RAID详情
使用以下代码可以查询指定套餐的RAID方式及磁盘大小

```go
// 设置你要操作的flavorId
flavorId := "your-choose-flavor-id"
if res, err := bbcClient.GetFlavorRaid(testFlavorId); err != nil {
    fmt.Println("Get raid failed: ", err)
} else {
    fmt.Println("Get raid success, result: ", res)
}
```

## 镜像
### 通过实例创建自定义镜像
- 用于创建自定义镜像，默认每个账号配额20个，创建后的镜像可用于创建实例
- 只有 Running 或 Stopped 状态的实例才可以执行成功
使用以下代码可以从指定的实例创建镜像
```go
// 用于创建镜像的实例ID
instanceId := "i-3EavdPl8"
// 设置创建镜像的名称
imageName := "testCreateImage"
queryArgs := &CreateImageArgs{
    ImageName:  testImageName,
    InstanceId: testInstanceId,
}
if res, err := bbcClient.CreateImageFromInstanceId(queryArgs); err != nil {
    fmt.Println("Create image failed: ", err)
} else {
    fmt.Println("Create image success, result: ", res)
}
```
### 查询镜像列表
- 用于查询用户所有的镜像信息
- 查询的镜像信息中包括系统镜像、自定义镜像和服务集成镜像
- 支持按 imageType 来过滤查询，此参数非必需，未设置时默认为 All,即查询所有类型的镜像
使用以下代码可以查询镜像列表
```go
// 指定要查询何种类型的镜像
// All(所有)
// System(系统镜像/公共镜像)
// Custom(自定义镜像)
// Integration(服务集成镜像)
// Sharing(共享镜像)
imageType := "All"
// 批量获取列表的查询的起始位置 
marker := "your-marker"
// 每页包含的最大数量
maxKeys := 100
queryArgs := &ListImageArgs{
    Marker:    marker, 
    MaxKeys:   maxKeys,
    ImageType: imageType,
}
if res, err := bbcClient.ListImage(queryArgs); err != nil {
    fmt.Println("List image failed: ", err)
} else {
    fmt.Println("List image success, result: ", res)
}
```
### 查询镜像详情
- 用于根据指定镜像ID查询单个镜像的详细信息
使用以下代码可以查询镜像详情
```go
// 待查询镜像ID
image_id :="your-choose-image-id"
if res, err := bbcClient.GetImageDetail(testImageId); err != nil {
    fmt.Println("Get image failed: ", err)
} else {
    fmt.Println("Get image success, result: ", res)
}
```
### 删除自定义镜像
- 用于删除用户自己的指定的自定义镜像，仅限自定义镜像，系统镜像和服务集成镜像不能删除
- 镜像删除后无法恢复，不能再用于创建、重置实例
使用以下代码可以删除指定镜像

```go
// 待删除镜像ID
imageId := "your-choose-image-id"
if err := bbcClient.DeleteImage(testImageId); err != nil {
    fmt.Println("Delete image failed: ", err)
}
```

## 操作日志
### 查询操作日志
通过以下代码查询指定操作日志

```go
// 批量获取列表的查询的起始位置，是一个由系统生成的字符串
marker := "your-marker"
// 每页包含的最大数量，最大数量通常不超过1000。缺省值为100
maxKeys := 100
// 需查询物理机操作的起始时间（UTC时间），格式 yyyy-MM-dd'T'HH:mm:ss'Z' ，为空则查询当日操作日志
startTime := ""
// 需查询物理机操作的终止时间（UTC时间），格式 yyyy-MM-dd'T'HH:mm:ss'Z' ，为空则查询当日操作日志
endTime := ""
queryArgs := &GetOperationLogArgs{
    Marker:    marker,
    MaxKeys:   maxKeys,
    StartTime: startTime,
    EndTime:   endTime,
}
if res, err := bbcClient.ListImage(queryArgs); err != nil {
    fmt.Println("Get Operation Log failed: ", err)
} else {
    fmt.Println("Get Operation Log success, result: ", res)
}
```

## 部署集
### 创建部署集
通过以下代码根据指定的部署集策略和并发度创建部署集
```go
// 设置创建部署集的名称
deploySetName := "your-deploy-set-name"
// 设置创建的部署集的描述信息
deployDesc := "your-deploy-set-desc"
// 设置部署集并发度，范围 [1,5]
concurrency := 1
// 设置创建部署集的策略，BBC实例策略只支持："tor_ha"
strategy := "tor_ha"
queryArgs := &CreateDeploySetArgs{
    Strategy:    strategy,
    Concurrency: concurrency,
    Name:        deploySetName,
    Desc:        deployDesc,
}
if res, err := bbcClient.CreateDeploySet(queryArgs); err != nil {
    fmt.Println("Create deploy set failed: ", err)
} else {
    fmt.Println("Create deploy set success, result: ", res)
}
```

### 查询部署集列表
使用以下代码查询所有部署集实例的列表及详情信息

```go
if res, err := bbcClient.ListDeploySets(); err != nil {
    fmt.Println("List deploy sets failed: ", err)
} else {
    fmt.Println("List deploy sets success, result: ", res)
}
```
### 查询部署集详情
使用以下代码可以查询指定套餐的详细信息

```go
// 设置你要查询的deploySetID
deploySetID := "your-choose-deploy-set-id"
if res, err := bbcClient.GetDeploySet(deploySetID); err != nil {
    fmt.Println("Get deploy set failed: ", err)
} else {
    fmt.Println("Get deploy set success, result: ", res)
}
```
### 删除指定的部署集
使用以下代码删除用户自己的指定的部署集

```go
// 设置你要删除的deploySetID
deploySetID := "your-choose-deploy-set-id"
if err := bbcClient.DeleteDeploySet(deploySetID); err != nil {
    fmt.Println("Delete deploy set failed: ", err)
}
```

