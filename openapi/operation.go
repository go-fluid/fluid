package openapi

type OperationObject struct {
	Tags         []string                    `json:"tags,omitempty"`
	Summary      string                      `json:"summary,omitempty"`
	Description  string                      `json:"description,omitempty"`
	ExternalDocs ExternalDocumentationObject `json:"externalDocs,omitempty"`
	OperationId  string                      `json:"operationId,omitempty"`
	Parameters   []ParameterObject           `json:"parameters,omitempty"`
	RequestBody  RequestBody                 `json:"requestBody,omitempty"`
	Responses    ResponseObjectMap           `json:"responses,omitempty"`
	Deprecated   bool                        `json:"deprecated,omitempty"`
	Security     SecurityRequirementObject   `json:"security,omitempty"`
	Servers      []ServerObject              `json:"servers,omitempty"`
}
