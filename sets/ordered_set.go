package sets

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/PlayerR9/mygo-data/common"
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

// Has implements BasicSet.
func (s *OrderedSet[E]) Has(e E) bool {
	if s == nil {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := slices.BinarySearch(s.elems, e)
	return ok
}

// Insert implements BasicSet.
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

// Reset implements common.Collection.
func (s *OrderedSet[E]) Reset() error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.elems) == 0 {
		return nil
	}

	clear(s.elems)
	s.elems = nil

	return nil
}

// Slice implements common.Collection.
func (s *OrderedSet[E]) Slice() []E {
	if s == nil {
		return nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	slice := make([]E, len(s.elems))
	copy(slice, s.elems)

	return slice
}

// Size implements common.Collection.
func (s *OrderedSet[E]) Size() uint {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return uint(len(s.elems))
}

// IsEmpty implements common.Collection.
func (s *OrderedSet[E]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.elems) == 0
}

// Add implements common.Collection.
func (s *OrderedSet[E]) Add(elem E) error {
	if s == nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	pos, ok := slices.BinarySearch(s.elems, elem)
	if ok {
		return nil
	}

	s.elems = slices.Insert(s.elems, pos, elem)

	return nil
}

// String implements fmt.Stringer.
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
