package common

// Collection is an interface for a collection.
type Collection[E any] interface {
	// IsEmpty checks whether the collection is empty.
	//
	// Returns:
	//   - bool: True if the collection is empty, false otherwise.
	//
	// If the receiver is nil, the function returns true.
	IsEmpty() bool

	// Size returns the number of elements in the collection.
	//
	// Returns:
	//   - uint: The number of elements in the collection.
	//
	// If the receiver is nil, the function returns 0.
	Size() uint

	// Reset resets the object, allowing it to be used again.
	//
	// Returns:
	//   - error: An error if the reset fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error: Implementation-specific.
	Reset() error

	// Slice returns the elements of the collection as a slice.
	//
	// Returns:
	//   - []E: The elements of the collection as a slice.
	//
	// Behaviors:
	//   - If the receiver is nil, the function returns nil.
	//   - The order of the elements is implementation-specific.
	Slice() []E
}
