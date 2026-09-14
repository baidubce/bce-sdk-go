package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// InstallDefaultPlugin installs a system default plugin on a cluster module.
func InstallDefaultPlugin(cli bce.Client, region string, request *DefaultPluginRequest) (*DefaultPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if err := checkDefaultPluginRequest(request, "install default plugin"); err != nil {
		return nil, err
	}

	uri := URI_DEFAULT_PLUGIN + "/install"

	result := &DefaultPluginResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UninstallDefaultPlugin uninstalls a system default plugin from a cluster module.
func UninstallDefaultPlugin(cli bce.Client, region string, request *DefaultPluginRequest) (*DefaultPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if err := checkDefaultPluginRequest(request, "uninstall default plugin"); err != nil {
		return nil, err
	}

	uri := URI_DEFAULT_PLUGIN + "/uninstall"

	result := &DefaultPluginResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// InstallCustomPlugin installs a previously uploaded custom plugin on a cluster module.
func InstallCustomPlugin(cli bce.Client, region string, request *CustomPluginRequest) (*CustomPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if err := checkCustomPluginRequest(request, "install custom plugin"); err != nil {
		return nil, err
	}

	uri := URI_PLUGIN + "/install"

	result := &CustomPluginResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UninstallCustomPlugin uninstalls a custom plugin from a cluster module.
func UninstallCustomPlugin(cli bce.Client, region string, request *CustomPluginRequest) (*CustomPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if err := checkCustomPluginRequest(request, "uninstall custom plugin"); err != nil {
		return nil, err
	}

	uri := URI_PLUGIN + "/uninstall"

	result := &CustomPluginResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCustomPlugin deletes an uploaded custom plugin package.
func DeleteCustomPlugin(cli bce.Client, region string, request *CustomPluginRequest) (*DeleteCustomPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if err := checkCustomPluginRequest(request, "delete custom plugin"); err != nil {
		return nil, err
	}

	uri := URI_PLUGIN + "/delete"

	result := &DeleteCustomPluginResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UploadCustomPlugin uploads a custom plugin package (zip) to a cluster module. The request
// body is multipart/form-data carrying a single file field, matching the Java SDK's uploadFile.
func UploadCustomPlugin(cli bce.Client, region string, request *UploadCustomPluginRequest) (*UploadCustomPluginResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("upload custom plugin request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("upload custom plugin request clusterId should not be empty")
	}
	if request.ModuleType == "" {
		return nil, errors.New("upload custom plugin request moduleType should not be empty")
	}
	if request.FilePath == "" {
		return nil, errors.New("upload custom plugin request filePath should not be empty")
	}

	file, fileInfo, err := openUploadFile(request.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body, contentType, err := buildMultipartBody(nil, "file", file, fileInfo.Name(), fileInfo.Size())
	if err != nil {
		return nil, err
	}

	uri := URI_PLUGIN + "/upload/" + request.ClusterId + "/" + request.ModuleType

	result := &UploadCustomPluginResponse{}
	if err := createMultipartRequest(cli, region, nethttp.MethodPost, uri, body, contentType, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPluginInfo gets a cluster's default and custom plugin lists.
func GetPluginInfo(cli bce.Client, region string, request *GetPluginInfoRequest) (*GetPluginInfoResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get plugin info request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get plugin info request clusterId should not be empty")
	}

	uri := URI_PLUGIN + "/info"

	result := &GetPluginInfoResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetNLPDict views a cluster's NLP dict configuration.
func GetNLPDict(cli bce.Client, region string, request *GetNLPDictRequest) (*GetNLPDictResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get nlp dict request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get nlp dict request clusterId should not be empty")
	}

	uri := URI_PREFIX_V2 + "/nlp_dict/display"

	result := &GetNLPDictResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateNLPDict uploads an NLP dict file to a cluster (the upload_dict flow). The request body
// is multipart/form-data carrying clusterId, separator and a file field, matching the Java SDK's
// updateNlpDict.
func UpdateNLPDict(cli bce.Client, region string, request *UpdateNLPDictRequest) (*UpdateNLPDictResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update nlp dict request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update nlp dict request clusterId should not be empty")
	}
	if request.Separator == "" {
		return nil, errors.New("update nlp dict request separator should not be empty")
	}
	if request.FilePath == "" {
		return nil, errors.New("update nlp dict request filePath should not be empty")
	}

	file, fileInfo, err := openUploadFile(request.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fields := []multipartField{
		{name: "clusterId", value: request.ClusterId},
		{name: "separator", value: request.Separator},
	}
	body, contentType, err := buildMultipartBody(fields, "file", file, fileInfo.Name(), fileInfo.Size())
	if err != nil {
		return nil, err
	}

	uri := URI_V2 + "/nlp_dict/update"

	result := &UpdateNLPDictResponse{}
	if err := createMultipartRequest(cli, region, nethttp.MethodPost, uri, body, contentType, result); err != nil {
		return nil, err
	}
	return result, nil
}

func checkDefaultPluginRequest(request *DefaultPluginRequest, action string) error {
	if request == nil {
		return errors.New(action + " request should not be nil")
	}
	if request.ClusterId == "" {
		return errors.New(action + " request clusterId should not be empty")
	}
	if request.PluginName == "" {
		return errors.New(action + " request pluginName should not be empty")
	}
	if request.ModuleType == "" {
		return errors.New(action + " request moduleType should not be empty")
	}
	return nil
}

func checkCustomPluginRequest(request *CustomPluginRequest, action string) error {
	if request == nil {
		return errors.New(action + " request should not be nil")
	}
	if request.ClusterId == "" {
		return errors.New(action + " request clusterId should not be empty")
	}
	if request.PluginName == "" {
		return errors.New(action + " request pluginName should not be empty")
	}
	if request.ModuleType == "" {
		return errors.New(action + " request moduleType should not be empty")
	}
	return nil
}
