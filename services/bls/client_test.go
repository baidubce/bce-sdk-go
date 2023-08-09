package bls

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bls/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BLS_CLIENT          *Client
	EXIST_LOGSTORE      = "bls-rd-wzy"
	LOGSTREAM_PATTERN   = "wzy"
	JSON_LOGSTREAM      = LOGSTREAM_PATTERN + "-JSON"
	TEXT_LOGSTREAM      = LOGSTREAM_PATTERN + "-TEXT"
	FASTQUERY_NAME      = "speedo"
	LOGSHIPPER_NAME     = "test-bls-sdk"
	TEST_STARTTIME      = "2021-07-06T19:01:00Z"
	LOGSHIPPER_ID       = "MjI3NDY5OTE5MTk3MzE1MDcy"
	DEAFAULT_LOGSTORE   = "ng-log"
	DEAFAULT_BOSPATH    = "/bls-test/bls-sdk/"
	DEFAULT_TEST_DOMAIN = "10.132.106.242"
)

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given: %+v\n", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &BlsClientConfiguration{}
	err = decoder.Decode(confObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	// default use http protocol
	BLS_CLIENT, _ = NewClient(confObj.Ak, confObj.Sk, DEFAULT_TEST_DOMAIN)
	// log.SetLogHandler(log.STDERR | log.FILE)
	// log.SetRotateType(log.ROTATE_SIZE)
	log.SetLogLevel(log.WARN)
	// log.SetLogHandler(log.STDERR)
	// log.SetLogLevel(log.DEBUG)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	actual interface{}, expected interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

// LogStore test
func TestCreateLogStore(t *testing.T) {
	err := BLS_CLIENT.CreateLogStore(EXIST_LOGSTORE, 3)
	ExpectEqual(t.Errorf, err, nil)
	err = BLS_CLIENT.CreateLogStore(EXIST_LOGSTORE, 36)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusConflict)
	}
	res, err := BLS_CLIENT.DescribeLogStore(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.LogStoreName, EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res.Retention, 3)
}

func TestUpdateLogStore(t *testing.T) {
	err := BLS_CLIENT.UpdateLogStore(EXIST_LOGSTORE, 8)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BLS_CLIENT.DescribeLogStore(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.LogStoreName, EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res.Retention, 8)
	err = BLS_CLIENT.UpdateLogStore("not"+EXIST_LOGSTORE, 22)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestDescribeLogStore(t *testing.T) {
	res, err := BLS_CLIENT.DescribeLogStore(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.LogStoreName, EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res.Retention, 8)
	res, err = BLS_CLIENT.DescribeLogStore("not" + EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res, nil)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestListLogStore(t *testing.T) {
	args := &api.QueryConditions{
		NamePattern: "bls-rd",
		Order:       "asc",
		OrderBy:     "creationDateTime",
		PageNo:      1,
		PageSize:    10}
	res, err := BLS_CLIENT.ListLogStore(args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(res.Result) > 0, true)
}

// LogRecord test
func TestPushLogRecord(t *testing.T) {
	jsonRecords := []api.LogRecord{
		{
			Message:   "{\"body_bytes_sent\":184,\"bytes_sent\":398,\"client_ip\":\"120.193.204.39\",\"connection\":915958195,\"hit\":1,\"host\":\"cdbb.wonter.net\"}",
			Timestamp: time.Now().UnixNano() / 1e6,
			Sequence:  1,
		},
		{
			Message:   "{\"body_bytes_sent\":14,\"bytes_sent\":408,\"client_ip\":\"120.193.222.39\",\"connection\":91567195,\"hit\":1,\"host\":\"cdbb.wonter.net\"}",
			Timestamp: time.Now().UnixNano() / 1e6,
			Sequence:  2,
		},
	}
	err := BLS_CLIENT.PushLogRecord(EXIST_LOGSTORE, JSON_LOGSTREAM, "JSON", jsonRecords)
	ExpectEqual(t.Errorf, err, nil)
	textRecords := []api.LogRecord{
		{
			Message:   "You know, for test",
			Timestamp: time.Now().UnixNano() / 1e6,
			Sequence:  3,
		},
		{
			Message:   "Baidu Log Service",
			Timestamp: time.Now().UnixNano() / 1e6,
			Sequence:  4,
		},
	}
	err = BLS_CLIENT.PushLogRecord(EXIST_LOGSTORE, TEXT_LOGSTREAM, "TEXT", textRecords)
	ExpectEqual(t.Errorf, err, nil)
	tooNewRecord := []api.LogRecord{
		{
			"LogRecord from future",
			time.Now().UnixNano() / 1e4,
			12,
		},
	}
	err = BLS_CLIENT.PushLogRecord(EXIST_LOGSTORE, TEXT_LOGSTREAM, "TEXT", tooNewRecord)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusBadRequest)
	}
	err = BLS_CLIENT.PushLogRecord("not"+EXIST_LOGSTORE, TEXT_LOGSTREAM, "TEXT", textRecords)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestPullLogRecord(t *testing.T) {
	args := &api.PullLogRecordArgs{
		LogStreamName: JSON_LOGSTREAM,
		StartDateTime: "2021-01-01T10:11:44Z",
		EndDateTime:   "2021-12-10T16:11:44Z",
		Limit:         500,
		Marker:        "",
	}
	res, err := BLS_CLIENT.PullLogRecord(EXIST_LOGSTORE, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(res.Result) >= 2, true)
	res, err = BLS_CLIENT.PullLogRecord("not"+EXIST_LOGSTORE, args)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
		ExpectEqual(t.Errorf, res, nil)
	}
}

func TestQueryLogRecord(t *testing.T) {
	args := &api.QueryLogRecordArgs{
		LogStreamName: JSON_LOGSTREAM,
		Query:         "select count(*)",
		StartDateTime: "2021-01-01T10:11:44Z",
		EndDateTime:   "2021-12-10T16:11:44Z",
	}
	res, err := BLS_CLIENT.QueryLogRecord(EXIST_LOGSTORE, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.ResultSet.Columns[0], "count(*)")
}

// LogStream test
func TestListLogStream(t *testing.T) {
	args := &api.QueryConditions{
		NamePattern: LOGSTREAM_PATTERN,
		Order:       "asc",
		OrderBy:     "creationDateTime",
		PageNo:      1,
		PageSize:    23,
	}
	res, err := BLS_CLIENT.ListLogStream(EXIST_LOGSTORE, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.PageSize, 23)
	res, err = BLS_CLIENT.ListLogStream("not"+EXIST_LOGSTORE, args)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
		ExpectEqual(t.Errorf, res, nil)
	}
}

// FastQuery test
func TestCreateFastQuery(t *testing.T) {
	args := &api.CreateFastQueryBody{
		FastQueryName: FASTQUERY_NAME,
		Query:         "select count(*)",
		Description:   "calculate record number",
		LogStoreName:  EXIST_LOGSTORE,
		LogStreamName: JSON_LOGSTREAM,
	}
	err := BLS_CLIENT.CreateFastQuery(args)
	ExpectEqual(t.Errorf, err, nil)
	// Not specify logStream
	err = BLS_CLIENT.CreateFastQuery(&api.CreateFastQueryBody{
		FastQueryName: FASTQUERY_NAME,
		Query:         "select count(*)",
		Description:   "duplicate",
		LogStoreName:  EXIST_LOGSTORE,
		LogStreamName: "",
	})
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusConflict)
	}
}

func TestDescribeFastQuery(t *testing.T) {
	res, err := BLS_CLIENT.DescribeFastQuery(FASTQUERY_NAME)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.FastQueryName, FASTQUERY_NAME)
	res, err = BLS_CLIENT.DescribeFastQuery("not" + FASTQUERY_NAME)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
		ExpectEqual(t.Errorf, res, nil)
	}
}

func TestUpdateFastQuery(t *testing.T) {
	args := &api.UpdateFastQueryBody{
		Query:         "select * limit 3",
		Description:   "Top 3",
		LogStoreName:  EXIST_LOGSTORE,
		LogStreamName: JSON_LOGSTREAM,
	}
	err := BLS_CLIENT.UpdateFastQuery(FASTQUERY_NAME, args)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BLS_CLIENT.DescribeFastQuery(FASTQUERY_NAME)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Query, "select * limit 3")
	err = BLS_CLIENT.UpdateFastQuery("not"+FASTQUERY_NAME, &api.UpdateFastQueryBody{
		Query:         "select * limit 3",
		Description:   "return top 3 records",
		LogStoreName:  EXIST_LOGSTORE,
		LogStreamName: "",
	})
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestListFastQuery(t *testing.T) {
	args := &api.QueryConditions{
		NamePattern: "s",
	}
	res, err := BLS_CLIENT.ListFastQuery(args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(res.Result) >= 1, true)
}

func TestDeleteFastQuery(t *testing.T) {
	err := BLS_CLIENT.DeleteFastQuery(FASTQUERY_NAME)
	ExpectEqual(t.Errorf, err, nil)
	res, err := BLS_CLIENT.DescribeFastQuery(FASTQUERY_NAME)
	ExpectEqual(t.Errorf, res, nil)
	err = BLS_CLIENT.DeleteFastQuery("not" + FASTQUERY_NAME)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

// Index test
func TestCreateIndex(t *testing.T) {
	fields := map[string]api.LogField{
		"age": {
			Type: "long",
		},
		"salary": {
			Type: "text",
		},
		"name": {
			Type: "object",
			Fields: map[string]api.LogField{
				"firstName": {
					Type: "text",
				},
				"lastName": {
					Type: "text",
				},
			},
		},
	}
	err := BLS_CLIENT.CreateIndex(EXIST_LOGSTORE, true, fields)
	ExpectEqual(t.Errorf, err, nil)
	err = BLS_CLIENT.CreateIndex(EXIST_LOGSTORE, true, fields)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusConflict)
	}
}

func TestUpdateIndex(t *testing.T) {
	fields := map[string]api.LogField{
		"age": {
			Type: "long",
		},
		"wage": {
			Type: "float",
		},
		"name": {
			Type: "object",
			Fields: map[string]api.LogField{
				"firstName": {
					Type: "text",
				},
				"midName": {
					Type: "text",
				},
				"lastName": {
					Type: "text",
				},
			},
		},
	}
	err := BLS_CLIENT.UpdateIndex(EXIST_LOGSTORE, false, fields)
	ExpectEqual(t.Errorf, err, nil)
	err = BLS_CLIENT.UpdateIndex("not"+EXIST_LOGSTORE, false, fields)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestDescribeIndex(t *testing.T) {
	res, err := BLS_CLIENT.DescribeIndex(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Fields["name"].Fields["midName"].Type, "text")
	ExpectEqual(t.Errorf, res.Fields["wage"].Type, "float")
	res, err = BLS_CLIENT.DescribeIndex("not" + EXIST_LOGSTORE)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
		ExpectEqual(t.Errorf, res, nil)
	}
}

func TestDeleteIndex(t *testing.T) {
	err := BLS_CLIENT.DeleteIndex(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, err, nil)
	res, _ := BLS_CLIENT.DescribeIndex(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res, nil)
	err = BLS_CLIENT.DeleteIndex(EXIST_LOGSTORE) // delete twice
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestDeleteLogStore(t *testing.T) {
	res, err := BLS_CLIENT.DescribeLogStore(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.LogStoreName, EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res.Retention, 8)
	err = BLS_CLIENT.DeleteLogStore(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, err, nil)
	res, err = BLS_CLIENT.DescribeLogStore(EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res, nil)
	res, err = BLS_CLIENT.DescribeLogStore("not" + EXIST_LOGSTORE)
	ExpectEqual(t.Errorf, res, nil)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestClient_CreateLogShipper(t *testing.T) {
	args := &api.CreateLogShipperBody{
		LogShipperName: LOGSHIPPER_NAME,
		LogStoreName:   DEAFAULT_LOGSTORE,
		StartTime:      TEST_STARTTIME,
		DestConfig: &api.ShipperDestConfig{
			BOSPath: DEAFAULT_BOSPATH,
		},
	}
	id, err := BLS_CLIENT.CreateLogShipper(args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(id) > 0, true)
	args = &api.CreateLogShipperBody{
		LogShipperName: "invalid",
		LogStoreName:   "not-exist",
		DestConfig:     &api.ShipperDestConfig{},
	}
	id, err = BLS_CLIENT.CreateLogShipper(args)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusBadRequest)
		ExpectEqual(t.Errorf, id, "")
	}
}

func TestClient_UpdateLogShipper(t *testing.T) {
	args := &api.UpdateLogShipperBody{
		LogShipperName: "shipper-sdk",
		DestConfig: &api.ShipperDestConfig{
			PartitionFormatLogStream: true,
			MaxObjectSize:            50,
			CompressType:             "snappy",
			DeliverInterval:          30,
			StorageFormat:            "json",
		},
	}
	err := BLS_CLIENT.UpdateLogShipper(LOGSHIPPER_ID, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestClient_GetLogShipper(t *testing.T) {
	logShipper, err := BLS_CLIENT.GetLogShipper(LOGSHIPPER_ID)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, logShipper.LogStoreName, DEAFAULT_LOGSTORE)
	ExpectEqual(t.Errorf, logShipper.DestConfig.BOSPath, DEAFAULT_BOSPATH)
}

func TestClient_ListLogShipper(t *testing.T) {
	args := &api.ListLogShipperCondition{
		LogShipperID: LOGSHIPPER_ID,
		LogStoreName: DEAFAULT_LOGSTORE,
		Status:       "Running",
	}
	shipperInfo, err := BLS_CLIENT.ListLogShipper(args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, shipperInfo.TotalCount, 1)
}

func TestClient_ListLogShipperRecord(t *testing.T) {
	args := &api.ListShipperRecordCondition{
		SinceHours: 20 * 24,
	}
	records, err := BLS_CLIENT.ListLogShipperRecord(LOGSHIPPER_ID, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, records.TotalCount > 0, true)
	ExpectEqual(t.Errorf, records.Result[0].FinishedCount > 0, true)
}

func TestClient_SetSingleLogShipperStatus(t *testing.T) {
	args := &api.SetSingleShipperStatusCondition{DesiredStatus: "Paused"}
	err := BLS_CLIENT.SetSingleLogShipperStatus(LOGSHIPPER_ID, args)
	ExpectEqual(t.Errorf, err, nil)
	logShipper, err := BLS_CLIENT.GetLogShipper(LOGSHIPPER_ID)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, logShipper.Status, "Paused")
}

func TestClient_BulkSetLogShipperStatus(t *testing.T) {
	args := &api.BulkSetShipperStatusCondition{
		LogShipperIDs: []string{LOGSHIPPER_ID},
		DesiredStatus: "Running",
	}
	err := BLS_CLIENT.BulkSetLogShipperStatus(args)
	ExpectEqual(t.Errorf, err, nil)
	logShipper, err := BLS_CLIENT.GetLogShipper(LOGSHIPPER_ID)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, logShipper.Status, "Running")
}

func TestClient_DeleteSingleLogShipper(t *testing.T) {
	args := &api.CreateLogShipperBody{
		LogShipperName: LOGSHIPPER_NAME,
		LogStoreName:   DEAFAULT_LOGSTORE,
		StartTime:      TEST_STARTTIME,
		DestConfig: &api.ShipperDestConfig{
			BOSPath: DEAFAULT_BOSPATH,
		},
	}
	id, err := BLS_CLIENT.CreateLogShipper(args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(id) > 0, true)
	time.Sleep(2 * time.Second)
	err = BLS_CLIENT.DeleteSingleLogShipper(id)
	ExpectEqual(t.Errorf, err, nil)
	logShipper, err := BLS_CLIENT.GetLogShipper(id)
	ExpectEqual(t.Errorf, logShipper, nil)
	if realErr, ok := err.(*bce.BceServiceError); ok {
		ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
	}
}

func TestClient_BulkDeleteLogShipper(t *testing.T) {
	var ids []string
	for i := 0; i < 3; i++ {
		args := &api.CreateLogShipperBody{
			LogShipperName: LOGSHIPPER_NAME,
			LogStoreName:   DEAFAULT_LOGSTORE,
			StartTime:      TEST_STARTTIME,
			DestConfig: &api.ShipperDestConfig{
				BOSPath: DEAFAULT_BOSPATH,
			},
		}
		id, err := BLS_CLIENT.CreateLogShipper(args)
		ExpectEqual(t.Errorf, err, nil)
		ExpectEqual(t.Errorf, len(id) > 0, true)
		ids = append(ids, id)
	}
	time.Sleep(time.Second * 2)
	args := &api.BulkDeleteShipperCondition{LogShipperIDs: ids}
	err := BLS_CLIENT.BulkDeleteLogShipper(args)
	ExpectEqual(t.Errorf, err, nil)
	for _, id := range ids {
		logShipper, err := BLS_CLIENT.GetLogShipper(id)
		ExpectEqual(t.Errorf, logShipper, nil)
		if realErr, ok := err.(*bce.BceServiceError); ok {
			ExpectEqual(t.Errorf, realErr.StatusCode, http.StatusNotFound)
		}
	}
}
