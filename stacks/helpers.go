package stacks

// Reverse returns a new slice with the elements of the given slice in reverse
// order. If the given slice is empty, the function returns nil.
//
// Parameters:
//   - s: The slice to reverse.
//
// Returns:
//   - S: The reversed slice.
func Reverse[S ~[]E, E any](s S) S {
	if len(s) == 0 {
		return nil
	}

	reverse := make(S, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, s[i])
	}

	return reverse
}
