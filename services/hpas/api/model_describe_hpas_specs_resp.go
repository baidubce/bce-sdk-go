/*
 * Copyright 2025 Baidu, Inc.
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

package api

type DescribeHpasSpecsResp struct {
	RequestId string     `json:"requestId,omitempty"`
	HpasSpecs []HpasSpec `json:"hpasSpecs,omitempty"`
}

type HpasSpec struct {
	AppType             string `json:"appType,omitempty"`
	Description         string `json:"description,omitempty"`
	Name                string `json:"name,omitempty"`
	AppPerformanceLevel string `json:"appPerformanceLevel,omitempty"`
	ChargeType          string `json:"chargeType,omitempty"`
	AppImageFamily      string `json:"appImageFamily,omitempty"`
	ZoneName            string `json:"zoneName,omitempty"`
	EniQuota            int    `json:"eniQuota,omitempty"`
	EriQuota            int    `json:"eriQuota,omitempty"`
	EphemeralDiskInGb   int    `json:"ephemeralDiskInGb,omitempty"`
	EphemeralDiskCount  int    `json:"ephemeralDiskCount,omitempty"`
	EphemeralDiskType   string `json:"ephemeralDiskType,omitempty"`
	NicIpv4Quota        int    `json:"nicIpv4Quota,omitempty"`
	NicIpv6Quota        int    `json:"nicIpv6Quota,omitempty"`
	CpuCount        	int    `json:"cpuCount,omitempty"`
	MemoryCapacityInGB  int	   `json:"memoryCapacityInGB,omitempty"`
	GpuCardCount        int    `json:"gpuCardCount,omitempty"`
}
