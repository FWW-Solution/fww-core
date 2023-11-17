package tools

import (
	"fww-core/internal/data/dto"
)

func ResponseBuilder(data interface{}, meta interface{}) dto.BaseResponse {
	return dto.BaseResponse{
		Meta: meta,
		Data: data,
	}
}

// ResponseBadRequest is a function to build response bad request.
func ResponseBadRequest(errMessage error) dto.BaseResponse {
	meta := dto.MetaResponse{
		StatusCode: "ERR400",
		IsSuccess:  false,
		Message:    errMessage.Error(),
	}

	return ResponseBuilder(nil, meta)
}

func ResponseInternalServerError(err error) dto.BaseResponse {
	return dto.BaseResponse{
		Meta: dto.MetaResponse{
			StatusCode: "ERR500",
			IsSuccess:  false,
			Message:    err.Error(),
		},
		Data: nil,
	}

}
