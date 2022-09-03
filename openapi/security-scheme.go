package openapi

type SecuritySchemeObject struct {
	Type             string `json:"type,omitempty"`
	Description      string `json:"description,omitempty"`
	Name             string `json:"name,omitempty"`
	In               string `json:"in,omitempty"`
	Scheme           string `json:"scheme,omitempty"`
	BearerFormat     string `json:"bearerFormat,omitempty"`
	OpenIdConnectUrl string `json:"openIdConnectUrl,omitempty"`
}
