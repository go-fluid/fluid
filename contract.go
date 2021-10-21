package fluid

const (
	ContractTypeRequest    = "request"
	ContractTypeResponse   = "response"
	ContractTypeParameters = "parameters"

	ContractFieldTypeFile     = "file"   // only available on request type
	ContractFieldTypeBinary   = "binary" // only available on response type
	ContractFieldTypeString   = "string"
	ContractFieldTypeUuid     = "uuid"
	ContractFieldTypeDate     = "date"
	ContractFieldTypeDateTime = "date-time"
	ContractFieldTypeTime     = "time"
	ContractFieldTypeInteger  = "integer"
	ContractFieldTypeDecimal  = "decimal"
	ContractFieldTypeBoolean  = "boolean"
	ContractFieldTypeMoney    = "money"
)

type ContractField struct {
	Name                       string `json:"name" validate:"required"`
	Description                string `json:"description" validate:"required"`
	Label                      string `json:"label,omitempty"`
	EnableMultipleValueSupport bool   `json:"enableMultipleValueSupport"`
	Type                       string `json:"type" validate:"required,contractFieldType"`
}

type Contract struct {
	Key    string          `json:"key" validate:"required"`
	Name   string          `json:"name" validate:"required"`
	Type   string          `json:"type" validate:"required,contractType"`
	Fields []ContractField `json:"fields" validate:"required,dive"`
}
