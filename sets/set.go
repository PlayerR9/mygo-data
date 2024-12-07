package sets

import (
	"fmt"

	"github.com/PlayerR9/mygo-data/common"
)

// BasicSet is the interface for a basic set.
type BasicSet[E any] interface {
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
}

// Insert adds multiple elements to the given BasicSet.
//
// Parameters:
//   - s: The BasicSet to which elements will be added.
//   - elems: The elements to be inserted into the set.
//
// Returns:
//   - uint: The number of elements successfully inserted into the set.
//   - error: An error if insertion fails at any point.
//
// Errors:
//   - common.ErrBadParam: If the BasicSet is nil.
//   - any other error: Implementation-specific.
func Insert[E any](s BasicSet[E], elems ...E) (uint, error) {
	if s == nil {
		return 0, common.NewErrNilParam("s")
	}

	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	}

	for i, elem := range elems {
		err := s.Insert(elem)
		if err != nil {
			return uint(i), err
		}
	}

	return lenElems, nil
}

// Set is a set of elements.
type Set[E any] interface {
	BasicSet[E]

	common.Collection[E]

	fmt.Stringer
}
