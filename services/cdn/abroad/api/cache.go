package api

import (
	"errors"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util"
)

type PurgedId string
type PrefetchId string

// CStatusQueryData defined a struct for the query conditions about the tasks' progress
type CStatusQueryData struct {
	EndTime   string
	StartTime string
	Url       string
	Marker    string
	Id        string
}

// CRecordQueryData defined a struct for the query conditions about the operated records
type CRecordQueryData struct {
	EndTime   string
	StartTime string
	Url       string
	Marker    string
	TaskType  string
}

// PurgeTask defined a struct for purged task
type PurgeTask struct {
	Url  string `json:"url"`
	Type string `json:"type,omitempty"`
}

// PrefetchTask defined a struct for prefetch task
type PrefetchTask struct {
	Url string `json:"url"`
}

// CachedDetail defined a struct for task details
type CachedDetail struct {
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	Progress   int64  `json:"progress"`
}

// PurgedDetail defined a struct for purged task information
type PurgedDetail struct {
	*CachedDetail
	Task PurgeTask `json:"task"`
}

// PrefetchDetail defined a struct for prefetch task information
type PrefetchDetail struct {
	*CachedDetail
	Task PrefetchTask `json:"task"`
}

// PurgedStatus defined a struct for purged status
type PurgedStatus struct {
	Details     []PurgedDetail `json:"details"`
	IsTruncated bool           `json:"isTruncated"`
	NextMarker  string         `json:"nextMarker"`
}

// PrefetchStatus defined a struct for prefetch status
type PrefetchStatus struct {
	Details     []PrefetchDetail `json:"details"`
	IsTruncated bool             `json:"isTruncated"`
	NextMarker  string           `json:"nextMarker"`
}

// QuotaDetail defined a struct for query quota
type QuotaDetail struct {
	UrlPurgeQuota  int64 `json:"urlPurgeQuota"`
	UrlPurgeRemain int64 `json:"urlPurgeRemain"`
	DirPurgeQuota  int64 `json:"dirPurgeQuota"`
	DirPurgeRemain int64 `json:"dirPurgeRemain"`
	PrefetchQuota  int64 `json:"prefetchQuota"`
	PrefetchRemain int64 `json:"prefetchRemain"`
}

// Purge - tells the CDN system to purge the specified files
// For more details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Zkbsy0k8j
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - tasks: the tasks about purging the files from the CDN nodes
// RETURNS:
//     - PurgedId: an ID representing a purged task, using it to search the task progress
//     - error: nil if success otherwise the specific error
func Purge(cli bce.Client, tasks []PurgeTask) (PurgedId, error) {
	respObj := &struct {
		Id string `json:"id"`
	}{}

	err := httpRequest(cli, "POST", "/v2/abroad/cache/purge", nil, &struct {
		Tasks []PurgeTask `json:"tasks"`
	}{
		Tasks: tasks,
	}, respObj)
	if err != nil {
		return "", err
	}

	return PurgedId(respObj.Id), nil
}

// GetPurgedStatus - get the purged progress
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ikbsy9cvb
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryData: querying conditions, it contains the time interval, the task ID and the specified url
// RETURNS:
//     - *PurgedStatus: the details about the purged
//     - error: nil if success otherwise the specific error
func GetPurgedStatus(cli bce.Client, queryData *CStatusQueryData) (*PurgedStatus, error) {
	if queryData == nil {
		queryData = &CStatusQueryData{}
	}

	params := map[string]string{}
	if queryData.Id != "" {
		params["id"] = queryData.Id
	}
	if err := getTimeParams(params, queryData.StartTime, queryData.EndTime); err != nil {
		return nil, err
	}

	if queryData.Url != "" {
		params["url"] = queryData.Url
	}
	if queryData.Marker != "" {
		params["marker"] = queryData.Marker
	}

	respObj := &PurgedStatus{}
	err := httpRequest(cli, "GET", "/v2/abroad/cache/purge", params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// Prefetch - tells the CDN system to prefetch the specified files
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Dlj5ch09q
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - tasks: the tasks about prefetch the files from the CDN nodes
//     - error: nil if success otherwise the specific error
func Prefetch(cli bce.Client, tasks []PrefetchTask) (PrefetchId, error) {
	respObj := &struct {
		Id string `json:"id"`
	}{}

	err := httpRequest(cli, "POST", "/v2/abroad/cache/prefetch", nil, &struct {
		Tasks []PrefetchTask `json:"tasks"`
	}{
		Tasks: tasks,
	}, respObj)
	if err != nil {
		return "", err
	}

	return PrefetchId(respObj.Id), nil
}

// GetPrefetchStatus - get the prefetch progress
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Mlj5e9y0i
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryData: querying conditions, it contains the time interval, the task ID and the specified url
// RETURNS:
//     - *PrefetchStatus: the details about the prefetch
//     - error: nil if success otherwise the specific error
func GetPrefetchStatus(cli bce.Client, queryData *CStatusQueryData) (*PrefetchStatus, error) {
	if queryData == nil {
		queryData = &CStatusQueryData{}
	}

	params := map[string]string{}
	if queryData.Id != "" {
		params["id"] = queryData.Id
	}
	if err := getTimeParams(params, queryData.StartTime, queryData.EndTime); err != nil {
		return nil, err
	}

	if queryData.Url != "" {
		params["url"] = queryData.Url
	}
	if queryData.Marker != "" {
		params["marker"] = queryData.Marker
	}

	respObj := &PrefetchStatus{}
	err := httpRequest(cli, "GET", "/v2/abroad/cache/prefetch", params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// GetQuota - get the quota about purge and prefetch
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/flnoakciq
//
// RETURNS:
//     - cli: the client agent which can perform sending request
//     - QuotaDetail: the quota details about a specified user
//     - error: nil if success otherwise the specific error
func GetQuota(cli bce.Client) (*QuotaDetail, error) {
	respObj := &QuotaDetail{}
	err := httpRequest(cli, "GET", "/v2/abroad/cache/quota", nil, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

func getTimeParams(params map[string]string, startTime, endTime string) error {
	// get "endTime"
	endTs := int64(0)
	if endTime == "" {
		// default current time
		endTs = time.Now().Unix()
		params["endTime"] = util.FormatISO8601Date(endTs)
	} else {
		t, err := util.ParseISO8601Date(endTime)
		if err != nil {
			return err
		}
		endTs = t.Unix()
		params["endTime"] = endTime
	}

	// get "startTime", the default "startTime" is one day later than the "endTime"
	startTs := int64(0)
	if startTime == "" {
		startTs = endTs - 24*60*60
		params["startTime"] = util.FormatISO8601Date(startTs)
	} else {
		t, err := util.ParseISO8601Date(startTime)
		if err != nil {
			return err
		}
		startTs = t.Unix()
		params["startTime"] = startTime
	}

	// the "startTime should be less than the "endTime"
	// if we set "startTime" but not "endTime", we normally assign the current time to the "endTime",
	// in this condition, we might get "startTs > endTs"
	if startTs > endTs {
		return errors.New("error time range, the startTime should be less than the endTime")
	}

	return nil
}
