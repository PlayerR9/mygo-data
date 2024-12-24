package stack

import (
	"github.com/PlayerR9/mygo-data/errors"
)

// RefusableStack is a stack that can be reset.
type RefusableStack[E any] struct {
	// stack is the stack that is being reset.
	stack Stack[E]

	// popped is the stack that is being reset.
	popped []E
}

// Push implements CoreStack.
func (s *RefusableStack[E]) Push(e E) error {
	if s == nil {
		return errors.ErrNilReceiver
	}

	err := s.stack.Push(e)
	return err
}

// Pop implements CoreStack.
func (s *RefusableStack[E]) Pop() (E, error) {
	if s == nil {
		return *new(E), errors.ErrNilReceiver
	}

	top, err := s.stack.Pop()
	if err != nil {
		return *new(E), err
	}

	s.popped = append(s.popped, top)

	return top, nil
}

// Slice implements Collection.
func (s RefusableStack[E]) Slice() []E {
	elems := s.stack.Slice()
	return elems
}

// Reset implements common.Resetter.
func (s *RefusableStack[E]) Reset() error {
	if s == nil {
		return errors.ErrNilReceiver
	}

	err := s.stack.Reset()
	if err != nil {
		return err
	}

	if len(s.popped) == 0 {
		return nil
	}

	clear(s.popped)
	s.popped = nil

	return nil
}

// RefusableOf creates a new RefusableStack from the provided stack.
//
// Parameters:
//   - stack: The stack to be wrapped in a RefusableStack.
//
// Returns:
//   - *RefusableStack[E]: A pointer to the newly created RefusableStack.
//   - error: An error if the provided stack is nil.
//
// Errors:
//   - common.ErrBadParam: If the stack parameter is nil.
func RefusableOf[E any](stack Stack[E]) (*RefusableStack[E], error) {
	if stack == nil {
		err := errors.NewErrNilParam("stack")
		return nil, err
	}

	s := &RefusableStack[E]{
		stack:  stack,
		popped: nil,
	}

	return s, nil
}

// Accept resets the popped stack, effectively "accepting" the popped elements.
//
// Returns:
//   - error: An error if the receiver is nil or if the popped stack could not be reset.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (s *RefusableStack[E]) Accept() error {
	if s == nil {
		return errors.ErrNilReceiver
	}

	if len(s.popped) == 0 {
		return nil
	}

	clear(s.popped)
	s.popped = nil

	return nil
}

// Refuse transfers all elements from the stack to the popped stack,
// effectively "refusing" the stack.
//
// Returns:
//   - error: An error if the receiver is nil.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (s *RefusableStack[E]) Refuse() error {
	if s == nil {
		return errors.ErrNilReceiver
	}

	for len(s.popped) > 0 {
		top := s.popped[len(s.popped)-1]
		s.popped = s.popped[:len(s.popped)-1]

		err := s.stack.Push(top)
		if err != nil {
			break
		}
	}

	return nil
}

// Popped returns the elements that were popped from the stack.
//
// Returns:
//   - []E: The elements that were popped from the stack.
func (s RefusableStack[E]) Popped() []E {
	if len(s.popped) == 0 {
		return nil
	}

	slice := make([]E, len(s.popped))
	copy(slice, s.popped)

	j := len(slice) - 1

	for i := 0; i < j; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}

	return slice
}
