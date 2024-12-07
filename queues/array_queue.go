package queues

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// ArrayQueue is a generic implementation of a queue using an array. This is
// thread-safe.
//
// An empty array queue can be created by using the `queue := new(ArrayQueue[T])`
// constructor.
type ArrayQueue[T any] struct {
	// queue is the array used to store the values.
	queue []T

	// mu is the mutex used to synchronize access to the queue.
	mu sync.RWMutex
}

// Enqueue implements Queue.
func (a *ArrayQueue[T]) Enqueue(value T) error {
	if a == nil {
		return common.ErrNilReceiver
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.queue = append(a.queue, value)

	return nil
}

// Dequeue implements Queue.
func (a *ArrayQueue[T]) Dequeue() (T, error) {
	if a == nil {
		return *new(T), common.ErrNilReceiver
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.queue) == 0 {
		return *new(T), ErrEmptyQueue
	}

	value := a.queue[0]
	a.queue = a.queue[1:]

	return value, nil
}

// IsEmpty implements Queue.
func (a *ArrayQueue[T]) IsEmpty() bool {
	if a == nil {
		return true
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	return len(a.queue) == 0
}

// Front implements Queue.
func (a *ArrayQueue[T]) Front() (T, error) {
	if a == nil {
		return *new(T), common.ErrNilReceiver
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	if len(a.queue) == 0 {
		return *new(T), ErrEmptyQueue
	}

	return a.queue[0], nil
}

// Slice implements Queue.
func (a *ArrayQueue[T]) Slice() []T {
	if a == nil {
		return nil
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	elems := make([]T, len(a.queue))
	copy(elems, a.queue)

	return elems
}

// Size implements Queue.
func (a *ArrayQueue[T]) Size() uint {
	if a == nil {
		return 0
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	return uint(len(a.queue))
}

// Reset implements Queue.
func (a *ArrayQueue[T]) Reset() error {
	if a == nil {
		return common.ErrNilReceiver
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.queue) == 0 {
		return nil
	}

	clear(a.queue)
	a.queue = nil

	return nil
}
