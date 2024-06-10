package utils

import (
	"github.com/merge-hotel-data/errors"
	"github.com/merge-hotel-data/model"
)

func FormErrorMessage(err errors.ErrorInterface) model.ErrorResponse {
	return model.ErrorResponse{
		Code:         -1,
		ErrorCode:    err.Code(),
		ErrorMessage: err.Error(),
	}
}
