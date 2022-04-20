package doc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/baidubce/bce-sdk-go/services/doc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	DOC_CLIENT *Client
	BOS_CLIENT *bos.Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK string
	SK string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given: %+v\n", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	DOC_CLIENT, _ = NewClient(confObj.AK, confObj.SK)
	BOS_CLIENT, _ = bos.NewClient(confObj.AK, confObj.SK, "")
	//log.SetLogHandler(log.STDERR | log.FILE)
	//log.SetRotateType(log.ROTATE_SIZE)
	log.SetLogLevel(log.WARN)
	//log.SetLogHandler(log.STDERR)
	//log.SetLogLevel(log.DEBUG)
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

func TestGetDocHTML(t *testing.T) {
	regParam := &api.RegDocumentParam{
		Title:  "test.txt",
		Format: "txt",
	}
	res, err := DOC_CLIENT.RegisterDocument(regParam)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)

	etag, err := BOS_CLIENT.PutObjectFromString(res.Bucket, res.Object, "test\nline", nil)
	ExpectEqual(t.Errorf, nil, err)

	t.Logf("%+v", etag)

	err = DOC_CLIENT.PublishDocument(res.DocumentId)
	ExpectEqual(t.Errorf, nil, err)

	for {
		time.Sleep(time.Duration(1) * time.Second)
		qRes, err := DOC_CLIENT.QueryDocument(res.DocumentId, &api.QueryDocumentParam{Https: false})
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", qRes)
		if qRes.Status == "PUBLISHED" {
			break
		}
	}
	rRes, err := DOC_CLIENT.ReadDocument(res.DocumentId, &api.ReadDocumentParam{ExpireInSeconds: 3600})
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "BCEDOC", rRes.Host)
	t.Logf("%+v", rRes)
}

func TestGetDocImages(t *testing.T) {
	regParam := &api.RegDocumentParam{
		Title:      "sudoku.pdf",
		Format:     "pdf",
		TargetType: "image",
	}
	res, err := DOC_CLIENT.RegisterDocument(regParam)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)

	etag, err := BOS_CLIENT.PutObjectFromFile(res.Bucket, res.Object, "./sudoku.pdf", nil)
	ExpectEqual(t.Errorf, nil, err)

	t.Logf("%+v", etag)

	err = DOC_CLIENT.PublishDocument(res.DocumentId)
	ExpectEqual(t.Errorf, nil, err)

	for {
		time.Sleep(time.Duration(1) * time.Second)
		qRes, err := DOC_CLIENT.QueryDocument(res.DocumentId, &api.QueryDocumentParam{Https: false})
		ExpectEqual(t.Errorf, nil, err)
		t.Logf("%+v", qRes)
		if qRes.Status == "PUBLISHED" {
			break
		}
	}
	rRes, err := DOC_CLIENT.GetImages(res.DocumentId)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", rRes)
}

func TestListDocs(t *testing.T) {
	listParam := &api.ListDocumentsParam{
		Status:  api.DOC_STATUS_PUBLISHED,
		MaxSize: 2,
	}
	res, err := DOC_CLIENT.ListDocuments(listParam)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Docs))
	t.Logf("%+v", res)

	listParam.Marker = res.NextMarker
	listParam.MaxSize = 3
	res, err = DOC_CLIENT.ListDocuments(listParam)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 3, len(res.Docs))
}

func TestDeleteDocs(t *testing.T) {
	regParam := &api.RegDocumentParam{
		Title:  "test.txt",
		Format: "txt",
	}
	res, err := DOC_CLIENT.RegisterDocument(regParam)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)

	err = DOC_CLIENT.DeleteDocument(res.DocumentId)
	ExpectEqual(t.Errorf, nil, err)
}
