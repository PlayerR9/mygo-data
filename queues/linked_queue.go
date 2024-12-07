package queues

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
	"github.com/PlayerR9/mygo-data/queues/internal"
)

// LinkedQueue is a queue implemented using a linked list. This is
// thread-safe.
//
// An empty LinkedQueue can be created by using the `queue := new(LinkedQueue[T])`
// constructor.
type LinkedQueue[T any] struct {
	// head is the first node in the queue.
	head *internal.Node[T]

	// head_mu is the mutex for the head field.
	head_mu sync.RWMutex

	// tail is the last node in the queue.
	tail *internal.Node[T]

	// tail_mu is the mutex for the tail field.
	tail_mu sync.RWMutex
}

// Enqueue implements BasicQueue.
func (l *LinkedQueue[T]) Enqueue(elem T) error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.tail_mu.Lock()
	defer l.tail_mu.Unlock()

	node := internal.NewNode(elem)

	if l.tail == nil {
		l.head_mu.Lock()
		l.head = node
		l.head_mu.Unlock()
	} else {
		_ = l.tail.SetNext(node)
	}

	l.tail = node

	return nil
}

// Dequeue implements BasicQueue.
func (l *LinkedQueue[T]) Dequeue() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	} else if l.head == nil {
		return *new(T), common.ErrEmptyCollection
	}

	l.head_mu.Lock()
	defer l.head_mu.Unlock()

	top := l.head
	l.head = l.head.GetNext()

	if l.head == nil {
		l.tail_mu.Lock()
		l.tail = nil
		l.tail_mu.Unlock()
	}

	_ = top.SetNext(nil) // garbage collection

	v, _ := top.GetValue()

	return v, nil
}

// Front implements BasicQueue.
func (l *LinkedQueue[T]) Front() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.head_mu.RLock()
	defer l.head_mu.RUnlock()

	if l.head == nil {
		return *new(T), common.ErrEmptyCollection
	}

	return l.head.GetValue()
}

// IsEmpty implements common.Collection.
func (l *LinkedQueue[T]) IsEmpty() bool {
	if l == nil {
		return true
	}

	l.head_mu.RLock()
	defer l.head_mu.RUnlock()

	return l.head == nil
}

// Slice implements common.Collection.
func (l *LinkedQueue[T]) Slice() []T {
	if l == nil {
		return nil
	}

	l.head_mu.RLock()
	defer l.head_mu.RUnlock()

	if l.head == nil {
		return nil
	}

	l.tail_mu.RLock()
	defer l.tail_mu.RUnlock()

	var elems []T

	for n := l.head; n != nil; n = n.GetNext() {
		v, _ := n.GetValue()

		elems = append(elems, v)
	}

	return elems
}

// Size implements common.Collection.
func (l *LinkedQueue[T]) Size() uint {
	if l == nil {
		return 0
	}

	l.head_mu.RLock()
	defer l.head_mu.RUnlock()

	if l.head == nil {
		return 0
	}

	l.tail_mu.RLock()
	defer l.tail_mu.RUnlock()

	var size uint

	for n := l.head; n != nil; n = n.GetNext() {
		size++
	}

	return size
}

// Reset implements common.Collection.
func (l *LinkedQueue[T]) Reset() error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.head_mu.Lock()
	defer l.head_mu.Unlock()

	if l.head == nil {
		return nil
	}

	l.tail_mu.Lock()
	defer l.tail_mu.Unlock()

	l.head.Release()
	l.head = nil

	l.tail = nil

	return nil
}

// Add implements common.Collection.
func (l *LinkedQueue[T]) Add(elem T) error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.tail_mu.Lock()
	defer l.tail_mu.Unlock()

	node := internal.NewNode(elem)

	if l.tail == nil {
		l.head_mu.Lock()
		l.head = node
		l.head_mu.Unlock()
	} else {
		_ = l.tail.SetNext(node)
	}

	l.tail = node

	return nil
}
