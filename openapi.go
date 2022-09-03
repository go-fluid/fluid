package fluid

import (
	"encoding/json"
	"github.com/go-fluid/fluid/openapi"
)

func (p Project) GenerateOpenApi() ([]byte, error) {
	scheme := openapi.SchemeObject{
		Openapi: "3.0.2",
		Info: openapi.InfoObject{
			Title:          p.Name,
			Version:        p.Version,
			Description:    p.Description,
			TermsOfService: "",                      // todo: populate termsOfService URL
			Contact:        openapi.ContactObject{}, // todo: populate contact object
			License:        openapi.LicenseObject{}, // todo: populate license object
		},
		Tags: []openapi.TagObject{
			{
				Name: "auth",
			},
		},
	}

	for _, entity := range p.Entities {

		key := camelCase(entity.NamePlural)
		scheme.Tags = append(
			scheme.Tags,
			openapi.TagObject{
				Name: key,
			},
		)
		scheme.Paths = append(
			scheme.Paths,
			openapi.PathItemObjectMapItem{
				Key:   key,
				Value: openapi.PathItemObject{},
			},
		)

	}

	output, err := json.MarshalIndent(scheme, "", "  ")
	if err != nil {
		return nil, err
	}
	return output, err
}
