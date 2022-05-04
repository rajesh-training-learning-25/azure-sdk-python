// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import "fmt"

// Code is an error code, usable by consuming code to work with
// programatically.
type Code string

const (
	// CodeConnectionLost means our connection was lost and all retry attempts failed.
	// This typically reflects an extended outage or connection disruption and may
	// require manual intervention.
	CodeConnectionLost Code = "connlost"

	// CodeLockLost means that the lock token you have for a message has expired.
	// This message will be available again after the lock period expires, or, potentially
	// go to the dead letter queue if delivery attempts have been exceeded.
	CodeLockLost Code = "locklost"
)

// Error represents a Service Bus specific error.
type Error struct {
	// Code is a stable error code which can be used as part of programatic error handling.
	// The codes can expand in the future, but the values (and their meaning) will remain the same.
	Code     Code
	innerErr error
}

// Error is an error message containing the code and a user friendly message, if any.
func (e Error) Error() string {
	return fmt.Sprintf("(%s): %s", e.Code, e.innerErr.Error())
}

// Unwrap is implemented so this error can be used with errors.As()
func (e Error) Unwrap() error {
	return e.innerErr
}

// NewError creates a new `Error` instance.
// NOTE: this function is only exported so it can be used by the `internal`
// package. It is not available for customers.
func NewError(code Code, innerErr error) error {
	return Error{
		Code:     code,
		innerErr: innerErr,
	}
}
