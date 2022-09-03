package openapi

type ComponentObject struct {
	Schemas   SchemaObjectMap         `json:"schemas,omitempty"`
	Responses ResponseObjectMap       `json:"responses,omitempty"`
	Headers   HeaderObjectMap         `json:"headers,omitempty"`
	Security  SecuritySchemeObjectMap `json:"securitySchemes,omitempty"`
}
