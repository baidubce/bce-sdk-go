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

func DescribeAppUDPListeners() {
    ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

    BlbClient, _ := appblb.NewClient(ak, sk, endpoint) // 初始化client

    BlbID := "blb id" // blb实例id
    describeUDPListenersArgs := &appblb.DescribeAppListenerArgs{
        ListenerPort: 53,    // listener端口号
    }

    response, err := BlbClient.DescribeAppUDPListeners(BlbID, describeUDPListenersArgs) //查询UDP Listeners

    if err != nil {
        panic(err)
    }
                    
    fmt.Println(response)
}
