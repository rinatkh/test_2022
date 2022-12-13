package constants

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type CodedError struct {
	err  error
	code int
}

func (ce *CodedError) Error() string {
	return ce.err.Error()
}

func (ce *CodedError) Code() int {
	return ce.code
}

func NewCodedError(errMessage string, code int) *CodedError {
	return &CodedError{errors.New(errMessage), code}
}

var (
	// Unathorized
	InputError      = &CodedError{errors.New("bad json request"), fiber.StatusBadRequest}
	ErrDBNotFound   = &CodedError{errors.New("not found in the database"), fiber.StatusBadRequest}
	AuthError       = &CodedError{errors.New("Invalid public api key"), fiber.StatusUnauthorized}
	ErrGenerateUUID = &CodedError{errors.New("failed to generate UUID"), fiber.StatusInternalServerError}
)
