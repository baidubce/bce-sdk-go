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

// Package afd Automated Fraud Detection
package afd

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// Sync api: /rcs/sync
func (c *Client) Sync(args *SyncArgs) (res *SyncResponse, err error) {
	err = bce.NewRequestBuilder(c).
		WithURL(bce.URI_PREFIX + "rcs/sync").
		WithMethod(http.POST).
		WithBody(args).
		WithResult(&res).
		Do()

	return res, err
}

// Factor api: /rcs/factor-saas
func (c *Client) Factor(args *FactorArgs) (res *FactorResponse, err error) {
	err = bce.NewRequestBuilder(c).
		WithURL(bce.URI_PREFIX + "rcs/factor-saas").
		WithMethod(http.POST).
		WithBody(args).
		WithResult(&res).
		Do()

	return res, err
}
