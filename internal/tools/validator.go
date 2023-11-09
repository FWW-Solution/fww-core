package tools

import (
	"fww-core/internal/data/dto"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func validateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateVariable(data interface{}) *dto.BaseResponse {
	errValidator := validateStruct(data)
	if errValidator != nil {
		response := dto.BaseResponse{
			Meta: dto.MetaResponse{
				StatusCode: "ERR400",
				IsSuccess:  false,
				Message:    errValidator[0].FailedField,
			},
			Data: nil,
		}

		return &response
	}
	return nil
}
