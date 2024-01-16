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
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/blb"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充
func DescribeHTTPListeners() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	blbClient, _ := blb.NewClient(ak, sk, endpoint) // 初始化client

	BlbID := "Your blb's id"

	DescribeListenerArgs := &blb.DescribeListenerArgs{
		ListenerPort: 80,   // 要查询的监听器端口
		Marker:       "",   // 分页查询的起始位置
		MaxKeys:      1000, // 每页包含的最大数量，最大数量不超过1000。缺省值为1000
	}

	result, err := blbClient.DescribeHTTPListeners(BlbID, DescribeListenerArgs)
	if err != nil {
		fmt.Println("describe http listeners error: ", err)
		return
	}

	// 获取listener的列表信息
	for _, listener := range result.ListenerList {
		fmt.Println("ListenerPort: ", listener.ListenerPort)
		fmt.Println("Scheduler: ", listener.Scheduler)
	}
}
