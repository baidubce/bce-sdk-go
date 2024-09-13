package bls

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bls/api"
)

func (c *Client) CreateProject(request CreateProjectRequest) error {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return err
	}
	return api.CreateProject(c, body)
}

func (c *Client) UpdateProject(request UpdateProjectRequest) error {
	param, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(param))
	if err != nil {
		return err
	}
	return api.UpdateProject(c, body)
}

func (c *Client) DescribeProject(request DescribeProjectRequest) (*api.Project, error) {
	return api.DescribeProject(c, request.UUID)
}

func (c *Client) DeleteProject(request DeleteProjectRequest) error {
	return api.DeleteProject(c, request.UUID)
}

func (c *Client) ListProject(request ListProjectRequest) (*api.ListProjectResult, error) {
	param, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromString(string(param))
	if err != nil {
		return nil, err
	}
	return api.ListProject(c, body)
}

func (c *Client) CreateLogStoreV2(request CreateLogStoreRequest) error {
	params, jsonErr := json.Marshal(&api.LogStore{
		Project:      request.Project,
		LogStoreName: request.LogStoreName,
		Retention:    request.Retention,
		Tags:         request.Tags,
	})
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return err
	}
	return api.CreateLogStore(c, body)
}

func (c *Client) UpdateLogStoreV2(request UpdateLogStoreRequest) error {
	param, jsonErr := json.Marshal(&api.LogStore{
		Retention: request.Retention,
	})
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(param))
	if err != nil {
		return err
	}
	return api.UpdateLogStore(c, request.Project, request.LogStoreName, body)
}

func (c *Client) DescribeLogStoreV2(request DescribeLogStoreRequest) (*api.LogStore, error) {
	return api.DescribeLogStore(c, request.Project, request.LogStoreName)
}

func (c *Client) DeleteLogStoreV2(request DeleteLogStoreRequest) error {
	return api.DeleteLogStore(c, request.Project, request.LogStoreName)
}

func (c *Client) ListLogStoreV2(request ListLogStoreRequest) (*api.ListLogStoreResult, error) {
	args := &api.QueryConditions{
		NamePattern: request.NamePattern,
		Order:       request.Order,
		OrderBy:     request.OrderBy,
		PageNo:      request.PageNo,
		PageSize:    request.PageSize,
	}
	return api.ListLogStore(c, request.Project, args)
}

func (c *Client) ListLogStreamV2(request ListLogStreamRequest) (*api.ListLogStreamResult, error) {
	args := &api.QueryConditions{
		NamePattern: request.NamePattern,
		Order:       request.Order,
		OrderBy:     request.OrderBy,
		PageNo:      request.PageNo,
		PageSize:    request.PageSize,
	}
	return api.ListLogStream(c, request.Project, request.LogStoreName, args)
}

func (c *Client) PushLogRecordV2(request PushLogRecordRequest) error {
	params, jsonErr := json.Marshal(&api.PushLogRecordBody{
		LogStreamName: request.LogStreamName,
		Type:          request.LogType,
		LogRecords:    request.LogRecords,
	})
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return err
	}
	return api.PushLogRecord(c, request.Project, request.LogStoreName, body)
}

func (c *Client) PullLogRecordV2(request PullLogRecordRequest) (*api.PullLogRecordResult, error) {
	args := &api.PullLogRecordArgs{
		LogStreamName: request.LogStreamName,
		StartDateTime: api.DateTime(request.StartDateTime),
		EndDateTime:   api.DateTime(request.EndDateTime),
		Limit:         request.Limit,
		Marker:        request.Marker,
	}
	return api.PullLogRecord(c, request.Project, request.LogStoreName, args)
}

func (c *Client) QueryLogRecordV2(request QueryLogRecordRequest) (*api.QueryLogResult, error) {
	args := &api.QueryLogRecordArgs{
		LogStreamName: request.LogStreamName,
		Query:         request.Query,
		StartDateTime: api.DateTime(request.StartDateTime),
		EndDateTime:   api.DateTime(request.EndDateTime),
		Limit:         request.Limit,
	}
	return api.QueryLogRecord(c, request.Project, request.LogStoreName, args)
}

func (c *Client) CreateFastQueryV2(request CreateFastQueryRequest) error {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return nil
	}
	return api.CreateFastQuery(c, body)
}

func (c *Client) UpdateFastQueryV2(request UpdateFastQueryRequest) error {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return nil
	}
	return api.UpdateFastQuery(c, body, request.FastQueryName)
}

func (c *Client) DescribeFastQueryV2(request DescribeFastQueryRequest) (*api.FastQuery, error) {
	return api.DescribeFastQuery(c, request.FastQueryName)
}

func (c *Client) DeleteFastQueryV2(request DeleteFastQueryRequest) error {
	return api.DeleteFastQuery(c, request.FastQueryName)
}

func (c *Client) ListFastQueryV2(request ListFastQueryRequest) (*api.ListFastQueryResult, error) {
	args := &api.QueryConditions{
		NamePattern: request.NamePattern,
		Order:       request.Order,
		OrderBy:     request.OrderBy,
		PageNo:      request.PageNo,
		PageSize:    request.PageSize,
	}
	return api.ListFastQuery(c, request.Project, request.LogStoreName, args)
}

func (c *Client) CreateIndexV2(request CreateIndexRequest) error {
	params, jsonErr := json.Marshal(&api.IndexFields{
		FullText: request.Fulltext,
		Fields:   request.Fields,
	})
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return err
	}
	return api.CreateIndex(c, request.Project, request.LogStoreName, body)
}

func (c *Client) UpdateIndexV2(request UpdateIndexRequest) error {
	params, jsonErr := json.Marshal(&api.IndexFields{
		FullText: request.Fulltext,
		Fields:   request.Fields,
	})
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return err
	}
	return api.UpdateIndex(c, request.Project, request.LogStoreName, body)
}

func (c *Client) DeleteIndexV2(request DeleteIndexRequest) error {
	return api.DeleteIndex(c, request.Project, request.LogStoreName)
}

func (c *Client) DescribeIndexV2(request DescribeIndexRequest) (*api.IndexFields, error) {
	return api.DescribeIndex(c, request.Project, request.LogStoreName)
}

func (c *Client) CreateLogShipperV2(request CreateLogShipperRequest) (string, error) {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return "", jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return "", nil
	}
	return api.CreateLogShipper(c, body)
}

func (c *Client) UpdateLogShipperV2(request UpdateLogShipperRequest) error {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return nil
	}
	return api.UpdateLogShipper(c, body, request.LogShipperID)
}

func (c *Client) GetLogShipperV2(request GetLogShipperRequest) (*api.LogShipper, error) {
	return api.GetLogShipper(c, request.LogShipperID)
}

func (c *Client) DeleteLogShipperV2(request DeleteLogShipperRequest) error {
	return api.DeleteSingleLogShipper(c, request.LogShipperID)
}

func (c *Client) BulkDeleteLogShipperV2(request BulkDeleteLogShipperRequest) error {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return nil
	}
	return api.BulkDeleteLogShipper(c, body)
}

func (c *Client) UpdateLogShipperStatusV2(request UpdateLogShipperStatusRequest) error {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return nil
	}
	return api.SetSingleLogShipperStatus(c, request.LogShipperID, body)
}

func (c *Client) BulkUpdateLogShipperStatusV2(request BulkUpdateLogShipperStatusRequest) error {
	params, jsonErr := json.Marshal(request)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromString(string(params))
	if err != nil {
		return nil
	}
	return api.BulkSetLogShipperStatus(c, body)
}

func (c *Client) ListLogShipperV2(request ListLogShipperRequest) (*api.ListShipperResult, error) {
	args := &api.ListLogShipperCondition{
		LogShipperID:   request.LogShipperID,
		LogShipperName: request.LogShipperName,
		Project:        request.Project,
		LogStoreName:   request.LogStoreName,
		DestType:       request.DestType,
		Status:         request.Status,
		Order:          request.Order,
		OrderBy:        request.OrderBy,
		PageNo:         request.PageNo,
		PageSize:       request.PageSize,
	}
	return api.ListLogShipper(c, args)
}

func (c *Client) ListLogShipperRecordV2(request ListShipperRecordRequest) (*api.ListShipperRecordResult, error) {
	args := &api.ListShipperRecordCondition{
		SinceHours: request.SinceHours,
		PageNo:     request.PageNo,
		PageSize:   request.PageSize,
	}
	return api.ListLogShipperRecord(c, request.LogShipperID, args)
}
