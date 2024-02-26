package oos

import (
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/oos/model"
)

// CreateTemplate create template
func (c *Client) CreateTemplate(req *model.Template) (*model.BaseTemplateResponse, error) {
	if err := checkTemplate(req); err != nil {
		return nil, err
	}

	req.Type = model.TemplateTypeIndividual
	result := &model.BaseTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/template").
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// CheckTemplate check template
func (c *Client) CheckTemplate(req *model.Template) (*model.CheckTemplateResponse, error) {
	if err := checkTemplate(req); err != nil {
		return nil, err
	}

	req.Type = model.TemplateTypeIndividual
	result := &model.CheckTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/template/check").
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// UpdateTemplate update template
func (c *Client) UpdateTemplate(req *model.Template) (*model.BaseResponse, error) {
	if err := checkTemplate(req); err != nil {
		return nil, err
	}

	req.Type = model.TemplateTypeIndividual
	result := &model.BaseResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/template").
		WithBody(req).
		WithMethod(http.PUT).
		WithResult(result).
		Do()
	return result, err
}

// DeleteTemplate delete template
func (c *Client) DeleteTemplate(templateId string) (*model.BaseResponse, error) {
	if len(templateId) <= 0 {
		return nil, errors.New("template id should not be empty")
	}

	result := &model.BaseResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/template").
		WithMethod(http.DELETE).
		WithQueryParam("id", templateId).
		WithResult(result).
		Do()
	return result, err
}

// GetTemplateDetail get template detail by template name
func (c *Client) GetTemplateDetail(templateName, templateType string) (*model.BaseTemplateResponse, error) {
	if len(templateName) <= 0 {
		return nil, errors.New("template name should not be empty")
	}

	if len(templateType) <= 0 {
		return nil, errors.New("template type should not be empty")
	}

	result := &model.BaseTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/template").
		WithMethod(http.GET).
		WithQueryParam("name", templateName).
		WithQueryParam("type", templateType).
		WithResult(result).
		Do()
	return result, err
}

// GetTemplateList get template list
func (c *Client) GetTemplateList(req *model.GetTemplateListRequest) (*model.GetTemplateListResponse, error) {
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should gt 0")
	}
	if req.PageSize <= 0 || req.PageSize > 100 {
		return nil, errors.New("pageSize should gt 0 and lt 100")
	}
	if len(req.Sort) <= 0 {
		req.Sort = "createTime"
	}

	result := &model.GetTemplateListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/template/list").
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetOperatorList get operator list
func (c *Client) GetOperatorList(req *model.BasePageRequest) (*model.GetOperatorListResponse, error) {
	if req.PageNo <= 0 {
		return nil, errors.New("pageNo should gt 0")
	}
	if req.PageSize <= 0 || req.PageSize > 100 {
		return nil, errors.New("pageSize should gt 0 and lt 100")
	}
	if len(req.Sort) <= 0 {
		req.Sort = "createTime"
	}

	result := &model.GetOperatorListResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v1/operator/list").
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// CreateExecution create execution
func (c *Client) CreateExecution(req *model.Execution) (*model.BaseExecutionResponse, error) {
	if req == nil {
		return nil, errors.New("request should not be null")
	}
	if req.Template == nil {
		return nil, errors.New("template should not be null")
	}
	if len(req.Template.Name) <= 0 && len(req.Template.Ref) <= 0 {
		return nil, errors.New("neither template ref nor template name is set")
	}

	result := &model.BaseExecutionResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/execution").
		WithBody(req).
		WithMethod(http.POST).
		WithResult(result).
		Do()
	return result, err
}

// GetExecutionDetail get execution detail
func (c *Client) GetExecutionDetail(executionId string) (*model.BaseExecutionResponse, error) {
	if len(executionId) <= 0 {
		return nil, errors.New("executionId should not be empty")
	}

	result := &model.BaseExecutionResponse{}
	err := bce.NewRequestBuilder(c).
		WithURL("/api/logic/oos/v2/execution").
		WithQueryParam("id", executionId).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func checkTemplate(req *model.Template) error {
	if req == nil {
		return errors.New("request should not be null")
	}
	if len(req.Name) <= 0 {
		return errors.New("name should not be empty")
	}
	if len(req.Operators) <= 0 {
		return errors.New("operators should not be empty")
	}
	if len(req.Links) <= 0 && !req.Linear {
		return errors.New("linear is false, links should not be empty")
	}
	return nil
}
