package bls

import (
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bls/api"
)

func TestProject(t *testing.T) {
	createProjectRequest := CreateProjectRequest{
		Name: "sdk-project-test",
	}
	err := BLS_CLIENT.CreateProject(createProjectRequest)
	ExpectEqual(t.Errorf, err, nil)

	listProjectRequest := ListProjectRequest{
		Name: createProjectRequest.Name,
	}
	pr, err := BLS_CLIENT.ListProject(listProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(pr.Projects))
	ExpectEqual(t.Errorf, createProjectRequest.Name, pr.Projects[0].Name)
	ExpectEqual(t.Errorf, "", pr.Projects[0].Description)
	ExpectEqual(t.Errorf, false, pr.Projects[0].Top)

	updateProjectRequest := UpdateProjectRequest{
		UUID:        pr.Projects[0].UUID,
		Description: "test",
		Top:         true,
	}
	err = BLS_CLIENT.UpdateProject(updateProjectRequest)
	ExpectEqual(t.Errorf, err, nil)

	describeProjectRequest := DescribeProjectRequest{
		UUID: pr.Projects[0].UUID,
	}
	p, err := BLS_CLIENT.DescribeProject(describeProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, createProjectRequest.Name, p.Name)
	ExpectEqual(t.Errorf, "test", p.Description)
	ExpectEqual(t.Errorf, true, p.Top)

	deleteProjectRequest := DeleteProjectRequest{
		UUID: pr.Projects[0].UUID,
	}
	err = BLS_CLIENT.DeleteProject(deleteProjectRequest)
	ExpectEqual(t.Errorf, err, nil)

	p, err = BLS_CLIENT.DescribeProject(describeProjectRequest)
	ExpectEqual(t.Errorf, nil, p)
	ExpectEqual(t.Errorf, "ProjectNotFound", (err.(*bce.BceServiceError)).Code)
}

func TestLogStoreAndIndexV2(t *testing.T) {
	createLogStoreRequest := CreateLogStoreRequest{
		Project:      DefaultProject,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err := BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	// test index api
	createIndexRequest := CreateIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
		Fulltext:     true,
		Fields: map[string]api.LogField{
			"test1": {
				Type: "float",
			},
		},
	}
	err = BLS_CLIENT.CreateIndexV2(createIndexRequest)
	ExpectEqual(t.Errorf, err, nil)

	updateIndexRequest := UpdateIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
		Fulltext:     true,
		Fields: map[string]api.LogField{
			"test1": {
				Type: "float",
			},
			"test2": {
				Type: "bool",
			},
		},
	}
	err = BLS_CLIENT.UpdateIndexV2(updateIndexRequest)
	ExpectEqual(t.Errorf, err, nil)

	describeIndexRequest := DescribeIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	idx, err := BLS_CLIENT.DescribeIndexV2(describeIndexRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, true, idx.FullText)
	ExpectEqual(t.Errorf, 2, len(idx.Fields))
	ExpectEqual(t.Errorf, "float", idx.Fields["test1"].Type)
	ExpectEqual(t.Errorf, "bool", idx.Fields["test2"].Type)

	deleteIndexRequest := DeleteIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteIndexV2(deleteIndexRequest)
	ExpectEqual(t.Errorf, err, nil)
	idx, err = BLS_CLIENT.DescribeIndexV2(describeIndexRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, false, idx.FullText)
	ExpectEqual(t.Errorf, 0, len(idx.Fields))

	// test logstore api
	updateLogStoreRequest := UpdateLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
		Retention:    2,
	}
	err = BLS_CLIENT.UpdateLogStoreV2(updateLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	describeLogStoreRequest := DescribeLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	ls, err := BLS_CLIENT.DescribeLogStoreV2(describeLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, createLogStoreRequest.Project, ls.Project)
	ExpectEqual(t.Errorf, createLogStoreRequest.LogStoreName, ls.LogStoreName)
	ExpectEqual(t.Errorf, updateLogStoreRequest.Retention, ls.Retention)
	ExpectEqual(t.Errorf, 0, len(ls.Tags))

	listLogStoreRequest := ListLogStoreRequest{
		Project:     createLogStoreRequest.Project,
		NamePattern: createLogStoreRequest.LogStoreName,
	}
	lr, err := BLS_CLIENT.ListLogStoreV2(listLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(lr.Result))

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	ls, err = BLS_CLIENT.DescribeLogStoreV2(describeLogStoreRequest)
	ExpectEqual(t.Errorf, nil, ls)
	ExpectEqual(t.Errorf, "LogStoreNotFound", (err.(*bce.BceServiceError)).Code)
}

func TestLogStoreAndIndexV2WithProject(t *testing.T) {
	createProjectRequest := CreateProjectRequest{
		Name: "sdk-project-test",
	}
	err := BLS_CLIENT.CreateProject(createProjectRequest)
	ExpectEqual(t.Errorf, err, nil)

	listProjectRequest := ListProjectRequest{
		Name: createProjectRequest.Name,
	}
	pr, err := BLS_CLIENT.ListProject(listProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(pr.Projects))
	ExpectEqual(t.Errorf, createProjectRequest.Name, pr.Projects[0].Name)

	createLogStoreRequest := CreateLogStoreRequest{
		Project:      createProjectRequest.Name,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err = BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	// test index api
	createIndexRequest := CreateIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
		Fulltext:     true,
		Fields: map[string]api.LogField{
			"test1": {
				Type: "float",
			},
		},
	}
	err = BLS_CLIENT.CreateIndexV2(createIndexRequest)
	ExpectEqual(t.Errorf, err, nil)

	updateIndexRequest := UpdateIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
		Fulltext:     true,
		Fields: map[string]api.LogField{
			"test1": {
				Type: "float",
			},
			"test2": {
				Type: "bool",
			},
		},
	}
	err = BLS_CLIENT.UpdateIndexV2(updateIndexRequest)
	ExpectEqual(t.Errorf, err, nil)

	describeIndexRequest := DescribeIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	idx, err := BLS_CLIENT.DescribeIndexV2(describeIndexRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, true, idx.FullText)
	ExpectEqual(t.Errorf, 2, len(idx.Fields))
	ExpectEqual(t.Errorf, "float", idx.Fields["test1"].Type)
	ExpectEqual(t.Errorf, "bool", idx.Fields["test2"].Type)

	deleteIndexRequest := DeleteIndexRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteIndexV2(deleteIndexRequest)
	ExpectEqual(t.Errorf, err, nil)
	idx, err = BLS_CLIENT.DescribeIndexV2(describeIndexRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, false, idx.FullText)
	ExpectEqual(t.Errorf, 0, len(idx.Fields))

	// test logstore api
	updateLogStoreRequest := UpdateLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
		Retention:    2,
	}
	err = BLS_CLIENT.UpdateLogStoreV2(updateLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	describeLogStoreRequest := DescribeLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	ls, err := BLS_CLIENT.DescribeLogStoreV2(describeLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, createLogStoreRequest.Project, ls.Project)
	ExpectEqual(t.Errorf, createLogStoreRequest.LogStoreName, ls.LogStoreName)
	ExpectEqual(t.Errorf, updateLogStoreRequest.Retention, ls.Retention)
	ExpectEqual(t.Errorf, 0, len(ls.Tags))

	listLogStoreRequest := ListLogStoreRequest{
		Project:     createLogStoreRequest.Project,
		NamePattern: createLogStoreRequest.LogStoreName,
	}
	lr, err := BLS_CLIENT.ListLogStoreV2(listLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(lr.Result))

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	ls, err = BLS_CLIENT.DescribeLogStoreV2(describeLogStoreRequest)
	ExpectEqual(t.Errorf, nil, ls)
	ExpectEqual(t.Errorf, "LogStoreNotFound", (err.(*bce.BceServiceError)).Code)

	deleteProjectRequest := DeleteProjectRequest{
		UUID: pr.Projects[0].UUID,
	}
	err = BLS_CLIENT.DeleteProject(deleteProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRecordAndStreamV2(t *testing.T) {
	createLogStoreRequest := CreateLogStoreRequest{
		Project:      DefaultProject,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err := BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	time.Sleep(2 * time.Second)
	pushLogRecordRequest := PushLogRecordRequest{
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		LogStreamName: "sdk-stream-test",
		LogType:       "JSON",
		LogRecords: []api.LogRecord{
			{
				Message:   "{\"@raw\": \"raw text\", \"level\": \"info\"}",
				Timestamp: time.Now().UnixMilli(),
			},
		},
	}
	err = BLS_CLIENT.PushLogRecordV2(pushLogRecordRequest)
	ExpectEqual(t.Errorf, err, nil)

	// 等待一段时间再查询
	time.Sleep(18 * time.Second)

	listLogStreamRequest := ListLogStreamRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	sr, err := BLS_CLIENT.ListLogStreamV2(listLogStreamRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(sr.Result))
	ExpectEqual(t.Errorf, pushLogRecordRequest.LogStreamName, sr.Result[0].LogStreamName)

	pullLogRecordRequest := PullLogRecordRequest{
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		StartDateTime: time.Now().Add(-10 * time.Minute).UTC().Format("2006-01-02T15:04:05Z"),
		EndDateTime:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	rr, err := BLS_CLIENT.PullLogRecordV2(pullLogRecordRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(rr.Result))
	ExpectEqual(t.Errorf, "{\"@raw\":\"raw text\",\"level\":\"info\"}", rr.Result[0].Message)

	queryLogRecordRequest := QueryLogRecordRequest{
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		Query:         "match *",
		StartDateTime: time.Now().Add(-20 * time.Second).UTC().Format("2006-01-02T15:04:05Z"),
		EndDateTime:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	qr, err := BLS_CLIENT.QueryLogRecordV2(queryLogRecordRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(qr.ResultSet.Rows))

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRecordAndStreamV2WithProject(t *testing.T) {
	createProjectRequest := CreateProjectRequest{
		Name: "sdk-project-test",
	}
	err := BLS_CLIENT.CreateProject(createProjectRequest)
	ExpectEqual(t.Errorf, err, nil)

	listProjectRequest := ListProjectRequest{
		Name: createProjectRequest.Name,
	}
	pr, err := BLS_CLIENT.ListProject(listProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(pr.Projects))
	ExpectEqual(t.Errorf, createProjectRequest.Name, pr.Projects[0].Name)

	createLogStoreRequest := CreateLogStoreRequest{
		Project:      createProjectRequest.Name,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err = BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	time.Sleep(2 * time.Second)
	pushLogRecordRequest := PushLogRecordRequest{
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		LogStreamName: "sdk-stream-test",
		LogType:       "JSON",
		LogRecords: []api.LogRecord{
			{
				Message:   "{\"@raw\": \"raw text\", \"level\": \"info\"}",
				Timestamp: time.Now().UnixMilli(),
			},
		},
	}
	err = BLS_CLIENT.PushLogRecordV2(pushLogRecordRequest)
	ExpectEqual(t.Errorf, err, nil)

	// 等待一段时间再查询
	time.Sleep(18 * time.Second)

	listLogStreamRequest := ListLogStreamRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	sr, err := BLS_CLIENT.ListLogStreamV2(listLogStreamRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(sr.Result))
	ExpectEqual(t.Errorf, pushLogRecordRequest.LogStreamName, sr.Result[0].LogStreamName)

	pullLogRecordRequest := PullLogRecordRequest{
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		StartDateTime: time.Now().Add(-20 * time.Second).UTC().Format("2006-01-02T15:04:05Z"),
		EndDateTime:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	rr, err := BLS_CLIENT.PullLogRecordV2(pullLogRecordRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(rr.Result))
	ExpectEqual(t.Errorf, "{\"@raw\":\"raw text\",\"level\":\"info\"}", rr.Result[0].Message)

	queryLogRecordRequest := QueryLogRecordRequest{
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		Query:         "match *",
		StartDateTime: time.Now().Add(-1 * time.Minute).UTC().Format("2006-01-02T15:04:05Z"),
		EndDateTime:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	qr, err := BLS_CLIENT.QueryLogRecordV2(queryLogRecordRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(qr.ResultSet.Rows))

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	deleteProjectRequest := DeleteProjectRequest{
		UUID: pr.Projects[0].UUID,
	}
	err = BLS_CLIENT.DeleteProject(deleteProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
}

func TestFastQueryV2(t *testing.T) {
	createLogStoreRequest := CreateLogStoreRequest{
		Project:      DefaultProject,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err := BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	createFastQueryRequest := CreateFastQueryRequest{
		FastQueryName: "sdk-fast-query-test",
		Query:         "match *",
		Description:   "",
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		StartDateTime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		EndDateTime:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	err = BLS_CLIENT.CreateFastQueryV2(createFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)

	updateFastQueryRequest := UpdateFastQueryRequest{
		FastQueryName: createFastQueryRequest.FastQueryName,
		Query:         "select * limit 10",
		Description:   "test",
		Project:       createFastQueryRequest.Project,
		LogStoreName:  createFastQueryRequest.LogStoreName,
		LogStreamName: createFastQueryRequest.LogStreamName,
		StartDateTime: createFastQueryRequest.StartDateTime,
		EndDateTime:   createFastQueryRequest.EndDateTime,
	}
	err = BLS_CLIENT.UpdateFastQueryV2(updateFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)

	describeFastQueryRequest := DescribeFastQueryRequest{
		FastQueryName: createFastQueryRequest.FastQueryName,
	}
	fq, err := BLS_CLIENT.DescribeFastQueryV2(describeFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, createFastQueryRequest.FastQueryName, fq.FastQueryName)
	ExpectEqual(t.Errorf, updateFastQueryRequest.Query, fq.Query)
	ExpectEqual(t.Errorf, updateFastQueryRequest.Description, fq.Description)
	ExpectEqual(t.Errorf, updateFastQueryRequest.Project, fq.Project)
	ExpectEqual(t.Errorf, updateFastQueryRequest.LogStoreName, fq.LogStoreName)
	ExpectEqual(t.Errorf, updateFastQueryRequest.LogStreamName, fq.LogStreamName)
	ExpectEqual(t.Errorf, updateFastQueryRequest.StartDateTime, fq.StartDateTime)
	ExpectEqual(t.Errorf, updateFastQueryRequest.EndDateTime, fq.EndDateTime)

	listFastQueryRequest := ListFastQueryRequest{
		Project:     createFastQueryRequest.Project,
		NamePattern: createFastQueryRequest.FastQueryName,
	}
	fr, err := BLS_CLIENT.ListFastQueryV2(listFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(fr.Result))
	ExpectEqual(t.Errorf, createFastQueryRequest.FastQueryName, fr.Result[0].FastQueryName)

	deleteFastQueryRequest := DeleteFastQueryRequest{
		FastQueryName: createFastQueryRequest.FastQueryName,
	}
	err = BLS_CLIENT.DeleteFastQueryV2(deleteFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)

	fq, err = BLS_CLIENT.DescribeFastQueryV2(describeFastQueryRequest)
	ExpectEqual(t.Errorf, nil, fq)
	ExpectEqual(t.Errorf, "FastQueryNotFound", (err.(*bce.BceServiceError)).Code)

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)
}

func TestFastQueryV2WithProject(t *testing.T) {
	createProjectRequest := CreateProjectRequest{
		Name: "sdk-project-test",
	}
	err := BLS_CLIENT.CreateProject(createProjectRequest)
	ExpectEqual(t.Errorf, err, nil)

	listProjectRequest := ListProjectRequest{
		Name: createProjectRequest.Name,
	}
	pr, err := BLS_CLIENT.ListProject(listProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(pr.Projects))
	ExpectEqual(t.Errorf, createProjectRequest.Name, pr.Projects[0].Name)

	createLogStoreRequest := CreateLogStoreRequest{
		Project:      createProjectRequest.Name,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err = BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	createFastQueryRequest := CreateFastQueryRequest{
		FastQueryName: "sdk-fast-query-test",
		Query:         "match *",
		Description:   "",
		Project:       createLogStoreRequest.Project,
		LogStoreName:  createLogStoreRequest.LogStoreName,
		StartDateTime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		EndDateTime:   time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	err = BLS_CLIENT.CreateFastQueryV2(createFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)

	updateFastQueryRequest := UpdateFastQueryRequest{
		FastQueryName: createFastQueryRequest.FastQueryName,
		Query:         "select * limit 10",
		Description:   "test",
		Project:       createFastQueryRequest.Project,
		LogStoreName:  createFastQueryRequest.LogStoreName,
		LogStreamName: createFastQueryRequest.LogStreamName,
		StartDateTime: createFastQueryRequest.StartDateTime,
		EndDateTime:   createFastQueryRequest.EndDateTime,
	}
	err = BLS_CLIENT.UpdateFastQueryV2(updateFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)

	describeFastQueryRequest := DescribeFastQueryRequest{
		FastQueryName: createFastQueryRequest.FastQueryName,
	}
	fq, err := BLS_CLIENT.DescribeFastQueryV2(describeFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, createFastQueryRequest.FastQueryName, fq.FastQueryName)
	ExpectEqual(t.Errorf, updateFastQueryRequest.Query, fq.Query)
	ExpectEqual(t.Errorf, updateFastQueryRequest.Description, fq.Description)
	ExpectEqual(t.Errorf, updateFastQueryRequest.Project, fq.Project)
	ExpectEqual(t.Errorf, updateFastQueryRequest.LogStoreName, fq.LogStoreName)
	ExpectEqual(t.Errorf, updateFastQueryRequest.LogStreamName, fq.LogStreamName)
	ExpectEqual(t.Errorf, updateFastQueryRequest.StartDateTime, fq.StartDateTime)
	ExpectEqual(t.Errorf, updateFastQueryRequest.EndDateTime, fq.EndDateTime)

	listFastQueryRequest := ListFastQueryRequest{
		Project:     createFastQueryRequest.Project,
		NamePattern: createFastQueryRequest.FastQueryName,
	}
	fr, err := BLS_CLIENT.ListFastQueryV2(listFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(fr.Result))
	ExpectEqual(t.Errorf, createFastQueryRequest.FastQueryName, fr.Result[0].FastQueryName)

	deleteFastQueryRequest := DeleteFastQueryRequest{
		FastQueryName: createFastQueryRequest.FastQueryName,
	}
	err = BLS_CLIENT.DeleteFastQueryV2(deleteFastQueryRequest)
	ExpectEqual(t.Errorf, err, nil)

	fq, err = BLS_CLIENT.DescribeFastQueryV2(describeFastQueryRequest)
	ExpectEqual(t.Errorf, nil, fq)
	ExpectEqual(t.Errorf, "FastQueryNotFound", (err.(*bce.BceServiceError)).Code)

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	deleteProjectRequest := DeleteProjectRequest{
		UUID: pr.Projects[0].UUID,
	}
	err = BLS_CLIENT.DeleteProject(deleteProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
}

func TestLogShipperV2(t *testing.T) {
	createLogStoreRequest := CreateLogStoreRequest{
		Project:      DefaultProject,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err := BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	createLogShipperRequest := CreateLogShipperRequest{
		LogShipperName: "sdk-log-shipper-test",
		Project:        createLogStoreRequest.Project,
		LogStoreName:   createLogStoreRequest.LogStoreName,
		StartTime:      time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		DestType:       "BOS",
		DestConfig: &api.ShipperDestConfig{
			BOSPath:                  "bls-test/sdk-log-shipper-test/",
			PartitionFormatTS:        "%Y/%m/%d/%H/%M/",
			PartitionFormatLogStream: false,
			MaxObjectSize:            64,
			CompressType:             "none",
			DeliverInterval:          5,
			StorageFormat:            "json",
			ShipperType:              "text",
		},
	}
	logShipperID, err := BLS_CLIENT.CreateLogShipperV2(createLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, true, len(logShipperID) > 0)

	listShipperRecordRequest := ListShipperRecordRequest{
		LogShipperID: logShipperID,
	}
	rr, err := BLS_CLIENT.ListLogShipperRecordV2(listShipperRecordRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 0, len(rr.Result))

	updateLogShipperRequest := UpdateLogShipperRequest{
		LogShipperID:   logShipperID,
		LogShipperName: createLogShipperRequest.LogShipperName,
		DestConfig: &api.ShipperDestConfig{
			BOSPath:                  "bls-test/sdk-log-shipper-test/",
			PartitionFormatTS:        "%Y/%m/%d/%H/%M/",
			PartitionFormatLogStream: false,
			MaxObjectSize:            128,
			CompressType:             "none",
			DeliverInterval:          10,
			StorageFormat:            "JSON",
		},
	}
	err = BLS_CLIENT.UpdateLogShipperV2(updateLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)

	getLogShipperRequest := GetLogShipperRequest{
		LogShipperID: logShipperID,
	}
	s, err := BLS_CLIENT.GetLogShipperV2(getLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, updateLogShipperRequest.LogShipperName, s.LogShipperName)
	ExpectEqual(t.Errorf, createLogShipperRequest.Project, s.Project)
	ExpectEqual(t.Errorf, createLogShipperRequest.LogStoreName, s.LogStoreName)
	ExpectEqual(t.Errorf, createLogShipperRequest.DestType, s.DestType)
	ExpectEqual(t.Errorf, createLogShipperRequest.StartTime, s.StartTime)
	ExpectEqual(t.Errorf, "Running", s.Status)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.BOSPath, s.DestConfig.BOSPath)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.PartitionFormatTS, s.DestConfig.PartitionFormatTS)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.PartitionFormatLogStream, s.DestConfig.PartitionFormatLogStream)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.MaxObjectSize, s.DestConfig.MaxObjectSize)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.CompressType, s.DestConfig.CompressType)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.DeliverInterval, s.DestConfig.DeliverInterval)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.StorageFormat, s.DestConfig.StorageFormat)

	listLogShipperRequest := ListLogShipperRequest{
		Project:      createLogShipperRequest.Project,
		LogStoreName: createLogShipperRequest.LogStoreName,
		PageNo:       1,
		PageSize:     10,
	}
	sr, err := BLS_CLIENT.ListLogShipperV2(listLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(sr.Result))
	ExpectEqual(t.Errorf, logShipperID, sr.Result[0].LogShipperID)
	ExpectEqual(t.Errorf, createLogShipperRequest.LogShipperName, sr.Result[0].LogShipperName)

	updateLogShipperStatusRequest := UpdateLogShipperStatusRequest{
		LogShipperID:  logShipperID,
		DesiredStatus: "Paused",
	}
	err = BLS_CLIENT.UpdateLogShipperStatusV2(updateLogShipperStatusRequest)
	ExpectEqual(t.Errorf, err, nil)

	s, err = BLS_CLIENT.GetLogShipperV2(getLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, updateLogShipperRequest.LogShipperName, s.LogShipperName)
	ExpectEqual(t.Errorf, "Paused", s.Status)

	deleteLogShipperRequest := DeleteLogShipperRequest{
		LogShipperID: logShipperID,
	}
	err = BLS_CLIENT.DeleteLogShipperV2(deleteLogShipperRequest)
	ExpectEqual(t.Errorf, nil, err)

	s, err = BLS_CLIENT.GetLogShipperV2(getLogShipperRequest)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Deleted", s.Status)

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)
}

func TestLogShipperV2WithProject(t *testing.T) {
	createProjectRequest := CreateProjectRequest{
		Name: "sdk-project-test",
	}
	err := BLS_CLIENT.CreateProject(createProjectRequest)
	ExpectEqual(t.Errorf, err, nil)

	listProjectRequest := ListProjectRequest{
		Name: createProjectRequest.Name,
	}
	pr, err := BLS_CLIENT.ListProject(listProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(pr.Projects))
	ExpectEqual(t.Errorf, createProjectRequest.Name, pr.Projects[0].Name)

	createLogStoreRequest := CreateLogStoreRequest{
		Project:      createProjectRequest.Name,
		LogStoreName: "sdk-logstore-test",
		Retention:    1,
	}
	err = BLS_CLIENT.CreateLogStoreV2(createLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	createLogShipperRequest := CreateLogShipperRequest{
		LogShipperName: "sdk-log-shipper-test",
		Project:        createLogStoreRequest.Project,
		LogStoreName:   createLogStoreRequest.LogStoreName,
		StartTime:      time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		DestType:       "BOS",
		DestConfig: &api.ShipperDestConfig{
			BOSPath:                  "bls-test/sdk-log-shipper-test/",
			PartitionFormatTS:        "%Y/%m/%d/%H/%M/",
			PartitionFormatLogStream: false,
			MaxObjectSize:            64,
			CompressType:             "none",
			DeliverInterval:          5,
			StorageFormat:            "json",
			ShipperType:              "text",
		},
	}
	logShipperID, err := BLS_CLIENT.CreateLogShipperV2(createLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, true, len(logShipperID) > 0)

	listShipperRecordRequest := ListShipperRecordRequest{
		LogShipperID: logShipperID,
	}
	rr, err := BLS_CLIENT.ListLogShipperRecordV2(listShipperRecordRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 0, len(rr.Result))

	updateLogShipperRequest := UpdateLogShipperRequest{
		LogShipperID:   logShipperID,
		LogShipperName: createLogShipperRequest.LogShipperName,
		DestConfig: &api.ShipperDestConfig{
			BOSPath:                  "bls-test/sdk-log-shipper-test/",
			PartitionFormatTS:        "%Y/%m/%d/%H/%M/",
			PartitionFormatLogStream: false,
			MaxObjectSize:            128,
			CompressType:             "none",
			DeliverInterval:          10,
			StorageFormat:            "JSON",
		},
	}
	err = BLS_CLIENT.UpdateLogShipperV2(updateLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)

	getLogShipperRequest := GetLogShipperRequest{
		LogShipperID: logShipperID,
	}
	s, err := BLS_CLIENT.GetLogShipperV2(getLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, updateLogShipperRequest.LogShipperName, s.LogShipperName)
	ExpectEqual(t.Errorf, createLogShipperRequest.Project, s.Project)
	ExpectEqual(t.Errorf, createLogShipperRequest.LogStoreName, s.LogStoreName)
	ExpectEqual(t.Errorf, createLogShipperRequest.DestType, s.DestType)
	ExpectEqual(t.Errorf, createLogShipperRequest.StartTime, s.StartTime)
	ExpectEqual(t.Errorf, "Running", s.Status)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.BOSPath, s.DestConfig.BOSPath)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.PartitionFormatTS, s.DestConfig.PartitionFormatTS)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.PartitionFormatLogStream, s.DestConfig.PartitionFormatLogStream)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.MaxObjectSize, s.DestConfig.MaxObjectSize)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.CompressType, s.DestConfig.CompressType)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.DeliverInterval, s.DestConfig.DeliverInterval)
	ExpectEqual(t.Errorf, updateLogShipperRequest.DestConfig.StorageFormat, s.DestConfig.StorageFormat)

	listLogShipperRequest := ListLogShipperRequest{
		Project:      createLogShipperRequest.Project,
		LogStoreName: createLogShipperRequest.LogStoreName,
		PageNo:       1,
		PageSize:     10,
	}
	sr, err := BLS_CLIENT.ListLogShipperV2(listLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 1, len(sr.Result))
	ExpectEqual(t.Errorf, logShipperID, sr.Result[0].LogShipperID)
	ExpectEqual(t.Errorf, createLogShipperRequest.LogShipperName, sr.Result[0].LogShipperName)

	bulkUpdateLogShipperStatusRequest := BulkUpdateLogShipperStatusRequest{
		LogShipperIDs: []string{logShipperID},
		DesiredStatus: "Paused",
	}
	err = BLS_CLIENT.BulkUpdateLogShipperStatusV2(bulkUpdateLogShipperStatusRequest)
	ExpectEqual(t.Errorf, err, nil)

	s, err = BLS_CLIENT.GetLogShipperV2(getLogShipperRequest)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, updateLogShipperRequest.LogShipperName, s.LogShipperName)
	ExpectEqual(t.Errorf, "Paused", s.Status)

	bulkDeleteLogShipperRequest := BulkDeleteLogShipperRequest{
		LogShipperIDs: []string{logShipperID},
	}
	err = BLS_CLIENT.BulkDeleteLogShipperV2(bulkDeleteLogShipperRequest)
	ExpectEqual(t.Errorf, nil, err)

	s, err = BLS_CLIENT.GetLogShipperV2(getLogShipperRequest)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Deleted", s.Status)

	deleteLogStoreRequest := DeleteLogStoreRequest{
		Project:      createLogStoreRequest.Project,
		LogStoreName: createLogStoreRequest.LogStoreName,
	}
	err = BLS_CLIENT.DeleteLogStoreV2(deleteLogStoreRequest)
	ExpectEqual(t.Errorf, err, nil)

	deleteProjectRequest := DeleteProjectRequest{
		UUID: pr.Projects[0].UUID,
	}
	err = BLS_CLIENT.DeleteProject(deleteProjectRequest)
	ExpectEqual(t.Errorf, err, nil)
}
