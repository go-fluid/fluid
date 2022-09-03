package openapi

type SchemeObject struct {
	Openapi      string                      `json:"openapi,omitempty"`
	Info         InfoObject                  `json:"info,omitempty"`
	Servers      []ServerObject              `json:"servers,omitempty"`
	Tags         []TagObject                 `json:"tags,omitempty"`
	Paths        PathItemObjectMap           `json:"paths,omitempty"`
	Components   ComponentObject             `json:"components,omitempty"`
	Security     SecurityRequirementObject   `json:"security,omitempty"`
	ExternalDocs ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
