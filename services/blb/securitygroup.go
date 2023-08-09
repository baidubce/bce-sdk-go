/*
 * Copyright 2022 Baidu, Inc.
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

// securitygroup.go - the securitygroup APIs definition supported by the BLB service

package blb

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// BindSecurityGroups - bind the blb security groups (normal/application/ipv6 LoadBalancer)
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: the parameter to update security groups
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) BindSecurityGroups(blbId string, args *UpdateSecurityGroupsArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.SecurityGroupIds) == 0 {
		return fmt.Errorf("unset security group ids")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSecurityGroupUri(blbId)).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UnbindSecurityGroups - unbind the blb security groups (normal/application/ipv6 LoadBalancer)
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: the parameter to update security groups
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UnbindSecurityGroups(blbId string, args *UpdateSecurityGroupsArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.SecurityGroupIds) == 0 {
		return fmt.Errorf("unset security group ids")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSecurityGroupUri(blbId)).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeSecurityGroups - describe all security groups of the specified LoadBalancer (normal/application/ipv6 LoadBalancer)
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//
// RETURNS:
//   - *DescribeSecurityGroupsResult: the result of describe all security groups
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeSecurityGroups(blbId string) (*DescribeSecurityGroupsResult, error) {

	result := &DescribeSecurityGroupsResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSecurityGroupUri(blbId)).
		WithResult(result)

	err := request.Do()
	return result, err
}
