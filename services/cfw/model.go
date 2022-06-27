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

package cfw

type BindCfwRequest struct {
	InstanceType string    `json:"instanceType"`
	Instances    []CfwBind `json:"instances"`
}

type BindCfwRequestInstances struct {
}

type Cfw struct {
	CfwId           string    `json:"cfwId"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	CreatedTime     string    `json:"createdTime"`
	BindInstanceNum int32     `json:"bindInstanceNum"`
	CfwRules        []CfwRule `json:"cfwRules"`
}

type CfwBind struct {
	Region     string `json:"region"`
	InstanceId string `json:"instanceId"`
	Role       string `json:"role"`
	MemberId   string `json:"memberId"`
}

type CfwRule struct {
	IpVersion     int32  `json:"ipVersion"`
	Priority      int32  `json:"priority"`
	Protocol      string `json:"protocol"`
	Direction     string `json:"direction"`
	SourceAddress string `json:"sourceAddress"`
	DestAddress   string `json:"destAddress"`
	SourcePort    string `json:"sourcePort"`
	DestPort      string `json:"destPort"`
	Action        string `json:"action"`
	Description   string `json:"description"`
	CfwId         string `json:"cfwId"`
	CfwRuleId     string `json:"cfwRuleId"`
}

type CreateCfwRequest struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CfwRules    []CreateRule `json:"cfwRules"`
}

type CreateCfwRequestCfwRules struct {
}

type CreateCfwResponse struct {
	CfwId string `json:"cfwId"`
}

type CreateCfwRuleRequest struct {
	CfwRules []CreateRule `json:"cfwRules"`
}

type CreateCfwRuleRequestCfwRules struct {
}

type CreateRule struct {
	IpVersion     int32  `json:"ipVersion"`
	Priority      int32  `json:"priority"`
	Protocol      string `json:"protocol"`
	Direction     string `json:"direction"`
	SourceAddress string `json:"sourceAddress"`
	DestAddress   string `json:"destAddress"`
	SourcePort    string `json:"sourcePort"`
	DestPort      string `json:"destPort"`
	Action        string `json:"action"`
	Description   string `json:"description"`
}

type DeleteCfwRuleRequest struct {
	CfwRuleIds []string `json:"cfwRuleIds"`
}

type DeleteCfwRuleRequestCfwRuleIds struct {
}

type DisableCfwRequest struct {
	InstanceId string `json:"instanceId"`
	Role       string `json:"role"`
	MemberId   string `json:"memberId"`
}

type EnableCfwRequest struct {
	InstanceId string `json:"instanceId"`
	Role       string `json:"role"`
	MemberId   string `json:"memberId"`
}

type GetCfwResponse struct {
	CfwId           string    `json:"cfwId"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	CreatedTime     string    `json:"createdTime"`
	BindInstanceNum int32     `json:"bindInstanceNum"`
	CfwRules        []CfwRule `json:"cfwRules"`
}

type GetCfwResponseCfwRules struct {
}

type Instance struct {
	InstanceId      string `json:"instanceId"`
	InstanceName    string `json:"instanceName"`
	Status          string `json:"status"`
	Region          string `json:"region"`
	CfwId           string `json:"cfwId"`
	CfwName         string `json:"cfwName"`
	VpcId           string `json:"vpcId"`
	VpcName         string `json:"vpcName"`
	PublicIp        string `json:"publicIp"`
	Role            string `json:"role"`
	LocalIfId       string `json:"localIfId"`
	LocalIfName     string `json:"localIfName"`
	PeerRegion      string `json:"peerRegion"`
	PeerVpcId       string `json:"peerVpcId"`
	PeerVpcName     string `json:"peerVpcName"`
	MemberId        string `json:"memberId"`
	MemberName      string `json:"memberName"`
	MemberAccountId string `json:"memberAccountId"`
}

type ListCfwResponse struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int32  `json:"maxKeys"`
	Cfws        []Cfw  `json:"cfws"`
}

type ListCfwResponseCfws struct {
}

type ListInstanceRequest struct {
	InstanceType string `json:"instanceType"`
	Marker       string `json:"marker"`
	MaxKeys      int    `json:"maxKeys"`
	Status       string `json:"status"`
	Region       string `json:"region"`
}

type ListInstanceRequestCfwRuleIds struct {
}

type ListInstanceResponse struct {
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
	Instances   []Instance `json:"instances"`
}

type ListInstanceResponseInstances struct {
}

type UnbindCfwRequest struct {
	InstanceType string    `json:"instanceType"`
	Instances    []CfwBind `json:"instances"`
}

type UnbindCfwRequestInstances struct {
}

type UpdateCfwRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCfwRuleRequest struct {
	IpVersion     int32  `json:"ipVersion"`
	Priority      int32  `json:"priority"`
	Protocol      string `json:"protocol"`
	Direction     string `json:"direction"`
	SourceAddress string `json:"sourceAddress"`
	DestAddress   string `json:"destAddress"`
	SourcePort    string `json:"sourcePort"`
	DestPort      string `json:"destPort"`
	Action        string `json:"action"`
	Description   string `json:"description"`
}

type ListCfwArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}
