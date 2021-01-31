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

// ddc.go - the ddc APIs definition supported by the DDC service
package ddc

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

const (
	KEY_CLIENT_TOKEN = "clientToken"
	KEY_MARKER       = "marker"
	KEY_MAXKEYS      = "maxKeys"
)

// CreateInstance - create a Instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a instance
// RETURNS:
//     - *InstanceIds: the result of create RDS, contains new RDS's instanceIds
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *CreateInstanceArgs) (*CreateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	result := &CreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDdcInstanceUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateDeploySet - create a deploy set
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) CreateDeploySet(poolId string, args *CreateDeployRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if !(args.Strategy == "distributed" || args.Strategy == "centralized") {
		return fmt.Errorf("Only support strategy distributed/centralized, current strategy: %v", args.Strategy)
	}

	result := &bce.BceResponse{}
	err := bce.NewRequestBuilder(cli).
		WithMethod(http.POST).
		WithURL(getDeploySetUri(poolId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return err
}

// ListDdcInstance - list all instances
// RETURNS:
//     - *ListDdcResult: the result of list instances with marker
//     - error: nil if success otherwise the specific error
func (c *Client) ListDdcInstance(marker *Marker) (*ListDdcResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDdcInstanceUri() + "/list")
	req.SetMethod(http.GET)
	if marker != nil {
		req.SetParam(KEY_MARKER, marker.Marker)
		req.SetParam(KEY_MAXKEYS, strconv.Itoa(marker.MaxKeys))
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListDdcResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetDetail - get details of the instance
//
// PARAMS:
//     - instanceId: the id of the instance
// RETURNS:
//     - *InstanceModelResult: the detail of the instance
//     - error: nil if success otherwise the specific error
func (c *Client) GetDetail(instanceId string) (*InstanceModelResult, error) {
	result := &InstanceModelResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteDdsInstance - delete instances
//
// PARAMS:
//     - DeleteDdcArgs: the id list of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteDdcInstance(instanceIds string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDdcInstanceUri()+"/delete").
		WithQueryParam("instanceIds", instanceIds).
		Do()
}

// UpdateInstanceNameArgs - update name of a specified instance
//
// PARAMS:
//     - instanceId: id of the instance
//     - args: the arguments to update instanceName
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(instanceId string, args *UpdateInstanceNameArgs) error {

	result := &bce.BceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instanceId)+"/updateName").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return err
}

// GetBackupList - get backup list of the instance
//
// PARAMS:
//     - instanceId: id of the instance
// RETURNS:
//     - *GetBackupListResult: result of the backup list
//     - error: nil if success otherwise the specific error
func (c *Client) GetBackupList(instanceId string) (*GetBackupListResult, error) {

	result := &GetBackupListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/snapshot").
		WithResult(result).
		Do()

	return result, err
}

// GetZoneList - list all zone
//
// PARAMS:
//     - c: the client agent which can perform sending request
// RETURNS:
//     - *GetZoneListResult: result of the zone list
//     - error: nil if success otherwise the specific error
func (c *Client) GetZoneList() (*GetZoneListResult, error) {
	result := &GetZoneListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(URI_PREFIX + "/zone").
		WithResult(result).
		Do()

	return result, err
}

// ListsSubnet - list all Subnets
//
// PARAMS:
//     - c: the client agent which can perform sending request
//     - args: the arguments to list all subnets, not necessary
// RETURNS:
//     - *ListSubnetsResult: result of the subnet list
//     - error: nil if success otherwise the specific error
func (c *Client) ListSubnets() (*ListSubnetsResult, error) {
	result := &ListSubnetsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(URI_PREFIX + "/subnet").
		WithResult(result).
		Do()

	return result, err
}

// ListDeploySets - list all deploy sets
// RETURNS:
//     - *ListResultWithMarker: the result of list deploy sets with marker
//     - error: nil if success otherwise the specific error
func (cli *Client) ListDeploySets(poolId string, marker *Marker) (*ListDeploySetResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUri(poolId))
	req.SetMethod(http.GET)
	if marker != nil {
		req.SetParam(KEY_MARKER, marker.Marker)
		req.SetParam(KEY_MAXKEYS, strconv.Itoa(marker.MaxKeys))
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListDeploySetResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// DeleteDeploySet - delete a deploy set
//
// PARAMS:
//     - poolId: the id of the pool
//     - deploySetId: the id of the deploy set
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) DeleteDeploySet(poolId string, deploySetId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUriWithId(poolId, deploySetId))
	req.SetMethod(http.DELETE)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()

	return nil
}

// GetDeploySet - get details of the deploy set
//
// PARAMS:
//     - poolId: the id of the pool
//     - cli: the client agent which can perform sending request
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - *DeploySet: the detail of the deploy set
//     - error: nil if success otherwise the specific error
func (cli *Client) GetDeploySet(poolId string, deploySetId string) (*DeploySet, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUriWithId(poolId, deploySetId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeploySet{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetSecurityIps - get all SecurityIps
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *GetSecurityIpsResult: all security IP
//     - error: nil if success otherwise the specific error
func (c *Client) GetSecurityIps(instanceId string) (*GetSecurityIpsResult, error) {
	result := &GetSecurityIpsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/authIp").
		WithResult(result).
		Do()

	return result, err
}

// UpdateSecurityIps - update SecurityIps
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - Args: all SecurityIps
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateSecurityIps(instacneId string, args *UpdateSecurityIpsArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instacneId)+"/updateAuthIp").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ListParameters - list all parameters of a RDS instance
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *ListParametersResult: the result of list all parameters
//     - error: nil if success otherwise the specific error
func (c *Client) ListParameters(instanceId string) (*ListParametersResult, error) {
	result := &ListParametersResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/parameter" + "/list").
		WithResult(result).
		Do()

	return result, err
}

// UpdateParameter - update Parameter
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - Args: *UpdateParameterArgs
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateParameter(instanceId string, args *UpdateParameterArgs) error {

	result := &bce.BceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instanceId)+"/parameter"+"/modify").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return err
}

// CreateBackup - create backup of the instance
//
// PARAMS:
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateBackup(instanceId string) error {

	result := &bce.BceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDdcUriWithInstanceId(instanceId)+"/snapshot").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return err
}

// GetBackupDetail - get details of the instance'Backup
//
// PARAMS:
//     - instanceId: the id of the instance
//     - snapshotId: the id of the backup
// RETURNS:
//     - *BackupDetailResult: the detail of the backup
//     - error: nil if success otherwise the specific error
func (c *Client) GetBackupDetail(instanceId string, snapshotId string) (*BackupDetailResult, error) {
	result := &BackupDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/snapshot" + "/" + snapshotId).
		WithResult(result).
		Do()

	return result, err
}

// ModifyBackupPolicy - update backupPolicy
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - Args: the specific rds Instance's BackupPolicy
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupPolicy(instanceId string, args *BackupPolicy) error {

	result := &bce.BceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instanceId)+"/snapshot"+"/update").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return err
}

// GetBackupList - get backup list of the instance
//
// PARAMS:
//     - instanceId: id of the instance
// RETURNS:
//     - *GetBackupListResult: result of the backup list
//     - error: nil if success otherwise the specific error
func (c *Client) GetBinlogList(instanceId string, datetime string) (*BinlogListResult, error) {

	result := &BinlogListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId)+"/binlog").
		WithQueryParam("datetime", datetime).
		WithResult(result).
		Do()

	return result, err
}

// GetBackupDetail - get details of the instance'Binlog
//
// PARAMS:
//     - instanceId: the id of the instance
//     - binlog: the id of the binlog
// RETURNS:
//     - *BackupDetailResult: the detail of the binlog
//     - error: nil if success otherwise the specific error
func (c *Client) GetBinlogDetail(instanceId string, binlog string) (*BinlogDetailResult, error) {
	result := &BinlogDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/binlog" + "/" + binlog).
		WithResult(result).
		Do()

	return result, err
}

// SwitchInstance - main standby switching of the instance
//
// PARAMS:
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) SwitchInstance(instanceId string) error {
	result := &bce.BceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/switchMaster").
		WithResult(result).
		Do()

	return err
}

// CreateDatabase - create a database with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//     - args: the arguments to create a account
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateDatabase(instanceId string, args *CreateDatabaseArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.DbName == "" {
		return fmt.Errorf("unset DbName")
	}

	if args.CharacterSetName == "" {
		return fmt.Errorf("unset CharacterSetName")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDatabaseUriWithInstanceId(instanceId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteDatabase - delete an database of a DDC instance
//
// PARAMS:
//     - instanceIds: the specific instanceIds
//     - dbName: the specific database's name
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteDatabase(instanceId, dbName string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDatabaseUriWithDbName(instanceId, dbName)).
		Do()
}

// UpdateDatabaseRemark - update a database remark with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//	   - dbName: the specific accountName
//     - args: the arguments to update a database remark
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateDatabaseRemark(instanceId string, dbName string, args *UpdateDatabaseRemarkArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	//if args.Remark == "" {
	//	return fmt.Errorf("unset Remark")
	//}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDatabaseUriWithDbName(instanceId, dbName)).
		WithQueryParam("remark", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// GetDatabase - get an database of a DDC instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - dbName: the specific database's name
// RETURNS:
//     - *Database: the database's meta
//     - error: nil if success otherwise the specific error
func (c *Client) GetDatabase(instanceId, dbName string) (*Database, error) {
	result := &DatabaseResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDatabaseUriWithDbName(instanceId, dbName)).
		WithResult(result).
		Do()

	return &result.Database, err
}

// ListDatabase - list all database of a DDC instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific ddc Instance's ID
// RETURNS:
//     - *ListDatabaseResult: the result of list all database, contains all databases' meta
//     - error: nil if success otherwise the specific error
func (c *Client) ListDatabase(instanceId string) (*ListDatabaseResult, error) {
	result := &ListDatabaseResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDatabaseUriWithInstanceId(instanceId)).
		WithResult(result).
		Do()

	return result, err
}

// CreateAccount - create a account with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//     - args: the arguments to create a account
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateAccount(instanceId string, args *CreateAccountArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.AccountName == "" {
		return fmt.Errorf("unset AccountName")
	}

	if args.Password == "" {
		return fmt.Errorf("unset Password")
	}

	if args.Type == "" {
		return fmt.Errorf("unset Type")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAccountUriWithInstanceId(instanceId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
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
		WithURL(getAccountUriWithAccountName(instanceId, accountName)).
		Do()
}

// UpdateAccountPassword - update a account password with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//	   - accountName: the specific accountName
//     - args: the arguments to update a account password
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountPassword(instanceId string, accountName string, args *UpdateAccountPasswordArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.Password == "" {
		return fmt.Errorf("unset Password")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAccountUriWithAccountName(instanceId, accountName)).
		WithQueryParam("password", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateAccountRemark - update a account remark with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//	   - accountName: the specific accountName
//     - args: the arguments to update a account remark
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountRemark(instanceId string, accountName string, args *UpdateAccountRemarkArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	//if args.Remark == "" {
	//	return fmt.Errorf("unset Remark")
	//}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAccountUriWithAccountName(instanceId, accountName)).
		WithQueryParam("remark", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateAccountPrivileges - update a account privileges with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//	   - accountName: the specific accountName
//     - args: the arguments to update a account privileges
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountPrivileges(instanceId string, accountName string, args *UpdateAccountPrivilegesArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAccountUriWithAccountName(instanceId, accountName)).
		WithQueryParam("privileges", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// GetAccount - get an account of a DDC instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - accountName: the specific account's name
// RETURNS:
//     - *Account: the account's meta
//     - error: nil if success otherwise the specific error
func (c *Client) GetAccount(instanceId, accountName string) (*Account, error) {
	result := &AccountResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAccountUriWithAccountName(instanceId, accountName)).
		WithResult(result).
		Do()

	return &result.Account, err
}

// ListAccount - list all account of a DDC instance with the specific parameters
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
		WithURL(getAccountUriWithInstanceId(instanceId)).
		WithResult(result).
		Do()

	return result, err
}

// ListRoGroup - list all roGroups of a DDC instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *ListRoGroupResult: All roGroups of the current instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListRoGroup(instanceId string) (*ListRoGroupResult, error) {
	result := &ListRoGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRoGroupUriWithInstanceId(instanceId) + "/list").
		WithResult(result).
		Do()

	return result, err
}

// ListVpc - list all Vpc
//
// PARAMS:
// RETURNS:
//     - *ListVpc: All vpc of
//     - error: nil if success otherwise the specific error
func (c *Client) ListVpc() (*[]VpcVo, error) {
	result := &[]VpcVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUri() + "/vpcList").
		WithResult(result).
		Do()

	return result, err
}
