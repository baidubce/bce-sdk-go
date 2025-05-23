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

func ListEip() {
	Init()

	listEipArgs := &EIP.ListEipArgs{
		IpVersion:    "ipv6",     // 指定 EIP IP类型
		Eip:          "",         // 指定EIP
		InstanceType: "",         // 指定实例类型
		InstanceId:   "",         // 指定实例ID
		Status:       "",         // 指定EIP状态
		EipIds:       []string{}, // 指定EIP短ID列表
		Marker:       "",         // 分页查询起始位置标识符
		MaxKeys:      1000,       // 分页查询每页最大数量
	}
	listEipResult, err := eipClient.ListEip(listEipArgs)

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(listEipResult, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
