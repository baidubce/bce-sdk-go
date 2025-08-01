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

type ListReservedHpasByMakerReq struct {
	ReservedHpasIds    []string `json:"reservedHpasIds,omitempty"`
	Name               string   `json:"name,omitempty"`
	ZoneName           string   `json:"zoneName,omitempty"`
	ReservedHpasStatus string   `json:"reservedHpasStatus,omitempty"`
	AppType            string   `json:"appType,omitempty"`
	HpasId             string   `json:"hpasId,omitempty"`
	Marker             string   `json:"marker,omitempty"`
	MaxKeys            int      `json:"maxKeys,omitempty"`
}
