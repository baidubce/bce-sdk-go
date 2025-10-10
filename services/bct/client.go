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

package bct

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bct/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "iam." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client of BCT service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk string) (*Client, error) {
	return NewClientWithEndpoint(ak, sk, DEFAULT_SERVICE_DOMAIN)
}

func NewClientWithEndpoint(ak, sk, endpoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{BceClient: bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

func (c *Client) QueryEventsV2(args *api.QueryEventsV2Request) (*api.QueryEventsV2Response, error) {
	body, err := NewBodyFromStruct(args)
	if err != nil {
		return nil, err
	}
	return api.QueryEventsV2(c, body)
}

func NewBodyFromStruct(args interface{}) (*bce.Body, error) {
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return body, nil
}
