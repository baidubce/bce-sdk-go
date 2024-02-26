package bcm

import (
	"errors"
	"strconv"
)

const (
	CustomAlarmBaseURL           = "/csm/api/v1/custom/alarm/configs"
	CreateCustomAlarmConfig      = CustomAlarmBaseURL + "/create"
	UpdateCustomAlarmConfig      = CustomAlarmBaseURL + "/update"
	BatchDeleteCustomAlarmPolicy = CustomAlarmBaseURL + "/delete"
	ListCustomAlarmPolicy        = CustomAlarmBaseURL + "/list"
	DetailCustomAlarmPolicy      = CustomAlarmBaseURL + "/detail"
	BlockCustomAlarmPolicy       = CustomAlarmBaseURL + "/block"
	UnblockCustomAlarmPolicy     = CustomAlarmBaseURL + "/unblock"
)

type AlarmLevel string

const (
	NOTICE   AlarmLevel = "NOTICE"
	WARNING  AlarmLevel = "WARNING"
	CRITICAL AlarmLevel = "CRITICAL"
	MAJOR    AlarmLevel = "MAJOR"
	CUSTOM   AlarmLevel = "CUSTOM"
)

type MetricDimensions struct {
	Name  string   `json:"name"`
	Value []string `json:"value,omitempty"`
}

type CustomAlarmRule struct {
	ID                 *int64             `json:"id,omitempty"`
	Index              int                `json:"index"`
	MetricName         string             `json:"metricName"`
	Dimensions         []MetricDimensions `json:"dimensions"`
	Statistics         string             `json:"statistics,omitempty"`
	Threshold          string             `json:"threshold,omitempty"`
	ComparisonOperator string             `json:"comparisonOperator,omitempty"`
	Cycle              int                `json:"cycle,omitempty"`
	Count              int                `json:"count,omitempty"`
	Function           string             `json:"function,omitempty"`
}

type CustomAlarmConfig struct {
	Comment             string            `json:"comment,omitempty"`
	UserID              string            `json:"userId"`
	AlarmName           string            `json:"alarmName"`
	Namespace           string            `json:"namespace"`
	Level               AlarmLevel        `json:"level"`
	ActionEnabled       bool              `json:"actionEnabled,omitempty"`
	PolicyEnabled       bool              `json:"policyEnabled,omitempty"`
	AlarmActions        []string          `json:"alarmActions"`
	OkActions           []string          `json:"okActions,omitempty"`
	InsufficientActions []string          `json:"insufficientActions"`
	InsufficientCycle   int               `json:"insufficientCycle,omitempty"`
	Rules               []CustomAlarmRule `json:"rules"`
	Region              string            `json:"region,omitempty"`
	CallbackURL         string            `json:"callbackUrl,omitempty"`
	CallbackToken       string            `json:"callbackToken,omitempty"`
	Tag                 string            `json:"tag,omitempty"`
	RepeatAlarmCycle    int               `json:"repeatAlarmCycle,omitempty"`
	MaxRepeatCount      int               `json:"maxRepeatCount,omitempty"`
}

func (config *CustomAlarmConfig) Validate() error {
	// 检查必填字段
	if config.UserID == "" {
		return errors.New("userId should not be empty")
	}
	if config.Region == "" {
		return errors.New("region should not be empty")
	}
	if config.Namespace == "" {
		return errors.New("namespace should not be empty")
	}
	if config.AlarmName == "" {
		return errors.New("alarmName should not be empty")
	}
	if config.Level == "" {
		return errors.New("level should not be empty")
	}
	if len(config.Rules) == 0 {
		return errors.New("rules should not be empty")
	}

	return nil
}

type ComponentType string

type AlarmPolicyBatch struct {
	UserId    string   `json:"userId,omitempty"`
	Scope     string   `json:"scope,omitempty"`
	AlarmName []string `json:"alarmName,omitempty"`
}

type EventAlarmPolicyBatch struct {
	AlarmPolicyBatch
	PolicyType ComponentType `json:"policyType,omitempty"`
}

type AlarmPolicyBatchList struct {
	MetricAlarmList      []AlarmPolicyBatch      `json:"metricAlarmList,omitempty"`
	EventAlarmList       []EventAlarmPolicyBatch `json:"eventAlarmList,omitempty"`
	CustomAlarmList      []AlarmPolicyBatch      `json:"customAlarmList,omitempty"`
	CustomEventAlarmList []AlarmPolicyBatch      `json:"customEventAlarmList,omitempty"`
}

func (a *AlarmPolicyBatchList) Validate(typeName string) error {
	if typeName == "CustomAlarmList" {
		if a.CustomAlarmList == nil {
			return errors.New("customAlarmList should not be empty")
		}
		for _, apb := range a.CustomAlarmList {
			if apb.UserId == "" {
				return errors.New("userId should not be empty")
			}
			if apb.Scope == "" {
				return errors.New("Scope should not be nil")
			}
			if apb.AlarmName == nil {
				return errors.New("AlarmName should not be empty")
			}
		}
	}
	return nil
}

type ListCustomPolicyPageResultResponse struct {
	OrderBy    string              `json:"orderBy,omitempty"`
	Order      string              `json:"order,omitempty"`
	PageNo     int                 `json:"pageNo,omitempty"`
	PageSize   int                 `json:"pageSize,omitempty"`
	TotalCount int                 `json:"totalCount,omitempty"`
	Result     []CustomAlarmConfig `json:"result,omitempty"`
}

type ListCustomAlarmPolicyParams struct {
	UserId        string `json:"userId,omitempty"`
	AlarmName     string `json:"alarmName,omitempty"`
	Namespace     string `json:"namespace,omitempty"`
	ActionEnabled bool   `json:"actionEnabled,omitempty"`
	PageNo        int    `json:"pageNo,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
}

func (params *ListCustomAlarmPolicyParams) ConvertToBceParams() map[string]string {
	bceParams := make(map[string]string)
	if params.UserId != "" {
		bceParams["userId"] = params.UserId
	}
	if params.AlarmName != "" {
		bceParams["alarmName"] = params.AlarmName
	}
	if params.Namespace != "" {
		bceParams["namespace"] = params.Namespace
	}
	if params.PageNo != 0 {
		bceParams["pageNo"] = strconv.Itoa(params.PageNo)
	}
	if params.PageSize != 0 {
		bceParams["pageSize"] = strconv.Itoa(params.PageSize)
	}
	return bceParams
}

type DetailCustomAlarmPolicyParams struct {
	UserId    string `json:"userId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
}

func (params *DetailCustomAlarmPolicyParams) ConvertToBceParams() map[string]string {
	bceParams := make(map[string]string)
	if params.UserId != "" {
		bceParams["userId"] = params.UserId
	}
	if params.AlarmName != "" {
		bceParams["alarmName"] = params.AlarmName
	}
	if params.Namespace != "" {
		bceParams["namespace"] = params.Namespace
	}
	return bceParams
}

type BlockCustomAlarmPolicyParams struct {
	UserId    string `json:"userId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
}

func (params *BlockCustomAlarmPolicyParams) ConvertToBceParams() map[string]string {
	bceParams := make(map[string]string)
	if params.UserId != "" {
		bceParams["userId"] = params.UserId
	}
	if params.AlarmName != "" {
		bceParams["alarmName"] = params.AlarmName
	}
	if params.Namespace != "" {
		bceParams["namespace"] = params.Namespace
	}
	return bceParams
}

type UnblockCustomAlarmPolicyParams struct {
	UserId    string `json:"userId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
}

func (params *UnblockCustomAlarmPolicyParams) ConvertToBceParams() map[string]string {
	bceParams := make(map[string]string)
	if params.UserId != "" {
		bceParams["userId"] = params.UserId
	}
	if params.AlarmName != "" {
		bceParams["alarmName"] = params.AlarmName
	}
	if params.Namespace != "" {
		bceParams["namespace"] = params.Namespace
	}
	return bceParams
}
