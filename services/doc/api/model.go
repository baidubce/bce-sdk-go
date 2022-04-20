/*
 * Copyright 2022 Baidu, Inc.
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

// model.go - definitions of the request arguments and results data structure model

package api

import (
	"encoding/json"
	"errors"
)

type StatusType string

const (
	DOC_TARGET_H5    = "h5"
	DOC_TARGET_IMAGE = "image"

	DOC_PUBLIC  = "PUBLIC"
	DOC_PRIVATE = "PRIVATE"

	DOC_STATUS_UPLOADING  StatusType = "UPLOADING"
	DOC_STATUS_PROCESSING StatusType = "PROCESSING"
	DOC_STATUS_PUBLISHED  StatusType = "PUBLISHED"
	DOC_STATUS_FAILED     StatusType = "FAILED"
)

type RegDocumentParam struct {
	Title        string `json:"title"`        // must
	Format       string `json:"format"`       // must，doc, docx, ppt, pptx, xls, xlsx, vsd, pot, pps, rtf, wps, et, dps, pdf, txt, epub
	TargetType   string `json:"targetType"`   // h5|image, default: h5
	Notification string `json:"notification"` // notification
	Access       string `json:"access"`       // PUBLIC|PRIVATE, default: PUBLIC
}

// String - 格式化为json格式
func (d *RegDocumentParam) String() (string, error) {
	if d.Title == "" || d.Format == "" {
		return "", errors.New("tile and format cannot be empty")
	}
	if d.TargetType == "" || (d.TargetType != DOC_TARGET_H5 && d.TargetType != DOC_TARGET_IMAGE) {
		d.TargetType = DOC_TARGET_H5
	}
	if d.Access != DOC_PUBLIC && d.Access != DOC_PRIVATE {
		d.Access = DOC_PUBLIC
	}

	j, e := json.Marshal(d)
	if e != nil {
		return "", e
	}

	// 如果 notification 为空，则去掉该参数，不然请求会报错
	if d.Notification == "" {
		var m map[string]string
		json.Unmarshal(j, &m)
		delete(m, "notification")
		j, _ = json.Marshal(m)
	}

	return string(j), nil
}

// RegDocumentResp - 注册文档请求响应
type RegDocumentResp struct {
	DocumentId  string `json:"documentId"`
	Bucket      string `json:"bucket"`
	Object      string `json:"object"`
	BosEndpoint string `json:"bosEndpoint"`
}

type GetImagesResp struct {
	Images []ImageResp `json:"images"`
}

type ImageResp struct {
	PageIndex int64  `json:"pageIndex"`
	Url       string `json:"url"`
}

type QueryDocumentParam struct {
	Https bool
}

type QueryDocumentResp struct {
	DocumentId   string            `json:"documentId"`
	Title        string            `json:"title"`
	Format       string            `json:"format"`
	TargetType   string            `json:"targetType"`
	Status       string            `json:"status"`
	UploadInfo   UploadInfoResp    `json:"uploadInfo"`
	PublishInfo  PublishInfoResp   `json:"publishInfo"`
	Notification string            `json:"notification"`
	Access       string            `json:"access"`
	CreateTime   string            `json:"createTime"`
	Error        DocumentErrorResp `json:"error"`
}

type UploadInfoResp struct {
	Bucket      string `json:"bucket"`
	Object      string `json:"object"`
	BosEndpoint string `json:"bosEndpoint"`
}

type PublishInfoResp struct {
	PageCount   int    `json:"pageCount"`
	SizeInBytes int    `json:"sizeInBytes"`
	CoverUrl    string `json:"coverUrl"`
	PublishTime string `json:"publishTime"`
}

type DocumentErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ReadDocumentParam struct {
	ExpireInSeconds int64
}

type ReadDocumentResp struct {
	DocumentId string `json:"documentId"`
	DocId      string `json:"docId"`
	Host       string `json:"host"`
	Token      string `json:"token"`
	CreateTime string `json:"createTime"`
	ExpireTime string `json:"expireTime"`
}

type ListDocumentsParam struct {
	Status  StatusType
	Marker  string
	MaxSize int64
}

func (l *ListDocumentsParam) Check() error {
	switch l.Status {
	case DOC_STATUS_UPLOADING:
	case DOC_STATUS_FAILED:
	case DOC_STATUS_PROCESSING:
	case DOC_STATUS_PUBLISHED:
	case "":
	default:
		return errors.New("invalid DOC status")
	}
	if l.MaxSize > 200 || l.MaxSize < 0 {
		return errors.New("invalid maxSize")
	}
	return nil
}

type ListDocumentsResp struct {
	Marker      string         `json:"marker"`
	IsTruncated bool           `json:"isTruncated"`
	NextMarker  string         `json:"nextMarker,omitempty"`
	Docs        []DocumentResp `json:"documents"`
}

type DocumentResp struct {
	DocumentId   string            `json:"documentId"`
	Title        string            `json:"title"`
	Format       string            `json:"format"`
	TargetType   string            `json:"targetType"`
	Status       string            `json:"status"`
	Notification string            `json:"notification"`
	Access       string            `json:"access"`
	CreateTime   string            `json:"createTime"`
	Error        DocumentErrorResp `json:"error"`
}
