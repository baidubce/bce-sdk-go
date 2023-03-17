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
	"fmt"
	"regexp"
)

const (
	RegularFunctionName = `^[a-zA-Z0-9-_:]+|\$LATEST$`
	RegularAliasName    = `^[a-zA-Z0-9-_]+$`
	RegularFunctionBRN  = `^(brn:(bce[a-zA-Z-]*):cfc:)([a-z]{2,5}[0-9]*:)([0-9a-z]{32}:)(function:)([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`
	RegularVersion      = `^\$LATEST|([0-9]+)$`

	RegularFlowName      = `^[a-zA-Z_]{1}[a-zA-Z0-9_-]{0,63}$`
	RegularExecutionName = `^[a-zA-Z_0-9]{1}[a-zA-Z0-9_-]{0,63}$`
	FlowType             = "FDL"

	MemoryBase     = 128
	minMemoryLimit = 128
	maxMemoryLimit = 3008
)

func getInvocationsUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s/invocations", functionName)
}

func getFunctionsUri() string {
	return fmt.Sprintf("/v1/functions")
}

func getFunctionUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s", functionName)
}

func getFunctionCodeUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s/code", functionName)
}

func getFunctionConfigurationUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s/configuration", functionName)
}

func getFunctionConCurrentUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s/concurrency", functionName)
}

func getFunctionVersionsUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s/versions", functionName)
}

func getFunctionAliasesUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s/aliases", functionName)
}

func getFunctionAliasUri(functionName string, aliasName string) string {
	return fmt.Sprintf("/v1/functions/%s/aliases/%s", functionName, aliasName)
}

func getEventSourceUri() string {
	return "/v1/event-source-mappings"
}

func getExecutionStartUri() string {
	return "/v1/execution/start"
}

func getExecutionStopUri() string {
	return "/v1/execution/stop"
}

func getExecutionsUri() string {
	return "/v1/executions"
}

func getExecutionUri() string {
	return "/v1/execution"
}

func getExecutionHistory() string {
	return "/v1/execution/history"
}

func getWorkflowsUri() string {
	return "/v1/flows"
}

func getWorkflowUri() string {
	return "/v1/flow"
}

func getOneEventSourceUri(uuid string) string {
	return fmt.Sprintf("/v1/event-source-mappings/%s", uuid)
}

func getTriggerUri() string {
	return fmt.Sprintf("/v1/console/relation")
}

func validateFunctionName(name string) bool {
	res, err := regexp.MatchString(RegularFunctionName, name)
	if err != nil {
		return false
	}
	return res
}

func validateFlowName(name string) bool {
	res, err := regexp.MatchString(RegularFlowName, name)
	if err != nil {
		return false
	}
	return res
}

func validateExecutionName(name string) bool {
	res, err := regexp.MatchString(RegularExecutionName, name)
	if err != nil {
		return false
	}
	return res
}

func validateAliasName(name string) bool {
	res, err := regexp.MatchString(RegularAliasName, name)
	if err != nil {
		return false
	}
	return res
}

func validateMemorySize(size int) error {
	if size%MemoryBase != 0 {
		return fmt.Errorf(memorySizeIllegal, size, MemoryBase)
	}
	if size > maxMemoryLimit {
		return fmt.Errorf(memoryRangeIllegal, size, minMemoryLimit, maxMemoryLimit)
	}
	if size < minMemoryLimit {
		return fmt.Errorf(memoryRangeIllegal, size, minMemoryLimit, maxMemoryLimit)
	}
	return nil
}

func validateFunctionBRN(brn string) bool {
	res, err := regexp.MatchString(RegularFunctionBRN, brn)
	if err != nil {
		return false
	}
	return res
}

func validateVersion(version string) bool {
	res, err := regexp.MatchString(RegularVersion, version)
	if err != nil {
		return false
	}
	return res
}

func validateQualifier(qualifier string) bool {
	var res bool
	res, err := regexp.MatchString(RegularAliasName, qualifier)
	if err == nil && res == true {
		return true
	}
	res, err = regexp.MatchString(RegularVersion, qualifier)
	if err != nil {
		return false
	}
	return res
}
