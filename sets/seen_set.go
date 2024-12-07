package sets

import (
	"fmt"
	"strings"
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

// String implements Set.
func (s *SeenSet[E]) String() string {
	if s == nil {
		return ""
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	elems := make([]string, 0, len(s.seen))

	for e := range s.seen {
		str := fmt.Sprint(e)
		elems = append(elems, str)
	}

	var builder strings.Builder

	_, _ = builder.WriteString("SeenSet[E][")
	_, _ = builder.WriteString(strings.Join(elems, ", "))
	_, _ = builder.WriteRune(']')

	return builder.String()
}

// Has implements Set.
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

// Insert implements Set.
//
// No other error is returned.
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

// Reset implements Set.
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

// Slice implements Set.
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

// IsEmpty implements Set.
func (s *SeenSet[E]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.seen) == 0
}

// Size implements Set.
func (s *SeenSet[E]) Size() uint {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return uint(len(s.seen))
}
