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
	PROJECT_PREFIX            = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "project"
	LIST_PROJECT_PREFIX       = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "project" + bce.URI_PREFIX + "list"
	LOGSTORE_PREFIX           = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "logstore"
	FASTQUERY_PREFIX          = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "fastquery"
	LOGSHIPPER_PREFIX         = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "logshipper"
	DOWNLOAD_TASK_PREFIX      = bce.URI_PREFIX + "v2" + bce.URI_PREFIX + "logstore" + bce.URI_PREFIX + "download"
	LIST_DOWNLOAD_TASK_PREFIX = DOWNLOAD_TASK_PREFIX + bce.URI_PREFIX + "list"
	BIND_PREFIX               = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "logstore" + bce.URI_PREFIX + "bind"
	UNBIND_PREFIX             = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "logstore" + bce.URI_PREFIX + "unbind"
	LIST_LOGSTORE_PREFIX      = bce.URI_PREFIX + "v2" + bce.URI_PREFIX + "logstore" + bce.URI_PREFIX + "list"
	BATCH_PREFIX              = bce.URI_PREFIX + "v1" + bce.URI_PREFIX + "logstore" + bce.URI_PREFIX + "batch"
)

func getProjectUri(UUID string) string {
	return PROJECT_PREFIX + bce.URI_PREFIX + UUID
}

func getLogStoreUri(logStoreName string) string {
	return LOGSTORE_PREFIX + bce.URI_PREFIX + logStoreName
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

func getLogShipperUri(logShipperID string) string {
	return LOGSHIPPER_PREFIX + bce.URI_PREFIX + logShipperID
}

func getLogShipperRecordUri(logShipperID string) string {
	return getLogShipperUri(logShipperID) + bce.URI_PREFIX + "record"
}

func getLogShipperStatusUri(logShipperID string) string {
	return getLogShipperUri(logShipperID) + bce.URI_PREFIX + "status"
}

func getBulkSetLogShipperStatusUri() string {
	return LOGSHIPPER_PREFIX + bce.URI_PREFIX + "status" + bce.URI_PREFIX + "batch"
}

func getDownloadTaskUri(UUID string) string {
	return DOWNLOAD_TASK_PREFIX + bce.URI_PREFIX + UUID
}

func getDownloadTaskLinkUri(UUID string) string {
	return DOWNLOAD_TASK_PREFIX + bce.URI_PREFIX + "link" + bce.URI_PREFIX + UUID
}
