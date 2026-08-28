// nolint
package bos

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
	"github.com/baidubce/bce-sdk-go/util"
)

const (
	metaTestRequestId = "mock-request-id"
	metaTestDebugId   = "mock-debug-id"
)

func newMetadataMockClient(t *testing.T) *Client {
	client, err := NewMockBosClient("ak", "sk", "bj.bcebos.com", "{}",
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.BCE_REQUEST_ID: metaTestRequestId,
			my_http.BCE_DEBUG_ID:   metaTestDebugId,
		}),
	)
	ExpectEqual(t.Errorf, nil, err)
	return client
}

// TestWithContextMethodsForwardOptions exercises the end to end metadata contract of
// every *WithContext method that used to drop its options: whatever the internal path,
// the caller supplied sink must be filled. The returned error is deliberately ignored,
// since the contract is that the metadata is filled as soon as an http response has been
// received. TestClientOptionsAreForwarded is what actually pins the forwarding itself.
func TestWithContextMethodsForwardOptions(t *testing.T) {
	bucket, object, uploadId := "test-bucket", "test-object", "test-upload-id"
	ctx := context.Background()

	file := filepath.Join(t.TempDir(), "part")
	if err := ioutil.WriteFile(file, []byte("hello"), 0600); err != nil {
		t.Fatal(err)
	}

	// a fresh handle per case, so a case consuming the reader cannot affect the others
	openFile := func() *os.File {
		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { f.Close() })
		return f
	}
	newBody := func() *bce.Body {
		body, err := bce.NewBodyFromString("hello")
		ExpectEqual(t.Errorf, nil, err)
		return body
	}

	cases := []struct {
		name string
		call func(c *Client, opt api.Option) error
	}{
		{"PutObjectWithContext", func(c *Client, opt api.Option) error {
			_, err := c.PutObjectWithContext(ctx, bucket, object, newBody(), nil, opt)
			return err
		}},
		{"PutObjectFromBytesWithContext", func(c *Client, opt api.Option) error {
			_, err := c.PutObjectFromBytesWithContext(ctx, bucket, object, []byte("hello"), nil, opt)
			return err
		}},
		{"PutObjectFromStringWithContext", func(c *Client, opt api.Option) error {
			_, err := c.PutObjectFromStringWithContext(ctx, bucket, object, "hello", nil, opt)
			return err
		}},
		{"PutObjectFromFileWithContext", func(c *Client, opt api.Option) error {
			_, err := c.PutObjectFromFileWithContext(ctx, bucket, object, file, nil, opt)
			return err
		}},
		{"PutObjectFromStreamWithContext", func(c *Client, opt api.Option) error {
			_, err := c.PutObjectFromStreamWithContext(ctx, bucket, object, strings.NewReader("hello"), nil, opt)
			return err
		}},
		{"CopyObjectWithContext", func(c *Client, opt api.Option) error {
			_, err := c.CopyObjectWithContext(ctx, bucket, object, "src-bucket", "src-object", nil, opt)
			return err
		}},
		{"GetObjectMetaWithContext", func(c *Client, opt api.Option) error {
			_, err := c.GetObjectMetaWithContext(ctx, bucket, object, opt)
			return err
		}},
		{"SelectObjectWithContext", func(c *Client, opt api.Option) error {
			_, err := c.SelectObjectWithContext(ctx, bucket, object, &api.SelectObjectArgs{SelectType: "csv"}, opt)
			return err
		}},
		{"FetchObjectWithContext", func(c *Client, opt api.Option) error {
			_, err := c.FetchObjectWithContext(ctx, bucket, object, "http://example.com/src", nil, opt)
			return err
		}},
		{"AppendObjectWithContext", func(c *Client, opt api.Option) error {
			_, err := c.AppendObjectWithContext(ctx, bucket, object, newBody(), nil, opt)
			return err
		}},
		{"UploadPartWithContext", func(c *Client, opt api.Option) error {
			_, err := c.UploadPartWithContext(ctx, bucket, object, uploadId, 1, newBody(), nil, opt)
			return err
		}},
		{"UploadPartFromSectionFileWithContext", func(c *Client, opt api.Option) error {
			_, err := c.UploadPartFromSectionFileWithContext(ctx, bucket, object, uploadId, 1,
				openFile(), 0, 5, nil, opt)
			return err
		}},
		{"UploadPartFromBytesWithContext", func(c *Client, opt api.Option) error {
			_, err := c.UploadPartFromBytesWithContext(ctx, bucket, object, uploadId, 1, []byte("hello"), nil, opt)
			return err
		}},
		{"UploadPartCopyWithContext", func(c *Client, opt api.Option) error {
			_, err := c.UploadPartCopyWithContext(ctx, bucket, object, "src-bucket", "src-object", uploadId, 1, nil, opt)
			return err
		}},
		{"ListPartsWithContext", func(c *Client, opt api.Option) error {
			_, err := c.ListPartsWithContext(ctx, bucket, object, uploadId, nil, opt)
			return err
		}},
		{"ListMultipartUploadsWithContext", func(c *Client, opt api.Option) error {
			_, err := c.ListMultipartUploadsWithContext(ctx, bucket, nil, opt)
			return err
		}},
	}

	for _, c := range cases {
		meta := &api.ResponseCommon{}
		if err := c.call(newMetadataMockClient(t), api.WithResponseCommon(meta)); err != nil {
			t.Logf("%s returned %v, the metadata must be filled anyway", c.name, err)
		}
		if meta.RequestId != metaTestRequestId || meta.StatusCode != 200 {
			t.Errorf("%s does not forward the options: got %+v", c.name, *meta)
		}
	}
}

// TestClientOptionsAreForwarded pins the invariant that a Client method accepting
// options actually passes them down. An end to end test cannot catch every violation on
// its own: the methods also calling HandleBosClientOptions register the metadata sink on
// the BosContext, which masks the missing forwarding.
func TestClientOptionsAreForwarded(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "client.go", nil, 0)
	ExpectEqual(t.Errorf, nil, err)

	checked := 0
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || fn.Body == nil {
			continue
		}
		last := fn.Type.Params.List[len(fn.Type.Params.List)-1]
		ellipsis, ok := last.Type.(*ast.Ellipsis)
		if !ok || len(last.Names) != 1 || last.Names[0].Name != "options" {
			continue
		}
		if selector, ok := ellipsis.Elt.(*ast.SelectorExpr); !ok || selector.Sel.Name != "Option" {
			continue
		}
		checked++
		if !forwardsOptions(fn.Body) {
			t.Errorf("%s accepts options but never forwards them with options...", fn.Name.Name)
		}
	}
	// guard against the walk silently matching nothing at all
	if checked < 150 {
		t.Errorf("only %d methods taking options were inspected, the walk is broken", checked)
	}
}

// forwardsOptions reports whether the body contains a call passing options on, either
// spread as a variadic argument or handed over as a whole slice to HandleBosClientOptions.
func forwardsOptions(body *ast.BlockStmt) bool {
	found := false
	ast.Inspect(body, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok || call.Ellipsis == token.NoPos {
			return true
		}
		for _, arg := range call.Args {
			if ident, ok := arg.(*ast.Ident); ok && ident.Name == "options" {
				found = true
			}
		}
		return true
	})
	return found
}

// TestClientResponseCommon covers the two shapes of api which cannot expose the request
// id through their result: the ones only returning an error and the ones returning a bare
// string.
func TestClientResponseCommon(t *testing.T) {
	meta := &api.ResponseCommon{}
	err := newMetadataMockClient(t).DeleteBucket("test-bucket", api.WithResponseCommon(meta))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, metaTestRequestId, meta.RequestId)
	ExpectEqual(t.Errorf, metaTestDebugId, meta.DebugId)
	ExpectEqual(t.Errorf, 200, meta.StatusCode)

	meta = &api.ResponseCommon{}
	_, err = newMetadataMockClient(t).PutBucket("test-bucket", api.WithResponseCommon(meta))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, metaTestRequestId, meta.RequestId)
	ExpectEqual(t.Errorf, 200, meta.StatusCode)
}

// TestClientResponseCommonOnFailure pins that the metadata of a failed request matches
// the one carried by the returned service error.
func TestClientResponseCommonOnFailure(t *testing.T) {
	client, err := NewMockBosClient("ak", "sk", "bj.bcebos.com", "{}",
		util.SetStatusCode(403),
		util.SetStatusMsg("403 Forbidden"),
		util.AddHeaders(map[string]string{
			my_http.BCE_REQUEST_ID: metaTestRequestId,
			my_http.BCE_DEBUG_ID:   metaTestDebugId,
		}),
	)
	ExpectEqual(t.Errorf, nil, err)

	meta := &api.ResponseCommon{}
	err = client.DeleteBucket("test-bucket", api.WithResponseCommon(meta))
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, serviceErr.RequestId, meta.RequestId)
	ExpectEqual(t.Errorf, 403, meta.StatusCode)
}

// TestClientResponseCommonNotSentUntouched checks that a failure happening before the
// request is sent leaves the caller supplied struct as it was.
func TestClientResponseCommonNotSentUntouched(t *testing.T) {
	meta := &api.ResponseCommon{RequestId: "untouched", StatusCode: 999}
	err := newMetadataMockClient(t).DeleteBucket("INVALID_BUCKET", api.WithResponseCommon(meta))
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, "untouched", meta.RequestId)
	ExpectEqual(t.Errorf, 999, meta.StatusCode)
}

// TestWithOptionsMethods covers the three methods whose trailing variadic parameter made
// it impossible to accept options, and which therefore got an additive sibling taking the
// variadic part as a slice instead.
func TestWithOptionsMethods(t *testing.T) {
	cases := []struct {
		name string
		call func(c *Client, opt api.Option) error
	}{
		{"PutBucketCopyrightProtectionWithOptions", func(c *Client, opt api.Option) error {
			return c.PutBucketCopyrightProtectionWithOptions("test-bucket",
				[]string{"test-bucket/test-object"}, opt)
		}},
		{"PutObjectAclGrantReadWithOptions", func(c *Client, opt api.Option) error {
			return c.PutObjectAclGrantReadWithOptions("test-bucket", "test-object",
				[]string{"user-id"}, opt)
		}},
		{"PutObjectAclGrantFullControlWithOptions", func(c *Client, opt api.Option) error {
			return c.PutObjectAclGrantFullControlWithOptions("test-bucket", "test-object",
				[]string{"user-id"}, opt)
		}},
	}
	for _, c := range cases {
		meta := &api.ResponseCommon{}
		ExpectEqual(t.Errorf, nil, c.call(newMetadataMockClient(t), api.WithResponseCommon(meta)))
		if meta.RequestId != metaTestRequestId || meta.DebugId != metaTestDebugId ||
			meta.StatusCode != 200 {
			t.Errorf("%s does not forward the options: got %+v", c.name, *meta)
		}
	}
}

// TestCompositeApiResponseCommon covers the composite apis issuing several requests: the
// sink is handed only to the last request running on the caller goroutine, never to the
// fan out, so it must be filled without racing. Run under -race to pin the latter.
func TestCompositeApiResponseCommon(t *testing.T) {
	file := filepath.Join(t.TempDir(), "super")
	if err := ioutil.WriteFile(file, make([]byte, MIN_MULTIPART_SIZE), 0600); err != nil {
		t.Fatal(err)
	}

	meta := &api.ResponseCommon{}
	err := newMetadataMockClient(t).UploadSuperFileWithOptions("test-bucket", "test-object", file, "",
		api.WithResponseCommon(meta))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, metaTestRequestId, meta.RequestId)
	ExpectEqual(t.Errorf, 200, meta.StatusCode)

	// the mock reports no Content-Length, so no range is downloaded and GetObjectMeta is
	// the only request issued
	meta = &api.ResponseCommon{}
	err = newMetadataMockClient(t).DownloadSuperFileWithOptions("test-bucket", "test-object",
		filepath.Join(t.TempDir(), "out"), api.WithResponseCommon(meta))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, metaTestRequestId, meta.RequestId)
	ExpectEqual(t.Errorf, 200, meta.StatusCode)

	// ParallelUpload and ParallelCopy expose the metadata of their last request through
	// the returned CompleteMultipartUploadResult, so they need no sink at all
	res, err := newMetadataMockClient(t).ParallelUpload("test-bucket", "test-object", file, "", nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, metaTestRequestId, res.RequestId)
	ExpectEqual(t.Errorf, 200, res.StatusCode)
}
