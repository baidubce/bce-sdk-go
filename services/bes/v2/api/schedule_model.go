package api

import "encoding/json"

// CreateScheduleRequest is the request of creating a scheduled task.
type CreateScheduleRequest struct {
	ClusterId    string                 `json:"clusterId"`
	Schedule     string                 `json:"schedule"`
	ScheduleName string                 `json:"scheduleName"`
	TaskType     string                 `json:"taskType"`
	Task         map[string]interface{} `json:"task"`
}

// ScheduleStringResultResponse is shared by schedule APIs whose response is
// {success: Boolean, status: Integer, result: String}.
type ScheduleStringResultResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// UpdateScheduleRequest is the request of updating a scheduled task.
type UpdateScheduleRequest struct {
	ClusterId    string                 `json:"clusterId"`
	Schedule     string                 `json:"schedule"`
	ScheduleName string                 `json:"scheduleName"`
	TaskType     string                 `json:"taskType"`
	Task         map[string]interface{} `json:"task"`
}

// ListSchedulesRequest is the request of listing a cluster's scheduled tasks.
type ListSchedulesRequest struct {
	ClusterId string `json:"clusterId"`
}

// ListSchedulesResponse is the response of listing a cluster's scheduled tasks.
type ListSchedulesResponse struct {
	Success bool                `json:"success"`
	Status  int                 `json:"status"`
	Result  *ScheduleListResult `json:"result"`
}

// ScheduleListResult carries a cluster's scheduled tasks.
type ScheduleListResult struct {
	Schedules []ScheduleItem `json:"schedules"`
}

// UnmarshalJSON tolerates the server answering result as an empty string instead of an object.
func (r *ScheduleListResult) UnmarshalJSON(data []byte) error {
	if isEmptyJSONResult(data) {
		return nil
	}
	type plain ScheduleListResult
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*r = ScheduleListResult(decoded)
	return nil
}

// ScheduleItem describes a single scheduled task.
type ScheduleItem struct {
	ClusterId         string                 `json:"clusterId"`
	Schedule          string                 `json:"schedule"`
	ScheduleName      string                 `json:"scheduleName"`
	TaskType          string                 `json:"taskType"`
	Task              map[string]interface{} `json:"task"`
	TaskRunTimes      int                    `json:"taskRunTimes"`
	TaskFailTimes     int                    `json:"taskFailTimes"`
	LastExecuteMillis int64                  `json:"lastExecuteMillis"`
	ModifiedMillis    int64                  `json:"modifiedMillis"`
	ScheduleId        string                 `json:"scheduleId"`
	ModifiedVersion   int                    `json:"modifiedVersion"`
	Status            string                 `json:"status"`
}

// DeleteScheduleRequest is the request of deleting a scheduled task.
type DeleteScheduleRequest struct {
	ClusterId    string `json:"clusterId"`
	ScheduleName string `json:"scheduleName"`
}
