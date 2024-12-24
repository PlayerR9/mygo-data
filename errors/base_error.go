package errors

// This is done to avoid dependency on the "errors" package.

// baseError is the base error type.
type baseError struct {
	// msg is the error message.
	msg string
}

// Error implements error.
func (be baseError) Error() string {
	return be.msg
}

// New returns a new error with the given message.
//
// Parameters:
//   - msg: The error message. If empty, it defaults to DefaultErrorMessage.
//
// Returns:
//   - error: The new error. Never returns nil.
func New(msg string) error {
	if msg == "" {
		msg = DefaultErrorMessage
	}

	be := &baseError{
		msg: msg,
	}

	return be
}
