/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// init.go - just import the sub packages

// Package sdk imports all sub packages to build all of them when calling `go install', `go build'
// or `go get' commands.
package sdk

import (
	_ "github.com/baidubce/bce-sdk-go/auth"
	_ "github.com/baidubce/bce-sdk-go/bce"
	_ "github.com/baidubce/bce-sdk-go/http"
	_ "github.com/baidubce/bce-sdk-go/model"
	_ "github.com/baidubce/bce-sdk-go/services/appblb"
	_ "github.com/baidubce/bce-sdk-go/services/bbc"
	_ "github.com/baidubce/bce-sdk-go/services/bcc"
	_ "github.com/baidubce/bce-sdk-go/services/bie"
	_ "github.com/baidubce/bce-sdk-go/services/blb"
	_ "github.com/baidubce/bce-sdk-go/services/bos"
	_ "github.com/baidubce/bce-sdk-go/services/cdn"
	_ "github.com/baidubce/bce-sdk-go/services/cert"
	_ "github.com/baidubce/bce-sdk-go/services/cfc"
	_ "github.com/baidubce/bce-sdk-go/services/dcc"
	_ "github.com/baidubce/bce-sdk-go/services/ddc"
	_ "github.com/baidubce/bce-sdk-go/services/ddc/v2"
	_ "github.com/baidubce/bce-sdk-go/services/eip"
	_ "github.com/baidubce/bce-sdk-go/services/etGateway"
	_ "github.com/baidubce/bce-sdk-go/services/mms"
	_ "github.com/baidubce/bce-sdk-go/services/rds"
	_ "github.com/baidubce/bce-sdk-go/services/scs"
	_ "github.com/baidubce/bce-sdk-go/services/sms"
	_ "github.com/baidubce/bce-sdk-go/services/sts"
	_ "github.com/baidubce/bce-sdk-go/services/vca"
	_ "github.com/baidubce/bce-sdk-go/services/vcr"
	_ "github.com/baidubce/bce-sdk-go/services/vpc"
	_ "github.com/baidubce/bce-sdk-go/services/vpn"
	_ "github.com/baidubce/bce-sdk-go/util"
	_ "github.com/baidubce/bce-sdk-go/util/crypto"
	_ "github.com/baidubce/bce-sdk-go/util/log"
)
