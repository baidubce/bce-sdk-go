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

const (
	DEFAULT_PAGE_SIZE = 10
	DEFAULT_PAGE_NUM  = 1
)

// CreateRds - create a RDS with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a rds
//
// RETURNS:
//   - *InstanceIds: the result of create RDS, contains new RDS's instanceIds
//   - error: nil if success otherwise the specific error
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
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateReadReplica - create a readReplica RDS with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a readReplica rds
//
// RETURNS:
//   - *InstanceIds: the result of create a readReplica RDS, contains the readReplica RDS's instanceIds
//   - error: nil if success otherwise the specific error
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
		WithQueryParam("readReplica", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateRdsProxy - create a proxy RDS with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a readReplica rds
//
// RETURNS:
//   - *InstanceIds: the result of create a readReplica RDS, contains the readReplica RDS's instanceIds
//   - error: nil if success otherwise the specific error
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
		WithQueryParam("rdsproxy", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListRds - list all RDS with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all RDS
//
// RETURNS:
//   - *ListRdsResult: the result of list all RDS, contains all rds' meta
//   - error: nil if success otherwise the specific error
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
//   - instanceId: the specific rds Instance's ID
//
// RETURNS:
//   - *Instance: the specific rdsInstance's detail
//   - error: nil if success otherwise the specific error
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
//   - instanceIds: the specific instanceIds
//
// RETURNS:
//   - error: nil if success otherwise the specific error
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
//   - instanceId: the specific instanceId
//   - args: the arguments to resize an RDS
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeRds(instanceId string, args *ResizeRdsArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("resize", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// CreateAccount - create a account with the specific parameters
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to create a account
//
// RETURNS:
//   - error: nil if success otherwise the specific error
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

	cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.Password)
	if err != nil {
		return err
	}
	args.Password = cryptedPass

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/account").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ListAccount - list all account of a RDS instance with the specific parameters
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//
// RETURNS:
//   - *ListAccountResult: the result of list all account, contains all accounts' meta
//   - error: nil if success otherwise the specific error
func (c *Client) ListAccount(instanceId string) (*ListAccountResult, error) {
	result := &ListAccountResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/account").
		WithResult(result).
		Do()

	return result, err
}

// GetAccount - get an account of a RDS instance with the specific parameters
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//   - accountName: the specific account's name
//
// RETURNS:
//   - *Account: the account's meta
//   - error: nil if success otherwise the specific error
func (c *Client) GetAccount(instanceId, accountName string) (*Account, error) {
	result := &Account{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/account/" + accountName).
		WithResult(result).
		Do()

	return result, err
}

// AccountDesc - modify account's description
//
// PARAMS:
//   - instanceIds: the specific instanceIds
//   - accountName: the specific account's name
//   - args: the arguments used to modify account's description
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyAccountDesc(instanceId, accountName string, args *ModifyAccountDesc) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/account/"+accountName+"/desc").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateAccountPrivileges - upate account's privileges
//
// PARAMS:
//   - instanceIds: the specific instanceIds
//   - accountName: the specific account's name
//   - args: the arguments used to modify account's privileges
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountPrivileges(instanceId, accountName string, args *UpdateAccountPrivileges) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/account/"+accountName+"/privileges").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateAccountPassword - update account's password
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - accountName: the specific account's name
//   - args: the arguments to update account's password
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountPassword(instanceId, accountName string, args *UpdatePasswordArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
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
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/account/"+accountName+"/password").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteAccount - delete an account of a RDS instance
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - accountName: the specific account's name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAccount(instanceId, accountName string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/account/" + accountName).
		Do()
}

// RebootInstance - reboot a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance to be rebooted
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RebootInstance(instanceId string) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("reboot", "").
		Do()
}

// UpdateInstanceName - update name of a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to update instanceName
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(instanceId string, args *UpdateInstanceNameArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("rename", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateSyncMode - update sync mode of a specified instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to update syncMode
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifySyncMode(instanceId string, args *ModifySyncModeArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("modifySyncMode", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ModifyEndpoint - modify the prefix of endpoint
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to modify endpoint
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyEndpoint(instanceId string, args *ModifyEndpointArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("modifyEndpoint", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ModifyPublicAccess - modify public access
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to modify public access
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyPublicAccess(instanceId string, args *ModifyPublicAccessArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("modifyPublicAccess", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ModifyBackupPolicy - modify backup policy
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - args: the arguments to modify public access
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupPolicy(instanceId string, args *ModifyBackupPolicyArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)).
		WithQueryParam("modifyBackupPolicy", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// GetBackupList - get backup list of the instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//
// RETURNS:
//   - *GetBackupListResult: result of the backup list
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupList(instanceId string, args *GetBackupListArgs) (*GetBackupListResult, error) {

	if args == nil {
		args = &GetBackupListArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &GetBackupListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/backup").
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetBackupDetail - get backup detail of the instance's backup
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: id of the instance
//   - backupId: id of the backup
//
// RETURNS:
//   - *Snapshot: result of the backup detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupDetail(instanceId string, backupId string) (*BackupDetail, error) {
	result := &BackupDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/backup/" + backupId).
		WithResult(result).
		Do()

	return result, err
}

// DeleteBackup - delete backup detail
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - backupId: the specific backupId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteBackup(instanceId, backupId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/backup/" + backupId).
		Do()
}

// GetBinlogList - get binlog list of the instance
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - datetime: datetime of the binlog
//
// RETURNS:
//   - *BinlogResult: result of the binlog list
//   - error: nil if success otherwise the specific error
func (c *Client) GetBinlogList(instanceId, dateTime string) (*GetBinlogListResult, error) {
	result := &GetBinlogListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/binlogs/" + dateTime).
		WithResult(result).
		Do()

	return result, err
}

// GetBinlogInfo - get binlog info of the instance
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - logId: the specific logId
//   - downloadValidTimeInSec: download valid time in seconds
//
// RETURNS:
//   - *BinlogInfo: result of the binlog info
//   - error: nil if success otherwise the specific error
func (c *Client) GetBinlogInfo(instanceId, logId string, downloadValidTimeInSec string) (*GetBinlogInfoResult, error) {
	result := &GetBinlogInfoResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/binlogs/" + logId + "/" + downloadValidTimeInSec).
		WithResult(result).
		Do()

	return result, err
}

// RecoveryToSourceInstanceByDatetime - recover by datetime
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to recover the instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RecoveryToSourceInstanceByDatetime(instanceId string, args *RecoveryByDatetimeArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/recoveryToSourceInstanceByDatetime").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// RecoveryToSourceInstanceBySnapshot - recover by snapshot
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to recover the instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RecoveryToSourceInstanceBySnapshot(instanceId string, args *RecoveryBySnapshotArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/recoveryToSourceInstanceBySnapshot").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// GetZoneList - list all zone
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//
// RETURNS:
//   - *GetZoneListResult: result of the zone list
//   - error: nil if success otherwise the specific error
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
//   - cli: the client agent which can perform sending request
//   - args: the arguments to list all subnets, not necessary
//
// RETURNS:
//   - *ListSubnetsResult: result of the subnet list
//   - error: nil if success otherwise the specific error
func (c *Client) ListSubnets(args *ListSubnetsArgs) (*ListSubnetsResult, error) {
	if args == nil {
		args = &ListSubnetsArgs{}
	}

	result := &ListSubnetsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(URI_PREFIX+"/subnet").
		WithQueryParamFilter("vpcId", args.VpcId).
		WithQueryParamFilter("zoneName", args.ZoneName).
		WithResult(result).
		Do()

	return result, err
}

// GetSecurityIps - get all SecurityIps
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//
// RETURNS:
//   - *GetSecurityIpsResult: all security IP
//   - error: nil if success otherwise the specific error
func (c *Client) GetSecurityIps(instanceId string) (*GetSecurityIpsResult, error) {
	result := &GetSecurityIpsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/securityIp").
		WithResult(result).
		Do()

	return result, err
}

// UpdateSecurityIps - update SecurityIps
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//   - Etag: get latest etag by GetSecurityIps
//   - Args: all SecurityIps
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSecurityIps(instanceId, Etag string, args *UpdateSecurityIpsArgs) error {

	headers := map[string]string{"x-bce-if-match": Etag}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/securityIp").
		WithHeaders(headers).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ListParameters - list all parameters of a RDS instance
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//
// RETURNS:
//   - *ListParametersResult: the result of list all parameters
//   - error: nil if success otherwise the specific error
func (c *Client) ListParameters(instanceId string) (*ListParametersResult, error) {
	result := &ListParametersResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/parameter").
		WithResult(result).
		Do()

	return result, err
}

// UpdateParameter - update Parameter
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//   - Etag: get latest etag by ListParameters
//   - Args: *UpdateParameterArgs
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateParameter(instanceId, Etag string, args *UpdateParameterArgs) error {

	headers := map[string]string{"x-bce-if-match": Etag}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/parameter").
		WithHeaders(headers).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ParameterHistory - list all parameter history
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//
// RETURNS:
//   - *ParameterHistory: the result of paremeters history
//   - error: nil if success otherwise the specific error
func (c *Client) ParameterHistory(instanceId string) (*ParameterHistoryResult, error) {
	result := &ParameterHistoryResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/parameter/history").
		WithResult(result).
		Do()

	return result, err
}

// autoRenew - create autoRenew
//
// PARAMS:
//   - Args: *autoRenewArgs
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AutoRenew(args *AutoRenewArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUri()).
		WithQueryParam("autoRenew", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// getSlowLogDownloadTaskList
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//   - datetime: the log time. range(datetime, datetime + 24 hours)
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowLogDownloadTaskList(instanceId, datetime string) (*SlowLogDownloadTaskListResult, error) {
	fmt.Println(getRdsUriWithInstanceId(instanceId) + "/slowlogs/logList/" + datetime)
	result := &SlowLogDownloadTaskListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/slowlogs/logList/" + datetime).
		WithResult(result).
		Do()
	fmt.Println(result, err)
	return result, err
}

// getSlowLogDownloadDetail
//
// PARAMS:
//   - Args: *slowLogDownloadTaskListArgs
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowLogDownloadDetail(instanceId, logId, downloadValidTimeInSec string) (*SlowLogDownloadDetail, error) {
	result := &SlowLogDownloadDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/slowlogs/download_url/" + logId + "/" + downloadValidTimeInSec).
		WithResult(result).
		Do()
	return result, err
}

// maintaintime - update maintaintime
//
// PARAMS:
//   - Args: *maintainTimeArgs
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateMaintainTime(instanceId string, args *MaintainTimeArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/maintaintime").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// diskAutoResizeConfig - config disk auto resize
//
// PARAMS:
//   - Args: *diskAutoResizeArgs
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ConfigDiskAutoResize(instanceId, action string, args *DiskAutoResizeArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/diskAutoResize/config/"+action).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// AutoResizeConfig - show disk auto resize config
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//
// RETURNS:
//   - *AutoResizeConfigResult: the result of autoResizeConfig
//   - error: nil if success otherwise the specific error
func (c *Client) GetAutoResizeConfig(instanceId string) (*AutoResizeConfigResult, error) {
	result := &AutoResizeConfigResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/autoResizeConfig").
		WithResult(result).
		Do()

	return result, err
}

// EnableAutoExpansion - is support auto expansion
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//
// RETURNS:
//   - *supportEnableDiskAutoResize: the result of list all parameters
//   - error: nil if success otherwise the specific error
func (c *Client) EnableAutoExpansion(instanceId string) (*EnableAutoExpansionResult, error) {
	result := &EnableAutoExpansionResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/autoExpansion").
		WithResult(result).
		Do()

	return result, err
}

// AzoneMigration - azone migration
//
// PARAMS:
//   - instanceId: the specific rds Instance's ID
//   - args: the arguments to set azone migration
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AzoneMigration(instanceId string, args *AzoneMigration) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/azoneMigration").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateDatabasePort - update database port
//
// PARAMS:
//   - instanceId: id of the instance
//   - args: the arguments to update database port
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateDatabasePort(instanceId string, args *UpdateDatabasePortArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/port").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ListDatabases - list all databases
//
// PARAMS:
//   - instanceId: id of the instance
//
// RETURNS:
//   - *ListDatabasesResult: result of the database list
//   - error: nil if success otherwise the specific error
func (c *Client) ListDatabases(instanceId string) (*ListDatabasesResult, error) {
	result := &ListDatabasesResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/databases").
		WithResult(result).
		Do()

	return result, err
}

// DatabaseRemark - update database's remark
//
// PARAMS:
//   - instanceIds: the specific instanceIds
//   - dbName: the specific database's name
//   - args: the arguments used to modify database's description
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyDatabaseDesc(instanceId, dbName string, args *ModifyDatabaseDesc) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/databases/"+dbName+"/remark").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteDatabase - delete database of RDS instance
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - dbName: the specific database's name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteDatabase(instanceId, dbName string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/databases/" + dbName).
		Do()
}

// CreateDatabase - create a database with the specific parameters
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to create a database
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDatabase(instanceId string, args *CreateDatabaseArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/databases").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// TaskList - task list
//
// PARAMS:
//   - args: the arguments to list tasks
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) TaskList(args *TaskListArgs) (*TaskListResult, error) {
	result := &TaskListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/task").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListRecyclerInstance - list all recycler instance
//
// PARAMS:
//   - marker: the specific marker
//   - maxKeys: the specific max keys
//
// RETURNS:
//   - *RecyclerListResult: the result of list all recycler instance
//   - error: nil if success otherwise the specific error
func (c *Client) ListRecyclerInstance(args *ListRdsArgs) (*RecyclerListResult, error) {
	result := &RecyclerListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUri()+"/recycler/list").
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// RecyclerRecover - recover recycler instance
//
// PARAMS:
//   - args: the arguments used to recover recycler instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RecyclerRecover(args *RecyclerRecoverArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUri()+"/recycler/recover").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteRecyclerInstance - delete recycler instance
//
// PARAMS:
//   - instanceId: the specific instanceId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteRecyclerInstance(instanceId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUri() + "/recycler/" + instanceId).
		Do()
}

// CreateInstanceGroup - create instance group
//
// PARAMS:
//   - args: the arguments to group create
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceGroup(args *InstanceGroupArgs) (*CreateInstanceGroupResult, error) {
	result := &CreateInstanceGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListInstanceGroup - list all instace group
//
// PARAMS:
//   - manner: the specific manner
//   - order: asc or desc
//   - orderBy: the specific orderBy
//   - pageNo: the specific pageNo
//   - pageSize: the specific pageSize
//
// RETURNS:
//   - *InstanceGroupResult: the result of list all instance group
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstanceGroup(args *ListInstanceGroupArgs) (*InstanceGroupListResult, error) {
	result := &InstanceGroupListResult{}
	if args.PageSize == 0 {
		args.PageSize = DEFAULT_PAGE_SIZE
	}
	if args.PageNo == 0 {
		args.PageNo = DEFAULT_PAGE_NUM
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUri()+"/group").
		WithQueryParamFilter("manner", args.Manner).
		WithQueryParamFilter("order", args.Order).
		WithQueryParamFilter("orderBy", args.OrderBy).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()

	return result, err
}

// InstanceGroupDetail - show the detail of instance group
//
// PARAMS:
//
// RETURNS:
//   - *InstanceGroupDetailResult: the result of instance group detail
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupDetail(groupId string) (*InstanceGroupDetailResult, error) {
	result := &InstanceGroupDetailResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUri() + "/group/" + groupId).
		WithResult(result).
		Do()

	return result, err
}

// InstanceGroupCheckGtid - check the gtid of instance group
//
// PARAMS:
//   - args: the arguments to check gtid
//
// RETURNS:
//   - *CheckGtidResult: the result of check gtid
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupCheckGtid(args *CheckGtidArgs) (*CheckGtidResult, error) {
	result := &CheckGtidResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/checkGtid").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// InstanceGroupCheckPing - instance group connectivity check
//
// PARAMS:
//   - args: the arguments of connectivity check
//
// RETURNS:
//   - *CheckPingResult: the result of connectivity check
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupCheckPing(args *CheckPingArgs) (*CheckPingResult, error) {
	result := &CheckPingResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/checkPing").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// InstanceGroupCheckData - instance group data check
//
// PARAMS:
//   - args: the arguments of data check
//
// RETURNS:
//   - *CheckDataResult: the result of data check
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupCheckData(args *CheckDataArgs) (*CheckDataResult, error) {
	result := &CheckDataResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/checkData").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// InstanceGroupCheckVersion - instance group version check
//
// PARAMS:
//   - args: the arguments of version check
//
// RETURNS:
//   - *CheckVersionResult: the result of data check
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupCheckVersion(args *CheckVersionArgs) (*CheckVersionResult, error) {
	result := &CheckVersionResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/checkVersion").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateInstanceGroupName - update instance group name
//
// PARAMS:
//   - args: the arguments of update instance group name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceGroupName(groupId string, args *InstanceGroupNameArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/"+groupId+"/name").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// InstanceGroupAdd - add instance to instance group
//
// PARAMS:
//   - args: the arguments of add instance to instance group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupAdd(groupId string, args *InstanceGroupAddArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/"+groupId+"/instance").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// InstanceGroupBatchAdd - batch add instance to instance group
//
// PARAMS:
//   - args: the arguments of batch add instance to instance group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupBatchAdd(args *InstanceGroupBatchAddArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/batchjoin").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// InstanceGroupForceChange - force change instance
//
// PARAMS:
//   - args: the arguments used to force change instance
//
// RETURNS:
//   - *ForceChangeResult: the result of force change
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupForceChange(groupId string, args *ForceChangeArgs) (*ForceChangeResult, error) {
	result := &ForceChangeResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUri()+"/group/"+groupId+"/forceChange").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// InstanceGroupChangeLeader - change leader of instance group
//
// PARAMS:
//   - args: the arguments used to change leader of instance group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupLeaderChange(groupId string, args *GroupLeaderChangeArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUri()+"/group/"+groupId+"/instance").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// InstanceGroupRemove - remove instance to instance group
//
// PARAMS:
//   - args: the arguments of remove instance to instance group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceGroupRemove(groupId, instanceId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUri()+"/group/"+groupId+"/instance/"+instanceId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// DeleteInstanceGroup - delete instance group
//
// PARAMS:
//   - groupId: the instance group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceGroup(groupId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUri()+"/group/"+groupId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// InstanceMinorVersionList - list minor versions available for instance
//
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *MinorVersionListResult: the result of list minor versions available for instance
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceMinorVersionList(instanceId string) (*MinorVersionListResult, error) {
	result := &MinorVersionListResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL("/v1/rds/instance/" + instanceId + "/upgradeMinorVersionList").
		WithResult(result).
		Do()

	return result, err
}

// InstanceUpgradeMinorVersion - upgrade minor version
//
// PARAMS:
//   - instanceId: the instance id
//   - args: the arguments used to upgrade minor version
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceUpgradeMinorVersion(instanceId string, args *UpgradeMinorVersionArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL("/v1/rds/instance/"+instanceId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithQueryParam("upgradeMinorVersion", "").
		Do()

	return err
}

// SlowSqlFlowStatus - get slow sql flow status
//
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *SlowSqlFlowStatusResult: the result of slow sql flow status
//   - error: nil if success otherwise the specific error
func (c *Client) SlowSqlFlowStatus(instanceId string) (*SlowSqlFlowStatusResult, error) {
	result := &SlowSqlFlowStatusResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsqlflow").
		WithResult(result).
		Do()

	return result, err
}

// EnableSlowSqlFlow - enable slow sql flow
//
// PARAMS:
//   - instanceId: the instance id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) EnableSlowSqlFlow(instanceId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/smartdba/slowsqlflow").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// DisableSlowSqlFlow - disable slow sql flow
//
// PARAMS:
//   - instanceId: the instance id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DisableSlowSqlFlow(instanceId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUriWithInstanceId(instanceId)+"/smartdba/slowsqlflow").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// GetSlowSqlList - get slow sql list
//
// PARAMS:
//   - InstanceId: instance id
//   - *GetSlowSqlArgs: the arguments of slow sql list
//
// RETURNS:
//   - *SlowSqlListResult: the result of slow sql list
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlList(instanceId string, args *GetSlowSqlArgs) (*SlowSqlListResult, error) {
	result := &SlowSqlListResult{}
	if args.PageSize == 0 {
		args.PageSize = DEFAULT_PAGE_SIZE
	}
	if args.Page == 0 {
		args.Page = DEFAULT_PAGE_NUM
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/list").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlBySqlId - get slow sql by sql id
// PARAMS:
//   - InstanceId: instance id
//   - SqlId: the sql id
//
// RETURNS:
//   - *SlowSqlItem: the result of slow sql detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlBySqlId(instanceId, sqlId string) (*SlowSqlItem, error) {
	result := &SlowSqlItem{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/" + sqlId).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlExplain - get slow sql explain
// PARAMS:
//   - InstanceId: instance id
//   - SqlId: the sql id
//   - Schema: the schema
//
// RETURNS:
//   - *SlowSqlExplainResult: the result of slow sql explain
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlExplain(instanceId, sqlId, schema string) (*SlowSqlExplainResult, error) {
	result := &SlowSqlExplainResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/explain/" + sqlId + "/" + schema).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlStatsDigest - get slow sql stats digest
//
// PARAMS:
//   - InstanceId: instance id
//   - *GetSlowSqlArgs: the arguments of slow sql
//
// RETURNS:
//   - *SlowSqlDigestResult: the result of slow sql stats digest
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlStatsDigest(instanceId string, args *GetSlowSqlArgs) (*SlowSqlDigestResult, error) {
	result := &SlowSqlDigestResult{}
	if args.PageSize == 0 {
		args.PageSize = DEFAULT_PAGE_SIZE
	}
	if args.Page == 0 {
		args.Page = DEFAULT_PAGE_NUM
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/stats/digest").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlDuration - get slow sql duration
//
// PARAMS:
//   - InstanceId: instance id
//   - *GetSlowSqlDurationArgs: gs query arguments of slow sql duration
//
// RETURNS:
//   - *SlowSqlDurationResult: the result of slow sql duration
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlDuration(instanceId string, args *GetSlowSqlDurationArgs) (*SlowSqlDurationResult, error) {
	result := &SlowSqlDurationResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/stats/duration").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlSource - get slow sql source
//
// PARAMS:
//   - InstanceId: instance id
//   - *GetSlowSqlSourceArgs: gs query arguments of slow sql srouce
//
// RETURNS:
//   - *SlowSqlSourceResult: the result of slow sql souce
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlSource(instanceId string, args *GetSlowSqlSourceArgs) (*SlowSqlSourceResult, error) {
	result := &SlowSqlSourceResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/stats/source").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlSchema - get slow sql schema
// PARAMS:
//   - InstanceId: instance id
//   - SqlId: the sql id
//   - Schema: the schema
//
// RETURNS:
//   - *SlowSqlSchemaResult: the result of slow sql schema
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlSchema(instanceId, sqlId, schema string) (*SlowSqlSchemaResult, error) {
	result := &SlowSqlSchemaResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/" + sqlId + "/" + schema).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlTable - get slow sql table
// PARAMS:
//   - InstanceId: instance id
//   - SqlId: the sql id
//   - Schema: the schema
//   - Table: the table
//
// RETURNS:
//   - *SlowSqlTableResult: the result of slow sql table
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlTable(instanceId, sqlId, schema, table string) (*SlowSqlTableResult, error) {
	result := &SlowSqlTableResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/" + sqlId + "/" + schema + "/" + table).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlIndex - get slow sql index
// PARAMS:
//   - InstanceId: instance id
//   - *GetSlowSqlIndexArgs: the arguments of slow sql index
//
// RETURNS:
//   - *SlowSqlIndexResult: the result of slow sql index
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlIndex(instanceId string, args *GetSlowSqlIndexArgs) (*SlowSqlIndexResult, error) {
	result := &SlowSqlIndexResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/table/index").
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlTrend - get slow sql trend
// PARAMS:
//   - InstanceId: instance id
//   - *GetSlowSqlTrendArgs: the arguments of slow sql trend
//
// RETURNS:
//   - *SlowSqlTrendResult: the result of slow sql trend
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlTrend(instanceId string, args *GetSlowSqlTrendArgs) (*SlowSqlTrendResult, error) {
	result := &SlowSqlTrendResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/trend").
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlAdvice - get slow sql advice
// PARAMS:
//   - InstanceId: instance id
//   - SqlId: the sql id
//   - Schema: the schema
//
// RETURNS:
//   - *SlowSqlAdviceResult: the result of slow sql advice
//   - error: nil if success otherwise the specific error
func (c *Client) GetSlowSqlAdvice(instanceId, sqlId, schema string) (*SlowSqlAdviceResult, error) {
	result := &SlowSqlAdviceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/slowsql/tuning/" + sqlId + "/" + schema).
		WithResult(result).
		Do()

	return result, err
}

// GetDiskInfo - get disk info
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *DiskInfoResult: the result of disk info
//   - error: nil if success otherwise the specific error
func (c *Client) GetDiskInfo(instanceId string) (*DiskInfoResult, error) {
	result := &DiskInfoResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/disk/list").
		WithResult(result).
		Do()

	return result, err
}

// GetDbListSize - get database list size
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *DbListResult: the result of database list size
//   - error: nil if success otherwise the specific error
func (c *Client) GetDbListSize(instanceId string) (*DbListResult, error) {
	result := &DbListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/db/list").
		WithResult(result).
		Do()

	return result, err
}

// GetTableListInfo - get table list info
// PARAMS:
//   - InstanceId: instance id
//   - *GetTableListArgs: the arguments of table list info
//
// RETURNS:
//   - *TableListResult: the result of table list info
//   - error: nil if success otherwise the specific error
func (c *Client) GetTableListInfo(instanceId string, args *GetTableListArgs) (*TableListResult, error) {
	result := &TableListResult{}
	if args.PageSize == 0 {
		args.PageSize = DEFAULT_PAGE_SIZE
	}
	if args.PageNo == 0 {
		args.PageNo = DEFAULT_PAGE_NUM
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/tb/list").
		WithResult(result).
		Do()

	return result, err
}

// GetKillSessionTypes - get kill session types
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *KillSessionTypesResult: the result of the kill session types
//   - error: nil if success otherwise the specific error
func (c *Client) GetKillSessionTypes(instanceId string) (*KillSessionTypesResult, error) {
	result := &KillSessionTypesResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/session/kill/types").
		WithResult(result).
		Do()

	return result, err
}

// GetSessionSummary - get session summary
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *SessionSummaryResult: the result of the session summary
//   - error: nil if success otherwise the specific error
func (c *Client) GetSessionSummary(instanceId string) (*SessionSummaryResult, error) {
	result := &SessionSummaryResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/session/summary").
		WithResult(result).
		Do()

	return result, err
}

// GetSessionDetail - get session detail
// PARAMS:
//   - InstanceId: instance id
//   - *SessionDetailArgs: the arguments of session detail
//
// RETURNS:
//   - *SessionDetailResult: the result of the session detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSessionDetail(instanceId string, args *SessionDetailArgs) (*SessionDetailResult, error) {
	result := &SessionDetailResult{}
	if args.PageSize == 0 {
		args.PageSize = DEFAULT_PAGE_SIZE
	}
	if args.Page == 0 {
		args.Page = DEFAULT_PAGE_NUM
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/session/detail").
		WithResult(result).
		Do()

	return result, err
}

// CheckKillSessionAuth - check kill session auth
// PARAMS:
//   - InstanceId: instance id
//   - *KillSessionAuthArgs: the arguments of kill session auth
//
// RETURNS:
//   - *KillSessionAuthResult: the result of kill session auth
//   - error: nil if success otherwise the specific error
func (c *Client) CheckKillSessionAuth(instanceId string, args *KillSessionAuthArgs) (*KillSessionAuthResult, error) {
	result := &KillSessionAuthResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/session/kill/authority").
		WithResult(result).
		Do()

	return result, err
}

// GetKillSessionHistory - get kill session history
// PARAMS:
//   - InstanceId: instance id
//   - *KillSessionHistory: the arguments of kill session history
//
// RETURNS:
//   - *KillSessionHistoryResult: the result of kill session history
//   - error: nil if success otherwise the specific error
func (c *Client) GetKillSessionHistory(instanceId string, args *KillSessionHistory) (*KillSessionHistoryResult, error) {
	result := &KillSessionHistoryResult{}
	if args.PageSize == 0 {
		args.PageSize = DEFAULT_PAGE_SIZE
	}
	if args.Page == 0 {
		args.Page = DEFAULT_PAGE_NUM
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/session/kill/history").
		WithResult(result).
		Do()

	return result, err
}

// KillSession - kill session
// PARAMS:
//   - InstanceId: instance id
//   - *KillSessionArgs: the arguments of kill session
//
// RETURNS:
//   - *KillSessionResult: the result of kill session
//   - error: nil if success otherwise the specific error
func (c *Client) KillSession(instanceId string, args *KillSessionArgs) (*KillSessionResult, error) {
	result := &KillSessionResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/session/kill").
		WithResult(result).
		Do()

	return result, err
}

// GetSessionStatistics - get session statistics
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *SessionStatisticsResult: the result of the session statistics
//   - error: nil if success otherwise the specific error
func (c *Client) GetSessionStatistics(instanceId string) (*SessionStatisticsResult, error) {
	result := &SessionStatisticsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/session/statistics").
		WithResult(result).
		Do()

	return result, err
}

// GetErrorLogStatus - get error log status
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *ErrorLogStatusResult: the result of error log status
//   - error: nil if success otherwise the specific error
func (c *Client) GetErrorLogStatus(instanceId string) (*ErrorLogStatusResult, error) {
	result := &ErrorLogStatusResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/errorlogflow").
		WithResult(result).
		Do()

	return result, err
}

// EnableErrorLog - enable error log
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *ErrorLogResult: the result of enable error log
//   - error: nil if success otherwise the specific error
func (c *Client) EnableErrorLog(instanceId string) (*ErrorLogResult, error) {
	result := &ErrorLogResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/errorlogflow").
		WithResult(result).
		Do()

	return result, err
}

// DisableErrorLog - disable error log
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *ErrorLogResult: the result of disable error log
//   - error: nil if success otherwise the specific error
func (c *Client) DisableErrorLog(instanceId string) (*ErrorLogResult, error) {
	result := &ErrorLogResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/errorlogflow").
		WithResult(result).
		Do()

	return result, err
}

// GetErrorLogList - get error log list
// PARAMS:
//   - InstanceId: instance id
//   - *ErrorLogListArgs: error log list arguments
//
// RETURNS:
//   - *ErrorLogListResult: the result of error log list
//   - error: nil if success otherwise the specific error
func (c *Client) GetErrorLogList(instanceId string, args *ErrorLogListArgs) (*ErrorLogListResult, error) {
	result := &ErrorLogListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/smartdba/errorlog/detail").
		WithResult(result).
		Do()

	return result, err
}

// GetSqlFilterList - get sql filter list
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *SqlFilterListResult: the result of sql filter list
//   - error: nil if success otherwise the specific error
func (c *Client) GetSqlFilterList(instanceId string) (*SqlFilterListResult, error) {
	result := &SqlFilterListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/sqlfilter").
		WithResult(result).
		Do()

	return result, err
}

// GetSqlFilterDetail - get sql filter detail
// PARAMS:
//   - InstanceId: instance id
//   - SqlFilterId: sql filter id
//
// RETURNS:
//   - *SqlFilterDetailResult: the result of sql filter detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSqlFilterDetail(instanceId, sqlFilterId string) (*SqlFilterItem, error) {
	result := &SqlFilterItem{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/sqlfilter/" + sqlFilterId).
		WithResult(result).
		Do()

	return result, err
}

// AddSqlFilter - add sql filter
// PARAMS:
//   - InstanceId: instance id
//   - *AddSqlFilterArgs: add sql filter arguments
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AddSqlFilter(instanceId string, args *SqlFilterArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/sqlfilter").
		Do()

	return err
}

// UpdateSqlFilter - update sql filter
// PARAMS:
//   - InstanceId: instance id
//   - SqlFilterId: sql filter id
//   - *AddSqlFilterArgs: add sql filter arguments
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSqlFilter(instanceId, sqlFilterId string, args *SqlFilterArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/sqlfilter/" + sqlFilterId).
		Do()

	return err
}

// StartOrStopSqlFilter - start or stop sql filter
//   - InstanceId: instance id
//   - SqlFilterId: sql filter id
//   - *StartOrStopSqlFilterArgs: start or stop sql filter arguments
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) StartOrStopSqlFilter(instanceId, sqlFilterId string, args *StartOrStopSqlFilterArgs) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/sqlfilter/" + sqlFilterId).
		Do()

	return err
}

// DeleteSqlFilter - delete sql filter
//   - InstanceId: instance id
//   - SqlFilterId: sql filter id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSqlFilter(instanceId, sqlFilterId string) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/sqlfilter/" + sqlFilterId).
		Do()

	return err
}

// IsAllowedSqlFilter - check if sql filter is allowed
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *IsAllowedResult: the result of sql filter allowed
//   - error: nil if success otherwise the specific error
func (c *Client) IsAllowedSqlFilter(instanceId string) (*IsAllowedResult, error) {
	result := &IsAllowedResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/sqlfilter/allowed").
		WithResult(result).
		Do()

	return result, err
}

// ProcessKill - kill process
// PARAMS:
//   - InstanceId: instance id
//   - *ProcessArgs: process arguments
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ProcessKill(instanceId string, args *ProcessArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/performance/process/kill").
		Do()

	return err
}

// InnodbStatus - get innodb status
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *InnodbStatusResult: the result of innodb status
//   - error: nil if success otherwise the specific error
func (c *Client) InnodbStatus(instanceId string) (*InnodbStatusResult, error) {
	result := &InnodbStatusResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/performance/innodbstatus").
		WithResult(result).
		Do()

	return result, err
}

// ProcessList - get process list
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *ProcessListResult: the result of process list
//   - error: nil if success otherwise the specific error
func (c *Client) ProcessList(instanceId string) (*ProcessListResult, error) {
	result := &ProcessListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/performance/processlist").
		WithResult(result).
		Do()

	return result, err
}

// TransactionList - get transaction list
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *TransactionListResult: the result of transaction list
//   - error: nil if success otherwise the specific error
func (c *Client) TransactionList(instanceId string) (*TransactionListResult, error) {
	result := &TransactionListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/performance/transaction").
		WithResult(result).
		Do()

	return result, err
}

// ConnectionList - get connection list
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *ConnectionListResult: the result of transaction list
//   - error: nil if success otherwise the specific error
func (c *Client) ConnectionList(instanceId string) (*ConnectionListResult, error) {
	result := &ConnectionListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRdsUriWithInstanceId(instanceId) + "/performance/connection").
		WithResult(result).
		Do()

	return result, err
}

// FailInjectWhiteList - get fail inject white list
//
// RETURNS:
//   - *FailInjectWhiteListResult: the result of transaction list
//   - error: nil if success otherwise the specific error
func (c *Client) FailInjectWhiteList() (*FailInjectWhiteListResult, error) {
	result := &FailInjectWhiteListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL("/v1/failinject/whitelist").
		WithResult(result).
		Do()

	return result, err
}

// AddToFailInjectWhiteList - add instance to failinject whitelist
// PARAMS:
//   - *AddFailInjectArgs: add failinject args
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AddToFailInjectWhiteList(args *FailInjectArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL("/v1/failinject/whitelist").
		Do()

	return err
}

// RemoveFailInjectWhiteList - remove instance to failinject whitelist
// PARAMS:
//   - *RemoveFailInjectArgs: remove failinject args
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RemoveFailInjectWhiteList(args *FailInjectArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL("/v1/failinject/whitelist/remove").
		Do()

	return err
}

// FailInjectStart - start failinject instance
// PARAMS:
//   - InstanceId: instance id
//
// RETURNS:
//   - *TaskResult: the result of task
//   - error: nil if success otherwise the specific error
func (c *Client) FailInjectStart(instanceId string) (*TaskResult, error) {
	result := &TaskResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL("/v1/failinject/" + instanceId).
		WithResult(result).
		Do()

	return result, err
}

// GetOrderStatus - get order status
// PARAMS:
//   - OrderId: order id
//
// RETURNS:
//   - *OrderStatusResult: the result of order status
//   - error: nil if success otherwise the specific error
func (c *Client) GetOrderStatus(orderId string) (*OrderStatusResult, error) {
	result := &OrderStatusResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL("/v1/instance/order/" + orderId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) Request(method, uri string, body interface{}) (interface{}, error) {
	res := struct{}{}
	req := bce.NewRequestBuilder(c).
		WithMethod(method).
		WithURL(uri).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	var err error
	if body != nil {
		err = req.
			WithBody(body).
			WithResult(&res).
			Do()
	} else {
		err = req.
			WithResult(&res).
			Do()
	}

	return res, err
}
