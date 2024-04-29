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
	"github.com/baidubce/bce-sdk-go/model"

	EIP "github.com/baidubce/bce-sdk-go/services/eip"
)

func CreateEipBp() {
	Init()

	createEipBpArgs := &EIP.CreateEipBpArgs{
		Name:            "test-eipbp",           // EipBp的名称
		Eip:             "x.x.x.x",              // EipBp绑定的EIP地址
		EipGroupId:      "eg-xxxxxxxx",          // EipBp绑定的EIPGroup ID，EIP和EIPGroup不能同时存在，如果同时存在，则EIP的优先级更高
		BandwidthInMbps: 200,                    // EipBp的带宽，单位Mbps，该值与所绑定资源的带宽总和不大于200Mbps
		Type:            "BandwidthPackage",     // EipBp的类型
		AutoReleaseTime: "2023-12-13T20:38:43Z", // EipBp的自动释放时间，格式为UTC时间，如不填写则随所绑定资源的到期一并释放
		Tags: []model.TagModel{ //标签信息
			{
				TagKey:   "tagK",
				TagValue: "tagV",
			},
		},
		ResourceGroupId: "RESG-Xnuw3joXLcy", //资源组ID
	}
	createEipBpResult, err := eipClient.CreateEipBp(createEipBpArgs)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(createEipBpResult, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
