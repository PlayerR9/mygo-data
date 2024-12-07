package internal

import "github.com/PlayerR9/mygo-data/common"

// Node is a node in a queue.
type Node[T any] struct {
	// value is the value of the node.
	value T

	// next is the next node in the queue.
	next *Node[T]
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
		next:  nil,
	}
}

// GetValue returns the value of the node.
//
// Returns:
//   - T: The value of the node.
//   - error: nil if the value was returned, or common.ErrNilReceiver if the receiver is nil.
func (n *Node[T]) GetValue() (T, error) {
	if n == nil {
		return *new(T), common.ErrNilReceiver
	}

	return n.value, nil
}

// GetNext returns the next node in the queue.
//
// Returns:
//   - *Node[T]: The next node in the queue, or nil if the receiver is nil.
func (n *Node[T]) GetNext() *Node[T] {
	if n == nil {
		return nil
	}

	return n.next
}

// SetNext sets the next node in the queue.
//
// Parameters:
//   - next: The next node in the queue.
//
// Returns:
//   - error: An error if the next node could not be set.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *Node[T]) SetNext(next *Node[T]) error {
	if n == nil {
		return common.ErrNilReceiver
	}

	n.next = next

	return nil
}
