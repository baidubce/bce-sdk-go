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
    // 部署区域,支持创建vpc虚机
    DeployInstances: &[]api.DeploymentInstance{api.DeploymentInstance{City: "HANGZHOU", Region: "EAST_CHINA", ServiceProvider: "CHINA_MOBILE", Replicas: 1,NetworkType: "vpc",}}, 
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
	DeployInstances: &[]bec.api.DeploymentInstance{bec.api.DeploymentInstance{Region: "SOUTH_CHINA", City: "GUANGZHOU", Replicas: 1, ServiceProvider: api.ServiceChinaUnicom,NetworkType: "vpc",}}
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

### 更新BEC服务

通过以下代码，可以更新BEC服务
```go
getReq := &api.UpdateServiceArgs{ServiceName: "s-f9ngbkbc", Type: api.UpdateServiceTypeReplicas, DeployInstances: &[]api.DeploymentInstance{
    api.DeploymentInstance{Region: api.RegionEastChina, Replicas: 1, City: "HANGZHOU", ServiceProvider: api.ServiceChinaMobile},
}}
res, err := CLIENT.UpdateService("s-f9ngbkbc", getReq)
if err != nil {
    fmt.Println("update bec service failed:", err)
} else {
    fmt.Println("update bec service success: ")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[更新BEC服务](https://cloud.baidu.com/doc/BEC/s/Tk3zaz9by)

### 启动BEC服务

通过以下代码，可以启动BEC服务
```go
res, err := CLIENT.ServiceAction("s-xxx", api.ServiceActionStart)
if err != nil {
    fmt.Println("start bec service failed:", err)
} else {
    fmt.Println("start bec service success: ")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[启动BEC服务](https://cloud.baidu.com/doc/BEC/s/sk3zb12kh)

### 停止BEC服务

通过以下代码，可以停止BEC服务
```go
res, err := CLIENT.ServiceAction("s-xxx",  api.ServiceActionStop)
if err != nil {
    fmt.Println("stop bec service failed:", err)
} else {
    fmt.Println("stop bec service success: ")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[停止BEC服务](https://cloud.baidu.com/doc/BEC/s/gk412paz1)




### 获取BEC容器服务详情

通过以下代码，可以获取BEC容器服务详情
```go
result, err := client.GetService(serviceId)
if err != nil {
    fmt.Println("get bec pod service failed:", err)
} else {
    fmt.Println("get bec pod service success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[GetService获取BEC容器服务详情](https://cloud.baidu.com/doc/BEC/s/Hk413e70s)

### 查询部署详情

通过以下代码，可以查询部署详情
```go
result, err := CLIENT.GetPodDeployment("sts-xxxx")
if err != nil {
    fmt.Println("get bec pod deployment failed:", err)
} else {
    fmt.Println("get bec pod deployment success: ", result)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询部署详情](https://cloud.baidu.com/doc/BEC/s/Lk3fmy6v2)

### 查询部署资源监控

通过以下代码，可以获取查询部署资源监控
```go
res, err := CLIENT.GetPodDeploymentMetrics("sts-xxxx", api.MetricsTypeMemory, "", 1661270400, 1661356800, 0)
if err != nil {
    fmt.Println("get bec pod deployment failed:", err)
} else {
    fmt.Println("get bec pod deployment success: ", res)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询部署资源监控](https://cloud.baidu.com/doc/BEC/s/Wk3h9n5n1)
### 更新部署副本数

通过以下代码，可以获取更新部署副本数
```go
getReq := &api.UpdateDeploymentReplicasRequest{
    Replicas: 2,
}
res, err := CLIENT.UpdatePodDeploymentReplicas("sts-xxxx", getReq)
if err != nil {
    fmt.Println("update bec pod deployment replicas failed:", err)
} else {
    fmt.Println("update bec pod deployment replicas success: ", res)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[更新部署副本数](https://cloud.baidu.com/doc/BEC/s/Wk3h9n5n1)

### 删除部署

通过以下代码，可以删除部署
```go
getReq := &[]string{"sts-xxxx"}
res, err := CLIENT.DeletePodDeployment(getReq)
if err != nil {
    fmt.Println("delete bec pod deployment failed:", err)
} else {
    fmt.Println("delete bec pod deployment success: ", res)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除部署](https://cloud.baidu.com/doc/BEC/s/jk3h9bcr5)

### 查询pod列表

通过以下代码，可以查询pod列表
```go
res, err := CLIENT.GetPodList(1, 100, "", "", "", "", "")
if err != nil {
    fmt.Println("get bec pod list failed:", err)
} else {
    fmt.Println("get bec pod list success: ", res)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询pod列表](https://cloud.baidu.com/doc/BEC/s/2k3i3ucrh)

### 查询pod资源监控

通过以下代码，可以查询pod资源监控
```go
res, err := CLIENT.GetPodMetrics("sts-xxx-0", api.MetricsTypeMemory, "", 1661270400, 1661356800, 0)

if err != nil {
    fmt.Println("get bec pod metrics failed:", err)
} else {
    fmt.Println("get bec pod metrics success: ", res)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询pod资源监控](https://cloud.baidu.com/doc/BEC/s/Rk3ibrrdu)

### 查询pod详情

通过以下代码，可以查询pod详情
```go
res, err := CLIENT.GetPodDetail("sts-xzzxxxx-0")
if err != nil {
    fmt.Println("get bec pod detail failed:", err)
} else {
    fmt.Println("get bec pod detail success: ", res)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询pod详情](https://cloud.baidu.com/doc/BEC/s/Ok3i3vgl7)

### 重启容器组

通过以下代码，可以重启容器组
```go
err := CLIENT.RestartPod("sts-xxxxx-0")
if err != nil {
    fmt.Println("restart bec pod failed:", err)
} else {
    fmt.Println("restart bec pod success: ", res)
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[重启容器组](https://cloud.baidu.com/doc/BEC/s/Jk3ib7df8)


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
	// 创建applb
    NetworkType: "vpc",
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

## 部署集相关
### 创建部署集

通过以下代码，可以创建部署集
```go
getReq := &api.CreateDeploySetArgs{
Name: "xxx_test",
Desc: "xxx-test",
}
res, err := CLIENT.CreateDeploySet(getReq)
if err != nil {
    fmt.Println("create deploy set failed:", err)
} else {
    fmt.Println("create deploy set success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[创建部署集](https://cloud.baidu.com/doc/BEC/s/Sl0t89s48)

### 修改部署集

通过以下代码，可以修改部署集
```go
getReq := &api.CreateDeploySetArgs{
    Name: "xxx_test",
    Desc: "xxx-test",
}
err := CLIENT.UpdateDeploySet("dset-xxx", getReq)
res, err := CLIENT.CreateDeploySet(getReq)
if err != nil {
    fmt.Println("update deploy set failed:", err)
} else {
    fmt.Println("update deploy set success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[修改部署集](https://cloud.baidu.com/doc/BEC/s/Ml0tb0d5s)

### 获取部署集列表

通过以下代码，可以获取部署集列表
```go
getReq := &api.ListRequest{}
res, err := CLIENT.GetDeploySetList(getReq)
if err != nil {
    fmt.Println("get deploy set list failed:", err)
} else {
    fmt.Println("get deploy set list success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[获取部署集列表](https://cloud.baidu.com/doc/BEC/s/Cl0t8xm4g)

### 获取部署集详情

通过以下代码，可以获取部署集详情
```go
res, err := CLIENT.GetDeploySetDetail("dset-xxxx")
if err != nil {
    fmt.Println("get deploy set details failed:", err)
} else {
    fmt.Println("get deploy set details success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[获取部署集详情](https://cloud.baidu.com/doc/BEC/s/9l0talan1)

### 删除部署集

通过以下代码，可以删除部署集
```go
err := CLIENT.DeleteDeploySet("dset-y4tumnel")
if err != nil {
    fmt.Println("delete deploy set failed:", err)
} else {
    fmt.Println("delete deploy set success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除部署集](https://cloud.baidu.com/doc/BEC/s/1l0ulgpsv)

### 虚机实例调整部署集

通过以下代码，可以调整虚机实例的部署集
```go
getReq := &api.UpdateVmDeploySetArgs{
InstanceId:      "vm-xxxx",
DeploysetIdList: []string{"dset-xxxx"},
}
err := CLIENT.UpdateVmInstanceDeploySet(getReq)
if err != nil {
    fmt.Println("update vm instance deploy set failed:", err)
} else {
    fmt.Println("update vm instance deploy set success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[虚机实例调整部署集](https://cloud.baidu.com/doc/BEC/s/nl0um3hvs)

### 部署集移除虚机实例

通过以下代码，可以将虚机实例从部署集移除
```go
getReq := &api.DeleteVmDeploySetArgs{
    DeploysetId:    "dset-y4tumnel",
    InstanceIdList: []string{"vm-dstkrmda-cn-langfang-ct-4thbz"},
}
err := CLIENT.DeleteVmInstanceFromDeploySet(getReq)
if err != nil {
    fmt.Println("remove vm instance from deploy set failed:", err)
} else {
    fmt.Println("remove vm instance from deploy set success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[部署集移除实例](https://cloud.baidu.com/doc/BEC/s/Ml0ulrz5z)
## APPBLB相关
### 创建APPBLB实例

通过以下代码，可以创建APPBLB实例
```go
getReq := &api.CreateAppBlbRequest{
    Name:         "xxx_test_applb",
    Desc:         "xxx-test",
    RegionId:     "cn-hangzhou-cm",
    NeedPublicIp: true,
    SubnetId:     "sbn-xx",
    VpcId:        "vpc-xx",
}
res, err := CLIENT.CreateAppBlb("testCreateAppBlb", getReq)

if err != nil {
    fmt.Println("create app blb instance failed:", err)
} else {
    fmt.Println("create app blb instance success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[创建APPBLB实例](https://cloud.baidu.com/doc/BEC/s/zl4nug4yg)

### 修改APPBLB实例
通过以下代码，可以修改APPBLB实例
```go
getReq := &api.ModifyBecBlbRequest{
    Name: "xx_test_applb",
    Desc: "xx-test1",
}
err := CLIENT.UpdateAppBlb("testUpdateAppBlb", "applb-xx", getReq)

if err != nil {
    fmt.Println("update app blb instance failed:", err)
} else {
    fmt.Println("update app blb instance success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[修改APPBLB实例](https://cloud.baidu.com/doc/BEC/s/Ul4nuv8n2)
### 查询APPBLB实例列表
通过以下代码，可以查询APPBLB实例列表
```go
getReq := &api.MarkerRequest{}
res, err := CLIENT.GetAppBlbList(getReq)

if err != nil {
    fmt.Println("get app blb instance list failed:", err)
} else {
    fmt.Println("get app blb instance list success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询APPBLB实例列表](https://cloud.baidu.com/doc/BEC/s/9l4nv22ji)

### 查询APPBLB实例详情
通过以下代码，可以查询APPBLB实例详情
```go
res, err := CLIENT.GetAppBlbDetails("applb-xxxx")

if err != nil {
    fmt.Println("get app blb instance detail failed:", err)
} else {
    fmt.Println("get app blb instance detail success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询APPBLB实例详情](https://cloud.baidu.com/doc/BEC/s/Ul4nvz6d8)

### 删除APPBLB实例
通过以下代码，可以删除APPBLB实例
```go
err := CLIENT.DeleteAppBlbInstance("applb-xxx", "")

if err != nil {
    fmt.Println("delete app blb instance  failed:", err)
} else {
    fmt.Println("delete app blb instance  success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除APPBLB实例](https://cloud.baidu.com/doc/BEC/s/7l4nwir90)

### 创建TCP监听器
通过以下代码，可以创建TCP监听器
```go
getReq := &api.CreateBecAppBlbTcpListenerRequest{
    ListenerPort:      80,
    Scheduler:         "RoundRobin",
    TcpSessionTimeout: 1000,
}
err := CLIENT.CreateTcpListener("testCreateTcpListener", "applb-xxx", getReq)
if err != nil {
    fmt.Println("create app tcp listener failed:", err)
} else {
    fmt.Println("create app tcp listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[创建TCP监听器](https://cloud.baidu.com/doc/BEC/s/il4nwx3ts)

### 创建UDP监听器
通过以下代码，可以创建UDP监听器
```go
getReq := &api.CreateBecAppBlbUdpListenerRequest{
    ListenerPort:      80,
    Scheduler:         "RoundRobin",
    UdpSessionTimeout: 1000,
}
err := CLIENT.CreateUdpListener("testCreateTcpListener", "applb-xxxx", getReq)
if err != nil {
    fmt.Println("create app udp listener failed:", err)
} else {
    fmt.Println("create app udp listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[创建UDP监听器](https://cloud.baidu.com/doc/BEC/s/sl4p73e1g)

### 新建监听器策略
通过以下代码，可以新建监听器策略
```go
getReq := &api.CreateAppBlbPoliciesRequest{
    ListenerPort: 80,
    AppPolicyVos: []api.AppPolicyVo{
        {
            AppIpGroupId: "bec_ip_group-xxx",
            Priority:     1,
            Desc:         "xxx-test",
        },
    },
}
err := CLIENT.CreateListenerPolicy("", "applb-xxx", getReq)
if err != nil {
    fmt.Println("create app blb listener policy failed:", err)
} else {
    fmt.Println("create app blb listener policy success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[新建监听器策略](https://cloud.baidu.com/doc/BEC/s/Hl4qgo88o)

### 查询监听器策略
通过以下代码，可以查询监听器策略
```go
getReq := &api.GetBlbListenerPolicyRequest{
    Port: 80,
}
res, err := CLIENT.GetListenerPolicy("applb-xxx", getReq)
if err != nil {
    fmt.Println("get app blb listener policy failed:", err)
} else {
    fmt.Println("get app blb listener policy success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询监听器策略](https://cloud.baidu.com/doc/BEC/s/cl4qgvkbp)

### 删除监听器策略
通过以下代码，可以删除监听器策略
```go
getReq := &api.DeleteAppBlbPoliciesRequest{
    Port: 80,
    PolicyIdList: []string{
        "bec_policy-scr9cwtk",
    },
}
err := CLIENT.DeleteListenerPolicy("", "applb-xxxxx", getReq)
if err != nil {
    fmt.Println("delete app blb listener policy failed:", err)
} else {
    fmt.Println("delete app blb listener policy success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除监听器策略](https://cloud.baidu.com/doc/BEC/s/Bl4qhc0vy)

### 修改TCP监听器
通过以下代码，可以修改TCP监听器
```go
getReq := &api.UpdateBecAppBlbTcpListenerRequest{
    Scheduler:         "RoundRobin",
    TcpSessionTimeout: 800,
}
err := CLIENT.UpdateTcpListener("testUpdateTcpListener", "applb-xxx", "80", getReq)
if err != nil {
    fmt.Println("update app tcp listener failed:", err)
} else {
    fmt.Println("update app tcp listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[修改TCP监听器](https://cloud.baidu.com/doc/BEC/s/sl4p73e1g)

### 修改UDP监听器
通过以下代码，可以修改UDP监听器
```go
getReq := &api.UpdateBecAppBlbUdpListenerRequest{
    Scheduler:         "RoundRobin",
    UdpSessionTimeout: 800,
}
err := CLIENT.UpdateUdpListener("testUpdateUdpListener", "applb-xxx", "80", getReq)
if err != nil {
    fmt.Println("update app udp listener failed:", err)
} else {
    fmt.Println("update app udp listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[修改UDP监听器](https://cloud.baidu.com/doc/BEC/s/al4p7c82y)

### 查询TCP监听器
通过以下代码，可以查询TCP监听器
```go
getReq := &api.GetBecAppBlbListenerRequest{
    ListenerPort: 80,
}
res, err := CLIENT.GetTcpListener("applb-xxxx", getReq)
if err != nil {
    fmt.Println("get app udp listener failed:", err)
} else {
    fmt.Println("get app udp listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询TCP监听器](https://cloud.baidu.com/doc/BEC/s/Ul4nxiz6l)

### 查询TCP监听器
通过以下代码，可以查询UDP监听器
```go
getReq := &api.GetBecAppBlbListenerRequest{
    ListenerPort: 80,
}
res, err := CLIENT.GetUdpListener("applb-xxxx", getReq)
if err != nil {
    fmt.Println("get app udp listener failed:", err)
} else {
    fmt.Println("get app udp listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询UDP监听器](https://cloud.baidu.com/doc/BEC/s/4l4p7ih5s)

### 删除监听器
通过以下代码，可以删除监听器
```go
getReq := &api.DeleteBlbListenerRequest{
PortTypeList: []api.PortTypeList{
        {
            Port: 80,
            Type: "TCP",
        },
        {
            Port: 80,
            Type: "UDP",
        },
    },
}
err := CLIENT.DeleteAppBlbListener("applb-xxx", "deleteApplbInstance", getReq)
if err != nil {
    fmt.Println("delete app listener failed:", err)
} else {
    fmt.Println("delete app listener success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除监听器](https://cloud.baidu.com/doc/BEC/s/dl4p7sfyb)

### 创建IP组
通过以下代码，可以创建IP组
```go
getReq := &api.CreateBlbIpGroupRequest{
    Name: "xxx-testIpGroup",
    Desc: "xxx-test",
}
res, err := CLIENT.CreateIpGroup("testIpGroup", "applb-xxx", getReq)
if err != nil {
    fmt.Println("create app blb ip group failed:", err)
} else {
    fmt.Println("create app blb ip group success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[创建IP组](https://cloud.baidu.com/doc/BEC/s/nl4p9vtw2)

### 更新IP组
通过以下代码，可以更新IP组
```go
getReq := &api.UpdateBlbIpGroupRequest{
    Name:      "xxx-testIpGroupupdate",
    Desc:      "xxx-testupdate",
    IpGroupId: "bec_ip_group-xxx",
}
err := CLIENT.UpdateIpGroup("testIpGroup", "applb-xxx", getReq)
if err != nil {
    fmt.Println("update app blb ip group failed:", err)
} else {
    fmt.Println("update app blb ip group success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[更新IP组](https://cloud.baidu.com/doc/BEC/s/wl4pahvlw)

### 查询IP组列表
通过以下代码，可以查询IP组列表
```go
getReq := &api.GetBlbIpGroupListRequest{}
res, err := CLIENT.GetIpGroup("applb-xxxx", getReq)
if err != nil {
    fmt.Println("get app blb ip group list failed:", err)
} else {
    fmt.Println("get app blb ip group list success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询IP组列表](https://cloud.baidu.com/doc/BEC/s/El4pan7ox)

### 删除IP组
通过以下代码，可以删除IP组
```go
getReq := &api.DeleteBlbIpGroupRequest{
    IpGroupId: "bec_ip_group-ukadxdrq",
}
err := CLIENT.DeleteIpGroup("testDeleteIpGroup", "applb-xxxx", getReq)
if err != nil {
    fmt.Println("delete app blb ip group failed:", err)
} else {
    fmt.Println("delete app blb ip group success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除IP组](https://cloud.baidu.com/doc/BEC/s/Ml4qem709)

### 创建IP组协议
通过以下代码，可以创建IP组协议
```go
getReq := &api.CreateBlbIpGroupBackendPolicyRequest{
    IpGroupId:                   "bec_ip_group-xxx,
    Type:                        "TCP",
    HealthCheck:                 "TCP",
    HealthCheckPort:             80,
    HealthCheckTimeoutInSecond:  10,
    HealthCheckIntervalInSecond: 3,
    HealthCheckDownRetry:        4,
    HealthCheckUpRetry:          5,
}
res, err := CLIENT.CreateIpGroupPolicy("", "applb-xxx", getReq)
if err != nil {
    fmt.Println("create app blb ip group policy failed:", err)
} else {
    fmt.Println("create app blb ip group policy success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[创建IP组协议](https://cloud.baidu.com/doc/BEC/s/Ml4qem709)

### 更新IP组协议
通过以下代码，可以更新IP组协议
```go
getReq := &api.UpdateBlbIpGroupBackendPolicyRequest{
    IpGroupId:       "bec_ip_group-xxx",
    Id:              "bec_ip_group_policy-xxx",
    HealthCheckPort: 80,
}
err := CLIENT.UpdateIpGroupPolicy("", "applb-xxx", getReq)
if err != nil {
    fmt.Println("update app blb ip group policy failed:", err)
} else {
    fmt.Println("update app blb ip group policy success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[更新IP组协议](https://cloud.baidu.com/doc/BEC/s/5l4qexv01)

### 查询IP组协议列表
通过以下代码，可以查询IP组协议列表
```go
getReq := &api.GetBlbIpGroupPolicyListRequest{
    IpGroupId: "bec_ip_group-xxx",
}
res, err := CLIENT.GetIpGroupPolicyList("applb-xxx", getReq)
if err != nil {
    fmt.Println("get app blb ip group policy list failed:", err)
} else {
    fmt.Println("get app blb ip group policy list success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询IP组协议列表](https://cloud.baidu.com/doc/BEC/s/2l4qf4jv6)

### 删除IP组协议
通过以下代码，可以删除IP组协议
```go
getReq := &api.DeleteBlbIpGroupBackendPolicyRequest{
    IpGroupId:           "bec_ip_group-xxx",
    BackendPolicyIdList: []string{"bec_ip_group_policy-xx"},
}
err := CLIENT.DeleteIpGroupPolicy("", "applb-xxx", getReq)
if err != nil {
    fmt.Println("delete app blb ip group policy failed:", err)
} else {
    fmt.Println("delete app blb ip group policy success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除IP组协议](https://cloud.baidu.com/doc/BEC/s/Ll4qg1hun)

### 创建IP组成员
通过以下代码，可以创建IP组成员
```go
getReq := &api.CreateBlbIpGroupMemberRequest{
    IpGroupId: "bec_ip_group-ukadxdrq",
    MemberList: []api.BlbIpGroupMember{
        {
            Ip:     "172.16.240.25",
            Port:   90,
            Weight: 100,
        },
    },
}
res, err := CLIENT.CreateIpGroupMember("", "applb-xxxx", getReq)
if err != nil {
    fmt.Println("create app blb ip group member failed:", err)
} else {
    fmt.Println("create app blb ip group member success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[创建IP组成员](https://cloud.baidu.com/doc/BEC/s/yl4qf9xbw)

### 更新IP组成员
通过以下代码，可以更新IP组成员
```go
getReq := &api.UpdateBlbIpGroupMemberRequest{
    IpGroupId: "bec_ip_group-ukadxdrq",
    MemberList: []api.UpdateBlbIpGroupMember{
        {
            MemberId: "bec_ip_member-ouiinabp",
            Port:     8080,
            Weight:   100,
        },
    },
}
err := CLIENT.UpdateIpGroupMember("", "applb-xxxx", getReq)
if err != nil {
    fmt.Println("update app blb ip group member failed:", err)
} else {
    fmt.Println("update app blb ip group member success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[更新IP组成员](https://cloud.baidu.com/doc/BEC/s/Gl4qfi31t)

### 查询IP组成员列表
通过以下代码，可以查询IP组成员列表
```go
getReq := &api.GetBlbIpGroupMemberListRequest{
    IpGroupId: "bec_ip_group-xxx",
}
res, err := CLIENT.GetIpGroupMemberList("applb-xxxx", getReq)
if err != nil {
    fmt.Println("get app blb ip group member list failed:", err)
} else {
    fmt.Println("get app blb ip group member list success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[查询IP组成员列表](https://cloud.baidu.com/doc/BEC/s/Il4qfnftj)

### 删除IP组成员
通过以下代码，可以删除IP组成员
```go
getReq := &api.DeleteBlbIpGroupBackendMemberRequest{
    IpGroupId:    "bec_ip_group-xxx",
    MemberIdList: []string{"bec_ip_member-xxx"},
}
err := CLIENT.DeleteIpGroupMember("", "applb-xxx", getReq)
if err != nil {
    fmt.Println("delete app blb ip group member failed:", err)
} else {
    fmt.Println("delete app blb ip group member success")
}
```
> **提示：**
>
> - 详细的参数配置及限制条件，可以参考BEC API 文档[删除IP组成员](https://cloud.baidu.com/doc/BEC/s/Hl4qftd63)






## 客户端异常

客户端异常表示客户端尝试向BEC发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError；当上传文件时发生IO异常时，也会抛出BceClientError。

## 服务端异常

当BEC服务端出现异常时，BEC服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[BEC错误返回](https://cloud.baidu.com/doc/BEC/s/5k4106ncs)

# 版本变更记录

## 更新日期 [2022-08-25]

更新内容

- 更新虚机相关的参数、支持创建VPC虚机、更新获取虚机监控的接口
- 更新负载均衡相关参数、支持创建APPLB、更新获取负载均衡监控的接口
- 更新了容器服务接口的相关参数

新增内容

- 创建、删除、列表、更新、详情虚机部署集
- 将虚机移入、移除部署集
- 监控接口新增stepInMin参数
- 创建、修改、查询、删除APPBLB
- 创建、修改、查询、删除APPBLB监听器以及监听器策略
- 创建、修改、查询、删除APPBLB IP组、IP组协议、IP组成员
- 查询容器部署详情、查询容器部署监控、更新容器部署副本、删除容器部署
- 查询POD列表、详情、资源监控，重启容器组



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

