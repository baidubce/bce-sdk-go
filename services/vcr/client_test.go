package vcr

import (
	"github.com/baidubce/bce-sdk-go/services/vcr/api"
	"github.com/baidubce/bce-sdk-go/util/log"
	"testing"
)

var CLIENT *Client

const (
	AK           = "YourAK"
	SK           = "YourSK"
	MEDIA_SOURCE = "YourTestMedia" // e.g. "bos://YourBucket/dir/video.mp4
	TEXT         = "YourTestText"
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
	desc := "vcr_test"
	args := &api.PutMediaArgs{Source: MEDIA_SOURCE, Preset: preset, Description: desc}
	if err := CLIENT.PutMedia(args); err != nil {
		t.Log("put media error")
	} else {
		t.Log("put media success")
	}
}

func TestSimplePutMedia(t *testing.T) {
	if err := CLIENT.SimplePutMedia(MEDIA_SOURCE, "", "", ""); err != nil {
		t.Log("simple put media error")
	} else {
		t.Log("simple put media success")
	}
}

func TestGetMedia(t *testing.T) {
	if res, err := CLIENT.GetMedia(MEDIA_SOURCE); err != nil {
		t.Log("get media error")
	} else {
		t.Logf("%+v", res)
	}
}

func TestPutText(t *testing.T) {
	args := &api.PutTextArgs{Text: TEXT}
	if res, err := CLIENT.PutText(args); err != nil {
		t.Log("simple put text error")
	} else {
		t.Logf("%+v", res)
	}
}

func TestSimplePutText(t *testing.T) {
	if res, err := CLIENT.SimplePutText(TEXT); err != nil {
		t.Log("simple put text error")
	} else {
		t.Logf("%+v", res)
	}
}
