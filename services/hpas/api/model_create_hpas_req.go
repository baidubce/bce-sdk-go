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

type CreateHpasReq struct {
	AppType             string                   `json:"appType,omitempty"`
	AppPerformanceLevel string                   `json:"appPerformanceLevel,omitempty"`
	Name                string                   `json:"name,omitempty"`
	ApplicationName     string                   `json:"applicationName,omitempty"`
	AutoSeqSuffix       bool                     `json:"autoSeqSuffix,omitempty"`
	Description         string                   `json:"description,omitempty"`
	PurchaseNum         int                      `json:"purchaseNum,omitempty"`
	ZoneName            string                   `json:"zoneName,omitempty"`
	ImageId             string                   `json:"imageId,omitempty"`
	SubnetId            string                   `json:"subnetId,omitempty"`
	Password            string                   `json:"password,omitempty"`
	KeypairId           string                   `json:"keypairId,omitempty"`
	EhcClusterId        string                   `json:"ehcClusterId,omitempty"`
	SecurityGroupType   string                   `json:"securityGroupType,omitempty"`
	SecurityGroupIds    []string                 `json:"securityGroupIds,omitempty"`
	InternalIps         []string                 `json:"internalIps,omitempty"`
	BillingModel        BillingModel             `json:"billingModel,omitempty"`
	Tags                []TagModel               `json:"tags,omitempty"`
	UserData            string                   `json:"userData,omitempty"`
	ReservedInstance    *CreateCombinedCouponReq `json:"reservedInstance,omitempty"`
}
