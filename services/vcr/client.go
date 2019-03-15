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

// client.go - define the client for VCR service

// Package vcr defines the VCR services of BCE. The supported APIs are all defined in sub-package
package vcr

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/vcr/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "vcr." + bce.DEFAULT_REGION + "." + bce.DEFAULT_DOMAIN
)

// Client of VCR service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the VCR service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
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

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

func (c *Client) PutMedia(args *api.PutMediaArgs) error {
	return api.PutMedia(c, args)
}

func (c *Client) SimplePutMedia(source string, description string, preset string, notification string) error {
	args := &api.PutMediaArgs{Source: source, Description: description, Preset: preset, Notification: notification}
	return api.PutMedia(c, args)
}

func (c *Client) GetMedia(source string) (*api.GetMediaResult, error) {
	return api.GetMedia(c, source)
}

func (c *Client) PutText(args *api.PutTextArgs) (*api.PutTextResult, error) {
	return api.PutText(c, args)
}

func (c *Client) SimplePutText(text string) (*api.PutTextResult, error) {
	args := &api.PutTextArgs{Text: text}
	return api.PutText(c, args)
}
