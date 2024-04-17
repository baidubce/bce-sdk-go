package api

import (
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
)

// The time interval of statistics data merged by, we defined 3 valid values here.
var (
	Period300   = createPeriod(300)
	Period3600  = createPeriod(3600)
	Period86400 = createPeriod(86400)
)

type Period interface {
	Value() int64
}

type period int64

func (p *period) Value() int64 {
	return int64(*p)
}

func createPeriod(value int64) Period {
	x := period(value)
	return &x
}

// QueryStatOption defined a method for setting optional configurations for query statistics.
type QueryStatOption func(interface{})

// QueryStatByDomains defined a method to pass the CDN domains you are interested in.
func QueryStatByDomains(domains []string) QueryStatOption {
	return func(o interface{}) {
		cfg, ok := o.(*queryStatOption)
		if ok && len(domains) > 0 {
			cfg.Domains = domains
		}
	}
}

// QueryStatByTimeRange the beginning and the end of the time range to query,
// The time values are in the ISO 8601 standard with the yyyy-MM-ddTHH:mm:ssZ format,
// e.g. 2024-04-15T00:00:00Z.
func QueryStatByTimeRange(startTime, endTime string) QueryStatOption {
	return func(o interface{}) {
		cfg, ok := o.(*queryStatOption)
		if ok {
			cfg.StartTime = startTime
			cfg.EndTime = endTime
		}
	}
}

// QueryStatByPeriod the time interval of statistics data merged by,
// valid values are 300, 3600 and 86400 represents by Period300, Period3060 and Period86400 respectively
func QueryStatByPeriod(period Period) QueryStatOption {
	return func(o interface{}) {
		if period == nil {
			return
		}
		cfg, ok := o.(*queryStatOption)
		if ok {
			cfg.Period = period.Value()
		}
	}
}

// QueryStatByCountry the specific Country you want to know about.
// The country value is one of the GEC codes.
// More details about GEC, you can read this article:
// https://baike.baidu.com/item/%E4%B8%96%E7%95%8C%E5%90%84%E5%9B%BD%E5%92%8C%E5%9C%B0%E5%8C%BA%E5%90%8D%E7%A7%B0%E4%BB%A3%E7%A0%81/6560023
// and pay attention of the column of the key world "两字母代码".
func QueryStatByCountry(country string) QueryStatOption {
	return func(o interface{}) {
		cfg, ok := o.(*queryStatOption)
		if ok {
			cfg.Country = country
		}
	}
}

type queryStatOption struct {
	Domains   []string
	EndTime   string
	StartTime string
	Period    int64
	Country   string
}

func (queryOption *queryStatOption) makeParams() map[string]string {
	params := map[string]string{}
	if len(queryOption.Domains) > 0 {
		params["domain"] = strings.Join(queryOption.Domains, ",")
	}
	if queryOption.StartTime != "" {
		params["startTime"] = queryOption.StartTime
	}
	if queryOption.EndTime != "" {
		params["endTime"] = queryOption.EndTime
	}
	if queryOption.Period > 0 {
		params["period"] = strconv.FormatInt(queryOption.Period, 10)
	}
	if queryOption.Country != "" {
		params["region"] = queryOption.Country
	}
	return params
}

// FlowDetail hold details of statistics of traffic.
type FlowDetail struct {
	Timestamp string `json:"timestamp"`
	Flow      int64  `json:"flow"`
	Bps       int64  `json:"bps"`
	Key       string `json:"key"`
}

// GetFlow - get the statistics of traffic(flow).
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Bkbszintg
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
//                4. QueryStatByCountry
// RETURNS:
//     - []FlowDetail: the details about traffic
//     - error: nil if success otherwise the specific error
func GetFlow(cli bce.Client, options ...QueryStatOption) ([]FlowDetail, error) {
	var queryOptions queryStatOption
	for _, opt := range options {
		opt(&queryOptions)
	}

	respObj := &struct {
		Status  string       `json:"status"`
		Count   int64        `json:"count"`
		Details []FlowDetail `json:"details"`
	}{}

	err := httpRequest(cli, "GET", "/v2/abroad/stat/flow", queryOptions.makeParams(), nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// PvDetail hold details of statistics of pv/qps.
type PvDetail struct {
	Timestamp string `json:"timestamp"`
	Pv        int64  `json:"pv"`
	Qps       int64  `json:"qps"`
	Key       string `json:"key"`
}

// GetPv - get the statistics of pv/qps.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/dkbszg48s
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
//                4. QueryStatByCountry
// RETURNS:
//     - []PvDetail: the details about pv/qps
//     - error: nil if success otherwise the specific error
func GetPv(cli bce.Client, options ...QueryStatOption) ([]PvDetail, error) {
	var queryOptions queryStatOption
	for _, opt := range options {
		opt(&queryOptions)
	}

	respObj := &struct {
		Status  string     `json:"status"`
		Count   int64      `json:"count"`
		Details []PvDetail `json:"details"`
	}{}

	err := httpRequest(cli, "GET", "/v2/abroad/stat/pv", queryOptions.makeParams(), nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// GetSrcFlow - get the statistics of traffic(flow) to your origin server.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/rkbsznt4v
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
// RETURNS:
//     - []FlowDetail: the details about traffic to your origin server.
//     - error: nil if success otherwise the specific error
func GetSrcFlow(cli bce.Client, options ...QueryStatOption) ([]FlowDetail, error) {
	var queryOptions queryStatOption
	for _, opt := range options {
		opt(&queryOptions)
	}

	respObj := &struct {
		Status  string       `json:"status"`
		Count   int64        `json:"count"`
		Details []FlowDetail `json:"details"`
	}{}

	err := httpRequest(cli, "GET", "/v2/abroad/stat/flow", queryOptions.makeParams(), nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// KvCounter defined a struct for name and count pairs.
type KvCounter struct {
	Name  int64 `json:"name"`
	Count int64 `json:"count"`
}

// HttpCodeDetail hold details of statistics of accessing HTTP codes.
type HttpCodeDetail struct {
	Timestamp string      `json:"timestamp"`
	Counters  []KvCounter `json:"counters"`
	Key       string      `json:"key"`
}

// GetHttpCode - get the statistics of accessing HTTP codes.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ekbszvxv5
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
// RETURNS:
//     - []HttpCodeDetail: the details about accessing HTTP codes.
//     - error: nil if success otherwise the specific error
func GetHttpCode(cli bce.Client, options ...QueryStatOption) ([]HttpCodeDetail, error) {
	var queryOptions queryStatOption
	for _, opt := range options {
		opt(&queryOptions)
	}

	respObj := &struct {
		Status  string           `json:"status"`
		Count   int64            `json:"count"`
		Details []HttpCodeDetail `json:"details"`
	}{}

	err := httpRequest(cli, "GET", "/v2/abroad/stat/httpcode", queryOptions.makeParams(), nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}

// HitDetail defined a struct for hitting ratio of Edge accessing.
type HitDetail struct {
	Timestamp string  `json:"timestamp"`
	HitRate   float64 `json:"hitrate"`
	Key       string  `json:"key"`
}

// GetRealHit - get the statistics of hit rate of accessing by traffic.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ckbszuehh
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
// RETURNS:
//     - []HitDetail: the details about traffic hit rate.
//     - error: nil if success otherwise the specific error
func GetRealHit(cli bce.Client, options ...QueryStatOption) ([]HitDetail, error) {
	var queryOptions queryStatOption
	for _, opt := range options {
		opt(&queryOptions)
	}

	respObj := &struct {
		Status  string      `json:"status"`
		Count   int64       `json:"count"`
		Details []HitDetail `json:"details"`
	}{}

	err := httpRequest(cli, "GET", "/v2/abroad/stat/realhit", queryOptions.makeParams(), nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Details, nil
}
