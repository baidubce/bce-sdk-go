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

// GetDomainLog -get one domain's log urls
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/mjwvxj0ec
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - timeInterval: the specified time interval
// RETURNS:
//     - []LogEntry: the log detail list
//     - error: nil if success otherwise the specific error
func GetDomainLog(cli bce.Client, domain string, timeInterval TimeInterval) ([]LogEntry, error) {
	if err := checkTimeInterval(timeInterval, 14*24*60*60); err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("/v2/abroad/log/%s", domain)
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
