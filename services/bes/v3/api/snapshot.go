package api

import (
	"errors"
	"strconv"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// ListIndicesByPattern lists snapshottable indices matching a pattern.
func ListIndicesByPattern(cli bce.Client, region string, request *ListIndicesByPatternRequest) (*ListIndicesByPatternResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list indices by pattern request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list indices by pattern request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_INDICES, URI_BACKUP_LIST_BY_PATTERN)

	result := &ListIndicesByPatternResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListRestores lists restore tasks of a snapshot.
func ListRestores(cli bce.Client, region string, request *ListRestoresRequest) (*ListRestoresResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list restores request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list restores request clusterId should not be empty")
	}
	if request.SnapshotId == "" {
		return nil, errors.New("list restores request snapshotId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOTS, request.SnapshotId, URI_BACKUP_RESTORES)

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
	if request.RestoreStatus != "" {
		query["restoreStatus"] = request.RestoreStatus
	}

	result := &ListRestoresResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListRestoreClusters lists clusters eligible as restore targets for a snapshot.
func ListRestoreClusters(cli bce.Client, region string, request *ListRestoreClustersRequest) (*ListRestoreClustersResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list restore clusters request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list restore clusters request clusterId should not be empty")
	}
	if request.SnapshotId == "" {
		return nil, errors.New("list restore clusters request snapshotId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOTS, request.SnapshotId, URI_BACKUP_RESTORE_CLUSTERS)

	query := map[string]string{}
	if request.IncludeSelf != nil {
		query["includeSelf"] = strconv.FormatBool(*request.IncludeSelf)
	}
	if request.Keyword != "" {
		query["keyword"] = request.Keyword
	}

	result := &ListRestoreClustersResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSnapshotRule queries the snapshot rule detail of a snapshot.
func GetSnapshotRule(cli bce.Client, region string, request *GetSnapshotRuleRequest) (*GetSnapshotRuleResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get snapshot rule request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get snapshot rule request clusterId should not be empty")
	}
	if request.SnapshotId == "" {
		return nil, errors.New("get snapshot rule request snapshotId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOTS, request.SnapshotId, URI_BACKUP_RULE)

	query := map[string]string{}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}
	if request.OrderBy != "" {
		query["orderBy"] = request.OrderBy
	}
	if request.Order != "" {
		query["order"] = request.Order
	}
	if request.IndexName != "" {
		query["indexName"] = request.IndexName
	}
	if request.RestoreIndexPattern != "" {
		query["restoreIndexPattern"] = request.RestoreIndexPattern
	}

	result := &GetSnapshotRuleResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListSnapshots lists snapshots of a cluster.
func ListSnapshots(cli bce.Client, region string, request *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list snapshots request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list snapshots request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOTS)

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
	if request.SnapshotName != "" {
		query["snapshotName"] = request.SnapshotName
	}
	if request.Type != "" {
		query["type"] = request.Type
	}
	if request.StartTimestamp > 0 {
		query["startTimestamp"] = strconv.FormatInt(request.StartTimestamp, 10)
	}
	if request.EndTimestamp > 0 {
		query["endTimestamp"] = strconv.FormatInt(request.EndTimestamp, 10)
	}
	if request.SnapshotStatus != "" {
		query["snapshotStatus"] = request.SnapshotStatus
	}
	if request.RestoreStatus != "" {
		query["restoreStatus"] = request.RestoreStatus
	}

	result := &ListSnapshotsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSnapshotConfig queries a snapshot config's detail.
func GetSnapshotConfig(cli bce.Client, region string, request *GetSnapshotConfigRequest) (*GetSnapshotConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get snapshot config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get snapshot config request clusterId should not be empty")
	}
	if request.ConfigId == "" {
		return nil, errors.New("get snapshot config request configId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOT_CONFIGS, request.ConfigId)

	result := &GetSnapshotConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSnapshot queries a snapshot's detail.
func GetSnapshot(cli bce.Client, region string, request *GetSnapshotRequest) (*GetSnapshotResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get snapshot request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get snapshot request clusterId should not be empty")
	}
	if request.SnapshotId == "" {
		return nil, errors.New("get snapshot request snapshotId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOTS, request.SnapshotId)

	result := &GetSnapshotResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListAutoSnapshotConfigs lists auto snapshot configs of a cluster.
func ListAutoSnapshotConfigs(cli bce.Client, region string, request *ListAutoSnapshotConfigsRequest) (*ListAutoSnapshotConfigsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list auto snapshot configs request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list auto snapshot configs request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_AUTO_SNAPSHOT_CFG)

	result := &ListAutoSnapshotConfigsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRestore creates a restore task from a snapshot.
func CreateRestore(cli bce.Client, region string, request *CreateRestoreRequest) (*CreateRestoreResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create restore request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create restore request clusterId should not be empty")
	}
	if request.SnapshotId == "" {
		return nil, errors.New("create restore request snapshotId should not be empty")
	}
	if request.DestClusterId == "" {
		return nil, errors.New("create restore request destClusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOTS, request.SnapshotId, URI_BACKUP_RESTORES)

	result := &CreateRestoreResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateManualSnapshot creates a manual snapshot.
func CreateManualSnapshot(cli bce.Client, region string, request *CreateManualSnapshotRequest) (*CreateManualSnapshotResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create manual snapshot request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create manual snapshot request clusterId should not be empty")
	}
	if request.Name == "" {
		return nil, errors.New("create manual snapshot request name should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_MANUAL_SNAPSHOT_CFG)

	result := &CreateManualSnapshotResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAutoSnapshotConfig creates an auto snapshot config.
func CreateAutoSnapshotConfig(cli bce.Client, region string, request *CreateAutoSnapshotConfigRequest) (*CreateAutoSnapshotConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create auto snapshot config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create auto snapshot config request clusterId should not be empty")
	}
	if request.Name == "" {
		return nil, errors.New("create auto snapshot config request name should not be empty")
	}
	if request.Period == "" {
		return nil, errors.New("create auto snapshot config request period should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_AUTO_SNAPSHOT_CFG)

	result := &CreateAutoSnapshotConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAutoSnapshotConfig updates an auto snapshot config.
func UpdateAutoSnapshotConfig(cli bce.Client, region string, request *UpdateAutoSnapshotConfigRequest) (*UpdateAutoSnapshotConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update auto snapshot config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update auto snapshot config request clusterId should not be empty")
	}
	if request.ConfigId == "" {
		return nil, errors.New("update auto snapshot config request configId should not be empty")
	}
	if request.Name == "" {
		return nil, errors.New("update auto snapshot config request name should not be empty")
	}
	if request.Period == "" {
		return nil, errors.New("update auto snapshot config request period should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_AUTO_SNAPSHOT_CFG, request.ConfigId)

	result := &UpdateAutoSnapshotConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSnapshot deletes a snapshot.
func DeleteSnapshot(cli bce.Client, region string, request *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete snapshot request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete snapshot request clusterId should not be empty")
	}
	if request.SnapshotId == "" {
		return nil, errors.New("delete snapshot request snapshotId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_SNAPSHOTS, request.SnapshotId)

	result := &DeleteSnapshotResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodDelete, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAutoSnapshotConfig deletes an auto snapshot config.
func DeleteAutoSnapshotConfig(cli bce.Client, region string, request *DeleteAutoSnapshotConfigRequest) (*DeleteAutoSnapshotConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete auto snapshot config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete auto snapshot config request clusterId should not be empty")
	}
	if request.ConfigId == "" {
		return nil, errors.New("delete auto snapshot config request configId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_AUTO_SNAPSHOT_CFG, request.ConfigId)

	result := &DeleteAutoSnapshotConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodDelete, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// BatchEnableAutoSnapshotConfigs batch enables or disables auto snapshot configs of a cluster.
func BatchEnableAutoSnapshotConfigs(cli bce.Client, region string, request *BatchEnableAutoSnapshotConfigsRequest) (*BatchEnableAutoSnapshotConfigsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("batch enable auto snapshot configs request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("batch enable auto snapshot configs request clusterId should not be empty")
	}
	if request.Enabled == nil {
		return nil, errors.New("batch enable auto snapshot configs request enabled should not be nil")
	}
	if len(request.ConfigIdList) == 0 {
		return nil, errors.New("batch enable auto snapshot configs request configIdList should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_BACKUP, URI_BACKUP_AUTO_SNAPSHOT_CFG)

	result := &BatchEnableAutoSnapshotConfigsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPatch, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
