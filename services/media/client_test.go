package media

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/media/api"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	MEDIA_CLIENT *Client
)

type Conf struct {
	AK       string
	SK       string
	Endpoint string
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

	MEDIA_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
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

func TestCreatePipline(t *testing.T) {
	err := MEDIA_CLIENT.CreatePipeline("test1", "go-sdk-test", "go-sdk-test", 10)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%s", "done")
}

func TestListPipelines(t *testing.T) {
	pipelines, err := MEDIA_CLIENT.ListPipelines()
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", pipelines)
}

func TestGetPipeline(t *testing.T) {
	pipeline, err := MEDIA_CLIENT.GetPipeline("go_sdk_test")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", pipeline)
}

func TestDeletePipeline(t *testing.T) {
	err := MEDIA_CLIENT.DeletePipeline("test11")
	ExpectEqual(t.Errorf, err, nil)
}

func TestUpdatePipeline(t *testing.T) {
	args, _ := MEDIA_CLIENT.GetPipelineUpdate("test1")
	args.Description = "update"
	args.TargetBucket = "vwdemo"
	args.SourceBucket = "vwdemo"

	config := &api.UpdatePipelineConfig{}
	config.Capacity = 2
	config.Notification = "zz"
	args.UpdatePipelineConfig = config
	err := MEDIA_CLIENT.UpdatePipeline("test1", args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateJob(t *testing.T) {
	jobResponse, err := MEDIA_CLIENT.CreateJob("go_sdk_test", "01.mp4", "01_go_02.mp4",
		"videoworks_system_preprocess_360p")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestCreateJobCustomize(t *testing.T) {
	args := &api.CreateJobArgs{}
	args.PipelineName = "go_sdk_test"
	source := &api.Source{Clips: &[]api.SourceClip{{
		SourceKey:             "01.mp4",
		EnableDelogo:          false,
		DurationInMillisecond: 6656,
		StartTimeInSecond:     2}}}
	args.Source = source
	target := &api.Target{}
	targetKey := "clips_playback_watermark_delogo_crop2.mp4"
	watermarkId := "wmk-pcgqidaj13iv1eyf"
	target.TargetKey = targetKey
	watermarkIdSlice := append(target.WatermarkIds, watermarkId)
	target.WatermarkIds = watermarkIdSlice
	presetName := "go_test_customize_audio_video"
	target.PresetName = presetName

	delogoArea := &api.Area{}
	delogoArea.X = 10
	delogoArea.Y = 10
	delogoArea.Width = 30
	delogoArea.Height = 40
	target.DelogoArea = delogoArea

	args.Target = target

	jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestCreateJobCustomizeFont(t *testing.T) {
	args := &api.CreateJobArgs{}
	args.PipelineName = "go_sdk_test"
	source := &api.Source{Clips: &[]api.SourceClip{{
		SourceKey:             "01.mp4",
		EnableDelogo:          false,
		DurationInMillisecond: 6656,
		StartTimeInSecond:     2}}}
	args.Source = source
	target := &api.Target{}
	targetKey := "clips_playback_watermark_delogo_crop2.mp4"
	target.TargetKey = targetKey
	presetName := "go_test_customize_audio_video"
	target.PresetName = presetName

	font := &api.Font{Color: "#FFFFFF"}
	text := "hello world"
	insert := &api.Insert{
		Text: &text,
		Font: font}
	target.Inserts = &[]api.Insert{*insert}
	args.Target = target

	jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestCreateJobCustomizeDelogoCrop(t *testing.T) {
	args := &api.CreateJobArgs{}
	args.PipelineName = "go_sdk_test"
	source := &api.Source{Clips: &[]api.SourceClip{{
		SourceKey:             "01.mp4",
		EnableDelogo:          false,
		DurationInMillisecond: 6656,
		StartTimeInSecond:     2}}}
	args.Source = source
	target := &api.Target{}
	targetKey := "clips_playback_watermark_delogo_crop2.mp4"
	watermarkId := "wmk-pcgqidaj13iv1eyf"
	target.TargetKey = targetKey
	watermarkIdSlice := append(target.WatermarkIds, watermarkId)
	target.WatermarkIds = watermarkIdSlice
	presetName := "go_test_customize_audio_video"
	target.PresetName = presetName

	delogoArea := &api.Area{}
	delogoArea.X = 10
	delogoArea.Y = 10
	delogoArea.Width = 30
	delogoArea.Height = 40
	target.DelogoArea = delogoArea

	args.Target = target

	jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestCreateJobWithNotify(t *testing.T) {
	args := &api.CreateJobArgs{}
	args.PipelineName = "go_sdk_test"
	source := &api.Source{Clips: &[]api.SourceClip{{
		SourceKey: "01.mp4",
	}}}
	args.Source = source
	target := &api.Target{}
	targetKey := "clips_playback_watermark_delogo_crop2.mp4"
	target.TargetKey = targetKey
	presetName := "go_test_customize_audio_video"
	target.PresetName = presetName
	cfg := &api.JobCfg{}
	cfg.Notification = "http://hello.lyq.com:80/job"
	target.JobCfg = cfg
	args.Target = target

	jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestCreateJobMetaDate(t *testing.T) {
	args := &api.CreateJobArgs{}
	args.PipelineName = "normal"
	source := &api.Source{Clips: &[]api.SourceClip{{
		SourceKey: "123123.mp4",
	}}}
	args.Source = source
	target := &api.Target{}
	targetKey := "metadata123123.mp4"
	target.TargetKey = targetKey
	presetName := "mct.video_mp4_640x360_600kbps"
	target.PresetName = presetName
	data := "{\"test\":\"hello\",\"data\":1}"
	target.MetaData = data
	args.Target = target

	jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestListTranscodingJobs(t *testing.T) {
	listTranscodingJobsResponse, err := MEDIA_CLIENT.ListTranscodingJobs("go_sdk_test")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", listTranscodingJobsResponse)
}

func TestGetTranscodingJob(t *testing.T) {
	getTranscodingJobResponse, err := MEDIA_CLIENT.GetTranscodingJob("job-qfan7kgbfh6tc2pp")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", getTranscodingJobResponse)
	t.Logf("%+v", getTranscodingJobResponse.Target.JobCfg.Notification)
}

func TestGetTranscodingJobOutput(t *testing.T) {
	getTranscodingJobResponse, err := MEDIA_CLIENT.GetTranscodingJob("job-rgwn1ijwmkmwca5i")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", getTranscodingJobResponse)
	t.Logf("%+v", getTranscodingJobResponse.JobOutputInfo)
}

func TestListPresets(t *testing.T) {
	listPresetsResponse, err := MEDIA_CLIENT.ListPresets()
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", listPresetsResponse)
}

func TestGetPreset(t *testing.T) {
	getPresetResponse, err := MEDIA_CLIENT.GetPreset("q_go_test_customize_audio_video")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", getPresetResponse)
	video := getPresetResponse.Video
	t.Logf("%+v", video)
	audio := getPresetResponse.Audio
	t.Logf("%+v", audio)
}

func TestCreatePreset(t *testing.T) {
	err := MEDIA_CLIENT.CreatePreset("go_sdk_test_preset3", "测试go创建模板3", "mp4")
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestCustomizeAudio(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "go_test_customize"
	preset.Description = "自定义创建模板"
	preset.Container = "mp4"

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	preset.Audio = audio

	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestCustomizeAudioEncryptionC(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "go_test_customize_encryption_clip"
	preset.Description = "自定义创建模板"
	preset.Container = "mp3"

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	preset.Audio = audio

	clip := &api.Clip{}
	clip.StartTimeInSecond = 2
	clip.DurationInSecond = 10
	preset.Clip = clip

	encryption := &api.Encryption{}
	encryption.Strategy = "PlayerBinding"
	preset.Encryption = encryption

	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestCustomizeAudioEncryption(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "go_test_customize_encryption"
	preset.Description = "自定义创建模板"
	preset.Container = "mp3"

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	preset.Audio = audio

	encryption := &api.Encryption{}
	encryption.Strategy = "PlayerBinding"
	preset.Encryption = encryption

	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestCustomizeAudioVideo(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "q_go_test_customize_audio_video"
	preset.Description = "自定义创建模板"
	preset.Container = "mp4"

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	audio.Codec = "opus"
	preset.Audio = audio

	video := &api.Video{}
	//video.BitRateInBps = 1024000
	video.MaxHeigtInPixel = 1920
	video.MaxWidthInPixel = 1280
	video.CodecEnhance = true
	video.Crf = 23
	video.RateControl = "crf"
	preset.Video = video

	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestCustomizeClpAudVidEncWat(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "go_test_customize_clp_aud_vid_en_wat"
	preset.Description = "自定义创建模板"
	preset.Container = "mp4"

	clip := &api.Clip{}
	clip.StartTimeInSecond = 0
	clip.DurationInSecond = 60
	preset.Clip = clip

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	preset.Audio = audio

	video := &api.Video{}
	video.BitRateInBps = 1024000
	preset.Video = video

	encryption := &api.Encryption{}
	encryption.Strategy = "PlayerBinding"
	preset.Encryption = encryption

	preset.WatermarkID = "wmk-pc0rdhzbm8ff99qw"

	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestCustomizeFullArgs(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "go_test_customize_full_args"
	preset.Description = "全参数"
	preset.Container = "hls"
	preset.Transmux = false

	clip := &api.Clip{}
	clip.StartTimeInSecond = 0
	clip.DurationInSecond = 60
	preset.Clip = clip

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	preset.Audio = audio

	video := &api.Video{}
	video.BitRateInBps = 1024000
	preset.Video = video

	encryption := &api.Encryption{}
	encryption.Strategy = "PlayerBinding"
	preset.Encryption = encryption

	water := &api.Watermarks{}
	water.Image = []string{"wmk-pc0rdhzbm8ff99qw"}
	preset.Watermarks = water

	transCfg := &api.TransCfg{}
	transCfg.TransMode = "normal"
	preset.TransCfg = transCfg

	extraCfg := &api.ExtraCfg{}
	extraCfg.SegmentDurationInSecond = 6.66
	preset.ExtraCfg = extraCfg

	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestClient_UpdatePreset(t *testing.T) {
	preset, _ := MEDIA_CLIENT.GetPreset("go_test_customize")
	preset.Description = "测试update-v2"
	err := MEDIA_CLIENT.UpdatePreset(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetMediaInfoOfFile(t *testing.T) {
	info, err := MEDIA_CLIENT.GetMediaInfoOfFile("go-test", "01.mp4")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", info)
	videoInfo := info.VideoInfo
	t.Logf("%+v", videoInfo)
	audioInfo := info.AudioInfo
	t.Logf("%+v", audioInfo)
}

func TestCreateThumbnailJob(t *testing.T) {
	target := &api.ThumbnailTarget{}
	target.Format = "jpg"
	target.SizingPolicy = "keep"
	capture := &api.ThumbnailCapture{}
	capture.Mode = "manual"
	capture.StartTimeInSecond = 0.0
	capture.EndTimeInSecond = 5.0
	capture.IntervalInSecond = 1.0
	// params piplineName sourceKey target capture
	createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob("go_sdk_test", "01.mp4",
		TargetOp(target), CaptureOp(capture))
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createJobResponse)
}

func TestCreateThumbnailJobTargetKeyPrefix(t *testing.T) {
	//source := &api.ThumbnailSource{}
	//source.Key = "01.mp4"
	target := &api.ThumbnailTarget{}
	target.KeyPrefix = "taget_key_prefix_test"

	// pipelineName presetName sourceKey targetKeyPrefix
	createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob("go_sdk_test", "01_go_02.mp4",
		PresetNameOp("test"), TargetOp(target))
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createJobResponse)
}

func TestCreateThumbnailJobDelogo(t *testing.T) {
	target := &api.ThumbnailTarget{}
	target.KeyPrefix = "taget_key_prefix_test_delogo3"
	delogo := &api.Area{}
	delogo.X = 20
	delogo.Y = 20
	delogo.Height = 50
	delogo.Width = 80
	// piplineName sourceKey target capture delogo
	createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob("go_sdk_test", "01_go_02.mp4",
		TargetOp(target), DelogoAreaOp(delogo))
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createJobResponse)
}

func TestCreateThumbnailJobDelogoCrop(t *testing.T) {
	target := &api.ThumbnailTarget{}
	target.KeyPrefix = "taget_key_prefix_test_delogo_crop"
	delogo := &api.Area{}
	delogo.X = 20
	delogo.Y = 20
	delogo.Height = 50
	delogo.Width = 80

	crop := &api.Area{}
	crop.X = 120
	crop.Y = 120
	crop.Height = 100
	crop.Width = 80
	// piplineName sourceKey target capture delogo crop
	createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob("go_sdk_test", "01_go_02.mp4",
		TargetOp(target), DelogoAreaOp(delogo), CropOp(crop))
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createJobResponse)
}

func TestCreateThumbnailJobCaptureDelogo(t *testing.T) {
	target := &api.ThumbnailTarget{}
	target.Format = "jpg"
	target.SizingPolicy = "keep"

	capture := &api.ThumbnailCapture{}
	capture.Mode = "split"
	capture.FrameNumber = 30

	delogo := &api.Area{}
	delogo.X = 20
	delogo.Y = 20
	delogo.Height = 50
	delogo.Width = 80
	// params piplineName sourceKey target capture delogo
	createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob("go_sdk_test", "01.mp4",
		TargetOp(target), CaptureOp(capture), DelogoAreaOp(delogo))
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createJobResponse)
}

func TestCreateThumbnailJobCaptureDelogoCrop(t *testing.T) {
	target := &api.ThumbnailTarget{}
	target.Format = "jpg"
	target.SizingPolicy = "keep"

	capture := &api.ThumbnailCapture{}
	capture.Mode = "splitss0"
	capture.FrameNumber = 10

	delogo := &api.Area{}
	delogo.X = 20
	delogo.Y = 20
	delogo.Height = 50
	delogo.Width = 80

	crop := &api.Area{}
	crop.X = 120
	crop.Y = 120
	crop.Height = 100
	crop.Width = 80
	// params piplineName sourceKey target capture delogo
	createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob("go_sdk_test", "01.mp4",
		TargetOp(target), CaptureOp(capture), DelogoAreaOp(delogo), CropOp(crop))
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createJobResponse)
}

func TestCreateThumbnailJobSimple(t *testing.T) {
	// params piplineName sourceKey target capture delogo
	createJobResponse, err := MEDIA_CLIENT.CreateThumbnailJob("go_sdk_test", "01.mp4")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createJobResponse)
}

func TestGetThumbanilJob(t *testing.T) {
	jobResponse, err := MEDIA_CLIENT.GetThumbanilJob("job-pcduuweehm1qd0et")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
	t.Logf("%+v", jobResponse.Source)
	t.Logf("%+v", jobResponse.Target)
	t.Logf("%+v", jobResponse.Capture)
	t.Logf("%+v", jobResponse.DelogoArea)
	t.Logf("%+v", jobResponse.Error)
}

func TestListThumbnailJobs(t *testing.T) {
	listThumbnailJobsResponse, err := MEDIA_CLIENT.ListThumbnailJobs("go_sdk_test")
	ExpectEqual(t.Errorf, err, nil)
	for _, job := range listThumbnailJobsResponse.Thumbnails {
		t.Logf("%+v", job)
	}
}

func TestCreateWaterMark(t *testing.T) {
	args := &api.CreateWaterMarkArgs{}
	// bucket, key, horizontalOffsetInPixel, verticalOffsetInPixel
	args.Bucket = "go-test"
	args.Key = "01.jpg"
	args.HorizontalOffsetInPixel = 20
	args.VerticalOffsetInPixel = 10
	createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createWaterMarkResponse)
}

func TestCreateWaterMarkHV(t *testing.T) {
	args := &api.CreateWaterMarkArgs{}
	// bucket, key, horizontalAlignment, verticalAlignment
	args.Bucket = "go-test"
	args.Key = "01.jpg"
	args.HorizontalAlignment = "right"
	args.VerticalAlignment = "top"
	createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createWaterMarkResponse)
}

func TestCreateWaterMarkHHVV(t *testing.T) {
	args := &api.CreateWaterMarkArgs{}
	// bucket, key, horizontalAlignment, verticalAlignment
	args.Bucket = "go-test"
	args.Key = "01.jpg"
	args.HorizontalOffsetInPixel = 200
	args.HorizontalAlignment = "left"
	args.VerticalOffsetInPixel = 200
	args.VerticalAlignment = "bottom"
	createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createWaterMarkResponse)
}

func TestCreateWaterMarkHVXY(t *testing.T) {
	args := &api.CreateWaterMarkArgs{}
	// bucket, key, horizontalAlignment, verticalAlignment
	args.Bucket = "go-test"
	args.Key = "01.jpg"
	args.HorizontalAlignment = "center"
	args.VerticalAlignment = "center"
	args.Dy = "0.1"
	args.Dy = "0.2"
	createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createWaterMarkResponse)
}

func TestCreateWaterMarkHVXYWH(t *testing.T) {
	args := &api.CreateWaterMarkArgs{}
	// bucket, key, horizontalAlignment, verticalAlignment
	args.Bucket = "go-test"
	args.Key = "01.jpg"
	args.HorizontalAlignment = "center"
	args.VerticalAlignment = "center"
	args.Dy = "0.1"
	args.Dy = "0.2"
	args.Width = "0.15"
	args.Height = "0.11"
	createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createWaterMarkResponse)
}

func TestCreateWaterMarkHVTRA(t *testing.T) {
	args := &api.CreateWaterMarkArgs{}
	// bucket, key, horizontalAlignment, verticalAlignment
	args.Bucket = "go-test"
	args.Key = "01.jpg"
	args.HorizontalAlignment = "left"
	args.VerticalAlignment = "top"
	args.HorizontalOffsetInPixel = 20
	args.VerticalOffsetInPixel = 10
	timeline := &api.Timeline{}
	timeline.StartTimeInMillisecond = 1000
	timeline.DurationInMillisecond = 3000
	args.Timeline = timeline
	args.Repeated = 1
	args.AllowScaling = true
	createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createWaterMarkResponse)
}

func TestCreateWaterMarkHVTXYWHTR(t *testing.T) {
	args := &api.CreateWaterMarkArgs{}
	// bucket, key, horizontalAlignment, verticalAlignment
	args.Bucket = "go-test"
	args.Key = "tupian.jpg"
	args.HorizontalAlignment = "center"
	args.VerticalAlignment = "center"
	args.Dy = "0.1"
	args.Dy = "0.2"
	args.Width = "150"
	args.Height = "110"
	timeline := &api.Timeline{}
	timeline.StartTimeInMillisecond = 1
	timeline.DurationInMillisecond = 5
	args.Timeline = timeline
	args.Repeated = 10
	createWaterMarkResponse, err := MEDIA_CLIENT.CreateWaterMark(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", createWaterMarkResponse)
}

func TestGetWaterMark(t *testing.T) {
	response, err := MEDIA_CLIENT.GetWaterMark("wmk-pcep0x4vvmvvx84r")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
	t.Logf("%+v", response.Timeline)
}

func TestListWaterMark(t *testing.T) {
	response, err := MEDIA_CLIENT.ListWaterMark()
	ExpectEqual(t.Errorf, err, nil)
	for _, watermark := range response.Watermarks {
		t.Logf("%+v", watermark)
	}
}

func TestDeleteWaterMark(t *testing.T) {
	err := MEDIA_CLIENT.DeleteWaterMark("wmk-pcep0x4vvmvvx84r")
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateNotification(t *testing.T) {
	err := MEDIA_CLIENT.CreateNotification("test", "http://www.baidu.com")
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetNotification(t *testing.T) {
	response, err := MEDIA_CLIENT.GetNotification("zz")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestListNotification(t *testing.T) {
	response, err := MEDIA_CLIENT.ListNotification()
	ExpectEqual(t.Errorf, err, nil)
	for _, notification := range response.Notifications {
		t.Logf("%+v", notification)
	}
}

func TestDeleteNotification(t *testing.T) {
	err := MEDIA_CLIENT.DeleteNotification("test")
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateDigitalWmPresetImage(t *testing.T) {
	preset := &api.DigitalWmPreset{}
	preset.Description = "test"
	preset.Bucket = "go-test"
	preset.Key = "logo.png"
	preset.DigitalWmType = "image"
	response, err := MEDIA_CLIENT.CreateDigitalWmPreset(preset)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestCreateDigitalWmPresetText(t *testing.T) {
	preset := &api.DigitalWmPreset{}
	preset.Description = "文字水印模板1"
	//preset.Bucket = "go-test"
	//preset.Key = "logo.png"
	//preset.DigitalWmId = ""
	preset.DigitalWmType = "text"
	preset.TextContent = "文字水印模板1"
	response, err := MEDIA_CLIENT.CreateDigitalWmPreset(preset)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestCreateDigitalWmImagePreset(t *testing.T) {
	response, err := MEDIA_CLIENT.CreateDigitalWmImagePreset("", "图片水印模板", "go-test", "logo.png")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestCreateDigitalWmTextPreset(t *testing.T) {
	response, err := MEDIA_CLIENT.CreateDigitalWmTextPreset("", "文字水印模板", "这是文字水印")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetDigitalWmPreset(t *testing.T) {
	response, err := MEDIA_CLIENT.GetDigitalWmPreset("dwm-qbkfswta3hy3f9ta")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestListDigitalWatermarks(t *testing.T) {
	response, err := MEDIA_CLIENT.ListDigitalWmPreset()
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestDeleteDigitalWmPreset(t *testing.T) {
	err := MEDIA_CLIENT.DeleteDigitalWmPreset("dwm-qbkfv5s5k1mnkiww")
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateDwmSecretkeyPreset(t *testing.T) {
	preset := &api.DwmSecretkeyPreset{}
	preset.Description = "数字水印密钥2"
	preset.SecretKey = "password2"
	response, err := MEDIA_CLIENT.CreateDwmSecretkeyPreset(preset)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetDwmSecretkeyPreset(t *testing.T) {
	response, err := MEDIA_CLIENT.GetDwmSecretkeyPreset("key-qbktvt6xbb6a716w")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestListDwmSecretkeyPresets(t *testing.T) {
	response, err := MEDIA_CLIENT.ListDwmSecretkeyPresets()
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestDeleteDwmSecretkeyPreset(t *testing.T) {
	err := MEDIA_CLIENT.DeleteDwmSecretkeyPreset("key-qbknpihxg9ds23v9")
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestWithDigitalArgsImage(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "go_test_digital_preset"
	preset.Description = "全参数"
	preset.Container = "mp4"

	clip := &api.Clip{}
	clip.StartTimeInSecond = 0
	preset.Clip = clip

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	audio.Codec = "opus"
	preset.Audio = audio

	video := &api.Video{}
	video.Codec = "h264"
	codecOptions := &api.CodecOptions{}
	codecOptions.Profile = "baseline"
	video.CodecOptions = codecOptions
	video.BitRateInBps = 1024000
	video.MaxFrameRate = 30
	video.MaxWidthInPixel = 4096
	video.MaxHeigtInPixel = 3072
	video.SizingPolicy = "keep"
	preset.Video = video

	preset.DigitalWmId = "dwm-qbkfs0uq6shgmbp9"
	preset.DigitalWmSecretKey = "key-qbktvt6xbb6a716w"
	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreatePrestWithDigitalArgsText(t *testing.T) {
	preset := &api.Preset{}
	preset.PresetName = "go_test_digital_preset_text"
	preset.Description = "全参数"
	preset.Container = "mp4"

	clip := &api.Clip{}
	clip.StartTimeInSecond = 0
	preset.Clip = clip

	audio := &api.Audio{}
	audio.BitRateInBps = 256000
	audio.Codec = "opus"
	preset.Audio = audio

	video := &api.Video{}
	video.Codec = "h264"
	codecOptions := &api.CodecOptions{}
	codecOptions.Profile = "baseline"
	video.CodecOptions = codecOptions
	video.BitRateInBps = 1024000
	video.MaxFrameRate = 30
	video.MaxWidthInPixel = 4096
	video.MaxHeigtInPixel = 3072
	video.SizingPolicy = "keep"
	preset.Video = video

	preset.DigitalWmId = "dwm-qbkfswta3hy3f9ta"
	preset.DigitalWmSecretKey = "key-qbktvt6xbb6a716w"
	err := MEDIA_CLIENT.CreatePrestCustomize(preset)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateJobCustomizeWithDigitalImage(t *testing.T) {
	args := &api.CreateJobArgs{}
	args.PipelineName = "go_sdk_test"
	source := &api.Source{Clips: &[]api.SourceClip{{
		SourceKey: "dog.mp4",
		Bucket:    "go-test",
	}}}
	args.Source = source
	target := &api.Target{}
	targetKey := "dog-result.mp4"

	target.TargetKey = targetKey
	presetName := "go_test_digital_preset"
	target.PresetName = presetName
	target.DigitalWmSecretKeyId = "key-qbktvt6xbb6a716w"

	args.Target = target

	jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestCreateJobCustomizeWithDigitalText(t *testing.T) {
	args := &api.CreateJobArgs{}
	args.PipelineName = "go_sdk_test"
	source := &api.Source{Clips: &[]api.SourceClip{{
		SourceKey: "dog.mp4",
		Bucket:    "go-test",
	}}}
	args.Source = source
	target := &api.Target{}
	targetKey := "dog-text-result.mp4"

	target.TargetKey = targetKey
	presetName := "go_test_digital_preset_text"
	target.PresetName = presetName
	target.DigitalWmSecretKeyId = "key-qbktvt6xbb6a716w"
	target.DigitalWmTextContent = "this is a demo"
	args.Target = target

	jobResponse, err := MEDIA_CLIENT.CreateJobCustomize(args)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", jobResponse)
}

func TestCreateDwmDetectJobImage(t *testing.T) {
	dwmDetect := &api.Dwmdetect{}
	dwmDetect.PipelineName = "go_sdk_test"
	dwmDetect.DigitalWmType = "image"
	dwmDetect.DigitalWmSecretKeyId = "key-qbktvt6xbb6a716w"
	dwmDetect.DigitalWmId = "dwm-qbkfs0uq6shgmbp9"
	source := &api.DwmSource{}
	source.Bucket = "go-test"
	source.Key = "dog-result.mp4"
	dwmDetect.Source = source
	target := &api.DwmTarget{}
	target.Bucket = "go-test"
	dwmDetect.Target = target
	ref := []api.RefResolutions{
		{OriginalVideoWidth: 1920, OriginalVideoHeight: 1080},
	}
	dwmDetect.RefResolutions = ref
	response, err := MEDIA_CLIENT.CreateDwmDetectJob(dwmDetect)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestCreateDwmDetectJobText(t *testing.T) {
	dwmDetect := &api.Dwmdetect{}
	dwmDetect.PipelineName = "go_sdk_test"
	dwmDetect.DigitalWmType = "text"
	dwmDetect.DigitalWmSecretKeyId = "key-qbktvt6xbb6a716w"
	dwmDetect.DigitalWmId = "dwm-qbkfswta3hy3f9ta"
	source := &api.DwmSource{}
	source.Bucket = "go-test"
	source.Key = "dog-text-result.mp4"
	dwmDetect.Source = source

	ref := []api.RefResolutions{
		{OriginalVideoWidth: 1920, OriginalVideoHeight: 1080},
	}
	dwmDetect.RefResolutions = ref
	dwmDetect.TextWmContent = "this"
	response, err := MEDIA_CLIENT.CreateDwmDetectJob(dwmDetect)

	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetDwmdetectResult(t *testing.T) {
	response, err := MEDIA_CLIENT.GetDwmdetectResult("job-qbmvxg64e292aauv")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
	t.Logf("%+v", response.Source)
	t.Logf("%+v", response.Target)
}

func TestCreateImagedwmJobT(t *testing.T) {
	imagedwm := &api.Imagedwm{}
	imagedwm.PipelineName = "go_sdk_test"
	source := &api.DwmSource{}
	source.Bucket = "go-test"
	source.Key = "jbs.jpeg"
	imagedwm.Source = source
	target := &api.DwmTarget{}
	target.Bucket = "go-test"
	target.Key = "jbs-dwm.jpeg"
	target.Quality = 90
	imagedwm.Target = target
	imagedwm.TaskType = "embed"
	strength := 0.7
	imagedwm.Strength = &strength
	digitalWm := &api.DigitalWm{}
	digitalWm.TextContent = "baidu"
	imagedwm.DigitalWm = digitalWm
	algorithmVersion := 1
	imagedwm.AlgorithmVersion = &algorithmVersion

	response, err := MEDIA_CLIENT.CreateImagedwmJob(imagedwm)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestCreateImagedwmJobI(t *testing.T) {
	imagedwm := &api.Imagedwm{}
	imagedwm.PipelineName = "go_sdk_test"
	source := &api.DwmSource{}
	source.Bucket = "go-test"
	source.Key = "demo.png"
	imagedwm.Source = source
	target := &api.DwmTarget{}
	target.Bucket = "go-test"
	target.Key = "demo-dwm-i.png"
	imagedwm.Target = target
	imagedwm.TaskType = "embed"
	strength := 0.8
	imagedwm.Strength = &strength
	algorithmVersion := 0
	imagedwm.AlgorithmVersion = &algorithmVersion
	digitalWm := &api.DigitalWm{}
	digitalWm.ImageBucket = "go-test"
	digitalWm.ImageKey = "logo.png"

	imagedwm.DigitalWm = digitalWm

	response, err := MEDIA_CLIENT.CreateImagedwmJob(imagedwm)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetImagedwmResult(t *testing.T) {
	response, err := MEDIA_CLIENT.GetImagedwmResult("job-rc2v7i5y6pinc58h")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
	t.Logf("%+v", response.Source)
	t.Logf("%+v", response.Target)
}

func TestCreateImagedwmDetectJob(t *testing.T) {
	imagedwm := &api.Imagedwm{}
	imagedwm.PipelineName = "go_sdk_test"
	source := &api.DwmSource{}
	source.Bucket = "go-test"
	source.Key = "demo-dwm-i.png"
	imagedwm.Source = source
	target := &api.DwmTarget{}
	target.Bucket = "go-test"
	target.Key = "demo-dwm-i-di.png"
	target.Quality = 80
	imagedwm.Target = target
	imagedwm.TaskType = "extract"

	response, err := MEDIA_CLIENT.CreateImagedwmDetectJob(imagedwm)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", response)
}

func TestGetImagedwmDetectResult(t *testing.T) {
	response, err := MEDIA_CLIENT.GetImagedwmDetectResult("job-rc2uz90cxrk3zxer")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("response:%v", response)
	t.Logf("%+v", response.Source)
	t.Logf("%+v", response.Target)
}
