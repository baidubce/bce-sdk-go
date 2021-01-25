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

package mms

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/mms/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var CLIENT *Client

const (
	AK           = ""
	SK           = ""
	VIDEO_LIB    = ""
	IMAGE_LIB    = ""
	VIDEO_SOURCE = "http://xxx.mp4"
	IMAGE_SOURCE = "http://xxx.png"
	ENDPOINT     = "http://xxx"
)

func init() {
	CLIENT, _ = NewClient(AK, SK, ENDPOINT)
	log.SetLogHandler(log.STDERR)
	log.SetLogLevel(log.DEBUG)
}

func TestInsertVideo(t *testing.T) {
	args := &api.BaseRequest{Source: VIDEO_SOURCE, Notification: "test", Description: "..."}
	res, err := CLIENT.InsertVideo(VIDEO_LIB, args)
	if err != nil {
		t.Fatalf("InsertVideo failed: %s", err)
	}
	if res == nil {
		t.Fatal("TestInsertVideo Failed. ")
	}
}

func TestGetInsertVideoResult(t *testing.T) {
	res, err := CLIENT.GetInsertVideoResult(VIDEO_LIB, VIDEO_SOURCE)
	if err != nil {
		t.Fatalf("TestGetInsertVideoResult failed: %s", err)
	}
	if res.Source != VIDEO_SOURCE {
		t.Fatal("TestGetInsertVideoResult Failed.")
	}
}

func TestDeleteVideo(t *testing.T) {
	res, err := CLIENT.DeleteVideo(VIDEO_LIB, VIDEO_SOURCE)
	if err != nil {
		t.Fatalf("TestDeleteVideo failed: %s", err)
	}
	if res == nil {
		t.Fatal("TestDeleteVideo Failed.")
	}
}

func TestInsertImage(t *testing.T) {
	args := &api.BaseRequest{Source: IMAGE_SOURCE, Notification: "test", Description: "..."}
	res, err := CLIENT.InsertImage(IMAGE_LIB, args)
	if err != nil {
		t.Fatalf("TestInsertImage failed: %s", err)
	}
	if res == nil {
		t.Fatal("TestInsertImage Failed. ")
	}
}

func TestDeleteImage(t *testing.T) {
	res, err := CLIENT.DeleteVideo(IMAGE_LIB, IMAGE_SOURCE)
	if err != nil {
		t.Fatalf("TestDeleteImage failed: %s", err)
	}
	if res == nil {
		t.Fatal("TestDeleteImage Failed.")
	}
}

func TestSearchImageByImage(t *testing.T) {
	args := &api.BaseRequest{Source: IMAGE_SOURCE, Notification: "test", Description: "..."}
	res, err := CLIENT.SearchImageByImage(IMAGE_LIB, args)
	if err != nil {
		t.Fatalf("TestSearchImageByImage failed: %s", err)
	}
	if res == nil {
		t.Fatal(res)
	}
}

func TestSearchVideoByImage(t *testing.T) {
	args := &api.BaseRequest{Source: IMAGE_SOURCE, Notification: "test", Description: "..."}
	res, err := CLIENT.SearchVideoByImage(VIDEO_LIB, args)
	if err != nil {
		t.Fatalf("TestSearchVideoByImage failed: %s", err)
	}
	if res == nil {
		t.Fatal(res)
	}
}

func TestSearchVideoByVideo(t *testing.T) {
	args := &api.BaseRequest{Source: VIDEO_SOURCE, Notification: "test", Description: "..."}
	res, err := CLIENT.SearchVideoByVideo(VIDEO_LIB, args)
	if err != nil {
		t.Fatalf("TestSearchVideoByVideo failed: %s", err)
	}
	if res == nil {
		t.Fatal(res)
	}
}

func TestGetSearchVideoByVideoResult(t *testing.T) {
	res, err := CLIENT.GetSearchVideoByVideoResult(VIDEO_LIB, VIDEO_SOURCE)
	if err != nil {
		t.Fatalf("TestGetSearchVideoByVideoResult failed: %s", err)
	}
	if res == nil {
		t.Fatal(res)
	}
}
