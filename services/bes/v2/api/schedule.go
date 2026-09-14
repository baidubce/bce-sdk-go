package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// CreateSchedule creates a scheduled task on a cluster.
func CreateSchedule(cli bce.Client, region string, request *CreateScheduleRequest) (*ScheduleStringResultResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create schedule request should not be nil")
	}
	if err := checkScheduleTaskRequest(request.ClusterId, request.Schedule, request.ScheduleName, request.TaskType, request.Task, "create schedule"); err != nil {
		return nil, err
	}

	uri := URI_SCHEDULE + "/create"

	result := &ScheduleStringResultResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateSchedule updates a scheduled task on a cluster.
func UpdateSchedule(cli bce.Client, region string, request *UpdateScheduleRequest) (*ScheduleStringResultResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update schedule request should not be nil")
	}
	if err := checkScheduleTaskRequest(request.ClusterId, request.Schedule, request.ScheduleName, request.TaskType, request.Task, "update schedule"); err != nil {
		return nil, err
	}

	uri := URI_SCHEDULE + "/update"

	result := &ScheduleStringResultResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListSchedules lists a cluster's scheduled tasks.
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

	uri := URI_SCHEDULE + "/list"

	result := &ListSchedulesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSchedule deletes a scheduled task from a cluster.
func DeleteSchedule(cli bce.Client, region string, request *DeleteScheduleRequest) (*ScheduleStringResultResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete schedule request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete schedule request clusterId should not be empty")
	}
	if request.ScheduleName == "" {
		return nil, errors.New("delete schedule request scheduleName should not be empty")
	}

	uri := URI_SCHEDULE + "/delete"

	result := &ScheduleStringResultResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

func checkScheduleTaskRequest(clusterId, schedule, scheduleName, taskType string, task map[string]interface{}, action string) error {
	if clusterId == "" {
		return errors.New(action + " request clusterId should not be empty")
	}
	if schedule == "" {
		return errors.New(action + " request schedule should not be empty")
	}
	if scheduleName == "" {
		return errors.New(action + " request scheduleName should not be empty")
	}
	if taskType == "" {
		return errors.New(action + " request taskType should not be empty")
	}
	if len(task) == 0 {
		return errors.New(action + " request task should not be empty")
	}
	return nil
}
