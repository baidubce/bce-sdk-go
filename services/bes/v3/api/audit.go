package api

import (
	"errors"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// GetAuditStatus queries the data-plane audit status of a cluster.
func GetAuditStatus(cli bce.Client, region string, request *GetAuditStatusRequest) (*AuditStatusResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get audit status request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get audit status request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_AUDIT)

	result := &AuditStatusResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// EnableAudit enables data-plane audit for a cluster.
func EnableAudit(cli bce.Client, region string, request *EnableAuditRequest) (*EnableAuditResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("enable audit request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("enable audit request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_AUDIT)

	result := &EnableAuditResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListAuditEventTypes lists the searchable audit event types of a cluster.
func ListAuditEventTypes(cli bce.Client, region string, request *ListAuditEventTypesRequest) (*ListAuditEventTypesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list audit event types request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list audit event types request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_AUDIT, URI_AUDIT_EVENT_TYPES)

	result := &ListAuditEventTypesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SearchAuditEvents searches data-plane audit events of a cluster.
func SearchAuditEvents(cli bce.Client, region string, request *SearchAuditEventsRequest) (*SearchAuditEventsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("search audit events request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("search audit events request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, URI_AUDIT, URI_AUDIT_EVENTS)

	query := map[string]string{}
	if request.StartTime != "" {
		query["startTime"] = request.StartTime
	}
	if request.EndTime != "" {
		query["endTime"] = request.EndTime
	}
	if request.Category != "" {
		query["category"] = request.Category
	}
	if request.IndexName != "" {
		query["indexName"] = request.IndexName
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}
	if request.Cursor != "" {
		query["cursor"] = request.Cursor
	}

	result := &SearchAuditEventsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBctAuthSwitch queries the cloud audit (BCT) authorization switch status.
func GetBctAuthSwitch(cli bce.Client, region string, request *Request) (*GetBctAuthSwitchResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request != nil {
		region = setRegion(region, request.Region)
	}

	result := &GetBctAuthSwitchResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, URI_BCT_AUTH_SWITCH, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateBctAuthSwitch updates the cloud audit (BCT) authorization switch status.
func UpdateBctAuthSwitch(cli bce.Client, region string, request *UpdateBctAuthSwitchRequest) (*BctAuthSwitchResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update bct auth switch request should not be nil")
	}
	if request.Enabled == nil {
		return nil, errors.New("update bct auth switch request enabled should not be nil")
	}
	region = setRegion(region, request.Region)

	result := &BctAuthSwitchResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, URI_BCT_AUTH_SWITCH, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SearchBctEvents searches cloud audit (BCT) control-plane events.
func SearchBctEvents(cli bce.Client, region string, request *SearchBctEventsRequest) (*SearchBctEventsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("search bct events request should not be nil")
	}
	if request.StartTime == "" {
		return nil, errors.New("search bct events request startTime should not be empty")
	}
	if request.EndTime == "" {
		return nil, errors.New("search bct events request endTime should not be empty")
	}
	region = setRegion(region, request.Region)

	result := &SearchBctEventsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, URI_BCT_EVENTS_QUERY, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
