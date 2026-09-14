package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// AuthorizeInspect authorizes intelligent inspection for a cluster.
func AuthorizeInspect(cli bce.Client, region string, request *AuthorizeInspectRequest) (*AuthorizeInspectResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("authorize inspect request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("authorize inspect request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/authorize"

	result := &AuthorizeInspectResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SwitchAutoInspect turns auto inspection on or off for a cluster.
func SwitchAutoInspect(cli bce.Client, region string, request *SwitchAutoInspectRequest) (*InspectStringSuccessCommonResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("switch auto inspect request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("switch auto inspect request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/switch_auto"

	result := &InspectStringSuccessCommonResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CheckAutoInspect checks whether auto inspection is enabled for a cluster.
func CheckAutoInspect(cli bce.Client, region string, request *CheckAutoInspectRequest) (*CheckAutoInspectResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("check auto inspect request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("check auto inspect request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/check_auto"

	result := &CheckAutoInspectResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateManualInspectTask submits a manual inspection task for a cluster.
func CreateManualInspectTask(cli bce.Client, region string, request *CreateManualInspectTaskRequest) (*InspectStringSuccessCommonResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create manual inspect task request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create manual inspect task request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/create_manual"

	result := &InspectStringSuccessCommonResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CheckInspectBusy checks whether a new inspection task can be submitted for a cluster.
func CheckInspectBusy(cli bce.Client, region string, request *CheckInspectBusyRequest) (*CheckInspectBusyResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("check inspect busy request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("check inspect busy request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/check_busy"

	result := &CheckInspectBusyResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetManualInspectCount gets today's completed manual inspection count for a cluster.
func GetManualInspectCount(cli bce.Client, region string, request *GetManualInspectCountRequest) (*GetManualInspectCountResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get manual inspect count request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get manual inspect count request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/manual_count"

	result := &GetManualInspectCountResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetManualInspectConfig views a cluster's manual inspection configuration.
func GetManualInspectConfig(cli bce.Client, region string, request *GetManualInspectConfigRequest) (*GetManualInspectConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get manual inspect config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get manual inspect config request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/get_manual_conf"

	result := &GetManualInspectConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateManualInspectConfig updates a cluster's manual inspection configuration.
func UpdateManualInspectConfig(cli bce.Client, region string, request *UpdateManualInspectConfigRequest) (*InspectStringSuccessCommonResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update manual inspect config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update manual inspect config request clusterId should not be empty")
	}
	if request.Indices == "" {
		return nil, errors.New("update manual inspect config request indices should not be empty")
	}

	uri := URI_INSPECT + "/update_manual_conf"

	result := &InspectStringSuccessCommonResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListInspectItems lists all selectable inspection items.
func ListInspectItems(cli bce.Client, region string) (*ListInspectItemsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}

	uri := URI_INSPECT + "/list_items"

	result := &ListInspectItemsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, struct{}{}, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInspectTask gets a single inspection task's execution status and result.
func GetInspectTask(cli bce.Client, region string, request *GetInspectTaskRequest) (*GetInspectTaskResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get inspect task request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get inspect task request clusterId should not be empty")
	}
	if request.TaskId == "" {
		return nil, errors.New("get inspect task request taskId should not be empty")
	}

	uri := URI_INSPECT + "/get_task"

	result := &GetInspectTaskResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListInspectTasks lists completed inspection tasks in the last 7 days.
func ListInspectTasks(cli bce.Client, region string, request *ListInspectTasksRequest) (*ListInspectTasksResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list inspect tasks request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list inspect tasks request clusterId should not be empty")
	}
	if request.PageNo == 0 {
		return nil, errors.New("list inspect tasks request pageNo should not be empty")
	}
	if request.PageSize == 0 {
		return nil, errors.New("list inspect tasks request pageSize should not be empty")
	}

	uri := URI_INSPECT + "/list_tasks"

	result := &ListInspectTasksResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLatestInspectOverview gets the latest inspection overview for a cluster.
func GetLatestInspectOverview(cli bce.Client, region string, request *GetLatestInspectOverviewRequest) (*GetLatestInspectOverviewResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get latest inspect overview request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get latest inspect overview request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/overview/latest"

	result := &GetLatestInspectOverviewResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetWeeklyInspectOverview gets the last 7 days' inspection overview for a cluster.
func GetWeeklyInspectOverview(cli bce.Client, region string, request *GetWeeklyInspectOverviewRequest) (*GetWeeklyInspectOverviewResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get weekly inspect overview request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get weekly inspect overview request clusterId should not be empty")
	}

	uri := URI_INSPECT + "/overview/weekly"

	result := &GetWeeklyInspectOverviewResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
