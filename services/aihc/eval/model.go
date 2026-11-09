package eval

// Resource 资源：name 取 cpu、memory、sharedMemory 或卡型（如 baidu.com/a800_80g_cgpu），memory、sharedMemory 单位 GB
type Resource struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}

// DataSource 数据源挂载，首版创建只支持 type=bos；平台挂载由后端自动补，不用传
type DataSource struct {
	Type       string             `json:"type"`
	Name       string             `json:"name,omitempty"`
	SourcePath string             `json:"sourcePath,omitempty"`
	MountPath  string             `json:"mountPath,omitempty"`
	Id         string             `json:"id,omitempty"`
	Options    *DataSourceOptions `json:"options,omitempty"`
}

// DataSourceOptions 数据源挂载选项，字段与训练任务的数据源选项一一对应
type DataSourceOptions struct {
	ReadOnly          bool   `json:"readOnly,omitempty"`
	CfsInstanceId     string `json:"cfsInstanceId,omitempty"`
	CfsMountPoint     string `json:"cfsMountPoint,omitempty"`
	NfsServer         string `json:"nfsServer,omitempty"`
	DatasetVersion    string `json:"datasetVersion,omitempty"`
	RapidfsSubPath    string `json:"rapidfsSubPath,omitempty"`
	PfsL3InstanceId   string `json:"pfsL3InstanceId,omitempty"`
	PfsL3MountPoint   string `json:"pfsL3MountPoint,omitempty"`
	PfsL3MountOptions string `json:"pfsL3MountOptions,omitempty"`
}

// ModelConfig 模型服务配置，创建请求、详情回显、模型元数据的 defaultModelConfig 共用
type ModelConfig struct {
	// ModelType public / customCKPT / customImage
	ModelType string `json:"modelType"`
	// ModelName public、customCKPT 必填；customImage 必须为空
	ModelName    string       `json:"modelName,omitempty"`
	CkptPath     string       `json:"ckptPath,omitempty"`
	Replicas     int32        `json:"replicas"`
	MaxBatchSize int32        `json:"maxBatchSize"`
	Resources    []Resource   `json:"resources"`
	CustomImage  *CustomImage `json:"customImage,omitempty"`
}

// CustomImage 模型服务自定义镜像，modelType 为 customImage 时必填
type CustomImage struct {
	Image   string   `json:"image"`
	Command string   `json:"command"`
	Envs    []EnvVar `json:"envs,omitempty"`
	// ImageConfig 私有仓库凭据；详情里 password 恒为 ******，原样拿去创建会被拒
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
}

// EnvVar 环境变量，详情原样返回，不要放密钥
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ImageConfig 私有镜像仓库凭据
type ImageConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// String 不输出密码，避免凭据经 %v / %+v 打进日志
func (i ImageConfig) String() string {
	return "{Username:" + i.Username + " Password:******}"
}

// BenchmarkConfig 评测集配置，创建请求、详情回显、评测集元数据的 defaultBenchmarkConfig 共用
type BenchmarkConfig struct {
	Name            string     `json:"name"`
	SubBench        string     `json:"subBench"`
	Resources       []Resource `json:"resources"`
	ParallelShards  int32      `json:"parallelShards"`
	EpisodesPerTask int32      `json:"episodesPerTask"`
	// EvalType sync / realtime，不传按 sync
	EvalType string `json:"evalType,omitempty"`
	// Params 覆盖评测集默认参数，是 JSON 对象；存量数据不是合法 JSON 时详情里会是字符串
	Params        interface{}    `json:"params,omitempty"`
	RecordVideo   bool           `json:"recordVideo"`
	ModelOverride *ModelOverride `json:"modelOverride,omitempty"`
}

// ModelOverride 评测集级模型参数，chunkSize 必填且大于 0
type ModelOverride struct {
	CkptPath  string `json:"ckptPath,omitempty"`
	ChunkSize int32  `json:"chunkSize"`
	// ExtraArgs 是 JSON 对象，可填的键见模型元数据的 extraArgsSchema
	ExtraArgs interface{} `json:"extraArgs,omitempty"`
}

// CreateEmbodiedEvalArgs 创建评测任务的参数
type CreateEmbodiedEvalArgs struct {
	// ClientToken 幂等键，放在 query；不传时 SDK 为本次调用生成一个。调用方自己重试同一次创建时传回同一个值，不同任务用不同的值
	ClientToken      string `json:"-"`
	Name             string `json:"name"`
	ResourcePoolId   string `json:"resourcePoolId"`
	QueueId          string `json:"queueId"`
	ResourcePoolType string `json:"resourcePoolType,omitempty"`
	Priority         string `json:"priority,omitempty"`
	// VisibleScope 0 仅创建人可见，1 队列内可见；nil 按 1
	VisibleScope     *int32            `json:"visibleScope,omitempty"`
	DataSources      []DataSource      `json:"dataSources,omitempty"`
	ModelConfig      *ModelConfig      `json:"modelConfig"`
	BenchmarkConfigs []BenchmarkConfig `json:"benchmarkConfigs"`
}

// CreateEmbodiedEvalResult 创建评测任务的返回
type CreateEmbodiedEvalResult struct {
	RequestId string `json:"requestId"`
	EvalId    string `json:"evalId"`
}

// DescribeEmbodiedEvalsArgs 查询评测任务列表的参数，零值字段不下发
type DescribeEmbodiedEvalsArgs struct {
	PageNumber int
	PageSize   int
	// KeywordType name / evalId / creator，传 Keyword 时必须传
	KeywordType string
	Keyword     string
	// Status、ModelName、Priority 都可以逗号分隔多选
	Status    string
	ModelName string
	Priority  string
	// OrderBy createdAt / finishedAt / runTimeSeconds / priority
	OrderBy string
	// Order ASC / DESC，只收大写
	Order string
}

// EmbodiedEvalBenchmark 任务实际跑的评测集与子集
type EmbodiedEvalBenchmark struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	SubBench    string `json:"subBench"`
}

// EmbodiedEvalSummary 评测任务列表的每项
type EmbodiedEvalSummary struct {
	EvalId           string                  `json:"evalId"`
	Name             string                  `json:"name"`
	Status           string                  `json:"status"`
	ResourcePoolId   string                  `json:"resourcePoolId"`
	QueueId          string                  `json:"queueId"`
	ResourcePoolType string                  `json:"resourcePoolType"`
	Priority         string                  `json:"priority"`
	VisibleScope     int32                   `json:"visibleScope"`
	Creator          string                  `json:"creator"`
	ModelName        string                  `json:"modelName"`
	ModelDisplayName string                  `json:"modelDisplayName"`
	Benchmarks       []EmbodiedEvalBenchmark `json:"benchmarks"`
	Resources        []Resource              `json:"resources"`
	RunTimeSeconds   int64                   `json:"runTimeSeconds"`
	CreatedAt        string                  `json:"createdAt"`
	FinishedAt       string                  `json:"finishedAt"`
	Region           string                  `json:"region"`
}

// DescribeEmbodiedEvalsResult 查询评测任务列表的返回
type DescribeEmbodiedEvalsResult struct {
	RequestId  string                `json:"requestId"`
	TotalCount int64                 `json:"totalCount"`
	Evals      []EmbodiedEvalSummary `json:"evals"`
}

// DescribeEmbodiedEvalDetailResult 评测任务详情，字段与创建请求同构，另加只读字段。
// 不叫 DescribeEmbodiedEvalResult：那是另一个 Action（评测结果）的名字
type DescribeEmbodiedEvalDetailResult struct {
	RequestId        string            `json:"requestId"`
	EvalId           string            `json:"evalId"`
	Name             string            `json:"name"`
	JobId            string            `json:"jobId"`
	Status           string            `json:"status"`
	ResourcePoolId   string            `json:"resourcePoolId"`
	QueueId          string            `json:"queueId"`
	ResourcePoolType string            `json:"resourcePoolType"`
	Priority         string            `json:"priority"`
	VisibleScope     int32             `json:"visibleScope"`
	Creator          string            `json:"creator"`
	Region           string            `json:"region"`
	RunTimeSeconds   int64             `json:"runTimeSeconds"`
	CreatedAt        string            `json:"createdAt"`
	FinishedAt       string            `json:"finishedAt"`
	UpdatedAt        string            `json:"updatedAt"`
	DataSources      []DataSource      `json:"dataSources"`
	ModelConfig      *ModelConfig      `json:"modelConfig"`
	BenchmarkConfigs []BenchmarkConfig `json:"benchmarkConfigs"`
}

// StopEmbodiedEvalArgs 停止评测任务的参数
type StopEmbodiedEvalArgs struct {
	EvalId      string
	ClientToken string
}

// DeleteEmbodiedEvalArgs 删除评测任务的参数
type DeleteEmbodiedEvalArgs struct {
	EvalId      string
	ClientToken string
}

// EmbodiedEvalIdResult 停止、删除评测任务的返回
type EmbodiedEvalIdResult struct {
	RequestId string `json:"requestId"`
	EvalId    string `json:"evalId"`
}

// DescribeEmbodiedEvalResultResult 评测结果；任务没成功、结果还没产出时 Result 为 nil，只承诺顶层结构
type DescribeEmbodiedEvalResultResult struct {
	RequestId string      `json:"requestId"`
	EvalId    string      `json:"evalId"`
	Result    interface{} `json:"result"`
}

// DescribeEmbodiedEvalVideoArgs 查询回放视频地址的参数，两个字段都必填
type DescribeEmbodiedEvalVideoArgs struct {
	EvalId  string
	VideoId string
}

// DescribeEmbodiedEvalVideoResult 回放视频地址，StreamUrl 是预签名地址，ExpiresIn 秒后失效
type DescribeEmbodiedEvalVideoResult struct {
	RequestId string `json:"requestId"`
	VideoId   string `json:"videoId"`
	StreamUrl string `json:"streamUrl"`
	ExpiresIn int32  `json:"expiresIn"`
	MimeType  string `json:"mimeType"`
	SizeBytes int64  `json:"sizeBytes"`
}

// DescribeEmbodiedEvalModelsArgs 查询可评测模型的参数；Source 取 public / custom，不传按 public
type DescribeEmbodiedEvalModelsArgs struct {
	Source string
}

// EmbodiedEvalBenchmarkItem 评测集及其子集
type EmbodiedEvalBenchmarkItem struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	SubBenches  []string `json:"subBenches"`
}

// EmbodiedEvalModelItem 可评测模型及其评测集
type EmbodiedEvalModelItem struct {
	ModelName   string                      `json:"modelName"`
	DisplayName string                      `json:"displayName"`
	Benchmarks  []EmbodiedEvalBenchmarkItem `json:"benchmarks"`
}

// DescribeEmbodiedEvalModelsResult 查询可评测模型的返回
type DescribeEmbodiedEvalModelsResult struct {
	RequestId string                  `json:"requestId"`
	Models    []EmbodiedEvalModelItem `json:"models"`
}

// MinHardware 最低硬件配置，memory 单位 GB
type MinHardware struct {
	AcceleratorTypes []string `json:"acceleratorTypes"`
	AcceleratorCount int32    `json:"acceleratorCount"`
	Cpu              int32    `json:"cpu"`
	Memory           int32    `json:"memory"`
}

// DescribeEmbodiedEvalModelArgs 查询模型元数据的参数，ModelName 必填
type DescribeEmbodiedEvalModelArgs struct {
	ModelName string
}

// DescribeEmbodiedEvalModelResult 模型元数据；chunkSizeLimit、maxBatchSizeLimit 只是建议上限，后端不强校验
type DescribeEmbodiedEvalModelResult struct {
	RequestId          string       `json:"requestId"`
	ModelName          string       `json:"modelName"`
	DisplayName        string       `json:"displayName"`
	MinHardware        *MinHardware `json:"minHardware"`
	DefaultModelConfig *ModelConfig `json:"defaultModelConfig"`
	ExtraArgsSchema    interface{}  `json:"extraArgsSchema"`
	ChunkSizeLimit     int32        `json:"chunkSizeLimit"`
	MaxBatchSizeLimit  int32        `json:"maxBatchSizeLimit"`
}

// DescribeEmbodiedEvalBenchmarksArgs 查询评测集元数据的参数，ModelName 必填，其余可选
type DescribeEmbodiedEvalBenchmarksArgs struct {
	ModelName     string
	BenchmarkName string
	SubBench      string
	// Source public / custom，不传按 public
	Source string
}

// EmbodiedEvalBenchmarkMetadata 一个 (模型, 评测集, 子集) 的默认配置；MaxEpisodesPerTask 为 0 表示不限
type EmbodiedEvalBenchmarkMetadata struct {
	ModelName              string           `json:"modelName"`
	Name                   string           `json:"name"`
	SubBench               string           `json:"subBench"`
	MinHardware            *MinHardware     `json:"minHardware"`
	DefaultBenchmarkConfig *BenchmarkConfig `json:"defaultBenchmarkConfig"`
	ParamsSchema           interface{}      `json:"paramsSchema"`
	MaxEpisodesPerTask     int32            `json:"maxEpisodesPerTask"`
}

// DescribeEmbodiedEvalBenchmarksResult 查询评测集元数据的返回
type DescribeEmbodiedEvalBenchmarksResult struct {
	RequestId  string                          `json:"requestId"`
	Benchmarks []EmbodiedEvalBenchmarkMetadata `json:"benchmarks"`
}
