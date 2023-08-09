package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util"
)

type LogBase struct {
	Domain string `json:"domain"`
	Url    string `json:"url"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
}

// TimeInterval defined a struct contains the started time and the end time
type TimeInterval struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// LogEntry defined a struct for log information
type LogEntry struct {
	*LogBase
	*TimeInterval
}

// LogQueryData defined a struct for query conditions
type LogQueryData struct {
	TimeInterval
	Type     int      `json:"type,omitempty"`
	Domains  []string `json:"domains,omitempty"`
	PageNo   int      `json:"pageNo,omitempty"`
	PageSize int      `json:"pageSize,omitempty"`
}

// GetDomainLog -get one domain's log urls
// For details, please refer https://cloud.baidu.com/doc/CDN/s/cjwvyf0r9
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - timeInterval: the specified time interval
//
// RETURNS:
//   - []LogEntry: the log detail list
//   - error: nil if success otherwise the specific error
func GetDomainLog(cli bce.Client, domain string, timeInterval TimeInterval) ([]LogEntry, error) {
	if err := checkTimeInterval(timeInterval, 14*24*60*60); err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("/v2/log/%s/log", domain)
	params := map[string]string{}
	params["startTime"] = timeInterval.StartTime
	params["endTime"] = timeInterval.EndTime

	respObj := &struct {
		Logs []LogEntry `json:"logs"`
	}{}

	if err := httpRequest(cli, "GET", urlPath, params, nil, respObj); err != nil {
		return nil, err
	}

	for i, _ := range respObj.Logs {
		respObj.Logs[i].Domain = domain
	}

	return respObj.Logs, nil
}

// GetMultiDomainLog - get many domains' log urls
// For details, please refer https://cloud.baidu.com/doc/CDN/API.html#.49.B0.4F.9D.D3.1A.FB.6F.59.A6.8A.B6.08.E9.BC.EF
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryData: the querying conditions
//   - error: nil if success otherwise the specific error
func GetMultiDomainLog(cli bce.Client, queryData *LogQueryData) ([]LogEntry, error) {
	if queryData == nil {
		return nil, errors.New("queryData could not be nil")
	}
	if err := checkTimeInterval(queryData.TimeInterval, 180*24*60*60); err != nil {
		return nil, err
	}

	if queryData.PageNo == 0 {
		queryData.PageNo = 1
	}
	if queryData.PageNo < 0 {
		return nil, errors.New("invalid queryData.PageNo, it should be larger than 0")
	}

	if queryData.PageSize < 0 {
		return nil, errors.New("invalid queryData.PageSize, it should be larger than 0")
	}

	respObj := &struct {
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
		TotalCount int    `json:"totalCount"`
		Urls       []struct {
			*LogBase
			LogTimeBegin string `json:"logTimeBegin"`
			LogTimeEnd   string `json:"logTimeEnd"`
		} `json:"urls"`
	}{}

	err := httpRequest(cli, "POST", "/v2/log/list", nil, queryData, respObj)
	if err != nil {
		return nil, err
	}

	result := make([]LogEntry, 0, len(respObj.Urls))

	for i, _ := range respObj.Urls {
		log := LogEntry{
			LogBase: respObj.Urls[i].LogBase,
			TimeInterval: &TimeInterval{
				StartTime: respObj.Urls[i].LogTimeBegin,
				EndTime:   respObj.Urls[i].LogTimeEnd,
			},
		}
		result = append(result, log)
	}

	return result, nil
}

func checkTimeInterval(timeInterval TimeInterval, maxTimeRange int64) error {
	if timeInterval.StartTime == "" {
		return errors.New("lack of startTime")
	}

	if timeInterval.EndTime == "" {
		return errors.New("lack of endTime")
	}

	st, err := util.ParseISO8601Date(timeInterval.StartTime)
	if err != nil {
		return errors.New(fmt.Sprintf("invalid startTime, %s", err.Error()))
	}
	startTs := st.Unix()

	et, err := util.ParseISO8601Date(timeInterval.EndTime)
	if err != nil {
		return errors.New(fmt.Sprintf("invalid endTime, %s", err.Error()))
	}
	endTs := et.Unix()

	if startTs > endTs {
		return errors.New(fmt.Sprintf("startTime shouble be less than endTime"))
	}

	curTs := time.Now().Unix()
	if curTs-startTs > maxTimeRange {
		return errors.New(fmt.Sprintf("only the first %d seconds of log files can be downloaded, please reset startTime", maxTimeRange))
	}

	if startTs > curTs {
		return errors.New("startTime could not larger than current time")
	}

	return nil
}
