/*
 * Copyright 2026 Baidu, Inc.
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

package besexamples

import bes "github.com/baidubce/bce-sdk-go/services/bes/v3"

var (
	besClient *bes.Client
	AK        = "Your AK"
	SK        = "Your SK"
	ENDPOINT  = "https://bes.bj.baidubce.com"

	// STS 临时凭证，可通过 services/sts 获取
	STS_AK    = "Your STS AK"
	STS_SK    = "Your STS SK"
	STS_TOKEN = "Your STS session token"
)

// Init 直接通过 AK/SK凭证创建客户端
func Init() error {
	client, err := bes.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		return err
	}
	besClient = client
	return nil
}

// InitWithSTS 使用 STS 临时凭证创建客户端。凭证过期后需要用新凭证重新创建客户端。
// 两种方法选其一创建即可
func InitWithSTS() error {
	client, err := bes.NewClientWithSTS(STS_AK, STS_SK, STS_TOKEN, ENDPOINT)
	if err != nil {
		return err
	}
	besClient = client
	return nil
}
