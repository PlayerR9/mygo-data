package stacks

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// CapacityStack is a stack that has a fixed capacity.
type CapacityStack[T any] struct {
	// stack is the underlying stack.
	stack Stack[T]

	// size is the number of elements in the stack.
	size uint

	// capacity is the maximum number of elements in the stack.
	capacity uint

	// mu is the mutex for the stack.
	mu sync.RWMutex
}

// Push implements BasicStack.
func (c *CapacityStack[T]) Push(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size >= c.capacity {
		return common.ErrFullCollection
	}

	err := c.stack.Push(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Pop implements BasicStack.
func (c *CapacityStack[T]) Pop() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return *new(T), common.ErrEmptyCollection
	}

	elem, err := c.stack.Pop()
	if err != nil {
		return *new(T), err
	}

	c.size--

	return elem, nil
}

// Peek implements BasicStack.
func (c *CapacityStack[T]) Peek() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.size == 0 {
		return *new(T), common.ErrEmptyCollection
	}

	top, err := c.stack.Peek()
	return top, err
}

// IsEmpty implements common.Collection.
func (c *CapacityStack[T]) IsEmpty() bool {
	if c == nil {
		return true
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size == 0
}

// Slice implements common.Collection.
func (c *CapacityStack[T]) Slice() []T {
	if c == nil {
		return nil
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	slice := c.stack.Slice()
	return slice
}

// Reset implements common.Collection.
func (c *CapacityStack[T]) Reset() error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.stack.Reset()
	if err != nil {
		return err
	}

	c.size = 0

	return nil
}

// Size implements common.Collection.
func (c *CapacityStack[T]) Size() uint {
	if c == nil {
		return 0
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size
}

// WithCapacity returns a new stack with the specified capacity.
//
// If stack is nil, a new ArrayStack is created.
//
// Parameters:
//   - capacity: The capacity of the stack.
//   - stack: The stack to create a CapacityStack from.
//
// Returns:
//   - *CapacityStack[T]: The created CapacityStack. Never returns nil.
func WithCapacity[T any](capacity uint, stack Stack[T]) *CapacityStack[T] {
	if stack == nil {
		stack = new(ArrayStack[T])
	}

	size := stack.Size()

	return &CapacityStack[T]{
		stack:    stack,
		size:     size,
		capacity: capacity,
	}
}
