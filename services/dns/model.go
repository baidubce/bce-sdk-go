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

package dns

type AddLineGroupRequest struct {
	Name  string   `json:"name"`
	Lines []string `json:"lines"`
}

type AddLineGroupRequestLines struct {
}

type Billing struct {
	PaymentTiming string      `json:"paymentTiming"`
	Reservation   Reservation `json:"reservation"`
}

type CreatePaidZoneRequest struct {
	Names          []string `json:"names"`
	ProductVersion string   `json:"productVersion"`
	Billing        Billing  `json:"billing"`
}

type CreateRecordRequest struct {
	Rr          string  `json:"rr"`
	Type        string  `json:"type"`
	Value       string  `json:"value"`
	Ttl         *int32  `json:"ttl"`
	Line        *string `json:"line"`
	Description *string `json:"description"`
	Priority    *int32  `json:"priority"`
}

type CreateZoneRequest struct {
	Name string `json:"name"`
}

type DeleteRecordRequest struct {
	Rr          string  `json:"rr"`
	Type        string  `json:"type"`
	Value       string  `json:"value"`
	Ttl         *int32  `json:"ttl"`
	Description *string `json:"description"`
	Priority    *int32  `json:"priority"`
}

type DeleteZoneRequest struct {
	Names          []string `json:"names"`
	ProductVersion string   `json:"productVersion"`
	Billing        Billing  `json:"billing"`
}

type DeleteZoneRequestNames struct {
}

type Line struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Lines              []string `json:"lines"`
	RelatedZoneCount   int32    `json:"relatedZoneCount"`
	RelatedRecordCount int32    `json:"relatedRecordCount"`
}

type ListLineGroupRequest struct {
	Marker  string
	MaxKeys int
}

type ListLineGroupResponse struct {
	Marker      *string `json:"marker"`
	IsTruncated *bool   `json:"isTruncated"`
	NextMarker  *string `json:"nextMarker"`
	MaxKeys     *int32  `json:"maxKeys"`
	LineList    []Line  `json:"lineList"`
}

type ListRecordRequest struct {
	Rr      string
	Id      string
	Marker  string
	MaxKeys int
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

type ListZoneRequest struct {
	Name    string
	Marker  string
	MaxKeys int
}

type ListZoneResponse struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int32  `json:"maxKeys"`
	Zones       []Zone `json:"zones"`
}

type ListZoneResponseZones struct {
}

type Record struct {
	Id          string `json:"id"`
	Rr          string `json:"rr"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Ttl         int32  `json:"ttl"`
	Line        string `json:"line"`
	Description string `json:"description"`
	Priority    int32  `json:"priority"`
}

type RenewZoneRequest struct {
	Billing Billing `json:"billing"`
}

type Reservation struct {
	ReservationLength int32 `json:"reservationLength"`
}

type TagModel struct {
	TagKey   *string `json:"tagKey"`
	TagValue *string `json:"tagValue"`
}

type UpdateLineGroupRequest struct {
	Name  string   `json:"name"`
	Lines []string `json:"lines"`
}

type UpdateLineGroupRequestLines struct {
}

type UpdateRecordDisableRequest struct {
	Rr          string  `json:"rr"`
	Type        string  `json:"type"`
	Value       string  `json:"value"`
	Ttl         *int32  `json:"ttl"`
	Description *string `json:"description"`
	Priority    *int32  `json:"priority"`
}

type UpdateRecordEnableRequest struct {
	Rr          string  `json:"rr"`
	Type        string  `json:"type"`
	Value       string  `json:"value"`
	Ttl         *int32  `json:"ttl"`
	Description *string `json:"description"`
	Priority    *int32  `json:"priority"`
}

type UpdateRecordRequest struct {
	Rr          string  `json:"rr"`
	Type        string  `json:"type"`
	Value       string  `json:"value"`
	Ttl         *int32  `json:"ttl"`
	Description *string `json:"description"`
	Priority    *int32  `json:"priority"`
}

type UpgradeZoneRequest struct {
	Names   []string `json:"names"`
	Billing Billing  `json:"billing"`
}

type UpgradeZoneRequestNames struct {
}

type Zone struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	Status         string     `json:"status"`
	ProductVersion string     `json:"productVersion"`
	CreateTime     string     `json:"createTime"`
	ExpireTime     string     `json:"expireTime"`
	Tags           []TagModel `json:"tags"`
}
