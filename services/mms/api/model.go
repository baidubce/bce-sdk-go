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

package api

import (
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v2"
	VIDEO_URI  = "/videolib"
	IMAGE_URI  = "/imagelib"
)

type BaseRequest struct {
	Source       string `json:"source"`
	Description  string `json:"description"`
	Notification string `json:"notification"`
	NeedTag      bool   `json:"needTag"`
}

type BaseResponse struct {
	TaskID      string  `json:"taskId"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
	Duration    float64 `json:"duration"`
	Error       struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Lib        string    `json:"lib"`
	Source     string    `json:"source"`
	UpdateTime time.Time `json:"updateTime"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	CreateTime time.Time `json:"createTime,"`
}

type MatchFrame struct {
	Distance  float64 `json:"distance"`
	Position  int     `json:"position"`
	Timestamp float64 `json:"timestamp"`
}

type VideoClip struct {
	Clip            bool    `json:"clip"`
	ClipNum         int     `json:"clipNum"`
	Distance        float64 `json:"distance"`
	FrameNum        int     `json:"frameNum"`
	InputEndPos     int     `json:"inputEndPos"`
	InputEndTime    float64 `json:"inputEndTime"`
	InputStartPos   int     `json:"inputStartPos"`
	InputStartTime  float64 `json:"inputStartTime"`
	InputSumTime    float64 `json:"inputSumTime"`
	MatchNum        int     `json:"matchNum"`
	OutputEndPos    int     `json:"outputEndPos"`
	OutputEndTime   float64 `json:"outputEndTime"`
	OutputStartPos  int     `json:"outputStartPos"`
	OutputStartTime float64 `json:"outputStartTime"`
	OutputSumTime   float64 `json:"outputSumTime"`
	PreIdx          int     `json:"preIdx"`
}

type SearchTaskResult struct {
	Cover       string       `json:"cover"`
	Description string       `json:"description"`
	Distance    float64      `json:"distance"`
	Duration    float64      `json:"duration"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Score       float64      `json:"score"`
	Source      string       `json:"source"`
	Type        string       `json:"type"`
	Frames      []MatchFrame `json:"frames"`
	Clips       []VideoClip  `json:"clips"`
}

type SearchTaskResultResponse struct {
	BaseResponse
	Results    []SearchTaskResult `json:"results"`
	TagResults []SearchTaskResult `json:"tagResults"`
}
