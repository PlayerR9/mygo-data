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
	Push(elem T) error

	// Pop removes and returns the top element from the stack.
	//
	// Returns:
	//   - T: The top element from the stack.
	//   - error: An error if the stack operation fails.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - ErrStackEmpty: If the stack is empty.
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
	//   - ErrStackEmpty: If the stack is empty.
	Peek() (T, error)

	// Slice returns the elements of the stack as a slice.
	//
	// Returns:
	//   - []T: The elements of the stack as a slice.
	//
	// The returned slice has the property that, the last element of the
	// slice is the bottom element of the stack.
	Slice() []T

	// Reset empties the stack for reuse.
	Reset()

	// Size returns the number of elements in the stack.
	//
	// Returns:
	//   - uint: The number of elements in the stack.
	Size() uint
}

// Push pushes multiple elements onto the stack.
//
// Parameters:
//   - stack: The stack to push onto.
//   - elems: The elements to push onto the stack.
//
// Returns:
//   - error: An error if the stack operation fails.
//
// Errors:
//   - common.ErrBadParam: If the stack is nil.
func Push[T any](stack Stack[T], elems ...T) error {
	if stack == nil {
		return common.NewErrNilParam("stack")
	} else if len(elems) == 0 {
		return nil
	}

	elems = Reverse(elems)

	for _, elem := range elems {
		_ = stack.Push(elem)
	}

	return nil
}
