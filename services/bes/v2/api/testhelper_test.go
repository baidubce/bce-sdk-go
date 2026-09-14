package api

import (
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	nethttp "net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	testAK       = "ak-test-0000000000000000"
	testSK       = "sk-test-0000000000000000"
	testRegion   = "bj"
	testPassword = "PLACEHOLDER-NOT-A-PASSWORD"
)

func boolPtr(v bool) *bool { return &v }

// newTestClient builds a bce.Client pointed at an httptest server. The configuration mirrors
// newClientWithCredentials in services/bes/v2/client.go so signing and headers behave exactly
// as they do through the v2 Client wrapper. The wrapper's own construction and region-inference
// logic stays covered by v2/client_test.go; these api tests only exercise the api layer.
func newTestClient(t *testing.T, handler nethttp.HandlerFunc) (bce.Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)

	credentials, err := auth.NewBceCredentials(testAK, testSK)
	if err != nil {
		server.Close()
		t.Fatalf("new credentials failed: %v", err)
	}

	conf := &bce.BceClientConfiguration{
		Endpoint:    server.URL,
		Region:      testRegion,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption: &auth.SignOptions{
			HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
			ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
		},
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
	}

	return bce.NewBceClient(conf, &auth.BceV1Signer{}), server
}

func writeJSONResponse(t *testing.T, w nethttp.ResponseWriter, value interface{}) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		t.Fatalf("write response failed: %v", err)
	}
}

func readJSONBody(t *testing.T, r *nethttp.Request) map[string]interface{} {
	t.Helper()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("read request body failed: %v", err)
	}
	if len(body) == 0 {
		return map[string]interface{}{}
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("unmarshal request body failed: %v", err)
	}
	return result
}

// writeTempUploadFile creates a temporary file with the given content and returns its path.
// The file is removed automatically when the test finishes.
func writeTempUploadFile(t *testing.T, name string, content []byte) string {
	t.Helper()
	dir := t.TempDir()
	path := dir + string(os.PathSeparator) + name
	if err := os.WriteFile(path, content, 0600); err != nil {
		t.Fatalf("write temp upload file failed: %v", err)
	}
	return path
}

// readMultipartBody parses a multipart/form-data request body and returns the plain form field
// values and the content of the given file field.
func readMultipartBody(t *testing.T, r *nethttp.Request, fileFieldName string) (fields map[string]string, fileContent []byte, fileName string) {
	t.Helper()
	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		t.Fatalf("parse content-type failed: %v", err)
	}
	boundary, ok := params["boundary"]
	if !ok {
		t.Fatalf("missing multipart boundary in content-type: %s", r.Header.Get("Content-Type"))
	}

	fields = map[string]string{}
	reader := multipart.NewReader(r.Body, boundary)
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("read multipart part failed: %v", err)
		}
		data, err := io.ReadAll(part)
		if err != nil {
			t.Fatalf("read multipart part content failed: %v", err)
		}
		if part.FormName() == fileFieldName {
			fileContent = data
			fileName = part.FileName()
			continue
		}
		fields[part.FormName()] = string(data)
	}
	return fields, fileContent, fileName
}
