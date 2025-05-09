package errors

import (
	"fmt"
	"net/http"
)

// ErrorCode represents application-specific error codes
type ErrorCode string

// Error codes used throughout the application
const (
	// General errors
	ErrInvalid       ErrorCode = "INVALID_INPUT"
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrInternal      ErrorCode = "INTERNAL_ERROR"
	ErrUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrForbidden     ErrorCode = "FORBIDDEN"
	ErrAlreadyExists ErrorCode = "ALREADY_EXISTS"
	ErrTimeout       ErrorCode = "TIMEOUT"

	// Business logic errors
	ErrInvalidTransaction ErrorCode = "INVALID_TRANSACTION"
	ErrInsufficientFunds  ErrorCode = "INSUFFICIENT_FUNDS"
	ErrInvalidStatus      ErrorCode = "INVALID_STATUS"
)

// AppError is a structured error for consistent API error responses
// It implements the error interface with additional context
type AppError struct {
	// Original error that caused this error
	Err error
	// HTTP status code to return
	HTTPCode int
	// Error code for client application to handle programmatically
	Code ErrorCode
	// Human-readable message
	Message string
	// Additional details for debugging (not exposed to client)
	Details map[string]interface{}
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		HTTPCode: codeToHTTP(code),
		Code:     code,
		Message:  message,
	}
}

// Wrap wraps an existing error in an AppError
func Wrap(err error, code ErrorCode, message string) *AppError {
	return &AppError{
		Err:      err,
		HTTPCode: codeToHTTP(code),
		Code:     code,
		Message:  message,
	}
}

// WithDetails adds context information to the error
func (e *AppError) WithDetails(details map[string]interface{}) *AppError {
	e.Details = details
	return e
}

// Helper function to convert error codes to HTTP status codes
func codeToHTTP(code ErrorCode) int {
	switch code {
	case ErrInvalid:
		return http.StatusBadRequest
	case ErrNotFound:
		return http.StatusNotFound
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrForbidden:
		return http.StatusForbidden
	case ErrAlreadyExists:
		return http.StatusConflict
	case ErrTimeout:
		return http.StatusRequestTimeout
	case ErrInvalidTransaction, ErrInsufficientFunds, ErrInvalidStatus:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// Convenience error creation functions

// NewNotFound creates a not found error
func NewNotFound(resource string, id interface{}) *AppError {
	return New(ErrNotFound, fmt.Sprintf("%s with ID '%v' not found", resource, id))
}

// NewInvalidInput creates an invalid input error
func NewInvalidInput(message string) *AppError {
	return New(ErrInvalid, message)
}

// NewInternalError creates an internal server error
func NewInternalError(err error) *AppError {
	return Wrap(err, ErrInternal, "Internal server error")
}

// NewUnauthorized creates an unauthorized error
func NewUnauthorized(message string) *AppError {
	if message == "" {
		message = "Unauthorized access"
	}
	return New(ErrUnauthorized, message)
}

// NewForbidden creates a forbidden error
func NewForbidden(message string) *AppError {
	if message == "" {
		message = "Access forbidden"
	}
	return New(ErrForbidden, message)
}

// NewAlreadyExists creates an already exists error
func NewAlreadyExists(resource string, id interface{}) *AppError {
	return New(ErrAlreadyExists, fmt.Sprintf("%s with ID '%v' already exists", resource, id))
}