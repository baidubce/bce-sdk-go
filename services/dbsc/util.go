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

// util.go - define the utilities for api package of BBC service
package dbsc

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_PREFIX_V2      = bce.URI_PREFIX + "v2"
	REQUEST_VOLUME_URI = "/volume"

	REQUEST_CREATE_URI            = "/cluster"
	REQUEST_AUTO_RENEW_URI        = "/cluster/autoRenew"
	REQUEST_CANCEL_AUTO_RENEW_URI = "/cluster/cancelAutoRenew"
)

func getVolumeClusterUri() string {
	return URI_PREFIX_V2 + REQUEST_VOLUME_URI
}
