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

package blbexamples

import (
    "github.com/baidubce/bce-sdk-go/services/blb"
)

func UpdateLoadBalancer() {
    ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

    BlbClient, _ := blb.NewClient(ak, sk, endpoint) // 初始化client
    
    BlbID := "blb id"                      // blb ID
    updateBlbArgs := &blb.UpdateLoadBalancerArgs{
        ClientToken:  "client token",      // token
        Name:         "Test-BLB",          // blb名称 
        Description:  "blb description",   // blb描述
    }

    err := BlbClient.UpdateLoadBalancer(BlbID, updateBlbArgs) // 更新blb

    if err != nil {
        panic(err)
    }
}
