package openapi

type InfoObject struct {
	Title          string        `json:"title,omitempty"`
	Version        string        `json:"version,omitempty"`
	Description    string        `json:"description,omitempty"`
	TermsOfService string        `json:"termsOfService,omitempty"`
	Contact        ContactObject `json:"contact,omitempty"`
	License        LicenseObject `json:"license,omitempty"`
}
