package sets

import (
	"fmt"

	"github.com/PlayerR9/mygo-data/common"
)

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

	common.Collection[E]

	fmt.Stringer
}
