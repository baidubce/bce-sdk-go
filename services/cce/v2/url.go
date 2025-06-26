// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package v2

import (
	"encoding/base64"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
)

const (
	URI_PREFIX        = bce.URI_PREFIX + "api/cce/service/v2"
	REMEDY_URI_PREFIX = bce.URI_PREFIX + "api/cce/remedy/v1"

	DEFAULT_ENDPOINT = "cce." + bce.DEFAULT_REGION + ".baidubce.com"

	REQUEST_CLUSTER_URL = "/cluster"

	REQUEST_CLUSTER_LIST_URL = "/clusters"

	REQUEST_INSTANCE_URL = "/instance"

	REQUEST_INSTANCE_LIST_URL = "/instances"

	REQUEST_INSTANCEGROUP_URL = "/instancegroup"

	REQUEST_Addon_URL = "/addon"

	REQUEST_INSTANCEGROUP_LIST_URL = "/instancegroups"

	REQUEST_INSTANCEGROUP_AUTOSCALER_URL = "/autoscaler"

	REQUEST_INSTANCEGROUP_REPLICAS_URL = "/replicas"

	REQUEST_INSTANCEGROUP_SCALE_UP_URL = "/scaleup"

	REQUEST_INSTANCEGROUP_SCALE_DOWN_URL = "/scaledown"

	REQUEST_INSTANCEGROUP_ATTACH_INSTANCE_URL = "/attachInstances"

	REQUEST_QUOTA_URL = "/quota"

	REQUEST_NODE_URL = "/node"

	REQUEST_NET_URL = "/net"

	REQUEST_NET_CHECK_CONTAINER_NETWORK_CIDR_URL = "/check_container_network_cidr"

	REQUEST_NET_CHECK_CLUSTERIP_CIDR_URL = "/check_clusterip_cidr"

	REQUEST_NET_RECOMMEND_CLSUTERIP_CIDR_URL = "/recommend_clusterip_cidr"

	REQUEST_NET_RECOMMEND_CONTAINER_CIDR_URL = "/recommend_container_cidr"

	REQUEST_AUTOSCALER = "/autoscaler"

	REQUEST_KUBECONFIG = "/kubeconfig/%s/%s"

	REQUEST_TASK_URL = "/task"

	REQUEST_TASK_LIST_URL = "/tasks"

	REQUEST_EVENT_URL = "/event"

	REQUEST_SYNC_URL = "/sync"

	REQUEST_RBAC_URL = "/rbac"

	REMEDY_RULE_URL = "/remedyrules"

	CHECK_WEBHOOK_URL = "/webhook/check"

	REMEDYATION_URL = "/remediation"

	REMEDY_TASK_URL = "/remedytasks"

	AUTH_REPAIR_URL = "/authRepair"

	CONFIRM_REPAIR_URL = "/confirm"

	REQUIREREPAIRAUTH_URL = "/requireRepairAuth"

	REQUEST_FORBIDDELETE_URL = "/forbiddelete"

	// backup
	BACKUP_REPO_URL          = "/backuprepositorys"
	BACKUP_Task_URL          = "/backuptasks"
	BACKUP_Schedule_Task_URL = "/backupScheduleRules"
	BACKUP_RESTORE_Task_URL  = "/restoreTasks"

	// 修改节点缩容保护状态
	REQUEST_INSTANCE_SCALEDOWN_PROTECTION_URL = "/instanceScaleDown"
)

var _ Interface = &Client{}

// Client 实现 ccev2.Interface
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getClusterURI() string {
	return URI_PREFIX + REQUEST_CLUSTER_URL
}

func getClusterUriWithIDURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID
}

func getClusterEventStepsURI(clusterID string) string {
	return URI_PREFIX + REQUEST_EVENT_URL + REQUEST_CLUSTER_URL + "/" + clusterID
}

func getInstanceEventStepsURI(instanceID string) string {
	return URI_PREFIX + REQUEST_EVENT_URL + REQUEST_INSTANCE_URL + "/" + instanceID
}

func getSyncInstancesURI(clusterID string) string {
	return URI_PREFIX + REQUEST_SYNC_URL + REQUEST_CLUSTER_URL + clusterID + REQUEST_INSTANCE_LIST_URL
}

func getClusterListURI() string {
	return URI_PREFIX + REQUEST_CLUSTER_LIST_URL
}

func getUpdateClusterForbidDeleteURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_FORBIDDELETE_URL
}

func getClusterInstanceListURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCE_LIST_URL
}

func getClusterInstanceURI(clusterID, instanceID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCE_URL + "/" + instanceID
}

func getClusterInstanceListWithInstanceGroupIDURI(clusterID, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID + REQUEST_INSTANCE_LIST_URL
}

func getInstanceGroupURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL
}

func getInstanceGroupListURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL
}

func getInstanceGroupWithIDURI(clusterID, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID
}

func getInstanceGroupAutoScalerURI(clusterID, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID + REQUEST_INSTANCEGROUP_AUTOSCALER_URL
}

func getScaleUpInstanceGroupURI(clusterID, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID + REQUEST_INSTANCEGROUP_SCALE_UP_URL
}

func getScaleDownInstanceGroupURI(clusterID, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID + REQUEST_INSTANCEGROUP_SCALE_DOWN_URL
}

func getAttachInstancesToInstanceGroupURI(clusterID string, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID + REQUEST_INSTANCEGROUP_ATTACH_INSTANCE_URL
}

func getInstanceGroupReplicasURI(clusterID, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID + REQUEST_INSTANCEGROUP_REPLICAS_URL
}

func getInstanceScaleDownURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCE_SCALEDOWN_PROTECTION_URL
}

func getNetCheckContainerNetworkCIDRURI() string {
	return URI_PREFIX + REQUEST_NET_URL + REQUEST_NET_CHECK_CONTAINER_NETWORK_CIDR_URL
}

func getNetCheckClusterIPCIDRURL() string {
	return URI_PREFIX + REQUEST_NET_URL + REQUEST_NET_CHECK_CLUSTERIP_CIDR_URL
}

func getNetRecommendClusterIpCidrURI() string {
	return URI_PREFIX + REQUEST_NET_URL + REQUEST_NET_RECOMMEND_CLSUTERIP_CIDR_URL
}

func getNetRecommendContainerCidrURI() string {
	return URI_PREFIX + REQUEST_NET_URL + REQUEST_NET_RECOMMEND_CONTAINER_CIDR_URL
}

func getQuotaURI() string {
	return URI_PREFIX + REQUEST_QUOTA_URL + REQUEST_CLUSTER_URL
}

func getQuotaNodeURI(clusterID string) string {
	return URI_PREFIX + REQUEST_QUOTA_URL + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_NODE_URL
}

func getAutoscalerURI(clusterID string) string {
	return URI_PREFIX + REQUEST_AUTOSCALER + "/" + clusterID
}

func getKubeconfigURI(clusterID string, kubeConfigType KubeConfigType) string {
	return URI_PREFIX + fmt.Sprintf(REQUEST_KUBECONFIG, clusterID, kubeConfigType)
}

func getTaskWithIDURI(taskType types.TaskType, taskID string) string {
	return URI_PREFIX + REQUEST_TASK_URL + "/" + string(taskType) + "/" + taskID
}

func getTaskListURI(taskType types.TaskType) string {
	return URI_PREFIX + REQUEST_TASK_LIST_URL + "/" + string(taskType)
}

func genGetInstanceCRDURI(clusterID, cceInstanceID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCE_URL + "/" + cceInstanceID + "/crd"
}

func genUpdateInstanceCRDURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCE_URL + "/crd"
}

func genGetClusterCRDURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + "/crd"
}

func genAddonURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_Addon_URL
}

func genAddonUpgradeURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_Addon_URL + "/upgrade"
}

func genUpdateClusterCRDURI(clusterID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + "/crd"
}

func getRBACURI() string {
	return URI_PREFIX + REQUEST_RBAC_URL
}

func remedyRuleURI(clusterID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REMEDY_RULE_URL
}

func getRemedyRuleURI(clusterID, ruleID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REMEDY_RULE_URL + "/" + ruleID
}

func getCheckWebhookAddressURI(clusterID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REMEDY_RULE_URL + CHECK_WEBHOOK_URL
}

func getInstanceGroupRemediationURI(clusterID, instanceGroupID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDYATION_URL
}

func instanceGroupRemedyRuleURI(clusterID, instanceGroupID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDY_RULE_URL
}

func getInstanceGroupRemedyRuleURI(clusterID, instanceGroupID, ruleID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDY_RULE_URL + "/" + ruleID
}

func instanceGroupRemedyTaskURI(clusterID, instanceGroupID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDY_TASK_URL
}

func getInstanceGroupRemedyTaskURI(clusterID, instanceGroupID, ruleID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDY_TASK_URL + "/" + ruleID
}

func getInstanceGroupRemedyTaskAuthRepairURI(clusterID, instanceGroupID, ruleID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDY_TASK_URL + "/" + ruleID + AUTH_REPAIR_URL
}

func getInstanceGroupRemedyTaskConfirmURI(clusterID, instanceGroupID, ruleID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDY_TASK_URL + "/" + ruleID + CONFIRM_REPAIR_URL
}

func getInstanceGroupRemedyTaskRequestAuthURI(clusterID, instanceGroupID, ruleID string) string {
	return REMEDY_URI_PREFIX + REQUEST_CLUSTER_LIST_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_LIST_URL + "/" + instanceGroupID + REMEDY_TASK_URL + "/" + ruleID + REQUIREREPAIRAUTH_URL
}

// backup url
func getBackupRepoURL(id string) string {
	if id == "" {
		return URI_PREFIX + BACKUP_REPO_URL
	}
	return URI_PREFIX + BACKUP_REPO_URL + "/" + id
}

func getBackupTaskURL(clusterID, taskID string) string {
	if taskID == "" {
		return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + BACKUP_Task_URL
	}
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + BACKUP_Task_URL + "/" + taskID
}

func getScheduleBackupTaskURL(clusterID, taskID string) string {
	if taskID == "" {
		return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + BACKUP_Schedule_Task_URL
	}
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + BACKUP_Schedule_Task_URL + "/" + taskID
}

func getRestoreBackupTaskURL(clusterID, taskID string) string {
	if taskID == "" {
		return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + BACKUP_RESTORE_Task_URL
	}
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + BACKUP_RESTORE_Task_URL + "/" + taskID
}

func encodeUserScriptInInstanceSet(instancesSets []*InstanceSet) error {
	if instancesSets == nil {
		return nil
	}
	for _, instanceSet := range instancesSets {
		encodeUserScript(&instanceSet.InstanceSpec)
	}
	return nil
}

func encodeUserScript(instanceSpec *types.InstanceSpec) {
	if instanceSpec == nil {
		return
	}
	if instanceSpec.DeployCustomConfig.PreUserScript != "" {
		base64Str := base64.StdEncoding.EncodeToString([]byte(instanceSpec.DeployCustomConfig.PreUserScript))
		instanceSpec.DeployCustomConfig.PreUserScript = base64Str
	}
	if instanceSpec.DeployCustomConfig.PostUserScript != "" {
		base64Str := base64.StdEncoding.EncodeToString([]byte(instanceSpec.DeployCustomConfig.PostUserScript))
		instanceSpec.DeployCustomConfig.PostUserScript = base64Str
	}
}
