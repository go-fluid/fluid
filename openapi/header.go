package openapi

type HeaderObject struct {
	Description string       `json:"description,omitempty"`
	Schema      SchemaObject `json:"schema,omitempty"`
	Reference   string       `json:"$ref,omitempty"`
}
