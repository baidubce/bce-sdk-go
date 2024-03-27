package bcm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bcm/model"
)

const (
	// GetMetricDataPath userId - scope - metricName
	GetMetricDataPath = "/json-api/v1/metricdata/%s/%s/%s"
	// BatchMetricDataPath userId - scope
	BatchMetricDataPath             = "/json-api/v1/metricdata/metricName/%s/%s"
	CreateNamespacePath             = "/csm/api/v1/userId/%s/custom/namespaces/create"
	BatchDeleteNamespacesPath       = "/csm/api/v1/userId/%s/custom/namespaces/delete"
	UpdateNamespacePath             = "/csm/api/v1/userId/%s/custom/namespaces/update"
	ListNamespacesPath              = "/csm/api/v1/userId/%s/custom/namespaces/list"
	CreateNamespaceMetricPath       = "/csm/api/v1/userId/%s/custom/namespaces/%s/metrics/create"
	BatchDeleteNamespaceMetricsPath = "/csm/api/v1/userId/%s/custom/namespaces/%s/metrics/delete"
	UpdateNamespaceMetricPath       = "/csm/api/v1/userId/%s/custom/namespaces/%s/metrics/%s"
	ListNamespaceMetricsPath        = "/csm/api/v1/userId/%s/custom/namespaces/metrics"
	GetCustomMetricPath             = "/csm/api/v1/userId/%s/custom/namespaces/%s/metrics/%s"
	CreateNamespaceEventPath        = "/csm/api/v1/custom/event/configs/create"
	BatchDeleteNamespaceEventsPath  = "/csm/api/v1/custom/event/configs/delete"
	UpdateNamespaceEventPath        = "/csm/api/v1/custom/event/configs/update"
	ListNamespaceEventsPath         = "/csm/api/v1/custom/event/configs/list"
	GetCustomEventPath              = "/csm/api/v1/custom/event/configs/detail"

	ApplicationInfoPath                 = "/csm/api/v1/userId/%s/application"
	ApplicationInstanceListPath         = "/csm/api/v1/userId/%s/instances/all"
	ApplicationInstanceCreatePath       = "/csm/api/v1/userId/%s/application/instance/bind"
	ApplicationInstanceCreatedListPath  = "/csm/api/v1/userId/%s/application/%s/instance/list"
	ApplicationInstanceDeletePath       = "/csm/api/v1/userId/%s/application/instance"
	ApplicationMonitorTaskCreatePath    = "/csm/api/v1/userId/%s/application/task/create"
	ApplicationMonitorTaskDetailPath    = "/csm/api/v1/userId/%s/application/%s/task/%s"
	ApplicationMonitorTaskListPath      = "/csm/api/v1/userId/%s/application/%s/task/list"
	ApplicationMonitorTaskDeletePath    = "/csm/api/v1/userId/%s/application/task/delete"
	ApplicationMonitorTaskUpdatePath    = "/csm/api/v1/userId/%s/application/task/update"
	ApplicationDimensionTableCreatePath = "/csm/api/v1/userId/%s/application/dimensionMap/create"
	ApplicationDimensionTableListPath   = "/csm/api/v1/userId/%s/application/%s/dimensionMap/list"
	ApplicationDimensionTableDeletePath = "/csm/api/v1/userId/%s/application/dimensionMap/delete"
	ApplicationDimensionTableUpdatePath = "/csm/api/v1/userId/%s/application/dimensionMap/update"

	EventCloudListPath    = "/event-api/v1/bce-event/list"
	EventPlatformListPath = "/event-api/v1/platform-event/list"
	EventPolicyPath       = "/event-api/v1/accounts/%s/services/%s/alarm-policies"

	InstanceGroupPath     = "/csm/api/v1/userId/%s/instance-group"
	InstanceGroupIdPath   = "/csm/api/v1/userId/%s/instance-group/%s"
	InstanceGroupListPath = "/csm/api/v1/userId/%s/instance-group/list"

	IG_INSTANCE_ADD_PATH               = "/csm/api/v1/userId/%s/instance-group/%s/instance/add"
	IG_INSTANCE_REMOVE_PATH            = "/csm/api/v1/userId/%s/instance-group/%s/instance/remove"
	IG_INSTANCE_LIST_PATH              = "/csm/api/v1/userId/%s/instance-group/instance/list"
	IG_QUERY_INSTANCE_LIST_PATH        = "/csm/api/v1/userId/%s/instance/list"
	IG_QUERY_INSTANCE_LIST_FILTER_PATH = "/csm/api/v1/userId/%s/instance/filteredList"

	MultiDimensionLatestMetricsPath = "/csm/api/v2/userId/%s/services/%s/data/metricData/latest/batch"
	MetricsByPartialDimensionsPath  = "/csm/api/v2/userId/%s/services/%s/data/metricData/PartialDimension"
	MultiMetricAllDataPath          = "/csm/api/v2/data/metricAllData"
	MultiMetricAllDataBatchPath     = "/csm/api/v2/data/metricAllData/batch"

	Average     = "average"
	Maximum     = "maximum"
	Minimum     = "minimum"
	Sum         = "sum"
	SampleCount = "sampleCount"

	NoticeEventLevel     = "NOTICE"
	WarningEventLevel    = "WARNING"
	MajorEventLevel      = "MAJOR"
	CriticalEventLevel   = "CRITICAL"
	DimensionNumberLimit = 100

	// CreateDashboardPath userId
	CreateDashboardPath = "/csm/api/v1/dashboard/products/%s/dashboards"

	// DashboardPath userId - dashboardName
	DashboardPath = "/csm/api/v1/dashboard/products/%s/dashboards/%s"

	// DuplicateDashboardPath userId - dashboardName
	DuplicateDashboardPath = "/csm/api/v1/dashboard/products/%s/dashboards/%s/duplicate"

	// DashboardWidgetPath userId - dashboardName - widgetName
	DashboardWidgetPath = "/csm/api/v1/dashboard/products/%s/dashboards/%s/widgets/%s"

	// CreateDashboardWidgetPath userId - dashboardName
	CreateDashboardWidgetPath = "/csm/api/v1/dashboard/products/%s/dashboards/%s/widgets"

	// DuplicateDashboardWidgetPath userId - dashboardName - widgetName
	DuplicateDashboardWidgetPath = "/csm/api/v1/dashboard/products/%s/dashboards/%s/widgets/%s/duplicate"

	ReportDataPath = "/csm/api/v1/dashboard/metric/report"

	TrendDataPath = "/csm/api/v1/dashboard/metric/trend"

	GaugeChartDataPath = "/csm/api/v1/dashboard/metric/gaugechart"

	BillboardDataPath = "/csm/api/v1/dashboard/metric/billboard"

	TrendSeniorDataPath = "/csm/api/v1/dashboard/metric/trend/senior"

	// DimensionsPath userId - serviceName - region
	DimensionsPath = "/csm/api/v1/userId/%s/services/%s/region/%s/metric/dimensions"

	SiteCreateHttpTaskPath  = "/csm/api/v1/userId/%s/site/http/create"
	SiteUpdateHttpTaskPath  = "/csm/api/v1/userId/%s/site/http/update"
	SiteGetHttpTaskPath     = "/csm/api/v1/userId/%s/site/http/detail"
	SiteCreateHttpsTaskPath = "/csm/api/v1/userId/%s/site/https/create"
	SiteUpdateHttpsTaskPath = "/csm/api/v1/userId/%s/site/https/update"
	SiteGetHttpsTaskPath    = "/csm/api/v1/userId/%s/site/https/detail"
	SiteCreatePingTaskPath  = "/csm/api/v1/userId/%s/site/ping/create"
	SiteUpdatePingTaskPath  = "/csm/api/v1/userId/%s/site/ping/update"
	SiteGetPingTaskPath     = "/csm/api/v1/userId/%s/site/ping/detail"
	SiteCreateTcpTaskPath   = "/csm/api/v1/userId/%s/site/tcp/create"
	SiteUpdateTcpTaskPath   = "/csm/api/v1/userId/%s/site/tcp/update"
	SiteGetTcpTaskPath      = "/csm/api/v1/userId/%s/site/tcp/detail"
	SiteCreateUdpTaskPath   = "/csm/api/v1/userId/%s/site/udp/create"
	SiteUpdateUdpTaskPath   = "/csm/api/v1/userId/%s/site/udp/update"
	SiteGetUdpTaskPath      = "/csm/api/v1/userId/%s/site/udp/detail"
	SiteCreateFtpTaskPath   = "/csm/api/v1/userId/%s/site/ftp/create"
	SiteUpdateFtpTaskPath   = "/csm/api/v1/userId/%s/site/ftp/update"
	SiteGetFtpTaskPath      = "/csm/api/v1/userId/%s/site/ftp/detail"
	SiteCreateDnsTaskPath   = "/csm/api/v1/userId/%s/site/dns/create"
	SiteUpdateDnsTaskPath   = "/csm/api/v1/userId/%s/site/dns/update"
	SiteGetDnsTaskPath      = "/csm/api/v1/userId/%s/site/dns/detail"
	SiteGetTaskListPath     = "/csm/api/v1/userId/%s/site/list"
	SiteDeleteTaskPath      = "/csm/api/v1/userId/%s/site/delete"
	SiteGetTaskDetailPath   = "/csm/api/v1/userId/%s/site/%s"

	SiteCreateAlarmConfigPath    = "/csm/api/v1/userId/%s/site/alarm/config/create"
	SiteUpdateAlarmConfigPath    = "/csm/api/v1/userId/%s/site/alarm/config/update"
	SiteDeleteAlarmConfigPath    = "/csm/api/v1/userId/%s/site/alarm/config/delete"
	SiteGetAlarmConfigDetailPath = "/csm/api/v1/userId/%s/site/alarm/config/detail"
	SiteGetAlarmConfigListPath   = "/csm/api/v1/userId/%s/site/alarm/config/list"
	SiteAlarmBlockPath           = "/csm/api/v1/userId/%s/site/alarm/config/block"
	SiteAlarmUnBlockPath         = "/csm/api/v1/userId/%s/site/alarm/config/unblock"
	SiteGetTaskByAlarmNamePath   = "/csm/api/v1/userId/%s/site/alarm/config/%s"
	SiteGetMetricDataPath        = "/csm/api/v1/userId/%s/site/metricSiteData"
	SiteGetOverallViewPath       = "/csm/api/v1/userId/%s/site/idc/overallView"
	SiteGetProvincialViewPath    = "/csm/api/v1/userId/%s/site/idc/provincialView"
	SiteAgentListPath            = "/csm/api/v1/userId/%s/site/agent/list"
	SiteGetAgentByTaskIdPath     = "/csm/api/v1/userId/%s/site/agent/idcIsp"
)

var eventLevel = map[string]bool{
	NoticeEventLevel:   true,
	WarningEventLevel:  true,
	MajorEventLevel:    true,
	CriticalEventLevel: true,
}

// GetMetricData get metric data
func (c *Client) GetMetricData(req *model.GetMetricDataRequest) (*model.GetMetricDataResponse, error) {
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.MetricName) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}
	if len(req.Statistics) <= 0 {
		return nil, errors.New("statistics should not be empty")
	}
	if req.PeriodInSecond < 10 {
		return nil, errors.New("periodInSecond should not be greater 10")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if len(req.Dimensions) <= 0 {
		return nil, errors.New("dimension should not be empty")
	}
	dimensionStr := ""
	for key, value := range req.Dimensions {
		dimensionStr = dimensionStr + key + ":" + value + ";"
	}
	dimensionStr = strings.TrimRight(dimensionStr, ";")

	result := &model.GetMetricDataResponse{}
	url := fmt.Sprintf(GetMetricDataPath, req.UserId, req.Scope, req.MetricName)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("dimensions", dimensionStr).
		WithQueryParam("statistics[]", strings.Join(req.Statistics, ",")).
		WithQueryParam("periodInSecond", strconv.Itoa(req.PeriodInSecond)).
		WithQueryParam("startTime", req.StartTime).
		WithQueryParam("endTime", req.EndTime).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// BatchGetMetricData batch get metric data
func (c *Client) BatchGetMetricData(req *model.BatchGetMetricDataRequest) (*model.BatchGetMetricDataResponse, error) {
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.MetricNames) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}
	if len(req.Statistics) <= 0 {
		return nil, errors.New("statistics should not be empty")
	}
	if req.PeriodInSecond < 10 {
		return nil, errors.New("periodInSecond should not be greater 10")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if len(req.Dimensions) <= 0 {
		return nil, errors.New("dimension should not be empty")
	}
	dimensionStr := ""
	for key, value := range req.Dimensions {
		dimensionStr = dimensionStr + key + ":" + value + ";"
	}
	dimensionStr = strings.TrimRight(dimensionStr, ";")

	result := &model.BatchGetMetricDataResponse{}
	url := fmt.Sprintf(BatchMetricDataPath, req.UserId, req.Scope)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("metricName[]", strings.Join(req.MetricNames, ",")).
		WithQueryParam("dimensions", dimensionStr).
		WithQueryParam("statistics[]", strings.Join(req.Statistics, ",")).
		WithQueryParam("periodInSecond", strconv.Itoa(req.PeriodInSecond)).
		WithQueryParam("startTime", req.StartTime).
		WithQueryParam("endTime", req.EndTime).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// CreateApplicationData create application data
func (c *Client) CreateApplicationData(req *model.ApplicationInfoRequest) (*model.ApplicationInfoResponse, error) {
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Name) <= 0 {
		return nil, errors.New("name should not be empty")
	}
	if len(req.Type) <= 0 {
		return nil, errors.New("type should not be empty")
	}
	result := &model.ApplicationInfoResponse{}
	url := fmt.Sprintf(ApplicationInfoPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetApplicationDataList get application data list
func (c *Client) GetApplicationDataList(userId string, searchName string, pageSize int, pageNo int) (*model.ApplicationInfoListResponse, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be invalid")
	}
	if pageSize <= 0 || pageSize > 100 {
		return nil, errors.New("pageSize should not be invalid")
	}
	if pageNo <= 0 {
		return nil, errors.New("pageNo should not be invalid")
	}
	result := &model.ApplicationInfoListResponse{}
	url := fmt.Sprintf(ApplicationInfoPath, userId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("pageSize", strconv.Itoa(pageSize)).
		WithQueryParam("pageNo", strconv.Itoa(pageNo)).
		WithQueryParam("searchName", searchName).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// UpdateApplicationData update application data
func (c *Client) UpdateApplicationData(req *model.ApplicationInfoUpdateRequest) (*model.ApplicationInfoResponse, error) {
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.ID == 0 {
		return nil, errors.New("id should not be empty")
	}
	if len(req.Name) <= 0 {
		return nil, errors.New("name should not be empty")
	}
	if len(req.Type) <= 0 {
		return nil, errors.New("type should not be empty")
	}
	result := &model.ApplicationInfoResponse{}
	url := fmt.Sprintf(ApplicationInfoPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// DeleteApplicationData delete application data
func (c *Client) DeleteApplicationData(userId string, req *model.ApplicationInfoDeleteRequest) error {
	if len(userId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Name) <= 0 {
		return errors.New("name should not be empty")
	}
	url := fmt.Sprintf(ApplicationInfoPath, userId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.DELETE).
		Do()
	return err
}

// GetApplicationInstanceList get application instance list
func (c *Client) GetApplicationInstanceList(userId string, req *model.ApplicationInstanceListRequest) (*model.ApplicationInstanceListResponse, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.SearchName) <= 0 {
		return nil, errors.New("searchName should not be empty")
	}
	if req.PageNo == 0 || req.PageSize == 0 {
		req.PageNo = 1
		req.PageSize = 10
	}
	result := &model.ApplicationInstanceListResponse{}
	url := fmt.Sprintf(ApplicationInstanceListPath, userId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// CreateApplicationInstance create application instance
func (c *Client) CreateApplicationInstance(req *model.ApplicationInstanceCreateRequest) error {
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return errors.New("appName should not be empty")
	}
	if len(req.HostList) <= 0 {
		return errors.New("hostList should not be empty")
	}
	url := fmt.Sprintf(ApplicationInstanceCreatePath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		Do()
	return err
}

// GetApplicationInstanceCreatedList get application instance created list
func (c *Client) GetApplicationInstanceCreatedList(req *model.ApplicationInstanceCreatedListRequest) (*model.ApplicationInstanceCreatedListResponse, error) {
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	url := fmt.Sprintf(ApplicationInstanceCreatedListPath, req.UserID, req.AppName)
	result := &model.ApplicationInstanceCreatedListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("region", req.Region).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// DeleteApplicationInstance delete application instance
func (c *Client) DeleteApplicationInstance(userId string, req *model.ApplicationInstanceDeleteRequest) error {
	if len(userId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.ID) <= 0 {
		return errors.New("id should not be empty")
	}
	if len(req.AppName) <= 0 {
		return errors.New("appName should not be empty")
	}
	url := fmt.Sprintf(ApplicationInstanceDeletePath, userId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.DELETE).
		Do()
	return err
}
func (c *Client) CreateApplicationInstanceTask(userId string, req *model.ApplicationMonitorTaskInfoRequest) (*model.ApplicationMonitorTaskResponse, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.Type < 0 || req.Type > 3 {
		return nil, errors.New("type should be 0-3")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.AliasName) <= 0 {
		return nil, errors.New("aliasName should not be empty")
	}
	if len(req.Target) <= 0 {
		return nil, errors.New("target should not be empty")
	}
	url := fmt.Sprintf(ApplicationMonitorTaskCreatePath, userId)
	result := &model.ApplicationMonitorTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// CreateApplicationMonitorLogTask Create the application monitor LOG task
func (c *Client) CreateApplicationMonitorLogTask(userId string,
	req *model.ApplicationMonitorTaskInfoLogRequest) (*model.ApplicationMonitorTaskInfoLogResponse, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.Type < 0 || req.Type > 3 {
		return nil, errors.New("type should be 0-3")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.AliasName) <= 0 {
		return nil, errors.New("aliasName should not be empty")
	}
	if len(req.Target) <= 0 {
		return nil, errors.New("target should not be empty")
	}
	if req.Type == 2 {
		if len(req.LogExample) <= 0 {
			return nil, errors.New("logExample should not be empty")
		}
		if len(req.MatchRule) <= 0 {
			return nil, errors.New("matchRule should not be empty")
		}
		if req.Rate == 0 {
			return nil, errors.New("rate should not be empty")
		}
		if len(req.ExtractResult) <= 0 {
			return nil, errors.New("extractResult should not be empty")
		}
		if len(req.Metrics) <= 0 {
			return nil, errors.New("metrics should not be empty")
		}
	}
	url := fmt.Sprintf(ApplicationMonitorTaskCreatePath, userId)
	result := &model.ApplicationMonitorTaskInfoLogResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetApplicationMonitorTaskDetail Get the application monitor task detail
func (c *Client) GetApplicationMonitorTaskDetail(req *model.ApplicationMonitorTaskDetailRequest) (*model.ApplicationMonitorTaskInfoLogResponse, error) {
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return nil, errors.New("taskName should not be empty")
	}
	url := fmt.Sprintf(ApplicationMonitorTaskDetailPath, req.UserID, req.AppName, req.TaskName)
	result := &model.ApplicationMonitorTaskInfoLogResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetApplicationMonitorTaskList Get the application monitor task list
func (c *Client) GetApplicationMonitorTaskList(req *model.ApplicationMonitorTaskListRequest) ([]*model.ApplicationMonitorTaskInfoListResponse, error) {
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	url := fmt.Sprintf(ApplicationMonitorTaskListPath, req.UserID, req.AppName)
	var result []*model.ApplicationMonitorTaskInfoListResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("type", req.Type).
		WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

func (c *Client) UpdateApplicationMonitorTask(userId string,
	req *model.ApplicationMonitorTaskInfoUpdateRequest) (*model.ApplicationMonitorTaskInfoNormalResponse, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.Type < 0 || req.Type > 3 {
		return nil, errors.New("type should be 0-3")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.AliasName) <= 0 {
		return nil, errors.New("aliasName should not be empty")
	}
	if len(req.Target) <= 0 {
		return nil, errors.New("target should not be empty")
	}
	if len(req.Name) <= 0 {
		return nil, errors.New("name should not be empty")
	}
	url := fmt.Sprintf(ApplicationMonitorTaskUpdatePath, userId)
	result := &model.ApplicationMonitorTaskInfoNormalResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err

}

// UpdateApplicationMonitorLogTask Update the application monitor task
func (c *Client) UpdateApplicationMonitorLogTask(userId string,
	req *model.ApplicationMonitorTaskInfoUpdateRequest) (*model.ApplicationMonitorTaskInfoLogResponse, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.Type < 0 || req.Type > 3 {
		return nil, errors.New("type should be 0-3")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.AliasName) <= 0 {
		return nil, errors.New("aliasName should not be empty")
	}
	if len(req.Target) <= 0 {
		return nil, errors.New("target should not be empty")
	}
	if req.Type == 2 {
		if len(req.LogExample) <= 0 {
			return nil, errors.New("logExample should not be empty")
		}
		if len(req.MatchRule) <= 0 {
			return nil, errors.New("matchRule should not be empty")
		}
		if req.Rate == 0 {
			return nil, errors.New("rate should not be empty")
		}
		if len(req.ExtractResult) <= 0 {
			return nil, errors.New("extractResult should not be empty")
		}
		if len(req.Metrics) <= 0 {
			return nil, errors.New("metrics should not be empty")
		}
	}
	url := fmt.Sprintf(ApplicationMonitorTaskUpdatePath, userId)
	result := &model.ApplicationMonitorTaskInfoLogResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// DeleteApplicationMonitorTask Delete the application monitor task
func (c *Client) DeleteApplicationMonitorTask(req *model.ApplicationMonitorTaskDeleteRequest) error {
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return errors.New("appName should not be empty")
	}
	if len(req.Name) <= 0 {
		return errors.New("name should not be empty")
	}
	url := fmt.Sprintf(ApplicationMonitorTaskDeletePath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.DELETE).
		Do()
	return err
}

// CreateApplicationDimensionTable Create the application dimension table
func (c *Client) CreateApplicationDimensionTable(req *model.ApplicationDimensionTableInfoRequest) (*model.ApplicationDimensionTableInfoResponse, error) {
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.TableName) <= 0 {
		return nil, errors.New("tableName should not be empty")
	}
	if len(req.MapContentJSON) <= 0 {
		return nil, errors.New("mapContentJson should not be empty")
	}
	url := fmt.Sprintf(ApplicationDimensionTableCreatePath, req.UserID)
	result := &model.ApplicationDimensionTableInfoResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetApplicationDimensionTableList Get the application dimension table list
func (c *Client) GetApplicationDimensionTableList(req *model.ApplicationDimensionTableListRequest) ([]*model.ApplicationDimensionTableInfoResponse, error) {
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	url := fmt.Sprintf(ApplicationDimensionTableListPath, req.UserID, req.AppName)
	var result []*model.ApplicationDimensionTableInfoResponse
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("searchName", req.SearchName).
		WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

// UpdateApplicationDimensionTable Update the application dimension table
func (c *Client) UpdateApplicationDimensionTable(req *model.ApplicationDimensionTableInfoRequest) error {
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return errors.New("appName should not be empty")
	}
	if len(req.TableName) <= 0 {
		return errors.New("tableName should not be empty")
	}
	if len(req.MapContentJSON) <= 0 {
		return errors.New("mapContentJson should not be empty")
	}
	url := fmt.Sprintf(ApplicationDimensionTableUpdatePath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		Do()
	return err
}

// DeleteApplicationDimensionTable Delete the application dimension table
func (c *Client) DeleteApplicationDimensionTable(req *model.ApplicationDimensionTableDeleteRequest) error {
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return errors.New("appName should not be empty")
	}
	if len(req.TableName) <= 0 {
		return errors.New("tableName should not be empty")
	}
	url := fmt.Sprintf(ApplicationDimensionTableDeletePath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.DELETE).
		Do()
	return err
}

// CreateNamespace create custom monitor namespace
func (c *Client) CreateNamespace(ns *model.Namespace) error {
	if len(ns.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(ns.Name) <= 0 {
		return errors.New("name should not be empty")
	}

	url := fmt.Sprintf(CreateNamespacePath, ns.UserId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(ns).
		Do()
	return err
}

// BatchDeleteNamespaces batch delete custom monitor namespace
func (c *Client) BatchDeleteNamespaces(cns *model.CustomBatchNames) error {
	if len(cns.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(cns.Names) <= 0 {
		return errors.New("names should not be empty")
	}

	url := fmt.Sprintf(BatchDeleteNamespacesPath, cns.UserId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(cns).
		Do()
	return err
}

// UpdateNamespace update custom monitor namespace
func (c *Client) UpdateNamespace(ns *model.Namespace) error {
	if len(ns.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(ns.Name) <= 0 {
		return errors.New("name should not be empty")
	}

	url := fmt.Sprintf(UpdateNamespacePath, ns.UserId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.PUT).
		WithBody(ns).
		Do()
	return err
}

// ListNamespaces list custom monitor namespaces
func (c *Client) ListNamespaces(req *model.ListNamespacesRequest) (*model.ListNamespacesResponse, error) {
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo == 0 || req.PageSize == 0 {
		req.PageNo = 1
		req.PageSize = 10
	}

	result := &model.ListNamespacesResponse{}
	url := fmt.Sprintf(ListNamespacesPath, req.UserId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithQueryParam("name", req.Name).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

// CreateNamespaceMetric create metric in custom monitor namespace
func (c *Client) CreateNamespaceMetric(nm *model.NamespaceMetric) error {
	if len(nm.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(nm.Namespace) <= 0 {
		return errors.New("namespace should not be empty")
	}
	if len(nm.MetricName) <= 0 {
		return errors.New("metricName should not be empty")
	}
	if nm.Cycle <= 0 {
		return errors.New("cycle should not greater 0")
	}
	if nm.Dimensions == nil {
		nm.Dimensions = []model.NamespaceMetricDimension{}
	}

	url := fmt.Sprintf(CreateNamespaceMetricPath, nm.UserId, nm.Namespace)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(nm).
		Do()
	return err
}

// BatchDeleteNamespaceMetrics batch delete metrics in custom monitor namespace
func (c *Client) BatchDeleteNamespaceMetrics(cis *model.CustomBatchIds) error {
	if len(cis.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(cis.Namespace) <= 0 {
		return errors.New("namespace should not be empty")
	}
	if len(cis.Ids) <= 0 {
		return errors.New("ids should not be empty")
	}

	url := fmt.Sprintf(BatchDeleteNamespaceMetricsPath, cis.UserId, cis.Namespace)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.POST).
		WithBody(cis).
		Do()
	return err
}

// UpdateNamespaceMetric update metric in custom monitor namespace
func (c *Client) UpdateNamespaceMetric(nm *model.NamespaceMetric) error {
	if len(nm.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(nm.Namespace) <= 0 {
		return errors.New("namespace should not be empty")
	}
	if len(nm.MetricName) <= 0 {
		return errors.New("metricName should not be empty")
	}
	if nm.Cycle <= 0 {
		return errors.New("cycle should not greater 0")
	}
	if nm.Dimensions == nil {
		nm.Dimensions = []model.NamespaceMetricDimension{}
	}

	url := fmt.Sprintf(UpdateNamespaceMetricPath, nm.UserId, nm.Namespace, nm.MetricName)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.PUT).
		WithBody(nm).
		Do()
	return err
}

// ListNamespaceMetrics list metrics in custom monitor namespace
func (c *Client) ListNamespaceMetrics(req *model.ListNamespaceMetricsRequest) (*model.ListNamespaceMetricsResponse, error) {
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Namespace) <= 0 {
		return nil, errors.New("namespace should not be empty")
	}
	if req.PageNo == 0 || req.PageSize == 0 {
		req.PageNo = 1
		req.PageSize = 10
	}

	result := &model.ListNamespaceMetricsResponse{}
	url := fmt.Sprintf(ListNamespaceMetricsPath, req.UserId)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithQueryParam("namespace", req.Namespace).
		WithQueryParam("metricName", req.MetricName).
		WithQueryParam("metricAlias", req.MetricAlias).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

// GetCustomMetric get metric detail in custom monitor namespace
func (c *Client) GetCustomMetric(userId string, namespace string, metricName string) (*model.NamespaceMetric, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(namespace) <= 0 {
		return nil, errors.New("namespace should not be empty")
	}
	if len(metricName) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}

	result := &model.NamespaceMetric{}
	url := fmt.Sprintf(GetCustomMetricPath, userId, namespace, metricName)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// CreateNamespaceEvent create event in custom monitor namespace
func (c *Client) CreateNamespaceEvent(ne *model.NamespaceEvent) error {
	if len(ne.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(ne.Namespace) <= 0 {
		return errors.New("namespace should not be empty")
	}
	if len(ne.EventName) <= 0 {
		return errors.New("eventName should not be empty")
	}
	if len(ne.EventLevel) <= 0 {
		return errors.New("eventLevel should not be empty")
	}
	if !eventLevel[ne.EventLevel] {
		return errors.New(fmt.Sprintf("eventLevel is one of [%s, %s, %s, %s]",
			NoticeEventLevel, WarningEventLevel, MajorEventLevel, CriticalEventLevel))
	}

	err := bce.NewRequestBuilder(c).
		WithURL(CreateNamespaceEventPath).
		WithMethod(http.POST).
		WithBody(ne).
		Do()
	return err
}

// BatchDeleteNamespaceEvents batch delete events in custom monitor namespace
func (c *Client) BatchDeleteNamespaceEvents(ces *model.CustomBatchEventNames) error {
	if len(ces.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(ces.Namespace) <= 0 {
		return errors.New("namespace should not be empty")
	}
	if len(ces.Names) <= 0 {
		return errors.New("names should not be empty")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(BatchDeleteNamespaceEventsPath).
		WithMethod(http.POST).
		WithBody(ces).
		Do()
	return err
}

// UpdateNamespaceEvent update event in custom monitor namespace
func (c *Client) UpdateNamespaceEvent(ne *model.NamespaceEvent) error {
	if len(ne.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(ne.Namespace) <= 0 {
		return errors.New("namespace should not be empty")
	}
	if len(ne.EventName) <= 0 {
		return errors.New("eventName should not be empty")
	}
	if len(ne.EventLevel) <= 0 {
		return errors.New("eventLevel should not be empty")
	}
	if !eventLevel[ne.EventLevel] {
		return errors.New(fmt.Sprintf("eventLevel must be one of [%s, %s, %s, %s]",
			NoticeEventLevel, WarningEventLevel, MajorEventLevel, CriticalEventLevel))
	}

	err := bce.NewRequestBuilder(c).
		WithURL(UpdateNamespaceEventPath).
		WithMethod(http.POST).
		WithBody(ne).
		Do()
	return err
}

// ListNamespaceEvents list events in custom monitor namespace
func (c *Client) ListNamespaceEvents(req *model.ListNamespaceEventsRequest) (*model.ListNamespaceEventsResponse, error) {
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Namespace) <= 0 {
		return nil, errors.New("namespace should not be empty")
	}
	if len(req.EventLevel) > 0 && !eventLevel[req.EventLevel] {
		return nil, errors.New(fmt.Sprintf("eventLevel must be one of [%s, %s, %s, %s]",
			NoticeEventLevel, WarningEventLevel, MajorEventLevel, CriticalEventLevel))
	}
	if req.PageNo == 0 || req.PageSize == 0 {
		req.PageNo = 1
		req.PageSize = 10
	}

	result := &model.ListNamespaceEventsResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(ListNamespaceEventsPath).
		WithMethod(http.GET).
		WithQueryParam("userId", req.UserId).
		WithQueryParam("namespace", req.Namespace).
		WithQueryParam("name", req.Name).
		WithQueryParam("eventLevel", req.EventLevel).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

// GetCustomEvent get event detail in custom monitor namespace
func (c *Client) GetCustomEvent(userId string, namespace string, eventName string) (*model.NamespaceEvent, error) {
	if len(userId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(namespace) <= 0 {
		return nil, errors.New("namespace should not be empty")
	}
	if len(eventName) <= 0 {
		return nil, errors.New("eventName should not be empty")
	}

	result := &model.NamespaceEvent{}
	err := bce.NewRequestBuilder(c).
		WithURL(GetCustomEventPath).
		WithMethod(http.GET).
		WithQueryParam("userId", userId).
		WithQueryParam("namespace", namespace).
		WithQueryParam("eventName", eventName).
		WithResult(result).
		Do()
	return result, err
}

// ListNotifyGroups ListCustomEvents list events in custom monitor namespace
func (c *Client) ListNotifyGroups(req *model.ListNotifyGroupsRequest) (*model.ListNotifyGroupsResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should be greater than 0")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should be greater than 0")
	}
	result := &model.ListNotifyGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/json-api/v1/alarm/notify/group/list").
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// ListNotifyParty ListNotifyParties list notify parties
func (c *Client) ListNotifyParty(req *model.ListNotifyPartiesRequest) (*model.ListNotifyPartiesResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should be greater than 0")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should be greater than 0")
	}
	result := &model.ListNotifyPartiesResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/json-api/v1/alarm/notify/party/list").
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// CreateAction create action
func (c *Client) CreateAction(req *model.CreateActionRequest) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Notifications) <= 0 {
		return errors.New("notifications should not be empty")
	}
	if len(req.Members) <= 0 {
		return errors.New("members should not be empty")
	}
	if len(req.Alias) <= 0 {
		return errors.New("alias should not be empty")
	}
	if req.ActionCallBacks == nil {
		req.ActionCallBacks = make([]model.ActionCallBack, 0)
	}
	if req.DisableTimes == nil {
		req.DisableTimes = []model.ActionDisableTime{{From: "00:00:00", To: "00:00:00"}}
	}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/action/create", req.UserId)).
		WithBody(req).
		WithMethod(http.POST).
		Do()

	return err
}

// DeleteAction delete action
func (c *Client) DeleteAction(req *model.DeleteActionRequest) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Name) <= 0 {
		return errors.New("name should not be empty")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/action/delete", req.UserId)).
		WithQueryParam("name", req.Name).
		WithMethod(http.DELETE).
		Do()
	return err
}

// ListActions list actions
func (c *Client) ListActions(req *model.ListActionsRequest) (*model.ListActionsResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should be greater than 0")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should be greater than 0")
	}
	result := &model.ListActionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/action/actionList", req.UserId)).
		WithBody(req).
		WithResult(result).
		WithMethod(http.POST).
		Do()
	return result, err
}

// UpdateAction update action
func (c *Client) UpdateAction(req *model.UpdateActionRequest) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Name) <= 0 {
		return errors.New("name should not be empty")
	}
	if len(req.Notifications) <= 0 {
		return errors.New("notifications should not be empty")
	}
	if len(req.Members) <= 0 {
		return errors.New("members should not be empty")
	}
	if len(req.Alias) <= 0 {
		return errors.New("alias should not be empty")
	}
	if req.ActionCallBacks == nil {
		req.ActionCallBacks = make([]model.ActionCallBack, 0)
	}
	if req.DisableTimes == nil {
		req.DisableTimes = []model.ActionDisableTime{{From: "00:00:00", To: "00:00:00"}}
	}
	req.Source = "USER"
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/action/update", req.UserId)).
		WithBody(req).
		WithMethod(http.PUT).
		Do()
	return err
}

// GetMetricMetaForApplication get metric meta for application
func (c *Client) GetMetricMetaForApplication(req *model.GetMetricMetaForApplicationRequest) (map[string][]string, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return nil, errors.New("taskName should not be empty")
	}
	if len(req.MetricName) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}
	if len(req.DimensionKeys) <= 0 {
		return nil, errors.New("dimensionKeys should not be empty")
	}
	url := fmt.Sprintf("/csm/api/v1/userId/%s/application/%s/task/%s/metricMeta", req.UserId, req.Instances, req.TaskName)
	result := make(map[string][]string)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("metricName", req.MetricName).
		WithQueryParam("dimensionKeys", strings.Join(req.DimensionKeys, ",")).
		WithQueryParamFilter("instances", strings.Join(req.Instances, ",")).
		WithResult(&result).
		WithMethod(http.GET).
		Do()
	return result, err
}

// GetMetricDataForApplication get metric data for application
func (c *Client) GetMetricDataForApplication(req *model.GetMetricDataForApplicationRequest) ([]*model.GetMetricDataForApplicationResult, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return nil, errors.New("taskName should not be empty")
	}
	if len(req.MetricName) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if !req.AggrData && len(req.Instances) == 0 {
		return nil, errors.New("instances should not be empty when aggrData is false")
	}
	url := fmt.Sprintf("/csm/api/v1/userId/%s/application/%s/task/%s/metricData", req.UserId, req.AppName, req.TaskName)
	rb := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("startTime", req.StartTime).
		WithQueryParamFilter("endTime", req.EndTime).
		WithQueryParam("metricName", req.MetricName).
		WithQueryParamFilter("statistics", strings.Join(req.Statistics, ",")).
		WithQueryParamFilter("instances", strings.Join(req.Instances, ",")).
		WithQueryParam("aggrData", strconv.FormatBool(req.AggrData))

	if req.Cycle > 0 {
		rb.WithQueryParam("cycle", strconv.Itoa(req.Cycle))
	}
	if len(req.Dimensions) > 0 {
		dims := make([]string, 0, len(req.Dimensions))
		for name, values := range req.Dimensions {
			dims = append(dims, name+":"+strings.Join(values, "___"))
		}
		rb.WithQueryParam("dimensions", strings.Join(dims, ","))
	}
	result := make([]*model.GetMetricDataForApplicationResult, 0)
	err := rb.WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

// GetAlarmMetricsForApplication get alarm metrics for application
func (c *Client) GetAlarmMetricsForApplication(req *model.GetAppMonitorAlarmMetricsRequest) ([]*model.AppMetric, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return nil, errors.New("taskName should not be empty")
	}
	url := fmt.Sprintf("/csm/api/v1/userId/%s/application/%s/%s/alarm/metrics", req.UserId, req.AppName, req.TaskName)
	result := make([]*model.AppMetric, 0)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParamFilter("searchName", req.SearchName).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// CreateAppMonitorAlarmConfig create application monitor alarm config
func (c *Client) CreateAppMonitorAlarmConfig(req *model.AppMonitorAlarmConfig) (*model.AppMonitorAlarmConfig, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return nil, errors.New("alarmName should not be empty")
	}
	if req.MonitorObjectType != model.MonitorObjectApp && req.MonitorObjectType != model.MonitorObjectService {
		return nil, errors.New("monitorObjectType must be APP or SERVICE")
	}
	if len(req.SrcName) <= 0 {
		return nil, errors.New("srcName should not be empty")
	}
	if req.SrcType != model.SrcTypePort && req.SrcType != model.SrcTypeLog && req.SrcType != model.SrcTypeProc && req.SrcType != model.SrcTypeSCR {
		return nil, errors.New("srcType must be one of PROC,SCR,PORT,LOG")
	}
	if req.Type != model.AlarmTypeInstance && req.Type != model.AlarmTypeService {
		return nil, errors.New("type must be INSTANCE or SERVICE")
	}
	if len(req.Rules) <= 0 {
		return nil, errors.New("rules should not be empty")
	}

	result := &model.AppMonitorAlarmConfig{}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/application/alarm/config/create", req.UserId)).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateAppMonitorAlarmConfig update application monitor alarm config
func (c *Client) UpdateAppMonitorAlarmConfig(req *model.AppMonitorAlarmConfig) (*model.AppMonitorAlarmConfig, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return nil, errors.New("alarmName should not be empty")
	}
	if req.MonitorObjectType != model.MonitorObjectApp && req.MonitorObjectType != model.MonitorObjectService {
		return nil, errors.New("monitorObjectType must be APP or SERVICE")
	}
	if len(req.SrcName) <= 0 {
		return nil, errors.New("srcName should not be empty")
	}
	if req.SrcType != model.SrcTypePort && req.SrcType != model.SrcTypeLog && req.SrcType != model.SrcTypeProc && req.SrcType != model.SrcTypeSCR {
		return nil, errors.New("srcType must be one of PROC,SCR,PORT,LOG")
	}
	if req.Type != model.AlarmTypeInstance && req.Type != model.AlarmTypeService {
		return nil, errors.New("type must be INSTANCE or SERVICE")
	}
	if len(req.Rules) <= 0 {
		return nil, errors.New("rules should not be empty")
	}

	result := &model.AppMonitorAlarmConfig{}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/application/alarm/config/update", req.UserId)).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// ListAppMonitorAlarmConfigs list application monitor alarm configs
func (c *Client) ListAppMonitorAlarmConfigs(req *model.ListAppMonitorAlarmConfigsRequest) (*model.ListAppMonitorAlarmConfigsResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should be greater than 0")
	}

	result := &model.ListAppMonitorAlarmConfigsResponse{}
	rb := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/application/alarm/config/list", req.UserId)).
		WithQueryParamFilter("appName", req.AppName).
		WithQueryParamFilter("alarmName", req.AlarmName).
		WithQueryParamFilter("taskName", req.TaskName).
		WithQueryParamFilter("srcType", string(req.SrcType)).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo))
	if req.ActionEnabled != nil {
		rb.WithQueryParam("actionEnabled", strconv.FormatBool(*req.ActionEnabled))
	} else {
		rb.WithQueryParam("actionEnabled", "")
	}
	if req.PageSize > 0 {
		rb.WithQueryParam("pageSize", strconv.Itoa(req.PageSize))
	}
	err := rb.WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// DeleteAppMonitorAlarmConfig delete application monitor alarm config
func (c *Client) DeleteAppMonitorAlarmConfig(req *model.DeleteAppMonitorAlarmConfigRequest) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return errors.New("appName should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return errors.New("alarmName should not be empty")
	}

	return bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/application/alarm/config", req.UserId)).
		WithBody(req).
		WithMethod(http.DELETE).
		Do()
}

// GetAppMonitorAlarmConfig get application monitor alarm config
func (c *Client) GetAppMonitorAlarmConfig(req *model.GetAppMonitorAlarmConfigDetailRequest) (*model.AppMonitorAlarmConfig, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AppName) <= 0 {
		return nil, errors.New("appName should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return nil, errors.New("alarmName should not be empty")
	}

	result := &model.AppMonitorAlarmConfig{}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/application/alarm/%s/config", req.UserId, req.AlarmName)).
		WithQueryParam("appName", req.AppName).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err

}

// LogExtract log extract
func (c *Client) LogExtract(req *model.LogExtractRequest) ([]*model.LogExtractResult, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.ExtractRule) <= 0 {
		return nil, errors.New("extractRule should not be null")
	}
	if len(req.LogExample) <= 0 {
		return nil, errors.New("logExample should not be null")
	}
	result := make([]*model.LogExtractResult, 0)
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v1/userId/%s/application/logextract", req.UserId)).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(&result).
		Do()
	return result, err
}

// CreateDashboard 向服务端创建仪表盘
//
// 参数：
//
//	req *DashboardRequest: 创建仪表盘请求的参数结构体
//
// 返回值：
//
//	*DashboardBaseResponse: 创建仪表盘返回的结果结构体
//	error: 错误信息，如果操作成功则为 nil
func (c *Client) CreateDashboard(req *model.DashboardRequest) (*model.DashboardBaseResponse, error) {
	err := CheckCreateDashboardRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardBaseResponse{}
	url := fmt.Sprintf(CreateDashboardPath, req.UserID)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// DeleteDashboard 删除仪表盘
// 参数：
//
//	req *DashboardRequest：删除仪表盘请求参数的结构体
//
// 返回值：
//
//	*DashboardResponse：删除仪表盘返回结果的结构体
//	error：请求过程中出现的错误，如果没有错误则为nil
func (c *Client) DeleteDashboard(req *model.DashboardRequest) (*model.DashboardResponse, error) {
	err := CheckDashboardRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardResponse{}
	url := fmt.Sprintf(DashboardPath, req.UserID, req.DashboardName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.DELETE).
		WithResult(result).
		Do()
	return result, err
}

// UpdateDashboard 向服务端更新仪表盘信息
// 参数 req 表示更新仪表盘的请求，包括 UserID 和 DashboardName
// 返回更新后的仪表盘信息及错误信息
func (c *Client) UpdateDashboard(req *model.DashboardRequest) (*model.DashboardResponse, error) {
	err := CheckUpdateDashboardRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardResponse{}
	url := fmt.Sprintf(DashboardPath, req.UserID, req.DashboardName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboard 用于获取仪表盘信息
// 参数 req 是 DashboardRequest 类型指针，表示请求体
// 返回值是 DashboardBaseResponse 类型指针和一个 error 类型值
func (c *Client) GetDashboard(req *model.DashboardRequest) (*model.DashboardBaseResponse, error) {
	err := CheckDashboardRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardBaseResponse{}
	url := fmt.Sprintf(DashboardPath, req.UserID, req.DashboardName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// DuplicateDashboard 方法用于复制仪表盘
// 参数req表示复制仪表盘请求，包含用户ID和仪表盘名称
// 返回复制仪表盘的响应和错误信息
func (c *Client) DuplicateDashboard(req *model.DashboardRequest) (*model.DashboardResponse, error) {
	err := CheckDashboardRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardResponse{}
	url := fmt.Sprintf(DuplicateDashboardPath, req.UserID, req.DashboardName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboardWidget 从云监控获取仪表板部件详情
// req: 仪表板部件请求参数
// 返回获取到的仪表板部件详情和错误信息
func (c *Client) GetDashboardWidget(req *model.DashboardWidgetRequest) (*model.DashboardBaseResponse, error) {
	err := CheckWidgetRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardBaseResponse{}
	url := fmt.Sprintf(DashboardWidgetPath, req.UserID, req.DashboardName, req.WidgetName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// CreateDashboardWidget 用于创建仪表盘Widget
// 参数 req 表示创建Widget请求的参数
// 返回值 result 表示创建Widget的返回结果，error 表示请求错误
func (c *Client) CreateDashboardWidget(req *model.DashboardWidgetRequest) (*model.DashboardResponse, error) {
	err := CheckCreateWidgetRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardResponse{}
	url := fmt.Sprintf(CreateDashboardWidgetPath, req.UserID, req.DashboardName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// DeleteDashboardWidget 从用户指定的仪表盘删除指定的Widget
//
// 参数：
//
//	req *DashboardRequest: 包含用户ID，仪表盘名称和Widget名称的请求体
//
// 返回值：
//
//	*DashboardResponse: 包含操作结果的响应体指针
//	error: 错误信息，如果操作成功则为nil
func (c *Client) DeleteDashboardWidget(req *model.DashboardWidgetRequest) (*model.DashboardResponse, error) {
	err := CheckWidgetRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardResponse{}
	url := fmt.Sprintf(DashboardWidgetPath, req.UserID, req.DashboardName, req.WidgetName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.DELETE).
		WithResult(result).
		Do()
	return result, err
}

// UpdateDashboardWidget 用于更新指定用户的指定仪表盘页面的指定小控件。
//
// req 表示更新小控件请求的 DashboardRequest 结构体指针。
//
// 返回更新小控件请求的 DashboardResponse 结构体指针和 error。
func (c *Client) UpdateDashboardWidget(req *model.DashboardWidgetRequest) (*model.DashboardResponse, error) {
	err := CheckUpdateWidgetRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardResponse{}
	url := fmt.Sprintf(DashboardWidgetPath, req.UserID, req.DashboardName, req.WidgetName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// DuplicateDashboardWidget 函数用于复制仪表板小部件
// 参数req表示复制请求，包括UserID，DashboardName和WidgetName三个字段
// 返回值 DashboardResponse 为复制后的小部件信息，error为请求过程中发生的错误
func (c *Client) DuplicateDashboardWidget(req *model.DashboardWidgetRequest) (*model.DashboardResponse, error) {
	err := CheckWidgetRequest(req)
	if err != nil {
		return nil, err
	}
	result := &model.DashboardResponse{}
	url := fmt.Sprintf(DuplicateDashboardWidgetPath, req.UserID, req.DashboardName, req.WidgetName)
	err = bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboardReportData 是一个 Client 类型的方法，用于获取仪表板表格数据。
// 它接收一个 model.DashboardDataRequest 指针类型的请求参数 req 和返回一个 model.DashboardReportDataResponse 指针类型的结果和错误信息。
func (c *Client) GetDashboardReportData(req *model.DashboardDataRequest) (*model.DashboardReportDataResponse, error) {
	result := &model.DashboardReportDataResponse{}
	url := fmt.Sprintf(ReportDataPath)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboardTrendData 从云监控获取仪表板趋势数据
// req: 包含请求信息的结构体
// 返回获取到的仪表板趋势数据和错误信息
func (c *Client) GetDashboardTrendData(req *model.DashboardDataRequest) (*model.DashboardTrendResponse, error) {
	result := &model.DashboardTrendResponse{}
	url := fmt.Sprintf(TrendDataPath)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboardGaugeChartData 获取仪表盘的仪表盘数据
// 参数 req 表示获取数据的请求，类型为 *model.DashboardDataRequest
// 返回获取到的仪表盘数据和可能产生的错误
func (c *Client) GetDashboardGaugeChartData(req *model.DashboardDataRequest) (*model.DashboardBillboardResponse, error) {
	result := &model.DashboardBillboardResponse{}
	url := fmt.Sprintf(GaugeChartDataPath)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboardBillboardData 获取 Billboard 数据
// 参数：
//
//	req *model.DashboardDataRequest: 包含请求信息的结构体，包括起始时间和结束时间等
//
// 返回值：
//
//	*model.DashboardBillboardResponse: 包含 Billboard 数据的结构体
//	error: 错误信息，请求失败时返回非空错误信息
func (c *Client) GetDashboardBillboardData(req *model.DashboardDataRequest) (*model.DashboardBillboardResponse, error) {
	result := &model.DashboardBillboardResponse{}
	url := fmt.Sprintf(BillboardDataPath)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboardTrendSeniorData 获取监控数据趋势图-高级能力
// req: 请求参数，包含时间范围、指标等
// 返回值:
//   - *model.DashboardTrendSeniorResponse: 返回请求结果，包括趋势图数据等
//   - error: 返回请求过程中的错误，如果没有错误则为nil
func (c *Client) GetDashboardTrendSeniorData(req *model.DashboardDataRequest) (*model.DashboardTrendSeniorResponse, error) {
	result := &model.DashboardTrendSeniorResponse{}
	url := fmt.Sprintf(TrendSeniorDataPath)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetDashboardDimensions 是 Client 类型的一个方法，用于获取仪表盘维度信息。
// 参数 req 是请求体，类型为 *model.DashboardDimensionsRequest。
// 返回值类型为 *model.DashboardDimensionsResponse 和 error。
func (c *Client) GetDashboardDimensions(req *model.DashboardDimensionsRequest) (*map[string][]string, error) {
	if len(req.UserID) == 0 {
		return nil, fmt.Errorf("userID is nil")
	}
	if len(req.Service) == 0 {
		return nil, fmt.Errorf("service is nil")
	}
	if len(req.Region) == 0 {
		return nil, fmt.Errorf("region is nil")
	}
	if len(req.MetricName) == 0 {
		return nil, fmt.Errorf("metricname is nil")
	}
	if len(req.ResourceID) == 0 {
		return nil, fmt.Errorf("resourceid is nil")
	}
	result := &map[string][]string{}
	url := fmt.Sprintf(DimensionsPath, req.UserID, req.Service, req.Region)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("showId", req.ResourceID).
		WithQueryParam("metricName", req.MetricName).
		WithQueryParam("dimensions", req.Dimensions).
		WithQueryParam("region", req.Region).
		WithQueryParam("service", req.Service).
		WithQueryParam("userId", req.UserID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckUpdateWidgetRequest(r *model.DashboardWidgetRequest) error {
	if len(r.UserID) == 0 {
		return fmt.Errorf("userID is nil")
	}
	if len(r.Title) == 0 {
		return fmt.Errorf("title is nil")
	}
	if len(r.Type) == 0 {
		return fmt.Errorf("type is nil")
	}
	if len(r.WidgetName) == 0 {
		return fmt.Errorf("widgetName is nil")
	}
	if len(r.DashboardName) == 0 {
		return fmt.Errorf("dashboardName is nil")
	}
	return nil
}

func CheckCreateDashboardRequest(d *model.DashboardRequest) error {
	if len(d.UserID) == 0 {
		return fmt.Errorf("userID is nil")
	}
	if len(d.Configure) == 0 {
		return fmt.Errorf("configure is nil")
	}
	if len(d.Title) == 0 {
		return fmt.Errorf("title is nil")
	}
	if len(d.Type) == 0 {
		return fmt.Errorf("type is nil")
	}
	return nil
}

func CheckWidgetRequest(r *model.DashboardWidgetRequest) error {
	if len(r.UserID) == 0 {
		return fmt.Errorf("userID is nil")
	}
	if len(r.DashboardName) == 0 {
		return fmt.Errorf("dashboardName is nil")
	}
	if len(r.WidgetName) == 0 {
		return fmt.Errorf("widgetName is nil")
	}
	return nil
}

func CheckCreateWidgetRequest(r *model.DashboardWidgetRequest) error {
	if len(r.UserID) == 0 {
		return fmt.Errorf("userID is nil")
	}
	if len(r.DashboardName) == 0 {
		return fmt.Errorf("dashboardName is nil")
	}
	return nil
}

func CheckDashboardRequest(r *model.DashboardRequest) error {
	if len(r.UserID) == 0 {
		return fmt.Errorf("userID is nil")
	}
	if len(r.DashboardName) == 0 {
		return fmt.Errorf("dashboardName is nil")
	}
	return nil
}

func CheckUpdateDashboardRequest(r *model.DashboardRequest) error {
	if len(r.UserID) == 0 {
		return fmt.Errorf("userID is nil")
	}
	if len(r.DashboardName) == 0 {
		return fmt.Errorf("dashboardName is nil")
	}
	if len(r.Configure) == 0 {
		return fmt.Errorf("configure is nil")
	}
	return nil
}

// CreateAlarmPolicy create alarm config
func (c *Client) CreateAlarmPolicy(req *model.AlarmConfig) error {
	if err := checkAlarmConfig(req); err != nil {
		return err
	}

	err := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config/create").
		WithBody(req).
		WithMethod(http.POST).
		Do()
	return err
}

// DeleteAlarmPolicy delete alarm config
func (c *Client) DeleteAlarmPolicy(req *model.CommonAlarmConfigRequest) error {
	if err := checkCommonAlarmConfigRequest(req); err != nil {
		return err
	}

	err := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config/delete").
		WithQueryParam("userId", req.UserId).
		WithQueryParam("scope", req.Scope).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.POST).
		Do()
	return err
}

// UpdateAlarmPolicy update alarm config
func (c *Client) UpdateAlarmPolicy(req *model.AlarmConfig) error {
	if err := checkAlarmConfig(req); err != nil {
		return err
	}
	if len(req.AlarmName) <= 0 {
		return errors.New("alarmName should not be empty")
	}

	err := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config/update").
		WithBody(req).
		WithMethod(http.POST).
		Do()
	return err
}

// BlockAlarmPolicy block alarm config
func (c *Client) BlockAlarmPolicy(req *model.CommonAlarmConfigRequest) error {
	if err := checkCommonAlarmConfigRequest(req); err != nil {
		return err
	}

	err := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config/block").
		WithQueryParam("userId", req.UserId).
		WithQueryParam("scope", req.Scope).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.POST).
		Do()
	return err
}

// UnblockAlarmPolicy unblock alarm config
func (c *Client) UnblockAlarmPolicy(req *model.CommonAlarmConfigRequest) error {
	if err := checkCommonAlarmConfigRequest(req); err != nil {
		return err
	}

	err := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config/unblock").
		WithQueryParam("userId", req.UserId).
		WithQueryParam("scope", req.Scope).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.POST).
		Do()
	return err
}

// GetAlarmPolicyDetail get alarm config detail
func (c *Client) GetAlarmPolicyDetail(req *model.CommonAlarmConfigRequest) (*model.AlarmConfig, error) {
	if err := checkCommonAlarmConfigRequest(req); err != nil {
		return nil, err
	}

	result := &model.AlarmConfig{}
	err := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config").
		WithQueryParam("userId", req.UserId).
		WithQueryParam("scope", req.Scope).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// ListSingleInstanceAlarmConfigs list single instance alarm config list
func (c *Client) ListSingleInstanceAlarmConfigs(req *model.ListSingleInstanceAlarmConfigsRequest) (*model.ListSingleInstanceAlarmConfigsResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.Region) <= 0 {
		req.Region = "bj"
	}
	if len(req.Order) <= 0 {
		req.Order = "desc"
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should be greater than 0")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should be greater than 0")
	}

	result := &model.ListSingleInstanceAlarmConfigsResponse{}
	rb := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config/list").
		WithQueryParamFilter("userId", req.UserId).
		WithQueryParamFilter("scope", req.Scope).
		WithQueryParamFilter("region", req.Region).
		WithQueryParamFilter("dimensions", req.Dimensions).
		WithQueryParamFilter("order", req.Order).
		WithQueryParamFilter("alarmNamePrefix", req.AlarmNamePrefix).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize))
	if req.ActionEnabled != nil {
		rb.WithQueryParam("actionEnabled", strconv.FormatBool(*req.ActionEnabled))
	} else {
		rb.WithQueryParam("actionEnabled", "")
	}
	err := rb.WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// ListAlarmMetrics list alarm metric list
func (c *Client) ListAlarmMetrics(req *model.ListAlarmMetricsRequest) ([]*model.AlarmMetric, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}

	result := make([]*model.AlarmMetric, 0)
	rb := bce.NewRequestBuilder(c).
		WithURL("/csm/api/v1/services/alarm/config/metrics").
		WithQueryParamFilter("userId", req.UserId).
		WithQueryParamFilter("scope", req.Scope).
		WithQueryParamFilter("region", req.Region).
		WithQueryParamFilter("dimensions", req.Dimensions).
		WithQueryParamFilter("type", req.Type).
		WithQueryParamFilter("locale", req.Locale)
	err := rb.WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

func checkAlarmConfig(req *model.AlarmConfig) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return errors.New("scope should not be empty")
	}
	if len(req.AliasName) <= 0 {
		return errors.New("aliasName should not be empty")
	}
	if _, ok := model.AlarmLevelMap[req.Level]; !ok {
		return errors.New("alarmLevel is invalid")
	}
	if len(req.AlarmType) <= 0 {
		return errors.New("type should not be empty")
	}
	if req.MonitorObject == nil {
		return errors.New("monitorObject should not be null")
	}
	if len(req.IncidentActions) <= 0 {
		return errors.New("alarmActions should not be empty")
	}
	if len(req.Rules) <= 0 {
		return errors.New("rules should not be empty")
	}
	return nil
}

func checkCommonAlarmConfigRequest(req *model.CommonAlarmConfigRequest) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return errors.New("scope should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return errors.New("alarmName should not be empty")
	}
	return nil
}

// CreateAlarmPolicyV2 create alarm config v2
func (c *Client) CreateAlarmPolicyV2(req *model.AlarmConfigV2) (*model.CreateAlarmPolicyV2Response, error) {
	if err := checkAlarmConfigV2(req); err != nil {
		return nil, err
	}
	if len(req.ResourceType) <= 0 {
		req.ResourceType = "Instance"
	}

	result := &model.CreateAlarmPolicyV2Response{}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v2/userId/%s/services/%s/alarm/config/create", req.UserId, req.Scope)).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateAlarmPolicyV2 update alarm config v2
func (c *Client) UpdateAlarmPolicyV2(req *model.AlarmConfigV2) error {
	if err := checkAlarmConfigV2(req); err != nil {
		return err
	}
	if len(req.AlarmName) <= 0 {
		return errors.New("alarmName should not be empty")
	}
	if len(req.ResourceType) <= 0 {
		req.ResourceType = "Instance"
	}

	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v2/userId/%s/services/%s/alarm/config/update", req.UserId, req.Scope)).
		WithBody(req).
		WithMethod(http.PUT).
		Do()
	return err
}

// BlockAlarmPolicyV2 block alarm config v2
func (c *Client) BlockAlarmPolicyV2(req *model.CommonAlarmConfigRequest) error {
	if err := checkCommonAlarmConfigRequest(req); err != nil {
		return err
	}

	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v2/userId/%s/services/%s/alarm/config/block", req.UserId, req.Scope)).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.POST).
		Do()
	return err
}

// UnblockAlarmPolicyV2 unblock alarm config v2
func (c *Client) UnblockAlarmPolicyV2(req *model.CommonAlarmConfigRequest) error {
	if err := checkCommonAlarmConfigRequest(req); err != nil {
		return err
	}

	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v2/userId/%s/services/%s/alarm/config/unblock", req.UserId, req.Scope)).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.POST).
		Do()
	return err
}

// GetAlarmPolicyDetailV2 get alarm config detail v2
func (c *Client) GetAlarmPolicyDetailV2(req *model.CommonAlarmConfigRequest) (*model.AlarmConfigV2, error) {
	if err := checkCommonAlarmConfigRequest(req); err != nil {
		return nil, err
	}

	result := &model.AlarmConfigV2{}
	err := bce.NewRequestBuilder(c).
		WithURL(fmt.Sprintf("/csm/api/v2/userId/%s/services/%s/alarm/config", req.UserId, req.Scope)).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func checkAlarmConfigV2(req *model.AlarmConfigV2) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserId) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return errors.New("scope should not be empty")
	}
	if len(req.AliasName) <= 0 {
		return errors.New("aliasName should not be empty")
	}
	if _, ok := model.AlarmLevelMap[req.AlarmLevel]; !ok {
		return errors.New("alarmLevel is invalid")
	}
	if _, ok := model.TargetTypeMap[req.TargetType]; !ok {
		return errors.New("targetType is invalid")
	}
	if req.TargetType == model.TargetTypeInstanceGroup && len(req.TargetInstanceGroups) <= 0 {
		return errors.New("targetInstanceGroups should not be empty")
	}
	if req.TargetType == model.TargetTypeMultiInstances && len(req.TargetInstances) <= 0 {
		return errors.New("targetInstances should not be empty")
	}
	if req.TargetType == model.TargetTypeInstanceTags && len(req.TargetInstanceTags) <= 0 {
		return errors.New("targetInstanceTags should not be empty")
	}
	if len(req.Policies) <= 0 {
		return errors.New("policies should not be empty")
	}
	if len(req.Actions) <= 0 {
		return errors.New("actions should not be empty")
	}
	return nil
}

// CreateSiteHttpTask - create an Http type site task
//
// PARAMS:
//   - args: the arguments to create Http task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of create Http task, contains new task ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSiteHttpTask(req *model.CreateHTTPTask) (*model.CreateTaskResponse, error) {
	checkError := CheckCreateHttpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteCreateHttpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSiteHttpTask - update an Http type site task
//
// PARAMS:
//   - args: the arguments to update Http task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of update Http task, contains task ID
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSiteHttpTask(req *model.CreateHTTPTask) (*model.CreateTaskResponse, error) {
	checkError := CheckUpdateHttpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteUpdateHttpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteHttpTask - get an Http type site task detail
//
// PARAMS:
//   - args: the arguments to get Http task
//
// RETURNS:
//   - *model.GetHttpTask: the result of get Http task detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteHttpTask(req *model.GetTaskDetailRequest) (*model.CreateHTTPTaskResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	result := &model.CreateHTTPTaskResponse{}
	url := fmt.Sprintf(SiteGetHttpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckCreateHttpRequest(req *model.CreateHTTPTask) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Address) <= 0 {
		return errors.New("address should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return errors.New("taskName should not be empty")
	}
	if len(req.Method) <= 0 {
		return errors.New("method should not be empty")
	}
	if req.Cycle <= 0 {
		return errors.New("cycle should not be less than 0")
	}
	if len(req.Idc) <= 0 {
		return errors.New("idc should not be empty")
	}
	if req.Timeout <= 0 {
		return errors.New("timeout should not be less than 0")
	}
	return nil
}

func CheckUpdateHttpRequest(req *model.CreateHTTPTask) error {
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	return CheckCreateHttpRequest(req)
}

// CreateSiteHttpsTask - create an Https type site task
//
// PARAMS:
//   - args: the arguments to create Https task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of create https task, contains new task ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSiteHttpsTask(req *model.CreateHTTPSTask) (*model.CreateTaskResponse, error) {
	checkError := CheckCreateHttpsRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteCreateHttpsTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSiteHttpsTask - update an Https type site task
//
// PARAMS:
//   - args: the arguments to update Https task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of update https task, contains task ID
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSiteHttpsTask(req *model.CreateHTTPSTask) (*model.CreateTaskResponse, error) {
	checkError := CheckUpdateHttpsRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteUpdateHttpsTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteHttpsTask - get an Https type site task detail
//
// PARAMS:
//   - args: the arguments to get Https task
//
// RETURNS:
//   - *model.GetHttpsTask: the result of get https task detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteHttpsTask(req *model.GetTaskDetailRequest) (*model.CreateHTTPSTaskResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	result := &model.CreateHTTPSTaskResponse{}
	url := fmt.Sprintf(SiteGetHttpsTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckCreateHttpsRequest(req *model.CreateHTTPSTask) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Address) <= 0 {
		return errors.New("address should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return errors.New("taskName should not be empty")
	}
	if len(req.Method) <= 0 {
		return errors.New("method should not be empty")
	}
	if req.Cycle <= 0 {
		return errors.New("cycle should not be less than 0")
	}
	if len(req.Idc) <= 0 {
		return errors.New("idc should not be empty")
	}
	if req.Timeout <= 0 {
		return errors.New("timeout should not be less than 0")
	}
	if len(req.IPType) <= 0 {
		return errors.New("ipType should not be empty")
	}
	return nil
}

func CheckUpdateHttpsRequest(req *model.CreateHTTPSTask) error {
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	return CheckCreateHttpsRequest(req)
}

// CreateSitePingTask - create an Ping type site task
//
// PARAMS:
//   - args: the arguments to create Ping task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of create ping task, contains new task ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSitePingTask(req *model.CreatePingTask) (*model.CreateTaskResponse, error) {
	checkError := CheckCreatePingRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteCreatePingTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSitePingTask - update an Ping type site task
//
// PARAMS:
//   - args: the arguments to update Ping task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of update ping task, contains task ID
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSitePingTask(req *model.CreatePingTask) (*model.CreateTaskResponse, error) {
	checkError := CheckUpdatePingRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteUpdatePingTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetSitePingTask - get an Ping type site task detail
//
// PARAMS:
//   - args: the arguments to get Ping task
//
// RETURNS:
//   - *model.GetPingTask: the result of get ping task detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSitePingTask(req *model.GetTaskDetailRequest) (*model.CreatePingTask, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	result := &model.CreatePingTask{}
	url := fmt.Sprintf(SiteGetPingTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckCreatePingRequest(req *model.CreatePingTask) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Address) <= 0 {
		return errors.New("address should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return errors.New("taskName should not be empty")
	}
	if req.Cycle <= 0 {
		return errors.New("cycle should not be less than 0")
	}
	if len(req.Idc) <= 0 {
		return errors.New("idc should not be empty")
	}
	if req.Timeout <= 0 {
		return errors.New("timeout should not be less than 0")
	}
	if len(req.IPType) <= 0 {
		return errors.New("ipType should not be empty")
	}
	return nil
}

func CheckUpdatePingRequest(req *model.CreatePingTask) error {
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	return CheckCreatePingRequest(req)
}

// CreateSiteTcpTask - create an Tcp type site task
//
// PARAMS:
//   - args: the arguments to create Tcp task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of create Tcp task, contains new task ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSiteTcpTask(req *model.CreateTCPTask) (*model.CreateTaskResponse, error) {
	checkError := CheckCreateTcpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteCreateTcpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSiteTcpTask - update an Tcp type site task
//
// PARAMS:
//   - args: the arguments to update Tcp task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of update Tcp task, contains task ID
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSiteTcpTask(req *model.CreateTCPTask) (*model.CreateTaskResponse, error) {
	checkError := CheckUpdateTcpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteUpdateTcpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteTcpTask - get an tcp type site task detail
//
// PARAMS:
//   - args: the arguments to get tcp task
//
// RETURNS:
//   - *model.GetTcpTask: the result of get tcp task detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteTcpTask(req *model.GetTaskDetailRequest) (*model.CreateTCPTask, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	result := &model.CreateTCPTask{}
	url := fmt.Sprintf(SiteGetTcpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckCreateTcpRequest(req *model.CreateTCPTask) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Address) <= 0 {
		return errors.New("address should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return errors.New("taskName should not be empty")
	}
	if req.Port <= 0 {
		return errors.New("port should not be less than 0")
	}
	if req.Cycle <= 0 {
		return errors.New("cycle should not be less than 0")
	}
	if len(req.Idc) <= 0 {
		return errors.New("idc should not be empty")
	}
	if req.Timeout <= 0 {
		return errors.New("timeout should not be less than 0")
	}
	if len(req.IPType) <= 0 {
		return errors.New("ipType should not be empty")
	}
	return nil
}

func CheckUpdateTcpRequest(req *model.CreateTCPTask) error {
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	return CheckCreateTcpRequest(req)
}

// CreateSiteUdpTask - create an Udp type site task
//
// PARAMS:
//   - args: the arguments to create Udp task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of create Udp task, contains new task ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSiteUdpTask(req *model.CreateUDPTask) (*model.CreateTaskResponse, error) {
	checkError := CheckCreateUdpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteCreateUdpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSiteUdpTask - update an Udp type site task
//
// PARAMS:
//   - args: the arguments to update Udp task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of update Udp task, contains task ID
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSiteUdpTask(req *model.CreateUDPTask) (*model.CreateTaskResponse, error) {
	checkError := CheckUpdateUdpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteUpdateUdpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteUdpTask - get an Udp type site task detail
//
// PARAMS:
//   - args: the arguments to get Udp task
//
// RETURNS:
//   - *model.GetUdpTask: the result of get Udp task detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteUdpTask(req *model.GetTaskDetailRequest) (*model.CreateUDPTask, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	result := &model.CreateUDPTask{}
	url := fmt.Sprintf(SiteGetUdpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckCreateUdpRequest(req *model.CreateUDPTask) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Address) <= 0 {
		return errors.New("address should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return errors.New("taskName should not be empty")
	}
	if req.Port <= 0 {
		return errors.New("port should not be less than 0")
	}
	if req.Cycle <= 0 {
		return errors.New("cycle should not be less than 0")
	}
	if len(req.Idc) <= 0 {
		return errors.New("idc should not be empty")
	}
	if req.Timeout <= 0 {
		return errors.New("timeout should not be less than 0")
	}
	if len(req.IPType) <= 0 {
		return errors.New("ipType should not be empty")
	}
	return nil
}

func CheckUpdateUdpRequest(req *model.CreateUDPTask) error {
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	return CheckCreateUdpRequest(req)
}

// CreateSiteFtpTask - create an Ftp type site task
//
// PARAMS:
//   - args: the arguments to create Ftp task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of create Ftp task, contains new task ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSiteFtpTask(req *model.CreateFtpTask) (*model.CreateTaskResponse, error) {
	checkError := CheckCreateFtpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteCreateFtpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSiteFtpTask - update an Ftp type site task
//
// PARAMS:
//   - args: the arguments to update Ftp task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of update Ftp task, contains task ID
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSiteFtpTask(req *model.CreateFtpTask) (*model.CreateTaskResponse, error) {
	checkError := CheckUpdateFtpRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteUpdateFtpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteFtpTask - get an Ftp type site task detail
//
// PARAMS:
//   - args: the arguments to get Ftp task
//
// RETURNS:
//   - *model.GetFtpTask: the result of get Ftp task detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteFtpTask(req *model.GetTaskDetailRequest) (*model.CreateFtpTask, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	result := &model.CreateFtpTask{}
	url := fmt.Sprintf(SiteGetFtpTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckCreateFtpRequest(req *model.CreateFtpTask) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Address) <= 0 {
		return errors.New("address should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return errors.New("taskName should not be empty")
	}
	if req.Port <= 0 {
		return errors.New("port should not be less than 0")
	}
	if req.Cycle <= 0 {
		return errors.New("cycle should not be less than 0")
	}
	if len(req.Idc) <= 0 {
		return errors.New("idc should not be empty")
	}
	if req.Timeout <= 0 {
		return errors.New("timeout should not be less than 0")
	}
	if len(req.IPType) <= 0 {
		return errors.New("ipType should not be empty")
	}
	return nil
}

func CheckUpdateFtpRequest(req *model.CreateFtpTask) error {
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	return CheckCreateFtpRequest(req)
}

// CreateSiteDnsTask - create an Dns type site task
//
// PARAMS:
//   - args: the arguments to create Dns task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of create Dns task, contains new task ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSiteDnsTask(req *model.CreateDNSTask) (*model.CreateTaskResponse, error) {
	checkError := CheckCreateDnsRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteCreateDnsTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSiteDnsTask - update an Dns type site task
//
// PARAMS:
//   - args: the arguments to update Dns task
//
// RETURNS:
//   - *model.CreateTaskResponse: the result of update Dns task, contains task ID
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSiteDnsTask(req *model.CreateDNSTask) (*model.CreateTaskResponse, error) {
	checkError := CheckUpdateDnsRequest(req)
	if checkError != nil {
		return nil, checkError
	}
	result := &model.CreateTaskResponse{}
	url := fmt.Sprintf(SiteUpdateDnsTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteDnsTask - get an Dns type site task detail
//
// PARAMS:
//   - args: the arguments to get Dns task
//
// RETURNS:
//   - *model.GetDnsTask: the result of get Dns task detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteDnsTask(req *model.GetTaskDetailRequest) (*model.CreateDNSTask, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	result := &model.CreateDNSTask{}
	url := fmt.Sprintf(SiteGetDnsTaskPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func CheckCreateDnsRequest(req *model.CreateDNSTask) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.Address) <= 0 {
		return errors.New("address should not be empty")
	}
	if len(req.TaskName) <= 0 {
		return errors.New("taskName should not be empty")
	}
	if req.Cycle <= 0 {
		return errors.New("cycle should not be less than 0")
	}
	if len(req.Idc) <= 0 {
		return errors.New("idc should not be empty")
	}
	if req.Timeout <= 0 {
		return errors.New("timeout should not be less than 0")
	}
	if len(req.IPType) <= 0 {
		return errors.New("ipType should not be empty")
	}
	return nil
}

func CheckUpdateDnsRequest(req *model.CreateDNSTask) error {
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	return CheckCreateDnsRequest(req)
}

// GetSiteTaskList - get site task list
//
// PARAMS:
//   - args: the arguments to get task list
//
// RETURNS:
//   - *model.GetDnsTask: the result of get task list
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteTaskList(req *model.GetTaskListRequest) (*model.GetTaskListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be less than 0")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be less than 0")
	}
	if len(req.Query) <= 0 {
		req.Query = "NAME:"
	}
	url := fmt.Sprintf(SiteGetTaskListPath, req.UserID)
	result := &model.GetTaskListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithQueryParam("query", req.Query).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// DeleteSiteTask - delete site task
//
// PARAMS:
//   - args: the arguments to delete task
//
// RETURNS:
//   - *model.DeleteTaskResponse: the result of delete task
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSiteTask(req *model.GetTaskDetailRequest) (*model.DeleteTaskResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	url := fmt.Sprintf(SiteDeleteTaskPath, req.UserID)
	result := &model.DeleteTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.DELETE).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteTaskDetail - get site task detail
//
// PARAMS:
//   - args: the arguments to get task detail
//
// RETURNS:
//   - *model.GetTaskDetailResponse: the result of get task
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteTaskDetail(req *model.GetTaskDetailRequest) (*model.GetTaskDetailResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}
	url := fmt.Sprintf(SiteGetTaskDetailPath, req.UserID, req.TaskID)
	result := &model.GetTaskDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// CreateSiteAlarmConfig - create site alarm config
//
// PARAMS:
//   - args: the arguments to create site alarm config
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSiteAlarmConfig(req *model.CreateSiteAlarmConfigRequest) error {
	checkError := checkCreateSiteAlarmConfig(req)
	if checkError != nil {
		return checkError
	}
	url := fmt.Sprintf(SiteCreateAlarmConfigPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		Do()
	return err
}

// UpdateSiteAlarmConfig - update site alarm config
//
// PARAMS:
//   - args: the arguments to update site alarm config
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSiteAlarmConfig(req *model.CreateSiteAlarmConfigRequest) error {
	checkError := checkUpdateSiteAlarmConfig(req)
	if checkError != nil {
		return checkError
	}
	url := fmt.Sprintf(SiteUpdateAlarmConfigPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PUT).
		Do()
	return err
}

func checkCreateSiteAlarmConfig(req *model.CreateSiteAlarmConfigRequest) error {
	if req == nil {
		return errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return errors.New("taskId should not be empty")
	}
	if len(req.AliasName) <= 0 {
		return errors.New("aliasName should not be empty")
	}
	if len(req.Namespace) <= 0 {
		return errors.New("namespace should not be empty")
	}
	if req.Level != model.LevelNotice && req.Level != model.LevelWarning && req.Level != model.LevelCritical &&
		req.Level != model.LevelMajor && req.Level != model.LevelCustom {
		return errors.New("level must be one of NOTICE,WARNING,CRITICAL,MAJOR,CUSTOM")
	}
	if len(req.Rules) <= 0 {
		return errors.New("rules should not be empty")
	}

	return nil
}

func checkUpdateSiteAlarmConfig(req *model.CreateSiteAlarmConfigRequest) error {
	if len(req.AlarmName) <= 0 {
		return errors.New("incidentAction should not be empty")
	}
	return checkCreateSiteAlarmConfig(req)
}

// DeleteSiteAlarmConfig - delete site alarm config
//
// PARAMS:
//   - args: the arguments to delete site alarm config
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSiteAlarmConfig(req *model.DeleteSiteAlarmConfigRequest) error {
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	url := fmt.Sprintf(SiteDeleteAlarmConfigPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.DELETE).
		Do()
	return err
}

// GetSiteAlarmConfigDetail - delete site alarm config detail
//
// PARAMS:
//   - args: the arguments to get site alarm config detail
//
// RETURNS:
//   - *model.CreateSiteAlarmConfigRequest: the result of get site alarm config detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteAlarmConfigDetail(req *model.GetSiteAlarmConfigRequest) (*model.CreateSiteAlarmConfigResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return nil, errors.New("alarmName should not be empty")
	}
	url := fmt.Sprintf(SiteGetAlarmConfigDetailPath, req.UserID)
	result := &model.CreateSiteAlarmConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("alarmName", req.AlarmName).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteAlarmConfigList - delete site alarm config list
//
// PARAMS:
//   - args: the arguments to get site alarm config list
//
// RETURNS:
//   - *model.CreateSiteAlarmConfigRequest: the result of get site alarm config list
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteAlarmConfigList(req *model.GetSiteAlarmConfigListRequest) (*model.GetSiteAlarmConfigListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be empty")
	}
	url := fmt.Sprintf(SiteGetAlarmConfigListPath, req.UserID)
	result := &model.GetSiteAlarmConfigListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithQueryParam("alarmName", req.AliasName).
		WithQueryParam("actionEnabled", strconv.FormatBool(req.ActionEnabled)).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// BlockSiteAlarmConfig - block site alarm config
//
// PARAMS:
//   - args: the arguments to block site alarm config
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) BlockSiteAlarmConfig(req *model.GetSiteAlarmConfigRequest) error {
	if req == nil {
		return errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return errors.New("alarmName should not be empty")
	}
	url := fmt.Sprintf(SiteAlarmBlockPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("alarmName", req.AlarmName).
		WithQueryParam("namespace", req.Namespace).
		WithMethod(http.POST).
		Do()
	return err
}

// UnBlockSiteAlarmConfig - block site alarm config
//
// PARAMS:
//   - args: the arguments to unblock site alarm config
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UnBlockSiteAlarmConfig(req *model.GetSiteAlarmConfigRequest) error {
	if req == nil {
		return errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return errors.New("userId should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return errors.New("alarmName should not be empty")
	}
	url := fmt.Sprintf(SiteAlarmUnBlockPath, req.UserID)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("alarmName", req.AlarmName).
		WithQueryParam("namespace", req.Namespace).
		WithMethod(http.POST).
		Do()
	return err
}

// GetTaskByAlarmName - get site alarm config by alarmName
//
// PARAMS:
//   - args: the arguments to get site alarm config by alarmName
//
// RETURNS:
//   - *model.GetTaskDetailResponse
//   - error: nil if success otherwise the specific error
func (c *Client) GetTaskByAlarmName(req *model.GetSiteAlarmConfigRequest) (*model.GetTaskDetailResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.AlarmName) <= 0 {
		return nil, errors.New("alarmName should not be empty")
	}
	url := fmt.Sprintf(SiteGetTaskByAlarmNamePath, req.UserID, req.AlarmName)
	result := &model.GetTaskDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetSiteMetricData - get site metric data
//
// PARAMS:
//   - args: the arguments to get site metric data
//
// RETURNS:
//   - *model.GetSiteMetricDataResponse
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteMetricData(req *model.GetSiteMetricDataRequest) ([]*model.GetSiteMetricDataResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.MetricName) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}
	if len(req.Statistics) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if req.Cycle <= 0 {
		return nil, errors.New("cycle should not be less than 0")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}

	url := fmt.Sprintf(SiteGetMetricDataPath, req.UserID)
	result := make([]*model.GetSiteMetricDataResponse, 0)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithQueryParam("metricName", req.MetricName).
		WithQueryParam("dimensions", req.Dimensions).
		WithQueryParam("statistics", strings.Join(req.Statistics, ",")).
		WithQueryParam("cycle", strconv.Itoa(req.Cycle)).
		WithQueryParam("startTime", req.StartTime).
		WithQueryParam("endTime", req.EndTime).
		WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

// GetSiteOverallView - get site over all view
//
// PARAMS:
//   - args: the arguments to get site over all view
//
// RETURNS:
//   - *model.GetSiteViewResponse
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteOverallView(req *model.GetTaskDetailRequest) ([]*model.GetSiteViewResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}

	url := fmt.Sprintf(SiteGetOverallViewPath, req.UserID)
	result := make([]*model.GetSiteViewResponse, 0)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

// GetSiteProvincialView - get site provincial view
//
// PARAMS:
//   - args: the arguments to get site provincial view
//
// RETURNS:
//   - *model.GetSiteViewResponse
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteProvincialView(req *model.GetTaskDetailRequest) ([]*model.GetSiteViewResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}

	url := fmt.Sprintf(SiteGetProvincialViewPath, req.UserID)
	result := make([]*model.GetSiteViewResponse, 0)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithQueryParam("isp", req.Isp).
		WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

// GetSiteAgentList - get site agent list
//
// PARAMS:
//   - args: the arguments to get site agent list
//
// RETURNS:
//   - *model.GetSiteViewResponse
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteAgentList(req *model.GetSiteAgentListRequest) ([]*model.GetSiteAgentListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}

	url := fmt.Sprintf(SiteAgentListPath, req.UserID)
	result := make([]*model.GetSiteAgentListResponse, 0)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithResult(&result).
		Do()
	return result, err
}

// GetSiteAgentListByTaskId - get site agent list by taskId
//
// PARAMS:
//   - args: the arguments to get site agent list by taskId
//
// RETURNS:
//   - *model.GetSiteAgentByTaskIDResponse
//   - error: nil if success otherwise the specific error
func (c *Client) GetSiteAgentListByTaskId(req *model.GetTaskDetailRequest) (*model.GetSiteAgentByTaskIDResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.TaskID) <= 0 {
		return nil, errors.New("taskId should not be empty")
	}

	url := fmt.Sprintf(SiteGetAgentByTaskIdPath, req.UserID)
	result := &model.GetSiteAgentByTaskIDResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("taskId", req.TaskID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetCloudEventData - get cloud event data
func (c *Client) GetCloudEventData(req *model.EventDataRequest) (*model.CloudEventResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be empty")
	}
	if len(req.AccountID) <= 0 {
		return nil, errors.New("accountId should not be empty")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if !isUtcTime(req.StartTime) {
		return nil, errors.New("startTime should be utc time")
	}
	if !isUtcTime(req.EndTime) {
		return nil, errors.New("endTime should be utc time")
	}

	url := fmt.Sprintf(EventCloudListPath)
	result := &model.CloudEventResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithQueryParam("startTime", req.StartTime).
		WithQueryParam("endTime", req.EndTime).
		WithQueryParam("accountId", req.AccountID).
		WithQueryParam("ascending", strconv.FormatBool(req.Ascending)).
		WithQueryParam("scope", req.Scope).
		WithQueryParam("region", req.Region).
		WithQueryParam("eventLevel", req.EventLevel).
		WithQueryParam("eventName", req.EventName).
		WithQueryParam("eventAlias", req.EventAlias).
		WithQueryParam("resourceType", req.ResourceType).
		WithQueryParam("resourceId", req.ResourceID).
		WithQueryParam("eventId", req.EventID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetPlatformEventData - get platform event data
func (c *Client) GetPlatformEventData(req *model.EventDataRequest) (*model.PlatformEventResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be empty")
	}
	if len(req.AccountID) <= 0 {
		return nil, errors.New("accountId should not be empty")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if !isUtcTime(req.StartTime) {
		return nil, errors.New("startTime should be utc time")
	}
	if !isUtcTime(req.EndTime) {
		return nil, errors.New("endTime should be utc time")
	}

	url := fmt.Sprintf(EventPlatformListPath)
	result := &model.PlatformEventResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithQueryParam("startTime", req.StartTime).
		WithQueryParam("endTime", req.EndTime).
		WithQueryParam("accountId", req.AccountID).
		WithQueryParam("ascending", strconv.FormatBool(req.Ascending)).
		WithQueryParam("region", req.Region).
		WithQueryParam("eventLevel", req.EventLevel).
		WithQueryParam("eventName", req.EventName).
		WithQueryParam("eventAlias", req.EventAlias).
		WithQueryParam("eventId", req.EventID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// CreateEventPolicy - create event policy
func (c *Client) CreateEventPolicy(req *model.EventPolicy) error {
	if req == nil {
		return errors.New("req should not be empty")
	}
	if len(req.AccountID) <= 0 {
		return errors.New("accountId should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return errors.New("serviceName should not be empty")
	}
	if len(req.Name) <= 0 {
		return errors.New("name should not be empty")
	}
	if len(req.BlockStatus) <= 0 {
		return errors.New("blockStatus should not be empty")
	}

	if len(req.EventFilter.EventLevel) <= 0 {
		return errors.New("eventLevel should not be empty")
	}
	if len(req.EventFilter.EventTypeList) <= 0 {
		return errors.New("eventTypeList should not be empty")
	}

	if len(req.Resource.Region) <= 0 {
		return errors.New("resource.region should not be empty")
	}
	if len(req.Resource.Type) <= 0 {
		return errors.New("resource.type should not be empty")
	}

	if len(req.IncidentActions) <= 0 {
		return errors.New("incidentActions should not be empty")
	}

	url := fmt.Sprintf(EventPolicyPath, req.AccountID, req.ServiceName)
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		Do()
	return err
}

// CreateInstanceGroup - create instance group
func (c *Client) CreateInstanceGroup(req *model.MergedGroup) (*model.InstanceGroup, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return nil, errors.New("serviceName should not be empty")
	}
	if len(req.TypeName) <= 0 {
		return nil, errors.New("typeName should not be empty")
	}
	if len(req.Name) <= 0 {
		return nil, errors.New("name should not be empty")
	}

	url := fmt.Sprintf(InstanceGroupPath, req.UserId)
	result := &model.InstanceGroup{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateInstanceGroup - update instance group
func (c *Client) UpdateInstanceGroup(req *model.InstanceGroup) (*model.InstanceGroup, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if req.ID <= 0 {
		return nil, errors.New("id should not be empty")
	}
	if len(req.Name) <= 0 {
		return nil, errors.New("name should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return nil, errors.New("serviceName should not be empty")
	}
	if len(req.TypeName) <= 0 {
		return nil, errors.New("typeName should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}

	url := fmt.Sprintf(InstanceGroupPath, req.UserID)
	result := &model.InstanceGroup{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.PATCH).
		WithResult(result).
		Do()
	return result, err
}

// DeleteInstanceGroup - delete instance group
func (c *Client) DeleteInstanceGroup(req *model.InstanceGroupBase) (*model.InstanceGroup, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.ID) <= 0 {
		return nil, errors.New("id should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}

	url := fmt.Sprintf(InstanceGroupIdPath, req.UserID, req.ID)
	result := &model.InstanceGroup{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.DELETE).
		WithResult(result).
		Do()
	return result, err
}

// GetInstanceGroup - get instance group
func (c *Client) GetInstanceGroup(req *model.InstanceGroupBase) (*model.InstanceGroup, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.ID) <= 0 {
		return nil, errors.New("id should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}

	url := fmt.Sprintf(InstanceGroupIdPath, req.UserID, req.ID)
	result := &model.InstanceGroup{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetInstanceGroupList - get instance group list
func (c *Client) GetInstanceGroupList(req *model.InstanceGroupQuery) (*model.InstanceGroupListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be empty")
	}

	url := fmt.Sprintf(InstanceGroupListPath, req.UserID)
	result := &model.InstanceGroupListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("name", req.Name).
		WithQueryParam("serviceName", req.ServiceName).
		WithQueryParam("region", req.Region).
		WithQueryParam("typeName", req.TypeName).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// AddInstanceToInstanceGroup - add instance to instance group
func (c *Client) AddInstanceToInstanceGroup(req *model.MergedGroup) (*model.InstanceGroup, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.ID) <= 0 {
		return nil, errors.New("id should not be empty")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return nil, errors.New("serviceName should not be empty")
	}
	if len(req.TypeName) <= 0 {
		return nil, errors.New("typeName should not be empty")
	}
	if len(req.Name) <= 0 {
		return nil, errors.New("name should not be empty")
	}

	url := fmt.Sprintf(IG_INSTANCE_ADD_PATH, req.UserId, req.ID)
	result := &model.InstanceGroup{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// RemoveInstanceFromInstanceGroup - remove instance from instance group
func (c *Client) RemoveInstanceFromInstanceGroup(req *model.MergedGroup) (*model.InstanceGroup, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.ID) <= 0 {
		return nil, errors.New("id should not be empty")
	}
	if len(req.UserId) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return nil, errors.New("serviceName should not be empty")
	}
	if len(req.TypeName) <= 0 {
		return nil, errors.New("typeName should not be empty")
	}
	if len(req.Name) <= 0 {
		return nil, errors.New("name should not be empty")
	}

	url := fmt.Sprintf(IG_INSTANCE_REMOVE_PATH, req.UserId, req.ID)
	result := &model.InstanceGroup{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetInstanceGroupInstanceList - get instance group instance list
func (c *Client) GetInstanceGroupInstanceList(req *model.IGInstanceQuery) (*model.IGInstanceListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return nil, errors.New("serviceName should not be empty")
	}
	if len(req.TypeName) <= 0 {
		return nil, errors.New("typeName should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}
	if len(req.ViewType) <= 0 {
		return nil, errors.New("viewType should not be empty")
	}
	if len(req.ID) <= 0 {
		return nil, errors.New("id should not be empty")
	}

	url := fmt.Sprintf(IG_INSTANCE_LIST_PATH, req.UserID)
	result := &model.IGInstanceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("serviceName", req.ServiceName).
		WithQueryParam("typeName", req.TypeName).
		WithQueryParam("region", req.Region).
		WithQueryParam("viewType", req.ViewType).
		WithQueryParam("id", req.ID).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetAllInstanceForInstanceGroup - get all instance for instance group
func (c *Client) GetAllInstanceForInstanceGroup(req *model.IGInstanceQuery) (*model.IGInstanceListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return nil, errors.New("serviceName should not be empty")
	}
	if len(req.TypeName) <= 0 {
		return nil, errors.New("typeName should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}
	if len(req.ViewType) <= 0 {
		return nil, errors.New("viewType should not be empty")
	}

	url := fmt.Sprintf(IG_QUERY_INSTANCE_LIST_PATH, req.UserID)
	result := &model.IGInstanceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("serviceName", req.ServiceName).
		WithQueryParam("typeName", req.TypeName).
		WithQueryParam("region", req.Region).
		WithQueryParam("viewType", req.ViewType).
		WithQueryParam("keywordType", req.KeywordType).
		WithQueryParam("keyword", req.Keyword).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetFilterInstanceForInstanceGroup - get filter instance for instance group
func (c *Client) GetFilterInstanceForInstanceGroup(req *model.IGInstanceQuery) (*model.IGInstanceListResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should not be empty")
	}
	if req.PageSize <= 0 {
		return nil, errors.New("pageSize should not be empty")
	}
	if len(req.ServiceName) <= 0 {
		return nil, errors.New("serviceName should not be empty")
	}
	if len(req.TypeName) <= 0 {
		return nil, errors.New("typeName should not be empty")
	}
	if len(req.Region) <= 0 {
		return nil, errors.New("region should not be empty")
	}
	if len(req.ViewType) <= 0 {
		return nil, errors.New("viewType should not be empty")
	}
	if len(req.ID) <= 0 {
		return nil, errors.New("id should not be empty")
	}
	if len(req.UUID) <= 0 {
		return nil, errors.New("uuid should not be empty")
	}

	url := fmt.Sprintf(IG_QUERY_INSTANCE_LIST_FILTER_PATH, req.UserID)
	result := &model.IGInstanceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithQueryParam("serviceName", req.ServiceName).
		WithQueryParam("typeName", req.TypeName).
		WithQueryParam("region", req.Region).
		WithQueryParam("viewType", req.ViewType).
		WithQueryParam("keywordType", req.KeywordType).
		WithQueryParam("keyword", req.Keyword).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithQueryParam("id", req.ID).
		WithQueryParam("uuid", req.UUID).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// GetMultiDimensionLatestMetrics Get Multi-Dimension latest metrics
func (c *Client) GetMultiDimensionLatestMetrics(req *model.MultiDimensionalLatestMetricsRequest) (*model.MultiDimensionalMetricsResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.MetricNames) <= 0 {
		return nil, errors.New("metricNames should not be empty")
	}
	if len(req.Dimensions) > DimensionNumberLimit {
		return nil, errors.New("dimensions should not be more than 100")
	}
	url := fmt.Sprintf(MultiDimensionLatestMetricsPath, req.UserID, req.Scope)
	result := &model.MultiDimensionalMetricsResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetMetricsByPartialDimensions Get metrics according to partial dimensions
func (c *Client) GetMetricsByPartialDimensions(req *model.MetricsByPartialDimensionsRequest) (*model.MetricsByPartialDimensionsPageResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.MetricName) <= 0 {
		return nil, errors.New("metricName should not be empty")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if len(req.Statistics) <= 0 {
		return nil, errors.New("statistics should not be empty")
	}
	if len(req.Dimensions) > DimensionNumberLimit {
		return nil, errors.New("dimensions should not be more than 100")
	}
	url := fmt.Sprintf(MetricsByPartialDimensionsPath, req.UserID, req.Scope)

	result := &model.MetricsByPartialDimensionsPageResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(url).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetMetricsAllDataV2 metric all data
func (c *Client) GetMetricsAllDataV2(req *model.TsdbMetricAllDataQueryRequest) (*model.MultiDimensionalMetricsResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.MetricNames) <= 0 {
		return nil, errors.New("metricNames should not be empty")
	}
	if len(req.Dimensions) <= 0 {
		return nil, errors.New("dimensions should not be empty")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if len(req.Dimensions) > DimensionNumberLimit {
		return nil, errors.New("dimensions should not be more than 100")
	}
	result := &model.MultiDimensionalMetricsResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(MultiMetricAllDataPath).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// BatchGetMetricsAllDataV2 metric all data
func (c *Client) BatchGetMetricsAllDataV2(req *model.TsdbMetricAllDataQueryRequest) (*model.MultiDimensionalMetricsResponse, error) {
	if req == nil {
		return nil, errors.New("req should not be empty")
	}
	if len(req.UserID) <= 0 {
		return nil, errors.New("userId should not be empty")
	}
	if len(req.Scope) <= 0 {
		return nil, errors.New("scope should not be empty")
	}
	if len(req.MetricNames) <= 0 {
		return nil, errors.New("metricNames should not be empty")
	}
	if len(req.Dimensions) <= 0 {
		return nil, errors.New("dimensions should not be empty")
	}
	if len(req.Dimensions) > DimensionNumberLimit {
		return nil, errors.New("dimensions should not be more than 100")
	}
	if len(req.StartTime) <= 0 {
		return nil, errors.New("startTime should not be empty")
	}
	if len(req.EndTime) <= 0 {
		return nil, errors.New("endTime should not be empty")
	}
	if len(req.Type) == 0 {
		req.Type = "Instance"
	}
	if req.Cycle <= 0 {
		req.Cycle = 60
	}
	result := &model.MultiDimensionalMetricsResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL(MultiMetricAllDataBatchPath).
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// 判断传入字符串是否为utc时间，格式为：2023-12-12T00:00:00Z
func isUtcTime(str string) bool {
	_, err := time.Parse(time.RFC3339, str)
	return err == nil
}
