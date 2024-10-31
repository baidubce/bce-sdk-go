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

package api

import "time"

type UserModel struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreateTime  time.Time `json:"createTime"`
	Description string    `json:"description"`
	Enabled     bool      `json:"enabled"`
}

type CreateUserArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateUserResult UserModel

type GetUserResult UserModel

type UpdateUserArgs struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateUserResult UserModel

type ListUserResult struct {
	Users []UserModel `json:"users"`
}

type AccessKeyModel struct {
	Id           string    `json:"id"`
	Secret       string    `json:"secret"`
	CreateTime   time.Time `json:"createTime"`
	LastUsedTime time.Time `json:"lastUsedTime"`
	Enabled      bool      `json:"enabled"`
	Description  string    `json:"description"`
}

type CreateAccessKeyResult AccessKeyModel

type UpdateAccessKeyResult AccessKeyModel

type ListAccessKeyResult struct {
	AccessKeys []AccessKeyModel `json:"accessKeys"`
}

type LoginProfileModel struct {
	Password          string `json:"password,omitempty"`
	NeedResetPassword bool   `json:"needResetPassword"`
	EnabledLogin      bool   `json:"enabledLogin"`
	EnabledLoginMfa   bool   `json:"enabledLoginMfa"`
	LoginMfaType      string `json:"loginMfaType,omitempty"`
	ThirdPartyType    bool   `json:"thirdPartyType,omitempty"`
	ThirdPartyAccount bool   `json:"thirdPartyAccount,omitempty"`
}

type UpdateUserLoginProfileArgs LoginProfileModel

type UpdateUserLoginProfileResult LoginProfileModel

type GetUserLoginProfileResult LoginProfileModel

type GroupModel struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreateTime  time.Time `json:"createTime"`
	Description string    `json:"description"`
}

type CreateGroupArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateGroupResult GroupModel

type GetGroupResult GroupModel

type UpdateGroupArgs struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateGroupResult GroupModel

type ListGroupResult struct {
	Groups []GroupModel `json:"groups"`
}

type ListUsersInGroupResult ListUserResult

type ListGroupsForUserResult ListGroupResult

type AclEntry struct {
	Eid        string    `json:"eid,omitempty"`
	Service    string    `json:"service"`
	Region     string    `json:"region"`
	Permission []string  `json:"permission"`
	Resource   []string  `json:"resource,omitempty"`
	Grantee    []Grantee `json:"grantee,omitempty"`
	Effect     string    `json:"effect"`
}

type Grantee struct {
	ID string `json:"id"`
}

type Acl struct {
	Id                string     `json:"id,omitempty"`
	AccessControlList []AclEntry `json:"accessControlList"`
}

type PolicyModel struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	CreateTime  time.Time `json:"createTime"`
	Description string    `json:"description"`
	Document    string    `json:"document"`
}

type CreatePolicyArgs struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Document    string `json:"document"`
}

type UpdatePolicyArgs struct {
	PolicyName  string `json:"policyName"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Document    string `json:"document"`
}

type CreatePolicyResult PolicyModel

type UpdatePolicyResult PolicyModel

type GetPolicyResult PolicyModel

type ListPolicyResult struct {
	Policies []PolicyModel `json:"policies"`
}

type AttachPolicyToUserArgs struct {
	UserName   string `json:"userName"`
	PolicyName string `json:"policyName"`
	PolicyType string `json:"policyType,omitempty"`
}

type DetachPolicyFromUserArgs struct {
	UserName   string `json:"userName"`
	PolicyName string `json:"policyName"`
	PolicyType string `json:"policyType,omitempty"`
}

type AttachPolicyToGroupArgs struct {
	GroupName  string `json:"groupName"`
	PolicyName string `json:"policyName"`
	PolicyType string `json:"policyType,omitempty"`
}

type DetachPolicyFromGroupArgs struct {
	GroupName  string `json:"groupName"`
	PolicyName string `json:"policyName"`
	PolicyType string `json:"policyType,omitempty"`
}

type AttachPolicyToRoleArgs struct {
	RoleName   string `json:"roleName"`
	PolicyName string `json:"policyName"`
	PolicyType string `json:"policyType,omitempty"`
}

type DetachPolicyToRoleArgs struct {
	RoleName   string `json:"roleName"`
	PolicyName string `json:"policyName"`
	PolicyType string `json:"policyType,omitempty"`
}

type RoleModel struct {
	Id                       string    `json:"id"`
	Name                     string    `json:"name"`
	CreateTime               time.Time `json:"createTime"`
	Description              string    `json:"description"`
	AssumeRolePolicyDocument string    `json:"assumeRolePolicyDocument"`
}

type CreateRoleArgs struct {
	Name                     string `json:"name"`
	Description              string `json:"description,omitempty"`
	AssumeRolePolicyDocument string `json:"assumeRolePolicyDocument"`
}

type UpdateRoleArgs struct {
	Description              string `json:"description"`
	AssumeRolePolicyDocument string `json:"assumeRolePolicyDocument,omitempty"`
}

type CreateRoleResult RoleModel

type GetRoleResult RoleModel

type UpdateRoleResult RoleModel

type ListRoleResult struct {
	Roles []RoleModel `json:"roles"`
}

type UserSwitchMfaArgs struct {
	UserName   string `json:"userName,omitempty"`
	EnabledMfa bool   `json:"enabledMfa"`
	MfaType    string `json:"mfaType,omitempty"`
}

type UpdateSubUserArgs struct {
	Password string `json:"password,omitempty"`
	Provider string `json:"provider,omitempty"`
	Enable   bool   `json:"enable"`
}

type PolicyAttachedEntity struct {
	Id         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Type       string    `json:"type,omitempty"`
	AttachTime time.Time `json:"attachTime,omitempty"`
}

type ListPolicyAttachedEntityResult struct {
	PolicyAttachedEntities []PolicyAttachedEntity `json:"entities"`
}

type IdpWithStatus struct {
	Status          string    `json:"status,omitempty"`
	AuxiliaryDomain string    `json:"auxiliaryDomain,omitempty"`
	DomainId        string    `json:"domainId,omitempty"`
	EncodeMetadata  string    `json:"encodeMetadata,omitempty"`
	FileName        string    `json:"fileName,omitempty"`
	CreateTime      time.Time `json:"createTime,omitempty"`
	UpdateTime      time.Time `json:"updateTime,omitempty"`
}

type UpdateSubUserIdpRequest struct {
	FileName        string `json:"fileName,omitempty"`
	EncodeMetadata  string `json:"encodeMetadata,omitempty"`
	AuxiliaryDomain string `json:"auxiliaryDomain,omitempty"`
}
