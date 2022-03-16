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

// util.go - define the utilities for api package of BBC service
package bbc

import (
	"encoding/hex"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util/crypto"
)

const (
	URI_PREFIX_V1 = bce.URI_PREFIX + "v1"
	URI_PREFIX_V2 = bce.URI_PREFIX + "v2"

	REQUEST_INSTANCE_URI     	= "/instance"
	REQUEST_INSTANCE_LABEL_URI	= "/instanceByLabel"
	REQUEST_BATCH_DELETE_URI 	= "/batchDelete"
	REQUEST_RECYCLE_URI      	= "/recycle"
	REQUEST_RECOVERY_URI     	= "/recovery"
	REQUEST_SUBNET_URI       	= "/vpcSubnet"
	REQUEST_VPC_URI          	= "/vpc"
	SECURITY_GROUP_URI       	= "/securitygroup"

	REQUEST_IMAGE_URI                        = "/image"
	REQUEST_BATCHADDIP_URI                   = "/batchAddIp"
	REQUEST_BATCHADDIPCROSSSUBNET_URI        = "/batchAddIpCrossSubnet"
	REQUEST_BATCHDELIP_URI                   = "/batchDelIp"
	REQUEST_BATCH_CREATE_AUTORENEW_RULES_URI = "/batchCreateAutoRenewRules"
	REQUEST_BATCH_Delete_AUTORENEW_RULES_URI = "/batchDeleteAutoRenewRules"
	REQUEST_BATCH_REBUILD_INSTANCE_URI       = "/batchRebuild"

	REQUEST_FLAVOR_URI           = "/flavor"
	REQUEST_FLAVOR_RAID_URI      = "/flavorRaid"
	REQUEST_COMMON_IMAGE_URI     = "/flavor/image"
	REQUEST_CUSTOM_IMAGE_URI     = "/customFlavor/image"
	REQUEST_IMAGE_SHAREDUSER_URI = "/sharedUsers"

	REQUEST_FLAVOR_ZONE_URI = "/order/flavorZone"
	REQUEST_FLAVORS_URI     = "/order/flavor"

	REQUEST_OPERATION_LOG_URI = "/operationLog"

	REQUEST_DEPLOY_SET_URI    = "/deployset"
	REQUEST_INSTANCE_PORT_URI = "/vpcPort"

	REQUEST_REPAIR_TASK_URI        = "/task"
	REQUEST_REPAIR_CLOSED_TASK_URI = "/closedTask"

	REQUEST_RULE_URI    = "/rule"
	REQUEST_CREATE_URI  = "/create"
	REQUEST_DELETE_URI  = "/delete"
	REQUEST_DISABLE_URI = "/disable"
	REQUEST_ENABLE_URI  = "/enable"
	REQUEST_VOLUME_URI  = "/volume"
)

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
	return URI_PREFIX_V1 + REQUEST_VOLUME_URI
}

func geBbcStockWithDeploySetUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/getStockWithDeploySet"
}
