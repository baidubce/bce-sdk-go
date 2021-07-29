/*
 * Copyright 2021 Baidu, Inc.
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

// util.go - define the utilities for api package of BLS service

package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_PREFIX   = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "logstore"
	FASTQUERY_PREFIX = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "fastquery"
)

func getLogStoreUri(logStoreName string) string {
	return DEFAULT_PREFIX + bce.URI_PREFIX + logStoreName
}

func getLogStreamName(logStoreName string) string {
	return getLogStoreUri(logStoreName) + bce.URI_PREFIX + "logstream"
}

func getLogRecordUri(logStoreName string) string {
	return getLogStoreUri(logStoreName) + bce.URI_PREFIX + "logrecord"
}

func getFastQueryUri(fastQuery string) string {
	return FASTQUERY_PREFIX + bce.URI_PREFIX + fastQuery
}

func getIndexUri(logStoreName string) string {
	return getLogStoreUri(logStoreName) + bce.URI_PREFIX + "index"
}
