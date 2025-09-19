/*
 * Copyright 2017 Baidu, Inc.
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

// util.go - define the utilities for api package of BCC service
package api

import (
	"encoding/hex"
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util/crypto"
)

const (
	URI_PREFIXV3 = bce.URI_PREFIX + "v3"
	URI_PREFIXV2 = bce.URI_PREFIX + "v2"
	URI_PREFIXV1 = bce.URI_PREFIX + "v1"

	REQUEST_ASP_URI              = "/asp"
	REQUEST_BATCHADDIP_URI       = "/batchAddIp"
	REQUEST_BATCHDELIP_URI       = "/batchDelIp"
	REQUEST_CREATE_URI           = "/create"
	REQUEST_DELETE_URI           = "/delete"
	REQUEST_UPDATE_URI           = "/updateRelation"
	REQUEST_DEL_URI              = "/delRelation"
	REQUEST_DEPLOYSET_URI        = "/deployset"
	REQUEST_IMAGE_URI            = "/image"
	REQUEST_IMAGE_SHAREDUSER_URI = "/sharedUsers"
	REQUEST_IMAGE_OS_URI         = "/os"
	REQUEST_INSTANCE_URI         = "/instance"
	REQUEST_REGION_URI           = "/region"
	REQUEST_BCC_RESERVED_TAG_URI = "/bcc/reserved/tag"
	REQUEST_ServiceType_TAG_URI  = "/bcc/tag"
	REQUEST_INSTANCE_LABEL_URI   = "/instanceByLabel"
	REQUEST_LIST_URI             = "/list"
	REQUEST_SECURITYGROUP_URI    = "/securityGroup"
	REQUEST_SNAPSHOT_URI         = "/snapshot"
	REQUEST_CHAIN_URI            = "/chain"
	REQUEST_SPEC_URI             = "/instance/spec"
	REQUEST_SUBNET_URI           = "/subnet"
	REQUEST_VPC_URI              = "/vpc"
	REQUEST_VNC_SUFFIX           = "/vnc"
	REQUEST_VOLUME_URI           = "/volume"
	REQUEST_ZONE_URI             = "/zone"
	REQUEST_RECYCLE              = "/recycle"
	REQUEST_DELETEPREPAY         = "/volume/deletePrepay"
	REQUEST_PROGRESS             = "/volume/progress"

	//
	REQUEST_FLAVOR_SPEC_URI                        = "/instance/flavorSpec"
	REQUEST_STOCK_GET_SORTED_INST_FLAVORS_URI      = "/stock/getSortedInstFlavors"
	REQUEST_STOCK_GET_INST_OCCUPY_STOCKS_OF_VM_URI = "/stock/getInstOccupyStocksOfVM"
	REQUEST_PRICE_URI                              = "/price"
	REQUEST_AUTO_RENEW_URI                         = "/autoRenew"
	REQUEST_CANCEL_AUTO_RENEW_URI                  = "/cancelAutoRenew"
	REQUEST_BID_PRICE_URI                          = "/bidPrice"
	REQUEST_BID_FLAVOR_URI                         = "/bidFlavor"

	//
	REQUEST_INSTANCE_PRICE_URI               = "/instance/price"
	REQUEST_INSTANCE_BY_SPEC_URI             = "/instanceBySpec"
	REQUEST_VOLUME_DISK_URI                  = "/volume/disk"
	REQUEST_VOLUME_DISK_QUOTA_URI            = "/volume/disk/quota"
	REQUEST_TYPE_ZONE_URI                    = "/instance/flavorZones"
	REQUEST_ENI_URI                          = "/eni"
	REQUEST_KEYPAIR_URI                      = "/keypair"
	REQUEST_REBUILD_URI                      = "/rebuild"
	REQUEST_TAG_URI                          = "/tag"
	REQUEST_NOCHARGE_URI                     = "/noCharge"
	REQUEST_BID_URI                          = "/bid"
	REQUEST_RECOVERY_URI                     = "/recovery"
	REQUEST_CANCEL_BIDORDER_URI              = "/cancelBidOrder"
	REQUEST_BATCH_CREATE_AUTORENEW_RULES_URI = "/batchCreateAutoRenewRules"
	REQUEST_BATCH_Delete_AUTORENEW_RULES_URI = "/batchDeleteAutoRenewRules"
	REQUEST_GET_ALL_STOCKS                   = "/getAllStocks"
	REQUEST_GET_STOCK_WITH_DEPLOYSET         = "/getStockWithDeploySet"
	REQUEST_GET_STOCK_WITH_SPEC              = "/getStockWithSpec"
	REQUEST_GET_AVAILABLE_STOCK_WITH_SPEC    = "/getAvailableStockWithSpec"
	REQUEST_DELETION_PROTECTION              = "/deletionProtection"
	REQUEST_TRANSFER_CREATE_URI              = "/reserved/transfer/create"
	REQUEST_TRANSFER_REVOKE_URI              = "/reserved/transfer/revoke"
	REQUEST_TRANSFER_REFUSE_URI              = "/reserved/transfer/refuse"
	REQUEST_TRANSFER_ACCEPT_URI              = "/reserved/transfer/accept"
	REQUEST_TRANSFER_IN_URI                  = "/reserved/transfer/in/list"
	REQUEST_TRANSFER_OUT_URI                 = "/reserved/transfer/out/list"
	REQUEST_RESERVED_LIST_URI                = "/reserved/list"
	REQUEST_RELATED_DELETE_POLICY            = "/modifyRelatedDeletePolicy"
	REQUEST_VOLUME_PRICE_URI                 = "/volume/getPrice"

	REQUEST_DESCRIBE_REGIONS_URI       = "/describeRegions"
	REQUEST_EHC_CLUSTER_CREATE_URI     = "/ehc/cluster/create"
	REQUEST_EHC_CLUSTER_LIST_URI       = "/ehc/cluster/list"
	REQUEST_EHC_CLUSTER_MODIFY_URI     = "/ehc/cluster/modify"
	REQUEST_EHC_CLUSTER_DELETE_URI     = "/ehc/cluster/delete"
	REQUEST_INSTANCE_USER_DATA_URI     = "/attribute/getUserdata"
	REQUEST_ENTER_RESCUE_MODE_URI      = "/rescue/mode/enter"
	REQUEST_EXIT_RESCUE_MODE_URI       = "/rescue/mode/exit"
	REQUEST_BIND_SECURITY_GROUP__URI   = "/securitygroup/bind"
	REQUEST_UNBIND_SECURITY_GROUP__URI = "/securitygroup/unbind"
	REQUEST_REPLACE_SECURITY_GROUP_URI = "/securitygroup/replace"
)

func getInstanceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI
}

func getInstanceByLabelUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_LABEL_URI
}

func getInstanceUriV3() string {
	return URI_PREFIXV3 + REQUEST_INSTANCE_URI
}

func getRecycleInstanceListUri() string {
	return URI_PREFIXV2 + REQUEST_RECYCLE + REQUEST_INSTANCE_URI
}

func getServersByMarkerV3Uri() string {
	return URI_PREFIXV3 + REQUEST_INSTANCE_URI + REQUEST_LIST_URI
}

func getInstanceBySpecUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_BY_SPEC_URI
}

func getInstanceUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/" + id
}

func getPrepaidInstanceDeleteUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_DELETE_URI
}

func getRecoveryInstanceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_RECOVERY_URI
}

func getBatchAddIpUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_BATCHADDIP_URI
}

func getBatchDelIpUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_BATCHDELIP_URI
}

func getBidInstancePriceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_BID_PRICE_URI
}

func listBidFlavorUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_BID_FLAVOR_URI
}

func getInstanceVNCUri(id string) string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/" + id + REQUEST_VNC_SUFFIX
}

func getInstanceDeletionProtectionUri(id string) string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/" + id + REQUEST_DELETION_PROTECTION
}

func getInstanceRelatedDeletePolicy(id string) string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/" + id + REQUEST_RELATED_DELETE_POLICY
}

func Aes128EncryptUseSecreteKey(sk string, data string) (string, error) {
	if len(sk) < 16 {
		return "", errors.New("error secrete key")
	}

	crypted, err := crypto.EBCEncrypto([]byte(sk[:16]), []byte(data))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypted), nil
}

func getVolumeUri() string {
	return URI_PREFIXV2 + REQUEST_VOLUME_URI
}

func getVolumeV3Uri() string {
	return URI_PREFIXV3 + REQUEST_VOLUME_URI
}

func getVolumeUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_VOLUME_URI + "/" + id
}

func getDeletePrepayVolumeUri() string {
	return URI_PREFIXV2 + REQUEST_DELETEPREPAY
}

func getVolumeProgressUri(volumeId string) string {
	return URI_PREFIXV2 + REQUEST_PROGRESS + "/" + volumeId
}

func getVolumeV3UriWithId(id string) string {
	return URI_PREFIXV3 + REQUEST_VOLUME_URI + "/" + id
}

func getAutoRenewVolumeUri() string {
	return URI_PREFIXV2 + REQUEST_VOLUME_URI + REQUEST_AUTO_RENEW_URI
}

func getCancelAutoRenewVolumeUri() string {
	return URI_PREFIXV2 + REQUEST_VOLUME_URI + REQUEST_CANCEL_AUTO_RENEW_URI
}

func getAvailableDiskInfo() string {
	return URI_PREFIXV2 + REQUEST_VOLUME_DISK_URI
}

func listPurchasableDisksInfo() string {
	return URI_PREFIXV2 + REQUEST_VOLUME_DISK_QUOTA_URI
}

func getSecurityGroupUri() string {
	return URI_PREFIXV2 + REQUEST_SECURITYGROUP_URI
}

func getSecurityGroupUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_SECURITYGROUP_URI + "/" + id
}

func getSecurityGroupRuleUri() string {
	return URI_PREFIXV2 + REQUEST_SECURITYGROUP_URI + "/rule"
}

func getImageUri() string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI
}

func getImageUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI + "/" + id
}

func getRenameImageUri() string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI + "/rename"
}

func getImageSharedUserUri(id string) string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI + "/" + id + REQUEST_IMAGE_SHAREDUSER_URI
}

func getImageOsUri() string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI + REQUEST_IMAGE_OS_URI
}

func getSnapshotUri() string {
	return URI_PREFIXV2 + REQUEST_SNAPSHOT_URI
}

func getSnapshotChainUri() string {
	return URI_PREFIXV2 + REQUEST_SNAPSHOT_URI + REQUEST_CHAIN_URI
}

func getSnapshotUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_SNAPSHOT_URI + "/" + id
}

func getASPUri() string {
	return URI_PREFIXV2 + REQUEST_ASP_URI
}

func getASPUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_ASP_URI + "/" + id
}

func getSpecUri() string {
	return URI_PREFIXV2 + REQUEST_SPEC_URI
}

func getZoneUri() string {
	return URI_PREFIXV2 + REQUEST_ZONE_URI
}

func getPriceBySpecUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_PRICE_URI
}

func getInstanceTypeZoneUri() string {
	return URI_PREFIXV1 + REQUEST_TYPE_ZONE_URI
}

func getChangeSubnetUri() string {
	return URI_PREFIXV2 + REQUEST_SUBNET_URI + "/changeSubnet"
}

func getChangeVpcUri() string {
	return URI_PREFIXV2 + REQUEST_VPC_URI + "/changeVpc"
}

func getInstanceEniUri(instanceId string) string {
	return URI_PREFIXV2 + REQUEST_ENI_URI + "/" + instanceId
}

func getKeypairUri() string {
	return URI_PREFIXV2 + REQUEST_KEYPAIR_URI
}

func getKeypairWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_KEYPAIR_URI + "/" + id
}

func getAllStocks() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_GET_ALL_STOCKS
}

func getStockWithDeploySet() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_GET_STOCK_WITH_DEPLOYSET
}

func getStockWithSpec() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_GET_STOCK_WITH_SPEC
}

func getAvailableStockWithSpec() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_GET_AVAILABLE_STOCK_WITH_SPEC
}

func getSortedInstFlavors() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_STOCK_GET_SORTED_INST_FLAVORS_URI
}

func getInstOccupyStocksOfVm() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_STOCK_GET_INST_OCCUPY_STOCKS_OF_VM_URI
}

func getCreateInstanceStock() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/stock/createInstance"
}

func getResizeInstanceStock() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/stock/resizeInstance"
}

func getFlavorSpecUri() string {
	return URI_PREFIXV2 + REQUEST_FLAVOR_SPEC_URI
}

func getResizeInstanceBySpec(id string) string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_BY_SPEC_URI + "/" + id
}

func getRebuildBatchInstanceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_REBUILD_URI
}

func getChangeToPrepaidUri(id string) string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/" + id
}

func getbindInstanceToTagsUri(id string) string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/" + id + REQUEST_TAG_URI
}

func GetBccReservedToTagsUri() string {
	return URI_PREFIXV2 + REQUEST_BCC_RESERVED_TAG_URI
}

func getCreateReservedInstanceUri() string {
	return URI_PREFIXV2 + "/instance/reserved/create"
}

func getModifyReservedInstancesUri() string {
	return URI_PREFIXV2 + "/instance/reserved/modify"
}

func getRenewReservedInstancesUri() string {
	return URI_PREFIXV2 + "/instance/reserved/renew"
}

func GetServiceTypeTagsUri() string {
	return URI_PREFIXV3 + REQUEST_ServiceType_TAG_URI
}

func GetInstanceNoChargeListUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_NOCHARGE_URI
}

func GetCreateBidInstanceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_BID_URI
}

func GetCancelBidOrderUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_CANCEL_BIDORDER_URI
}

func getBatchCreateAutoRenewRulesUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_BATCH_CREATE_AUTORENEW_RULES_URI
}

func getBatchDeleteAutoRenewRulesUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_BATCH_Delete_AUTORENEW_RULES_URI
}

func getDeleteInstanceDeleteIngorePaymentUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/delete"
}

func getDeleteRecycledInstanceUri(id string) string {
	return URI_PREFIXV2 + "/recycle" + REQUEST_INSTANCE_URI + "/" + id
}

func getDescribeRegionsUri() string {
	return URI_PREFIXV2 + REQUEST_REGION_URI + REQUEST_DESCRIBE_REGIONS_URI
}

func getListInstancesByIdsUrl() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/listByInstanceId"
}

func getBatchDeleteInstanceWithRelatedResourceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/batchDelete"
}

func getBatchStartInstanceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/batchAction"
}

func getBatchStopInstanceUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/batchAction"
}

func getListInstanceTypesUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/types"
}

func getListIdMappingsUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/id/mapping"
}

func getBatchResizeInstanceUri() string {
	return URI_PREFIXV2 + "/instanceBatchBySpec"
}

func getInstanceDeleteProgress() string {
	return URI_PREFIXV2 + "/instance/deleteProgress"
}

func getTagVolumeUri(id string) string {
	return URI_PREFIXV2 + REQUEST_VOLUME_URI + "/" + id + REQUEST_TAG_URI
}

func getUntagVolumeUri(id string) string {
	return URI_PREFIXV2 + REQUEST_VOLUME_URI + "/" + id + REQUEST_TAG_URI
}

func getTagSnapshotChainUri(id string) string {
	return URI_PREFIXV2 + REQUEST_SNAPSHOT_URI + REQUEST_CHAIN_URI + "/" + id + REQUEST_TAG_URI
}

func getUntagSnapshotChainUri(id string) string {
	return URI_PREFIXV2 + REQUEST_SNAPSHOT_URI + REQUEST_CHAIN_URI + "/" + id + REQUEST_TAG_URI
}

func getListInstancesByOrderIdUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/getServersByOrderId"
}

func getCreateInstanceReturnOrderIdUri() string {
	return URI_PREFIXV2 + "/instanceReturnOrderId"
}

func getListAvailableResizeSpecsUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI
}

func getBatchChangeInstanceToPrepay() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/batch/charging"
}

func getBatchChangeInstanceToPostpay() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/batch/charging"
}

func getResizeInstanceReturnOrderId(instanceId string) string {
	return URI_PREFIXV2 + "/instanceReturnOrderId" + instanceId
}

func listInstanceRoles() string {
	return URI_PREFIXV2 + "/instance/role/list"
}

func postInstanceRole() string {
	return URI_PREFIXV2 + "/instance/role"
}

func deleteIpv6() string {
	return URI_PREFIXV2 + "/instance/delIpv6"
}

func addIpv6() string {
	return URI_PREFIXV2 + "/instance/addIpv6"
}

func getImageToTagsUri(id string) string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI + "/" + id + REQUEST_TAG_URI
}

func getRemoteCopySnapshotUri(id string) string {
	return URI_PREFIXV2 + REQUEST_SNAPSHOT_URI + "/remote_copy/" + id
}

func getImportCustomImageUri() string {
	return URI_PREFIXV2 + "/image/import"
}

func getBatchRefundResourceUri() string {
	return URI_PREFIXV2 + "/instance/batchRefundResource"
}

func getAvailableImagesBySpecUri() string {
	return URI_PREFIXV2 + "/image/getAvailableImagesBySpec"
}

func getListReservedInstancesUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_RESERVED_LIST_URI
}

func getCreateTransferReservedInstanceOrderUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_TRANSFER_CREATE_URI
}

func getRevokeTransferReservedInstanceOrderUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_TRANSFER_REVOKE_URI
}

func getRefuseTransferReservedInstanceOrderUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_TRANSFER_REFUSE_URI
}

func getAcceptTransferReservedInstanceOrderUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_TRANSFER_ACCEPT_URI
}

func getTransferInReservedInstanceOrdersUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_TRANSFER_IN_URI
}

func getTransferOutReservedInstanceOrdersUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_TRANSFER_OUT_URI
}

func getCdsPriceUri() string {
	return URI_PREFIXV2 + REQUEST_VOLUME_PRICE_URI
}

func getCreateEhcClusterUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_EHC_CLUSTER_CREATE_URI
}

func getEhcClusterListUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_EHC_CLUSTER_LIST_URI
}

func getEhcClusterModifyUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_EHC_CLUSTER_MODIFY_URI
}

func getEhcClusterDeleteUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_EHC_CLUSTER_DELETE_URI
}

func getInstanceUserDataUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_INSTANCE_USER_DATA_URI
}

func getEnterRescueModeUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_ENTER_RESCUE_MODE_URI
}

func getExitRescueModeUri() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + REQUEST_EXIT_RESCUE_MODE_URI
}

func getBindSecurityGroupUrl() string {
	return URI_PREFIXV2 + REQUEST_BIND_SECURITY_GROUP__URI
}

func getUnbindSecurityGroupUrl() string {
	return URI_PREFIXV2 + REQUEST_UNBIND_SECURITY_GROUP__URI
}

func getReplaceSecurityGroupUrl() string {
	return URI_PREFIXV2 + REQUEST_REPLACE_SECURITY_GROUP_URI
}

func getSnapshotShareUrl() string {
	return URI_PREFIXV2 + "/snapshot/share"
}

func getSnapshotUnShareUrl() string {
	return URI_PREFIXV2 + "/snapshot/unShare"
}

func getSnapshotShareListUrl() string {
	return URI_PREFIXV2 + "/snapshot/snapshotShare/list"
}

func getTaskDetailUrl() string {
	return URI_PREFIXV2 + "/task/detail"
}

func listTasklUrl() string {
	return URI_PREFIXV2 + "/task/list"
}
