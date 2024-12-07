package common

import (
	"sync"
)

// ArrayCollection is an array collection.
//
// An empty array collection can be created by using the `new(ArrayCollection[T])`
// constructor.
type ArrayCollection[E any] struct {
	// elems is a slice of elements.
	elems []E

	// mu is a read-write mutex.
	mu sync.RWMutex
}

// IsEmpty implements Collection.
func (a *ArrayCollection[E]) IsEmpty() bool {
	if a == nil {
		return true
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	return len(a.elems) == 0
}

// Size implements Collection.
func (a *ArrayCollection[E]) Size() uint {
	if a == nil {
		return 0
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	return uint(len(a.elems))
}

// Reset implements Collection.
func (a *ArrayCollection[E]) Reset() error {
	if a == nil {
		return ErrNilReceiver
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

// Slice implements Collection.
func (a *ArrayCollection[E]) Slice() []E {
	if a == nil {
		return nil
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	slice := make([]E, len(a.elems))
	copy(slice, a.elems)

	return slice
}

// Add implements Collection.
func (a *ArrayCollection[E]) Add(elem E) error {
	if a == nil {
		return ErrNilReceiver
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.elems = append(a.elems, elem)

	return nil
}
