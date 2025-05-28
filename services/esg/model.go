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

package esg

import "github.com/baidubce/bce-sdk-go/model"

type CreateEsgArgs struct {
	ClientToken string                        `json:"-"`
	Name        string                        `json:"name"`
	Desc        string                        `json:"desc"`
	Rules       []EnterpriseSecurityGroupRule `json:"rules"`
	Tags        []model.TagModel              `json:"tags,omitempty"`
}

type CreateEsgResult struct {
	EnterpriseSecurityGroupId string `json:"enterpriseSecurityGroupId"`
}

type EnterpriseSecurityGroup struct {
	Id          string                        `json:"id"`
	Name        string                        `json:"name"`
	Desc        string                        `json:"desc"`
	Rules       []EnterpriseSecurityGroupRule `json:"rules"`
	Tags        []model.TagModel              `json:"tags"`
	CreatedTime string                        `json:"createdTime"`
}

type EnterpriseSecurityGroupRule struct {
	EnterpriseSecurityGroupRuleId string `json:"enterpriseSecurityGroupRuleId"`
	Remark                        string `json:"remark"`
	Direction                     string `json:"direction"`
	Ethertype                     string `json:"ethertype"`
	PortRange                     string `json:"portRange"`
	SourcePortRange               string `json:"sourcePortRange"`
	Protocol                      string `json:"protocol"`
	SourceIp                      string `json:"sourceIp"`
	DestIp                        string `json:"destIp"`
	LocalIp                       string `json:"localIp"`
	RemoteIpSet                   string `json:"remoteIpSet"`
	RemoteIpGroup                 string `json:"remoteIpGroup"`
	Action                        string `json:"action"`
	Priority                      int    `json:"priority"`
	CreatedTime                   string `json:"createdTime"`
	UpdatedTime                   string `json:"updatedTime"`
}

type ListEsgArgs struct {
	InstanceId string
	Marker     string
	MaxKeys    int
}

type ListEsgResult struct {
	EnterpriseSecurityGroups []EnterpriseSecurityGroup `json:"enterpriseSecurityGroups"`
	Marker                   string                    `json:"marker"`
	IsTruncated              bool                      `json:"isTruncated"`
	NextMarker               string                    `json:"nextMarker"`
	MaxKeys                  int                       `json:"maxKeys"`
}

type DeleteEsgArgs struct {
	EnterpriseSecurityGroupId string
	ClientToken               string
}

type CreateEsgRuleArgs struct {
	ClientToken               string                        `json:"-"`
	EnterpriseSecurityGroupId string                        `json:"enterpriseSecurityGroupId"`
	Rules                     []EnterpriseSecurityGroupRule `json:"rules"`
}

type DeleteEsgRuleArgs struct {
	EnterpriseSecurityGroupRuleId string
}

type UpdateEsgRuleArgs struct {
	ClientToken                   string
	EnterpriseSecurityGroupRuleId string `json:"-"`
	Remark                        string `json:"remark"`
	PortRange                     string `json:"portRange"`
	Protocol                      string `json:"protocol"`
	SourceIp                      string `json:"sourceIp"`
	DestIp                        string `json:"destIp"`
	Action                        string `json:"action"`
	Priority                      int    `json:"priority"`
}
