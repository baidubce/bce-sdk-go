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

// model.go - definitions of Kafka request and response models

package kafka

import (
	"encoding/json"
	"reflect"
)

type KafkaRequest struct{}

type AuthMode string

const (
	AuthModeNone      AuthMode = "NONE"
	AuthModeSSL       AuthMode = "SSL"
	AuthModeSASLIAM   AuthMode = "SASL_IAM"
	AuthModeSASLSCRAM AuthMode = "SASL_SCRAM"
	AuthModeSASLPlain AuthMode = "SASL_PLAIN"
)

type ClusterConfigOverrideMode string

const (
	ClusterConfigOverrideModeRequired ClusterConfigOverrideMode = "REQUIRED"
	ClusterConfigOverrideModeOptional ClusterConfigOverrideMode = "OPTIONAL"
)

type MaintainPeriod string

const (
	MaintainPeriodMonday    MaintainPeriod = "MONDAY"
	MaintainPeriodTuesday   MaintainPeriod = "TUESDAY"
	MaintainPeriodWednesday MaintainPeriod = "WEDNESDAY"
	MaintainPeriodThursday  MaintainPeriod = "THURSDAY"
	MaintainPeriodFriday    MaintainPeriod = "FRIDAY"
	MaintainPeriodSaturday  MaintainPeriod = "SATURDAY"
	MaintainPeriodSunday    MaintainPeriod = "SUNDAY"
)

type Mode string

const (
	ModeHA Mode = "HA"
	ModeHP Mode = "HP"
)

type StoragePolicyType string

const (
	StoragePolicyTypeNone             StoragePolicyType = "NONE"
	StoragePolicyTypeAutoDelete       StoragePolicyType = "AUTO_DELETE"
	StoragePolicyTypeAutoExpand       StoragePolicyType = "AUTO_EXPAND"
	StoragePolicyTypeDynamicRetention StoragePolicyType = "DYNAMIC_RETENTION"
)

type StorageType string

const (
	StorageTypeSSD            StorageType = "SSD"
	StorageTypeEnhancedSSDPL1 StorageType = "ENHANCED_SSD_PL1"
)

type Type string

const (
	TypeProvisioned Type = "PROVISIONED"
	TypeServerless  Type = "SERVERLESS"
)

type ListRequest struct {
	KafkaRequest
	Marker  string `json:"marker,omitempty"`
	MaxKeys int    `json:"maxKeys,omitempty"`
}

type ListResponse struct {
	Marker      string `json:"marker,omitempty"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker,omitempty"`
	MaxKeys     *int   `json:"maxKeys,omitempty"`
}

type PageListRequest struct {
	KafkaRequest
	PageNo   *int `json:"pageNo,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

type PageListResponse struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

type Acl struct {
	Username     string `json:"username,omitempty"`
	PatternType  string `json:"patternType,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	Operation    string `json:"operation,omitempty"`
}

type CreateAclRequest struct {
	KafkaRequest
	ClusterID    string   `json:"clusterId,omitempty"`
	Username     string   `json:"username,omitempty"`
	PatternType  string   `json:"patternType,omitempty"`
	ResourceType string   `json:"resourceType,omitempty"`
	ResourceName string   `json:"resourceName,omitempty"`
	Operations   []string `json:"operations,omitempty"`
}

func (r CreateAclRequest) MarshalJSON() ([]byte, error) {
	type requestAlias CreateAclRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"operations": r.Operations})
}

type CreateAclResponse struct {
	Username string `json:"username,omitempty"`
}

type DeleteAclRequest struct {
	KafkaRequest
	ClusterID    string `json:"clusterId,omitempty"`
	Username     string `json:"username,omitempty"`
	PatternType  string `json:"patternType,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	Operation    string `json:"operation,omitempty"`
}

type DeleteAclResponse struct {
	Username string `json:"username,omitempty"`
}

type ListAclRequest struct {
	KafkaRequest
	ClusterID    string `json:"-"`
	Username     string `json:"-"`
	PatternType  string `json:"-"`
	ResourceType string `json:"-"`
	ResourceName string `json:"-"`
}

type ListAclResponse struct {
	Acls []Acl `json:"acls,omitempty"`
}

type AccessEndpoint struct {
	SecurityProtocol string `json:"securityProtocol,omitempty"`
	Endpoints        string `json:"endpoints,omitempty"`
	Network          string `json:"network,omitempty"`
}

type Authentication struct {
	Mode    string `json:"mode,omitempty"`
	Context string `json:"context,omitempty"`
}

type Billing struct {
	Payment             string   `json:"payment,omitempty"`
	TimeLength          int      `json:"timeLength"`
	TimeUnit            string   `json:"timeUnit,omitempty"`
	ExpireTime          string   `json:"expireTime,omitempty"`
	AutoRenewEnabled    bool     `json:"autoRenewEnabled"`
	AutoRenewTimeLength int      `json:"autoRenewTimeLength"`
	AutoRenewTimeUnit   string   `json:"autoRenewTimeUnit,omitempty"`
	CouponIds           []string `json:"couponIds,omitempty"`
	IsAutoPay           *bool    `json:"isAutoPay,omitempty"`
}

func NewBilling() *Billing {
	autoPay := true
	return &Billing{
		TimeUnit:          "month",
		AutoRenewTimeUnit: "month",
		IsAutoPay:         &autoPay,
	}
}

func (b Billing) MarshalJSON() ([]byte, error) {
	type billingAlias Billing
	value := struct {
		billingAlias
		TimeUnit          string `json:"timeUnit"`
		AutoRenewTimeUnit string `json:"autoRenewTimeUnit"`
		IsAutoPay         bool   `json:"isAutoPay"`
	}{
		billingAlias:      billingAlias(b),
		TimeUnit:          b.TimeUnit,
		AutoRenewTimeUnit: b.AutoRenewTimeUnit,
		IsAutoPay:         true,
	}
	if value.TimeUnit == "" {
		value.TimeUnit = "month"
	}
	if value.AutoRenewTimeUnit == "" {
		value.AutoRenewTimeUnit = "month"
	}
	if b.IsAutoPay != nil {
		value.IsAutoPay = *b.IsAutoPay
	}
	return marshalWithPresentCollections(value, map[string]interface{}{"couponIds": b.CouponIds})
}

type Cluster struct {
	ClusterID           string   `json:"clusterId,omitempty"`
	ClusterSid          string   `json:"clusterSid,omitempty"`
	Name                string   `json:"name,omitempty"`
	Region              string   `json:"region,omitempty"`
	Type                string   `json:"type,omitempty"`
	Mode                string   `json:"mode,omitempty"`
	State               string   `json:"state,omitempty"`
	KafkaVersion        string   `json:"kafkaVersion,omitempty"`
	LogicalZones        []string `json:"logicalZones,omitempty"`
	Payment             string   `json:"payment,omitempty"`
	ACLEnabled          *bool    `json:"aclEnabled,omitempty"`
	PublicIPEnabled     *bool    `json:"publicIpEnabled,omitempty"`
	IntranetIPEnabled   *bool    `json:"intranetIpEnabled,omitempty"`
	AdvertisedIPEnabled *bool    `json:"advertisedIpEnabled,omitempty"`
	AuthenticationModes []string `json:"authenticationModes,omitempty"`
	Tags                []Tag    `json:"tags,omitempty"`
	CreateTime          string   `json:"createTime,omitempty"`
	ExpireTime          string   `json:"expireTime,omitempty"`
	DeleteTime          string   `json:"deleteTime,omitempty"`
}

type ClusterConfigOption struct {
	Name         string                    `json:"name,omitempty"`
	Description  string                    `json:"description,omitempty"`
	UpdateMode   string                    `json:"updateMode,omitempty"`
	Scope        []interface{}             `json:"scope,omitempty"`
	DefaultValue interface{}               `json:"defaultValue,omitempty"`
	CurrentValue interface{}               `json:"currentValue,omitempty"`
	Type         string                    `json:"type,omitempty"`
	Unit         string                    `json:"unit,omitempty"`
	Category     string                    `json:"category,omitempty"`
	OverrideMode ClusterConfigOverrideMode `json:"overrideMode,omitempty"`
}

type ClusterDetail struct {
	ClusterID   string       `json:"clusterId,omitempty"`
	ClusterSid  string       `json:"clusterSid,omitempty"`
	Name        string       `json:"name,omitempty"`
	Region      string       `json:"region,omitempty"`
	Type        string       `json:"type,omitempty"`
	Mode        string       `json:"mode,omitempty"`
	VPCMode     string       `json:"vpcMode,omitempty"`
	State       string       `json:"state,omitempty"`
	Provisioned *Provisioned `json:"provisioned,omitempty"`
	Tags        []Tag        `json:"tags,omitempty"`
	CreateTime  string       `json:"createTime,omitempty"`
	DeleteTime  string       `json:"deleteTime,omitempty"`
}

type ConfigMeta struct {
	ConfigID   string            `json:"configId,omitempty"`
	RevisionID string            `json:"revisionId,omitempty"`
	Context    map[string]string `json:"context,omitempty"`
}

func NewConfigMeta() *ConfigMeta {
	return &ConfigMeta{Context: map[string]string{}}
}

func (c ConfigMeta) MarshalJSON() ([]byte, error) {
	type configMetaAlias ConfigMeta
	value := struct {
		configMetaAlias
		Context map[string]string `json:"context"`
	}{configMetaAlias: configMetaAlias(c), Context: c.Context}
	if value.Context == nil {
		value.Context = map[string]string{}
	}
	return json.Marshal(value)
}

type Controller struct {
	BrokerID   *int   `json:"brokerId,omitempty"`
	ChangeTime string `json:"changeTime,omitempty"`
	PublicIP   string `json:"publicIp,omitempty"`
	InternalIP string `json:"internalIp,omitempty"`
}

type CreateClusterRequest struct {
	KafkaRequest
	Name         string       `json:"name,omitempty"`
	Mode         Mode         `json:"mode,omitempty"`
	Type         Type         `json:"type,omitempty"`
	EngineType   string       `json:"engineType,omitempty"`
	MetadataMode string       `json:"metadataMode,omitempty"`
	Provisioned  *Provisioned `json:"provisioned,omitempty"`
	Tags         []Tag        `json:"tags,omitempty"`
}

func (r CreateClusterRequest) MarshalJSON() ([]byte, error) {
	type requestAlias CreateClusterRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"tags": r.Tags})
}

type CreateClusterResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
}

type DecreaseBrokerCountRequest struct {
	KafkaRequest
	ClusterID           string `json:"clusterId,omitempty"`
	NumberOfBrokerNodes *int   `json:"numberOfBrokerNodes,omitempty"`
}

type DecreaseBrokerCountResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type DeleteClusterRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type DeleteClusterResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
}

type ExpandBrokerDiskCapacityRequest struct {
	KafkaRequest
	ClusterID   string   `json:"clusterId,omitempty"`
	StorageSize *int64   `json:"storageSize,omitempty"`
	CouponIds   []string `json:"couponIds,omitempty"`
	IsAutoPay   *bool    `json:"isAutoPay,omitempty"`
}

func (r ExpandBrokerDiskCapacityRequest) MarshalJSON() ([]byte, error) {
	type requestAlias ExpandBrokerDiskCapacityRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"couponIds": r.CouponIds})
}

type ExpandBrokerDiskCapacityResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type GetClusterAccessEndpointsRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type GetClusterAccessEndpointsResponse struct {
	AccessEndpoints []AccessEndpoint `json:"accessEndpoints,omitempty"`
}

type GetClusterConfigurationsRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type GetClusterConfigurationsResponse struct {
	Context []ClusterConfigOption `json:"context,omitempty"`
}

type GetClusterCurrentControllerRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type GetClusterCurrentControllerResponse struct {
	Controller *Controller `json:"controller,omitempty"`
}

type GetClusterDeletionRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type GetClusterDetailRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type GetClusterDetailResponse struct {
	Cluster *ClusterDetail `json:"cluster,omitempty"`
}

type GetClusterHistoryControllerRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type GetClusterHistoryControllerResponse struct {
	Controllers []Controller `json:"controllers,omitempty"`
}

type GetClusterNodesRequest struct {
	KafkaRequest
	ListRequest
	ClusterID string `json:"clusterId,omitempty"`
	State     string `json:"state,omitempty"`
}

type GetClusterNodesResponse struct {
	ListResponse
	Nodes []Node `json:"nodes,omitempty"`
}

type GetZkPasswordRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type GetZkPasswordResponse struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type IncreaseBrokerCountRequest struct {
	KafkaRequest
	ClusterID           string   `json:"clusterId,omitempty"`
	NumberOfBrokerNodes *int     `json:"numberOfBrokerNodes,omitempty"`
	CouponIds           []string `json:"couponIds,omitempty"`
	IsAutoPay           *bool    `json:"isAutoPay,omitempty"`
}

func (r IncreaseBrokerCountRequest) MarshalJSON() ([]byte, error) {
	type requestAlias IncreaseBrokerCountRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"couponIds": r.CouponIds})
}

type IncreaseBrokerCountResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type ListClustersRequest struct {
	KafkaRequest
	ListRequest
	ClusterName  string  `json:"clusterName,omitempty"`
	State        string  `json:"state,omitempty"`
	Mode         string  `json:"mode,omitempty"`
	KafkaVersion string  `json:"kafkaVersion,omitempty"`
	Payment      string  `json:"payment,omitempty"`
	TagKey       string  `json:"tagKey,omitempty"`
	TagValue     *string `json:"tagValue,omitempty"`
}

type ListClustersResponse struct {
	ListResponse
	Clusters []Cluster `json:"clusters,omitempty"`
}

type MigrateClusterAzRequest struct {
	KafkaRequest
	ClusterID           string   `json:"clusterId,omitempty"`
	ResizeType          *int     `json:"resizeType,omitempty"`
	CouponIds           []string `json:"couponIds,omitempty"`
	IsAutoPay           *bool    `json:"isAutoPay,omitempty"`
	LogicalZones        []string `json:"logicalZones,omitempty"`
	Subnets             []string `json:"subnets,omitempty"`
	NumberOfBrokerNodes *int     `json:"numberOfBrokerNodes,omitempty"`
	BatchSize           *int     `json:"batchSize,omitempty"`
	InterBrokerThrottle *int64   `json:"interBrokerThrottle,omitempty"`
}

func (r MigrateClusterAzRequest) MarshalJSON() ([]byte, error) {
	type requestAlias MigrateClusterAzRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"couponIds": r.CouponIds, "logicalZones": r.LogicalZones, "subnets": r.Subnets,
	})
}

type MigrateClusterAzResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type Node struct {
	BrokerID   *int   `json:"brokerId,omitempty"`
	Host       string `json:"host,omitempty"`
	NodeID     string `json:"nodeId,omitempty"`
	State      string `json:"state,omitempty"`
	PublicIP   string `json:"publicIp,omitempty"`
	InternalIP string `json:"internalIp,omitempty"`
}

type Provisioned struct {
	KafkaVersion               string           `json:"kafkaVersion,omitempty"`
	Billing                    *Billing         `json:"billing,omitempty"`
	VPC                        *Vpc             `json:"vpc,omitempty"`
	Subnets                    []Subnet         `json:"subnets,omitempty"`
	LogicalZones               []string         `json:"logicalZones,omitempty"`
	SecurityGroups             []SecurityGroup  `json:"securityGroups,omitempty"`
	SecurityGroup              []SecurityGroup  `json:"securityGroup,omitempty"`
	VPCID                      string           `json:"vpcId,omitempty"`
	SubnetIds                  []string         `json:"subnetIds,omitempty"`
	SecurityGroupIds           []string         `json:"securityGroupIds,omitempty"`
	PublicIPEnabled            bool             `json:"publicIpEnabled"`
	PublicIPBandwidth          int              `json:"publicIpBandwidth"`
	IntranetIPEnabled          bool             `json:"intranetIpEnabled"`
	Authentications            []Authentication `json:"authentications,omitempty"`
	ACLEnabled                 bool             `json:"aclEnabled"`
	NumberOfBrokerNodes        int              `json:"numberOfBrokerNodes"`
	NumberOfBrokerNodesPerZone *int             `json:"numberOfBrokerNodesPerZone,omitempty"`
	NodeType                   string           `json:"nodeType,omitempty"`
	StorageMeta                *StorageMeta     `json:"storageMeta,omitempty"`
	StoragePolicyEnabled       *bool            `json:"storagePolicyEnabled,omitempty"`
	StoragePolicy              *StoragePolicy   `json:"storagePolicy,omitempty"`
	ConfigMeta                 *ConfigMeta      `json:"configMeta,omitempty"`
	DeploySetEnabled           bool             `json:"deploySetEnabled"`
	RemoteStorageEnabled       *bool            `json:"remoteStorageEnabled,omitempty"`
	MaintenancePeriods         []MaintainPeriod `json:"maintenancePeriods,omitempty"`
	MaintenanceStartTime       string           `json:"maintenanceStartTime,omitempty"`
	MaintenanceDurationHours   *int             `json:"maintenanceDurationHours,omitempty"`
	AdvertisedIPEnabled        *bool            `json:"advertisedIpEnabled,omitempty"`
	DomainEnabled              *bool            `json:"domainEnabled,omitempty"`
}

func (p Provisioned) MarshalJSON() ([]byte, error) {
	type provisionedAlias Provisioned
	return marshalWithPresentCollections(provisionedAlias(p), map[string]interface{}{
		"subnets": p.Subnets, "logicalZones": p.LogicalZones,
		"securityGroups": p.SecurityGroups, "securityGroup": p.SecurityGroup,
		"subnetIds": p.SubnetIds, "securityGroupIds": p.SecurityGroupIds,
		"authentications": p.Authentications, "maintenancePeriods": p.MaintenancePeriods,
	})
}

type ResizeClusterEipBandwidthRequest struct {
	KafkaRequest
	ClusterID         string   `json:"clusterId,omitempty"`
	PublicIPBandwidth *int     `json:"publicIpBandwidth,omitempty"`
	CouponIds         []string `json:"couponIds,omitempty"`
	IsAutoPay         *bool    `json:"isAutoPay,omitempty"`
}

func (r ResizeClusterEipBandwidthRequest) MarshalJSON() ([]byte, error) {
	type requestAlias ResizeClusterEipBandwidthRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"couponIds": r.CouponIds})
}

type ResizeClusterEipBandwidthResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type RestartBrokerRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	NodeID    string `json:"nodeId,omitempty"`
}

type RestartBrokerResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	NodeID    string `json:"nodeId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type RestartClusterRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type RestartClusterResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type SecurityGroup struct {
	SecurityGroupID   string `json:"securityGroupId,omitempty"`
	SecurityGroupUUID string `json:"securityGroupUuid,omitempty"`
	Name              string `json:"name,omitempty"`
	VPCID             string `json:"vpcId,omitempty"`
	VPCUUID           string `json:"vpcUuid,omitempty"`
}

type StartClusterRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type StartClusterResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type StopClusterRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
}

type StopClusterResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type StorageAutoDelete struct {
	DiskUsedThresholdPercent *int   `json:"diskUsedThresholdPercent,omitempty"`
	LogMinRetentionMs        *int64 `json:"logMinRetentionMs,omitempty"`
	LogMinRetentionBytes     *int64 `json:"logMinRetentionBytes,omitempty"`
}

type StorageAutoExpand struct {
	DiskUsedThresholdPercent *int `json:"diskUsedThresholdPercent,omitempty"`
	StepForwardPercent       *int `json:"stepForwardPercent,omitempty"`
	StepForwardSize          *int `json:"stepForwardSize,omitempty"`
	MaxStorageSize           *int `json:"maxStorageSize,omitempty"`
}

type StorageDynamicRetention struct {
	DiskUsedThresholdPercent *int   `json:"diskUsedThresholdPercent,omitempty"`
	StepForwardPercent       *int   `json:"stepForwardPercent,omitempty"`
	LogMinRetentionMs        *int64 `json:"logMinRetentionMs,omitempty"`
}

type StorageMeta struct {
	StorageType  StorageType `json:"storageType,omitempty"`
	StorageSize  int         `json:"storageSize"`
	NumberOfDisk int         `json:"numberOfDisk"`
}

func NewStorageMeta() *StorageMeta {
	return &StorageMeta{NumberOfDisk: 1}
}

func (s StorageMeta) MarshalJSON() ([]byte, error) {
	type storageMetaAlias StorageMeta
	numberOfDisk := s.NumberOfDisk
	if numberOfDisk == 0 {
		numberOfDisk = 1
	}
	return json.Marshal(struct {
		storageMetaAlias
		NumberOfDisk int `json:"numberOfDisk"`
	}{storageMetaAlias: storageMetaAlias(s), NumberOfDisk: numberOfDisk})
}

type StoragePolicy struct {
	Type             StoragePolicyType        `json:"type,omitempty"`
	AutoDelete       *StorageAutoDelete       `json:"autoDelete,omitempty"`
	AutoExpand       *StorageAutoExpand       `json:"autoExpand,omitempty"`
	DynamicRetention *StorageDynamicRetention `json:"dynamicRetention,omitempty"`
}

type Subnet struct {
	SubnetID   string `json:"subnetId,omitempty"`
	SubnetUUID string `json:"subnetUuid,omitempty"`
	Name       string `json:"name,omitempty"`
	SubnetType string `json:"subnetType,omitempty"`
	Zone       string `json:"zone,omitempty"`
	VPCID      string `json:"vpcId,omitempty"`
	Cidr       string `json:"cidr,omitempty"`
}

type SwitchClusterAdvertisedIpRequest struct {
	KafkaRequest
	ClusterID           string `json:"clusterId,omitempty"`
	AdvertisedIPEnabled *bool  `json:"advertisedIpEnabled,omitempty"`
}

type SwitchClusterAdvertisedIpResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type SwitchClusterDomainRequest struct {
	KafkaRequest
	ClusterID     string `json:"clusterId,omitempty"`
	DomainEnabled *bool  `json:"domainEnabled,omitempty"`
}

type SwitchClusterDomainResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type SwitchClusterEipRequest struct {
	KafkaRequest
	ClusterID          string     `json:"clusterId,omitempty"`
	PublicIPEnabled    *bool      `json:"publicIpEnabled,omitempty"`
	PublicIPBandwidth  *int       `json:"publicIpBandwidth,omitempty"`
	ACLEnabled         *bool      `json:"aclEnabled,omitempty"`
	AuthenticationMode []AuthMode `json:"authenticationMode,omitempty"`
	CouponIds          []string   `json:"couponIds,omitempty"`
	IsAutoPay          *bool      `json:"isAutoPay,omitempty"`
}

func (r SwitchClusterEipRequest) MarshalJSON() ([]byte, error) {
	type requestAlias SwitchClusterEipRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"authenticationMode": r.AuthenticationMode, "couponIds": r.CouponIds,
	})
}

type SwitchClusterEipResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type SwitchClusterIntranetIpRequest struct {
	KafkaRequest
	ClusterID          string     `json:"clusterId,omitempty"`
	IntranetIPEnabled  *bool      `json:"intranetIpEnabled,omitempty"`
	AuthenticationMode []AuthMode `json:"authenticationMode,omitempty"`
}

func (r SwitchClusterIntranetIpRequest) MarshalJSON() ([]byte, error) {
	type requestAlias SwitchClusterIntranetIpRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"authenticationMode": r.AuthenticationMode,
	})
}

type SwitchClusterIntranetIpResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type Tag struct {
	TagKey   string `json:"tagKey,omitempty"`
	TagValue string `json:"tagValue,omitempty"`
}

type UnifyClusterEndpointRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type UnifyClusterEndpointResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type UpdateAccessConfigRequest struct {
	KafkaRequest
	ClusterID       string           `json:"clusterId,omitempty"`
	ACLEnabled      *bool            `json:"aclEnabled,omitempty"`
	Authentications []Authentication `json:"authentications,omitempty"`
}

func (r UpdateAccessConfigRequest) MarshalJSON() ([]byte, error) {
	type requestAlias UpdateAccessConfigRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"authentications": r.Authentications,
	})
}

type UpdateAccessConfigResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type UpdateBrokerNodeTypeRequest struct {
	KafkaRequest
	ClusterID string   `json:"clusterId,omitempty"`
	NodeType  string   `json:"nodeType,omitempty"`
	CouponIds []string `json:"couponIds,omitempty"`
	IsAutoPay *bool    `json:"isAutoPay,omitempty"`
}

func (r UpdateBrokerNodeTypeRequest) MarshalJSON() ([]byte, error) {
	type requestAlias UpdateBrokerNodeTypeRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"couponIds": r.CouponIds})
}

type UpdateBrokerNodeTypeResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type UpdateKafkaConfigRequest struct {
	KafkaRequest
	ClusterID  string `json:"clusterId,omitempty"`
	ConfigID   string `json:"configId,omitempty"`
	RevisionID *int   `json:"revisionId,omitempty"`
}

type UpdateKafkaConfigResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type UpdateMaintenanceDurationRequest struct {
	KafkaRequest
	ClusterID                string           `json:"clusterId,omitempty"`
	MaintenancePeriods       []MaintainPeriod `json:"maintenancePeriods,omitempty"`
	MaintenanceStartTime     string           `json:"maintenanceStartTime,omitempty"`
	MaintenanceDurationHours *int             `json:"maintenanceDurationHours,omitempty"`
}

func (r UpdateMaintenanceDurationRequest) MarshalJSON() ([]byte, error) {
	type requestAlias UpdateMaintenanceDurationRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"maintenancePeriods": r.MaintenancePeriods,
	})
}

type UpdateMaintenanceDurationResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
}

type UpdateSecurityGroupRequest struct {
	KafkaRequest
	ClusterID        string   `json:"clusterId,omitempty"`
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`
}

func (r UpdateSecurityGroupRequest) MarshalJSON() ([]byte, error) {
	type requestAlias UpdateSecurityGroupRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"securityGroupIds": r.SecurityGroupIds,
	})
}

type UpdateSecurityGroupResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type UpdateStoragePolicyRequest struct {
	KafkaRequest
	ClusterID            string         `json:"clusterId,omitempty"`
	StoragePolicyEnabled *bool          `json:"storagePolicyEnabled,omitempty"`
	StoragePolicy        *StoragePolicy `json:"storagePolicy,omitempty"`
}

type UpdateStoragePolicyResponse struct {
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type Vpc struct {
	VPCID   string `json:"vpcId,omitempty"`
	VPCUUID string `json:"vpcUuid,omitempty"`
	Name    string `json:"name,omitempty"`
	Cidr    string `json:"cidr,omitempty"`
}

type ClusterConfig struct {
	ConfigID    string `json:"configId,omitempty"`
	Name        string `json:"name,omitempty"`
	State       string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
}

type ClusterConfigRevision struct {
	RevisionID  *int   `json:"revisionId,omitempty"`
	State       string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
}

type ClusterConfigRevisionDetail struct {
	RevisionID  *int                  `json:"revisionId,omitempty"`
	State       string                `json:"state,omitempty"`
	Description string                `json:"description,omitempty"`
	CreateTime  string                `json:"createTime,omitempty"`
	Context     []ClusterConfigOption `json:"context,omitempty"`
}

type CreateClusterConfigRequest struct {
	KafkaRequest
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Context     map[string]string `json:"context,omitempty"`
}

func (r CreateClusterConfigRequest) MarshalJSON() ([]byte, error) {
	type requestAlias CreateClusterConfigRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"context": r.Context})
}

type CreateClusterConfigResponse struct {
	ConfigID string `json:"configId,omitempty"`
}

type CreateClusterConfigRevisionRequest struct {
	KafkaRequest
	ConfigID    string            `json:"configId,omitempty"`
	RevisionID  *int              `json:"revisionId,omitempty"`
	Description string            `json:"description,omitempty"`
	Context     map[string]string `json:"context,omitempty"`
}

func (r CreateClusterConfigRevisionRequest) MarshalJSON() ([]byte, error) {
	type requestAlias CreateClusterConfigRevisionRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"context": r.Context})
}

type CreateClusterConfigRevisionResponse struct {
	RevisionID *int `json:"revisionId,omitempty"`
}

type DeleteClusterConfigRequest struct {
	KafkaRequest
	ConfigID string `json:"configId,omitempty"`
}

type DeleteClusterConfigResponse struct {
	ConfigID string `json:"configId,omitempty"`
}

type GetClusterConfigRequest struct {
	KafkaRequest
	ConfigID string `json:"configId,omitempty"`
}

type GetClusterConfigResponse struct {
	Config *ClusterConfig `json:"config,omitempty"`
}

type GetClusterConfigRevisionRequest struct {
	KafkaRequest
	ConfigID   string `json:"configId,omitempty"`
	RevisionID *int   `json:"revisionId,omitempty"`
}

type GetClusterConfigRevisionResponse struct {
	Revision *ClusterConfigRevisionDetail `json:"revision,omitempty"`
}

type ListClusterConfigRevisionsRequest struct {
	KafkaRequest
	ConfigID string `json:"configId,omitempty"`
	State    string `json:"state,omitempty"`
}

type ListClusterConfigRevisionsResponse struct {
	Revisions []ClusterConfigRevision `json:"revisions,omitempty"`
}

type ListClusterConfigsRequest struct {
	KafkaRequest
	ListRequest
	ConfigName string `json:"configName,omitempty"`
	State      string `json:"state,omitempty"`
}

type ListClusterConfigsResponse struct {
	ListResponse
	Configs []ClusterConfig `json:"configs,omitempty"`
}

type DeleteConsumerGroupRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

type DeleteConsumerGroupResponse struct {
	GroupName string `json:"groupName,omitempty"`
}

type GetSubscribedTopicOverviewRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

type GetSubscribedTopicOverviewResponse struct {
	Overview *SubscribedTopicOverview `json:"overview,omitempty"`
}

type Group struct {
	GroupName          string `json:"groupName,omitempty"`
	UpdateTime         string `json:"updateTime,omitempty"`
	State              string `json:"state,omitempty"`
	GroupCoordinatorID *int   `json:"groupCoordinatorId,omitempty"`
}

type ListConsumerGroupRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

type ListConsumerGroupResponse struct {
	Groups []Group `json:"groups,omitempty"`
}

type ListSubscribedTopicsRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

type ListSubscribedTopicsResponse struct {
	Topics []string `json:"topics,omitempty"`
}

type ResetConsumerGroupRequest struct {
	KafkaRequest
	ClusterID     string `json:"clusterId,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	TopicName     string `json:"topicName,omitempty"`
	Partitions    []int  `json:"partitions,omitempty"`
	ResetStrategy string `json:"resetStrategy,omitempty"`
	ResetValue    string `json:"resetValue,omitempty"`
}

func (r ResetConsumerGroupRequest) MarshalJSON() ([]byte, error) {
	type requestAlias ResetConsumerGroupRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"partitions": r.Partitions})
}

type ResetConsumerGroupResponse struct {
	GroupName string `json:"groupName,omitempty"`
}

type SubscribedTopicOverview struct {
	SubscribedTopicNum *int   `json:"subscribedTopicNum,omitempty"`
	LastConsumeTime    string `json:"lastConsumeTime,omitempty"`
}

type CancelJobRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type CancelJobResponse struct {
	ActionID string `json:"actionId,omitempty"`
}

type GetJobDetailRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type GetJobDetailResponse struct {
	Job *Job `json:"job,omitempty"`
}

type GetOperationDetailRequest struct {
	KafkaRequest
	ClusterID   string `json:"clusterId,omitempty"`
	ActionID    string `json:"actionId,omitempty"`
	OperationID string `json:"operationId,omitempty"`
}

type GetOperationDetailResponse struct {
	Operation *OperationDetail `json:"operation,omitempty"`
}

type Job struct {
	Name       string      `json:"name,omitempty"`
	ActionID   string      `json:"actionId,omitempty"`
	Status     string      `json:"status,omitempty"`
	Operations []Operation `json:"operations,omitempty"`
}

type JobGroup struct {
	GroupName   string `json:"groupName,omitempty"`
	GroupNameCN string `json:"groupNameCN,omitempty"`
	State       string `json:"state,omitempty"`
	Analysis    bool   `json:"analysis"`
}

type ListJobsRequest struct {
	KafkaRequest
	ListRequest
	ClusterID string `json:"clusterId,omitempty"`
	Name      string `json:"name,omitempty"`
}

type ListJobsResponse struct {
	ListResponse
	Jobs []Job `json:"jobs,omitempty"`
}

type Operation struct {
	ActionID    string `json:"actionId,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	OperationID string `json:"operationId,omitempty"`
	Type        string `json:"type,omitempty"`
	State       string `json:"state,omitempty"`
	Process     *int   `json:"process,omitempty"`
	Schedule    string `json:"schedule,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	StartTime   string `json:"startTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	Started     *bool  `json:"started,omitempty"`
}

type OperationDetail struct {
	ActionID      string     `json:"actionId,omitempty"`
	OperationID   string     `json:"operationId,omitempty"`
	Type          string     `json:"type,omitempty"`
	State         string     `json:"state,omitempty"`
	Process       int        `json:"process"`
	Schedule      string     `json:"schedule,omitempty"`
	Started       *bool      `json:"started,omitempty"`
	Groups        []JobGroup `json:"groups,omitempty"`
	SourceContext string     `json:"sourceContext,omitempty"`
	TargetContext string     `json:"targetContext,omitempty"`
	CreateTime    string     `json:"createTime,omitempty"`
	StartTime     string     `json:"startTime,omitempty"`
	EndTime       string     `json:"endTime,omitempty"`
}

type ResumeJobRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type ResumeJobResponse struct {
	ActionID string `json:"actionId,omitempty"`
}

type StartJobRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type StartJobResponse struct {
	ActionID string `json:"actionId,omitempty"`
}

type SuspendJobRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	ActionID  string `json:"actionId,omitempty"`
}

type SuspendJobResponse struct {
	ActionID string `json:"actionId,omitempty"`
}

type CreateQuotaRequest struct {
	KafkaRequest
	ClusterID        string `json:"clusterId,omitempty"`
	Username         string `json:"username,omitempty"`
	UserDefault      *bool  `json:"userDefault,omitempty"`
	ClientID         string `json:"clientId,omitempty"`
	ClientDefault    *bool  `json:"clientDefault,omitempty"`
	ProducerByteRate *int64 `json:"producerByteRate,omitempty"`
	ConsumerByteRate *int64 `json:"consumerByteRate,omitempty"`
}

type CreateQuotaResponse struct {
	Quota *Quota `json:"quota,omitempty"`
}

type DeleteQuotaRequest struct {
	KafkaRequest
	ClusterID     string `json:"clusterId,omitempty"`
	Username      string `json:"username,omitempty"`
	UserDefault   *bool  `json:"userDefault,omitempty"`
	ClientID      string `json:"clientId,omitempty"`
	ClientDefault *bool  `json:"clientDefault,omitempty"`
}

type DeleteQuotaResponse struct {
	Username      string `json:"username,omitempty"`
	UserDefault   *bool  `json:"userDefault,omitempty"`
	ClientID      string `json:"clientId,omitempty"`
	ClientDefault *bool  `json:"clientDefault,omitempty"`
}

type ListQuotasRequest struct {
	KafkaRequest
	ClusterID  string `json:"clusterId,omitempty"`
	EntityType string `json:"entityType,omitempty"`
}

type ListQuotasResponse struct {
	Quotas []Quota `json:"quotas,omitempty"`
}

type Quota struct {
	Username         string `json:"username,omitempty"`
	UserDefault      *bool  `json:"userDefault,omitempty"`
	ClientID         string `json:"clientId,omitempty"`
	ClientDefault    *bool  `json:"clientDefault,omitempty"`
	ProducerByteRate *int64 `json:"producerByteRate,omitempty"`
	ConsumerByteRate *int64 `json:"consumerByteRate,omitempty"`
}

type UpdateQuotaRequest struct {
	KafkaRequest
	ClusterID        string `json:"clusterId,omitempty"`
	Username         string `json:"username,omitempty"`
	UserDefault      *bool  `json:"userDefault,omitempty"`
	ClientID         string `json:"clientId,omitempty"`
	ClientDefault    *bool  `json:"clientDefault,omitempty"`
	ProducerByteRate *int64 `json:"producerByteRate,omitempty"`
	ConsumerByteRate *int64 `json:"consumerByteRate,omitempty"`
}

type UpdateQuotaResponse struct {
	Quota *Quota `json:"quota,omitempty"`
}

type CreateTopicRequest struct {
	KafkaRequest
	ClusterID         string            `json:"clusterId,omitempty"`
	TopicName         string            `json:"topicName,omitempty"`
	PartitionNum      int               `json:"partitionNum"`
	ReplicationFactor int               `json:"replicationFactor"`
	OtherConfigs      map[string]string `json:"otherConfigs,omitempty"`
}

func (r CreateTopicRequest) MarshalJSON() ([]byte, error) {
	type requestAlias CreateTopicRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"otherConfigs": r.OtherConfigs})
}

type CreateTopicResponse struct {
	TopicName string `json:"topicName,omitempty"`
}

type DeleteTopicRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"topicName,omitempty"`
}

type DeleteTopicResponse struct {
	TopicName string `json:"topicName,omitempty"`
}

type GetSubscribedGroupDetailRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"topicName,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

type GetSubscribedGroupDetailResponse struct {
	SubscribePartitions []GroupTopicPartition `json:"subscribePartitions,omitempty"`
}

type GetSubscribedGroupOverviewRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"topicName,omitempty"`
}

type GetSubscribedGroupOverviewResponse struct {
	Overview *SubscribedGroupOverview `json:"overview,omitempty"`
}

type GetTopicDetailRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"topicName,omitempty"`
}

type GetTopicDetailResponse struct {
	Topic *TopicDetail `json:"topic,omitempty"`
}

type GetTopicPartitionDetailRequest struct {
	KafkaRequest
	ClusterID   string `json:"clusterId,omitempty"`
	TopicName   string `json:"topicName,omitempty"`
	PartitionID string `json:"partitionId,omitempty"`
}

type GetTopicPartitionDetailResponse struct {
	TopicName string          `json:"topicName,omitempty"`
	Partition *TopicPartition `json:"partition,omitempty"`
}

type GetTopicPartitionOverviewRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"topicName,omitempty"`
}

type GetTopicPartitionOverviewResponse struct {
	Overview *TopicPartitionOverview `json:"overview,omitempty"`
}

type GroupTopicPartition struct {
	TopicName       string `json:"topicName,omitempty"`
	PartitionID     int    `json:"partitionId"`
	ConsumerID      string `json:"consumerId,omitempty"`
	ClientID        string `json:"clientId,omitempty"`
	Host            string `json:"host,omitempty"`
	MaxOffset       int64  `json:"maxOffset"`
	CommittedOffset int64  `json:"committedOffset"`
	Lag             int64  `json:"lag"`
	LastConsumeTime string `json:"lastConsumeTime,omitempty"`
}

type ListSubscribedGroupsRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"topicName,omitempty"`
}

type ListSubscribedGroupsResponse struct {
	Groups []string `json:"groups,omitempty"`
}

type ListTopicConfigOptionsRequest struct {
	KafkaRequest
}

type ListTopicConfigOptionsResponse struct {
	TopicConfigs []TopicConfigOption `json:"topicConfigs,omitempty"`
}

type ListTopicPartitionsRequest struct {
	KafkaRequest
	PageListRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"topicName,omitempty"`
}

type ListTopicPartitionsResponse struct {
	PageListResponse
	Partitions []TopicPartition `json:"partitions,omitempty"`
}

type ListTopicRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	TopicName string `json:"-"`
}

type ListTopicResponse struct {
	Topics []Topic `json:"topics,omitempty"`
}

type QueryTopicMessagesByStartOffsetRequest struct {
	KafkaRequest
	ClusterID   string `json:"clusterId,omitempty"`
	TopicName   string `json:"topicName,omitempty"`
	PartitionID int    `json:"partitionId"`
	StartOffset int64  `json:"startOffset"`
}

type QueryTopicMessagesByStartOffsetResponse struct {
	Messages []QueryTopicRecord `json:"messages,omitempty"`
}

type QueryTopicMessagesByStartTimeRequest struct {
	KafkaRequest
	ClusterID   string `json:"clusterId,omitempty"`
	TopicName   string `json:"topicName,omitempty"`
	PartitionID *int   `json:"partitionId,omitempty"`
	StartTime   int64  `json:"startTime"`
}

type QueryTopicMessagesByStartTimeResponse struct {
	Messages []QueryTopicRecord `json:"messages,omitempty"`
}

type QueryTopicRecord struct {
	TopicName   string              `json:"topicName,omitempty"`
	PartitionID int                 `json:"partitionId"`
	Offset      int64               `json:"offset"`
	Timestamp   int64               `json:"timestamp"`
	Key         *string             `json:"key,omitempty"`
	Value       *string             `json:"value,omitempty"`
	Size        int                 `json:"size"`
	Headers     []TopicRecordHeader `json:"headers,omitempty"`
}

type SendTopicMessageRequest struct {
	KafkaRequest
	ClusterID   string  `json:"clusterId,omitempty"`
	TopicName   string  `json:"topicName,omitempty"`
	PartitionID *int    `json:"partitionId,omitempty"`
	Key         *string `json:"key,omitempty"`
	Value       *string `json:"value,omitempty"`
}

type SendTopicMessageResponse struct {
	Message *SendTopicRecord `json:"message,omitempty"`
}

type SendTopicRecord struct {
	TopicName           string  `json:"topicName,omitempty"`
	PartitionID         int     `json:"partitionId"`
	Offset              int64   `json:"offset"`
	Timestamp           int64   `json:"timestamp"`
	Key                 *string `json:"key,omitempty"`
	Value               *string `json:"value,omitempty"`
	SerializedKeySize   int     `json:"serializedKeySize"`
	SerializedValueSize int     `json:"serializedValueSize"`
}

type SubscribedGroupOverview struct {
	SubscribedGroupNum *int   `json:"subscribedGroupNum,omitempty"`
	LastConsumeTime    string `json:"lastConsumeTime,omitempty"`
}

type Topic struct {
	TopicName    string `json:"topicName,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	ReadOnly     *bool  `json:"readOnly,omitempty"`
	PartitionNum *int   `json:"partitionNum,omitempty"`
	ReplicaNum   *int   `json:"replicaNum,omitempty"`
}

type TopicConfig struct {
	Key   string      `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Unit  string      `json:"unit,omitempty"`
}

type TopicConfigOption struct {
	Name         string        `json:"name,omitempty"`
	DefaultValue interface{}   `json:"defaultValue,omitempty"`
	Description  string        `json:"description,omitempty"`
	Type         string        `json:"type,omitempty"`
	Unit         string        `json:"unit,omitempty"`
	ValueScope   []interface{} `json:"valueScope,omitempty"`
}

type TopicDetail struct {
	TopicName           string        `json:"topicName,omitempty"`
	PartitionNum        int           `json:"partitionNum"`
	ReplicationFactor   int           `json:"replicationFactor"`
	BrokersSkewed       float64       `json:"brokersSkewed"`
	BrokersLeaderSkewed float64       `json:"brokersLeaderSkewed"`
	BrokersSpread       float64       `json:"brokersSpread"`
	PreferredReplicas   float64       `json:"preferredReplicas"`
	UnderReplicated     float64       `json:"underReplicated"`
	OtherConfigs        []TopicConfig `json:"otherConfigs,omitempty"`
	ReadOnly            *bool         `json:"readOnly,omitempty"`
}

type TopicPartition struct {
	TopicName      string `json:"topicName,omitempty"`
	PartitionID    int    `json:"partitionId"`
	LeaderID       int    `json:"leaderId"`
	Replicas       []int  `json:"replicas,omitempty"`
	InSyncReplicas []int  `json:"inSyncReplicas,omitempty"`
	MinOffset      int64  `json:"minOffset"`
	MaxOffset      int64  `json:"maxOffset"`
	MessageNum     int64  `json:"messageNum"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
}

type TopicPartitionOverview struct {
	TotalMessageNum int64  `json:"totalMessageNum"`
	LastUpdateTime  string `json:"lastUpdateTime,omitempty"`
}

type TopicRecordHeader struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type UpdateTopicRequest struct {
	KafkaRequest
	ClusterID    string            `json:"clusterId,omitempty"`
	TopicName    string            `json:"topicName,omitempty"`
	PartitionNum string            `json:"partitionNum,omitempty"`
	OtherConfigs map[string]string `json:"otherConfigs,omitempty"`
}

func (r UpdateTopicRequest) MarshalJSON() ([]byte, error) {
	type requestAlias UpdateTopicRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{"otherConfigs": r.OtherConfigs})
}

type UpdateTopicResponse struct {
	TopicName string `json:"topicName,omitempty"`
}

type CreateUserRequest struct {
	KafkaRequest
	ClusterID      string   `json:"clusterId,omitempty"`
	Username       string   `json:"username,omitempty"`
	Password       string   `json:"password,omitempty"`
	SASLMechanisms []string `json:"saslMechanisms,omitempty"`
}

func (r CreateUserRequest) MarshalJSON() ([]byte, error) {
	type requestAlias CreateUserRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"saslMechanisms": r.SASLMechanisms,
	})
}

type CreateUserResponse struct {
	Username string `json:"username,omitempty"`
}

type DeleteUserRequest struct {
	KafkaRequest
	ClusterID string `json:"clusterId,omitempty"`
	Username  string `json:"username,omitempty"`
}

type DeleteUserResponse struct {
	Username string `json:"username,omitempty"`
}

type ListUserResponse struct {
	Users []User `json:"users,omitempty"`
}

type ListUsersRequest struct {
	KafkaRequest
	ClusterID string `json:"-"`
}

type ResetUserPasswordRequest struct {
	KafkaRequest
	ClusterID      string   `json:"clusterId,omitempty"`
	Username       string   `json:"username,omitempty"`
	Password       string   `json:"password,omitempty"`
	SASLMechanisms []string `json:"saslMechanisms,omitempty"`
}

func (r ResetUserPasswordRequest) MarshalJSON() ([]byte, error) {
	type requestAlias ResetUserPasswordRequest
	return marshalWithPresentCollections(requestAlias(r), map[string]interface{}{
		"saslMechanisms": r.SASLMechanisms,
	})
}

type ResetUserPasswordResponse struct {
	Username string `json:"username,omitempty"`
}

type User struct {
	Username       string   `json:"username,omitempty"`
	CreateTime     string   `json:"createTime,omitempty"`
	SASLMechanisms []string `json:"saslMechanisms,omitempty"`
}

func marshalWithPresentCollections(value interface{}, collections map[string]interface{}) ([]byte, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	object := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &object); err != nil {
		return nil, err
	}
	for name, collection := range collections {
		value := reflect.ValueOf(collection)
		if !value.IsValid() || ((value.Kind() == reflect.Map || value.Kind() == reflect.Slice) && value.IsNil()) {
			continue
		}
		encoded, err := json.Marshal(collection)
		if err != nil {
			return nil, err
		}
		object[name] = encoded
	}
	return json.Marshal(object)
}
