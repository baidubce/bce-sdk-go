// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

// ServerStatus BCC 虚机状态
type ServerStatus string

const (
	// ServerStatusActive 虚机运行中
	ServerStatusActive ServerStatus = "ACTIVE"

	// ServerStatusBuild 虚机创建中
	ServerStatusBuild ServerStatus = "BUILD"

	// ServerStatusRebuild 虚机重装系统中
	ServerStatusRebuild ServerStatus = "REBUILD"

	// ServerStatusDeleted 虚机已删除
	ServerStatusDeleted ServerStatus = "DELETED"

	// ServerStatusSnapshot 创建快照
	ServerStatusSnapshot ServerStatus = "SNAPSHOT"

	// ServerStatusDeleteSnapshot 删除快照
	ServerStatusDeleteSnapshot ServerStatus = "DELETE_SNAPSHOT"

	// ServerStatusVolumeResize VOLUME_RESIZE
	ServerStatusVolumeResize ServerStatus = "VOLUME_RESIZE"

	// ServerStatusError 虚机异常
	ServerStatusError ServerStatus = "ERROR"

	// ServerStatusExpired 虚机欠费释放
	ServerStatusExpired ServerStatus = "EXPIRED"

	// ServerStatusReboot 虚机重启
	ServerStatusReboot ServerStatus = "REBOOT"

	// ServerStatusRecharge 虚机续费
	ServerStatusRecharge ServerStatus = "RECHARGE"

	// ServerStatusShutoff 虚机关机
	ServerStatusShutoff ServerStatus = "SHUTOFF"

	// ServerStatusStopped 虚机关机
	ServerStatusStopped ServerStatus = "STOPPED"

	// ServerStatusUnknown 虚机状态未知
	ServerStatusUnknown ServerStatus = "UNKNOWN"
)
