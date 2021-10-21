package fluid

import "net/http"

const (
	EntityActionMethodGet    = http.MethodGet
	EntityActionMethodPost   = http.MethodPost
	EntityActionMethodPut    = http.MethodPut
	EntityActionMethodDelete = http.MethodDelete

	EntityActionTypeList   = "list"
	EntityActionTypeRecord = "record"

	EntityFieldTypePassword  = "password"
	EntityFieldTypeBinary    = "binary"
	EntityFieldTypeString    = "string"
	EntityFieldTypeUuid      = "uuid"
	EntityFieldTypeDate      = "date"
	EntityFieldTypeDateTime  = "date-time"
	EntityFieldTypeTime      = "time"
	EntityFieldTypeInteger   = "integer"
	EntityFieldTypeDecimal   = "decimal"
	EntityFieldTypeBoolean   = "boolean"
	EntityFieldTypeMoney     = "money"     // money will be stored as an integer to 10^5 precision
	EntityFieldTypeAttribute = "attribute" // attribute denotes an EAV type field
)

type EntityAction struct {
	Name                            string `json:"name" validate:"required"`
	Description                     string `json:"description" validate:"required"`
	Method                          string `json:"method" validate:"required,entityActionMethod"`
	Type                            string `json:"type" validate:"required,entityActionType"`
	EnableFileDownloadResponse      bool   `json:"enableFileDownloadResponse,omitempty"`
	DisableAuthorizationRequirement bool   `json:"disableAuthorizationRequirement,omitempty"`
	RequestParametersContractKey    string `json:"requestParametersContractKey,omitempty"` // optional and only available if method is GET
	RequestBodyContractKey          string `json:"requestBodyContractKey,omitempty"`       // optional and only available if method is not GET
	ResponseBodyContractKey         string `json:"responseBodyContractKey,omitempty"`      // optional and only available if enableFileDownloadResponse is false
}

type EntityFieldOption struct {
	Value interface{} `json:"value" validate:"required,entityFieldOptionValue"` // integer or string based on field type
	Text  string      `json:"text" validate:"required"`                         // human friendly text
}

type EntityField struct {
	Group                      string              `json:"group,omitempty"`
	Name                       string              `json:"name" validate:"required"`
	Description                string              `json:"description" validate:"required"`
	Label                      string              `json:"label,omitempty"`
	Hint                       string              `json:"hint,omitempty"`             // field hint to display on front end
	EnableMultipleValueSupport bool                `json:"enableMultipleValueSupport"` // [not available for password type]
	Type                       string              `json:"type" validate:"required,entityFieldType"`
	EnableOptionsSupport       bool                `json:"enableOptionsSupport,omitempty"`    // [only available for string, integer type]
	Pattern                    string              `json:"pattern,omitempty"`                 // [not available for password type] regex which is only available for 'string', 'password' types and only if enableOptionsSupport is false (remember to show testing field and help on front-end)
	DefaultValue               string              `json:"defaultValue,omitempty"`            // (validate on screen using enums, pattern, etc.)
	NotHeader                  bool                `json:"notHeader,omitempty"`               // [not available for password type]
	NotSortable                bool                `json:"notSortable,omitempty"`             // [not available for password type] only available if notHeader is false otherwise force to false
	NotFilterable              bool                `json:"notFilterable,omitempty"`           // [not available for password type]
	NotEditable                bool                `json:"notEditable,omitempty"`             // [not available for password type]
	IsOptional                 bool                `json:"isOptional,omitempty"`              // [not available for password type]
	IsProtected                bool                `json:"isProtected,omitempty"`             // [only available for string type] denotes a piece of personal information that must be protected (encryption and audit trail)
	IsEncrypted                bool                `json:"isEncrypted,omitempty"`             // [only available for string type] only available if isProtect is false otherwise force to true
	IsIdentifier               bool                `json:"isIdentifier,omitempty"`            // [only available for string type] only available if isProtect is false otherwise force to true
	Options                    []EntityFieldOption `json:"options,omitempty" validate:"dive"` // [only available for string, integer type]
}

type Entity struct {
	NameSingular         string         `json:"nameSingular" validate:"required"`
	NamePlural           string         `json:"namePlural" validate:"required"`
	EnableFreeTextSearch bool           `json:"enableFreeTextSearch,omitempty"`
	DisableList          bool           `json:"disableList,omitempty"`
	DisableCreate        bool           `json:"disableCreate,omitempty"`
	DisableUpdate        bool           `json:"disableUpdate,omitempty"`
	DisableDelete        bool           `json:"disableDelete,omitempty"`
	Fields               []EntityField  `json:"fields" validate:"required,dive"`
	Actions              []EntityAction `json:"actions,omitempty" validate:"dive"`
}
