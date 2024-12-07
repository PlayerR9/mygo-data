package internal

import "github.com/PlayerR9/mygo-data/common"

// Node is a node in a stack.
type Node[T any] struct {
	// value is the value of the node.
	value T

	// prev is the previous node in the stack.
	prev *Node[T]
}

// NewNode creates a new node with the given value and returns it.
//
// Parameters:
//   - value: The value of the node.
//
// Returns:
//   - *Node[T]: The new node. Never returns nil.
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
		prev:  nil,
	}
}

// SetPrev sets the previous node in the stack.
//
// Parameters:
//   - prev: The previous node in the stack.
//
// Returns:
//   - error: An error if the previous node could not be set.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *Node[T]) SetPrev(prev *Node[T]) error {
	if n == nil {
		return common.ErrNilReceiver
	}

	n.prev = prev

	return nil
}

// GetPrev returns the previous node in the stack.
//
// Returns:
//   - *Node[T]: The previous node in the stack, or nil if the receiver is nil.nil if the value was returned, or common.ErrNilReceiver if the receiver is nil.
func (n *Node[T]) GetPrev() *Node[T] {
	if n == nil {
		return nil
	}

	return n.prev
}

// GetValue returns the value of the node.
//
// Returns:
//   - T: The value of the node.
//   - error: An error if the value could not be returned.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *Node[T]) GetValue() (T, error) {
	if n == nil {
		return *new(T), common.ErrNilReceiver
	}

	return n.value, nil
}

// Release releases the resources held by the node.
func (n *Node[T]) Release() {
	if n == nil {
		return
	}

	n.value = *new(T)

	n.prev.Release()
	n.prev = nil
}
