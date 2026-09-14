package api

import (
	"errors"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// ListSystemPlugins lists the system plugins available to a BES cluster.
func ListSystemPlugins(cli bce.Client, region string, request *ListSystemPluginsRequest) (*ListSystemPluginsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list system plugins request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list system plugins request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_SYSTEM_PLUGINS)

	result := &ListSystemPluginsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateSystemPlugin installs or uninstalls system plugins on a BES cluster.
func UpdateSystemPlugin(cli bce.Client, region string, request *UpdateSystemPluginRequest) (*UpdateSystemPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update system plugin request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update system plugin request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_SYSTEM_PLUGINS)

	result := &UpdateSystemPluginResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListCustomPlugins lists the custom plugins installed on a BES cluster.
func ListCustomPlugins(cli bce.Client, region string, request *ListCustomPluginsRequest) (*ListCustomPluginsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list custom plugins request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list custom plugins request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_CUSTOM_PLUGINS)

	result := &ListCustomPluginsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCustomPlugin installs or uninstalls custom plugins on a BES cluster.
func UpdateCustomPlugin(cli bce.Client, region string, request *UpdateCustomPluginRequest) (*UpdateCustomPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update custom plugin request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update custom plugin request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_CUSTOM_PLUGINS)

	result := &UpdateCustomPluginResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCustomPluginVersion deletes a specific version of a custom plugin from a BES cluster.
func DeleteCustomPluginVersion(cli bce.Client, region string, request *DeleteCustomPluginVersionRequest) (*DeleteCustomPluginVersionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete custom plugin version request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete custom plugin version request clusterId should not be empty")
	}
	if request.PluginName == "" {
		return nil, errors.New("delete custom plugin version request pluginName should not be empty")
	}
	if request.PluginVersion == "" {
		return nil, errors.New("delete custom plugin version request pluginVersion should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_CUSTOM_PLUGINS, request.PluginName, URI_VERSIONS, request.PluginVersion)

	result := &DeleteCustomPluginVersionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodDelete, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UploadPluginFile uploads a plugin package or other plugin-management file to a BES cluster.
func UploadPluginFile(cli bce.Client, region string, request *UploadPluginFileRequest) (*UploadPluginFileResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("upload plugin file request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("upload plugin file request clusterId should not be empty")
	}
	if request.FilePath == "" {
		return nil, errors.New("upload plugin file request filePath should not be empty")
	}
	region = setRegion(region, request.Region)

	file, fileInfo, err := openUploadFile(request.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fields []multipartField
	if request.Type != "" {
		fields = append(fields, multipartField{name: "type", value: request.Type})
	}
	if request.Subtype != "" {
		fields = append(fields, multipartField{name: "subtype", value: request.Subtype})
	}

	body, contentType, err := buildMultipartBody(fields, "file", file, fileInfo.Name(), fileInfo.Size())
	if err != nil {
		return nil, err
	}

	uri := clusterURI(request.ClusterId, URI_FILES)

	result := &UploadPluginFileResponse{}
	if err := createMultipartRequest(cli, region, nethttp.MethodPost, uri, body, contentType, result); err != nil {
		return nil, err
	}
	return result, nil
}
