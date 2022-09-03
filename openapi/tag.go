package openapi

type TagObject struct {
	Name         string                       `json:"name,omitempty"`
	Description  string                       `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
