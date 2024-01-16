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
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package blbexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/blb"
)

// DescribeSSLListeners 函数用于获取指定负载均衡实例的SSL监听列表
func DescribeSSLListeners() {
	ak, sk, endpoint := "Your AK", "Your SK", "blb.bj.baidubce.com"
	BlbID := "your blb id"
	sslListenerPort := uint16(80)

	client, _ := blb.NewClient(ak, sk, endpoint) // 初始化client
	args := &blb.DescribeListenerArgs{
		ListenerPort: sslListenerPort, // SSL监听端口
	}
	response, err := client.DescribeSSLListeners(BlbID, args)

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
