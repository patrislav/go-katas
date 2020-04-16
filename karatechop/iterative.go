package karatechop

// IterativeChopper is an implementation of the Chopper interface using a looping approach.
type IterativeChopper struct{}

// Iterative is a default IterativeChopper, instantiated and exported for ease of use.
var Iterative = IterativeChopper{}

// Chop returns the index of an item (needle) in a slice (haystack) or -1 if it was not found.
//
// This implementation maintains bounds of a window within the haystack that is being searched,
// moving either the lower of upper bound on each iteration until the value is found.
func (IterativeChopper) Chop(needle int, haystack []int) int {
	left := 0
	right := len(haystack)-1

	for right >= left {
		midIndex := (left + right) / 2
		midValue := haystack[midIndex]

		switch {
		case needle < midValue:
			right = midIndex-1
		case needle > midValue:
			left = midIndex+1
		default:
			return midIndex
		}
	}
	return -1
}
