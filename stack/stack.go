package stack

import "github.com/PlayerR9/mygo-data/common"

// CoreStack is a generic stack interface.
type CoreStack[E any] interface {
	// Push pushes an element onto the stack.
	//
	// Parameters:
	//   - e: The element to be pushed onto the stack.
	//
	// Returns:
	//   - error: An error if the element could not be pushed onto the stack.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	Push(e E) error

	// Pop pops an element from the stack.
	//
	// Returns:
	//   - E: The element that was popped from the stack.
	//   - error: An error if the element could not be popped from the stack.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - ErrEmptyStack: If the stack is empty.
	Pop() (E, error)

	// IsEmpty checks if the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	IsEmpty() bool
}

// Stack is an interface that extends CoreStack and Collection.
type Stack[E any] interface {
	CoreStack[E]
	Collection[E]
	common.Resetter
}
