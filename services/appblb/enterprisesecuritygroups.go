/*
 * Copyright 2024 Baidu, Inc.
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

// enterprisesecuritygroups.go - the enterprisesecuritygroup APIs definition supported by the APPBLB service

package appblb

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// BindEnterpriseSecurityGroups - bind the blb enterprise security groups (normal/application/ipv6 LoadBalancer)
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: the parameter to update enterprise security groups
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) BindEnterpriseSecurityGroups(blbId string, args *UpdateEnterpriseSecurityGroupsArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.EnterpriseSecurityGroupIds) == 0 {
		return fmt.Errorf("unset enterprise security group ids")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnterpriseSecurityGroupUri(blbId)).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UnbindEnterpriseSecurityGroups - unbind the blb enterprise security groups (normal/application/ipv6 LoadBalancer)
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: the parameter to update enterprise security groups
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UnbindEnterpriseSecurityGroups(blbId string, args *UpdateEnterpriseSecurityGroupsArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.EnterpriseSecurityGroupIds) == 0 {
		return fmt.Errorf("unset enterprise security group ids")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnterpriseSecurityGroupUri(blbId)).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeEnterpriseSecurityGroups - describe all enterprise security groups of the specified LoadBalancer (normal/application/ipv6 LoadBalancer)
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//
// RETURNS:
//   - *DescribeEnterpriseSecurityGroupsResult: the result of describe all enterprise security groups
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeEnterpriseSecurityGroups(blbId string) (*DescribeEnterpriseSecurityGroupsResult, error) {

	result := &DescribeEnterpriseSecurityGroupsResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEnterpriseSecurityGroupUri(blbId)).
		WithResult(result)

	err := request.Do()
	return result, err
}
