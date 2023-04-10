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

package eni

type CreateEniArgs struct {
	ClientToken                string      `json:"-"`
	Name                       string      `json:"name"`
	SubnetId                   string      `json:"subnetId"`
	InstanceId                 string      `json:"instanceId,omitempty"`
	SecurityGroupIds           []string    `json:"securityGroupIds"`
	EnterpriseSecurityGroupIds []string    `json:"enterpriseSecurityGroupIds"`
	PrivateIpSet               []PrivateIp `json:"privateIpSet"`
	Ipv6PrivateIpSet           []PrivateIp `json:"ipv6PrivateIpSet,omitempty"`
	Description                string      `json:"description,omitempty"`
}

type CreateEniResult struct {
	EniId string `json:"eniId"`
}

type UpdateEniArgs struct {
	EniId       string `json:"-"`
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteEniArgs struct {
	EniId       string
	ClientToken string
}

type ListEniArgs struct {
	VpcId            string
	InstanceId       string
	Name             string
	Marker           string
	MaxKeys          int
	PrivateIpAddress []string `json:"privateIpAddress,omitempty"`
}

type ListEniResult struct {
	Eni         []Eni  `json:"enis"`
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

type Eni struct {
	EniId                      string      `json:"eniId"`
	Name                       string      `json:"name"`
	ZoneName                   string      `json:"zoneName"`
	Description                string      `json:"description"`
	InstanceId                 string      `json:"instanceId"`
	MacAddress                 string      `json:"macAddress"`
	VpcId                      string      `json:"vpcId"`
	SubnetId                   string      `json:"subnetId"`
	Status                     string      `json:"status"`
	PrivateIpSet               []PrivateIp `json:"privateIpSet"`
	Ipv6PrivateIpSet           []PrivateIp `json:"ipv6PrivateIpSet"`
	SecurityGroupIds           []string    `json:"securityGroupIds"`
	EnterpriseSecurityGroupIds []string    `json:"enterpriseSecurityGroupIds"`
	CreatedTime                string      `json:"createdTime"`
}

type PrivateIp struct {
	PublicIpAddress  string `json:"publicIpAddress"`
	Primary          bool   `json:"primary"`
	PrivateIpAddress string `json:"privateIpAddress"`
}

type EniPrivateIpArgs struct {
	EniId            string `json:"-"`
	ClientToken      string `json:"-"`
	IsIpv6           bool   `json:"isIpv6,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress"`
}

type EniBatchPrivateIpArgs struct {
	EniId                 string   `json:"-"`
	ClientToken           string   `json:"-"`
	IsIpv6                bool     `json:"isIpv6,omitempty"`
	PrivateIpAddresses    []string `json:"privateIpAddresses"`
	PrivateIpAddressCount int      `json:"privateIpAddressCount,omitempty"`
}

type EniBatchAddPrivateIpCrossSubnetArgs struct {
	EniId                 string          `json:"-"`
	ClientToken           string          `json:"-"`
	SubnetId              string          `json:"subnetId"`
	IsIpv6                bool            `json:"isIpv6,omitempty"`
	PrivateIps            []PrivateIpArgs `json:"privateIps"`
	PrivateIpAddressCount int             `json:"privateIpAddressCount,omitempty"`
}

type PrivateIpArgs struct {
	PrivateIpAddress string `json:"privateIpAddress"`
	SubnetId         string `json:"subnetId"`
}

type AddPrivateIpResult struct {
	PrivateIpAddress string `json:"privateIpAddress"`
}

type BatchAddPrivateIpResult struct {
	PrivateIpAddresses []string `json:"privateIpAddresses"`
}

type EniInstance struct {
	EniId       string `json:"-"`
	InstanceId  string `json:"instanceId"`
	ClientToken string `json:"-"`
}

type BindEniPublicIpArgs struct {
	EniId            string `json:"-"`
	ClientToken      string `json:"-"`
	PrivateIpAddress string `json:"privateIpAddress"`
	PublicIpAddress  string `json:"publicIpAddress"`
}

type UnBindEniPublicIpArgs struct {
	EniId           string `json:"-"`
	ClientToken     string `json:"-"`
	PublicIpAddress string `json:"publicIpAddress"`
}

type UpdateEniSecurityGroupArgs struct {
	EniId            string   `json:"-"`
	ClientToken      string   `json:"-"`
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type UpdateEniEnterpriseSecurityGroupArgs struct {
	EniId                      string   `json:"-"`
	ClientToken                string   `json:"-"`
	EnterpriseSecurityGroupIds []string `json:"enterpriseSecurityGroupIds"`
}

type EniQuoteArgs struct {
	EniId      string `json:"-"`
	InstanceId string `json:"-"`
}

type EniQuoteInfo struct {
	TotalQuantity     int `json:"totalQuantity"`
	AvailableQuantity int `json:"availableQuantity"`
}
