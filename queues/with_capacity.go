package queues

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// CapacityQueue is a queue that has a fixed capacity.
type CapacityQueue[T any] struct {
	// queue is the underlying queue.
	queue Queue[T]

	// size is the number of elements in the queue.
	size uint

	// capacity is the maximum number of elements in the queue.
	capacity uint

	// mu is the mutex for the queue.
	mu sync.RWMutex
}

// Enqueue implements Queue.
func (c *CapacityQueue[T]) Enqueue(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size >= c.capacity {
		return common.ErrFullCollection
	}

	err := c.queue.Enqueue(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Dequeue implements Queue.
func (c *CapacityQueue[T]) Dequeue() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return *new(T), common.ErrEmptyCollection
	}

	elem, err := c.queue.Dequeue()
	if err != nil {
		return *new(T), err
	}

	c.size--

	return elem, nil
}

// IsEmpty implements Queue.
func (c *CapacityQueue[T]) IsEmpty() bool {
	if c == nil {
		return true
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size == 0
}

// Front implements Queue.
func (c *CapacityQueue[T]) Front() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.size == 0 {
		return *new(T), common.ErrEmptyCollection
	}

	top, err := c.queue.Front()
	return top, err
}

// Slice implements Queue.
func (c *CapacityQueue[T]) Slice() []T {
	if c == nil {
		return nil
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	slice := c.queue.Slice()
	return slice
}

// Reset implements Queue.
func (c *CapacityQueue[T]) Reset() error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.queue.Reset()
	if err != nil {
		return err
	}

	c.size = 0

	return nil
}

// Size implements Queue.
func (c *CapacityQueue[T]) Size() uint {
	if c == nil {
		return 0
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size
}

// WithCapacity returns a new queue with the specified capacity.
//
// If queue is nil, a new ArrayQueue is created.
//
// Parameters:
//   - capacity: The capacity of the queue.
//   - queue: The queue to create a CapacityQueue from.
//
// Returns:
//   - *CapacityQueue[T]: The created CapacityQueue. Never returns nil.
func WithCapacity[T any](capacity uint, queue Queue[T]) *CapacityQueue[T] {
	if queue == nil {
		queue = new(ArrayQueue[T])
	}

	size := queue.Size()

	return &CapacityQueue[T]{
		queue:    queue,
		size:     size,
		capacity: capacity,
	}
}
