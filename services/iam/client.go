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

// client.go - define the client for IAM service which is derived from BceClient

// Package iam defines the IAM service of BCE.
// It contains the model sub package to implement the concrete request and response of the
// IAM user/accessKey/policy API
package iam

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/iam/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "iam." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client of IAM service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk string) (*Client, error) {
	return NewClientWithEndpoint(ak, sk, DEFAULT_SERVICE_DOMAIN)
}

func NewClientWithEndpoint(ak, sk, endpoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

func (c *Client) CreateUser(args *api.CreateUserArgs) (*api.CreateUserResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.CreateUser(c, body)
}

func (c *Client) GetUser(name string) (*api.GetUserResult, error) {
	return api.GetUser(c, name)
}

func (c *Client) UpdateUser(name string, args *api.UpdateUserArgs) (*api.UpdateUserResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.UpdateUser(c, name, body)
}

func (c *Client) DeleteUser(name string) error {
	return api.DeleteUser(c, name)
}

func (c *Client) ListUser() (*api.ListUserResult, error) {
	return api.ListUser(c)
}

func (c *Client) UpdateUserLoginProfile(name string, args *api.UpdateUserLoginProfileArgs) (
	*api.UpdateUserLoginProfileResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.UpdateUserLoginProfile(c, name, body)
}

func (c *Client) GetUserLoginProfile(name string) (*api.GetUserLoginProfileResult, error) {
	return api.GetUserLoginProfile(c, name)
}

func (c *Client) DeleteUserLoginProfile(name string) error {
	return api.DeleteUserLoginProfile(c, name)
}

func (c *Client) CreateGroup(args *api.CreateGroupArgs) (*api.CreateGroupResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.CreateGroup(c, body)
}

func (c *Client) GetGroup(name string) (*api.GetGroupResult, error) {
	return api.GetGroup(c, name)
}

func (c *Client) UpdateGroup(name string, args *api.UpdateGroupArgs) (*api.UpdateGroupResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.UpdateGroup(c, name, body)
}

func (c *Client) DeleteGroup(name string) error {
	return api.DeleteGroup(c, name)
}

func (c *Client) ListGroup() (*api.ListGroupResult, error) {
	return api.ListGroup(c)
}

func (c *Client) AddUserToGroup(userName string, groupName string) error {
	return api.AddUserToGroup(c, userName, groupName)
}

func (c *Client) DeleteUserFromGroup(userName string, groupName string) error {
	return api.DeleteUserFromGroup(c, userName, groupName)
}

func (c *Client) ListUsersInGroup(name string) (*api.ListUsersInGroupResult, error) {
	return api.ListUsersInGroup(c, name)
}

func (c *Client) ListGroupsForUser(name string) (*api.ListGroupsForUserResult, error) {
	return api.ListGroupsForUser(c, name)
}

func (c *Client) CreatePolicy(args *api.CreatePolicyArgs) (*api.CreatePolicyResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.CreatePolicy(c, body)
}

func (c *Client) GetPolicy(name, policyType string) (*api.GetPolicyResult, error) {
	return api.GetPolicy(c, name, policyType)
}

func (c *Client) DeletePolicy(name string) error {
	return api.DeletePolicy(c, name)
}

func (c *Client) ListPolicy(nameFilter, policyType string) (*api.ListPolicyResult, error) {
	return api.ListPolicy(c, nameFilter, policyType)
}

func (c *Client) AttachPolicyToUser(args *api.AttachPolicyToUserArgs) error {
	return api.AttachPolicyToUser(c, args)
}

func (c *Client) DetachPolicyFromUser(args *api.DetachPolicyFromUserArgs) error {
	return api.DetachPolicyFromUser(c, args)
}

func (c *Client) ListUserAttachedPolicies(name string) (*api.ListPolicyResult, error) {
	return api.ListUserAttachedPolicies(c, name)
}

func (c *Client) AttachPolicyToGroup(args *api.AttachPolicyToGroupArgs) error {
	return api.AttachPolicyToGroup(c, args)
}

func (c *Client) DetachPolicyFromGroup(args *api.DetachPolicyFromGroupArgs) error {
	return api.DetachPolicyFromGroup(c, args)
}

func (c *Client) ListGroupAttachedPolicies(name string) (*api.ListPolicyResult, error) {
	return api.ListGroupAttachedPolicies(c, name)
}

func (c *Client) CreateAccessKey(userName string) (*api.CreateAccessKeyResult, error) {
	return api.CreateAccessKey(c, userName)
}

func (c *Client) DisableAccessKey(userName, accessKeyId string) (*api.UpdateAccessKeyResult, error) {
	return api.DisableAccessKey(c, userName, accessKeyId)
}

func (c *Client) EnableAccessKey(userName, accessKeyId string) (*api.UpdateAccessKeyResult, error) {
	return api.EnableAccessKey(c, userName, accessKeyId)
}

func (c *Client) DeleteAccessKey(userName, accessKeyId string) error {
	return api.DeleteAccessKey(c, userName, accessKeyId)
}

func (c *Client) ListAccessKey(userName string) (*api.ListAccessKeyResult, error) {
	return api.ListAccessKey(c, userName)
}

func (c *Client) CreateRole(args *api.CreateRoleArgs) (*api.CreateRoleResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.CreateRole(c, body)
}

func (c *Client) GetRole(roleName string) (*api.GetRoleResult, error) {
	return api.GetRole(c, roleName)
}

func (c *Client) UpdateRole(roleName string, args *api.UpdateRoleArgs) (*api.UpdateRoleResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.UpdateRole(c, roleName, body)
}

func (c *Client) DeleteRole(roleName string) error {
	return api.DeleteRole(c, roleName)
}

func (c *Client) ListRole() (*api.ListRoleResult, error) {
	return api.ListRole(c)
}

func (c *Client) AttachPolicyToRole(args *api.AttachPolicyToRoleArgs) error {
	return api.AttachPolicyToRole(c, args)
}

func (c *Client) DetachPolicyFromRole(args *api.DetachPolicyToRoleArgs) error {
	return api.DetachPolicyFromRole(c, args)
}

func (c *Client) ListRoleAttachedPolicies(name string) (*api.ListPolicyResult, error) {
	return api.ListRoleAttachedPolicies(c, name)
}

func (c *Client) ListPolicyAttachedEntities(policyId string) (*api.ListPolicyAttachedEntityResult, error) {
	return api.ListPolicyAttachedEntities(c, policyId)
}

func (c *Client) UserOperationMfaSwitch(args *api.UserSwitchMfaArgs) error {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return err
	}
	return api.UserOperationMfaSwitch(c, body)
}

func (c *Client) SubUserUpdate(userName string, args *api.UpdateSubUserArgs) (*api.UpdateUserResult, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.SubUserUpdate(c, body, userName)
}

func NewBodyFromStruct(args interface{}) (*bce.Body, error) {
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return body, nil
}
