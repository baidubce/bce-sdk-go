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
package quotacenter

type QuotaCenterQueryArgs struct {
	Type        string `json:"type"`
	ServiceType string `json:"serviceType"`
	Region      string `json:"region"`
	Name        string `json:"name,omitempty"`
	Marker      string `json:"marker,omitempty"`
	MaxKeys     int    `json:"maxKeys,omitempty"`
}

type ListQuotaResult struct {
	Marker      string       `json:"marker"`
	MaxKeys     int          `json:"maxKeys"`
	NextMarker  string       `json:"nextMarker"`
	IsTruncated bool         `json:"isTruncated"`
	Result      []QuotaModel `json:"result"`
}

type QuotaModel struct {
	ProductType string `json:"productType"`
	ServiceType string `json:"serviceType"`
	Type        string `json:"type"`
	Region      string `json:"region"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Used        string `json:"used"`
}

type ProductQueryArgs struct {
	ProductType string `json:"productType,omitempty"`
	Marker      string `json:"marker,omitempty"`
	MaxKeys     int    `json:"maxKeys,omitempty"`
}

type ListProductResult struct {
	Marker      string         `json:"marker"`
	MaxKeys     int            `json:"maxKeys"`
	NextMarker  string         `json:"nextMarker"`
	IsTruncated bool           `json:"isTruncated"`
	Result      []ProductModel `json:"result"`
}

type ProductModel struct {
	ProductType string `json:"productType"`
	ServiceType string `json:"serviceType"`
}

type RegionQueryArgs struct {
	ProductType string `json:"productType"`
	ServiceType string `json:"serviceType"`
	Type        string `json:"type"`
}

type ListRegionResult struct {
	Regions []string `json:"regions"`
}

type InfoQueryArgs struct {
	ServiceType string `json:"serviceType,omitempty"`
	Region      string `json:"region,omitempty"`
	Marker      string `json:"marker,omitempty"`
	MaxKeys     int    `json:"maxKeys,omitempty"`
}

type ListInfoResult struct {
	Marker      string      `json:"marker"`
	MaxKeys     int         `json:"maxKeys"`
	NextMarker  string      `json:"nextMarker"`
	IsTruncated bool        `json:"isTruncated"`
	Result      []InfoModel `json:"result"`
}

type InfoModel struct {
	ProductType string `json:"productType"`
	ServiceType string `json:"serviceType"`
	Type        string `json:"type"`
	Region      string `json:"region"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Apply       bool   `json:"apply"`
}

type ApplicationCreateModel struct {
	ProductType string `json:"productType"`
	ServiceType string `json:"serviceType"`
	Type        string `json:"type"`
	Region      string `json:"region"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Reason      string `json:"reason"`
}

type IdModel struct {
	id string `json:"id"`
}

type ApplicationQueryArgs struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	ProductType string `json:"status,productType"`
	ServiceType string `json:"serviceType"`
	Type        string `json:"type"`
	Region      string `json:"region"`
	Marker      string `json:"marker,omitempty"`
	MaxKeys     int    `json:"maxKeys,omitempty"`
}

type ListApplicationResult struct {
	Marker      string             `json:"marker"`
	MaxKeys     int                `json:"maxKeys"`
	NextMarker  string             `json:"nextMarker"`
	IsTruncated bool               `json:"isTruncated"`
	Result      []ApplicationModel `json:"result"`
}

type ApplicationModel struct {
	Id          string `json:"id"`
	ProductType string `json:"productType"`
	ServiceType string `json:"serviceType"`
	Type        string `json:"type"`
	Region      string `json:"region"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Reason      string `json:"reason"`
	Status      string `json:"status"`
	Conclusion  string `json:"conclusion"`
	CreateTime  string `json:"createTime"`
	EffectTime  string `json:"effectTime,omitempty"`
	ApproveTime string `json:"approveTime,omitempty"`
}
