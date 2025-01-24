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

type BillingModel struct {
	ChargeType          string `json:"chargeType,omitempty"`
	Period              int32  `json:"period,omitempty"`
	PeriodUnit          string `json:"periodUnit,omitempty"`
	AutoRenew           bool   `json:"autoRenew,omitempty"`
	AutoRenewPeriodUnit string `json:"autoRenewPeriodUnit,omitempty"`
	AutoRenewPeriod     int32  `json:"autoRenewPeriod,omitempty"`
	SpotPriceLimit      string `json:"spotPriceLimit,omitempty"`
	SpotStrategy        string `json:"spotStrategy,omitempty"`
}
