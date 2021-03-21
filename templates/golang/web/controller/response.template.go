package controller

// response constants
const (
	Response = `
	package controllers

import (
	"context"
)

// APIResponse : generic api response
type APIResponse struct {
	Code      APIResponseCode ` + "`" + `json:"code"` + "`" + `
	Data      interface{}     ` + "`" + `json:"data"` + "`" + `
	RequestID string          ` + "`" + `json:"req_id"` + "`" + `
}

// NewAPIResponse :
func NewAPIResponse(ctx context.Context, code APIResponseCode, message interface{}) APIResponse {
	return APIResponse{
		Code:      code,
		Data:      message,
		RequestID: filter.GetRequestID(ctx),
	}
}

// APIResponseCode :
type APIResponseCode int32

// Success error codes
const (
	StatusOK APIResponseCode = iota + 20000
)

// Client error codes
const (
	ErrInvalidRequest APIResponseCode = iota + 40000
	ErrInvalidAccountActivationRequest
	ErrInvalidOrExpiredAccountActivationRequest
	ErrUnauthorized
	ErrEntityDoesNotExist
	ErrEntityAlreadyExist
	ErrUserRegistrationIncomplete
)

// Server error codes
const (
	ErrInternalServerError APIResponseCode = iota + 50000
	ErrAccountCreatedEmailNotSent
	ErrAccountActivationCompleteEmailNotSent
	ErrResendActivationEmailNotSent
)
`
)
