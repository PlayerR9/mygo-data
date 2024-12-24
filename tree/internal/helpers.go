package internal

// RejectNils filters out all nil elements from the given slice of pointers in-place.
//
// This function takes a slice of pointers and returns a new slice containing
// only the non-nil elements from the input slice. If the input slice contains
// only nil elements, it returns nil.
//
// Parameters:
//   - s: The slice of pointers to filter.
//
// Returns:
//   - uint: The number of non-nil elements in the filtered slice.
func RejectNils[S ~[]*E, E any](s *S) uint {
	var end uint

	for _, e := range *s {
		if e != nil {
			(*s)[end] = e
			end++
		}
	}

	n := uint(len(*s)) - end

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return n
}
