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
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package lbdcexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/model"

	"github.com/baidubce/bce-sdk-go/services/lbdc"
)

func CreateLbdc() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := lbdc.NewClient(ak, sk, endpoint)             // 初始化client

	args := &lbdc.CreateLbdcArgs{
		Name:     "name",
		Type:     "4Layer",
		CcuCount: 4,
		Billing: &lbdc.Billing{
			PaymentTiming: "Prepaid",
			Reservation: &lbdc.Reservation{
				ReservationLength: 1,
			},
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",
				TagValue: "tagValue",
			},
		},
	}
	res, err := client.CreateLbdc(args) // 创建lbdc
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
