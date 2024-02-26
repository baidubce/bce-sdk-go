package model

type Operator struct {
	Name                  string                 `json:"name,omitempty"`
	Description           string                 `json:"description,omitempty"`
	Operator              string                 `json:"operator,omitempty"`
	Retries               int                    `json:"retries,omitempty"`
	RetryInterval         int                    `json:"retryInterval,omitempty"`
	Timeout               int                    `json:"timeout,omitempty"`
	ParallelismControl    *RateControl           `json:"parallelismControl,omitempty"`
	AllowedFailureControl *RateControl           `json:"allowedFailureControl,omitempty"`
	Manually              bool                   `json:"manually,omitempty"`
	ScheduleDelayMilli    int                    `json:"scheduleDelayMilli,omitempty"`
	PauseOnFailure        bool                   `json:"pauseOnFailure,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`
	InitContext           map[string]interface{} `json:"initContext,omitempty"`
}

type OperatorSpec struct {
	Name                string                 `json:"name,omitempty"`
	Label               string                 `json:"label,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Operator            string                 `json:"operator,omitempty"`
	Retries             int                    `json:"retries,omitempty"`
	RetryInterval       int                    `json:"retryInterval,omitempty"`
	Timeout             int                    `json:"timeout,omitempty"`
	ParallelismRatio    float64                `json:"parallelismRatio,omitempty"`
	ParallelismCount    int                    `json:"parallelismCount,omitempty"`
	AllowedFailureRatio float64                `json:"allowedFailureRatio,omitempty"`
	AllowedFailureCount int                    `json:"allowedFailureCount,omitempty"`
	Manually            bool                   `json:"manually,omitempty"`
	ScheduleDelayMilli  int                    `json:"scheduleDelayMilli,omitempty"`
	PauseOnFailure      bool                   `json:"pauseOnFailure,omitempty"`
	Properties          []*Property            `json:"properties,omitempty"`
	InitContext         map[string]interface{} `json:"initContext,omitempty"`
}

type RateControl struct {
	Ratio float64 `json:"ratio,omitempty"`
	Count int     `json:"count,omitempty"`
}
