package stacks

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// ArrayStack is a stack implemented using a slice. This is thread-safe.
//
// An empty ArrayStack can be created by either using the `new(ArrayStack[T])`
// constructor.
type ArrayStack[T any] struct {
	// elems is a stack of elements.
	elems []T

	// mu is a mutex for the stack.
	mu sync.RWMutex
}

// Push implements Stack.
func (a *ArrayStack[T]) Push(elem T) error {
	if a == nil {
		return common.ErrNilReceiver
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.elems = append(a.elems, elem)

	return nil
}

// Pop implements Stack.
func (a *ArrayStack[T]) Pop() (T, error) {
	if a == nil {
		return *new(T), common.ErrNilReceiver
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.elems) == 0 {
		return *new(T), ErrEmptyStack
	}

	top := a.elems[len(a.elems)-1]
	a.elems = a.elems[:len(a.elems)-1]

	return top, nil
}

// IsEmpty implements Stack.
func (a *ArrayStack[T]) IsEmpty() bool {
	if a == nil {
		return true
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	return len(a.elems) == 0
}

// Peek implements Stack.
func (a *ArrayStack[T]) Peek() (T, error) {
	if a == nil {
		return *new(T), common.ErrNilReceiver
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	if len(a.elems) == 0 {
		return *new(T), ErrEmptyStack
	}

	return a.elems[len(a.elems)-1], nil
}

// Slice implements Stack.
func (a *ArrayStack[T]) Slice() []T {
	if a == nil {
		return nil
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	elems := make([]T, len(a.elems))
	copy(elems, a.elems)

	return elems
}

// Reset implements Stack.
func (a *ArrayStack[T]) Reset() error {
	if a == nil {
		return common.ErrNilReceiver
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.elems) == 0 {
		return nil
	}

	clear(a.elems)
	a.elems = nil

	return nil
}

// Size implements Stack.
func (a *ArrayStack[T]) Size() uint {
	if a == nil {
		return 0
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	return uint(len(a.elems))
}
