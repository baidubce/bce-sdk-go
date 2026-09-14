package api

// DefaultPluginRequest is the request shared by install/uninstall default plugin APIs.
type DefaultPluginRequest struct {
	ClusterId  string `json:"clusterId"`
	PluginName string `json:"pluginName"`
	ModuleType string `json:"moduleType"`
}

// DefaultPluginResponse is the response shared by install/uninstall default plugin APIs.
type DefaultPluginResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// CustomPluginRequest is the request shared by install/uninstall/delete custom plugin APIs.
type CustomPluginRequest struct {
	ClusterId  string `json:"clusterId"`
	PluginName string `json:"pluginName"`
	ModuleType string `json:"moduleType"`
}

// CustomPluginResponse is the response shared by install/uninstall custom plugin APIs.
type CustomPluginResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// DeleteCustomPluginResponse is the response of deleting a custom plugin.
type DeleteCustomPluginResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// UploadCustomPluginRequest is the request of uploading a custom plugin package.
// FilePath must point to a local .zip file, per the API doc's "只支持zip" constraint.
type UploadCustomPluginRequest struct {
	ClusterId  string
	ModuleType string
	FilePath   string
}

// UploadCustomPluginResponse is the response of uploading a custom plugin package.
type UploadCustomPluginResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// GetPluginInfoRequest is the request of getting a cluster's default and custom plugin lists.
type GetPluginInfoRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetPluginInfoResponse is the response of getting a cluster's default and custom plugin lists.
type GetPluginInfoResponse struct {
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Result  *PluginInfoResult `json:"result"`
}

// PluginInfoResult carries the custom and default plugin lists of a cluster.
type PluginInfoResult struct {
	Plugins        []PluginItem        `json:"plugins"`
	DefaultPlugins []DefaultPluginItem `json:"defaultPlugins"`
}

// PluginItem describes a single custom plugin.
type PluginItem struct {
	ModuleType   string `json:"moduleType"`
	PluginDesc   string `json:"pluginDesc"`
	PluginName   string `json:"pluginName"`
	PluginStatus string `json:"pluginStatus"`
}

// DefaultPluginItem describes a single default (system) plugin.
type DefaultPluginItem struct {
	ModuleType      string            `json:"moduleType"`
	PluginDesc      string            `json:"pluginDesc"`
	PluginName      string            `json:"pluginName"`
	PluginStatus    string            `json:"pluginStatus"`
	PluginOperation []PluginOperation `json:"pluginOperation,omitempty"`
}

// PluginOperation describes the operation currently available for a default plugin.
type PluginOperation struct {
	Enable    bool   `json:"enable"`
	Operation string `json:"operation"`
	Text      string `json:"text"`
}

// GetNLPDictRequest is the request of viewing a cluster's NLP dict configuration.
type GetNLPDictRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetNLPDictResponse is the response of viewing a cluster's NLP dict configuration.
type GetNLPDictResponse struct {
	Success bool           `json:"success"`
	Status  int            `json:"status"`
	Result  *NLPDictResult `json:"result"`
}

// NLPDictResult describes a cluster's currently configured NLP dict.
type NLPDictResult struct {
	OperationType string `json:"operationType"`
	Separator     string `json:"separator"`
	FileName      string `json:"fileName,omitempty"`
	DictContent   string `json:"dictContent,omitempty"`
}

// UpdateNLPDictRequest is the request of uploading an NLP dict file (upload_dict mode via a
// plain file upload, as opposed to the JSON-only bucket/text modes of the legacy endpoint).
type UpdateNLPDictRequest struct {
	ClusterId string
	Separator string
	FilePath  string
}

// NLPDictUpdateResult 是上传 nlp 词典响应里的 result 对象。
type NLPDictUpdateResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// UpdateNLPDictResponse is the response of uploading an NLP dict file.
type UpdateNLPDictResponse struct {
	Success bool                 `json:"success"`
	Status  int                  `json:"status"`
	Result  *NLPDictUpdateResult `json:"result"`
}
