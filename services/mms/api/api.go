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
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func InsertVideo(cli bce.Client, lib string, args *BaseRequest) (*BaseResponse, error) {
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(cli, http.PUT, URI_PREFIX+VIDEO_URI+"/"+lib, jsonBytes, map[string]string{})
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil

}

func GetInsertVideoResult(cli bce.Client, lib, source string) (*BaseResponse, error) {

	params := map[string]string{
		"source": source,
	}

	resp, err := sendRequest(cli, http.GET, URI_PREFIX+VIDEO_URI+"/"+lib, []byte{}, params)
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func GetInsertVideoResultById(cli bce.Client, libId, mediaId string) (*BaseResponse, error) {

	params := map[string]string{
		"mediaId":               mediaId,
		"getInsertResponseById": "",
	}

	resp, err := sendRequest(cli, http.GET, URI_PREFIX+VIDEO_URI+"/"+libId, []byte{}, params)
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteVideo(cli bce.Client, lib, source string) (*BaseResponse, error) {

	params := map[string]string{
		"source":      source,
		"deleteVideo": "",
	}

	resp, err := sendRequest(cli, http.POST, URI_PREFIX+VIDEO_URI+"/"+lib, []byte{}, params)
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteVideoById(cli bce.Client, libId, mediaId string) (*BaseResponse, error) {

	params := map[string]string{
		"mediaId":         mediaId,
		"deleteVideoById": "",
	}

	resp, err := sendRequest(cli, http.POST, URI_PREFIX+VIDEO_URI+"/"+libId, []byte{}, params)
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func InsertImage(cli bce.Client, lib string, args *BaseRequest) (*BaseResponse, error) {
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(cli, http.PUT, URI_PREFIX+IMAGE_URI+"/"+lib, jsonBytes, map[string]string{})
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteImage(cli bce.Client, lib, source string) (*BaseResponse, error) {

	params := map[string]string{
		"source":      source,
		"deleteImage": "",
	}

	resp, err := sendRequest(cli, http.POST, URI_PREFIX+IMAGE_URI+"/"+lib, []byte{}, params)
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteImageById(cli bce.Client, libId, mediaId string) (*BaseResponse, error) {

	params := map[string]string{
		"mediaId":         mediaId,
		"deleteImageById": "",
	}

	resp, err := sendRequest(cli, http.POST, URI_PREFIX+IMAGE_URI+"/"+libId, []byte{}, params)
	res := &BaseResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func SearchImageByImage(cli bce.Client, lib string, args *BaseRequest) (*SearchTaskResultResponse, error) {

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"searchByImage": "",
	}

	resp, err := sendRequest(cli, http.POST, URI_PREFIX+IMAGE_URI+"/"+lib, jsonBytes, params)
	res := &SearchTaskResultResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func SearchVideoByImage(cli bce.Client, lib string, args *BaseRequest) (*SearchTaskResultResponse, error) {

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"searchByImage": "",
	}

	resp, err := sendRequest(cli, http.POST, URI_PREFIX+VIDEO_URI+"/"+lib, jsonBytes, params)
	res := &SearchTaskResultResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func SearchVideoByVideo(cli bce.Client, lib string, args *BaseRequest) (*SearchTaskResultResponse, error) {

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"searchByVideo": "",
	}

	resp, err := sendRequest(cli, http.POST, URI_PREFIX+VIDEO_URI+"/"+lib, jsonBytes, params)
	res := &SearchTaskResultResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func GetSearchVideoByVideoResult(cli bce.Client, lib, source string) (*SearchTaskResultResponse, error) {

	params := map[string]string{
		"searchByVideo": "",
		"source":        source,
	}

	resp, err := sendRequest(cli, http.GET, URI_PREFIX+VIDEO_URI+"/"+lib, []byte{}, params)
	res := &SearchTaskResultResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func GetSearchVideoByVideoResultById(cli bce.Client, lib, taskId string) (*SearchTaskResultResponse, error) {

	params := map[string]string{
		"getSearchResponseByTaskId": "",
		"taskId":                    taskId,
	}

	resp, err := sendRequest(cli, http.GET, URI_PREFIX+VIDEO_URI+"/"+lib, []byte{}, params)
	res := &SearchTaskResultResponse{}
	if err != nil {
		return nil, err
	}

	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func sendRequest(cli bce.Client, httpMethod, url string, bodyJson []byte, params map[string]string) (*bce.BceResponse, error) {
	req := &bce.BceRequest{}
	req.SetHeader(http.CONTENT_TYPE, "application/json;charset=utf-8")
	req.SetUri(url)
	req.SetMethod(httpMethod)
	req.SetParams(params)

	body, err := bce.NewBodyFromBytes(bodyJson)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	return resp, nil
}
