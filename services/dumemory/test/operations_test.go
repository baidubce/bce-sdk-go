// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import (
	"testing"

	"github.com/baidubce/bce-sdk-go/services/dumemory/api"
)

func TestListOperations(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)
	out, err := c.ListOperations(ctx(t), bankID(), dumemory.ListOperationsOptions{Limit: 10})
	if err != nil {
		t.Fatalf("ListOperations: %v", err)
	}
	t.Logf("ListOperations => %+v", out)
}

func TestGetOperationStatus(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)

	resp, err := c.ListOperations(ctx(t), bankID(), dumemory.ListOperationsOptions{Limit: 1})
	if err != nil {
		t.Skipf("ListOperations: %v", err)
	}
	if resp == nil || len(resp.Operations) == 0 {
		t.Skip("no operation available")
	}

	opID := resp.Operations[0].Id
	out, err := c.GetOperationStatus(ctx(t), bankID(), opID)
	if err != nil {
		t.Fatalf("GetOperationStatus: %v", err)
	}
	if out == nil {
		t.Fatal("GetOperationStatus returned nil response")
	}
	t.Logf("GetOperationStatus => %+v", out)
}

func TestCancelOperation(t *testing.T) {
	c := newClient(t)
	ensureBank(t, c)

	resp, err := c.ListOperations(ctx(t), bankID(), dumemory.ListOperationsOptions{Limit: 1, Status: "running"})
	if err != nil {
		t.Skipf("ListOperations: %v", err)
	}
	if resp == nil || len(resp.Operations) == 0 {
		t.Skip("no running operation to cancel")
	}
	opID := resp.Operations[0].Id
	out, err := c.CancelOperation(ctx(t), bankID(), opID)
	if err != nil {
		t.Fatalf("CancelOperation: %v", err)
	}
	t.Logf("CancelOperation => %+v", out)
}
