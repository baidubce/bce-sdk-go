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
	"fmt"

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
	REQUEST_UPDATE_URI           = "/updateRelation"
	REQUEST_DEL_URI              = "/delRelation"
	REQUEST_DEPLOYSET_URI        = "/deployset"
	REQUEST_IMAGE_URI            = "/image"
	REQUEST_IMAGE_SHAREDUSER_URI = "/sharedUsers"
	REQUEST_IMAGE_OS_URI         = "/os"
	REQUEST_INSTANCE_URI         = "/instance"
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
	//
	REQUEST_FLAVOR_SPEC_URI       = "/instance/flavorSpec"
	REQUEST_PRICE_URI             = "/price"
	REQUEST_AUTO_RENEW_URI        = "/autoRenew"
	REQUEST_CANCEL_AUTO_RENEW_URI = "/cancelAutoRenew"
	REQUEST_BID_PRICE_URI         = "/bidPrice"
	REQUEST_BID_FLAVOR_URI        = "/bidFlavor"
	//
	REQUEST_INSTANCE_PRICE_URI               = "/instance/price"
	REQUEST_INSTANCE_BY_SPEC_URI             = "/instanceBySpec"
	REQUEST_VOLUME_DISK_URI                  = "/volume/disk"
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
	REQUEST_DELETION_PROTECTION              = "/deletionProtection"
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

func Aes128EncryptUseSecreteKey(sk string, data string) (string, error) {
	if len(sk) < 16 {
		return "", fmt.Errorf("error secrete key")
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

func getSecurityGroupUri() string {
	return URI_PREFIXV2 + REQUEST_SECURITYGROUP_URI
}

func getSecurityGroupUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_SECURITYGROUP_URI + "/" + id
}

func getImageUri() string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI
}

func getImageUriWithId(id string) string {
	return URI_PREFIXV2 + REQUEST_IMAGE_URI + "/" + id
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

func getListInstancesByIdsUrl() string {
	return URI_PREFIXV2 + REQUEST_INSTANCE_URI + "/listByInstanceId"
}
