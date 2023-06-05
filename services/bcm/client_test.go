package bcm

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	bcmClient *Client
	bcmConf   *Conf
)

type Conf struct {
	AK         string `json:"AK"`
	SK         string `json:"SK"`
	Endpoint   string `json:"Endpoint"`
	InstanceId string `json:"InstanceId"`
	UserId     string `json:"UserId"`
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	bcmConf = &Conf{}
	_ = decoder.Decode(bcmConf)

	bcmClient, _ = NewClient(bcmConf.AK, bcmConf.SK, bcmConf.Endpoint)
	log.SetLogLevel(log.WARN)
}

func TestClient_GetMetricData(t *testing.T) {
	dimensions := map[string]string{
		"InstanceId": bcmConf.InstanceId,
	}
	req := &GetMetricDataRequest{
		UserId:         bcmConf.UserId,
		Scope:          "BCE_BCC",
		MetricName:     "vCPUUsagePercent",
		Dimensions:     dimensions,
		Statistics:     strings.Split(Average+","+SampleCount+","+Sum+","+Minimum+","+Maximum, ","),
		PeriodInSecond: 60,
		StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
		EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
	}
	resp, err := bcmClient.GetMetricData(req)
	if err != nil {
		t.Errorf("bcm get metric data error with %v\n", err)
	}
	if resp.Code != "OK" {
		t.Errorf("bcm get metric data response code error with %v\n", resp.Code)
	}
	if len(resp.DataPoints) < 1 {
		t.Error("bcm get metric data response dataPoints size not be greater 0\n")
	} else {
		if resp.DataPoints[0].Sum <= 0 {
			t.Error("bcm get metric data response dataPoints[0] sum not be greater 0\n")
		}
		if resp.DataPoints[0].Average <= 0 {
			t.Error("bcm get metric data response dataPoints[0] average not be greater 0\n")
		}
		if resp.DataPoints[0].SampleCount <= 0 {
			t.Error("bcm get metric data response dataPoints[0] sampleCount not be greater 0\n")
		}
		if resp.DataPoints[0].Minimum <= 0 {
			t.Error("bcm get metric data response dataPoints[0] minimum not be greater 0\n")
		}
		if resp.DataPoints[0].Maximum <= 0 {
			t.Error("bcm get metric data response dataPoints[0] maximum not be greater 0\n")
		}
	}
}

func TestClient_BatchGetMetricData(t *testing.T) {
	dimensions := map[string]string{
		"InstanceId": bcmConf.InstanceId,
	}
	req := &BatchGetMetricDataRequest{
		UserId:         bcmConf.UserId,
		Scope:          "BCE_BCC",
		MetricNames:    []string{"vCPUUsagePercent", "CpuIdlePercent"},
		Dimensions:     dimensions,
		Statistics:     strings.Split(Average+","+SampleCount+","+Sum+","+Minimum+","+Maximum, ","),
		PeriodInSecond: 60,
		StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
		EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
	}
	resp, err := bcmClient.BatchGetMetricData(req)
	if err != nil {
		t.Errorf("bcm batch get metric data error with %v\n", err)
	}
	if resp.Code != "OK" {
		t.Errorf("bcm batch get metric data response code error with %v\n", resp.Code)
	}
	if len(resp.ErrorList) > 0 {
		t.Error("bcm batch get metric data response errorList size not be lower 1\n")
	}
	if len(resp.SuccessList) < 2 {
		t.Error("bcm batch get metric data response successList size not be greater 1\n")
	} else {
		if len(resp.SuccessList[0].DataPoints) <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints size not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Sum <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] sum not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Average <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] average not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].SampleCount <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] sampleCount not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Minimum <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] minimum not be greater 0\n")
		}
		if resp.SuccessList[0].DataPoints[0].Maximum <= 0 {
			t.Error("bcm batch get metric data response successList dataPoints[0] maximum not be greater 0\n")
		}
	}

}
