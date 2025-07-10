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

// client.go - define the client for CFC service

// Package cfc defines the CFC services of BCE. The supported APIs are all defined in sub-package

package cfc

import (
	"errors"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cfc/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "cfc." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client of CFC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the CFC service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func newClient(ak, sk, token, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 || len(sk) == 0 {
		return nil, errors.New("ak sk illegal")
	}
	if len(token) == 0 {
		credentials, err = auth.NewBceCredentials(ak, sk)
	} else {
		credentials, err = auth.NewSessionBceCredentials(ak, sk, token)
	}
	if err != nil {
		return nil, err
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}
	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

func NewClient(ak, sk, endpoint string) (*Client, error) {
	return newClient(ak, sk, "", endpoint)
}

// Invocations - invocation a cfc function with specific parameters
//
// PARAMS:
//   - args: the arguments to invocation cfc function
//
// RETURNS:
//   - *api.InvocationsResult: the result of invocation cfc function
//   - error: nil if success otherwise the specific error
func (c *Client) Invocations(args *api.InvocationsArgs) (*api.InvocationsResult, error) {
	return api.Invocations(c, args)
}

// Invoke - invoke a cfc function, the same as Invocations
//
// PARAMS:
//   - args: the arguments to invocation cfc function
//
// RETURNS:
//   - *api.InvocationsResult: the result of invocation cfc function
//   - error: nil if success otherwise the specific error
func (c *Client) Invoke(args *api.InvocationsArgs) (*api.InvocationsResult, error) {
	return api.Invocations(c, args)
}

// ListFunctions - list all functions with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all functions
//
// RETURNS:
//   - *api.ListFunctionsResult: the result of list all functions
//   - error: nil if success otherwise the specific error
func (c *Client) ListFunctions(args *api.ListFunctionsArgs) (*api.ListFunctionsResult, error) {
	return api.ListFunctions(c, args)
}

// GetFunction - get a specific cfc function
//
// PARAMS:
//   - args: the arguments to get a specific cfc function
//
// RETURNS:
//   - *api.GetFunctionResult: the result of get function
//   - error: nil if success otherwise the specific error
func (c *Client) GetFunction(args *api.GetFunctionArgs) (*api.GetFunctionResult, error) {
	return api.GetFunction(c, args)
}

// CreateFunction - create a cfc function with specific parameters
//
// PARAMS:
//   - args: the arguments to create a cfc function
//
// RETURNS:
//   - *api.CreateFunctionResult: the result of create a cfc function, it contains function information
//   - error: nil if success otherwise the specific error
func (c *Client) CreateFunction(args *api.CreateFunctionArgs) (*api.CreateFunctionResult, error) {
	return api.CreateFunction(c, args)
}

// CreateFunctionByBlueprint - create a cfc function by a blueprint
//
// PARAMS:
//   - args: the arguments to create a cfc function
//
// RETURNS:
//   - *api.CreateFunctionResult: the result of create a cfc function, it contains function information
//   - error: nil if success otherwise the specific error
func (c *Client) CreateFunctionByBlueprint(args *api.CreateFunctionByBlueprintArgs) (*api.CreateFunctionResult, error) {
	return api.CreateFunctionByBlueprint(c, args)
}

// DeleteFunction - delete a specific cfc function
//
// PARAMS:
//   - args: the arguments to delete cfc function
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteFunction(args *api.DeleteFunctionArgs) error {
	return api.DeleteFunction(c, args)
}

// UpdateFunctionCode - update a cfc function code
//
// PARAMS:
//   - args: the arguments to update function code
//
// RETURNS:
//   - *api.UpdateFunctionCodeResult: the result of update function code
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateFunctionCode(args *api.UpdateFunctionCodeArgs) (*api.UpdateFunctionCodeResult, error) {
	return api.UpdateFunctionCode(c, args)
}

// GetFunctionConfiguration - get a specific cfc function configuration
//
// PARAMS:
//   - args: the arguments to get function configuration
//
// RETURNS:
//   - *api.GetFunctionConfigurationResult: the result of function configuration
//   - error: nil if success otherwise the specific error
func (c *Client) GetFunctionConfiguration(args *api.GetFunctionConfigurationArgs) (*api.GetFunctionConfigurationResult, error) {
	return api.GetFunctionConfiguration(c, args)
}

// UpdateFunctionConfiguration - update a specific cfc function configuration
//
// PARAMS:
//   - args: the arguments to update cfc function
//
// RETURNS:
//   - *api.UpdateFunctionConfigurationResult: the result of update function configuration
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateFunctionConfiguration(args *api.UpdateFunctionConfigurationArgs) (*api.UpdateFunctionConfigurationResult, error) {
	return api.UpdateFunctionConfiguration(c, args)
}

// ListVersionsByFunction - list all versions about a specific cfc function
//
// PARAMS:
//   - args: the arguments to list all versions
//
// RETURNS:
//   - *api.ListVersionsByFunctionResult: the result of all versions information
//   - error: nil if success otherwise the specific error
func (c *Client) ListVersionsByFunction(args *api.ListVersionsByFunctionArgs) (*api.ListVersionsByFunctionResult, error) {
	return api.ListVersionsByFunction(c, args)
}

// PublishVersion - publish a cfc function as a new version
//
// PARAMS:
//   - args: the arguments to publish a version
//
// RETURNS:
//   - *api.PublishVersionResult: the result of publish a function version
//   - error: nil if success otherwise the specific error
func (c *Client) PublishVersion(args *api.PublishVersionArgs) (*api.PublishVersionResult, error) {
	return api.PublishVersion(c, args)
}

// ListAliases - list all alias about a specific cfc function with specific parameters
//
// PARAMS:
//   - args: the arguments to list all alias
//
// RETURNS:
//   - *api.ListAliasesResult: the result of list all alias
//   - error: nil if success otherwise the specific error
func (c *Client) ListAliases(args *api.ListAliasesArgs) (*api.ListAliasesResult, error) {
	return api.ListAliases(c, args)
}

// CreateAlias - create an alias which bind one specific cfc function version
//
// PARAMS:
//   - args: the arguments to create an alias
//
// RETURNS:
//   - *api.CreateAliasResult: the result of create alias
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAlias(args *api.CreateAliasArgs) (*api.CreateAliasResult, error) {
	return api.CreateAlias(c, args)
}

// GetAlias - get alias information which bind one cfc function
//
// PARAMS:
//   - args: the arguments to get an alias
//
// RETURNS:
//   - *api.GetAliasResult: the result of get alias
//   - error: nil if success otherwise the specific error
func (c *Client) GetAlias(args *api.GetAliasArgs) (*api.GetAliasResult, error) {
	return api.GetAlias(c, args)
}

// UpdateAlias - update an alias configuration
//
// PARAMS:
//   - args: the arguments to update an alias
//
// RETURNS:
//   - *api.UpdateAliasResult: the result of update an alias
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAlias(args *api.UpdateAliasArgs) (*api.UpdateAliasResult, error) {
	return api.UpdateAlias(c, args)
}

// DeleteAlias - delete an alias
//
// PARAMS:
//   - args: the arguments to delete an alias
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteAlias(args *api.DeleteAliasArgs) error {
	return api.DeleteAlias(c, args)
}

// ListTriggers - list all triggers in one cfc function version
//
// PARAMS:
//   - args: the arguments to list all triggers
//
// RETURNS:
//   - *api.ListTriggersResult: the result of list all triggers
//   - error: nil if success otherwise the specific error
func (c *Client) ListTriggers(args *api.ListTriggersArgs) (*api.ListTriggersResult, error) {
	return api.ListTriggers(c, args)
}

// CreateTrigger - create a specific trigger
//
// PARAMS:
//   - args: the arguments to create a trigger
//
// RETURNS:
//   - *api.CreateTriggerResult: the result of create a trigger
//   - error: nil if success otherwise the specific error
func (c *Client) CreateTrigger(args *api.CreateTriggerArgs) (*api.CreateTriggerResult, error) {
	return api.CreateTrigger(c, args)
}

// UpdateTrigger - update a trigger
//
// PARAMS:
//   - args: the arguments to update a trigger
//
// RETURNS:
//   - *api.UpdateTriggerResult: the result of update a trigger
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateTrigger(args *api.UpdateTriggerArgs) (*api.UpdateTriggerResult, error) {
	return api.UpdateTrigger(c, args)
}

// DeleteTrigger - delete a trigger
//
// PARAMS:
//   - args: the arguments to delete a trigger
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteTrigger(args *api.DeleteTriggerArgs) error {
	return api.DeleteTrigger(c, args)
}

// SetReservedConcurrentExecutions - set a cfc function reserved concurrent executions
//
// PARAMS:
//   - args: the arguments to set reserved concurrent executions
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) SetReservedConcurrentExecutions(args *api.ReservedConcurrentExecutionsArgs) error {
	return api.SetReservedConcurrentExecutions(c, args)
}

// DeleteReservedConcurrentExecutions - delete one cfc function reserved concurrent executions setting
//
// PARAMS:
//   - args: the arguments to delete reserved concurrent executions setting
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteReservedConcurrentExecutions(args *api.DeleteReservedConcurrentExecutionsArgs) error {
	return api.DeleteReservedConcurrentExecutions(c, args)
}

// ListEventSource - list all event source mapping settings in one cfc function version
//
// PARAMS:
//   - args: the arguments to list all event source mapping settings
//
// RETURNS:
//   - *api.ListEventSourceResult: the result of list all event source mapping settings
//   - error: nil if success otherwise the specific error
func (c *Client) ListEventSource(args *api.ListEventSourceArgs) (*api.ListEventSourceResult, error) {
	return api.ListEventSource(c, args)
}

// GetEventSource - get info for a event source mapping setting
//
// PARAMS:
//   - args: the arguments to get a event source mapping setting
//
// RETURNS:
//   - *api.GetEventSourceResult: the result of get a event source mapping
//   - error: nil if success otherwise the specific error
func (c *Client) GetEventSource(args *api.GetEventSourceArgs) (*api.GetEventSourceResult, error) {
	return api.GetEventSource(c, args)
}

// UpdateEventSource - update a event source mapping setting
//
// PARAMS:
//   - args: the arguments to update a event source mapping
//
// RETURNS:
//   - *api.UpdateEventSourceResult: the result of update a event source mapping
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateEventSource(args *api.UpdateEventSourceArgs) (*api.UpdateEventSourceResult, error) {
	return api.UpdateEventSource(c, args)
}

// CreateEventSource - create a event source mapping setting
//
// PARAMS:
//   - args: the arguments to create a event source mapping setting
//
// RETURNS:
//   - *api.CreateEventSourceResult: the result of create event source mapping setting
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEventSource(args *api.CreateEventSourceArgs) (*api.CreateEventSourceResult, error) {
	return api.CreateEventSource(c, args)
}

// DeleteEventSource - delete one cfc event source mapping setting
//
// PARAMS:
//   - args: the arguments to delete cfc event source mapping setting
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteEventSource(args *api.DeleteEventSourceArgs) error {
	return api.DeleteEventSource(c, args)
}

// StartExecution - start an execution of a workflow
//
// PARAMS:
//   - args: the arguments to start execution, include flowName、executionName(optional) and input
//     input is a serialized json string, for example: "{\"fruits\":[\"apple\", \"banana\"]}".
//
// RETURNS:
//   - *api.Execution: information about the execution
//   - error: nil if success otherwise the specific error
func (c *Client) StartExecution(args *api.StartExecutionArgs) (*api.Execution, error) {
	return api.StartExecution(c, args)
}

// StopExecution - stop a running execution of a workflow
//
// PARAMS:
//   - flowName: which flow the execution belongs to
//   - executionName: which execution to stop
//
// RETURNS:
//   - *api.Execution: information about the execution
//   - error: nil if success otherwise the specific error
func (c *Client) StopExecution(flowName, executionName string) (*api.Execution, error) {
	return api.StopExecution(c, flowName, executionName)
}

// ListExecutions - list all executions of a flow
//
// PARAMS:
//   - flowName: which flow to list executions
//
// RETURNS:
//   - *api.ListExecutionsResult: information about the executions
//   - error: nil if success otherwise the specific error
func (c *Client) ListExecutions(flowName string) (*api.ListExecutionsResult, error) {
	return api.ListExecutions(c, flowName)
}

// DescribeExecution - describe detail info of an execution
//
// PARAMS:
//   - flowName: which flow the execution belongs to
//   - executionName: which execution to get
//
// RETURNS:
//   - *api.Execution: information about the execution
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeExecution(flowName, executionName string) (*api.Execution, error) {
	return api.DescribeExecution(c, flowName, executionName)
}

// GetExecutionHistory - get all state event detail info of an execution
//
// PARAMS:
//   - args: args to get history
//
// RETURNS:
//   - *api.GetExecutionHistoryResult: state history of the execution
//   - error: nil if success otherwise the specific error
func (c *Client) GetExecutionHistory(args *api.GetExecutionHistoryArgs) (*api.GetExecutionHistoryResult, error) {
	return api.GetExecutionHistory(c, args)
}

// ListFlow - list all flows
//
// RETURNS:
//   - *api.ListFlowResult: list of all flows
//   - error: nil if success otherwise the specific error
func (c *Client) ListFlow() (*api.ListFlowResult, error) {
	return api.ListFlow(c)
}

// CreateFlow - create a flow
//
// PARAMS:
//   - args: args to create a flow
//
// RETURNS:
//   - *api.Flow: the created flow information
//   - error: nil if success otherwise the specific error
func (c *Client) CreateFlow(args *api.CreateUpdateFlowArgs) (*api.Flow, error) {
	return api.CreateFlow(c, args)
}

// UpdateFlow - update a flow
//
// PARAMS:
//   - args: args to update a flow, args.Name is the name of flow to update
//
// RETURNS:
//   - *api.Flow: flow information after update
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateFlow(args *api.CreateUpdateFlowArgs) (*api.Flow, error) {
	return api.UpdateFlow(c, args)
}

// DescribeFlow - get information of a flow
//
// PARAMS:
//   - flowName: name of the flow to describe
//
// RETURNS:
//   - *api.Flow: flow information
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeFlow(flowName string) (*api.Flow, error) {
	return api.DescribeFlow(c, flowName)
}

// DeleteFlow - delete a flow
//
// PARAMS:
//   - flowName: name of the flow to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteFlow(flowName string) error {
	return api.DeleteFlow(c, flowName)
}

// PublishLayer - publish a layer
//
// PARAMS:
//   - layerName: name of the layer to publish
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) PublishLayer(args *api.PublishLayerVersionInput) (*api.PublishLayerVersionOutput, error) {
	return api.PublishLayerVersion(c, args)
}

// GetLayerVersion - get a layer version
//
// PARAMS:
//   - args: args to get a layer version
//
// RETURNS:
//   - *api.GetLayerVersionOutput: layer version information
//   - error: nil if success otherwise the specific error
func (c *Client) GetLayerVersion(args *api.GetLayerVersionArgs) (*api.GetLayerVersionOutput, error) {
	return api.GetLayerVersion(c, args)
}

// GetLayerVersionByBrn - get a layer version by brn
//
// PARAMS:
//   - args: args to get a layer version by brn
//
// RETURNS:
//   - *api.GetLayerVersionOutput: layer version information
func (c *Client) GetLayerVersionByBrn(args *api.GetLayerVersionArgs) (*api.GetLayerVersionOutput, error) {
	return api.GetLayerVersionByBrn(c, args)
}

// ListLayers - list all layers
//
// PARAMS:
//   - args: args to list all layers
//
// RETURNS:
//   - *api.ListLayersOutput: list of all layers
//   - error: nil if success otherwise the specific error
func (c *Client) ListLayers(args *api.ListLayerInput) (*api.ListLayersOutput, error) {
	return api.ListLayers(c, args)
}

// ListLayers - list all layers
//
// PARAMS:
//   - args: args to list all layers
//
// RETURNS:
//   - *api.ListLayerVersionsOutput: list of all layers
//   - error: nil if success otherwise the specific error
func (c *Client) ListLayerVersions(args *api.ListLayerVersionsInput) (*api.ListLayerVersionsOutput, error) {
	return api.ListLayerVersions(c, args)
}

// DeleteLayerVersion - delete a layer version
//
// PARAMS:
//   - args: args to delete a layer version
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteLayerVersion(args *api.DeleteLayerVersionArgs) error {
	return api.DeleteLayerVersion(c, args)
}

// DeleteLayer - delete a layer
//
// PARAMS:
//   - args: args to delete a layer
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteLayer(args *api.DeleteLayerArgs) error {
	return api.DeleteLayer(c, args)
}

// ListServices - list all services
//
// RETURNS:
//   - *api.ListServicesResult: result containing slice of services
//   - error: nil if success otherwise the specific error
func (c *Client) ListServices() (*api.ListServicesResult, error) {
	return api.ListServices(c)
}

// CreateService - create a service
//
// PARAMS:
//   - args: args to create a service
//
// RETURNS:
//   - *api.Service: the created service information
//   - error: nil if success otherwise the specific error
func (c *Client) CreateService(args *api.CreateServiceArgs) (*api.CreateServiceResult, error) {
	return api.CreateService(c, args)
}

// UpdateService - update a service
//
// PARAMS:
//   - args: args to update a service
//
// RETURNS:
//   - *api.Service: the updated service information
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateService(args *api.UpdateServiceArgs) (*api.UpdateServiceResult, error) {
	return api.UpdateService(c, args)
}

// DeleteService - delete a service
//
// PARAMS:
//   - DeleteServiceArgs: args of the service to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteService(args *api.DeleteServiceArgs) error {
	return api.DeleteService(c, args)
}

// GetService - get a service
//
// PARAMS:
//   - GetServiceArgs: args of the service to get
//
// RETURNS:
//   - *api.Service: the service information
//   - error: nil if success otherwise the specific error
func (c *Client) GetService(args *api.GetServiceArgs) (*api.GetServiceResult, error) {
	return api.GetService(c, args)
}
