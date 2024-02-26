package as

import (
	"strconv"
	"time"

	"github.com/baidubce/bce-sdk-go/model"
)

type ListAsGroupRequest struct {
	GroupId     string `json:"groupId,omitempty"`
	OrderBy     string `json:"orderBy,omitempty"`
	Order       string `json:"order,omitempty"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	KeyWord     string `json:"keyWord,omitempty"`
	KeyWordType string `json:"keyWordType,omitempty"`
}

type ListAsGroupResponse struct {
	OrderBy    string      `json:"orderBy,omitempty"`
	Order      string      `json:"order,omitempty"`
	PageNo     int         `json:"pageNo,omitempty"`
	PageSize   int         `json:"pageSize,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	Result     []GroupInfo `json:"result,omitempty"`
}

type AsGroupStatus string

const (
	CREATING       AsGroupStatus = "CREATING"
	RUNNING        AsGroupStatus = "RUNNING"
	SCALING_UP     AsGroupStatus = "SCALING_UP"
	SCALING_DOWN   AsGroupStatus = "SCALING_DOWN"
	ATTACHING_NODE AsGroupStatus = "ATTACHING_NODE"
	DETACHING_NODE AsGroupStatus = "DETACHING_NODE"
	DELETING       AsGroupStatus = "DELETING"
	BINDING_BLB    AsGroupStatus = "BINDING_BLB"
	UNBINDING_BLB  AsGroupStatus = "UNBINDING_BLB"
	COOLDOWN       AsGroupStatus = "COOLDOWN"
	PAUSE          AsGroupStatus = "PAUSE"
	DELETED        AsGroupStatus = "DELETED"
)

type ZoneInfo struct {
	Zone       string `json:"zone,omitempty"`
	SubnetID   string `json:"subnetId,omitempty"`
	SubnetUUID string `json:"subnetUuid,omitempty"`
	SubnetName string `json:"subnetName,omitempty"`
	SubnetType int16  `json:"subnetType,omitempty"`
	NodeCount  int    `json:"nodeCount,omitempty"`
}

type GroupConfig struct {
	MinNodeNum    int `json:"minNodeNum,omitempty"`
	MaxNodeNum    int `json:"maxNodeNum,omitempty"`
	CooldownInSec int `json:"cooldownInSec,omitempty"`
	ExpectNum     int `json:"expectNum,omitempty"`
}

type GroupInfo struct {
	GroupId    string        `json:"groupId,omitempty"`
	GroupName  string        `json:"groupName,omitempty"`
	Region     string        `json:"region,omitempty"`
	Status     AsGroupStatus `json:"status,omitempty"`
	VpcId      string        `json:"vpcId,omitempty"`
	NodeNum    int           `json:"nodeNum,omitempty"`
	CreateTime string        `json:"createTime,omitempty"`
	ZoneInfo   []ZoneInfo    `json:"zoneInfo,omitempty"`
	Config     GroupConfig   `json:"config,omitempty"`
	BlbId      string        `json:"blbId,omitempty"`
}

type GetAsGroupRequest struct {
	GroupId string `json:"groupId,omitempty"`
}

type VpcInfo struct {
	VpcId   string `json:"vpcId,omitempty"`
	VpcName string `json:"vpcName,omitempty"`
	VpcUUID string `json:"vpcUuid,omitempty"`
}

type GetAsGroupResponse struct {
	GroupID           string      `json:"groupId,omitempty"`
	GroupName         string      `json:"groupName,omitempty"`
	Region            string      `json:"region,omitempty"`
	Status            string      `json:"status,omitempty"`
	VpcInfo           VpcInfo     `json:"vpcInfo,omitempty"`
	ZoneInfo          []ZoneInfo  `json:"zoneInfo,omitempty"`
	Config            GroupConfig `json:"config,omitempty"`
	BlbID             string      `json:"blbId,omitempty"`
	NodeNum           int         `json:"nodeNum,omitempty"`
	CreateTime        string      `json:"createTime,omitempty"`
	RdsIDs            string      `json:"rdsIds,omitempty"`
	ScsIDs            string      `json:"scsIds,omitempty"`
	ExpansionStrategy string      `json:"expansionStrategy,omitempty"`
	ShrinkageStrategy string      `json:"shrinkageStrategy,omitempty"`
}

type IncreaseAsGroupRequest struct {
	ClientToken       string   `json:"-"`
	GroupId           string   `json:"groupId,omitempty"`
	NodeCount         int      `json:"nodeCount,omitempty"`
	Zone              []string `json:"zone,omitempty"`
	ExpansionStrategy string   `json:"expansionStrategy,omitempty"`
}

type DecreaseAsGroupRequest struct {
	ClientToken string   `json:"-"`
	GroupId     string   `json:"groupId,omitempty"`
	Nodes       []string `json:"nodes,omitempty"`
}

type AdjustAsGroupRequest struct {
	ClientToken string `json:"-"`
	GroupId     string `json:"groupId,omitempty"`
	AdjustNum   int    `json:"adjustNum,omitempty"`
}

type ListAsNodeRequest struct {
	GroupId string `json:"groupId,omitempty"`
	Marker  string `json:"marker,omitempty"`
	MaxKeys int    `json:"maxKeys,omitempty"`
}

type AsEipModel struct {
	BandwidthInMbps int    `json:"bandwidthInMbps,omitempty"`
	EipId           string `json:"eipId,omitempty"`
	Address         string `json:"address,omitempty"`
	EipStatus       string `json:"eipStatus,omitempty"`
	EipAllocationId string `json:"eipAllocationId,omitempty"`
}

type NodeModel struct {
	InstanceId         string           `json:"instanceId,omitempty"`
	InstanceUuid       string           `json:"instanceUuid,omitempty"`
	InstanceName       string           `json:"instanceName,omitempty"`
	FloatingIp         string           `json:"floatingIp,omitempty"`
	InternalIp         string           `json:"internalIp,omitempty"`
	Status             string           `json:"status,omitempty"`
	Payment            string           `json:"payment,omitempty"`
	CpuCount           int64            `json:"cpuCount,omitempty"`
	MemoryCapacityInGB int64            `json:"memoryCapacityInGB,omitempty"`
	InstanceType       string           `json:"instanceType,omitempty"`
	SysDiskInGB        int              `json:"sysDiskInGB,omitempty"`
	CreateTime         string           `json:"createTime,omitempty"`
	Eip                AsEipModel       `json:"eip,omitempty"`
	SubnetType         string           `json:"subnetType,omitempty"`
	IsProtected        bool             `json:"isProtected,omitempty"`
	NodeType           string           `json:"nodeType,omitempty"`
	Tags               []model.TagModel `json:"tags,omitempty"`
	GroupId            string           `json:"groupId,omitempty"`
}

type ListAsNodeResponse struct {
	OrderBy    string      `json:"orderBy,omitempty"`
	Order      string      `json:"order,omitempty"`
	PageNo     int         `json:"pageNo,omitempty"`
	PageSize   int         `json:"pageSize,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	Result     []NodeModel `json:"result,omitempty"`
}

type CreateAsGroupRequest struct {
	GroupName         string        `json:"groupName"`
	Config            Config        `json:"config"`
	HealthCheck       HealthCheck   `json:"healthCheck"`
	KeypairId         string        `json:"keypairId,omitempty"`
	KeypairName       string        `json:"keypairName,omitempty"`
	Blb               []BlbInfo     `json:"blb"`
	Rds               []string      `json:"rds,omitempty"`
	Scs               []string      `json:"scs,omitempty"`
	ExpansionStrategy string        `json:"expansionStrategy,omitempty"`
	ShrinkageStrategy string        `json:"shrinkageStrategy,omitempty"`
	ZoneInfo          []ZoneInfo    `json:"zoneInfo"`
	AssignTagInfo     AssignTagInfo `json:"assignTagInfo"`
	Nodes             []NodeInfo    `json:"nodes"`
	Eip               Eip           `json:"eip"`
	Billing           Billing       `json:"billing"`
	CmdConfig         CmdConfig     `json:"cmdConfig"`
	BccNameConfig     BccNameConfig `json:"bccNameConfig"`
}

type Config struct {
	MinNodeNum    int `json:"minNodeNum"`
	ExpectNum     int `json:"expectNum"`
	MaxNodeNum    int `json:"maxNodeNum"`
	CooldownInSec int `json:"cooldownInSec"`
}

type HealthCheck struct {
	HealthCheckInterval int `json:"healthCheckInterval"`
	GraceTime           int `json:"graceTime"`
}

type BlbInfo struct {
	BlbId   string   `json:"blbId,omitempty"`
	BlbName string   `json:"blbName,omitempty"`
	BlbType string   `json:"blbType,omitempty"`
	SgIds   []string `json:"sgIds,omitempty"`
}

type AssignTagInfo struct {
	RelationTag bool  `json:"relationTag"`
	Tags        []Tag `json:"tags,omitempty"`
}

type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

type NodeInfo struct {
	CpuCount           int             `json:"cpuCount"`
	MemoryCapacityInGB int             `json:"memoryCapacityInGB"`
	SysDiskType        string          `json:"sysDiskType"`
	SysDiskInGB        int             `json:"sysDiskInGB"`
	InstanceType       int             `json:"instanceType"`
	ImageId            string          `json:"imageId"`
	ImageType          string          `json:"imageType"`
	SecurityGroupId    string          `json:"securityGroupId"`
	Spec               string          `json:"spec"`
	EphemeralDisks     []EphemeralDisk `json:"ephemeralDisks,omitempty"`
	AspId              string          `json:"aspId,omitempty"`
	UserData           string          `json:"userData,omitempty"`
	Priorities         int             `json:"priorities"`
	ZoneSubnet         string          `json:"zoneSubnet,omitempty"`
	TotalCount         int             `json:"totalCount"`
	Cds                []CdsInfo       `json:"cds,omitempty"`
	BidModel           string          `json:"bidModel"`
	BidPrice           float64         `json:"bidPrice,omitempty"`
	ProductType        string          `json:"productType"`
	OsType             string          `json:"osType"`
}

type CdsInfo struct {
	VolumeType   string `json:"volumeType,omitempty"`
	SizeInGB     int    `json:"sizeInGB,omitempty"`
	SnapshotId   string `json:"snapshotId,omitempty"`
	SnapshotName string `json:"snapshotName,omitempty"`
}

type EphemeralDisk struct {
	StorageType string `json:"storageType,omitempty"`
	SizeInGB    int    `json:"sizeInGB,omitempty"`
}

type Eip struct {
	IfBindEip       bool   `json:"ifBindEip"`
	BandwidthInMbps int    `json:"bandwidthInMbps"`
	EipProductType  string `json:"eipProductType"`
	PurchaseType    string `json:"purchaseType"`
}

type Billing struct {
	PaymentTiming string `json:"paymentTiming"`
}

type CmdConfig struct {
	HasDecreaseCmd bool   `json:"hasDecreaseCmd,omitempty"`
	DecCmdStrategy string `json:"decCmdStrategy,omitempty"`
	DecCmdData     string `json:"decCmdData,omitempty"`
	DecCmdTimeout  int    `json:"decCmdTimeout,omitempty"`
	DecCmdManual   bool   `json:"decCmdManual,omitempty"`
	HasIncreaseCmd bool   `json:"hasIncreaseCmd,omitempty"`
	IncCmdStrategy string `json:"incCmdStrategy,omitempty"`
	IncCmdData     string `json:"incCmdData,omitempty"`
	IncCmdTimeout  int    `json:"incCmdTimeout,omitempty"`
	IncCmdManual   bool   `json:"incCmdManual,omitempty"`
}

type BccNameConfig struct {
	BccName            string `json:"bccName,omitempty"`
	BccHostname        string `json:"bccHostname,omitempty"`
	AutoSeqSuffix      bool   `json:"autoSeqSuffix,omitempty"`
	OpenHostnameDomain bool   `json:"openHostnameDomain,omitempty"`
}

type CreateAsGroupResponse struct {
	GroupId string `json:"groupId,omitempty"`
}

type DeleteAsGroupRequest struct {
	GroupIds []string `json:"groupIds,omitempty"`
}

const (
	ListRecordsUrl = "/v1/record"
	BaseURL        = "/v1/group/%s"
	BaseV2URL      = "/v2/group/%s"
	RuleURL        = "/v1/rule"
	RuleIdURL      = "/v1/rule/%s"
)

type ListRecordsRequest struct {
	GroupID   string    `json:"groupid"`
	PageNo    int       `json:"pageNo"`
	OrderBy   string    `json:"orderBy"`
	Order     string    `json:"order"`
	PageSize  int       `json:"pageSize"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func (request *ListRecordsRequest) GetBceQueryParams() map[string]string {
	result := make(map[string]string)
	if request.GroupID != "" {
		result["groupid"] = request.GroupID
	}
	if request.PageNo != 0 {
		result["pageNo"] = strconv.Itoa(request.PageNo)
	}
	if request.OrderBy != "" {
		result["orderBy"] = request.OrderBy
	}
	if request.Order != "" {
		result["order"] = request.Order
	}
	if request.PageSize != 0 {
		result["pageSize"] = strconv.Itoa(request.PageSize)
	}
	if !request.StartTime.IsZero() {
		result["startTime"] = request.StartTime.Format(time.RFC3339)
	}
	if !request.EndTime.IsZero() {
		result["endTime"] = request.EndTime.Format(time.RFC3339)
	}

	return result
}

type AsRecord struct {
	GroupID         string         `json:"groupId"`
	RecordID        string         `json:"recordId"`
	StartTime       time.Time      `json:"startTime"`
	Result          string         `json:"result"`
	ActualScaleNode []string       `json:"actualScaleNode"`
	RemainedNode    []string       `json:"remainedNode"`
	Action          string         `json:"action"`
	ScaleCondition  ScaleCondition `json:"scaleCondition"`
	RuleID          string         `json:"ruleId"`
	Message         string         `json:"message"`
	ExpectAction    ExpectAction   `json:"expectAction"`
	ExecuteType     string         `json:"executeType"`
	DagID           string         `json:"dagId"`
}

type ScaleCondition struct {
	TargetType         string `json:"targetType"`
	TargetID           string `json:"targetId"`
	Indicator          string `json:"indicator"`
	Threshold          string `json:"threshold"`
	Unit               string `json:"unit"`
	ComparisonOperator string `json:"comparisonOperator"`
	CronTime           string `json:"cronTime"`
	Type               string `json:"type"`
	PeriodType         string `json:"periodType"`
	PeriodValue        int    `json:"periodValue"`
}

type ExpectAction struct {
	ActionType  string `json:"actionType"`
	ActionNum   int    `json:"actionNum"`
	AdjustToNum int    `json:"adjustToNum"`
}

type ListRecordsResponse struct {
	OrderBy    string     `json:"orderBy"`
	Order      string     `json:"order"`
	PageNo     int        `json:"pageNo"`
	PageSize   int        `json:"pageSize"`
	TotalCount int        `json:"totalCount"`
	Result     []AsRecord `json:"result"`
}
type CreateDagResponse struct {
	Success bool     `json:"success,omitempty"`
	Msg     string   `json:"msg,omitempty"`
	Result  DagModel `json:"result,omitempty"`
}

type DagModel struct {
	ID                string                         `json:"id,omitempty"`
	Description       string                         `json:"description,omitempty"`
	Revision          int64                          `json:"revision,omitempty"`
	CreatedTimestamp  int64                          `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  int64                          `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp int64                          `json:"finishedTimestamp,omitempty"`
	Namespace         string                         `json:"namespace,omitempty"`
	State             string                         `json:"state,omitempty"`
	DagSpec           DagSpecModel                   `json:"dagSpec,omitempty"`
	InitContext       map[string]interface{}         `json:"initContext,omitempty"`
	Parallelism       int                            `json:"parallelism,omitempty"`
	Manually          bool                           `json:"manually,omitempty"`
	WorkerSelectors   []TagSelector                  `json:"workerSelectors,omitempty"`
	Tasks             []TaskModel                    `json:"tasks,omitempty"`
	User              UserModel                      `json:"user,omitempty"`
	OperatorActions   map[string]OperatorActionModel `json:"operatorActions,omitempty"`
	DagActions        DagActionModel                 `json:"dagActions,omitempty"`
	Semaphore         SemaphoreModel                 `json:"semaphore,omitempty"`
}

type DagSpecModel struct {
	Ref         string                 `json:"ref,omitempty"`
	Namespace   string                 `json:"namespace,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
	Operators   []OperatorModel        `json:"operators,omitempty"`
	Linear      bool                   `json:"linear,omitempty"`
	Links       []LinkModel            `json:"links,omitempty"`
	Inputs      []InputModel           `json:"inputs,omitempty"`
	Outputs     []OutputModel          `json:"outputs,omitempty"`
	Parallelism int                    `json:"parallelism,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
}

type OperatorModel struct {
	Name                   string                 `json:"name,omitempty"`
	Description            string                 `json:"description,omitempty"`
	Tags                   map[string]string      `json:"tags,omitempty"`
	Operator               string                 `json:"operator,omitempty"`
	DagSpec                DagSpecModel           `json:"dagSpec,omitempty"`
	Inline                 bool                   `json:"inline,omitempty"`
	Retries                int                    `json:"retries,omitempty"`
	RetryInterval          int                    `json:"retryInterval,omitempty"`
	Timeout                int                    `json:"timeout,omitempty"`
	InitContext            map[string]interface{} `json:"initContext,omitempty"`
	Loops                  map[string]string      `json:"loops,omitempty"`
	ParallelismRatio       float64                `json:"parallelismRatio,omitempty"`
	ParallelismCount       int                    `json:"parallelismCount,omitempty"`
	AllowedFailureRatio    float64                `json:"allowedFailureRatio,omitempty"`
	AllowedFailureCount    int                    `json:"allowedFailureCount,omitempty"`
	Manually               bool                   `json:"manually,omitempty"`
	ScheduleDelayMilli     int                    `json:"scheduleDelayMilli,omitempty"`
	WorkerSelectors        []TagSelector          `json:"workerSelectors,omitempty"`
	CollectChildrenContext string                 `json:"collectChildrenContext,omitempty"`
	TriggerRule            string                 `json:"triggerRule,omitempty"`
	WaitOnAgentMilli       int                    `json:"waitOnAgentMilli,omitempty"`
	Condition              map[string]string      `json:"condition,omitempty"`
}

type TagSelector struct {
	Expressions []TagExpression `json:"expressions,omitempty"`
}

type TagExpression struct {
	Key    string   `json:"key,omitempty"`
	Op     string   `json:"op,omitempty"`
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

type LinkModel struct {
	Src string `json:"src,omitempty"`
	Dst string `json:"dst,omitempty"`
}

type InputModel struct {
	Name         string   `json:"name,omitempty"`
	Required     bool     `json:"required,omitempty"`
	Type         string   `json:"type,omitempty"`
	Description  string   `json:"description,omitempty"`
	Options      []string `json:"options,omitempty"`
	DefaultValue string   `json:"default,omitempty"`
}

type OutputModel struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type OperatorActionModel struct {
	RunManually bool `json:"runManually,omitempty"`
	Ignore      bool `json:"ignore,omitempty"`
}

type DagActionModel struct {
	Pause  bool `json:"pause,omitempty"`
	Resume bool `json:"resume,omitempty"`
}

type SemaphoreModel struct {
	Name     string `json:"name,omitempty"`
	Conflict string `json:"conflict,omitempty"`
}

type UserModel struct {
	ID string `json:"id,omitempty"`
}

type TaskModel struct {
	ID                string                 `json:"id,omitempty"`
	LoopIndex         int                    `json:"loopIndex,omitempty"`
	Namespace         string                 `json:"namespace,omitempty"`
	Dag               DagModel               `json:"dag,omitempty"`
	Dags              []DagModel             `json:"dags,omitempty"`
	Revision          int64                  `json:"revision,omitempty"`
	CreatedTimestamp  int64                  `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  int64                  `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp int64                  `json:"finishedTimestamp,omitempty"`
	State             string                 `json:"state,omitempty"`
	Operator          OperatorModel          `json:"operator,omitempty"`
	Reason            string                 `json:"reason,omitempty"`
	InitContext       map[string]interface{} `json:"initContext,omitempty"`
	Context           map[string]interface{} `json:"context,omitempty"`
	OutputContext     map[string]interface{} `json:"outputContext,omitempty"`
	Tries             int                    `json:"tries,omitempty"`
	Children          []TaskModel            `json:"children,omitempty"`
}

type ScalingUpRequest struct {
	NodeCount         int      `json:"nodeCount,omitempty"`
	Zone              []string `json:"zone,omitempty"`
	ExpansionStrategy string   `json:"expansionStrategy,omitempty"`
}
type ScalingDownRequest struct {
	Nodes []string `json:"nodes,omitempty"`
}

type NodeRequest struct {
	Nodes []string `json:"nodes,omitempty"`
}

type RuleRequest struct {
	RuleID             string `json:"ruleId,omitempty"`
	RuleName           string `json:"ruleName,omitempty"`
	GroupID            string `json:"groupId,omitempty"`
	State              string `json:"state,omitempty"`
	Type               string `json:"type,omitempty"`
	TargetType         string `json:"targetType,omitempty"`
	TargetID           string `json:"targetId,omitempty"`
	Indicator          string `json:"indicator,omitempty"`
	Threshold          string `json:"threshold,omitempty"`
	Unit               string `json:"unit,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	ActionType         string `json:"actionType,omitempty"`
	ActionNum          int    `json:"actionNum,omitempty"`
	CronTime           string `json:"cronTime,omitempty"`
	CooldownInSec      int    `json:"cooldownInSec,omitempty"`
	PeriodType         string `json:"periodType,omitempty"`
	PeriodValue        int    `json:"periodValue,omitempty"`
	PeriodStartTime    string `json:"periodStartTime,omitempty"`
	PeriodEndTime      string `json:"periodEndTime,omitempty"`
}

type CreateRuleResult struct {
	RuleID string `json:"ruleId,omitempty"`
}

type RuleListQuery struct {
	GroupID     string `json:"groupId,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	KeywordType string `json:"keywordType,omitempty"`
	Order       string `json:"order,omitempty"`
	OrderBy     string `json:"orderBy,omitempty"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type RuleVOListResponse struct {
	OrderBy    string   `json:"orderBy,omitempty"`
	Order      string   `json:"order,omitempty"`
	PageNo     int      `json:"pageNo,omitempty"`
	PageSize   int      `json:"pageSize,omitempty"`
	TotalCount int      `json:"totalCount,omitempty"`
	Result     []RuleVO `json:"result,omitempty"`
}

type RuleVO struct {
	RuleID             string `json:"ruleId,omitempty"`
	RuleName           string `json:"ruleName,omitempty"`
	GroupID            string `json:"groupId,omitempty"`
	AccountID          string `json:"accountId,omitempty"`
	State              string `json:"state,omitempty"`
	Type               string `json:"type,omitempty"`
	TargetType         string `json:"targetType,omitempty"`
	TargetID           string `json:"targetId,omitempty"`
	Indicator          string `json:"indicator,omitempty"`
	Threshold          string `json:"threshold,omitempty"`
	Unit               string `json:"unit,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	CronTime           string `json:"cronTime,omitempty"`
	ActionType         string `json:"actionType,omitempty"`
	ActionNum          int    `json:"actionNum,omitempty"`
	CooldownInSec      int    `json:"cooldownInSec,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
	LastExecutionTime  string `json:"lastExecutionTime,omitempty"`
	PeriodStartTime    string `json:"periodStartTime,omitempty"`
	PeriodEndTime      string `json:"periodEndTime,omitempty"`
	PeriodType         string `json:"periodType,omitempty"`
	PeriodValue        int    `json:"periodValue,omitempty"`
}

type RuleDelRequest struct {
	RuleIds  []string `json:"ruleIds,omitempty"`
	GroupIds []string `json:"groupIds,omitempty"`
}
