package queues

import "errors"

var (
	// ErrEmptyQueue occurs when an operation is performed on an empty queue. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"queue is empty"
	ErrEmptyQueue error

	// ErrFullQueue occurs when an operation is performed on a full queue. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"queue is full"
	ErrFullQueue error
)

func init() {
	ErrEmptyQueue = errors.New("queue is empty")
	ErrFullQueue = errors.New("queue is full")
}
