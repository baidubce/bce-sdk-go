/*
 * Copyright 2024 Baidu, Inc.
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

package userservice

type AuthStatus string

const (
	ServiceAuthAllow AuthStatus = "allow"
	ServiceAuthDeny  AuthStatus = "deny"
)

type CreateUserServiceArgs struct {
	ClientToken string                 `json:"-"`
	InstanceId  string                 `json:"instanceId"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	ServiceName string                 `json:"serviceName"`
	AuthList    []UserServiceAuthModel `json:"authList"`
}

type CreateUserServiceResult struct {
	Service string `json:"service"`
}

type UserServiceAuthModel struct {
	Uid  string     `json:"uid"`
	Auth AuthStatus `json:"auth"`
}

type UpdateServiceArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserServiceBindArgs struct {
	ClientToken string `json:"-"`
	InstanceId  string `json:"instanceId"`
}

type UserServiceUnBindArgs struct {
	ClientToken string `json:"-"`
}

type UserServiceAuthArgs struct {
	ClientToken string                 `json:"-"`
	AuthList    []UserServiceAuthModel `json:"authList"`
}

type UserServiceRemoveAuthArgs struct {
	ClientToken string   `json:"-"`
	UidList     []string `json:"uidList"`
}

type DescribeUserServicesArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type DescribeUserServicesResult struct {
	Services    []UserServiceModel `json:"services"`
	Marker      string             `json:"marker"`
	IsTruncated bool               `json:"isTruncated"`
	NextMarker  string             `json:"nextMarker"`
	MaxKeys     int                `json:"maxKeys"`
}

type UserServiceModel struct {
	ServiceId     string                 `json:"serviceId"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	ServiceName   string                 `json:"serviceName"`
	BindType      string                 `json:"bindType"`
	InstanceId    string                 `json:"instanceId"`
	Status        string                 `json:"status"`
	Service       string                 `json:"service"`
	CreateTime    string                 `json:"createTime"`
	EndpointCount int                    `json:"endpointCount"`
	EndpointList  []RelatedEndpointModel `json:"endpointList"`
	AuthList      []UserServiceAuthModel `json:"authList"`
}

type RelatedEndpointModel struct {
	EndpointId string `json:"endpointId"`
	Uid        string `json:"uid"`
	AttachTime string `json:"attachTime"`
}

type DescribeUserServiceDetailResult struct {
	UserServiceModel
}
