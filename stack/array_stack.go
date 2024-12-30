package stack

import (
	"sync"

	common "github.com/PlayerR9/mygo-data/common"
	"github.com/PlayerR9/mygo-data/stack/internal"
)

// ArrayStack is a generic stack implemented using an array.
//
// An empty stack can be created with the `as := new(ArrayStack[E])`
// constructor.
type ArrayStack[E any] struct {
	// elems is the underlying array.
	elems []E

	// mu is the mutex for the stack.
	mu sync.RWMutex
}

// Push implements CoreStack.
func (as *ArrayStack[E]) Push(e E) error {
	if as == nil {
		return common.ErrNilReceiver
	}

	as.mu.Lock()
	defer as.mu.Unlock()

	as.elems = append(as.elems, e)

	return nil
}

// Pop implements CoreStack.
func (as *ArrayStack[E]) Pop() (E, error) {
	if as == nil {
		return *new(E), common.ErrNilReceiver
	}

	as.mu.Lock()
	defer as.mu.Unlock()

	if len(as.elems) == 0 {
		return *new(E), ErrEmptyStack
	}

	e := as.elems[len(as.elems)-1]
	as.elems = as.elems[:len(as.elems)-1]

	return e, nil
}

// IsEmpty implements CoreStack.
func (as *ArrayStack[E]) IsEmpty() bool {
	if as == nil {
		return true
	}

	as.mu.RLock()
	defer as.mu.RUnlock()

	ok := len(as.elems) == 0
	return ok
}

// Slice implements Collection.
func (as *ArrayStack[E]) Slice() []E {
	if as == nil {
		return nil
	}

	as.mu.RLocker()
	defer as.mu.RUnlock()

	if len(as.elems) == 0 {
		return nil
	}

	slice := make([]E, 0, len(as.elems))
	copy(slice, as.elems)

	internal.Reverse(slice)

	return slice
}

// Reset implements Collection.
func (as *ArrayStack[E]) Reset() error {
	if as == nil {
		return common.ErrNilReceiver
	}

	as.mu.Lock()
	defer as.mu.Unlock()

	if len(as.elems) == 0 {
		return nil
	}

	clear(as.elems)
	as.elems = nil

	return nil
}

// PushMany pushes all elements in the slice onto the stack in the order they are given in the slice.
//
// Parameters:
//   - elems: The elements to push onto the stack.
//
// Returns:
//   - error: An error if the elements could not be pushed onto the stack.
//
// Errors:
//   - common.ErrNilReceiver: If the stack is nil.
func (as *ArrayStack[E]) PushMany(elems []E) error {
	if as == nil {
		return common.ErrNilReceiver
	}

	as.mu.Lock()
	defer as.mu.Unlock()

	if len(elems) == 0 {
		return nil
	}

	tmp := make([]E, len(elems)) // Prevent side-effects
	copy(tmp, elems)

	internal.Reverse(tmp)

	as.elems = append(as.elems, tmp...)

	return nil
}
