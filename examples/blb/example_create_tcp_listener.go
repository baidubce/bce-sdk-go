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
package blbexamples

import (
	"github.com/baidubce/bce-sdk-go/services/blb"
)

func CreateTCPListener() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	blbClient, _ := blb.NewClient(ak, sk, endpoint) // 初始化client

	blbID := "Your BlbID" // blb实例ID
	createTCPListenerArgs := &blb.CreateTCPListenerArgs{
		ListenerPort: 80,           // 监听端口
		BackendPort:  80,           // 后端端口
		Scheduler:    "RoundRobin", // 调度算法
	}

	// 创建blb tcp监听器
	if err := blbClient.CreateTCPListener(blbID, createTCPListenerArgs); err != nil {
		panic(err)
	}
}
