/*
 * Copyright  Baidu, Inc.
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

package localDns

type AddRecordRequest struct {
	ClientToken string `json:"-"`
	Rr          string `json:"rr"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Ttl         int32  `json:"ttl,omitempty"`
	Priority    int32  `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
}

type AddRecordResponse struct {
	RecordId string `json:"recordId"`
}

type BindVpcRequest struct {
	ClientToken string   `json:"-"`
	Region      string   `json:"region"`
	VpcIds      []string `json:"vpcIds"`
}

type BindVpcRequestVpcIds struct {
}

type ListPrivateZoneRequest struct {
	Marker  string
	MaxKeys int
}

type CreatePrivateZoneRequest struct {
	ClientToken string `json:"-"`
	ZoneName    string `json:"zoneName"`
}

type CreatePrivateZoneResponse struct {
	ZoneId string `json:"zoneId"`
}

type DeletePrivateZoneRequest struct {
	ZoneName string `json:"zoneName"`
}

type DeleteRecordRequest struct {
	RecordId    string  `json:"recordId"`
	Rr          string  `json:"rr"`
	Value       string  `json:"value"`
	Type        string  `json:"type"`
	Ttl         *int32  `json:"ttl"`
	Priority    *int32  `json:"priority"`
	Description *string `json:"description"`
}

type GetPrivateZoneResponse struct {
	ZoneId      string `json:"zoneId"`
	ZoneName    string `json:"zoneName"`
	RecordCount int32  `json:"recordCount"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	BindVpcs    []Vpc  `json:"bindVpcs"`
}

type GetPrivateZoneResponseBindVpcs struct {
}

type ListPrivateZoneResponse struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int32  `json:"maxKeys"`
	Zones       []Zone `json:"zones"`
}

type ListPrivateZoneResponseZones struct {
}

type ListRecordResponse struct {
	Marker      string   `json:"marker"`
	IsTruncated bool     `json:"isTruncated"`
	NextMarker  string   `json:"nextMarker"`
	MaxKeys     int32    `json:"maxKeys"`
	Records     []Record `json:"records"`
}

type ListRecordResponseRecords struct {
}

type Record struct {
	RecordId    string `json:"recordId"`
	Rr          string `json:"rr"`
	Value       string `json:"value"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Ttl         int32  `json:"ttl"`
	Priority    int32  `json:"priority"`
	Description string `json:"description"`
}

type UnbindVpcRequest struct {
	ClientToken string   `json:"-"`
	Region      string   `json:"region"`
	VpcIds      []string `json:"vpcIds"`
}

type UnbindVpcRequestVpcIds struct {
}

type UpdateRecordRequest struct {
	ClientToken string `json:"-"`
	Rr          string `json:"rr"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Ttl         int32  `json:"ttl,omitempty"`
	Priority    int32  `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
}

type Vpc struct {
	VpcId     string `json:"vpcId"`
	VpcName   string `json:"vpcName"`
	VpcRegion string `json:"vpcRegion"`
}

type Zone struct {
	ZoneId      string `json:"zoneId"`
	ZoneName    string `json:"zoneName"`
	RecordCount int32  `json:"recordCount"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}
