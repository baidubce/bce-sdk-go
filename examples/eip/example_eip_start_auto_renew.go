/*
 * Copyright 2017 Baidu, Inc.
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

package eipexamples

import (
	"fmt"

	EIP "github.com/baidubce/bce-sdk-go/services/eip"
)

func StartAutoRenew() {
	Init()

	// 开启自动续费的EIP
	eip := "x.x.x.x"
	startAutoRenewArgs := &EIP.StartAutoRenewArgs{
		AutoRenewTimeUnit: "month", // 续费周期
		AutoRenewTime:     1,       // 续费周期数
		ClientToken:       "",
	}
	err := eipClient.StartAutoRenew(eip, startAutoRenewArgs)
	if err != nil {
		panic(err)
	}

	fmt.Println("StartAutoRenew success")
}
