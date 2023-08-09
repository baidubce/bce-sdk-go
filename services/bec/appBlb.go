/*
 * Copyright 2021 Baidu, Inc.
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

// client.go - define the client for BEC service

// Package bec defines the BEC services of BCE. The supported APIs are all defined in sub-package

package bec

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
	"strconv"
)

// CreateAppBlb - create app lb  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to create  app lb
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of create app lb
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateAppBlb(clientToken string, args *api.CreateAppBlbRequest) (*api.CreateAppBlbResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}

	result := &api.CreateAppBlbResponse{}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI(), Result: result, Body: args, Params: params}
	err := api.Post(c, req)

	return result, err
}

// UpdateAppBlb - update app lb  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update app lb
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of create app lb
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppBlb(clientToken, blbId string, args *api.ModifyBecBlbRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}

	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId, Body: args, Params: params}
	err := api.Put(c, req)

	return err
}

// GetAppBlbList - get app lb list with the specific parameters
//
// PARAMS:
//   - args: the arguments to get app lb list
//
// RETURNS:
//   - *AppBlbListResponse: the result of  app lb list
//   - error: nil if ok otherwise the specific error
func (c *Client) GetAppBlbList(args *api.MarkerRequest) (*api.AppBlbListResponse, error) {

	params := make(map[string]string)
	if args.Marker != "" {
		params["marker"] = args.Marker
	}
	if args.MaxKeys != 0 {
		params["maxKeys"] = strconv.Itoa(args.MaxKeys)
	}
	res := &api.AppBlbListResponse{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI(), Params: params, Result: res}
	err := api.Get(c, req)

	return res, err
}

// GetAppBlbDetails - get app lb detail with the specific parameters
//
// PARAMS:
//   - blbId: the arguments to get app lb details
//
// RETURNS:
//   - *AppBlbListResponse: the result of  app lb detail
//   - error: nil if ok otherwise the specific error
func (c *Client) GetAppBlbDetails(blbId string) (*api.AppBlbDetails, error) {

	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	res := &api.AppBlbDetails{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI() + "/" + blbId, Result: res}
	err := api.Get(c, req)
	return res, err
}

// DeleteAppBlbInstance - delete app lb  with the specific parameters
//
// PARAMS:
//   - blbId: the arguments to delete app lb
//
// RETURNS:
//   - *AppBlbListResponse: the result of  app lb detail
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppBlbInstance(blbId, clientToken string) error {

	if blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId, Params: params}
	err := api.Delete(c, req)
	return err
}

// CreateTcpListener - create app lb tcp listener  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to create  app lb tcp listener
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of create app lb tcp listener
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateTcpListener(clientToken, blbId string, args *api.CreateBecAppBlbTcpListenerRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}

	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/TCPlistener", Body: args, Params: params}
	err := api.Post(c, req)

	return err
}

// UpdateTcpListener - update app lb tcp listener  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update  app lb tcp listener
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of update app lb tcp listener
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateTcpListener(clientToken, blbId, listenerPort string, args *api.UpdateBecAppBlbTcpListenerRequest) error {
	if args == nil || blbId == "" || listenerPort == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	params["listenerPort"] = listenerPort
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/TCPlistener", Body: args, Params: params}
	err := api.Put(c, req)

	return err
}

// GetTcpListener - get app lb tcp listener  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to get  app lb tcp listener
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of get app lb tcp listener
//   - error: nil if ok otherwise the specific error
func (c *Client) GetTcpListener(blbId string, args *api.GetBecAppBlbListenerRequest) (*api.GetBecAppBlbTcpListenerResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if args.ListenerPort != 0 {
		params["listenerPort"] = strconv.Itoa(args.ListenerPort)
	}
	if args.MaxKeys != 0 {
		params["maxKeys"] = strconv.Itoa(args.MaxKeys)
	}
	if args.Marker != "" {
		params["marker"] = args.Marker
	}
	res := &api.GetBecAppBlbTcpListenerResponse{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/TCPlistener", Params: params, Result: res}
	err := api.Get(c, req)

	return res, err
}

// CreateUdpListener - create app lb tcp listener  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to create  app lb udp listener
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of create app lb udp listener
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateUdpListener(clientToken, blbId string, args *api.CreateBecAppBlbUdpListenerRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}

	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/UDPlistener", Body: args, Params: params}
	err := api.Post(c, req)

	return err
}

// UpdateUdpListener - update app lb udp listener  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update  app lb udp listener
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of update app lb udp listener
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateUdpListener(clientToken, blbId, listenerPort string, args *api.UpdateBecAppBlbUdpListenerRequest) error {
	if args == nil || blbId == "" || listenerPort == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	params["listenerPort"] = listenerPort
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/UDPlistener", Body: args, Params: params}
	err := api.Put(c, req)
	return err
}

// GetUdpListener - get app lb udp listener  with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to get  app lb udp listener
//
// RETURNS:
//   - *CreateAppBlbResponse: the result of get app lb udp listener
//   - error: nil if ok otherwise the specific error
func (c *Client) GetUdpListener(blbId string, args *api.GetBecAppBlbListenerRequest) (*api.GetBecAppBlbUdpListenerResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if args.ListenerPort != 0 {
		params["listenerPort"] = strconv.Itoa(args.ListenerPort)
	}
	if args.MaxKeys != 0 {
		params["maxKeys"] = strconv.Itoa(args.MaxKeys)
	}
	if args.Marker != "" {
		params["marker"] = args.Marker
	}
	res := &api.GetBecAppBlbUdpListenerResponse{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/UDPlistener", Params: params, Result: res}
	err := api.Get(c, req)

	return res, err
}

// DeleteAppBlbListener - delete app lb listener with the specific parameters
//
// PARAMS:
//   - blbId: the arguments to delete app lb listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppBlbListener(blbId, clientToken string, args *api.DeleteBlbListenerRequest) error {

	if blbId == "" || args == nil {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	params["batchdelete"] = ""
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/listener", Params: params, Body: args}
	err := api.Put(c, req)
	return err
}

// CreateIpGroup - create app lb ip group with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to create  app lb ip group
//
// RETURNS:
//   - *api.CreateBlbIpGroupResponse the result of app lb ip group
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateIpGroup(clientToken, blbId string, args *api.CreateBlbIpGroupRequest) (*api.CreateBlbIpGroupResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	res := &api.CreateBlbIpGroupResponse{}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup", Body: args, Params: params, Result: res}
	err := api.Post(c, req)

	return res, err
}

// UpdateIpGroup - update app lb ip group with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update  app lb ip group
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateIpGroup(clientToken, blbId string, args *api.UpdateBlbIpGroupRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup", Body: args, Params: params}
	err := api.Put(c, req)

	return err
}

// GetIpGroup - get app lb ip group with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to get  app lb ip group
//
// RETURNS:
//   - *api.GetBlbIpGroupListResponse the result of app lb ip group
//   - error: nil if ok otherwise the specific error
func (c *Client) GetIpGroup(blbId string, args *api.GetBlbIpGroupListRequest) (*api.GetBlbIpGroupListResponse, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if args.Name != "" {
		params["name"] = args.Name
	}
	if args.ExactlyMatch != false {
		params["exactlyMatch"] = strconv.FormatBool(args.ExactlyMatch)
	} else {
		params["exactlyMatch"] = strconv.FormatBool(false)
	}
	if args.Marker != "" {
		params["maker"] = args.Marker
	}
	if args.MaxKeys != 0 {
		params["maxKeys"] = strconv.Itoa(args.MaxKeys)
	}
	res := &api.GetBlbIpGroupListResponse{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup", Params: params, Result: res}
	err := api.Get(c, req)

	return res, err
}

// DeleteIpGroup - delete app lb ip group with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to delete  app lb ip group
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteIpGroup(clientToken, blbId string, args *api.DeleteBlbIpGroupRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	params["delete"] = ""
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup", Body: args, Params: params}
	err := api.Put(c, req)

	return err
}

// CreateIpGroupPolicy - create app lb ip group Policy with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to create  app lb ip group Policy
//
// RETURNS:
//   - *api.CreateBlbIpGroupResponse the result of app lb ip group Policy
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateIpGroupPolicy(clientToken, blbId string, args *api.CreateBlbIpGroupBackendPolicyRequest) (*api.CreateBlbIpGroupBackendPolicyResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	res := &api.CreateBlbIpGroupBackendPolicyResponse{}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/backendpolicy", Body: args, Params: params, Result: res}
	err := api.Post(c, req)

	return res, err
}

// UpdateIpGroupPolicy - update app lb ip group Policy with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update  app lb ip group Policy
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateIpGroupPolicy(clientToken, blbId string, args *api.UpdateBlbIpGroupBackendPolicyRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/backendpolicy", Body: args, Params: params}
	err := api.Put(c, req)

	return err
}

// GetIpGroupPolicyList - update app lb ip group Policy list with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update  app lb ip group Policy
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) GetIpGroupPolicyList(blbId string, args *api.GetBlbIpGroupPolicyListRequest) (*api.GetBlbIpGroupPolicyListResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if args.IpGroupId != "" {
		params["ipGroupId"] = args.IpGroupId
	}

	if args.Marker != "" {
		params["maker"] = args.Marker
	}
	if args.MaxKeys != 0 {
		params["maxKeys"] = strconv.Itoa(args.MaxKeys)
	}
	res := &api.GetBlbIpGroupPolicyListResponse{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/backendpolicy", Result: res, Params: params}
	err := api.Get(c, req)

	return res, err
}

// DeleteIpGroupPolicy - delete app lb ip group policy with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to delete  app lb ip group
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteIpGroupPolicy(clientToken, blbId string, args *api.DeleteBlbIpGroupBackendPolicyRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/backendpolicy", Body: args, Params: params}
	err := api.Delete(c, req)
	return err
}

// CreateIpGroupMember - create app lb ip group member with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to create  app lb ip group member
//
// RETURNS:
//   - *api.CreateBlbIpGroupResponse the result of app lb ip group member
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateIpGroupMember(clientToken, blbId string, args *api.CreateBlbIpGroupMemberRequest) (*api.CreateBlbIpGroupMemberResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	res := &api.CreateBlbIpGroupMemberResponse{}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/member", Body: args, Params: params, Result: res}
	err := api.Post(c, req)

	return res, err
}

// UpdateIpGroupMember - update app lb ip group member with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update  app lb ip group member
//
// RETURNS:
//   - *api.CreateBlbIpGroupResponse the result of app lb ip group member
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateIpGroupMember(clientToken, blbId string, args *api.UpdateBlbIpGroupMemberRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/member", Body: args, Params: params}
	err := api.Put(c, req)

	return err
}

// GetIpGroupMemberList - get app lb ip group member list with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to update  app lb ip group member
//
// RETURNS:
//   - *api.CreateBlbIpGroupResponse the result of app lb ip group member
//   - error: nil if ok otherwise the specific error
func (c *Client) GetIpGroupMemberList(blbId string, args *api.GetBlbIpGroupMemberListRequest) (*api.GetBlbIpGroupMemberListResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if args.IpGroupId != "" {
		params["ipGroupId"] = args.IpGroupId
	}

	if args.Marker != "" {
		params["maker"] = args.Marker
	}
	if args.MaxKeys != 0 {
		params["maxKeys"] = strconv.Itoa(args.MaxKeys)
	}
	res := &api.GetBlbIpGroupMemberListResponse{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/member", Result: res, Params: params}
	err := api.Get(c, req)

	return res, err
}

// DeleteIpGroupMember - delete app lb ip group member with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to delete  app lb ip group member
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteIpGroupMember(clientToken, blbId string, args *api.DeleteBlbIpGroupBackendMemberRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	params["delete"] = ""
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/ipgroup/member", Body: args, Params: params}
	err := api.Put(c, req)

	return err
}

// CreateListenerPolicy - create app lb listener policy with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to create  app lb ip listener policy
//
// RETURNS:
//   - *api.CreateBlbIpGroupResponse the result of app lb iplistener policy
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateListenerPolicy(clientToken, blbId string, args *api.CreateAppBlbPoliciesRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/policys", Body: args, Params: params}
	err := api.Post(c, req)

	return err
}

// GetListenerPolicy - get app lb listener policy with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to get  app lb ip listener policy
//
// RETURNS:
//   - *api.CreateBlbIpGroupResponse the result of app lb ip listener policy
//   - error: nil if ok otherwise the specific error
func (c *Client) GetListenerPolicy(blbId string, args *api.GetBlbListenerPolicyRequest) (*api.GetBlbListenerPolicyResponse, error) {
	if args == nil || blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if args.Type != "" {
		params["type"] = args.Type
	}
	if args.Port != 0 {
		params["port"] = strconv.Itoa(args.Port)
	}

	if args.Marker != "" {
		params["maker"] = args.Marker
	}
	if args.MaxKeys != 0 {
		params["maxKeys"] = strconv.Itoa(args.MaxKeys)
	}
	res := &api.GetBlbListenerPolicyResponse{}
	req := &api.GetHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/policys", Result: res, Params: params}
	err := api.Get(c, req)

	return res, err
}

// DeleteListenerPolicy - delete app lb listener policy with the specific parameters
//
// PARAMS:
//   - clientToken:idempotent token，an ASCII string no longer than 64 bits，unnecessary
//   - args: the arguments to delete  app lb listener policy
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteListenerPolicy(clientToken, blbId string, args *api.DeleteAppBlbPoliciesRequest) error {
	if args == nil || blbId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	params["batchdelete"] = ""
	if clientToken != "" {
		params["clientToken"] = clientToken
	}
	req := &api.PostHttpReq{Url: api.GetAppBlbURI() + "/" + blbId + "/policys", Body: args, Params: params}
	err := api.Put(c, req)

	return err
}
