package auth

import (
	"runtime"
	"strings"
	"testing"

	my_http "github.com/baidubce/bce-sdk-go/http"
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

func TestGetCanonicalHeaders(t *testing.T) {
	headersToSign := DEFAULT_HEADERS_TO_SIGN
	headersToSign[my_http.AUTHORIZATION] = struct{}{}
	headers := map[string]string{
		my_http.HOST:                              "host",
		my_http.BCE_REQUEST_ID:                    "request-id",
		my_http.BCE_STORAGE_CLASS:                 "storage-class",
		my_http.BCE_USER_METADATA_PREFIX + "key1": "",
		my_http.BCE_USER_METADATA_PREFIX + "key2": "value2",
	}
	expCanonicalHeaders := []string{
		strings.ToLower(my_http.HOST), my_http.BCE_USER_METADATA_PREFIX + "key2", my_http.BCE_STORAGE_CLASS,
	}
	_, signedHeadersArr := getCanonicalHeaders(headers, headersToSign)
	ExpectEqual(t, expCanonicalHeaders, signedHeadersArr)
}
