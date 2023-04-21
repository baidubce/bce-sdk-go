# BVW服务

# 概述

本文档主要介绍videoworks（以下均简称BVW）GO SDK的使用。在使用本文档前，您需要先了解BVW的一些基本知识，并已开通了BVW服务。若您还不了解BVW，可以参考[产品描述](https://cloud.baidu.com/doc/VideoWorks/s/gjys24iww)。

# 初始化

## 确认Endpoint

在确认您使用SDK时配置的Endpoint时，可先阅读开发人员指南中关于[服务域名](https://cloud.baidu.com/doc/VideoWorks/s/Xjyo34h6o)的部分，理解Endpoint相关的概念。

目前BVW服务仅支持“华北-北京”该区域。北京区域：`http://bvw.bj.baidubce.com`，对应信息为：

| 访问区域 | 对应Endpoint        |
| -------- | ------------------- |
| BJ       | bvw.bj.baidubce.com |

## 获取密钥

要使用百度云BVW，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BVW做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## 新建BVW Client

BVW Client是BVW服务的客户端，为开发者与BVW服务进行交互提供了一系列的方法。

### 使用AK/SK新建BVW Client

通过AK/SK方式访问BVW，用户可以参考如下代码新建一个BVW Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/bvw"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个BVWClient
	MEDIA_CLIENT, err := bvw.NewClient(AK, SK, ENDPOINT)
}
```

在上面代码中，`ACCESS_KEY_ID`对应控制台中的“Access Key ID”，`SECRET_ACCESS_KEY`对应控制台中的“Access Key Secret”，获取方式请参考《相关参考 [如何获取AKSK](https://cloud.baidu.com/doc/Reference/s/9jwvz2egb)》。第三个参数`ENDPOINT`支持用户自己指定域名，如果设置为空字符串，会使用默认域名作为BVW的服务地址。

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`http://bvw.bj.baidubce.com`。

# 媒资库

## 普通素材

### 上传素材

用户上传音频/视频/图片到素材库，创作视频时可从素材中心导入

```go
req := &api.MatlibUploadRequest{}
req.Key = "dog.mp4"
req.Bucket = "bvw-console"
req.Title = "split_result_267392.mp4"
req.MediaType = "video"
uploadResponse, err := BVW_CLIENT.UploadMaterial(req)
if err != nil {
    fmt.Printf("upload error: %+v\n", err)
    return
}
fmt.Printf("upload success : %+v\n", uploadResponse)
```

> 注意：这里上传的媒资类型需和MediaType严格对应，MediaType可选：video：视频， audio：音频，image：图片

### 查询素材

使用如下代码可以从媒资库查询一个素材。

```go
materialId := "d9b9f08ef1e0a28967fa0f7b5819db30"
materialGetResponse, err := BVW_CLIENT.GetMaterial(materialId)
if err != nil {
    fmt.Printf("get material error: %+v\n", err)
    return
}
fmt.Printf("get material success : %+v\n", materialGetResponse)
```

### 搜索素材

使用如下代码可以从媒资库搜索指定条件的素材。

```go
req := &api.MaterialSearchRequest{}
req.Size = 5
req.PageNo = 1
req.Status = "FINISHED"
req.MediaType = "video"
req.Begin = "2023-01-11T16:00:00Z"
req.End = "2023-04-12T15:59:59Z"
materialResp, err := BVW_CLIENT.SearchMaterial(req)
if err != nil {
    fmt.Printf("search material error: %+v\n", err)
    return
}
fmt.Printf("search material success : %+v\n", materialResp)
```

### 删除素材

使用如下代码可以从媒资库删除一个素材。

```go
materialId := "d9b9f08ef1e0a28967fa0f7b5819db30"
err := BVW_CLIENT.DeleteMaterial(materialId)
if err != nil {
    fmt.Printf("delete material error: %+v\n", err)
    return
}
fmt.Println("delete material success")
```

## 预置素材

预置素材分为音乐/贴图/背景/字幕/转场，除了系统自带的预置素材外，用户可以自定义音乐/贴图预置素材。创作视频时，不需要导入，可直接使用

### 上传预置素材

使用如下代码可以上传一个音乐/贴图预置素材到媒资库。

```go
req := &api.MatlibUploadRequest{}
req.Key = "item2.jpeg"
req.Bucket = "bvw-console"
req.Title = "item2.jpeg"
req.MediaType = "image"
type := "PICTURE"
uploadResponse, err := BVW_CLIENT.UploadMaterialPreset(type, req)
if err != nil {
    fmt.Printf("upload preset material error: %+v\n", err)
    return
}
fmt.Println("upload preset success : %+v\n", uploadResponse)
```

> 注意：预置素材类型type，只支持2种类型自定义：MUSIC：音乐，PICTURE：贴图。

### 搜索预置素材

使用如下代码可以查询一个指定预置素材。

```go
req := &api.MaterialPresetSearchRequest{}
req.PageSize = "10"
req.Status = "FINISHED"
req.PageNo = "1"
req.SourceType = "USER"
req.MediaType = "PICTURE"
response, err := BVW_CLIENT.SearchMaterialPreset(req)
if err != nil {
    fmt.Printf("search preset material error: %+v\n", err)
    return
}
fmt.Println("search preset success : %+v\n", response)
```

### 查询预置素材

使用如下代码可以从媒资库搜索指定条件的预置素材。

```go
id := "cc0aabdc71421abaa17e80a26caa009f"
response, err := BVW_CLIENT.GetMaterialPreset(id)
if err != nil {
    fmt.Printf("get preset material error: %+v\n", err)
    return
}
fmt.Println("get preset success : %+v\n", response)
```

### 删除预置素材

使用如下代码可以删除一个指定预置素材。

```go
id := "cc0aabdc71421abaa17e80a26caa009f"
err := BVW_CLIENT.DeleteMaterialPreset(id)
if err != nil {
    fmt.Printf("delete preset material error: %+v\n", err)
    return
}
fmt.Println("delete preset material success")
```

## 媒资库设置

媒资库相关操作需要预先设置bucket。

### 创建媒资库设置

使用如下代码可以创建媒资库设置，设置媒资库交互中使用的默认bucket。

```go
request := &api.MatlibConfigBaseRequest{
  Bucket: "go-test"}
response, err := BVW_CLIENT.CreateMatlibConfig(request)
if err != nil {
    fmt.Printf("create config error: %+v\n", err)
    return
}
fmt.Println("create config success : %+v\n", response)
```

### 更新媒资库设置

使用如下代码可以更新媒资库设置，设置媒资库交互中使用的默认bucket。

```go
err := BVW_CLIENT.UpdateMatlibConfig(&api.MatlibConfigUpdateRequest{Bucket: "go-test"})
if err != nil {
    fmt.Printf("update config error: %+v\n", err)
    return
}
fmt.Println("update config success")
```

### 查询媒资库设置

使用如下代码可以查询媒资库设置，设置媒资库交互中使用的默认bucket。

```go
response, err := BVW_CLIENT.GetMatlibConfig()
if err != nil {
    fmt.Printf("get config error: %+v\n", err)
    return
}
fmt.Println("get config success : %+v\n", response)
```

# 快编

## 创建草稿

在进行视频合成前需要先创建草稿，以获取任务id用于后续视频合成任务

```go
response, err := BVW_CLIENT.CreateDraft(&api.CreateDraftRequest{
		Duration: "0",
		Titile:   "testCreateDraft3",
		Ratio:    "hori16x9"})
if err != nil {    
  fmt.Printf("create draft error: %+v\n", err)    
  return
}
fmt.Println("create draft success taskId :", response.Id)
```

## 更新草稿

如下代码可更新所创建的草稿

```go
var request api.MatlibTaskRequest
jsonStr := "{\"title\":\"updatesucess\",\"draftId\":\"1017890\",\"timeline\":{\"timeline\":{\"video\":[{\"key\"" +
":\"name\",\"isMaster\":true,\"list\":[{\"type\":\"video\",\"start\":0,\"end\":5.859375,\"showStart\":0," +
"\"showEnd\":5.859375,\"xpos\":0,\"ypos\":0,\"width\":1280,\"height\":720,\"duration\":5.859375,\"extInfo" +
"\":{\"style\":\"\",\"lightness\":0,\"gifMode\":0,\"contrast\":0,\"saturation\":0,\"hue\":0,\"speed\":1" +
",\"transitionStart\":\"\",\"transitionEnd\":\"black\",\"transitionDuration\":1,\"volume\":1,\"rotate\":0," +
"\"mirror\":\"\",\"blankArea\":\"\"},\"mediaInfo\":{\"fileType\":\"video\",\"sourceType\":\"USER\",\"" +
"sourceUrl\":\"https://bj.bcebos.com/v1/bvw-console/360p/dog.mp4?x-bce-security-token=" +
"ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7x" +
"G%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hb" +
"SJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRp" +
"OC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ" +
"%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2" +
"F%2Fb227edbf73344bdfc9fed00ba491c5c0c3abe229792d7b3d026604cfbe541b68\",\"audioUrl\":\"" +
"https://bj.bcebos.com/v1/bvw-console/audio/dog.mp3?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjY" +
"yYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ" +
"5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3" +
"OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzT" +
"Cgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=" +
"bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F3dcc823c9497aca1154f" +
"f0007eca86af4e682363c3ceddba0b3c74ca14e2d154\",\"bucket\":\"bvw-console\",\"key\":\"dog.mp4\",\"audioKey\":" +
"\"audio/dog.mp3\",\"coverImage\":\"https://bj.bcebos.com/v1/bvw-console/thumbnail/dog00000500.jpg" +
"?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZp" +
"Ppsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQz" +
"FlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6P" +
"z36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQj" +
"Wm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cb" +
"c0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F21a92744dfb9fc3f46745e75d095da327bb04677f9028fb85e00ff5dc7df6" +
"daf\",\"duration\":18.73,\"width\":1920,\"height\":1080,\"status\":\"FINISHED\",\"name\":\"dog.mp4\"," +
"\"thumbnailPrefix\":\"\",\"thumbnailKeys\":[\"thumbnail/dog00000500.jpg\"],\"mediaId\":\"" +
"1f10ce0db10b8eb5b2f2755daf544900\",\"offstandard\":false},\"uid\":\"a081f1c6-9dc9-4e7b-a00e-6eb5217e771d" +
"\"},{\"type\":\"video\",\"start\":5.859375,\"end\":18.73,\"showStart\":5.859375,\"showEnd\":18.73,\"xpos" +
"\":0,\"ypos\":0,\"width\":1280,\"height\":720,\"duration\":12.870625,\"extInfo\":{\"style\":\"\",\"lightness" +
"\":0,\"gifMode\":0,\"contrast\":0,\"saturation\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"black\",\"" +
"transitionEnd\":\"\",\"transitionDuration\":1,\"volume\":1,\"rotate\":0,\"mirror\":\"\",\"blankArea\":\"\"}," +
"\"mediaInfo\":{\"fileType\":\"video\",\"sourceType\":\"USER\",\"sourceUrl\":" +
"\"https://bj.bcebos.com/v1/bvw-console/360p/dog.mp4?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyN" +
"jk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9" +
"RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7D" +
"GO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df" +
"5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%" +
"3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2Fb" +
"227edbf73344bdfc9fed00ba491c5c0c3abe229792d7b3d026604cfbe541b68\",\"audioUrl\":" +
"\"https://bj.bcebos.com/v1/bvw-console/audio/dog.mp3?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjY" +
"yYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5" +
"Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3O" +
"HDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCg" +
"h8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=" +
"bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F3dcc823c9497aca1" +
"154ff0007eca86af4e682363c3ceddba0b3c74ca14e2d154\",\"bucket\":\"bvw-console\",\"key\":\"dog.mp4\",\"" +
"audioKey\":\"audio/dog.mp3\",\"coverImage\":\"https://bj.bcebos.com/v1/bvw-console/thumbnail/dog00000500" +
".jpg?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZp" +
"Ppsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQz" +
"FlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz" +
"36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm" +
"231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0" +
"fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F21a92744dfb9fc3f46745e75d095da327bb04677f9028fb85e00ff5dc7df6da" +
"f\",\"duration\":18.73,\"width\":1920,\"height\":1080,\"status\":\"FINISHED\",\"name\":\"dog.mp4\",\"thumb" +
"nailPrefix\":\"\",\"thumbnailKeys\":[\"thumbnail/dog00000500.jpg\"],\"mediaId\":\"1f10ce0db10b8eb5b2f2755d" +
"af544900\",\"offstandard\":false},\"uid\":\"70af482e-0bf5-4c38-8f05-03c9e3b8ae03\"}],\"unlinkMaster\":true" +
"}],\"audio\":[{\"key\":\"\",\"isMaster\":false,\"list\":[{\"start\":0,\"end\":155.99,\"showStart\":" +
"0.234375,\"showEnd\":156.224375,\"duration\":155.99,\"xpos\":0,\"ypos\":0,\"hidden\":false,\"mediaInfo\"" +
":{\"fileType\":\"audio\",\"sourceUrl\":\"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset" +
"/music/audio/%E5%8F%A4%E9%A3%8E%E9%A3%98%E6%89%AC.mp3?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772" +
"a1409590%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F2ddcf78c92de29ae3d7c3166a4e17e7c5d07fa38dcefd24c29c4f4d5b" +
"5ba46fe\",\"audioUrl\":\"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset/music/audio/" +
"%E5%8F%A4%E9%A3%8E%E9%A3%98%E6%89%AC.mp3?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772a1409590%2F2" +
"023-04-14T07%3A16%3A35Z%2F86400%2F%2F2ddcf78c92de29ae3d7c3166a4e17e7c5d07fa38dcefd24c29c4f4d5b5ba46fe\"," +
"\"bucket\":\"videoworks-system-preprocess\",\"key\":\"systemPreset/music/古风飘扬.aac\",\"audioKey\":\"" +
"systemPreset/music/audio/古风飘扬.mp3\",\"coverImage\":\"\",\"duration\":155.99,\"name\":\"\",\"thumbnailList" +
"\":[],\"mediaId\":\"\",\"offstandard\":false},\"type\":\"audio\",\"uid\":\"bd52be7f-1f19-4368-8c41-44e991af" +
"8164\",\"name\":\"古风飘扬\",\"extInfo\":{\"style\":\"\",\"lightness\":0,\"gifMode\":0,\"contrast\":0,\"" +
"saturation\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"\",\"transitionEnd\":\"\",\"transitionDuration" +
"\":1,\"volume\":1,\"rotate\":0},\"boxDataLeft\":4,\"dragBoxWidth\":1996.6720000000003,\"lineType\":\"audio" +
"\"}]}],\"subtitle\":[{\"key\":\"\",\"list\":[{\"duration\":3,\"hidden\":false,\"name\":\"time-place\",\"" +
"tagExtInfo\":{\"marginBottom\":0,\"textFadeIn\":1,\"textFadeOut\":1,\"textOutMaskDur\":1},\"showStart\":" +
"5.859375,\"showEnd\":8.859,\"templateId\":\"6764ce3331ea7e406e4ab4475d1dff18\",\"type\":\"subtitle\",\"uid" +
"\":\"5aaa35f4-8fae-4b8c-b7ed-761a54550244\",\"xpos\":\"0\",\"ypos\":\"309\",\"config\":[{\"alpha\":0,\"" +
"fontColor\":\"#ffffff\",\"fontSize\":50,\"fontStyle\":\"normal\",\"fontType\":\"方正时代宋 简 Extrabold\"" +
",\"lineHeight\":1.2,\"name\":\"时间\",\"text\":\"haha\",\"backgroundColor\":\"#2468F2\",\"backgroundAlpha" +
"\":0,\"fontx\":0,\"fonty\":0,\"invisible\":false},{\"alpha\":0,\"fontColor\":\"#000000\",\"fontSize\":50," +
"\"fontStyle\":\"normal\",\"fontType\":\"方正时代宋 简 Extrabold\",\"lineHeight\":1.2,\"name\":\"地点\",\"" +
"text\":\"cd\",\"backgroundColor\":\"#ffffff\",\"backgroundAlpha\":0,\"fontx\":0,\"fonty\":0,\"invisible\"" +
":false}],\"boxDataLeft\":76,\"dragBoxWidth\":38.400000000000006,\"lineType\":\"subtitle\"}],\"master\":" +
"false}],\"sticker\":[{\"key\":\"\",\"isMaster\":false,\"list\":[{\"showStart\":0,\"showEnd\":3,\"duration\"" +
":3,\"xpos\":0,\"ypos\":0,\"width\":215,\"height\":120.9375,\"hidden\":false,\"mediaInfo\":{\"sourceUrl\":\"" +
"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset/picture/%E9%9D%A2%E5%8C%85%E3%80%81%E8" +
"%82%A0.png?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772a1409590%2F2023-04-14T07%3A16%3A35Z%2F8640" +
"0%2F%2F84867e4874cc94eb374898017cb4367ed8c24a5750a9d8ebd14d8ca989cf2e53\",\"audioUrl\":\"" +
"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset/picture/audio/%E9%9D%A2%E5%8C%85%E3%80%" +
"81%E8%82%A0.mp3?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772a1409590%2F2023-04-14T07%3A16%3A35Z%2F" +
"86400%2F%2F1807d3005f5f8fc270fa17238ec550c20162252c19952e6a45ae497ef7148086\",\"bucket\":\"" +
"videoworks-system-preprocess\",\"key\":\"systemPreset/picture/面包、肠.png\",\"audioKey\":\"systemPreset" +
"/picture/audio/面包、肠.mp3\",\"coverImage\":\"\",\"width\":215,\"height\":120,\"name\":\"\",\"thumbnailList" +
"\":[],\"mediaId\":\"\",\"offstandard\":false},\"type\":\"image\",\"uid\":\"e419b583-aedb-4265-" +
"9a67-7e21fc621f85\",\"name\":\"面包、肠\",\"extInfo\":{\"style\":\"\",\"lightness\":0,\"gifMode\":0," +
"\"contrast\":0,\"saturation\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"\",\"transitionEnd\":\"\"," +
"\"transitionDuration\":1,\"volume\":1,\"rotate\":0},\"lineType\":\"sticker\",\"boxDataLeft\":1,\"" +
"dragBoxWidth\":38.400000000000006}]}]}},\"ratio\":\"hori16x9\",\"resourceList\":[{\"id\":\"1f10ce0db10" +
"b8eb5b2f2755daf544900\",\"userId\":\"e7e47aa53fbb47dfb1e4c86424bb7ad3\",\"mediaType\":\"video\",\"" +
"sourceType\":\"USER\",\"status\":\"FINISHED\",\"title\":\"dog.mp4\",\"sourceUrl\":\"https://bj.bcebos.com" +
"/v1/bvw-console/dog.mp4?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM" +
"1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Z" +
"zgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkh" +
"MLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7" +
"D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1" +
"a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F4a3c399db912e35f6b6c008faffebe5752c47e36fcd21d4bf03" +
"bc908c3a29e5e\",\"sourceUrl360p\":\"https://bj.bcebos.com/v1/bvw-console/360p/dog.mp4?x-bce-security-toke" +
"n=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZT" +
"NAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwb" +
"rCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F" +
"9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujH" +
"KfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07" +
"%3A16%3A35Z%2F86400%2F%2Fb227edbf73344bdfc9fed00ba491c5c0c3abe229792d7b3d026604cfbe541b68\",\"audioUrl\"" +
":\"https://bj.bcebos.com/v1/bvw-console/audio/dog.mp3?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1Z" +
"jYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8" +
"sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46Nx" +
"S3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOF" +
"zTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorizatio" +
"n=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F3dcc823c9497aca1" +
"154ff0007eca86af4e682363c3ceddba0b3c74ca14e2d154\",\"thumbnailList\":[\"https://bj.bcebos.com/v1/bvw-cons" +
"ole/thumbnail/dog00000500.jpg?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAA" +
"BgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9" +
"xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSIS" +
"jn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCK" +
"r%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9" +
"411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F21a92744dfb9fc3f46745e75d095da327bb04677f902" +
"8fb85e00ff5dc7df6daf\"],\"subtitleUrls\":[],\"createTime\":\"2023-04-11 16:55:32\",\"updateTime\":\"2023-0" +
"4-11 16:55:43\",\"duration\":18.73,\"height\":1080,\"width\":1920,\"fileSizeInByte\":8948434,\"thumbnailK" +
"eys\":[\"thumbnail/dog00000500.jpg\"],\"subtitles\":[\"\"],\"bucket\":\"bvw-console\",\"key\":\"dog.mp4\"" +
",\"key360p\":\"360p/dog.mp4\",\"key720p\":\"720p/dog.mp4\",\"audioKey\":\"audio/dog.mp3\",\"used\":true}]" +
",\"coverBucket\":\"bvw-console\",\"coverKey\":\"thumbnail/dog00000500.jpg\"}"
jsonErr := json.Unmarshal([]byte(jsonStr), &request)
ExpectEqual(t.Errorf, jsonErr, nil)
err := BVW_CLIENT.UpdateDraft(1017890, &request)
if err != nil {    
  fmt.Printf("update draft error: %+v\n", err)    
  return
}
fmt.Println("update draft success")
```

## 获取草稿列表

使用如下代码可以获取所创建了所有快编草稿

```go
response, err := BVW_CLIENT.GetDraftList(&api.DraftListRequest{
  PageNo:    1,
  PageSize:  20,
  BeginTime: "2023-01-11T16:00:00Z",
  EndTime:   "2023-04-12T15:59:59Z"})
if err != nil {
    fmt.Printf("get draft list error: %+v\n", err)
    return
}
fmt.Println("get draft list success : %+v\n", response)
```

## 获取单条草稿

如下代码可以通过快编任务id获取草稿信息

```go
taskId := 123456
response, err := BVW_CLIENT.GetSingleDraft(taskId)
if err != nil {
    fmt.Printf("get draft error: %+v\n", err)
    return
}
fmt.Println("get draft success : %+v\n", response)
```

## 视频合成

视频合成功能支持将一段视频编辑的Timeline（不同媒体分类组成的时间轴数据）编码合成输出。

使用如下代码可以发起合成。

```go
var request api.VideoEditCreateRequest
jsonStr := "{\"title\":\"新建作品-202304141603\",\"taskId\":\"1017895\",\"bucket\":\"vwdemo\",\"cmd\":{\"timeline" +
"\":{\"video\":[{\"list\":[{\"type\":\"video\",\"start\":0,\"end\":7.65625,\"showStart\":0,\"showEnd\":" +
"7.65625,\"xpos\":0,\"ypos\":0,\"width\":1,\"height\":1,\"duration\":7.65625,\"extInfo\":{\"style\":\"\"" +
",\"lightness\":0,\"contrast\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"\",\"transitionEnd\":\"black\"" +
",\"transitionDuration\":1,\"mirror\":\"\",\"rotate\":0,\"blankArea\":\"\",\"volume\":1},\"mediaInfo\":{\"" +
"mediaId\":\"1f10ce0db10b8eb5b2f2755daf544900\",\"key\":\"dog.mp4\",\"bucket\":\"bvw-console\",\"fileType\"" +
":\"video\",\"width\":1920,\"height\":1080}},{\"type\":\"video\",\"start\":7.65625,\"end\":18.73,\"showStart" +
"\":7.65625,\"showEnd\":18.73,\"xpos\":0,\"ypos\":0,\"width\":1,\"height\":1,\"duration\":11.07375,\"" +
"extInfo\":{\"style\":\"\",\"lightness\":0,\"contrast\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"black\"" +
",\"transitionEnd\":\"\",\"transitionDuration\":1,\"mirror\":\"\",\"rotate\":0,\"blankArea\":\"\",\"volume\":" +
"1},\"mediaInfo\":{\"mediaId\":\"1f10ce0db10b8eb5b2f2755daf544900\",\"key\":\"dog.mp4\",\"bucket\":\"" +
"bvw-console\",\"fileType\":\"video\",\"width\":1920,\"height\":1080}}]}],\"audio\":[{\"list\":[{\"name\":" +
"\"古风飘扬\",\"start\":0,\"end\":155.99,\"duration\":155.99,\"showStart\":0.078125,\"showEnd\":156.068125,\"" +
"uid\":\"cc8c1ecc-fcd3-493d-be5b-8cce8c15ed15\",\"extInfo\":{\"volume\":\"1.0\",\"transitions\":[]},\"" +
"mediaInfo\":{\"fileType\":\"audio\",\"key\":\"systemPreset/music/古风飘扬.aac\",\"bucket\":\"videoworks-" +
"system-preprocess\",\"name\":\"古风飘扬\"}}]}],\"subtitle\":[{\"list\":[{\"templateId\":\"6764ce3331ea7e406e" +
"4ab4475d1dff18\",\"showStart\":0,\"showEnd\":3,\"duration\":3,\"uid\":\"05e59686-b23e-4b33-96d7-040eab63" +
"85b6\",\"tag\":\"time-place\",\"xpos\":0,\"ypos\":0.431,\"config\":[{\"text\":\"时间\",\"fontSize\":50,\"" +
"fontType\":\"方正时代宋 简 Extrabold\",\"fontColor\":\"#ffffff\",\"alpha\":0,\"fontStyle\":\"normal\",\"" +
"backgroundColor\":\"#2468F2\",\"backgroundAlpha\":0,\"fontx\":0.039,\"fonty\":0.028,\"rectx\":0,\"recty\"" +
":0.431,\"rectWidth\":0.156,\"rectHeight\":0.139},{\"text\":\"地点\",\"fontSize\":50,\"fontType\":\"" +
"方正时代宋 简 Extrabold\",\"fontColor\":\"#000000\",\"alpha\":0,\"fontStyle\":\"normal\",\"backgroundColor" +
"\":\"#ffffff\",\"backgroundAlpha\":0,\"fontx\":0.039,\"fonty\":0.028,\"rectx\":0.156,\"recty\":0.431,\"" +
"rectWidth\":0.156,\"rectHeight\":0.139}],\"tagExtInfo\":{\"glExtInfo\":{}}}]}],\"sticker\":[{\"list\":" +
"[{\"type\":\"image\",\"showStart\":0,\"showEnd\":3,\"duration\":3,\"xpos\":0,\"ypos\":0,\"width\":0.168" +
",\"height\":0.168,\"extInfo\":{},\"mediaInfo\":{\"key\":\"systemPreset/picture/面包、肠.png\",\"bucket\":" +
"\"videoworks-system-preprocess\",\"width\":215,\"height\":120.9375,\"fileType\":\"image\"}}]}]}},\"extInfo" +
"\":{\"aspect\":\"hori16x9\",\"resolutionRatio\":\"v720p\"}}"
jsonErr := json.Unmarshal([]byte(jsonStr), &request)
fmt.Println("jsonError: ", jsonErr)
t.Logf("%+v", &request)
response, err := BVW_CLIENT.CreateVideoEdit(&request)
if err != nil {
    fmt.Printf("create edit job error: %+v\n", err)
    return
}
fmt.Println("create edit job success : %+v\n", response)
```

> 上述代码模拟了前端请求后端数据结构进行视频合成任务，若不想采用上述方式也可自行构建VideoEditCreateRequest结构体进行视频合成任务。

## 查询视频合成结果

使用如下代码可以查询视频合成结果。

```json
editId := 123456
response, err := BVW_CLIENT.PollingVideoEdit(editId)
if err != nil {
    fmt.Printf("get edit job error: %+v\n", err)
    return
}
fmt.Println("get edit job success : %+v\n", response)
```

# 错误处理

GO语言以error类型标识错误，BVW支持两种错误见下表：

| 错误类型        | 说明               |
| --------------- | ------------------ |
| BceClientError  | 用户操作产生的错误 |
| BceServiceError | BVW服务返回的错误  |

用户使用SDK调用BVW相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误进行处理。实例如下：

```go
// MEDIA_CLIENT 为已创建的BVW Client对象
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

客户端异常表示客户端尝试向BVW发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回BceClientError。

## 服务端异常

当BVW服务端出现异常时，BVW服务端会返回给用户相应的错误信息，以便定位问题。

## SDK日志

BVW GO SDK支持六个级别、三种输出（标准输出、标准错误、文件）、基本格式设置的日志模块，导入路径为`github.com/baidubce/bce-sdk-go/util/log`。输出为文件时支持设置五种日志滚动方式（不滚动、按天、按小时、按分钟、按大小），此时还需设置输出日志文件的目录。

### 默认日志

BVW GO SDK自身使用包级别的全局日志对象，该对象默认情况下不记录日志，如果需要输出SDK相关日志需要用户自定指定输出方式和级别，详见如下示例：

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

```
1. 日志默认输出级别为`DEBUG`
2. 如果设置为输出到文件，默认日志输出目录为`/tmp`，默认按小时滚动
3. 如果设置为输出到文件且按大小滚动，默认滚动大小为1GB
4. 默认的日志输出格式为：`FMT_LEVEL, FMT_LTIME, FMT_LOCATION, FMT_MSG`
```

### 项目使用

该日志模块无任何外部依赖，用户使用GO SDK开发项目，可以直接引用该日志模块自行在项目中使用，用户可以继续使用GO SDK使用的包级别的日志对象，也可创建新的日志对象，详见如下示例：

```go
// 直接使用包级别全局日志对象（会和GO SDK自身日志一并输出）
log.SetLogHandler(log.STDERR)
log.Debugf("%s", "logging message using the log package in the BVW go sdk")

// 创建新的日志对象（依据自定义设置输出日志，与GO SDK日志输出分离）
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the BVW go sdk")
```

# 版本变更记录

首次发布:

- BVW支持go-sdk啦，现在您可以通过golang调用BVW-SDK服务。当前SDK能力支持快编（视频编辑合成）服务、媒资库服务。