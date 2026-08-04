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

package kafkaexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/kafka"
)

func CreateTopic(clusterID, topicName string) {
	if kafkaClient == nil {
		if err := Init(); err != nil {
			fmt.Printf("Failed to new kafka client, err: %v.\n", err)
			return
		}
	}
	response, err := kafkaClient.CreateTopic(&kafka.CreateTopicRequest{
		ClusterID:         clusterID,
		TopicName:         topicName,
		PartitionNum:      3,
		ReplicationFactor: 3,
	})
	if err != nil {
		fmt.Printf("Failed to create kafka topic, err: %v.\n", err)
		return
	}
	fmt.Printf("create topic success, topicName: %s\n", response.TopicName)
}
