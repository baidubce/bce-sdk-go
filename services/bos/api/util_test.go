package api

import (
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
)

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !util.Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
}

func TestUtil(t *testing.T) {
	_, ok := VALID_RESTORE_TIER[RESTORE_TIER_STANDARD]
	ExpectEqual(t, true, ok)
	_, ok = VALID_RESTORE_TIER[RESTORE_TIER_EXPEDITED]
	ExpectEqual(t, true, ok)
	_, ok = VALID_RESTORE_TIER[RESTORE_TIER_LOWCOST]
	ExpectEqual(t, true, ok)
	_, ok = VALID_RESTORE_TIER["restore_tier_unknown"]
	ExpectEqual(t, false, ok)

	bucket := "test-bucket"
	object := "test-object"
	ExpectEqual(t, "/"+bucket, getBucketUri(bucket))
	ExpectEqual(t, "/"+bucket+"/"+object, getObjectUri(bucket, object))
	ExpectEqual(t, "", getCnameUri(""))
	ExpectEqual(t, "/", getCnameUri("/"))
	ExpectEqual(t, "/", getCnameUri("/path"))
	ExpectEqual(t, "/dir", getCnameUri("/path/dir"))

	ExpectEqual(t, true, validMetadataDirective(METADATA_DIRECTIVE_COPY))
	ExpectEqual(t, true, validMetadataDirective(METADATA_DIRECTIVE_COPY))
	ExpectEqual(t, false, validMetadataDirective("unknown_metadata_directive"))

	tooLengthTagging := make([]byte, 4096)
	for i := 0; i < len(tooLengthTagging); i++ {
		tooLengthTagging[i] = 't'
	}
	tooLengthKey := make([]byte, 150)
	for i := 0; i < len(tooLengthKey); i++ {
		tooLengthKey[i] = 'k'
	}
	tooLengthVal := make([]byte, 300)
	for i := 0; i < len(tooLengthVal); i++ {
		tooLengthVal[i] = 'v'
	}

	testTagging := []struct {
		tag string
		ok  bool
		res string
	}{
		{"", false, ""},
		{"k=v=vv", false, ""},
		{"testtagging", false, ""},
		{string(tooLengthTagging), false, ""},
		{string(tooLengthKey) + "=val", false, ""},
		{"key=" + string(tooLengthVal), false, ""},
		{"key1=val1&key2=val2", true, "key1=val1&key2=val2"},
	}

	for _, v := range testTagging {
		ok, res := validObjectTagging(v.tag)
		ExpectEqual(t, v.ok, ok)
		ExpectEqual(t, v.res, res)
	}

	testTagsToStr := []struct {
		tags map[string]string
		res  string
	}{
		{make(map[string]string), ""},
		{map[string]string{string(tooLengthKey): "val"}, ""},
		{map[string]string{"key": string(tooLengthVal)}, ""},
		{map[string]string{"key1": "val1", "key2": "val2"}, "key2=val1&key1=val1"},
	}
	for _, v := range testTagsToStr {
		ExpectEqual(t, v.res, taggingMapToStr(v.tags))
	}

	httpKey := []struct {
		input  string
		output string
	}{
		{"x-bce-header1", "X-Bce-Header1"},
		{"", ""},
		{"1-2-3", "1-2-3"},
	}
	for _, v := range httpKey {
		ExpectEqual(t, v.output, toHttpHeaderKey(v.input))
	}
}
