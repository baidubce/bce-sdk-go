/*
 * Copyright 2020 Baidu, Inc.
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

// client.go - define the client for BES service

// Package bes defines the BES services of BCE. The supported APIs are all defined in sub-package
package bes

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const DEFAULT_SERVICE_DOMAIN = "bes." + bce.DEFAULT_REGION + ".baidubce.com"

// Client of BES service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the BES service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	if endPoint == "" {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endPoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// CreateInstance - create an instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *CreateInstanceResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateCluster(args *ESClusterRequest) (*ESClusterResponse, error) {

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return CreateCluster(c, args, body)
}
func (c *Client) DeleteCluster(args *GetESClusterRequest) (*DeleteESClusterResponse, error) {

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return DeleteCluster(c, args, body)
}
func (c *Client) GetCluster(args *GetESClusterRequest) (*DetailESClusterResponse, error) {

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return GetCluster(c, args, body)
}
func getCreateUri() string {
	return "/api/bes/cluster/create"
}
func getReadUri() string {
	return "/api/bes/cluster/detail"
}
func getDeleteUri() string {
	return "/api/bes/cluster/delete"
}

func GetCluster(cli bce.Client, args *GetESClusterRequest, reqBody *bce.Body) (*DetailESClusterResponse, error) {
	//clientToken := args.ClientToken
	//requestToken := args.RequestToken
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getReadUri())
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")
	req.SetHeader("X-Region", cli.GetBceClientConfig().Region)
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	//req.SetHeader("x-request-token", requestToken)

	//if clientToken != "" {
	//	req.SetParam("clientToken", clientToken)
	//}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DetailESClusterResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func DeleteCluster(cli bce.Client, args *GetESClusterRequest, reqBody *bce.Body) (*DeleteESClusterResponse, error) {
	//clientToken := args.ClientToken
	//requestToken := args.RequestToken
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeleteUri())
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")
	req.SetHeader("X-Region", cli.GetBceClientConfig().Region)
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	//req.SetHeader("x-request-token", requestToken)

	//if clientToken != "" {
	//	req.SetParam("clientToken", clientToken)
	//}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeleteESClusterResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
func CreateCluster(cli bce.Client, args *ESClusterRequest, reqBody *bce.Body) (*ESClusterResponse, error) {
	//clientToken := args.ClientToken
	//requestToken := args.RequestToken
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCreateUri())
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")
	req.SetHeader("X-Region", cli.GetBceClientConfig().Region)
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	//req.SetHeader("x-request-token", requestToken)

	//if clientToken != "" {
	//	req.SetParam("clientToken", clientToken)
	//}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ESClusterResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
