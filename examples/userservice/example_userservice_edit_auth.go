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

// UserServiceEditAuth 是一个函数，用于编辑用户服务授权。
func UserServiceEditAuth() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	serviceDomain := "your service domain"
	uid := "the uid will be added"

	client, _ := userservice.NewClient(ak, sk, endpoint) // 初始化client

	args := &userservice.UserServiceAuthArgs{
		ClientToken: util.NewUUID(),
		AuthList: []userservice.UserServiceAuthModel{
			{
				Uid:  uid,                          // 将要添加权限的uid
				Auth: userservice.ServiceAuthAllow, // 权限类型，允许或拒绝，对应userservice.ServiceAuthAllow和userservice.ServiceAuthDeny
			}},
	}

	err := client.UserServiceEditAuth(serviceDomain, args)

	if err != nil {
		panic(err)
	}
}
