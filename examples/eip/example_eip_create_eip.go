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
	"github.com/baidubce/bce-sdk-go/model"

	EIP "github.com/baidubce/bce-sdk-go/services/eip"
)

func CreateEip() {
	Init()
	// 创建后付费EIP
	billing := &EIP.Billing{
		PaymentTiming: "Postpaid",    //后付费
		BillingMethod: "ByBandwidth", //按带宽计费
	}
	args := &EIP.CreateEipArgs{
		Name:            "test-eip", //EIP名称
		BandWidthInMbps: 10,         //带宽，单位Mbps
		Billing:         billing,    //订单信息
		Tags: []model.TagModel{ //标签信息
			{
				TagKey:   "tagK",
				TagValue: "tagV",
			},
		},
		ResourceGroupId: "RESG-Xnuw3joXLcy", //资源组ID
	}
	resp, err := eipClient.CreateEip(args)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
