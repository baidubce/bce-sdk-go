package api

import (
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

const testPassword = "PLACEHOLDER-NOT-A-PASSWORD"

// TestAes128EncryptWithFirst16Char cross-checks the Go implementation against an independent
// reference computed with `openssl enc -aes-128-ecb`, using an obviously-fake placeholder key
// (not a real credential). This guards the protocol contract with BES: the server decrypts
// password fields with the same algorithm keyed by the caller's real secret key, so any drift
// here would silently corrupt passwords in production.
func TestAes128EncryptWithFirst16Char(t *testing.T) {
	// Reference value: `printf 'PLACEHOLDER-NOT-A-PASSWORD' | openssl enc -aes-128-ecb -K $(printf
	// 'sk-test-00000000' | xxd -p)` (first 16 chars of the placeholder key below).
	got, err := aes128EncryptWithFirst16Char(testPassword, "sk-test-0000000000000000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "f0a7f1578ed3938154b296a7683e83c5babe6ec1cefe1eb8cd1a6c045f1ad83b"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestAes128EncryptWithFirst16CharShortKey(t *testing.T) {
	if _, err := aes128EncryptWithFirst16Char(testPassword, "short"); err == nil {
		t.Fatal("expected error for secret key shorter than 16 bytes")
	}
}

// newMultipartTestBody builds a minimal non-nil *bce.Body. bce.BceRequest.SetBody dereferences its
// argument (bce/request.go:382), so createMultipartRequest cannot be called with a nil body.
func newMultipartTestBody(t *testing.T) *bce.Body {
	t.Helper()
	body, err := bce.NewBodyFromString("--abc\r\nContent-Disposition: form-data\r\n\r\n--abc--\r\n")
	if err != nil {
		t.Fatalf("new body failed: %v", err)
	}
	return body
}

// TestCreateMultipartRequestRejectsFailedResponse covers the bceResp.IsFail() guard. It is
// unreachable through *bce.BceClient, which converts a failing status into the error returned by
// SendRequest, so it takes a stub client (clusterFailingResponseClient, defined in cluster_test.go)
// to reach the branch the guard actually protects for other bce.Client implementations.
func TestCreateMultipartRequestRejectsFailedResponse(t *testing.T) {
	err := createMultipartRequest(clusterFailingResponseClient{}, testRegion, nethttp.MethodPost,
		"/v3/clusters/search-xxxx/plugins/file", newMultipartTestBody(t),
		"multipart/form-data; boundary=abc", nil)
	if err == nil {
		t.Fatal("expected the service error of the failed response")
	}
}

// TestCreateMultipartRequestWithNilResult covers the `resp == nil` early return, taken when a
// caller does not want the response body decoded.
func TestCreateMultipartRequestWithNilResult(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		if got := r.Header.Get("Content-Type"); got != "multipart/form-data; boundary=abc" {
			t.Fatalf("unexpected content type: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"ignored": true})
	}))
	defer server.Close()

	err := createMultipartRequest(cli, testRegion, nethttp.MethodPost,
		"/v3/clusters/search-xxxx/plugins/file", newMultipartTestBody(t),
		"multipart/form-data; boundary=abc", nil)
	if err != nil {
		t.Fatalf("create multipart request failed: %v", err)
	}
}
