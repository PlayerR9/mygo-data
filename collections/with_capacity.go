package collections

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// CapacityCollection is a collection that has a fixed capacity.
type CapacityCollection[T any] struct {
	// collection is the underlying collection.
	collection Collection[T]

	// size is the number of elements in the collection.
	size uint

	// capacity is the fixed capacity of the collection.
	capacity uint

	// mu is the mutex for the collection.
	mu sync.RWMutex
}

// IsEmpty implements Collection.
func (c *CapacityCollection[T]) IsEmpty() bool {
	if c == nil {
		return true
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size == 0
}

// Size implements Collection.
func (c *CapacityCollection[T]) Size() uint {
	if c == nil {
		return 0
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size
}

// Reset implements Collection.
func (c *CapacityCollection[T]) Reset() error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.collection.Reset()
	if err != nil {
		return err
	}

	c.size = 0

	return nil
}

// Slice implements Collection.
func (c *CapacityCollection[T]) Slice() []T {
	if c == nil {
		return nil
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	slice := c.collection.Slice()
	return slice
}

// Add implements Collection.
func (c *CapacityCollection[T]) Add(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size >= c.capacity {
		return ErrFullCollection
	}

	err := c.collection.Add(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// WithCapacity returns a new CapacityCollection with the specified capacity.
//
// If the provided collection is nil, a new ArrayCollection is created.
//
// Parameters:
//   - collection: The collection to create a CapacityCollection from.
//   - capacity: The capacity of the collection.
//
// Returns:
//   - *CapacityCollection[T]: The created CapacityCollection. Never returns nil.
func WithCapacity[T any](collection Collection[T], capacity uint) *CapacityCollection[T] {
	if collection == nil {
		collection = new(ArrayCollection[T])
	}

	return &CapacityCollection[T]{
		collection: collection,
		capacity:   capacity,
	}
}
