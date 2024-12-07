package collections

import (
	"fmt"
	"strings"

	"github.com/PlayerR9/mygo-data/common"
)

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
		return 0, common.NewErrNilParam("c")
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

// func IsEmpty(collection any) (bool, error) {
// 	if collection == nil {
// 		return false, common.NewErrNilParam("collection")
// 	}

// 	switch c := collection.(type) {
// 	case interface{ IsEmpty() bool }:
// 		ok := c.IsEmpty()
// 		return ok, nil
// 	case interface{ Size() uint }:
// 		size := c.Size()
// 		return size == 0, nil
// 	default:
// 		return false, fmt.Errorf("collection type %T does not implement IsEmpty() or Size()", collection)
// 	}
// }

// String returns a string representation of the given collection.
//
// Parameters:
//   - type_: The name of the type of elements in the collection.
//   - collection: The collection for which to generate the string.
//
// Returns:
//   - string: The string representation of the collection.
//
// Notes:
//   - If the collection is empty, the string representation is the type name
//     followed by "[]".
//   - If the collection is not empty, the string representation is the type name
//     followed by "[" and the elements of the collection (separated by ", ") and
//     "]".
func String[E any](type_ string, collection Collection[E]) string {
	if collection == nil {
		return "<nil>"
	}

	slice := collection.Slice()
	if len(slice) == 0 {
		return type_ + "[]"
	}

	elems := make([]string, 0, len(slice))
	for _, e := range slice {
		elems = append(elems, fmt.Sprint(e))
	}

	return type_ + "[" + strings.Join(elems, ", ") + "]"
}
