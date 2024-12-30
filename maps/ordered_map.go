package maps

import (
	"cmp"
	"iter"
	"slices"

	common "github.com/PlayerR9/mygo-data/common"
)

// OrderedMap is a map that is ordered by the keys.
type OrderedMap[K cmp.Ordered, V any] struct {
	// table is the underlying map.
	table map[K]V

	// keys is the slice of keys.
	keys []K
}

// Set sets the value for the given key. If the key does not exist,
// it is inserted at the correct position in the slice of keys. If
// the key already exists, its value is updated.
//
// Parameters:
//   - k: The key to set.
//   - v: The value to set.
//
// Returns:
//   - error: An error if there is an error while setting the value.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (om *OrderedMap[K, V]) Set(k K, v V) error {
	if om == nil {
		return common.ErrNilReceiver
	}

	pos, ok := slices.BinarySearch(om.keys, k)
	if !ok {
		if om.table == nil {
			om.table = make(map[K]V)
		}

		om.keys = slices.Insert(om.keys, pos, k)
	}

	om.table[k] = v

	return nil
}

// Entry returns an iterator over the key-value pairs in the ordered map.
//
// Returns:
//   - iter.Seq2[K, V]: An iterator over the key-value pairs in the ordered
//     map. Never returns nil.
func (om OrderedMap[K, V]) Entry() iter.Seq2[K, V] {
	if len(om.keys) == 0 {
		return func(yield func(K, V) bool) {}
	}

	fn := func(yield func(K, V) bool) {
		for _, k := range om.keys {
			v := om.table[k]
			ok := yield(k, v)
			if !ok {
				return
			}
		}
	}

	return fn
}

// HasKey returns a boolean indicating whether the key exists in the ordered map.
//
// Parameters:
//   - k: The key to check for.
//
// Returns:
//   - bool: A boolean indicating whether the key exists in the ordered map.
func (om OrderedMap[K, V]) HasKey(k K) bool {
	if len(om.keys) == 0 {
		return false
	}

	_, ok := slices.BinarySearch(om.keys, k)
	return ok
}

// Get returns the value associated with the given key and a boolean indicating
// whether the key exists in the ordered map.
//
// Parameters:
//   - k: The key to retrieve the value for.
//
// Returns:
//   - V: The value associated with the given key. If the key does not exist,
//     returns a zero value.
//   - bool: A boolean indicating whether the key exists in the ordered map.
func (om OrderedMap[K, V]) Get(k K) (V, bool) {
	if len(om.keys) == 0 {
		return *new(V), false
	}

	v, ok := om.table[k]
	return v, ok
}

// Keys returns a slice of all keys in the ordered map.
//
// Returns:
//   - []K: A slice of all keys in the ordered map.
func (om OrderedMap[K, V]) Keys() []K {
	if len(om.keys) == 0 {
		return nil
	}

	keys := make([]K, len(om.keys))
	copy(keys, om.keys)

	return keys
}
