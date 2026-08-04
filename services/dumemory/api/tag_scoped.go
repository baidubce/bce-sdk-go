// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight

package dumemory

import (
	"context"
	"errors"

	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ErrMissingEntityScope is returned by scoped APIs when no entity id is set.
var ErrMissingEntityScope = errors.New("At least one entity ID (user_id, agent_id, app_id, or run_id) is required.")

const defaultScopedTagsMatch = "all_strict"

// EntityScope identifies the entity boundary used by the scoped APIs.
// At least one field must be set. Non-empty fields are converted into tags:
// user_id:<id>, agent_id:<id>, app_id:<id>, and run_id:<id>.
type EntityScope struct {
	UserID  string
	AgentID string
	AppID   string
	RunID   string
}

// Validate verifies that at least one entity id is present.
func (s EntityScope) Validate() error {
	if s.UserID == "" && s.AgentID == "" && s.AppID == "" && s.RunID == "" {
		return ErrMissingEntityScope
	}
	return nil
}

// Tags returns the normalized scope tags for this entity boundary.
func (s EntityScope) Tags() ([]string, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}
	tags := make([]string, 0, 4)
	if s.UserID != "" {
		tags = append(tags, "user_id:"+s.UserID)
	}
	if s.AgentID != "" {
		tags = append(tags, "agent_id:"+s.AgentID)
	}
	if s.AppID != "" {
		tags = append(tags, "app_id:"+s.AppID)
	}
	if s.RunID != "" {
		tags = append(tags, "run_id:"+s.RunID)
	}
	return tags, nil
}

// ListTagsOptions are optional filters for ListTagsWithScope.
type ListTagsOptions struct {
	Q      string
	Source string
	Limit  int32
	Offset int32
}

// ListTagsResponse is the response body for ListTagsWithScope.
type ListTagsResponse = hindsight.ListTagsResponse

// TagItem is a single tag entry returned by ListTagsWithScope.
type TagItem = hindsight.TagItem

// RetainWithScope writes memories after appending entity scope tags to each item
// and to document_tags. document_tags is exposed by the current generated client;
// adding scope tags there keeps batch-level documents aligned with their memory
// units. If callers already set document_tags or per-item tags, those tags are
// preserved and de-duplicated with the scope tags.
func (c *Client) RetainWithScope(ctx context.Context, bankID string, scope EntityScope, req hindsight.RetainRequest) (*hindsight.RetainResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	for i := range req.Items {
		req.Items[i].Tags = mergeTags(req.Items[i].Tags, scopeTags)
	}
	req.DocumentTags = mergeTags(req.DocumentTags, scopeTags)
	return c.Retain(ctx, bankID, req)
}

// RecallWithScope recalls memories within an entity scope. When TagsMatch is
// empty or still at the upstream constructor default "any", scoped filtering
// uses "all_strict" so all scope tags must match and untagged global memories
// are excluded.
func (c *Client) RecallWithScope(ctx context.Context, bankID string, scope EntityScope, req hindsight.RecallRequest) (*hindsight.RecallResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	req.Tags = mergeTags(req.Tags, scopeTags)
	req.TagsMatch = scopedTagsMatch(req.TagsMatch)
	return c.Recall(ctx, bankID, req)
}

// ReflectWithScope synthesizes an answer using memories within an entity scope.
// When TagsMatch is empty or still at the upstream constructor default "any",
// scoped filtering uses "all_strict" so all scope tags must match and untagged
// global memories are excluded.
func (c *Client) ReflectWithScope(ctx context.Context, bankID string, scope EntityScope, req hindsight.ReflectRequest) (*hindsight.ReflectResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	req.Tags = mergeTags(req.Tags, scopeTags)
	req.TagsMatch = scopedTagsMatch(req.TagsMatch)
	return c.Reflect(ctx, bankID, req)
}

// GetMemoryWithScope fetches a memory and lets callers verify the returned tags
// contain the requested scope. The upstream endpoint does not accept tag filters.
func (c *Client) GetMemoryWithScope(ctx context.Context, bankID, memoryID string, scope EntityScope) (interface{}, error) {
	if err := scope.Validate(); err != nil {
		return nil, err
	}
	out, _, err := c.hindsight.MemoryAPI.GetMemory(ctx, bankID, memoryID).Execute()
	return out, err
}

// ListTagsWithScope lists tags visible within an entity scope. If Q is empty,
// the query is scoped to the first scope tag with a wildcard (for example,
// user_id:123*). If Q is set, it is preserved and the scope still must be valid.
func (c *Client) ListTagsWithScope(ctx context.Context, bankID string, scope EntityScope, opts ListTagsOptions) (*hindsight.ListTagsResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	call := c.hindsight.MemoryAPI.ListTags(ctx, bankID)
	if opts.Q != "" {
		call = call.Q(opts.Q)
	} else if len(scopeTags) > 0 {
		call = call.Q(scopeTags[0] + "*")
	}
	if opts.Source != "" {
		call = call.Source(opts.Source)
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

// ListMemoriesWithScope lists memory units filtered by entity scope. When
// TagsMatch is empty or still at the upstream/default "any", scoped filtering
// uses "all_strict" so all scope tags must match and untagged global memories
// are excluded. Other ListMemoriesOptions filters are preserved.
func (c *Client) ListMemoriesWithScope(ctx context.Context, bankID string, scope EntityScope, opts ListMemoriesOptions) (*hindsight.ListMemoryUnitsResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	opts.Tags = mergeTags(opts.Tags, scopeTags)
	opts.TagsMatch = scopedTagsMatchValue(opts.TagsMatch)
	return c.ListMemories(ctx, bankID, opts)
}

// ListDocumentsWithScope lists documents filtered by entity scope. When
// TagsMatch is empty or still at the upstream/default "any", scoped filtering
// uses "all_strict" so all scope tags must match and untagged global documents
// are excluded.
func (c *Client) ListDocumentsWithScope(ctx context.Context, bankID string, scope EntityScope, opts ListDocumentsOptions) (*hindsight.ListDocumentsResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	opts.Tags = mergeTags(opts.Tags, scopeTags)
	opts.TagsMatch = scopedTagsMatchValue(opts.TagsMatch)
	return c.ListDocuments(ctx, bankID, opts)
}

// UpdateDocumentTagsWithScope updates document tags after appending entity scope
// tags. The server propagates document tags to associated memory units and marks
// derived observations stale for recomputation.
func (c *Client) UpdateDocumentTagsWithScope(ctx context.Context, bankID, documentID string, scope EntityScope, req hindsight.UpdateDocumentRequest) (*hindsight.UpdateDocumentResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	req.Tags = mergeTags(req.Tags, scopeTags)
	return c.UpdateDocument(ctx, bankID, documentID, req)
}

// ListDirectivesWithScope lists directives filtered by entity scope. When
// TagsMatch is empty or still at the upstream/default "any", scoped filtering
// uses "all_strict" so all scope tags must match and untagged global directives
// are excluded.
func (c *Client) ListDirectivesWithScope(ctx context.Context, bankID string, scope EntityScope, opts ListDirectivesOptions) (*hindsight.DirectiveListResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	opts.Tags = mergeTags(opts.Tags, scopeTags)
	opts.TagsMatch = scopedTagsMatchValue(opts.TagsMatch)
	return c.ListDirectives(ctx, bankID, opts)
}

// CreateDirectiveWithScope creates a directive after appending entity scope tags.
func (c *Client) CreateDirectiveWithScope(ctx context.Context, bankID string, scope EntityScope, req hindsight.CreateDirectiveRequest) (*hindsight.DirectiveResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	req.Tags = mergeTags(req.Tags, scopeTags)
	return c.CreateDirective(ctx, bankID, req)
}

// UpdateDirectiveWithScope updates a directive after appending entity scope tags.
func (c *Client) UpdateDirectiveWithScope(ctx context.Context, bankID, directiveID string, scope EntityScope, req hindsight.UpdateDirectiveRequest) (*hindsight.DirectiveResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	req.Tags = mergeTags(req.Tags, scopeTags)
	return c.UpdateDirective(ctx, bankID, directiveID, req)
}

// ListMentalModelsWithScope lists mental models filtered by entity scope. When
// TagsMatch is empty or still at the upstream/default "any", scoped filtering
// uses "all_strict" so all scope tags must match and untagged global mental
// models are excluded.
func (c *Client) ListMentalModelsWithScope(ctx context.Context, bankID string, scope EntityScope, opts ListMentalModelsOptions) (*hindsight.MentalModelListResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	opts.Tags = mergeTags(opts.Tags, scopeTags)
	opts.TagsMatch = scopedTagsMatchValue(opts.TagsMatch)
	return c.ListMentalModels(ctx, bankID, opts)
}

// CreateMentalModelWithScope creates a mental model after appending entity scope tags.
func (c *Client) CreateMentalModelWithScope(ctx context.Context, bankID string, scope EntityScope, req hindsight.CreateMentalModelRequest) (*hindsight.CreateMentalModelResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	req.Tags = mergeTags(req.Tags, scopeTags)
	return c.CreateMentalModel(ctx, bankID, req)
}

// UpdateMentalModelWithScope updates a mental model after appending entity scope tags.
func (c *Client) UpdateMentalModelWithScope(ctx context.Context, bankID, modelID string, scope EntityScope, req hindsight.UpdateMentalModelRequest) (*hindsight.MentalModelResponse, error) {
	scopeTags, err := scope.Tags()
	if err != nil {
		return nil, err
	}
	req.Tags = mergeTags(req.Tags, scopeTags)
	return c.UpdateMentalModel(ctx, bankID, modelID, req)
}

func scopedTagsMatch(tagsMatch *string) *string {
	if tagsMatch != nil && *tagsMatch != "" && *tagsMatch != "any" {
		return tagsMatch
	}
	value := defaultScopedTagsMatch
	return &value
}

func scopedTagsMatchValue(tagsMatch string) string {
	if tagsMatch != "" && tagsMatch != "any" {
		return tagsMatch
	}
	return defaultScopedTagsMatch
}

func mergeTags(base, extra []string) []string {
	if len(base) == 0 && len(extra) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(base)+len(extra))
	out := make([]string, 0, len(base)+len(extra))
	for _, tag := range base {
		if tag == "" {
			continue
		}
		if _, ok := seen[tag]; ok {
			continue
		}
		seen[tag] = struct{}{}
		out = append(out, tag)
	}
	for _, tag := range extra {
		if tag == "" {
			continue
		}
		if _, ok := seen[tag]; ok {
			continue
		}
		seen[tag] = struct{}{}
		out = append(out, tag)
	}
	return out
}
