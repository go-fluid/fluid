package openapi

type SchemaObject struct {
	Type       string          `json:"type,omitempty"`
	Format     string          `json:"format,omitempty"`
	Default    interface{}     `json:"default,omitempty"`
	Enum       []string        `json:"enum,omitempty"`
	Reference  string          `json:"$ref,omitempty"`
	Items      *SchemaObject   `json:"items,omitempty"`
	Properties SchemaObjectMap `json:"properties,omitempty"`
	Required   []string        `json:"required,omitempty"`
}
