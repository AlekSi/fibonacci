package errs

import (
	"fmt"
)

// APIErrorCode is a type for API error codes.
// It exists to facilitate usage of constants defined in this package instead of magic numbers.
type APIErrorCode int

// Defined error codes.
const (
	InvalidParameter APIErrorCode = 400001
)

// APIError is API error returned to clients.
type APIError struct {
	Code    APIErrorCode `json:"code"`            // Code like 400001, 403001, etc. See constants.
	Message string       `json:"message"`         // Message for clients.
	Err     error        `json:"error,omitempty"` // Original error, exposed in debug mode.
}

// check interface
var _ error = &APIError{}

// New creates new error.
func New(code APIErrorCode, message string, err error) *APIError {
	return &APIError{Code: code, Message: message, Err: err}
}

// Error implements error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("%d: %s (%v)", e.Code, e.Message, e.Err)
}

// Cause returns underlying error. Implements github.com/pkg/errors.causer
func (e *APIError) Cause() error {
	return e.Err
}
