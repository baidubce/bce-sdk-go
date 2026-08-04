/*
 * Copyright 2026 Baidu, Inc.
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

// kafka.go - the Kafka APIs definition supported by the Kafka service

package kafka

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	version               = "v2"
	clustersPrefix        = "clusters"
	topicsPrefix          = "topics"
	consumerGroupsPrefix  = "consumer-groups"
	usersPrefix           = "users"
	aclsPrefix            = "acls"
	accessEndpointsPrefix = "access-endpoints"
	nodesPrefix           = "nodes"
	offsetsPrefix         = "offsets"
	partitionsPrefix      = "partitions"
	jobsPrefix            = "jobs"
	operationsPrefix      = "operations"
	configurationsPrefix  = "configurations"
	configsPrefix         = "configs"
	revisionsPrefix       = "revisions"
	messagesPrefix        = "messages"
	quotasPrefix          = "quotas"

	pageNoKey        = "pageNo"
	pageSizeKey      = "pageSize"
	markerKey        = "marker"
	maxKeysKey       = "maxKeys"
	stateKey         = "state"
	modeKey          = "mode"
	nameKey          = "name"
	kafkaVersionKey  = "kafkaVersion"
	paymentKey       = "payment"
	tagKeyKey        = "tagKey"
	tagValueKey      = "tagValue"
	startTimeKey     = "startTime"
	startOffsetKey   = "startOffset"
	clusterNameKey   = "clusterName"
	clusterIDKey     = "clusterId"
	topicNameKey     = "topicName"
	groupNameKey     = "groupName"
	configNameKey    = "configName"
	actionIDKey      = "actionId"
	operationIDKey   = "operationId"
	configIDKey      = "configId"
	revisionIDKey    = "revisionId"
	nodeIDKey        = "nodeId"
	partitionIDKey   = "partitionId"
	usernameKey      = "username"
	passwordKey      = "password"
	patternTypeKey   = "patternType"
	resourceTypeKey  = "resourceType"
	resourceNameKey  = "resourceName"
	operationKey     = "operation"
	entityTypeKey    = "entityType"
	userDefaultKey   = "userDefault"
	clientIDKey      = "clientId"
	clientDefaultKey = "clientDefault"
)

var errRequestNil = errors.New("request should not be nil")

func kafkaURI(parts ...string) string {
	all := append([]string{"", version}, parts...)
	return strings.Join(all, "/")
}

func (c *Client) request(method, url string, result, body interface{}, params map[string]string) error {
	builder := bce.NewRequestBuilder(c).
		WithMethod(method).
		WithURL(url).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	if len(params) > 0 {
		builder.WithQueryParams(params)
	}
	if body != nil {
		builder.WithBody(body)
	}
	if result != nil {
		builder.WithResult(result)
	}
	return builder.Do()
}

func checkStringNotEmpty(value, key string) error {
	if len(value) == 0 {
		return fmt.Errorf("request %s should not be null or empty", key)
	}
	return nil
}

func checkNotNil(value interface{}, key string) error {
	if value == nil {
		return fmt.Errorf("request %s should not be nil", key)
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v.IsNil() {
			return fmt.Errorf("request %s should not be nil", key)
		}
	}
	return nil
}

func checkStringSliceNotEmpty(value []string, key string) error {
	if len(value) == 0 {
		return fmt.Errorf("request %s should not be null or empty", key)
	}
	return nil
}

func addStringParam(params map[string]string, key, value string) {
	if len(value) > 0 {
		params[key] = value
	}
}

func addBoolPtrParam(params map[string]string, key string, value *bool) {
	if value != nil {
		params[key] = strconv.FormatBool(*value)
	}
}

func addIntPtrParam(params map[string]string, key string, value *int) {
	if value != nil {
		params[key] = strconv.Itoa(*value)
	}
}

func addIntParam(params map[string]string, key string, value int) {
	params[key] = strconv.Itoa(value)
}

func addInt64Param(params map[string]string, key string, value int64) {
	params[key] = strconv.FormatInt(value, 10)
}

func addMaxKeysParam(params map[string]string, maxKeys int) {
	if maxKeys > 0 && maxKeys <= 1000 {
		params[maxKeysKey] = strconv.Itoa(maxKeys)
	}
}

func encryptPasswordWithSecret(password, secret string) (string, error) {
	if len(secret) < aes.BlockSize {
		return "", errors.New("secret access key should be at least 16 bytes")
	}
	block, err := aes.NewCipher([]byte(secret[:aes.BlockSize]))
	if err != nil {
		return "", err
	}
	padding := aes.BlockSize - len(password)%aes.BlockSize
	plain := make([]byte, len(password)+padding)
	copy(plain, []byte(password))
	for i := len(password); i < len(plain); i++ {
		plain[i] = byte(padding)
	}
	encrypted := make([]byte, len(plain))
	for start := 0; start < len(plain); start += aes.BlockSize {
		block.Encrypt(encrypted[start:start+aes.BlockSize], plain[start:start+aes.BlockSize])
	}
	return strings.ToUpper(hex.EncodeToString(encrypted)), nil
}

func (c *Client) encryptPassword(password string) (string, error) {
	if c == nil || c.Config == nil {
		return "", errors.New("client configuration should not be nil")
	}
	credentials := c.Config.Credentials
	if credentials == nil {
		return "", errors.New("client credentials should not be nil")
	}
	return encryptPasswordWithSecret(password, credentials.SecretAccessKey)
}

// CreateCluster creates a Kafka cluster.
func (c *Client) CreateCluster(request *CreateClusterRequest) (*CreateClusterResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.Name, nameKey); err != nil {
		return nil, err
	}
	result := &CreateClusterResponse{}
	err := c.request(http.POST, kafkaURI(clustersPrefix), result, request, nil)
	return result, err
}

// DeleteCluster releases a Kafka cluster.
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &DeleteClusterResponse{}
	err := c.request(http.DELETE, kafkaURI(clustersPrefix, request.ClusterID), result, nil, nil)
	return result, err
}

// ListClusters lists Kafka clusters.
func (c *Client) ListClusters(request *ListClustersRequest) (*ListClustersResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	params := map[string]string{}
	addStringParam(params, markerKey, request.Marker)
	addMaxKeysParam(params, request.MaxKeys)
	addStringParam(params, clusterNameKey, request.ClusterName)
	addStringParam(params, stateKey, request.State)
	addStringParam(params, modeKey, request.Mode)
	addStringParam(params, kafkaVersionKey, request.KafkaVersion)
	addStringParam(params, paymentKey, request.Payment)
	if len(request.TagKey) > 0 {
		if request.TagValue == nil {
			return nil, errors.New("request tagValue should not be nil")
		}
		params[tagKeyKey] = request.TagKey
		params[tagValueKey] = *request.TagValue
	} else if request.TagValue != nil {
		return nil, errors.New("request tagKey should not be null or empty")
	}
	result := &ListClustersResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix), result, nil, params)
	return result, err
}

// GetClusterDetail gets Kafka cluster details.
func (c *Client) GetClusterDetail(request *GetClusterDetailRequest) (*GetClusterDetailResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterDetailResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID), result, nil, nil)
	return result, err
}

// GetClusterDeletion gets Kafka cluster deletion details.
func (c *Client) GetClusterDeletion(request *GetClusterDeletionRequest) (*GetClusterDetailResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterDetailResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, "deletion"), result, nil, nil)
	return result, err
}

// GetClusterAccessEndpoints gets cluster access endpoints.
func (c *Client) GetClusterAccessEndpoints(request *GetClusterAccessEndpointsRequest) (*GetClusterAccessEndpointsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterAccessEndpointsResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, accessEndpointsPrefix), result, nil, nil)
	return result, err
}

// GetClusterNodes lists cluster nodes.
func (c *Client) GetClusterNodes(request *GetClusterNodesRequest) (*GetClusterNodesResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, markerKey, request.Marker)
	addMaxKeysParam(params, request.MaxKeys)
	addStringParam(params, stateKey, request.State)
	result := &GetClusterNodesResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, nodesPrefix), result, nil, params)
	return result, err
}

// GetClusterConfigurations gets cluster configurations.
func (c *Client) GetClusterConfigurations(request *GetClusterConfigurationsRequest) (*GetClusterConfigurationsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterConfigurationsResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, configurationsPrefix), result, nil, nil)
	return result, err
}

// IncreaseBrokerCount increases broker count.
func (c *Client) IncreaseBrokerCount(request *IncreaseBrokerCountRequest) (*IncreaseBrokerCountResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &IncreaseBrokerCountResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "increase-broker-count"), result, request, nil)
	return result, err
}

// DecreaseBrokerCount decreases broker count.
func (c *Client) DecreaseBrokerCount(request *DecreaseBrokerCountRequest) (*DecreaseBrokerCountResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &DecreaseBrokerCountResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "decrease-broker-count"), result, request, nil)
	return result, err
}

// MigrateClusterAz migrates cluster availability zones.
func (c *Client) MigrateClusterAz(request *MigrateClusterAzRequest) (*MigrateClusterAzResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &MigrateClusterAzResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "brokers", "migration"), result, request, nil)
	return result, err
}

// UnifyClusterEndpoint unifies cluster access endpoint.
func (c *Client) UnifyClusterEndpoint(request *UnifyClusterEndpointRequest) (*UnifyClusterEndpointResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UnifyClusterEndpointResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, accessEndpointsPrefix), result, request, nil)
	return result, err
}

// UpdateBrokerNodeType updates broker node type.
func (c *Client) UpdateBrokerNodeType(request *UpdateBrokerNodeTypeRequest) (*UpdateBrokerNodeTypeResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UpdateBrokerNodeTypeResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "update-broker-node-type"), result, request, nil)
	return result, err
}

// ExpandBrokerDiskCapacity expands broker disk capacity.
func (c *Client) ExpandBrokerDiskCapacity(request *ExpandBrokerDiskCapacityRequest) (*ExpandBrokerDiskCapacityResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &ExpandBrokerDiskCapacityResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "expand-broker-disk-capacity"), result, request, nil)
	return result, err
}

// UpdateAccessConfig updates cluster access config.
func (c *Client) UpdateAccessConfig(request *UpdateAccessConfigRequest) (*UpdateAccessConfigResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UpdateAccessConfigResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "update-access-config"), result, request, nil)
	return result, err
}

// StartCluster starts a cluster.
func (c *Client) StartCluster(request *StartClusterRequest) (*StartClusterResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &StartClusterResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "start"), result, request, nil)
	return result, err
}

// StopCluster stops a cluster.
func (c *Client) StopCluster(request *StopClusterRequest) (*StopClusterResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &StopClusterResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "stop"), result, request, nil)
	return result, err
}

// ResizeClusterEipBandwidth resizes cluster public IP bandwidth.
func (c *Client) ResizeClusterEipBandwidth(request *ResizeClusterEipBandwidthRequest) (*ResizeClusterEipBandwidthResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &ResizeClusterEipBandwidthResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "eip-bandwidths/resize"), result, request, nil)
	return result, err
}

// SwitchClusterEip switches cluster public IP access.
func (c *Client) SwitchClusterEip(request *SwitchClusterEipRequest) (*SwitchClusterEipResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &SwitchClusterEipResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "eips/switch"), result, request, nil)
	return result, err
}

// UpdateStoragePolicy updates cluster storage policy.
func (c *Client) UpdateStoragePolicy(request *UpdateStoragePolicyRequest) (*UpdateStoragePolicyResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UpdateStoragePolicyResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "update-storage-policy"), result, request, nil)
	return result, err
}

// UpdateKafkaConfig updates cluster Kafka configuration.
func (c *Client) UpdateKafkaConfig(request *UpdateKafkaConfigRequest) (*UpdateKafkaConfigResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UpdateKafkaConfigResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "update-kafka-config"), result, request, nil)
	return result, err
}

// UpdateSecurityGroup updates cluster security groups.
func (c *Client) UpdateSecurityGroup(request *UpdateSecurityGroupRequest) (*UpdateSecurityGroupResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UpdateSecurityGroupResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "security-groups"), result, request, nil)
	return result, err
}

// UpdateMaintenanceDuration updates cluster maintenance duration.
func (c *Client) UpdateMaintenanceDuration(request *UpdateMaintenanceDurationRequest) (*UpdateMaintenanceDurationResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UpdateMaintenanceDurationResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "update-maintenance-duration"), result, request, nil)
	return result, err
}

// SwitchClusterIntranetIp switches cluster intranet IP access.
func (c *Client) SwitchClusterIntranetIp(request *SwitchClusterIntranetIpRequest) (*SwitchClusterIntranetIpResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &SwitchClusterIntranetIpResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "intranet-ips/switch"), result, request, nil)
	return result, err
}

// GetClusterCurrentController gets current controller information.
func (c *Client) GetClusterCurrentController(request *GetClusterCurrentControllerRequest) (*GetClusterCurrentControllerResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterCurrentControllerResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, "controller"), result, nil, nil)
	return result, err
}

// GetClusterHistoryController gets historical controller information.
func (c *Client) GetClusterHistoryController(request *GetClusterHistoryControllerRequest) (*GetClusterHistoryControllerResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterHistoryControllerResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, "controller/history"), result, nil, nil)
	return result, err
}

// RestartCluster restarts the Kafka service on a cluster.
func (c *Client) RestartCluster(request *RestartClusterRequest) (*RestartClusterResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &RestartClusterResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "restart"), result, nil, nil)
	return result, err
}

// RestartBroker restarts a broker node.
func (c *Client) RestartBroker(request *RestartBrokerRequest) (*RestartBrokerResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.NodeID, nodeIDKey); err != nil {
		return nil, err
	}
	result := &RestartBrokerResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, nodesPrefix, request.NodeID, "restart-broker"), result, nil, nil)
	return result, err
}

// SwitchClusterAdvertisedIp switches advertised IP access.
func (c *Client) SwitchClusterAdvertisedIp(request *SwitchClusterAdvertisedIpRequest) (*SwitchClusterAdvertisedIpResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &SwitchClusterAdvertisedIpResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "advertised-ips/switch"), result, request, nil)
	return result, err
}

// SwitchClusterDomain switches cluster domain access.
func (c *Client) SwitchClusterDomain(request *SwitchClusterDomainRequest) (*SwitchClusterDomainResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &SwitchClusterDomainResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, "domains/switch"), result, request, nil)
	return result, err
}

// GetZkPassword gets ZooKeeper username and password.
func (c *Client) GetZkPassword(request *GetZkPasswordRequest) (*GetZkPasswordResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &GetZkPasswordResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, "zookeeper-password"), result, nil, nil)
	return result, err
}

// CreateClusterConfig creates a reusable cluster configuration.
func (c *Client) CreateClusterConfig(request *CreateClusterConfigRequest) (*CreateClusterConfigResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	result := &CreateClusterConfigResponse{}
	err := c.request(http.POST, kafkaURI(configsPrefix), result, request, nil)
	return result, err
}

// ListClusterConfigs lists reusable cluster configurations.
func (c *Client) ListClusterConfigs(request *ListClusterConfigsRequest) (*ListClusterConfigsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	params := map[string]string{}
	addStringParam(params, markerKey, request.Marker)
	addMaxKeysParam(params, request.MaxKeys)
	addStringParam(params, configNameKey, request.ConfigName)
	addStringParam(params, stateKey, request.State)
	result := &ListClusterConfigsResponse{}
	err := c.request(http.GET, kafkaURI(configsPrefix), result, nil, params)
	return result, err
}

// DeleteClusterConfig deletes a reusable cluster configuration.
func (c *Client) DeleteClusterConfig(request *DeleteClusterConfigRequest) (*DeleteClusterConfigResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ConfigID, configIDKey); err != nil {
		return nil, err
	}
	result := &DeleteClusterConfigResponse{}
	err := c.request(http.DELETE, kafkaURI(configsPrefix, request.ConfigID), result, nil, nil)
	return result, err
}

// GetClusterConfig gets reusable cluster configuration details.
func (c *Client) GetClusterConfig(request *GetClusterConfigRequest) (*GetClusterConfigResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ConfigID, configIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterConfigResponse{}
	err := c.request(http.GET, kafkaURI(configsPrefix, request.ConfigID), result, nil, nil)
	return result, err
}

// CreateClusterConfigRevision creates a new revision for a reusable cluster configuration.
func (c *Client) CreateClusterConfigRevision(request *CreateClusterConfigRevisionRequest) (*CreateClusterConfigRevisionResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ConfigID, configIDKey); err != nil {
		return nil, err
	}
	result := &CreateClusterConfigRevisionResponse{}
	err := c.request(http.POST, kafkaURI(configsPrefix, request.ConfigID, revisionsPrefix), result, request, nil)
	return result, err
}

// ListClusterConfigRevisions lists revisions of a reusable cluster configuration.
func (c *Client) ListClusterConfigRevisions(request *ListClusterConfigRevisionsRequest) (*ListClusterConfigRevisionsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ConfigID, configIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, stateKey, request.State)
	result := &ListClusterConfigRevisionsResponse{}
	err := c.request(http.GET, kafkaURI(configsPrefix, request.ConfigID, revisionsPrefix), result, nil, params)
	return result, err
}

// GetClusterConfigRevision gets details of a reusable cluster configuration revision.
func (c *Client) GetClusterConfigRevision(request *GetClusterConfigRevisionRequest) (*GetClusterConfigRevisionResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ConfigID, configIDKey); err != nil {
		return nil, err
	}
	if err := checkNotNil(request.RevisionID, revisionIDKey); err != nil {
		return nil, err
	}
	result := &GetClusterConfigRevisionResponse{}
	err := c.request(http.GET, kafkaURI(configsPrefix, request.ConfigID, revisionsPrefix, strconv.Itoa(*request.RevisionID)), result, nil, nil)
	return result, err
}

// UpdateTopic updates topic partition count or configs.
func (c *Client) UpdateTopic(request *UpdateTopicRequest) (*UpdateTopicResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	if request.PartitionNum == "" && len(request.OtherConfigs) == 0 {
		return nil, errors.New("request partitionNum and otherConfigs should not be both empty")
	}
	result := &UpdateTopicResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName), result, request, nil)
	return result, err
}

// GetSubscribedGroupDetail gets topic subscription details for a consumer group.
func (c *Client) GetSubscribedGroupDetail(request *GetSubscribedGroupDetailRequest) (*GetSubscribedGroupDetailResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.GroupName, groupNameKey); err != nil {
		return nil, err
	}
	result := &GetSubscribedGroupDetailResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, consumerGroupsPrefix, request.GroupName, "subscribe-details"), result, nil, nil)
	return result, err
}

// ListTopicPartitions lists topic partition statuses.
func (c *Client) ListTopicPartitions(request *ListTopicPartitionsRequest) (*ListTopicPartitionsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	pageNo := 1
	if request.PageNo != nil {
		pageNo = *request.PageNo
	}
	pageSize := 10
	if request.PageSize != nil {
		pageSize = *request.PageSize
	}
	addIntParam(params, pageNoKey, pageNo)
	if pageSize > 0 {
		addIntParam(params, pageSizeKey, pageSize)
	}
	result := &ListTopicPartitionsResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, partitionsPrefix, "statuses"), result, nil, params)
	return result, err
}

// GetTopicPartitionDetail gets topic partition details.
func (c *Client) GetTopicPartitionDetail(request *GetTopicPartitionDetailRequest) (*GetTopicPartitionDetailResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.PartitionID, partitionIDKey); err != nil {
		return nil, err
	}
	result := &GetTopicPartitionDetailResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, partitionsPrefix, request.PartitionID, "statuses"), result, nil, nil)
	return result, err
}

// ListTopic lists topics in a cluster.
func (c *Client) ListTopic(request *ListTopicRequest) (*ListTopicResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, topicNameKey, request.TopicName)
	result := &ListTopicResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix), result, nil, params)
	return result, err
}

// GetTopicDetail gets topic details.
func (c *Client) GetTopicDetail(request *GetTopicDetailRequest) (*GetTopicDetailResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	result := &GetTopicDetailResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName), result, nil, nil)
	return result, err
}

// CreateTopic creates a topic.
func (c *Client) CreateTopic(request *CreateTopicRequest) (*CreateTopicResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	result := &CreateTopicResponse{}
	err := c.request(http.POST, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix), result, request, nil)
	return result, err
}

// ListSubscribedGroups lists consumer groups that subscribe a topic.
func (c *Client) ListSubscribedGroups(request *ListSubscribedGroupsRequest) (*ListSubscribedGroupsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	result := &ListSubscribedGroupsResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, consumerGroupsPrefix), result, nil, nil)
	return result, err
}

// DeleteTopic deletes a topic.
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (*DeleteTopicResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	result := &DeleteTopicResponse{}
	err := c.request(http.DELETE, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName), result, nil, nil)
	return result, err
}

// GetTopicPartitionOverview gets topic partition overview.
func (c *Client) GetTopicPartitionOverview(request *GetTopicPartitionOverviewRequest) (*GetTopicPartitionOverviewResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	result := &GetTopicPartitionOverviewResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, partitionsPrefix, "statuses/overview"), result, nil, nil)
	return result, err
}

// GetSubscribedGroupOverview gets topic subscription relationship overview.
func (c *Client) GetSubscribedGroupOverview(request *GetSubscribedGroupOverviewRequest) (*GetSubscribedGroupOverviewResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	result := &GetSubscribedGroupOverviewResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, consumerGroupsPrefix, "overview"), result, nil, nil)
	return result, err
}

// ListTopicConfigOptions lists supported topic config options.
func (c *Client) ListTopicConfigOptions() (*ListTopicConfigOptionsResponse, error) {
	result := &ListTopicConfigOptionsResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, topicsPrefix, "config-options"), result, nil, nil)
	return result, err
}

// SendTopicMessage sends a message to a topic.
func (c *Client) SendTopicMessage(request *SendTopicMessageRequest) (*SendTopicMessageResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	if request.Value == nil {
		return nil, errors.New("request value should not be null or empty")
	}
	if err := checkStringNotEmpty(*request.Value, "value"); err != nil {
		return nil, err
	}
	result := &SendTopicMessageResponse{}
	err := c.request(http.POST, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, messagesPrefix), result, request, nil)
	return result, err
}

// QueryTopicMessagesByStartTime queries messages from a start timestamp.
func (c *Client) QueryTopicMessagesByStartTime(request *QueryTopicMessagesByStartTimeRequest) (*QueryTopicMessagesByStartTimeResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	if request.StartTime <= 0 {
		return nil, errors.New("startTime must be positive")
	}
	params := map[string]string{}
	addIntPtrParam(params, partitionIDKey, request.PartitionID)
	addInt64Param(params, startTimeKey, request.StartTime)
	result := &QueryTopicMessagesByStartTimeResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, messagesPrefix, "query-by-start-time"), result, nil, params)
	return result, err
}

// QueryTopicMessagesByStartOffset queries messages from a start offset.
func (c *Client) QueryTopicMessagesByStartOffset(request *QueryTopicMessagesByStartOffsetRequest) (*QueryTopicMessagesByStartOffsetResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	if request.PartitionID < 0 {
		return nil, errors.New("partitionId must be non-negative")
	}
	if request.StartOffset < 0 {
		return nil, errors.New("startOffset must be non-negative")
	}
	params := map[string]string{}
	addIntParam(params, partitionIDKey, request.PartitionID)
	addInt64Param(params, startOffsetKey, request.StartOffset)
	result := &QueryTopicMessagesByStartOffsetResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, topicsPrefix, request.TopicName, messagesPrefix, "query-by-start-offset"), result, nil, params)
	return result, err
}

// ListSubscribedTopics lists topics subscribed by a consumer group.
func (c *Client) ListSubscribedTopics(request *ListSubscribedTopicsRequest) (*ListSubscribedTopicsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.GroupName, groupNameKey); err != nil {
		return nil, err
	}
	result := &ListSubscribedTopicsResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, consumerGroupsPrefix, request.GroupName, topicsPrefix), result, nil, nil)
	return result, err
}

// ListConsumerGroup lists consumer groups in a cluster.
func (c *Client) ListConsumerGroup(request *ListConsumerGroupRequest) (*ListConsumerGroupResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, groupNameKey, request.GroupName)
	result := &ListConsumerGroupResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, consumerGroupsPrefix), result, nil, params)
	return result, err
}

// DeleteConsumerGroup deletes a consumer group.
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (*DeleteConsumerGroupResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.GroupName, groupNameKey); err != nil {
		return nil, err
	}
	result := &DeleteConsumerGroupResponse{}
	err := c.request(http.DELETE, kafkaURI(clustersPrefix, request.ClusterID, consumerGroupsPrefix, request.GroupName), result, nil, nil)
	return result, err
}

// ResetConsumerGroup resets consumer group offsets.
func (c *Client) ResetConsumerGroup(request *ResetConsumerGroupRequest) (*ResetConsumerGroupResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.GroupName, groupNameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.TopicName, topicNameKey); err != nil {
		return nil, err
	}
	if len(request.Partitions) == 0 {
		return nil, errors.New("request partitions should not be null or empty")
	}
	if err := checkStringNotEmpty(request.ResetStrategy, "resetStrategy"); err != nil {
		return nil, err
	}
	result := &ResetConsumerGroupResponse{}
	err := c.request(http.POST, kafkaURI(clustersPrefix, request.ClusterID, consumerGroupsPrefix, request.GroupName, offsetsPrefix), result, request, nil)
	return result, err
}

// GetSubscribedTopicOverview gets consumer group subscription overview.
func (c *Client) GetSubscribedTopicOverview(request *GetSubscribedTopicOverviewRequest) (*GetSubscribedTopicOverviewResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.GroupName, groupNameKey); err != nil {
		return nil, err
	}
	result := &GetSubscribedTopicOverviewResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, consumerGroupsPrefix, request.GroupName, "topics/overview"), result, nil, nil)
	return result, err
}

// CreateUser creates a Kafka user. Password is encrypted with the client's SK before sending.
func (c *Client) CreateUser(request *CreateUserRequest) (*CreateUserResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Username, usernameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Password, passwordKey); err != nil {
		return nil, err
	}
	password, err := c.encryptPassword(request.Password)
	if err != nil {
		return nil, err
	}
	payload := *request
	payload.Password = password
	result := &CreateUserResponse{}
	err = c.request(http.POST, kafkaURI(clustersPrefix, request.ClusterID, usersPrefix), result, &payload, nil)
	return result, err
}

// DeleteUser deletes a Kafka user.
func (c *Client) DeleteUser(request *DeleteUserRequest) (*DeleteUserResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Username, usernameKey); err != nil {
		return nil, err
	}
	result := &DeleteUserResponse{}
	err := c.request(http.DELETE, kafkaURI(clustersPrefix, request.ClusterID, usersPrefix, request.Username), result, nil, nil)
	return result, err
}

// ResetUserPassword resets a Kafka user's password. Password is encrypted with the client's SK before sending.
func (c *Client) ResetUserPassword(request *ResetUserPasswordRequest) (*ResetUserPasswordResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Username, usernameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Password, passwordKey); err != nil {
		return nil, err
	}
	password, err := c.encryptPassword(request.Password)
	if err != nil {
		return nil, err
	}
	payload := *request
	payload.Password = password
	result := &ResetUserPasswordResponse{}
	err = c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, usersPrefix, request.Username), result, &payload, nil)
	return result, err
}

// ListUsers lists Kafka users. request may be a cluster ID string, ListUsersRequest,
// or *ListUsersRequest to mirror the Java SDK overloads.
func (c *Client) ListUsers(request interface{}) (*ListUserResponse, error) {
	listRequest, err := toListUsersRequest(request)
	if err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(listRequest.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &ListUserResponse{}
	err = c.request(http.GET, kafkaURI(clustersPrefix, listRequest.ClusterID, usersPrefix), result, nil, nil)
	return result, err
}

func toListUsersRequest(request interface{}) (*ListUsersRequest, error) {
	switch value := request.(type) {
	case string:
		return &ListUsersRequest{ClusterID: value}, nil
	case *ListUsersRequest:
		if value == nil {
			return nil, errRequestNil
		}
		return value, nil
	case ListUsersRequest:
		return &value, nil
	default:
		return nil, fmt.Errorf("request should be string or ListUsersRequest")
	}
}

// CreateAcl creates Kafka ACLs.
func (c *Client) CreateAcl(request *CreateAclRequest) (*CreateAclResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Username, usernameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.PatternType, patternTypeKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ResourceType, resourceTypeKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ResourceName, resourceNameKey); err != nil {
		return nil, err
	}
	if err := checkStringSliceNotEmpty(request.Operations, operationKey); err != nil {
		return nil, err
	}
	result := &CreateAclResponse{}
	err := c.request(http.POST, kafkaURI(clustersPrefix, request.ClusterID, aclsPrefix), result, request, nil)
	return result, err
}

// DeleteAcl deletes a Kafka ACL.
func (c *Client) DeleteAcl(request *DeleteAclRequest) (*DeleteAclResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Username, usernameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.PatternType, patternTypeKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ResourceType, resourceTypeKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ResourceName, resourceNameKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.Operation, operationKey); err != nil {
		return nil, err
	}
	params := map[string]string{
		usernameKey:     request.Username,
		patternTypeKey:  request.PatternType,
		resourceTypeKey: request.ResourceType,
		resourceNameKey: request.ResourceName,
		operationKey:    request.Operation,
	}
	result := &DeleteAclResponse{}
	err := c.request(http.DELETE, kafkaURI(clustersPrefix, request.ClusterID, aclsPrefix), result, nil, params)
	return result, err
}

// ListAcls lists Kafka ACLs. request may be a cluster ID string, ListAclRequest,
// or *ListAclRequest to mirror the Java SDK overloads.
func (c *Client) ListAcls(request interface{}) (*ListAclResponse, error) {
	listRequest, err := toListAclRequest(request)
	if err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(listRequest.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, usernameKey, listRequest.Username)
	addStringParam(params, patternTypeKey, listRequest.PatternType)
	addStringParam(params, resourceTypeKey, listRequest.ResourceType)
	addStringParam(params, resourceNameKey, listRequest.ResourceName)
	result := &ListAclResponse{}
	err = c.request(http.GET, kafkaURI(clustersPrefix, listRequest.ClusterID, aclsPrefix), result, nil, params)
	return result, err
}

func toListAclRequest(request interface{}) (*ListAclRequest, error) {
	switch value := request.(type) {
	case string:
		return &ListAclRequest{ClusterID: value}, nil
	case *ListAclRequest:
		if value == nil {
			return nil, errRequestNil
		}
		return value, nil
	case ListAclRequest:
		return &value, nil
	default:
		return nil, fmt.Errorf("request should be string or ListAclRequest")
	}
}

// ListJobs lists cluster jobs.
func (c *Client) ListJobs(request *ListJobsRequest) (*ListJobsResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, markerKey, request.Marker)
	addMaxKeysParam(params, request.MaxKeys)
	addStringParam(params, nameKey, request.Name)
	result := &ListJobsResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, jobsPrefix), result, nil, params)
	return result, err
}

// GetJob gets cluster job details.
func (c *Client) GetJob(request *GetJobDetailRequest) (*GetJobDetailResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ActionID, actionIDKey); err != nil {
		return nil, err
	}
	result := &GetJobDetailResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, jobsPrefix, request.ActionID), result, nil, nil)
	return result, err
}

// GetOperation gets a job operation details.
func (c *Client) GetOperation(request *GetOperationDetailRequest) (*GetOperationDetailResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ActionID, actionIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.OperationID, operationIDKey); err != nil {
		return nil, err
	}
	result := &GetOperationDetailResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, jobsPrefix, request.ActionID, operationsPrefix, request.OperationID), result, nil, nil)
	return result, err
}

// StartJob starts a job.
func (c *Client) StartJob(request *StartJobRequest) (*StartJobResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ActionID, actionIDKey); err != nil {
		return nil, err
	}
	result := &StartJobResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, jobsPrefix, request.ActionID, "start"), result, nil, nil)
	return result, err
}

// CancelJob cancels a job.
func (c *Client) CancelJob(request *CancelJobRequest) (*CancelJobResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ActionID, actionIDKey); err != nil {
		return nil, err
	}
	result := &CancelJobResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, jobsPrefix, request.ActionID, "cancel"), result, nil, nil)
	return result, err
}

// SuspendJob suspends a job.
func (c *Client) SuspendJob(request *SuspendJobRequest) (*SuspendJobResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ActionID, actionIDKey); err != nil {
		return nil, err
	}
	result := &SuspendJobResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, jobsPrefix, request.ActionID, "suspend"), result, nil, nil)
	return result, err
}

// ResumeJob resumes a suspended job.
func (c *Client) ResumeJob(request *ResumeJobRequest) (*ResumeJobResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	if err := checkStringNotEmpty(request.ActionID, actionIDKey); err != nil {
		return nil, err
	}
	result := &ResumeJobResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, jobsPrefix, request.ActionID, "resume"), result, nil, nil)
	return result, err
}

// ListQuotas lists Kafka quotas.
func (c *Client) ListQuotas(request *ListQuotasRequest) (*ListQuotasResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, entityTypeKey, request.EntityType)
	result := &ListQuotasResponse{}
	err := c.request(http.GET, kafkaURI(clustersPrefix, request.ClusterID, quotasPrefix), result, nil, params)
	return result, err
}

// CreateQuota creates a Kafka quota.
func (c *Client) CreateQuota(request *CreateQuotaRequest) (*CreateQuotaResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &CreateQuotaResponse{}
	err := c.request(http.POST, kafkaURI(clustersPrefix, request.ClusterID, quotasPrefix), result, request, nil)
	return result, err
}

// UpdateQuota updates a Kafka quota.
func (c *Client) UpdateQuota(request *UpdateQuotaRequest) (*UpdateQuotaResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	result := &UpdateQuotaResponse{}
	err := c.request(http.PUT, kafkaURI(clustersPrefix, request.ClusterID, quotasPrefix), result, request, nil)
	return result, err
}

// DeleteQuota deletes a Kafka quota.
func (c *Client) DeleteQuota(request *DeleteQuotaRequest) (*DeleteQuotaResponse, error) {
	if request == nil {
		return nil, errRequestNil
	}
	if err := checkStringNotEmpty(request.ClusterID, clusterIDKey); err != nil {
		return nil, err
	}
	params := map[string]string{}
	addStringParam(params, usernameKey, request.Username)
	addBoolPtrParam(params, userDefaultKey, request.UserDefault)
	addStringParam(params, clientIDKey, request.ClientID)
	addBoolPtrParam(params, clientDefaultKey, request.ClientDefault)
	result := &DeleteQuotaResponse{}
	err := c.request(http.DELETE, kafkaURI(clustersPrefix, request.ClusterID, quotasPrefix), result, nil, params)
	return result, err
}
