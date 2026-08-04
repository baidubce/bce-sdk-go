/*
 * Copyright 2026 Baidu, Inc.
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

// client.go - define the client for Kafka service

package kafka

import (
	"errors"
	"fmt"
	"time"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_ENDPOINT                     = "kafka." + bce.DEFAULT_REGION + ".baidubce.com"
	DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS = 50 * 1000
	DEFAULT_SOCKET_TIMEOUT_IN_MILLIS     = 50 * 1000
)

// Client of Kafka service is a kind of BceClient, so derived from BceClient.
type Client struct {
	*bce.BceClient
}

// NewClient creates a Kafka service client with AK/SK credentials.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	return NewClientWithConfig(&KafkaClientConfiguration{
		Endpoint:    endPoint,
		Credentials: credentials,
	})
}

// NewClientWithSTS creates a Kafka service client with temporary session credentials.
func NewClientWithSTS(ak, sk, token, endPoint string) (*Client, error) {
	credentials, err := auth.NewSessionBceCredentials(ak, sk, token)
	if err != nil {
		return nil, err
	}
	return NewClientWithConfig(&KafkaClientConfiguration{
		Endpoint:    endPoint,
		Credentials: credentials,
	})
}

// NewClientWithConfig creates a Kafka service client from a Kafka client configuration.
func NewClientWithConfig(config *KafkaClientConfiguration) (*Client, error) {
	if config == nil {
		return nil, errors.New("kafka client configuration should not be nil")
	}
	conf := *config
	if len(conf.Region) == 0 {
		conf.Region = bce.DEFAULT_REGION
	}
	if len(conf.Endpoint) == 0 {
		conf.Endpoint = fmt.Sprintf("kafka.%s.baidubce.com", conf.Region)
	}
	if len(conf.UserAgent) == 0 {
		conf.UserAgent = bce.DEFAULT_USER_AGENT
	}
	conf.SignOption = &auth.SignOptions{
		HeadersToSign: map[string]struct{}{
			"host":       {},
			"x-bce-date": {},
		},
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
	}
	if conf.Retry == nil {
		conf.Retry = bce.DEFAULT_RETRY_POLICY
	}
	if conf.ConnectionTimeoutInMillis == 0 {
		conf.ConnectionTimeoutInMillis = DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS
	}
	if conf.DialTimeout == nil {
		timeout := time.Duration(conf.ConnectionTimeoutInMillis) * time.Millisecond
		conf.DialTimeout = &timeout
	}
	if conf.ReadTimeout == nil {
		timeout := time.Duration(DEFAULT_SOCKET_TIMEOUT_IN_MILLIS) * time.Millisecond
		conf.ReadTimeout = &timeout
	}
	bceClient, err := bce.NewBceClientWithExclusiveHTTPClient(&conf, &auth.BceV1Signer{})
	if err != nil {
		return nil, err
	}
	return &Client{bceClient}, nil
}
