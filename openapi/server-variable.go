package openapi

type ServerVariableObject struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
}
