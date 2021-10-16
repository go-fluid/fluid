package fluid

import "net/http"

type MaterialColor struct {
	Main string `json:"main"`
	Contrast string `json:"contrast"`
	Shade string `json:"shade"`
	Tint string `json:"tint"`
}

type Theme struct {
	Primary MaterialColor `json:"primary"`
	Secondary MaterialColor `json:"secondary"`
	Tertiary MaterialColor `json:"tertiary"`
	Success MaterialColor `json:"success"`
	Warning MaterialColor `json:"warning"`
	Error MaterialColor `json:"error"`
	Medium MaterialColor `json:"medium"`
	Light MaterialColor `json:"light"`
	Dark MaterialColor `json:"dark"`
}

const (
	PortalTypeIonic = "ionic"
	PortalTypeVuetify = "vuetify"
)

type Portal struct {
	Name string `json:"name"`
	Type string `json:"type"`
	AccountEntityKeys []string `json:"accountEntityKeys"` // list of different account entities that can log into this platform
	LightTheme Theme `json:"lightTheme"`
	DarkTheme Theme `json:"darkTheme"`
}

const (
	ContractFieldTypeFile = "file" // only available on request type
	ContractFieldTypeBinary = "binary" // only available on response type
	ContractFieldTypeString = "string"
	ContractFieldTypeUuid = "uuid"
	ContractFieldTypeDate = "date"
	ContractFieldTypeDateTime = "date-time"
	ContractFieldTypeTime = "time"
	ContractFieldTypeInteger = "integer"
	ContractFieldTypeDecimal = "decimal"
	ContractFieldTypeBoolean = "boolean"
	ContractFieldTypeMoney = "money"
	ContractFieldTypeUnsignedInteger = "unsigned integer"

)

type ContractField struct {
	Name string `json:"name"`
	Alias string `json:"alias"`
	EnableMultipleValueSupport bool `json:"enableMultipleValueSupport"`
	Type string `json:"type"`
}

const (
	ContractTypeRequest = "request"
	ContractTypeResponse = "response"
	ContractTypeParameters = "parameters"
)

type Contract struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

const (
	EntityActionMethodGet = http.MethodGet
	EntityActionMethodPost = http.MethodPost
	EntityActionMethodPut = http.MethodPut
	EntityActionMethodDelete = http.MethodDelete

	EntityActionTypeList = "list"
	EntityActionTypeRecord = "record"
)

type EntityAction struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Method string `json:"method"`
	Type string `json:"type"`
	EnableFileDownloadResponse bool `json:"enableFileDownloadResponse"`
	DisableAuthorizationRequirement bool `json:"disableAuthorizationRequirement"`
	RequestParametersContractKey string `json:"requestParametersContractKey"`
	RequestBodyContractKey string `json:"requestBodyContractKey"`
	ResponseBodyContractKey string `json:"responseBodyContractKey"` // optional and only available if enableFileDownloadResponse is false
}

type EntityFieldOption struct {
	Order int `json:"order"` // value for integer type
	Name string `json:"name"` // value for string type
	Text string `json:"text"` // human friendly text
}

const (
	EntityFieldTypePassword = "password"
	EntityFieldTypeBinary = "binary"
	EntityFieldTypeString = "string"
	EntityFieldTypeUuid = "uuid"
	EntityFieldTypeDate = "date"
	EntityFieldTypeDateTime = "date-time"
	EntityFieldTypeTime = "time"
	EntityFieldTypeInteger = "integer"
	EntityFieldTypeDecimal = "decimal"
	EntityFieldTypeBoolean = "boolean"
	EntityFieldTypeMoney = "money" // money will be stored as an integer to 10^5 precision
	EntityFieldTypeAttribute = "attribute" // attribute denotes an EAV type field
)

type EntityField struct {
	Group string `json:"group"`
	Order int `json:"order"`
	Name string `json:"name"`
	Label string `json:"label"`
	EnableMultipleValueSupport bool `json:"enableMultipleValueSupport"`
	Type string `json:"type"`
	EnableOptionsSupport bool `json:"enableOptionsSupport"`
	Pattern string `json:"pattern"` // regex which is only available for 'string', 'password' types and only if enableOptionsSupport is false (remember to show testing field and help on front-end)
	DefaultValue string `json:"defaultValue"` // (validate on screen using enums, pattern, etc.)
	NotHeader bool `json:"notHeader"` // [not available for password type]
	NotSortable bool `json:"notSortable"` // [not available for password type] only available if notHeader is false otherwise force to false
	NotFilterable bool `json:"notFilterable"` // [not available for password type]
	NotEditable bool `json:"notEditable"` // [not available for password type]
	IsOptional bool `json:"isOptional"` // [not available for password type]
	IsProtected bool `json:"isProtected"` // [only available for string type] denotes a piece of personal information that must be protected (encryption and audit trail)
	IsEncrypted bool `json:"isEncrypted"` // [only available for string type] only available if isProtect is false otherwise force to true
	IsIdentifier bool `json:"isIdentifier"` // [only available for string type] only available if isProtect is false otherwise force to true
	Options map[string]EntityFieldOption `json:"options"`
}

type Entity struct {
	NameSingular string `json:"nameSingular"`
	NamePlural string `json:"namePlural"`
	EnableFreeTextSearch bool `json:"enableFreeTextSearch"`
	DisableList bool `json:"disableList"`
	DisableCreate bool `json:"disableCreate"`
	DisableUpdate bool `json:"disableUpdate"`
	DisableDelete bool `json:"disableDelete"`
	Fields map[string]EntityField `json:"fields"`
	Actions map[string]EntityAction `json:"actions"`
}

type Project struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Version string `json:"version"`
	Portals map[string]Portal `json:"portals"`
	Entities map[string]Entity `json:"entities"`
	Contracts map[string]Contract `json:"contracts"`
}
