package stacks

import "errors"

var (
	// ErrEmptyStack occurs when the stack is empty. This can be checked using
	// the == operator.
	//
	// Format:
	// 	"stack is empty"
	ErrEmptyStack error

	// ErrFullStack occurs when the stack is full. This can be checked using
	// the == operator.
	//
	// Format:
	// 	"stack is full"
	ErrFullStack error
)

func init() {
	ErrEmptyStack = errors.New("stack is empty")
	ErrFullStack = errors.New("stack is full")
}
