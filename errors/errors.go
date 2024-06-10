package errors

import (
	"github.com/merge-hotel-data/constants"
)

type ErrorInterface interface {
	Error() string
	Code() uint32
}

type baseError struct {
	code    uint32
	message string
}

func (err baseError) Error() string {
	return err.message
}

func (err baseError) Code() uint32 {
	return err.code
}

func New(code uint32, message string) ErrorInterface {
	return &baseError{
		code:    code,
		message: message,
	}
}

func InternalServerError() ErrorInterface {
	return New(constants.InternalServerErrorCode, constants.InternalServerError)
}

func LoadingConfigurationFileError() ErrorInterface {
	return New(constants.LoadingConfigurationFileErrorCode, constants.LoadingConfigurationFileError)
}
