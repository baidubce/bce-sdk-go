package bcm

import (
	"strings"
	"time"
)

type SiteOnceProtocol string

const (
	SiteOnceBaseURL     = "/csm/api/v1/site/once/"
	CreateSiteOnce      = SiteOnceBaseURL + "%s/taskCreate"
	ListSiteOnceTasks   = SiteOnceBaseURL + "/taskList"
	DeleteSiteOnceTasks = SiteOnceBaseURL + "/taskDelete"
	LoadSiteOnceData    = SiteOnceBaseURL + "/loadData"
	DetailSiteOnceTask  = SiteOnceBaseURL + "/groupTask"
	AgainSiteOnce       = SiteOnceBaseURL + "/createFromTask"
	ListSiteOnceHistory = SiteOnceBaseURL + "/groupTaskList"
	GetSiteAgent        = SiteOnceBaseURL + "/siteAgent"
)

type ResolveTypeEnum string

type SiteOnceConfig struct {
	Method         string          `json:"method,omitempty"`
	PostContent    string          `json:"postContent,omitempty"`
	ResolveType    ResolveTypeEnum `json:"resolveType,omitempty"`
	Server         string          `json:"server,omitempty"`
	KidnapWhite    string          `json:"kidnapWhite,omitempty"`
	PacketCount    int             `json:"packetCount,omitempty"`
	Port           int             `json:"port,omitempty"`
	InputType      int             `json:"inputType,omitempty"`
	Input          string          `json:"input,omitempty"`
	OutputType     int             `json:"outputType,omitempty"`
	ExpectedOutput string          `json:"expectedOutput,omitempty"`
	AnonymousLogin bool            `json:"anonymousLogin,omitempty"`
	Username       string          `json:"username,omitempty"`
	Password       string          `json:"password,omitempty"`
}

type SiteAdvancedConfig struct {
	Cookies        string `json:"cookies,omitempty"`
	UserAgent      string `json:"userAgent,omitempty"`
	Host           string `json:"host,omitempty"`
	ResponseCode   string `json:"responseCode,omitempty"`
	ResponseCheck  string `json:"responseCheck,omitempty"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	InputType      int    `json:"inputType,omitempty"`
	Input          string `json:"input,omitempty"`
	OutputType     int    `json:"outputType,omitempty"`
	ExpectedOutput string `json:"expectedOutput,omitempty"`
}

type SiteOnceRequest struct {
	UserID         string             `json:"userId,omitempty"`
	Address        string             `json:"address,omitempty"`
	IpType         string             `json:"ipType,omitempty"`
	GroupId        string             `json:"groupId,omitempty"`
	Idc            string             `json:"idc,omitempty"`
	ProtocolType   SiteOnceProtocol   `json:"protocolType,omitempty"`
	TaskType       string             `json:"taskType,omitempty"`
	Timeout        int                `json:"timeout,omitempty"`
	OnceConfig     SiteOnceConfig     `json:"onceConfig,omitempty"`
	AdvancedFlag   bool               `json:"advancedFlag,omitempty"`
	AdvancedConfig SiteAdvancedConfig `json:"advancedConfig,omitempty"`
}

type SiteOnceBaseResponse struct {
	RequestID string `json:"requestId,omitempty"`
	Message   string `json:"message,omitempty"`
	Success   bool   `json:"success,omitempty"`
	Result    struct {
		SiteID  string `json:"siteId,omitempty"`
		GroupID string `json:"groupId,omitempty"`
	} `json:"result"`
	Code int `json:"code,omitempty"`
}
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05.000Z0700"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

type SiteOnceTaskRequest struct {
	SiteID     string   `json:"siteId,omitempty"`
	UserID     string   `json:"userId,omitempty"`
	GroupID    string   `json:"groupId,omitempty"`
	PageNo     int      `json:"pageNo,omitempty"`
	PageSize   int      `json:"pageSize,omitempty"`
	FilterArea string   `json:"filterArea,omitempty"`
	FilterISP  string   `json:"filterIsp,omitempty"`
	Order      string   `json:"order,omitempty"`
	OrderBy    string   `json:"orderBy,omitempty"`
	URL        string   `json:"url,omitempty"`
	SiteIDs    []string `json:"siteIds,omitempty"`
}

type SiteOnceTaskListResponse struct {
	Result    SiteOnceTaskList `json:"result,omitempty"`
	Code      int              `json:"code,omitempty"`
	RequestID string           `json:"requestId,omitempty"`
	Message   string           `json:"message,omitempty"`
	Success   bool             `json:"success,omitempty"`
}

type SiteOnceTaskList struct {
	PageNo     int              `json:"pageNo,omitempty"`
	PageSize   int              `json:"pageSize,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
	TaskList   []SiteOnceRecord `json:"taskList,omitempty"`
}

type SiteOnceRecord struct {
	SiteID       string     `json:"siteId,omitempty"`
	GroupID      string     `json:"groupId,omitempty"`
	UserID       string     `json:"userId,omitempty"`
	TaskType     string     `json:"taskType,omitempty"`
	IPType       string     `json:"ipType,omitempty"`
	ProtocolType string     `json:"protocolType,omitempty"`
	URL          string     `json:"url,omitempty"`
	SumSampleNum int        `json:"sumSampleNum,omitempty"`
	AgentNum     int        `json:"agentNum,omitempty"`
	Success      float64    `json:"success,omitempty"`
	MonitorTime  CustomTime `json:"monitorTime,omitempty"`
	CreateTime   CustomTime `json:"createTime,omitempty"`
	Status       string     `json:"status,omitempty"`
}

type LoadDataResponse struct {
	Result    SiteOnceTask `json:"result,omitempty"`
	Code      int          `json:"code,omitempty"`
	RequestId string       `json:"requestId,omitempty"`
	Message   string       `json:"message,omitempty"`
	Success   bool         `json:"success,omitempty"`
}

type SiteOnceTask struct {
	TotalNum     int              `json:"totalNum,omitempty"`
	PageNo       int              `json:"pageNo,omitempty"`
	PageSize     int              `json:"pageSize,omitempty"`
	Order        string           `json:"order,omitempty"`
	OrderBy      string           `json:"orderBy,omitempty"`
	FilterArea   string           `json:"filterArea,omitempty"`
	FilterIsp    string           `json:"filterIsp,omitempty"`
	Status       string           `json:"status,omitempty"`
	ProtocolType SiteOnceProtocol `json:"protocolType,omitempty"`
	URL          string           `json:"url,omitempty"`
	TaskType     string           `json:"taskType,omitempty"`
	AgentNum     int              `json:"agentNum,omitempty"`
	MonitorTime  CustomTime       `json:"monitorTime,omitempty"`
	CreateTime   CustomTime       `json:"createTime,omitempty"`
	SiteId       string           `json:"siteId,omitempty"`
	JobId        string           `json:"jobId,omitempty"`
	GroupId      string           `json:"groupId,omitempty"`
	UserId       string           `json:"userId,omitempty"`
	AllAreas     []string         `json:"allAreas,omitempty"`
	MetricOrder  []string         `json:"metricOrder,omitempty"`
	OverviewData SiteOnceOverview `json:"overviewData,omitempty"`
	DetailData   []SiteOnceDetail `json:"detailData,omitempty"`
	TaskConfig   SiteOnceRequest  `json:"taskConfig,omitempty"`
}

type SiteOnceOverview struct {
	Success        float64            `json:"success,omitempty"`
	Metrics        map[string]float64 `json:"metrics,omitempty"`
	SumSampleNum   int                `json:"sumSampleNum,omitempty"`
	RightSampleNum int                `json:"rightSampleNum,omitempty"`
	ErrSampleNum   int                `json:"errSampleNum,omitempty"`
}

type SiteOnceDetail struct {
	Id             int64              `json:"id,omitempty"`
	Region         string             `json:"region,omitempty"`
	AgentProv      string             `json:"agentProv,omitempty"`
	AgentIsp       string             `json:"agentIsp,omitempty"`
	ClientIp       string             `json:"clientIp,omitempty"`
	ClientCity     string             `json:"clientCity,omitempty"`
	RemoteAddr     string             `json:"remoteAddr,omitempty"`
	RemoteArea     string             `json:"remoteArea,omitempty"`
	RemoteCity     string             `json:"remoteCity,omitempty"`
	RemoteCounty   string             `json:"remoteCounty,omitempty"`
	AnalysisResult []string           `json:"analysisResult,omitempty"`
	IpProtocol     string             `json:"ipProtocol,omitempty"`
	Metrics        map[string]float64 `json:"metrics,omitempty"`
	Success        float64            `json:"success,omitempty"`
	MonitorTime    CustomTime         `json:"monitorTime,omitempty"`
	Status         string             `json:"status,omitempty"`
}

type SiteOnceGroupTask struct {
	TotalNum     int              `json:"totalNum,omitempty"`
	SumSampleNum int              `json:"sumSampleNum,omitempty"`
	PageNo       int              `json:"pageNo,omitempty"`
	PageSize     int              `json:"pageSize,omitempty"`
	Order        string           `json:"order,omitempty"`
	OrderBy      string           `json:"orderBy,omitempty"`
	FilterArea   string           `json:"filterArea,omitempty"`
	FilterIsp    string           `json:"filterIsp,omitempty"`
	ProtocolType SiteOnceProtocol `json:"protocolType,omitempty"`
	URL          string           `json:"url,omitempty"`
	TaskType     string           `json:"taskType,omitempty"`
	GroupId      string           `json:"groupId,omitempty"`
	MetricOrder  []string         `json:"metricOrder,omitempty"`
	AllAreas     []string         `json:"allAreas,omitempty"`
	OverviewData SiteOnceOverview `json:"overviewData,omitempty"`
	DetailData   []SiteOnceDetail `json:"detailData,omitempty"`
}

type GroupTaskResponseWrapper struct {
	Result    SiteOnceGroupTask `json:"result,omitempty"`
	Code      int               `json:"code,omitempty"`
	RequestId string            `json:"requestId,omitempty"`
	Message   string            `json:"message,omitempty"`
	Success   bool              `json:"success,omitempty"`
}

type SiteAgent struct {
	AgentId    string `json:"agentId,omitempty"`
	AgentName  string `json:"agentName,omitempty"`
	Status     int    `json:"status,omitempty"`
	Ipv6Status int    `json:"ipv6Status,omitempty"`
	Region     string `json:"region,omitempty"` // Assuming SiteAgentRegion is a struct
}

type SiteOnceAgent struct {
	WhiteUser  bool        `json:"whiteUser,omitempty"`
	SiteAgents []SiteAgent `json:"siteAgents,omitempty"`
}

type SiteAgentResponseWrapper struct {
	Result    SiteOnceAgent `json:"result,omitempty"`
	Code      int           `json:"code,omitempty"`
	RequestId string        `json:"requestId,omitempty"`
	Message   string        `json:"message,omitempty"`
	Success   bool          `json:"success,omitempty"`
}
