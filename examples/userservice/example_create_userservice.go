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
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/userservice"
	"github.com/baidubce/bce-sdk-go/util"
)

// CreateUserService 用于创建用户服务
func CreateUserService() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	client, _ := userservice.NewClient(ak, sk, endpoint) // 初始化client

	createArgs := &userservice.CreateUserServiceArgs{
		ClientToken: util.NewUUID(),
		InstanceId:  "lb-xxxxxxxx",                   //BLB实例ID
		Name:        "test_name",                     //名称
		Description: "test_user_service_description", //UserService备注
		ServiceName: "testUserServiceName",           //UserService名称，系统根据此名称生成服务名
		AuthList: []userservice.UserServiceAuthModel{ //用户权限列表
			{
				Uid:  "12345678901234567890123456789012", //用户ID
				Auth: userservice.ServiceAuthAllow,       //用户权限，允许访问
			}},
	}

	response, err := client.CreateUserService(createArgs)

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
