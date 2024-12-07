package stacks

import (
	"slices"
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// RefusableStack is a stack that can be refused. This is thread-safe.
type RefusableStack[T any] struct {
	// stack is the underlying stack.
	stack Stack[T]

	// popped is the stack of elements that have been popped.
	popped []T

	// mu is the mutex for the RefusableStack.
	mu sync.RWMutex
}

// Push implements BasicStack.
func (r *RefusableStack[T]) Push(elem T) error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.stack.Push(elem)
	return err
}

// Pop implements BasicStack.
func (r *RefusableStack[T]) Pop() (T, error) {
	if r == nil {
		return *new(T), common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	top, err := r.stack.Pop()
	if err != nil {
		return *new(T), err
	}

	r.popped = append(r.popped, top)

	return top, nil
}

// Peek implements BasicStack.
func (r *RefusableStack[T]) Peek() (T, error) {
	if r == nil {
		return *new(T), common.ErrNilReceiver
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	top, err := r.stack.Peek()
	if err != nil {
		return *new(T), err
	}

	return top, nil
}

// IsEmpty implements common.Collection.
func (r *RefusableStack[T]) IsEmpty() bool {
	if r == nil {
		return true
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	ok := r.stack.IsEmpty()
	return ok
}

// Slice implements common.Collection.
func (r *RefusableStack[T]) Slice() []T {
	if r == nil {
		return nil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	slice := r.stack.Slice()
	return slice
}

// Reset implements common.Collection.
func (r *RefusableStack[T]) Reset() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.stack.Reset()
	if err != nil {
		return err
	}

	clear(r.popped)
	r.popped = nil

	return nil
}

// Size implements common.Collection.
func (r *RefusableStack[T]) Size() uint {
	if r == nil {
		return 0
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	size := r.stack.Size()
	return size
}

// RefusableOf creates a RefusableStack from the given Stack.
//
// Parameters:
//   - stack: The stack to create a RefusableStack from.
//
// Returns:
//   - *RefusableStack[T]: The created RefusableStack. Never returns nil.
//
// If the given stack is nil, a new ArrayStack will be created.
func RefusableOf[T any](stack Stack[T]) *RefusableStack[T] {
	if stack == nil {
		stack = new(ArrayStack[T])
	}

	return &RefusableStack[T]{
		stack:  stack,
		popped: nil,
	}
}

// Popped returns a slice containing the elements that have been popped from the stack
// but not yet accepted or refused.
//
// Returns:
//   - []T: A slice of the popped elements.
//
// If the receiver is nil, the function returns nil.
func (r *RefusableStack[T]) Popped() []T {
	if r == nil {
		return nil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	popped := make([]T, 0, len(r.popped))
	copy(popped, r.popped)

	return popped
}

// Accept clears the popped elements from the stack, effectively accepting them.
//
// Returns:
//   - error: An error if the stack operation fails.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (r *RefusableStack[T]) Accept() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.popped) == 0 {
		return nil
	}

	clear(r.popped)
	r.popped = nil

	return nil
}

// Refuse refuses the popped elements, putting them back onto the stack.
//
// Returns:
//   - error: An error if the stack operation fails.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - ErrFullStack: If the stack is full.
//   - any other error: Implementation-specific.
func (r *RefusableStack[T]) Refuse() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.popped) == 0 {
		return nil
	}

	slices.Reverse(r.popped)

	for _, elem := range r.popped {
		err := r.stack.Push(elem)
		if err != nil {
			return err
		}
	}

	clear(r.popped)
	r.popped = nil

	return nil
}

// Add implements common.Collection.
func (r *RefusableStack[T]) Add(elem T) error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.stack.Add(elem)
	return err
}
