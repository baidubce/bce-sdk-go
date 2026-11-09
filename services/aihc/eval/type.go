package eval

// Interface 具身评测 OpenAPI 接口，一个方法对应一个 Action
type Interface interface {
	CreateEmbodiedEval(*CreateEmbodiedEvalArgs) (*CreateEmbodiedEvalResult, error)
	DescribeEmbodiedEvals(*DescribeEmbodiedEvalsArgs) (*DescribeEmbodiedEvalsResult, error)
	DescribeEmbodiedEval(evalId string) (*DescribeEmbodiedEvalDetailResult, error)
	StopEmbodiedEval(*StopEmbodiedEvalArgs) (*EmbodiedEvalIdResult, error)
	DeleteEmbodiedEval(*DeleteEmbodiedEvalArgs) (*EmbodiedEvalIdResult, error)
	DescribeEmbodiedEvalResult(evalId string) (*DescribeEmbodiedEvalResultResult, error)
	DescribeEmbodiedEvalVideo(*DescribeEmbodiedEvalVideoArgs) (*DescribeEmbodiedEvalVideoResult, error)
	DescribeEmbodiedEvalModels(*DescribeEmbodiedEvalModelsArgs) (*DescribeEmbodiedEvalModelsResult, error)
	DescribeEmbodiedEvalModel(*DescribeEmbodiedEvalModelArgs) (*DescribeEmbodiedEvalModelResult, error)
	DescribeEmbodiedEvalBenchmarks(*DescribeEmbodiedEvalBenchmarksArgs) (*DescribeEmbodiedEvalBenchmarksResult, error)
}
