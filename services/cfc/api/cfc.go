/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func ListFunctions(cli bce.Client, args *ListFunctionsArgs) (*ListFunctionsResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionsUri(),
		HTTPMethod: GET,
	}

	if args.MaxItems <= 0 {
		args.MaxItems = 10000
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"FunctionVersion": args.FunctionVersion,
			"Marker":          args.Marker,
			"MaxItems":        args.MaxItems,
		},
	}
	result := &cfcResult{
		Result: &ListFunctionsResult{},
	}

	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListFunctionsResult); ok {
		return value, nil
	}
	return nil, nil
}

func GetFunction(cli bce.Client, args *GetFunctionArgs) (*GetFunctionResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionUri(args.FunctionName),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"Qualifier": args.Qualifier,
		},
	}
	result := &cfcResult{
		Result: &GetFunctionResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*GetFunctionResult); ok {
		return value, nil
	}
	return nil, nil
}

func CreateFunction(cli bce.Client, args *CreateFunctionArgs) (*CreateFunctionResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionsUri(),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &CreateFunctionResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*CreateFunctionResult); ok {
		return value, nil
	}
	return nil, nil
}

func CreateFunctionByBlueprint(cli bce.Client, args *CreateFunctionByBlueprintArgs) (*CreateFunctionResult, error) {
	op := &Operation{
		HTTPUri:    getBlueprintUri(args.BlueprintID),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &CreateFunctionResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*CreateFunctionResult); ok {
		return value, nil
	}
	return nil, nil
}

func DeleteFunction(cli bce.Client, args *DeleteFunctionArgs) error {
	op := &Operation{
		HTTPUri:    getFunctionUri(args.FunctionName),
		HTTPMethod: DELETE,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"Qualifier": args.Qualifier,
		},
	}
	result := &cfcResult{
		Result: nil,
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFunctionCode(cli bce.Client, args *UpdateFunctionCodeArgs) (*UpdateFunctionCodeResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionCodeUri(args.FunctionName),
		HTTPMethod: PUT,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &UpdateFunctionCodeResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*UpdateFunctionCodeResult); ok {
		return value, nil
	}
	return nil, nil
}

func GetFunctionConfiguration(cli bce.Client, args *GetFunctionConfigurationArgs) (*GetFunctionConfigurationResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionConfigurationUri(args.FunctionName),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"Qualifier": args.Qualifier,
		},
	}
	result := &cfcResult{
		Result: &GetFunctionConfigurationResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*GetFunctionConfigurationResult); ok {
		return value, nil
	}
	return nil, nil
}

func UpdateFunctionConfiguration(cli bce.Client, args *UpdateFunctionConfigurationArgs) (*UpdateFunctionConfigurationResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionConfigurationUri(args.FunctionName),
		HTTPMethod: PUT,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &UpdateFunctionConfigurationResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*UpdateFunctionConfigurationResult); ok {
		return value, nil
	}
	return nil, nil
}

func SetReservedConcurrentExecutions(cli bce.Client, args *ReservedConcurrentExecutionsArgs) error {
	op := &Operation{
		HTTPUri:    getFunctionConCurrentUri(args.FunctionName),
		HTTPMethod: PUT,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: nil,
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return err
	}
	return nil
}

func DeleteReservedConcurrentExecutions(cli bce.Client, args *DeleteReservedConcurrentExecutionsArgs) error {
	op := &Operation{
		HTTPUri:    getFunctionConCurrentUri(args.FunctionName),
		HTTPMethod: DELETE,
	}
	request := &cfcRequest{
		Args: args,
	}
	result := &cfcResult{}
	err := caller(cli, op, request, result)
	if err != nil {
		return err
	}
	return nil
}

func Invocations(cli bce.Client, args *InvocationsArgs) (*InvocationsResult, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}
	if len(args.InvocationType) == 0 {
		args.InvocationType = InvocationTypeRequestResponse
	}
	if len(args.LogType) == 0 {
		args.LogType = LogTypeNone
	}

	req := &bce.BceRequest{}
	http.SetResponseHeaderTimeout(DefaultMaxFunctionTimeout)
	req.SetRequestId(args.RequestId)
	req.SetUri(getInvocationsUri(args.FunctionName))
	req.SetMethod(http.POST)
	req.SetParam("invocationType", string(args.InvocationType))
	req.SetParam("logType", string(args.LogType))
	if args.Qualifier != "" {
		req.SetParam("qualifier", args.Qualifier)
	}

	if args.Payload != nil {
		var payloadBytes []byte
		var err error
		switch args.Payload.(type) {
		case string:
			payloadBytes = []byte(args.Payload.(string))
		case []byte:
			payloadBytes = args.Payload.([]byte)
		default:
			payloadBytes, err = json.Marshal(args.Payload)
			if err != nil {
				return nil, err
			}
		}
		var js interface{}
		if json.Unmarshal([]byte(payloadBytes), &js) != nil {
			return nil, ParseJsonError
		}
		requestBody, err := bce.NewBodyFromBytes(payloadBytes)
		if err != nil {
			return nil, err
		}
		req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
		req.SetBody(requestBody)
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	defer resp.Body().Close()
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	body, err := ioutil.ReadAll(resp.Body())
	if err != nil {
		return nil, err
	}
	result := &InvocationsResult{
		Payload: string(body),
	}
	errorStr := resp.Header("x-bce-function-error")
	if len(errorStr) > 0 {
		result.FunctionError = errorStr
	}
	logResult := resp.Header("x-bce-log-result")
	if len(logResult) > 0 {
		decodeBytes, err := base64.StdEncoding.DecodeString(string(logResult))
		if err != nil {
			return nil, err
		}
		result.LogResult = string(decodeBytes)
	}
	return result, nil
}

func ListVersionsByFunction(cli bce.Client, args *ListVersionsByFunctionArgs) (*ListVersionsByFunctionResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionVersionsUri(args.FunctionName),
		HTTPMethod: GET,
	}
	if args.MaxItems <= 0 {
		args.MaxItems = 10000
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"Marker":   args.Marker,
			"MaxItems": args.MaxItems,
		},
	}
	result := &cfcResult{
		Result: &ListVersionsByFunctionResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListVersionsByFunctionResult); ok {
		return value, nil
	}
	return nil, nil
}

func PublishVersion(cli bce.Client, args *PublishVersionArgs) (*PublishVersionResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionVersionsUri(args.FunctionName),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: map[string]interface{}{
			"Description": args.Description,
			"CodeSha256":  args.CodeSha256,
		},
	}
	result := &cfcResult{
		Result: &PublishVersionResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*PublishVersionResult); ok {
		return value, nil
	}
	return nil, nil
}

func ListAliases(cli bce.Client, args *ListAliasesArgs) (*ListAliasesResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionAliasesUri(args.FunctionName),
		HTTPMethod: GET,
	}
	if args.MaxItems <= 0 {
		args.MaxItems = 10000
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"FunctionVersion": args.FunctionVersion,
			"Marker":          args.Marker,
			"MaxItems":        args.MaxItems,
		},
	}
	result := &cfcResult{
		Result: &ListAliasesResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListAliasesResult); ok {
		return value, nil
	}
	return nil, nil
}

func CreateAlias(cli bce.Client, args *CreateAliasArgs) (*CreateAliasResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionAliasesUri(args.FunctionName),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: map[string]interface{}{
			"FunctionVersion": args.FunctionVersion,
			"Name":            args.Name,
			"Description":     args.Description,
		},
	}
	result := &cfcResult{
		Result: &CreateAliasResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*CreateAliasResult); ok {
		return value, nil
	}
	return nil, nil
}

func GetAlias(cli bce.Client, args *GetAliasArgs) (*GetAliasResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionAliasUri(args.FunctionName, args.AliasName),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
	}
	result := &cfcResult{
		Result: &GetAliasResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*GetAliasResult); ok {
		return value, nil
	}
	return nil, nil
}

func UpdateAlias(cli bce.Client, args *UpdateAliasArgs) (*UpdateAliasResult, error) {
	op := &Operation{
		HTTPUri:    getFunctionAliasUri(args.FunctionName, args.AliasName),
		HTTPMethod: PUT,
	}
	request := &cfcRequest{
		Args: args,
		Body: map[string]interface{}{
			"FunctionVersion": args.FunctionVersion,
			"Description":     args.Description,
		},
	}
	result := &cfcResult{
		Result: &UpdateAliasResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*UpdateAliasResult); ok {
		return value, nil
	}
	return nil, nil
}

func DeleteAlias(cli bce.Client, args *DeleteAliasArgs) error {
	op := &Operation{
		HTTPUri:    getFunctionAliasUri(args.FunctionName, args.AliasName),
		HTTPMethod: DELETE,
	}
	request := &cfcRequest{
		Args: args,
	}
	result := &cfcResult{}
	err := caller(cli, op, request, result)
	if err != nil {
		return err
	}
	return nil
}

func ListTriggers(cli bce.Client, args *ListTriggersArgs) (*ListTriggersResult, error) {
	op := &Operation{
		HTTPUri:    getTriggerUri(),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"FunctionBrn": args.FunctionBrn,
			"ScopeType":   args.ScopeType,
		},
	}
	result := &cfcResult{
		Result: &ListTriggersResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListTriggersResult); ok {
		return value, nil
	}
	return nil, nil
}

func CreateTrigger(cli bce.Client, args *CreateTriggerArgs) (*CreateTriggerResult, error) {
	op := &Operation{
		HTTPUri:    getTriggerUri(),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &CreateTriggerResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*CreateTriggerResult); ok {
		return value, nil
	}
	return nil, nil

}

func UpdateTrigger(cli bce.Client, args *UpdateTriggerArgs) (*UpdateTriggerResult, error) {
	op := &Operation{
		HTTPUri:    getTriggerUri(),
		HTTPMethod: PUT,
	}
	request := &cfcRequest{
		Args: args,
		Body: map[string]interface{}{
			"RelationId": args.RelationId,
			"Target":     args.Target,
			"Source":     args.Source,
			"Data":       args.Data,
		},
	}
	result := &cfcResult{
		Result: &UpdateTriggerResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*UpdateTriggerResult); ok {
		return value, nil
	}
	return nil, nil
}

func DeleteTrigger(cli bce.Client, args *DeleteTriggerArgs) error {
	op := &Operation{
		HTTPUri:    getTriggerUri(),
		HTTPMethod: DELETE,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"RelationId": args.RelationId,
			"Target":     args.Target,
			"Source":     string(args.Source),
		},
	}
	result := &cfcResult{}
	err := caller(cli, op, request, result)
	if err != nil {
		return err
	}
	return nil
}

func caller(cli bce.Client, op *Operation, request *cfcRequest, response *cfcResult) error {
	if request.Args != nil {
		if err := request.Args.(Validator).Validate(); err != nil {
			return err
		}
	}
	req := new(bce.BceRequest)
	if op.HTTPUri == "" {
		return URIIllegal
	}
	req.SetUri(op.HTTPUri)
	if op.HTTPMethod == "" {
		return MethodIllegal
	}
	req.SetMethod(op.HTTPMethod)
	if request.Params != nil {
		for key, value := range request.Params {
			rv := reflect.ValueOf(value)
			switch rv.Kind() {
			case reflect.String:
				req.SetParam(key, value.(string))
			case reflect.Int:
				req.SetParam(key, strconv.Itoa(value.(int)))
			case reflect.Struct:
				if valueBytes, err := json.Marshal(value); err != nil {
					req.SetParam(key, string(valueBytes))
				}
			}
		}
	}
	if request.Body != nil {
		argsBytes, err := json.Marshal(request.Body)
		if err != nil {
			return err
		}
		requestBody, err := bce.NewBodyFromBytes(argsBytes)
		if err != nil {
			return err
		}
		req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
		req.SetBody(requestBody)
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	defer resp.Body().Close()
	if resp.IsFail() {
		return resp.ServiceError()
	}

	if response.Result != nil {
		if err := resp.ParseJsonBody(response.Result); err != nil {
			return err
		}
	}
	return nil
}

// ListEventSource
func ListEventSource(cli bce.Client, args *ListEventSourceArgs) (*ListEventSourceResult, error) {
	op := &Operation{
		HTTPUri:    getEventSourceUri(),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"FunctionName": args.FunctionName,
			"Marker":       args.Marker,
			"MaxItems":     args.MaxItems,
		},
	}
	result := &cfcResult{
		Result: &ListEventSourceResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListEventSourceResult); ok {
		return value, nil
	}
	return nil, nil
}

// GetEventSource
func GetEventSource(cli bce.Client, args *GetEventSourceArgs) (*GetEventSourceResult, error) {
	op := &Operation{
		HTTPUri:    getOneEventSourceUri(args.UUID),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
	}
	result := &cfcResult{
		Result: &GetEventSourceResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*GetEventSourceResult); ok {
		return value, nil
	}
	return nil, nil
}

// CreateEventSource
func CreateEventSource(cli bce.Client, args *CreateEventSourceArgs) (*CreateEventSourceResult, error) {
	op := &Operation{
		HTTPUri:    getEventSourceUri(),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &CreateEventSourceResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*CreateEventSourceResult); ok {
		return value, nil
	}
	return nil, nil
}

// UpdateEventSource
func UpdateEventSource(cli bce.Client, args *UpdateEventSourceArgs) (*UpdateEventSourceResult, error) {
	op := &Operation{
		HTTPUri:    getOneEventSourceUri(args.UUID),
		HTTPMethod: PUT,
	}
	request := &cfcRequest{
		Args: args,
		Body: args.FuncEventSource,
	}
	result := &cfcResult{
		Result: &UpdateEventSourceResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*UpdateEventSourceResult); ok {
		return value, nil
	}
	return nil, nil
}

func DeleteEventSource(cli bce.Client, args *DeleteEventSourceArgs) error {
	op := &Operation{
		HTTPUri:    getOneEventSourceUri(args.UUID),
		HTTPMethod: DELETE,
	}
	request := &cfcRequest{
		Args: args,
	}
	result := &cfcResult{}
	err := caller(cli, op, request, result)
	if err != nil {
		return err
	}
	return nil
}
