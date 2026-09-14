package api

import (
	"errors"
	"net/url"
	"strconv"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	bcehttp "github.com/baidubce/bce-sdk-go/http"
)

func createJSONRequest(cli bce.Client, region, method, uri string, query map[string]string, payload interface{}, resp interface{}) error {
	return createJSONRequestWithHeaders(cli, region, method, uri, query, nil, payload, resp)
}

func createJSONRequestWithHeaders(cli bce.Client, region, method, uri string, query map[string]string,
	headers map[string]string, payload interface{}, resp interface{}) error {
	builder := bce.NewRequestBuilder(cli).
		WithMethod(method).
		WithURL(uri).
		WithHeader(bcehttp.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithHeader(HEADER_REGION, region)

	for key, value := range headers {
		if value != "" {
			builder = builder.WithHeader(key, value)
		}
	}
	for key, value := range query {
		if value != "" {
			builder = builder.WithQueryParamFilter(key, value)
		}
	}

	if payload != nil {
		builder = builder.WithBody(payload)
	}
	if resp != nil {
		builder = builder.WithResult(resp)
	}
	return builder.Do()
}

// credentialsOf returns the AK/SK configured on the client, used to encrypt password fields
// before they are sent. BES decrypts with the SK looked up by the AK sent in the
// X-Bce-Accesskey header, so both must be taken from the same credentials.
func credentialsOf(cli bce.Client) (*auth.BceCredentials, error) {
	conf := cli.GetBceClientConfig()
	if conf == nil || conf.Credentials == nil {
		return nil, errors.New("client credentials are not configured")
	}
	return conf.Credentials, nil
}

func setRegion(region string, requestRegion string) string {
	if requestRegion != "" {
		return requestRegion
	}
	return region
}

func intString(v int) string {
	if v <= 0 {
		return ""
	}
	return strconv.Itoa(v)
}

func clusterURI(parts ...string) string {
	if len(parts) == 0 {
		return URI_CLUSTERS
	}
	uri := URI_CLUSTERS
	for _, part := range parts {
		uri += "/" + url.PathEscape(part)
	}
	return uri
}

// CreateCluster creates a BES cluster.
func CreateCluster(cli bce.Client, region string, request *CreateClusterRequest) (*CreateClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create cluster request should not be nil")
	}
	if request.Payment == "" {
		return nil, errors.New("create cluster request payment should not be empty")
	}
	if request.Name == "" {
		return nil, errors.New("create cluster request name should not be empty")
	}
	if request.Engine == "" {
		return nil, errors.New("create cluster request engine should not be empty")
	}
	if request.EngineType == "" {
		return nil, errors.New("create cluster request engine type should not be empty")
	}
	if request.Version == "" {
		return nil, errors.New("create cluster request version should not be empty")
	}
	if request.AdminPassword == "" {
		return nil, errors.New("create cluster request admin password should not be empty")
	}
	if request.EnableHttps == nil {
		return nil, errors.New("create cluster request enable https should not be nil")
	}
	if len(request.LogicalZones) == 0 {
		return nil, errors.New("create cluster request logical zones should not be empty")
	}
	if request.VpcId == "" {
		return nil, errors.New("create cluster request vpc id should not be empty")
	}
	if request.SubnetId == "" {
		return nil, errors.New("create cluster request subnet id should not be empty")
	}
	if len(request.NodeSpecs) == 0 {
		return nil, errors.New("create cluster request node specs should not be empty")
	}
	region = setRegion(region, request.Region)

	credentials, err := credentialsOf(cli)
	if err != nil {
		return nil, err
	}
	encryptedPassword, err := aes128EncryptWithFirst16Char(request.AdminPassword, credentials.SecretAccessKey)
	if err != nil {
		return nil, err
	}
	// Send a copy with the encrypted password so the caller's request is left untouched.
	body := *request
	body.AdminPassword = encryptedPassword
	headers := map[string]string{HEADER_X_BCE_ACCESSKEY: credentials.AccessKeyId}

	result := &CreateClusterResponse{}
	if err := createJSONRequestWithHeaders(cli, region, nethttp.MethodPost, URI_CLUSTERS, nil, headers, &body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListClusters lists BES clusters.
func ListClusters(cli bce.Client, region string, request *ListClustersRequest) (*ListClustersResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	region = setRegion(region, request.Region)

	query := map[string]string{}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}
	if request.Order != "" {
		query["order"] = request.Order
	}
	if request.OrderBy != "" {
		query["orderBy"] = request.OrderBy
	}
	if request.Payment != "" {
		query["payment"] = request.Payment
	}
	if request.Engine != "" {
		query["engine"] = request.Engine
	}
	if request.Status != "" {
		query["status"] = request.Status
	}
	if request.ClusterHealth != "" {
		query["clusterHealth"] = request.ClusterHealth
	}
	if request.LogicalZone != "" {
		query["logicalZone"] = request.LogicalZone
	}
	if request.ClusterName != "" {
		query["clusterName"] = request.ClusterName
	}
	if request.ClusterId != "" {
		query["clusterId"] = request.ClusterId
	}
	if request.TagKey != "" {
		query["tagKey"] = request.TagKey
	}
	if request.TagValue != "" {
		query["tagValue"] = request.TagValue
	}

	result := &ListClustersResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, URI_CLUSTERS, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCluster gets BES cluster detail.
func GetCluster(cli bce.Client, region string, request *GetClusterRequest) (*GetClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get cluster request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId)

	result := &GetClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListClusterNodes lists nodes for a BES cluster.
func ListClusterNodes(cli bce.Client, region string, request *ListClusterNodesRequest) (*ListClusterNodesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list cluster nodes request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list cluster nodes request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "nodes")

	result := &ListClusterNodesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAvailableSpecs queries available cluster specs.
func QueryAvailableSpecs(cli bce.Client, region string, request *QueryAvailableSpecsRequest) (*QueryAvailableSpecsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request != nil {
		region = setRegion(region, request.Region)
	}

	result := &QueryAvailableSpecsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, URI_CLUSTERS+"/specs", nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAvailableKernels queries available cluster kernels.
func QueryAvailableKernels(cli bce.Client, region string, request *QueryAvailableKernelsRequest) (*QueryAvailableKernelsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request != nil {
		region = setRegion(region, request.Region)
	}

	result := &QueryAvailableKernelsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, URI_CLUSTERS+"/kernels", nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DownloadClusterCert downloads the cluster certificate zip stream.
// The returned Body must be closed by the caller.
func DownloadClusterCert(cli bce.Client, region string, request *DownloadClusterCertRequest) (*DownloadClusterCertResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("download cluster cert request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("download cluster cert request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "certs")

	req := &bce.BceRequest{}
	req.SetMethod(nethttp.MethodGet)
	req.SetUri(uri)
	req.SetHeader("Accept", CONTENT_TYPE_ZIP)
	req.SetHeader(HEADER_REGION, region)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	result := &DownloadClusterCertResponse{
		Body:                resp.Body(),
		ContentType:         resp.Header(bcehttp.CONTENT_TYPE),
		ContentDisposition:  resp.Header(bcehttp.CONTENT_DISPOSITION),
		CacheControl:        resp.Header(bcehttp.CACHE_CONTROL),
		Pragma:              resp.Header(HEADER_PRAGMA),
		Expires:             resp.Header(bcehttp.EXPIRES),
		XContentTypeOptions: resp.Header(HEADER_X_CONTENT_TYPE_OPTIONS),
	}
	if cl := resp.Header(bcehttp.CONTENT_LENGTH); cl != "" {
		if n, err := strconv.ParseInt(cl, 10, 64); err == nil {
			result.ContentLength = n
		}
	}
	return result, nil
}

// UpdateClusterName updates a cluster name.
func UpdateClusterName(cli bce.Client, region string, request *UpdateClusterNameRequest) (*UpdateClusterNameResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster name request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster name request clusterId should not be empty")
	}
	if request.NewName == "" {
		return nil, errors.New("update cluster name request new name should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "name")

	result := &UpdateClusterNameResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterMaintenance updates the maintenance window.
func UpdateClusterMaintenance(cli bce.Client, region string, request *UpdateClusterMaintenanceRequest) (*UpdateClusterMaintenanceResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster maintenance request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster maintenance request clusterId should not be empty")
	}
	if request.MaintenanceTimeZone == "" {
		return nil, errors.New("update cluster maintenance request maintenance time zone should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "maintenance-duration")

	result := &UpdateClusterMaintenanceResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterDeletionProtection updates deletion protection.
func UpdateClusterDeletionProtection(cli bce.Client, region string, request *UpdateClusterDeletionProtectionRequest) (*UpdateClusterDeletionProtectionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster deletion protection request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster deletion protection request clusterId should not be empty")
	}
	if request.DeletionProtection == nil {
		return nil, errors.New("update cluster deletion protection request deletion protection should not be nil")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "deletion-protection")

	result := &UpdateClusterDeletionProtectionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RecoverCluster recovers a deleted cluster.
func RecoverCluster(cli bce.Client, region string, request *RecoverClusterRequest) (*RecoverClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("recover cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("recover cluster request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "recovery")

	result := &RecoverClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCluster deletes a BES cluster.
func DeleteCluster(cli bce.Client, region string, request *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete cluster request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId)
	query := map[string]string{}
	if request.Recycle != nil {
		query["recycle"] = strconv.FormatBool(*request.Recycle)
	}

	result := &DeleteClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodDelete, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ResizeCluster performs a partial resize of a BES cluster.
func ResizeCluster(cli bce.Client, region string, request *ResizeClusterRequest) (*ResizeClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("resize cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("resize cluster request clusterId should not be empty")
	}
	if request.Mode == "" {
		return nil, errors.New("resize cluster request mode should not be empty")
	}
	if len(request.NodeSpecs) > 0 && len(request.DeleteNodeTypes) > 0 {
		return nil, errors.New("resize cluster request nodeSpecs and deleteNodeTypes should not be set together")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId)

	result := &ResizeClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListSuggestedMigrationNodes lists suggested nodes for data migration.
func ListSuggestedMigrationNodes(cli bce.Client, region string, request *ListSuggestedMigrationNodesRequest) (*ListMigrationNodesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list suggested migration nodes request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list suggested migration nodes request clusterId should not be empty")
	}
	if request.Type == "" {
		return nil, errors.New("list suggested migration nodes request type should not be empty")
	}
	if request.MigrateCount <= 0 {
		return nil, errors.New("list suggested migration nodes request migrateCount should be positive")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "migrations", "suggested-nodes")
	query := map[string]string{}
	if request.Type != "" {
		query["type"] = request.Type
	}
	if request.MigrateCount > 0 {
		query["migrateCount"] = intString(request.MigrateCount)
	}

	result := &ListMigrationNodesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListMigratableNodes lists nodes that are eligible for data migration.
func ListMigratableNodes(cli bce.Client, region string, request *ListMigratableNodesRequest) (*ListMigrationNodesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list migratable nodes request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list migratable nodes request clusterId should not be empty")
	}
	if request.Type == "" {
		return nil, errors.New("list migratable nodes request type should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "migrations", "nodes")
	query := map[string]string{}
	if request.Type != "" {
		query["type"] = request.Type
	}

	result := &ListMigrationNodesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAvailableUpgrades queries the versions and kernels a cluster can upgrade to.
func QueryAvailableUpgrades(cli bce.Client, region string, request *QueryAvailableUpgradesRequest) (*QueryAvailableUpgradesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("query available upgrades request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("query available upgrades request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "upgrades", "versions")

	result := &QueryAvailableUpgradesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterPublicAccess switches the cluster public access.
func UpdateClusterPublicAccess(cli bce.Client, region string, request *UpdateClusterPublicAccessRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster public access request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster public access request clusterId should not be empty")
	}
	if request.Enabled == nil {
		return nil, errors.New("update cluster public access request enabled should not be nil")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "public-access")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterCerebro enables Cerebro for a cluster. The API only supports enabling,
// so request.Enabled must be true.
func UpdateClusterCerebro(cli bce.Client, region string, request *UpdateClusterCerebroRequest) (*UpdateClusterCerebroResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster cerebro request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster cerebro request clusterId should not be empty")
	}
	if request.Enabled == nil || !*request.Enabled {
		return nil, errors.New("update cluster cerebro request enabled must be true")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "cerebro")

	result := &UpdateClusterCerebroResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StartCluster resumes a suspended BES cluster.
func StartCluster(cli bce.Client, region string, request *StartClusterRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("start cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("start cluster request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "start")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StartClusterNodes starts the given nodes of a cluster.
func StartClusterNodes(cli bce.Client, region string, request *ClusterNodesRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("start cluster nodes request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("start cluster nodes request clusterId should not be empty")
	}
	if len(request.Nodes) == 0 {
		return nil, errors.New("start cluster nodes request nodes should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "nodes", "start")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// MigrateClusterNodeData migrates data off the given nodes.
func MigrateClusterNodeData(cli bce.Client, region string, request *MigrateClusterNodeDataRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("migrate cluster node data request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("migrate cluster node data request clusterId should not be empty")
	}
	if len(request.Nodes) == 0 {
		return nil, errors.New("migrate cluster node data request nodes should not be empty")
	}
	if request.Type == "" {
		return nil, errors.New("migrate cluster node data request type should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "migrations")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterAccessWhitelist sets the access whitelist for a cluster and returns the
// cluster's full current whitelist state.
func UpdateClusterAccessWhitelist(cli bce.Client, region string, request *UpdateClusterAccessWhitelistRequest) (*UpdateClusterAccessWhitelistResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster access whitelist request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster access whitelist request clusterId should not be empty")
	}
	if request.AccessType == "" {
		return nil, errors.New("update cluster access whitelist request accessType should not be empty")
	}
	if request.NetworkType == "" {
		return nil, errors.New("update cluster access whitelist request networkType should not be empty")
	}
	if len(request.IpWhitelist) == 0 {
		return nil, errors.New("update cluster access whitelist request ipWhitelist should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "access-white-ips")

	result := &UpdateClusterAccessWhitelistResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpgradeCluster checks or executes an engine/kernel upgrade for a cluster.
func UpgradeCluster(cli bce.Client, region string, request *UpgradeClusterRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("upgrade cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("upgrade cluster request clusterId should not be empty")
	}
	if request.UpgradeType == "" {
		return nil, errors.New("upgrade cluster request upgradeType should not be empty")
	}
	if request.Operation == "" {
		return nil, errors.New("upgrade cluster request operation should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "upgrades")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StopCluster suspends a running BES cluster.
func StopCluster(cli bce.Client, region string, request *StopClusterRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("stop cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("stop cluster request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "stop")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StopClusterNodes stops the given nodes of a cluster.
func StopClusterNodes(cli bce.Client, region string, request *ClusterNodesRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("stop cluster nodes request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("stop cluster nodes request clusterId should not be empty")
	}
	if len(request.Nodes) == 0 {
		return nil, errors.New("stop cluster nodes request nodes should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "nodes", "stop")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateClusterProtocol switches HTTPS on or off for a cluster.
func UpdateClusterProtocol(cli bce.Client, region string, request *UpdateClusterProtocolRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update cluster protocol request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update cluster protocol request clusterId should not be empty")
	}
	if request.EnableHttps == nil {
		return nil, errors.New("update cluster protocol request enableHttps should not be nil")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "protocol")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RestartCluster restarts a cluster or a subset of its nodes.
func RestartCluster(cli bce.Client, region string, request *RestartClusterRequest) (*ClusterActionResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("restart cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("restart cluster request clusterId should not be empty")
	}
	if request.RestartType == "" {
		return nil, errors.New("restart cluster request restartType should not be empty")
	}
	if request.Mode == "" {
		return nil, errors.New("restart cluster request mode should not be empty")
	}
	switch request.RestartType {
	case RESTART_TYPE_CLUSTER:
		if len(request.Nodes) != 0 {
			return nil, errors.New("restart cluster request nodes should be empty when restartType is CLUSTER")
		}
	case RESTART_TYPE_NODE:
		if len(request.Nodes) == 0 {
			return nil, errors.New("restart cluster request nodes should not be empty when restartType is NODE")
		}
	}
	switch request.Mode {
	case RESTART_MODE_BLUE_GREEN, RESTART_MODE_ROLLING:
		if request.BatchCount != nil && *request.BatchCount != 0 {
			return nil, errors.New("restart cluster request batchCount should be empty or 0 when mode is BLUE_GREEN or ROLLING")
		}
		if request.BatchUnit != "" {
			return nil, errors.New("restart cluster request batchUnit should be empty when mode is BLUE_GREEN or ROLLING")
		}
	case RESTART_MODE_FORCE:
		if request.BatchUnit == "" {
			return nil, errors.New("restart cluster request batchUnit should not be empty when mode is FORCE")
		}
		if request.BatchCount == nil {
			return nil, errors.New("restart cluster request batchCount should not be nil when mode is FORCE")
		}
	}
	if request.BatchUnit == BATCH_UNIT_PERCENT && request.BatchCount != nil {
		if *request.BatchCount <= 0 || *request.BatchCount > 100 {
			return nil, errors.New("restart cluster request batchCount should be within (0, 100] when batchUnit is PERCENT")
		}
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "restart")

	result := &ClusterActionResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
