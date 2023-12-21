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
	"encoding/json"
	"fmt"

	EIP "github.com/baidubce/bce-sdk-go/services/eip"
)

func ListEipGroup() {
	Init()

	listEipGroupArgs := &EIP.ListEipGroupArgs{
		Id:      "",   // 指定ID查询
		Name:    "",   // 指定名称查询
		Marker:  "",   // 分页查询的起始位置
		MaxKeys: 1000, // 分页查询的个数
		Status:  "",   // 指定状态查询
	}
	eipGroupList, err := eipClient.ListEipGroup(listEipGroupArgs)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(eipGroupList, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
