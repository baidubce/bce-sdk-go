// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

// BLBType for load balancer type
type BLBType string

const (
	// BLBTypeNormal 普通 BLB 类型
	BLBTypeNormal BLBType = "normal"

	// BLBTypeApplication 应用型 BLB 类型
	BLBTypeApplication BLBType = "application"
)
