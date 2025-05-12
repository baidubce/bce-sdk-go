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
	eipPackage "github.com/baidubce/bce-sdk-go/services/eip"
)

func ListDdos() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := eipPackage.NewClient(ak, sk, endpoint)     // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	listDdosRequest := &eipPackage.ListDdosRequest{
		Ips:     "",
		Type:    "",
		Marker:  "",
		MaxKeys: int32(100),
	}
	result := &eipPackage.ListDdosResponse{}
	result, err = client.ListDdos(listDdosRequest)
	data, e := json.Marshal(result)
	if e != nil {
		fmt.Println("json marshal failed!")
	}
	fmt.Printf("%s", data)
}
