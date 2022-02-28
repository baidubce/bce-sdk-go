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

package cfs

import ()

type FSStatus string

const (
	FSStatusAvailable   FSStatus = "available"
	FSStatusPaused      FSStatus = "paused"
	FSStatusUnavailable FSStatus = "unavailable"
)

type CreateFSArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"fsName"`
	Type        string `json:"type"`
	Protocol    string `json:"protocol"`
	VpcID       string `json:"vpcId,omitempty"`
	Zone        string `json:"zone,omitempty"`
}

type CreateFSResult struct {
	FSID string `json:"fsId"`
}

type UpdateFSArgs struct {
	FSID   string `json:"-"`
	FSName string `json:"fsName"`
}

type DescribeResultMeta struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

type DescribeFSArgs struct {
	FSID    string
	UserId  string
	Marker  string
	MaxKeys int
}

type MoutTargetModel struct {
	AccessGroupName string `json:"accessGroupName"`
	MountID         string `json:"mountId"`
	Ovip            string `json:"ovip"`
	Domain          string `json:"domain"`
	SubnetID        string `json:"subnetId"`
}

type FSModel struct {
	FSID        string            `json:"fsId"`
	Name        string            `json:"fsName"`
	VpcID       string            `json:"vpcId"`
	Type        string            `json:"type"`
	Protocol    string            `json:"protocol"`
	Usage       string            `json:"fsUsage"`
	Status      FSStatus          `json:"status"`
	MoutTargets []MoutTargetModel `json:"MountTargetList"`
}

type DescribeFSResult struct {
	FSList []FSModel `json:"FileSystemList"`
	DescribeResultMeta
}

type DropFSArgs struct {
	FSID string
}

type CreateMountTargetArgs struct {
	FSID            string `json:"-"`
	VpcID           string `json:"vpcId"`
	SubnetId        string `json:"subnetId"`
	AccessGroupName string `json:"accessGroupName"`
}

type CreateMountTargetResult struct {
	MountID string `json:"mountId"`
	Domain  string `json:"domain"`
}

type DropMountTargetArgs struct {
	FSID    string
	MountId string
}

type DescribeMountTargetArgs struct {
	FSID    string
	MountID string
	Marker  string
	MaxKeys int
}

type DescribeMountTargetResult struct {
	MountTargetList []MoutTargetModel `json:"MountTargetList"`
	DescribeResultMeta
}
