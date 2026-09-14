package api

// PluginItem describes a single plugin to install or uninstall.
type PluginItem struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// SystemPluginVersionInfo describes a single version of a system plugin.
type SystemPluginVersionInfo struct {
	Version      string `json:"version,omitempty"`
	Changelog    string `json:"changelog,omitempty"`
	ChangelogUrl string `json:"changelogUrl,omitempty"`
	Status       string `json:"status,omitempty"`
}

// SystemPluginInfo describes a system plugin available to a cluster.
type SystemPluginInfo struct {
	PluginName     string                    `json:"pluginName,omitempty"`
	Engine         string                    `json:"engine,omitempty"`
	Description    string                    `json:"description,omitempty"`
	DescriptionUrl string                    `json:"descriptionUrl,omitempty"`
	Order          int                       `json:"order,omitempty"`
	VersionInfos   []SystemPluginVersionInfo `json:"versionInfos,omitempty"`
}

// ListSystemPluginsRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/system-plugins.
type ListSystemPluginsRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListSystemPluginsResponse contains the response body for
// GET /v3/clusters/{clusterId}/system-plugins.
type ListSystemPluginsResponse struct {
	DevPlugins []SystemPluginInfo `json:"devPlugins,omitempty"`
	PrePlugins []SystemPluginInfo `json:"prePlugins,omitempty"`
}

// UpdateSystemPluginRequest contains the request body for
// POST /v3/clusters/{clusterId}/system-plugins.
type UpdateSystemPluginRequest struct {
	Request
	ClusterId string       `json:"-"`
	Install   []PluginItem `json:"install,omitempty"`
	Uninstall []PluginItem `json:"uninstall,omitempty"`
}

// UpdateSystemPluginResponse contains the response body for
// POST /v3/clusters/{clusterId}/system-plugins.
type UpdateSystemPluginResponse struct {
	ClusterId   string `json:"clusterId,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	ActionId    string `json:"actionId,omitempty"`
}

// CustomPluginVersionInfo describes a single version of a custom plugin.
type CustomPluginVersionInfo struct {
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	UploadTime  int64  `json:"uploadTime,omitempty"`
}

// CustomPluginInfo describes a custom plugin installed on a cluster.
type CustomPluginInfo struct {
	PluginName   string                    `json:"pluginName,omitempty"`
	VersionInfos []CustomPluginVersionInfo `json:"versionInfos,omitempty"`
}

// ListCustomPluginsRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/custom-plugins.
type ListCustomPluginsRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListCustomPluginsResponse contains the response body for
// GET /v3/clusters/{clusterId}/custom-plugins.
type ListCustomPluginsResponse struct {
	CustomPlugins []CustomPluginInfo `json:"customPlugins,omitempty"`
}

// UpdateCustomPluginRequest contains the request body for
// POST /v3/clusters/{clusterId}/custom-plugins.
type UpdateCustomPluginRequest struct {
	Request
	ClusterId string       `json:"-"`
	Install   []PluginItem `json:"install,omitempty"`
	Uninstall []PluginItem `json:"uninstall,omitempty"`
}

// UpdateCustomPluginResponse contains the response body for
// POST /v3/clusters/{clusterId}/custom-plugins.
type UpdateCustomPluginResponse struct {
	ClusterId   string `json:"clusterId,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	ActionId    string `json:"actionId,omitempty"`
}

// DeleteCustomPluginVersionRequest contains the path parameters for
// DELETE /v3/clusters/{clusterId}/custom-plugins/{pluginName}/versions/{pluginVersion}.
type DeleteCustomPluginVersionRequest struct {
	Request
	ClusterId     string `json:"-"`
	PluginName    string `json:"-"`
	PluginVersion string `json:"-"`
}

// DeleteCustomPluginVersionResponse contains the response body for
// DELETE /v3/clusters/{clusterId}/custom-plugins/{pluginName}/versions/{pluginVersion}.
type DeleteCustomPluginVersionResponse struct {
	PluginName string `json:"pluginName,omitempty"`
	Version    string `json:"version,omitempty"`
}

// UploadPluginFileRequest contains the multipart/form-data fields for
// POST /v3/clusters/{clusterId}/files.
type UploadPluginFileRequest struct {
	Request
	ClusterId string `json:"-"`
	FilePath  string `json:"-"`
	Type      string `json:"-"`
	Subtype   string `json:"-"`
}

// UploadPluginFileResponse contains the response body for
// POST /v3/clusters/{clusterId}/files.
type UploadPluginFileResponse struct {
	Url  string `json:"url,omitempty"`
	Size int64  `json:"size,omitempty"`
}
