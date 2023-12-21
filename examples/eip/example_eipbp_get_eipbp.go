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
)

func GetEipBp() {
	Init()

	// 要查询eipbp详情的id
	id := "bw-xxxxxxxx"
	eipBpDetail, err := eipClient.GetEipBp(id, "")
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(eipBpDetail, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
