package api

import (
	"errors"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// ListWeeklyOverview lists the diagnosis risk items found within the last 7 days.
func ListWeeklyOverview(cli bce.Client, region string, request *ListWeeklyOverviewRequest) (*ListWeeklyOverviewResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list weekly overview request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list weekly overview request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_OVERVIEW, URI_DIAGNOSIS_OVERVIEW_WEEKLY)

	result := &ListWeeklyOverviewResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetManualDiagnosisCount gets the number of manual diagnoses that can be or have been created.
func GetManualDiagnosisCount(cli bce.Client, region string, request *GetManualDiagnosisCountRequest) (*GetManualDiagnosisCountResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get manual diagnosis count request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get manual diagnosis count request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_MANUAL, URI_DIAGNOSIS_MANUAL_COUNT)

	result := &GetManualDiagnosisCountResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetManualDiagnosisConfig gets the manual diagnosis configuration of a cluster.
func GetManualDiagnosisConfig(cli bce.Client, region string, request *GetManualDiagnosisConfigRequest) (*GetManualDiagnosisConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get manual diagnosis config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get manual diagnosis config request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_MANUAL)

	result := &GetManualDiagnosisConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateManualDiagnosisConfig updates the manual diagnosis configuration of a cluster.
func UpdateManualDiagnosisConfig(cli bce.Client, region string, request *UpdateManualDiagnosisConfigRequest) (*UpdateManualDiagnosisConfigResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update manual diagnosis config request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update manual diagnosis config request clusterId should not be empty")
	}
	if request.Items == nil {
		return nil, errors.New("update manual diagnosis config request items should not be nil")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_MANUAL)

	result := &UpdateManualDiagnosisConfigResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateManualDiagnosis creates a manual diagnosis task on a cluster.
func CreateManualDiagnosis(cli bce.Client, region string, request *CreateManualDiagnosisRequest) (*CreateManualDiagnosisResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create manual diagnosis request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create manual diagnosis request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_MANUAL)

	result := &CreateManualDiagnosisResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListDiagnosisReports lists the diagnosis reports of a cluster.
func ListDiagnosisReports(cli bce.Client, region string, request *ListDiagnosisReportsRequest) (*ListDiagnosisReportsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list diagnosis reports request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list diagnosis reports request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_REPORTS)
	query := map[string]string{}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}

	result := &ListDiagnosisReportsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDiagnosisReport gets the detail of a diagnosis report.
func GetDiagnosisReport(cli bce.Client, region string, request *GetDiagnosisReportRequest) (*GetDiagnosisReportResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get diagnosis report request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get diagnosis report request clusterId should not be empty")
	}
	if request.ReportId == "" {
		return nil, errors.New("get diagnosis report request reportId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_REPORTS, request.ReportId)

	result := &GetDiagnosisReportResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDiagnosisBusyStatus gets whether the diagnosis service is busy for a cluster.
func GetDiagnosisBusyStatus(cli bce.Client, region string, request *GetDiagnosisBusyStatusRequest) (*GetDiagnosisBusyStatusResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get diagnosis busy status request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get diagnosis busy status request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_BUSY)

	result := &GetDiagnosisBusyStatusResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDiagnosisAuthorizationStatus gets whether diagnosis is authorized for a cluster.
func GetDiagnosisAuthorizationStatus(cli bce.Client, region string, request *GetDiagnosisAuthorizationStatusRequest) (*GetDiagnosisAuthorizationStatusResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get diagnosis authorization status request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get diagnosis authorization status request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_AUTHORIZE)

	result := &GetDiagnosisAuthorizationStatusResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// AuthorizeDiagnosis authorizes diagnosis for a cluster.
func AuthorizeDiagnosis(cli bce.Client, region string, request *AuthorizeDiagnosisRequest) (*AuthorizeDiagnosisResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("authorize diagnosis request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("authorize diagnosis request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_AUTHORIZE)

	result := &AuthorizeDiagnosisResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListDiagnosisItems lists the diagnosis items supported by the service.
func ListDiagnosisItems(cli bce.Client, region string, request *ListDiagnosisItemsRequest) (*ListDiagnosisItemsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list diagnosis items request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list diagnosis items request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_ITEMS)

	result := &ListDiagnosisItemsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAutoDiagnosisStatus gets the auto diagnosis switch status of a cluster.
func GetAutoDiagnosisStatus(cli bce.Client, region string, request *GetAutoDiagnosisStatusRequest) (*GetAutoDiagnosisStatusResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get auto diagnosis status request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get auto diagnosis status request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_AUTO)

	result := &GetAutoDiagnosisStatusResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAutoDiagnosis updates the auto diagnosis switch of a cluster.
func UpdateAutoDiagnosis(cli bce.Client, region string, request *UpdateAutoDiagnosisRequest) (*UpdateAutoDiagnosisResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update auto diagnosis request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update auto diagnosis request clusterId should not be empty")
	}
	if request.Enabled == nil {
		return nil, errors.New("update auto diagnosis request enabled should not be nil")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_AUTO)

	result := &UpdateAutoDiagnosisResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLatestOverview gets the latest diagnosis overview of a cluster.
func GetLatestOverview(cli bce.Client, region string, request *GetLatestOverviewRequest) (*GetLatestOverviewResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get latest overview request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get latest overview request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_DIAGNOSIS, URI_DIAGNOSIS_OVERVIEW, URI_DIAGNOSIS_OVERVIEW_LATEST)

	result := &GetLatestOverviewResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}
