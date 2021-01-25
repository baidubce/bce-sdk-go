# MMS - 多模态媒资检索

## 介绍

- 百度智能云多模态媒资（Multimodal Media Search，简称 MMS）基于视频指纹特征与视频内容理解，实现多模态的搜索能力，主要包含以视频搜视频、以图搜视频、以图搜图等功能，赋予用户多模态的高效、精准、智能的搜索能力。

## 接口调用准备

### Endpoint

- 目前统一为: mms.bj.baidubce.com

### AK/SK

- 要使用百度云 SMS，您需要拥有一个有效的 AK(Access Key ID)和 SK(Secret Access Key)用来进行签名认证。AK/SK 是由系统分配给用户的，均为字符串，用于标识用户，为访问 SMS 做签名验证。
- 可以通过如下步骤获得并了解您的 AK/SK 信息：[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)、[创建 AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

### AK/SK 应用及客户端初始化

```golang
	AK           := ""              // 填充Access Key ID
    SK           := ""              // 填充Secret Access Key
    ENDPOINT     := "http://xxx"    // 填充Endpoint
    CLIENT, err := mms.NewClient(AK, SK, ENDPOINT)  // 初始化客户端
```

## 视频入库

### 接口描述

- 本接口用于向视频库中插入视频特征。
- 入库接口为异步接口，可通过[查询视频入库结果](#查询视频入库结果)接口查询入库结果。或通过通知服务回调结果。

### 请求结构

    PUT /v{version}/videolib/{libName}
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": videoUrl,
        "description": desc,
        "notification": notificationName
    }

### 请求参数

| 参数名称     | 类型   | 是否必需 | 参数位置  | 描述                 |
| ------------ | ------ | -------- | --------- | -------------------- |
| version      | String | 是       | URL 参数  | API 版本号           |
| libName      | String | 是       | URL 参数  | 用户的视频库名称     |
| source       | String | 是       | Body 参数 | 入库视频的 URL       |
| description  | String | 否       | Body 参数 | 用户对此次请求的描述 |
| notification | String | 否       | Body 参数 | 入库结果通知的名称   |

- 注：如使用 notification 参数，需提前配置通知名称及对于的回调地址。

### 响应参数

| 参数名称 | 类型   | 描述     |
| -------- | ------ | -------- |
| status   | String | 请求结果 |

### 请求示例

    PUT /v2/videolib/baiduyun_test
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": "http://test.mp4",
        "description": "test",
        "notification": "notification_name"
    }

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status": "success"
    }

### 通知服务回调结果示例

    {
        "messageId": "a114f8e1-0de0-473f-9f6c-d47e33df5d7d",
        "messageBody": "{\"taskId\":\"n7CCcHIBTmikKXpp-AS8\",\"status\":\"success\",\"source\":\"http://test.mp4\",\"duration\":6.5,\"description\":\"\",\"createTime\":\"2020-06-01T15:32:11Z\",\"startTime\":\"2020-06-01T15:32:11Z\",\"updateTime\":\"2020-06-01T15:32:13Z\",\"finishTime\":\"2020-06-01T15:32:13Z\"}"
    }

## 查询视频入库结果

### 接口描述

- 本接口用于查询视频入库结果。

### 请求结构

    GET /v{version}/videolib/{libName}?source=videoUrl
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 请求参数

| 参数名称 | 类型   | 是否必需 | 参数位置 | 描述             |
| -------- | ------ | -------- | -------- | ---------------- |
| version  | String | 是       | URL 参数 | API 版本号       |
| libName  | String | 是       | URL 参数 | 用户的视频库名称 |
| source   | String | 是       | URL 参数 | 入库视频的 URL   |

### 响应参数

| 参数名称    | 类型   | 描述                                                                                    |
| ----------- | ------ | --------------------------------------------------------------------------------------- |
| status      | String | 入库任务状态，取值为 provision/processing/success/failed，分别为排队中/处理中/成功/失败 |
| description | String | 用户入库请求传入的 description 字段                                                     |
| taskId      | String | 视频入库任务 ID                                                                         |
| source      | String | 入库视频的 URL                                                                          |

### 请求示例

    GET /v2/videolib/baiduyun_test?source=http://test.mp4
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 响应示例

- 入库中

  HTTP/1.1 200 OK
  Content-Type: application/json;charset=UTF-8
  {
  "createTime": "2020-05-13T07:57:25Z",
  "description": "",
  "source": "http://test.mp4",
  "startTime": "2020-05-13T07:57:26Z",
  "status": "processing",
  "taskId": "VUcJDXIBrTeiQx_QzcWe"
  }

- 入库完成

  HTTP/1.1 200 OK
  Content-Type: application/json;charset=UTF-8
  {
  "createTime": "2020-05-13T07:55:55Z",
  "description": "",
  "duration": 6.5,
  "finishTime": "2020-05-13T07:55:57Z",
  "source": "http://test.mp4",
  "startTime": "2020-05-13T07:55:55Z",
  "status": "success",
  "taskId": "VUcJDXIBrTeiQx_QzcWe"
  }

## 图片入库

### 接口描述

- 本接口用于向图片库中插入图片特征。

### 请求结构

    PUT /v{version}/imagelib/{libName}
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": imageUrl,
        "description": desc
    }

### 请求参数

| 参数名称    | 类型   | 是否必需 | 参数位置  | 描述                 |
| ----------- | ------ | -------- | --------- | -------------------- |
| version     | String | 是       | URL 参数  | API 版本号           |
| libName     | String | 是       | URL 参数  | 用户的图片库名称     |
| source      | String | 是       | Body 参数 | 入库图片的 URL       |
| description | String | 否       | Body 参数 | 用户对此次请求的描述 |

### 响应参数

| 参数名称 | 类型   | 描述     |
| -------- | ------ | -------- |
| status   | String | 入库结果 |

### 请求示例

    PUT /v2/imagelib/baiduyun_test
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": "http://test.jpg",
        "description": "test"
    }

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status": "success"
    }

## 删除视频库中的视频

### 接口描述

- 本接口用于删除视频库中某个视频

### 请求结构

    POST /v{version}/videolib/{libName}?deleteVideo=&source=videoUrl
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 请求参数

| 参数名称    | 类型   | 是否必需 | 参数位置 | 描述                             |
| ----------- | ------ | -------- | -------- | -------------------------------- |
| version     | String | 是       | URL 参数 | API 版本号                       |
| libName     | String | 是       | URL 参数 | 用户的视频库名称                 |
| source      | String | 是       | URL 参数 | 要删除视频的 URL（入库时的 URL） |
| deleteVideo | String | 是       | URL 参数 | 标识参数，无内容                 |

### 响应参数

| 参数名称 | 类型   | 描述     |
| -------- | ------ | -------- |
| status   | String | 删除结果 |

### 请求示例

    POST /v2/videolib/baiduyun_test?deleteVideo=&source=http://test.mp4
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status": "success"
    }

## 删除图片库中的图片

### 接口描述

- 本接口用于删除图片库中某张图片

### 请求结构

    POST /v{version}/imagelib/{libName}?deleteImage=&source=imageUrl
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 请求参数

| 参数名称    | 类型   | 是否必需 | 参数位置 | 描述                             |
| ----------- | ------ | -------- | -------- | -------------------------------- |
| version     | String | 是       | URL 参数 | API 版本号                       |
| libName     | String | 是       | URL 参数 | 用户的图片库名称                 |
| source      | String | 是       | URL 参数 | 要删除图片的 URL（入库时的 URL） |
| deleteImage | String | 是       | URL 参数 | 标识参数，无内容                 |

### 响应参数

| 参数名称 | 类型   | 描述     |
| -------- | ------ | -------- |
| status   | String | 删除结果 |

### 请求示例

    POST /v2/imagelib/baiduyun_test?deleteImage=&source=http://test.jpg
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status": "success"
    }

## 视频检索视频

### 接口描述

- 本接口使用视频来检索库中存在的相似视频。
- 本接口为异步接口，可通过[查询视频检索结果](#查询视频检索视频结果)接口查询检索结果。或通过通知服务回调结果。

### 请求结构

    POST /v{version}/videolib/{libName}?searchByVideo
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": videoUrl,
        "description": desc,
        "notification": notificationName
    }

### 请求参数

| 参数名称     | 类型   | 是否必需 | 参数位置  | 描述                 |
| ------------ | ------ | -------- | --------- | -------------------- |
| version      | String | 是       | URL 参数  | API 版本号           |
| libName      | String | 是       | URL 参数  | 用户的视频库名称     |
| source       | String | 是       | Body 参数 | 检索视频的 URL       |
| description  | String | 否       | Body 参数 | 用户对此次请求的描述 |
| notification | String | 否       | Body 参数 | 检索结果通知的名称   |

- 注：如使用 notification 参数，需提前配置通知名称及对于的回调地址。

### 响应参数

| 参数名称 | 类型   | 描述            |
| -------- | ------ | --------------- |
| status   | String | 请求结果        |
| taskId   | String | 视频检索任务 ID |

### 请求示例

    POST /v2/videolib/baiduyun_test?searchByVideo
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": "https://test.mp4",
        "description": "test",
        "notification": "notification_name"
    }

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status": "success",
        "taskId": "VkcZDXIBrTeiQx_QrcXT"
    }

### 通知服务回调结果示例

    {
    "messageId": "360d15ea-14ae-440e-9be1-f1f431beae19",
    "messageBody": "{\"taskId\":\"ybjMcnIBFaqg3FXUVQrA\",\"status\":\"success\",\"lib\":\"video_lib\",\"source\":\"https://test.mp4\",\"duration\":6.5,\"description\":\"\",\"createTime\":\"2020-06-02T02:11:33Z\",\"startTime\":\"2020-06-02T02:12:05Z\",\"updateTime\":\"2020-06-02T02:12:08Z\",\"finishTime\":\"2020-06-02T02:12:08Z\",\"results\":[{\"id\":\"n7CCcHIBTmikKXpp-AS8\",\"name\":\"search_hit.mp4\",\"source\":\"http://hit.mp4\",\"duration\":6.5,\"description\":\"\",\"type\":\"SEARCH_VIDEO_BY_VIDEO\",\"score\":100,\"clips\":[{\"inputStartTime\":0.0,\"inputEndTime\":6.5,\"outputStartTime\":0.0,\"outputEndTime\":6.5}]}]}"
    }

## 查询视频检索视频结果

### 接口描述

- 本接口用于查询视频检索视频任务的结果。

请求结构

    GET /v{version}/videolib/{libName}?searchByVideo&source=videoUrl
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 请求参数

| 参数名称      | 类型   | 是否必需 | 参数位置 | 描述                   |
| ------------- | ------ | -------- | -------- | ---------------------- |
| version       | String | 是       | URL 参数 | API 版本号             |
| libName       | String | 是       | URL 参数 | 用户的视频库名称       |
| source        | String | 是       | URL 参数 | 发起检索任务视频的 URL |
| searchByVideo | String | 是       | URL 参数 | 标识参数，无内容       |

### 响应参数

| 参数名称          | 类型   | 描述                                                               |
| ----------------- | ------ | ------------------------------------------------------------------ |
| status            | String | 任务状态，取值为 processing/success/failed，分别为处理中/成功/失败 |
| lib               | String | 检索的视频库名称                                                   |
| source            | String | 检索视频的 URL                                                     |
| description       | String | 用户传入的请求描述信息                                             |
| results           | List   | 检索视频的结果                                                     |
| +score            | Double | 检索视频的相似度，取值范围为[0, 100]                               |
| +source           | String | 结果视频的 URL                                                     |
| +description      | String | 结果视频的描述                                                     |
| +clips            | List   | 请求成功时才会有此值                                               |
| ++inputStartTime  | Double | 检索视频片段的开始时间，单位：秒                                   |
| ++inputEndTime    | Double | 检索视频片段的结束时间，单位：秒                                   |
| ++outputStartTime | Double | 底库视频片段的开始时间，单位：秒                                   |
| ++outputEndTime   | Double | 底库视频片段的结束时间，单位：秒                                   |
| error             | Object | 请求失败时才会有此值                                               |
| +code             | String | 请求失败时才会有此值，表示错误码                                   |
| +message          | String | 请求失败时才会有此值，表示错误信息                                 |

### 请求示例

    GET /v2/videolib/baiduyun_test?searchByVideo&source=https://test.mp4
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status":"success",
        "lib":"baiduyun_test",
        "source":"https://test.mp4",
        "description":"test",
        "results":[
            {
                "source":"https://test.mp4",
                "description":"test",
                "score":100,
                "clips":[
                    {
                        "inputStartTime":0.08,
                        "inputEndTime":20.16,
                        "outputStartTime":0.08,
                        "outputEndTime":20.16
                    }
                ]
            }
        ]
    }

## 图片检索图片

### 接口描述

- 本接口使用图片来检索库中存在的相似图片。

### 请求结构

    POST /v{version}/imagelib/{libName}?searchByImage
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": imageUrl,
        "description": desc
    }

### 请求参数

| 参数名称    | 类型   | 是否必需 | 参数位置  | 描述                 |
| ----------- | ------ | -------- | --------- | -------------------- |
| version     | String | 是       | URL 参数  | API 版本号           |
| libName     | String | 是       | URL 参数  | 用户的图片库名称     |
| source      | String | 是       | Body 参数 | 检索图片的 URL       |
| description | String | 否       | Body 参数 | 用户对此次请求的描述 |

### 响应参数

| 参数名称     | 类型   | 描述                                           |
| ------------ | ------ | ---------------------------------------------- |
| status       | String | 请求结果                                       |
| lib          | String | 检索的图片库名称                               |
| source       | String | 用户传入的图片 URL                             |
| description  | String | 用户传入的请求描述信息                         |
| results      | List   | 检索图片的结果                                 |
| +distance    | Double | 检索图片的相似度，取值范围为[0, 1]，越小越相似 |
| +source      | String | 结果图片的 URL                                 |
| +description | String | 结果图片的描述                                 |
| error        | Object | 请求失败时才会有此值                           |
| +code        | String | 请求失败时才会有此值，表示错误码               |
| +message     | String | 请求失败时才会有此值，表示错误信息             |

### 请求示例

    POST /v2/imagelib/baiduyun_test?searchByImage
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": "http://test.jpg",
        "description": "nothing to desc"
    }

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status": "success",
        "lib":" baiduyun_test",
        "source": "http://test.jpg",
        "description": "nothing to desc",
        "results": [
            {
                "distance": 0.12,
                "source": "http://test2.jpg",
                "description":"nothing to desc"
            }
        ]
    }

## 图片检索视频

### 接口描述

- 本接口使用图片来检索库中存在的包含相似图片的视频。

### 请求结构

    POST /v{version}/videolib/{libName}?searchByImage
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": imageUrl,
        "description": desc
    }

### 请求参数

| 参数名称    | 类型   | 是否必需 | 参数位置  | 描述                 |
| ----------- | ------ | -------- | --------- | -------------------- |
| version     | String | 是       | URL 参数  | API 版本号           |
| libName     | String | 是       | URL 参数  | 用户的视频库名称     |
| source      | String | 是       | Body 参数 | 检索图片的 URL       |
| description | String | 否       | Body 参数 | 用户对此次请求的描述 |

### 响应参数

| 参数名称     | 类型   | 描述                                               |
| ------------ | ------ | -------------------------------------------------- |
| status       | String | 请求结果                                           |
| lib          | String | 检索的视频库名称                                   |
| source       | String | 用户传入的图片 URL                                 |
| description  | String | 用户传入的请求描述信息                             |
| results      | List   | 检索视频的结果                                     |
| +source      | String | 结果视频的 URL                                     |
| +distance    | Double | 结果视频中命中最相似图片的相似度，取值范围为[0, 1] |
| +description | String | 结果视频的描述                                     |
| +frames      | List   | 结果视频中对应的图片                               |
| ++distance   | Double | 视频中对应图片的相似度，取值范围为[0, 1]           |
| ++timestamp  | Double | 视频中对应图片的时间戳，单位为秒（s）              |
| error        | Object | 请求失败时才会有此值                               |
| +code        | String | 请求失败时才会有此值，表示错误码                   |
| +message     | String | 请求失败时才会有此值，表示错误信息                 |

### 请求示例

    POST /v2/videolib/baiduyun_test?searchByImage
    host: mms.bj.baidubce.com
    Authorization: <bce-authorization-string>
    {
        "source": "http://test.jpg",
        "description": "nothing to desc"
    }

### 响应示例

    HTTP/1.1 200 OK
    Content-Type: application/json;charset=UTF-8
    {
        "status": "success",
        "lib":" baiduyun_test",
        "source": "http://test.jpg",
        "description": "nothing to desc",
        "results": [
            {
                "source": "http://test2.jpg",
                "description": "nothing to desc",
                "distance": 0.12,
                "frames": [
                    {
                        "distance": 0.12,
                        "timestamp": 3.4333
                    }
                ]
            }
        ]
    }
