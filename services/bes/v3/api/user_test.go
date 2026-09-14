package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

func TestResetAdminPassword(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/users/admins/passwords" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		got, _ := body["newPassword"].(string)
		if got == "" || got == testPassword {
			t.Fatalf("newPassword should be encrypted, not sent as plaintext: %v", body)
		}
		if decryptAES128WithFirst16Char(t, got, testSK) != testPassword {
			t.Fatalf("newPassword did not decrypt to the original plaintext: %v", got)
		}
		if headerAK := r.Header.Get("X-Bce-Accesskey"); headerAK != testAK {
			t.Fatalf("unexpected X-Bce-Accesskey: %s", headerAK)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := ResetAdminPassword(cli, testRegion, &ResetAdminPasswordRequest{
		Request:     Request{Region: "sh"},
		ClusterId:   "search-xxxxxxxx",
		NewPassword: testPassword,
	})
	if err != nil {
		t.Fatalf("reset admin password failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestResetAdminPasswordRejectsNilClient(t *testing.T) {
	if _, err := ResetAdminPassword(nil, testRegion, &ResetAdminPasswordRequest{}); !errors.Is(err, ErrNilClient) {
		t.Fatalf("got %v, want ErrNilClient", err)
	}
}

func TestResetAdminPasswordRejectsInvalidRequests(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"nilRequest": func() error {
			_, err := ResetAdminPassword(cli, testRegion, nil)
			return err
		},
		"missingClusterId": func() error {
			_, err := ResetAdminPassword(cli, testRegion, &ResetAdminPasswordRequest{NewPassword: testPassword})
			return err
		},
		"missingNewPassword": func() error {
			_, err := ResetAdminPassword(cli, testRegion, &ResetAdminPasswordRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// newUserTestClientWithCredentials builds a client that never reaches the network, so the
// credential-dependent failures before the request is sent can be exercised in isolation.
func newUserTestClientWithCredentials(t *testing.T, credentials *auth.BceCredentials) bce.Client {
	t.Helper()
	return bce.NewBceClient(&bce.BceClientConfiguration{
		Endpoint:    "http://127.0.0.1:1",
		Region:      testRegion,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption: &auth.SignOptions{
			HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
			ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
		},
	}, &auth.BceV1Signer{})
}

// TestResetAdminPasswordRejectsUnusableCredentials covers the two credential error paths that
// precede the HTTP call: missing credentials on the client, and a secret key too short to key
// the AES-128 encryption of the password field.
func TestResetAdminPasswordRejectsUnusableCredentials(t *testing.T) {
	shortKeyCredentials, err := auth.NewBceCredentials(testAK, "short")
	if err != nil {
		t.Fatalf("new credentials failed: %v", err)
	}

	cases := map[string]bce.Client{
		"missingCredentials": newUserTestClientWithCredentials(t, nil),
		"shortSecretKey":     newUserTestClientWithCredentials(t, shortKeyCredentials),
	}
	for name, cli := range cases {
		_, err := ResetAdminPassword(cli, testRegion, &ResetAdminPasswordRequest{
			ClusterId:   "search-xxxxxxxx",
			NewPassword: testPassword,
		})
		if err == nil {
			t.Fatalf("%s: expected error", name)
		}
	}
}

// TestResetAdminPasswordPropagatesServerError uses 400 rather than 500 on purpose: the helper
// mirrors the production retry policy, which retries 5xx with backoff and would make this test
// take seconds.
func TestResetAdminPasswordPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		w.WriteHeader(nethttp.StatusBadRequest)
		writeJSONResponse(t, w, map[string]interface{}{
			"code":      "InvalidParameter",
			"message":   "invalid parameter",
			"requestId": "req-xxxx",
		})
	}))
	defer server.Close()

	if _, err := ResetAdminPassword(cli, testRegion, &ResetAdminPasswordRequest{
		ClusterId:   "search-xxxxxxxx",
		NewPassword: testPassword,
	}); err == nil {
		t.Fatal("expected error for 400 response")
	}
}
