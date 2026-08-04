// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package dumemory

import (
	"context"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ListMemoriesOptions are optional filters for ListMemories.
type ListMemoriesOptions struct {
	Type               string
	Q                  string
	ConsolidationState string
	State              string
	DocumentID         string
	EntityID           string
	Tags               []string
	TagsMatch          string
	Limit              int32
	Offset             int32
}

// Retain synchronously writes a memory via POST /memories/retain.
// The upstream API is keyed by bankId on the path; pass the same bank in req.
func (c *Client) Retain(ctx context.Context, bankID string, req hindsight.RetainRequest) (*hindsight.RetainResponse, error) {
	out, _, err := c.hindsight.MemoryAPI.RetainMemories(ctx, bankID).RetainRequest(req).Execute()
	return out, err
}

// RetainAsync asynchronously writes a memory.
// Hindsight exposes a single RetainMemories endpoint; async is selected via
// the request body. This wrapper sets req.Async = true (when supported by
// the deployed server) and otherwise behaves like Retain.
func (c *Client) RetainAsync(ctx context.Context, bankID string, req hindsight.RetainRequest) (*hindsight.RetainResponse, error) {
	asynchronous := true
	req.Async = &asynchronous
	out, _, err := c.hindsight.MemoryAPI.RetainMemories(ctx, bankID).RetainRequest(req).Execute()
	return out, err
}

// Recall searches memories via POST /recall.
func (c *Client) Recall(ctx context.Context, bankID string, req hindsight.RecallRequest) (*hindsight.RecallResponse, error) {
	out, _, err := c.hindsight.MemoryAPI.RecallMemories(ctx, bankID).RecallRequest(req).Execute()
	return out, err
}

// Reflect synthesises an answer via POST /reflect.
func (c *Client) Reflect(ctx context.Context, bankID string, req hindsight.ReflectRequest) (*hindsight.ReflectResponse, error) {
	out, _, err := c.hindsight.MemoryAPI.Reflect(ctx, bankID).ReflectRequest(req).Execute()
	return out, err
}

// ListMemories lists memories via GET /list.
func (c *Client) ListMemories(ctx context.Context, bankID string, opts ListMemoriesOptions) (*hindsight.ListMemoryUnitsResponse, error) {
	call := c.hindsight.MemoryAPI.ListMemories(ctx, bankID)
	if opts.Type != "" {
		call = call.Type_(opts.Type)
	}
	if opts.Q != "" {
		call = call.Q(opts.Q)
	}
	if opts.ConsolidationState != "" {
		call = call.ConsolidationState(opts.ConsolidationState)
	}
	if opts.State != "" {
		call = call.State(opts.State)
	}
	if opts.DocumentID != "" {
		call = call.DocumentId(opts.DocumentID)
	}
	if opts.EntityID != "" {
		call = call.EntityId(opts.EntityID)
	}
	if len(opts.Tags) > 0 {
		call = call.Tags(opts.Tags)
	}
	if opts.TagsMatch != "" {
		call = call.TagsMatch(opts.TagsMatch)
	}
	if opts.Limit != 0 {
		call = call.Limit(opts.Limit)
	}
	if opts.Offset != 0 {
		call = call.Offset(opts.Offset)
	}
	out, _, err := call.Execute()
	return out, err
}
