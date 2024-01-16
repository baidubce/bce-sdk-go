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
func DeleteAppServerGroup() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	blbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

	blbId := "Your app blb's id"
	deleteAppServerGroupArgs := &appblb.DeleteAppServerGroupArgs{
		SgId:        "your app server group id",
		ClientToken: "test_token",
	}

	err := blbClient.DeleteAppServerGroup(blbId, deleteAppServerGroupArgs)
	if err != nil {
		fmt.Println("delete app server group error : ", err)
		return
	}
}
