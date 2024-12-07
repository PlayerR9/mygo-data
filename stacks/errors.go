package stacks

import "errors"

var (
	// ErrEmptyStack occurs when an operation is performed on an empty stack. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"stack is empty"
	ErrEmptyStack error

	// ErrFullStack occurs when an operation is performed on a full stack. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"stack is full"
	ErrFullStack error
)

func init() {
	ErrEmptyStack = errors.New("stack is empty")
	ErrFullStack = errors.New("stack is full")
}
