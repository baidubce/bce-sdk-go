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

package csn

type AttachInstanceRequest struct {
	InstanceType      string  `json:"instanceType"`
	InstanceId        string  `json:"instanceId"`
	InstanceRegion    string  `json:"instanceRegion"`
	InstanceAccountId *string `json:"instanceAccountId,omitempty"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}

type BindCsnBpRequest struct {
	CsnId string `json:"csnId"`
}

type CreateAssociationRequest struct {
	AttachId    string  `json:"attachId"`
	Description *string `json:"description,omitempty"`
}

type CreateCsnBpLimitRequest struct {
	LocalRegion string `json:"localRegion"`
	PeerRegion  string `json:"peerRegion"`
	Bandwidth   int32  `json:"bandwidth"`
}

type CreateCsnBpRequest struct {
	Name          string  `json:"name"`
	InterworkType *string `json:"interworkType,omitempty"`
	Bandwidth     int32   `json:"bandwidth"`
	GeographicA   string  `json:"geographicA"`
	GeographicB   string  `json:"geographicB"`
	Billing       Billing `json:"billing"`
}

type CreateCsnBpResponse struct {
	CsnBpId string `json:"csnBpId"`
}

type CreateCsnRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type CreateCsnResponse struct {
	CsnId string `json:"csnId"`
}

type CreatePropagationRequest struct {
	AttachId    string  `json:"attachId"`
	Description *string `json:"description,omitempty"`
}

type CreateRouteRuleRequest struct {
	AttachId    string `json:"attachId"`
	DestAddress string `json:"destAddress"`
	RouteType   string `json:"routeType"`
}

type Csn struct {
	CsnId       string `json:"csnId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	InstanceNum int32  `json:"instanceNum"`
	CsnBpNum    int32  `json:"csnBpNum"`
}

type CsnBp struct {
	CsnBpId         string `json:"csnBpId"`
	Name            string `json:"name"`
	Bandwidth       int32  `json:"bandwidth"`
	UsedBandwidth   int32  `json:"usedBandwidth"`
	CsnId           string `json:"csnId"`
	InterworkType   string `json:"interworkType"`
	InterworkRegion string `json:"interworkRegion"`
	Status          string `json:"status"`
	PaymentTiming   string `json:"paymentTiming"`
	ExpireTime      string `json:"expireTime"`
	CreatedTime     string `json:"createdTime"`
}

type CsnBpLimit struct {
	CsnBpId     string `json:"csnBpId"`
	CsnId       string `json:"csnId"`
	LocalRegion string `json:"localRegion"`
	PeerRegion  string `json:"peerRegion"`
	Bandwidth   int32  `json:"bandwidth"`
}

type CsnRouteTable struct {
	CsnRtId     string `json:"csnRtId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type CsnRtAssociation struct {
	AttachId       string `json:"attachId"`
	Description    string `json:"description"`
	InstanceId     string `json:"instanceId"`
	InstanceName   string `json:"instanceName"`
	InstanceRegion string `json:"instanceRegion"`
	InstanceType   string `json:"instanceType"`
	Status         string `json:"status"`
}

type CsnRtPropagation struct {
	AttachId       string `json:"attachId"`
	Description    string `json:"description"`
	InstanceId     string `json:"instanceId"`
	InstanceName   string `json:"instanceName"`
	InstanceRegion string `json:"instanceRegion"`
	InstanceType   string `json:"instanceType"`
	Status         string `json:"status"`
}

type CsnRtRule struct {
	RuleId        string `json:"ruleId"`
	RouteType     string `json:"routeType"`
	CsnId         string `json:"csnId"`
	CsnRtId       string `json:"csnRtId"`
	Description   string `json:"description"`
	FromAttachId  string `json:"fromAttachId"`
	Status        string `json:"status"`
	SourceAddress string `json:"sourceAddress"`
	DestAddress   string `json:"destAddress"`
	NextHopId     string `json:"nextHopId"`
	NextHopName   string `json:"nextHopName"`
	NextHopRegion string `json:"nextHopRegion"`
	NextHopType   string `json:"nextHopType"`
	AsPath        string `json:"asPath"`
	Community     string `json:"community"`
	BlackHole     bool   `json:"blackHole"`
}

type DeleteCsnBpLimitRequest struct {
	LocalRegion string `json:"localRegion"`
	PeerRegion  string `json:"peerRegion"`
}

type DetachInstanceRequest struct {
	InstanceType      string  `json:"instanceType"`
	InstanceId        string  `json:"instanceId"`
	InstanceRegion    string  `json:"instanceRegion"`
	InstanceAccountId *string `json:"instanceAccountId,omitempty"`
}

type GetCsnResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CsnId       string `json:"csnId"`
	Status      string `json:"status"`
	InstanceNum int32  `json:"instanceNum"`
	CsnBpNum    int32  `json:"csnBpNum"`
}

type GetCsnBpResponse struct {
	CsnBpId         string `json:"csnBpId"`
	Name            string `json:"name"`
	Bandwidth       int32  `json:"bandwidth"`
	UsedBandwidth   int32  `json:"usedBandwidth"`
	CsnId           string `json:"csnId"`
	InterworkType   string `json:"interworkType"`
	InterworkRegion string `json:"interworkRegion"`
	Status          string `json:"status"`
	PaymentTiming   string `json:"paymentTiming"`
	ExpireTime      string `json:"expireTime"`
	CreatedTime     string `json:"createdTime"`
}

type Instance struct {
	AttachId          string `json:"attachId"`
	InstanceType      string `json:"instanceType"`
	InstanceId        string `json:"instanceId"`
	InstanceName      string `json:"instanceName"`
	InstanceRegion    string `json:"instanceRegion"`
	InstanceAccountId string `json:"instanceAccountId"`
	Status            string `json:"status"`
}

type ListAssociationResponse struct {
	Associations []CsnRtAssociation `json:"associations"`
}

type ListAssociationResponseAssociations struct {
}

type ListCsnBpLimitByCsnIdRequest struct {
	LocalRegion string `json:"localRegion"`
	PeerRegion  string `json:"peerRegion"`
}

type ListCsnBpLimitByCsnIdResponse struct {
	BpLimits []CsnBpLimit `json:"bpLimits"`
}

type ListCsnBpLimitByCsnIdResponseBpLimits struct {
}

type ListCsnBpLimitResponse struct {
	BpLimits []CsnBpLimit `json:"bpLimits"`
}

type ListCsnBpLimitResponseBpLimits struct {
}

type ListCsnBpResponse struct {
	CsnBps      []CsnBp `json:"csnBps"`
	Marker      *string `json:"marker,omitempty"`
	IsTruncated bool    `json:"isTruncated"`
	NextMarker  *string `json:"nextMarker,omitempty"`
	MaxKeys     int32   `json:"maxKeys"`
}

type ListCsnBpResponseCsnBps struct {
}

type ListCsnResponse struct {
	Csns        []Csn   `json:"csns"`
	Marker      *string `json:"marker,omitempty"`
	IsTruncated bool    `json:"isTruncated"`
	NextMarker  *string `json:"nextMarker,omitempty"`
	MaxKeys     int32   `json:"maxKeys"`
}

type ListCsnResponseCsns struct {
}

type ListInstanceResponse struct {
	Instances   []Instance `json:"instances"`
	Marker      *string    `json:"marker,omitempty"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  *string    `json:"nextMarker,omitempty"`
	MaxKeys     int32      `json:"maxKeys"`
}

type ListInstanceResponseInstances struct {
}

type ListPropagationResponse struct {
	Propagations []CsnRtPropagation `json:"propagations"`
}

type ListPropagationResponsePropagations struct {
}

type ListRouteRuleResponse struct {
	CsnRtRules  []CsnRtRule `json:"csnRtRules"`
	Marker      *string     `json:"marker,omitempty"`
	IsTruncated bool        `json:"isTruncated"`
	NextMarker  *string     `json:"nextMarker,omitempty"`
	MaxKeys     int32       `json:"maxKeys"`
}

type ListRouteRuleResponseCsnRtRules struct {
}

type ListRouteTableResponse struct {
	CsnRts      []CsnRouteTable `json:"csnRts"`
	Marker      *string         `json:"marker,omitempty"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  *string         `json:"nextMarker,omitempty"`
	MaxKeys     int32           `json:"maxKeys"`
}

type ListRouteTableResponseCsnRts struct {
}

type ListTgwResponse struct {
	Tgws        []Tgw   `json:"tgws"`
	Marker      *string `json:"marker,omitempty"`
	IsTruncated bool    `json:"isTruncated"`
	NextMarker  *string `json:"nextMarker,omitempty"`
	MaxKeys     int32   `json:"maxKeys"`
}

type ListTgwRuleResponse struct {
	TgwRtRules  []TgwRtRule `json:"tgwRtRules"`
	Marker      *string     `json:"marker,omitempty"`
	IsTruncated bool        `json:"isTruncated"`
	NextMarker  *string     `json:"nextMarker,omitempty"`
	MaxKeys     int32       `json:"maxKeys"`
}

type ListTgwRuleResponseTgwRtRules struct {
}

type Reservation struct {
	ReservationLength   int32  `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type ResizeCsnBpRequest struct {
	Bandwidth int32 `json:"bandwidth"`
}

type Tgw struct {
	TgwId       string `json:"tgwId"`
	CsnId       string `json:"csnId"`
	Name        string `json:"name"`
	Region      string `json:"region"`
	Description string `json:"description"`
}

type TgwRtRule struct {
	RuleId        string `json:"ruleId"`
	RouteType     string `json:"routeType"`
	CsnId         string `json:"csnId"`
	CsnRtId       string `json:"csnRtId"`
	FromAttachId  string `json:"fromAttachId"`
	Status        string `json:"status"`
	DestAddress   string `json:"destAddress"`
	NextHopId     string `json:"nextHopId"`
	NextHopName   string `json:"nextHopName"`
	NextHopRegion string `json:"nextHopRegion"`
	NextHopType   string `json:"nextHopType"`
	AsPath        string `json:"asPath"`
	Community     string `json:"community"`
	BlackHole     bool   `json:"blackHole"`
}

type UnbindCsnBpRequest struct {
	CsnId string `json:"csnId"`
}

type UpdateCsnBpLimitRequest struct {
	LocalRegion string `json:"localRegion"`
	PeerRegion  string `json:"peerRegion"`
	Bandwidth   int32  `json:"bandwidth"`
}

type UpdateCsnBpRequest struct {
	Name string `json:"name"`
}

type UpdateCsnRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateTgwRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ListCsnArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListCsnBpArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListInstanceArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListRouteRuleArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListRouteTableArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListTgwArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListTgwRuleArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}
