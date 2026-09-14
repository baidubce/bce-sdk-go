/*
 * Copyright 2026 Baidu, Inc.
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

package besexamples

import (
	"fmt"

	bes "github.com/baidubce/bce-sdk-go/services/bes/v3"
)

func ListIndices(clusterId string) {
	if besClient == nil {
		if err := Init(); err != nil {
			fmt.Printf("Failed to new bes client, err: %v.\n", err)
			return
		}
	}
	includeSystem := false
	response, err := besClient.ListIndices(&bes.ListIndicesRequest{
		ClusterId:     clusterId,
		IncludeSystem: &includeSystem,
		PageNo:        1,
		PageSize:      20,
	})
	if err != nil {
		fmt.Printf("Failed to list bes indices, err: %v.\n", err)
		return
	}
	fmt.Printf("totalCount: %d\n", response.TotalCount)
	for _, item := range response.Indices {
		fmt.Printf("indexName: %s, health: %s, status: %s, shards: %s, replicas: %s, docs: %s, size: %s\n",
			item.IndexName, item.Health, item.Status,
			item.PrimaryShards, item.Replicas, item.DocumentCount, item.StorageSize)
	}
}
