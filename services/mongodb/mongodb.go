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

// mongodb.go - the mongodb APIs definition supported by the MONGODB service
package mongodb

import (
	"errors"
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	DEFAULT_PAGE_SIZE = 10
	DEFAULT_PAGE_NUM  = 1
	S_SHARDING        = "sharding"
	S_REPLICA         = "replica"
	S_POSTPAID        = "Postpaid"
	S_PREPAID         = "Prepaid"
)

// ListMongodb - list MONGODB with the specific parameters
//
// PARAMS:
//   - args: the arguments to list MONGODB
//
// RETURNS:
//   - *ListMongodbResult: the result of list MONGODB, contains mongodb' meta
//   - error: nil if success otherwise the specific error
func (c *Client) ListMongodb(args *ListMongodbArgs) (*ListMongodbResult, error) {
	if args == nil {
		args = &ListMongodbArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListMongodbResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getMongodbUri()).
		WithQueryParamFilter("manner", "marker").
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParamFilter("engineVersion", args.EngineVersion).
		WithQueryParamFilter("storageEngine", args.StorageEngine).
		WithQueryParamFilter("dbInstanceType", args.DbInstanceType).
		WithResult(result).
		Do()

	return result, err
}

// GetInstanceDetail - get a specific mongodb Instance's detail
//
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//
// RETURNS:
//   - *Instance: the specific mongodbInstance's detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceId string) (*InstanceDetail, error) {
	result := &InstanceDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithResult(result).
		Do()
	return result, err
}

// CreateReplica - create a Replica MONGODB with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a mongodb
//
// RETURNS:
//   - *CreateResult: the result of create MONGODB
//   - error: nil if success otherwise the specific error
func (c *Client) CreateReplica(args *CreateReplicaArgs) (*CreateResult, error) {
	if args == nil {
		return nil, errors.New("unset args")
	}

	if args.StorageEngine == "" {
		return nil, errors.New("unset StorageEngine")
	}

	if args.EngineVersion == "" {
		return nil, errors.New("unset EngineVersion")
	}

	if args.DbInstanceCpuCount <= 0 {
		return nil, errors.New("unset DbInstanceCpuCount")
	}

	if args.DbInstanceMemoryCapacity <= 0 {
		return nil, errors.New("unset DbInstanceMemoryCapacity")
	}

	if args.DbInstanceStorage <= 0 {
		return nil, errors.New("unset DbInstanceStorage")
	}

	if args.VotingMemberNum <= 0 {
		return nil, errors.New("unset VotingMemberNum")
	}

	if args.Billing.PaymentTiming == "" {
		return nil, errors.New("unset PaymentTiming")
	}

	if args.Billing.PaymentTiming == S_POSTPAID {
		args.Billing.Reservation.ReservationLength = 0
		args.Billing.Reservation.ReservationTimeUnit = "Month"
	}

	if len(args.AccountPassword) > 0 {
		cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AccountPassword)
		if err != nil {
			return nil, err
		}
		args.AccountPassword = cryptedPass
	}

	args.DbInstanceType = S_REPLICA

	result := &CreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMongodbUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateSharding - create a Sharding MONGODB with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a mongodb
//
// RETURNS:
//   - *CreateResult: the result of create MONGODB
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSharding(args *CreateShardingArgs) (*CreateResult, error) {
	if args == nil {
		return nil, errors.New("unset args")
	}

	if args.StorageEngine == "" {
		return nil, errors.New("unset StorageEngine")
	}

	if args.EngineVersion == "" {
		return nil, errors.New("unset EngineVersion")
	}

	if args.MongosCpuCount <= 0 {
		return nil, errors.New("unset MongosCpuCount")
	}

	if args.ShardCpuCount <= 0 {
		return nil, errors.New("unset ShardCpuCount")
	}

	if args.ShardMemoryCapacity <= 0 {
		return nil, errors.New("unset ShardMemoryCapacity")
	}

	if args.ShardStorage <= 0 {
		return nil, errors.New("unset ShardStorage")
	}

	if args.Billing.PaymentTiming == "" {
		return nil, errors.New("unset PaymentTiming")
	}

	if args.Billing.PaymentTiming == S_POSTPAID {
		args.Billing.Reservation.ReservationLength = 0
		args.Billing.Reservation.ReservationTimeUnit = "Month"
	}

	if len(args.AccountPassword) > 0 {
		cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AccountPassword)
		if err != nil {
			return nil, err
		}
		args.AccountPassword = cryptedPass
	}

	args.DbInstanceType = S_SHARDING

	result := &CreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMongodbUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ReleaseMongodbs - release MONGODBs / 将若干个MONGODB实例放入回收站
//
// PARAMS:
//   - instanceIds: the specific instanceIds
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ReleaseMongodbs(instanceIds []string) error {
	ids := strings.Join(instanceIds, ",")
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getMongodbUri()).
		WithQueryParamFilter("dbInstanceIds", ids).
		Do()
}

// ReleaseMongodb - release a MONGODB / 将一个MONGODB实例放入回收站
//
// PARAMS:
//   - instanceId: the specific instanceId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ReleaseMongodb(instanceId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		Do()
}

// DeleteMongodbs - delete MONGODBs / 从回收站删除若干个MONGODB实例
//
// PARAMS:
//   - instanceIds: the specific instanceIds
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteMongodbs(instanceIds []string) error {
	ids := strings.Join(instanceIds, ",")
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMongodbUri()+"/deletePermanent").
		WithQueryParamFilter("dbInstanceIds", ids).
		Do()
}

// DeleteMongodb - delete a MONGODB / 从回收站删除一个MONGODB实例
//
// PARAMS:
//   - instanceId: the specific instanceId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteMongodb(instanceId string) error {
	ids := []string{instanceId}
	return c.DeleteMongodbs(ids)
}

// RecoverMongodbs - delete MONGODBs / 从回收站恢复若干个MONGODB实例
//
// PARAMS:
//   - instanceIds: the specific instanceIds
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RecoverMongodbs(instanceIds []string) error {
	ids := strings.Join(instanceIds, ",")
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMongodbUri()+"/recover").
		WithQueryParamFilter("dbInstanceIds", ids).
		Do()
}

// ShardingAddComponent - create a Component with the specific parameters / 分片集实例新增组件
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to create a Component
//
// RETURNS:
//   - *ShardingAddComponentResult: the result of Add Component
//   - error: nil if success otherwise the specific error
func (c *Client) ShardingAddComponent(instanceId string, args *ShardingAddComponentArgs) (*ShardingAddComponentResult, error) {
	if args == nil {
		return nil, errors.New("unset args")
	}

	if args.NodeType == "" {
		return nil, errors.New("unset NodeType")
	}

	if args.NodeCpuCount <= 0 {
		return nil, errors.New("unset NodeCpuCount")
	}

	if args.NodeMemoryCapacity <= 0 {
		return nil, errors.New("unset NodeMemoryCapacity")
	}

	if args.NodeStorage <= 0 && args.NodeType == S_SHARDING {
		return nil, errors.New("unset NodeStorage for sharding")
	}

	result := &ShardingAddComponentResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNodeUriWithInstanceId(instanceId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ReplicaResize - Resize a Replica with the specific parameters / 副本集实例改配
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to Resize a Replica
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ReplicaResize(instanceId string, args *ReplicaResizeArgs) error {
	if args == nil {
		return errors.New("unset args")
	}

	if args.DbInstanceCpuCount <= 0 {
		return errors.New("unset DbInstanceCpuCount")
	}

	if args.DbInstanceMemoryCapacity <= 0 {
		return errors.New("unset DbInstanceMemoryCapacity")
	}

	if args.DbInstanceStorage <= 0 {
		return errors.New("unset DbInstanceStorage")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ShardingComponentResize - Resize a Sharding Component with the specific parameters / 分片集实例组件改配
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - nodeId: the specific nodeId
//   - args: the arguments to Resize a Sharding Component
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ShardingComponentResize(instanceId string, nodeId string, args *ShardingComponentResizeArgs) error {
	if args == nil {
		return errors.New("unset args")
	}

	if args.NodeCpuCount <= 0 {
		return errors.New("unset NodeCpuCount")
	}

	if args.NodeMemoryCapacity <= 0 {
		return errors.New("unset NodeMemoryCapacity")
	}
	args.DbInstanceId = instanceId
	args.NodeId = nodeId

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getNodeUriWithNodeId(instanceId, nodeId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// RestartMongodbs - Restart MONGODBs
//
// PARAMS:
//   - instanceIds: the specific instanceIds
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RestartMongodbs(instanceIds []string) error {
	args := &RestartMongodbsArgs{
		DbInstanceIds: instanceIds,
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUri()).
		WithQueryParam("restart", "").
		WithBody(args).
		Do()
}

// RestartMongodb - Restart a MONGODB
//
// PARAMS:
//   - instanceId: the specific instanceId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RestartMongodb(instanceId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("restart", "").
		Do()
}

// RestartShardingComponent - Restart a Sharding Component
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - nodeId: the specific nodeId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RestartShardingComponent(instanceId string, nodeId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getNodeUriWithNodeId(instanceId, nodeId)).
		WithQueryParam("restart", "").
		Do()
}

// UpdateInstanceName - update name of a specified instance
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to update instanceName
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(instanceId string, args *UpdateInstanceNameArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("modifyName", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateShardingComponentName - update name of a specified Component
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - nodeId: the specific nodeId
//   - args: the arguments to update Component Name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateShardingComponentName(instanceId string, nodeId string, args *UpdateComponentNameArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getNodeUriWithNodeId(instanceId, nodeId)).
		WithQueryParam("modifyName", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// UpdateAccountPassword - update account's password
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to update account's password
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountPassword(instanceId string, args *UpdatePasswordArgs) error {
	if args == nil {
		return errors.New("unset args")
	}

	if args.AccountPassword == "" {
		return errors.New("unset AccountPassword")
	}

	cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AccountPassword)
	if err != nil {
		return err
	}
	args.AccountPassword = cryptedPass

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("resetPassword", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// ReplicaSwitch - Switch a Replica / 副本集实例主从切换
//
// PARAMS:
//   - instanceId: the specific instanceId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ReplicaSwitch(instanceId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("switchHA", "").
		Do()
}

// ShardingComponentSwitch - Switch a Sharding / 分片集实例组件主从切换
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - nodeId: the specific nodeId
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ShardingComponentSwitch(instanceId string, nodeId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getNodeUriWithNodeId(instanceId, nodeId)).
		WithQueryParam("switchHA", "").
		Do()
}

// ReplicaAddReadonlyNodes - Replica Add some ReadonlyNodes with the specific parameters / 副本集实例添加只读节点
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to Replica Add some ReadonlyNodes
//
// RETURNS:
//   - *ReplicaAddReadonlyNodesResult: the result of Replica Add some ReadonlyNodes
//   - error: nil if success otherwise the specific error
func (c *Client) ReplicaAddReadonlyNodes(instanceId string, args *ReplicaAddReadonlyNodesArgs) (*ReplicaAddReadonlyNodesResult, error) {
	if args == nil {
		return nil, errors.New("unset args")
	}

	if args.ReadonlyNodeNum <= 0 {
		return nil, errors.New("unset ReadonlyNodeNum")
	}

	result := &ReplicaAddReadonlyNodesResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// GetReadonlyNodes - get ReadonlyNodes list
//
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//
// RETURNS:
//   - *GetReadonlyNodesResult: the specific ReadonlyNodes list
//   - error: nil if success otherwise the specific error
func (c *Client) GetReadonlyNodes(instanceId string) (*GetReadonlyNodesResult, error) {
	result := &GetReadonlyNodesResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getMongodbUriWithInstanceId(instanceId) + "/getReadonlyNodes").
		WithResult(result).
		Do()
	return result, err
}

// MigrateAzone - Migrate Azone of MONGODB with the specific parameters / 迁移可用区
//
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//   - args: the arguments to create a mongodb
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) MigrateAzone(instanceId string, args *MigrateAzoneArgs) error {
	if args == nil {
		return errors.New("unset args")
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("migrateAzone", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// InstanceAssignTags - update tags of instance / 全量更新实例绑定的标签，即覆盖更新
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//   - tags: tag args
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceAssignTags(instanceId string, tags []TagModel) error {
	args := &AssignTagArgs{
		Resources: []LogicAssignResource{
			{
				ResourceId:  instanceId,
				ServiceType: "MONGODB",
				Tags:        tags,
			},
		},
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL(getMongodbUri() + "/tag/assign").
		Do()
	return err
}

// InstanceBindTags - add tags of instance / 增加实例绑定的标签，即追加更新
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//   - tags: tag args
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceBindTags(instanceId string, tags []TagModel) error {
	args := &UpdateTagArgs{
		DbInstanceId: instanceId,
		Tags:         tags,
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("bindTag", "").
		Do()
	return err
}

// InstanceUnbindTags - unbind tags of instance / 减少实例绑定的标签
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//   - tags: tag args
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceUnbindTags(instanceId string, tags []TagModel) error {
	args := &UpdateTagArgs{
		DbInstanceId: instanceId,
		Tags:         tags,
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithBody(args).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("unBindTag", "").
		Do()
	return err
}

// CreateBackup - Create a Backup of mongodb Instance
//
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//   - backupMethod: backup method
//   - backupDescription: backup description
//
// RETURNS:
//   - *CreateBackupResult: Backup Result
//   - error: nil if success otherwise the specific error
func (c *Client) CreateBackup(instanceId string, backupMethod string, backupDescription string) (*CreateBackupResult, error) {
	result := &CreateBackupResult{}
	args := CreateBackupArgs{
		BackupMethod:      backupMethod,
		BackupDescription: backupDescription,
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBackupUriWithInstanceId(instanceId)).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// ListBackup - list all Backup with the specific parameters
//
// PARAMS:
//   - instanceId: the specific mongodb Instance's ID
//   - args: the arguments to list all MONGODB backup
//
// RETURNS:
//   - *ListBackupResult: the result of list all Backup
//   - error: nil if success otherwise the specific error
func (c *Client) ListBackup(instanceId string, args *ListBackupArgs) (*ListBackupResult, error) {
	if args == nil {
		args = &ListBackupArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListBackupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackupUriWithInstanceId(instanceId)).
		WithQueryParamFilter("manner", "marker").
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetBackupDetail - get backup detail of the instance's backup
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - backupId: id of the backup
//
// RETURNS:
//   - *BackupDetail: result of the backup detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupDetail(instanceId string, backupId string) (*BackupDetail, error) {
	result := &BackupDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackupUriWithBackupId(instanceId, backupId)).
		WithResult(result).
		Do()

	return result, err
}

// ModifyBackupDescription - modify backup description
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - backupId: id of the backup
//   - args: the arguments to modify backup description
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupDescription(instanceId string, backupId string, args *ModifyBackupDescriptionArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBackupUriWithBackupId(instanceId, backupId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteBackup - delete Backup
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - backupId: id of the backup
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteBackup(instanceId string, backupId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getBackupUriWithBackupId(instanceId, backupId)).
		Do()
}

// GetBackupPolicy - get BackupPolicy of a specific mongodb Instance
//
// PARAMS:
//   - instanceId: the specific instanceId
//
// RETURNS:
//   - *BackupPolicy: the Backup Policy of the instance
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupPolicy(instanceId string) (*BackupPolicy, error) {
	result := &BackupPolicy{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getMongodbUriWithInstanceId(instanceId) + "/backupPolicy").
		WithResult(result).
		Do()
	return result, err
}

// ModifyBackupPolicy - modify backup policy
//
// PARAMS:
//   - instanceId: the specific instanceId
//   - args: the arguments to modify backup policy
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupPolicy(instanceId string, args *BackupPolicy) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)+"/backupPolicy").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// GetSecurityIps - get all SecurityIps
//
// PARAMS:
//   - instanceId: the specific Instance's ID
//
// RETURNS:
//   - *SecurityIpModel: all security IP
//   - error: nil if success otherwise the specific error
func (c *Client) GetSecurityIps(instanceId string) (*SecurityIpModel, error) {
	result := &SecurityIpModel{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("describeSecurityIps", "").
		WithResult(result).
		Do()

	return result, err
}

// AddSecurityIps - add SecurityIps
//
// PARAMS:
//   - instanceId: the specific Instance's ID
//   - Args: SecurityIps
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AddSecurityIps(instanceId string, args *SecurityIpModel) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("addSecurityIps", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// DeleteSecurityIps - delete SecurityIps
//
// PARAMS:
//   - instanceId: the specific Instance's ID
//   - args: the arguments to delete SecurityIps
//
// RETURNS:
//   - *SecurityIpModel: security IP
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityIps(instanceId string, args *SecurityIpModel) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("deleteSecurityIps", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
}

// StartLogging - start Logging
// PARAMS:
//   - instanceId: the specific Instance's ID
//   - args: the arguments to start Logging
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) StartLogging(instanceId string, args *StartLoggingArgs) error {
	if args == nil {
		return errors.New("unset args")
	}
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMongodbUriWithInstanceId(instanceId)).
		WithQueryParam("startLogging", "").
		WithBody(args).
		Do()
}

// ListLogFiles - get log list
// PARAMS:
//   - instanceId: the specific Instance's ID
//   - args: the arguments to get log list
//
// RETURNS:
//   - *ListLogFilesResult: the result of log list
//   - error: nil if success otherwise the specific error
func (c *Client) ListLogFiles(instanceId string, args *ListLogFilesArgs) (*ListLogFilesResult, error) {
	result := &ListLogFilesResult{}
	if args == nil {
		return result, errors.New("unset args")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithBody(args).
		WithURL(getMongodbUriWithInstanceId(instanceId)+"/log").
		WithQueryParam("listLogFiles", "").
		WithQueryParamFilter("memberId", args.MemberId).
		WithQueryParamFilter("type", args.Type).
		WithQueryParamFilter("startTime", args.StartTime).
		WithQueryParamFilter("endTime", args.EndTime).
		WithResult(result).
		Do()

	return result, err
}
