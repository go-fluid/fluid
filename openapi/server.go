package openapi

type ServerObject struct {
	Url         string                  `json:"url,omitempty"`
	Description string                  `json:"description,omitempty"`
	Variables   ServerVariableObjectMap `json:"variables,omitempty"`
}
