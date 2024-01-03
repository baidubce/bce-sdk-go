/*
 * Copyright 2023 Baidu, Inc.
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

// client.go - define the client for ET service

// Package et defines the et services of BCE.
// The supported APIs are all defined in different files.
package et

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	DEFAULT_ENDPOINT = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

	REQUEST_ET_URL                    = "/et"
	REQUEST_ET_CHANNEL_URL            = "/channel"
	REQUEST_ET_CHANNEL_ROUTE_URL      = "/route"
	REQUEST_ET_CHANNEL_ROUTE_RULE_URL = "/rule"
)

// Client of ET service is a kind of BceClient, so derived from BceClient
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

func getURLForEt() string {
	return URI_PREFIX + REQUEST_ET_URL
}

func getURLForEtId(etId string) string {
	return getURLForEt() + "/" + etId
}

func getURLForEtChannel(etId string) string {
	return getURLForEtId(etId) + REQUEST_ET_CHANNEL_URL
}

func getURLForEtChannelId(etId string, etChannelId string) string {
	return getURLForEtChannel(etId) + "/" + etChannelId
}

func getURLForEtChannelRoute(etId string, etChannelId string) string {
	return getURLForEtChannelId(etId, etChannelId) + REQUEST_ET_CHANNEL_ROUTE_URL
}

func getURLForEtChannelRouteRule(etId string, etChannelId string) string {
	return getURLForEtChannelRoute(etId, etChannelId) + REQUEST_ET_CHANNEL_ROUTE_RULE_URL
}

func getURLForEtChannelRouteRuleId(etId string, etChannelId string, routeRuleId string) string {
	return getURLForEtChannelRouteRule(etId, etChannelId) + "/" + routeRuleId
}
