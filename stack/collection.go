package stack

// Collection is an interface that represents a collection of elements.
type Collection[E any] interface {
	// Slice returns a slice of the elements in the collection.
	//
	// Returns:
	//   - []E: A slice of the elements in the collection.
	Slice() []E
}
