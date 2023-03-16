package types

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	Error() string
	RefineError(err ...interface{}) Error
}

type BizError struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func NewBusinessError(code uint32, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

func (e *BizError) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (e *BizError) RefineError(err ...interface{}) Error {
	return NewBusinessError(e.Code, e.Message+fmt.Sprint(err...))
}
