package eval

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

const (
	DefaultVersion = "v2"
	DefaultUrl     = "/"

	ApiCreateEmbodiedEval             = "CreateEmbodiedEval"
	ApiDescribeEmbodiedEvals          = "DescribeEmbodiedEvals"
	ApiDescribeEmbodiedEval           = "DescribeEmbodiedEval"
	ApiStopEmbodiedEval               = "StopEmbodiedEval"
	ApiDeleteEmbodiedEval             = "DeleteEmbodiedEval"
	ApiDescribeEmbodiedEvalResult     = "DescribeEmbodiedEvalResult"
	ApiDescribeEmbodiedEvalVideo      = "DescribeEmbodiedEvalVideo"
	ApiDescribeEmbodiedEvalModels     = "DescribeEmbodiedEvalModels"
	ApiDescribeEmbodiedEvalModel      = "DescribeEmbodiedEvalModel"
	ApiDescribeEmbodiedEvalBenchmarks = "DescribeEmbodiedEvalBenchmarks"
)

// newRequest 具身评测 OpenAPI 的公共请求：POST|GET /?action=<Action>，header version: v2
func newRequest(method, action string) *bce.BceRequest {
	req := &bce.BceRequest{}
	req.SetUri(DefaultUrl)
	req.SetMethod(method)
	req.SetHeader("version", DefaultVersion)
	req.SetParam("action", action)
	return req
}

// setParamIfNotEmpty 零值参数不下发，交给服务端默认值
func setParamIfNotEmpty(req *bce.BceRequest, key, value string) {
	if value != "" {
		req.SetParam(key, value)
	}
}

// send 发请求并把成功体解析到 result；HTTP 状态 >= 400 时返回带 code/message/requestId 的 *bce.BceServiceError
func send(cli bce.Client, req *bce.BceRequest, result interface{}) error {
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return resp.ParseJsonBody(result)
}

// newClientToken 生成幂等键；单测替换此变量
var newClientToken = util.NewUUID

func requireArg(name, value string) error {
	if value == "" {
		return errors.New(name + " is required")
	}
	return nil
}

func CreateEmbodiedEval(cli bce.Client, args *CreateEmbodiedEvalArgs) (*CreateEmbodiedEvalResult, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req := newRequest(http.POST, ApiCreateEmbodiedEval)
	// clientToken 只认 query，body 里的不读。不传时为本次调用生成一个：SDK 自动重试复用同一个请求，
	// 不会因为首次响应超时而重复建任务；不回写 args，避免调用方复用 args 建下一个任务时撞上幂等冲突
	clientToken := args.ClientToken
	if clientToken == "" {
		clientToken = newClientToken()
	}
	req.SetParam("clientToken", clientToken)
	req.SetBody(body)
	result := &CreateEmbodiedEvalResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func DescribeEmbodiedEvals(cli bce.Client, args *DescribeEmbodiedEvalsArgs) (*DescribeEmbodiedEvalsResult, error) {
	if args == nil {
		args = &DescribeEmbodiedEvalsArgs{}
	}
	req := newRequest(http.GET, ApiDescribeEmbodiedEvals)
	if args.PageNumber != 0 {
		req.SetParam("pageNumber", strconv.Itoa(args.PageNumber))
	}
	if args.PageSize != 0 {
		req.SetParam("pageSize", strconv.Itoa(args.PageSize))
	}
	setParamIfNotEmpty(req, "keywordType", args.KeywordType)
	setParamIfNotEmpty(req, "keyword", args.Keyword)
	setParamIfNotEmpty(req, "status", args.Status)
	setParamIfNotEmpty(req, "modelName", args.ModelName)
	setParamIfNotEmpty(req, "priority", args.Priority)
	setParamIfNotEmpty(req, "orderBy", args.OrderBy)
	setParamIfNotEmpty(req, "order", args.Order)
	result := &DescribeEmbodiedEvalsResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func DescribeEmbodiedEval(cli bce.Client, evalId string) (*DescribeEmbodiedEvalDetailResult, error) {
	if err := requireArg("evalId", evalId); err != nil {
		return nil, err
	}
	req := newRequest(http.GET, ApiDescribeEmbodiedEval)
	req.SetParam("evalId", evalId)
	result := &DescribeEmbodiedEvalDetailResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

// operateEmbodiedEval Stop、Delete 共用：evalId 必填，clientToken 可选，body 不读
func operateEmbodiedEval(cli bce.Client, action, evalId, clientToken string) (*EmbodiedEvalIdResult, error) {
	if err := requireArg("evalId", evalId); err != nil {
		return nil, err
	}
	req := newRequest(http.POST, action)
	req.SetParam("evalId", evalId)
	setParamIfNotEmpty(req, "clientToken", clientToken)
	result := &EmbodiedEvalIdResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func StopEmbodiedEval(cli bce.Client, args *StopEmbodiedEvalArgs) (*EmbodiedEvalIdResult, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}
	return operateEmbodiedEval(cli, ApiStopEmbodiedEval, args.EvalId, args.ClientToken)
}

func DeleteEmbodiedEval(cli bce.Client, args *DeleteEmbodiedEvalArgs) (*EmbodiedEvalIdResult, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}
	return operateEmbodiedEval(cli, ApiDeleteEmbodiedEval, args.EvalId, args.ClientToken)
}

func DescribeEmbodiedEvalResult(cli bce.Client, evalId string) (*DescribeEmbodiedEvalResultResult, error) {
	if err := requireArg("evalId", evalId); err != nil {
		return nil, err
	}
	req := newRequest(http.GET, ApiDescribeEmbodiedEvalResult)
	req.SetParam("evalId", evalId)
	result := &DescribeEmbodiedEvalResultResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func DescribeEmbodiedEvalVideo(cli bce.Client, args *DescribeEmbodiedEvalVideoArgs) (*DescribeEmbodiedEvalVideoResult, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}
	if err := requireArg("evalId", args.EvalId); err != nil {
		return nil, err
	}
	if err := requireArg("videoId", args.VideoId); err != nil {
		return nil, err
	}
	req := newRequest(http.GET, ApiDescribeEmbodiedEvalVideo)
	req.SetParam("evalId", args.EvalId)
	req.SetParam("videoId", args.VideoId)
	result := &DescribeEmbodiedEvalVideoResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func DescribeEmbodiedEvalModels(cli bce.Client, args *DescribeEmbodiedEvalModelsArgs) (*DescribeEmbodiedEvalModelsResult, error) {
	if args == nil {
		args = &DescribeEmbodiedEvalModelsArgs{}
	}
	req := newRequest(http.GET, ApiDescribeEmbodiedEvalModels)
	setParamIfNotEmpty(req, "source", args.Source)
	result := &DescribeEmbodiedEvalModelsResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func DescribeEmbodiedEvalModel(cli bce.Client, args *DescribeEmbodiedEvalModelArgs) (*DescribeEmbodiedEvalModelResult, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}
	if err := requireArg("modelName", args.ModelName); err != nil {
		return nil, err
	}
	req := newRequest(http.GET, ApiDescribeEmbodiedEvalModel)
	req.SetParam("modelName", args.ModelName)
	result := &DescribeEmbodiedEvalModelResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func DescribeEmbodiedEvalBenchmarks(cli bce.Client, args *DescribeEmbodiedEvalBenchmarksArgs) (*DescribeEmbodiedEvalBenchmarksResult, error) {
	if args == nil {
		return nil, errors.New("args is required")
	}
	if err := requireArg("modelName", args.ModelName); err != nil {
		return nil, err
	}
	req := newRequest(http.GET, ApiDescribeEmbodiedEvalBenchmarks)
	req.SetParam("modelName", args.ModelName)
	setParamIfNotEmpty(req, "benchmarkName", args.BenchmarkName)
	setParamIfNotEmpty(req, "subBench", args.SubBench)
	setParamIfNotEmpty(req, "source", args.Source)
	result := &DescribeEmbodiedEvalBenchmarksResult{}
	if err := send(cli, req, result); err != nil {
		return nil, err
	}
	return result, nil
}
