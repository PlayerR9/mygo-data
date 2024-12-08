package sets

import (
	"sync"

	"github.com/PlayerR9/mygo-data/common"
)

// SeenSet is a set of seen elements.
//
// An empty set can be created by either using the `var set SeenSet[T]` syntax or the `new(SeenSet[T])` constructor.
type SeenSet[E comparable] struct {
	// seen is a map of seen elements.
	seen map[E]struct{}

	// mu is a mutex to protect the seen map.
	mu sync.RWMutex
}

// Has implements BasicSet.
func (s *SeenSet[E]) Has(e E) bool {
	if s == nil {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.seen) == 0 {
		return false
	}

	_, ok := s.seen[e]
	return ok
}

// Insert implements BasicSet.
func (s *SeenSet[E]) Insert(e E) error {
	if s == nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.seen == nil {
		s.seen = make(map[E]struct{})
	}

	s.seen[e] = struct{}{}

	return nil
}

// Reset implements common.Collection.
func (s *SeenSet[E]) Reset() error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.seen) == 0 {
		return nil
	}

	clear(s.seen)
	s.seen = nil

	return nil
}

// Slice implements common.Collection.
func (s *SeenSet[E]) Slice() []E {
	if s == nil {
		return nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	slice := make([]E, 0, len(s.seen))

	for e := range s.seen {
		slice = append(slice, e)
	}

	return slice
}

// IsEmpty implements common.Collection.
func (s *SeenSet[E]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.seen) == 0
}

// Size implements common.Collection.
func (s *SeenSet[E]) Size() uint {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return uint(len(s.seen))
}

// Add implements common.Collection.
func (s *SeenSet[E]) Add(elem E) error {
	if s == nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.seen == nil {
		s.seen = make(map[E]struct{})
	}

	s.seen[elem] = struct{}{}

	return nil
}
