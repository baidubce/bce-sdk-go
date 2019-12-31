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

package api

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type InstanceStatus string

const (
	InstanceStatusRunning            InstanceStatus = "Running"
	InstanceStatusStarting           InstanceStatus = "Starting"
	InstanceStatusStopping           InstanceStatus = "Stopping"
	InstanceStatusStopped            InstanceStatus = "Stopped"
	InstanceStatusDeleted            InstanceStatus = "Deleted"
	InstanceStatusScaling            InstanceStatus = "Scaling"
	InstanceStatusExpired            InstanceStatus = "Expired"
	InstanceStatusError              InstanceStatus = "Error"
	InstanceStatusSnapshotProcessing InstanceStatus = "SnapshotProcessing"
	InstanceStatusImageProcessing    InstanceStatus = "ImageProcessing"
)

type PaymentTimingType string

const (
	PaymentTimingPrePaid  PaymentTimingType = "Prepaid"
	PaymentTimingPostPaid PaymentTimingType = "Postpaid"
)

// Instance define instance model
type InstanceModel struct {
	InstanceId            string           `json:"id"`
	InstanceName          string           `json:"name"`
	Status                InstanceStatus   `json:"status"`
	Description           string           `json:"desc"`
	PaymentTiming         string           `json:"paymentTiming"`
	CreationTime          string           `json:"createTime"`
	ExpireTime            string           `json:"expireTime"`
	PublicIP              string           `json:"publicIp"`
	InternalIP            string           `json:"internalIp"`
	FlavorId              string           `json:"flavorId"`
	ImageId               string           `json:"imageId"`
	NetworkCapacityInMbps int              `json:"networkCapacityInMbps"`
	ZoneName              string           `json:"zone"`
	Region                string           `json:"region"`
	Tags                  []model.TagModel `json:"tags"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"` // [1,2,3,4,5,6,7,8,9,12,24,36]
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

type CreateInstanceArgs struct {
	FlavorId         string  `json:"flavorId,omitempty"`
	ImageId          string  `json:"imageId,omitempty"`
	RaidId           string  `json:"raidId,omitempty"`
	RootDiskSizeInGb int     `json:"rootDiskSizeInGb,omitempty"`
	PurchaseCount    int     `json:"purchaseCount,omitempty"`
	ZoneName         string  `json:"zoneName,omitempty"`
	SubnetId         string  `json:"subnetId,omitempty"`
	Billing          Billing `json:"billing,omitempty"`
	Name             string  `json:"name"`
	AdminPass        string  `json:"adminPass"`
	ClientToken      string  `json:"-"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
}

// all is query parms
type ListInstanceArgs struct {
	Marker     string
	MaxKeys    int
	InternalIp string
}

type ListInstanceResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type GetInstanceDetailResult struct {
	Instance InstanceModel `json:"instance"`
}
