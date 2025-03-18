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

type HpasResponse struct {
	HpasId              string        `json:"hpasId"`
	Name                string        `json:"name"`
	ZoneName            string        `json:"zoneName"`
	Status              string        `json:"status"`
	AppType             string        `json:"appType"`
	AppPerformanceLevel string        `json:"appPerformanceLevel"`
	ChargeType          string        `json:"chargeType"`
	VpcId               string        `json:"vpcId"`
	VpcName             string        `json:"vpcName"`
	VpcCidr             string        `json:"vpcCidr"`
	InternalIp          string        `json:"internalIp"`
	SubnetId            string        `json:"subnetId"`
	SubnetName          string        `json:"subnetName"`
	EhcClusterId        string        `json:"ehcClusterId"`
	EhcClusterName      string        `json:"ehcClusterName"`
	ImageId             string        `json:"imageId"`
	ImageName           string        `json:"imageName"`
	CreateTime          string        `json:"createTime"`
	Tags                []TagResponse `json:"tags"`
	RdmaNicTopo         []RdmaNicTopo `json:"rdmaNicTopo"`
	NicInfo             []NicInfo     `json:"nicInfo"`
	AppImageFamily      string        `json:"appImageFamily"`
	RdmaUnitID          string        `json:"rdmaUnitID"`
	RdmaPodName         string        `json:"rdmaPodName"`
}

type RdmaNicTopo struct {
	RdmaIp     string `json:"rdmaIp"`
	SwitchName string `json:"switchName"`
	SwitchPort string `json:"switchPort"`
}

type NicInfo struct {
	NicId             string     `json:"nicId"`
	NicName           string     `json:"nicName"`
	NicType           string     `json:"nicType"`
	SubnetId          string     `json:"subnetId"`
	Description       string     `json:"description"`
	HpasId            string     `json:"hpasId"`
	NicStatus         string     `json:"nicStatus"`
	MacAddress        string     `json:"macAddress"`
	VpcId             string     `json:"vpcId"`
	CreatedTime       string     `json:"createdTime"`
	Ipv4s             []Ipv4Info `json:"ipv4s"`
	Ipv6s             []Ipv6Info `json:"ipv6s"`
	SecurityGroupType string     `json:"securityGroupType"`
	SecurityGroupIds  []string   `json:"securityGroupIds"`
}

type Ipv4Info struct {
	PrivateIP string `json:"privateIp"`
	Primary   bool   `json:"primary"`
}

type Ipv6Info struct {
	Ipv6Address string `json:"ipv6Address"`
}