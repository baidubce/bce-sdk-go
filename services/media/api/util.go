package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	MEDIA_PREFIX              = bce.URI_PREFIX + "v3/"
	PIPLINE                   = "pipeline"
	TRANSCODING_JOB           = "job/transcoding"
	PRESET                    = "preset"
	MEDIA_INFO                = "mediainfo"
	THUMBNAIL                 = "job/thumbnail"
	WATERMARK                 = "watermark"
	NOTIFICATION              = "notification"
	DIGITALWATERMARK          = "digitalwatermark"
	DIGITALWATERMARKSECRETKEY = "digitalwatermark/secretkey"
	DWMDETECT                 = "job/dwmdetect"
	IMAGEDWM                  = "job/imagedwm"
)

func getPipLineUrl() string {
	return MEDIA_PREFIX + PIPLINE
}

func getTrandCodingJobUrl() string {
	return MEDIA_PREFIX + TRANSCODING_JOB
}

func getPresetUrl() string {
	return MEDIA_PREFIX + PRESET
}

func getMediaInfoUrl() string {
	return MEDIA_PREFIX + MEDIA_INFO
}

func getThumbnailUrl() string {
	return MEDIA_PREFIX + THUMBNAIL
}

func getWatermarkUrl() string {
	return MEDIA_PREFIX + WATERMARK
}

func getNotification() string {
	return MEDIA_PREFIX + NOTIFICATION
}

func getDigitalwatermarkUrl() string {
	return MEDIA_PREFIX + DIGITALWATERMARK
}

func getDigitalwatermarkSecretkeyUrl() string {
	return MEDIA_PREFIX + DIGITALWATERMARKSECRETKEY
}

func getDwmdetectUrl() string {
	return MEDIA_PREFIX + DWMDETECT
}

func getImagedwmUrl() string {
	return MEDIA_PREFIX + IMAGEDWM
}
