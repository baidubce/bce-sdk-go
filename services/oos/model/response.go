package model

type BaseResponse struct {
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type BaseTemplateResponse struct {
	BaseResponse
	Result *Template `json:"result,omitempty"`
}

type BaseExecutionResponse struct {
	BaseResponse
	Result *Execution `json:"result,omitempty"`
}

type CheckTemplateResult struct {
	IsValid bool   `json:"isValid,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

type CheckTemplateResponse struct {
	BaseResponse
	Result *CheckTemplateResult `json:"result,omitempty"`
}

type BasePageResponse struct {
	PageNo     int    `json:"pageNo,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	TotalCount int    `json:"totalCount,omitempty"`
	OrderBy    string `json:"orderBy,omitempty"`
	Order      string `json:"order,omitempty"`
}

type GetTemplateListResult struct {
	BasePageRequest
	Templates []*Template `json:"templates,omitempty"`
}

type GetOperatorListResult struct {
	BasePageRequest
	Operators []*OperatorSpec `json:"operators,omitempty"`
}

type GetTemplateListResponse struct {
	BaseResponse
	Result *GetTemplateListResult `json:"result,omitempty"`
}

type GetOperatorListResponse struct {
	BaseResponse
	Result *GetOperatorListResult `json:"result,omitempty"`
}
