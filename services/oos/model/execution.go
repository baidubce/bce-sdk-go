package model

type Execution struct {
	ID                string                 `json:"id,omitempty"`
	Description       string                 `json:"description,omitempty"`
	Template          *Template              `json:"template,omitempty"`
	CreatedTimestamp  int64                  `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  int64                  `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp int64                  `json:"finishedTimestamp,omitempty"`
	State             string                 `json:"state,omitempty"`
	Properties        map[string]interface{} `json:"properties,omitempty"`
	Tags              []*KV                  `json:"tags,omitempty"`
	Tasks             []*Task                `json:"tasks,omitempty"`
	Trigger           string                 `json:"trigger,omitempty"`
}

type Task struct {
	ID                string                 `json:"id,omitempty"`
	LoopIndex         int                    `json:"loopIndex,omitempty"`
	Revision          int64                  `json:"revision,omitempty"`
	CreatedTimestamp  int64                  `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  int64                  `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp int64                  `json:"finishedTimestamp,omitempty"`
	State             string                 `json:"state,omitempty"`
	Operator          *Operator              `json:"operator,omitempty"`
	Reason            string                 `json:"reason,omitempty"`
	InitContext       map[string]interface{} `json:"initContext,omitempty"`
	Context           map[string]interface{} `json:"context,omitempty"`
	OutputContext     map[string]interface{} `json:"outputContext,omitempty"`
	Tries             int                    `json:"tries,omitempty"`
	Children          []*Task                `json:"children,omitempty"`
	Log               []*Log                 `json:"log,omitempty"`
}

type Log struct {
	Timestamp string            `json:"timestamp,omitempty"`
	Level     string            `json:"level,omitempty"`
	Msg       string            `json:"msg,omitempty"`
	Tags      map[string]string `json:"tags,omitempty"`
}
