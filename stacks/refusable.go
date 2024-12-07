package stacks

import (
	"slices"
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// RefusableStack is a stack that can be refused. This is thread-safe.
//
// A RefusableStack can be created by using the `stack := new(RefusableStack[T])`
// constructor.
type RefusableStack[T any] struct {
	// stack is the underlying stack.
	stack Stack[T]

	// popped is the stack of elements that have been popped.
	popped []T

	// mu is the mutex for the RefusableStack.
	mu sync.RWMutex
}

// Push implements Stack.
func (r *RefusableStack[T]) Push(elem T) error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	_ = r.stack.Push(elem)

	return nil
}

// Pop implements Stack.
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

// IsEmpty implements Stack.
func (r *RefusableStack[T]) IsEmpty() bool {
	if r == nil {
		return true
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.stack.IsEmpty()
}

// Peek implements Stack.
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

// Slice implements Stack.
func (r *RefusableStack[T]) Slice() []T {
	if r == nil {
		return nil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.stack.Slice()
}

// Reset implements Stack.
func (r *RefusableStack[T]) Reset() {
	if r == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.stack.Reset()

	clear(r.popped)
	r.popped = nil
}

// Size implements Stack.
func (r *RefusableStack[T]) Size() uint {
	if r == nil {
		return 0
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.stack.Size()
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
//
// This method is thread-safe.
func (r *RefusableStack[T]) Refuse() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	slices.Reverse(r.popped)

	for _, elem := range r.popped {
		_ = r.stack.Push(elem)
	}

	clear(r.popped)
	r.popped = nil

	return nil
}
