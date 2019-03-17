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

// model.go - definitions of the request arguments and results data structure model for VCA

package api

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"
	MEDIA_URI  = "/media"
)

type PutMediaArgs struct {
	Source       string `json:"source"`
	Auth         string `json:"auth"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Preset       string `json:"preset"`
	Notification string `json:"notification"`
}

type GetMediaResult struct {
	Source       string          `json:"source"`
	MediaId      string          `json:"mediaId"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Preset       string          `json:"preset"`
	Status       string          `json:"status"`
	Percent      int             `json:"percent"`
	Notification string          `json:"notification"`
	CreateTime   string          `json:"createTime"`
	StartTime    string          `json:"startTime"`
	PublishTime  string          `json:"publishTime"`
	Results      []AnalyzeResult `json:"results"`
	Error        AnalyzeError    `json:"error"`
}

type AnalyzeResult struct {
	Type       string             `json:"type"`
	Attributes []AnalyzeAttribute `json:"result"`
}

type AnalyzeAttribute struct {
	Attribute     string       `json:"attribute"`
	Confidence    float32      `json:"confidence"`
	Source        string       `json:"source"`
	AttributeTime []TimePeriod `json:"time"`
}

type TimePeriod struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type AnalyzeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
