package stacks

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
	"github.com/PlayerR9/mygo-data/stacks/internal"
)

// LinkedStack is a stack implemented using a linked list.
//
// An empty LinkedStack can be created by either using the `new(LinkedStack[T])`
// constructor.
type LinkedStack[T any] struct {
	// head is the top of the stack
	head *internal.Node[T]

	// mu is the mutex for the stack
	mu sync.RWMutex
}

// Push implements Stack.
func (l *LinkedStack[T]) Push(elem T) error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	node := internal.NewNode(elem)
	_ = node.SetPrev(l.head)

	l.head = node

	return nil
}

// Pop implements Stack.
func (l *LinkedStack[T]) Pop() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return *new(T), ErrEmptyStack
	}

	top := l.head
	l.head = l.head.GetPrev()

	_ = top.SetPrev(nil) // garbage collection

	v, _ := top.GetValue()

	return v, nil
}

// IsEmpty implements Stack.
func (l *LinkedStack[T]) IsEmpty() bool {
	if l == nil {
		return true
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.head == nil
}

// Peek implements Stack.
func (l *LinkedStack[T]) Peek() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.head == nil {
		return *new(T), ErrEmptyStack
	}

	v, _ := l.head.GetValue()

	return v, nil
}

// Slice implements Stack.
func (l *LinkedStack[T]) Slice() []T {
	if l == nil {
		return nil
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	var elems []T

	for n := l.head; n != nil; n = n.GetPrev() {
		v, _ := n.GetValue()

		elems = append(elems, v)
	}

	return elems
}

// Reset implements Stack.
func (l *LinkedStack[T]) Reset() error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return nil
	}

	l.head.Release()
	l.head = nil

	return nil
}

// Size implements Stack.
func (l *LinkedStack[T]) Size() uint {
	if l == nil {
		return 0
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.head == nil {
		return 0
	}

	var size uint

	for n := l.head; n != nil; n = n.GetPrev() {
		size++
	}

	return size
}
