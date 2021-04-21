# BEC服务

# 概述

本文档主要介绍BEC GO SDK的使用。在使用本文档前，您需要先了解BEC的一些基本知识。若您还不了解BEC，可以参考[产品描述](https://cloud.baidu.com/doc/BEC/s/xk0p0rsgc)和[入门指南](https://cloud.baidu.com/doc/BEC/s/wk0q5mrcc)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[BEC访问域名](https://cloud.baidu.com/doc/BEC/s/Tk41077qy)的部分，理解Endpoint相关的概念。

## 获取密钥

要使用百度云BEC，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BEC做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建普通型BEC Client

BEC Client是BEC服务的客户端，为开发者与BEC服务进行交互提供了一系列的方法。

### 使用AK/SK新建普通型BEC Client

通过AK/SK方式访问BEC，用户可以参考如下代码新建一个BEC Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bec"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BECClient
	becClient, err := bec.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《操作指南 [管理ACCESSKEY](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb)》。

> **注意：**`ENDPOINT`BEC域名为bec.baidubce.com，更多信息参考：https://cloud.baidu.com/doc/BEC/s/Tk41077qy


## 配置HTTPS协议访问BEC

BEC支持HTTPS传输协议，您可以通过在创建BEC Client对象时指定的Endpoint中指明HTTPS的方式，在BEC GO SDK中使用HTTPS访问BEC服务：

```go
// import "github.com/baidubce/bce-sdk-go/services/bec"

ENDPOINT := "https://bec.baidubce.com" //指明使用HTTPS协议
AK, SK := <your-access-key-id>, <your-secret-access-key>
becClient, _ := bec.NewClient(AK, SK, ENDPOINT)
```

# 主要接口

bec服务为用户提供边缘测vm和边缘测容器的计算服务，支持绑定lb，配置cpu，内存等资源，镜像服务等功能。

## vm实例管理

### 创建虚机实例

通过以下代码，可以创建bec虚机服务
```go
args := &api.CreateVmServiceArgs{
	// 虚机服务名称
	ServiceName: "xxxxxxx@", 
	// 镜像id
	ImageId: "im-dikfttnj-3-u-guangzhou", 
	// 密码
	AdminPass: "x123xxx@",
	// 系统盘配置
    SystemVolume: &api.SystemVolumeConfig{VolumeType: api.DiskTypeNVME}, 
    // cpu大小,必须大于1
    Cpu: 1, 
    // memory大小,必须大于1
    Memory: 2,
    // 部署区域
    DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{City: "HANGZHOU", Region: "EAST_CHINA", ServiceProvider: "CHINA_MOBILE", Replicas: 1}}, 
    // 镜像类型（默认为bcc、仅使用bec虚机自定义镜像时为bec）
    ImageType: "bec", 
    // 密码或密钥配置
    KeyConfig: &api.KeyConfig{Type: "password", AdminPass: "xxxx123@"}
}

result, err := client.CreateVmService(args)
if err != nil {
    fmt.Println("create bec failed:", err)
} else {
    fmt.Println("create bec success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[CreateVmService创建虚机实例](https://cloud.baidu.com/doc/BEC/s/Mkbsrt2yv)

### 更新BEC虚机服务

通过以下代码，可以更新BEC虚机服务
```go
args := &api.UpdateVmServiceArgs{
	UpdateBecVmForm: api.UpdateBecVmForm{
		// 更新类型，包括 password,replicas,resource，serviceName
		Type: bec.api.UpdateVmTypeServiceName, 
		// 虚机名称
		VmName: "vm-xxxxx"}, 
	// 服务名称
	ServiceName: "xxxxtest-2", 
	// 部署区域列表
	DeployInstances: &[]bec.api.DeploymentInstance{bec.api.DeploymentInstance{Region: "SOUTH_CHINA", City: "GUANGZHOU", Replicas: 1, ServiceProvider: api.ServiceChinaUnicom}}
}

err := client.UpdateVmService(serviceId, args)
if err != nil {
    fmt.Println("update bec failed:", err)
} else {
    fmt.Println("update bec success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[UpdateVmService更新BEC虚机服务](https://cloud.baidu.com/doc/BEC/s/jkbsxwduk)

### 获取BEC虚机实例详情

通过以下代码，可以获取BEC虚机实例详情
```go

result, err := client.GetVmServiceDetail(serviceId)
if err != nil {
    fmt.Println("get bec failed:", err)
} else {
    fmt.Println("get bec success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[GetVmServiceDetail获取BEC虚机实例详情](https://cloud.baidu.com/doc/BEC/s/Ekbst85ty)

### 获取BEC虚机服务列表

通过以下代码，可以获取BEC虚机服务列表
```go
args := &api.ListVmServiceArgs{
	// 第几页，默认值1，最小值1
    PageNo: 1,
    // 每页个数，默认值1000，最小1，最大1000
    pageSize: 100,
}

result, err := client.GetVmServiceList(args)
if err != nil {
    fmt.Println("list bec failed:", err)
} else {
    fmt.Println("list bec success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[DescribeLoadBalancerDetail获取BEC虚机服务列表](https://cloud.baidu.com/doc/BEC/s/Okbssrahc)

### 删除BEC虚机服务

通过以下代码，可以删除BEC虚机服务
```go
err := client.DeleteVmService(serviceId)
if err != nil {
    fmt.Println("delete bec failed:", err)
} else {
    fmt.Println("delete bec success")
}
```

### 操作BEC虚机实例

通过以下代码，可以操作BEC虚机实例
```go
err := client.OperateVmDeployment(vmId, action)
if err != nil {
    fmt.Println("operate bec failed:", err)
} else {
    fmt.Println("operate bec success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[OperateVmDeployment操作BEC虚机实例](https://cloud.baidu.com/doc/BEC/s/Fkbroisf0)

### 重装BEC虚机实例系统

通过以下代码，可以重装BEC虚机实例系统
```go
args := &api.ReinstallVmInstanceArg{
	// 镜像id
	ImageId: "im-dikfxxxx", 
	// 	密码
	AdminPass: "1xxAxxx@", 
	// 镜像类型（默认为bcc、仅使用bec虚机自定义镜像时为bec）
	ImageType: api.ImageTypeBec,
	KeyConfig: &api.KeyConfig{Type: "password", AdminPass: "1xxAxxx@"}
}

err := client.ReinstallVmInstance(vmId, args)
if err != nil {
    fmt.Println("reinstall bec failed:", err)
} else {
    fmt.Println("reinstall bec success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[ReinstallVmInstance重装BEC虚机实例系统](https://cloud.baidu.com/doc/BEC/s/Ekbrqqv9j)

### 更新BEC虚机实例

通过以下代码，可以更新BEC虚机实例
```go
args := &api.UpdateVmDeploymentArgs{
	// 更新类型(password,replicas,resource)
	Type: "resource", 
	// cpu大小
	Cpu: 2,
	VmName: "xxxxx-test",
    // 系统盘配置
    SystemVolume: &api.SystemVolumeConfig{VolumeType: api.DiskTypeNVME, Name: "sys", PvcName: "lvm-xxxxxx-rootfs"}
}

err := client.UpdateVmDeployment(vmId, args)
if err != nil {
    fmt.Println("update bec instance failed:", err)
} else {
    fmt.Println("update bec instance bec success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[UpdateVmDeployment更新BEC虚机实例](https://cloud.baidu.com/doc/BEC/s/wkbroqu5a)

### 获取BEC虚机实例详情

通过以下代码，可以获取BEC虚机实例详情
```go
err := client.GetVirtualMachine(vmId)
if err != nil {
    fmt.Println("get bec instance failed:", err)
} else {
    fmt.Println("get bec instance bec success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[GetVirtualMachine获取BEC虚机实例详情](https://cloud.baidu.com/doc/BEC/s/Fkbro0yld)

### 获取BEC虚机实例列表

通过以下代码，可以获取BEC虚机实例列表
```go
args := &api.ListRequest{
	// 查询实例的关键字类型，instanceId或serviceId,缺省为serviceId
	KeywordType: "instanceId", 
	// 查询实例的关键字
	Keyword: "vm-xxx"
}

err := client.GetVmInstanceList(args)
if err != nil {
    fmt.Println("get bec instance list failed:", err)
} else {
    fmt.Println("get bec instance list success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[GetVmInstanceList更新BEC虚机实例](https://cloud.baidu.com/doc/BEC/s/Qkbrebfbz)

## 容器相关服务

### 创建BEC容器服务

通过以下代码，可以创建BEC容器服务
```go
args := &api.CreateServiceArgs{
	// 服务名称
	ServiceName: "xxxx-1-test",
	// 付费方式，只支持后付费postpay
	PaymentMethod: "postpay", 
	// 容器组名称
	ContainerGroupName: "cg1",
	// 当needPublicIp为true时，用于设置外网带宽，范围为1Mps到2048Mps，可开通白名单上限增加到5120Mps
    Bandwidth: 100,
    // 是否购买公网IP，缺省否
    NeedPublicIp: false,
    // 容器组信息
    Containers: &[]api.ContainerDetails{
        api.ContainerDetails{
            Name: "container01",
            Cpu:          1,
            Memory:       2,
            ImageAddress: "hub.baidubce.com/public/mysql",
            ImageVersion: "5.7",
            Commands: []string{"sh",
                "-c",
                "echo OK!&& sleep 3660"},
            VolumeMounts: &[]api.V1VolumeMount{
                api.V1VolumeMount{
                MountPath: "/temp",
                Name:      "emptydir01",
            },
        },},
    },
    // 部署地域信息
    DeployInstances: &[]api.DeploymentInstance{
    api.DeploymentInstance{
        Region:          "EAST_CHINA",
        Replicas:        1,
        City:            "SHANGHAI",
        ServiceProvider: "CHINA_TELECOM"},
    },
    // 存储卷信息
    Volumes: &api.Volume{
    EmptyDir: &[]api.EmptyDir{
        api.EmptyDir{Name: "emptydir01"},
        },
    },
    // 标签信息
    Tags: &[]api.Tag{
        api.Tag{
            TagKey:   "a",
            TagValue: "1"
            },
    },
}
err := client.CreateService(args)
if err != nil {
    fmt.Println("create bec service failed:", err)
} else {
    fmt.Println("create bec service success: ")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[CreateService创建BEC容器服务](https://cloud.baidu.com/doc/BEC/s/Wk3zawt4x)

### 删除BEC容器服务

通过以下代码，可以删除BEC容器服务
```go
err := client.DeleteService(serviceId)
if err != nil {
    fmt.Println("delete bec service failed:", err)
} else {
    fmt.Println(""delete bec service success: ")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[DeleteService删除BEC容器服务](https://cloud.baidu.com/doc/BEC/s/Uk3zb1nwe)

### 获取BEC容器服务详情

通过以下代码，可以获取BEC容器服务详情
```go
result, err := client.GetService(serviceId)
if err != nil {
    fmt.Println("describe backend servers failed:", err)
} else {
    fmt.Println("describe backend servers success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[GetService获取BEC容器服务详情](https://cloud.baidu.com/doc/BEC/s/Hk413e70s)

### 删除BEC容器服务

通过以下代码，可以删除BEC容器服务
```go
result, err := client.DeleteService(serviceId)
if err != nil {
    fmt.Println("describe health status failed:", err)
} else {
    fmt.Println("describe health status success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[DeleteService删除BEC容器服务](https://cloud.baidu.com/doc/BEC/s/Uk3zb1nwe)

## 负载均衡器

### 创建负载均衡

通过以下代码，可以创建负载均衡
```go
args := &api.CreateBlbArgs{
	// 负载均衡名称
	BlbName: "xxxx-test", 
	// 负载均衡所在区域信息
	Region: api.RegionEastChina,
	// 负载均衡所在城市信息
    City: "HANGZHOU", 
    LbType: "vm", 
    ServiceProvider: api.ServiceChinaMobile
}
err := client.CreateBlb(args)
if err != nil {
    fmt.Println("remove backend servers failed:", err)
} else {
    fmt.Println("remove backend servers success: ")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[CreateBlb创建负载均衡](https://cloud.baidu.com/doc/BEC/s/zkbrnvg6w)

### 删除负载均衡

通过以下代码，可以删除负载均衡
```go

err := client.DeleteBlb(BLBID)
if err != nil {
    fmt.Println("create TCP Listener failed:", err)
} else {
    fmt.Println("create TCP Listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[DeleteBlb删除负载均衡](https://cloud.baidu.com/doc/BEC/s/qkbrnx86o)

### 查看负载均衡详情

通过以下代码，可以查看负载均衡详情
```go
err := client.GetBlbDetail(BLBID)
if err != nil {
    fmt.Println("update TCP Listener failed:", err)
} else {
    fmt.Println("update TCP Listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[GetBlbDetail查看负载均衡详情](https://cloud.baidu.com/doc/BEC/s/ekbrnvz0c)

### 创建监听设置

通过以下代码，可以创建监听设置
```go
args := &api.BlbMonitorArgs{
	// 转发规则
	LbMode: api.LbModeWrr, 
	// 负载均衡端口
	FrontendPort: &api.Port{Protocol: api.ProtocolUdp,Port: 80},
	// 后端端口
    BackendPort: 80, 
    // 健康检查设置
    HealthCheck: &api.HealthCheck{HealthCheckString: "", HealthCheckType: "udp", HealthyThreshold: 1000, 
        UnhealthyThreshold: 1000, 
        TimeoutInSeconds: 900, 
        IntervalInSeconds: 3}
}

result, err := client.CreateBlbMonitorPort(BLBID, args)
if err != nil {
    fmt.Println("describe TCP Listener failed:", err)
} else {
    fmt.Println("describe TCP Listener success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[CreateBlbMonitorPort创建监听设置](https://cloud.baidu.com/doc/BEC/s/ikbrnxqnj)

### 删除监听设置

通过以下代码，可以删除监听设置
```go
// 负载均衡监听端口列表
args := &[]api.Port{{
	Protocol: api.ProtocolUdp, 
	Port: 80},
}
err := client.DeleteBlbMonitorPort(BLBID, args)
if err != nil {
    fmt.Println("create UDP Listener failed:", err)
} else {
    fmt.Println("create UDP Listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[DeleteBlbMonitorPort删除监听设置](https://cloud.baidu.com/doc/BEC/s/dkbrnxzcw)

### 创建后端服务器

通过以下代码，可以创建后端服务器
```go
args := &&api.CreateBlbBindingArgs{
	// 创建后端服务器请求
	BindingForms: &[]api.BlbBindingForm{api.BlbBindingForm{DeploymentId: "xxxx", PodWeight: &[]api.Backends{api.Backends{Name: "xxxxxx", Ip: "172.xx.x.xx", Weight: 100}},
}}}
err := client.CreateBlbBinding(BLBID, args)
if err != nil {
    fmt.Println("update UDP Listener failed:", err)
} else {
    fmt.Println("update UDP Listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[CreateBlbBinding创建后端服务器](https://cloud.baidu.com/doc/BEC/s/Ckbrnyb8q)

### 删除后端服务器

通过以下代码，可以删除后端服务器
```go
args := &api.DeleteBlbBindPodArgs{
	// 删除后端服务器请求
	PodWeightList: &[]api.Backends{
        api.Backends{Name: "vm-xxx", Ip: "172.16.9xxx.xxx", Weight: 10}},
        DeploymentIds: []string{"vmrs-xxxx"},
}

result, err := client.DeleteBlbBindPod(BLBID, args)
if err != nil {
    fmt.Println("describe UDP Listener failed:", err)
} else {
    fmt.Println("describe UDP Listener success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[DeleteBlbBindPod删除后端服务器](https://cloud.baidu.com/doc/BEC/s/2kbrnz2wi)

## 镜像相关

### 创建BEC虚机镜像

通过以下代码，可以创建BEC虚机镜像
```go
args := &api.CreateVmImageArgs{
	// 虚机实例id
	VmId: "vm-xxxx-1", 
	// 镜像名称(长度1-65)
	Name: "xxxx-test"
}
err := client.CreateVmImage(args)
if err != nil {
    fmt.Println("create HTTP Listener failed:", err)
} else {
    fmt.Println("create HTTP Listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[CreateVmImage创建BEC虚机镜像](https://cloud.baidu.com/doc/BEC/s/Uklg7azo0)

### 批量删除BEC虚机镜像

通过以下代码，可以批量删除BEC虚机镜像
```go
args := []string{"xxxxxx-1", "xxxxxx-2"}
err := client.DeleteVmImage(args)
if err != nil {
    fmt.Println("update HTTP Listener failed:", err)
} else {
    fmt.Println("update HTTP Listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[DeleteVmImage批量删除BEC虚机镜像](https://cloud.baidu.com/doc/BEC/s/Sklgbqn7g)

```

## 客户端异常

客户端异常表示客户端尝试向BEC发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当BEC服务端出现异常时，BEC服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[BEC错误返回](https://cloud.baidu.com/doc/BEC/s/5k4106ncs)

# 版本变更记录

## v0.9.11 [2021-04-21]

首次发布：

 - 创建、列表、更新、删除VM镜像
 - 创建、列表、更新、删除、详情、批量创建、批量删除blb、获取负载均衡监控信息
 - 创建、删除、列表、更新、详情、批量创建blb监听
 - 获取负载均衡已绑定资源列表、获取负载均衡可绑定的部署列表、获取部署中可绑定资源列表、创建后端服务器、删除后端服务器、修改负载均衡已绑定资源权重
 - 创建、更新、列表、详情、删除、操作、批量删除、批量操作虚机服务、获取BEC虚机服务监控
 - 创建无实例的虚机服务、创建虚机实例、添加辅助ip、删除辅助ip 
 - 列表、详情、删除、更新、重装、操作虚机实例、获取虚机实例配置、获取所在节点的BEC虚机列表、获取虚机监控信息
 - 创建、列表、详情、操作、更新、删除、批量操作、批量删除容器服务、获取BEC容器服务监控

```