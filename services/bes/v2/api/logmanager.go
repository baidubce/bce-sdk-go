package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// UpdateLogSettings toggles a cluster's slow-log/audit-log switches.
func UpdateLogSettings(cli bce.Client, region string, request *UpdateLogSettingsRequest) (*UpdateLogSettingsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update log settings request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update log settings request clusterId should not be empty")
	}

	uri := URI_LOG_VIEW + "/settings"

	result := &UpdateLogSettingsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SearchLog searches a cluster's logs by time range and type.
func SearchLog(cli bce.Client, region string, request *SearchLogRequest) (*SearchLogResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("search log request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("search log request clusterId should not be empty")
	}
	if request.PageNo == 0 {
		return nil, errors.New("search log request pageNo should not be empty")
	}
	if request.PageSize == 0 {
		return nil, errors.New("search log request pageSize should not be empty")
	}
	if request.StartTime == "" {
		return nil, errors.New("search log request startTime should not be empty")
	}
	if request.EndTime == "" {
		return nil, errors.New("search log request endTime should not be empty")
	}
	if request.Type == "" {
		return nil, errors.New("search log request type should not be empty")
	}

	uri := URI_LOG_VIEW + "/search"

	result := &SearchLogResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateLogExportTask creates a log export task for one or more log types.
func CreateLogExportTask(cli bce.Client, region string, request *CreateLogExportTaskRequest) (*CreateLogExportTaskResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		request = &CreateLogExportTaskRequest{}
	}

	uri := URI_LOG_VIEW + "/export"

	result := &CreateLogExportTaskResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLogExportRecord gets the progress and configuration of a log export task.
func GetLogExportRecord(cli bce.Client, region string, request *GetLogExportRecordRequest) (*GetLogExportRecordResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		request = &GetLogExportRecordRequest{}
	}

	uri := URI_LOG_VIEW + "/export_record"

	result := &GetLogExportRecordResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
