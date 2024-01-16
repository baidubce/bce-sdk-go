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
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/appblb"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充
func CreatePolicysServerGroup() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	appblbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

	AppBlbID := "Your appblb's id"

	args := &appblb.CreatePolicysArgs{
		ListenerPort: 98,     // 对应所在BLB下监听器端口号
		Type:         "HTTP", // 对应所在BLB下监听器协议
		AppPolicyVos: []appblb.AppPolicy{ // 监听器绑定策略列表
			{
				Description:      "bb",          // 策略描述
				Priority:         90,            // 策略优先级
				AppServerGroupId: "sg-8a8xxxx6", // 策略绑定服务器组标识符
				BackendPort:      80,            // 目标端口号
				RuleList: []appblb.AppRule{ // 策略规则列表
					{
						Key:   "*", // 规则的类型
						Value: "*", // 通配符匹配字符串
					},
				},
			},
		},
	}

	if err := appblbClient.CreatePolicys(AppBlbID, args); err != nil {
		fmt.Println("create policy with server group error: ", err)
		return
	}
}
