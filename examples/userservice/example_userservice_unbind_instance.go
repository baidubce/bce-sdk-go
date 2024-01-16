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

package userserviceexamples

import (
	"github.com/baidubce/bce-sdk-go/services/userservice"
	"github.com/baidubce/bce-sdk-go/util"
)

// UserServiceUnBindInstance 从用户服务中解绑实例
func UserServiceUnBindInstance() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	serviceDomain := "Your service domain"

	client, _ := userservice.NewClient(ak, sk, endpoint) // 初始化client

	args := &userservice.UserServiceUnBindArgs{
		ClientToken: util.NewUUID(),
	}

	err := client.UserServiceUnBindInstance(serviceDomain, args)

	if err != nil {
		panic(err)
	}
}
