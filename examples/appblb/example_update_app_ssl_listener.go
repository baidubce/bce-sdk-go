/*
 * Copyright 2024 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing
 * permissions
 * and limitations under the License.
 */
package appblbexamples

import (
	"github.com/baidubce/bce-sdk-go/services/appblb"
)

func UpdateAppSSLListener() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	appBlbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

	blbID := "Your BlbID" // blb实例ID
	updateAppSSLListenerArgs := &appblb.UpdateAppSSLListenerArgs{
		ListenerPort: 8082,                          // 监听端口
		Scheduler:    "RoundRobin",                  // 调度算法
		CertIds:      []string{"cert-xxxxxxxxxxxx"}, // 证书ID列表
	}

	// 更新appblb ssl监听器
	if err := appBlbClient.UpdateAppSSLListener(blbID, updateAppSSLListenerArgs); err != nil {
		panic(err)
	}
}
