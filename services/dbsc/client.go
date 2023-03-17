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

package dbsc

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const DEFAULT_SERVICE_DOMAIN = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

// Client of BCC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the BCC service client with default configuration.
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

func (c *Client) CreateVolumeCluster(args *CreateVolumeClusterArgs) (*CreateVolumeClusterResult, error) {
	return CreateVolumeCluster(c, args)
}

func (c *Client) ListVolumeCluster(queryArgs *ListVolumeClusterArgs) (*ListVolumeClusterResult, error) {
	return ListVolumeCluster(c, queryArgs)
}

func (c *Client) GetVolumeClusterDetail(clusterId string) (*VolumeClusterDetail, error) {
	return GetVolumeClusterDetail(c, clusterId)
}

func (c *Client) ResizeVolumeCluster(clusterId string, args *ResizeVolumeClusterArgs) error {
	return ResizeVolumeCluster(c, clusterId, args)
}

func (c *Client) PurchaseReservedVolumeCluster(clusterId string, args *PurchaseReservedVolumeClusterArgs) error {
	return PurchaseReservedVolumeCluster(c, clusterId, args)
}

func (c *Client) AutoRenewVolumeCluster(args *AutoRenewVolumeClusterArgs) error {
	return AutoRenewVolumeCluster(c, args)
}

func (c *Client) CancelAutoRenewVolumeCluster(args *CancelAutoRenewVolumeClusterArgs) error {
	return CancelAutoRenewVolumeCluster(c, args)
}
