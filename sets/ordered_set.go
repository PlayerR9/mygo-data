package sets

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
	"strings"
	"sync"
)

// OrderedSet is a set of elements that are ordered.
//
// An empty set can be created by either using the `var set OrderedSet[E]` syntax or the `new(OrderedSet[E])` constructor.
type OrderedSet[E cmp.Ordered] struct {
	// elems is a slice of elements.
	elems []E

	// mu is a mutex to protect the ordered map.
	mu sync.RWMutex
}

// String implements Set.
func (s *OrderedSet[E]) String() string {
	if s == nil {
		return ""
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	elems := make([]string, 0, len(s.elems))

	for _, e := range s.elems {
		str := fmt.Sprint(e)
		elems = append(elems, str)
	}

	var builder strings.Builder

	_, _ = builder.WriteString("OrderedSet[E][")
	_, _ = builder.WriteString(strings.Join(elems, ", "))
	_, _ = builder.WriteRune(']')

	return builder.String()
}

// Has implements Set.
func (s *OrderedSet[E]) Has(e E) bool {
	if s == nil {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := slices.BinarySearch(s.elems, e)
	return ok
}

// Insert implements Set.
//
// No other error is returned.
func (s *OrderedSet[E]) Insert(e E) error {
	if s == nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	pos, ok := slices.BinarySearch(s.elems, e)
	if ok {
		return nil
	}

	s.elems = slices.Insert(s.elems, pos, e)

	return nil
}

// Reset resets the set, allowing it to be used again.
func (s *OrderedSet[E]) Reset() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	clear(s.elems)
	s.elems = nil
}

// Elem returns an iterator function that yields all the elements in the set.
//
// The iterator function takes a yield function as an argument and calls it for
// each element in the set. If the yield function returns false, the iteration
// stops early.
//
// Returns:
//   - iter.Seq[E]: An iterator function for the elements in the set.
func (s *OrderedSet[E]) Elem() iter.Seq[E] {
	if s == nil {
		return func(yield func(E) bool) {}
	}

	return func(yield func(E) bool) {
		s.mu.RLock()
		defer s.mu.RUnlock()

		for _, e := range s.elems {
			if !yield(e) {
				break
			}
		}
	}
}

// Size implements Set.
func (s *OrderedSet[E]) Size() uint {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return uint(len(s.elems))
}

// IsEmpty implements Set.
func (s *OrderedSet[E]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.elems) == 0
}
