/*
 * Copyright 2022 Baidu, Inc.
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

// model.go - definitions of the request arguments and results data structure model

package lbdc

import "github.com/baidubce/bce-sdk-go/model"

type Reservation struct {
	ReservationLength int `json:"reservationLength"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming"`
	Reservation   *Reservation `json:"reservation"`
}

type ReservationForCreate struct {
	ReservationLength int `json:"reservationLength,omitempty"`
}

type BillingForRenew struct {
	Reservation *Reservation `json:"reservation"`
}

type Cluster struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	CcuCount    int    `json:"ccuCount"`
	CreateTime  string `json:"createTime"`
	ExpireTime  string `json:"expireTime"`
	Description string `json:"desc"`
}

// CreateLbdcArgs defines the structure of input parameters for the CreateLbdc api
type CreateLbdcArgs struct {
	ClientToken      string           `json:"-"`
	Name             string           `json:"name"`
	Type             string           `json:"type"`
	CcuCount         int              `json:"ccuCount"`
	Description      *string          `json:"desc,omitempty"`
	Billing          *Billing         `json:"billing"`
	RenewReservation *Reservation     `json:"renewReservation"`
	Tags             []model.TagModel `json:"tags,omitempty"`
}

// CreateLbdcResult defines the structure of output parameters for the CreateLbdc api
type CreateLbdcResult struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"desc"`
}

// UpgradeLbdcArgs defines the structure of input parameters for the UpgradeLbdc api
type UpgradeLbdcArgs struct {
	ClientToken string `json:"-"`
	Id          string `json:"id"`
	CcuCount    int    `json:"ccuCount"`
}

// RenewLbdcArgs defines the structure of input parameters for the RenewLbdc api
type RenewLbdcArgs struct {
	ClientToken string           `json:"-"`
	Id          string           `json:"id"`
	Billing     *BillingForRenew `json:"billing"`
}

// ListLbdcArgs defines the structure of input parameters for the ListLbdc api
type ListLbdcArgs struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ListLbdcResult defines the structure of output parameters for the ListLbdc api
type ListLbdcResult struct {
	Marker      string    `json:"marker"`
	IsTruncated bool      `json:"isTruncated"`
	NextMarker  string    `json:"nextMarker"`
	MaxKeys     int       `json:"maxKeys"`
	ClusterList []Cluster `json:"clusterList"`
}

// GetLbdcDetailResult defines the structure of output parameters for the GetLbdcDetail api
type GetLbdcDetailResult struct {
	// 4 layer
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Status            string `json:"status"`
	CcuCount          int    `json:"ccuCount"`
	CreateTime        string `json:"createTime"`
	ExpireTime        string `json:"expireTime"`
	TotalConnectCount int64  `json:"totalConnectCount"`
	NewConnectCps     *int64 `json:"newConnectCps,omitempty"`
	NetworkInBps      int64  `json:"networkInBps"`
	NetworkOutBps     int64  `json:"networkOutBps"`

	// 7layer
	HttpsQps           *int64           `json:"httpsQps,omitempty"`
	HttpQps            *int64           `json:"httpQps,omitempty"`
	HttpNewConnectCps  *int64           `json:"httpNewConnectCps,omitempty"`
	HttpsNewConnectCps *int64           `json:"httpsNewConnectCps,omitempty"`
	SslNewConnectCps   *int64           `json:"sslNewConnectCps,omitempty"`
	Tags               []model.TagModel `json:"tags,omitempty"`
}

// UpdateLbdcArgs defines the structure of input parameters for the UpdateLbdc api
type UpdateLbdcArgs struct {
	ClientToken    string          `json:"-"`
	Id             string          `json:"id"`
	UpdateLbdcBody *UpdateLbdcBody `json:"updateLbdcBody"`
}

// UpdateLbdcBody defines the structure of input parameters for the UpdateLbdc api request body
type UpdateLbdcBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"desc,omitempty"`
}

type AssociateBlb struct {
	BlbId        string `json:"blbId"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	BlbType      string `json:"blbType"`
	PublicIp     string `json:"publicIp"`
	EipRouteType string `json:"eipRouteType"`
	Bandwidth    int    `json:"bandwidth"`
	Address      string `json:"address"`
	Ipv6         string `json:"ipv6"`
	VpcId        string `json:"vpcId"`
	SubnetId     string `json:"subnetId"`
}

// GetBoundBlBListOfLbdcResult defines the structure of output parameters for the GetBoundBlBListOfLbdc api
type GetBoundBlBListOfLbdcResult struct {
	BlbList []AssociateBlb `json:"blbList"`
}
