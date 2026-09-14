package api

import (
	"errors"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_LOGS          = "logs"
	URI_LOG_SWITCH    = "switch"
	URI_LOG_USAGE     = "usage"
	URI_LOG_COLLECTOR = "collector"
	URI_LOG_STATUS    = "status"
	URI_LOG_SEARCH    = "search"
)

// GetLogSwitch queries the log collection switch state of a cluster.
func GetLogSwitch(cli bce.Client, region string, request *GetLogSwitchRequest) (*LogSwitchResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get log switch request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get log switch request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_LOGS, URI_LOG_SWITCH)
	result := &LogSwitchResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateLogSwitch updates the log collection switch state of a cluster.
func UpdateLogSwitch(cli bce.Client, region string, request *UpdateLogSwitchRequest) (*LogSwitchResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update log switch request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update log switch request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_LOGS, URI_LOG_SWITCH)
	result := &LogSwitchResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLogUsage queries aggregated log index usage for a cluster.
func GetLogUsage(cli bce.Client, region string, request *GetLogUsageRequest) (*LogUsageResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get log usage request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get log usage request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_LOGS, URI_LOG_USAGE)
	result := &LogUsageResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLogCollectorStatus queries the log collector deployment and operation state.
func GetLogCollectorStatus(cli bce.Client, region string, request *GetLogCollectorStatusRequest) (*LogCollectorStatusResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get log collector status request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get log collector status request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_LOGS, URI_LOG_COLLECTOR, URI_LOG_STATUS)
	result := &LogCollectorStatusResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SearchLogs searches raw log records of a cluster.
func SearchLogs(cli bce.Client, region string, request *SearchLogsRequest) (*SearchLogsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("search logs request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("search logs request clusterId should not be empty")
	}
	if request.LogType == "" {
		return nil, errors.New("search logs request logType should not be empty")
	}
	if request.StartTime == "" {
		return nil, errors.New("search logs request startTime should not be empty")
	}
	if request.EndTime == "" {
		return nil, errors.New("search logs request endTime should not be empty")
	}
	if request.Page < 0 || request.Page > 50 {
		return nil, errors.New("search logs request page should be between 1 and 50 when set")
	}
	if request.PageSize < 0 || request.PageSize > 200 {
		return nil, errors.New("search logs request pageSize should be between 1 and 200 when set")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_LOGS, URI_LOG_SEARCH)
	result := &SearchLogsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
