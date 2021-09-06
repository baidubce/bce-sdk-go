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

// model.go - definitions of the request arguments and results data structure model for VCR

package api

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"
	MEDIA_URI  = "/media"
	TEXT_URI   = "/text"
	IMAGE_URI  = "/image"
)

type PutMediaArgs struct {
	Source       string `json:"source"`
	Auth         string `json:"auth"`
	Description  string `json:"description"`
	Preset       string `json:"preset"`
	Notification string `json:"notification"`
}

type GetMediaResult struct {
	Source       string        `json:"source"`
	MediaId      string        `json:"mediaId"`
	Description  string        `json:"description"`
	Preset       string        `json:"preset"`
	Status       string        `json:"status"`
	Percent      int           `json:"percent"`
	Notification string        `json:"notification"`
	CreateTime   string        `json:"createTime"`
	FinishTime   string        `json:"finishTime"`
	Label        string        `json:"label"`
	Results      []CheckResult `json:"results"`
	Error        CheckError    `json:"error"`
}

type CheckResult struct {
	Type  string       `json:"type"`
	Items []ResultItem `json:"items"`
}

type ResultItem struct {
	SubType       string   `json:"subType"`
	Target        string   `json:"target"`
	TimeInSeconds int      `json:"timeInSeconds"`
	Confidence    float32  `json:"confidence"`
	Label         string   `json:"label"`
	Extra         string   `json:"extra"`
	Evidence      Evidence `json:"evidence"`
}

type Evidence struct {
	Thumbnail string   `json:"thumbnail"`
	Location  Location `json:"location"`
	Text      string   `json:"text"`
}

type Location struct {
	LeftOffsetInPixel int `json:"leftOffsetInPixel"`
	TopOffsetInPixel  int `json:"topOffsetInPixel"`
	WidthInPixel      int `json:"widthInPixel"`
	HeightInPixel     int `json:"heightInPixel"`
}

type CheckError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PutTextArgs struct {
	Text   string `json:"text"`
	Preset string `json:"preset"`
}

type PutTextResult struct {
	Text    string        `json:"text"`
	Preset  string        `json:"preset"`
	Results []CheckResult `json:"results"`
}

type PutImageSyncArgs struct {
	Source string `json:"source"`
	Preset string `json:"preset"`
}

type PutImageSyncResult struct {
	Source  string        `json:"source"`
	Label   string        `json:"label"`
	Preset  string        `json:"preset"`
	Status  string        `json:"status"`
	Results []CheckResult `json:"results"`
}
