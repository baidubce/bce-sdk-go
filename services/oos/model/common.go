package model

type TemplateType string

const (
	TemplateTypeGlobal     TemplateType = "GLOBAL"
	TemplateTypeIndividual TemplateType = "INDIVIDUAL"
)

type KV struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
