/*
 * Copyright 2020 Baidu, Inc.
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

// rds.go - the rds APIs definition supported by the RDS service
package rds

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateRds - create a RDS with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a rds
// RETURNS:
//     - *InstanceIds: the result of create RDS, contains new RDS's instanceIds
//     - error: nil if success otherwise the specific error
func (c *Client) CreateRds(args *CreateRdsArgs) (*CreateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.Engine == "" {
		return nil, fmt.Errorf("unset Engine")
	}

	if args.EngineVersion == "" {
		return nil, fmt.Errorf("unset EngineVersion")
	}

	if args.Billing.PaymentTiming == "" {
		return nil, fmt.Errorf("unset PaymentTiming")
	}

	result := &CreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateReadReplica - create a readReplica RDS with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a readReplica rds
// RETURNS:
//     - *InstanceIds: the result of create a readReplica RDS, contains the readReplica RDS's instanceIds
//     - error: nil if success otherwise the specific error
func (c *Client) CreateReadReplica(args *CreateReadReplicaArgs) (*CreateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.SourceInstanceId == "" {
		return nil, fmt.Errorf("unset SourceInstanceId")
	}

	if args.Billing.PaymentTiming == "" {
		return nil, fmt.Errorf("unset PaymentTiming")
	}

	result := &CreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("readReplica","").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateRdsProxy - create a proxy RDS with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a readReplica rds
// RETURNS:
//     - *InstanceIds: the result of create a readReplica RDS, contains the readReplica RDS's instanceIds
//     - error: nil if success otherwise the specific error
func (c *Client) CreateRdsProxy(args *CreateRdsProxyArgs) (*CreateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.SourceInstanceId == "" {
		return nil, fmt.Errorf("unset SourceInstanceId")
	}

	if args.Billing.PaymentTiming == "" {
		return nil, fmt.Errorf("unset PaymentTiming")
	}

	result := &CreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("rdsproxy","").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListRds - list all RDS with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all RDS
// RETURNS:
//     - *ListRdsResult: the result of list all RDS, contains all rds' meta
//     - error: nil if success otherwise the specific error
func (c *Client) ListRds(args *ListRdsArgs) (*ListRdsResult, error) {
	if args == nil {
		args = &ListRdsArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListRdsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUri()).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetDetail - get a specific rds Instance's detail
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *Instance: the specific rdsInstance's detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetDetail(instanceId string) (*Instance, error) {
	result := &Instance{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteRds - delete a rds
//
// PARAMS:
//     - instanceIds: the specific instanceIds
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRds(instanceIds string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUri()).
		WithQueryParamFilter("instanceIds", instanceIds).
		Do()
}

// ResizeRds - resize an RDS with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//     - args: the arguments to resize an RDS
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeRds(instanceId string, args *ResizeRdsArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("resize", "").
		WithBody(args).
		Do()
}

// CreateAccount - create a account with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//     - args: the arguments to create a account
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateAccount(instanceId string, args *CreateAccountArgs)  error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.AccountName == "" {
		return fmt.Errorf("unset AccountName")
	}

	if args.Password == "" {
		return fmt.Errorf("unset Password")
	}

	cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.Password)
	if err != nil {
		return err
	}
	args.Password = cryptedPass

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/account").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// ListAccount - list all account of a RDS instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *ListAccountResult: the result of list all account, contains all accounts' meta
//     - error: nil if success otherwise the specific error
func (c *Client) ListAccount(instanceId string) (*ListAccountResult, error) {
	result := &ListAccountResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/account").
		WithResult(result).
		Do()

	return result, err
}

// GetAccount - get an account of a RDS instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - accountName: the specific account's name
// RETURNS:
//     - *Account: the account's meta
//     - error: nil if success otherwise the specific error
func (c *Client) GetAccount(instanceId,accountName string) (*Account, error) {
	result := &Account{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/account/"+accountName).
		WithResult(result).
		Do()

	return result, err
}

// DeleteAccount - delete an account of a RDS instance
//
// PARAMS:
//     - instanceIds: the specific instanceIds
//     - accountName: the specific account's name
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteAccount(instanceId, accountName string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/account/" + accountName).
		Do()
}