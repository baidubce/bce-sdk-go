/*
 * Copyright  Baidu, Inc.
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

package havip

type CreateHaVipArgs struct {
	ClientToken      string `json:"-"`
	SubnetId         string `json:"subnetId"`
	Name             string `json:"name"`
	PrivateIpAddress string `json:"privateIpAddress"`
	Description      string `json:"description"`
}

type CreateHavipResult struct {
	HaVipId string `json:"haVipId"`
}

type HaVipBindedInstance struct {
	InstanceId   string `json:"instanceId"`
	InstanceType string `json:"instanceType"`
	Master       bool   `json:"master"`
}

type HaVip struct {
	HaVipId          string                `json:"haVipId"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	VpcId            string                `json:"vpcId"`
	SubnetId         string                `json:"subnetId"`
	Status           string                `json:"status"`
	PrivateIpAddress string                `json:"privateIpAddress"`
	PublicIpAddress  string                `json:"publicIpAddress,omitempty"`
	BindedInstances  []HaVipBindedInstance `json:"bindedInstances,omitempty"`
	CreatedTime      string                `json:"createdTime"`
}

type ListHaVipArgs struct {
	VpcId   string
	Marker  string
	MaxKeys int
}

type ListHaVipResult struct {
	HaVips      []HaVip `json:"haVips"`
	Marker      string  `json:"marker"`
	IsTruncated bool    `json:"isTruncated"`
	NextMarker  string  `json:"nextMarker"`
	MaxKeys     int     `json:"maxKeys"`
}

type UpdateHaVipArgs struct {
	HaVipId     string `json:"-"`
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type DeleteHaVipArgs struct {
	HaVipId     string
	ClientToken string
}

type HaVipInstanceArgs struct {
	HaVipId      string   `json:"-"`
	InstanceIds  []string `json:"instanceIds"`
	InstanceType string   `json:"instanceType"`
	ClientToken  string   `json:"-"`
}

type HaVipBindPublicIpArgs struct {
	HaVipId         string `json:"-"`
	ClientToken     string `json:"-"`
	PublicIpAddress string `json:"publicIpAddress"`
}

type HaVipUnbindPublicIpArgs struct {
	HaVipId     string
	ClientToken string
}
