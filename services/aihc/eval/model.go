package eval

import (
	"fmt"
	"net/url"
)

// Resource 资源：name 取 cpu、memory、sharedMemory 或卡型（如 baidu.com/a800_80g_cgpu），memory、sharedMemory 单位 GB
type Resource struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}

// DataSource 数据源挂载，创建只收 type=bos、dataset（dataset 时 id 必填，版本放 options.datasetVersion）；平台挂载由后端自动补，不用传
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
	// ModelType public / customCKPT / customImage / selfHosted
	ModelType string `json:"modelType"`
	// ModelName public、customCKPT 必填；customImage 必须为空；selfHosted 为空或 selfHosted
	ModelName string `json:"modelName,omitempty"`
	CkptPath  string `json:"ckptPath,omitempty"`
	// Replicas、MaxBatchSize、Resources selfHosted 不能传（保持零值）
	Replicas     int32        `json:"replicas"`
	MaxBatchSize int32        `json:"maxBatchSize"`
	Resources    []Resource   `json:"resources"`
	CustomImage  *CustomImage `json:"customImage,omitempty"`
	// SelfHosted modelType 为 selfHosted 时必填，其他类型不能传
	SelfHosted *SelfHosted `json:"selfHosted,omitempty"`
}

// SelfHosted 自建推理服务：平台不拉起模型服务，评测节点直连 endpoints 里的地址
type SelfHosted struct {
	// Endpoints 本期只支持 1 个
	Endpoints []SelfHostedEndpoint `json:"endpoints"`
}

// SelfHostedEndpoint 自建推理服务地址
type SelfHostedEndpoint struct {
	// Url ws:// 或 wss://；详情里 query 整段换成 ?***，拿详情重建时要重新传完整 url
	Url string `json:"url"`
	// ServerTimeoutSec 单帧推理超时 1～900 秒，0 按 60；详情里是生效值
	ServerTimeoutSec int32 `json:"serverTimeoutSec,omitempty"`
	// AttachTimeoutSec 等待服务就绪的总时长 30～900 秒，0 按 300；详情里是生效值
	AttachTimeoutSec int32 `json:"attachTimeoutSec,omitempty"`
	// BearerToken 建连时放进 Authorization: Bearer <token>，不要带 Bearer 前缀；只写不读，详情里不返回
	BearerToken string `json:"bearerToken,omitempty"`
	// HasBearerToken 只在详情里返回：创建时带了 bearerToken。为 true 时拿详情重建要重新传 BearerToken
	HasBearerToken bool `json:"hasBearerToken,omitempty"`
}

// String 不输出 bearerToken 和 url query，避免凭据经 %v / %+v 打进日志
func (e SelfHostedEndpoint) String() string {
	token := ""
	if e.BearerToken != "" {
		token = "******"
	}
	return fmt.Sprintf("{Url:%s ServerTimeoutSec:%d AttachTimeoutSec:%d BearerToken:%s HasBearerToken:%t}",
		maskURLQuery(e.Url), e.ServerTimeoutSec, e.AttachTimeoutSec, token, e.HasBearerToken)
}

// maskURLQuery 把 query 整段换成 ?***、去掉 userinfo 和 fragment；解析失败不回显原串
func maskURLQuery(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return "<invalid url>"
	}
	hasQuery := u.RawQuery != "" || u.ForceQuery
	// 流水线是 Go 1.13，没有 RawFragment；只清 Fragment 时 String() 不再输出 #fragment
	u.User, u.RawQuery, u.ForceQuery, u.Fragment = nil, "", false, ""
	s := u.String()
	if hasQuery {
		s += "?***"
	}
	return s
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
	// OutputConf 评测输出位置；nil 不下发，写默认结果桶
	OutputConf *OutputConf `json:"outputConf,omitempty"`
}

// OutputConf 评测输出位置，创建请求与 OutputInfo.OutputConf 共用。SDK 不校验，规则以服务端为准
type OutputConf struct {
	// Mode 空或 default：默认结果桶；custom：写到本次评测 DataSources 里的一个用户挂载
	Mode string `json:"mode"`
	// MountPath custom 必填，等于 DataSources 里某个挂载的 mountPath（数据集没填 mountPath 时按版本默认挂载路径比对），
	// 不能是平台挂载点。只能命中非只读的 bos，或底层为 BOS、文件夹格式、当前用户可写的 dataset；mode 为空或 default 时不能传。
	// 输出根 <mountPath>[/<subdir>] 不能等于或落在其他挂载（含平台挂载、/mnt/aihc-eval-result）之下；
	// 结果文件的 <桶>/<前缀>/<evalId>/result/... 不能超过 1024 字节
	MountPath string `json:"mountPath,omitempty"`
	// Subdir 选填，仅 custom：挂载内的相对子目录，只含字母、数字和 . _ - /，不以 / 开头或结尾，path.Clean 后不变，
	// 不含 ..，不能是 .，不超过 128 字节；mode 为空或 default 时不能传
	Subdir string `json:"subdir,omitempty"`
}

// OutputInfo 服务端解析出的输出位置，详情和结果接口回显。
// 服务端取不出输出配置记录（记录损坏、mode 未知或自定义记录不完整）时不返回，此时为 nil，其余字段照常返回。
// 老任务（输出位置功能上线前创建，没有输出配置记录）按默认位置回显，结果文件直接在输出根下：
// ResultDir 是 /mnt/aihc-eval-result，StorageLocation 是 bos://<默认桶>/results/，都不含 evalId
type OutputInfo struct {
	// OutputConf 创建时的 outputConf，可直接填回 CreateEmbodiedEvalArgs.OutputConf。
	// 没传、mode 为空和老任务都回显 Mode=default；MountPath、Subdir 为空时省略
	OutputConf *OutputConf `json:"outputConf,omitempty"`
	// StorageType default（默认结果桶，含老任务）/ bos / dataset
	StorageType string `json:"storageType"`
	// ResultDir 结果文件目录的 Pod 内路径，新任务是 <输出根>/<evalId>/result
	ResultDir string `json:"resultDir"`
	// VideoDir 视频目录的 Pod 内路径，<输出根>/<evalId>/videos
	VideoDir string `json:"videoDir"`
	// StorageLocation 输出在 BOS 上的位置，新任务是 bos://<桶>/<前缀>/<evalId>/，前缀为空时 bos://<桶>/<evalId>/
	StorageLocation string `json:"storageLocation"`
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
	// OutputInfo 输出位置，取不出时为 nil。输出配置不在顶层回显，复制任务时把 OutputInfo.OutputConf 填到创建请求
	OutputInfo *OutputInfo `json:"outputInfo,omitempty"`
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
	// OutputInfo 输出位置，语义同 DescribeEmbodiedEvalDetailResult.OutputInfo
	OutputInfo *OutputInfo `json:"outputInfo,omitempty"`
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
