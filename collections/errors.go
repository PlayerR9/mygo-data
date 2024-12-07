package collections

import "errors"

var (
	// ErrEmptyCollection occurs when a collection is empty. This error can be
	// checked with the == operator.
	//
	// Format:
	// 	"collection is empty"
	ErrEmptyCollection error

	// ErrFullCollection occurs when a collection is full. This error can be
	// checked with the == operator.
	//
	// Format:
	// 	"collection is full"
	ErrFullCollection error
)

func init() {
	ErrEmptyCollection = errors.New("collection is empty")
	ErrFullCollection = errors.New("collection is full")
}
