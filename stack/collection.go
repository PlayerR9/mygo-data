package stack

// Collection is an interface that represents a collection of elements.
type Collection[E any] interface {
	// Slice returns a slice of the elements in the collection.
	//
	// Returns:
	//   - []E: A slice of the elements in the collection.
	Slice() []E

	// Reset resets the object.
	//
	// Returns:
	//   - error: An error if the object could not be reset.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error: Implementation-specific.
	Reset() error
}
