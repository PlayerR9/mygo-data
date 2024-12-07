package queues

import (
	"github.com/PlayerR9/mygo-data/collections"
	"github.com/PlayerR9/mygo-data/common"
)

type BasicQueue[E any] interface {
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
	//   - ErrFullQueue: If the queue is full.
	//   - any other error: Implementation-specific.
	Enqueue(elem E) error

	// Dequeue removes and returns the first value from the queue.
	//
	// Returns:
	//   - E: The value removed from the queue.
	//   - error: An error if the value could not be removed from the queue.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the queue is nil.
	//   - ErrEmptyQueue: If the queue is empty.
	//   - any other error: Implementation-specific.
	Dequeue() (E, error)

	// Front returns the first value in the queue without removing it.
	//
	// Returns:
	//   - T: The first value in the queue.
	//   - error: An error if the value could not be returned.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the queue is nil.
	//   - ErrEmptyQueue: If the queue is empty.
	//   - any other error: Implementation-specific.
	Front() (E, error)
}

// Queue is a generic interface for a queue.
type Queue[T any] interface {
	BasicQueue[T]

	collections.Collection[T]
}

// Enqueue enqueues multiple values to the given queue.
//
// Parameters:
//   - queue: The queue to enqueue to.
//   - elems: The values to enqueue.
//
// Returns:
//   - uint: The number of values that were enqueued.
//   - error: An error if the queue is nil or if a value could not be enqueued.
//
// Errors:
//   - common.ErrBadParam: If the queue is nil.
//   - ErrFullQueue: If the queue is full.
//   - any other error: Implementation-specific.
func Enqueue[T any](queue BasicQueue[T], elems ...T) (uint, error) {
	if queue == nil {
		return 0, common.NewErrNilParam("queue")
	}

	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	}

	for i, elem := range elems {
		err := queue.Enqueue(elem)
		if err != nil {
			return uint(i), err
		}
	}

	return lenElems, nil
}
