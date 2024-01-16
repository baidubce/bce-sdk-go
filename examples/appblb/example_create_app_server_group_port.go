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
 * either express or implied. See the License for the specific language governing
 * permissions
 * and limitations under the License.
 */

package appblbexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/appblb"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充
func CreateAppServerGroupPort() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	blbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

	BlbID := "Your app blb's id"
	CreateAppServerGroupPortArgs := &appblb.CreateAppServerGroupPortArgs{
		ClientToken:                 "you client token",
		SgId:                        "Your app server group id",
		Port:                        80,
		Type:                        "TCP",
		HealthCheck:                 "TCP",
		HealthCheckPort:             80,
		HealthCheckTimeoutInSecond:  5,
		HealthCheckIntervalInSecond: 5,
		HealthCheckDownRetry:        3,
		HealthCheckUpRetry:          3,
		HealthCheckNormalStatus:     "http_2xx|http_3xx",
		HealthCheckUrlPath:          "/",
		UdpHealthCheckString:        "",
	}

	result, err := blbClient.CreateAppServerGroupPort(BlbID, CreateAppServerGroupPortArgs)
	if err != nil {
		fmt.Println("create app server group port error : ", err)
		return
	}

	// 打印返回结果
	fmt.Println("id: ", result.Id)
	fmt.Println("name: ", result.Name)
}
