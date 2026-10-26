// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight
//
// This file re-exports the upstream Hindsight types and constructors used by
// the public SDK surface so that callers (e.g. example code) do not need to
// import the upstream module directly. Aliases preserve the upstream methods
// (SetXxx / GetXxx) without wrapping each one.

package dumemory

import (
	hindsight "github.com/vectorize-io/hindsight/hindsight-clients/go"
)

// ---- Bank ----

// CreateBankRequest is the request body for CreateBank.
type CreateBankRequest = hindsight.CreateBankRequest

// NewCreateBankRequest builds an empty CreateBankRequest.
func NewCreateBankRequest() *CreateBankRequest { return hindsight.NewCreateBankRequest() }

// BankConfigUpdate is the request body for UpdateBankConfig.
type BankConfigUpdate = hindsight.BankConfigUpdate

// NewBankConfigUpdate builds a BankConfigUpdate carrying the given config overrides.
func NewBankConfigUpdate(config map[string]interface{}) *BankConfigUpdate {
	return hindsight.NewBankConfigUpdate(config)
}

// ConsolidationRequest is the request body for ConsolidateBank.
type ConsolidationRequest = hindsight.ConsolidationRequest

// NewConsolidationRequest builds an empty ConsolidationRequest.
func NewConsolidationRequest() *ConsolidationRequest { return hindsight.NewConsolidationRequest() }

// ---- Memory ----

// MemoryItem is a single item carried by RetainRequest.
type MemoryItem = hindsight.MemoryItem

// NewMemoryItem builds a MemoryItem with the given content string.
func NewMemoryItem(content string) *MemoryItem {
	return hindsight.NewMemoryItem(hindsight.Content{String: &content})
}

// RetainRequest is the request body for Retain / RetainAsync.
type RetainRequest = hindsight.RetainRequest

// NewRetainRequest builds a RetainRequest carrying the given items.
func NewRetainRequest(items []MemoryItem) *RetainRequest { return hindsight.NewRetainRequest(items) }

// RecallRequest is the request body for Recall.
type RecallRequest = hindsight.RecallRequest

// NewRecallRequest builds a RecallRequest with the given query string.
func NewRecallRequest(query string) *RecallRequest { return hindsight.NewRecallRequest(query) }

// ReflectRequest is the request body for Reflect.
type ReflectRequest = hindsight.ReflectRequest

// NewReflectRequest builds a ReflectRequest with the given query string.
func NewReflectRequest(query string) *ReflectRequest { return hindsight.NewReflectRequest(query) }

// ---- Documents ----

// UpdateDocumentRequest is the request body for UpdateDocument.
type UpdateDocumentRequest = hindsight.UpdateDocumentRequest

// NewUpdateDocumentRequest builds an empty UpdateDocumentRequest.
func NewUpdateDocumentRequest() *UpdateDocumentRequest { return hindsight.NewUpdateDocumentRequest() }

// ---- Mental models ----

// CreateMentalModelRequest is the request body for CreateMentalModel.
type CreateMentalModelRequest = hindsight.CreateMentalModelRequest

// NewCreateMentalModelRequest builds a CreateMentalModelRequest with required fields.
func NewCreateMentalModelRequest(name, sourceQuery string) *CreateMentalModelRequest {
	return hindsight.NewCreateMentalModelRequest(name, sourceQuery)
}

// UpdateMentalModelRequest is the request body for UpdateMentalModel.
type UpdateMentalModelRequest = hindsight.UpdateMentalModelRequest

// NewUpdateMentalModelRequest builds an empty UpdateMentalModelRequest.
func NewUpdateMentalModelRequest() *UpdateMentalModelRequest {
	return hindsight.NewUpdateMentalModelRequest()
}

// ---- Directives ----

// DirectiveList is the paginated response body for ListDirectives.
type DirectiveList = hindsight.DirectiveListResponse

// CreateDirectiveRequest is the request body for CreateDirective.
type CreateDirectiveRequest = hindsight.CreateDirectiveRequest

// NewCreateDirectiveRequest builds a CreateDirectiveRequest with required fields.
func NewCreateDirectiveRequest(tag, content string) *CreateDirectiveRequest {
	return hindsight.NewCreateDirectiveRequest(tag, content)
}

// UpdateDirectiveRequest is the request body for UpdateDirective.
type UpdateDirectiveRequest = hindsight.UpdateDirectiveRequest

// NewUpdateDirectiveRequest builds an empty UpdateDirectiveRequest.
func NewUpdateDirectiveRequest() *UpdateDirectiveRequest {
	return hindsight.NewUpdateDirectiveRequest()
}
