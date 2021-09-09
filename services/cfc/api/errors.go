package api

import "errors"

var (
	URIIllegal      = errors.New("invalid request uri")
	MethodIllegal   = errors.New("invalid request method")
	InvalidArgument = errors.New("invalid arguments")
	ParseJsonError  = errors.New("could not parse payload into json")
)

const (
	requiredIllegal     = "the %s field is required"
	memoryRangeIllegal  = "memory size %d must in %d MB ~ %d MB"
	memorySizeIllegal   = "memory size %d must be a multiple of %d MB"
	functionNameInvalid = "the function name of %s must match " + RegularFunctionName
	AliasNameInvalid    = "the alias name of %s must match " + RegularAliasName
	functionBRNInvalid  = "the brn %s must match " + RegularFunctionBRN
	FunctionCodeInvalid = "the code of function is invalidate"
	VersionInvalid      = "the version of function must match " + RegularVersion
	QualifierInvalid    = "the qualifier is not the function's version or alias"
	PaginateInvalid     = "the pagination must greater than 0"
    EventSourceTypeNotSupport = "the event source type: %s not support"
)
