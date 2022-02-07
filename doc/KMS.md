# 密钥管理服务（KMS）

## 概述

本文档主要介绍 KMS GO SDK 的使用。在使用本文档前，您需要先了解 KMS 的一些基本知识，并已开通了 KMS 服务。若您还不了解 KMS，可以参考[产品描述](https://cloud.baidu.com/doc/KMS/s/2jwvxk0sn)和[操作指南](https://cloud.baidu.com/doc/KMS/s/Hjwvxk5tq)。

## 初始化

### 确认Endpoint

在确认您使用 SDK 时配置的 Endpoint 时，可先阅读开发人员指南中关于[KMS服务域名](https://cloud.baidu.com/doc/KMS/s/vjwvxk7am)的部分，理解 Endpoint 相关的概念。百度云目前开放了多区域支持，请参考[区域选择说明](https://cloud.baidu.com/doc/Reference/s/2jwvz23xx/)。

目前KMS服务支持的地域（Region）有华北-北京、华南-广州、华东-苏州，分别对应的 Region 为 bj、gz、su。

### 获取密钥

要使用百度云 KMS，您需要拥有一个有效的 AK(Access Key ID) 和 SK(Secret Access Key) 用来进行签名认证。AK/SK 是由系统分配给用户的，均为字符串，用于标识用户，为访问 KMS 做签名验证。

可以通过如下步骤获得并了解您的 AK/SK 信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

### 新建 KMS Client

KMS Client是KMS服务的客户端，为开发者与KMS服务进行交互提供了一系列的方法。
通过AK/SK方式访问KMS，用户可以参考如下代码新建一个KMS Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/kms"
)

func main() {
	// 用户的 Access Key ID 和 Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的 Region，可以选择 bj、gz、su 三个地域
	REGION := <region-name>

	// 初始化一个kmsClient
	kmsClient, err := kms.NewClient(AK, SK, ENDPOINT)
}
```

创建完 KMS Client，即可以使用 KMS 提供的一系列接口，[接口文档](https://cloud.baidu.com/doc/KMS/s/Ojwvxk7pl)。使用 KMS 的功能需要创建密钥，可以通过两种方法创建密钥。在公有云平台可以手动创建密钥，也可以通过接口进行创建密钥，即接口文档中的 CMK。

### 使用 KMS Client 加密
根据[接口文档](https://cloud.baidu.com/doc/KMS/s/Ojwvxk7pl)中 Encrypt 接口说明可知，加密的明文长度需小于 4096B，加密时需要对该明文使用 base64 进行编码，方可进行加密。

**示例代码**

```go
keyID := <公有云上创建的 KeyID>
strInBase64 := base64.StdEncoding.EncodeToString([]byte("你要加密的明文"))

encryptRes, err := kmsClient.Encrypt(&EncryptReq{KeyID: keyID, Plaintext: strInBase64})
	
```

### 使用 KMS Client 解密

根据[接口文档](https://cloud.baidu.com/doc/KMS/s/Ojwvxk7pl)中 Decrypt 接口说明可知，解密后获取的是加密的明文的 base64，因此需要使用 base64 进行解码。

**示例代码**

```go
ciphertext := <加密的密文>

decryptRes, err := kmsClient.Decrypt(&DecryptReq{Ciphertext: ciphertext})

str err := base64.StdEncoding.DecodeString(decryptRes.Plaintext)
	
```

## 错误处理

当用户访问出错时，请参考 KMS 返回[错误码](https://cloud.baidu.com/doc/KMS/s/vjwvxk7am)。