/*
 * Copyright 2021 Baidu, Inc.
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

package endpoint

import "github.com/baidubce/bce-sdk-go/model"

type Endpoint struct {
	EndpointId  string           `json:"endpointId"`
	Name        string           `json:"name"`
	IpAddress   string           `json:"ipAddress"`
	Status      string           `json:"status"`
	Service     string           `json:"service"`
	SubnetId    string           `json:"subnetId"`
	Description string           `json:"description"`
	CreateTime  string           `json:"createTime"`
	VpcId       string           `json:"vpcId"`
	ProductType string           `json:"productType"`
	Tags        []model.TagModel `json:"tags,omitempty"`
}

type ListEndpointArgs struct {
	VpcId     string
	Name      string
	IpAddress string
	Status    string
	SubnetId  string
	Service   string
	Marker    string
	MaxKeys   int
}

type ListEndpointResult struct {
	Endpoints   []Endpoint `json:"endpoints"`
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
}

type ListServiceResult struct {
	Services []string `json:"services"`
}

type UpdateEndpointArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateEndpointNSGArgs struct {
	SecurityGroupIds []string `json:"securityGroupIds"`
	ClientToken      string   `json:"-"`
}
type UpdateEndpointESGArgs struct {
	EnterpriseSecurityGroupIds []string `json:"enterpriseSecurityGroupIds"`
	ClientToken                string   `json:"-"`
}

type CreateEndpointArgs struct {
	ClientToken string           `json:"-"`
	VpcId       string           `json:"vpcId"`
	Name        string           `json:"name"`
	SubnetId    string           `json:"subnetId"`
	Service     string           `json:"service"`
	Description string           `json:"description,omitempty"`
	IpAddress   string           `json:"ipAddress,omitempty"`
	Billing     *Billing         `json:"billing"`
	Tags        []model.TagModel `json:"tags,omitempty"`
}

type CreateEndpointResult struct {
	Id        string `json:"id"`
	IpAddress string `json:"ipAddress"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

type (
	PaymentTimingType string
)

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}
