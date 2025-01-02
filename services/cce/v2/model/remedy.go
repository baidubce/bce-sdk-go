package model

import nrv1 "github.com/baidubce/bce-sdk-go/services/cce/v2/model/remedy/api/v1"

type ConditionType string
type MsgTplType string

type RemedyRule struct {
	ClusterID    string `json:"clusterID"`
	RemedyRuleID string `json:"remedyRuleID,omitempty"`

	ObjectMeta `json:"objectMeta,omitempty"`
	Spec       *RemedyRuleSpec `json:"spec"`
}

type ObjectMeta struct {
	ID                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	CreationTimestamp nrv1.Time         `json:"creationTimestamp,omitempty"`
	UpdateTimestamp   *nrv1.Time        `json:"updateTimestamp,omitempty"`
	DeletionTimestamp *nrv1.Time        `json:"deletionTimestamp,omitempty"`
}

type RemedyRuleSpec struct {
	Conditions     []RemedyCondition    `json:"conditions,omitempty"`
	NodeSelector   nrv1.LabelSelector   `json:"labelSelector,omitempty"`
	InstanceGroups []InstanceGroupTuple `json:"instanceGroups,omitempty"`
}

type InstanceGroupTuple struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RemedyCondition struct {
	Type               ConditionType       `json:"type"`
	ConfigType         nrv1.ConfigType     `json:"configType"`
	EnableCheck        bool                `json:"enableCheck"`
	EnableRemediation  bool                `json:"enableRemediation"`
	Steps              []RemedyStep        `json:"steps,omitempty"`
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`
}

type NotificationConfig struct {
	WebhookURL string `json:"webhookURL"`
}

type RemedyStep struct {
	Name                       string                            `json:"name"`
	EnableNotification         bool                              `json:"enableNotification"`
	DeleteNodeConfig           *RemedyDeleteNodeConfig           `json:"deleteNodeConfig,omitempty"`
	RequestAuthorizationConfig *RemedyRequestAuthorizationConfig `json:"requestAuthorizationConfig,omitempty"`
}

type RemedyDeleteNodeConfig struct {
	KeepInstanceCount    bool `json:"keepInstanceCount,omitempty"`
	KeepPostPaidInstance bool `json:"keepPostPaidInstance,omitempty"`
}

type RemedyRequestAuthorizationConfig struct {
	Automatic bool `json:"automatic"`

	RequestAuthWebhookURL string `json:"requestAuthWebhookURL,omitempty"`
}

type RemedyRuleBinding struct {
	ClusterID string `json:"clusterID,omitempty"`

	InstanceGroupID string `json:"instanceGroupID,omitempty"`
	RemedyRuleID    string `json:"remedyRuleID,omitempty"`
}

type CreateRemedyRuleResponse struct {
	RemedyRule *RemedyRule `json:"remedyRule,omitempty"`
	RequestID  string      `json:"requestID"`
}

type RemedyRuleOptions struct {
	ClusterID    string `json:"clusterID,omitempty"`
	RemedyRuleID string `json:"remedyRuleID,omitempty"`
}

type GetRemedyRuleResponse struct {
	RemedyRule *RemedyRule `json:"remedyRule,omitempty"`
	RequestID  string      `json:"requestID"`
}

type RemedyRulePage struct {
	OrderBy     string        `json:"orderBy,omitempty"`
	Order       string        `json:"order,omitempty"`
	PageNo      int           `json:"pageNo,omitempty"`
	PageSize    int           `json:"pageSize,omitempty"`
	TotalCount  int           `json:"totalCount,omitempty"`
	RemedyRules []*RemedyRule `json:"remedyRules,omitempty"`
}

type ListRemedyRulesResponse struct {
	RemedyRulePage *RemedyRulePage `json:"remedyRulePage,omitempty"`
	RequestID      string          `json:"requestID" json:"RequestID,omitempty"`
}

type ListRemedyRulesOptions struct {
	ClusterID string `json:"clusterID,omitempty"`

	OrderBy  string `json:"orderBy,omitempty"`
	Order    string `json:"order,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type ListRemedyTaskOptions struct {
	ClusterID string `json:"clusterID,omitempty"`

	InstanceGroupID string `json:"instanceGroupID,omitempty"`
	PageNo          int    `json:"pageNo,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	StartTime       string `json:"startTime,omitempty"`
	StopTime        string `json:"stopTime,omitempty"`
}

type GetRemedyTaskOptions struct {
	ClusterID       string `json:"clusterID,omitempty"`
	InstanceGroupID string `json:"instanceGroupID,omitempty"`
	RemedyTaskID    string `json:"remedyTaskID,omitempty"`
}

type RemedyTaskPage struct {
	TotalCount  int          `json:"totalCount,omitempty"`
	PageNo      int          `json:"pageNo,omitempty"`
	PageSize    int          `json:"pageSize,omitempty"`
	RemedyTasks []RemedyTask `json:"remedyTasks,omitempty"`
}

type RemedyTask struct {
	nrv1.TypeMeta   `json:",inline"`
	nrv1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RemedyTaskSpec   `json:"spec,omitempty"`
	Status RemedyTaskStatus `json:"status,omitempty"`
}

type RemedyTaskSpec = nrv1.RemedyTaskSpec

type RemedyTaskStatus struct {
	ReconcileSteps       []RemedyTaskStep   `json:"reconcileSteps,omitempty"`
	ReconcileDeleteSteps []RemedyTaskStep   `json:"reconcileDeleteSteps,omitempty"`
	InstancePhase        nrv1.InstancePhase `json:"instancePhase,omitempty"`
	LastFinished         *nrv1.Time         `json:"lastFinished,omitempty"`
	// 最近发起维修时间
	LastStarted *nrv1.Time `json:"lastStarted,omitempty"`
	Message     string     `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	Healthy     bool       `json:"healthy,omitempty"`
}

type RemedyTaskStep struct {
	Name      nrv1.StepName `json:"name"`
	nrv1.Step `json:",inline"`
}

type UpdateRemedyRuleOptions struct {
	ClusterID    string `json:"clusterID,omitempty"`
	RemedyRuleID string `json:"remedyRuleID,omitempty"`
}

type UpdateRemedyRuleResponse struct {
	RemedyRule *RemedyRule `json:"remedyRule,omitempty"`
	RequestID  string      `json:"requestID"`
}

type DeleteRemedyRuleResponse struct {
	RequestID string `json:"requestID"`
}

type BindInstanceGroupResponse struct {
	RequestID string `json:"requestID"`
}

type UnbindInstanceGroupResponse struct {
	RequestID string `json:"requestID"`
}

type ListRemedyTasksResponse struct {
	RemedyTaskPage *RemedyTaskPage `json:"remedyTaskPage,omitempty"`
	RequestID      string          `json:"requestID"`
}

type GetRemedyTaskResponse struct {
	RemedyTask *RemedyTask `json:"remedyTask,omitempty"`
	RequestID  string      `json:"requestID"`
}

// 确认任务恢复
type ConfirmRemedyTask struct {
	ClusterID       string `json:"clusterID,omitempty"`
	InstanceGroupID string `json:"instanceGroupID,omitempty"`
	RemedyTaskID    string `json:"remedyTaskID,omitempty"`

	IsConfirmed bool `json:"isConfirmed,omitempty"`
}

type ConfirmRemedyTaskResponse struct {
	IsConfirmed bool   `json:"isConfirmed,omitempty"`
	RequestID   string `json:"requestID"`
}

// 检查任务恢复状态
type CheckRemedyTaskConfirmedStatusResponse struct {
	IsConfirmed bool   `json:"isConfirmed,omitempty"`
	RequestID   string `json:"requestID"`
}

// 授权任务
type AuthorizeTaskArgs struct {
	ClusterID       string `json:"clusterID,omitempty"`
	InstanceGroupID string `json:"instanceGroupID,omitempty"`
	RemedyTaskID    string `json:"remedyTaskID,omitempty"`

	IsAuthorized bool `json:"isAuthorized,omitempty"`
}

type AuthorizeTaskResponse struct {
	IsAuthorized bool   `json:"isAuthorized,omitempty"`
	RequestID    string `json:"requestID"`
}

// webhook 请求链接
type RequestAuthorizeArgs struct {
	ClusterID       string `json:"clusterID,omitempty"`
	InstanceGroupID string `json:"instanceGroupID,omitempty"`
	RemedyTaskID    string `json:"remedyTaskID,omitempty"`

	Webhook    string          `json:"webhook,omitempty"`
	MsgTplType MsgTplType      `json:"msgTplType,omitempty"`
	Content    *WebhookContent `json:"content,omitempty"`
}

type WebhookContent struct {
	TaskID               string `json:"taskId,omitempty"`
	RemedyTaskName       string `json:"remedyTaskName,omitempty"`
	EventInfo            string `json:"eventInfo,omitempty"`
	InstanceInfo         string `json:"instanceInfo,omitempty"`
	InstanceGroupID      string `json:"instanceGroupId,omitempty"`
	InstanceGroupName    string `json:"instanceGroupName,omitempty"`
	InstanceType         string `json:"instanceType,omitempty"`
	Region               string `json:"region,omitempty"`
	TaskTime             string `json:"taskTime,omitempty"`
	ClusterID            string `json:"clusterId,omitempty"`
	Message              string `json:"message,omitempty"`
	KeepPostPaidInstance bool   `json:"keepPostPaidInstance,omitempty"`
	KeepInstanceCount    bool   `json:"keepInstanceCount,omitempty"`
}

type RequestAuthorizeResponse struct {
	RequestID string `json:"requestID"`
}

type AuthRemedyTaskReq struct {
	InstanceGroupID string `json:"instanceGroupID,omitempty"`
	RemedyTaskID    string `json:"remedyTaskID,omitempty"`
}

type AuthRemedyTaskResponse struct {
	RequestID string `json:"requestID"`
}

type CheckWebhookAddressRequest struct {
	ClusterID      string `json:"clusterId,omitempty"`
	WebhookAddress string `json:"webhookAddress"`
}

type CheckWebhookAddressResult struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
}

type CheckWebhookAddressResponse struct {
	Result    *CheckWebhookAddressResult `json:"result,omitempty"`
	RequestID string                     `json:"requestID"`
}

type BindingOrUnBindingRequest struct {
	ClusterID       string `json:"clusterID,omitempty"`
	InstanceGroupID string `json:"instanceGroupID,omitempty"`

	RemedyRulesBinding *RemedyRulesBinding `json:"remedyRulesBinding,omitempty"`
}

type RemedyRulesBinding struct {
	RemedyRuleID         string `json:"remedyRuleID,omitempty"`
	EnableCheckANDRemedy bool   `json:"enableCheckANDRemedy,omitempty"`
}

type BindingOrUnBindingResponse struct {
	Result    *BindingOrUnBindingRequest `json:"result,omitempty"`
	RequestID string                     `json:"requestID"`
}
