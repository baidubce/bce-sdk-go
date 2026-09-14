package api

import (
	"crypto/aes"
	"encoding/hex"
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

func boolPtr(v bool) *bool { return &v }

func intPtr(v int) *int { return &v }

const (
	testAK     = "ak-test-0000000000000000"
	testSK     = "sk-test-0000000000000000"
	testRegion = "bj"
)

// newTestClient builds a bce.Client pointed at an httptest server. The configuration mirrors
// newClientWithCredentials in services/bes/v3/client.go so signing and headers behave exactly
// as they do through the v3 Client wrapper. The wrapper's own construction and region-inference
// logic stays covered by v3/client_test.go; these api tests only exercise the api layer.
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

// decryptAES128WithFirst16Char reverses aes128EncryptWithFirst16Char (AES-128/ECB/PKCS5Padding,
// keyed by the first 16 bytes of secretKey) so tests can assert on the plaintext password
// instead of hardcoding an expected ciphertext.
func decryptAES128WithFirst16Char(t *testing.T, cipherHex, secretKey string) string {
	t.Helper()
	key := []byte(secretKey[:16])
	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatalf("new cipher failed: %v", err)
	}
	encrypted, err := hex.DecodeString(cipherHex)
	if err != nil {
		t.Fatalf("decode hex failed: %v", err)
	}
	if len(encrypted) == 0 || len(encrypted)%block.BlockSize() != 0 {
		t.Fatalf("invalid ciphertext length: %d", len(encrypted))
	}
	decrypted := make([]byte, len(encrypted))
	blockSize := block.BlockSize()
	for start := 0; start < len(encrypted); start += blockSize {
		block.Decrypt(decrypted[start:start+blockSize], encrypted[start:start+blockSize])
	}
	padSize := int(decrypted[len(decrypted)-1])
	if padSize <= 0 || padSize > blockSize {
		t.Fatalf("invalid PKCS5 padding size: %d", padSize)
	}
	return string(decrypted[:len(decrypted)-padSize])
}
