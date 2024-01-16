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
func DescribePolicysServerGroup() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	appblbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

	AppBlbID := "Your appblb's id"

	args := &appblb.DescribePolicysArgs{
		Port: 98,
		Type: "HTTP",
	}

	result, err := appblbClient.DescribePolicys(AppBlbID, args)
	if err != nil {
		fmt.Println("describe ploicy error: ", err)
		return
	}

	// 获取policy的列表信息
	for _, policy := range result.PolicyList {
		fmt.Println("id : ", policy.Id)
		fmt.Println("appServerGroupId : ", policy.AppServerGroupId)
		fmt.Println("appServerGroupName : ", policy.AppServerGroupName)
		fmt.Println("priority : ", policy.Priority)
	}
}
