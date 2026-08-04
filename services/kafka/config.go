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

package kafka

import "github.com/baidubce/bce-sdk-go/bce"

// KafkaClientConfiguration keeps the Kafka client configuration surface aligned with the Java SDK.
// It intentionally reuses the shared BCE client configuration without adding Kafka-specific fields.
type KafkaClientConfiguration = bce.BceClientConfiguration
