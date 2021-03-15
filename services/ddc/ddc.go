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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
	"strings"
)

const (
	KEY_CLIENT_TOKEN = "clientToken"
	KEY_MARKER       = "marker"
	KEY_MAX_KEYS     = "maxKeys"
	COMMA            = ","
)

// Convert marker to request params
func getMarkerParams(marker *Marker) map[string]string {
	if marker == nil {
		marker = &Marker{Marker: "-1"}
	}
	params := make(map[string]string, 2)
	params[KEY_MARKER] = marker.Marker
	if marker.MaxKeys > 0 {
		params[KEY_MAX_KEYS] = strconv.Itoa(marker.MaxKeys)
	}
	return params
}

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

// CreateRds - create a DDC with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a ddc
// RETURNS:
//     - *InstanceIds: the result of create DDC, contains new DDC's instanceIds
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

	newArgs := &CreateInstanceArgs{
		InstanceType: "RDS",
		Number:       args.PurchaseCount,
		ClientToken:  args.ClientToken,
		Instance: CreateInstance{
			InstanceName:         args.InstanceName,
			Engine:               args.Engine,
			EngineVersion:        args.EngineVersion,
			CpuCount:             args.CpuCount,
			AllocatedMemoryInGB:  args.MemoryCapacity,
			AllocatedStorageInGB: args.VolumeCapacity,
			DiskIoType:           "ssd",
			DeployId:             args.DeployId,
			PoolId:               args.PoolId,
			IsDirectPay:          args.IsDirectPay,
			Billing:              args.Billing,
			AutoRenewTime:        args.AutoRenewTime,
			AutoRenewTimeUnit:    args.AutoRenewTimeUnit,
			Tags:                 args.Tags,
			Category:             args.Category,
		},
	}

	info, err := c.SupplyVpcInfo(newArgs, args)
	if err != nil {
		return nil, err
	}
	result, err2 := c.CreateInstance(info)
	if err2 != nil {
		return nil, err2
	}
	//list, err4 := c.GetZoneList()
	//if err4 != nil {
	//	fmt.Printf("get zone list error: %+v\n", err4)
	//	return nil, err4
	//}
	//if args.ZoneNames != nil {
	//	newZoneName := ""
	//	for _, e1 := range list.Zones {
	//		if strings.Join(args.ZoneNames, ",") == strings.Join(e1.ZoneNames, ",") {
	//			newZoneName = strings.Join(e1.ZoneNames, ",")
	//		}
	//	}
	//	newArgs.Instance.AZone = newZoneName
	//}
	return result, nil
}

func (c *Client) SupplyVpcInfo(newArgs *CreateInstanceArgs, args *CreateRdsArgs) (*CreateInstanceArgs, error) {

	info := newArgs
	vpc, err := c.ListVpc()
	if err != nil {
		fmt.Printf("list vpc error: %+v\n", err)
		return nil, err
	}
	defaultVpcId := ""
	if args.VpcId == "" {
		for _, e := range *vpc {
			defaultVpcId = e.VpcId
			info.Instance.VpcId = e.VpcId
			args.VpcId = e.VpcId
		}
	}
	if args.VpcId == defaultVpcId {
		for _, e := range *vpc {
			if e.VpcId == args.VpcId {
				info.Instance.VpcId = e.VpcId
				args.VpcId = e.VpcId
			}
		}
		info, err = c.UnDefaultVpcInfo(info, args)
		if err != nil {
			fmt.Printf("set vpc error: %+v\n", err)
			return nil, err
		}
	} else {
		for _, e := range *vpc {
			if args.VpcId == e.ShortId {
				info.Instance.VpcId = e.VpcId
				args.VpcId = e.VpcId
			}
		}
		info, err = c.UnDefaultVpcInfo(info, args)
		if err != nil {
			fmt.Printf("supply zoneAndSubnet info error: %+v\n", err)
			return nil, err
		}
	}
	return info, nil
}

func (c *Client) SupplyZoneAndSubnetInfo(newArgs *CreateInstanceArgs, args *CreateRdsArgs) (*CreateInstanceArgs, error) {
	newZoneName := ""
	if args.Subnets != nil {
		for _, e := range args.Subnets {
			newZoneName += e.ZoneName + ","
		}
		if newZoneName != "" && newZoneName[0:len(newZoneName)-1] != strings.Join(args.ZoneNames, ",") {
			fmt.Printf("subnets and zoneNames not matcher: %+v\n", nil)
			return nil, errors.New("subnets and zoneNames not matcher")
		} else {
			subnets, err1 := c.ListSubnets(nil)
			if err1 != nil {
				fmt.Printf("list subnets error: %+v\n", err1)
				return nil, err1
			}
			subnetId := ""
			if args.Subnets != nil {
				for _, e := range subnets.Subnets {
					for _, e1 := range args.Subnets {
						if e.ShortId == e1.SubnetId {
							subnetId += e.Az + ":" + e.SubnetId + ","
						}
					}
				}
				if subnetId == "" {
					return nil, errors.New("subnetId no match")
				}
				newArgs.Instance.SubnetId = subnetId[0 : len(subnetId)-1]
			}
		}
	} else {
		var subnetStr string
		for _, e := range args.ZoneNames {
			if !strings.Contains(e, ",") {
				listSubnetsArgs := &ListSubnetsArgs{
					VpcId:    args.VpcId,
					ZoneName: e,
				}
				subnets, err := c.ListSubnets(listSubnetsArgs)
				if err != nil {
					fmt.Printf("list subnets error: %+v\n", err)
					return nil, err
				}
				if subnets != nil && len(subnets.Subnets) > 0 {
					subnetId := subnets.Subnets[0].SubnetId
					subnetStr += e + ":" + subnetId + ","
				}
			}
		}
		if len(subnetStr) < 1 {
			return nil, errors.New("Have no available subnet")
		}
		newArgs.Instance.SubnetId = subnetStr[:len(subnetStr)-1]
	}
	return newArgs, nil
}

func (c *Client) UnDefaultVpcInfo(newArgs *CreateInstanceArgs, args *CreateRdsArgs) (*CreateInstanceArgs, error) {
	info := newArgs
	list, err2 := c.GetZoneList()
	if err2 != nil {
		fmt.Printf("get zone list error: %+v\n", err2)
		return nil, err2
	}
	newZoneName := ""
	if args.ZoneNames == nil {
		if args.Subnets == nil {
			subnets, _ := c.ListSubnets(nil)
			if subnets == nil || len(subnets.Subnets) == 0 {
				return nil, errors.New("Have no available subnet or zone")
			}

			for _, e := range subnets.Subnets {
				info.Instance.AZone = e.Az
				args.ZoneNames = append(args.ZoneNames, e.Az)
				break
			}
			//for _, e := range list.Zones {
			//	info.Instance.AZone = e.ApiZoneNames[0]
			//	args.ZoneNames = append(args.ZoneNames, e.ApiZoneNames[0])
			//	break
			//}
		} else {
			for _, e := range args.Subnets {
				newZoneName += e.ZoneName + ","
				args.ZoneNames = append(args.ZoneNames, e.ZoneName)
			}
			for _, e := range list.Zones {
				if newZoneName == strings.Join(e.ApiZoneNames, ",") {
					info.Instance.AZone = strings.Join(e.ApiZoneNames, ",")
				}
			}
		}
		if args.ZoneNames == nil {
			return nil, errors.New("Have no available zone for your operation.")
		}
	} else {
		newZoneName = ""
		for _, e1 := range list.Zones {
			if strings.Join(args.ZoneNames, ",") == strings.Join(e1.ZoneNames, ",") {
				newZoneName = strings.Join(e1.ApiZoneNames, ",")
			}
		}
		info.Instance.AZone = newZoneName
	}

	info, err2 = c.SupplyZoneAndSubnetInfo(info, args)
	if err2 != nil {
		return nil, err2
	}
	return info, nil
}

// CreateReadReplica - create a readReplica ddc with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a readReplica ddc
// RETURNS:
//     - *InstanceIds: the result of create a readReplica ddc, contains the readReplica DDC's instanceIds
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
	detail, err2 := c.GetDdcDetail(args.SourceInstanceId)
	if err2 != nil {
		return nil, err2
	}
	newArgs := &CreateInstanceArgs{
		InstanceType: "RDS",
		Number:       args.PurchaseCount,
		ClientToken:  args.ClientToken,
		Instance: CreateInstance{
			InstanceName:         args.InstanceName,
			Engine:               detail.Instance.Engine,
			EngineVersion:        detail.Instance.EngineVersion,
			CpuCount:             args.CpuCount,
			AllocatedMemoryInGB:  args.MemoryCapacity,
			AllocatedStorageInGB: args.VolumeCapacity,
			DiskIoType:           "ssd",
			DeployId:             args.DeployId,
			PoolId:               args.PoolId,
			RoGroupId:            args.RoGroupId,
			RoGroupWeight:        args.RoGroupWeight,
			EnableDelayOff:       args.EnableDelayOff,
			DelayThreshold:       args.DelayThreshold,
			LeastInstanceAmount:  args.LeastInstanceAmount,
			Billing:              args.Billing,
			IsDirectPay:          args.IsDirectPay,
			Tags:                 args.Tags,
		},
	}

	createRdsArgs := &CreateRdsArgs{
		VpcId:     args.VpcId,
		Subnets:   args.Subnets,
		ZoneNames: args.ZoneNames,
	}

	info, err := c.SupplyVpcInfo(newArgs, createRdsArgs)
	if err != nil {
		return nil, err
	}
	result, err2 := c.CreateInstance(info)
	if err2 != nil {
		return nil, err2
	}

	return result, err
}

// UpdateRoGroup - update a roGroup
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateRoGroup(roGroupId string, args *UpdateRoGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRoGroupUriWithId(roGroupId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// UpdateRoGroupReplicaWeight- update repica weight in roGroup
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateRoGroupReplicaWeight(roGroupId string, args *UpdateRoGroupWeightArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRoGroupWeightUriWithId(roGroupId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// ReBalanceRoGroup- Initiate a rebalance for foGroup
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ReBalanceRoGroup(roGroupId string) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getReBalanceRoGroupUriWithId(roGroupId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
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

// UpdateDeploySet - update a deploy set
//
// PARAMS:
//     - body: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateDeploySet(poolId string, deployId string, args *UpdateDeployRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if !(args.Strategy == "distributed" || args.Strategy == "centralized") {
		return fmt.Errorf("Only support strategy distributed/centralized, current strategy: %v", args.Strategy)
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeploySetUriWithId(poolId, deployId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// ListRds - list all instances
// RETURNS:
//     - *ListRdsResult: the result of list instances with marker
//     - error: nil if success otherwise the specific error
func (c *Client) ListRds(marker *ListRdsArgs) (*ListRdsResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDdcInstanceUri() + "/list")
	req.SetMethod(http.GET)
	if marker != nil {
		req.SetParam(KEY_MARKER, marker.Marker)
		req.SetParam(KEY_MAX_KEYS, strconv.Itoa(marker.MaxKeys))
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListRdsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetDdcDetail - get details of the instance
//
// PARAMS:
//     - instanceId: the id of the instance
// RETURNS:
//     - *InstanceModelResult: the detail of the instance
//     - error: nil if success otherwise the specific error
func (c *Client) GetDdcDetail(instanceId string) (*InstanceModelResult, error) {
	result := &InstanceModelResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId)).
		WithResult(result).
		Do()

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
	detail, err := c.GetDdcDetail(instanceId)
	instanceJson, _ := json.Marshal(detail.Instance)
	args := &Instance{
		InstanceName:       detail.Instance.InstanceName,
		InstanceId:         detail.Instance.InstanceId,
		SourceInstanceId:   detail.Instance.SourceInstanceId,
		Endpoint:           detail.Instance.Endpoint,
		Engine:             detail.Instance.Engine,
		EngineVersion:      detail.Instance.EngineVersion,
		InstanceStatus:     detail.Instance.InstanceStatus,
		CpuCount:           detail.Instance.CpuCount,
		MemoryCapacity:     detail.Instance.AllocatedMemoryInGB,
		VolumeCapacity:     detail.Instance.AllocatedStorageInGB,
		UsedStorage:        detail.Instance.UsedStorageInGB,
		InstanceType:       detail.Instance.Type,
		InstanceCreateTime: detail.Instance.InstanceCreateTime,
		InstanceExpireTime: detail.Instance.InstanceExpireTime,
		PublicAccessStatus: detail.Instance.PublicAccessStatus,
		PaymentTiming:      detail.Instance.PaymentTiming,
		SyncMode:           detail.Instance.SyncMode,
		Region:             detail.Instance.Region,
		VpcId:              detail.Instance.VpcId,
		BackupPolicy:       detail.Instance.BackupPolicy,
		RoGroupList:        detail.Instance.RoGroupList,
		NodeMaster:         detail.Instance.NodeMaster,
		NodeSlave:          detail.Instance.NodeSlave,
		NodeReadReplica:    detail.Instance.NodeReadReplica,
		Subnets:            detail.Instance.Subnets,
		DeployId:           detail.Instance.DeployId,
		ZoneNames:          detail.Instance.ZoneNames,
	}
	err = json.Unmarshal(instanceJson, &args)
	return args, err
}

// DeleteRds - delete instances
//
// PARAMS:
//    - instanceIds: id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRds(instanceIds string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDdcInstanceUri()+"/delete").
		WithQueryParam("instanceIds", instanceIds).
		Do()
}

// RebootInstance - reboot a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be rebooted
//     - args: reboot args
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebootInstance(instanceId string, args *RebootArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instanceId)).
		WithQueryParam("reboot", "").
		WithBody(args).
		Do()
}

// UpdateInstanceName - update name of a specified instance
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
func (c *Client) ListSubnets(args *ListSubnetsArgs) (*ListSubnetsResult, error) {
	result := &ListSubnetsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(URI_PREFIX + "/subnet").
		WithResult(result).
		Do()
	if args == nil {
		return result, err
	}
	if args.ZoneName == "" {
		return result, err
	}
	// to compat rds api, filter by zone and vpcId
	if result.Subnets != nil && len(result.Subnets) > 0 {
		var filterd = []Subnet{}
		for _, subnet := range result.Subnets {
			// subnet az is logical zone
			if subnet.Az == args.ZoneName {
				if args.VpcId == "" || args.VpcId == subnet.VpcId {
					filterd = append(filterd, subnet)
				}
			}
		}
		result.Subnets = filterd
	}
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
		req.SetParam(KEY_MAX_KEYS, strconv.Itoa(marker.MaxKeys))
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
	rowResult := &SecurityIpsRawResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/authIp").
		WithResult(rowResult).
		Do()
	// to compat rds api,json annotations for SecurityIps are different
	result := &GetSecurityIpsResult{
		SecurityIps: rowResult.SecurityIps,
	}
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

// GetBinlogList - get backup list of the instance
//
// PARAMS:
//     - instanceId: id of the instance
// RETURNS:
//     - *BinlogListResult: result of the backup list
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

// GetBinlogDetail - get details of the instance'Binlog
//
// PARAMS:
//     - instanceId: the id of the instance
//     - binlog: the id of the binlog
// RETURNS:
//     - *BinlogDetailResult: the detail of the binlog
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
//     - args: switch now or wait to the maintain time
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) SwitchInstance(instanceId string, args *SwitchArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	result := &bce.BceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instanceId)+"/switchMaster").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
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

	if args.AccountType == "" {
		args.AccountType = "common"
	}
	if args.AccountType == "Super" {
		args.AccountType = "rdssuper"
	}
	if args.AccountType == "Common" {
		args.AccountType = "common"
	}
	if args.DatabasePrivileges != nil {
		for idx, _ := range args.DatabasePrivileges {
			if args.DatabasePrivileges[idx].AuthType == "ReadOnly" {
				args.DatabasePrivileges[idx].AuthType = "readOnly"
			} else if args.DatabasePrivileges[idx].AuthType == "ReadWrite" {
				args.DatabasePrivileges[idx].AuthType = "readWrite"
			}
		}
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

// UpdateAccountDesc - update a account desc with the specific parameters
//
// PARAMS:
//     - instanceId: the specific instanceId
//	   - accountName: the specific accountName
//     - args: the arguments to update a account remark
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountDesc(instanceId string, accountName string, args *UpdateAccountDescArgs) error {
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

	for idx, _ := range args.DatabasePrivileges {
		if args.DatabasePrivileges[idx].AuthType == "ReadOnly" {
			args.DatabasePrivileges[idx].AuthType = "readOnly"
		} else if args.DatabasePrivileges[idx].AuthType == "ReadWrite" {
			args.DatabasePrivileges[idx].AuthType = "readWrite"
		}
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

	if result.Account.AccountType == "common" {
		result.Account.AccountType = "Common"
	} else if result.Account.AccountType == "rdssuper" {
		result.Account.AccountType = "Super"
	}

	for idx, _ := range result.Account.DatabasePrivileges {
		if result.Account.DatabasePrivileges[idx].AuthType == "readOnly" {
			result.Account.DatabasePrivileges[idx].AuthType = "ReadOnly"
		} else if result.Account.DatabasePrivileges[idx].AuthType == "readWrite" {
			result.Account.DatabasePrivileges[idx].AuthType = "ReadWrite"
		}
	}
	result.Account.Status = strings.Title(result.Account.Status)
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
	for idx, _ := range result.Accounts {
		if result.Accounts[idx].AccountType == "common" {
			result.Accounts[idx].AccountType = "Common"
		} else if result.Accounts[idx].AccountType == "rdssuper" {
			result.Accounts[idx].AccountType = "Super"
		}
		result.Accounts[idx].Status = strings.Title(result.Accounts[idx].Status)

		for iidx, _ := range result.Accounts[idx].DatabasePrivileges {
			if result.Accounts[idx].DatabasePrivileges[iidx].AuthType == "readOnly" {
				result.Accounts[idx].DatabasePrivileges[iidx].AuthType = "ReadOnly"
			} else if result.Accounts[idx].DatabasePrivileges[iidx].AuthType == "readWrite" {
				result.Accounts[idx].DatabasePrivileges[iidx].AuthType = "ReadWrite"
			}
		}
	}
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
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getMaintainTimeUriWithInstanceId(instanceId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &MaintainWindow{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return &jsonBody.MaintainTime, nil
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

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateMaintainTimeUriWithInstanceId(instanceId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// ListRecycleInstances - list all instances in recycler with marker
//
// PARAMS:
//     - marker: marker page
// RETURNS:
//     - *RecyclerInstanceList: the result of instances in recycler
//     - error: nil if success otherwise the specific error
func (c *Client) ListRecycleInstances(marker *Marker) (*RecyclerInstanceList, error) {
	result := &RecyclerInstanceList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithQueryParams(getMarkerParams(marker)).
		WithURL(getRecyclerUrl()).
		WithResult(result).
		Do()

	return result, err
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

	args := &BatchInstanceIds{
		InstanceIds: strings.Join(instanceIds, COMMA),
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRecyclerRecoverUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
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

	// delete use query params
	instanceIdsParam := strings.Join(instanceIds, COMMA)
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRecyclerDeleteUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParam("instanceIds", instanceIdsParam).
		Do()
	return err
}

// ListSecurityGroupByVpcId - list security groups by vpc id
//
// PARAMS:
//     - vpcId: id of vpc
// RETURNS:
//     - *[]SecurityGroup:security groups of vpc
//     - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroupByVpcId(vpcId string) (*[]SecurityGroup, error) {
	if len(vpcId) < 1 {
		return nil, fmt.Errorf("unset vpcId")
	}
	result := &[]SecurityGroup{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSecurityGroupWithVpcIdUrl(vpcId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// ListSecurityGroupByInstanceId - list security groups by instance id
//
// PARAMS:
//     - instanceId: id of instance
// RETURNS:
//     - *ListSecurityGroupResult: list secrity groups result of instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroupByInstanceId(instanceId string) (*ListSecurityGroupResult, error) {
	if len(instanceId) < 1 {
		return nil, fmt.Errorf("unset instanceId")
	}
	result := &ListSecurityGroupResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSecurityGroupWithInstanceIdUrl(instanceId)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
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

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindSecurityGroupWithUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
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

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnBindSecurityGroupWithUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
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

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getReplaceSecurityGroupWithUrl()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

// ListLogByInstanceId - list error or slow logs of instance
//
// PARAMS:
//     - instanceId: id of instance
// RETURNS:
//     - *[]Log:logs of instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListLogByInstanceId(instanceId string, args *ListLogArgs) (*[]Log, error) {
	if len(instanceId) < 1 {
		return nil, fmt.Errorf("unset instanceId")
	}
	if args == nil {
		return nil, fmt.Errorf("unset list args")
	}
	if "error" != args.LogType && "slow" != args.LogType {
		return nil, fmt.Errorf("invalid logType, should be 'error' or 'slow'")
	}
	result := &[]Log{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getLogsUrlWithInstanceId(instanceId)).
		WithQueryParam("logType", strings.ToLower(args.LogType)).
		WithQueryParam("datetime", args.Datetime).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// GetLogById - get log's detail of instance
//
// PARAMS:
//     - instanceId: id of instance
// RETURNS:
//     - *Log:log's detail of instance
//     - error: nil if success otherwise the specific error
func (c *Client) GetLogById(instanceId, logId string, args *GetLogArgs) (*LogDetail, error) {
	if len(instanceId) < 1 {
		return nil, fmt.Errorf("unset instanceId")
	}
	if len(logId) < 1 {
		return nil, fmt.Errorf("unset logId")
	}
	if args == nil {
		return nil, fmt.Errorf("unset get log args")
	}

	result := &LogDetail{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getLogsUrlWithLogId(instanceId, logId)).
		WithQueryParam("downloadValidTimeInSec", strconv.Itoa(args.ValidSeconds)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
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
	if len(instanceId) < 1 {
		return fmt.Errorf("unset instanceId")
	}
	if len(dbName) < 1 {
		return fmt.Errorf("unset dbName")
	}
	if len(tableName) < 1 {
		return fmt.Errorf("unset tableName")
	}

	args := &CreateTableHardLinkArgs{
		TableName: tableName,
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateTableHardLinkUrl(instanceId, dbName)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
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
	if len(instanceId) < 1 {
		return fmt.Errorf("unset instanceId")
	}
	if len(dbName) < 1 {
		return fmt.Errorf("unset dbName")
	}
	if len(tableName) < 1 {
		return fmt.Errorf("unset tableName")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getTableHardLinkUrl(instanceId, dbName, tableName)).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
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
	if args == nil {
		return fmt.Errorf("unset args")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDdcUriWithInstanceId(instanceId) + "/resize").
		WithBody(args).
		Do()
}
