/*
 * Copyright 2017 Baidu, Inc.
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
package eip

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type Reservation struct {
	ReservationLength   int    `json:"reservationLength,omitempty"`
	ReservationTimeUnit string `json:"reservationTimeUnit,omitempty"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming,omitempty"`
	BillingMethod string       `json:"billingMethod,omitempty"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}

type CreateEipArgs struct {
	Name              string           `json:"name,omitempty"`
	BandWidthInMbps   int              `json:"bandwidthInMbps"`
	Billing           *Billing         `json:"billing"`
	Tags              []model.TagModel `json:"tags"`
	AutoRenewTimeUnit string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int              `json:"autoRenewTime,omitempty"`
	RouteType         string           `json:"routeType,omitempty"`
	Idc               string           `json:"idc,omitempty"`
	ClientToken       string           `json:"-"`
}

type BatchCreateEipArgs struct {
	Name              string           `json:"name,omitempty"`
	BandWidthInMbps   int              `json:"bandwidthInMbps"`
	Billing           *Billing         `json:"billing"`
	Tags              []model.TagModel `json:"tags"`
	AutoRenewTimeUnit string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int              `json:"autoRenewTime,omitempty"`
	RouteType         string           `json:"routeType,omitempty"`
	Idc               string           `json:"idc,omitempty"`
	Continuous        bool             `json:"continuous,omitempty"`
	Count             int              `json:"count,omitempty"`
	ClientToken       string           `json:"-"`
}

type CreateEipResult struct {
	Eip string `json:"eip"`
}

type BatchCreateEipResult struct {
	Eips []string `json:"eips"`
}

type ResizeEipArgs struct {
	NewBandWidthInMbps int    `json:"newBandwidthInMbps"`
	ClientToken        string `json:"-"`
}

type BindEipArgs struct {
	InstanceType string `json:"instanceType"`
	InstanceId   string `json:"instanceId"`
	ClientToken  string `json:"-"`
}

type ListEipArgs struct {
	Eip          string
	InstanceType string
	InstanceId   string
	Marker       string
	MaxKeys      int
	Status       string
}

type ListEipResult struct {
	Marker      string     `json:"marker"`
	MaxKeys     int        `json:"maxKeys"`
	NextMarker  string     `json:"nextMarker"`
	IsTruncated bool       `json:"isTruncated"`
	EipList     []EipModel `json:"eipList"`
}

type EipModel struct {
	Name            string           `json:"name"`
	Eip             string           `json:"eip"`
	EipId           string           `json:"eipId"`
	Status          string           `json:"status"`
	EipInstanceType string           `json:"eipInstanceType"`
	InstanceType    string           `json:"instanceType"`
	InstanceId      string           `json:"instanceId"`
	ShareGroupId    string           `json:"shareGroupId"`
	ClusterId       string           `json:"clusterId"`
	BandWidthInMbps int              `json:"bandwidthInMbps"`
	PaymentTiming   string           `json:"paymentTiming"`
	BillingMethod   string           `json:"billingMethod"`
	CreateTime      string           `json:"createTime"`
	ExpireTime      string           `json:"expireTime"`
	Tags            []model.TagModel `json:"tags"`
}

type ListRecycleEipArgs struct {
	Eip     string
	Name    string
	Marker  string
	MaxKeys int
}

type ListRecycleEipResult struct {
	Marker      string            `json:"marker"`
	MaxKeys     int               `json:"maxKeys"`
	NextMarker  string            `json:"nextMarker"`
	IsTruncated bool              `json:"isTruncated"`
	EipList     []RecycleEipModel `json:"eipList"`
}

type RecycleEipModel struct {
	Name                string `json:"name"`
	Eip                 string `json:"eip"`
	EipId               string `json:"eipId"`
	Status              string `json:"status"`
	RouteType           string `json:"routeType"`
	BandWidthInMbps     int    `json:"bandwidthInMbps"`
	PaymentTiming       string `json:"paymentTiming"`
	BillingMethod       string `json:"billingMethod"`
	RecycleTime         string `json:"recycleTime"`
	ScheduledDeleteTime string `json:"scheduledDeleteTime"`
}

type PurchaseReservedEipArgs struct {
	Billing     *Billing `json:"billing"`
	ClientToken string   `json:"clientToken"`
}

type StartAutoRenewArgs struct {
	AutoRenewTimeUnit string `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int    `json:"autoRenewTime,omitempty"`
	ClientToken       string `json:"-"`
}

type ListClusterResult struct {
	Marker      string         `json:"marker"`
	MaxKeys     int            `json:"maxKeys"`
	NextMarker  string         `json:"nextMarker"`
	IsTruncated bool           `json:"isTruncated"`
	ClusterList []ClusterModel `json:"clusterList"`
}

type ClusterModel struct {
	ClusterId     string `json:"clusterId"`
	ClusterName   string `json:"clusterName"`
	ClusterRegion string `json:"clusterRegion"`
	ClusterAz     string `json:"clusterAz"`
}

type ClusterDetail struct {
	ClusterId     string `json:"clusterId"`
	ClusterName   string `json:"clusterName"`
	ClusterRegion string `json:"clusterRegion"`
	ClusterAz     string `json:"clusterAz"`
	NetworkInBps  int64  `json:"networkInBps"`
	NetworkOutBps int64  `json:"networkOutBps"`
	NetworkInPps  int64  `json:"networkInPps"`
	NetworkOutPps int64  `json:"networkOutPps"`
}

type Package struct {
	Id           string `json:"id,omitempty"`
	DeductPolicy string `json:"deductPolicy,omitempty"`
	PackageType  string `json:"packageType,omitempty"`
	Status       string `json:"status,omitempty"`
	Capacity     string `json:"capacity,omitempty"`
	UsedCapacity string `json:"usedCapacity,omitempty"`
	ActiveTime   string `json:"activeTime"`
	ExpireTime   string `json:"expireTime"`
	CreateTime   string `json:"createTime"`
}

type CreateEipTpArgs struct {
	ReservationLength int    `json:"reservationLength,omitempty"`
	Capacity          string `json:"capacity,omitempty"`
	DeductPolicy      string `json:"deductPolicy,omitempty"`
	PackageType       string `json:"packageType,omitempty"`
	ClientToken       string `json:"-"`
}

type CreateEipTpResult struct {
	Id string `json:"id,omitempty"`
}

type ListEipTpArgs struct {
	Id           string `json:"id,omitempty"`
	DeductPolicy string `json:"deductPolicy,omitempty"`
	Status       string `json:"status,omitempty"`
	Marker       string `json:"marker"`
	MaxKeys      int    `json:"maxKeys"`
}

type ListEipTpResult struct {
	Marker      string    `json:"marker"`
	MaxKeys     int       `json:"maxKeys"`
	NextMarker  string    `json:"nextMarker"`
	IsTruncated bool      `json:"isTruncated"`
	PackageList []Package `json:"packageList"`
}

type EipTpDetail struct {
	Id           string `json:"id,omitempty"`
	DeductPolicy string `json:"deductPolicy,omitempty"`
	PackageType  string `json:"packageType,omitempty"`
	Status       string `json:"status,omitempty"`
	Capacity     string `json:"capacity,omitempty"`
	UsedCapacity string `json:"usedCapacity,omitempty"`
	ActiveTime   string `json:"activeTime,omitempty"`
	ExpireTime   string `json:"expireTime,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
}

type CreateEipGroupArgs struct {
	Name            string           `json:"name,omitempty"`
	EipCount        int              `json:"eipCount"`
	BandWidthInMbps int              `json:"bandwidthInMbps"`
	Billing         *Billing         `json:"billing"`
	Tags            []model.TagModel `json:"tags"`
	RouteType       string           `json:"routeType,omitempty"`
	Idc             string           `json:"idc,omitempty"`
	Continuous      bool             `json:"continuous,omitempty"`
	ClientToken     string           `json:"-"`
}

type CreateEipGroupResult struct {
	Id string `json:"id"`
}

type ResizeEipGroupArgs struct {
	BandWidthInMbps int    `json:"bandwidthInMbps"`
	ClientToken     string `json:"-"`
}

type GroupAddEipCountArgs struct {
	EipAddCount int    `json:"eipAddCount"`
	ClientToken string `json:"-"`
}

type ReleaseEipGroupIpsArgs struct {
	ReleaseIps  []string `json:"releaseIps"`
	ClientToken string   `json:"-"`
}

type RenameEipGroupArgs struct {
	Name        string `json:"name"`
	ClientToken string `json:"-"`
}

type ListEipGroupArgs struct {
	Id      string
	Name    string
	Marker  string
	MaxKeys int
	Status  string
}

type ListEipGroupResult struct {
	Marker      string          `json:"marker"`
	MaxKeys     int             `json:"maxKeys"`
	NextMarker  string          `json:"nextMarker"`
	IsTruncated bool            `json:"isTruncated"`
	EipGroup    []EipGroupModel `json:"eipgroups"`
}

type EipGroupModel struct {
	Name                      string     `json:"name"`
	Status                    string     `json:"status"`
	Id                        string     `json:"id"`
	BandWidthInMbps           int        `json:"bandwidthInMbps"`
	DefaultDomesticBandwidth  int        `json:"defaultDomesticBandwidth"`
	BwShortId                 string     `json:"bwShortId"`
	BwBandwidthInMbps         int        `json:"bwBandwidthInMbps"`
	DomesticBwShortId         string     `json:"domesticBwShortId"`
	DomesticBwBandwidthInMbps int        `json:"domesticBwBandwidthInMbps"`
	PaymentTiming             string     `json:"paymentTiming"`
	BillingMethod             string     `json:"billingMethod"`
	CreateTime                string     `json:"createTime"`
	ExpireTime                string     `json:"expireTime"`
	Region                    string     `json:"region"`
	RouteType                 string     `json:"routeType"`
	Eips                      []EipModel `json:"eips"`
}

type EipGroupMoveOutArgs struct {
	MoveOutEips []MoveOutEip `json:"moveOutEips"`
	ClientToken string       `json:"-"`
}

type MoveOutEip struct {
	Eip             string   `json:"eip"`
	BandWidthInMbps int      `json:"bandwidthInMbps"`
	Billing         *Billing `json:"billing"`
}

type EipGroupMoveInArgs struct {
	Eips        []string `json:"eips"`
	ClientToken string   `json:"-"`
}
