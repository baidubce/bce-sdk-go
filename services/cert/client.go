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

// client.go - define the client for Cert service

// Package cert defines the Cert services of BCE. The supported APIs are all defined in sub-package
package cert

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX       = bce.URI_PREFIX + "v1"
	DEFAULT_ENDPOINT = "certificate.baidubce.com"
	REQUEST_CERT_URL = "/certificate"
)

// Client of Cert service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getCertUri() string {
	return URI_PREFIX + REQUEST_CERT_URL
}

func getCertUriWithId(id string) string {
	return URI_PREFIX + REQUEST_CERT_URL + "/" + id
}
