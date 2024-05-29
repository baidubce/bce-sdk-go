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

package endpointexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/model"

	"github.com/baidubce/bce-sdk-go/services/endpoint"
)

func CreateEndpoint() {
	ak, sk, Endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := endpoint.NewClient(ak, sk, Endpoint)         // 初始化client

	args := &endpoint.CreateEndpointArgs{
		VpcId:       "vpc id",
		SubnetId:    "subnet id",
		Name:        "name",
		Service:     "service",
		Description: "desc",
		Billing: &endpoint.Billing{
			PaymentTiming: "Postpaid",
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",
				TagValue: "tagValue",
			},
		},
	}
	result, err := client.CreateEndpoint(args) // 创建endpoint
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
