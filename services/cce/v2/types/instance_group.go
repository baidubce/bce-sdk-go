package types

const (
	DefaultShrinkPolicy               = PriorityShrinkPolicy
	PriorityShrinkPolicy ShrinkPolicy = "Priority"
	RandomShrinkPolicy   ShrinkPolicy = "Random"

	DefaultUpdatePolicy                  = ConcurrencyUpdatePolicy
	RollingUpdatePolicy     UpdatePolicy = "Rolling"
	ConcurrencyUpdatePolicy UpdatePolicy = "Concurrency"

	DefaultCleanPolicy             = RemainCleanPolicy
	RemainCleanPolicy  CleanPolicy = "Remain"
	DeleteCleanPolicy  CleanPolicy = "Delete"
)

type InstanceGroupSpec struct {
	CCEInstanceGroupID string `json:"cceInstanceGroupID,omitempty" `
	InstanceGroupName  string `json:"instanceGroupName" `

	ClusterID   string      `json:"clusterID,omitempty" `
	ClusterRole ClusterRole `json:"clusterRole,omitempty" `

	Selector *InstanceSelector `json:"selector" `

	ShrinkPolicy ShrinkPolicy `json:"shrinkPolicy,omitempty" `

	UpdatePolicy UpdatePolicy `json:"updatePolicy,omitempty" `

	CleanPolicy CleanPolicy `json:"cleanPolicy,omitempty" `

	// Deprecated: Use InstanceTemplates instead.
	InstanceTemplate  InstanceTemplate   `json:"instanceTemplate" `
	InstanceTemplates []InstanceTemplate `json:"instanceTemplates,omitempty" `

	Replicas int `json:"replicas" `

	ClusterAutoscalerSpec *ClusterAutoscalerSpec `json:"clusterAutoscalerSpec,omitempty" `

	SecurityGroups []SecurityGroupV2 `json:"securityGroups,omitempty" `
}

type InstanceTemplate struct {
	InstanceSpec `json:",inline"`
}

type InstanceSelector struct {
	LabelSelector `json:",inline"`
}

type ShrinkPolicy string
type UpdatePolicy string
type CleanPolicy string

type ClusterAutoscalerSpec struct {
	Enabled              bool `json:"enabled" `
	MinReplicas          int  `json:"minReplicas" `
	MaxReplicas          int  `json:"maxReplicas" `
	ScalingGroupPriority int  `json:"scalingGroupPriority" `
}

type InstanceGroupStatus struct {
	ReadyReplicas       int                 `json:"readyReplicas" `
	UndeliveredMachines UndeliveredMachines `json:"undeliveredMachines,omitempty" `
	Pause               *PauseDetail        `json:"pause,omitempty" `
}

type UndeliveredMachines struct {
	FailedMachines  []string `json:"failedMachines,omitempty"`
	PendingMachines []string `json:"pendingMachines,omitempty"`
}

type PauseDetail struct {
	Paused bool   `json:"paused"`
	Reason string `json:"reason"`
}

type LabelSelector struct {
	MatchLabels      map[string]string          `json:"matchLabels,omitempty" protobuf:"bytes,1,rep,name=matchLabels"`
	MatchExpressions []LabelSelectorRequirement `json:"matchExpressions,omitempty" protobuf:"bytes,2,rep,name=matchExpressions"`
}

type LabelSelectorRequirement struct {
	Key      string                `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	Operator LabelSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=LabelSelectorOperator"`
	Values   []string              `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

type LabelSelectorOperator string

const (
	LabelSelectorOpIn           LabelSelectorOperator = "In"
	LabelSelectorOpNotIn        LabelSelectorOperator = "NotIn"
	LabelSelectorOpExists       LabelSelectorOperator = "Exists"
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
)
