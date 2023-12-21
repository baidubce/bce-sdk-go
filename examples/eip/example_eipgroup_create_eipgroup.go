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

func CreateEipGroup() {
	Init()

	createEipGroupArgs := &EIP.CreateEipGroupArgs{
		Name:            "test-eipgroup-go", // 共享带宽名称
		EipCount:        5,                  // 共享带宽的IP数量
		BandWidthInMbps: 100,                // 共享带宽的带宽
		Billing: &EIP.Billing{
			PaymentTiming: "Prepaid", // 后付费
			Reservation: &EIP.Reservation{
				ReservationLength:   1,       // 预付费的周期数
				ReservationTimeUnit: "month", // 预付费的周期单位
			},
		},
	}
	eipGroupResult, err := eipClient.CreateEipGroup(createEipGroupArgs)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(eipGroupResult, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
