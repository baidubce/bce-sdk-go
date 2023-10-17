package as

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type ListAsGroupRequest struct {
	GroupName string `json:"groupName,omitempty"`
	Marker    string `json:"marker,omitempty"`
	MaxKeys   int    `json:"maxKeys,omitempty"`
}

type ListAsGroupResponse struct {
	Marker      string      `json:"marker"`
	IsTruncated bool        `json:"isTruncated"`
	NextMarker  string      `json:"nextMarker"`
	MaxKeys     int         `json:"maxKeys"`
	Result      []GroupInfo `json:"result"`
}

type AsGroupStatus string

const (
	CREATING       AsGroupStatus = "CREATING"
	RUNNING        AsGroupStatus = "RUNNING"
	SCALING_UP     AsGroupStatus = "SCALING_UP"
	SCALING_DOWN   AsGroupStatus = "SCALING_DOWN"
	ATTACHING_NODE AsGroupStatus = "ATTACHING_NODE"
	DETACHING_NODE AsGroupStatus = "DETACHING_NODE"
	DELETING       AsGroupStatus = "DELETING"
	BINDING_BLB    AsGroupStatus = "BINDING_BLB"
	UNBINDING_BLB  AsGroupStatus = "UNBINDING_BLB"
	COOLDOWN       AsGroupStatus = "COOLDOWN"
	PAUSE          AsGroupStatus = "PAUSE"
	DELETED        AsGroupStatus = "DELETED"
)

type ZoneInfo struct {
	Zone       string `json:"zone,omitempty"`
	SubnetID   string `json:"subnetId,omitempty"`
	SubnetUUID string `json:"subnetUuid,omitempty"`
	SubnetName string `json:"subnetName,omitempty"`
	SubnetType int16  `json:"subnetType,omitempty"`
	NodeCount  int    `json:"nodeCount,omitempty"`
}

type GroupConfig struct {
	MinNodeNum    int `json:"minNodeNum,omitempty"`
	MaxNodeNum    int `json:"maxNodeNum,omitempty"`
	CooldownInSec int `json:"cooldownInSec,omitempty"`
	ExpectNum     int `json:"expectNum,omitempty"`
}

type GroupInfo struct {
	GroupId    string        `json:"groupId,omitempty"`
	GroupName  string        `json:"groupName,omitempty"`
	Region     string        `json:"region,omitempty"`
	Status     AsGroupStatus `json:"status,omitempty"`
	VpcId      string        `json:"vpcId,omitempty"`
	NodeNum    int           `json:"nodeNum,omitempty"`
	CreateTime string        `json:"createTime,omitempty"`
	ZoneInfo   []ZoneInfo    `json:"zoneInfo,omitempty"`
	Config     GroupConfig   `json:"config,omitempty"`
	BlbId      string        `json:"blbId,omitempty"`
}

type GetAsGroupRequest struct {
	GroupId string `json:"groupId,omitempty"`
}

type VpcInfo struct {
	VpcId   string `json:"vpcId,omitempty"`
	VpcName string `json:"vpcName,omitempty"`
	VpcUUID string `json:"vpcUuid,omitempty"`
}

type GetAsGroupResponse struct {
	GroupID           string      `json:"groupId,omitempty"`
	GroupName         string      `json:"groupName,omitempty"`
	Region            string      `json:"region,omitempty"`
	Status            string      `json:"status,omitempty"`
	VpcInfo           VpcInfo     `json:"vpcInfo,omitempty"`
	ZoneInfo          []ZoneInfo  `json:"zoneInfo,omitempty"`
	Config            GroupConfig `json:"config,omitempty"`
	BlbID             string      `json:"blbId,omitempty"`
	NodeNum           int         `json:"nodeNum,omitempty"`
	CreateTime        string      `json:"createTime,omitempty"`
	RdsIDs            string      `json:"rdsIds,omitempty"`
	ScsIDs            string      `json:"scsIds,omitempty"`
	ExpansionStrategy string      `json:"expansionStrategy,omitempty"`
	ShrinkageStrategy string      `json:"shrinkageStrategy,omitempty"`
}

type IncreaseAsGroupRequest struct {
	ClientToken       string   `json:"-"`
	GroupId           string   `json:"groupId,omitempty"`
	NodeCount         int      `json:"nodeCount,omitempty"`
	Zone              []string `json:"zone,omitempty"`
	ExpansionStrategy string   `json:"expansionStrategy,omitempty"`
}

type DecreaseAsGroupRequest struct {
	ClientToken string   `json:"-"`
	GroupId     string   `json:"groupId,omitempty"`
	Nodes       []string `json:"nodes,omitempty"`
}

type AdjustAsGroupRequest struct {
	ClientToken string `json:"-"`
	GroupId     string `json:"groupId,omitempty"`
	AdjustNum   int    `json:"adjustNum,omitempty"`
}

type ListAsNodeRequest struct {
	GroupId string `json:"groupId,omitempty"`
	Marker  string `json:"marker,omitempty"`
	MaxKeys int    `json:"maxKeys,omitempty"`
}

type AsEipModel struct {
	BandwidthInMbps int    `json:"bandwidthInMbps,omitempty"`
	EipId           string `json:"eipId,omitempty"`
	Address         string `json:"address,omitempty"`
	EipStatus       string `json:"eipStatus,omitempty"`
	EipAllocationId string `json:"eipAllocationId,omitempty"`
}

type NodeModel struct {
	InstanceId         string           `json:"instanceId,omitempty"`
	InstanceUuid       string           `json:"instanceUuid,omitempty"`
	InstanceName       string           `json:"instanceName,omitempty"`
	FloatingIp         string           `json:"floatingIp,omitempty"`
	InternalIp         string           `json:"internalIp,omitempty"`
	Status             string           `json:"status,omitempty"`
	Payment            string           `json:"payment,omitempty"`
	CpuCount           int64            `json:"cpuCount,omitempty"`
	MemoryCapacityInGB int64            `json:"memoryCapacityInGB,omitempty"`
	InstanceType       string           `json:"instanceType,omitempty"`
	SysDiskInGB        int              `json:"sysDiskInGB,omitempty"`
	CreateTime         string           `json:"createTime,omitempty"`
	Eip                AsEipModel       `json:"eip,omitempty"`
	SubnetType         string           `json:"subnetType,omitempty"`
	IsProtected        bool             `json:"isProtected,omitempty"`
	NodeType           string           `json:"nodeType,omitempty"`
	Tags               []model.TagModel `json:"tags,omitempty"`
	GroupId            string           `json:"groupId,omitempty"`
}

type ListAsNodeResponse struct {
	Marker      string      `json:"marker"`
	IsTruncated bool        `json:"isTruncated"`
	NextMarker  string      `json:"nextMarker"`
	MaxKeys     int         `json:"maxKeys"`
	Result      []NodeModel `json:"result,omitempty"`
}
