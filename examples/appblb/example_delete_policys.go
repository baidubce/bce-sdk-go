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
func DeletePolicys() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	appblbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

	AppBlbID := "Your appblb's id"

	args := &appblb.DeletePolicysArgs{
		Port: 98,
		Type: "HTTP",
		PolicyIdList: []string{
			"policy-c4bxxxx6",
		},
	}

	if err := appblbClient.DeletePolicys(AppBlbID, args); err != nil {
		fmt.Println("delete policy error: ", err)
		return
	}
}
