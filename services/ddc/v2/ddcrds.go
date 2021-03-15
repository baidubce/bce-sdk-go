package ddcrds

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/ddc/ddc_util"
	"github.com/baidubce/bce-sdk-go/services/rds"
	"strconv"
	"strings"
)

// Int int point helper
func Int(value string) *int {
	if len(value) < 1 {
		return nil
	}
	intPtr, err := strconv.Atoi(value)
	if err != nil {
		panic("please pass valid int value ")
	}
	return &intPtr
}

func Json(v interface{}) string {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		panic("convert to json faild")
	}
	return string(jsonStr)
}

func isDDC(productType string) bool {
	return productType == "ddc" || productType == "DDC"
}

// 根据实例ID判断产品类型是否为DDC
func isDDCId(instanceId string) bool {
	if strings.HasPrefix(instanceId, "rds") {
		return false
	}
	return true
}

func getDDCAndRDSIds(instanceIds string) (string, string) {
	instanceIdArr := strings.Split(instanceIds, ",")
	ddcIds, rdsIds := "", ""
	if len(instanceIdArr) > 0 {
		for _, id := range instanceIdArr {
			if isDDCId(id) {
				ddcIds += id + ","
			} else {
				rdsIds += id + ","
			}
		}
	}
	if strings.HasSuffix(ddcIds, ",") {
		ddcIds = ddcIds[:len(ddcIds)-1]
	}
	if strings.HasSuffix(rdsIds, ",") {
		rdsIds = rdsIds[:len(rdsIds)-1]
	}
	return ddcIds, rdsIds
}

// CreateInstance - create a Instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a instance
// RETURNS:
//     - *InstanceIds: the result of create RDS, contains new RDS's instanceIds
//     - error: nil if success otherwise the specific error
func (c *Client) CreateRds(args *CreateRdsArgs, productType string) (*CreateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}
	// 限制InstanceName
	if strings.Contains(args.InstanceName, ".") {
		return nil, fmt.Errorf("invalid InstanceName:. not support")
	}
	var result *CreateResult
	var err error
	if isDDC(productType) {
		result, err = c.ddcClient.CreateRds(args)
	} else {
		rdsArgs := &rds.CreateRdsArgs{}
		// copy 请求参数
		err = ddc_util.SimpleCopyProperties(rdsArgs, args)
		if err != nil {
			return nil, err
		}
		var rdsRes *rds.CreateResult
		rdsRes, err = c.rdsClient.CreateRds(rdsArgs)
		if rdsRes != nil {
			result = &CreateResult{InstanceIds: rdsRes.InstanceIds}
		}
	}

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
	if len(args.SourceInstanceId) < 1 {
		return nil, fmt.Errorf("unset SourceInstanceId")
	}

	// 默认创建1个只读实例
	if args.PurchaseCount < 1 {
		args.PurchaseCount = 1
	}
	var result *CreateResult
	var err error
	if isDDCId(args.SourceInstanceId) {
		result, err = c.ddcClient.CreateReadReplica(args)
	} else {
		rdsArgs := &rds.CreateReadReplicaArgs{}
		// copy 请求参数
		err = ddc_util.SimpleCopyProperties(rdsArgs, args)
		if err != nil {
			return nil, err
		}
		var rdsRes *rds.CreateResult
		rdsRes, err = c.rdsClient.CreateReadReplica(rdsArgs)
		if rdsRes != nil {
			result = &CreateResult{InstanceIds: rdsRes.InstanceIds}
		}
	}
	return result, err
}

// UpdateRoGroup - update a roGroup
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateRoGroup(roGroupId string, args *UpdateRoGroupArgs, productType string) error {
	if isDDC(productType) {
		return c.ddcClient.UpdateRoGroup(roGroupId, args)
	}
	return RDSNotSupportError()
}

// UpdateRoGroupReplicaWeight- update repica weight in roGroup
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateRoGroupReplicaWeight(roGroupId string, args *UpdateRoGroupWeightArgs, productType string) error {
	if isDDC(productType) {
		return c.ddcClient.UpdateRoGroupReplicaWeight(roGroupId, args)
	}
	return RDSNotSupportError()
}

// ReBalanceRoGroup- Initiate a rebalance for foGroup
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ReBalanceRoGroup(roGroupId, productType string) error {
	if len(roGroupId) < 1 {
		return fmt.Errorf("unset roGroupId")
	}
	if isDDC(productType) {
		return c.ddcClient.ReBalanceRoGroup(roGroupId)
	}
	return RDSNotSupportError()
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
	if len(args.SourceInstanceId) < 1 {
		return nil, fmt.Errorf("unset SourceInstanceId")
	}

	if isDDCId(args.SourceInstanceId) {
		return nil, DDCNotSupportError()
	}
	// copy请求参数
	rdsArgs := &rds.CreateRdsProxyArgs{}
	err := ddc_util.SimpleCopyProperties(rdsArgs, args)
	if err != nil {
		return nil, err
	}
	// copy返回结果
	rdsRes, err := c.rdsClient.CreateRdsProxy(rdsArgs)
	if err != nil {
		return nil, err
	}
	result := &CreateResult{InstanceIds: rdsRes.InstanceIds}
	return result, err
}

func (c *Client) listRdsInstance(marker *ListRdsArgs) (*ListRdsResult, error) {
	rdsArgs := &rds.ListRdsArgs{}
	// copy请求参数
	err := ddc_util.SimpleCopyProperties(rdsArgs, marker)
	if err != nil {
		return nil, err
	}
	var rdsRes *rds.ListRdsResult
	rdsRes, err = c.rdsClient.ListRds(rdsArgs)
	// copy返回结果
	result := &ListRdsResult{}
	if rdsRes != nil {
		err = ddc_util.SimpleCopyProperties(result, rdsRes)
	}
	if result.Instances != nil &&
		len(result.Instances) > 0 {
		for i := range result.Instances {
			err = convertRdsInstance(&result.Instances[i])
		}
	}
	return result, err
}

// ListRds - list all instances
// RETURNS:
//     - *ListRdsResult: the result of list instances with marker
//     - error: nil if success otherwise the specific error
func (c *Client) ListRds(marker *ListRdsArgs) (*ListRdsResult, error) {
	var result *ListRdsResult
	var err error
	// 先获取DDC列表
	if len(marker.Marker) < 1 || marker.Marker == "-1" || isDDCId(marker.Marker) {
		result, err = c.ddcClient.ListRds(marker)
		if err != nil {
			return nil, err
		}
		// 数量不够时获取RDS列表
		if len(result.Instances) < marker.MaxKeys {
			// 修改marker
			marker.MaxKeys = marker.MaxKeys - len(result.Instances)
			rdsResult, err := c.listRdsInstance(marker)
			if err != nil {
				return nil, err
			}
			// 合并结果到result
			if rdsResult != nil && len(rdsResult.Instances) > 0 {
				result.Instances = append(result.Instances, rdsResult.Instances...)
				result.Marker = rdsResult.Marker
				result.IsTruncated = rdsResult.IsTruncated
				result.NextMarker = rdsResult.NextMarker
			}
		} else if !result.IsTruncated {
			// 使用IsTruncated判断DDC实例是否已查询完
			marker.MaxKeys = 1
			rdsResult, err := c.listRdsInstance(marker)
			if err != nil {
				return nil, err
			}
			if rdsResult != nil && len(rdsResult.Instances) > 0 {
				result.NextMarker = rdsResult.Instances[0].InstanceId
				result.IsTruncated = true
			}
		}
		return result, err
	}
	// marker 到达rds时，直接取rds列表
	result, err = c.listRdsInstance(marker)
	return result, err
}

// GetDetail - get a specific ddc Instance's detail
//
// PARAMS:
//     - instanceId: the specific ddc Instance's ID
// RETURNS:
//     - *Instance: the specific ddc Instance's detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetDetail(instanceId string) (*Instance, error) {
	var result *Instance
	var err error
	if isDDCId(instanceId) {
		result, err = c.ddcClient.GetDetail(instanceId)
	} else {
		rdsRes, err1 := c.rdsClient.GetDetail(instanceId)
		if err1 != nil {
			return nil, err
		}
		// copy返回结果
		if rdsRes != nil {
			result = &Instance{}
			err = ddc_util.SimpleCopyProperties(result, rdsRes)
			if err != nil {
				return nil, err
			}
			err = convertRdsInstance(result)
		}
	}

	return result, err
}

// 修改RDS中参数类型无法匹配的字段
func convertRdsInstance(result *Instance) error {
	// rds公网访问
	// Closed	未开放外网权限
	// Creating	公网开通中，成功后状态为Available
	// Available	已开通公网
	result.PubliclyAccessible = result.PublicAccessStatus == "Available"
	days := result.BackupPolicy.ExpireInDays
	var err error
	if len(days) > 0 {
		result.BackupPolicy.ExpireInDaysInt, err = strconv.Atoi(days)
	}
	return err
}

// DeleteRds - delete instances
//
// PARAMS:
//    - instanceIds: id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRds(instanceIds string) error {
	var err error
	ddcIds, rdsIds := getDDCAndRDSIds(instanceIds)
	if len(ddcIds) > 0 {
		err = c.ddcClient.DeleteRds(ddcIds)
	}
	if err != nil {
		return err
	}
	if len(rdsIds) > 0 {
		err = c.rdsClient.DeleteRds(rdsIds)
	}
	return err
}

// ResizeRds - resize an RDS with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//     - args: the arguments to resize an RDS
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeRds(instanceId string, args *ResizeRdsArgs) error {
	if isDDCId(instanceId) {
		return c.ddcClient.ResizeRds(instanceId, args)
	}
	// copy请求参数
	resReq := &rds.ResizeRdsArgs{}
	err := ddc_util.SimpleCopyProperties(resReq, args)
	if err != nil {
		return err
	}
	err = c.rdsClient.ResizeRds(instanceId, resReq)
	return err
}

// CreateAccount - create a account with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//     - args: the arguments to create a account
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateAccount(instanceId string, args *CreateAccountArgs) error {
	var err error
	if isDDCId(instanceId) {
		err = c.ddcClient.CreateAccount(instanceId, args)
	} else {
		rdsArgs := &rds.CreateAccountArgs{}
		// copy请求参数
		err = ddc_util.SimpleCopyProperties(rdsArgs, args)
		if err != nil {
			return err
		}
		err = c.rdsClient.CreateAccount(instanceId, rdsArgs)
	}

	return err
}

// ListAccount - list all account of a RDS instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *ListAccountResult: the result of list all account, contains all accounts' meta
//     - error: nil if success otherwise the specific error
func (c *Client) ListAccount(instanceId string) (*ListAccountResult, error) {
	var result *ListAccountResult
	var err error
	if isDDCId(instanceId) {
		result, err = c.ddcClient.ListAccount(instanceId)
	} else {
		var rdsRes *rds.ListAccountResult
		rdsRes, err = c.rdsClient.ListAccount(instanceId)
		// copy返回结果
		if rdsRes != nil {
			result = &ListAccountResult{}
			err = ddc_util.SimpleCopyProperties(result, rdsRes)
		}
	}

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
func (c *Client) GetAccount(instanceId, accountName string) (*Account, error) {
	var result *Account
	var err error
	if isDDCId(instanceId) {
		result, err = c.ddcClient.GetAccount(instanceId, accountName)
	} else {
		var rdsRes *rds.Account
		rdsRes, err = c.rdsClient.GetAccount(instanceId, accountName)
		// copy返回结果
		if rdsRes != nil {
			result = &Account{}
			err = ddc_util.SimpleCopyProperties(result, rdsRes)
		}
	}

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
	var err error
	if isDDCId(instanceId) {
		err = c.ddcClient.DeleteAccount(instanceId, accountName)
	} else {
		err = c.rdsClient.DeleteAccount(instanceId, accountName)
	}
	return err
}

// RebootInstance - reboot a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be rebooted
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebootInstance(instanceId string) error {
	if isDDCId(instanceId) {
		args := &RebootArgs{IsRebootNow: true}
		return c.ddcClient.RebootInstanceWithArgs(instanceId, args)
	}
	err := c.rdsClient.RebootInstance(instanceId)
	return err
}

// RebootInstance - reboot a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be rebooted
//     - args: reboot args
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebootInstanceWithArgs(instanceId string, args *RebootArgs) error {
	if isDDCId(instanceId) {
		return c.ddcClient.RebootInstanceWithArgs(instanceId, args)
	}
	return RDSNotSupportError()
}

// UpdateInstanceName - update name of a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//     - args: the arguments to update instanceName
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(instanceId string, args *UpdateInstanceNameArgs) error {
	var err error
	if isDDCId(instanceId) {
		err = c.ddcClient.UpdateInstanceName(instanceId, args)
	} else {
		rdsArgs := &rds.UpdateInstanceNameArgs{
			InstanceName: args.InstanceName,
		}
		err = c.rdsClient.UpdateInstanceName(instanceId, rdsArgs)
	}
	return err
}

// UpdateSyncMode - update sync mode of a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//     - args: the arguments to update syncMode
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifySyncMode(instanceId string, args *ModifySyncModeArgs) error {
	if isDDCId(instanceId) {
		return DDCNotSupportError()
	}
	rdsArgs := &rds.ModifySyncModeArgs{SyncMode: args.SyncMode}
	err := c.rdsClient.ModifySyncMode(instanceId, rdsArgs)
	return err
}

// ModifyEndpoint - modify the prefix of endpoint
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//     - args: the arguments to modify endpoint
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyEndpoint(instanceId string, args *ModifyEndpointArgs) error {
	if isDDCId(instanceId) {
		return DDCNotSupportError()
	}
	rdsArgs := &rds.ModifyEndpointArgs{Address: args.Address}
	err := c.rdsClient.ModifyEndpoint(instanceId, rdsArgs)
	return err
}

// ModifyPublicAccess - modify public access
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//     - args: the arguments to modify public access
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyPublicAccess(instanceId string, args *ModifyPublicAccessArgs) error {
	if isDDCId(instanceId) {
		return DDCNotSupportError()
	}
	rdsArgs := &rds.ModifyPublicAccessArgs{PublicAccess: args.PublicAccess}
	err := c.rdsClient.ModifyPublicAccess(instanceId, rdsArgs)
	return err
}

// GetBackupList - get backup list of the instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
// RETURNS:
//     - *GetBackupListResult: result of the backup list
//     - error: nil if success otherwise the specific error
func (c *Client) GetBackupList(instanceId string, args *GetBackupListArgs) (*GetBackupListResult, error) {

	var result *GetBackupListResult
	var err error
	if isDDCId(instanceId) {
		result, err = c.ddcClient.GetBackupList(instanceId, args)
	} else {
		rdsArgs := &rds.GetBackupListArgs{}
		// copy请求参数
		err = ddc_util.SimpleCopyProperties(rdsArgs, args)
		if err != nil {
			return nil, err
		}
		var rdsRes *rds.GetBackupListResult
		rdsRes, err = c.rdsClient.GetBackupList(instanceId, rdsArgs)
		// 转换返回结果
		if rdsRes != nil {
			result = &GetBackupListResult{}
			err = ddc_util.SimpleCopyProperties(result, rdsRes)
		}
	}

	return result, err
}

// GetZoneList - list all zone
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *GetZoneListResult: result of the zone list
//     - error: nil if success otherwise the specific error
func (c *Client) GetZoneList(productType string) (*GetZoneListResult, error) {
	if isDDC(productType) {
		return c.ddcClient.GetZoneList()
	}
	rdsRes, err := c.rdsClient.GetZoneList()
	if err != nil {
		return nil, err
	}
	result := &GetZoneListResult{}
	err = ddc_util.SimpleCopyProperties(result, rdsRes)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ListsSubnet - list all Subnets
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list all subnets, not necessary
// RETURNS:
//     - *ListSubnetsResult: result of the subnet list
//     - error: nil if success otherwise the specific error
func (c *Client) ListSubnets(args *ListSubnetsArgs, productType string) (*ListSubnetsResult, error) {
	if isDDC(productType) {
		return c.ddcClient.ListSubnets(args)
	}
	rdsArgs := &rds.ListSubnetsArgs{}
	err := ddc_util.SimpleCopyProperties(rdsArgs, args)
	if err != nil {
		return nil, err
	}
	rdsRes, err := c.rdsClient.ListSubnets(rdsArgs)
	if err != nil {
		return nil, err
	}
	result := &ListSubnetsResult{}
	err = ddc_util.SimpleCopyProperties(result, rdsRes)
	return result, err
}

// GetSecurityIps - get all SecurityIps
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *GetSecurityIpsResult: all security IP
//     - error: nil if success otherwise the specific error
func (c *Client) GetSecurityIps(instanceId string) (*GetSecurityIpsResult, error) {
	if isDDCId(instanceId) {
		return c.ddcClient.GetSecurityIps(instanceId)
	}
	rdsRes, err := c.rdsClient.GetSecurityIps(instanceId)
	if err != nil {
		return nil, err
	}
	result := &GetSecurityIpsResult{}
	err = ddc_util.SimpleCopyProperties(result, rdsRes)
	return result, err
}

// UpdateSecurityIps - update SecurityIps
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - Etag: get latest etag by GetSecurityIps
//     - Args: all SecurityIps
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateSecurityIps(instanceId, Etag string, args *UpdateSecurityIpsArgs) error {
	if isDDCId(instanceId) {
		return c.ddcClient.UpdateSecurityIps(instanceId, args)
	}
	rdsArgs := &rds.UpdateSecurityIpsArgs{SecurityIps: args.SecurityIps}
	err := c.rdsClient.UpdateSecurityIps(instanceId, Etag, rdsArgs)
	return err
}

// ListParameters - list all parameters of a RDS instance
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *ListParametersResult: the result of list all parameters
//     - error: nil if success otherwise the specific error
func (c *Client) ListParameters(instanceId string) (*ListParametersResult, error) {
	if isDDCId(instanceId) {
		return c.ddcClient.ListParameters(instanceId)
	}
	rdsRes, err := c.rdsClient.ListParameters(instanceId)
	result := &ListParametersResult{}
	err1 := ddc_util.SimpleCopyProperties(result, rdsRes)
	if err1 != nil {
		return nil, err1
	}
	// 兼容rds处理
	if result.Parameters != nil && len(result.Parameters) > 0 {
		for i, parameter := range result.Parameters {
			result.Parameters[i].IsDynamic = parameter.Dynamic == "true"
			result.Parameters[i].ISModifiable = parameter.Modifiable == "true"
		}
	}
	return result, err
}

// UpdateParameter - update Parameter
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - Etag: get latest etag by ListParameters
//     - Args: *UpdateParameterArgs
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateParameter(instanceId, Etag string, args *UpdateParameterArgs) error {
	if isDDCId(instanceId) {
		return c.ddcClient.UpdateParameter(instanceId, args)
	}
	rdsArgs := &rds.UpdateParameterArgs{}
	err1 := ddc_util.SimpleCopyProperties(rdsArgs, args)
	if err1 != nil {
		return err1
	}
	err := c.rdsClient.UpdateParameter(instanceId, Etag, rdsArgs)
	return err
}

// CreateDeploySet - create a deploy set
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateDeploySet(poolId string, args *CreateDeployRequest) (*CreateDeployResult, error) {
	return c.ddcClient.CreateDeploySet(poolId, args)
}

// UpdateDeploySet - create a deploy set
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateDeploySet(poolId, deployId string, args *UpdateDeployRequest) error {
	return c.ddcClient.UpdateDeploySet(poolId, deployId, args)
}

// ListDeploySets - list all deploy sets
// RETURNS:
//     - *ListResultWithMarker: the result of list deploy sets with marker
//     - error: nil if success otherwise the specific error
func (c *Client) ListDeploySets(poolId string, marker *Marker) (*ListDeploySetResult, error) {
	return c.ddcClient.ListDeploySets(poolId, marker)
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
func (c *Client) GetDeploySet(poolId string, deploySetId string) (*DeploySet, error) {
	return c.ddcClient.GetDeploySet(poolId, deploySetId)
}

// DeleteDeploySet - delete a deploy set
//
// PARAMS:
//     - poolId: the id of the pool
//     - deploySetId: the id of the deploy set
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteDeploySet(poolId string, deploySetId string) error {
	return c.ddcClient.DeleteDeploySet(poolId, deploySetId)
}

// CreateBackup - create backup of the instance
//
// PARAMS:
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateBackup(instanceId string) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.CreateBackup(instanceId)
}

// ModifyBackupPolicy - update backupPolicy
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
//     - Args: the specific rds Instance's BackupPolicy
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupPolicy(instanceId string, args *BackupPolicy) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.ModifyBackupPolicy(instanceId, args)
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
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.GetBackupDetail(instanceId, snapshotId)
}

// GetBinlogList - get backup list of the instance
//
// PARAMS:
//     - instanceId: id of the instance
// RETURNS:
//     - *BinlogListResult: result of the backup list
//     - error: nil if success otherwise the specific error
func (c *Client) GetBinlogList(instanceId string, datetime string) (*BinlogListResult, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.GetBinlogList(instanceId, datetime)
}

// GetBinlogDetail - get details of the instance'Binlog
//
// PARAMS:
//     - instanceId: the id of the instance
//     - binlog: the id of the binlog
// RETURNS:
//     - *BinlogDetailResult: the detail of the binlog
//     - error: nil if success otherwise the specific error
func (c *Client) GetBinlogDetail(instanceId string, binlog string) (*BinlogDetailResult, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.GetBinlogDetail(instanceId, binlog)
}

// SwitchInstance - main standby switching of the instance
//
// PARAMS:
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) SwitchInstance(instanceId string, args *SwitchArgs) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.SwitchInstance(instanceId, args)
}

// CreateDatabase - create a database with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//     - args: the arguments to create a account
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateDatabase(instanceId string, args *CreateDatabaseArgs) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.CreateDatabase(instanceId, args)
}

// DeleteDatabase - delete an database of a DDC instance
//
// PARAMS:
//     - instanceIds: the specific instanceIds
//     - dbName: the specific database's name
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteDatabase(instanceId, dbName string) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.DeleteDatabase(instanceId, dbName)
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
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.UpdateDatabaseRemark(instanceId, dbName, args)
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
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.GetDatabase(instanceId, dbName)
}

// ListDatabase - list all database of a DDC instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific ddc Instance's ID
// RETURNS:
//     - *ListDatabaseResult: the result of list all database, contains all databases' meta
//     - error: nil if success otherwise the specific error
func (c *Client) ListDatabase(instanceId string) (*ListDatabaseResult, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.ListDatabase(instanceId)
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
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.UpdateAccountPassword(instanceId, accountName, args)
}

// UpdateAccountDesc - update a account desc with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//	   - accountName: the specific accountName
//     - args: the arguments to update a account remark
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountDesc(instanceId string, accountName string, args *UpdateAccountDescArgs) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.UpdateAccountDesc(instanceId, accountName, args)
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
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.UpdateAccountPrivileges(instanceId, accountName, args)
}

// ListRoGroup - list all roGroups of a DDC instance with the specific parameters
//
// PARAMS:
//     - instanceId: the specific rds Instance's ID
// RETURNS:
//     - *ListRoGroupResult: All roGroups of the current instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListRoGroup(instanceId string) (*ListRoGroupResult, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.ListRoGroup(instanceId)
}

// ListVpc - list all Vpc
//
// PARAMS:
// RETURNS:
//     - *ListVpc: All vpc of
//     - error: nil if success otherwise the specific error
func (c *Client) ListVpc(productType string) (*[]VpcVo, error) {
	if !isDDC(productType) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.ListVpc()
}

// autoRenew - create autoRenew
//
// PARAMS:
//     - Args: *autoRenewArgs
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) AutoRenew(args *AutoRenewArgs, productType string) error {
	if isDDC(productType) {
		return DDCNotSupportError()
	}
	// 实例列表不能为空
	if len(args.InstanceIds) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	rdsArgs := &rds.AutoRenewArgs{}
	err := ddc_util.SimpleCopyProperties(rdsArgs, args)
	if err != nil {
		return err
	}
	return c.rdsClient.AutoRenew(rdsArgs)
}

// GetMaintenTime - get details of the maintenTime
//
// PARAMS:
//     - poolId: the id of the pool
//     - cli: the client agent which can perform sending request
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - *DeploySet: the detail of the deploy set
//     - error: nil if success otherwise the specific error
func (c *Client) GetMaintainTime(instanceId string) (*MaintainTime, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.GetMaintainTime(instanceId)
}

// UpdateMaintenTime - update UpdateMaintenTime of instance
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateMaintainTime(instanceId string, args *MaintainTime) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.UpdateMaintainTime(instanceId, args)
}

// ListRecycleInstances - list all instances in recycler with marker
//
// PARAMS:
//     - marker: marker page
// RETURNS:
//     - *RecyclerInstanceList: the result of instances in recycler
//     - error: nil if success otherwise the specific error
func (c *Client) ListRecycleInstances(marker *Marker, productType string) (*RecyclerInstanceList, error) {
	if !isDDC(productType) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.ListRecycleInstances(marker)
}

// RecoverRecyclerInstances - batch recover instances that in recycler
//
// PARAMS:
//     - instanceIds: instanceId list to recover
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RecoverRecyclerInstances(instanceIds []string) error {
	if instanceIds == nil || len(instanceIds) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	if len(instanceIds) > 10 {
		return fmt.Errorf("the instanceIds length max value is 10")
	}

	for _, id := range instanceIds {
		if !isDDCId(id) {
			return RDSNotSupportError()
		}
	}
	return c.ddcClient.RecoverRecyclerInstances(instanceIds)
}

// DeleteRecyclerInstances - batch delete instances that in recycler
//
// PARAMS:
//     - instanceIds: instanceId list to delete
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRecyclerInstances(instanceIds []string) error {
	if instanceIds == nil || len(instanceIds) < 1 {
		return fmt.Errorf("unset instanceIds")
	}
	if len(instanceIds) > 10 {
		return fmt.Errorf("the instanceIds length max value is 10")
	}

	for _, id := range instanceIds {
		if !isDDCId(id) {
			return RDSNotSupportError()
		}
	}
	return c.ddcClient.DeleteRecyclerInstances(instanceIds)
}

// ListSecurityGroupByVpcId - list security groups by vpc id
//
// PARAMS:
//     - vpcId: id of vpc
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroupByVpcId(vpcId string) (*[]SecurityGroup, error) {
	return c.ddcClient.ListSecurityGroupByVpcId(vpcId)
}

// ListSecurityGroupByInstanceId - list security groups by instance id
//
// PARAMS:
//     - instanceId: id of instance
// RETURNS:
//     - *ListSecurityGroupResult: list secrity groups result of instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroupByInstanceId(instanceId string) (*ListSecurityGroupResult, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.ListSecurityGroupByInstanceId(instanceId)
}

// BindSecurityGroups - bind SecurityGroup to instances
//
// PARAMS:
//     - args: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindSecurityGroups(args *SecurityGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	for _, id := range args.InstanceIds {
		if !isDDCId(id) {
			return RDSNotSupportError()
		}
	}
	return c.ddcClient.BindSecurityGroups(args)
}

// UnBindSecurityGroups - unbind SecurityGroup to instances
//
// PARAMS:
//     - args: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindSecurityGroups(args *SecurityGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	for _, id := range args.InstanceIds {
		if !isDDCId(id) {
			return RDSNotSupportError()
		}
	}
	return c.ddcClient.UnBindSecurityGroups(args)
}

// ReplaceSecurityGroups - replace SecurityGroup to instances
//
// PARAMS:
//     - args: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ReplaceSecurityGroups(args *SecurityGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	for _, id := range args.InstanceIds {
		if !isDDCId(id) {
			return RDSNotSupportError()
		}
	}
	return c.ddcClient.ReplaceSecurityGroups(args)
}

// ListLogByInstanceId - list error or slow logs of instance
//
// PARAMS:
//     - instanceId: id of instance
// RETURNS:
//     - *[]Log:logs of instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListLogByInstanceId(instanceId string, args *ListLogArgs) (*[]Log, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.ListLogByInstanceId(instanceId, args)
}

// GetLogById - list log's detail of instance
//
// PARAMS:
//     - instanceId: id of instance
// RETURNS:
//     - *Log:log's detail of instance
//     - error: nil if success otherwise the specific error
func (c *Client) GetLogById(instanceId, logId string, args *GetLogArgs) (*LogDetail, error) {
	if !isDDCId(instanceId) {
		return nil, RDSNotSupportError()
	}
	return c.ddcClient.GetLogById(instanceId, logId, args)
}

// LazyDropCreateHardLink - create a hard link for specified large table
//
// PARAMS:
//     - instanceId: id of instance
//     - dbName: name of database
//     - tableName: name of table
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) LazyDropCreateHardLink(instanceId, dbName, tableName string) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.LazyDropCreateHardLink(instanceId, dbName, tableName)
}

// LazyDropDeleteHardLink - delete the hard link for specified large table
//
// PARAMS:
//     - instanceId: id of instance
//     - dbName: name of database
//     - tableName: name of table
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) LazyDropDeleteHardLink(instanceId, dbName, tableName string) error {
	if !isDDCId(instanceId) {
		return RDSNotSupportError()
	}
	return c.ddcClient.LazyDropDeleteHardLink(instanceId, dbName, tableName)
}
