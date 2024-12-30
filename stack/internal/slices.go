package internal

// Reverse reverses the given slice in-place.
//
// Parameters:
//   - slice: The slice to be reversed.
func Reverse[E any](slice []E) {
	j := len(slice) - 1

	for i := 0; i < j; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}
}
