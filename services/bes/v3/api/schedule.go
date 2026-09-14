package api

import (
	"errors"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// ListSchedules lists the scheduled tasks of a cluster.
func ListSchedules(cli bce.Client, region string, request *ListSchedulesRequest) (*ListSchedulesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list schedules request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list schedules request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_SCHEDULES)
	query := map[string]string{}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}

	result := &ListSchedulesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSchedule creates a scheduled task on a cluster.
func CreateSchedule(cli bce.Client, region string, request *CreateScheduleRequest) (*CreateScheduleResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create schedule request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create schedule request clusterId should not be empty")
	}
	if request.ScheduleName == "" {
		return nil, errors.New("create schedule request scheduleName should not be empty")
	}
	if request.ScheduleType == "" {
		return nil, errors.New("create schedule request scheduleType should not be empty")
	}
	if request.CronExpr == "" {
		return nil, errors.New("create schedule request cronExpr should not be empty")
	}
	if request.Params == nil {
		return nil, errors.New("create schedule request params should not be nil")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_SCHEDULES)

	result := &CreateScheduleResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateSchedule updates a scheduled task on a cluster.
func UpdateSchedule(cli bce.Client, region string, request *UpdateScheduleRequest) (*UpdateScheduleResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update schedule request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update schedule request clusterId should not be empty")
	}
	if request.ScheduleId == "" {
		return nil, errors.New("update schedule request scheduleId should not be empty")
	}
	if request.CronExpr == "" {
		return nil, errors.New("update schedule request cronExpr should not be empty")
	}
	if request.Params == nil {
		return nil, errors.New("update schedule request params should not be nil")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_SCHEDULES, request.ScheduleId)

	result := &UpdateScheduleResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSchedule deletes a scheduled task from a cluster.
func DeleteSchedule(cli bce.Client, region string, request *DeleteScheduleRequest) (*DeleteScheduleResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete schedule request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete schedule request clusterId should not be empty")
	}
	if request.ScheduleId == "" {
		return nil, errors.New("delete schedule request scheduleId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_SCHEDULES, request.ScheduleId)

	result := &DeleteScheduleResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodDelete, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}
