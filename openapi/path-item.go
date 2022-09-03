package openapi

type PathItemObject struct {
	Reference   string           `json:"$ref,omitempty"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	Get         *OperationObject `json:"get,omitempty"`
	Put         *OperationObject `json:"put,omitempty"`
	Post        *OperationObject `json:"post,omitempty"`
	Delete      *OperationObject `json:"delete,omitempty"`
	Options     *OperationObject `json:"options,omitempty"`
	Head        *OperationObject `json:"head,omitempty"`
	Patch       *OperationObject `json:"patch,omitempty"`
	Trace       *OperationObject `json:"trace,omitempty"`
	Servers     []ServerObject   `json:"servers,omitempty"`
	//Parameters []ParameterObject `json:"parameters,omitempty"`
}
