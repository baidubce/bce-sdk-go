package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// GetClusterConfig views a cluster's extra ES/Kibana configuration.
func GetClusterConfig(cli bce.Client, region string, request *GetClusterConfigRequest) (*GetClusterConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get cluster config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get cluster config request clusterId should not be empty")
	}

	uri := URI_PREFIX_V2 + "/config/info"

	result := &GetClusterConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterConfig performs a full update of a cluster's extra ES/Kibana configuration.
func UpdateClusterConfig(cli bce.Client, region string, request *UpdateClusterConfigRequest) (*UpdateClusterConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster config request clusterId should not be empty")
	}
	if request.EsConfigsMap == nil {
		return nil, errors.New("update cluster config request esConfigsMap should not be nil")
	}
	if request.KibanaConfigsMap == nil {
		return nil, errors.New("update cluster config request kibanaConfigsMap should not be nil")
	}

	uri := URI_PREFIX_V2 + "/config/update"

	result := &UpdateClusterConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListSynonymDicts lists a cluster's uploaded synonym dict files.
func ListSynonymDicts(cli bce.Client, region string, request *ListSynonymDictsRequest) (*ListSynonymDictsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list synonym dicts request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list synonym dicts request clusterId should not be empty")
	}

	uri := URI_PREFIX_V2 + "/synonym_dict/list"

	result := &ListSynonymDictsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSynonymDict deletes a cluster's synonym dict file.
func DeleteSynonymDict(cli bce.Client, region string, request *DeleteSynonymDictRequest) (*DeleteSynonymDictResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete synonym dict request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete synonym dict request clusterId should not be empty")
	}
	if request.SynonymDictName == "" {
		return nil, errors.New("delete synonym dict request synonymDictName should not be empty")
	}

	uri := URI_PREFIX_V2 + "/synonym_dict/delete"

	result := &DeleteSynonymDictResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UploadSynonymDict uploads a synonym dict file to a cluster. The request body is
// multipart/form-data carrying a single file field, matching the Java SDK's uploadSynonymDict.
func UploadSynonymDict(cli bce.Client, region string, request *UploadSynonymDictRequest) (*UploadSynonymDictResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("upload synonym dict request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("upload synonym dict request clusterId should not be empty")
	}
	if request.FilePath == "" {
		return nil, errors.New("upload synonym dict request filePath should not be empty")
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

	uri := URI_PREFIX_V2 + "/" + request.ClusterId + "/synonym_dict/upload"

	result := &UploadSynonymDictResponse{}
	if err := createMultipartRequest(cli, region, nethttp.MethodPost, uri, body, contentType, result); err != nil {
		return nil, err
	}
	return result, nil
}
