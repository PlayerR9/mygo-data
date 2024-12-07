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

// Enqueue implements BasicQueue.
func (c *CapacityQueue[T]) Enqueue(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size >= c.capacity {
		return ErrFullQueue
	}

	err := c.queue.Enqueue(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Dequeue implements BasicQueue.
func (c *CapacityQueue[T]) Dequeue() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return *new(T), ErrEmptyQueue
	}

	elem, err := c.queue.Dequeue()
	if err != nil {
		return *new(T), err
	}

	c.size--

	return elem, nil
}

// Front implements BasicQueue.
func (c *CapacityQueue[T]) Front() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.size == 0 {
		return *new(T), ErrEmptyQueue
	}

	top, err := c.queue.Front()
	return top, err
}

// IsEmpty implements common.Collection.
func (c *CapacityQueue[T]) IsEmpty() bool {
	if c == nil {
		return true
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size == 0
}

// Slice implements common.Collection.
func (c *CapacityQueue[T]) Slice() []T {
	if c == nil {
		return nil
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	slice := c.queue.Slice()
	return slice
}

// Reset implements common.Collection.
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

// Size implements common.Collection.
func (c *CapacityQueue[T]) Size() uint {
	if c == nil {
		return 0
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size
}

// Add implements common.Collection.
func (c *CapacityQueue[T]) Add(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size >= c.capacity {
		return ErrFullQueue
	}

	err := c.queue.Add(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
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
