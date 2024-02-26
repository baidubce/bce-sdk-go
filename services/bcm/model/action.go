package model

type ActionNotificationType string

const (
	ActionNotificationTypeEmail ActionNotificationType = "EMAIL"
	ActionNotificationTypePhone ActionNotificationType = "PHONE"
	ActionNotificationTypeSms   ActionNotificationType = "SMS"
)

type ActionMember struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ActionDisableTime struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type ActionCallBack struct {
	Url string `json:"url,omitempty"`
}

type ActionNotification struct {
	Type ActionNotificationType `json:"type,omitempty"`
	//AliasName string                 `json:"aliasName,omitempty"`
	Receiver string `json:"receiver,omitempty"`
}

type NotifyGroup struct {
	Id          string `json:"id,omitempty"`
	DomainId    string `json:"domainId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}

type NotifyParty struct {
	Id       string `json:"id,omitempty"`
	DomainId string `json:"domainId,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Type     string `json:"type,omitempty"`
}

type ActionUserInfo struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Type  string `json:"type,omitempty"`
}

type ListNotifyGroupsRequest struct {
	Name     string `json:"name,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type ListNotifyGroupsResponse struct {
	Success bool                       `json:"success,omitempty"`
	Status  int                        `json:"status,omitempty"`
	Page    ListNotifyGroupsPageResult `json:"page,omitempty"`
}

type ListNotifyGroupsPageResult struct {
	OrderBy    string        `json:"orderBy,omitempty"`
	Order      string        `json:"order,omitempty"`
	PageNo     int           `json:"pageNo,omitempty"`
	PageSize   int           `json:"pageSize,omitempty"`
	TotalCount int           `json:"totalCount,omitempty"`
	Result     []NotifyGroup `json:"result,omitempty"`
}

type ListNotifyPartiesRequest struct {
	Name     string `json:"name,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type ListNotifyPartiesResponse struct {
	Success bool                        `json:"success,omitempty"`
	Status  int                         `json:"status,omitempty"`
	Page    ListNotifyPartiesPageResult `json:"page,omitempty"`
}

type ListNotifyPartiesPageResult struct {
	OrderBy    string        `json:"orderBy,omitempty"`
	Order      string        `json:"order,omitempty"`
	PageNo     int           `json:"pageNo,omitempty"`
	PageSize   int           `json:"pageSize,omitempty"`
	TotalCount int           `json:"totalCount,omitempty"`
	Result     []NotifyParty `json:"result,omitempty"`
}

type CreateActionRequest struct {
	UserId          string               `json:"userId,omitempty"`
	Notifications   []ActionNotification `json:"notifications,omitempty"`
	Members         []ActionMember       `json:"members,omitempty"`
	Alias           string               `json:"alias,omitempty"`
	DisableTimes    []ActionDisableTime  `json:"disableTimes,omitempty"`
	ActionCallBacks []ActionCallBack     `json:"actionCallBacks"`
}

type CreateActionResponse struct {
	Success bool        `json:"success,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type DeleteActionRequest struct {
	UserId string `json:"userId,omitempty"`
	Name   string `json:"name,omitempty"`
}

type DeleteActionResponse CreateActionResponse

type ListActionsRequest struct {
	UserId   string `json:"userId,omitempty"`
	Name     string `json:"name,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	OrderBy  string `json:"orderBy,omitempty"`
	Order    string `json:"order,omitempty"`
}

type ListActionsResponse struct {
	RequestId string                `json:"requestId,omitempty"`
	Message   string                `json:"message,omitempty"`
	Success   bool                  `json:"success,omitempty"`
	Code      int                   `json:"code,omitempty"`
	Result    ListActionsPageResult `json:"result,omitempty"`
}

type ListActionsPageResult struct {
	OrderBy    string   `json:"orderBy,omitempty"`
	Order      string   `json:"order,omitempty"`
	PageNo     int      `json:"pageNo,omitempty"`
	PageSize   int      `json:"pageSize,omitempty"`
	TotalCount int      `json:"totalCount,omitempty"`
	Result     []Action `json:"result,omitempty"`
}

type Action struct {
	ProductName      string                       `json:"productName,omitempty"`
	Name             string                       `json:"name,omitempty"`
	Alias            string                       `json:"alias,omitempty"`
	Source           string                       `json:"source,omitempty"`
	Type             string                       `json:"type,omitempty"`
	DisableTimes     []ActionDisableTime          `json:"disableTimes,omitempty"`
	Notifications    []ActionNotification         `json:"notifications,omitempty"`
	ActionCallBacks  []ActionCallBack             `json:"actionCallBacks,omitempty"`
	Members          []ActionMember               `json:"members,omitempty"`
	UserInfos        []ActionUserInfo             `json:"userInfos,omitempty"`
	GroupInfos       map[string][]*ActionUserInfo `json:"groupInfos,omitempty"`
	LastModifiedDate string                       `json:"lastModifiedDate,omitempty"`
}

type UpdateActionRequest struct {
	UserId          string               `json:"userId,omitempty"`
	Name            string               `json:"name,omitempty"`
	Notifications   []ActionNotification `json:"notifications,omitempty"`
	Members         []ActionMember       `json:"members,omitempty"`
	Alias           string               `json:"alias,omitempty"`
	DisableTimes    []ActionDisableTime  `json:"disableTimes,omitempty"`
	ActionCallBacks []ActionCallBack     `json:"actionCallBacks"`
	Source          string               `json:"source,omitempty"`
}

type UpdateActionResponse CreateActionResponse
