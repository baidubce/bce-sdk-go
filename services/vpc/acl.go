/*
 * Copyright 2017 Baidu, Inc.
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

// acl.go - the acl APIs definition supported by the VPC service

package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListAclEntrys - list all acl entrys of the given VPC
//
// PARAMS:
//   - vpcId: the id of the specific VPC
//
// RETURNS:
//   - *ListAclEntrysResult: the result of all acl entrys
//   - error: nil if success otherwise the specific error
func (c *Client) ListAclEntrys(vpcId string) (*ListAclEntrysResult, error) {
	result := &ListAclEntrysResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForAclEntry()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", vpcId).
		WithResult(result).
		Do()

	return result, err
}

// CreateAclRule - create a new acl rule with the specific parameters
//
// PARAMS:
//   - args: the arguments to create acl rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAclRule(args *CreateAclRuleArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForAclRule()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// ListAclRules - list all acl rules with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all acl rules
//
// RETURNS:
//   - *ListAclRulesResult: the result of all acl rules
//   - error: nil if success otherwise the specific error
func (c *Client) ListAclRules(args *ListAclRulesArgs) (*ListAclRulesResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The listAclRulesArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListAclRulesResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForAclRule()).
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParam("subnetId", args.SubnetId).
		WithResult(result).
		Do()

	return result, err
}

// UpdateAclRule - udpate acl rule with the specific parameters
//
// PARAMS:
//   - aclRuleId: the id of the specific acl rule
//   - args: the arguments to update acl rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAclRule(aclRuleId string, args *UpdateAclRuleArgs) error {
	if args == nil {
		args = &UpdateAclRuleArgs{}
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForAclRuleId(aclRuleId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// DeleteAclRule - delete the specific acl rule
//
// PARAMS:
//   - aclRuleId: the id of the specific acl rule
//   - clientToken: the idempotent token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAclRule(aclRuleId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForAclRuleId(aclRuleId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}
