package bcm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	// GetMetricDataPath userId - scope - metricName
	GetMetricDataPath = "/json-api/v1/metricdata/%s/%s/%s"
	// BatchMetricDataPath userId - scope
	BatchMetricDataPath = "/json-api/v1/metricdata/metricName/%s/%s"

	Average     = "average"
	Maximum     = "maximum"
	Minimum     = "minimum"
	Sum         = "sum"
	SampleCount = "sampleCount"
)

// GetMetricData get metric data
func (c *Client) GetMetricData(req *GetMetricDataRequest) (*GetMetricDataResponse, error) {
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

	result := &GetMetricDataResponse{}
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
func (c *Client) BatchGetMetricData(req *BatchGetMetricDataRequest) (*BatchGetMetricDataResponse, error) {
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

	result := &BatchGetMetricDataResponse{}
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
