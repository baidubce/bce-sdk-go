package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

func PublishLayerVersion(cli bce.Client, args *PublishLayerVersionInput) (*PublishLayerVersionOutput, error) {
	op := &Operation{
		HTTPUri:    layerVersionUri(args.LayerName),
		HTTPMethod: POST,
	}
	request := &cfcRequest{
		Args: args,
		Body: args,
	}
	result := &cfcResult{
		Result: &PublishLayerVersionOutput{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*PublishLayerVersionOutput); ok {
		return value, nil
	}
	return nil, nil
}

func GetLayerVersion(cli bce.Client, args *GetLayerVersionArgs) (*GetLayerVersionOutput, error) {
	op := &Operation{
		HTTPUri:    getLayerVersionUri(args.LayerName, args.VersionNumber),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"find": args.find,
			"Brn":  args.Brn,
		},
	}
	result := &cfcResult{
		Result: &GetLayerVersionOutput{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*GetLayerVersionOutput); ok {
		return value, nil
	}
	return nil, nil
}

func GetLayerVersionByBrn(cli bce.Client, args *GetLayerVersionArgs) (*GetLayerVersionOutput, error) {
	args.find = "LayerVersion"
	return GetLayerVersion(cli, args)
}

func ListLayers(cli bce.Client, args *ListLayerInput) (*ListLayersOutput, error) {
	op := &Operation{
		HTTPUri:    listLayersUri(),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"CompatibleRuntime": args.CompatibleRuntime,
			"Marker":            args.Marker,
			"MaxItems":          args.MaxItems,
			"PageNo":            args.PageNo,
			"PageSize":          args.PageSize,
		},
	}
	result := &cfcResult{
		Result: &ListLayersOutput{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListLayersOutput); ok {
		return value, nil
	}
	return nil, nil
}

func ListLayerVersions(cli bce.Client, args *ListLayerVersionsInput) (*ListLayerVersionsOutput, error) {
	op := &Operation{
		HTTPUri:    layerVersionUri(args.LayerName),
		HTTPMethod: GET,
	}
	request := &cfcRequest{
		Args: args,
		Params: map[string]interface{}{
			"CompatibleRuntime": args.CompatibleRuntime,
			"Marker":            args.Marker,
			"MaxItems":          args.MaxItems,
			"PageNo":            args.PageNo,
			"PageSize":          args.PageSize,
		},
	}
	result := &cfcResult{
		Result: &ListLayerVersionsOutput{},
	}
	err := caller(cli, op, request, result)
	if err != nil {
		return nil, err
	}
	if value, ok := result.Result.(*ListLayerVersionsOutput); ok {
		return value, nil
	}
	return nil, nil
}

func DeleteLayerVersion(cli bce.Client, args *DeleteLayerVersionArgs) error {
	op := &Operation{
		HTTPUri:    getLayerVersionUri(args.LayerName, args.VersionNumber),
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

func DeleteLayer(cli bce.Client, args *DeleteLayerArgs) error {
	op := &Operation{
		HTTPUri:    deleteLayerUri(args.LayerName),
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
