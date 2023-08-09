// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

import (
	bccapi "github.com/baidubce/bce-sdk-go/services/bcc/api"
)

// SupportedInstanceType - CCE 支持的 Instance 类型
var SupportedInstanceType = map[bccapi.InstanceType]string{
	bccapi.InstanceTypeN1: "",
	bccapi.InstanceTypeN2: "",
	bccapi.InstanceTypeN3: "",
	bccapi.InstanceTypeN4: "",
	bccapi.InstanceTypeN5: "",
	bccapi.InstanceTypeC1: "",
	bccapi.InstanceTypeC2: "",
	bccapi.InstanceTypeS1: "",
	bccapi.InstanceTypeG1: "",
	bccapi.InstanceTypeF1: "",
	// 以下为 CCE 自行定义
	InstanceTypeDCC:    "",
	InstanceTypeBBC:    "",
	InstanceTypeBBCGPU: "",
}

// SupportedStorageType - CCE 支持的 Storage 类型
var SupportedStorageType = map[bccapi.StorageType]string{
	bccapi.StorageTypeStd1:     "",
	bccapi.StorageTypeHP1:      "",
	bccapi.StorageTypeCloudHP1: "",
}

// SupportedRootDiskStorageType - CCE 支持的 RootDiskStorage 类型
var SupportedRootDiskStorageType = map[bccapi.StorageType]string{
	bccapi.StorageTypeHP1:      "",
	bccapi.StorageTypeCloudHP1: "",
}

// SupportedGPUType - CCE 支持的 GPU 类型
var SupportedGPUType = map[GPUType]string{
	GPUTypeV100_32: "",
	GPUTypeV100_16: "",
	GPUTypeP40:     "",
	GPUTypeP4:      "",
	GPUTypeK40:     "",
	GPUTypeDLCard:  "",
}

// SupportedK8SVersions - CCE 支持的 K8s 版本
var SupportedK8SVersions = map[K8SVersion]string{
	K8S_1_13_10: "",
	K8S_1_16_8:  "",
}

// SupportedClusterHA - CCE 支持的 ClusterHA 类型
var SupportedClusterHA = map[ClusterHA]string{
	ClusterHALow:    "",
	ClusterHAMedium: "",
	ClusterHAHigh:   "",
}

// SupportedMasterType - CCE 支持 Master 类型
var SupportedMasterType = map[MasterType]string{
	MasterTypeManaged:    "",
	MasterTypeCustom:     "",
	MasterTypeServerless: "",
}

// SupportedImageType - CCE 支持镜像类型
var SupportedImageType = map[bccapi.ImageType]string{
	bccapi.ImageTypeSystem:      "",
	bccapi.ImageTypeCustom:      "",
	bccapi.ImageTypeGPUSystem:   "",
	bccapi.ImageTypeGPUCustom:   "",
	bccapi.ImageTypeSharing:     "",
	bccapi.ImageTypeIntegration: "",
	ImageTypeService:            "",
	// ImageTypeBBCSystem BBC 公有
	bccapi.ImageTypeBBCSystem: "",
	// ImageTypeBBCCustom BBC 自定义
	bccapi.ImageTypeBBCCustom: "",
}

// SupportedContainerNetworkMode - CCE 支持的容器网络类型
var SupportedContainerNetworkMode = map[ContainerNetworkMode]string{
	ContainerNetworkModeKubenet:                  "",
	ContainerNetworkModeVPCCNI:                   "",
	ContainerNetworkModeVPCRouteAutoDetect:       "",
	ContainerNetworkModeVPCRouteVeth:             "",
	ContainerNetworkModeVPCRouteIPVlan:           "",
	ContainerNetworkModeVPCSecondaryIPAutoDetect: "",
	ContainerNetworkModeVPCSecondaryIPVeth:       "",
	ContainerNetworkModeVPCSecondaryIPIPVlan:     "",
}

var SupportedRuntimeType = map[RuntimeType]string{
	RuntimeTypeDocker: "",
}

var SupportedKubeProxyMode = map[KubeProxyMode]string{
	KubeProxyModeIptables: "",
	KubeProxyModeIPVS:     "",
}
