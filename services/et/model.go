/*
 * Copyright 2023 Baidu, Inc.
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

package et

type GetEtChannelArgs struct {
	ClientToken string `json:"clientToken,omitempty"`
	EtId        string `json:"etId"`
}

type RecommitEtChannelArgs struct {
	ClientToken string                  `json:"clientToken,omitempty"`
	EtId        string                  `json:"etId"`
	EtChannelId string                  `json:"etChannelId"`
	Result      RecommitEtChannelResult `json:"etChannelResult"`
}

type UpdateEtChannelArgs struct {
	ClientToken string                `json:"clientToken,omitempty"`
	EtId        string                `json:"etId"`
	EtChannelId string                `json:"etChannelId"`
	Result      UpdateEtChannelResult `json:"UpdateEtChannelResult"`
}

type DeleteEtChannelArgs struct {
	ClientToken string `json:"clientToken,omitempty"`
	EtId        string `json:"etId"`
	EtChannelId string `json:"etChannelId"`
}

type EnableEtChannelIPv6Args struct {
	ClientToken string                    `json:"clientToken,omitempty"`
	EtId        string                    `json:"etId"`
	EtChannelId string                    `json:"etChannelId"`
	Result      EnableEtChannelIPv6Result `json:"enableEtChannelIpv6Result"`
}

type GetEtChannelsResult struct {
	EtChannels []EtChannelResult `json:"etChannels"`
}

type RecommitEtChannelResult struct {
	AuthorizedUsers     []string `json:"authorizedUsers"`
	Description         string   `json:"description"`
	BaiduAddress        string   `json:"baiduAddress"`
	Name                string   `json:"name"`
	Networks            []string `json:"networks"`
	CustomerAddress     string   `json:"customerAddress"`
	RouteType           string   `json:"routeType"`
	VlanId              string   `json:"vlanId"`
	Id                  string   `json:"id"`
	Status              string   `json:"status"`
	EnableIpv6          uint32   `json:"enableIpv6"`
	BaiduIpv6Address    string   `json:"baiduIpv6Address"`
	Ipv6Networks        []string `json:"ipv6Networks"`
	CustomerIpv6Address string   `json:"CustomerIpv6Address"`
}

type UpdateEtChannelResult struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EnableEtChannelIPv6Result struct {
	BaiduIpv6Address    string   `json:"baiduIpv6Address"`
	CustomerIpv6Address string   `json:"CustomerIpv6Address"`
	Ipv6Networks        []string `json:"ipv6Networks"`
}

type EtChannelResult struct {
	AuthorizedUsers     []string `json:"authorizedUsers"`
	Description         string   `json:"description"`
	BaiduAddress        string   `json:"baiduAddress"`
	Name                string   `json:"name"`
	Networks            []string `json:"networks"`
	BGPAsn              string   `json:"bgpAsn"`
	BGPKey              string   `json:"bgpKey"`
	BGPStatus           string   `json:"bgpStatus"`
	Ipv6BGPStatus       string   `json:"ipv6BgpStatus"`
	CustomerAddress     string   `json:"customerAddress"`
	RouteType           string   `json:"routeType"`
	VlanId              string   `json:"vlanId"`
	Id                  string   `json:"id"`
	Status              string   `json:"status"`
	EnableIpv6          uint32   `json:"enableIpv6"`
	BaiduIpv6Address    string   `json:"baiduIpv6Address"`
	Ipv6Networks        []string `json:"ipv6Networks"`
	CustomerIpv6Address string   `json:"CustomerIpv6Address"`
	Tags                []Tag    `json:"tags"`
}

type CreateEtDcphyArgs struct {
	ClientToken string `json:"clientToken,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Isp         string `json:"isp"`
	IntfType    string `json:"intfType"`
	ApType      string `json:"apType"`
	ApAddr      string `json:"apAddr"`
	UserName    string `json:"userName"`
	UserPhone   string `json:"userPhone"`
	UserEmail   string `json:"userEmail"`
	UserIdc     string `json:"userIdc"`
	Tags        []Tag  `json:"tags"`
}

type CreateEtDcphyResult struct {
	Id string `json:"id"`
}

type UpdateEtDcphyArgs struct {
	ClientToken string `json:"clientToken,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UserName    string `json:"userName,omitempty"`
	UserPhone   string `json:"userPhone,omitempty"`
	UserEmail   string `json:"userEmail,omitempty"`
}

type ListEtDcphyArgs struct {
	Marker  string
	MaxKeys int
	Status  string
}

type ListEtDcphyResult struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
	Ets         []Et   `json:"ets"`
}

type Et struct {
	Id          string `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ExpireTime  string `json:"expireTime"`
	Isp         string `json:"isp"`
	IntfType    string `json:"intfType"`
	ApType      string `json:"apType"`
	ApAddr      string `json:"apAddr"`
	UserName    string `json:"userName"`
	UserPhone   string `json:"userPhone"`
	UserEmail   string `json:"userEmail"`
	UserIdc     string `json:"userIdc"`
	Tags        []Tag  `json:"tags"`
}

type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

type EtDcphyDetail struct {
	Id          string `json:"clientToken,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ExpireTime  string `json:"expireTime"`
	Isp         string `json:"isp"`
	IntfType    string `json:"intfType"`
	ApType      string `json:"apType"`
	ApAddr      string `json:"apAddr"`
	UserName    string `json:"userName"`
	UserPhone   string `json:"userPhone"`
	UserEmail   string `json:"userEmail"`
	UserIdc     string `json:"userIdc"`
	Tags        []Tag  `json:"tags"`
}

type CreateEtChannelArgs struct {
	ClientToken         string   `json:"clientToken,omitempty"`
	EtId                string   `json:"etId"`
	AuthorizedUsers     []string `json:"authorizedUsers,omitempty"`
	Description         string   `json:"description,omitempty"`
	BaiduAddress        string   `json:"baiduAddress"`
	Name                string   `json:"name"`
	Networks            []string `json:"networks,omitempty"`
	CustomerAddress     string   `json:"customerAddress"`
	RouteType           string   `json:"routeType"`
	VlanId              int      `json:"vlanId"`
	BgpAsn              string   `json:"bgpAsn,omitempty"`
	BgpKey              string   `json:"bgpKey,omitempty"`
	EnableIpv6          int      `json:"enableIpv6,omitempty"`
	BaiduIpv6Address    string   `json:"baiduIpv6Address,omitempty"`
	CustomerIpv6Address string   `json:"customerIpv6Address,omitempty"`
	Ipv6Networks        []string `json:"ipv6Networks,omitempty"`
	Tags                []Tag    `json:"tags"`
}

type CreateEtChannelResult struct {
	Id string `json:"id"`
}

type DisableEtChannelIPv6Args struct {
	ClientToken string `json:"clientToken,omitempty"`
	EtId        string `json:"etId"`
	EtChannelId string `json:"etChannelId"`
}

type CreateEtChannelRouteRuleArgs struct {
	EtId        string `json:"etId"`
	EtChannelId string `json:"etChannelId"`
	ClientToken string `json:"clientToken,omitempty"`
	IpVersion   int    `json:"ipVersion,omitempty"`
	DestAddress string `json:"destAddress"`
	NextHopType string `json:"nexthopType"`
	NextHopId   string `json:"nexthopId"`
	Description string `json:"description,omitempty"`
}

type CreateEtChannelRouteRuleResult struct {
	RouteRuleId string `json:"routeRuleId"`
}

type ListEtChannelRouteRuleArgs struct {
	EtId        string `json:"etId"`
	EtChannelId string `json:"etChannelId"`
	Marker      string `json:"marker,omitempty"`
	MaxKeys     int    `json:"maxKeys,omitempty"`
	DestAddress string `json:"destAddress,omitempty"`
}

type ListEtChannelRouteRuleResult struct {
	Marker     string               `json:"marker"`
	IsTrucated bool                 `json:"isTruncated"`
	NextMarker string               `json:"nextMarker"`
	MaxKeys    int                  `json:"maxKeys"`
	RouteRules []EtChannelRouteRule `json:"routeRules"`
}

type EtChannelRouteRule struct {
	RouteRuleId     string `json:"routeRuleId"`
	IpVersion       int    `json:"ipVersion"`
	DestAddress     string `json:"destAddress"`
	NextHopType     string `json:"nexthopType"`
	NextHopId       string `json:"nexthopId"`
	Description     string `json:"description"`
	RouteProto      string `json:"routeProto"`
	AsPaths         string `json:"asPaths"`
	LocalPreference int    `json:"localPreference"`
	Med             int    `json:"med"`
	Origin          string `json:"origin"`
}

type UpdateEtChannelRouteRuleArgs struct {
	ClientToken string `json:"clientToken,omitempty"`
	EtId        string `json:"etId"`
	EtChannelId string `json:"etChannelId"`
	RouteRuleId string `json:"routeRuleId"`
	Description string `json:"description"`
}

type DeleteEtChannelRouteRuleArgs struct {
	ClientToken string `json:"clientToken,omitempty"`
	EtId        string `json:"etId"`
	EtChannelId string `json:"etChannelId"`
	RouteRuleId string `json:"routeRuleId"`
}

type AssociateEtChannelArgs struct {
	ClientToken    string `json:"clientToken,omitempty"`
	EtId           string `json:"etId"`
	EtChannelId    string `json:"etChannelId"`
	ExtraChannelId string `json:"extraChannelId"`
}

type DisAssociateEtChannelArgs struct {
	ClientToken    string `json:"clientToken,omitempty"`
	EtId           string `json:"etId"`
	EtChannelId    string `json:"etChannelId"`
	ExtraChannelId string `json:"extraChannelId"`
}
