package helper

import "github.com/go-playground/validator/v10"

type ErrorStruct struct {
	Code    int
	Message error
}

var Validate = validator.New()
