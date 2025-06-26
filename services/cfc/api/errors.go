package api

import "errors"

var (
	URIIllegal      = errors.New("invalid request uri")
	MethodIllegal   = errors.New("invalid request method")
	InvalidArgument = errors.New("invalid arguments")
	ParseJsonError  = errors.New("could not parse payload into json")
)

const (
	requiredIllegal           = "the %s field is required"
	memoryRangeIllegal        = "memory size %d must in %d MB ~ %d MB"
	memorySizeIllegal         = "memory size %d must be a multiple of %d MB"
	functionNameInvalid       = "the function name of %s must match " + RegularFunctionName
	AliasNameInvalid          = "the alias name of %s must match " + RegularAliasName
	functionBRNInvalid        = "the brn %s must match " + RegularFunctionBRN
	FunctionCodeInvalid       = "the code of function is invalidate"
	VersionInvalid            = "the version of function must match " + RegularVersion
	QualifierInvalid          = "the qualifier is not the function's version or alias"
	PaginateInvalid           = "the pagination must greater than 0"
	EventSourceTypeNotSupport = "the event source type: %s not support"

	// Layer related error constants
	layerNameInvalid     = "the layer name %s must match " + RegularLayerName
	layerVersionInvalid  = "the layer version number %s must be a positive integer"
	layerContentInvalid  = "the layer content is required and must contain either ZipFile or ZipFileBytes or BosBucket/BosObject"
	layerRuntimesInvalid = "the compatible runtimes field is required and cannot be empty"

	// Service related error constants
	serviceNameInvalid = "the service name %s must match " + RegularServiceName
)

const (
	getExecutionHistoryLimitIllegal = "the limit field should be greater than 0"
	flowNameInvalid                 = "the flow name %s must match " + RegularFlowName
	executionNameInvalid            = "the execution name %s must match " + RegularExecutionName
	flowTypeInvalid                 = "the flow type %s not supported, only support: " + FlowType
	executionInputInvalid           = "the execution input is not valid json, err: %s"
)
