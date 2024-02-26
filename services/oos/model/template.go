package model

type Template struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Ref         string       `json:"ref,omitempty"`
	Type        TemplateType `json:"type,omitempty"`
	Description string       `json:"description,omitempty"`
	Tags        []*KV        `json:"tags,omitempty"`
	Operators   []*Operator  `json:"operators,omitempty"`
	Properties  []*Property  `json:"properties,omitempty"`
	Links       []*Link      `json:"links,omitempty"`
	Linear      bool         `json:"linear,omitempty"`
}

type Link struct {
	Src string `json:"src,omitempty"`
	Dst string `json:"dst,omitempty"`
}

type Property struct {
	Name         string        `json:"name,omitempty"`
	Type         string        `json:"type,omitempty"`
	Required     bool          `json:"required,omitempty"`
	Multiple     bool          `json:"multiple,omitempty"`
	Label        string        `json:"label,omitempty"`
	Description  string        `json:"description,omitempty"`
	Options      []interface{} `json:"options,omitempty"`
	Value        interface{}   `json:"value,omitempty"`
	DefaultValue interface{}   `json:"defaultValue,omitempty"`
	Unit         string        `json:"unit,omitempty"`
}
