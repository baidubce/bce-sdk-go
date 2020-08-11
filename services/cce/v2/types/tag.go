// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

// Tag represents TagModel in BCE
type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

// Quota - CCE Cluster/Node Quota
type Quota struct {
	Quota int `json:"quota"`
	Used  int `json:"used"`
}
