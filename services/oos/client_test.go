package oos

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/oos/model"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	oosClient *Client
	oosConf   *Conf
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
	oosConf = &Conf{}
	_ = decoder.Decode(oosConf)

	oosClient, _ = NewClient(oosConf.AK, oosConf.SK, oosConf.Endpoint)
	log.SetLogLevel(log.WARN)
}

func getTemplate(des string) *model.Template {
	template := &model.Template{
		Name:        "test_template_011",
		Description: des,
		Tags:        nil,
		Operators: []*model.Operator{
			{
				Name:          "stop_bcc",
				Description:   "停止BCC实例",
				Operator:      "BCE::BCC::StopInstance",
				Retries:       3,
				RetryInterval: 60000,
				Timeout:       3600000,
				ParallelismControl: &model.RateControl{
					Ratio: 0,
					Count: 1,
				},
				Properties: map[string]interface{}{
					"instances": []map[string]string{
						{
							"instanceId": "i-jOS24U0l",
							"name":       "instance-3nd0wnlr",
						},
					},
				},
			},
			{
				Name:        "start_bcc",
				Description: "启动BCC实例",
				Operator:    "BCE::BCC::StartInstance",
				Properties: map[string]interface{}{
					"instances": []map[string]string{
						{
							"instanceId": "i-jOS24U0l",
							"name":       "instance-3nd0wnlr",
						},
					},
				},
			},
		},
		Properties: []*model.Property{
			{
				Name:     "test_param",
				Type:     "string",
				Required: false,
			},
		},
		Linear: true,
	}
	return template
}

func Test_CreateTemplate(t *testing.T) {
	req := getTemplate("创建模板测试")
	result, err := oosClient.CreateTemplate(req)
	if err != nil || result.Code != 200 {
		t.Errorf("oos create template error with %v\n", err)
	}
	fmt.Println(result.Result.ID)
}

func Test_CheckTemplate(t *testing.T) {
	req := getTemplate("校验模板测试")
	result, err := oosClient.CheckTemplate(req)
	if err != nil || result.Code != 200 {
		t.Errorf("oos check template error with %v\n", err)
	}
}

func Test_UpdateTemplate(t *testing.T) {
	req := getTemplate("更改模板测试")
	req.ID = "tpl-nXzDX2Bn"
	result, err := oosClient.UpdateTemplate(req)
	if err != nil || result.Code != 200 {
		t.Errorf("oos update template error with %v\n", err)
	}
}

func Test_DeleteTemplate(t *testing.T) {
	result, err := oosClient.DeleteTemplate("tpl-nXzDX2Bn")
	if err != nil || result.Code != 200 {
		t.Errorf("oos delete template error with %v\n", err)
	}
}

func Test_GetTemplateDetail(t *testing.T) {
	result, err := oosClient.GetTemplateDetail("test_template_011", string(model.TemplateTypeIndividual))
	if err != nil || result.Code != 200 {
		t.Errorf("oos get template detail error with %v\n", err)
	}
}

func Test_GetTemplateList(t *testing.T) {
	req := &model.GetTemplateListRequest{
		BasePageRequest: model.BasePageRequest{
			Ascending: false,
			PageNo:    1,
			PageSize:  100,
		},
	}
	_, err := oosClient.GetTemplateList(req)
	if err != nil {
		t.Errorf("oos get template list error with %v\n", err)
	}
}

func Test_GetOperatorList(t *testing.T) {
	req := &model.BasePageRequest{
		Ascending: false,
		PageNo:    1,
		PageSize:  100,
	}
	_, err := oosClient.GetOperatorList(req)
	if err != nil {
		t.Errorf("oos get operator list error with %v\n", err)
	}
}

func Test_CreateExecution1(t *testing.T) {
	req := &model.Execution{
		Description: "创建执行测试1",
		Template: &model.Template{
			Ref:    "tpl-nXzDX2Bn",
			Linear: true,
		},
	}
	_, err := oosClient.CreateExecution(req)
	if err != nil {
		t.Errorf("oos create execution error with %v\n", err)
	}
}

func Test_CreateExecution2(t *testing.T) {
	req := &model.Execution{
		Description: "创建执行测试2",
		Template: &model.Template{
			Name: "test_template_01",
			Operators: []*model.Operator{
				{
					Name:          "stop_bcc",
					Description:   "停止BCC实例",
					Operator:      "BCE::BCC::StopInstance",
					Retries:       3,
					RetryInterval: 60000,
					Timeout:       3600000,
					ParallelismControl: &model.RateControl{
						Ratio: 0,
						Count: 1,
					},
					Properties: map[string]interface{}{
						"instances": []map[string]string{
							{
								"instanceId": "i-jOS24U0l",
								"name":       "instance-3nd0wnlr",
							},
						},
					},
				},
				{
					Name:        "start_bcc",
					Description: "启动BCC实例",
					Operator:    "BCE::BCC::StartInstance",
					Properties: map[string]interface{}{
						"instances": []map[string]string{
							{
								"instanceId": "i-jOS24U0l",
								"name":       "instance-3nd0wnlr",
							},
						},
					},
				},
			},
			Linear: true,
		},
	}
	_, err := oosClient.CreateExecution(req)
	if err != nil {
		t.Errorf("oos create execution error with %v\n", err)
	}
}

func Test_GetExecutionDetail(t *testing.T) {
	result, err := oosClient.GetExecutionDetail("d-dSPRITR4tx2r")
	if err != nil || result.Code != 200 {
		t.Errorf("oos get execution detail error with %v\n", err)
	}
}
