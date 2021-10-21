package fluid

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-uniform/uniform"
	"strconv"
	"strings"
)

func (p Project) Validate() validator.ValidationErrors {
	v := validator.New()

	if err := v.RegisterValidation("portalType", func(fl validator.FieldLevel) bool {
		return uniform.Contains([]string{
			PortalTypeIonic,
			PortalTypeVuetify,
		}, fl.Field().String(), true)
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("entityFieldType", func(fl validator.FieldLevel) bool {
		return uniform.Contains([]string{
			EntityFieldTypeString,
			EntityFieldTypeInteger,
			EntityFieldTypeDecimal,
			EntityFieldTypeDateTime,
			EntityFieldTypeDate,
			EntityFieldTypeTime,
			EntityFieldTypeBoolean,
			EntityFieldTypeMoney,
			EntityFieldTypeUuid,
			EntityFieldTypeAttribute,
			EntityFieldTypeBinary,
			EntityFieldTypePassword,
		}, fl.Field().String(), true)
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("entityActionMethod", func(fl validator.FieldLevel) bool {
		return uniform.Contains([]string{
			EntityActionMethodGet,
			EntityActionMethodPost,
			EntityActionMethodPut,
			EntityActionMethodDelete,
		}, fl.Field().String(), true)
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("entityActionType", func(fl validator.FieldLevel) bool {
		return uniform.Contains([]string{
			EntityActionTypeList,
			EntityActionTypeRecord,
		}, fl.Field().String(), true)
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("contractType", func(fl validator.FieldLevel) bool {
		return uniform.Contains([]string{
			ContractTypeRequest,
			ContractTypeResponse,
			ContractTypeParameters,
		}, fl.Field().String(), true)
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("contractFieldType", func(fl validator.FieldLevel) bool {
		return uniform.Contains([]string{
			ContractFieldTypeString,
			ContractFieldTypeInteger,
			ContractFieldTypeDecimal,
			ContractFieldTypeDateTime,
			ContractFieldTypeDate,
			ContractFieldTypeTime,
			ContractFieldTypeBoolean,
			ContractFieldTypeMoney,
			ContractFieldTypeUuid,
			ContractFieldTypeBinary,
			ContractFieldTypeFile,
		}, fl.Field().String(), true)
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("color", func(fl validator.FieldLevel) bool {
		hexValue := strings.TrimPrefix(fl.Field().String(), "#")

		if len(hexValue) == 0 {
			return true
		}

		if len(hexValue) == 3 {
			hexValue = fmt.Sprintf(
				"%s%s%s%s%s%s",
				hexValue[0:1],
				hexValue[0:1],
				hexValue[1:2],
				hexValue[1:2],
				hexValue[2:3],
				hexValue[2:3],
			)
		}

		if len(hexValue) != 6 {
			return false
		}

		if _, err := strconv.ParseInt(hexValue, 16, 32); err != nil {
			return false
		}

		return true
	}); err != nil {
		panic(err)
	}

	if err := v.Struct(p); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return validationErrors
		}
	}

	return nil
}
