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

func ListClusters() {
	if besClient == nil {
		if err := Init(); err != nil {
			fmt.Printf("Failed to new bes client, err: %v.\n", err)
			return
		}
	}
	response, err := besClient.ListClusters(&bes.ListClustersRequest{
		PageNo:   1,
		PageSize: 10,
	})
	if err != nil {
		fmt.Printf("Failed to list bes clusters, err: %v.\n", err)
		return
	}
	fmt.Printf("totalCount: %d\n", response.TotalCount)
	for _, item := range response.Clusters {
		fmt.Printf("clusterId: %s, name: %s, status: %s, health: %s, version: %s\n",
			item.ClusterId, item.Name, item.Status, item.ClusterHealth.Status, item.Version)
	}
}
