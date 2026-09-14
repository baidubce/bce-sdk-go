package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// StartInstance starts a single cluster instance.
func StartInstance(cli bce.Client, region string, request *InstanceOperationRequest) (*InstanceOperationResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("start instance request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("start instance request clusterId should not be empty")
	}
	if request.InstanceId == "" {
		return nil, errors.New("start instance request instanceId should not be empty")
	}

	uri := URI_INSTANCE + "/start"

	result := &InstanceOperationResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StopInstance stops a single cluster instance.
func StopInstance(cli bce.Client, region string, request *InstanceOperationRequest) (*InstanceOperationResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("stop instance request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("stop instance request clusterId should not be empty")
	}
	if request.InstanceId == "" {
		return nil, errors.New("stop instance request instanceId should not be empty")
	}

	uri := URI_INSTANCE + "/stop"

	result := &InstanceOperationResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// BatchStartInstances starts multiple cluster instances at once.
func BatchStartInstances(cli bce.Client, region string, request *BatchInstanceOperationRequest) (*BatchInstanceOperationResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("batch start instances request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("batch start instances request clusterId should not be empty")
	}
	if len(request.InstanceIdList) == 0 {
		return nil, errors.New("batch start instances request instanceIdList should not be empty")
	}

	uri := URI_INSTANCE + "/batchStart"

	result := &BatchInstanceOperationResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// BatchStopInstances stops multiple cluster instances at once.
func BatchStopInstances(cli bce.Client, region string, request *BatchInstanceOperationRequest) (*BatchInstanceOperationResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("batch stop instances request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("batch stop instances request clusterId should not be empty")
	}
	if len(request.InstanceIdList) == 0 {
		return nil, errors.New("batch stop instances request instanceIdList should not be empty")
	}

	uri := URI_INSTANCE + "/batchStop"

	result := &BatchInstanceOperationResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteInstances deletes (scales in) cluster instances.
func DeleteInstances(cli bce.Client, region string, request *DeleteInstancesRequest) (*DeleteInstancesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete instances request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete instances request clusterId should not be empty")
	}
	if len(request.InstanceIdList) == 0 {
		return nil, errors.New("delete instances request instanceIdList should not be empty")
	}
	if request.ModuleType == "" {
		return nil, errors.New("delete instances request moduleType should not be empty")
	}

	uri := URI_INSTANCE + "/delete"

	result := &DeleteInstancesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListScaleInInstances lists scale-in candidate instances for a cluster module.
func ListScaleInInstances(cli bce.Client, region string, request *ListScaleInInstancesRequest) (*ListScaleInInstancesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list scale-in instances request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list scale-in instances request clusterId should not be empty")
	}
	if request.ModuleType == "" {
		return nil, errors.New("list scale-in instances request moduleType should not be empty")
	}

	uri := URI_V2 + "/scalein/instances"

	result := &ListScaleInInstancesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ConfirmDataMigration confirms a pending data migration for the given instances.
func ConfirmDataMigration(cli bce.Client, region string, request *ConfirmDataMigrationRequest) (*ConfirmDataMigrationResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("confirm data migration request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("confirm data migration request clusterId should not be empty")
	}
	if request.ModuleType == "" {
		return nil, errors.New("confirm data migration request moduleType should not be empty")
	}
	if len(request.InstanceIdList) == 0 {
		return nil, errors.New("confirm data migration request instanceIdList should not be empty")
	}

	uri := URI_MIGRATE_V1 + "/confirm"

	result := &ConfirmDataMigrationResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RollbackDataMigration rolls back a previously confirmed data migration.
func RollbackDataMigration(cli bce.Client, region string, request *RollbackDataMigrationRequest) (*RollbackDataMigrationResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("rollback data migration request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("rollback data migration request clusterId should not be empty")
	}
	if request.RequestId == "" {
		return nil, errors.New("rollback data migration request requestId should not be empty")
	}

	uri := URI_MIGRATE_V1 + "/rollback"

	result := &RollbackDataMigrationResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListDataMigrationInstances lists data migration candidate instances for a cluster module.
func ListDataMigrationInstances(cli bce.Client, region string, request *ListDataMigrationInstancesRequest) (*ListDataMigrationInstancesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list data migration instances request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list data migration instances request clusterId should not be empty")
	}
	if request.ModuleType == "" {
		return nil, errors.New("list data migration instances request moduleType should not be empty")
	}

	uri := URI_MIGRATE_V1 + "/instances"

	result := &ListDataMigrationInstancesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SuggestDataMigrationInstances gets the system-suggested instances for a data migration plan.
func SuggestDataMigrationInstances(cli bce.Client, region string, request *SuggestDataMigrationInstancesRequest) (*SuggestDataMigrationInstancesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("suggest data migration instances request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("suggest data migration instances request clusterId should not be empty")
	}
	if request.ModuleType == "" {
		return nil, errors.New("suggest data migration instances request moduleType should not be empty")
	}
	if request.MigrateCount == 0 {
		return nil, errors.New("suggest data migration instances request migrateCount should not be empty")
	}

	uri := URI_MIGRATE_V1 + "/suggestInstances"

	result := &SuggestDataMigrationInstancesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
