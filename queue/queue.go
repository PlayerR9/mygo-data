package queue

import "github.com/PlayerR9/mygo-data/common"

// Queue is a generic interface for a queue.
type Queue[T any] interface {
	// Enqueue adds a value to the queue.
	//
	// Parameters:
	//   - elem: The value to add to the queue.
	//
	// Returns:
	//   - error: An error if the value could not be added to the queue.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the queue is nil.
	Enqueue(elem T) error

	// Dequeue removes and returns the first value from the queue.
	//
	// Returns:
	//   - T: The value removed from the queue.
	//   - error: An error if the value could not be removed from the queue.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the queue is nil.
	//   - ErrEmptyQueue: If the queue is empty.
	Dequeue() (T, error)

	// IsEmpty checks if the queue is empty.
	//
	// Returns:
	//   - bool: True if the queue is empty, false otherwise.
	IsEmpty() bool

	// Front returns the first value in the queue without removing it.
	//
	// Returns:
	//   - T: The first value in the queue.
	//   - error: An error if the value could not be returned.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the queue is nil.
	//   - ErrEmptyQueue: If the queue is empty.
	Front() (T, error)

	// Slice returns a slice of the values in the queue.
	//
	// Returns:
	//   - []T: A slice of the values in the queue.
	//
	// The returned slice has the property that, the first element of the
	// slice is the first value in the queue.
	Slice() []T

	// Size returns the number of values in the queue.
	//
	// Returns:
	//   - uint: The number of values in the queue.
	Size() uint
}

// Enqueue enqueues multiple values to the given queue.
//
// Parameters:
//   - queue: The queue to enqueue to.
//   - elems: The values to enqueue.
//
// Returns:
//   - error: An error if the queue is nil or if a value could not be enqueued.
//
// Errors:
//   - common.ErrBadParam: If the queue is nil.
func Enqueue[T any](queue Queue[T], elems ...T) error {
	if queue == nil {
		return common.NewErrNilParam("queue")
	} else if len(elems) == 0 {
		return nil
	}

	for _, elem := range elems {
		_ = queue.Enqueue(elem)
	}

	return nil
}
