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

// userservice.go - the User Service APIs definition supported by the User Service service

package userservice

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateUserService - create a User Service
//
// PARAMS:
//   - args: parameters to create User Service
//
// RETURNS:
//   - *CreateUserServiceResult: the result of create User Service, contains new Service Domain
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateUserService(args *CreateUserServiceArgs) (*CreateUserServiceResult, error) {
	if args == nil || len(args.Name) == 0 {
		return nil, fmt.Errorf("unset name")
	}

	if len(args.ServiceName) == 0 {
		return nil, fmt.Errorf("unset service name")
	}

	if len(args.InstanceId) == 0 {
		return nil, fmt.Errorf("unset instance id")
	}

	if len(args.AuthList) > 0 {
		for i := range args.AuthList {
			authModel := args.AuthList[i]
			if authModel.Uid == "" {
				return nil, fmt.Errorf("unset uid")
			}
			if authModel.Auth != ServiceAuthAllow && authModel.Auth != ServiceAuthDeny {
				return nil, fmt.Errorf("invalid auth")
			}
		}
	} else {
		args.AuthList = []UserServiceAuthModel{
			{
				Uid:  "*",
				Auth: ServiceAuthDeny,
			}}
	}

	result := &CreateUserServiceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUserServiceUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateUserService - update a User Service
//
// PARAMS:
//   - service: Service Domain
//   - args: parameters to update User Service
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateUserService(service string, args *UpdateServiceArgs) error {
	if args == nil {
		args = &UpdateServiceArgs{}
	}

	if len(service) == 0 {
		return fmt.Errorf("unset service")
	}

	return c.userServiceRequest("modifyAttribute", http.PUT, getUserServiceIdUri(service), args.ClientToken, args)
}

// UserServiceBindInstance - User Service bind BLB instance
//
// PARAMS:
//   - service: Service Domain
//   - args: parameters to bind blb instance id
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UserServiceBindInstance(service string, args *UserServiceBindArgs) error {
	if args == nil {
		args = &UserServiceBindArgs{}
	}

	if len(service) == 0 {
		return fmt.Errorf("unset service")
	}

	if len(args.InstanceId) == 0 {
		return fmt.Errorf("unset instance id")
	}

	return c.userServiceRequest("bind", http.PUT, getUserServiceIdUri(service), args.ClientToken, args)
}

// UserServiceUnBindInstance - User Service unbind BLB instance
//
// PARAMS:
//   - service: Service Domain
//   - args: parameters to unbind blb instance id
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UserServiceUnBindInstance(service string, args *UserServiceUnBindArgs) error {
	if args == nil {
		args = &UserServiceUnBindArgs{}
	}

	if len(service) == 0 {
		return fmt.Errorf("unset service")
	}

	return c.userServiceRequest("unbind", http.PUT, getUserServiceIdUri(service), args.ClientToken, args)
}

// UserServiceAddAuth - add auth to User Service
//
// PARAMS:
//   - service: Service Domain
//   - args: parameters to add auth
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UserServiceAddAuth(service string, args *UserServiceAuthArgs) error {
	if args == nil {
		args = &UserServiceAuthArgs{}
	}

	if len(args.AuthList) == 0 {
		return fmt.Errorf("unset auth list")
	}

	for i := range args.AuthList {
		authModel := args.AuthList[i]
		if authModel.Uid == "" {
			return fmt.Errorf("unset uid")
		}
		if authModel.Auth != ServiceAuthAllow && authModel.Auth != ServiceAuthDeny {
			return fmt.Errorf("invalid auth")
		}
	}

	return c.userServiceRequest("addAuth", http.PUT, getUserServiceIdUri(service), args.ClientToken, args)
}

// UserServiceEditAuth - edit auth to User Service
//
// PARAMS:
//   - service: Service Domain
//   - args: parameters to edit auth
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UserServiceEditAuth(service string, args *UserServiceAuthArgs) error {
	if args == nil {
		args = &UserServiceAuthArgs{}
	}

	if len(service) == 0 {
		return fmt.Errorf("unset service")
	}

	if len(args.AuthList) == 0 {
		return fmt.Errorf("unset auth list")
	}

	for i := range args.AuthList {
		authModel := args.AuthList[i]
		if authModel.Uid == "" {
			return fmt.Errorf("unset uid")
		}
		if authModel.Auth != ServiceAuthAllow && authModel.Auth != ServiceAuthDeny {
			return fmt.Errorf("invalid auth")
		}
	}

	return c.userServiceRequest("editAuth", http.PUT, getUserServiceIdUri(service), args.ClientToken, args)

}

// UserServiceRemoveAuth - Remove Auth to User Service
//
// PARAMS:
//   - service: Service Domain
//   - args: parameters to remove auth
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UserServiceRemoveAuth(service string, args *UserServiceRemoveAuthArgs) error {
	if args == nil {
		args = &UserServiceRemoveAuthArgs{}
	}

	if len(service) == 0 {
		return fmt.Errorf("unset service")
	}

	if len(args.UidList) == 0 {
		return fmt.Errorf("unset auth list")
	}

	for i := range args.UidList {
		uid := args.UidList[i]
		if uid == "" {
			return fmt.Errorf("unset uid")
		}
	}

	return c.userServiceRequest("removeAuth", http.PUT, getUserServiceIdUri(service), args.ClientToken, args)
}

// DescribeUserServices - describe all User Services
//
// PARAMS:
//   - args: parameters to describe all User Services
//
// RETURNS:
//   - *DescribeUserServicesResult: the result all User Services's detail
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeUserServices(args *DescribeUserServicesArgs) (*DescribeUserServicesResult, error) {
	if args == nil {
		args = &DescribeUserServicesArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeUserServicesResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getUserServiceUri()).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	err := request.Do()
	return result, err
}

// DescribeUserServiceDetail - describe a User Service
//
// PARAMS:
//   - service: describe Service Domain
//
// RETURNS:
//   - *DescribeServiceDetailResult: the result Service detail
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeUserServiceDetail(service string) (*DescribeUserServiceDetailResult, error) {
	result := &DescribeUserServiceDetailResult{}

	if len(service) == 0 {
		return nil, fmt.Errorf("unset service")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getUserServiceIdUri(service)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteUserService - delete a User Service
//
// PARAMS:
//   - blbId: parameters to delete Service Domain
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteUserService(service string) error {
	if len(service) == 0 {
		return fmt.Errorf("unset service")
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getUserServiceIdUri(service)).
		Do()
}
