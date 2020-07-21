package api

import (
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	statisticsObjectKey  = "/v2/stat/query"
	statisticsBillingKey = "/v2/billing"
)

// QueryCondition defined a struct for query condition
type QueryCondition struct {
	EndTime   string   `json:"endTime,omitempty"`
	StartTime string   `json:"startTime,omitempty"`
	Period    int      `json:"period,omitempty"`
	KeyType   int      `json:"key_type"`
	Key       []string `json:"key,omitempty"`
	GroupBy   string   `json:"groupBy,omitempty"`
}

type DetailBase struct {
	Timestamp string `json:"timestamp"`
	Key       string `json:"key,omitempty"`
}

// NetSite defined a struct for the ISP's information
type NetSite struct {
	Location string `json:"location"`
	Isp      string `json:"isp"`
}

type AvgSpeedRegionData struct {
	*NetSite
	AvgSpeed int64 `json:"avgspeed"`
}

type AvgSpeedDetail struct {
	*DetailBase
	AvgSpeed int64 `json:"avgspeed"`
}

type AvgSpeedRegionDetail struct {
	*DetailBase
	Distribution []AvgSpeedRegionData `json:"distribution"`
}

type PvDetail struct {
	*DetailBase
	Pv  int64 `json:"pv"`
	Qps int64 `json:"qps"`
}

type PVRegionData struct {
	*NetSite
	Pv  int64 `json:"pv"`
	Qps int64 `json:"qps"`
}

type PvRegionDetail struct {
	*DetailBase
	Distribution []PVRegionData `json:"distribution"`
}

type UvDetail struct {
	*DetailBase
	Uv int64 `json:"uv"`
}

type FlowDetail struct {
	*DetailBase
	Flow float64 `json:"flow"`
	Bps  int64   `json:"bps"`
}

type FlowRegionData struct {
	*NetSite
	Flow float64 `json:"flow"`
	Bps  int64   `json:"bps"`
}

type FlowRegionDetail struct {
	*DetailBase
	Distribution []FlowRegionData `json:"distribution"`
}

type HitDetail struct {
	*DetailBase
	HitRate float64 `json:"hitrate"`
}

type KvCounter struct {
	Name  int64 `json:"name"`
	Count int64 `json:"count"`
}

type HttpCodeDetail struct {
	*DetailBase
	Counters []KvCounter `json:"counters"`
}

type HttpCodeRegionData struct {
	*NetSite
	Counters []KvCounter `json:"counters"`
}

type HttpCodeRegionDetail struct {
	*DetailBase
	Distribution []HttpCodeRegionData `json:"distribution"`
}

type TopNCounter struct {
	Name string  `json:"name"`
	Pv   int64   `json:"pv"`
	Flow float64 `json:"flow"`
}

type TopNDetail struct {
	*DetailBase
	TotalPv   int64         `json:"total_pv"`
	TotalFlow float64       `json:"total_flow"`
	Counters  []TopNCounter `json:"counters"`
}

type ErrorKvCounter struct {
	Code     string      `json:"code"`
	Counters []KvCounter `json:"counters"`
}

type ErrorDetail struct {
	*DetailBase
	Counters []ErrorKvCounter `json:"counters"`
}

// GetAvgSpeed - get the average speed
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E5%B9%B3%E5%9D%87%E9%80%9F%E7%8E%87
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []AvgSpeedDetail: the detail list about the average speed
//     - error: nil if success otherwise the specific error
func GetAvgSpeed(cli bce.Client, queryCondition *QueryCondition) ([]AvgSpeedDetail, error) {

	respObj := &struct {
		Status  string           `json:"status"`
		Count   int64            `json:"count"`
		Details []AvgSpeedDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Metric:         "avg_speed",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetAvgSpeed - get the average speed filter by location
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83%E6%9F%A5%E8%AF%A2%E5%B9%B3%E5%9D%87%E9%80%9F%E7%8E%87
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []AvgSpeedRegionDetail: the detail list about the average speed
//     - error: nil if success otherwise the specific error
func GetAvgSpeedByRegion(cli bce.Client, queryCondition *QueryCondition, prov string, isp string) ([]AvgSpeedRegionDetail, error) {

	respObj := &struct {
		Status  string                 `json:"status"`
		Count   int64                  `json:"count"`
		Details []AvgSpeedRegionDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Prov   string `json:"prov,omitempty"`
		Isp    string `json:"isp,omitempty"`
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Prov:           prov,
		Isp:            isp,
		Metric:         "avg_speed_region",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetPv - get the PV data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#pvqps%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - level: the node level, the available values are "edge", "internal" and "all"
// RETURNS:
//     - []PvDetail: the detail list about page view
//     - error: nil if success otherwise the specific error
func GetPv(cli bce.Client, queryCondition *QueryCondition, level string) ([]PvDetail, error) {

	respObj := &struct {
		Status  string     `json:"status"`
		Count   int64      `json:"count"`
		Details []PvDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Level  string `json:"level,omitempty"`
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Level:          level,
		Metric:         "pv",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetSrcPv - get the PV data in back to the sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%9B%9E%E6%BA%90pvqps%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []PvDetail: the detail list about page view
//     - error: nil if success otherwise the specific error
func GetSrcPv(cli bce.Client, queryCondition *QueryCondition) ([]PvDetail, error) {

	respObj := &struct {
		Status  string     `json:"status"`
		Count   int64      `json:"count"`
		Details []PvDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Metric:         "pv_src",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetAvgPvByRegion - get the PV data filter by location
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2pvqps%E5%88%86%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []PvRegionDetail: the detail list about page view
//     - error: nil if success otherwise the specific error
func GetPvByRegion(cli bce.Client, queryCondition *QueryCondition, prov string, isp string) ([]PvRegionDetail, error) {

	respObj := &struct {
		Status  string           `json:"status"`
		Count   int64            `json:"count"`
		Details []PvRegionDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Prov   string `json:"prov,omitempty"`
		Isp    string `json:"isp,omitempty"`
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Prov:           prov,
		Isp:            isp,
		Metric:         "pv_region",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetUv - get the UV data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#uv%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []UvDetail: the detail list about unique visitor
//     - error: nil if success otherwise the specific error
func GetUv(cli bce.Client, queryCondition *QueryCondition) ([]UvDetail, error) {

	respObj := &struct {
		Status  string     `json:"status"`
		Count   int64      `json:"count"`
		Details []UvDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Level  string `json:"level,omitempty"`
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Metric:         "uv",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetFlow - get the flow data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%B5%81%E9%87%8F%E3%80%81%E5%B8%A6%E5%AE%BD
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []FlowDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func GetFlow(cli bce.Client, queryCondition *QueryCondition, level string) ([]FlowDetail, error) {

	respObj := &struct {
		Status  string       `json:"status"`
		Count   int64        `json:"count"`
		Details []FlowDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Level  string `json:"level,omitempty"`
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Level:          level,
		Metric:         "flow",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetFlowByProtocol - get the flow data filter by protocol
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%B5%81%E9%87%8F%E3%80%81%E5%B8%A6%E5%AE%BD%E5%88%86%E5%8D%8F%E8%AE%AE
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - protocol: the specified HTTP protocol, like "http" or "https", "all" means both "http" and "https"
// RETURNS:
//     - []FlowDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func GetFlowByProtocol(cli bce.Client, queryCondition *QueryCondition, protocol string) ([]FlowDetail, error) {

	respObj := &struct {
		Status  string       `json:"status"`
		Count   int64        `json:"count"`
		Details []FlowDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Protocol string `json:"protocol,omitempty"`
		Metric   string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Protocol:       protocol,
		Metric:         "flow_protocol",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetFlowByRegion - get the flow data filter by location
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%B5%81%E9%87%8F%E3%80%81%E5%B8%A6%E5%AE%BD%EF%BC%88%E5%88%86%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83%EF%BC%89
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []FlowRegionDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func GetFlowByRegion(cli bce.Client, queryCondition *QueryCondition, prov string, isp string) ([]FlowRegionDetail, error) {

	respObj := &struct {
		Status  string             `json:"status"`
		Count   int64              `json:"count"`
		Details []FlowRegionDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Prov   string `json:"prov,omitempty"`
		Isp    string `json:"isp,omitempty"`
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Prov:           prov,
		Isp:            isp,
		Metric:         "flow_region",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetSrcFlow - get the flow data in backed to sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E5%9B%9E%E6%BA%90%E6%B5%81%E9%87%8F%E3%80%81%E5%9B%9E%E6%BA%90%E5%B8%A6%E5%AE%BD
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []FlowDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func GetSrcFlow(cli bce.Client, queryCondition *QueryCondition) ([]FlowDetail, error) {

	respObj := &struct {
		Status  string       `json:"status"`
		Count   int64        `json:"count"`
		Details []FlowDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Protocol string `json:"protocol,omitempty"`
		Metric   string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Metric:         "src_flow",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

func getHit(cli bce.Client, queryCondition *QueryCondition, metric string) ([]HitDetail, error) {
	respObj := &struct {
		Status  string      `json:"status"`
		Count   int64       `json:"count"`
		Details []HitDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Protocol string `json:"protocol,omitempty"`
		Metric   string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Metric:         metric,
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetRealHit - get the detail about byte hit rate
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%AD%97%E8%8A%82%E5%91%BD%E4%B8%AD%E7%8E%87%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HitDetail: the detail list about byte rate
//     - error: nil if success otherwise the specific error
func GetRealHit(cli bce.Client, queryCondition *QueryCondition) ([]HitDetail, error) {
	return getHit(cli, queryCondition, "real_hit")
}

// GetPvHit - get the detail about PV hit rate
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E8%AF%B7%E6%B1%82%E5%91%BD%E4%B8%AD%E7%8E%87%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HitDetail: the detail list about pv rate
//     - error: nil if success otherwise the specific error
func GetPvHit(cli bce.Client, queryCondition *QueryCondition) ([]HitDetail, error) {
	return getHit(cli, queryCondition, "pv_hit")
}

func getHttpCode(cli bce.Client, queryCondition *QueryCondition, metric string) ([]HttpCodeDetail, error) {
	respObj := &struct {
		Status  string           `json:"status"`
		Count   int64            `json:"count"`
		Details []HttpCodeDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Protocol string `json:"protocol,omitempty"`
		Metric   string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Metric:         metric,
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetHttpCode - get the http code's statistics
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E7%8A%B6%E6%80%81%E7%A0%81%E7%BB%9F%E8%AE%A1%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HttpCodeDetail: the detail list about http code
//     - error: nil if success otherwise the specific error
func GetHttpCode(cli bce.Client, queryCondition *QueryCondition) ([]HttpCodeDetail, error) {
	return getHttpCode(cli, queryCondition, "httpcode")
}

// GetSrcHttpCode - get the http code's statistics in backed to sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%9B%9E%E6%BA%90%E7%8A%B6%E6%80%81%E7%A0%81%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HttpCodeDetail: the detail list about http code
//     - error: nil if success otherwise the specific error
func GetSrcHttpCode(cli bce.Client, queryCondition *QueryCondition) ([]HttpCodeDetail, error) {
	return getHttpCode(cli, queryCondition, "src_httpcode")
}

// GetHttpCodeByRegion - get the http code's statistics filter by location
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E7%8A%B6%E6%80%81%E7%A0%81%E7%BB%9F%E8%AE%A1%E6%9F%A5%E8%AF%A2%EF%BC%88%E5%88%86%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83%EF%BC%89
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []HttpCodeRegionDetail: the detail list about http code
//     - error: nil if success otherwise the specific error
func GetHttpCodeByRegion(cli bce.Client, queryCondition *QueryCondition, prov string, isp string) ([]HttpCodeRegionDetail, error) {

	respObj := &struct {
		Status  string                 `json:"status"`
		Count   int64                  `json:"count"`
		Details []HttpCodeRegionDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Prov   string `json:"prov,omitempty"`
		Isp    string `json:"isp,omitempty"`
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Prov:           prov,
		Isp:            isp,
		Metric:         "httpcode_region",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

func getTopN(cli bce.Client, queryCondition *QueryCondition, httpCode string, metric string) ([]TopNDetail, error) {
	extra, err := strconv.ParseInt(httpCode, 10, 64)
	if err != nil {
		extra = 0
	}

	respObj := &struct {
		Status  string       `json:"status"`
		Count   int64        `json:"count"`
		Details []TopNDetail `json:"details"`
	}{}

	err = httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Protocol string `json:"protocol,omitempty"`
		Metric   string `json:"metric"`
		Extra    int64  `json:"extra,omitempty"`
	}{
		QueryCondition: queryCondition,
		Metric:         metric,
		Extra:          extra,
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetTopNUrls - get the top N urls that requested
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#topn-urls
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - httpCode: the specified HTTP code, like "200"
// RETURNS:
//     - []TopNDetail: the top N urls' detail
//     - error: nil if success otherwise the specific error
func GetTopNUrls(cli bce.Client, queryCondition *QueryCondition, httpCode string) ([]TopNDetail, error) {
	return getTopN(cli, queryCondition, httpCode, "top_urls")
}

// GetTopNReferers - get the top N urls that brought by requested
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#topn-referers
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - httpCode: the specified HTTP code, like "200"
// RETURNS:
//     - []TopNDetail: the top N referer urls' detail
//     - error: nil if success otherwise the specific error
func GetTopNReferers(cli bce.Client, queryCondition *QueryCondition, httpCode string) ([]TopNDetail, error) {
	return getTopN(cli, queryCondition, httpCode, "top_referers")
}

// GetTopNDomains - get the top N domains that requested
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#topn-domains
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
//     - httpCode: the specified HTTP code, like "200"
// RETURNS:
//     - []TopNDetail: the top N domains' detail
//     - error: nil if success otherwise the specific error
func GetTopNDomains(cli bce.Client, queryCondition *QueryCondition, httpCode string) ([]TopNDetail, error) {
	return getTopN(cli, queryCondition, httpCode, "top_domains")
}

// GetError - get the error code's data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#cdn%E9%94%99%E8%AF%AF%E7%A0%81%E5%88%86%E7%B1%BB%E7%BB%9F%E8%AE%A1%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryCondition: the querying conditions
// RETURNS:
//     - []ErrorDetail: the top N error details
//     - error: nil if success otherwise the specific error
func GetError(cli bce.Client, queryCondition *QueryCondition) ([]ErrorDetail, error) {
	respObj := &struct {
		Status  string        `json:"status"`
		Count   int64         `json:"count"`
		Details []ErrorDetail `json:"details"`
	}{}

	err := httpRequest(cli, "POST", statisticsObjectKey, nil, &struct {
		*QueryCondition
		Metric string `json:"metric"`
	}{
		QueryCondition: queryCondition,
		Metric:         "error",
	}, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetPeak95Bandwidth - get peak 95 bandwidth for the specified tags or domains.
// For details, pleader refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%9C%8895%E5%B3%B0%E5%80%BC%E5%B8%A6%E5%AE%BD
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - startTime: start time which in `YYYY-mm-ddTHH:ii:ssZ` style
//     - endTime: end time which in `YYYY-mm-ddTHH:ii:ssZ` style
//     - domains: a list of domains, only one of `tags` and `domains` can contains item
//     - tags: a list of tag names, only one of `tags` and `domains` can contains item
// RETURNS:
//     - string: the peak95 time which in `YYYY-mm-ddTHH:ii:ssZ` style
//     - int64: peak95 bandwidth
//     - error: nil if success otherwise the specific error
func GetPeak95Bandwidth(cli bce.Client, startTime, endTime string, domains, tags []string) (peak95Time string, peak95Band int64, err error) {
	respObj := &struct {
		Details struct {
			Bandwidth int64  `json:"bill_band"`
			Time      string `json:"bill_time"`
		} `json:"billing_details"`
	}{}

	tagOrDomains, withTag := domains, false
	if len(tags) != 0 {
		tagOrDomains, withTag = tags, true
	}
	err = httpRequest(cli, "POST", statisticsBillingKey, nil, map[string]interface{}{
		"domains":   strings.Join(tagOrDomains, ","),
		"type":      "peak95",
		"withTag":   withTag,
		"byTime":    true,
		"startTime": startTime,
		"endTime":   endTime,
	}, respObj)
	if err != nil {
		return
	}
	return respObj.Details.Time, respObj.Details.Bandwidth, nil
}
