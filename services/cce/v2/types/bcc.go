// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

import bccapi "github.com/baidubce/bce-sdk-go/services/bcc/api"

const (
	// InstanceTypeDCC DCC 类型
	InstanceTypeDCC bccapi.InstanceType = "DCC"
	// InstanceTypeBBC BBC 类型
	InstanceTypeBBC bccapi.InstanceType = "BBC"
	// InstanceTypeBBCGPU BBC GPU 类型
	InstanceTypeBBCGPU bccapi.InstanceType = "BBC_GPU"
	// InstanceTypeHPAS HPAS 类型
	InstanceTypeHPAS bccapi.InstanceType = "HPAS"
)

// SecurityGroupRule 安全组规则
type SecurityGroupRule struct {
	SecurityGroupID string    `json:"securityGroupId"`
	EtherType       EtherType `json:"ethertype"`
	Direction       Direction `json:"direction"`
	Protocol        Protocol  `json:"protocol"`
	SourceGroupID   string    `json:"sourceGroupId"`
	SourceIP        string    `json:"sourceIp"`
	DestGroupID     string    `json:"destGroupId"`
	DestIP          string    `json:"destIp"`
	PortRange       string    `json:"portRange"`
	Remark          string    `json:"remark"`
}

type Direction string

const (
	DirectionIngress Direction = "ingress"
	DirectionEgress  Direction = "egress"
)

type EtherType string

const (
	EtherTypeIPv4 EtherType = "IPv4"
	EtherTypeIPv6 EtherType = "IPv6"
)

type Protocol string

const (
	ProtocolAll  Protocol = "all"
	ProtocolTCP  Protocol = "tcp"
	ProtocolUDP  Protocol = "udp"
	ProtocolICMP Protocol = "icmp"
)

// GPUType GPU 类型
type GPUType string

const (
	// GPUTypeV100_32 NVIDIA Tesla V100-32G
	GPUTypeV100_32 GPUType = "V100-32"
	// GPUTypeV100_16 NVIDIA Tesla V100-16G
	GPUTypeV100_16 GPUType = "V100-16"
	// GPUTypeP40 P40 NVIDIA Tesla P40
	GPUTypeP40 GPUType = "P40"
	// GPUTypeP4 P4 NVIDIA Tesla P4
	GPUTypeP4 GPUType = "P4"
	// GPUTypeK40 K40 NVIDIA Tesla K40
	GPUTypeK40 GPUType = "K40"
	// GPUTypeDLCard NVIDIA 深度学习开发卡
	GPUTypeDLCard GPUType = "DLCard"
)
