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

func UpdateEipBpName() {
	Init()

	// 更新eipbp name的id
	id := "bw-xxxxxxxx"
	updateEipBpNameArgs := &EIP.UpdateEipBpNameArgs{
		Name: "new-name",
	}
	err := eipClient.UpdateEipBpName(id, updateEipBpNameArgs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Update eipbp name success")
}
