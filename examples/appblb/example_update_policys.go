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

func UpdatePolicys() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	appblbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

	AppBlbID := "Your appblb's id"

	args := &appblb.UpdatePolicysArgs{
		Port: 80,
		Type: "HTTP",
		PolicyList: []appblb.PolicyForUpdate{
			appblb.PolicyForUpdate{
				PolicyId:    "Your policy id",
				Priority:    100,
				Description: "policy description for update",
			},
			{
				PolicyId:    "Your policy id",
				Priority:    101,
				Description: "policy description for update",
			},
		},
	}

	if err := appblbClient.UpdatePolicys(AppBlbID, args); err != nil {
		fmt.Println("update policy error: ", err)
		return
	}
}
