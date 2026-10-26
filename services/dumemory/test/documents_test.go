// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import (
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/dumemory/api"
)

// pickFirstDocumentID returns a document id from the bank, or skips the
// caller test when none exist.
func pickFirstDocumentID(t *testing.T, c *dumemory.Client) string {
	t.Helper()
	resp, err := c.ListDocuments(ctx(t), bankID(), dumemory.ListDocumentsOptions{Limit: 1})
	if err != nil {
		t.Skipf("ListDocuments failed: %v", err)
	}
	if resp == nil || len(resp.Items) == 0 {
		t.Skip("no documents available; create one via FilesRetain first")
	}
	id := resp.Items[0].Id
	if id == "" {
		t.Skipf("first document has no id field: %v", resp.Items[0])
	}
	return id
}

func TestListDocuments(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	out, err := c.ListDocuments(ctx(t), bankID(), dumemory.ListDocumentsOptions{Limit: 10})
	if err != nil {
		t.Fatalf("ListDocuments: %v", err)
	}
	t.Logf("ListDocuments => %+v", out)
}

func TestGetDocument(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	docID := pickFirstDocumentID(t, c)
	out, err := c.GetDocument(ctx(t), bankID(), docID)
	if err != nil {
		t.Fatalf("GetDocument: %v", err)
	}
	t.Logf("GetDocument => %+v", out)
}

func TestUpdateDocument(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	docID := pickFirstDocumentID(t, c)
	req := dumemory.NewUpdateDocumentRequest()
	req.SetTags([]string{"bce-sdk-it"})
	out, err := c.UpdateDocument(ctx(t), bankID(), docID, *req)
	if err != nil {
		t.Fatalf("UpdateDocument: %v", err)
	}
	t.Logf("UpdateDocument => %+v", out)
}

func TestListDocumentChunks(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	docID := pickFirstDocumentID(t, c)
	out, err := c.ListDocumentChunks(ctx(t), bankID(), docID, 10, 0)
	if err != nil {
		t.Fatalf("ListDocumentChunks: %v", err)
	}
	t.Logf("ListDocumentChunks => %+v", out)
}

func TestDeleteDocument(t *testing.T) {
	if testing.Short() {
		t.Skip("skip destructive DeleteDocument in -short")
	}
	c := newClient(t)
	ensureBank(t, c)
	docID := pickFirstDocumentID(t, c)
	out, err := c.DeleteDocument(ctx(t), bankID(), docID)
	if err != nil {
		t.Fatalf("DeleteDocument: %v", err)
	}
	t.Logf("DeleteDocument => %s", fmt.Sprintf("%+v", out))
}
