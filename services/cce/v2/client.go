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
	URI_PREFIX = bce.URI_PREFIX + "api/cce/service/v2"

	DEFAULT_ENDPOINT = "cce." + bce.DEFAULT_REGION + ".baidubce.com"

	REQUEST_CLUSTER_URL = "/cluster"

	REQUEST_CLUSTER_LIST_URL = "/clusters"

	REQUEST_INSTANCE_URL = "/instance"

	REQUEST_INSTANCE_LIST_URL = "/instances"

	REQUEST_INSTANCEGROUP_URL = "/instancegroup"

	REQUEST_INSTANCEGROUP_LIST_URL = "/instancegroups"

	REQUEST_INSTANCEGROUP_AUTOSCALER_URL = "/autoscaler"

	REQUEST_INSTANCEGROUP_REPLICAS_URL = "/replicas"

	REQUEST_QUOTA_URL = "/quota"

	REQUEST_NODE_URL = "/node"

	REQUEST_NET_URL = "/net"

	REQUEST_NET_CHECK_CONTAINER_NETWORK_CIDR_URL = "/check_container_network_cidr"

	REQUEST_NET_CHECK_CLUSTERIP_CIDR_URL = "/check_clusterip_cidr"

	REQUEST_NET_RECOMMEND_CLSUTERIP_CIDR_URL = "/recommend_clusterip_cidr"

	REQUEST_NET_RECOMMEND_CONTAINER_CIDR_URL = "/recommend_container_cidr"

	REQUEST_AUTOSCALER = "/autoscaler"

	REQUEST_KUBECONFIG = "/kubeconfig/%s/%s"
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

func getClusterListURI() string {
	return URI_PREFIX + REQUEST_CLUSTER_LIST_URL
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

func getInstanceGroupReplicasURI(clusterID, instanceGroupID string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterID + REQUEST_INSTANCEGROUP_URL + "/" + instanceGroupID + REQUEST_INSTANCEGROUP_REPLICAS_URL
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
