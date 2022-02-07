/*
 * Copyright 2022 Baidu, Inc.
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

// client.go - define the client for KMS service

// Package kms defines the KMS services of BCE. The supported APIs are all defined in sub-package
package kms

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	// DEFAULTENDPOINT KMS endpoint
	DEFAULTENDPOINT = "https://bkm.%s.baidubce.com"
	// DEFAULTREGION default region
	DEFAULTREGION = "bj"
)

// Client of KMS service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient return BCE Client
func NewClient(ak, sk, region string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error

	credentials, err = auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}

	// if region is not given use DEFAULT_REGION
	if len(region) == 0 {
		region = DEFAULTREGION
	}
	// 设置签名头为默认签名头
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	// 设置 bce 配置
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  fmt.Sprintf(DEFAULTENDPOINT, region),
		Region:                    region,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}

	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}
