package openapi

type ResponseObject struct {
	Reference   string             `json:"$ref,omitempty"`
	Description string             `json:"description,omitempty"`
	Headers     HeaderObjectMap    `json:"headers,omitempty"`
	Content     MediaTypeObjectMap `json:"content,omitempty"`
}
