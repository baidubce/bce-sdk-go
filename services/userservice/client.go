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

// client.go - define the client for User Service service

// Package userservice defines the User Service services of BCE. The supported APIs are all defined in sub-package
package userservice

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_SERVICE_DOMAIN = "blb." + bce.DEFAULT_REGION + ".baidubce.com"
	URI_PREFIX             = bce.URI_PREFIX + "v1"
	REQUEST_SERVICE_URL    = "/service"
)

type Client struct {
	*bce.BceClient
}

// NewClient 是一个函数，用于创建一个新的客户端对象
func NewClient(ak, sk, endPoint string) (*Client, error) {
	if endPoint == "" {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err

	}
	return &Client{client}, nil
}

// userServiceRequest 发起用户服务请求
func (c *Client) userServiceRequest(action, method, uri, clientToken string, args interface{}) error {
	req := &bce.BceRequest{}
	req.SetMethod(method)
	req.SetUri(uri)
	req.SetParam(action, "")
	req.SetParam("clientToken", clientToken)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	jsonBody, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(jsonBody)
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// getUserServiceUri 函数返回用户服务的URI地址
func getUserServiceUri() string {
	return URI_PREFIX + REQUEST_SERVICE_URL
}

// getUserServiceIdUri 返回指定服务的用户服务ID的URI
func getUserServiceIdUri(service string) string {
	return URI_PREFIX + REQUEST_SERVICE_URL + "/" + service
}
