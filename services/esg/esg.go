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

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateEsg - create an esg with the specific parameters
//
// PARAMS:
//   - args: the arguments to create an esg
//
// RETURNS:
//   - *CreateEsgResult: the result of create esg
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEsg(args *CreateEsgArgs) (*CreateEsgResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createEsgArgs cannot be nil.")
	}

	result := &CreateEsgResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEsg()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListEsg - list all esg with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all esg
//
// RETURNS:
//   - *ListEsgResult: the result of list all esg
//   - error: nil if success otherwise the specific error
func (c *Client) ListEsg(args *ListEsgArgs) (*ListEsgResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The ListEsgArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListEsgResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForEsg()).
		WithMethod(http.GET).
		WithQueryParamFilter("instanceId", args.InstanceId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys))

	err := builder.WithResult(result).Do()

	return result, err
}

// DeleteEsg - delete an esg
//
// PARAMS:
//   - DeleteEsgArgs: the arguments to delete an esg
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteEsg(args *DeleteEsgArgs) error {
	if args == nil {
		return fmt.Errorf("The deleteEsgArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEsgId(args.EnterpriseSecurityGroupId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// CreateEsgRules - create esg rules
//
// PARAMS:
//   - CreateEsgRuleArgs: the arguments to create esg rules
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEsgRules(args *CreateEsgRuleArgs) error {
	if args == nil {
		return fmt.Errorf("The createEsgRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEsgId(args.EnterpriseSecurityGroupId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("authorizeRule", "").
		Do()
}

// DeleteEsgRule - delete an esg rule
//
// PARAMS:
//   - DeleteEsgArgs: the arguments to delete an esg rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteEsgRule(args *DeleteEsgRuleArgs) error {
	if args == nil {
		return fmt.Errorf("The deleteEsgRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEsgRuleId(args.EnterpriseSecurityGroupRuleId)).
		WithMethod(http.DELETE).
		Do()
}

// UpdateEsgRule - update esg rule
//
// PARAMS:
//   - CreateEsgRuleArgs: the arguments to update esg rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateEsgRule(args *UpdateEsgRuleArgs) error {
	if args == nil {
		return fmt.Errorf("The updateEsgRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEsgRuleId(args.EnterpriseSecurityGroupRuleId)).
		WithMethod(http.PUT).
		WithBody(args).
		Do()
}
