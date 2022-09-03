package openapi

type ParameterObject struct {
	Name            string       `json:"name,omitempty"`
	In              string       `json:"in,omitempty"`
	Description     string       `json:"description,omitempty"`
	Required        bool         `json:"required,omitempty"`
	Deprecated      bool         `json:"deprecated,omitempty"`
	AllowEmptyValue bool         `json:"allowEmptyValue,omitempty"`
	Schema          SchemaObject `json:"schema,omitempty"`
}
