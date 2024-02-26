package model

type BasePageRequest struct {
	Ascending bool   `json:"ascending,omitempty"`
	Sort      string `json:"sort,omitempty"`
	PageNo    int    `json:"pageNo,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

type GetTemplateListRequest struct {
	BasePageRequest
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
