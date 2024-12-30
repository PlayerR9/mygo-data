package stack

import common "github.com/PlayerR9/mygo-data/common"

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

// Push pushes all elements in the slice onto the stack in the order they are given in the slice.
//
// Parameters:
//   - stack: The stack to push the elements onto.
//   - elems: The elements to push onto the stack.
//
// Returns:
//   - error: An error if the stack is nil, or if any of the elements could not be pushed onto the stack.
//
// Errors:
//   - common.ErrNilParam: If the stack is nil.
//   - any other error: Implementation-specific.
func Push[E any](stack CoreStack[E], elems []E) error {
	if stack == nil {
		return common.NewErrNilParam("stack")
	} else if len(elems) == 0 {
		return nil
	}

	if v, ok := stack.(interface{ PushMany([]E) error }); ok {
		return v.PushMany(elems)
	}

	for i := len(elems) - 1; i >= 0; i-- {
		err := stack.Push(elems[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// Stack is an interface that extends CoreStack and Collection.
type Stack[E any] interface {
	CoreStack[E]
	Collection[E]
}
