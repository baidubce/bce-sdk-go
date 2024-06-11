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

// gaiadb.go - the gaiadb APIs definition supported by the GAIADB service
package gaiadb

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

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
func getGaiadbUri() string {
	return "/v1/gaiadb"
}
func Json(v interface{}) string {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		panic("convert to json faild")
	}
	return string(jsonStr)
}

// Convert struct to request params
func getQueryParams(val interface{}) (map[string]string, error) {
	var params map[string]string
	if val != nil {
		err := json.Unmarshal([]byte(Json(val)), &params)
		if err != nil {
			return nil, err
		}
	}
	return params, nil
}

// GetAvailableSubnetList - get available subnetList
//
// PARAMS:
//   - vpcId: the vpc id which you want to query
//
// RETURNS:
//   - []AvailableSubnet: the result of get available subnetList
//   - error: nil if success otherwise the specific error
func (c *Client) GetAvailableSubnetList(vpcId string) ([]AvailableSubnet, error) {
	result := &[]AvailableSubnet{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/cluster/subnetId").
		WithQueryParamFilter("vpcId", vpcId).
		WithResult(result).
		Do()

	return *result, err
}

// CreateCluster - create gaiadb cluster
//
// PARAMS:
//   - args: the arguments to create gaiadb cluster
//
// RETURNS:
//   - *CreateResult: the result of create gaiadb cluster
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCluster(args *CreateClusterArgs) (*CreateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	result := &CreateResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/cluster").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// DeleteCluster - delete gaiadb cluster
//
// PARAMS:
//   - clusterId: the cluster id which you want to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteCluster(clusterId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// RenameCluster - rename gaiadb cluster
//
// PARAMS:
//   - clusterId: the cluster id which you want to delete
//   - *ClusterName: the new cluster name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenameCluster(clusterId string, args *ClusterName) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId).
		WithQueryParam("clusterName", "").
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// ResizeCluster - resize gaiadb cluster
//
// PARAMS:
//   - clusterId: the cluster id which you want to delete
//   - *ResizeClusterArgs: the arguments to resize gaiadb cluster
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *ResizeResult: the result of resize gaiadb cluster
func (c *Client) ResizeCluster(clusterId string, args *ResizeClusterArgs) (*OrderId, error) {
	result := &OrderId{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId).
		WithQueryParam("resize", "").
		WithBody(args).
		WithResult(result).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return result, err
}

// GetClusterList - get gaiadb cluster list
//
// PARAMS:
//   - *Markder: the arguments to get cluster list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *ClusterListResult: the result of get cluster list
func (c *Client) GetClusterList(args *Marker) (*ClusterListResult, error) {
	result := &ClusterListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/cluster").
		WithBody(args).
		WithResult(result).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return result, err
}

// GetClusterCapacity - get cluster Capacity
// PARAMS:
//   - ClusterId: cluster id
//
// RETURNS:
//   - *ClusterCapacityResult: the result of cluster detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetClusterDetail(clusterId string) (*ClusterDetailResult, error) {
	result := &ClusterDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri() + "/cluster/" + clusterId).
		WithResult(result).
		Do()

	return result, err
}

// GetClusterCapacity - get cluster capacity
// PARAMS:
//   - ClusterId: cluster id
//
// RETURNS:
//   - *ClusterCapacityResult: the result of cluster capacity
//   - error: nil if success otherwise the specific error
func (c *Client) GetClusterCapacity(clusterId string) (*ClusterCapacityResult, error) {
	result := &ClusterCapacityResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri() + "/" + clusterId + "/capacity").
		WithResult(result).
		Do()

	return result, err
}

// QueryClusterPrice - query cluster price
//
// PARAMS:
//   - args: the arguments to query cluster price
//
// RETURNS:
//   - *PriceResult: the result of query cluster price
//   - error: nil if success otherwise the specific error
func (c *Client) QueryClusterPrice(args *QueryPriceArgs) (*PriceResult, error) {
	result := &PriceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/price").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// QueryResizeClusterPrice - query resize cluster price
//
// PARAMS:
//   - args: the arguments to query resize cluster price
//
// RETURNS:
//   - *PriceResult: the result of query resize cluster price
//   - error: nil if success otherwise the specific error
func (c *Client) QueryResizeClusterPrice(args *QueryResizePriceArgs) (*PriceResult, error) {
	result := &PriceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/price/diff").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// RebootInstance - reboot instance
//
// PARAMS:
//   - args: the arguments to reboot instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RebootInstance(clusterId, instanceId string, args *RebootInstanceArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId+"/instance/"+instanceId+"/reboot").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// BindTags - bind tags
//
// PARAMS:
//   - args: the arguments to bind tags
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) BindTags(args *BindTagsArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL("/v1/tags").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// ClusterSwitch - cluster switch
//
// PARAMS:
//   - clusterId: the cluster id
//   - args: the arguments to switch cluster
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *SwitchResult: the result of switch cluster
func (c *Client) ClusterSwitch(clusterId string, args *ClusterSwitchArgs) (*ClusterSwitchResult, error) {
	result := &ClusterSwitchResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId+"/switch").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetInterfaceList - get interface list
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - *InterfaceListResult: the result of interface list
//   - error: nil if success otherwise the specific error
func (c *Client) GetInterfaceList(clusterId string) (*InterfaceListResult, error) {
	result := &InterfaceListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri() + "/" + clusterId + "/interface").
		WithResult(result).
		Do()

	return result, err
}

// UpdateDnsName - update dns name
//
// PARAMS:
//   - clusterId: the cluster id
//   - args: the arguments to update dns name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateDnsName(clusterId string, args *UpdateDnsNameArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/interface/dns-name").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// UpdateInterface - update interface
//
// PARAMS:
//   - clusterId: cluster id
//   - args: the arguments to update interface
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInterface(clusterId string, args *UpdateInterfaceArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/interface").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// NewInstanceAutoJoin - new instance auto join
//
// PARAMS:
//   - clusterId: cluster id
//   - args: the arguments to new instance auto join
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) NewInstanceAutoJoin(clusterId string, args *NewInstanceAutoJoinArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/interface/new-instance-auto-join").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// CreateAccount - create account
//
// PARAMS:
//   - clusterId: cluster id
//   - args: the arguments to create account
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAccount(clusterId string, args *CreateAccountArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// DeleteAccount - delete account
//
// PARAMS:
//   - clusterId: cluster id
//   - accountName: account name to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAccount(clusterId, accountName string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account/"+accountName).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// GetAccountDetail - get account detail
//
// PARAMS:
//   - clusterId: cluster id
//   - accountName: account name to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetAccountDetail(clusterId, accountName string) (*AccountDetail, error) {
	result := &AccountDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account/"+accountName).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// GetAccountList - get account list
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetAccountList(clusterId string) (*AccountList, error) {
	result := &AccountList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// UpdateAccountRemark - update account remark
//
// PARAMS:
//   - clusterId: cluster id
//   - accountName: account name to update
//   - args: the arguments to update account
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountRemark(clusterId, accountName string, args *RemarkArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account/"+accountName).
		WithQueryParam("remark", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// UpdateAccountAuthIp - update account auth ip
//
// PARAMS:
//   - clusterId: cluster id
//   - accountName: account name to update
//   - args: the arguments to update account auth ip
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountAuthIp(clusterId, accountName string, args *AuthIpArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account/"+accountName).
		WithQueryParam("authip", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// UpdateAccountPrivileges - update account privileges
//
// PARAMS:
//   - clusterId: cluster id
//   - accountName: account name to update
//   - args: the arguments to update account privileges
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountPrivileges(clusterId, accountName string, args *PrivilegesArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account/"+accountName).
		WithQueryParam("privileges", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// UpdateAccountPassword - update account password
//
// PARAMS:
//   - clusterId: cluster id
//   - accountName: account name to update
//   - args: the arguments to update account password
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAccountPassword(clusterId, accountName string, args *PasswordArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/account/"+accountName).
		WithQueryParam("password", "").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// CreateDatabase - create database
//
// PARAMS:
//   - clusterId: cluster id
//   - args: the arguments to create database
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDatabase(clusterId string, args *CreateDatabaseArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/"+clusterId+"/database").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// DeleteDatabase - delete database
//
// PARAMS:
//   - clusterId: cluster id
//   - dbName: the database name to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteDatabase(clusterId, dbName string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getGaiadbUri()+"/"+clusterId+"/database/"+dbName).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// ListDatabase - list database
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *DatabaseList: the database list
func (c *Client) ListDatabase(clusterId string) (*DatabaseList, error) {
	result := &DatabaseList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/database").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// CreateSnapshot - create snapshot
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSnapshot(clusterId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/"+clusterId+"/snapshot").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// ListSnapshot - list snapshot
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *SnapshotList: the snapshot list
func (c *Client) ListSnapshot(clusterId string) (*SnapshotList, error) {
	result := &SnapshotList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/snapshot").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// UpdateSnapshotPolicy - update snapshot policy
//
// PARAMS:
//   - clusterId: cluster id
//   - args: the arguments to update snapshot policy
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSnapshotPolicy(clusterId string, args *UpdateSnapshotPolicyArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/snapshot/policy").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// GetSnapshotPolicy - get snapshot policy
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *SnapshotPolicy: the snapshot policy
func (c *Client) GetSnapshotPolicy(clusterId string) (*SnapshotPolicy, error) {
	result := &SnapshotPolicy{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/snapshot/policy").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// UpdateWhiteList - update white list
//
// PARAMS:
//   - clusterId : cluster id
//   - args: the arguments to update white list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateWhiteList(clusterId string, args *WhiteList) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/whitelist").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// GetWhiteList - get white list
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *WhiteListResult: the result of white list
func (c *Client) GetWhiteList(clusterId string) (*WhiteList, error) {
	result := &WhiteList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/whitelist").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// CreateMultiactiveGroup - create multiactive group
//
// PARAMS:
//   - args: the arguments to create multiactive group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *CreateMultiactiveGroupResult: multiactive group result
func (c *Client) CreateMultiactiveGroup(args *CreateMultiactiveGroupArgs) (*CreateMultiactiveGroupResult, error) {
	result := &CreateMultiactiveGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/multiactivegroup").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// DeleteMultiactiveGroup - elete multiactive group
//
// PARAMS:
//   - groupId: multiactive group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteMultiactiveGroup(groupId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getGaiadbUri()+"/multiactivegroup/"+groupId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// RenameMultiactiveGroup - rename multiactive group
//
// PARAMS:
//   - groupId: multiactive group id
//   - args: the arguments to rename multiactive group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenameMultiactiveGroup(groupId string, args *RenameMultiactiveGroupArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/multiactivegroup/"+groupId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParam("multiActiveGroupName", "").
		WithBody(args).
		Do()

	return err
}

// MultiactiveGroupList - list multiactive group
//
// PARAMS:
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *MultiactiveGroupListResult: multiactive group list result
func (c *Client) MultiactiveGroupList() (*MultiactiveGroupListResult, error) {
	result := &MultiactiveGroupListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/multiactivegroup").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// MultiactiveGroupDetail - get multiactive group detail
//
// PARAMS:
//   - groupId: multiactive group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *MultiactiveGroupDetailResult: multiactive group Detail result
func (c *Client) MultiactiveGroupDetail(groupId string) (*MultiactiveGroupDetailResult, error) {
	result := &MultiactiveGroupDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/multiactivegroup/"+groupId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// GetSyncStatus - get multiactive group sync status
//
// PARAMS:
//   - groupId: multiactive group id
//   - followerClusterId: the follower cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GetSyncStatusResult: multiactive group Detail result
func (c *Client) GetSyncStatus(groupId, followerClusterId string) (*GetSyncStatusResult, error) {
	result := &GetSyncStatusResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/multiactivegroup/"+groupId+"/syncStatus").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParam("followerClusterId", followerClusterId).
		WithResult(result).
		Do()

	return result, err
}

// GroupExchange - group exchange
//
// PARAMS:
//   - groupId: multiactive group id
//   - args: the arguments to group exchange
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GroupExchange(groupId string, args *ExchangeArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/multiactivegroup/"+groupId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParam("exchange", "").
		WithBody(args).
		Do()

	return err
}

// GetParamsList - get paremeters list
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GetParamsListResult: get paremeters list result
func (c *Client) GetParamsList(clusterId string) (*GetParamsListResult, error) {
	result := &GetParamsListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId+"/compute/params").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// GetParamsHistory - get paremeters history
//
// PARAMS:
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GetParamsHistoryResult: get paremeters history result
func (c *Client) GetParamsHistory(clusterId string) (*GetParamsHistoryResult, error) {
	result := &GetParamsHistoryResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId+"/compute/params/history").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// UpdateParams -  update paremeters
//
// PARAMS:
//   - clusterId: cluster id
//   - args: the arguments to update paremeters
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateParams(clusterId string, args *UpdateParamsArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/cluster/"+clusterId+"/compute/params").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// ListParamTemplate - list param template
//
// PARAMS:
//   - args:  the arguments to list param template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *ListParamTemplateResult: get paremeters list result
func (c *Client) ListParamTemplate(args *ListParamTempArgs) (*ListParamTempResult, error) {
	result := &ListParamTempResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/paramTemplate/listParaTemplate").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParamFilter("detail", strconv.Itoa(args.Detail)).
		WithQueryParamFilter("type", args.Type).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()

	return result, err
}

// SaveAsParamTemplate - save as params template
//
// PARAMS:
//   - args: the arguments to save as params template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SaveAsParamTemplate(args *ParamTempArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/paramTemplate").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// GetTemplateApplyRecords - get template apply records
//
// PARAMS:
//   - templateId: template id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *GetTemplateApplyRecordsResult: get template apply records result
func (c *Client) GetTemplateApplyRecords(templateId string) (*GetTemplateApplyRecordsResult, error) {
	result := &GetTemplateApplyRecordsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId+"/apply").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// DeleteParamsFromTemp - delete params from template
//
// PARAMS:
//   - templateId: template id
//   - args: the params to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteParamsFromTemp(templateId string, args *Params) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId+"/delParams").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// UpdateParamTemplate - update param template
//
// PARAMS:
//   - templateId: template id
//   - args: the params to update
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateParamTemplate(templateId string, args *UpdateParamTplArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// ModifyParams - modify params
//
// PARAMS:
//   - templateId: template id
//   - args: the params to modify
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyParams(templateId string, args *ModifyParamsArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId+"/modifyParams").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// DeleteParamTemplate - delete param template
//
// PARAMS:
//   - templateId: template id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteParamTemplate(templateId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()

	return err
}

// CreateParamTemplate - create param template
//
// PARAMS:
//   - args: the params to create param template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateParamTemplate(args *CreateParamTemplateArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/paramTemplate/create").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// GetParamTemplateDetail - get param template detail
//
// PARAMS:
//   - templateId: template id
//   - detail: detail type
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetParamTemplateDetail(templateId, detail string) (*ParamTemplateDetail, error) {
	result := &ParamTemplateDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId).
		WithQueryParamFilter("detail", detail).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// GetParamTemplateHistory - get param template history
//
// PARAMS:
//   - templateId: template id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetParamTemplateHistory(templateId, action string) (*ParamTemplateHistory, error) {
	result := &ParamTemplateHistory{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId+"/history").
		WithQueryParamFilter("action", action).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// ApplyParamTemplate - apply param template
//
// PARAMS:
//
//	-templateId: template id
//	- args: the params to apply param template
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ApplyParamTemplate(templateId string, args *ApplyParamTemplateArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGaiadbUri()+"/paramTemplate/"+templateId+"/apply").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// UpdateMaintenTime - update maintenTime
//
// PARAMS:
//
//   - clusterId: cluster id
//   - args: the params to update maintenTime
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateMaintenTime(clusterId string, args *UpdateMaintenTimeArgs) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/"+clusterId+"/maintentime").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()

	return err
}

// GetMaintenTime - get maintenTime
//
// PARAMS:
//
//   - clusterId: cluster id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *MaintenTime: maintenTime data
func (c *Client) GetMaintenTime(clusterId string) (*MaintenTimeDetail, error) {
	result := &MaintenTimeDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/maintentime").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// GetSlowSqlDetail - get slow sql detail
//
// PARAMS:
//
//   - args: the params to get slow sql detail
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *SlowSqlDetail:  slow sql detail
func (c *Client) GetSlowSqlDetail(clusterId string, args *GetSlowSqlArgs) (*SlowSqlDetailDetail, error) {
	result := &SlowSqlDetailDetail{}
	params, err2 := getQueryParams(args)
	if err2 != nil {
		return nil, err2
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/slowsql/gaiadb-s").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// SlowSqlAdvice - slow sql advice
//
// PARAMS:
//
//   - clusterId: cluster id
//   - sqlId : sql id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *SlowSqlAdviceDetail:  slow sql advice detail
func (c *Client) SlowSqlAdvice(clusterId, sqlId string) (*SlowSqlAdviceDetail, error) {
	result := &SlowSqlAdviceDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/"+clusterId+"/slowsql/gaiadb-s/"+sqlId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

// GetBinlogDetail - get binlog detail
//
// PARAMS:
//   - binlogId : binlog id
//   - args: the params to get binlog detail
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *BinlogDetail:  binlog detail
func (c *Client) GetBinlogDetail(binlogId string, args *GetBinlogArgs) (*BinlogDetail, error) {
	result := &BinlogDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/binlog/"+binlogId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParamFilter("appID", args.AppID).
		WithQueryParamFilter("logBackupType", args.LogBackupType).
		WithResult(result).
		Do()

	return result, err
}

// GetBinlogList - get binlog list
//
// PARAMS:
//
//   - args: the params to get binlog list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *BinlogList:  binlog List
func (c *Client) GetBinlogList(args *GetBinlogListArgs) (*BinlogList, error) {
	result := &BinlogList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/binlog/list").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParamFilter("appID", args.AppID).
		WithQueryParamFilter("logBackupType", args.LogBackupType).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithQueryParamFilter("startDateTime", args.StartDateTime).
		WithQueryParamFilter("endDateTime", args.EndDateTime).
		WithResult(result).
		Do()

	return result, err
}

// ExecuteTaskNow - execute task now
//
// PARAMS:
//   - taskId : task id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ExecuteTaskNow(taskId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/task/"+taskId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParam("executeNow", "").
		Do()

	return err
}

// CancelTask - cancel task
//
// PARAMS:
//   - taskId : task id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CancelTask(taskId string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getGaiadbUri()+"/task/"+taskId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParam("cancel", "").
		Do()

	return err
}

// GetTaskList - get task list
//
// PARAMS:
//   - args : the params to get task list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *TaskList:  task List
func (c *Client) GetTaskList(args *TaskListArgs) (*TaskList, error) {
	result := &TaskList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/task").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParamFilter("region", args.Region).
		WithQueryParamFilter("startTime", args.StartTime).
		WithQueryParamFilter("endTime", args.EndTime).
		WithResult(result).
		Do()

	return result, err
}

// GetClusterByVpcId - get cluster by vpc id
//
// PARAMS:
//   - vpcId : the params to get cluster by vpc id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *ClusterList:  cluster List
func (c *Client) GetClusterByVpcId(vpcId string) (*ClusterList, error) {
	result := &ClusterList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/security/byVpc").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParamFilter("vpcId", vpcId).
		WithResult(result).
		Do()

	return result, err
}

// GetClusterByLbId - get cluster by lb id
//
// PARAMS:
//   - LbId : the params to get cluster by lb id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *ClusterList:  cluster List
func (c *Client) GetClusterByLbId(lbIds string) (*ClusterList, error) {
	result := &ClusterList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/security/byLb").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithQueryParamFilter("lbIds", lbIds).
		WithResult(result).
		Do()

	return result, err
}

// GetOrderInfo - get order info
//
// PARAMS:
//   - orderId : the params to get order info
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - *OrderInfo:  order info
func (c *Client) GetOrderInfo(orderId string) (*OrderInfo, error) {
	result := &OrderInfo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGaiadbUri()+"/order/"+orderId).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}
