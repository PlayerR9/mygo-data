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

	// Add adds an element to the collection.
	//
	// Parameters:
	//   - elem: The element to add to the collection.
	//
	// Returns:
	//   - error: An error if the addition fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - ErrFullCollection: If the collection is full.
	//   - any other error: Implementation-specific.
	Add(elem E) error
}

// Add adds multiple elements to the given collection.
//
// Parameters:
//   - c: The collection to which elements will be added.
//   - elems: The elements to be added to the collection.
//
// Returns:
//   - uint: The number of elements successfully added to the collection.
//   - error: An error if addition fails at any point.
//
// Errors:
//   - common.ErrBadParam: If the collection is nil.
//   - ErrFullCollection: If the collection is full.
//   - any other error: Implementation-specific.
func Add[E any](c Collection[E], elems ...E) (uint, error) {
	if c == nil {
		return 0, NewErrNilParam("c")
	}

	lenElems := len(elems)
	if lenElems == 0 {
		return 0, nil
	}

	for i, elem := range elems {
		err := c.Add(elem)
		if err != nil {
			return uint(i), err
		}
	}

	return uint(lenElems), nil
}
