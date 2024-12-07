package common

// Resetter is an interface for objects that can be reset.
type Resetter interface {
	// Reset resets the object, allowing it to be used again.
	//
	// Returns:
	//   - error: An error if the reset fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error: Implementation-specific.
	Reset() error
}
