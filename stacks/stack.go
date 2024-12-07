package stacks

import "github.com/PlayerR9/mygo-data/common"

// Stack is a generic stack interface.
type Stack[T any] interface {
	// Push pushes an element onto the stack.
	//
	// Parameters:
	//   - elem: The element to push onto the stack.
	//
	// Returns:
	//   - error: An error if the stack operation fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - ErrFullStack: If the stack is full.
	//   - any other error: Implementation-specific.
	Push(elem T) error

	// Pop removes and returns the top element from the stack.
	//
	// Returns:
	//   - T: The top element from the stack.
	//   - error: An error if the stack operation fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - ErrEmptyStack: If the stack is empty.
	//   - any other error: Implementation-specific.
	Pop() (T, error)

	// IsEmpty checks whether the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	IsEmpty() bool

	// Peek returns the top element from the stack without removing it.
	//
	// Returns:
	//   - T: The top element from the stack.
	//   - error: An error if the stack operation fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - ErrEmptyStack: If the stack is empty.
	//   - any other error: Implementation-specific.
	Peek() (T, error)

	// Slice returns the elements of the stack as a slice.
	//
	// Returns:
	//   - []T: The elements of the stack as a slice.
	//
	// The returned slice has the property that, the last element of the
	// slice is the bottom element of the stack.
	Slice() []T

	// Size returns the number of elements in the stack.
	//
	// Returns:
	//   - uint: The number of elements in the stack.
	Size() uint

	common.Resetter
}

// Push pushes multiple elements onto the stack.
//
// Parameters:
//   - stack: The stack to push onto.
//   - elems: The elements to push onto the stack.
//
// Returns:
//   - uint: The number of elements pushed onto the stack.
//   - error: An error if the stack operation fails.
//
// Errors:
//   - common.ErrBadParam: If the stack is nil.
//   - ErrFullStack: If the stack is full.
//   - any other error: Implementation-specific.
func Push[T any](stack Stack[T], elems ...T) (uint, error) {
	if stack == nil {
		return 0, common.NewErrNilParam("stack")
	}

	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	}

	elems = Reverse(elems)

	for i, elem := range elems {
		err := stack.Push(elem)
		if err != nil {
			return uint(i), err
		}
	}

	return lenElems, nil
}
