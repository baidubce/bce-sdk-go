package api

// GetClusterConfigRequest is the request of viewing a cluster's extra configuration.
type GetClusterConfigRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetClusterConfigResponse is the response of viewing a cluster's extra configuration.
type GetClusterConfigResponse struct {
	Success bool               `json:"success"`
	Status  int                `json:"status"`
	Result  *ClusterConfigInfo `json:"result"`
}

// ClusterConfigInfo carries the extra ES/Kibana configuration of a cluster.
type ClusterConfigInfo struct {
	EsNodeExtraConfig string `json:"esNodeExtraConfig"`
	KibanaExtraConfig string `json:"kibanaExtraConfig"`
}

// UpdateClusterConfigRequest is the request of updating a cluster's extra configuration.
type UpdateClusterConfigRequest struct {
	ClusterId        string                 `json:"clusterId"`
	EsConfigsMap     map[string]interface{} `json:"esConfigsMap"`
	KibanaConfigsMap map[string]interface{} `json:"kibanaConfigsMap"`
}

// UpdateClusterConfigResponse is the response of updating a cluster's extra configuration.
type UpdateClusterConfigResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// ListSynonymDictsRequest is the request of listing a cluster's synonym dict files.
type ListSynonymDictsRequest struct {
	ClusterId string `json:"clusterId"`
}

// ListSynonymDictsResponse is the response of listing a cluster's synonym dict files.
type ListSynonymDictsResponse struct {
	Success bool          `json:"success"`
	Status  int           `json:"status"`
	Result  []SynonymDict `json:"result"`
}

// SynonymDict describes a single synonym dict file.
type SynonymDict struct {
	SynonymDictName string `json:"synonym_dict_name"`
}

// DeleteSynonymDictRequest is the request of deleting a cluster's synonym dict file.
type DeleteSynonymDictRequest struct {
	ClusterId       string `json:"clusterId"`
	SynonymDictName string `json:"synonymDictName"`
}

// DeleteSynonymDictResponse is the response of deleting a cluster's synonym dict file.
type DeleteSynonymDictResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// UploadSynonymDictRequest is the request of uploading a synonym dict file.
type UploadSynonymDictRequest struct {
	ClusterId string
	FilePath  string
}

// UploadSynonymDictResponse is the response of uploading a synonym dict file.
type UploadSynonymDictResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}
