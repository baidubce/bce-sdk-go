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
	"github.com/baidubce/bce-sdk-go/bce"
)

func StartExecution(cli bce.Client, args *StartExecutionArgs) (*Execution, error) {
	op := &Operation{
		HTTPUri:    getExecutionStartUri(),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &Execution{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*Execution); ok {
		return value, nil
	}
	return nil, nil
}

func StopExecution(cli bce.Client, flowName, executionName string) (*Execution, error) {
	op := &Operation{
		HTTPUri:    getExecutionStopUri(),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Body: map[string]string{"flowName": flowName, "executionName": executionName},
	}
	result := &cfcResult{
		Result: &Execution{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*Execution); ok {
		return value, nil
	}
	return nil, nil
}

func ListExecutions(cli bce.Client, flowName string) (*ListExecutionsResult, error) {
	op := &Operation{
		HTTPUri:    getExecutionsUri(),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Params: map[string]interface{}{"flowName": flowName},
	}
	result := &cfcResult{
		Result: &ListExecutionsResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListExecutionsResult); ok {
		return value, nil
	}
	return nil, nil
}

func DescribeExecution(cli bce.Client, flowName, executionName string) (*Execution, error) {
	op := &Operation{
		HTTPUri:    getExecutionUri(),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Params: map[string]interface{}{
			"flowName":      flowName,
			"executionName": executionName,
		},
	}
	result := &cfcResult{
		Result: &Execution{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*Execution); ok {
		return value, nil
	}
	return nil, nil
}

func GetExecutionHistory(cli bce.Client, args *GetExecutionHistoryArgs) (*GetExecutionHistoryResult, error) {
	op := &Operation{
		HTTPUri:    getExecutionHistory(),
		HTTPMethod: GET,
	}
	if args.Limit == 0 {
		args.Limit = 100
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"flowName":      args.FlowName,
			"executionName": args.ExecutionName,
			"limit":         args.Limit,
		},
	}
	result := &cfcResult{
		Result: &GetExecutionHistoryResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*GetExecutionHistoryResult); ok {
		return value, nil
	}
	return nil, nil
}

func ListFlow(cli bce.Client) (*ListFlowResult, error) {
	op := &Operation{
		HTTPUri:    getWorkflowsUri(),
		HTTPMethod: GET,
	}
	request := &cfcRequest{}
	result := &cfcResult{
		Result: &ListFlowResult{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListFlowResult); ok {
		return value, nil
	}
	return nil, nil
}

func CreateFlow(cli bce.Client, args *CreateUpdateFlowArgs) (*Flow, error) {
	op := &Operation{
		HTTPUri:    getWorkflowUri(),
		HTTPMethod: POST,
	}
	return createUpdateFlow(cli, args, op)
}

func UpdateFlow(cli bce.Client, args *CreateUpdateFlowArgs) (*Flow, error) {
	op := &Operation{
		HTTPUri:    getWorkflowUri(),
		HTTPMethod: PUT,
	}

	return createUpdateFlow(cli, args, op)
}

func createUpdateFlow(cli bce.Client, args *CreateUpdateFlowArgs, op *Operation) (*Flow, error) {
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &Flow{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*Flow); ok {
		return value, nil
	}
	return nil, nil
}

func DescribeFlow(cli bce.Client, flowName string) (*Flow, error) {
	op := &Operation{
		HTTPUri:    getWorkflowUri(),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Params: map[string]interface{}{"flowName": flowName},
	}
	result := &cfcResult{
		Result: &Flow{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*Flow); ok {
		return value, nil
	}
	return nil, nil
}

func DeleteFlow(cli bce.Client, flowName string) error {
	op := &Operation{
		HTTPUri:    getWorkflowUri(),
		HTTPMethod: DELETE,
	}
	request := &cfcRequest{
		Params: map[string]interface{}{"flowName": flowName},
	}
	result := &cfcResult{}
	err := caller(cli, op, request, result)
	if err != nil {
		return err
	}
	return nil
}
