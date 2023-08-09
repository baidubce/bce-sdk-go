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

package afd

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

// Client def
type Client struct {
	*bce.BceClient
}

// NewClient ak, sk, endPoint
func NewClient(args ...string) (*Client, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Invalid ak/sk")
	}

	endPoint := "afd.baidubce.com"
	if len(args) > 2 {
		endPoint = args[2]
	}

	client, err := bce.NewBceClientWithAkSk(args[0], args[1], endPoint)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
