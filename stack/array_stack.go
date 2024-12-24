package stack

import (
	"github.com/PlayerR9/mygo-data/errors"
)

// ArrayStack is a generic stack implemented using an array.
type ArrayStack[E any] struct {
	// elems is the underlying array.
	elems []E
}

// Push implements CoreStack.
func (as *ArrayStack[E]) Push(e E) error {
	if as == nil {
		return errors.ErrNilReceiver
	}

	as.elems = append(as.elems, e)

	return nil
}

// Pop implements CoreStack.
func (as *ArrayStack[E]) Pop() (E, error) {
	if as == nil {
		return *new(E), errors.ErrNilReceiver
	}

	if len(as.elems) == 0 {
		return *new(E), ErrEmptyStack
	}

	e := as.elems[len(as.elems)-1]
	as.elems = as.elems[:len(as.elems)-1]

	return e, nil
}

// IsEmpty implements CoreStack.
func (as ArrayStack[E]) IsEmpty() bool {
	ok := len(as.elems) == 0
	return ok
}

// Slice implements Collection.
func (as ArrayStack[E]) Slice() []E {
	if len(as.elems) == 0 {
		return nil
	}

	slice := make([]E, 0, len(as.elems))
	copy(slice, as.elems)

	j := len(slice) - 1

	for i := 0; i < j; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}

	return slice
}

// Reset implements Collection.
func (as *ArrayStack[E]) Reset() error {
	if as == nil {
		return errors.ErrNilReceiver
	}

	if len(as.elems) == 0 {
		return nil
	}

	clear(as.elems)
	as.elems = nil

	return nil
}
