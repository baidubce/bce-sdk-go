package api

import (
	"encoding/json"
	"fmt"
)

type Validator interface {
	Validate() error
}

func (args InvocationsArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if args.Qualifier != "" && !validateQualifier(args.Qualifier) {
		return fmt.Errorf(QualifierInvalid)
	}
	return nil
}

func (args ListFunctionsArgs) Validate() error {
	if args.Marker < 0 || args.MaxItems < 0 {
		return fmt.Errorf(PaginateInvalid)
	}
	if args.FunctionVersion != "" && !validateVersion(args.FunctionVersion) {
		return fmt.Errorf(VersionInvalid)
	}
	return nil
}

func (args GetFunctionArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if args.Qualifier != "" && !validateQualifier(args.Qualifier) {
		return fmt.Errorf(QualifierInvalid)
	}
	return nil
}

func (args CreateFunctionArgs) Validate() error {
	if args.Code == nil {
		return fmt.Errorf(requiredIllegal, "Code")
	}
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if args.Handler == "" {
		return fmt.Errorf(requiredIllegal, "Handler")
	}
	if args.Runtime == "" {
		return fmt.Errorf(requiredIllegal, "Runtime")
	}
	if err := validateMemorySize(args.MemorySize); err != nil {
		return err
	}
	if len(args.Code.ZipFile) == 0 && (len(args.Code.BosBucket) == 0 || len(args.Code.BosObject) == 0) {
		return fmt.Errorf(FunctionCodeInvalid)
	}
	return nil
}

func (args DeleteFunctionArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if args.Qualifier != "" && !validateQualifier(args.Qualifier) {
		return fmt.Errorf(QualifierInvalid)
	}
	return nil
}

func (args UpdateFunctionCodeArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	return nil
}

func (args GetFunctionConfigurationArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if args.Qualifier != "" && !validateQualifier(args.Qualifier) {
		return fmt.Errorf(QualifierInvalid)
	}
	return nil
}

func (args UpdateFunctionConfigurationArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	return nil
}

func (args ListVersionsByFunctionArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if args.Marker < 0 || args.MaxItems < 0 {
		return fmt.Errorf(PaginateInvalid)
	}
	return nil
}

func (args PublishVersionArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	return nil
}

func (args ListAliasesArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if args.Marker < 0 || args.MaxItems < 0 {
		return fmt.Errorf(PaginateInvalid)
	}
	if args.FunctionVersion != "" && !validateVersion(args.FunctionVersion) {
		return fmt.Errorf(VersionInvalid)
	}
	return nil
}

func (args CreateAliasArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if args.FunctionVersion == "" {
		return fmt.Errorf(requiredIllegal, "FunctionVersion")
	}
	if args.Name == "" {
		return fmt.Errorf(requiredIllegal, "Name")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if !validateVersion(args.FunctionVersion) {
		return fmt.Errorf(VersionInvalid)
	}
	if !validateAliasName(args.Name) {
		return fmt.Errorf(AliasNameInvalid, args.FunctionName)
	}
	return nil
}

func (args GetAliasArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if args.AliasName == "" {
		return fmt.Errorf(requiredIllegal, "AliasName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if !validateAliasName(args.AliasName) {
		return fmt.Errorf(AliasNameInvalid, args.FunctionName)
	}
	return nil
}

func (args UpdateAliasArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if args.AliasName == "" {
		return fmt.Errorf(requiredIllegal, "AliasName")
	}
	if args.FunctionVersion == "" {
		return fmt.Errorf(requiredIllegal, "FunctionVersion")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if !validateAliasName(args.AliasName) {
		return fmt.Errorf(AliasNameInvalid, args.FunctionName)
	}
	if !validateVersion(args.FunctionVersion) {
		return fmt.Errorf(VersionInvalid)
	}
	return nil
}

func (args DeleteAliasArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	if args.AliasName == "" {
		return fmt.Errorf(requiredIllegal, "AliasName")
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	if !validateAliasName(args.AliasName) {
		return fmt.Errorf(AliasNameInvalid, args.FunctionName)
	}
	return nil
}

func (args ListTriggersArgs) Validate() error {
	if args.FunctionBrn == "" {
		return fmt.Errorf(requiredIllegal, "FunctionBrn")
	}
	if !validateFunctionBRN(args.FunctionBrn) {
		return fmt.Errorf(functionBRNInvalid, args.FunctionBrn)
	}
	return nil
}

func (args CreateTriggerArgs) Validate() error {
	if args.Target == "" {
		return fmt.Errorf(requiredIllegal, "Target")
	}
	if args.Source == "" {
		return fmt.Errorf(requiredIllegal, "Source")
	}
	if !validateFunctionBRN(args.Target) {
		return fmt.Errorf(functionBRNInvalid, args.Target)
	}

	return nil
}

func (args UpdateTriggerArgs) Validate() error {
	if args.RelationId == "" {
		return fmt.Errorf(requiredIllegal, "RelationId")
	}
	if args.Target == "" {
		return fmt.Errorf(requiredIllegal, "Target")
	}
	if args.Source == "" {
		return fmt.Errorf(requiredIllegal, "Source")
	}
	if !validateFunctionBRN(args.Target) {
		return fmt.Errorf(functionBRNInvalid, args.Target)
	}

	return nil
}

func (args DeleteTriggerArgs) Validate() error {
	if args.RelationId == "" {
		return fmt.Errorf(requiredIllegal, "RelationId")
	}
	if args.Target == "" {
		return fmt.Errorf(requiredIllegal, "Target")
	}
	if args.Source == "" {
		return fmt.Errorf(requiredIllegal, "Source")
	}
	if !validateFunctionBRN(args.Target) {
		return fmt.Errorf(functionBRNInvalid, args.Target)
	}
	return nil
}

func (args ReservedConcurrentExecutionsArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}

	if args.ReservedConcurrentExecutions < 0 || args.ReservedConcurrentExecutions > 90 {
		return fmt.Errorf(requiredIllegal, "ReservedConcurrentExecutions")
	}

	return nil
}

func (args DeleteReservedConcurrentExecutionsArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}

	return nil
}

func (args ListEventSourceArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	// 先检查是否满足brn规则 再看是不是function name
	if !validateFunctionBRN(args.FunctionName) {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}

	if args.Marker < 0 || args.MaxItems < 0 {
		return fmt.Errorf(PaginateInvalid)
	}
	return nil
}
func (args GetEventSourceArgs) Validate() error {
	if args.UUID == "" {
		return fmt.Errorf(requiredIllegal, "UUID")
	}
	return nil
}

func (args CreateEventSourceArgs) Validate() error {
	if args.FunctionName == "" {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}
	// 先检查是否满足brn规则 再看是不是function name
	if !validateFunctionBRN(args.FunctionName) || !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(requiredIllegal, "FunctionName")
	}

	if args.Type == TypeEventSourceDatahubTopic || args.Type == TypeEventSourceBms {
		return nil
	} else {
		return fmt.Errorf(EventSourceTypeNotSupport, args.Type)
	}
}

func (args UpdateEventSourceArgs) Validate() error {
	if args.UUID == "" {
		return fmt.Errorf(requiredIllegal, "UUID")
	}
	if args.FuncEventSource.Type != TypeEventSourceDatahubTopic && args.FuncEventSource.Type != TypeEventSourceBms {
		return fmt.Errorf(EventSourceTypeNotSupport, args.FuncEventSource.Type)
	}
	return nil
}

func (args DeleteEventSourceArgs) Validate() error {
	if args.UUID == "" {
		return fmt.Errorf(requiredIllegal, "UUID")
	}
	return nil
}

func (args GetExecutionHistoryArgs) Validate() error {
	if args.Limit <= 0 {
		return fmt.Errorf(getExecutionHistoryLimitIllegal)
	}
	if !validateFlowName(args.FlowName) {
		return fmt.Errorf(flowNameInvalid, args.FlowName)
	}
	if !validateExecutionName(args.ExecutionName) {
		return fmt.Errorf(executionNameInvalid, args.ExecutionName)
	}
	return nil
}

func (args CreateUpdateFlowArgs) Validate() error {
	if !validateFlowName(args.Name) {
		return fmt.Errorf(flowNameInvalid, args.Name)
	}
	if args.Type != "" && args.Type != "FDL" {
		return fmt.Errorf(flowTypeInvalid, args.Type)
	}
	return nil
}

func (args StartExecutionArgs) Validate() error {
	if !validateFlowName(args.FlowName) {
		return fmt.Errorf(flowNameInvalid, args.FlowName)
	}
	if args.ExecutionName != "" {
		if !validateExecutionName(args.ExecutionName) {
			return fmt.Errorf(executionNameInvalid, args.ExecutionName)
		}
	}

	var input = map[string]interface{}{}
	err := json.Unmarshal([]byte(args.Input), &input)
	if err != nil {
		return fmt.Errorf(executionInputInvalid, err.Error())
	}
	return nil
}

func (args CreateFunctionByBlueprintArgs) Validate() error {
	if args.BlueprintID == "" {
		return fmt.Errorf(requiredIllegal, args.BlueprintID)
	}
	if !validateFunctionName(args.FunctionName) {
		return fmt.Errorf(functionNameInvalid, args.FunctionName)
	}
	return nil
}

func (args PublishLayerVersionInput) Validate() error {
	if args.LayerName == "" {
		return fmt.Errorf(requiredIllegal, "LayerName")
	}
	if !validateLayerName(args.LayerName) {
		return fmt.Errorf(layerNameInvalid, args.LayerName)
	}
	if len(args.CompatibleRuntimes) == 0 {
		return fmt.Errorf(layerRuntimesInvalid)
	}
	if args.Content == nil {
		return fmt.Errorf(layerContentInvalid)
	}
	// Check if Content has valid data
	if args.Content != nil {
		hasZipFile := len(args.Content.ZipFile) > 0
		hasBosFile := args.Content.BosBucket != "" && args.Content.BosObject != ""
		if !hasZipFile && !hasBosFile {
			return fmt.Errorf(layerContentInvalid)
		}
	}
	return nil
}

// GetLayerVersionArgs validation
func (args GetLayerVersionArgs) Validate() error {
	if args.LayerName == "" {
		return fmt.Errorf(requiredIllegal, "LayerName")
	}
	if !validateLayerName(args.LayerName) {
		return fmt.Errorf(layerNameInvalid, args.LayerName)
	}
	if args.VersionNumber == "" {
		return fmt.Errorf(requiredIllegal, "VersionNumber")
	}
	if !validateVersionNumber(args.VersionNumber) {
		return fmt.Errorf(layerVersionInvalid, args.VersionNumber)
	}
	return nil
}

// ListLayerVersionsInput validation
func (args ListLayerInput) Validate() error {
	return nil
}

func (args ListLayerVersionsInput) Validate() error {
	if args.LayerName != "" && !validateLayerName(args.LayerName) {
		return fmt.Errorf(layerNameInvalid, args.LayerName)
	}
	return nil
}

// DeleteLayerVersionArgs validation
func (args DeleteLayerVersionArgs) Validate() error {
	if args.LayerName == "" {
		return fmt.Errorf(requiredIllegal, "LayerName")
	}
	if !validateLayerName(args.LayerName) {
		return fmt.Errorf(layerNameInvalid, args.LayerName)
	}
	if args.VersionNumber == "" {
		return fmt.Errorf(requiredIllegal, "VersionNumber")
	}
	if !validateVersionNumber(args.VersionNumber) {
		return fmt.Errorf(layerVersionInvalid, args.VersionNumber)
	}
	return nil
}

// DeleteLayerArgs validation
func (args DeleteLayerArgs) Validate() error {
	if args.LayerName == "" {
		return fmt.Errorf(requiredIllegal, "LayerName")
	}
	if !validateLayerName(args.LayerName) {
		return fmt.Errorf(layerNameInvalid, args.LayerName)
	}
	return nil
}

// Service validation
func (args CreateServiceArgs) Validate() error {
	if args.ServiceName == "" {
		return fmt.Errorf(requiredIllegal, "ServiceName")
	}
	if !validateServiceName(args.ServiceName) {
		return fmt.Errorf(serviceNameInvalid, args.ServiceName)
	}
	return nil
}

func (args DeleteServiceArgs) Validate() error {
	if args.ServiceName == "" {
		return fmt.Errorf(requiredIllegal, "ServiceName")
	}
	return nil
}

func (args UpdateServiceArgs) Validate() error {
	if args.ServiceName == "" {
		return fmt.Errorf(requiredIllegal, "ServiceName")
	}
	return nil
}

func (args GetServiceArgs) Validate() error {
	if args.ServiceName == "" {
		return fmt.Errorf(requiredIllegal, "ServiceName")
	}
	return nil
}
