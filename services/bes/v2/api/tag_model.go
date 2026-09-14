package api

// TagListResponse is the response of listing all tags across clusters.
type TagListResponse struct {
	Success bool  `json:"success"`
	Status  int   `json:"status"`
	Result  []Tag `json:"result"`
}

// UpdateClusterTagsRequest is the request of updating the tags of a single cluster.
type UpdateClusterTagsRequest struct {
	ClusterId string `json:"clusterId"`
	Tags      []Tag  `json:"tags"`
}

// UpdateClusterTagsResponse is the response of updating the tags of a single cluster.
type UpdateClusterTagsResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// BatchInsertTagsRequest is the request of batch inserting tags into multiple clusters.
type BatchInsertTagsRequest struct {
	ClusterIdList []string `json:"clusterIdList"`
	InsertTags    []Tag    `json:"insertTags"`
}

// BatchInsertTagsResponse is the response of batch inserting tags into multiple clusters.
type BatchInsertTagsResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}
