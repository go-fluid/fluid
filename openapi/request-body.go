package openapi

type RequestBody struct {
	Required    bool               `json:"required,omitempty"`
	Description string             `json:"description,omitempty"`
	Content     MediaTypeObjectMap `json:"content,omitempty"`
}
