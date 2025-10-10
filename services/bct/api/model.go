/*
 * Copyright 2025 Baidu, Inc.
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

package api

import "time"

type QueryEventsV2Request CloudTrailQueryWithMarkerRequest

type QueryEventsV2Response PageWithMarkerResponse

type FieldFilter struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CloudTrailQueryWithMarkerRequest struct {
	DomainId   string        `json:"domainId,omitempty"`
	Filters    []FieldFilter `json:"filters,omitempty"`
	StartTime  time.Time     `json:"startTime"`
	EndTime    time.Time     `json:"endTime"`
	PageSize   int           `json:"pageSize"`
	NextMarker string        `json:"nextMarker,omitempty"`
}

type ResourceInfo struct {
	ResourceType string `json:"resourceType"`
	ResourceId   string `json:"resourceId"`
	ResourceName string `json:"resourceName"`
}

type UserIdentity struct {
	IamDomainId          string `json:"iamDomainId"`
	IamUserId            string `json:"iamUserId"`
	RoleId               string `json:"roleId"`
	LoginUserId          string `json:"loginUserId"`
	UserDisplayName      string `json:"userDisplayName"`
	OrganizationId       string `json:"organizationId"`
	OrganizationMasterId string `json:"organizationMasterId"`
	Accesskey            string `json:"accesskey"`
	ApiKey               string `json:"apiKey"`
	SecurityToken        string `json:"securityToken"`
}

type EventDTO struct {
	EventType               string         `json:"eventType"`
	EventSource             string         `json:"eventSource"`
	EventName               string         `json:"eventName"`
	EventTimeInMilliseconds int64          `json:"eventTimeInMilliseconds"`
	EventTime               string         `json:"eventTime"`
	UserIpAddress           string         `json:"userIpAddress"`
	UserAgent               string         `json:"userAgent"`
	RegionId                string         `json:"regionId"`
	RequestId               string         `json:"requestId"`
	OrderId                 string         `json:"orderId"`
	ApiVersion              string         `json:"apiVersion"`
	Description             string         `json:"description"`
	ErrorCode               string         `json:"errorCode"`
	ErrorMessage            string         `json:"errorMessage"`
	Success                 bool           `json:"success"`
	UserIdentity            UserIdentity   `json:"userIdentity"`
	Resources               []ResourceInfo `json:"resources"`
	OriginRequestParameters interface{}    `json:"originRequestParameters"`
	OriginResponse          interface{}    `json:"originResponse"`
}

type PageWithMarkerResponse struct {
	PageSize   int        `json:"pageSize"`
	Data       []EventDTO `json:"data"`
	NextMarker string     `json:"nextMarker"`
	Truncated  bool       `json:"truncated"`
}
