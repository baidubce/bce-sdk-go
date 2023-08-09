// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

// BillingMethod 计费方式
type BillingMethod string

const (
	// BillingMethodByTraffic 按照流量计费
	BillingMethodByTraffic BillingMethod = "ByTraffic"
	// BillingMethodByBandwidth 按带宽计费
	BillingMethodByBandwidth BillingMethod = "ByBandwidth"
)

// PaymentTiming 付费时间选择
type PaymentTiming string

const (
	// PaymentTimingPrepaid 预付费
	PaymentTimingPrepaid PaymentTiming = "Prepaid"
	// PaymentTimingPostpaid 后付费
	PaymentTimingPostpaid PaymentTiming = "Postpaid"
)
