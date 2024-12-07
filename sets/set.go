package sets

import "fmt"

// Set is a set of elements.
type Set[E any] interface {
	// Has checks if the given element is in the set.
	//
	// Parameters:
	//   - e: The element to check.
	//
	// Returns:
	//   - bool: True if the element is in the set, false otherwise.
	//
	// When the receiver is nil, the function returns false.
	Has(e E) bool

	// Insert adds the given element to the set.
	//
	// Parameters:
	//   - e: The element to insert.
	//
	// Returns:
	//   - error: An error if the insertion fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error: Implementation-specific.
	Insert(e E) error

	// IsEmpty checks if the set is empty.
	//
	// Returns:
	//   - bool: True if the set is empty, false otherwise.
	IsEmpty() bool

	// Size returns the number of elements in the set.
	//
	// Returns:
	//   - uint: The number of elements in the set.
	Size() uint

	fmt.Stringer
}
