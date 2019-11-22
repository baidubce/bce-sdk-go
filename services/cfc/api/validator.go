package api

import (
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
	if len(args.Code.ZipFile) == 0 && (len(args.Code.BosBucket) == 0 || len(args.Code.BosObject) == 0){
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
