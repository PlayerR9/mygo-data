package queues

import "errors"

var (
	// ErrEmptyQueue occurs when the queue is empty. This error can be checked with
	// the == operator.
	//
	// Format:
	// 	"queue is empty"
	ErrEmptyQueue error

	// ErrFullQueue occurs when the queue is full. This error can be checked with
	// the == operator.
	//
	// Format:
	// 	"queue is full"
	ErrFullQueue error
)

func init() {
	ErrEmptyQueue = errors.New("queue is empty")
	ErrFullQueue = errors.New("queue is full")
}
