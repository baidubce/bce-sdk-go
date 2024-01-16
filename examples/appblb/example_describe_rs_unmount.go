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

func DescribeRsUnMount() {
	//设置您的ak、sk和要访问的地域
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"
	client, err := appblb.NewClient(ak, sk, endpoint) // 初始化client
	if err != nil {
		panic(err)
	}
	blbID := "Your BlbID" //LB实例ID
	sgID := "sg-328xxxxx" //服务器组ID
	//发起请求并处理返回或异常
	resp, err := client.DescribeRsUnMount(blbID, sgID)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
