// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package dumemory

import (
	"context"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ListOperationsOptions are optional filters for ListOperations.
type ListOperationsOptions struct {
	Status         string
	Type           string
	Limit          int32
	Offset         int32
	ExcludeParents bool
}

// ListOperations lists background operations via GET /operations/{bankId}.
func (c *Client) ListOperations(ctx context.Context, bankID string, opts ListOperationsOptions) (*hindsight.OperationsListResponse, error) {
	call := c.hindsight.OperationsAPI.ListOperations(ctx, bankID)
	if opts.Status != "" {
		call = call.Status(opts.Status)
	}
	if opts.Type != "" {
		call = call.Type_(opts.Type)
	}
	if opts.Limit != 0 {
		call = call.Limit(opts.Limit)
	}
	if opts.Offset != 0 {
		call = call.Offset(opts.Offset)
	}
	if opts.ExcludeParents {
		call = call.ExcludeParents(true)
	}
	out, _, err := call.Execute()
	return out, err
}

// GetOperationStatus gets a single operation via GET /operations/{bankId}/{operationId}.
func (c *Client) GetOperationStatus(ctx context.Context, bankID, operationID string) (*hindsight.OperationStatusResponse, error) {
	out, _, err := c.hindsight.OperationsAPI.GetOperationStatus(ctx, bankID, operationID).Execute()
	return out, err
}

// CancelOperation cancels a single operation via DELETE /operations/{bankId}/{operationId}.
//
// Note: the Baidu reference table lists DELETE /operations/{bankId} as
// "cancel background operations". Hindsight implements this as a per-
// operation cancel keyed by operationId; pass the targeted operation id.
func (c *Client) CancelOperation(ctx context.Context, bankID, operationID string) (*hindsight.CancelOperationResponse, error) {
	out, _, err := c.hindsight.OperationsAPI.CancelOperation(ctx, bankID, operationID).Execute()
	return out, err
}
