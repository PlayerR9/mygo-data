package stacks

import "errors"

var (
	// ErrStackEmpty occurs when the stack is empty. This can be checked using
	// the == operator.
	//
	// Format:
	// 	"stack is empty"
	ErrStackEmpty error
)

func init() {
	ErrStackEmpty = errors.New("stack is empty")
}
