package models

import (
	"snappy-api/core/logger"

	jsoniter "github.com/json-iterator/go"
)

// Data s
type Data = interface{}

// ResponseMessage asd
type ResponseMessage struct {
	Message     string
	IsSucceeded bool
	Data        Data
}

// NewResponse s
func NewResponse(suc bool, d Data, mes string) ResponseMessage {
	return ResponseMessage{mes, suc, d}
}

// MustMarshal s
func (r *ResponseMessage) MustMarshal() []byte {
	logger := logger.GetLogInstance("", "")

	j, err := jsoniter.Marshal(r)

	if err != nil {
		logger.Error(err)
	}
	return j
}
