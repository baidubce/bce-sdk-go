# MCP服务

# 概述

本文档主要介绍MCP GO SDK的使用。在使用本文档前，您需要先了解MCP的一些基本知识，并已开通了MCP服务。若您还不了解MCP，可以参考[产品描述](https://cloud.baidu.com/doc/MCT/s/9jwvz4hes)和[入门指南](https://cloud.baidu.com/doc/MCT/s/mkd8ii2ck)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[使用须知](https://cloud.baidu.com/doc/MCT/s/Sjwvz5hq5)的部分，理解Endpoint相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/Regions.html)。

目前支持“华北-北京”、“华南-广州”和“华东-苏州”三个区域。北京区域：`http://media.bj.baidubce.com`，广州区域：`http://media.gz.baidubce.com`，苏州区域：`http://media.su.baidubce.com`。对应信息为：

| 访问区域 | 对应Endpoint          |
| -------- | --------------------- |
| BJ       | media.bj.baidubce.com |
| GZ       | media.gz.baidubce.com |
| SU       | media.su.baidubce.com |

## 获取密钥

要使用百度云MCP，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问MCP做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建MCP Client

MCP Client是MCP服务的客户端，为开发者与MCP服务进行交互提供了一系列的方法。

### 使用AK/SK新建MCP Client

通过AK/SK方式访问MCP，用户可以参考如下代码新建一个MCP Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/media"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个MCPClient
	MEDIA_CLIENT, err := media.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《相关参考 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为MCP的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`http://media.bj.baidubce.com`。

# pipeline队列

队列分为免费型与专享型：

- 免费型队列中的转码任务分享百度智能云为音视频转码所提供的约400路720P转码计算资源。
- 专享型队列需额外采购，以便更好的满足那些对于转码时效性和稳定性有更高要求的用户的业务需求。

用户可以利用队列实现任务优先级。用户通过创建多个队列达到区分任务优先级的目的，将大部分任务创建至普通优先级队列，将高优的任务放入高优先级的队列，以利用队列先到先服务的工作原理来实现任务的优先级调整。

## 新建Pipeline

如下代码可以新建一个Pipeline。

```go
pipelineName := "test"
sourceBucket := "testBucket"
targetBucket := "targetBucket"
capacity := 10
err := MEDIA_CLIENT.CreatePipeline(pipelineName, sourceBucket, targetBucket, capacity)
if err != nil {
    fmt.Printf("create Pipeline error: %+v\n", err)
    return
}
fmt.Println("create pipeline success")
```

## 列出全部pipline

如下代码可以列出用户所有的pipeline

```go
pipelines, err := MEDIA_CLIENT.ListPipelines()
if err != nil {
    fmt.Printf("list Pipeline error: %+v\n", err)
    return
}
fmt.Println("list pipeline success\n")
for _, pipeline := range pipelines.Pipelines {
		fmt.Printf("pipeline: %+v\n", pipeline)
}
```

## 查询指定pipeline

如下代码可以按照pipelineName查询pipeline。

```go
pipelineName := "test"
pipeline, err := MEDIA_CLIENT.GetPipeline(pipelineName)
if err != nil {
    fmt.Printf("list Pipeline error: %+v\n", err)
    return
}
fmt.Println("get pipeline success")
fmt.Printf("pipeline: %+v\n", pipeline)
```

## 删除pipeline

如下代码可以按照pipelineName删除pipeline。

```go
pipelineName := "test"
err := MEDIA_CLIENT.DeletePipeline(pipelineName)
if err != nil {
    fmt.Printf("delete Pipeline error: %+v\n", err)
    return
}
fmt.Println("delete pipeline success")
```

需要注意的是，如果Pipeline有关联的Job未完成，则Pipeline无法被删除，必须等Job执行结束后才能成功删除。

## 更新指定的Pipeline

如下代码可以对指定的pipeline进行更新。

```go
pipelineName := "test"
args, _ := MEDIA_CLIENT.GetPipelineUpdate(pipelineName)
//args := &api.UpdatePipelineArgs{}
args.Description = "update"
args.TargetBucket = "vwdemo"
args.SourceBucket = "vwdemo"

config := &api.UpdatePipelineConfig{}
config.Capacity = 2
config.Notification = "zz"
args.UpdatePipelineConfig = config
err := MEDIA_CLIENT.UpdatePipeline(pipelineName, args)
if err != nil {
    fmt.Printf("update Pipeline error: %+v\n", err)
    return
}
fmt.Println("update pipeline success")
```

# Transcoding-Job转码任务

Transcoding Job(任务)是音视频转码中最基本的执行单元，每个任务将一个原始的音视频资源转码成目标规格的音视频资源。因此，任务和转码的目标是一一对应的，也就是说如果用户需要将一个原始多媒体文件转换成三种目标规格，比如从AVI格式转码成FLV/MP4/HLS格式，那么用户将会需要创建三个任务。

## 创建Transcoding Job

用户在创建转码任务时，需要为转码任务指定所属的Pipeline、所需应用的Preset以及原始音视频资源的BOS Key以及目标音视频资源BOS Key。

如下代码创建一个Job, 并获取新创建的jobID：

```go
pipelineName := "go_sdk_test"
sourceKey := "test.mp4"
targetKey := "test-result.mp4"
presetName := "test_preset"
jobResponse, err := MEDIA_CLIENT.CreateJob(pipelineName, sourceKey, targetKey, presetName)
if err != nil {
    fmt.Printf("create job error: %+v\n", err)
    return
}
fmt.Println("create job success jobId:", jobResponse.JobId)
```

如下代码创建一个支持视频合并、去水印、加水印（Job上而不是Preset上指定watermarkId）的Job, 并获取新创建的jobID：

```go
args := &api.CreateJobArgs{}
args.PipelineName = "go_sdk_test"
source := &api.Source{Clips: &[]api.SourceClip{{
  SourceKey:             "01.mp4",
  EnableDelogo:          false,
  DurationInMillisecond: 6656,
  StartTimeInSecond:     2}}}
args.Source = source
target := &api.Target{}
targetKey := "clips_playback_watermark_delogo_crop2.mp4"
watermarkId := "wmk-xxxx"
target.TargetKey = targetKey
watermarkIdSlice := append(target.WatermarkIds, watermarkId)
target.WatermarkIds = watermarkIdSlice
presetName := "go_test_customize_audio_video"
target.PresetName = presetName

delogoArea := &api.Area{}
delogoArea.X = 10
delogoArea.Y = 10
delogoArea.Width = 30
delogoArea.Height = 40
target.DelogoArea = delogoArea

args.Target = target

jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
if err != nil {
    fmt.Printf("create job error: %+v\n", err)
    return
}
fmt.Println("create job success jobId:", jobResponse.JobId)
```

如下代码创建一个支持视频合并、去水印、加水印、去黑边、插入多样叠加效果（Insert）的Job, 并获取新创建的jobID：











## 列出指定Pipeline的所有Transcoding Job

如下代码通过指定pipelineName查询该Pipeline下的所有Job：

```go
pipelineName := "test"
listTranscodingJobsResponse, err := MEDIA_CLIENT.ListTranscodingJobs(pipelineName)
if err != nil {
    fmt.Printf("list job  error: %+v\n", err)
    return
}
fmt.Printf("list job success : %+v\n", listTranscodingJobsResponse)
```

## 查询指定的Transcoding Job信息

可以通过如下代码通过jobId读取某个Job：

```go
jobId := "job-xxxxxxxxx"
getTranscodingJobResponse, err := MEDIA_CLIENT.GetTranscodingJob(jobId)
if err != nil {
    fmt.Printf("get job error: %+v\n", err)
    return
}
fmt.Printf("get job success : %+v\n", getTranscodingJobResponse)
```

# Preset模板

模板是系统预设的对于一个视频资源在做转码计算时所需定义的集合。用户可以更简便的将一个模板应用于一个和多个视频的转码任务，以使这些任务输出相同规格的目标视频资源。

音视频转码为用户预设了丰富且完备的系统模板，以满足用户对于目标规格在格式、码率、分辨率、加解密、水印等诸多方向上的普遍需求，对于不希望过多了解音视频复杂技术背景的用户来说，是最佳的选择。百度为那些在音视频技术上有着丰富积累的用户，提供了可定制化的转码模板，以帮助他们满足复杂业务条件下的转码需求。

当用户仅需对于音视频的容器格式做变化时，百度提供Transmux模板帮助用户以秒级的延迟快速完成容器格式的转换，比如从MP4转换成HLS，而保持原音视频的属性不变。

## 查询当前用户Preset及所有系统Preset

用户可以通过如下代码查询所有的Preset

```go
listPresetsResponse, err := MEDIA_CLIENT.ListPresets()
if err != nil {
    fmt.Printf("list preset error: %+v\n", err)
    return
}
fmt.Printf("list preset success: %+v\n", listPresetsResponse)
```

## 查询指定的Preset信息

如下代码通过指定pipelineName查询该Pipeline下的所有Job：

```go
presetName := "test"
getPresetResponse, err := MEDIA_CLIENT.GetPreset(preset)
if err != nil {
    fmt.Printf("list preset error: %+v\n", err)
    return
}
fmt.Printf("list preset success: %+v\n", getPresetResponse)
```

## 创建Preset

如果系统预设的Preset无法满足用户的需求，用户可以自定义自己的Preset。根据不同的转码需求，可以使用不同的接口创建Preset。

### 创建仅支持容器格式转换的Preset

如下代码创建仅执行容器格式转换Preset

```go
presetName := "test"
description := "测试创建模板"
container := "mp4"
err := MEDIA_CLIENT.CreatePreset(presetName, description, container)
if err != nil {
    fmt.Printf("create preset error: %+v\n", err)
    return
}
fmt.Println("create preset success")
```

### 创建音频文件的转码Preset，不需要截取片段和加密

如果创建一个不需要截取片段和加密的音频文件转码Preset，可以参考如下代码

```go
preset := &api.Preset{}
preset.PresetName = "go_test_customize"
preset.Description = "自定义创建模板"
preset.Container = "mp3"

audio := &api.Audio{}
audio.BitRateInBps = 256000
preset.Audio = audio

err := MEDIA_CLIENT.CreatePrestCustomize(preset)
if err != nil {
    fmt.Printf("create preset error: %+v\n", err)
    return
}
fmt.Println("create preset success")
```

### 创建音频文件转码Preset，需要设置片段截取属性和加密属性

如果创建一个支持截取片段和加密的音频文件转码Preset，可以参考如下代码

```go
preset := &api.Preset{}
preset.PresetName = "go_test_customize_encryption_clip"
preset.Description = "自定义创建模板"
preset.Container = "mp3"

audio := &api.Audio{}
audio.BitRateInBps = 256000
preset.Audio = audio

clip := &api.Clip{}
clip.StartTimeInSecond = 2
clip.DurationInSecond = 10
preset.Clip = clip

encryption := &api.Encryption{}
encryption.Strategy = "PlayerBinding"
preset.Encryption = encryption

err := MEDIA_CLIENT.CreatePrestCustomize(preset)
err := MEDIA_CLIENT.CreatePrestCustomize(preset)
if err != nil {
    fmt.Printf("create preset error: %+v\n", err)
    return
}
fmt.Println("create preset success")
```

### 创建视频文件转码Preset，不需要截取片段、加密和水印属性

如果创建一个不需要截取片段，加密和水印的视频文件转码Preset，可以参考如下代码

```go
preset := &api.Preset{}
preset.PresetName = "go_test_customize_audio_video"
preset.Description = "自定义创建模板"
preset.Container = "mp4"

audio := &api.Audio{}
audio.BitRateInBps = 256000
preset.Audio = audio

video := &api.Video{}
video.BitRateInBps = 1024000
preset.Video = video

err := MEDIA_CLIENT.CreatePrestCustomize(preset)
if err != nil {
    fmt.Printf("create preset error: %+v\n", err)
    return
}
fmt.Println("create preset success")
```

### 创建视频文件转码Preset，需要设置片段截取、加密和水印属性

如果创建一个需要截取片段，加密和添加水印的视频文件转码Preset，可以参考如下代码

```go
preset := &api.Preset{}
preset.PresetName = "go_test_customize_clp_aud_vid_en_wat"
preset.Description = "自定义创建模板"
preset.Container = "mp4"

clip := &api.Clip{}
clip.StartTimeInSecond = 0
clip.DurationInSecond = 60
preset.Clip = clip

audio := &api.Audio{}
audio.BitRateInBps = 256000
preset.Audio = audio

video := &api.Video{}
video.BitRateInBps = 1024000
preset.Video = video

encryption := &api.Encryption{}
encryption.Strategy = "PlayerBinding"
preset.Encryption = encryption

preset.WatermarkID = "wmk-xxxxxx"

err := MEDIA_CLIENT.CreatePrestCustomize(preset)
if err != nil {
    fmt.Printf("create preset error: %+v\n", err)
    return
}
fmt.Println("create preset success")
```

### 创建Preset，指定所有的参数

如果需要定制所有配置参数，可以参考如下代码

```go
preset := &api.Preset{}
preset.PresetName = "go_test_customize_full_args"
preset.Description = "全参数"
preset.Container = "hls"
preset.Transmux = false

clip := &api.Clip{}
clip.StartTimeInSecond = 0
clip.DurationInSecond = 60
preset.Clip = clip

audio := &api.Audio{}
audio.BitRateInBps = 256000
preset.Audio = audio

video := &api.Video{}
video.BitRateInBps = 1024000
preset.Video = video

encryption := &api.Encryption{}
encryption.Strategy = "PlayerBinding"
preset.Encryption = encryption

water := &api.Watermarks{}
water.Image = []string{"wmk-pc0rdhzbm8ff99qw"}
preset.Watermarks = water

transCfg := &api.TransCfg{}
transCfg.TransMode = "normal"
preset.TransCfg = transCfg

extraCfg := &api.ExtraCfg{}
extraCfg.SegmentDurationInSecond = 6.66
preset.ExtraCfg = extraCfg

err := MEDIA_CLIENT.CreatePrestCustomize(preset)
if err != nil {
    fmt.Printf("create preset error: %+v\n", err)
    return
}
fmt.Println("create preset success")
```

## 更新Preset

用户可以根据模板名更新自己创建的模板：

```go
preset, _ := MEDIA_CLIENT.GetPreset("go_test_customize")
preset.Description = "test update preset"
err := MEDIA_CLIENT.UpdatePreset(preset)
if err != nil {
    fmt.Printf("update preset error: %+v\n", err)
    return
}
fmt.Println("update preset success")
```

# Mediainfo媒体信息

对于BOS中某个Object，可以通过下面代码获取媒体信息

```go
bucket := "bucekt"
key := "key"
info, err := MEDIA_CLIENT.GetMediaInfoOfFile(bucket, key)
if err != nil {
    fmt.Printf("get media information error: %+v\n", err)
    return
}
fmt.Printf("get media information success: %+v\n", info)
```

# Thumbnail-Job缩略图任务

缩略图是图片、视频经压缩方式处理后的小图。因其小巧，加载速度非常快，故用于快速浏览。缩略图任务可用于为BOS中的多媒体资源创建缩略图。

## 创建Thumbnail Job

通过pipeline，BOS Key以及其他配置信息为指定媒体生成缩略图，并获取返回的缩略图任务jobId。可以参考如下代码：

```go
pipelineName := "go_test"
sourcekey := "01.mp4"
target := &api.ThumbnailTarget{}
target.Format = "jpg"
target.SizingPolicy = "keep"
capture := &api.ThumbnailCapture{}
capture.Mode = "manual"
capture.StartTimeInSecond = 0.0
capture.EndTimeInSecond = 5.0
capture.IntervalInSecond = 1.0
createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob(pipelineName, sourcekey, TargetOp(target), CaptureOp(capture))
if err != nil {
    fmt.Printf("create thumbanil job error: %+v\n", err)
    return
}
fmt.Println("create thumbanil job success jobId: ", createJobResponse.JobId)
```

创建去水印的缩略图，可以参考如下代码：

```go
pipelineName := "go_test"
sourcekey := "01.mp4"
target := &api.ThumbnailTarget{}
target.KeyPrefix = "taget_key_prefix_test_delogo3"
delogo := &api.Area{}
delogo.X = 20
delogo.Y = 20
delogo.Height = 50
delogo.Width = 80

createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob(pipelineName, sourcekey, TargetOp(target), DelogoAreaOp(delogo))
if err != nil {
    fmt.Printf("create thumbanil job error: %+v\n", err)
    return
}
fmt.Println("create thumbanil job success jobId: ", createJobResponse.JobId)
```

创建去水印、去黑边的缩略图，可以参考如下代码：

```go
pipelineName := "go_test"
sourcekey := "01.mp4"
target := &api.ThumbnailTarget{}
target.KeyPrefix = "taget_key_prefix_test_delogo_crop"
delogo := &api.Area{}
delogo.X = 20
delogo.Y = 20
delogo.Height = 50
delogo.Width = 80

crop := &api.Area{}
crop.X = 120
crop.Y = 120
crop.Height = 100
crop.Width = 80

createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob(pipelineName, sourcekey,
                                                          TargetOp(target), DelogoAreaOp(delogo), CropOp(crop))
if err != nil {
    fmt.Printf("create thumbanil job error: %+v\n", err)
    return
}
fmt.Println("create thumbanil job success jobId: ", createJobResponse.JobId)
```

创建去水印缩略图任务，其中指定了缩略图格式为jpg、尺寸为与原视频保持一致（keep），抽帧模式（SizingPolicy）为split，根据指定的起止时间和张数截取缩略图，FrameNumber则指定了缩略图张数，代码如下：

```go
pipelineName := "go_test"
sourcekey := "01.mp4"
target := &api.ThumbnailTarget{}
target.Format = "jpg"
target.SizingPolicy = "keep"

capture := &api.ThumbnailCapture{}
capture.Mode = "split"
capture.FrameNumber = 30

delogo := &api.Area{}
delogo.X = 20
delogo.Y = 20
delogo.Height = 50
delogo.Width = 80

createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob(pipelineName, sourcekey,
                                                          TargetOp(target), CaptureOp(capture), DelogoAreaOp(delogo))
if err != nil {
    fmt.Printf("create thumbanil job error: %+v\n", err)
    return
}
fmt.Println("create thumbanil job success jobId: ", createJobResponse.JobId)
```

如果只想创建一个简单的缩略图任务可以参考如下代码：

```go
pipelineName := "go_test"
sourcekey := "01.mp4"
createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob(pipelineName, sourcekey)
if err != nil {
    fmt.Printf("create thumbanil job error: %+v\n", err)
    return
}
fmt.Println("create thumbanil job success jobId: ", createJobResponse.JobId)
```

## 查询指定Thumbnail Job

如果需要获取一个已创建的缩略图任务的信息，可以参考如下代码：

```go
jobId := "job-xxxxxxx"
jobResponse, err := MEDIA_CLIENT.GetThumbanilJob(jobId)
if err != nil {
    fmt.Printf("get thumbanil job error: %+v\n", err)
    return
}
fmt.Printf("get thumbanil job success job: %+v\n", jobResponse)
```

## 查询指定队列的Thumbnail Jobs

如果需要获取一个队列里的全部缩略图任务的信息，可以参考如下代码：

```go
pipelineName := "go_sdk_test"
listThumbnailJobsResponse, err := MEDIA_CLIENT.ListThumbnailJobs(pipelineName)
if err != nil {
    fmt.Printf("list thumbanil job error: %+v\n", err)
    return
}
for _, job := range listThumbnailJobsResponse.Thumbnails {
		fmt.Printf("list thumbanil job success : %+v\n", job)
}
```

# Watermark水印

数字水印是向数据多媒体（如图像、音频、视频信号等）中添加某些数字信息以达到文件真伪鉴别、版权保护等功能。嵌入的水印信息隐藏于宿主文件中，不影响原始文件的可观性和完整性。

用户可以将BOS中的一个Object创建为水印，获得对应的watermarkId。然后在转码任务中将此水印添加到目的多媒体文件。

## 创建水印

如果需要创建一个水印, 指定水印的位置, 并获得水印的唯一ID。其中bucket是水印文件所在bucket名称，key是水印文件在该bucket中的文件名。可以参考如下代码：

```go
args := &api.CreateWaterMarkArgs{}
args.Bucket = "go-test"
args.Key = "01.jpg"
args.HorizontalAlignment = "right"
args.VerticalAlignment = "top"
createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
if err != nil {
    fmt.Printf("create watermark job error: %+v\n", err)
    return
}
fmt.Println("create watermark job success Id: ", createWaterMarkResponse.WatermarkId)
```

如果需要创建一个水印, 指定水印的位置、显示时间段、重复显示次数（动态水印）、自动缩放, 并获得水印的唯一ID，可以参考如下代码：

```go
args := &api.CreateWaterMarkArgs{}
args.Bucket = "go-test"
args.Key = "01.jpg"
args.HorizontalAlignment = "left"
args.VerticalAlignment = "top"
args.HorizontalOffsetInPixel = 20
args.VerticalOffsetInPixel = 10
timeline := &api.Timeline{}
timeline.StartTimeInMillisecond = 1000
timeline.DurationInMillisecond = 3000
args.Timeline = timeline
args.Repeated = 1
args.AllowScaling = true
createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
if err != nil {
    fmt.Printf("create watermark job error: %+v\n", err)
    return
}
fmt.Println("create watermark job success Id: ", createWaterMarkResponse.WatermarkId)
```

## 查询指定水印

如果需要查询已创建的水印，可以参考如下代码：

```go
waterMarkId := "wmk-xxx"
response, err := MEDIA_CLIENT.GetWaterMark(waterMarkId)
if err != nil {
    fmt.Printf("get watermark job error: %+v\n", err)
    return
}
fmt.Printf("get watermark job success: %+v\n", response)
```

## 查询当前用户水印

如果需要查询出本用户所创建的全部水印，可以参考如下代码：

```go
response, err := MEDIA_CLIENT.ListWaterMark()
if err != nil {
    fmt.Printf("get watermark job error: %+v\n", err)
    return
}
for _, watermark := range response.Watermarks {
		fmt.Printf("watermark job: %+v\n", watermark)
}
```

## 删除水印

如果需要删除某个已知watermarkId的水印，可以参考如下代码：

```go
waterMarkId := "wmk-xxx"
err := MEDIA_CLIENT.DeleteWaterMark(waterMarkId)
if err != nil {
    fmt.Printf("delete watermark job error: %+v\n", err)
    return
}
fmt.Println("delete watermark success")
```

# Notification通知

通知功能可以在音视频转码任务状态转换时主动向开发者服务器推送消息。

## 创建通知

如果需要创建通知可以参考如下代码：

```go
name := "test"
endpoint := "http://www.baidu.com"
err := MEDIA_CLIENT.CreateNotification(name, endpoint)
if err != nil {
    fmt.Printf("create notification error: %+v\n", err)
    return
}
fmt.Println("create notification success")
```

## 查询指定通知

如果需要查询已创建的通知，可以参考如下代码：

```go
name := "test"
response, err := MEDIA_CLIENT.GetNotification(test)
if err != nil {
    fmt.Printf("get notification error: %+v\n", err)
    return
}
fmt.Printf("get notification success : %+v\n", response)
```

## 查询当前用户通知

如果需要查询出本用户所创建的全部通知，可以参考如下代码：

```go
response, err := MEDIA_CLIENT.ListNotification()
if err != nil {
    fmt.Printf("list user`s notification error: %+v\n", err)
    return
}
for _, notification := range response.Notifications {
		fmt.Printf("list notification success : %+v\n", notification)
}
```

## 删除通知

如果需要删除某个通知，可以参考如下代码：

```go
name := "test"
err := MEDIA_CLIENT.DeleteNotification(name)
if err != nil {
    fmt.Printf("delete notification error: %+v\n", err)
    return
}
fmt.Println("delete notification success")
```

# 错误处理

GO语言以error类型标识错误，MCP支持两种错误见下表：

| 错误类型        | 说明               |
| --------------- | ------------------ |
| BceClientError  | 用户操作产生的错误 |
| BceServiceError | MCP服务返回的错误  |

用户使用SDK调用MCP相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```go
// MEDIA_CLIENT 为已创建的MCP Client对象
result, err := MEDIA_CLIENT.ListPipelines()
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

客户端异常表示客户端尝试向MCP发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当MCP服务端出现异常时，MCP服务端会返回给用户相应的错误信息，以便定位问题。常见服务端异常可参见[MCP错误码](https://cloud.baidu.com/doc/MCT/s/bjwvz5h3i)

## SDK日志

MCP GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

MCP GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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
log.Debugf("%s", "logging message using the log package in the MCP go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the MCP go sdk")
```
# 版本变更记录

首次发布:

- MCP支持go-sdk啦，现在您可以通过golang调用MCP-SDK服务。当前SDK能力支持pipeline队列操作、Transcoding-Job转码任务操作、Preset模板操作、Thumbnail-Job缩略图任务操作、Watermark水印任务操作、MediaInfo媒资信息操作、Notification通知操作。