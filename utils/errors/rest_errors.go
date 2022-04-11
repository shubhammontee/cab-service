package errors

import (
	"errors"
	"net/http"
)

type (
	//RestErr ...
	RestErr struct {
		Message string
		Status  int
		Error   string
	}
)

//NewError ...
func NewError(msg string) error {
	return errors.New(msg)
}

//NewBadRequestError ...
func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

//NewInternalServerError ...
func NewInternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}
