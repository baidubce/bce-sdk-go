package vca

import (
	"github.com/baidubce/bce-sdk-go/services/vca/api"
	"github.com/baidubce/bce-sdk-go/util/log"
	"testing"
)

var CLIENT *Client

const (
	AK           = "YourAK"
	SK           = "YourSK"
	MEDIA_SOURCE = "YourTestMedia" // e.g. "bos://YourBucket/dir/video.mp4
)

func init() {
	CLIENT, _ = NewClient(AK, SK, "")
	//log.SetLogHandler(log.STDERR | log.FILE)
	//log.SetRotateType(log.ROTATE_SIZE)
	//log.SetLogLevel(log.WARN)

	log.SetLogHandler(log.STDERR)
	log.SetLogLevel(log.DEBUG)
}

func TestPutMedia(t *testing.T) {
	preset := "default"
	args := &api.PutMediaArgs{Source: MEDIA_SOURCE, Preset: preset}
	if res, err := CLIENT.PutMedia(args); err != nil {
		t.Log("put media error")
	} else {
		t.Logf("%+v", res)
	}
}

func TestGetMedia(t *testing.T) {
	if res, err := CLIENT.GetMedia(MEDIA_SOURCE); err != nil {
		t.Log("get media error")
	} else {
		t.Logf("%+v", res)
	}
}
