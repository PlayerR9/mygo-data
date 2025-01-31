package stack

import "errors"

var (
	// ErrEmptyStack occurs when the stack is empty. This error can be checked
	// with the == operator.
	//
	// Format:
	// 	"stack is empty"
	ErrEmptyStack error
)

func init() {
	ErrEmptyStack = errors.New("stack is empty")
}
