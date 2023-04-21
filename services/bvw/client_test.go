package bvw

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bvw/api"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	BVW_CLIENT *Client
)

type Conf struct {
	AK string
	SK string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given: %+v\n", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	BVW_CLIENT, _ = NewClient(confObj.AK, confObj.SK, "")
	log.SetLogLevel(log.DEBUG)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestUploadMaterial(t *testing.T) {
	req := &api.MatlibUploadRequest{}
	req.Key = "dog.mp4"
	req.Bucket = "bvw-console"
	req.Title = "split_result_267392.mp4"
	req.MediaType = "video"
	uploadResponse, err := BVW_CLIENT.UploadMaterial(req)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", uploadResponse)
}

func TestSearchMaterial(t *testing.T) {
	req := &api.MaterialSearchRequest{}
	req.Size = 5
	req.PageNo = 1
	req.Status = "FINISHED"
	req.MediaType = "video"
	req.Begin = "2023-01-11T16:00:00Z"
	req.End = "2023-04-12T15:59:59Z"
	materialResp, err := BVW_CLIENT.SearchMaterial(req)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", materialResp)
}

func TestGetMaterial(t *testing.T) {
	materialGetResponse, err := BVW_CLIENT.GetMaterial("d9b9f08ef1e0a28967fa0f7b5819db30")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", materialGetResponse)
}

func TestDeleteMaterial(t *testing.T) {
	err := BVW_CLIENT.DeleteMaterial("d9b9f08ef1e0a28967fa0f7b5819db30")
	ExpectEqual(t.Errorf, err, nil)
}

func TestUploadMaterialPreset(t *testing.T) {
	req := &api.MatlibUploadRequest{}
	req.Key = "item2.jpeg"
	req.Bucket = "bvw-console"
	req.Title = "item2.jpeg"
	req.MediaType = "image"
	uploadResponse, err := BVW_CLIENT.UploadMaterialPreset("PICTURE", req)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", uploadResponse)
}

func TestSearchMaterialPreset(t *testing.T) {
	req := &api.MaterialPresetSearchRequest{}
	req.PageSize = "10"
	req.Status = "FINISHED"
	req.PageNo = "1"
	req.SourceType = "USER"
	req.MediaType = "PICTURE"
	response, err := BVW_CLIENT.SearchMaterialPreset(req)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetMaterialPreset(t *testing.T) {
	response, err := BVW_CLIENT.GetMaterialPreset("cc0aabdc71421abaa17e80a26caa009f")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestDeleteMaterialPreset(t *testing.T) {
	err := BVW_CLIENT.DeleteMaterialPreset("cc0aabdc71421abaa17e80a26caa009f")
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateMatlibConfig(t *testing.T) {
	request := &api.MatlibConfigBaseRequest{
		Bucket: "go-test"}
	response, err := BVW_CLIENT.CreateMatlibConfig(request)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetMatlibConfig(t *testing.T) {
	response, err := BVW_CLIENT.GetMatlibConfig()
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestUpdateMatlibConfig(t *testing.T) {
	err := BVW_CLIENT.UpdateMatlibConfig(&api.MatlibConfigUpdateRequest{Bucket: "go-test"})
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateDraft(t *testing.T) {
	response, err := BVW_CLIENT.CreateDraft(&api.CreateDraftRequest{
		Duration: "0",
		Titile:   "testCreateDraft3",
		Ratio:    "hori16x9"})
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetSingleDraft(t *testing.T) {
	response, err := BVW_CLIENT.GetSingleDraft(1017834)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetDraftList(t *testing.T) {
	response, err := BVW_CLIENT.GetDraftList(&api.DraftListRequest{
		PageNo:    1,
		PageSize:  20,
		BeginTime: "2023-01-11T16:00:00Z",
		EndTime:   "2023-04-12T15:59:59Z"})
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestUpdateDraft(t *testing.T) {
	var request api.MatlibTaskRequest
	jsonStr := "{\"title\":\"updatesucess\",\"draftId\":\"1017890\",\"timeline\":{\"timeline\":{\"video\":[{\"key\"" +
		":\"name\",\"isMaster\":true,\"list\":[{\"type\":\"video\",\"start\":0,\"end\":5.859375,\"showStart\":0," +
		"\"showEnd\":5.859375,\"xpos\":0,\"ypos\":0,\"width\":1280,\"height\":720,\"duration\":5.859375,\"extInfo" +
		"\":{\"style\":\"\",\"lightness\":0,\"gifMode\":0,\"contrast\":0,\"saturation\":0,\"hue\":0,\"speed\":1" +
		",\"transitionStart\":\"\",\"transitionEnd\":\"black\",\"transitionDuration\":1,\"volume\":1,\"rotate\":0," +
		"\"mirror\":\"\",\"blankArea\":\"\"},\"mediaInfo\":{\"fileType\":\"video\",\"sourceType\":\"USER\",\"" +
		"sourceUrl\":\"https://bj.bcebos.com/v1/bvw-console/360p/dog.mp4?x-bce-security-token=" +
		"ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7x" +
		"G%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hb" +
		"SJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRp" +
		"OC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ" +
		"%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2" +
		"F%2Fb227edbf73344bdfc9fed00ba491c5c0c3abe229792d7b3d026604cfbe541b68\",\"audioUrl\":\"" +
		"https://bj.bcebos.com/v1/bvw-console/audio/dog.mp3?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjY" +
		"yYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ" +
		"5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3" +
		"OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzT" +
		"Cgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=" +
		"bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F3dcc823c9497aca1154f" +
		"f0007eca86af4e682363c3ceddba0b3c74ca14e2d154\",\"bucket\":\"bvw-console\",\"key\":\"dog.mp4\",\"audioKey\":" +
		"\"audio/dog.mp3\",\"coverImage\":\"https://bj.bcebos.com/v1/bvw-console/thumbnail/dog00000500.jpg" +
		"?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZp" +
		"Ppsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQz" +
		"FlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6P" +
		"z36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQj" +
		"Wm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cb" +
		"c0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F21a92744dfb9fc3f46745e75d095da327bb04677f9028fb85e00ff5dc7df6" +
		"daf\",\"duration\":18.73,\"width\":1920,\"height\":1080,\"status\":\"FINISHED\",\"name\":\"dog.mp4\"," +
		"\"thumbnailPrefix\":\"\",\"thumbnailKeys\":[\"thumbnail/dog00000500.jpg\"],\"mediaId\":\"" +
		"1f10ce0db10b8eb5b2f2755daf544900\",\"offstandard\":false},\"uid\":\"a081f1c6-9dc9-4e7b-a00e-6eb5217e771d" +
		"\"},{\"type\":\"video\",\"start\":5.859375,\"end\":18.73,\"showStart\":5.859375,\"showEnd\":18.73,\"xpos" +
		"\":0,\"ypos\":0,\"width\":1280,\"height\":720,\"duration\":12.870625,\"extInfo\":{\"style\":\"\",\"lightness" +
		"\":0,\"gifMode\":0,\"contrast\":0,\"saturation\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"black\",\"" +
		"transitionEnd\":\"\",\"transitionDuration\":1,\"volume\":1,\"rotate\":0,\"mirror\":\"\",\"blankArea\":\"\"}," +
		"\"mediaInfo\":{\"fileType\":\"video\",\"sourceType\":\"USER\",\"sourceUrl\":" +
		"\"https://bj.bcebos.com/v1/bvw-console/360p/dog.mp4?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyN" +
		"jk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9" +
		"RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7D" +
		"GO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df" +
		"5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%" +
		"3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2Fb" +
		"227edbf73344bdfc9fed00ba491c5c0c3abe229792d7b3d026604cfbe541b68\",\"audioUrl\":" +
		"\"https://bj.bcebos.com/v1/bvw-console/audio/dog.mp3?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjY" +
		"yYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5" +
		"Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3O" +
		"HDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCg" +
		"h8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=" +
		"bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F3dcc823c9497aca1" +
		"154ff0007eca86af4e682363c3ceddba0b3c74ca14e2d154\",\"bucket\":\"bvw-console\",\"key\":\"dog.mp4\",\"" +
		"audioKey\":\"audio/dog.mp3\",\"coverImage\":\"https://bj.bcebos.com/v1/bvw-console/thumbnail/dog00000500" +
		".jpg?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZp" +
		"Ppsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQz" +
		"FlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz" +
		"36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm" +
		"231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0" +
		"fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F21a92744dfb9fc3f46745e75d095da327bb04677f9028fb85e00ff5dc7df6da" +
		"f\",\"duration\":18.73,\"width\":1920,\"height\":1080,\"status\":\"FINISHED\",\"name\":\"dog.mp4\",\"thumb" +
		"nailPrefix\":\"\",\"thumbnailKeys\":[\"thumbnail/dog00000500.jpg\"],\"mediaId\":\"1f10ce0db10b8eb5b2f2755d" +
		"af544900\",\"offstandard\":false},\"uid\":\"70af482e-0bf5-4c38-8f05-03c9e3b8ae03\"}],\"unlinkMaster\":true" +
		"}],\"audio\":[{\"key\":\"\",\"isMaster\":false,\"list\":[{\"start\":0,\"end\":155.99,\"showStart\":" +
		"0.234375,\"showEnd\":156.224375,\"duration\":155.99,\"xpos\":0,\"ypos\":0,\"hidden\":false,\"mediaInfo\"" +
		":{\"fileType\":\"audio\",\"sourceUrl\":\"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset" +
		"/music/audio/%E5%8F%A4%E9%A3%8E%E9%A3%98%E6%89%AC.mp3?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772" +
		"a1409590%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F2ddcf78c92de29ae3d7c3166a4e17e7c5d07fa38dcefd24c29c4f4d5b" +
		"5ba46fe\",\"audioUrl\":\"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset/music/audio/" +
		"%E5%8F%A4%E9%A3%8E%E9%A3%98%E6%89%AC.mp3?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772a1409590%2F2" +
		"023-04-14T07%3A16%3A35Z%2F86400%2F%2F2ddcf78c92de29ae3d7c3166a4e17e7c5d07fa38dcefd24c29c4f4d5b5ba46fe\"," +
		"\"bucket\":\"videoworks-system-preprocess\",\"key\":\"systemPreset/music/古风飘扬.aac\",\"audioKey\":\"" +
		"systemPreset/music/audio/古风飘扬.mp3\",\"coverImage\":\"\",\"duration\":155.99,\"name\":\"\",\"thumbnailList" +
		"\":[],\"mediaId\":\"\",\"offstandard\":false},\"type\":\"audio\",\"uid\":\"bd52be7f-1f19-4368-8c41-44e991af" +
		"8164\",\"name\":\"古风飘扬\",\"extInfo\":{\"style\":\"\",\"lightness\":0,\"gifMode\":0,\"contrast\":0,\"" +
		"saturation\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"\",\"transitionEnd\":\"\",\"transitionDuration" +
		"\":1,\"volume\":1,\"rotate\":0},\"boxDataLeft\":4,\"dragBoxWidth\":1996.6720000000003,\"lineType\":\"audio" +
		"\"}]}],\"subtitle\":[{\"key\":\"\",\"list\":[{\"duration\":3,\"hidden\":false,\"name\":\"time-place\",\"" +
		"tagExtInfo\":{\"marginBottom\":0,\"textFadeIn\":1,\"textFadeOut\":1,\"textOutMaskDur\":1},\"showStart\":" +
		"5.859375,\"showEnd\":8.859,\"templateId\":\"6764ce3331ea7e406e4ab4475d1dff18\",\"type\":\"subtitle\",\"uid" +
		"\":\"5aaa35f4-8fae-4b8c-b7ed-761a54550244\",\"xpos\":\"0\",\"ypos\":\"309\",\"config\":[{\"alpha\":0,\"" +
		"fontColor\":\"#ffffff\",\"fontSize\":50,\"fontStyle\":\"normal\",\"fontType\":\"方正时代宋 简 Extrabold\"" +
		",\"lineHeight\":1.2,\"name\":\"时间\",\"text\":\"haha\",\"backgroundColor\":\"#2468F2\",\"backgroundAlpha" +
		"\":0,\"fontx\":0,\"fonty\":0,\"invisible\":false},{\"alpha\":0,\"fontColor\":\"#000000\",\"fontSize\":50," +
		"\"fontStyle\":\"normal\",\"fontType\":\"方正时代宋 简 Extrabold\",\"lineHeight\":1.2,\"name\":\"地点\",\"" +
		"text\":\"cd\",\"backgroundColor\":\"#ffffff\",\"backgroundAlpha\":0,\"fontx\":0,\"fonty\":0,\"invisible\"" +
		":false}],\"boxDataLeft\":76,\"dragBoxWidth\":38.400000000000006,\"lineType\":\"subtitle\"}],\"master\":" +
		"false}],\"sticker\":[{\"key\":\"\",\"isMaster\":false,\"list\":[{\"showStart\":0,\"showEnd\":3,\"duration\"" +
		":3,\"xpos\":0,\"ypos\":0,\"width\":215,\"height\":120.9375,\"hidden\":false,\"mediaInfo\":{\"sourceUrl\":\"" +
		"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset/picture/%E9%9D%A2%E5%8C%85%E3%80%81%E8" +
		"%82%A0.png?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772a1409590%2F2023-04-14T07%3A16%3A35Z%2F8640" +
		"0%2F%2F84867e4874cc94eb374898017cb4367ed8c24a5750a9d8ebd14d8ca989cf2e53\",\"audioUrl\":\"" +
		"https://bj.bcebos.com/v1/videoworks-system-preprocess/systemPreset/picture/audio/%E9%9D%A2%E5%8C%85%E3%80%" +
		"81%E8%82%A0.mp3?authorization=bce-auth-v1%2F66c557960e7a4822bd82c772a1409590%2F2023-04-14T07%3A16%3A35Z%2F" +
		"86400%2F%2F1807d3005f5f8fc270fa17238ec550c20162252c19952e6a45ae497ef7148086\",\"bucket\":\"" +
		"videoworks-system-preprocess\",\"key\":\"systemPreset/picture/面包、肠.png\",\"audioKey\":\"systemPreset" +
		"/picture/audio/面包、肠.mp3\",\"coverImage\":\"\",\"width\":215,\"height\":120,\"name\":\"\",\"thumbnailList" +
		"\":[],\"mediaId\":\"\",\"offstandard\":false},\"type\":\"image\",\"uid\":\"e419b583-aedb-4265-" +
		"9a67-7e21fc621f85\",\"name\":\"面包、肠\",\"extInfo\":{\"style\":\"\",\"lightness\":0,\"gifMode\":0," +
		"\"contrast\":0,\"saturation\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"\",\"transitionEnd\":\"\"," +
		"\"transitionDuration\":1,\"volume\":1,\"rotate\":0},\"lineType\":\"sticker\",\"boxDataLeft\":1,\"" +
		"dragBoxWidth\":38.400000000000006}]}]}},\"ratio\":\"hori16x9\",\"resourceList\":[{\"id\":\"1f10ce0db10" +
		"b8eb5b2f2755daf544900\",\"userId\":\"e7e47aa53fbb47dfb1e4c86424bb7ad3\",\"mediaType\":\"video\",\"" +
		"sourceType\":\"USER\",\"status\":\"FINISHED\",\"title\":\"dog.mp4\",\"sourceUrl\":\"https://bj.bcebos.com" +
		"/v1/bvw-console/dog.mp4?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM" +
		"1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Z" +
		"zgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkh" +
		"MLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7" +
		"D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1" +
		"a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F4a3c399db912e35f6b6c008faffebe5752c47e36fcd21d4bf03" +
		"bc908c3a29e5e\",\"sourceUrl360p\":\"https://bj.bcebos.com/v1/bvw-console/360p/dog.mp4?x-bce-security-toke" +
		"n=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZT" +
		"NAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwb" +
		"rCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F" +
		"9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujH" +
		"KfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07" +
		"%3A16%3A35Z%2F86400%2F%2Fb227edbf73344bdfc9fed00ba491c5c0c3abe229792d7b3d026604cfbe541b68\",\"audioUrl\"" +
		":\"https://bj.bcebos.com/v1/bvw-console/audio/dog.mp3?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1Z" +
		"jYyYjlkZjNjODB8AAAAAM0BAABgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8" +
		"sJ5Dc75AlLyVko5U4CFsiaEE9xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46Nx" +
		"S3OHDvxtr95gIpW592MxArSISjn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOF" +
		"zTCgh8qeej6RXfYjBKn0pvmWCKr%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorizatio" +
		"n=bce-auth-v1%2F4a2cac88da9411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F3dcc823c9497aca1" +
		"154ff0007eca86af4e682363c3ceddba0b3c74ca14e2d154\",\"thumbnailList\":[\"https://bj.bcebos.com/v1/bvw-cons" +
		"ole/thumbnail/dog00000500.jpg?x-bce-security-token=ZjkyZmQ2YmQxZTQ3NDcyNjk0ZTg1ZjYyYjlkZjNjODB8AAAAAM0BAA" +
		"BgAa0YM1kG5uQI39UZkqCZpPpsi8DEL63qLoYtl2x5OFqZTNAWS7xG%2FfhP%2BlWF9RNJhYFABpfrg8sJ5Dc75AlLyVko5U4CFsiaEE9" +
		"xGdGQU4r3Zzgl1fJothQzFlDKfhH9hh9NXykFPkd4OXwbrCmrl902hbSJu8e6Q7DGO0tOi444b9K46NxS3OHDvxtr95gIpW592MxArSIS" +
		"jn%2FpMVkhMLtymxh6Pz36iVdo0ErJnD1JIozvKo%2F9bV7pIjpIAysjRpOC8Df5Mh5cSG96BBwftUOFzTCgh8qeej6RXfYjBKn0pvmWCK" +
		"r%2BM6bV7D39wKiQjWm231giBr3teGDbG%2BfujHKfC4tNAYpzSrCwEFCyCQ%3D%3D&authorization=bce-auth-v1%2F4a2cac88da9" +
		"411edaf1a4f67d1cbc0fc%2F2023-04-14T07%3A16%3A35Z%2F86400%2F%2F21a92744dfb9fc3f46745e75d095da327bb04677f902" +
		"8fb85e00ff5dc7df6daf\"],\"subtitleUrls\":[],\"createTime\":\"2023-04-11 16:55:32\",\"updateTime\":\"2023-0" +
		"4-11 16:55:43\",\"duration\":18.73,\"height\":1080,\"width\":1920,\"fileSizeInByte\":8948434,\"thumbnailK" +
		"eys\":[\"thumbnail/dog00000500.jpg\"],\"subtitles\":[\"\"],\"bucket\":\"bvw-console\",\"key\":\"dog.mp4\"" +
		",\"key360p\":\"360p/dog.mp4\",\"key720p\":\"720p/dog.mp4\",\"audioKey\":\"audio/dog.mp3\",\"used\":true}]" +
		",\"coverBucket\":\"bvw-console\",\"coverKey\":\"thumbnail/dog00000500.jpg\"}"
	jsonErr := json.Unmarshal([]byte(jsonStr), &request)
	ExpectEqual(t.Errorf, jsonErr, nil)
	t.Logf("%+v", &request)
	err := BVW_CLIENT.UpdateDraft(1017890, &request)
	ExpectEqual(t.Errorf, err, nil)
}

func TestPollingVideoEdit(t *testing.T) {
	response, err := BVW_CLIENT.PollingVideoEdit(267539)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestCreateVideoEdit(t *testing.T) {
	var request api.VideoEditCreateRequest
	jsonStr := "{\"title\":\"新建作品-202304141603\",\"taskId\":\"1017895\",\"bucket\":\"vwdemo\",\"cmd\":{\"timeline" +
		"\":{\"video\":[{\"list\":[{\"type\":\"video\",\"start\":0,\"end\":7.65625,\"showStart\":0,\"showEnd\":" +
		"7.65625,\"xpos\":0,\"ypos\":0,\"width\":1,\"height\":1,\"duration\":7.65625,\"extInfo\":{\"style\":\"\"" +
		",\"lightness\":0,\"contrast\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"\",\"transitionEnd\":\"black\"" +
		",\"transitionDuration\":1,\"mirror\":\"\",\"rotate\":0,\"blankArea\":\"\",\"volume\":1},\"mediaInfo\":{\"" +
		"mediaId\":\"1f10ce0db10b8eb5b2f2755daf544900\",\"key\":\"dog.mp4\",\"bucket\":\"bvw-console\",\"fileType\"" +
		":\"video\",\"width\":1920,\"height\":1080}},{\"type\":\"video\",\"start\":7.65625,\"end\":18.73,\"showStart" +
		"\":7.65625,\"showEnd\":18.73,\"xpos\":0,\"ypos\":0,\"width\":1,\"height\":1,\"duration\":11.07375,\"" +
		"extInfo\":{\"style\":\"\",\"lightness\":0,\"contrast\":0,\"hue\":0,\"speed\":1,\"transitionStart\":\"black\"" +
		",\"transitionEnd\":\"\",\"transitionDuration\":1,\"mirror\":\"\",\"rotate\":0,\"blankArea\":\"\",\"volume\":" +
		"1},\"mediaInfo\":{\"mediaId\":\"1f10ce0db10b8eb5b2f2755daf544900\",\"key\":\"dog.mp4\",\"bucket\":\"" +
		"bvw-console\",\"fileType\":\"video\",\"width\":1920,\"height\":1080}}]}],\"audio\":[{\"list\":[{\"name\":" +
		"\"古风飘扬\",\"start\":0,\"end\":155.99,\"duration\":155.99,\"showStart\":0.078125,\"showEnd\":156.068125,\"" +
		"uid\":\"cc8c1ecc-fcd3-493d-be5b-8cce8c15ed15\",\"extInfo\":{\"volume\":\"1.0\",\"transitions\":[]},\"" +
		"mediaInfo\":{\"fileType\":\"audio\",\"key\":\"systemPreset/music/古风飘扬.aac\",\"bucket\":\"videoworks-" +
		"system-preprocess\",\"name\":\"古风飘扬\"}}]}],\"subtitle\":[{\"list\":[{\"templateId\":\"6764ce3331ea7e406e" +
		"4ab4475d1dff18\",\"showStart\":0,\"showEnd\":3,\"duration\":3,\"uid\":\"05e59686-b23e-4b33-96d7-040eab63" +
		"85b6\",\"tag\":\"time-place\",\"xpos\":0,\"ypos\":0.431,\"config\":[{\"text\":\"时间\",\"fontSize\":50,\"" +
		"fontType\":\"方正时代宋 简 Extrabold\",\"fontColor\":\"#ffffff\",\"alpha\":0,\"fontStyle\":\"normal\",\"" +
		"backgroundColor\":\"#2468F2\",\"backgroundAlpha\":0,\"fontx\":0.039,\"fonty\":0.028,\"rectx\":0,\"recty\"" +
		":0.431,\"rectWidth\":0.156,\"rectHeight\":0.139},{\"text\":\"地点\",\"fontSize\":50,\"fontType\":\"" +
		"方正时代宋 简 Extrabold\",\"fontColor\":\"#000000\",\"alpha\":0,\"fontStyle\":\"normal\",\"backgroundColor" +
		"\":\"#ffffff\",\"backgroundAlpha\":0,\"fontx\":0.039,\"fonty\":0.028,\"rectx\":0.156,\"recty\":0.431,\"" +
		"rectWidth\":0.156,\"rectHeight\":0.139}],\"tagExtInfo\":{\"glExtInfo\":{}}}]}],\"sticker\":[{\"list\":" +
		"[{\"type\":\"image\",\"showStart\":0,\"showEnd\":3,\"duration\":3,\"xpos\":0,\"ypos\":0,\"width\":0.168" +
		",\"height\":0.168,\"extInfo\":{},\"mediaInfo\":{\"key\":\"systemPreset/picture/面包、肠.png\",\"bucket\":" +
		"\"videoworks-system-preprocess\",\"width\":215,\"height\":120.9375,\"fileType\":\"image\"}}]}]}},\"extInfo" +
		"\":{\"aspect\":\"hori16x9\",\"resolutionRatio\":\"v720p\"}}"
	jsonErr := json.Unmarshal([]byte(jsonStr), &request)
	ExpectEqual(t.Errorf, jsonErr, nil)
	t.Logf("%+v", &request)
	response, err := BVW_CLIENT.CreateVideoEdit(&request)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}
